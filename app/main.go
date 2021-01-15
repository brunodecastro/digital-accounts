package app

import (
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/error"
	"github.com/brunodecastro/digital-accounts/app/persistence"
	"github.com/sirupsen/logrus"
)

func main() {
	//var log = logger.NewLogUtil()
	var log = logrus.New()
	log.Infoln("Starting Digital Accounts API...")


	apiConfig, err := config.LoadConfigs()
	error.MaybeFatal(err, "Unable to load api configuration")

	poolConfig := persistence.PoolConfig(&apiConfig.DatabaseConfig, log)
	defer poolConfig.Close()


}