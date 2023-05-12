package rar

type RarHandler interface {
	Extract(destination string) error
}
