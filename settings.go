package main

type SimSettings struct {
	Iterations   int
	Duration     float64
	PerSecond    float64
	ExecuteRange float64
}

var Settings = SimSettings{
	Iterations:   1,
	Duration:     120,
	ExecuteRange: 20,
	PerSecond:    0.01,
}
