package task

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/stats"
	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/jamethy/project-rising-heat/internal/util/ptr"
	"github.com/jamethy/project-rising-heat/internal/weather"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func DailyData(ctx context.Context, d *sql.DB, w weather.Client) error {

	if err := CreateToday(ctx, d, w); err != nil {
		return err
	}

	if err := CalculateSummaries(ctx, d); err != nil {
		return err
	}

	return nil
}

func CreateToday(ctx context.Context, d *sql.DB, w weather.Client) error {
	today := time.Now().UTC().Truncate(24 * time.Hour)

	existingCount, err := db.DailyData(db.DailyDatumWhere.Date.EQ(today)).Count(ctx, d)
	if err != nil {
		return fmt.Errorf("failed to check existance: %w", err)
	}
	if existingCount > 0 {
		log.Println("Record already exists for today")
	} else {
		drec, err := w.CreateDailyDBRecord(ctx)
		if err != nil {
			return err
		}

		err = drec.Insert(ctx, d, boil.Infer())
		if err != nil {
			return fmt.Errorf("failed to write to database: %w", err)
		}
		util.PrettyPrint(drec)
	}

	return nil
}

func CalculateSummaries(ctx context.Context, d *sql.DB) error {
	today := time.Now().UTC().Truncate(24 * time.Hour)

	drecs, err := db.DailyData(
		db.DailyDatumWhere.SummaryDate.IsNull(),
		db.DailyDatumWhere.Date.NEQ(today),
	).All(ctx, d)
	if err != nil {
		return fmt.Errorf("failed to get daily records: %w", err)
	}

	log.Printf("summarizing %d days\n", len(drecs))

	for _, drec := range drecs {
		err = updateDailyData(ctx, d, drec)
		if err != nil {
			return err
		}
	}

	log.Println(drecs)
	return nil
}

// average the temperature over five minutes around bed time
func getBedTimeTemp(ctx context.Context, d *sql.DB, bedTime time.Time) (*float32, error) {
	data, err := db.Upstairs(
		qm.Select(db.UpstairColumns.Temperature, db.UpstairColumns.Timestamp),
		db.UpstairWhere.Timestamp.GT(bedTime.UTC().Add(-150*time.Second)),
		db.UpstairWhere.Timestamp.LT(bedTime.UTC().Add(150*time.Second)),
		db.UpstairWhere.Temperature.IsNotNull(),
		qm.OrderBy(db.UpstairColumns.Timestamp),
	).All(ctx, d)
	if err != nil {
		return nil, err
	}
	temperatures := make(stats.Data, len(data))
	for i, t := range data {
		temperatures[i] = stats.DatumFrom(t.Timestamp, t.Temperature)
	}

	return ptr.Float32(temperatures.Avg()), nil
}

func updateDailyData(ctx context.Context, d *sql.DB, rec *db.DailyDatum) error {
	year, month, date := rec.Date.Date()
	timezone, _ := time.LoadLocation("America/New_York")
	bedTime := time.Date(year, month, date, 22, 0, 0, 0, timezone)

	bedTimeTemp, err := getBedTimeTemp(ctx, d, bedTime)
	if err != nil {
		return err
	}

	weatherRecs, err := db.Weathers(
		db.WeatherWhere.Timestamp.GT(rec.Sunrise.Time),
		db.WeatherWhere.Timestamp.LT(bedTime),
		qm.OrderBy(db.WeatherColumns.Timestamp),
	).All(ctx, d)

	temperatures := make(stats.Data, len(weatherRecs))
	feelsLikes := make(stats.Data, len(weatherRecs))
	uvs := make(stats.Data, len(weatherRecs))
	rains := make(stats.Data, len(weatherRecs))
	clouds := make(stats.Data, len(weatherRecs))

	for i, w := range weatherRecs {
		temperatures[i] = stats.DatumFrom(w.Timestamp, w.Temperature)
		feelsLikes[i] = stats.DatumFrom(w.Timestamp, w.FeelsLike)
		uvs[i] = stats.DatumFrom(w.Timestamp, w.UvIndex)
		rains[i] = stats.DatumFrom(w.Timestamp, w.RainLevel)
		clouds[i] = stats.DatumFrom(w.Timestamp, w.Clouds)
	}

	rec.SummaryDate = null.TimeFrom(time.Now())
	rec.BedTimeTemp = null.Float32FromPtr(bedTimeTemp)
	rec.FanOn = null.StringFrom("OFF")
	rec.TemperatureMax = null.Float32From(temperatures.Max())
	rec.TemperatureAvg = null.Float32From(temperatures.Avg())
	rec.TemperatureSum = null.Float32From(temperatures.AreaSum())
	rec.FeelsLikeMax = null.Float32From(feelsLikes.Max())
	rec.FeelsLikeAvg = null.Float32From(feelsLikes.Avg())
	rec.FeelsLikeSum = null.Float32From(feelsLikes.AreaSum())
	rec.UvMax = null.Float32From(uvs.Max())
	rec.UvAvg = null.Float32From(uvs.Avg())
	rec.UvSum = null.Float32From(uvs.AreaSum())
	rec.RainMax = null.Float32From(rains.Max())
	rec.RainAvg = null.Float32From(rains.Avg())
	rec.RainSum = null.Float32From(rains.AreaSum())
	rec.CloudMax = null.Float32From(clouds.Max())
	rec.CloudAvg = null.Float32From(clouds.Avg())
	rec.CloudSum = null.Float32From(clouds.AreaSum())

	_, err = rec.Update(ctx, d, boil.Infer())
	return err
}
