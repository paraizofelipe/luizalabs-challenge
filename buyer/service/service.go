package service

import "github.com/paraizofelipe/luizalabs-challenge/buyer/repository"

type Service interface {
	repository.Reader
	repository.Writer
}
