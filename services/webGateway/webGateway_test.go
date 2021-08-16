package main

import (
	"context"
	"fmt"
	"testing"
 
	serverPB "github.com/NicholasBunn/mastersCaseStudy/protoFiles/go/webGateway/v1"
)

func TestLogin(t *testing.T) {

	r.Tun("Testing standard stuff.", func(t *testing.T) {
		/* This tests some of the core stuff, I'll document this once I've properly written the tests
		*/

		server := server{}

		request := serverPB.LoginRequest{
			Username: "admin",
			Password: "adminPassword",
		}

		response, err := server.Login(context.Background(), &request)

		if (err != nil) {
			fmt.Println("We made a whoopsy: ", err)
		}

		fmt.Println(response, err)
	})
}