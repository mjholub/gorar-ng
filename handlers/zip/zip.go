package ziphandlers

type ZipHandler interface {
	Extract(destination string) error
}
