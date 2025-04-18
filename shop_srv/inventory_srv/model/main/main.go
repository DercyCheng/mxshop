package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"shop_srv/inventory_srv/model"
	"shop_srv/inventory_srv/global"
)

func main() {
	initialize.InitConfig()
	dsn := fmt.Sprintf("root:jiushi@tcp(%s:3306)/shop_inventory_srv?charset=utf8mb4&parseTime=True&loc=Local", global.ServerConfig.MysqlInfo.Host)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	//_ = db.AutoMigrate(&model.Inventory{}, &model.StockSellDetail{})
	//插入一条数据
	orderDetail := model.StockSellDetail{
		OrderSn: "1-bobby",
		Status:  1,
		Detail:  []model.GoodsDetail{{1, 2}, {2, 3}},
	}
	db.Create(&orderDetail)

	var sellDetail model.StockSellDetail
	db.Where(model.StockSellDetail{OrderSn: "1-bobby"}).First(&sellDetail)
	fmt.Println(sellDetail.Detail)
}
