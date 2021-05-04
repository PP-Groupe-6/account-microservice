package main

import (
	"net/http"
	"os"

	accountService "github.com/PP-Groupe-6/account-microservice/account_microservice"
	"github.com/go-kit/kit/log"
)

func main() {
	info := accountService.DbConnexionInfo{
		DbUrl:    "postgre://postgres",
		DbPort:   "5432",
		DbName:   "prix_banque_test",
		Username: "dev",
		Password: "dev",
	}

	service := accountService.NewAccountService(info)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	err := http.ListenAndServe(":8000", accountService.MakeHTTPHandler(service, logger))
	if err != nil {
		panic(err)
	}
}
