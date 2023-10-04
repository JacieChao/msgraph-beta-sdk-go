package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCloudPCsItemGetSupportedCloudPcRemoteActionsResponse 
// Deprecated: This class is obsolete. Use getSupportedCloudPcRemoteActionsGetResponse instead.
type ItemCloudPCsItemGetSupportedCloudPcRemoteActionsResponse struct {
    ItemCloudPCsItemGetSupportedCloudPcRemoteActionsGetResponse
}
// NewItemCloudPCsItemGetSupportedCloudPcRemoteActionsResponse instantiates a new ItemCloudPCsItemGetSupportedCloudPcRemoteActionsResponse and sets the default values.
func NewItemCloudPCsItemGetSupportedCloudPcRemoteActionsResponse()(*ItemCloudPCsItemGetSupportedCloudPcRemoteActionsResponse) {
    m := &ItemCloudPCsItemGetSupportedCloudPcRemoteActionsResponse{
        ItemCloudPCsItemGetSupportedCloudPcRemoteActionsGetResponse: *NewItemCloudPCsItemGetSupportedCloudPcRemoteActionsGetResponse(),
    }
    return m
}
// CreateItemCloudPCsItemGetSupportedCloudPcRemoteActionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCloudPCsItemGetSupportedCloudPcRemoteActionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCloudPCsItemGetSupportedCloudPcRemoteActionsResponse(), nil
}
// ItemCloudPCsItemGetSupportedCloudPcRemoteActionsResponseable 
// Deprecated: This class is obsolete. Use getSupportedCloudPcRemoteActionsGetResponse instead.
type ItemCloudPCsItemGetSupportedCloudPcRemoteActionsResponseable interface {
    ItemCloudPCsItemGetSupportedCloudPcRemoteActionsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
