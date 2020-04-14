package provisioner

// Cluster interface for deploying K8s clusters
type Cluster interface {
	CreateBootstrap() error
	InstallControlPlane() error
	CreatePermanent() error
	PivotControlPlane() error
	InstallAddons() error
	RequiredCommands() []string
	Events() chan interface{}
}
