package main

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	serverPB "github.com/NicholasBunn/mastersCaseStudy/services/authenticationService/proto/v1/generated"
)

func TestSave(t *testing.T) {
	fmt.Println("UNIMPLEMENTED")
}

func TestFind(t *testing.T) {
	fmt.Println("UNIMPLEMENTED")
}

func TestLoginAuth(t *testing.T) {

	t.Run("Testing Authentication Service: Integration Test: Existing User Test for Login Auth (Service Test)", func(t *testing.T) {
		/* This tests that the core functionality of the login auth service, verifying that it logs users in successfully.
		*/

		server := server{}

		request := serverPB.LoginAuthRequest {
			Username: "admin",
			Password: "myPassword",
		}

		response, err := server.LoginAuth(context.Background(), &request)

		if (err != nil) {
			t.Errorf("Failed to call 'LoginAuth' function")
		} else if (response.Permissions != "admin") || (strings.Split(response.AccessToken, ".")[0] != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9") {
			t.Errorf("FAILED")
		} else {
			fmt.Println("PASSED")
		}
	})

	t.Run("Testing Authentication Service: Integration Test: Non-Existing User Test for Login Auth (Service Test)", func(t *testing.T) {
		/* This tests that the core functionality of the login auth service, verifying that it returns guest permissions for non-existing user.
		*/

		server := server{}

		request := serverPB.LoginAuthRequest {
			Username: "nonExistingUser",
			Password: "myPassword",
		}

		response, err := server.LoginAuth(context.Background(), &request)

		errorStatus, _ := status.FromError(err)
		if (response != nil) {
			t.Errorf("Failed to call 'LoginAuth' function")
		} else if (errorStatus.Code() != codes.NotFound) {
			t.Errorf("FAILED")
		} else {
			fmt.Println("PASSED")
		}
	})
}

