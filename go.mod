module github.com/codyja/ftdata

go 1.19

replace github.com/codyja/focustronic/api => ../focustronic/api

require (
	github.com/codyja/focustronic/api v0.0.0-20210313185533-4072c4b9c8e8
	github.com/fatih/color v1.14.1
	github.com/rodaine/table v1.1.0
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	golang.org/x/sys v0.3.0 // indirect
)
