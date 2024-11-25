package helpers

func HashPassword(password string) string {
	hashed := ""
	for i, char := range password {
		if i%2 == 0 {
			hashed += string(char + 6)
		} else {
			hashed += string(char + 4)
		}
	}
	return hashed
}

func VerifyPassword(input, hashedPassword string) bool {
	original := ""
	for i, char := range hashedPassword {
		if i%2 == 0 {
			original += string(char - 6)
		} else {
			original += string(char - 4)
		}
	}
	return original == input
}
