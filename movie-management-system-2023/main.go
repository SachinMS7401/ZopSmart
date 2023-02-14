package main

import (
	_ "net/http"
	_ "time"

	"developer.zopsmart.com/go/gofr/cmd/gofr/migration"
	dbmigration "developer.zopsmart.com/go/gofr/cmd/gofr/migration/dbMigration"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-training/movie-management-system-2023/middlewares"
	_ "github.com/gorilla/mux"

	http1 "github.com/go-training/movie-management-system-2023/internels/http/movie"
	svcNew "github.com/go-training/movie-management-system-2023/internels/services/movie"
	storeNew "github.com/go-training/movie-management-system-2023/internels/stores/movie"
	"github.com/go-training/movie-management-system-2023/migrations"
)

func main() {
	store := storeNew.New()
	svc := svcNew.New(store)
	handler := http1.New(svc)
	app := gofr.New()

	if app.DB() != nil {
		runMigrations(app)
	} else {
		app.Logger.Warn("DB not configured, skipping migrations!")
	}

	app.Server.ValidateHeaders = false
	app.Server.UseMiddleware(middlewares.Middleware)
	app.REST("movies", handler)
	app.Start()
}

func runMigrations(app *gofr.Gofr) {
	db := dbmigration.NewGorm(app.GORM())

	if err := migration.Migrate("movie-management-system", db, migrations.All(), "UP", app.Logger); err != nil {
		app.Logger.Error("error while migrating")
		return
	}

	app.Logger.Info("Migrated successfully")
}
