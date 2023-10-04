package app

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnlineMeetingsGetAllRecordingsResponse 
// Deprecated: This class is obsolete. Use getAllRecordingsGetResponse instead.
type OnlineMeetingsGetAllRecordingsResponse struct {
    OnlineMeetingsGetAllRecordingsGetResponse
}
// NewOnlineMeetingsGetAllRecordingsResponse instantiates a new OnlineMeetingsGetAllRecordingsResponse and sets the default values.
func NewOnlineMeetingsGetAllRecordingsResponse()(*OnlineMeetingsGetAllRecordingsResponse) {
    m := &OnlineMeetingsGetAllRecordingsResponse{
        OnlineMeetingsGetAllRecordingsGetResponse: *NewOnlineMeetingsGetAllRecordingsGetResponse(),
    }
    return m
}
// CreateOnlineMeetingsGetAllRecordingsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnlineMeetingsGetAllRecordingsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlineMeetingsGetAllRecordingsResponse(), nil
}
// OnlineMeetingsGetAllRecordingsResponseable 
// Deprecated: This class is obsolete. Use getAllRecordingsGetResponse instead.
type OnlineMeetingsGetAllRecordingsResponseable interface {
    OnlineMeetingsGetAllRecordingsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
