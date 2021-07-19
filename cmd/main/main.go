package main

import (
	"github.com/githanselBonifacio/MiPrimeraAPIRestChi/internal/database"
	"github.com/githanselBonifacio/MiPrimeraAPIRestChi/internal/logs"
	"github.com/golang-migrate/migrate/v4"
)

func main() {
	_ = logs.InitLogger()
	client := database.NewSqlClient("root:root@/phones-review")
}
