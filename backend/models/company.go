package models

type Company struct {
	ID                string `json:"id,omitempty"`
	CompanyName       string `json:"company_name"`
	Status            string `json:"status"`
	LatestMemo        string `json:"latest_memo"`
	ProgressStep      string `json:"progress_step"`
	ShortMemo         string `json:"short_memo"`
	EntryDate         string `json:"entry_date"`
	ESDeadline        string `json:"es_deadline"`
	InterviewMemo     string `json:"interview_memo"`
	EntrySite         string `json:"entry_site"`
	RequiredDocuments string `json:"required_documents"`
}
