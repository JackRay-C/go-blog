package bo

import "blog/app/model/po"

type Post struct {
	Head *po.Head
	Repositories []*po.Repository
	Histories []*po.History
}

