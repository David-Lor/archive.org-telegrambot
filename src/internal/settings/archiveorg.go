package settings

import "time"

type ArchiveorgSettings struct {
	RawTimeout               string `yaml:"timeout"`
	ParsedTimeout            time.Duration
	ConsecutiveRequestsLimit int `yaml:"consecutive-requests-limit"`
}
