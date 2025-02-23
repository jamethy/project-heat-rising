// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// DailyDatum is an object representing the database table.
type DailyDatum struct {
	ID             int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt      time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Date           time.Time `boil:"date" json:"date" toml:"date" yaml:"date"`
	Sunrise        time.Time `boil:"sunrise" json:"sunrise" toml:"sunrise" yaml:"sunrise"`
	Sunset         time.Time `boil:"sunset" json:"sunset" toml:"sunset" yaml:"sunset"`
	SummaryDate    time.Time `boil:"summary_date" json:"summary_date" toml:"summary_date" yaml:"summary_date"`
	BedTimeTemp    float64   `boil:"bed_time_temp" json:"bed_time_temp" toml:"bed_time_temp" yaml:"bed_time_temp"`
	FanOn          string    `boil:"fan_on" json:"fan_on" toml:"fan_on" yaml:"fan_on"`
	TemperatureMax float64   `boil:"temperature_max" json:"temperature_max" toml:"temperature_max" yaml:"temperature_max"`
	TemperatureAvg float64   `boil:"temperature_avg" json:"temperature_avg" toml:"temperature_avg" yaml:"temperature_avg"`
	TemperatureSum float64   `boil:"temperature_sum" json:"temperature_sum" toml:"temperature_sum" yaml:"temperature_sum"`
	FeelsLikeMax   float64   `boil:"feels_like_max" json:"feels_like_max" toml:"feels_like_max" yaml:"feels_like_max"`
	FeelsLikeAvg   float64   `boil:"feels_like_avg" json:"feels_like_avg" toml:"feels_like_avg" yaml:"feels_like_avg"`
	FeelsLikeSum   float64   `boil:"feels_like_sum" json:"feels_like_sum" toml:"feels_like_sum" yaml:"feels_like_sum"`
	UvMax          float64   `boil:"uv_max" json:"uv_max" toml:"uv_max" yaml:"uv_max"`
	UvAvg          float64   `boil:"uv_avg" json:"uv_avg" toml:"uv_avg" yaml:"uv_avg"`
	UvSum          float64   `boil:"uv_sum" json:"uv_sum" toml:"uv_sum" yaml:"uv_sum"`
	RainMax        float64   `boil:"rain_max" json:"rain_max" toml:"rain_max" yaml:"rain_max"`
	RainAvg        float64   `boil:"rain_avg" json:"rain_avg" toml:"rain_avg" yaml:"rain_avg"`
	RainSum        float64   `boil:"rain_sum" json:"rain_sum" toml:"rain_sum" yaml:"rain_sum"`
	CloudMax       float64   `boil:"cloud_max" json:"cloud_max" toml:"cloud_max" yaml:"cloud_max"`
	CloudAvg       float64   `boil:"cloud_avg" json:"cloud_avg" toml:"cloud_avg" yaml:"cloud_avg"`
	CloudSum       float64   `boil:"cloud_sum" json:"cloud_sum" toml:"cloud_sum" yaml:"cloud_sum"`

	R *dailyDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dailyDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DailyDatumColumns = struct {
	ID             string
	CreatedAt      string
	Date           string
	Sunrise        string
	Sunset         string
	SummaryDate    string
	BedTimeTemp    string
	FanOn          string
	TemperatureMax string
	TemperatureAvg string
	TemperatureSum string
	FeelsLikeMax   string
	FeelsLikeAvg   string
	FeelsLikeSum   string
	UvMax          string
	UvAvg          string
	UvSum          string
	RainMax        string
	RainAvg        string
	RainSum        string
	CloudMax       string
	CloudAvg       string
	CloudSum       string
}{
	ID:             "id",
	CreatedAt:      "created_at",
	Date:           "date",
	Sunrise:        "sunrise",
	Sunset:         "sunset",
	SummaryDate:    "summary_date",
	BedTimeTemp:    "bed_time_temp",
	FanOn:          "fan_on",
	TemperatureMax: "temperature_max",
	TemperatureAvg: "temperature_avg",
	TemperatureSum: "temperature_sum",
	FeelsLikeMax:   "feels_like_max",
	FeelsLikeAvg:   "feels_like_avg",
	FeelsLikeSum:   "feels_like_sum",
	UvMax:          "uv_max",
	UvAvg:          "uv_avg",
	UvSum:          "uv_sum",
	RainMax:        "rain_max",
	RainAvg:        "rain_avg",
	RainSum:        "rain_sum",
	CloudMax:       "cloud_max",
	CloudAvg:       "cloud_avg",
	CloudSum:       "cloud_sum",
}

var DailyDatumTableColumns = struct {
	ID             string
	CreatedAt      string
	Date           string
	Sunrise        string
	Sunset         string
	SummaryDate    string
	BedTimeTemp    string
	FanOn          string
	TemperatureMax string
	TemperatureAvg string
	TemperatureSum string
	FeelsLikeMax   string
	FeelsLikeAvg   string
	FeelsLikeSum   string
	UvMax          string
	UvAvg          string
	UvSum          string
	RainMax        string
	RainAvg        string
	RainSum        string
	CloudMax       string
	CloudAvg       string
	CloudSum       string
}{
	ID:             "daily_data.id",
	CreatedAt:      "daily_data.created_at",
	Date:           "daily_data.date",
	Sunrise:        "daily_data.sunrise",
	Sunset:         "daily_data.sunset",
	SummaryDate:    "daily_data.summary_date",
	BedTimeTemp:    "daily_data.bed_time_temp",
	FanOn:          "daily_data.fan_on",
	TemperatureMax: "daily_data.temperature_max",
	TemperatureAvg: "daily_data.temperature_avg",
	TemperatureSum: "daily_data.temperature_sum",
	FeelsLikeMax:   "daily_data.feels_like_max",
	FeelsLikeAvg:   "daily_data.feels_like_avg",
	FeelsLikeSum:   "daily_data.feels_like_sum",
	UvMax:          "daily_data.uv_max",
	UvAvg:          "daily_data.uv_avg",
	UvSum:          "daily_data.uv_sum",
	RainMax:        "daily_data.rain_max",
	RainAvg:        "daily_data.rain_avg",
	RainSum:        "daily_data.rain_sum",
	CloudMax:       "daily_data.cloud_max",
	CloudAvg:       "daily_data.cloud_avg",
	CloudSum:       "daily_data.cloud_sum",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var DailyDatumWhere = struct {
	ID             whereHelperint
	CreatedAt      whereHelpertime_Time
	Date           whereHelpertime_Time
	Sunrise        whereHelpertime_Time
	Sunset         whereHelpertime_Time
	SummaryDate    whereHelpertime_Time
	BedTimeTemp    whereHelperfloat64
	FanOn          whereHelperstring
	TemperatureMax whereHelperfloat64
	TemperatureAvg whereHelperfloat64
	TemperatureSum whereHelperfloat64
	FeelsLikeMax   whereHelperfloat64
	FeelsLikeAvg   whereHelperfloat64
	FeelsLikeSum   whereHelperfloat64
	UvMax          whereHelperfloat64
	UvAvg          whereHelperfloat64
	UvSum          whereHelperfloat64
	RainMax        whereHelperfloat64
	RainAvg        whereHelperfloat64
	RainSum        whereHelperfloat64
	CloudMax       whereHelperfloat64
	CloudAvg       whereHelperfloat64
	CloudSum       whereHelperfloat64
}{
	ID:             whereHelperint{field: "\"prh\".\"daily_data\".\"id\""},
	CreatedAt:      whereHelpertime_Time{field: "\"prh\".\"daily_data\".\"created_at\""},
	Date:           whereHelpertime_Time{field: "\"prh\".\"daily_data\".\"date\""},
	Sunrise:        whereHelpertime_Time{field: "\"prh\".\"daily_data\".\"sunrise\""},
	Sunset:         whereHelpertime_Time{field: "\"prh\".\"daily_data\".\"sunset\""},
	SummaryDate:    whereHelpertime_Time{field: "\"prh\".\"daily_data\".\"summary_date\""},
	BedTimeTemp:    whereHelperfloat64{field: "\"prh\".\"daily_data\".\"bed_time_temp\""},
	FanOn:          whereHelperstring{field: "\"prh\".\"daily_data\".\"fan_on\""},
	TemperatureMax: whereHelperfloat64{field: "\"prh\".\"daily_data\".\"temperature_max\""},
	TemperatureAvg: whereHelperfloat64{field: "\"prh\".\"daily_data\".\"temperature_avg\""},
	TemperatureSum: whereHelperfloat64{field: "\"prh\".\"daily_data\".\"temperature_sum\""},
	FeelsLikeMax:   whereHelperfloat64{field: "\"prh\".\"daily_data\".\"feels_like_max\""},
	FeelsLikeAvg:   whereHelperfloat64{field: "\"prh\".\"daily_data\".\"feels_like_avg\""},
	FeelsLikeSum:   whereHelperfloat64{field: "\"prh\".\"daily_data\".\"feels_like_sum\""},
	UvMax:          whereHelperfloat64{field: "\"prh\".\"daily_data\".\"uv_max\""},
	UvAvg:          whereHelperfloat64{field: "\"prh\".\"daily_data\".\"uv_avg\""},
	UvSum:          whereHelperfloat64{field: "\"prh\".\"daily_data\".\"uv_sum\""},
	RainMax:        whereHelperfloat64{field: "\"prh\".\"daily_data\".\"rain_max\""},
	RainAvg:        whereHelperfloat64{field: "\"prh\".\"daily_data\".\"rain_avg\""},
	RainSum:        whereHelperfloat64{field: "\"prh\".\"daily_data\".\"rain_sum\""},
	CloudMax:       whereHelperfloat64{field: "\"prh\".\"daily_data\".\"cloud_max\""},
	CloudAvg:       whereHelperfloat64{field: "\"prh\".\"daily_data\".\"cloud_avg\""},
	CloudSum:       whereHelperfloat64{field: "\"prh\".\"daily_data\".\"cloud_sum\""},
}

// DailyDatumRels is where relationship names are stored.
var DailyDatumRels = struct {
}{}

// dailyDatumR is where relationships are stored.
type dailyDatumR struct {
}

// NewStruct creates a new relationship struct
func (*dailyDatumR) NewStruct() *dailyDatumR {
	return &dailyDatumR{}
}

// dailyDatumL is where Load methods for each relationship are stored.
type dailyDatumL struct{}

var (
	dailyDatumAllColumns            = []string{"id", "created_at", "date", "sunrise", "sunset", "summary_date", "bed_time_temp", "fan_on", "temperature_max", "temperature_avg", "temperature_sum", "feels_like_max", "feels_like_avg", "feels_like_sum", "uv_max", "uv_avg", "uv_sum", "rain_max", "rain_avg", "rain_sum", "cloud_max", "cloud_avg", "cloud_sum"}
	dailyDatumColumnsWithoutDefault = []string{"date", "sunrise", "sunset", "summary_date", "bed_time_temp", "fan_on", "temperature_max", "temperature_avg", "temperature_sum", "feels_like_max", "feels_like_avg", "feels_like_sum", "uv_max", "uv_avg", "uv_sum", "rain_max", "rain_avg", "rain_sum", "cloud_max", "cloud_avg", "cloud_sum"}
	dailyDatumColumnsWithDefault    = []string{"id", "created_at"}
	dailyDatumPrimaryKeyColumns     = []string{"id"}
	dailyDatumGeneratedColumns      = []string{}
)

type (
	// DailyDatumSlice is an alias for a slice of pointers to DailyDatum.
	// This should almost always be used instead of []DailyDatum.
	DailyDatumSlice []*DailyDatum
	// DailyDatumHook is the signature for custom DailyDatum hook methods
	DailyDatumHook func(context.Context, boil.ContextExecutor, *DailyDatum) error

	dailyDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dailyDatumType                 = reflect.TypeOf(&DailyDatum{})
	dailyDatumMapping              = queries.MakeStructMapping(dailyDatumType)
	dailyDatumPrimaryKeyMapping, _ = queries.BindMapping(dailyDatumType, dailyDatumMapping, dailyDatumPrimaryKeyColumns)
	dailyDatumInsertCacheMut       sync.RWMutex
	dailyDatumInsertCache          = make(map[string]insertCache)
	dailyDatumUpdateCacheMut       sync.RWMutex
	dailyDatumUpdateCache          = make(map[string]updateCache)
	dailyDatumUpsertCacheMut       sync.RWMutex
	dailyDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dailyDatumAfterSelectHooks []DailyDatumHook

var dailyDatumBeforeInsertHooks []DailyDatumHook
var dailyDatumAfterInsertHooks []DailyDatumHook

var dailyDatumBeforeUpdateHooks []DailyDatumHook
var dailyDatumAfterUpdateHooks []DailyDatumHook

var dailyDatumBeforeDeleteHooks []DailyDatumHook
var dailyDatumAfterDeleteHooks []DailyDatumHook

var dailyDatumBeforeUpsertHooks []DailyDatumHook
var dailyDatumAfterUpsertHooks []DailyDatumHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DailyDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dailyDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DailyDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dailyDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DailyDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dailyDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DailyDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dailyDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DailyDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dailyDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DailyDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dailyDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DailyDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dailyDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DailyDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dailyDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DailyDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dailyDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDailyDatumHook registers your hook function for all future operations.
func AddDailyDatumHook(hookPoint boil.HookPoint, dailyDatumHook DailyDatumHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		dailyDatumAfterSelectHooks = append(dailyDatumAfterSelectHooks, dailyDatumHook)
	case boil.BeforeInsertHook:
		dailyDatumBeforeInsertHooks = append(dailyDatumBeforeInsertHooks, dailyDatumHook)
	case boil.AfterInsertHook:
		dailyDatumAfterInsertHooks = append(dailyDatumAfterInsertHooks, dailyDatumHook)
	case boil.BeforeUpdateHook:
		dailyDatumBeforeUpdateHooks = append(dailyDatumBeforeUpdateHooks, dailyDatumHook)
	case boil.AfterUpdateHook:
		dailyDatumAfterUpdateHooks = append(dailyDatumAfterUpdateHooks, dailyDatumHook)
	case boil.BeforeDeleteHook:
		dailyDatumBeforeDeleteHooks = append(dailyDatumBeforeDeleteHooks, dailyDatumHook)
	case boil.AfterDeleteHook:
		dailyDatumAfterDeleteHooks = append(dailyDatumAfterDeleteHooks, dailyDatumHook)
	case boil.BeforeUpsertHook:
		dailyDatumBeforeUpsertHooks = append(dailyDatumBeforeUpsertHooks, dailyDatumHook)
	case boil.AfterUpsertHook:
		dailyDatumAfterUpsertHooks = append(dailyDatumAfterUpsertHooks, dailyDatumHook)
	}
}

// One returns a single dailyDatum record from the query.
func (q dailyDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DailyDatum, error) {
	o := &DailyDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: failed to execute a one query for daily_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DailyDatum records from the query.
func (q dailyDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DailyDatumSlice, error) {
	var o []*DailyDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db: failed to assign all query results to DailyDatum slice")
	}

	if len(dailyDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DailyDatum records in the query.
func (q dailyDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to count daily_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dailyDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db: failed to check if daily_data exists")
	}

	return count > 0, nil
}

// DailyData retrieves all the records using an executor.
func DailyData(mods ...qm.QueryMod) dailyDatumQuery {
	mods = append(mods, qm.From("\"prh\".\"daily_data\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"prh\".\"daily_data\".*"})
	}

	return dailyDatumQuery{q}
}

// FindDailyDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDailyDatum(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DailyDatum, error) {
	dailyDatumObj := &DailyDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"prh\".\"daily_data\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, dailyDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: unable to select from daily_data")
	}

	if err = dailyDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return dailyDatumObj, err
	}

	return dailyDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DailyDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db: no daily_data provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dailyDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dailyDatumInsertCacheMut.RLock()
	cache, cached := dailyDatumInsertCache[key]
	dailyDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dailyDatumAllColumns,
			dailyDatumColumnsWithDefault,
			dailyDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dailyDatumType, dailyDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dailyDatumType, dailyDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"prh\".\"daily_data\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"prh\".\"daily_data\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "db: unable to insert into daily_data")
	}

	if !cached {
		dailyDatumInsertCacheMut.Lock()
		dailyDatumInsertCache[key] = cache
		dailyDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DailyDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DailyDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return err
	}
	key := makeCacheKey(columns, nil)
	dailyDatumUpdateCacheMut.RLock()
	cache, cached := dailyDatumUpdateCache[key]
	dailyDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dailyDatumAllColumns,
			dailyDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("db: unable to update daily_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"prh\".\"daily_data\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dailyDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dailyDatumType, dailyDatumMapping, append(wl, dailyDatumPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "db: unable to update daily_data row")
	}

	if !cached {
		dailyDatumUpdateCacheMut.Lock()
		dailyDatumUpdateCache[key] = cache
		dailyDatumUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dailyDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "db: unable to update all for daily_data")
	}

	return nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DailyDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("db: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dailyDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"prh\".\"daily_data\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, dailyDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "db: unable to update all in dailyDatum slice")
	}

	return nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DailyDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db: no daily_data provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dailyDatumColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	dailyDatumUpsertCacheMut.RLock()
	cache, cached := dailyDatumUpsertCache[key]
	dailyDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dailyDatumAllColumns,
			dailyDatumColumnsWithDefault,
			dailyDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dailyDatumAllColumns,
			dailyDatumPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("db: unable to upsert daily_data, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dailyDatumPrimaryKeyColumns))
			copy(conflict, dailyDatumPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"prh\".\"daily_data\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(dailyDatumType, dailyDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dailyDatumType, dailyDatumMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "db: unable to upsert daily_data")
	}

	if !cached {
		dailyDatumUpsertCacheMut.Lock()
		dailyDatumUpsertCache[key] = cache
		dailyDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DailyDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DailyDatum) Delete(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil {
		return errors.New("db: no DailyDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dailyDatumPrimaryKeyMapping)
	sql := "DELETE FROM \"prh\".\"daily_data\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "db: unable to delete from daily_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return err
	}

	return nil
}

// DeleteAll deletes all matching rows.
func (q dailyDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if q.Query == nil {
		return errors.New("db: no dailyDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return errors.Wrap(err, "db: unable to delete all from daily_data")
	}

	return nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DailyDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) error {
	if len(o) == 0 {
		return nil
	}

	if len(dailyDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dailyDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"prh\".\"daily_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dailyDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	_, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "db: unable to delete all from dailyDatum slice")
	}

	if len(dailyDatumAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DailyDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDailyDatum(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DailyDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DailyDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dailyDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"prh\".\"daily_data\".* FROM \"prh\".\"daily_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dailyDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db: unable to reload all in DailyDatumSlice")
	}

	*o = slice

	return nil
}

// DailyDatumExists checks if the DailyDatum row exists.
func DailyDatumExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"prh\".\"daily_data\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db: unable to check if daily_data exists")
	}

	return exists, nil
}

// Exists checks if the DailyDatum row exists.
func (o *DailyDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DailyDatumExists(ctx, exec, o.ID)
}
