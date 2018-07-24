# grafeas-cli

Initial experiment -- not sure if it will be useful yet, just using this as a place to save my code across computers for now.

Temporary workaround for `vendor/` type issues:

```bash
go get ./...
# The `vendor` directory causes type errors
mv $GOPATH/src/github.com/grafeas/grafeas/vendor $GOPATH/src/github.com/grafeas/grafeas/vendor.bak
```