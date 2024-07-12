package helper

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "mySecureP@ssw0rd"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Error hashing password: %v", err)
	}

	if hashedPassword == "" {
		t.Error("Hashed password is empty")
	}

	// ตรวจสอบว่ารหัสผ่านที่แฮชแล้วสามารถถูกเปรียบเทียบได้อย่างถูกต้อง
	check := ComparePassword(hashedPassword, password)
	if !check {
		t.Error("Hashed password does not match the original password")
	}
}
