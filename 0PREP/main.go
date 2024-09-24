package main
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet",    func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	http.ListenAndServe("localhost:8000", nil)

}


// TERMINAL_1:
// go build main.go -o m
// run server: ./m         => it will start web server
// TERMINAL_2:
// curl http://localhost:8000/greet  <enter>
// OUT2: Hello World!