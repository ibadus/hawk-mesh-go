# hawk mesh go 🌟

```
status: working ✅
```

## 💬 Description 
 Wrapper to easily generate "X-Request-Auth" header for Mesh sites in golang. Based on [hawk-go](https://github.com/tent/hawk-go).

## Getting Started
### Import the module :
```go
import (
    "github.com/ibadus/hawk-mesh-go"
)
```
### 🤔 Usage :
```go
headers, err := hawk.GenerateHeaders("GET", url, key, secret)
headers, err := hawk.GenerateHeaders("PUT", url, key, secret)
headers, err := hawk.GenerateHeadersWithPayload("POST", url, key, secret, string(payload))
headers, err := hawk.GenerateHeadersWithPayload("PUT", url, key, secret, string(payload))
```
You can find ``key``, ``secret``, ``apikey`` in the **Mesh-Keys.txt** [click here](https://github.com/VastidDev/Mesh-Keys). Python solution also available in the repo linked.
### 🔔 Examples :
#### Create a GET "X-Request-Auth" header
```go
key := "f2188a5b06"
secret := "8bb6bd51c83f2ec9821e1bda5c77b25b"
url := "https://prod.jdgroupmesh.cloud/stores/jdsportsfr/" + "..."

headers, err := hawk.GenerateHeaders("GET", url, key, secret)
if err != nil {
	return
}
hawkHeader := map[string]string{"X-Request-Auth": headers}
```
#### Create a POST "X-Request-Auth" header
```go
key := "f2188a5b06"
secret := "8bb6bd51c83f2ec9821e1bda5c77b25b"
url := "https://prod.jdgroupmesh.cloud/stores/jdsportsfr/" + "..."
payload := []byte(fmt.Sprintf(`
    "Value1": "%s",
    "Value2": "%s",
`, "v1", "v2"))

headers, err := hawk.GenerateHeadersWithPayload("POST", url, key, secret, string(payload))
if err != nil {
	return
}
hawkHeader := map[string]string{"X-Request-Auth": headers}
```
#### Create a PUT "X-Request-Auth" header
```go
key := "f2188a5b06"
secret := "8bb6bd51c83f2ec9821e1bda5c77b25b"
url := "https://prod.jdgroupmesh.cloud/stores/jdsportsfr/" + "..."
payload := []byte(fmt.Sprintf(`
    "Value1": "%s",
    "Value2": "%s",
`, "v1", "v2"))

headers, err := hawk.GenerateHeadersWithPayload("PUT", url, key, secret, string(payload))
if err != nil {
	return
}
hawkHeader := map[string]string{"X-Request-Auth": headers}
```
#### Let's test
```go
package main

import (
	"fmt"
	"github.com/ibadus/hawk-mesh-go"
)

func main() {
	fmt.Println("Let's try")
	key := "f2188a5b06"
	secret := "8bb6bd51c83f2ec9821e1bda5c77b25b"
	url := "https://prod.jdgroupmesh.cloud/stores/jdsportsfr/"

	headers, err := hawk.GenerateHeaders("GET", url, key, secret)
	if err != nil {
		return
	}
	fmt.Println("X-Request-Auth:", headers)
}
```
