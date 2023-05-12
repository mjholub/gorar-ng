package ziphandlers

import (
	"github.com/154pinkchairs/gorar-ng/handlers/io"
)

type Handler interface {
	Extract(*iohandlers.Dir) error
}
