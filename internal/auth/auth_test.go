package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "testsecret"

	token, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT error: %v", err)
	}

	gotID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("ValidateJWT error: %v", err)
	}
	if gotID != userID {
		t.Errorf("expected userID %v, got %v", userID, gotID)
	}
}

func TestExpiredJWT(t *testing.T) {
	userID := uuid.New()
	secret := "testsecret"

	token, err := MakeJWT(userID, secret, -time.Second)
	if err != nil {
		t.Fatalf("MakeJWT error: %v", err)
	}

	_, err = ValidateJWT(token, secret)
	if err == nil {
		t.Fatal("expected error for expired token, got nil")
	}
}

func TestWrongSecretJWT(t *testing.T) {
	userID := uuid.New()

	token, err := MakeJWT(userID, "correctsecret", time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT error: %v", err)
	}

	_, err = ValidateJWT(token, "wrongsecret")
	if err == nil {
		t.Fatal("expected error for wrong secret, got nil")
	}
}
