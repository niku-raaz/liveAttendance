package main

import(
	"fmt"
	"net/http"
	"liveAt/utils"
	"liveAt/middleware"
	"github.com/go-chi/chi/v5"
)

func main(){

	fmt.Println("starting the server at port: 3000 ")
    
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.LogRequestTime)

	r.Get("/ping",pingHandler)

	r.Get("/health",healthHandler)

	r.Get("/version",func(w http.ResponseWriter, r *http.Request) {
		utils.Success(w,http.StatusOK,"v1.0.0")
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		utils.Error(w,http.StatusNotFound,"Route not found")
	})

	// always remember this blocks the main go-routine forever
	err:= http.ListenAndServe(":3000",r)
	if err!=nil{
		fmt.Println("Error starting server ",err)
	}

}

func pingHandler(w http.ResponseWriter, r *http.Request){

	utils.Success(w,http.StatusOK,"PONG")
}

func healthHandler(w http.ResponseWriter, r *http.Request){

	response:= map[string]any{
		"status": "ok",
	}

	utils.Success(w,http.StatusOK,response)
}

