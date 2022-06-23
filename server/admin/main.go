package main

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/teleport-network/teleport-analytics-dashboard/server/admin/models"
	"github.com/teleport-network/teleport-analytics-dashboard/server/admin/pages"
	"github.com/teleport-network/teleport-analytics-dashboard/server/admin/tables"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"              // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	_ "github.com/GoAdminGroup/themes/sword"                      // ui theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"
)

func main() {
	startServer()
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	template.AddComp(chartjs.NewChart())

	eng := engine.Default()
	eng.AddConfig(&config.Config{
		Databases: config.DatabaseList{
			"default": config.Database{
				Host:   os.Getenv("DATABASE_HOST"),
				Port:   os.Getenv("DATABASE_PORT"),
				User:   os.Getenv("DATABASE_USER"),
				Pwd:    os.Getenv("DATABASE_PWD"),
				Name:   os.Getenv("DATABASE_NAME"),
				Driver: os.Getenv("DATABASE_DRIVER"),
			},
		},
		AppID:     os.Getenv("APP_ID"),
		Language:  os.Getenv("LANGUAGE"),
		UrlPrefix: os.Getenv("URL_PREFIX"),
		Theme:     os.Getenv("THEME"),
		Store: config.Store{
			Path:   os.Getenv("STORE_PATH"),
			Prefix: os.Getenv("STORE_PREFIX"),
		},
		Title:              os.Getenv("TITLE"),
		Logo:               template.HTML(os.Getenv("LOGO")),
		MiniLogo:           template.HTML(os.Getenv("MINI_LOGO")),
		IndexUrl:           os.Getenv("INDEX_URL"),
		LoginUrl:           os.Getenv("LOGIN_URL"),
		Debug:              os.Getenv("DEBUG") == "true",
		Env:                os.Getenv("ENV"),
		InfoLogPath:        os.Getenv("INFO_LOG_PATH"),
		ErrorLogPath:       os.Getenv("ERROR_LOG_PATH"),
		AccessLogPath:      os.Getenv("ACCESS_LOG_PATH"),
		AccessAssetsLogOff: false,
		SqlLog:             false,
		AccessLogOff:       false,
		InfoLogOff:         false,
		ErrorLogOff:        false,
		SessionLifeTime:    86400,
		AssetUrl:           os.Getenv("ASSET_URL"),
		FileUploadEngine: config.FileUploadEngine{
			Name: os.Getenv("FILE_UPLOAD_ENGINE_NAME"),
		},
	})

	if err := eng.AddConfigFromYAML("./config.yml").
		AddGenerators(tables.Generators).
		Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/admin", pages.GetDashBoard)
	eng.HTMLFile("GET", "/admin/hello", "./html/hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})

	models.Init(eng.MysqlConnection())

	_ = r.Run(":" + os.Getenv("PORT"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.MysqlConnection().Close()
}
