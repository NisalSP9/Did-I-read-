package controllers

import (
	"github.com/NisalSP9/Did-I-read/commons"
	"github.com/NisalSP9/Did-I-read/dao"
)

func UserAuth(username, password string) (string, *commons.RequestError) {
	return dao.UserAuth(username, password)
}
