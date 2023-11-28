// Package app configures and runs application.
package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/good1hare/GolangTemplate/configs"
	v1 "github.com/good1hare/GolangTemplate/internal/controller/api/v1"
	"github.com/good1hare/GolangTemplate/internal/usecase"
	"github.com/good1hare/GolangTemplate/internal/usecase/repo"
	"github.com/good1hare/GolangTemplate/pkg/logger"
	"github.com/labstack/echo/v4"
	"os"
	"path/filepath"
)

func Run(cfg *configs.Config) {
	log := logger.New(cfg.Logger.Level)

	//Migrations and connect db
	projectDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	migrationsDir := filepath.Join(projectDir, "db/migrations")

	db, err := sql.Open(cfg.Mysql.DriverName, fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.Mysql.UserName, cfg.Mysql.Password, cfg.Mysql.Host+":"+cfg.Mysql.Port, cfg.Mysql.DatabaseName))
	if err != nil {
		panic(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsDir,
		cfg.Mysql.DriverName,
		driver,
	)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Println("No changes needed. Continuing program execution...")
		} else {
			panic(err)
		}
	}

	//Use case
	userUseCase := usecase.NewUserUseCase(
		repo.NewUserRepo(db),
	)

	// create a new echo instance
	e := echo.New()
	v1.NewRouter(e, log, userUseCase)

	//Start server
	echoErr := e.Start(":8080")
	if err != nil {
		panic(echoErr)
	}
}
