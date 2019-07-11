module github.com/lijiang2014/cwl

go 1.12

//replace 	github.com/buchanae/cwl  => .

require (
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/kr/pretty v0.1.0
	github.com/lijiang2014/tugboat v0.0.0-20180327011757-94d752d436bd
	github.com/lijiang2014/yamlast v0.0.0-20160529193950-1f01fc418da0
	github.com/robertkrimen/otto v0.0.0-20180617131154-15f95af6e78d
	github.com/rs/xid v1.2.1
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.4
	github.com/stvp/assert v0.0.0-20170616060220-4bc16443988b // indirect
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
)

replace github.com/lijiang2014/tugboat => ../tugboat
