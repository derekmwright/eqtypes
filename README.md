# EQTypes

This is currently a work in progress and does not have complete coverage of the Project EQ database.
It is being written mainly for item, character, quest, and tradeskill parsing to provide visibility tools for server admins and their communities.

If you would like to request a specific data type be created, please open an issue or submit a PR.

## Installation

    go get github.com/derekmwright/eqtypes

## Usage

This project contains a mix of objects taken from the database and custom types that implement special methods to assist in parsing/marshaling data.
Types such as `Item` contain `db:` tags to ensure they can be scanned properly by most DB tools.
I personally use `sql/database` and `sqlx` for querying and scanning from the database, so you will have the best luck using those tools.

Example Usage:

```go
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/derekmwright/eqtypes"
	"github.com/jmoiron/sqlx"
)

const driverName = "mysql"

func main() {
	db, err := sqlx.Connect(driverName, os.Getenv("DATABASE_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}

	// Use unsafe to skip scanning unknown or unused fields
	db = db.Unsafe()

	// Queries an item with ID from the DB and scans it into `item`
	query := `SELECT * FROM items WHERE id = ? LIMIT 1`
	item := eqtypes.Item{}
	if err = db.QueryRowxContext(context.Background(), query, 1001).StructScan(&item); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("item not found")
		}
	}
	
	fmt.Printf("%+v\n", item)
}
```
