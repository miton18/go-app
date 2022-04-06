# MongoDB client setup

```sh
go get go.clever-cloud.com/app/mongodb
```

```go
package main

import(
    "go.clever-cloud.com/app/mongodb"
)

func main() {
    client, err := mongodb.New()
    if err != nil {
        panic(err)
    }

    res, err := client.Collection("qux").InsertOne(context.Background(), bson.M{"hello": "world"})
    if err != nil {
        panic(err)
    }
}

```