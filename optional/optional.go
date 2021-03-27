package optional

// Option option
type Option struct {
	value interface{}
}

// None returns none Option
func None() Option {
	return Option{value: nil}
}

// Some returns a Option with a given value
func Some(value interface{}) Option {
	return Option{
		value: value,
	}
}

// Unwrap unwap the value
func (opt Option) Unwrap() interface{} {
	return opt.value
}

// UnwrapOrDefault returns some value of default value
func (opt Option) UnwrapOrDefault(defaultValue interface{}) interface{} {
	if opt.IsNone() {
		return defaultValue
	}

	return opt.value
}

// UnwrapOrElse returns some value or computes it from a func
func (opt Option) UnwrapOrElse(fn func() interface{}) interface{} {
	if opt.IsNone() {
		return fn()
	}

	return opt.value
}

// Or returns the option if it contains a value, otherwise returns optb.
func (opt Option) Or(or Option) Option {
	if opt.IsSome() {
		return opt
	}

	return or
}

// Or returns the option if it contains a value, otherwise calls f and returns the result.
func (opt Option) OrElse(fn func() Option) Option {
	if opt.IsSome() {
		return opt
	}

	return fn()
}

// IsSome return true means has some value
func (opt Option) IsSome() bool {
	return opt.value != nil
}

// IsNone return true means has no some value
func (opt Option) IsNone() bool {
	return !opt.IsSome()
}
