package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Celebrity struct {
	ID           string `gorm:"primaryKey"`
	FullName     string
	Nationality  string
	ReqPhotoPath string
}

var db *gorm.DB

func initDB() {

	dsn := "host=postgres user=postgres password=postgres dbname=celebrities port=5432 sslmode=disable"

	var err error

	db, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("Ошибка БД:", err)
	}

	db.AutoMigrate(&Celebrity{})

	log.Println("PostgreSQL подключена")
}

var celebrityType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Celebrity",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"fullName": &graphql.Field{
				Type: graphql.String,
			},
			"nationality": &graphql.Field{
				Type: graphql.String,
			},
			"reqPhotoPath": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{

			"list": &graphql.Field{
				Type: graphql.NewList(celebrityType),

				Resolve: func(
					p graphql.ResolveParams,
				) (interface{}, error) {

					var list []Celebrity

					db.Find(&list)

					return list, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{

			"add": &graphql.Field{
				Type: celebrityType,

				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(
							graphql.String,
						),
					},

					"fullName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(
							graphql.String,
						),
					},

					"nationality": &graphql.ArgumentConfig{
						Type: graphql.String,
					},

					"photo": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},

				Resolve: func(
					p graphql.ResolveParams,
				) (interface{}, error) {

					c := Celebrity{
						ID:           p.Args["id"].(string),
						FullName:     p.Args["fullName"].(string),
						Nationality:  p.Args["nationality"].(string),
						ReqPhotoPath: p.Args["photo"].(string),
					}

					err := db.Create(&c).Error

					return c, err
				},
			},
		},
	},
)

func main() {

	initDB()

	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)

	h := handler.New(
		&handler.Config{
			Schema:   &schema,
			Pretty:   true,
			GraphiQL: true,
		},
	)

	r := mux.NewRouter()

	r.Handle("/graphql", h)

	log.Println("Сервер на http://localhost:8080/graphql")

	log.Fatal(
		http.ListenAndServe(":8080", r),
	)
}
