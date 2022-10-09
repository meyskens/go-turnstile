# Cloudflare Turnstile Go library

This library is a [Turnstile](https://www.cloudflare.com/products/turnstile//) server side library for Go.
It allows to verify challange responses sent to the server.

## Usage

The library gives a struct set up to verify the challenges of a given private key set up with `New()`, responses received by the server can be verified with `Verify()`.

```go
import "github.com/meyskens/go-turnstile"

func handleRequest(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    tsResponse, _ := r.Form["cf-turnstile-response"]
    ts := turnstile.New("<insert secret key>")
    //Get IP from RemoteAddr
    ip, _, err := net.SplitHostPort(r.RemoteAddr)

    resp, err := ts.Verify(tsResponse[0], ip)
    // handle errors please!
    if resp.Success {
        // captcha OK!
    }
}
```
