package stats

import (
	"reflect"
	"testing"
	"time"

	"github.com/volatiletech/null"
)

func TestData_AreaSum(t *testing.T) {
	tests := []struct {
		name string
		d    Data
		want float32
	}{
		{
			name: "flat 60 seconds",
			d: Data{
				Datum{
					Time: time.Date(2020, 7, 14, 22, 3, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 7, 14, 22, 4, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 8, 14, 22, 3, 0, 0, time.UTC),
					Value: 100000000,
					Valid: false,
				},
			},
			want: 180,
		},
		{
			name: "up 60 seconds",
			d: Data{
				Datum{
					Time: time.Date(2020, 7, 14, 22, 3, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 7, 14, 22, 4, 0, 0, time.UTC),
					Value: 5,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 8, 14, 22, 3, 0, 0, time.UTC),
					Value: 100000000,
					Valid: false,
				},
			},
			want: 240,
		},
		{
			name: "multiple",
			d: Data{
				Datum{
					Time: time.Date(2020, 7, 14, 22, 3, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 7, 14, 22, 4, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 7, 14, 22, 5, 0, 0, time.UTC),
					Value: 5,
					Valid: true,
				},
			},
			want: 420,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.AreaSum(); got != tt.want {
				t.Errorf("AreaSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Avg(t *testing.T) {
	tests := []struct {
		name string
		d    Data
		want float32
	}{
		{
			name: "flat 60 seconds",
			d: Data{
				Datum{
					Time: time.Date(2020, 7, 14, 22, 3, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 7, 14, 22, 4, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 8, 14, 22, 3, 0, 0, time.UTC),
					Value: 100000000,
					Valid: false,
				},
			},
			want: 3,
		},
		{
			name: "up 60 seconds",
			d: Data{
				Datum{
					Time: time.Date(2020, 7, 14, 22, 3, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 7, 14, 22, 4, 0, 0, time.UTC),
					Value: 5,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 8, 14, 22, 3, 0, 0, time.UTC),
					Value: 100000000,
					Valid: false,
				},
			},
			want: 4,
		},
		{
			name: "multiple",
			d: Data{
				Datum{
					Time: time.Date(2020, 7, 14, 22, 3, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 7, 14, 22, 4, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 7, 14, 22, 5, 0, 0, time.UTC),
					Value: 5,
					Valid: true,
				},
			},
			want: 3.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Avg(); got != tt.want {
				t.Errorf("Avg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_Max(t *testing.T) {
	tests := []struct {
		name string
		d    Data
		want float32
	}{
		{
			name: "simple",
			d: Data{
				Datum{
					Time: time.Date(2020, 7, 14, 22, 3, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 8, 14, 22, 3, 0, 0, time.UTC),
					Value: 6,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 8, 14, 22, 3, 0, 0, time.UTC),
					Value: 100000000,
					Valid: false,
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Max(); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_OnlyValid(t *testing.T) {
	tests := []struct {
		name string
		d    Data
		want Data
	}{
		{
			name: "simple",
			d: Data{
				Datum{
					Time: time.Date(2020, 7, 14, 22, 3, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
				Datum{
					Time: time.Date(2020, 8, 14, 22, 3, 0, 0, time.UTC),
					Value: 0,
					Valid: false,
				},
			},
			want: Data{
				Datum{
					Time: time.Date(2020, 7, 14, 22, 3, 0, 0, time.UTC),
					Value: 3,
					Valid: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.OnlyValid(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnlyValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatumFrom(t *testing.T) {
	type args struct {
		t time.Time
		f null.Float32
	}
	tests := []struct {
		name string
		args args
		want Datum
	}{
		{
			name: "simple",
			args: args{
				t: time.Date(2020, 7, 14, 22, 3, 0, 0, time.UTC),
				f: null.Float32From(3),
			},
			want: Datum{
				Time: time.Date(2020, 7, 14, 22, 3, 0, 0, time.UTC),
				Value: 3,
				Valid: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DatumFrom(tt.args.t, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DatumFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatum_MinusTime(t *testing.T) {
	type fields struct {
		Time  time.Time
		Value float32
		Valid bool
	}
	type args struct {
		other Datum
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		{
			name: "one minute",
			fields: fields{
				Time: time.Date(2020, 7, 14, 22, 3, 0, 0, time.UTC),
			},
			args: args{
				other: Datum{
					Time: time.Date(2020, 7, 14, 22, 2, 0, 0, time.UTC),
				},
			},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Datum{
				Time:  tt.fields.Time,
				Value: tt.fields.Value,
				Valid: tt.fields.Valid,
			}
			if got := d.MinusTime(tt.args.other); got != tt.want {
				t.Errorf("MinusTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatum_MinusValue(t *testing.T) {
	type fields struct {
		Time  time.Time
		Value float32
		Valid bool
	}
	type args struct {
		other Datum
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		{
			name: "simple",
			fields: fields{
				Value: 7,
			},
			args: args{
				other: Datum{
					Value: 3,
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Datum{
				Time:  tt.fields.Time,
				Value: tt.fields.Value,
				Valid: tt.fields.Valid,
			}
			if got := d.MinusValue(tt.args.other); got != tt.want {
				t.Errorf("MinusValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
