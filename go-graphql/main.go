package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

type Person struct {
	ID   string `json:"id"`
	Name string `json:"text"`
	Age  int8   `json:"int"`
}

var Persons []Person

var personType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

func init() {
	tom := Person{ID: "a", Name: "Im Tom", Age: 12}
	hana := Person{ID: "b", Name: "Im Hana", Age: 23}
	bob := Person{ID: "c", Name: "Im Bob", Age: 23}
	Persons = append(Persons, tom, hana, bob)
}

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "persons",
	Fields: graphql.Fields{
		"persons": &graphql.Field{
			Type:        graphql.NewList(personType),
			Description: "List of persons",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return Persons, nil
			},
		},
	},
})

func main() {
	// Schema
	schemaConfig := graphql.SchemaConfig{Query: rootQuery}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			persons {
				id
				name
				age
			}
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
