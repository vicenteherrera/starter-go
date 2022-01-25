package analyzer

import log "github.com/sirupsen/logrus"

type Client interface {
	AnalyzeFile() (AnalyzerResponse, error)
}

func NewClient(filepath string) Client {
	return &client{
		filepath: filepath,
	}
}

type client struct {
	filepath string
}

func (s *client) AnalyzeFile() (AnalyzerResponse, error) {

	var response AnalyzerResponse
	response.AnalysisStatus = "analyzed"
	log.Info("Analyzing " + s.filepath)
	log.Info(s.filepath + " : " + response.AnalysisStatus)

	return response, nil
}
