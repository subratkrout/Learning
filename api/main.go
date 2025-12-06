package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	package main

	import (
		"encoding/json"
		"fmt"
		"log"
		"net/http"
		"os"
		"time"

		"gorm.io/driver/postgres"
		"gorm.io/gorm"
	)

	type User struct {
		ID        uint      `gorm:"primaryKey" json:"id"`
		Name      string    `json:"name"`
		DOB       time.Time `json:"dob"`
		CreatedAt time.Time `json:"created_at"`
	}

	var gormDB *gorm.DB

	func main() {
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			user := os.Getenv("POSTGRES_USER")
			pass := os.Getenv("POSTGRES_PASSWORD")
			db := os.Getenv("POSTGRES_DB")
			if user == "" {
				user = "appuser"
			}
			if pass == "" {
				pass = "apppassword"
			}
			if db == "" {
				db = "appdb"
			}
			dsn = fmt.Sprintf("postgres://%s:%s@db:5432/%s?sslmode=disable", user, pass, db)
		}

		var err error
		gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to db: %v", err)
		}

		// Auto-migrate the User model
		if err := gormDB.AutoMigrate(&User{}); err != nil {
			log.Fatalf("auto migrate failed: %v", err)
		}

		http.HandleFunc("/health", healthHandler)
		http.HandleFunc("/users", usersHandler)

		port := os.Getenv("PORT")
		if port == "" {
			port = "8000"
		}
		addr := ":" + port
		log.Printf("starting api on %s", addr)
		log.Fatal(http.ListenAndServe(addr, nil))
	}

	func healthHandler(w http.ResponseWriter, r *http.Request) {
		sqlDB, err := gormDB.DB()
		dbVersion := "unavailable"
		if err == nil {
			if err := sqlDB.Ping(); err == nil {
				// simple version query
				var ver string
				_ = gormDB.Raw("SELECT version();").Scan(&ver).Error
				if ver != "" {
					dbVersion = ver
				} else {
					dbVersion = "connected"
				}
			}
		}
		resp := map[string]string{"status": "ok", "db_version": dbVersion}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}

	func usersHandler(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			var users []User
			if err := gormDB.Order("id desc").Find(&users).Error; err != nil {
				http.Error(w, "failed to fetch users", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users)
			return

		case http.MethodPost:
			var payload struct {
				Name string `json:"name"`
				DOB  string `json:"dob"` // expect YYYY-MM-DD or RFC3339
			}
			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
				http.Error(w, "invalid json", http.StatusBadRequest)
				return
			}
			if payload.Name == "" || payload.DOB == "" {
				http.Error(w, "name and dob required", http.StatusBadRequest)
				return
			}
			// parse DOB
			dob, err := time.Parse("2006-01-02", payload.DOB)
			if err != nil {
				// try RFC3339
				dob, err = time.Parse(time.RFC3339, payload.DOB)
				if err != nil {
					http.Error(w, "dob must be YYYY-MM-DD or RFC3339", http.StatusBadRequest)
					return
				}
			}

			user := User{Name: payload.Name, DOB: dob}
			if err := gormDB.Create(&user).Error; err != nil {
				http.Error(w, "failed to create user", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(user)
			return

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
