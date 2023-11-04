package connect

import (
	"modules/app"

	"github.com/herb-go/util"
)

// ModuleName module name
const ModuleName = "900connect"

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		util.InitOrderByName(ModuleName)
		DefaultManager = NewManager()
		DefaultManager.Start(app.System)
		initSwitch()
	})
}
