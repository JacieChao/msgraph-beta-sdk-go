package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEndpointDeviceImagesGetSourceImagesResponse 
// Deprecated: This class is obsolete. Use getSourceImagesGetResponse instead.
type VirtualEndpointDeviceImagesGetSourceImagesResponse struct {
    VirtualEndpointDeviceImagesGetSourceImagesGetResponse
}
// NewVirtualEndpointDeviceImagesGetSourceImagesResponse instantiates a new VirtualEndpointDeviceImagesGetSourceImagesResponse and sets the default values.
func NewVirtualEndpointDeviceImagesGetSourceImagesResponse()(*VirtualEndpointDeviceImagesGetSourceImagesResponse) {
    m := &VirtualEndpointDeviceImagesGetSourceImagesResponse{
        VirtualEndpointDeviceImagesGetSourceImagesGetResponse: *NewVirtualEndpointDeviceImagesGetSourceImagesGetResponse(),
    }
    return m
}
// CreateVirtualEndpointDeviceImagesGetSourceImagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEndpointDeviceImagesGetSourceImagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEndpointDeviceImagesGetSourceImagesResponse(), nil
}
// VirtualEndpointDeviceImagesGetSourceImagesResponseable 
// Deprecated: This class is obsolete. Use getSourceImagesGetResponse instead.
type VirtualEndpointDeviceImagesGetSourceImagesResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEndpointDeviceImagesGetSourceImagesGetResponseable
}
