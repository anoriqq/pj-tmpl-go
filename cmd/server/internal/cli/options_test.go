package cli_test

import (
	"flag"
	"fmt"
	"testing"

	"github.com/anoriqq/pj-tmpl-go/cmd/server/internal/cli"
	"github.com/brianvoe/gofakeit/v7"
)

func TestNewOptions(t *testing.T) {
	// [flag] packageの操作がスレッドセーフでないので並列化しない

	t.Run("デフォルト値を取得できる", func(t *testing.T) {
		// Act
		got := cli.NewOptions()

		// Assert
		want := ""
		if got.Env() != want {
			t.Errorf("デフォルト値である`%s`を得るはずが`%s`を得た", want, got.Env())
		}
	})

	t.Run("環境変数からオプションを取得できる", func(t *testing.T) {
		want := gofakeit.Word()
		t.Setenv("ENV", want)
		t.Logf("環境変数を設定: %s", want)

		// Act
		got := cli.NewOptions()

		// Assert
		if got.Env() != want {
			t.Errorf("設定した`%s`を得るはずが`%s`を得た", want, got.Env())
		}
	})

	t.Run("フラグからオプションを取得できる", func(t *testing.T) {
		want := gofakeit.Word()
		if err := flag.CommandLine.Set("env", want); err != nil {
			t.Fatalf("flagの設定に失敗: %v", err)
		}
		t.Logf("flagを設定: %s", want)

		// Act
		got := cli.NewOptions()

		// Assert
		if got.Env() != want {
			t.Errorf("設定した`%s`を得るはずが`%s`を得た", want, got.Env())
		}
	})

	t.Run("環境変数とフラグの両方が設定されている場合はフラグが優先される", func(t *testing.T) {
		fromEnv := gofakeit.Word()
		t.Setenv("ENV", fmt.Sprint(fromEnv))
		t.Logf("環境変数を設定: %s", fromEnv)

		fromFlag := gofakeit.Word()
		if err := flag.CommandLine.Set("env", fmt.Sprint(fromFlag)); err != nil {
			t.Fatalf("flagの設定に失敗: %v", err)
		}
		t.Logf("flagを設定: %s", fromFlag)

		// Act
		got := cli.NewOptions()

		// Assert
		if got.Env() != fromFlag {
			t.Errorf("flagで設定した`%s`を得るはずが`%s`を得た", fromFlag, got.Env())
		}
	})
}
