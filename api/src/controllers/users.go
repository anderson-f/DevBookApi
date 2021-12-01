package controllers

import "net/http"

// CreateUser insert a new user on database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

// GetUsers search all users on database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

// GetUsers search one user on database
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

// UpdateUser change a user information on database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

// DeleteUser delete a user information on database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}
