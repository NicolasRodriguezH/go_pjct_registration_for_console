package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	Id       int
	Username string
	Email    string
	Age      int
}

var Id int
var users map[int]User

func crearUsuario() {

	clearConsole()
	fmt.Print("Ingresa un valor para username: ")
	Username := obtenerData()
	fmt.Print("Ingresa un valor para email: ")
	Email := obtenerData()
	fmt.Print("Ingresa un valor para age: ")
	Age, err := strconv.Atoi(obtenerData())
	if err != nil {
		panic("No fue posible convertir a entero")
	}

	Id++
	user := User{Id, Username, Email, Age}

	users[Id] = user

	fmt.Println(">>> Usuario creado exitosamente!\n")
}

func listarUsuario() {
	clearConsole()
	for k, v := range users {
		fmt.Println(k, "-", v.Username, v.Email, v.Age)
	}

	fmt.Println("\n")
}

func actualizarUsuario() {
	clearConsole()

	/* Conseguido! */
	fmt.Print("Ingrese el Id del usuario que desea actualizar: ")

	Id, err := strconv.Atoi(obtenerData())
	if err != nil {
		panic("No fue posible convertir a entero")
	}

	if _, ok := users[Id]; ok {
		fmt.Print("Ingresa un valor para username: ")
		Username := obtenerData()
		fmt.Print("Ingresa un valor para email: ")
		Email := obtenerData()
		fmt.Print("Ingresa un valor para age: ")
		Age, err := strconv.Atoi(obtenerData())
		if err != nil {
			panic("No fue posible convertir a entero")
		}

		user := User{Id, Username, Email, Age}
		users[Id] = user
	}

	fmt.Println(">>> Usuario actualizado exitosamente!\n")
}

func eliminarUsuario() {
	clearConsole()
	fmt.Print("Ingresa el Id del usuario a eliminar: ")

	Id, err := strconv.Atoi(obtenerData())
	if err != nil {
		panic("No fue posible convertir a entero")
	}

	if _, ok := users[Id]; ok {
		delete(users, Id)
	}

	fmt.Println(">>> Usuario eliminado exitosamente!\n")
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	cmd.Run()
}

func obtenerData() string {
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor")
	} else {
		return strings.TrimSuffix(option, "\n")
	}

}

func main() {

	users = make(map[int]User)

	var option string

	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("A", "Crear")
		fmt.Println("B", "Listar")
		fmt.Println("C", "Actualizar")
		fmt.Println("D", "Eliminar")

		fmt.Print("Ingresa una opcion valida: ")

		option = obtenerData()

		if option == "quit" || option == "q" {
			break
		}

		switch option {
		case "a", "crear":
			crearUsuario()
		case "b", "listar":
			listarUsuario()
		case "c", "actualizar":
			actualizarUsuario()
		case "d", "eliminar":
			eliminarUsuario()
		default:
			fmt.Println("Ninguno coincide")
		}
	}
	fmt.Println("El programa ha finalziado")
}
