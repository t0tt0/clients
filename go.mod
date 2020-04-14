module github.com/HyperService-Consortium/go-ves

go 1.13

replace (
	github.com/HyperService-Consortium/go-hexutil => github.com/HyperService-Consortium/go-hexutil v1.0.1
	github.com/HyperService-Consortium/go-uip => github.com/HyperService-Consortium/go-uip v0.0.0-20200408075657-d2425491ab24
)

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/DeanThompson/ginpprof v0.0.0-20190408063150-3be636683586
	github.com/HyperService-Consortium/NSB v0.7.4-0.20200414172424-b8a34f8e59d5

	github.com/HyperService-Consortium/go-ethabi v0.9.1
	github.com/HyperService-Consortium/go-mpt v1.1.1-0.20190903152503-b63ec8d75bd4
	github.com/HyperService-Consortium/go-rlp v1.0.1-0.20190903144357-b5693c05a6b8
	github.com/HyperService-Consortium/go-uip v0.0.0-20200408075657-d2425491ab24
	github.com/Joker/jade v1.0.0 // indirect
	github.com/Myriad-Dreamin/artisan v0.8.1-0.20200204040619-76955d01aad9
	github.com/Myriad-Dreamin/dorm v0.0.0-20191205101004-33dbc61bb34e
	github.com/Myriad-Dreamin/functional-go v0.0.0-20191104092509-c2b7f373dd31
	github.com/Myriad-Dreamin/gin-middleware v0.0.0-20191222015112-6e9c660fff45
	github.com/Myriad-Dreamin/go-magic-package v0.0.0-20191102120213-a407f918fece
	github.com/Myriad-Dreamin/go-model-traits v0.0.0-20191209220601-85cd28b274b0
	github.com/Myriad-Dreamin/go-parse-package v1.0.1
	github.com/Myriad-Dreamin/gvm v1.0.2
	github.com/Myriad-Dreamin/minimum-lib v0.0.0-20200117225041-ec905257618d
	github.com/Myriad-Dreamin/mydrest v1.0.1
	github.com/Myriad-Dreamin/screenrus v1.0.0
	github.com/Shopify/goreferrer v0.0.0-20181106222321-ec9c9a553398 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/aymerick/raymond v2.0.2+incompatible // indirect
	github.com/casbin/casbin/v2 v2.1.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flosch/pongo2 v0.0.0-20190707114632-bbf5a6c351f4 // indirect
	github.com/gavv/monotime v0.0.0-20190418164738-30dba4353424 // indirect
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.5.0
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.0
	github.com/golang/protobuf v1.3.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/go-querystring v1.0.0
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/gorilla/websocket v1.4.0
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/imroc/req v0.3.0
	github.com/iris-contrib/blackfriday v2.0.0+incompatible // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/iris-contrib/go.uuid v2.0.0+incompatible // indirect
	github.com/iris-contrib/httpexpect v0.0.0-20180314041918-ebe99fcebbce // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.10 // indirect
	github.com/kataras/iris v11.1.1+incompatible
	github.com/klauspost/compress v1.10.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/pelletier/go-toml v1.6.0
	github.com/ryanuber/columnize v2.1.0+incompatible // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/stretchr/testify v1.5.1
	github.com/syndtr/goleveldb v1.0.1-0.20190318030020-c3a204f8e965
	github.com/tidwall/gjson v1.4.0
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	go.uber.org/atomic v1.5.0
	go.uber.org/zap v1.13.0
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	golang.org/x/mod v0.2.0 // indirect
	golang.org/x/net v0.0.0-20191028085509-fe3aa8a45271
	golang.org/x/sys v0.0.0-20200212091648-12a6c2dcc1e4 // indirect
	golang.org/x/tools v0.0.0-20200214225126-5916a50871fb
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543 // indirect
	google.golang.org/grpc v1.23.0
	gopkg.in/yaml.v2 v2.2.8
)
