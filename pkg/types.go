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

func (w WithName) String() string {
	return w.Name
}

type WithTitle struct {
	Title string `json:"title"`
}

func (w WithTitle) String() string {
	return w.Title
}

type WithUrl struct {
	Url string `json:"url"`
}

func (w WithUrl) String() string {
	return w.Url
}

type WithBody struct {
	Body string `json:"body"`
}

func (w WithBody) String() string {
	return w.Body
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
	HtmlUrl   string `json:"html_url"`
	Head      Ref    `json:"head"`
	Base      Ref    `json:"base"`
	Number    int    `json:"number"`
}

type IssueCommentPayload struct {
	Action  string  `json:"action"`
	Issue   Issue   `json:"issue"`
	Comment Comment `json:"comment"`
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
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
}

type PullRequestReviewCommentEventPayload struct {
	Action      string      `json:"action"`
	Comment     Comment     `json:"comment"`
	PullRequest PullRequest `json:"pull_request"`
}

type IssuesEventPayload struct {
	Action string `json:"action"`
	Issue  Issue  `json:"issue"`
}
