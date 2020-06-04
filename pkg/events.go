package pkg

import (
	"fmt"
)

type Event interface {
	IsRelevantToReport(username string) bool
}

// -----------------

type BaseEvent struct {
	Type      string   `json:"type"`
	CreatedAt string   `json:"created_at"`
	Actor     Actor    `json:"actor"`
	Repo      WithName `json:"repo"`
}

func (e BaseEvent) String() string {
	return fmt.Sprintf("GENERIC %s - %s - %s at %s", e.CreatedAt, e.Actor, e.Type, e.Repo.Name)
}

// -----------------

var _ Event = IssueCommentEvent{}

type IssueCommentEvent struct {
	BaseEvent `json:",squash"`
	Payload   IssueCommentPayload `json:"payload"`
}

func (e IssueCommentEvent) String() string {
	res := fmt.Sprintf("%s - %s %s comment at %s\n", e.CreatedAt, e.Actor, e.Payload.Action, e.Repo.Name)
	res += fmt.Sprintf("\t%s - %s\n", e.Payload.Issue.HtmlUrl, e.Payload.Issue.Title)
	res += fmt.Sprintf("\t----\n%s\n\t----", LeftAdjust(e.Payload.Comment.Body, "\t"))
	return res
}

func (e IssueCommentEvent) IsRelevantToReport(username string) bool {
	return e.Actor.DisplayLogin == username
}

// -----------------

var _ Event = PushEvent{}

type PushEvent struct {
	BaseEvent `json:",squash"`
	Payload   PushEventPayload `json:"payload"`
}

func (e PushEvent) String() string {
	res := fmt.Sprintf("%s - %s pushed %d commits at %s on ref %s\n", e.CreatedAt, e.Actor, len(e.Payload.Commits), e.Repo.Name, e.Payload.Ref)
	for i, commit := range e.Payload.Commits {
		res += fmt.Sprintf("\t - %s\n", commit.Sha)
		res += fmt.Sprintf("\t    ----\n%s\n\t    ----", LeftAdjust(commit.Message, "\t    "))
		if i != len(e.Payload.Commits)-1 {
			res += "\n"
		}
	}
	return res
}

func (e PushEvent) IsRelevantToReport(username string) bool {
	return e.Actor.DisplayLogin == username
}

// -----------------

var _ Event = ForkEvent{}

type ForkEvent struct {
	BaseEvent `json:",squash"`
}

func (e ForkEvent) String() string {
	return fmt.Sprintf("%s - %s forked %s", e.CreatedAt, e.Actor, e.Repo.Name)
}

func (e ForkEvent) IsRelevantToReport(username string) bool {
	return e.Actor.DisplayLogin != username
}

// -----------------

var _ Event = WatchEvent{}

type WatchEvent struct {
	BaseEvent `json:",squash"`
	Payload   WatchEventPayload `json:"payload"`
}

func (e WatchEvent) String() string {
	return fmt.Sprintf("%s - %s %s %s", e.CreatedAt, e.Actor, e.Payload.Action, e.Repo.Name)
}

func (e WatchEvent) IsRelevantToReport(username string) bool {
	return e.Actor.DisplayLogin != username
}

// -----------------

var _ Event = CreateEvent{}

type CreateEvent struct {
	BaseEvent `json:",squash"`
	Payload   CreateDeleteEventPayload `json:"payload"`
}

func (e CreateEvent) String() string {
	return fmt.Sprintf("%s - %s created %s %s at %s", e.CreatedAt, e.Actor, e.Payload.RefType, e.Payload.Ref, e.Repo.Name)
}

func (e CreateEvent) IsRelevantToReport(username string) bool {
	return e.Payload.RefType != "branch"
}

// -----------------

var _ Event = DeleteEvent{}

type DeleteEvent struct {
	BaseEvent `json:",squash"`
	Payload   CreateDeleteEventPayload `json:"payload"`
}

func (e DeleteEvent) String() string {
	return fmt.Sprintf("%s - %s deleted %s %s at %s", e.CreatedAt, e.Actor, e.Payload.RefType, e.Payload.Ref, e.Repo.Name)
}

func (e DeleteEvent) IsRelevantToReport(username string) bool {
	return e.Payload.RefType != "branch"
}

// -----------------

var _ Event = PullRequestEvent{}

type PullRequestEvent struct {
	BaseEvent `json:",squash"`
	Payload   PullRequestEventPayload `json:"payload"`
}

func (e PullRequestEvent) String() string {
	res := ""
	res += fmt.Sprintf("%s - %s %s PR #%d at %s\n", e.CreatedAt, e.Actor, e.Payload.Action, e.Payload.PullRequest.Number, e.Repo.Name)
	res += fmt.Sprintf("\twants to merge %s into %s \n", e.Payload.PullRequest.Head.Label, e.Payload.PullRequest.Base.Label)
	res += fmt.Sprintf("\t%s", e.Payload.PullRequest.HtmlUrl)
	return res
}

func (e PullRequestEvent) IsRelevantToReport(username string) bool {
	return e.Actor.DisplayLogin == username
}

// -----------------

var _ Event = PullRequestReviewCommentEvent{}

type PullRequestReviewCommentEvent struct {
	BaseEvent `json:",squash"`
	Payload   PullRequestReviewCommentEventPayload `json:"payload"`
}

func (e PullRequestReviewCommentEvent) String() string {
	res := ""
	res += fmt.Sprintf("%s - %s %s PR review comment at %s#%d ", e.CreatedAt, e.Actor, e.Payload.Action, e.Repo.Name, e.Payload.PullRequest.Number)
	res += fmt.Sprintf("(%s <-- %s)\n", e.Payload.PullRequest.Head.Label, e.Payload.PullRequest.Base.Label)
	res += fmt.Sprintf("\t%s\n", e.Payload.Comment.HtmlUrl)
	res += fmt.Sprintf("\t----\n%s\n\t----", LeftAdjust(e.Payload.Comment.Body, "\t"))
	return res
}

func (e PullRequestReviewCommentEvent) IsRelevantToReport(username string) bool {
	return e.Actor.DisplayLogin == username
}

// -----------------

var _ Event = IssuesEvent{}

type IssuesEvent struct {
	BaseEvent `json:",squash"`
	Payload   IssuesEventPayload `json:"payload"`
}

func (e IssuesEvent) String() string {
	res := ""
	res += fmt.Sprintf("%s - %s %s issue at %s#%d - %s\n", e.CreatedAt, e.Actor, e.Payload.Action, e.Repo.Name, e.Payload.Issue.Number, e.Payload.Issue.Title)
	res += fmt.Sprintf("\t%s\n", e.Payload.Issue.HtmlUrl)
	res += fmt.Sprintf("\t----\n%s\n\t----", LeftAdjust(e.Payload.Issue.Body, "\t"))
	return res
}

func (e IssuesEvent) IsRelevantToReport(username string) bool {
	return e.Actor.DisplayLogin == username
}
