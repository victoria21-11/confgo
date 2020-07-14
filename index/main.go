package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
  "github.com/labstack/echo"

  "domain"
  "time"
  "fmt"
  // "context"

  _conferenceRepository "conference/repository/mysql"
  _conferenceUsecase "conference/usecase"
  _conferenceHandler "conference/delivery/http"

  _eventRepository "event/repository/mysql"
  _eventUsecase "event/usecase"
  _eventHandler "event/delivery/http"

  _roleRepository "role/repository/mysql"
  _roleUsecase "role/usecase"
  _roleHandler "role/delivery/http"

  _sectionRepository "section/repository/mysql"
  _sectionUsecase "section/usecase"
  _sectionHandler "section/delivery/http"

  "github.com/spf13/viper"
)

func main() {
  db, err := gorm.Open("mysql", "root:WK58iYfg6xEr@/testgo?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()

  e := echo.New()

  db.AutoMigrate(&domain.Section{})
  db.AutoMigrate(&domain.Role{})
  db.AutoMigrate(&domain.Event{})
  db.AutoMigrate(&domain.Conference{})

  timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

  eventRepository := _eventRepository.Create(db)
  eventUsecase := _eventUsecase.Create(eventRepository, timeoutContext)
  _eventHandler.Create(e, eventUsecase)

  conferenceRepository := _conferenceRepository.Create(db)
  conferenceUsecase := _conferenceUsecase.Create(conferenceRepository, timeoutContext)
  _conferenceHandler.Create(e, conferenceUsecase)

  roleRepository := _roleRepository.Create(db)
  roleUsecase := _roleUsecase.Create(roleRepository, timeoutContext)
  _roleHandler.Create(e, roleUsecase)

  sectionRepository := _sectionRepository.Create(db)
  sectionUsecase := _sectionUsecase.Create(sectionRepository, timeoutContext)
  _sectionHandler.Create(e, sectionUsecase)

  e.Logger.Fatal(e.Start(":4040"))

}

func createTestItem(db *gorm.DB) {
  // db.Create(&domain.Conference{Title: "Test", Description: "Test Description", Active: true, StartDate: time.Now(), EndDate: time.Now()})

  var conferences []domain.Conference
  db.Find(&conferences)
  fmt.Println(conferences)
}

func createSection(db *gorm.DB) {
  db.Create(&domain.Section{Title: "111", Description: "111", ConferenceID: 10, CoordinatorID: 1})
}

func deleteTestItem(db *gorm.DB) {
  db.Where("id = ?", 1).Delete(&domain.Conference{})
}