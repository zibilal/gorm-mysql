package app

import (
	"github.com/spf13/cobra"
	"gorm-mysql/api"
	"gorm-mysql/appctx"
	"gorm-mysql/businessrule/menurule/controller"
	"gorm-mysql/conf"
	"gorm-mysql/engine/apiengine"
	"gorm-mysql/engine/dbengine"
	"log"
	"os"
)

func InitApp() {
	f, err := os.Open("conf.yaml")
	if err != nil {
		log.Fatalf("Unable to open configuration file: %s", err.Error())
	}

	c, err := conf.GetConfiguration(f)
	if err != nil {
		log.Fatalf("Unable to initialize configuration object")
	}

	dbEngine, err := dbengine.InitDbEngine(
		c.Db.Hostname,
		c.Db.User,
		c.Db.Password,
		c.Db.Name,
		c.Db.Port,
	)
	if err != nil {
		log.Fatalf("Unable to connect to database")
	}

	apiEngine := apiengine.InitApiEngine()

	appctx.InitAppContext()

	appctx.AppContext.DbContext = dbEngine
	appctx.AppContext.Host = c.Host
	appctx.AppContext.ApiEngine = apiEngine

}

func RunApi() {
	InitApp()

	menuRouter := controller.NewMenuRouter(appctx.AppContext)

	a := api.NewApi(appctx.AppContext)
	a.RegisterRouter("menu", "1.0.0", menuRouter)
	a.Run()
}

var CmdRunApi = &cobra.Command{
	Use: "run [service to run]",
	Short: "run a service",
	Long: "run execute and run a service that specified in the run argument. Only one argument will be processed",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		switch args[0] {
		case "apiengine":

			RunApi()

		default:
			log.Fatalf("%s service not found", args[0])

		}


	},
}
