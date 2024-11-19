package dbrules

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	. "github.com/mjarkk/laravalidate"
)

type DB struct {
	conn          *sql.DB
	variableStyle QueryVariableStyle
}

type QueryVariableStyle uint8

const (
	DefaultStyle QueryVariableStyle = iota // ?
	PgStyle                                // $1, $2, ...
)

func AddRules(conn *sql.DB, variableStyle QueryVariableStyle) {
	if conn == nil {
		panic("DB connection cannot be nil")
	}

	db := &DB{conn, variableStyle}

	RegisterValidator("exists", db.Exists)

	BaseRegisterMessages(map[string]MessageResolver{
		"exists": BasicMessageResolver("The selected :attribute is invalid."),
	})

	LogValidatorsWithoutMessages()
}

func (b *DB) prepareQuery(in string) string {
	switch b.variableStyle {
	case PgStyle:
		for i := 1; i <= strings.Count(in, "?"); i++ {
			iStr := "$" + strconv.Itoa(i)
			in = strings.Replace(in, "?", iStr, 1)
		}
	}
	return in
}

func (b *DB) query(query string, args ...any) (*sql.Rows, error) {
	return b.conn.Query(b.prepareQuery(query), args...)
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

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ? LIMIT 1", column, tableName, column)
	result, err := b.query(
		query,
		ctx.Value.Interface(),
	)
	if err != nil {
		return "exists", false
	}

	defer result.Close()
	for result.Next() {
		resp := sql.RawBytes{}
		err := result.Scan(&resp)
		if err != nil {
			return "exists", false
		}

		return "", true
	}

	return "exists", false
}
