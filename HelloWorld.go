package main

import "fmt"
import "http"

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello Word")
	}
	
func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}