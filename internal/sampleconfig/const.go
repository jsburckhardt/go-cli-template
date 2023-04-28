// Package sampleconfig is a helper for configs in sample
package sampleconfig

// ConfigName is the struct that enforces config names to be set in code
type ConfigName string

//revive:disable
const (
	APP_MODE           ConfigName = "APP_MODE"
	DEVICE_HOSTNAME    ConfigName = "DEVICE_HOSTNAME"
	PI_REGISTRY_PORT   ConfigName = "PI_REGISTRY_PORT"
	PI_DOCKER_REGISTRY ConfigName = "PI_DOCKER_REGISTRY"
)
