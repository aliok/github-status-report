package pkg

type Actor struct {
	DisplayLogin string `json:"display_login"`
}

func (a Actor) String() string {
	return a.DisplayLogin
}

type WithName struct {
	Name string `json:"name"`
}

type WithTitle struct {
	Title string `json:"title"`
}

type WithUrl struct {
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

type WithBody struct {
	Body string `json:"body"`
}

type WithAction struct {
	Action string `json:"action"`
}

type Ref struct {
	Label string `json:"label"`
}

type Issue struct {
	WithUrl     `json:",squash"`
	WithTitle   `json:",squash"`
	WithBody    `json:",squash"`
	PullRequest WithUrl `json:"pull_request,omitempty"`
	Number      int     `json:"number"`
}

type Comment struct {
	WithUrl  `json:",squash"`
	WithBody `json:",squash"`
}

type Commit struct {
	Sha     string `json:"sha"`
	Message string `json:"message"`
}

type PullRequest struct {
	WithUrl   `json:",squash"`
	WithTitle `json:",squash"`
	WithBody  `json:",squash"`
	Head      Ref `json:"head"`
	Base      Ref `json:"base"`
	Number    int `json:"number"`
}

type IssueCommentPayload struct {
	WithAction `json:",squash"`
	Issue      Issue   `json:"issue"`
	Comment    Comment `json:"comment"`
}

type PushEventPayload struct {
	Ref     string   `json:"ref"`
	Commits []Commit `json:"commits"`
}

type CreateDeleteEventPayload struct {
	Ref     string `json:"ref"`
	RefType string `json:"ref_type"`
}

type PullRequestEventPayload struct {
	WithAction  `json:",squash"`
	PullRequest PullRequest `json:"pull_request"`
}

type PullRequestReviewCommentEventPayload struct {
	WithAction  `json:",squash"`
	Comment     Comment     `json:"comment"`
	PullRequest PullRequest `json:"pull_request"`
}

type IssuesEventPayload struct {
	Action string `json:"action"`
	Issue  Issue  `json:"issue"`
}

type WatchEventPayload struct {
	WithAction `json:",squash"`
}
