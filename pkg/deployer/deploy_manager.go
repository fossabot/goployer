package deployer

import (
	"github.com/DevopsArtFactory/goployer/pkg/builder"
)

type DeployManager interface {
	GetStackName() string
	Deploy(config builder.Config)
	HealthChecking(config builder.Config) map[string]bool
	FinishAdditionalWork(config builder.Config) error
	CleanPreviousVersion(config builder.Config) error
	TriggerLifecycleCallbacks(config builder.Config) error
	TerminateChecking(config builder.Config) map[string]bool
}
