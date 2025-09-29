# Project Overview: pj-tmpl-go

## Purpose
Go言語プロジェクトのテンプレート。新しいGoプロジェクトを作成するための基盤として使用される。

## Tech Stack
- **言語**: Go
- **ツール管理**: mise (tool version manager)
- **エラーハンドリング**: github.com/go-errors/errors (スタックトレース付き)
- **テスト**: 標準ライブラリ + github.com/google/go-cmp/cmp (比較)
- **ゴールデンファイルテスト**: github.com/tenntenn/golden
- **フェイクデータ**: github.com/brianvoe/gofakeit/v7
- **コード生成**: github.com/anoriqq/enumer (enum生成)

## Architecture
- HTTPサーバーサンプル（graceful shutdown付き）
- CLIラッパー
- ドメイン駆動設計風の構造
- 環境設定とポート管理のドメインオブジェクト

## Entry Points
- `cmd/main.go` - CLI application entry point
- `run.go` - Main application logic

## Tech Stack Details
- Build: CGO_ENABLED=0 for static binaries
- Hot reload: Air (.air.toml)
- Git hooks: lefthook
- Dependency management: Renovate