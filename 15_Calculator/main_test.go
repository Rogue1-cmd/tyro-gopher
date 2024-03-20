package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(3, 5)
	if result != 10 {
		t.Errorf("Add failed. Expected %d, got %d", 10, result)
	}
}

func TestSub(t *testing.T) {
	result := sub(8, 3)
	if result != 5 {
		t.Errorf("Sub failed. Expected %d, got %d", 5, result)
	}
}

func TestProduct(t *testing.T) {
	result := product(4, 6)
	if result != 24 {
		t.Errorf("Product failed. Expected %d, got %d", 24, result)
	}
}

func TestDiv(t *testing.T) {
	result := div(10, 2)
	if result != 5 {
		t.Errorf("Div failed. Expected %d, got %d", 5, result)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Settting up test")
}

func TestCalc(t *testing.T) {

	exp := 20
	result := add(10, 10)

	if result == exp {
		fmt.Println("Test Pass")
	} else {
		fmt.Println("Test Failed")
	}

	t.Run("Test Sub", func(t *testing.T) {
		exp := 12
		result := sub(20, 10)

		if result == exp {
			fmt.Println("Test Pass")
		} else {
			fmt.Println("Test Failed")
		}

	})

	t.Run("Test Sub", func(t *testing.T) {
		exp := 12
		result := sub(20, 10)

		if result == exp {
			fmt.Println("Test Pass")
		} else {
			fmt.Println("Test Failed")
		}

	})
}

//Table Driven Tests

func TestArithmeticOperations(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name   string
		a      int
		b      int
		choice int
		want   int
	}{
		// Test cases
		{"Addition", 3, 5, 1, 8},
		{"Subtraction", 8, 3, 2, 5},
		{"Product", 4, 6, 3, 24},
		{"Division", 10, 2, 4, 5},
	}

	// Execution
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got int
			switch tc.choice {
			case 1:
				got = add(tc.a, tc.b)
			case 2:
				got = sub(tc.a, tc.b)
			case 3:
				got = product(tc.a, tc.b)
			case 4:
				got = div(tc.a, tc.b)
			}

			if got != tc.want {
				t.Errorf("%s failed. Expected %d, got %d", tc.name, tc.want, got)
			}
		})
	}
}

// Benchmark Test
func BenchmarkAdd(b *testing.B) {
	x, y := 10, 5
	for i := 0; i < b.N; i++ {
		add(x, y)
	}
}

func BenchmarkSub(b *testing.B) {
	x, y := 10, 5
	for i := 0; i < b.N; i++ {
		sub(x, y)
	}
}

func BenchmarkProduct(b *testing.B) {
	x, y := 10, 5
	for i := 0; i < b.N; i++ {
		product(x, y)
	}
}

func BenchmarkDiv(b *testing.B) {
	x, y := 10, 5
	for i := 0; i < b.N; i++ {
		div(x, y)
	}
}
