package main

import(
	"fmt"
	"net/http"
	"liveAt/utils"
)

func main(){

	fmt.Println("starting the server at port: 3000 ")
    
	mux:= http.NewServeMux()

	mux.HandleFunc("/ping",pingHandler)
	mux.HandleFunc("/health",healthHandler)

	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		utils.Error(w,http.StatusNotFound,"Route not found")
	})

	// always remember this blocks the main go-routine forever
	err:= http.ListenAndServe(":3000",mux)
	if err!=nil{
		fmt.Println("Error starting server ",err)
	}

}

func pingHandler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	utils.Success(w,http.StatusOK,"PONG")
}
func healthHandler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet {
		utils.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	response:= map[string]any{
		"status": "ok",
	}

	utils.Success(w,http.StatusOK,response)
}

