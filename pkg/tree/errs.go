package tree

import "errors"

var errTreeIsNil = errors.New("tree is nil")
var errNodeAlreadyExist = errors.New("this node index already exists")
var errValueNotFound = errors.New("no value with such index")
