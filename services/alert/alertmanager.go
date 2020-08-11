package alert

import (
	"errors"

	"github.com/rudderlabs/rudder-server/config"
)

var (
	alertProvider       string
	pagerDutyRoutingKey string
	instanceName        string
	victorOpsRoutingKey string
)

const (
	CRITICAL = "CRITICAL"
	RESOLVED = "RECOVERY"
)

func init() {
	loadConfig()
}

func loadConfig() {
	alertProvider = config.GetEnv("ALERT_PROVIDER", "victorops")
	pagerDutyRoutingKey = config.GetEnv("PG_ROUTING_KEY", "")
	instanceName = config.GetEnv("INSTANCE_ID", "")
	victorOpsRoutingKey = config.GetEnv("VICTOROPS_ROUTING_KEY", "")
}

type PayloadT struct {
	IncidentID  string
	Message     string
	MessageType string
}

// AlertManager interface
type AlertManager interface {
	Alert(payload PayloadT)
}

// New returns FileManager backed by configured privider
func New() (AlertManager, error) {
	switch alertProvider {
	case "victorops":
		return &VictorOps{
			routingKey:   victorOpsRoutingKey,
			instanceName: instanceName,
		}, nil
	case "pagerduty":
		return &PagerDuty{
			routingKey:   pagerDutyRoutingKey,
			instanceName: instanceName,
		}, nil
	}
	return nil, errors.New("No provider configured for Alert Manager")
}
