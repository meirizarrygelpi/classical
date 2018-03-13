package classical

// Bit is a classical bit.
type Bit struct {
	value bool
}

// Value returns the boolean value of z.
func (z *Bit) Value() bool {
	return z.value
}

// SetValue sets the boolean value of z equal to a.
func (z *Bit) SetValue(a bool) {
	z.value = a
}

// String returns the string version of a classical bit.
func (z *Bit) String() string {
	if z.Value() {
		return "1"
	}
	return "0"
}

// NewBit returns a classical bit with boolean value a.
func NewBit(a bool) *Bit {
	return &Bit{a}
}

// Copy copies the value of z onto y.
func Copy(z, y *Bit) {
	y.SetValue(z.Value())
}

// Not gate (reversible).
func Not(z *Bit) *Bit {
	return NewBit(!z.Value())
}

// And gate.
func And(z, y *Bit) *Bit {
	return NewBit(z.Value() && y.Value())
}

// NAnd gate.
func NAnd(z, y *Bit) *Bit {
	return NewBit(Not(And(z, y)).Value())
}

// Or gate.
func Or(z, y *Bit) *Bit {
	return NewBit(z.Value() || y.Value())
}

// NOr gate.
func NOr(z, y *Bit) *Bit {
	return NewBit(Not(Or(z, y)).Value())
}

// XOr gate.
func XOr(z, y *Bit) *Bit {
	return NewBit(z.Value() != y.Value())
}

// NXOr gate.
func NXOr(z, y *Bit) *Bit {
	return NewBit(z.Value() == y.Value())
}

// FanOut gate.
func FanOut(z *Bit) (*Bit, *Bit) {
	return z, NewBit(z.Value())
}

// Swap gate (reversible).
func Swap(z, y *Bit) (*Bit, *Bit) {
	return y, z
}

// CNot gate (reversible).
func CNot(z, y *Bit) (*Bit, *Bit) {
	return z, XOr(z, y)
}

// Fredkin gate (reversible).
func Fredkin(z, y, x *Bit) (*Bit, *Bit, *Bit) {
	if z.Value() {
		return z, x, y
	}
	return z, y, x
}

// Toffoli gate (reversible).
func Toffoli(z, y, x *Bit) (*Bit, *Bit, *Bit) {
	return z, y, XOr(x, And(z, y))
}

// Swap243 gate (reversible).
func Swap243(z, y, x, w *Bit) (*Bit, *Bit, *Bit, *Bit) {
	z, x, w = Fredkin(z, x, w)
	z = Not(z)
	z, y, w = Fredkin(z, y, w)
	z = Not(z)
	return z, y, x, w
}

// CCSwap gate (reversible).
func CCSwap(z, y, x, w, v *Bit) (*Bit, *Bit, *Bit, *Bit, *Bit) {
	z, y, v = Toffoli(z, y, v)
	v, x, w = Fredkin(v, x, w)
	z, y, v = Toffoli(z, y, v)
	return z, y, x, w, v
}

// DeMultiplexer gate (reversible).
func DeMultiplexer(z, y, x, w *Bit) (*Bit, *Bit, *Bit, *Bit) {
	z, w = CNot(z, w)
	y, x = CNot(y, x)
	z, y, x = Toffoli(z, y, x)
	w = Not(w)
	w, y, z = Toffoli(w, y, z)
	z = Not(z)
	z, y = CNot(z, y)
	y = Not(y)
	y, w = CNot(y, w)
	w = Not(w)
	return z, y, x, w
}
