package main


import(
	"fmt"
	"io"
	"net/http"
)

func main(){

	// creating a new http client
	client := &http.Client{}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	// resp, err := client.Get("https://swapi.dev/api/people/1/")

	if err!= nil{
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	// reading and priting the response body
	body, err := io.ReadAll(resp.Body)
	if err!= nil{
		fmt.Println("Error making response body:", err)
		return
	}

	fmt.Println(string(body))
	
}