package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MohitArora1/student/models"
	"github.com/MohitArora1/student/utils"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//GetStudents API will return all the students
func GetStudents(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	session, err := utils.GetDatabaseSession()
	if err != nil {
		log.Printf("Unexpeted error :%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	database := session.DB(utils.Config.DatabaseName)
	students, err := getStudents(database)
	if err != nil {
		log.Printf("error in getting students records :%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, students)
}

func getStudents(database *mgo.Database) (students []models.Student, err error) {
	collection := database.C("students")
	err = collection.Find(nil).All(&students)
	return
}

//PostStudent will save data
func PostStudent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	session, err := utils.GetDatabaseSession()
	if err != nil {
		log.Printf("Unexpeted error :%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	database := session.DB(utils.Config.DatabaseName)
	student, err := postStudent(database, r)
	if err != nil {
		log.Printf("error in Post student records :%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, student)
}

func postStudent(database *mgo.Database, r *http.Request) (student models.Student, err error) {
	json.NewDecoder(r.Body).Decode(&student)
	collection := database.C("students")
	student.ID = bson.NewObjectId()
	err = collection.Insert(&student)
	return
}
