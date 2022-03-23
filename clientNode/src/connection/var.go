package connection

import (
	"image/jpeg"

	"github.com/ranon-rat/camdeyeGO/clientNode/src/core"
)

var (
	cameras = []core.Camera{}

	defaultJPEGConfig = &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	}
)
