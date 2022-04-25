package db

import (
	"GO_gRPC/models"
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(path string) *gorm.DB {

	//protocol://username:password@host:port/database

	conf := &dbConfig{}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&conf); err != nil {
		panic(err)
	}

	dbURL := conf.Protocol + "://" + conf.Username + ":" + conf.Password + "@" + conf.Host + ":" + conf.Port + "/" + conf.DB

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Plants{})
	db.AutoMigrate(&models.Rooms{})
	db.AutoMigrate(&models.Species{})
	db.AutoMigrate(&models.WateringDays{})

	var plant = models.Plants{}

	plant.Id = 1
	plant.Name = "Фиалка ночная"
	plant.Room = 5
	plant.Species = 2
	plant.WateringDay = 6
	db.Create(&plant)

	plant.Id = 2
	plant.Name = "Папоротник обыкновенный"
	plant.Room = 1
	plant.Species = 4
	plant.WateringDay = 1
	db.Create(&plant)

	var room = models.Rooms{}

	room.Id = 1
	room.Name = "Прихожая"
	db.Create(&room)
	room.Id = 2
	room.Name = "Коридор"
	db.Create(&room)
	room.Id = 3
	room.Name = "Кухня"
	db.Create(&room)
	room.Id = 4
	room.Name = "Гостинная"
	db.Create(&room)
	room.Id = 5
	room.Name = "Спальня"
	db.Create(&room)
	room.Id = 6
	room.Name = "Кабинет"
	db.Create(&room)
	room.Id = 7
	room.Name = "Балкон"
	db.Create(&room)

	var species = models.Species{}

	species.Id = 1
	species.Name = "Однодольные"
	db.Create(&species)
	species.Id = 2
	species.Name = "Двудольные"
	db.Create(&species)
	species.Id = 3
	species.Name = "Мхи"
	db.Create(&species)
	species.Id = 4
	species.Name = "Папоротники"
	db.Create(&species)

	var wdays = models.WateringDays{}

	wdays.Id = 1
	wdays.Name = "Понедельник"
	db.Create(&wdays)
	wdays.Id = 2
	wdays.Name = "Вторник"
	db.Create(&wdays)
	wdays.Id = 3
	wdays.Name = "Среда"
	db.Create(&wdays)
	wdays.Id = 4
	wdays.Name = "Четверг"
	db.Create(&wdays)
	wdays.Id = 5
	wdays.Name = "Пятница"
	db.Create(&wdays)
	wdays.Id = 6
	wdays.Name = "Суббота"
	db.Create(&wdays)
	wdays.Id = 7
	wdays.Name = "Воскресенье"
	db.Create(&wdays)

	return db
}

func NewDB() *gorm.DB {

	return InitDB("db/config.yml")

}
