package _structs

type TeleMessageRequest struct {
	ResponseId  string       `json:"responseId"`
	QueryResult QueryResults `json:"queryResult"`
}

type QueryResults struct {
	QueryText       string      `json:"queryText"`
	Intent          IntentQuery `json:"intent"`
	FulfillmentText string      `json:"fulfillmentText"`
}
type IntentQuery struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}
