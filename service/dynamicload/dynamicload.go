package dynamicload

import (
	"github.com/jasonrichardsmith/mongolar/wrapper"
)

type DynamicLoad{
	Target string
	Controller string
	Id	string
	Template string
}

func dynamic(d DynamicLoad, w *wrapper.Wrapper) {
	w.SetPayload("dynamic", d)
}