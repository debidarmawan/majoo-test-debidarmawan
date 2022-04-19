package main

import (
	"fmt"
	"majoo-test-debidarmawan/config"
	"majoo-test-debidarmawan/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

// @Title Majoo Technical Assessment Swagger - Debi Darmawan
// @Description This is API Documentation of Majoo Technical Assessment
// @Author Debi Darmawan
// @Contact.email debidarmawan1998@gmail.com
// @BasePath /v1
func main() {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	if err := godotenv.Load(exPath + "/.env"); err != nil {
		fmt.Println("MAIN FILE, Message : .env file is not loaded properly")
		os.Exit(2)
	}

	dbConn := config.InitConnection(os.Getenv("DB_URL"), 5)
	config.Migrate(dbConn.MajooDB)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		http.ServeHTTP(dbConn)
	}()
	waitGroup.Wait()
}
