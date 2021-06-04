// go run main.go
// curl -i http://localhost:8080

package main

import (
	// "fmt"
	"github.com/Countingup/tech-talks/bluesmith/helloswagger/restapi"
	"github.com/Countingup/tech-talks/bluesmith/helloswagger/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	//"github.com/go-openapi/swag"
)

func getGreeting(params operations.GetGreetingParams) middleware.Responder {

	// name := swag.StringValue(params.Name)
	// // name := params.Name
	// if name == "" {
	// 	name = "World"
	// }
	// return operations.NewGetGreetingOK().WithPayload(fmt.Sprintf("Hello %s!\n", name))

	return operations.NewGetGreetingOK().WithPayload("Hello World!\n")
}

func main() {
	swaggerSpec, _ := loads.Analyzed(restapi.SwaggerJSON, "")
	api := operations.NewHelloSwaggerAPI(swaggerSpec)

	// api.GetGreetingHandler = operations.GetGreetingHandlerFunc(getGreeting)

	server := restapi.NewServer(api)
	server.Port = 8080

	defer server.Shutdown()
	server.Serve()
}
