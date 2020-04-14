package provisioner

// Cluster interface for deploying K8s clusters
type Cluster interface {
	CreateBootstrap() error
	InstallCAPV() error
	CreatePermanent() error
	CAPvPivot() error
	InstallAddons() error
	RequiredCommands() []string
	Events() chan interface{}
}
