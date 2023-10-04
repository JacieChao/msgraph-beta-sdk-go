package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCloudPCsBulkResizeResponse 
// Deprecated: This class is obsolete. Use bulkResizePostResponse instead.
type ItemCloudPCsBulkResizeResponse struct {
    ItemCloudPCsBulkResizePostResponse
}
// NewItemCloudPCsBulkResizeResponse instantiates a new ItemCloudPCsBulkResizeResponse and sets the default values.
func NewItemCloudPCsBulkResizeResponse()(*ItemCloudPCsBulkResizeResponse) {
    m := &ItemCloudPCsBulkResizeResponse{
        ItemCloudPCsBulkResizePostResponse: *NewItemCloudPCsBulkResizePostResponse(),
    }
    return m
}
// CreateItemCloudPCsBulkResizeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCloudPCsBulkResizeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCloudPCsBulkResizeResponse(), nil
}
// ItemCloudPCsBulkResizeResponseable 
// Deprecated: This class is obsolete. Use bulkResizePostResponse instead.
type ItemCloudPCsBulkResizeResponseable interface {
    ItemCloudPCsBulkResizePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
