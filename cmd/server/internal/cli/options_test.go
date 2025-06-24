package cli_test

import (
	"testing"

	"github.com/anoriqq/pj-tmpl-go/cmd/server/internal/cli"
	"github.com/brianvoe/gofakeit/v7"
)

func TestNewOptions(t *testing.T) {
	t.Run("環境変数からオプションを取得できる", func(t *testing.T) {
		want := gofakeit.Word()
		t.Setenv("ENV", want)

		// Act
		got := cli.NewOptions()

		// Assert
		if got.Env() != want {
			t.Errorf("envは %s を得るはずが %s を得た", want, got.Env())
		}
	})

	t.Run("デフォルト値を取得できる", func(t *testing.T) {
		// Act
		got := cli.NewOptions()

		// Assert
		want := ""
		if got.Env() != want {
			t.Errorf("envは %s を得るはずが %s を得た", want, got.Env())
		}
	})
}
