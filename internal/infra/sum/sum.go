/*
Package sum provides a sum types.
*/
package sum

const (
	tag1 = iota
	tag2
	tag3
)

// Union2 2つの型のいずれかを保持することを表す型
type Union2[T1, T2 any] struct {
	tag uint8
	v1  T1 // ゼロ値とみなされる
	v2  T2
}

// Union2T1 2つの型のうち、最初の型を保持するUnion2を作成する
func Union2T1[T1, T2 any](v T1) Union2[T1, T2] {
	return Union2[T1, T2]{
		tag: tag1,
		v1:  v,
	}
}

// Union2T2 2つの型のうち、2番目の型を保持するUnion2を作成する
func Union2T2[T1, T2 any](v T2) Union2[T1, T2] {
	return Union2[T1, T2]{
		tag: tag2,
		v2:  v,
	}
}

// Match2 Union2の値に応じて異なる関数を呼び出す
func Match2[T1, T2, R any](
	u Union2[T1, T2],
	f1 func(T1) R,
	f2 func(T2) R,
) R {
	switch u.tag {
	case tag2:
		return f2(u.v2)
	case tag1:
		fallthrough
	default:
		return f1(u.v1)
	}
}

// Union3 3つの型のいずれかを保持することを表す型
type Union3[T1, T2, T3 any] struct {
	tag uint8
	v1  T1 // ゼロ値とみなされる
	v2  T2
	v3  T3
}

// Union3T1 3つの型のうち、最初の型を保持するUnion3を作成する
func Union3T1[T1, T2, T3 any](v T1) Union3[T1, T2, T3] {
	return Union3[T1, T2, T3]{
		tag: tag1,
		v1:  v,
	}
}

// Union3T2 3つの型のうち、2番目の型を保持するUnion3を作成する
func Union3T2[T1, T2, T3 any](v T2) Union3[T1, T2, T3] {
	return Union3[T1, T2, T3]{
		tag: tag2,
		v2:  v,
	}
}

// Union3T3 3つの型のうち、3番目の型を保持するUnion3を作成する
func Union3T3[T1, T2, T3 any](v T3) Union3[T1, T2, T3] {
	return Union3[T1, T2, T3]{
		tag: tag3,
		v3:  v,
	}
}

// Match3 Union3の値に応じて異なる関数を呼び出す
func Match3[T1, T2, T3, R any](
	u Union3[T1, T2, T3],
	f1 func(T1) R,
	f2 func(T2) R,
	f3 func(T3) R,
) R {
	switch u.tag {
	case tag3:
		return f3(u.v3)
	case tag2:
		return f2(u.v2)
	case tag1:
		fallthrough
	default:
		return f1(u.v1)
	}
}
