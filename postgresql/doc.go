// Package postgresql configure PostgreSQL connection
// Use it on production:
//  package main
//
//  import "go.clever-cloud.dev/app/postgresql"
//
//  func main() {
//    db, err := postgresql.New(context.Background())
//    if err != nil {
//      panic(err)
//    }
//    // use `db` ....
//  }
//
// On your desk for tests (require Docker)
//  package main
//
//  import "go.clever-cloud.dev/app/postgresql"
//
//  func main() {
//    // Star a local PostgreSQL
//    err := SetupLocalDB(context.Background())
//    if err != nil {
//      t.Fatalf(err.Error())
//    }
//
//    db, err := postgresql.New(context.Background())
//    if err != nil {
//      panic(err)
//    }
//    // use `db` ....
//  }
package postgresql
