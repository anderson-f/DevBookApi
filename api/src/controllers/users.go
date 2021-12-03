package controllers

import "net/http"

// CreateUser insert a new user on database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

// GetUsers search all users on database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuários"))
}

// GetUsers search one user on database
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um único usuário"))
}

// UpdateUser change a user information on database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

// DeleteUser delete a user information on database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
