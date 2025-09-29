# Go言語でテストを書く上でのベストプラクティス

## テスト名は条件と期待結果を含める

Go言語における、テスト関数名とサブテストのname引数の値を「テスト名」と呼ぶ。
SUT（System Under Test）の名前を含め、テストの条件と期待結果を明確にすることが重要。

```go
// 悪い例
// 何をテストしているのか不明
func Test(t *testing.T) {
    t.Parallel()
    if Divide(10, 2) != 5 {
        t.Error()
    }
}

// 良い例
func TestDivide(t *testing.T) { // 関数名にはSUTの名前を使う
    t.Parallel()

    t.Run("正の数で除算すると商を返す", func(t *testing.T) { // テスト名に条件と期待する結果を含める
        t.Parallel()

        // Act
        got := Divide(10, 2)

        // Assert
        want := 5
        if got != want {
            t.Errorf("func(10, 2) = %d, want %d", got, want) // リネームが大変なのでDivideなどの名前をリテラル型の中で使わない
        }
    })
}
```

## 別の仕様には別のテストを書く

```go
// 悪い例
// テストが複数の仕様を含んでおり、何を検証しているのか不明瞭
func TestUserRegistration(t *testing.T) {
    user := RegisterUser("test@example.com", "password123")
    if user == nil {
        t.Fatal("ユーザーが作成されなかった")
    }
    if user.Email != "test@example.com" {
        t.Errorf("メールアドレスが正しくない")
    }
    if !user.IsActive {
        t.Errorf("ユーザーがアクティブでない")
    }
}

// 良い例
// 仕様が明確に伝わるテスト
func TestUserRegistration(t *testing.T) {
    t.Parallel()

    t.Run("有効なメールアドレスとパスワードでユーザーを作成できる", func(t *testing.T) {
        t.Parallel()

        // Act
        user := RegisterUser("test@example.com", "password123")

        // Assert
        if user == nil {
            t.Fatal("ユーザーが作成されなかった")
        }
    })

    t.Run("作成されたユーザーのメールアドレスが正しく設定される", func(t *testing.T) {
        t.Parallel()

        // Act
        user := RegisterUser("test@example.com", "password123")

        // Assert
        if user.Email != "test@example.com" {
            t.Errorf("メールアドレスが期待値と異なる: got %s, want test@example.com", user.Email)
        }
    })

    t.Run("新規ユーザーはデフォルトでアクティブ状態になる", func(t *testing.T) {
        t.Parallel()

        // Act
        user := RegisterUser("test@example.com", "password123")

        // Assert
        if !user.IsActive {
            t.Error("新規ユーザーがアクティブ状態でない")
        }
    })
}
```

## エラーメッセージは詳細に

```go
// 悪い例
func TestParseConfig(t *testing.T) {
    config, err := ParseConfig("config.json")
    if err != nil {
        t.Error("エラーが発生した")  // 何のエラーか不明
    }
    if config.Port != 8080 {
        t.Error("ポートが違う")  // 期待値と実際の値が不明
    }
}

// 良い例
func TestParseConfig(t *testing.T) {
    t.Parallel()

    // Act
    got, err := ParseConfig("config.json")

    // Assert
    if err != nil {
        t.Errorf("設定ファイルの解析に失敗した: %v", err) // リネームが大変なのでParseConfigなどの名前をリテラル型の中で使わない
    }

    want := 8080
    if got.Port != want {
        t.Errorf("ポート番号 %d を得るはずが %d を得た", want, got.Port)
    }
}
```

## 「正常系」「異常系」といった分類は避ける

```go
// 悪い例
func TestUserService_NormalCase(t *testing.T) {
    // 「正常系」が何を指すのか不明確
}

func TestUserService_ErrorCase(t *testing.T) {
    // 「異常系」が何を指すのか不明確
}

// 良い例
func TestUserService(t *testing.T) {
    t.Parallel()

    t.Run("存在するユーザーIDで検索すると該当ユーザーを返す", func(t *testing.T) {
        t.Parallel()

        // Act
        user, err := GetUser(123)

        // Assert
        if err != nil {
            t.Fatalf("予期しないエラー: %v", err)
        }
        if user.ID != 123 {
            t.Errorf("ユーザーIDが期待値と異なる: got %d, want %d", user.ID, 123)
        }
    })

    t.Run("存在しないユーザーIDで検索するとNotFoundエラーを返す", func(t *testing.T) {
        t.Parallel()

        // Act
        _, err := GetUser(999)

        // Assert
        if err != ErrUserNotFound {
            t.Errorf("期待したエラーが返されなかった: got %v, want %v", err, ErrUserNotFound)
        }
    })

    t.Run("負のユーザーIDで検索するとInvalidIDエラーを返す", func(t *testing.T) {
        t.Parallel()

        // Act
        _, err := GetUser(-1)

        // Assert
        if err != ErrInvalidID {
            t.Errorf("期待したエラーが返されなかった: got %v, want %v", err, ErrInvalidID)
        }
    })
}
```


## テーブル駆動テストの活用

テーブルにはmapを使用し、テストの実行順序がランダムになるようにする。

```go
func TestCalculateTax(t *testing.T) {
    t.Parallel()

    tests := map[string]struct {
        price    int
        taxRate  float64
        expected int
    }{
        "消費税10%の場合": {
            price:    1000,
            taxRate:  0.10,
            expected: 1100,
        },
        "消費税8%の場合": {
            price:    1000,
            taxRate:  0.08,
            expected: 1080,
        },
        "0円の商品には税金がかからない": {
            price:    0,
            taxRate:  0.10,
            expected: 0,
        },
    }
    for name, tt := range tests {
        t.Run(name, func(t *testing.T) {
            t.Parallel()

            // Act
            got := CalculateTax(tt.price, tt.taxRate)

            // Assert
            if got != tt.expected {
                t.Errorf("func(%d, %.2f) = %d, want %d", // リネームが大変なのでCalculateTaxなどの名前をリテラル型の中で使わない
                    tt.price, tt.taxRate, got, tt.expected)
            }
        })
    }
}
```

## エラーケースでも全ての戻り値を検証する

エラーが期待される場合でも、関数が予期せず成功した際のデバッグのため、全ての戻り値を検証する。

```go
// 良い例
if tt.wantErr && gotErr == nil {
    t.Errorf("エラーが期待されたが成功した") // t.Fatal()ではなくt.Errorf()で継続
}
// エラーケースでも戻り値を検証（デバッグ情報として有用）
if got != tt.want {
    t.Errorf("got %v, want %v", got, tt.want)
}
```

**理由**: 予期しない成功時に「なぜ成功したか」の手がかりが得られる。
