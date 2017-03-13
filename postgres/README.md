# Postgres Package

Usage/Code:

```go
package main

import (
	"fmt"

	"github.build.ge.com/aviation-intelligent-network/px-go-micro-svc-tmp/rest"
	pg "github.build.ge.com/aviation-intelligent-network/px-go-micro-svc-tmp/postgres"

)

func main() {
	fmt.Println("-------------------------------------------")
	fmt.Println("Starting px-application-name ...")
	fmt.Println("-------------------------------------------")

	//init Postgres
	pg.Open_postgres("user=postgres password=predix dbname=postgres sslmode=disable", "schema")
	defer pg.Database.Close()

	//  start rest interface here
	rest.StartServer()
}
```
