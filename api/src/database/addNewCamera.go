package database

import (
	"math/rand"

	"github.com/ranon-rat/camdeyeGO/api/src/core"
	"golang.org/x/crypto/bcrypt"
)

func AddNewNodeCamera(camera core.AddNewNodeCamera) (err error) {
	db := connect()
	defer db.Close()
	camera.PublicKey = rand.Intn(10)
	hashPrivateKey, err := bcrypt.GenerateFromPassword([]byte(camera.UnencryptedPrivateKey), camera.PublicKey)
	camera.EncryptedPrivateKey = string(hashPrivateKey)
	token := GenerateToken(camera.Name)
	camera.Token = EncryptSha256(token)
	// i need to see if the camera is on the database
	_, err = db.Exec(`INSERT INTO camera(
		name,
		owner_id, 
		public_key,
		token,
		private_key)
	 VALUES($1, $2, $3, $4, $5)`,
		camera.Name,
		camera.OwnerID,
		camera.PublicKey,
		camera.Token,
		camera.EncryptedPrivateKey)
	return
}
