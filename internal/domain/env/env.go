/*
Package env provides the environment domain models used throughout the application.
*/
package env

// Env 環境を表す。
//
//go:generate go run github.com/anoriqq/enumer@latest -type=Env -transform=lower
type Env int

// 環境一覧。
const (
	PRD Env = iota
	STG
	DEV
	LCL
)

// FromStringZero 文字列から [Env] を取得する。
// 指定された文字列が不正な場合はゼロ値を返す。
func FromStringZero(s string) Env {
	e, _ := EnvString(s)

	return e
}
