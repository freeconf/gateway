module github.com/freeconf/gateway

go 1.18

require (
	github.com/freeconf/restconf v0.0.0-20221129143225-001b540b4e54
	github.com/freeconf/yang v0.0.0-20221129142318-a555424bf792
)

replace github.com/freeconf/restconf => "../restconf"