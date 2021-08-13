package authentication

import (
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {

	t.Run("Testing Users: Unit Test: Create User (Function Test)", func(t *testing.T) {
		/* This tests that the CreateUser function returns a user object (struct) with the expected fields. It ensures that the User struct and CreateUser function remain backwards-compatible during changes. NOTE: this test ensures that the password is encrpyted, but does not verify the encryption or the algorithm used.
		*/

		result, err := CreateUser("testUser", "testPassword", "testRole")
		fmt.Println(result)
		
		if (err != nil) {
			t.Errorf("Error from CreateUser function")
		} else if (result.Username != "testUser") || (result.HashedPassword == "testPassword") || (result.Role != "testRole") {
			t.Errorf("FAILED")
		} else {
			fmt.Println("PASSED")
		}
	})
}

func TestCheckPassword(t *testing.T) {

	myUser, err := CreateUser("testUser", "testPassword", "testRole")

	if (err != nil) {
		t.Errorf("Error from CreateUser function")
	}

	t.Run("Testing Users: Unit Test: Truth Testing Check Password (Function Test)", func(t *testing.T) {
		/* This tests that the Check Password function succesfully verifies a hashed password with the original password. It serves to validate that the create password function succesfully encrypts the password with the correct algorithm.
		*/

		result := myUser.CheckPassword("testPassword")

		if (result != true) {
			t.Errorf("FAILED")
		}

		fmt.Println("PASSED")
	})

	t.Run("Testing Users: Unit Test: False Testing Check Password (Function Test)", func(t *testing.T) {
		/* This tests that the Check Password function succesfully verifies a hashed password with the original password. It serves to validate that the create password function succesfully encrypts the password with the correct algorithm.
		*/

		result := myUser.CheckPassword("falseTestPassword")

		if (result != false) {
			t.Errorf("FAILED")
		}
		
		fmt.Println("PASSED")
	})
}