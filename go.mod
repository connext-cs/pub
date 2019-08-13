module github.com/connext-cs/pub

go 1.12

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20190108154635-47c0da630f72

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/astaxie/beego v1.12.0
	github.com/connext-cs/protocol v0.0.0-20190807032835-9fa672e79694
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/hashicorp/consul v1.4.4
	github.com/jmoiron/sqlx v1.2.0
	github.com/micro/go-api v0.7.0
	github.com/micro/go-config v1.1.0
	github.com/micro/go-micro v1.1.0
	go.etcd.io/etcd v3.3.13+incompatible
	gotest.tools v2.2.0+incompatible
)
