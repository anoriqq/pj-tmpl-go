# Code Style and Conventions

## Testing Conventions
1. **テスト名**: `Test_FunctionName_Condition_ExpectedResult` 形式
2. **テーブル駆動テスト**: mapを使用してランダム実行順序
3. **並行実行**: `t.Parallel()` を使用
4. **詳細なエラーメッセージ**: 期待値と実際の値を明示
5. **別仕様は別テスト**: 仕様ごとにサブテストを分離
6. **ゴールデンファイル**: `testdata/` ディレクトリに格納

## Error Handling
- `github.com/go-errors/errors` でスタックトレース付きエラー
- `errors.Wrap(err, 0)` でエラーラップ

## Code Generation
- `github.com/anoriqq/enumer` でenum生成
- `go generate ./...` で実行

## Project Structure Pattern
- `internal/domain/` - ドメインロジック
- `internal/infra/` - インフラストラクチャ
- `cmd/` - エントリーポイント

## Build Configuration
- Static binaries: `CGO_ENABLED=0`
- Release builds: `netgo` tag, static linking
- Development: race detector enabled

## Documentation
- README.md は日本語
- コメントは適切に日本語/英語を使い分け