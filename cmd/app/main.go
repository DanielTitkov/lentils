package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/DanielTitkov/lentils/cmd/app/prepare"
	"github.com/DanielTitkov/lentils/internal/app"
	"github.com/DanielTitkov/lentils/internal/configs"
	"github.com/DanielTitkov/lentils/internal/handler"
	"github.com/DanielTitkov/lentils/internal/logger"
	"github.com/DanielTitkov/lentils/internal/repository/entgo"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/twitter"

	_ "github.com/lib/pq"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("failed to load config", errors.New("config path is not provided"))
	}
	configPath := args[0]
	log.Println("loading config from "+configPath, "")

	cfg, err := configs.ReadConfigs(configPath)
	if err != nil {
		log.Fatal("failed to load config", err)
	}
	log.Println("loaded config")

	logger := logger.NewLogger(cfg.Env)
	defer logger.Sync()
	logger.Info("starting service", "")

	var dbOptions []ent.Option

	if cfg.Env == "dev" {
		if cfg.Debug.LogDBQueries {
			dbOptions = append(dbOptions, ent.Debug())
		}
	}
	db, err := ent.Open(cfg.DB.Driver, cfg.DB.URI, dbOptions...)
	if err != nil {
		logger.Fatal("failed connecting to database", err)
	}
	defer db.Close()
	logger.Info("connected to database", cfg.DB.Driver+", "+cfg.DB.URI)

	err = prepare.Migrate(context.Background(), db) // run db migration
	if err != nil {
		logger.Fatal("failed creating schema resources", err)
	}
	logger.Info("migrations done", "")

	repo := entgo.NewEntgoRepository(db, logger)

	store := prepare.Store(cfg)

	a, err := app.New(cfg, logger, repo, store.Store)
	if err != nil {
		logger.Fatal("failed to init app", err)
	}

	// email service
	// es, err := email.New(cfg.External.Sendinblue.Key)
	// if err != nil {
	// 	logger.Fatal("failed to init email service", err)
	// }

	// // body := sendinblue.CreateDoiContact{
	// // 	Email: "catie19924343@gmail.com",
	// // 	// Attributes:     attr,
	// // 	IncludeListIds: []int64{3},
	// // 	// TemplateId:     int64(1),
	// // 	RedirectionUrl: "https://lentils.live",
	// // }

	// body := sendinblue.CreateContact{
	// 	Email:   "catie19924343@gmail.com",
	// 	ListIds: []int64{3},
	// }

	// model, resp, err := es.Client.ContactsApi.CreateContact(context.Background(), body)
	// // resp, err := es.Client.ContactsApi.CreateDoiContact(context.Background(), body)
	// if err != nil {
	// 	fmt.Println("Error in ContactsApi->CreateDoiContact ", err.Error())
	// }
	// fmt.Println("CreateDoiContact response:", resp, model)
	// email end

	gothic.Store = store.Store
	goth.UseProviders(
		google.New(
			cfg.Auth.Google.Client,   // client
			cfg.Auth.Google.Secret,   // secret
			cfg.Auth.Google.Callback, // callback url
			"email", "profile",       // scopes
		),
		github.New(
			cfg.Auth.Github.Client,
			cfg.Auth.Github.Secret,
			cfg.Auth.Github.Callback,
			"email", "profile",
		),
		twitter.NewAuthenticate(
			cfg.Auth.Twitter.Client,
			cfg.Auth.Twitter.Secret,
			cfg.Auth.Twitter.Callback,
		),
	)

	h := handler.NewHandler(a, logger, "templates/")
	r := prepare.Mux(cfg, store, h)

	httpServer := prepare.Server(cfg, r)
	httpServer.Addr = cfg.Server.GetAddress()
	logger.Info("starting http server", cfg.Server.GetAddress())
	log.Fatal(httpServer.ListenAndServe())
}
