package notification_entity

import (
	"time"

	"github.com/google/uuid"
)

func NewNotification(rawDate string, comunicationChannelID string) (*Notification, error) {
	refDate := "2006-01-02 15:00"

	convDate, err := time.Parse(refDate, rawDate)
	if err != nil {
		return nil, err
	}

	n := &Notification{
		ID:                    uuid.New().String(),
		Date:                  convDate,
		ComunicationChannelID: comunicationChannelID,
		Status:                "pending",
		Content:               "",
		CreatedAt:             time.Now(),
	}

	if err := n.Validate(); err != nil {
		return nil, err
	}
	return n, nil
}

type Notification struct {
	ID                    string    `json:"id"`
	Date                  time.Time `json:"date"`
	ComunicationChannelID string    `json:"comunication_channel_id"`
	Status                string    `json:"status"`
	Content               string    `json:"content"`
	CreatedAt             time.Time `json:"created_at"`
}

func (n *Notification) Validate() error {
	return nil
}
