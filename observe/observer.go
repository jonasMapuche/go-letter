package observer

type Color struct {
	Red   float64
	Green float64
	Blue  float64
	Alpha float64
}

type User struct {
	Name  string
	Image string
	Color Color
}
