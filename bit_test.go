package classical_test

import (
	"fmt"

	"github.com/meirizarrygelpi/classical"
)

var booleans = [2]bool{false, true}

func ExampleCopy() {
	z, y := new(classical.Bit), new(classical.Bit)
	for _, a := range booleans {
		y = classical.NewBit(a)
		classical.Copy(y, z)
		fmt.Println(y, "|", z)
	}
	// Output:
	// 0 | 0
	// 1 | 1
}

func ExampleNot() {
	z := new(classical.Bit)
	for _, a := range booleans {
		z = classical.NewBit(a)
		fmt.Println(z, "|", classical.Not(z))
	}
	// Output:
	// 0 | 1
	// 1 | 0
}

func ExampleAnd() {
	z, y := new(classical.Bit), new(classical.Bit)
	for _, a := range booleans {
		for _, b := range booleans {
			z = classical.NewBit(a)
			y = classical.NewBit(b)
			fmt.Println(z, y, "|", classical.And(z, y))
		}
	}
	// Output:
	// 0 0 | 0
	// 0 1 | 0
	// 1 0 | 0
	// 1 1 | 1
}

func ExampleNAnd() {
	z, y := new(classical.Bit), new(classical.Bit)
	for _, a := range booleans {
		for _, b := range booleans {
			z = classical.NewBit(a)
			y = classical.NewBit(b)
			fmt.Println(z, y, "|", classical.NAnd(z, y))
		}
	}
	// Output:
	// 0 0 | 1
	// 0 1 | 1
	// 1 0 | 1
	// 1 1 | 0
}

func ExampleOr() {
	z, y := new(classical.Bit), new(classical.Bit)
	for _, a := range booleans {
		for _, b := range booleans {
			z = classical.NewBit(a)
			y = classical.NewBit(b)
			fmt.Println(z, y, "|", classical.Or(z, y))
		}
	}
	// Output:
	// 0 0 | 0
	// 0 1 | 1
	// 1 0 | 1
	// 1 1 | 1
}

func ExampleNOr() {
	z, y := new(classical.Bit), new(classical.Bit)
	for _, a := range booleans {
		for _, b := range booleans {
			z = classical.NewBit(a)
			y = classical.NewBit(b)
			fmt.Println(z, y, "|", classical.NOr(z, y))
		}
	}
	// Output:
	// 0 0 | 1
	// 0 1 | 0
	// 1 0 | 0
	// 1 1 | 0
}

func ExampleXOr() {
	z, y := new(classical.Bit), new(classical.Bit)
	for _, a := range booleans {
		for _, b := range booleans {
			z = classical.NewBit(a)
			y = classical.NewBit(b)
			fmt.Println(z, y, "|", classical.XOr(z, y))
		}
	}
	// Output:
	// 0 0 | 0
	// 0 1 | 1
	// 1 0 | 1
	// 1 1 | 0
}

func ExampleNXOr() {
	z, y := new(classical.Bit), new(classical.Bit)
	for _, a := range booleans {
		for _, b := range booleans {
			z = classical.NewBit(a)
			y = classical.NewBit(b)
			fmt.Println(z, y, "|", classical.NXOr(z, y))
		}
	}
	// Output:
	// 0 0 | 1
	// 0 1 | 0
	// 1 0 | 0
	// 1 1 | 1
}

func ExampleFanOut() {
	z := new(classical.Bit)
	for _, a := range booleans {
		z = classical.NewBit(a)
		fmt.Print(z)
		fmt.Print(" | ")
		fmt.Print(classical.FanOut(z))
		fmt.Print("\n")
	}
	// Output:
	// 0 | 0 0
	// 1 | 1 1
}

func ExampleSwap() {
	z, y := new(classical.Bit), new(classical.Bit)
	for _, a := range booleans {
		for _, b := range booleans {
			z = classical.NewBit(a)
			y = classical.NewBit(b)
			fmt.Print(z, y)
			fmt.Print(" | ")
			fmt.Print(classical.Swap(z, y))
			fmt.Print("\n")
		}
	}
	// Output:
	// 0 0 | 0 0
	// 0 1 | 1 0
	// 1 0 | 0 1
	// 1 1 | 1 1
}

func ExampleCNot() {
	z, y := new(classical.Bit), new(classical.Bit)
	for _, a := range booleans {
		for _, b := range booleans {
			z = classical.NewBit(a)
			y = classical.NewBit(b)
			fmt.Print(z, y)
			fmt.Print(" | ")
			fmt.Print(classical.CNot(z, y))
			fmt.Print("\n")
		}
	}
	// Output:
	// 0 0 | 0 0
	// 0 1 | 0 1
	// 1 0 | 1 1
	// 1 1 | 1 0
}

func ExampleFredkin() {
	z, y, x := new(classical.Bit), new(classical.Bit), new(classical.Bit)
	for _, a := range booleans {
		for _, b := range booleans {
			for _, c := range booleans {
				z = classical.NewBit(a)
				y = classical.NewBit(b)
				x = classical.NewBit(c)
				fmt.Print(z, y, x)
				fmt.Print(" | ")
				fmt.Print(classical.Fredkin(z, y, x))
				fmt.Print("\n")
			}
		}
	}
	// Output:
	// 0 0 0 | 0 0 0
	// 0 0 1 | 0 0 1
	// 0 1 0 | 0 1 0
	// 0 1 1 | 0 1 1
	// 1 0 0 | 1 0 0
	// 1 0 1 | 1 1 0
	// 1 1 0 | 1 0 1
	// 1 1 1 | 1 1 1

}

func ExampleToffoli() {
	z, y, x := new(classical.Bit), new(classical.Bit), new(classical.Bit)
	for _, a := range booleans {
		for _, b := range booleans {
			for _, c := range booleans {
				z = classical.NewBit(a)
				y = classical.NewBit(b)
				x = classical.NewBit(c)
				fmt.Print(z, y, x)
				fmt.Print(" | ")
				fmt.Print(classical.Toffoli(z, y, x))
				fmt.Print("\n")
			}
		}
	}
	// Output:
	// 0 0 0 | 0 0 0
	// 0 0 1 | 0 0 1
	// 0 1 0 | 0 1 0
	// 0 1 1 | 0 1 1
	// 1 0 0 | 1 0 0
	// 1 0 1 | 1 0 1
	// 1 1 0 | 1 1 1
	// 1 1 1 | 1 1 0

}

func ExampleSwap243() {
	z, y, x, w := new(classical.Bit), new(classical.Bit), new(classical.Bit), new(classical.Bit)
	for _, a := range booleans {
		for _, b := range booleans {
			for _, c := range booleans {
				for _, d := range booleans {
					z = classical.NewBit(a)
					y = classical.NewBit(b)
					x = classical.NewBit(c)
					w = classical.NewBit(d)
					fmt.Print(z, y, x, w)
					fmt.Print(" | ")
					fmt.Print(classical.Swap243(z, y, x, w))
					fmt.Print("\n")
				}
			}
		}
	}
	// Output:
	// 0 0 0 0 | 0 0 0 0
	// 0 0 0 1 | 0 1 0 0
	// 0 0 1 0 | 0 0 1 0
	// 0 0 1 1 | 0 1 1 0
	// 0 1 0 0 | 0 0 0 1
	// 0 1 0 1 | 0 1 0 1
	// 0 1 1 0 | 0 0 1 1
	// 0 1 1 1 | 0 1 1 1
	// 1 0 0 0 | 1 0 0 0
	// 1 0 0 1 | 1 0 1 0
	// 1 0 1 0 | 1 0 0 1
	// 1 0 1 1 | 1 0 1 1
	// 1 1 0 0 | 1 1 0 0
	// 1 1 0 1 | 1 1 1 0
	// 1 1 1 0 | 1 1 0 1
	// 1 1 1 1 | 1 1 1 1

}

func ExampleCCSwap() {
	z, y, x, w := new(classical.Bit), new(classical.Bit), new(classical.Bit), new(classical.Bit)
	v := classical.NewBit(false)
	for _, a := range booleans {
		for _, b := range booleans {
			for _, c := range booleans {
				for _, d := range booleans {
					z = classical.NewBit(a)
					y = classical.NewBit(b)
					x = classical.NewBit(c)
					w = classical.NewBit(d)
					fmt.Print(z, y, x, w, v)
					fmt.Print(" | ")
					fmt.Print(classical.CCSwap(z, y, x, w, v))
					fmt.Print("\n")
				}
			}
		}
	}
	// Output:
	// 0 0 0 0 0 | 0 0 0 0 0
	// 0 0 0 1 0 | 0 0 0 1 0
	// 0 0 1 0 0 | 0 0 1 0 0
	// 0 0 1 1 0 | 0 0 1 1 0
	// 0 1 0 0 0 | 0 1 0 0 0
	// 0 1 0 1 0 | 0 1 0 1 0
	// 0 1 1 0 0 | 0 1 1 0 0
	// 0 1 1 1 0 | 0 1 1 1 0
	// 1 0 0 0 0 | 1 0 0 0 0
	// 1 0 0 1 0 | 1 0 0 1 0
	// 1 0 1 0 0 | 1 0 1 0 0
	// 1 0 1 1 0 | 1 0 1 1 0
	// 1 1 0 0 0 | 1 1 0 0 0
	// 1 1 0 1 0 | 1 1 1 0 0
	// 1 1 1 0 0 | 1 1 0 1 0
	// 1 1 1 1 0 | 1 1 1 1 0

}

func ExampleDeMultiplexer() {
	z, y := new(classical.Bit), new(classical.Bit)
	x := classical.NewBit(false)
	w := classical.NewBit(false)
	for _, a := range booleans {
		for _, b := range booleans {
			z = classical.NewBit(a)
			y = classical.NewBit(b)
			fmt.Print(z, y, x, w)
			fmt.Print(" | ")
			fmt.Print(classical.DeMultiplexer(z, y, x, w))
			fmt.Print("\n")
		}
	}
	// Output:
	// 0 0 0 0 | 1 0 0 0
	// 0 1 0 0 | 0 0 1 0
	// 1 0 0 0 | 0 1 0 0
	// 1 1 0 0 | 0 0 0 1
}
