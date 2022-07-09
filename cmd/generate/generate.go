package main

import (
	"gorm.io/gen"
	mysql "inkfinance.xyz/conf"
)

func init() {
	mysql.DB = mysql.ConnectDB(mysql.MySQLDSN)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "dal/query",
	})
	g.UseDB(mysql.DB)

	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
