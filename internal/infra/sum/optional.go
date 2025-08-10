package sum

// None 値なしを表す型。
type None struct{}

// Option 値が無い場合を表現できる型。
type Option[T any] = Union2[T, None]

// SomeOf 値がある場合のOptionを生成する。
func SomeOf[T any](v T) Option[T] {
	return Union2T1[T, None](v)
}

// NoneOf 値がない場合のOptionを生成する。
func NoneOf[T any]() Option[T] {
	return Union2T2[T](None{})
}

// UnwrapOr 値がある場合はその値を、ない場合はデフォルト値を返す。
func UnwrapOr[T any](o Option[T], def T) T {
	return Match2(
		o,
		func(v T) T { return v },
		func(None) T { return def },
	)
}
