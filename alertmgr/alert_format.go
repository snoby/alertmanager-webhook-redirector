package alertmgr

import "time"

type AlertManagerNative struct {
	Alerts []struct {
		Annotations struct {
			Description string `json:"description"`
			Summary     string `json:"summary"`
		} `json:"annotations"`
		EndsAt       time.Time `json:"endsAt"`
		GeneratorURL string    `json:"generatorURL"`
		Labels       struct {
			Alertname string `json:"alertname"`
			Job       string `json:"job"`
			Service   string `json:"service"`
			Severity  string `json:"severity"`
			Verb      string `json:"verb"`
		} `json:"labels"`
		StartsAt time.Time `json:"startsAt"`
		Status   string    `json:"status"`
	} `json:"alerts"`
	CommonAnnotations struct {
		Summary string `json:"summary"`
	} `json:"commonAnnotations"`
	CommonLabels struct {
		Alertname string `json:"alertname"`
		Job       string `json:"job"`
		Service   string `json:"service"`
		Severity  string `json:"severity"`
	} `json:"commonLabels"`
	ExternalURL string `json:"externalURL"`
	GroupKey    string `json:"groupKey"`
	GroupLabels struct {
		Alertname string `json:"alertname"`
		Service   string `json:"service"`
	} `json:"groupLabels"`
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
	Version  string `json:"version"`
}
