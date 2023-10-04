package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEndpointCloudPCsBulkResizeResponse 
// Deprecated: This class is obsolete. Use bulkResizePostResponse instead.
type VirtualEndpointCloudPCsBulkResizeResponse struct {
    VirtualEndpointCloudPCsBulkResizePostResponse
}
// NewVirtualEndpointCloudPCsBulkResizeResponse instantiates a new VirtualEndpointCloudPCsBulkResizeResponse and sets the default values.
func NewVirtualEndpointCloudPCsBulkResizeResponse()(*VirtualEndpointCloudPCsBulkResizeResponse) {
    m := &VirtualEndpointCloudPCsBulkResizeResponse{
        VirtualEndpointCloudPCsBulkResizePostResponse: *NewVirtualEndpointCloudPCsBulkResizePostResponse(),
    }
    return m
}
// CreateVirtualEndpointCloudPCsBulkResizeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEndpointCloudPCsBulkResizeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpointCloudPCsBulkResizeResponse(), nil
}
// VirtualEndpointCloudPCsBulkResizeResponseable 
// Deprecated: This class is obsolete. Use bulkResizePostResponse instead.
type VirtualEndpointCloudPCsBulkResizeResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEndpointCloudPCsBulkResizePostResponseable
}
