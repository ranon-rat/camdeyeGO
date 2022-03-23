package core

type CameraIceOffer struct {
	Token    string `json:"token"` // with this i just avoid to login
	IceOffer string `json:"ice"`   // with this you can connect to the camera node
}

type CameraLogin struct {
	Password string `json:"password"`
	Email    string `json:"Email"`
}
