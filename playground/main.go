package main

func Do[T any](v T) {
	println(v)
}

type Custom struct {
	Value int
}

func main() {
	Do(10)
	Do(10.0)
	Do("10.0")
	Do(Custom{
		Value: 10,
	})
	Do(&Custom{
		Value: 10,
	})
}

// Default:
// go tool compile -N -l -S main.go > main.asm
//
// Monomorphized:
// go tool compile '-d=unified=1' -N -l  -S main.go > main.asm
