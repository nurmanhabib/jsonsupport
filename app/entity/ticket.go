package entity

type Ticket struct {
	CreatedAt                string       `json:"Created At"`
	ResolvedAt               string       `json:"Resolved At"`
	Title                    string       `json:"Title"`
	ReopenedAt               string       `json:"Reopened At"`
	FirstRespondedAt         string       `json:"First Responded At"`
	AverageResponseTimeSecs  interface{}  `json:"Average Response Time (secs)"`
	CustomerName             string       `json:"Customer Name"`
	Tags                     string       `json:"Tags"`
	CustomerPhoneNumber      string       `json:"Customer Phone Number"`
	FirstResponseTimeSecs    interface{}  `json:"First Response Time (secs)"`
	Status                   string       `json:"Status"`
	Priority                 string       `json:"Priority"`
	CustomerEmail            string       `json:"Customer Email"`
	AssignedAt               string       `json:"Assigned At"`
	AgentEmail               string       `json:"Agent Email"`
	NumberOfPosts            int          `json:"Number of Posts"`
	Channel                  string       `json:"Channel"`
	AgentName                string       `json:"Agent Name"`
	AssigneeFirstRespondedAt string       `json:"Assignee First Responded At"`
	CustomFieldsRaw          string       `json:"Custom Fields"`
	CustomFields             CustomFields `json:"-"`
}

type CustomFields map[string]string
