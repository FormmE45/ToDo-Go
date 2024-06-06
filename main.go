package main

import ("github.com/gin-gonic/gin"
		"net/http")

type todo struct{
	ID 			string 		"json:'id'"
	Item 		string		"json:'title'"
	Completed 	bool 		"json:'completed'"
}
var todos =[]todo{
	{ID: "1",Item:"Clean room",Completed:false},
	{ID: "2",Item:"Read book",Completed: false},
	{ID: "3",Item:"Record Video",Completed: true},
}
func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK,todos)
}
func main(){
	router := gin.Default()
	router.GET("/todos",getTodos)
	router.Run("localhost:9090")
}