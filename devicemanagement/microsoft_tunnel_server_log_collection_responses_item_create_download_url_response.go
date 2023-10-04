package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlResponse 
// Deprecated: This class is obsolete. Use createDownloadUrlPostResponse instead.
type MicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlResponse struct {
    MicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlPostResponse
}
// NewMicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlResponse instantiates a new MicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlResponse and sets the default values.
func NewMicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlResponse()(*MicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlResponse) {
    m := &MicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlResponse{
        MicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlPostResponse: *NewMicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlPostResponse(),
    }
    return m
}
// CreateMicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlResponse(), nil
}
// MicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlResponseable 
// Deprecated: This class is obsolete. Use createDownloadUrlPostResponse instead.
type MicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlResponseable interface {
    MicrosoftTunnelServerLogCollectionResponsesItemCreateDownloadUrlPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
