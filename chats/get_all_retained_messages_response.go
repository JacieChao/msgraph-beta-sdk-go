package chats

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetAllRetainedMessagesResponse 
// Deprecated: This class is obsolete. Use getAllRetainedMessagesGetResponse instead.
type GetAllRetainedMessagesResponse struct {
    GetAllRetainedMessagesGetResponse
}
// NewGetAllRetainedMessagesResponse instantiates a new GetAllRetainedMessagesResponse and sets the default values.
func NewGetAllRetainedMessagesResponse()(*GetAllRetainedMessagesResponse) {
    m := &GetAllRetainedMessagesResponse{
        GetAllRetainedMessagesGetResponse: *NewGetAllRetainedMessagesGetResponse(),
    }
    return m
}
// CreateGetAllRetainedMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetAllRetainedMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetAllRetainedMessagesResponse(), nil
}
// GetAllRetainedMessagesResponseable 
// Deprecated: This class is obsolete. Use getAllRetainedMessagesGetResponse instead.
type GetAllRetainedMessagesResponseable interface {
    GetAllRetainedMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
