package pkg

import (
	"fmt"
)

type Event struct {
	Type      string   `json:"type"`
	CreatedAt string   `json:"created_at"`
	Actor     Actor    `json:"actor"`
	Repo      WithName `json:"repo"`
}

func (e Event) String() string {
	return fmt.Sprintf("GENERIC %s - %s - %s at %s", e.CreatedAt, e.Actor, e.Type, e.Repo)
}

type IssueCommentEvent struct {
	Event   `json:",squash"`
	Payload IssueCommentPayload `json:"payload"`
}

func (e IssueCommentEvent) String() string {
	res := fmt.Sprintf("%s - %s %s comment at %s\n", e.CreatedAt, e.Actor, e.Payload.Action, e.Repo)
	res += fmt.Sprintf("\t%s - %s\n", e.Payload.Issue.Url, e.Payload.Issue.Title)
	res += fmt.Sprintf("\t----\n%s\n\t----", LeftAdjust(e.Payload.Comment.Body, "\t"))
	return res
}

type PushEvent struct {
	Event   `json:",squash"`
	Payload PushEventPayload `json:"payload"`
}

func (e PushEvent) String() string {
	res := fmt.Sprintf("%s - %s pushed %d commits at %s on ref %s\n", e.CreatedAt, e.Actor, len(e.Payload.Commits), e.Repo, e.Payload.Ref)
	for i, commit := range e.Payload.Commits {
		res += fmt.Sprintf("\t - %s\n", commit.Sha)
		res += fmt.Sprintf("\t    ----\n%s\n\t    ----", LeftAdjust(commit.Message, "\t    "))
		if i != len(e.Payload.Commits)-1 {
			res += "\n"
		}
	}
	return res
}

type ForkEvent struct {
	Event `json:",squash"`
}

func (e ForkEvent) String() string {
	return fmt.Sprintf("%s - %s forked %s", e.CreatedAt, e.Actor, e.Repo)
}

type CreateEvent struct {
	Event   `json:",squash"`
	Payload CreateDeleteEventPayload `json:"payload"`
}

func (e CreateEvent) String() string {
	return fmt.Sprintf("%s - %s created %s %s at %s", e.CreatedAt, e.Actor, e.Payload.RefType, e.Payload.Ref, e.Repo)
}

type DeleteEvent struct {
	Event   `json:",squash"`
	Payload CreateDeleteEventPayload `json:"payload"`
}

func (e DeleteEvent) String() string {
	return fmt.Sprintf("%s - %s deleted %s %s at %s", e.CreatedAt, e.Actor, e.Payload.RefType, e.Payload.Ref, e.Repo)
}

type PullRequestEvent struct {
	Event   `json:",squash"`
	Payload PullRequestEventPayload `json:"payload"`
}

func (e PullRequestEvent) String() string {
	res := ""
	res += fmt.Sprintf("%s - %s %s PR #%d at %s\n", e.CreatedAt, e.Actor, e.Payload.Action, e.Payload.PullRequest.Number, e.Repo)
	res += fmt.Sprintf("\twants to merge %s into %s \n", e.Payload.PullRequest.Head.Label, e.Payload.PullRequest.Base.Label)
	res += fmt.Sprintf("\t%s", e.Payload.PullRequest.HtmlUrl)
	return res
}

type PullRequestReviewCommentEvent struct {
	Event   `json:",squash"`
	Payload PullRequestReviewCommentEventPayload `json:"payload"`
}

func (e PullRequestReviewCommentEvent) String() string {
	res := ""
	res += fmt.Sprintf("%s - %s %s PR review comment at %s#%d ", e.CreatedAt, e.Actor, e.Payload.Action, e.Repo, e.Payload.PullRequest.Number)
	res += fmt.Sprintf("(%s <-- %s)\n", e.Payload.PullRequest.Head.Label, e.Payload.PullRequest.Base.Label)
	res += fmt.Sprintf("\t%s\n", e.Payload.Comment.Url)
	res += fmt.Sprintf("\t----\n%s\n\t----", LeftAdjust(e.Payload.Comment.Body, "\t"))
	return res
}

type IssuesEvent struct {
	Event   `json:",squash"`
	Payload IssuesEventPayload `json:"payload"`
}

func (e IssuesEvent) String() string {
	res := ""
	res += fmt.Sprintf("%s - %s %s issue at %s#%d - %s\n", e.CreatedAt, e.Actor, e.Payload.Action, e.Repo, e.Payload.Issue.Number, e.Payload.Issue.Title)
	res += fmt.Sprintf("\t%s\n", e.Payload.Issue.Url)
	res += fmt.Sprintf("\t----\n%s\n\t----", LeftAdjust(e.Payload.Issue.Body, "\t"))
	return res
}
