module github.com/symcn/api

go 1.16

require (
	github.com/prometheus/client_golang v1.9.0
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
	sigs.k8s.io/controller-runtime v0.8.2
)

replace k8s.io/client-go => k8s.io/client-go v0.20.4
