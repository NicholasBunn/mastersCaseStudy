package authentication

import (
	"fmt"
	"time"
	"strings"
	"testing"
)

func TestNewJWTManager(t *testing.T) {

	t.Run("Testing JWT Manager: Unit Test: New JWT Manager (Function Test)", func(t *testing.T) {
		/* This tests that the NewJWTManager function returns a struct with the expected fields. It ensures that the JWTManager struct and NewJWTManager function remain backwards-compatible during changes.
		*/

		result := NewJWTManager("secretKey", (15 * time.Minute))

		if (result.TokenDuration != (15 * time.Minute)) || (result.SecretKey != "secretKey") {
			t.Errorf("FAILED")
		}
		
		fmt.Println("PASSED")
	})
}

func TestGenerateToken(t *testing.T) {

	myManager := NewJWTManager("secretKey", (15 * time.Minute))

	myUser, err := CreateUser("testUser", "testPassword", "testRole")

	if (err != nil) {
		t.Errorf("Error from CreateUser function")
	}

	t.Run("Testing JWT Manager: Unit Test: Generate Manager (Function Test)", func(t *testing.T) {
		/* This tests that the correct header (encryption algorithm and token type) is/are being used for the authentication tokens.
		*/

		testToken, err := myManager.GenerateToken(myUser)

		if (err != nil) {
			t.Errorf("Failed to call 'Generate Token' function")
		}
	
		testTokenHeader := strings.Split(testToken, ".")[0]
	
		if (testTokenHeader != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9") {
			t.Errorf("FAILED")
		}
	
		fmt.Println("PASSED")
	})
}

func TestVerifyJWT(t *testing.T) {

	myManager := NewJWTManager("secretKey", (15 * time.Minute))

	myUser, err := CreateUser("testUser", "testPassword", "testRole")

	myToken, err := myManager.GenerateToken(myUser)

	if (err != nil) {
		t.Errorf("Error from CreateUser and/or GenerateToken function")
	}

	t.Run("Testing JWT Manager: Unit Test: Verify JWT (Function Test)", func(t *testing.T) {
		/* This tests thjat the VerifyJWT function is working correctly, and that it returns the correct claims when called.
		*/
		
		claims, err := myManager.VerifyJWT(myToken)

		if (err != nil) {
			t.Errorf("FAILED")
		} else if (claims.Username != "testUser") || (claims.Role != "testRole") {
			t.Errorf("FAILED")
		} else {
			fmt.Println("PASSED")
		}
	})
}