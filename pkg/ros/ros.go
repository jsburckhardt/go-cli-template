// Package ros is used to manage and configure ros packages
package ros

import (
	"sample/internal/directory"
	"sample/internal/sampleconfig"
	"sample/pkg/adapter"
)

// Service interacts with the flight computer and returns stuff
type Service struct {
	logger *adapter.Logger
	config *sampleconfig.Config
}

// NewRosService returns a new instance of ros service
func NewRosService(logger *adapter.Logger) *Service {
	return &Service{
		logger: logger,
		config: sampleconfig.GetConfig(),
	}
}

// Build builds a ros package based on the parameters,

// Run runs given script to run the ros package
func (s *Service) Run(pkgName string, wsPath string, nodeName string) error {
	// validate if path is relative or absolute
	wsPath, err := directory.CleanPathArgument(wsPath)
	if err != nil {
		return err
	}
	// retrieving scripts
	s.logger.Infof("validating ros environment for package %s in %s and nodename: ", pkgName, wsPath, nodeName)

	return nil
}
