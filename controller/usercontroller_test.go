package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/tanaypatankar/go_user_mgmt/models"
)

var tanay = models.User{
	Id:      100,
	Name:    "Tanay",
	Email:   "tanay@platform9.com",
	Age:     21,
	Gender:  "M",
	Country: "India",
	Status:  "Active",
}

var shubhamc = models.User{
	Id:      101,
	Name:    "Shubham C",
	Email:   "schoudhari@platform9.com",
	Age:     21,
	Gender:  "M",
	Country: "India",
	Status:  "Active",
}

func TestCreateUser(t *testing.T) {

	responseBody, _ := json.Marshal(tanay)

	res, err := http.Post("http://localhost:9010/users/", "application/json", bytes.NewBuffer(responseBody))
	if err != nil {
		t.Fatal(err)
	}
	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var user_rec models.User
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &user_rec)

	if user_rec != tanay {
		t.Errorf("Returned unexpected body: got %v want %v", user_rec, tanay)
	}

}

func TestGetUserByID(t *testing.T) {

	res, err := http.Get("http://localhost:9010/users/100")
	if err != nil {
		t.Fatal(err)
	}

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body, _ := io.ReadAll(res.Body)
	var user_rec models.User
	json.Unmarshal(body, &user_rec)
	if user_rec != tanay {
		t.Errorf("Returned unexpected body: got %v want %v", user_rec, tanay)
	}

}

func TestGetUsers(t *testing.T) {
	// Add second record for testing
	responseBody, _ := json.Marshal(shubhamc)
	_, err := http.Post("http://localhost:9010/users/", "application/json", bytes.NewBuffer(responseBody))
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Get("http://localhost:9010/users/")
	if err != nil {
		t.Fatal(err)
	}

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	users := []models.User{tanay, shubhamc}
	body, _ := io.ReadAll(res.Body)
	var user_rec []models.User
	json.Unmarshal(body, &user_rec)
	if user_rec[0] != tanay && user_rec[1] != shubhamc {
		t.Errorf("Returned unexpected body: got %v want %v", user_rec, users)
	}

}

func TestUpdateUser(t *testing.T) {
	tanay = models.User{
		Id:      100,
		Name:    "Tanay Patankar",
		Email:   "tanay@platform9.com",
		Age:     22,
		Gender:  "M",
		Country: "India",
		Status:  "Inactive",
	}

	responseBody, _ := json.Marshal(tanay)

	req, _ := http.NewRequest("PUT", "http://localhost:9010/users/", bytes.NewBuffer(responseBody))

	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body, _ := io.ReadAll(res.Body)
	var user_rec models.User
	json.Unmarshal(body, &user_rec)
	if user_rec != tanay {
		t.Errorf("Returned unexpected body: got %v want %v", user_rec, tanay)
	}

}

func TestDeleteUser(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "http://localhost:9010/users/101", nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body, _ := io.ReadAll(res.Body)
	var user_rec models.User
	json.Unmarshal(body, &user_rec)
	if user_rec != shubhamc {
		t.Errorf("Returned unexpected body: got %v want %v", user_rec, shubhamc)
	}

	// Delete Tanay for cleanup
	req2, _ := http.NewRequest("DELETE", "http://localhost:9010/users/100", nil)
	http.DefaultClient.Do(req2)

}
