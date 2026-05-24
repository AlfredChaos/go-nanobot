// 消息类型
package bus

import "time"

type InboundMessage struct {
	Channel   string         `json:"channel"`
	SenderID  string         `json:"sender_id"`
	ChatID    string         `json:"chat_id"`
	Content   string         `json:"content"`
	Timestamp time.Time      `json:"timestamp"`
	Media     []string       `json:"media,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

type OutboundMessage struct {
	Channel  string         `json:"channel"`
	ChatID   string         `json:"chat_id"`
	Content  string         `json:"content"`
	ReplyTo  string         `json:"reply_to,omitempty"`
	Media    []string       `json:"media,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
	Buttons  [][]string     `json:"buttons,omitempty"`
}
