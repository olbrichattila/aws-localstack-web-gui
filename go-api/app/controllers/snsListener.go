package controller

import (
	"net/http"
	"webuiApi/app/repositories/snslistener"

	"github.com/olbrichattila/gofra/pkg/app/gofraerror"
)

type portsResponse struct {
	Ports []snslistener.ListenerInfo `json:"ports"`
}

type SNSListenerController struct {
}

// SnsListenerAction function can take any parameters defined in the Di config
func (c *SNSListenerController) NewSNSListener(port int, sns snslistener.SNSListener) (string, error) {
	err := sns.Listen(port)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusBadRequest)
	}

	return "{}", nil
}

func (c *SNSListenerController) CloseSNSListener(port int, sns snslistener.SNSListener) (string, error) {
	err := sns.Close(port)
	if err != nil {
		return "", gofraerror.NewJSON(err.Error(), http.StatusBadRequest)
	}

	return "{}", nil
}

func (c *SNSListenerController) GetRequests(port int, sns snslistener.SNSListener) (map[string]any, error) {
	req, err := sns.GetRequests(port)
	if err != nil {
		return nil, gofraerror.NewJSON(err.Error(), http.StatusBadRequest)
	}

	response := map[string]any{
		"requests": req,
	}

	return response, nil
}

func (c *SNSListenerController) GetListeners(sns snslistener.SNSListener) portsResponse {
	ports := sns.GetListeningPorts()

	portItems := make([]snslistener.ListenerInfo, len(ports))
	for i, port := range ports {
		portItems[i] = port
	}

	return portsResponse{
		Ports: portItems,
	}
}
