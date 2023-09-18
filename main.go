package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

var (
	DB *gorm.DB
)

func init()  {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port string
	DB_Host string
	DB_Name string
}

func InitDB()  {
	config := Config {
		DB_Username: "root",
		DB_Password: "",
		DB_Port: "3306",
		DB_Host: "localhost",
		DB_Name: "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name string `json:"name" from:"name"`
	Email string `json:"email" from:"email"`
	Password string `json:"password" from:"password"`
}

func InitialMigration()  {
	DB.AutoMigrate(&User{})
}

// get all Users
func GetUsersController(c echo.Context) error {
	users := []User{}

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg": "Success get all users",
		"Data": users,
	})
}

// get User by Id
func GetUserController(c echo.Context) error {
	
}

// create new User
func CreateuserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code": 201,
		"msg": "Success create new user",
		"Data": user,
	})
}

// Delete User by Id
func DeleteUserController(c echo.Context) error {
	
}

// Update User by Id
func UpdateUserController(c echo.Context) error {
	
}

func main()  {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateuserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}