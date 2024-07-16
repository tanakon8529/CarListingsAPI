package tests

import (
	"bytes"
	"daveslist-emdpcv/api/database"
	"daveslist-emdpcv/api/models"
	"daveslist-emdpcv/api/routes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testRouter *gin.Engine
var testToken string
var idAdminUser uint
var idRegisteredUser uint
var idCategoryPublic uint
var idCategoryPrivate uint

func TestMain(m *testing.M) {
	// Initialize test database
	var err error
	database.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}
	database.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Listing{},
		&models.Picture{},
		&models.Reply{},
		&models.PrivateMessage{},
	)

	// Create a test user admin
	adminUser := models.User{
		Username: "admin_test_01",
		Email:    "admin_test_01@example.com",
		Password: "admin_test_01",
		Role:     "admin",
	}
	database.DB.Create(&adminUser)
	idAdminUser = adminUser.ID

	// Create a test user registered_user
	registeredUser := models.User{
		Username: "registered_user_test_01",
		Email:    "registered_user_test_01@example.com",
		Password: "registered_user_test_01",
		Role:     "registered_user",
	}
	database.DB.Create(&registeredUser)
	idRegisteredUser = registeredUser.ID

	// Create a test category public
	categoryPublic := models.Category{
		Name:     "Test Public Category",
		IsPublic: true,
	}
	database.DB.Create(&categoryPublic)
	idCategoryPublic = categoryPublic.ID

	// Create a test category private
	categoryPrivate := models.Category{
		Name:     "Test Private Category",
		IsPublic: false,
	}
	database.DB.Create(&categoryPrivate)
	idCategoryPrivate = categoryPrivate.ID

	// Setup the router
	testRouter = gin.Default()
	routes.Setup(testRouter)

	// Generate and store the token
	testToken = generateTestToken()

	// Run the tests
	code := m.Run()

	// Cleanup
	os.Exit(code)
}

func generateTestToken() string {
	req, _ := http.NewRequest("POST", "/api/v1/auth", nil)
	req.Header.Set("client_id", "admin_test_01")
	req.Header.Set("secret_id", "admin_test_01")

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	return response["token"]
}

func TestCreateListingPublic(t *testing.T) {
	newListing := models.Listing{
		Title:       "Test listing",
		Content:     "Test content",
		CategoryID:  idCategoryPublic,
		CreatedByID: idRegisteredUser,
		IsPublic:    true,
	}
	jsonValue, _ := json.Marshal(newListing)
	req, _ := http.NewRequest("POST", "/api/v1/listings", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken)

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %v", w.Code)
	}
}

func TestCreateListingPrivate(t *testing.T) {
	newListing := models.Listing{
		Title:       "Test private listing",
		Content:     "Test private content",
		CategoryID:  idCategoryPrivate,
		CreatedByID: idRegisteredUser,
		IsPublic:    false,
	}
	jsonValue, _ := json.Marshal(newListing)
	req, _ := http.NewRequest("POST", "/api/v1/listings", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken)

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %v", w.Code)
	}
}

func TestGetListing(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/listings/1", nil)
	req.Header.Set("Authorization", "Bearer "+testToken)

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %v", w.Code)
	}
}

func TestUpdateListing(t *testing.T) {
	updatedListing := models.Listing{
		Title:   "Updated listing",
		Content: "Updated content",
	}
	jsonValue, _ := json.Marshal(updatedListing)
	req, _ := http.NewRequest("PUT", "/api/v1/listings/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken)

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %v", w.Code)
	}
}

func TestDeleteListing(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/api/v1/listings/1", nil)
	req.Header.Set("Authorization", "Bearer "+testToken)

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %v", w.Code)
	}
}

func TestListAllListingsForRegisteredUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/listings", nil)
	req.Header.Set("Authorization", "Bearer "+testToken)

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %v", w.Code)
	}
}

func TestCreateReply(t *testing.T) {
	newReply := models.Reply{
		Content:   "Test reply",
		ListingID: 1,
	}
	jsonValue, _ := json.Marshal(newReply)
	req, _ := http.NewRequest("POST", "/api/v1/listings/1/replies", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken)

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %v", w.Code)
	}
}

func TestGetReply(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/replies/1", nil)
	req.Header.Set("Authorization", "Bearer "+testToken)

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %v", w.Code)
	}
}

func TestUpdateReply(t *testing.T) {
	updatedReply := models.Reply{
		Content: "Updated reply",
	}
	jsonValue, _ := json.Marshal(updatedReply)
	req, _ := http.NewRequest("PUT", "/api/v1/replies/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken)

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %v", w.Code)
	}
}

func TestDeleteReply(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/api/v1/replies/1", nil)
	req.Header.Set("Authorization", "Bearer "+testToken)

	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, but got %v", w.Code)
	}
}
