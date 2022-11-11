package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var orm *gorm.DB

func OpenDatabase() *gorm.DB {
	dsn := "treenod:xmflshem@tcp(ec2-13-125-229-201.ap-northeast-2.compute.amazonaws.com:3306)/study"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}
	return db
}

func SetupDatabase() {
	db := OpenDatabase()
	orm = db
	//여기서 리턴 안하는 이유는 한번만 해줘도 되는 커넥트를 계속 실행하게 되니까
	//새로 함수를 만들어서 한번 연결해둔 orm을 재활용함!
}

func GetORM() *gorm.DB {
	return orm
}
