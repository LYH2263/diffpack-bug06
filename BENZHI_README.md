# go-diffpack

基于 Go 实现的二进制差分包库组件，配套 diffpackd 管理页，完成块级匹配、差分构建、bundle 校验与作业 hunks 浏览。

## Build / Test

```text
go build ./...
go test ./... -count=1
go run ./cmd/diffpackd -addr :8240
```
