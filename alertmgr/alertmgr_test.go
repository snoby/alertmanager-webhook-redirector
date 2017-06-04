package alertmgr_test

import (
	"testing"

	"github.com/snoby/alertmanager-webhook-redirector/alertmgr"
)

const minimizedJson = `{"alerts":[{"annotations":{"description":"99th percentile Latency for proxy requests to the kube-apiserver is higher than 1s.","summary":"Kubernetes apiserver latency is high"},"endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://prometheus.int.ops.tropo.com/graph#%5B%7B%22expr%22%3A%22histogram_quantile%280.99%2C%20sum%28apiserver_request_latencies_bucket%7Bverb%21~%5C%22CONNECT%7CWATCHLIST%7CWATCH%5C%22%7D%29%20WITHOUT%20%28instance%2C%20node%2C%20resource%29%29%20%2F%201000000%20%3E%201%22%2C%22tab%22%3A0%7D%5D","labels":{"alertname":"K8SApiServerLatency","job":"kubernetes-cluster","service":"k8s","severity":"warning","verb":"proxy"},"startsAt":"2017-05-31T19:09:37.818Z","status":"firing"},{"annotations":{"description":"99th percentile Latency for POST requests to the kube-apiserver is higher than 1s.","summary":"Kubernetes apiserver latency is high"},"endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://prometheus.int.ops.tropo.com/graph#%5B%7B%22expr%22%3A%22histogram_quantile%280.99%2C%20sum%28apiserver_request_latencies_bucket%7Bverb%21~%5C%22CONNECT%7CWATCHLIST%7CWATCH%5C%22%7D%29%20WITHOUT%20%28instance%2C%20node%2C%20resource%29%29%20%2F%201000000%20%3E%201%22%2C%22tab%22%3A0%7D%5D","labels":{"alertname":"K8SApiServerLatency","job":"kubernetes-cluster","service":"k8s","severity":"warning","verb":"POST"},"startsAt":"2017-05-31T19:09:37.818Z","status":"firing"},{"annotations":{"description":"99th percentile Latency for PROXY requests to the kube-apiserver is higher than 1s.","summary":"Kubernetes apiserver latency is high"},"endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://prometheus.int.ops.tropo.com/graph#%5B%7B%22expr%22%3A%22histogram_quantile%280.99%2C%20sum%28apiserver_request_latencies_bucket%7Bverb%21~%5C%22CONNECT%7CWATCHLIST%7CWATCH%5C%22%7D%29%20WITHOUT%20%28instance%2C%20node%2C%20resource%29%29%20%2F%201000000%20%3E%201%22%2C%22tab%22%3A0%7D%5D","labels":{"alertname":"K8SApiServerLatency","job":"kubernetes-cluster","service":"k8s","severity":"warning","verb":"PROXY"},"startsAt":"2017-05-31T19:09:37.818Z","status":"firing"}],"commonAnnotations":{"summary":"Kubernetes apiserver latency is high"},"commonLabels":{"alertname":"K8SApiServerLatency","job":"kubernetes-cluster","service":"k8s","severity":"warning"},"externalURL":"http://prometheus.int.ops.tropo.com/alertmanager","groupKey":"{}:{alertname=\"K8SApiServerLatency\", service=\"k8s\"}","groupLabels":{"alertname":"K8SApiServerLatency","service":"k8s"},"receiver":"custom_webhook","status":"firing","version":"4"}`

// type AlertMgr struct {
// 	buf []byte
// }
//
// func (alert *AlertMgr) LoadRawData(buf []byte) (err error) {
// 	fmt.Println(" Loading buffer into AlertMgr memory")
// 	alert.buf = buf
// 	return nil
// }
func TestLoadRawData(t *testing.T) {
	var test alertmgr.AlertMgr
	buf := []byte(minimizedJson)

	err := test.LoadRawData(buf)
	if err != nil {
		t.Error("Not able to load data to be processed")
	}
}
