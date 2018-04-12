//reusable code
package utility

type APILogMessage struct {
	Method    string `json:"method"`
	URI       string `json:"uri"`
	Name      string `json:"name"`
	TimeTaken int64  `json:"time_taken"`
	Type      string `json:"type"`
}
