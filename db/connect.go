package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}
type pgxDb struct {
	*pgx.Conn
}
type gormDb struct {
	*gorm.DB
}
type Store struct {
	//	db pgxDb
	db gormDb
}

func Connect(ctx context.Context, cfg *Config) (*Store, error) {

	//db, err := newPgxDB(ctx, cfg)
	db, err := newGormDB(ctx, cfg)
	if err != nil {
		return nil, err
	}
	s := &Store{db: *db}
	return s, err
}

func newPgxDB(ctx context.Context, cfg *Config) (*pgxDb, error) {

	u := fmt.Sprintf("postgres://%s:%s@%s:%v/%s",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name)

	p, err := pgx.Connect(context.Background(), u)
	if err != nil {
		return nil, err
	}

	defer p.Close(ctx)
	pgxDb := &pgxDb{Conn: p}
	return pgxDb, nil

}

func newGormDB(ctx context.Context, cfg *Config) (*gormDb, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.Name)
	// Open the GORM database connection
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Retrieve the underlying SQL database connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Ping the database to check connectivity
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	gormDb := &gormDb{DB: db}
	return gormDb, nil
}
