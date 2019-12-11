package main
import (
	"KProject/database"
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
	"net/http"
)
func getProducts(c echo.Context) error {
	productsCollection := database.ProductsDB.C("Products")
	query := bson.M{}
	products := []database.Product{}
	productsCollection.Find(query).All(&products)
	var sb strings.Builder
	for _, p := range products{
		sb.WriteString(p.Category + " " + p.Part + " " + p.Company + " Price:" + strconv.Itoa(p.Price) + "\n")
	}
	return c.String(http.StatusOK, sb.String())
}
func registerUser(c echo.Context) error {
	usersCollection := database.ProductsDB.C("Users")
	username := c.QueryParam("username")
	login := c.QueryParam("login")
	password := c.QueryParam("password")
	query := bson.M{
		"login" : login,
	}
	var users database.User
	usersCollection.Find(query).One(&users)
	if login == "" || password == ""{return c.String(http.StatusOK, "Data entry error")}
	if users.Login == login{ return c.String(http.StatusOK, "Error, user with such login already exists")}
	unew := &database.User{Id:bson.NewObjectId(), Username: username, Login:login, Password:password}
	err := usersCollection.Insert(unew)
	if err != nil{
		fmt.Println(err)
	}
	return c.String(http.StatusOK, "Registration completed succesfully. Welcome, " + username )
}
func loginUser(c echo.Context) error {
	usersCollection := database.ProductsDB.C("Users")
	login := c.QueryParam("login")
	password := c.QueryParam("password")
	query := bson.M{
		"login" : login,
		"password" : password,
	}
	var users database.User
	usersCollection.Find(query).One(&users)
	if users.Login == login && users.Password == password { return c.String(http.StatusOK, "Welcome back, " + users.Username)}
	if login == "" || password == ""{return c.String(http.StatusOK, "Data entry error")}
	return c.String(http.StatusOK, "Incorrect login or password")
}
func productsFilter(c echo.Context) error {
	productsCollection := database.ProductsDB.C("Products")
	category := c.QueryParam("category")
	company := c.QueryParam("company")
	price := c.QueryParam("price")
	var query bson.M
	switch{
	case price != "" : {intprice, err := strconv.Atoi(price)
		if err != nil {
			fmt.Println(err)
		}
		switch {
		case company != "" && category != "":
			query = bson.M{
				"category": category,
				"company":  company,
				"price": bson.M{
					"$lt": intprice,
				},
			}
		case company == "" && category != "":
			query = bson.M{
				"category": category,
				"price": bson.M{
					"$lt": intprice,
				},
			}
		case company != "" && category == "":
			query = bson.M{
				"company":  company,
				"price": bson.M{
					"$lt": intprice,
				},
			}
		case company == "" && category == "":
			query = bson.M{
				"price": bson.M{
					"$lt": intprice,
				},
			}
		}
	}
	case price == "" :
		switch {
		case company != "" && category != "":
			query = bson.M{
				"category": category,
				"company":  company,
			}
		case company == "" && category != "":
			query = bson.M{
				"category": category,
			}
		case company != "" && category == "":
			query = bson.M{
				"company":  company,
			}
		}
	default : query = bson.M{}
	}
	products := []database.Product{}
	productsCollection.Find(query).All(&products)
	var sb strings.Builder
	for _, p := range products{
		sb.WriteString( p.Part + " " + p.Company + " Price:" + strconv.Itoa(p.Price) + "\n")
	}
	return c.String(http.StatusOK, sb.String())
}
func main() {
	database.Init()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Main page")
	})
	e.GET("/products", getProducts)
	e.GET("/register", registerUser)
	e.GET("/login", loginUser)
	e.GET("/filter", productsFilter)
	defer database.Session.Close()
	e.Logger.Fatal(e.Start(":27017"))
}