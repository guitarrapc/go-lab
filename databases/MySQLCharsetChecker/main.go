package main

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// InfomationSchemaTables definition for information_schema.Tables
type Tables struct {
	TableName      string `gorm:"column:TABLE_NAME"`
	TableCollation string `gorm:"column:TABLE_COLLATION"`
}

const (
	expectedCharset             = "utf8mb4"
	databaseName                = "<databaseNAME>"
	connectionString            = "<userName>:<Password>@tcp(<mysqlHostAddress>)/<yourDatabase>?parseTime=true&tls=true&loc=Asia%2FTokyo"
	informationschemaConnection = "<userName>:<Password>@tcp(<mysqlHostAddress>)/information_schema?parseTime=true&tls=true&loc=Asia%2FTokyo"
)

func main() {
	fmt.Println("connect database...")
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	dbInfo, err := gorm.Open("mysql", informationschemaConnection)
	if err != nil {
		panic(err.Error())
	}
	defer dbInfo.Close()

	fmt.Println("get table infos...")
	var tables []Tables
	// dbInfo.Raw("SELECT * FROM TABLES WHERE TABLE_SCHEMA = 'juiz_dot' WHERE ").Scan(&tables)
	dbInfo.Table("TABLES").Where("TABLE_SCHEMA = ?", databaseName).Select("*").Scan(&tables)
	//fmt.Println(tables)

	fmt.Println("get unexpected collation tables...")
	for _, table := range tables {
		if !strings.HasPrefix(table.TableCollation, expectedCharset) {
			fmt.Println(table.TableName)
		}
	}
}
