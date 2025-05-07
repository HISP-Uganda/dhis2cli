package analytics

// HTTPResponse struct represents the main payload structure
type HTTPResponse struct {
	HTTPStatus     string   `json:"httpStatus"`
	HTTPStatusCode int      `json:"httpStatusCode"`
	Message        string   `json:"message"`
	Response       Response `json:"response"`
	Status         string   `json:"status"`
}

// Response struct represents the nested "response" structure
type Response struct {
	Created                  string        `json:"created"`
	ID                       string        `json:"id"`
	JobParameters            JobParameters `json:"jobParameters"`
	JobStatus                string        `json:"jobStatus"`
	JobType                  string        `json:"jobType"`
	Name                     string        `json:"name"`
	RelativeNotifierEndpoint string        `json:"relativeNotifierEndpoint"`
	ResponseType             string        `json:"responseType"`
}

// JobParameters struct represents the nested "jobParameters" structure
type JobParameters struct {
	SkipOutliers       bool     `json:"skipOutliers"`
	SkipPrograms       []string `json:"skipPrograms"`
	SkipResourceTables bool     `json:"skipResourceTables"`
	SkipTableTypes     []string `json:"skipTableTypes"`
}

type SubTaskStatus struct {
	UID       string `json:"uid,omitempty"`
	Level     string `json:"level,omitempty"`
	Category  string `json:"category,omitempty"`
	Time      string `json:"time,omitempty"`
	Message   string `json:"message,omitempty"`
	Completed bool   `json:"completed,omitempty"`
	ID        string `json:"id,omitempty"`
}
