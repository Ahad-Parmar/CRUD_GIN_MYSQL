

package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type truck struct {
	TruckID       int      `json:"truckId"`
	DriverName    string   `json:"driverName"`
	CleanerName   string   `json:"cleanerName"`
	
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "golang"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	router := gin.Default()

	router.POST("/create", func(ctx *gin.Context) {
		var t truck
		if ctx.BindJSON(&s) == nil {
			ctx.JSON(200, gin.H{
				"driverName":   t.DriverName,
				"cleanerName": t.CleanerName,
				
			})
			db := dbConn()
			insert, err := db.Query("INSERT INTO truck(driverName, cleanerName) VALUES(?,?,?)", t.DriverName, t.CleanerName)
			if err != nil {
				panic(err.Error())
			}
			defer insert.Close()
			// insert.Exec(t.DriverName, t.CleanerName)
			fmt.Printf("driverName: %s, cleanerName: %s, t.DriverName, t.CleanerName)
		}
	})

	router.PUT("/update", func(ctx *gin.Context) {
		var t truck
		if ctx.BindJSON(&s) == nil {
			ctx.JSON(200, gin.H{
				"driverName":   t.DriverName,
				"cleanerName": t.CleanerName,
				
			})
			db := dbConn()
			update, err := db.Prepare("UPDATE truck SET driverName=?, cleanerName=? Where truckId=?")
			if err != nil {
				panic(err.Error())
			}
			update.Exec(t.DriverName, t.CleanerName, t.TruckId)
		}
	})

	router.GET("/read", func(ctx *gin.Context) {
		db := dbConn()
		selDB, err := db.Query("SELECT * FROM truck")
		if err != nil {
			panic(err.Error())
		}
		var truckId int
		var driverName, cleanerName string
		for selDB.Next() {
			err = selDB.Scan(&truckId, &driverName, &cleanerName)
			ctx.JSON(200, gin.H{
				"truckId":      truckId,
				"driverName":   driverName,
				"cleanerName":  cleanerName,
				
			})
			fmt.Printf("driverName: %s, cleanerName: %s, driverName, cleanerName)
			if err != nil {
				panic(err.Error())
			}
		}
	})

	router.DELETE("/delete", func(ctx *gin.Context) {
		var t truck
		if ctx.BindJSON(&s) == nil {
			db := dbConn()
			del, err := db.Prepare("DELETE FROM truck WHERE driverName=?")
			if err != nil {
				panic(err.Error())
			}
			del.Exec(t.DriverName)
			log.Println("DELETE")
			defer db.Close()
		}
	})

	router.Run(":8000")

}
