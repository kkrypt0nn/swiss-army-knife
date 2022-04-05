# Certificate

This package lets you create a valid x509 certificate for a local web server so that you can access it from another computer using the same certificate imported in the browser.

## Examples

There are some settings we need to provide, in the same order, to be able to generate the certificate:
- **serial**: The serial number of the certificate.
- **country**: The country the certificate is issued from.
- **organization**: The organization that is issuing the certificate.
- **commonName**: The common name of the certificate.
- **expire**: When the certificate is due to expire.
- **localIP**: The local IP of the future web server, pass `nil` if it should be automatically gathered from `GetLocalIP()`.

```go
package main

import (
    "fmt"
    "github.com/kkrypt0nn/swiss-army-knife/certificate"
    "time"
)

func main() {
	privateKey, certificate := certificate.GenerateCertificate(1337, "CH", "Krypton", "My Certificate", time.Now().AddDate(7, 7, 7), nil)
	fmt.Printf("%s\n", privateKey)
	fmt.Printf("%s\n", certificate)
}
```
