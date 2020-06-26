package ptr

func Float64(f float64) *float64 {
	return &f
}

func Float32(f float32) *float32 {
	return &f
}

func Str(s string) *string {
	return &s
}

func Bool(b bool) *bool {
	return &b
}

func True() *bool {
	return Bool(true)
}

func False() *bool {
	return Bool(false)
}
