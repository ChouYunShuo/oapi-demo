package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"

	"github.com/ChouYunShuo/oapi-demo/idm"
	"github.com/ChouYunShuo/oapi-demo/private_api"
	"github.com/ChouYunShuo/oapi-demo/public_api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tendant/chi-demo/server"
)

type DbConfig struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

func (c DbConfig) toDatabaseUrl() string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.User, c.Password),
		Host:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		Path:   c.Database,
	}
	return u.String()
}

type DbConf interface {
	toDbConfig() DbConfig
}

type DemoDbConfig struct {
	Host     string `env:"ONBOARD_DEMO_PG_HOST" env-default:"localhost"`
	Port     uint16 `env:"ONBOARD_DEMO_PG_PORT" env-default:"5432"`
	Database string `env:"ONBOARD_DEMO_PG_DATABASE" env-default:"onboard_demo_db"`
	User     string `env:"ONBOARD_DEMO_PG_USER" env-default:"demo_usr"`
	Password string `env:"ONBOARD_DEMO_PG_PASSWORD" env-default:"pwd"`
}

func (d DemoDbConfig) toDbConfig() DbConfig {
	return DbConfig{
		Host:     d.Host,
		Port:     d.Port,
		Database: d.Database,
		User:     d.User,
		Password: d.Password,
	}
}

type Config struct {
	ServerConfig server.Config
	DemoDb       DemoDbConfig
}

var tokenAuth *jwtauth.JWTAuth

func main() {
	var cfg Config
	cleanenv.ReadEnv(&cfg)

	connPool, dbQuery := connectDb(cfg)

	defer connPool.Close()

	sh_public := public_api.IdmStore{
		IdmService: &public_api.Service{
			Queries: dbQuery,
		},
	}
	sh_private := private_api.IdmStore{
		IdmService: &private_api.Service{
			Queries: dbQuery,
		},
	}
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	s := server.Default(cfg.ServerConfig)
	server.Routes(s.R)

	publicRouter := chi.NewRouter()
	privateRouter := chi.NewRouter()

	privateRouter.Use(jwtauth.Verifier(tokenAuth))
	privateRouter.Use(jwtauth.Authenticator)

	// Public routes
	public_api.HandlerFromMux(&sh_public, publicRouter)
	// Private routes
	private_api.HandlerFromMux(&sh_private, privateRouter)

	s.R.Mount("/public", publicRouter)
	s.R.Mount("/api", privateRouter)

	s.Run()

}

func connectDb(cfg Config) (*pgxpool.Pool, *idm.Queries) {
	connStr := cfg.DemoDb.toDbConfig().toDatabaseUrl()
	connPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		slog.Error("Failed creating dbpool", "db", "onboard_demo_db", "url", connStr)
		os.Exit(-1)
	}

	err = connPool.Ping(context.Background())
	if err != nil {
		slog.Error("Unable to ping the database: %v", err)
		os.Exit(-1)
	} else {
		slog.Info("Successfully connected to the database")
	}

	return connPool, idm.New(connPool)
}
