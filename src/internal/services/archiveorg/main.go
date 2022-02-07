package archiveorg

import (
	"github.com/David-Lor/archive.org-telegrambot/internal/settings"
	"github.com/gammazero/workerpool"
	"net/http"
)

type Client struct {
	client   *http.Client
	workPool *workerpool.WorkerPool
}

func NewArchiveOrgClient(config settings.ArchiveorgSettings) *Client {
	return &Client{
		client: &http.Client{
			Timeout: config.ParsedTimeout,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
		workPool: workerpool.New(config.ConsecutiveRequestsLimit),
	}
}
