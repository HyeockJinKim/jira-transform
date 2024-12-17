package servers

import "github.com/HyeockJinKim/jira-transform/data"

type JIRAMirrorRequest struct {
	Source      string   `param:"source"`
	Destination string   `param:"destination"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Labels      []string `json:"labels"`
}

func (m JIRAMirrorRequest) ToIssue() data.Issue {
	return data.Issue{
		Title:       m.Title,
		Description: JiraToMD(m.Description),
		Labels:      m.Labels,
	}
}
