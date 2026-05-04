package utilities

// добавить хеширование паролей
func ComparePasswords(appPassword, userPassword string) bool {
	return appPassword == userPassword
}
