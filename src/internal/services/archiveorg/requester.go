package archiveorg

import (
	"errors"
	"log"
	"net/url"
)

const (
	// archiveEndpoint endpoint for saving pages in web.archive.org. With ending slash / (after which is given the URL to save)
	archiveEndpoint = "https://web.archive.org/save/"
	// archivedResultLocationHeader is the header returned by the original response (which is a redirection), with the URL to the saved snapshot
	archivedResultLocationHeader = "Location"
)

var (
	InvalidArchiveResponse = errors.New("invalid response")
)

func (client *Client) ArchiveURL(_url string) (archivedResultURL string, err error) {
	client.workPool.SubmitWait(func() {
		archivedResultURL, err = client.archiveURLWork(_url)
	})
	return
}

func (client *Client) archiveURLWork(_url string) (archivedResultURL string, err error) {
	archiveURL, err := url.Parse(_url)
	if err != nil {
		return
	}

	requestURL := archiveEndpoint + archiveURL.String() // TODO urlencode archiveURL
	log.Printf("Archiving with request %s ...", requestURL)

	response, err := client.client.Get(requestURL)
	if err != nil {
		return
	}

	statuscode := response.StatusCode
	archivedResultURL = response.Header.Get(archivedResultLocationHeader)
	if statuscode >= 400 || archivedResultURL == "" {
		err = InvalidArchiveResponse
		return
	}

	return
}
