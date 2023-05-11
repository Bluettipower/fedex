# Fedex Api Client

# Useage

## Install
```console
$ go get github.com/bluettipower/fedex
```
## Use
```go
import "github.com/bluettipower/fedex"
```
## Crate Shipment 
```go
func main() {
  // new client
  client := fedex.NewFedExClient("your clientID","your clientSecret")

  // new Shipment 
  data := fedex.Shipment{
    // Your data ...
  }
  // Create Shipment
  shipment,err := client.Create(data)
  // ...
}
```

## With Sandbox
```go
  // ...
	client := NewFedExClient("your clientID", "your clientSecret", WithSandbox(true))
  // ...
```