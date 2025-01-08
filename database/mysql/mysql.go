package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
)

func (m *MySQLRepository) GetDB() *sql.DB {
	return m.DB
}

func (m *MySQLRepository) IsExist(ctx context.Context, query string, args ...interface{}) (bool, error) {
	var exists int
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return exists > 0, nil
}

func (m *MySQLRepository) Insert(ctx context.Context, query string, args ...interface{}) error {
	_, err := m.DB.ExecContext(ctx, query, args...)
	return err
}

func (m *MySQLRepository) Delete(ctx context.Context, query string, args ...interface{}) error {
	_, err := m.DB.ExecContext(ctx, query, args...)
	return err
}

func (m *MySQLRepository) Update(ctx context.Context, query string, args ...interface{}) error {
	_, err := m.DB.ExecContext(ctx, query, args...)
	return err
}

func (m *MySQLRepository) UpdateField(ctx context.Context, table, field, value, where string, args ...interface{}) (int64, error) {
	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s", table, field, where)
	result, err := m.DB.ExecContext(ctx, query, append([]interface{}{value}, args...)...)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	return rowsAffected, err
}

func (m *MySQLRepository) Upsert(ctx context.Context, query string, args ...interface{}) error {
	_, err := m.DB.ExecContext(ctx, query, args...)
	return err
}

func (m *MySQLRepository) GetAllData(ctx context.Context, query string, dest interface{}, args ...interface{}) error {
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))
	for i := range columns {
		valuePointers[i] = &values[i]
	}

	var results []map[string]interface{}
	for rows.Next() {
		err := rows.Scan(valuePointers...)
		if err != nil {
			return err
		}

		row := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			row[col] = val
		}
		results = append(results, row)
	}

	reflectValue := reflect.ValueOf(dest)
	if reflectValue.Kind() != reflect.Ptr || reflectValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("dest must be a pointer to a slice")
	}

	reflectValue.Elem().Set(reflect.ValueOf(results))
	return nil
}

func (m *MySQLRepository) GetData(ctx context.Context, query string, dest interface{}, args ...interface{}) error {
	return m.GetAllData(ctx, query, dest, args...)
}

func (m *MySQLRepository) GetMultiplePartitionData(ctx context.Context, query, orderBy string, offset, limit int, dest interface{}, args ...interface{}) error {
	queryWithPagination := fmt.Sprintf("%s ORDER BY %s LIMIT %d OFFSET %d", query, orderBy, limit, offset)
	return m.GetAllData(ctx, queryWithPagination, dest, args...)
}

func (m *MySQLRepository) CreateIndex(ctx context.Context, table, fieldName string) error {
	indexName := fmt.Sprintf("%s_%s_index", table, fieldName)
	query := fmt.Sprintf("CREATE INDEX %s ON %s (%s)", indexName, table, fieldName)
	_, err := m.DB.ExecContext(ctx, query)
	return err
}

func (m *MySQLRepository) DropIndex(ctx context.Context, table, indexName string) error {
	query := fmt.Sprintf("DROP INDEX %s ON %s", indexName, table)
	_, err := m.DB.ExecContext(ctx, query)
	return err
}

func (m *MySQLRepository) FindOne(ctx context.Context, query string, dest interface{}, args ...interface{}) error {
	row := m.DB.QueryRowContext(ctx, query, args...)
	return row.Scan(dest)
}

func (m *MySQLRepository) UpdateMany(ctx context.Context, query string, args ...interface{}) (int64, error) {
	result, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *MySQLRepository) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return m.DB.QueryContext(ctx, query, args...)
}

func (m *MySQLRepository) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return m.DB.ExecContext(ctx, query, args...)
}

func (m *MySQLRepository) Count(ctx context.Context, query string, args ...interface{}) (int64, error) {
	var count int64
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *MySQLRepository) Transaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := m.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit()
}

func (m *MySQLRepository) Close(ctx context.Context) {
	m.DB.Close()
}
