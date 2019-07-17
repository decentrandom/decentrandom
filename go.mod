module github.com/decentrandom/decentrandom

go 1.12

require (
	github.com/btcsuite/btcd v0.0.0-20190523000118-16327141da8c // indirect
	github.com/cosmos/cosmos-sdk v0.28.2-0.20190717162648-ae77f0080a72
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d // indirect
	github.com/gorilla/mux v1.7.2
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mattn/go-isatty v0.0.8 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/prometheus/common v0.4.1 // indirect
	github.com/prometheus/procfs v0.0.0-20190523193104-a7aeb8df3389 // indirect
	github.com/rakyll/statik v0.1.6
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.32.1
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f // indirect
	golang.org/x/sys v0.0.0-20190527104216-9cd6430ef91e // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/genproto v0.0.0-20190522204451-c2c4e71fbf69 // indirect
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5

replace github.com/cosmos/cosmos-sdk => github.com/cosmos/cosmos-sdk v0.28.2-0.20190622092459-7b5e6cee0787
