module github.com/lijiang2014/cwl

go 1.12

require (
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf
	//github.com/buchanae/cwl v0.0.0-20181219185852-c4d1d10d5f38
	github.com/buchanae/tugboat v0.0.0-20180327011757-94d752d436bd
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/kr/pretty v0.1.0
	github.com/robertkrimen/otto v0.0.0-20180617131154-15f95af6e78d
	github.com/rs/xid v1.2.1
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.4
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
)

replace github.com/commondream/yamlast => github.com/buchanae/yamlast latest

replace 	github.com/buchanae/cwl  => .