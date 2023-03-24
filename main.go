package main
import(
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Server")
	http.HandleFunc("/api",Handler)
	http.ListenAndServe(":8080",nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in the handler")
}

type APIResponse struct {
	
}