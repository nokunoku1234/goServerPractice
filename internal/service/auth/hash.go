package auth

import "golang.org/x/crypto/bcrypt"

// PasswordEncrypt 暗号化処理（hash化）
func PasswordEncrypt(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPassword), err
}

// CheckHashPassword 暗号化パスワードと元のパスワードを比較検証
func CheckHashPassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
