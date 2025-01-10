package images

import "math/rand"

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetLastXChars(s string, length int) (ext string) {
	ext = string(s[len(s)-length:])
	return ext
}

func CheckExtention(ext string) bool {
	extensions := []string{"png", "webm", "jpg", "jpeg"}

	for _, v := range extensions {
		if v == ext {
			return true
		}
	}
	return false
}
