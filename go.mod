module stu

go 1.13

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net => github.com/golang/net v0.0.0-20190628185345-da137c7871d7
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190712062909-fae7ac547cb7
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190719005602-e377ae9d6386
)

require (
	github.com/mattn/go-gtk v0.0.0-20191030024613-af2e013261f5
	github.com/mattn/go-pointer v0.0.0-20190911064623-a0a44394634f // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
)
