# Prometheus endpoint setup

```sh
go get go.clever-cloud.com/app/prometheus
```

```go
package main

import(
    "go.clever-cloud.com/app/prometheus"
)

func main() {
    prom, err := prometheus.Start()
    if err != nil {
        panic(err)
    }
    defer prom.Close()

    // register your metrics
}
```