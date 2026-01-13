package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/iamskyy666/golang-job-portal-api/internal/models"
)

// GenerateRandomPassword generates a random password of given length
func GenerateRandomPassword(charCount int) string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyz"
	var password strings.Builder

	rand.Seed(time.Now().UnixNano())
	password.Grow(charCount)

	for range charCount {
		password.WriteByte(charset[rand.Intn(len(charset))])
	}

	return password.String()
}

func ValidatePasswordStrength(password string)(bool, []string){
	validation:=models.PasswordValidation{
	MinLength: 8,
	HasUpper: true,
	HasLower: true,
	HasNumber: true,
	HasSpecial: true,
	}

	var validationErrors []string

	// Check min. length
	if len(password)<validation.MinLength{
		validationErrors = append(validationErrors, fmt.Sprintf("Password must be at least %d characters long!", validation.MinLength))
	}

	// Check for uppercase letters
	if validation.HasUpper && !strings.ContainsAny(password,"ABCDEFGHIJKLMNOPQRSTUVWXYZ"){
		validationErrors = append(validationErrors, "Password must have at least one uppercase letter!")
	}

	// Check for lowercase letters
	if validation.HasLower && !strings.ContainsAny(password,"abcdefghijklmnopqrstuvwxyz"){
		validationErrors = append(validationErrors, "Password must have at least one lowercase letter!")
	}
	

	// Check for numbers
	if validation.HasNumber && !strings.ContainsAny(password,"0123456789"){
		validationErrors = append(validationErrors, "Password must have at least one number!")
	}

	// Check for special chars.
	if validation.HasSpecial{
		hasSpecial:=false
		for _,char:=range password {
			if unicode.IsPunct(char) || unicode.IsSymbol(char){
				hasSpecial = true
				break
			}
		}
		if !hasSpecial{
			validationErrors = append(validationErrors, "Password must have at least one special character!")
		}
	}

	return len(validationErrors) == 0,validationErrors

}