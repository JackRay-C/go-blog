package service

import (
	"blog/app/pager"
)

type Service interface {
	SelectOne() error
	SelectAll(page *pager.Pager) error
	DeleteOne() error
	DeleteAll(ids []int) error
	CreateOne() error
	UpdateOne() error
	SaveOne() error
}
