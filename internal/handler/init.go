package handler

import (
	"github.com/rafli024/mytodo-app/internal/contract"
)

var app *contract.App

func Init(a *contract.App) {
	app = a
}
