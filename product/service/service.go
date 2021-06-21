package service

import "github.com/paraizofelipe/luizalabs-challenge/product/repository"

type Service interface {
	repository.Reader
	repository.Writer
}
