package temp

import (
	"go-salt/common/inject"
)

//type Aprinter interface {
//	inject.Obj.Common.UserAPI.LoadPolicy()
//}

func UserPolicy(id int)error {
	err := inject.Obj.Common.UserAPI.LoadPolicy(id)
	return err
}