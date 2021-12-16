package main

import "testing"

func TestMiSuma(t *testing.T) {
	type prueba struct {
		datos     []int
		respuesta int
	}

	pruebas := []prueba{
		{
			[]int{1, 2, 3, 4, 5}, 15,
		},
		{
			[]int{1, 2, 3, 4, 6}, 16,
		},
		{
			[]int{10, 2, 3, 4, 5}, 24,
		},
	}

	for _, x := range pruebas {
		v := miSuma(x.datos...)
		if v != x.respuesta {
			t.Error("Expected", x.respuesta, "Got", v)
		}
	}
}

func BenchmarkSuma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		miSuma(1, 2, 3, 4, 5, 6, 7, 8)
	}
}
