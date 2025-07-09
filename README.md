# anoriqq/pj-tmpl-go

Go言語プロジェクトのテンプレート。

## Usage

```sh
PKG=github.com/yourname/yourpj
ghq create ${PKG}
gonew github.com/anoriqq/pj-tmpl-go $(ghq list -e ${PKG}) $(ghq list -p -e ${PKG})/tmp
mv $(ghq list -p -e ${PKG})/tmp/* $(ghq list -p -e ${PKG})
rm -r $(ghq list -p -e ${PKG})/tmp
```

## Prerequisites

- [mise](https://mise.jdx.dev/)
