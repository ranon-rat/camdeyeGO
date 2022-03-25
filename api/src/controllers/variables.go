package controllers

import "github.com/ranon-rat/camdeyeGO/api/src/core"

var (
	camerasNodesIceOffer = make(map[string]core.CameraIceOffer)
	camerasIPs           = make(map[string]string) // i dont really want to store another ice offer for the same node
)
