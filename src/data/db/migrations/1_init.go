package migrations

import (
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/constants"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"github.com/alireza-fa/golang-car-shop/data/models"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()

	createTables(database)
	createDefaultUserInformation(database)
	createCountry(database)
	createPropertyCategory(database)
	createCarType(database)
}

func createTables(database *gorm.DB) {
	tables := []interface{}{}

	// Basic
	tables = addNewTable(database, models.Country{}, tables)
	tables = addNewTable(database, models.City{}, tables)
	tables = addNewTable(database, models.PersianYear{}, tables)
	tables = addNewTable(database, models.File{}, tables)

	// Property
	tables = addNewTable(database, models.PropertyCategory{}, tables)
	tables = addNewTable(database, models.Property{}, tables)

	// User
	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.UserRole{}, tables)

	// Car
	tables = addNewTable(database, models.Company{}, tables)
	tables = addNewTable(database, models.Gearbox{}, tables)
	tables = addNewTable(database, models.Color{}, tables)
	tables = addNewTable(database, models.CarType{}, tables)

	tables = addNewTable(database, models.CarModel{}, tables)
	tables = addNewTable(database, models.CarModelColor{}, tables)
	tables = addNewTable(database, models.CarModelYear{}, tables)
	tables = addNewTable(database, models.CarModelImage{}, tables)
	tables = addNewTable(database, models.CarModelPriceHistory{}, tables)
	tables = addNewTable(database, models.CarModelProperty{}, tables)
	tables = addNewTable(database, models.CarModelComment{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}
	logger.Info(logging.Postgres, logging.Migration, "created tables", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultUserInformation(database *gorm.DB) {
	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	adminUser := models.User{}
	adminUser.Username = constants.DefaultUserName
	adminUser.FirstName = "Test"
	adminUser.LastName = "Test"
	adminUser.MobileNumber = "09121111111"
	adminUser.Email = "admin@admin.com"
	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	adminUser.Password = string(hashedPassword)

	createAdminUserIfNotExist(database, &adminUser, adminRole.Id)
}

func createAdminUserIfNotExist(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.
		Model(models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createCountry(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Country{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Country{Name: "Iran", Cities: []models.City{
			{Name: "Tehran"},
			{Name: "Isfahan"},
			{Name: "Shiraz"},
			{Name: "Chalus"},
			{Name: "Ahwaz"},
		}, Companies: []models.Company{
			{Name: "Saipa"},
			{Name: "Iran Khodro"},
		}})
		database.Create(&models.Country{Name: "USA", Cities: []models.City{
			{Name: "New York"},
			{Name: "Washington"},
		}, Companies: []models.Company{
			{Name: "Opel"},
			{Name: "Benz"},
		}})
		database.Create(&models.Country{Name: "Germany", Cities: []models.City{
			{Name: "Berlin"},
			{Name: "Munich"},
		}, Companies: []models.Company{
			{Name: "Opel"},
			{Name: "Benz"},
		}})
		database.Create(&models.Country{Name: "China", Cities: []models.City{
			{Name: "Beijing"},
			{Name: "Shanghai"},
		}, Companies: []models.Company{
			{Name: "Chery"},
			{Name: "Geely"},
		}})
		database.Create(&models.Country{Name: "Italy", Cities: []models.City{
			{Name: "Roma"},
			{Name: "Turin"},
		}, Companies: []models.Company{
			{Name: "Ferrari"},
			{Name: "Fiat"},
		}})
		database.Create(&models.Country{Name: "France", Cities: []models.City{
			{Name: "Paris"},
			{Name: "Lyon"},
		}, Companies: []models.Company{
			{Name: "Renault"},
			{Name: "Bugatti"},
		}})
		database.Create(&models.Country{Name: "Japan", Cities: []models.City{
			{Name: "Tokyo"},
			{Name: "Kyoto"},
		}, Companies: []models.Company{
			{Name: "Toyota"},
			{Name: "Honda"},
		}})
		database.Create(&models.Country{Name: "South Korea", Cities: []models.City{
			{Name: "Seoul"},
			{Name: "Ulsan"},
		}, Companies: []models.Company{
			{Name: "Kia"},
			{Name: "Hyundai"},
		}})
	}
}

func createPropertyCategory(database *gorm.DB) {
	count := 0

	database.
		Model(&models.PropertyCategory{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.PropertyCategory{Name: "Body"})                     // بدنه
		database.Create(&models.PropertyCategory{Name: "Engine"})                   // موتور
		database.Create(&models.PropertyCategory{Name: "Drivetrain"})               // پیشرانه
		database.Create(&models.PropertyCategory{Name: "Suspension"})               // تعلیق
		database.Create(&models.PropertyCategory{Name: "Equipment"})                // تجهیزات
		database.Create(&models.PropertyCategory{Name: "Driver support systems"})   // سیستم های پشتیبانی راننده
		database.Create(&models.PropertyCategory{Name: "Lights"})                   // چراغ ها
		database.Create(&models.PropertyCategory{Name: "Multimedia"})               // چند رسانه ای
		database.Create(&models.PropertyCategory{Name: "Safety equipment"})         // تجهیزات ایمنی
		database.Create(&models.PropertyCategory{Name: "Seats and steering wheel"}) // صندلی و فرمان
		database.Create(&models.PropertyCategory{Name: "Windows and mirrors"})      // پنجره و آینه
	}
	createProperty(database, "Body")
	createProperty(database, "Engine")
	createProperty(database, "Drivetrain")
	createProperty(database, "Suspension")
	createProperty(database, "Comfort")
	createProperty(database, "Driver support systems")
	createProperty(database, "Lights")
	createProperty(database, "Multimedia")
	createProperty(database, "Safety equipment")
	createProperty(database, "Seats and steering wheel")
	createProperty(database, "Windows and mirrors")
}

func createProperty(database *gorm.DB, cat string) {
	count := 0
	catModel := models.PropertyCategory{}

	database.
		Model(models.PropertyCategory{}).
		Where("Name = ?", cat).
		Find(&catModel)

	database.
		Model(&models.Property{}).
		Select("count(*)").
		Where("category_id = ?", catModel.Id).
		Find(&count)

	if count > 0 || catModel.Id == 0 {
		return
	}
	var props *[]models.Property
	switch cat {
	case "Body":
		props = getBodyProperties(catModel.Id)

	case "Engine":
		props = getEngineProperties(catModel.Id)

	case "Drivetrain":
		props = getDrivetrainProperties(catModel.Id)

	case "Suspension":
		props = getSuspensionProperties(catModel.Id)

	case "Comfort":
		props = getComfortProperties(catModel.Id)

	case "Driver support systems":
		props = getDriverSupportSystemProperties(catModel.Id)

	case "Lights":
		props = getLightsProperties(catModel.Id)

	case "Multimedia":
		props = getMultimediaProperties(catModel.Id)

	case "Safety equipment":
		props = getSafetyEquipmentProperties(catModel.Id)

	case "Seats and steering wheel":
		props = getSeatsProperties(catModel.Id)

	case "Windows and mirrors":
		props = getWindowsProperties(catModel.Id)
	}

	for _, prop := range *props {
		database.Create(&prop)
	}
}

func createCarType(database *gorm.DB) {
	count := 0
	database.
		Model(&models.CarType{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.CarType{Name: "Crossover"})
		database.Create(&models.CarType{Name: "Sedan"})
		database.Create(&models.CarType{Name: "Sports"})
		database.Create(&models.CarType{Name: "Coupe"})
		database.Create(&models.CarType{Name: "Hatchback"})
	}
}
