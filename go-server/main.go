package main

import{
	"fmt"
	"log"
	"net/http"
}
func formHandler(w http.ResponseWriter, r *http.Request){
	if err :- r.ParseForm(); err != nil {
		fmt.printf(w, "Parseform() err: %v",err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n",name)
	fmt.Fprintf(w, "Address =%s\n",address)
}

func helloHandler(w http.ReponseWriter, r *http.Request){
	if r.URL.path != "/hello" {
		http.Error(w, "404 not found", http.statusNotFound)
		return
	}
if r.method != "GET"{
	http.Error(w,"method is not supported", http.StatusNotFound)
	return
}
}
fmt.Fprintf(w,"hello!")



func main(){
	fileserver := http.fileserver(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.Handlefunc("/form", formhandler)
	http.Handlefunc("./hello,", hellohandler) 

	fmt.println("port server 8000")
	if err :- http.ListenAndServer(":8000" , nil); err !=nil{
		log.Fatal(err)
	}
}