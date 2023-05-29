package orm

import (
	"RudderMaster/models/auth"
	"RudderMaster/models/system"
)

var TableModes []interface{} = []interface{}{
	auth.User{}, system.Role{},
}
