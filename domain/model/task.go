package model

import (
	"errors"

	"github.com/google/uuid"
)

func NewTask(title, assigneeId string) (Task, error) {
	if title == "" {
		return Task{}, errors.New("タイトルを入力してください。")
	}
	return Task{
		id:         uuid.New().String(),
		title:      title,
		status:     NEW,
		assigneeId: assigneeId,
	}, nil
}

func GenerateTask(id string, title string, status Status, assigneeId string) Task {
	return Task{
		id:         uuid.New().String(),
		title:      title,
		status:     NEW,
		assigneeId: assigneeId,
	}
}

type Task struct {
	id         string
	title      string
	status     Status
	assigneeId string
}

func (t *Task) Id() string {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Status() string {
	return string(t.status)
}

type Status string

func (Status) Convert(status string) (Status, error) {
	switch status {
	case string(NEW):
		return NEW, nil
	case string(WORKING):
		return WORKING, nil
	case string(DONE):
		return DONE, nil
	default:
		return "", errors.New("存在しないステータス")
	}
}

const (
	NEW     Status = "新規"
	WORKING        = "作業中"
	DONE           = "完了"
)
