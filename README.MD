Example of using go binary-only packages.

Requires go version go1.7beta2 or above.

Assuming you're on "darwin_amd64":

```
go install github.com/elliott5/gboplay
go run $GOPATH/src/github.com/elliott5/gboplay/test/*.go
echo should have failed
mkdir -p $GOPATH/pkg/darwin_amd64/github.com/elliott5/gboplay/test/vendor/github.com/elliott5 
cp $GOPATH/pkg/darwin_amd64/github.com/elliott5/gboplay.a $GOPATH/pkg/darwin_amd64/github.com/elliott5/gboplay/test/vendor/github.com/elliott5 
go run $GOPATH/src/github.com/elliott5/gboplay/test/*.go
echo should have worked
```

Obviously if you run it a second time it will all work, so delete directory $GOPATH/pkg/darwin_amd64/github.com/elliott5/gboplay/test/vendor/github.com/elliott5 to see it in action again ;)
