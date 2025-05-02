package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// ---------------------- 1. Object Creation (Struct Definition) ----------------------

// Define a custom object type: User
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" validate:"required,alphanum"`
	Email     string    `json:"email" validate:"required,email"`
	Age       int       `json:"age" validate:"gte=0,lte=150"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ---------------------- 2. Object Creation (Instantiation) ----------------------

func createUser(username, email string, age int) *User {
	now := time.Now()
	return &User{
		Username:  username,
		Email:     email,
		Age:       age,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// ---------------------- 3. Conversion (to DTO) ----------------------

// Define a Data Transfer Object (DTO) for a simplified user representation
type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func convertToDTO(user *User) *UserDTO {
	return &UserDTO{
		ID:       user.ID,
		Username: user.Username,
	}
}

// ---------------------- 4. Transformation (Adding a Greeting) ----------------------

type UserWithGreeting struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Greeting string `json:"greeting"`
}

func transformWithGreeting(dto *UserDTO) *UserWithGreeting {
	return &UserWithGreeting{
		ID:       dto.ID,
		Username: dto.Username,
		Greeting: fmt.Sprintf("Hello, %s!", dto.Username),
	}
}

// ---------------------- 5. Destructuring (Selecting Specific Fields) ----------------------

func extractUsernameAndEmail(user *User) (string, string) {
	return user.Username, user.Email
}

// ---------------------- 6. Validation ----------------------

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func validateUser(user *User) map[string]string {
	errors := make(map[string]string)
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := strings.ToLower(err.Field())
			tag := err.Tag()
			switch tag {
			case "required":
				errors[fieldName] = "is required"
			case "alphanum":
				errors[fieldName] = "must be alphanumeric"
			case "email":
				errors[fieldName] = "must be a valid email address"
			case "gte":
				errors[fieldName] = fmt.Sprintf("must be greater than or equal to %s", err.Param())
			case "lte":
				errors[fieldName] = fmt.Sprintf("must be less than or equal to %s", err.Param())
			default:
				errors[fieldName] = fmt.Sprintf("failed validation on tag %s", tag)
			}
		}
	}
	return errors
}

// ---------------------- 7. Serialization (to JSON) ----------------------

func serializeUserToJSON(user interface{}) (string, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return "", fmt.Errorf("failed to marshal to JSON: %w", err)
	}
	return string(jsonData), nil
}

// ---------------------- 8. Deserialization (from JSON) ----------------------

func deserializeJSONToUser(jsonData string) (*User, error) {
	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return &user, nil
}

func main() {
	// 1 & 2. Object Creation
	user := createUser("testuser123", "test@example.com", 25)
	fmt.Printf("Created User: %+v\n", user)

	// 6. Validation
	validationErrors := validateUser(user)
	if len(validationErrors) > 0 {
		fmt.Println("Validation Errors:", validationErrors)
	} else {
		fmt.Println("User is valid.")
	}

	invalidUser := createUser("invalid-user!", "not-an-email", -5)
	invalidErrors := validateUser(invalidUser)
	fmt.Println("Invalid User Validation Errors:", invalidErrors)

	// 3. Conversion to DTO
	userDTO := convertToDTO(user)
	fmt.Printf("Converted to DTO: %+v\n", userDTO)

	// 4. Transformation
	userWithGreeting := transformWithGreeting(userDTO)
	fmt.Printf("Transformed with Greeting: %+v\n", userWithGreeting)

	// 5. Destructuring
	username, email := extractUsernameAndEmail(user)
	fmt.Printf("Destructured - Username: %s, Email: %s\n", username, email)

	// 7. Serialization
	jsonString, err := serializeUserToJSON(user)
	if err != nil {
		fmt.Println("Serialization Error:", err)
	} else {
		fmt.Println("Serialized JSON:", jsonString)
	}

	// 8. Deserialization
	newUser, err := deserializeJSONToUser(jsonString)
	if err != nil {
		fmt.Println("Deserialization Error:", err)
	} else {
		fmt.Printf("Deserialized User: %+v\n", newUser)
	}
}