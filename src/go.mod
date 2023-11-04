module herb-go-app

go 1.20

replace modules => ./modules

require (
	github.com/herb-go/util v0.0.0-20230321163403-6a757b251968
	modules v0.0.0-00010101000000-000000000000
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/herb-go/datasource v0.0.0-20211122123843-a3546acd0d8d // indirect
	github.com/herb-go/datasource-drivers v0.0.0-20201011165914-7717acb90545 // indirect
	github.com/herb-go/herbconfig v0.0.0-20210201131438-44d8e331b703 // indirect
	github.com/herb-go/logger v0.0.0-20210115164802-8259bb9dcc90 // indirect
	github.com/herb-go/uniqueid v0.0.0-20210304163719-ac56f6357531 // indirect
	github.com/herb-go/worker v0.0.0-20210318151232-dbaf101c7d73 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
