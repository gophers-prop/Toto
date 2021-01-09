package controllers

import (
	"Toto/forms"
	"Toto/models"
	"goji.io/pat"
	"net/http"
	"encoding/json"
	
	"github.com/sirupsen/logrus"
	
	

)

var userModel = new(models.UserModel)

// UserController : struct for user controller
type UserController struct{}


// Create : function to create user 
func (user *UserController) Create(w http.ResponseWriter, r *http.Request)  {
	var data forms.CreateUserCommand
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		logrus.WithFields(logrus.Fields{"User": "walrus",}).Error("This is an error message")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = userModel.Create(data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"User": "walrus",}).Error("This is an error message")
		var errorMsg forms.ResponseMsg
		errorMsg.Message = "User not created"		
		json.NewEncoder(w).Encode(errorMsg)
	}else {
		var msg forms.ResponseMsg
		msg.Message = "User created"		
		json.NewEncoder(w).Encode(msg)
		
	}
}

// Get : funcation to get user
func (user *UserController) Get(w http.ResponseWriter, r *http.Request) {
	id := pat.Param(r,"id")
	profile, err := userModel.Get(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		logrus.WithFields(logrus.Fields{"User": "walrus",}).Error("This is an error message")
		var errorMsg forms.ResponseMsg
		errorMsg.Message = "User not found"		
		json.NewEncoder(w).Encode(errorMsg)
	} else {		
		json.NewEncoder(w).Encode(profile)
	}
}


// Find : List all users
func (user *UserController) Find(w http.ResponseWriter, r *http.Request) {
	list, err := userModel.Find()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		logrus.WithFields(logrus.Fields{"User": "walrus",}).Error("This is an error message")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		json.NewEncoder(w).Encode(list)
	}
}

// Update : Function to update users based on ID
func (user *UserController) Update(w http.ResponseWriter, r *http.Request) {
	id := pat.Param(r,"id")
	data := forms.UpdateUserCommand{}
	err := json.NewDecoder(r.Body).Decode(&data)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		logrus.WithFields(logrus.Fields{"User": "walrus",}).Error("This is an error message")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = userModel.Update(id, data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"User": "walrus",}).Error("This is an error message")
		http.Error(w, "User not updated", http.StatusBadRequest)
		return
	} else {
		var msg forms.ResponseMsg
		msg.Message = "User Updated"		
		json.NewEncoder(w).Encode(msg)
	}
}
// Delete : Controller function to delete user based on ID
func (user *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := pat.Param(r,"id")
	w.Header().Set("Content-Type", "application/json")
	err := userModel.Delete(id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"User": "walrus",}).Error("This is an error message")
		var errorMsg forms.ResponseMsg
		errorMsg.Message = "User not found"		
		json.NewEncoder(w).Encode(errorMsg)
	}else {
		var msg forms.ResponseMsg
		msg.Message = "User deleted"		
		json.NewEncoder(w).Encode(msg)
	}
}
