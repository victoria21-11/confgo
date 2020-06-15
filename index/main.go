package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
  "github.com/labstack/echo"

  "domain"
  "time"
  "fmt"
  // "context"

  // _articleHttpDelivery "article/delivery/http"
  // _articleRepo "article/repository/mysql"
  // _articleUcase "article/usecase"

  // _conferenceRepository "conference/repository/mysql"
  // _conferenceUsecase "conference/usecase"
  // _conferenceHandler "conference/delivery/http"

  _eventRepository "event/repository/mysql"
  _eventUsecase "event/usecase"
  _eventHandler "event/delivery/http"

  "github.com/spf13/viper"
)

func main() {
  db, err := gorm.Open("mysql", "root:WK58iYfg6xEr@/testgo?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()

  e := echo.New()

  db.AutoMigrate(&domain.Event{})


  timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

  eventRepository := _eventRepository.Create(db)
  eventUsecase := _eventUsecase.Create(eventRepository, timeoutContext)
  _eventHandler.Create(e, eventUsecase)

  e.Logger.Fatal(e.Start(":1323"))

  // createTestItem(db)

}

func createTestItem(db *gorm.DB) {
  db.Create(&domain.Conference{Title: "Test", Description: "Test Description", Active: true, StartDate: time.Now(), EndDate: time.Now()})

  var conference domain.Conference
  db.First(&conference, 1)
  fmt.Println(conference)
}