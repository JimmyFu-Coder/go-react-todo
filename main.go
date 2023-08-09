package main

import (
	"fmt"
	"log"
	"net/http"
	"go-react-todo"
)

func main(){
	r:=router.Router()
	fmt.Println("123")
	http.ListenAndServe(":9000", r)
}