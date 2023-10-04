package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCloudPCsValidateBulkResizeResponse 
// Deprecated: This class is obsolete. Use validateBulkResizePostResponse instead.
type ItemCloudPCsValidateBulkResizeResponse struct {
    ItemCloudPCsValidateBulkResizePostResponse
}
// NewItemCloudPCsValidateBulkResizeResponse instantiates a new ItemCloudPCsValidateBulkResizeResponse and sets the default values.
func NewItemCloudPCsValidateBulkResizeResponse()(*ItemCloudPCsValidateBulkResizeResponse) {
    m := &ItemCloudPCsValidateBulkResizeResponse{
        ItemCloudPCsValidateBulkResizePostResponse: *NewItemCloudPCsValidateBulkResizePostResponse(),
    }
    return m
}
// CreateItemCloudPCsValidateBulkResizeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCloudPCsValidateBulkResizeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCloudPCsValidateBulkResizeResponse(), nil
}
// ItemCloudPCsValidateBulkResizeResponseable 
// Deprecated: This class is obsolete. Use validateBulkResizePostResponse instead.
type ItemCloudPCsValidateBulkResizeResponseable interface {
    ItemCloudPCsValidateBulkResizePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
