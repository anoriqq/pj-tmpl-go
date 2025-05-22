# anoriqq/pj-tmpl-go

```sh
git clone https://github.com/anoriqq/pj-tmpl-go <path/to/pj>
cd <path/to/pj>
go mod edit -module <new/module/name>
go mod edit -go=$(go version | awk '{print $3}' | sed 's/go//')
```
