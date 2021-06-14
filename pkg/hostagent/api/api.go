package api

import (
	"time"
)

type Status struct {
	Running  bool `json:"running,omitempty"`
	Degraded bool `json:"degraded,omitempty"`
	Aborted  bool `json:"aborted,omitempty"`

	Errors []string `json:"errors,omitempty"`

	SSHLocalPort int `json:"sshLocalPort,omitempty"`
}

type Event struct {
	Time   time.Time `json:"time,omitempty"`
	Status Status    `json:"status,omitempty"`
}
