package helper

import "testing"

func TestGenerateToken(t *testing.T) {

	_, err := GenerateToken("teguh.afdilla138@gmail.com", "123456")

	if err != nil {
		panic("Gagal Generate Token")
	}
}

func TestValidateToken(t *testing.T) {
	_, err := ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhZmRpbGxhQGdtYWlsLmNvbSIsImlkIjoiNCJ9.GAEVYlZz5FdOGOfrA4Hl4ZMfQ_7ZCYPI2O_OrUkbGTU")

	if err != nil {
		panic("Gagal Validasi Token")
	}
}
