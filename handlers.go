package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index Handler
func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte("Welcome!"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// FindAllUsers return all user in database
func FindAllUsers(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	users, err := NewUser().GetAllUsers()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//SaveUsers create a user
func SaveUsers(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	body := req.Body
	usr := NewUser()
	err := json.NewDecoder(body).Decode(usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer body.Close()
	err = usr.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// FindUserByID return a user
func FindUserByID(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := params.ByName("id")
	usr, err := NewUser().FindByID(id)
	log.Println(usr.ID)
	if usr.ID == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// UpdateUser update an existing user
func UpdateUser(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ID := params.ByName("id")
	if ID == "" {
		http.Error(w, "ID can not be empty", http.StatusBadRequest)
		return
	}
	var u User
	err := json.NewDecoder(req.Body).Decode(&u)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if u.ID == 0 {
		http.Error(w, "Please set ID in User information", http.StatusBadRequest)
		return
	}
	_, err = NewUser().FindByID(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = NewUser().Update(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteUser delete a user from database
func DeleteUser(w http.ResponseWriter, req *http.Request, prm httprouter.Params) {
	ID := prm.ByName("id")
	if ID == "" {
		http.Error(w, "Please send ID of user", http.StatusBadRequest)
		return
	}
	usr, err := NewUser().FindByID(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if usr.ID == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	err = usr.Delete(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
