package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
        _ "github.com/cognoz/swag-ms/swagger/models"
        _ "github.com/cognoz/swag-ms/swagger/restapi"
        _ "github.com/cognoz/swag-ms/swagger/restapi/operations"
        _ "github.com/go-openapi/loads"
        _ "github.com/go-openapi/runtime/middleware"
        _ "github.com/go-openapi/swag"
)



type cliArgs struct {
	Port int `arg:"-p,help:port to listen to"`
}

var (
	args = &cliArgs{
		Port: 8080,
	}
)

func main () {
	arg.MustParse(args)
	fmt.Printf("port=%d\n", args.Port)
}
