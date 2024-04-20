package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-go-app/config"
	"simple-go-app/internal/handler"
	"simple-go-app/internal/repository"
	"simple-go-app/internal/router"
	"simple-go-app/internal/service"
	"simple-go-app/pkg/db/postgres"

	"github.com/gorilla/mux"
)

// _ "github.com/lib/pq"

// const (
// 	host     = "localhost"
// 	port     = "5432"
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "sgp"
// )

// var psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// var db *sql.DB

// func connectToPostgres() error {
// 	var err error
// 	db, err = sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// type Student struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// 	Age   int    `json:"age"`
// }

func main() {

	config, err := config.GetPostgresConfig()

	if err != nil {
		log.Fatalf("Config err: %v", err)
		return
	}

	db, err := postgres.NewPsqlDB(config)

	if err != nil {
		log.Fatalf("DB connection error: %v", err)
		return
	}

	defer db.Close()

	fmt.Println(db.Stats().InUse)

	repo := repository.StudentRepo{DB: db}
	service := service.StudentServ{SR: &repo}
	handler := handler.StudentHandler{StudentService: &service}

	r := mux.NewRouter()
	router.InitRouter(r, &handler)

	log.Fatal(http.ListenAndServe(":8080", r))

}

// err := connectToPostgres()
// if err != nil {
// 	return
// }
// defer db.Close()
// r := mux.NewRouter()
// r.HandleFunc("/", HomePageFunction).Methods("GET")
// r.HandleFunc("/items", GetAllItems).Methods("GET")
// r.HandleFunc("/items/{id}", GetItem).Methods("GET")
// r.HandleFunc("/items", CreateItem).Methods("POST")
// r.HandleFunc("/items/{id}", UpdateItem).Methods("PUT")
// r.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")

// log.Fatal(http.ListenAndServe(":8080", r))
// r := route.NewRouter()
// log.Fatal(http.ListenAndServe(":8080", r))

// func HomePageFunction(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Home page")
// }

// func GetAllItems(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var students []Student
// 	rows, err := db.Query("select * from student")
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("error in select query: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var student Student
// 		if err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Age); err != nil {
// 			http.Error(w, fmt.Sprintf("error in query all student: %v", err), 404)
// 			return
// 		}
// 		students = append(students, student)
// 	}

// 	if err := rows.Err(); err != nil {
// 		http.Error(w, fmt.Sprintf("error in query all student: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	json, err := json.Marshal(students)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(json)
// }

// func GetItem(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	idParam := vars["id"]
// 	var student Student

// 	if idParam == "" {
// 		http.Error(w, "id is empty", http.StatusBadRequest)
// 		return
// 	}

// 	id, err := strconv.Atoi(idParam)

// 	if err != nil {
// 		http.Error(w, "Invalid id param", http.StatusBadRequest)
// 		return
// 	}

// 	err = db.QueryRow("select id, name, email, age from student where id=$1", id).Scan(&student.ID, &student.Name, &student.Email, &student.Age)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			http.Error(w, fmt.Sprintf("student not found: %v", err), http.StatusBadRequest)
// 		} else {
// 			http.Error(w, fmt.Sprintf("Internal error: %v", err), http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	jsonStudent, err := json.Marshal(student)

// 	if err != nil {
// 		http.Error(w, "Error while parsing into JSON", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonStudent)
// }

// func CreateItem(w http.ResponseWriter, r *http.Request) {
// 	var student Student
// 	var id int

// 	err := json.NewDecoder(r.Body).Decode(&student)

// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("error occured while parsing request body: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	err = db.QueryRow("insert into student(name, email, age) values($1, $2, $3) returning id",
// 		student.Name, student.Email, student.Age).Scan(&id)

// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Errir while inserting object into database: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	jsonId, _ := json.Marshal(id)
// 	fmt.Fprintf(w, "%s", jsonId)
// }

// func UpdateItem(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	idParam := vars["id"]

// 	if idParam == "" {
// 		http.Error(w, "empty id", http.StatusBadRequest)
// 		return
// 	}

// 	id, err := strconv.Atoi(idParam)

// 	if err != nil {
// 		http.Error(w, "invalid id", http.StatusBadRequest)
// 		return
// 	}

// 	var student Student

// 	err = db.QueryRow("select id, name, email, age from student where id=$1", id).Scan(&student.ID, &student.Name, &student.Email, &student.Age)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			http.Error(w, fmt.Sprintf("student not found: %v", err), http.StatusBadRequest)
// 		} else {
// 			http.Error(w, fmt.Sprintf("Internal error: %v", err), http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	err = json.NewDecoder(r.Body).Decode(&student)

// 	if err != nil {
// 		http.Error(w, "error while parsing json", http.StatusInternalServerError)
// 		return
// 	}

// 	_, err = db.Exec("update student set name=$1, email=$2, age=$3 where id=$4", student.Name, student.Email, student.Age, student.ID)

// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("error occured while updating: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	updatedStudentJSON, _ := json.Marshal(student)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(updatedStudentJSON)
// }

// func DeleteItem(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	idParam := vars["id"]
// 	id, err := strconv.Atoi(idParam)

// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("error occured while parsing to int: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// _, err = db.Exec("delete from student where id=$1", id)
// 	// if err != nil {
// 	// 	http.Error(w, fmt.Sprintf("error occured while deleting: %v", err), http.StatusInternalServerError)
// 	// 	return
// 	// }
// 	err = db.QueryRow("delete from student where id=$1 returning id", id).Scan(&id)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("error occured while deleting: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	idJSON, _ := json.Marshal(id)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(idJSON)
// }
