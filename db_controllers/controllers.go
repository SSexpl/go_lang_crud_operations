package controllers

// this particular file will contain functions to do crud operations in golang -- mongodb
import (
	"context"
	"dbtest/db"
	"dbtest/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client = db.Db()

func AddUser(w http.ResponseWriter, r *http.Request) {

	var person models.User
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Fatal("there is an error in the received json body for addition")
		return
	} else {

		collection := client.Database("golang").Collection("user") // refernce to the collection to be used in the program.
		insertResult, err := collection.InsertOne(context.TODO(), person)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(insertResult.InsertedID) // return the //mongodb ID of generated document
	}

}

// one way is to define struct for getting the req and using the attributes
type ReqStruct struct {
	Attribute string `json : "attribute"`
	Value     string `json : "value"`
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// we will receive 2 things here 1 is attribute and second is value ..

	var req ReqStruct
	err := json.NewDecoder(r.Body).Decode(&req)
	fmt.Println(req.Attribute, " ", req.Value)
	if err != nil {
		log.Fatal("error in req. in delete route")

	} else {
		//here we will create a bson.D filter
		collection := client.Database("golang").Collection("user")
		filter := bson.D{{Key: req.Attribute, Value: req.Value}}

		res, err := collection.DeleteMany(context.TODO(), filter)
		if err != nil {
			log.Fatal("problem while deleting data")
		} else {
			json.NewEncoder(w).Encode(res.DeletedCount)
		}
	}

	//now we will delete

}
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	//default functioning defining attributes for ResponseWritter
	w.Header().Set("Content-Type", "application/json")

	collection := client.Database("golang").Collection("user") // reference to collection user for querying

	filter := bson.D{} //defining an empty filter for extracting all the users

	cursor, err := collection.Find(context.TODO(), filter) // cursor is like a refernce/pointer to all the doc that satisfied the cond. as specified in filter
	fmt.Printf("cursor = %T\n", cursor)
	if err != nil {
		fmt.Println(err)

	} else {
		var results []models.User // define a slice of type users to retrieve an slice from the cursor
		e := cursor.All(context.TODO(), &results)
		// 1.this converts the cursor to a slice of User type and stores that into results
		//2. also since here desirialisation is happening from the bson to User the type semantics are taken care of . so 'bson' -name ->User.Name
		if e != nil {
			log.Fatal("error found in cursor conversion")

		}

		json.NewEncoder(w).Encode(results) // returns a slice containing the results or a slice of maps Users
	}
}
func GetParticularUser(w http.ResponseWriter, r *http.Request) {
	//default functioning defining attributes for ResponseWritter
	w.Header().Set("Content-Type", "application/json")

	collection := client.Database("golang").Collection("user") // reference to collection user for querying
	params := mux.Vars(r)
	id := params["id"]
	docid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", docid}} //defining an empty filter for extracting all the users

	cursor, err := collection.Find(context.TODO(), filter) // cursor is like a refernce/pointer to all the doc that satisfied the cond. as specified in filter
	fmt.Printf("cursor = %T\n", cursor)
	if err != nil {
		fmt.Println(err)

	} else {
		var results []models.User // define a slice of type users to retrieve an slice from the cursor
		e := cursor.All(context.TODO(), &results)
		// 1.this converts the cursor to a slice of User type and stores that into results
		//2. also since here desirialisation is happening from the bson to User the type semantics are taken care of . so 'bson' -name ->User.Name
		if e != nil {
			log.Fatal("error found in cursor conversion")

		}

		json.NewEncoder(w).Encode(results) // returns a slice containing the results or a slice of maps Users
	}
}
func UpdateDetails(w http.ResponseWriter, r *http.Request) {

	// update the details of a particular user with a given id.
	// refernce to the user collection
	collection := client.Database("golang").Collection("user")
	params := mux.Vars(r)
	id := params["id"]
	//step 3 define the update query to update particular attributes.
	var updateData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updateData)
	log.Println(updateData)

	if err != nil {
		log.Fatal("json body cant be processed")
	} else {
		//as interface doesnot define the strict datatype we need type conversion explicit to get a string datat type for id.

		docid, _ := primitive.ObjectIDFromHex(id) // the id in the mongodb is of typeDoc type so it is not used as string
		//delete the attribute id and chnage the rest

		filter := bson.D{{"_id", docid}}

		update := bson.D{{"$set", updateData}}
		log.Println(updateData)

		updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			http.Error(w, "Failed to update document", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(updateResult)
	}

}
