package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessageReactionIdentitySet 
type ChatMessageReactionIdentitySet struct {
    IdentitySet
}
// NewChatMessageReactionIdentitySet instantiates a new chatMessageReactionIdentitySet and sets the default values.
func NewChatMessageReactionIdentitySet()(*ChatMessageReactionIdentitySet) {
    m := &ChatMessageReactionIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    return m
}
// CreateChatMessageReactionIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageReactionIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessageReactionIdentitySet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageReactionIdentitySet) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ChatMessageReactionIdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
