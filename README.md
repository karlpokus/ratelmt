# ratelmt
Rate limiter http middleware

# install
```bash
$ go get github.com/karlpokus/ratelmt
```

# usage
```go
import "github.com/karlpokus/ratelmt"

// the reply http.Handler is rate limited to 1 req/s
http.ListenAndServe(addr, ratelmt.Mw(1, reply))
```

# test
```bash
$ go test -v -race
```

# license
MIT
