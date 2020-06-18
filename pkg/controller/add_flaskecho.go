package controller

import (
	"github.com/tk-flask-go/tk-flask-operator/pkg/controller/flaskecho"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, flaskecho.Add)
}
