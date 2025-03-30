package biz

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/Fl0rencess720/Wittgenstein/pkgs/utils"
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	UID          string   `gorm:"primaryKey;column:uid;size:255"`
	Content      string   `gorm:"column:content;type:text"`
	State        State    `gorm:"-;"`
	Participants []string `gorm:"column:participants;type:json;serializer:json"`
	Speeches     []Speech `gorm:"foreignKey:TopicUID"`
	Title        string   `gorm:"column:title;type:varchar(255)"`
	TitleImage   string   `gorm:"column:title_image;type:varchar(255)"`
	Phone        string   `gorm:"column:phone;type:varchar(255)"`
}

type Speech struct {
	gorm.Model
	UID      string    `gorm:"primaryKey;column:uid;size:255"`
	TopicUID string    `gorm:"column:topic_uid;size:255;index"`
	RoleUID  string    `gorm:"column:role_uid;size:255"`
	Content  string    `gorm:"column:content;type:text"`
	Time     time.Time `gorm:"column:time"`
}

func (t *Topic) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &t.Participants)
}

func (t Topic) Value() (driver.Value, error) {
	return json.Marshal(t.Participants)
}

func NewTopic(content string, participants []string) (*Topic, error) {
	uid, err := utils.GetSnowflakeID(0)
	if err != nil {
		return nil, err
	}
	return &Topic{
		Content:      content,
		State:        &PreparingState{},
		UID:          uid,
		Participants: participants,
		Speeches:     []Speech{},
		Title:        "新主题",
	}, nil
}

func (topic *Topic) GetState() string {
	return topic.State.getState()
}

func (topic *Topic) Start() error {
	return topic.State.start(topic)
}

func (topic *Topic) Pause() error {
	return topic.State.pause(topic)
}

func (topic *Topic) Resume() error {
	return topic.State.resume(topic)
}
