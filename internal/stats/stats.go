package stats

import (
	"time"

	"github.com/volatiletech/null"
)

type Datum struct {
	Time  time.Time
	Value float32
	Valid bool
}

func (d Datum) MinusTime(other Datum) float32 {
	return float32(d.Time.Sub(other.Time) / time.Second)
}

func (d Datum) MinusValue(other Datum) float32 {
	return d.Value - other.Value
}

type Data []Datum

func DatumFrom(t time.Time, f null.Float32) Datum {
	return Datum{
		Time:  t,
		Value: f.Float32,
		Valid: f.Valid,
	}
}

func (d Data) OnlyValid() Data {
	data := make(Data, 0, len(d))
	for _, datum := range d {
		if datum.Valid {
			data = append(data, datum)
		}
	}
	return data
}

func (d Data) Max() float32 {
	var max float32
	for _, v := range d.OnlyValid() {
		if v.Valid && v.Value > max {
			max = v.Value
		}
	}
	return max
}

func (d Data) AreaSum() float32 {
	data := d.OnlyValid()
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0].Value
	}

	var sum float32

	for i := 0; i < len(data)-1; i++ {
		v1, v2 := data[i], data[i+1]

		sum += v2.MinusTime(v1) *
			(v1.Value + 0.5*(v2.MinusValue(v1)))
	}

	return sum
}

func (d Data) Avg() float32 {
	data := d.OnlyValid()
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0].Value
	}

	timeDiff := data[len(data)-1].MinusTime(data[0])
	return data.AreaSum() / timeDiff
}
