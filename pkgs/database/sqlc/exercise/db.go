// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package exercise

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.countDrinkHistoryByUserIdStmt, err = db.PrepareContext(ctx, countDrinkHistoryByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query CountDrinkHistoryByUserId: %w", err)
	}
	if q.countExerciseHistoryByUserIdStmt, err = db.PrepareContext(ctx, countExerciseHistoryByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query CountExerciseHistoryByUserId: %w", err)
	}
	if q.createStmt, err = db.PrepareContext(ctx, create); err != nil {
		return nil, fmt.Errorf("error preparing query Create: %w", err)
	}
	if q.createExerciseStmt, err = db.PrepareContext(ctx, createExercise); err != nil {
		return nil, fmt.Errorf("error preparing query CreateExercise: %w", err)
	}
	if q.createOrUpdateWaterStmt, err = db.PrepareContext(ctx, createOrUpdateWater); err != nil {
		return nil, fmt.Errorf("error preparing query CreateOrUpdateWater: %w", err)
	}
	if q.deleteExerciseStmt, err = db.PrepareContext(ctx, deleteExercise); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteExercise: %w", err)
	}
	if q.deleteHealthStmt, err = db.PrepareContext(ctx, deleteHealth); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteHealth: %w", err)
	}
	if q.fetchByDateTimeStmt, err = db.PrepareContext(ctx, fetchByDateTime); err != nil {
		return nil, fmt.Errorf("error preparing query FetchByDateTime: %w", err)
	}
	if q.fetchCategoriesStmt, err = db.PrepareContext(ctx, fetchCategories); err != nil {
		return nil, fmt.Errorf("error preparing query FetchCategories: %w", err)
	}
	if q.fetchExerciseByCategoryIdStmt, err = db.PrepareContext(ctx, fetchExerciseByCategoryId); err != nil {
		return nil, fmt.Errorf("error preparing query FetchExerciseByCategoryId: %w", err)
	}
	if q.fetchTodayExerciseByUserIdStmt, err = db.PrepareContext(ctx, fetchTodayExerciseByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query FetchTodayExerciseByUserId: %w", err)
	}
	if q.getExerciseByIdStmt, err = db.PrepareContext(ctx, getExerciseById); err != nil {
		return nil, fmt.Errorf("error preparing query GetExerciseById: %w", err)
	}
	if q.getTodayExerciseCountStmt, err = db.PrepareContext(ctx, getTodayExerciseCount); err != nil {
		return nil, fmt.Errorf("error preparing query GetTodayExerciseCount: %w", err)
	}
	if q.getWaterByUserIdStmt, err = db.PrepareContext(ctx, getWaterByUserId); err != nil {
		return nil, fmt.Errorf("error preparing query GetWaterByUserId: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.countDrinkHistoryByUserIdStmt != nil {
		if cerr := q.countDrinkHistoryByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countDrinkHistoryByUserIdStmt: %w", cerr)
		}
	}
	if q.countExerciseHistoryByUserIdStmt != nil {
		if cerr := q.countExerciseHistoryByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countExerciseHistoryByUserIdStmt: %w", cerr)
		}
	}
	if q.createStmt != nil {
		if cerr := q.createStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createStmt: %w", cerr)
		}
	}
	if q.createExerciseStmt != nil {
		if cerr := q.createExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createExerciseStmt: %w", cerr)
		}
	}
	if q.createOrUpdateWaterStmt != nil {
		if cerr := q.createOrUpdateWaterStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createOrUpdateWaterStmt: %w", cerr)
		}
	}
	if q.deleteExerciseStmt != nil {
		if cerr := q.deleteExerciseStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteExerciseStmt: %w", cerr)
		}
	}
	if q.deleteHealthStmt != nil {
		if cerr := q.deleteHealthStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteHealthStmt: %w", cerr)
		}
	}
	if q.fetchByDateTimeStmt != nil {
		if cerr := q.fetchByDateTimeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing fetchByDateTimeStmt: %w", cerr)
		}
	}
	if q.fetchCategoriesStmt != nil {
		if cerr := q.fetchCategoriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing fetchCategoriesStmt: %w", cerr)
		}
	}
	if q.fetchExerciseByCategoryIdStmt != nil {
		if cerr := q.fetchExerciseByCategoryIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing fetchExerciseByCategoryIdStmt: %w", cerr)
		}
	}
	if q.fetchTodayExerciseByUserIdStmt != nil {
		if cerr := q.fetchTodayExerciseByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing fetchTodayExerciseByUserIdStmt: %w", cerr)
		}
	}
	if q.getExerciseByIdStmt != nil {
		if cerr := q.getExerciseByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getExerciseByIdStmt: %w", cerr)
		}
	}
	if q.getTodayExerciseCountStmt != nil {
		if cerr := q.getTodayExerciseCountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTodayExerciseCountStmt: %w", cerr)
		}
	}
	if q.getWaterByUserIdStmt != nil {
		if cerr := q.getWaterByUserIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getWaterByUserIdStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                               DBTX
	tx                               *sql.Tx
	countDrinkHistoryByUserIdStmt    *sql.Stmt
	countExerciseHistoryByUserIdStmt *sql.Stmt
	createStmt                       *sql.Stmt
	createExerciseStmt               *sql.Stmt
	createOrUpdateWaterStmt          *sql.Stmt
	deleteExerciseStmt               *sql.Stmt
	deleteHealthStmt                 *sql.Stmt
	fetchByDateTimeStmt              *sql.Stmt
	fetchCategoriesStmt              *sql.Stmt
	fetchExerciseByCategoryIdStmt    *sql.Stmt
	fetchTodayExerciseByUserIdStmt   *sql.Stmt
	getExerciseByIdStmt              *sql.Stmt
	getTodayExerciseCountStmt        *sql.Stmt
	getWaterByUserIdStmt             *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                               tx,
		tx:                               tx,
		countDrinkHistoryByUserIdStmt:    q.countDrinkHistoryByUserIdStmt,
		countExerciseHistoryByUserIdStmt: q.countExerciseHistoryByUserIdStmt,
		createStmt:                       q.createStmt,
		createExerciseStmt:               q.createExerciseStmt,
		createOrUpdateWaterStmt:          q.createOrUpdateWaterStmt,
		deleteExerciseStmt:               q.deleteExerciseStmt,
		deleteHealthStmt:                 q.deleteHealthStmt,
		fetchByDateTimeStmt:              q.fetchByDateTimeStmt,
		fetchCategoriesStmt:              q.fetchCategoriesStmt,
		fetchExerciseByCategoryIdStmt:    q.fetchExerciseByCategoryIdStmt,
		fetchTodayExerciseByUserIdStmt:   q.fetchTodayExerciseByUserIdStmt,
		getExerciseByIdStmt:              q.getExerciseByIdStmt,
		getTodayExerciseCountStmt:        q.getTodayExerciseCountStmt,
		getWaterByUserIdStmt:             q.getWaterByUserIdStmt,
	}
}
