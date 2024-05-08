# go-metrics

Install:
```shell
go get github.com/artarts36/go-metrics
```

## Usage

```go
package main

import (
	"github.com/artarts36/go-metrics"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	cfg := &metrics.Config{
		Server: metrics.ServerConfig{
			Addr:    ":8000",
			Timeout: 30 * time.Second,
		},
		Namespace: "service_name",
	}

	registry := metrics.NewDefaultRegistry(cfg)
	registry.
		NewCounter(prometheus.CounterOpts{
		    Name: "clicks",
        }).
		Inc()
	
	server := metrics.NewServer(cfg)
	
	server.Serve()
}
```
