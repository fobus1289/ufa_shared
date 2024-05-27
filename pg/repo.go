package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type Query interface {
	/*
		type User struct {
			Login    string
			Password string
			Id       int64
		}

		func (u *userService) Create(user User) (int64, error) {
			var id int64

			err := u.Get(&id, "INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id",
				user.Login,
				user.Password,
			)
			if err != nil {
				return 0, err
			}

			return id, nil
		}

		func (u *userService) FindById(id int64) (*User, error) {
			var user User

			err := u.Get(&user, "SELECT * FROM users WHERE id = $1", id)
			if err != nil {
				return nil, err
			}

			return &user, nil
		}
	*/
	Get(dest any, query string, args ...any) error

	/*```go
	ctx = context.Background()

	jason = Person{}

	err = db.GetContext(ctx, &jason, "SELECT * FROM person WHERE first_name=$1", "Jason")
	*/
	GetContext(ctx context.Context, dest any, query string, args ...any) error

	/* ```go
	 sqlResult, err := db.Exec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	``` */
	Exec(query string, args ...any) (sql.Result, error)

	/* ```
	ctx := context.Background()

	sqlResult, err := db.ExecContext(ctx, "INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	``` */
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)

	/* ```go
	db.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})

	_, err = db.NamedExec(`INSERT INTO person (first_name,last_name,email) VALUES (:first,:last,:email)`,
	    map[string]interface{}{
	        "first": "Bin",
	        "last": "Smuth",
	        "email": "bensmith@allblacks.nz",
	})

	personStructs := []Person{
	    {FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
	    {FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
	    {FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	}

	_, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)`, personStructs)
	```*/
	NamedExec(query string, arg any) (sql.Result, error)

	/* ```go
	 p := Person{first_name:"jhone"}

	 rows, err := db.NamedQuery(`SELECT * FROM person WHERE first_name=:first_name`, &p)

	 rows, err := db.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`, map[string]interface{}{"fn": "Bin"})
	``` */
	NamedQuery(query string, arg any) (*sqlx.Rows, error)

	/* ```go
	 p := &Person{"Jane", "Citizen", "jane.citzen@example.com"}

	 db.NamedExec(
	   "INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)",
	   p,
	 )

	 personStructs := []Person{
	 	{FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
	 	{FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
	 	{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	 }

	 _, err := db.NamedExec(
	 	`INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)`,
	 	personStructs,
	 )
	``` */
	PrepareNamed(query string) (*sqlx.NamedStmt, error)

	/* ```go
	ctx := context.Background()

	p := &Person{"Jane", "Citizen", "jane.citzen@example.com"}

	db.NamedExec(
	  "INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)",
	  p,
	)

	personStructs := []Person{
		{FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
		{FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
		{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	}

	_, err := db.NamedExec(
		`INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)`,
		personStructs,
	)
	``` */
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)

	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)

	QueryRowx(query string, args ...any) *sqlx.Row
	QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row

	/* ```go

	rows, err := db.Queryx("SELECT * FROM place")

	log.Fatalln(err)

	for rows.Next() {
		var place Place

	  err := rows.StructScan(&place)

	  if err != nil {
	    log.Fatalln(err)
	  }

	  fmt.Printf("%#v\n", place)
	}

	``` */
	Queryx(query string, args ...any) (*sqlx.Rows, error)

	/* ```go

	ctx := context.Background()

	rows, err := db.QueryxContext(ctx, "SELECT * FROM place")

	log.Fatalln(err)

	for rows.Next() {
		var place Place

	  err := rows.StructScan(&place)

	  if err != nil {
	    log.Fatalln(err)
	  }

	  fmt.Printf("%#v\n", place)
	}

	``` */
	QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)

	/* ```go
	people := []Person{}

	err := db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	```*/
	Select(dest any, query string, args ...any) error

	/* ```go
	ctx := context.Background()

	people := []Person{}

	err := db.SelectContext(ctx ,&people, "SELECT * FROM person ORDER BY first_name ASC")
	```*/
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
}

type Transaction interface {
	Rollback() error
	Commit() error

	TTx() (RepositoryTx, error)
	WithTx(rtx RepositoryTx) (RepositoryTx, error)
	Txx(ctx context.Context, opts *sql.TxOptions) (RepositoryTx, error)
	WithTxx(ctx context.Context, rtx RepositoryTx, opts *sql.TxOptions) (RepositoryTx, error)
}

type Repository interface {
	Transaction
	Query
}

type RepositoryTx interface {
	Transaction
	Query
}

type repo struct {
	*sqlx.DB
}

type repoTx struct {
	*sqlx.Tx
}

func (r *repoTx) TTx() (RepositoryTx, error) {
	return nil, errors.New("not implemented")
}

func (r *repoTx) WithTx(rtx RepositoryTx) (RepositoryTx, error) {
	return nil, errors.New("not implemented")
}

func (r *repoTx) Txx(ctx context.Context, opts *sql.TxOptions) (RepositoryTx, error) {
	return nil, errors.New("not implemented")
}

func (r *repoTx) WithTxx(ctx context.Context, rtx RepositoryTx, opts *sql.TxOptions) (RepositoryTx, error) {
	return nil, errors.New("not implemented")
}

func (r *repo) TTx() (RepositoryTx, error) {
	tx, err := r.Beginx()
	if err != nil {
		return nil, err
	}

	rTx := &repoTx{tx}

	return rTx, nil
}

func (r *repo) Txx(ctx context.Context, opts *sql.TxOptions) (RepositoryTx, error) {
	tx, err := r.BeginTxx(ctx, opts)
	if err != nil {
		return nil, err
	}

	rTx := &repoTx{tx}

	return rTx, nil
}

func (r *repo) WithTx(rtx RepositoryTx) (RepositoryTx, error) {

	rr, ok := rtx.(*repoTx)

	if !ok {
		return nil, errors.New("wrong type")
	}

	rTx := &repoTx{rr.Tx}

	return rTx, nil
}

func (r *repo) WithTxx(ctx context.Context, rtx RepositoryTx, opts *sql.TxOptions) (RepositoryTx, error) {
	rr, ok := rtx.(*repoTx)

	if !ok {
		return nil, errors.New("wrong type")
	}

	rTx := &repoTx{rr.Tx}

	return rTx, nil
}

func (r *repo) DBX() *sqlx.DB {
	return r.DB
}

func (r *repo) Rollback() error {
	return errors.New("not implemented")
}

func (r *repo) Commit() error {
	return errors.New("not implemented")
}

func New(cfg Config) (Repository, error) {
	instance, err := RetryConnect(cfg, 20, 10)
	if err != nil {
		return nil, err
	}

	r := &repo{DB: instance.DB}

	return r, nil
}

type TransactionFn[T any] func(RepositoryTx) (T, error)

func Transact[T any](db Repository, transactionFn TransactionFn[T]) (result T, err error) {
	tx, err := db.TTx()

	if err != nil {
		return
	}

	defer func(err error) {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}

		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}(err)

	result, err = transactionFn(tx)

	if err != nil {
		return
	}

	return result, nil
}
