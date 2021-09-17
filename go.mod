module github.com/adodon2go/nodeexporter

go 1.15

replace github.com/anuvu/zot => /home/midgard/work/adodon2go/zot

replace github.com/aquasecurity/trivy => github.com/anuvu/trivy v0.9.2-0.20200731014147-c5f97b59c172

require (
	github.com/anuvu/zot v0.0.0-00010101000000-000000000000 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/rs/zerolog v1.22.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
)
