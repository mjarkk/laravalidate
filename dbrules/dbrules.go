package dbrules

import (
	"database/sql"
	"fmt"

	. "github.com/mjarkk/laravalidate"
)

type DB struct{ conn *sql.DB }

func AddRules(conn *sql.DB) {
	if conn == nil {
		panic("DB connection cannot be nil")
	}

	db := &DB{conn}

	RegisterValidator("exists", db.Exists)

	BaseRegisterMessages(map[string]MessageResolver{
		"exists": BasicMessageResolver("The selected :attribute is invalid."),
	})

	LogValidatorsWithoutMessages()
}

func (b *DB) Exists(ctx *ValidatorCtx) (string, bool) {
	if len(ctx.Args) == 0 {
		return "args", false
	}

	tableName := ctx.Args[0]
	if tableName == "" {
		return "args", false
	}

	column := "id"
	if len(ctx.Args) >= 2 && ctx.Args[1] != "" {
		column = ctx.Args[1]
	}

	ctx.UnwrapPointer()

	if !ctx.HasValue() {
		return "", true
	}

	row := b.conn.QueryRow(
		fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? LIMIT 1", column, tableName, column),
		ctx.Value.Interface(),
	)
	if row.Err() != nil {
		return "exists", false
	}

	resp := sql.RawBytes{}
	err := row.Scan(&resp)
	if err != nil {
		return "exists", false
	}

	return "", true
}
