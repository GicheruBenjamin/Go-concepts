package totti

/*
A user and a team 
State management and lifecycle of each.
User - id , firstName, lastName, email, password , role (leader, member)
Team - id, name, description, ownerId, members 
*/


type Email string

func validateEmail(email Email) bool {
	return IsValidEmail(string(email))
}

