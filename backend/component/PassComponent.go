package component

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"math/rand"
)


func HashPassword(pass string) (string) {
	storePass := pass
	hash, err := bcrypt.GenerateFromPassword([]byte(storePass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
  }
	return string(hash);
}


func CheckPassword(Password string, hashedPassword string) (error) {
	err := bcrypt.CompareHashAndPassword([]byte(Password), []byte(hashedPassword))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func RandString(num int) (string) {
	const rs1Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, num)
	for i := range b {
			b[i] =  rs1Letters[int(rand.Int63()%int64(len(rs1Letters)))]
	}
	return string(b)
}