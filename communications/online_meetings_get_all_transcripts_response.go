package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnlineMeetingsGetAllTranscriptsResponse 
// Deprecated: This class is obsolete. Use getAllTranscriptsGetResponse instead.
type OnlineMeetingsGetAllTranscriptsResponse struct {
    OnlineMeetingsGetAllTranscriptsGetResponse
}
// NewOnlineMeetingsGetAllTranscriptsResponse instantiates a new OnlineMeetingsGetAllTranscriptsResponse and sets the default values.
func NewOnlineMeetingsGetAllTranscriptsResponse()(*OnlineMeetingsGetAllTranscriptsResponse) {
    m := &OnlineMeetingsGetAllTranscriptsResponse{
        OnlineMeetingsGetAllTranscriptsGetResponse: *NewOnlineMeetingsGetAllTranscriptsGetResponse(),
    }
    return m
}
// CreateOnlineMeetingsGetAllTranscriptsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnlineMeetingsGetAllTranscriptsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlineMeetingsGetAllTranscriptsResponse(), nil
}
// OnlineMeetingsGetAllTranscriptsResponseable 
// Deprecated: This class is obsolete. Use getAllTranscriptsGetResponse instead.
type OnlineMeetingsGetAllTranscriptsResponseable interface {
    OnlineMeetingsGetAllTranscriptsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
