package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zubeyiro/secure-chat/internal/events"
)

func TestNewMessage(t *testing.T) {
	msg := NewMessage(events.NEW_USER_CONNECTED, "zubeyir", "content")

	assert.Equal(t, msg.Command, events.NEW_USER_CONNECTED)
	assert.Equal(t, msg.Owner, "zubeyir")
	assert.Equal(t, msg.Message, "content")
}

func TestSerialization(t *testing.T) {
	msg := NewMessage(events.NEW_USER_CONNECTED, "zubeyir", "content")

	assert.Equal(t, msg.Serialize(), "user_connected,zubeyir,content\n")
}

func TestDeserializeSuccess(t *testing.T) {
	msg := Deserialize("user_connected,zubeyir,content\n")

	assert.Equal(t, msg.Command, events.NEW_USER_CONNECTED)
	assert.Equal(t, msg.Owner, "zubeyir")
	assert.Equal(t, msg.Message, "content\n")
}

func TestDeserializeFail(t *testing.T) {
	msg := Deserialize("user_connected,zubeyir")

	assert.Equal(t, msg, (*Message)(nil))
}
