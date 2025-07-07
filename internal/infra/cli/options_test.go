package cli_test

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/anoriqq/pj-tmpl-go/internal/infra/cli"
	"github.com/brianvoe/gofakeit/v7"
)

func TestNewOptions(t *testing.T) {
	// 一部のサブテストが [testing.T.Setenv] を使っているので並列化しない

	environments := []string{"lcl", "dev", "stg", "prd"}

	t.Run("デフォルト値を取得できる", func(t *testing.T) {
		t.Parallel()

		args := []string{"-env", environments[rand.Intn(len(environments))]} // envだけは必須

		// Act
		got, err := cli.NewOptions(args)

		// Assert
		if !errors.Is(err, nil) {
			t.Fatalf("オプションの取得に失敗: %v", err)
		}
		want := uint64(8080)
		if got.Port() != want {
			t.Errorf("デフォルト値である`%d`を得るはずが`%d`を得た", want, got.Port())
		}
	})

	t.Run("環境変数からオプションを取得できる", func(t *testing.T) {
		// [testing.T.Setenv] を使っているので並列化しない

		want := environments[rand.Intn(len(environments))]
		t.Setenv("ENV", want)
		t.Logf("環境変数を設定: %s", want)

		// Act
		got, err := cli.NewOptions(nil)

		// Assert
		if !errors.Is(err, nil) {
			t.Fatalf("オプションの取得に失敗: %v", err)
		}
		if got.Env().String() != want {
			t.Errorf("設定した`%s`を得るはずが`%s`を得た", want, got.Env())
		}
	})

	t.Run("フラグからオプションを取得できる", func(t *testing.T) {
		t.Parallel()

		want := environments[rand.Intn(len(environments))]
		args := []string{"-env", want}

		// Act
		got, err := cli.NewOptions(args)

		// Assert
		if !errors.Is(err, nil) {
			t.Fatalf("オプションの取得に失敗: %v", err)
		}
		if got.Env().String() != want {
			t.Errorf("設定した`%s`を得るはずが`%s`を得た", want, got.Env())
		}
	})

	t.Run("環境変数とフラグの両方が設定されている場合はフラグが優先される", func(t *testing.T) {
		// [testing.T.Setenv] を使っているので並列化しない

		fromEnv := gofakeit.Word()
		t.Setenv("ENV", fmt.Sprint(fromEnv))
		t.Logf("環境変数を設定: %s", fromEnv)

		fromFlag := environments[rand.Intn(len(environments))]
		args := []string{"-env", fromFlag}

		// Act
		got, err := cli.NewOptions(args)

		// Assert
		if !errors.Is(err, nil) {
			t.Fatalf("オプションの取得に失敗: %v", err)
		}
		if got.Env().String() != fromFlag {
			t.Errorf("flagで設定した`%s`を得るはずが`%s`を得た", fromFlag, got.Env())
		}
	})
}
