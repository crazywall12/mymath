package mymath

import "math"

// Abs returns the absolute value of x.
func Abs(x float64) float64 {
    return math.Abs(x)
}

// Sqrt returns the square root of x.
func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}

// Pi returns the value of π.
func Pi() float64 {
    return math.Pi
}

// Cos returns the cosine of the radian argument x.
func Cos(x float64) float64 {
    return math.Cos(x)
}

// Sin returns the sine of the radian argument x.
func Sin(x float64) float64 {
    return math.Sin(x)
}

// Exp returns e**x, the base-e exponential of x.
func Exp(x float64) float64 {
    return math.Exp(x)
}