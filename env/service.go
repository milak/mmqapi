package env

import (

)

type Service interface {
	Start()
	GetName() string
	Stop()
}