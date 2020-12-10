module github.com/kamva/hexa-tuner

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Kamva/mgm/v3 v3.0.0 // indirect
	github.com/contribsys/faktory v1.3.0-1
	github.com/contribsys/faktory_worker_go v1.4.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/kamva/gutil v0.0.0-20201117080033-8568e68b386e
	github.com/kamva/hexa v0.0.0-20201210123552-58be39a23ee4
	github.com/kamva/hexa-arranger v0.0.0-20201210124912-a257459af177
	github.com/kamva/hexa-echo v0.0.0-20201210123726-fbd3fb69a712
	github.com/kamva/hexa-event v0.0.0-20201210152034-8ab4f2cfc935
	github.com/kamva/hexa-job v0.0.0-20201210160854-95e4ad9a432f
	github.com/kamva/hexa-rpc v0.0.0-20201210171118-d2153bd711e3
	github.com/kamva/tracer v0.0.0-20201115122932-ea39052d56cd
	github.com/labstack/echo/v4 v4.1.17
	github.com/nicksnyder/go-i18n/v2 v2.0.3
	github.com/spf13/viper v1.6.2
	go.uber.org/cadence v0.13.4 // indirect
	golang.org/x/text v0.3.4
	google.golang.org/grpc v1.33.1
)
