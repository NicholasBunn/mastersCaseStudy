package main

import (
	"context"
	"fmt"
	"testing"
 
	serverPB "github.com/NicholasBunn/mastersCaseStudy/services/authenticationService/proto/v1/generated"
)

func TestSave(t *testing.T) {
	fmt.Println("UNIMPLEMENTED")
}

func TestFind(t *testing.T) {
	fmt.Println("UNIMPLEMENTED")
}

func TestLoginAuth(t *testing.T) {
	t.Run("WHAT ARE WE TESTING", func(t *testing.T) {

		server := server{}

		request := serverPB.LoginAuthRequest {
			Username: "admin",
			Password: "myPassword",
		}

		response, err := server.LoginAuth(context.Background(), &request)

		if (err != nil) {
			fmt.Println("FAILED: ", err)
		}

		fmt.Println(response)
	})
}

