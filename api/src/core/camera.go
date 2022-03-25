package core

/*
create table camera(
    id SERIAL,
    name varchar(254) NOT NULL,       -- crypt/rand.
    owner_id integer not null,        -- you are the one that create a new camera id.
    public_key integer not null,      -- generate a new one when you log in with crypt/rand.
    token varchar(64) not null,       -- generate a new one when you log  .
    private_key varchar(64) not null, -- sha256(password+public_key).
    -- how did you generate all this stuff, well in the case that you want to use it
    -- the only thing that you need to use is log in, in our page and then go to the part of camera
    -- and then click in "new camera", and the it will appear you a token that you need to put in the node administration page
    -- "http://camdeye.local".

);


*/

// camera /new-camera route
type AddNewNodeCamera struct {
	Name                  string `json:"name"`        // Name of the camera, preferably unique
	UnencryptedPrivateKey string `json:"private_key"` // this is the camera private key
	OwnerID               string `json:"owner_id"`    // just for knowing who is
	OwnerToken            string `json:"owner_token"` // use it for log in
	EncryptedPrivateKey   string
	PublicKey             int
	Token                 string
}

// yes i need to this /ice-offer route
type CameraIceOffer struct {
	IP       string
	Token    string `json:"token"` // with this i just avoid to login again(and i will encrypt it )
	IceOffer string `json:"ice"`   // with this you can connect to the camera node(just send the string)
} // how i can delete if you cant connect ?

// login route
type CameraNodeLoginPost struct {
	ID       string `json:"id"`       // the id of the node camera
	Username string `json:"username"` //remember to save it
	Password string `json:"password"` //remember to save it

}
type CameraNodeLoginResponse struct {
	Token string `json:"token"` // remember to use it and save it
}
