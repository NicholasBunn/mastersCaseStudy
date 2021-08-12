package authentication

import (
	"fmt"
	"time"
	"testing"
)

func TestNewJWTManager(t *testing.T) {

}

func TestGenerateManager(t *testing.T) {

}

func TestVerifyJWT(t *testing.T) {
	myManager := newJWTManager("secretKey", (15 * time.Minute))

	fmt.Println(myManager)
}

// func TestLoginAuth(t *testing.T) {
// 	t.Run("WHAT ARE WE TESTING", func(t *testing.T) {

// 		server := server{}

// 		request := serverPB.LoginAuthRequest {
// 			Username: "admin",
// 			Password: "myPassword",
// 		}

// 		response, err := server.LoginAuth(context.Background(), &request)

// 		if err != nil {
// 			fmt.Println("FAILED: ", err)
// 		}

// 		fmt.Println(response)
// 	})
// }

