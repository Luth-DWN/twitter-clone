package postgres

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/clients/sqlinterface"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	sqlinterface.SQL
}

type client struct {
	ctx  context.Context
	conn *pgxpool.Pool
}

// New creates a new postgresql database instance
func New(ctx context.Context, connString string) (Client, error) {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &client{
		ctx:  ctx,
		conn: conn,
	}, nil
}

func (c *client) Exec(sql string, args ...interface{}) (int64, error) {
	result, err := c.conn.Exec(c.ctx, sql, args...)
	return result.RowsAffected(), err
}

func (c *client) Query(sql string, args ...interface{}) (sqlinterface.SQLRows, error) {
	rows, err := c.conn.Query(c.ctx, sql, args...)
	return newDatabaseRows(rows), err
}

func (c *client) QueryRow(sql string, args ...interface{}) sqlinterface.SQLRow {
	row := c.conn.QueryRow(c.ctx, sql, args...)
	return newDatabaseRow(row)
}

func (c *client) BeginTx() (sqlinterface.SQLTransaction, error) {
	tx, err := c.conn.BeginTx(c.ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}

	return newDatabaseTransaction(c.ctx, tx), nil
}

type clientRow struct {
	row sqlinterface.SQLRow
}

func newDatabaseRow(row sqlinterface.SQLRow) sqlinterface.SQLRow {
	return &clientRow{
		row: row,
	}
}

func (c *clientRow) Scan(dest ...interface{}) error {
	return c.row.Scan(dest...)
}

type clientRows struct {
	rows sqlinterface.SQLRows
}

func newDatabaseRows(rows sqlinterface.SQLRows) sqlinterface.SQLRows {
	return &clientRows{
		rows: rows,
	}
}

func (c *clientRows) Scan(dest ...interface{}) error {
	return c.rows.Scan(dest...)
}

func (c *clientRows) Next() bool {
	return c.rows.Next()
}

func (c *clientRows) Close() {
	c.rows.Close()
}

func (c *clientRows) Err() error {
	return c.rows.Err()
}

type transaction struct {
	ctx context.Context
	tx  pgx.Tx
}

func newDatabaseTransaction(ctx context.Context, tx pgx.Tx) sqlinterface.SQLTransaction {
	return &transaction{
		ctx: ctx,
		tx:  tx,
	}
}

func (c *transaction) Commit() error {
	return c.tx.Commit(c.ctx)
}

func (c *transaction) Rollback() error {
	return c.tx.Rollback(c.ctx)
}

func (c *transaction) Exec(sql string, args ...interface{}) (int64, error) {
	result, err := c.tx.Exec(c.ctx, sql, args...)
	return result.RowsAffected(), err
}

func (c *transaction) Query(sql string, args ...interface{}) (sqlinterface.SQLRows, error) {
	rows, err := c.tx.Query(c.ctx, sql, args...)
	return newDatabaseRows(rows), err
}

func (c *transaction) QueryRow(sql string, args ...interface{}) sqlinterface.SQLRow {
	row := c.tx.QueryRow(c.ctx, sql, args...)
	return newDatabaseRow(row)
}
