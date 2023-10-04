package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemDevicesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemDevicesDeltaResponse struct {
    ItemDevicesDeltaGetResponse
}
// NewItemDevicesDeltaResponse instantiates a new ItemDevicesDeltaResponse and sets the default values.
func NewItemDevicesDeltaResponse()(*ItemDevicesDeltaResponse) {
    m := &ItemDevicesDeltaResponse{
        ItemDevicesDeltaGetResponse: *NewItemDevicesDeltaGetResponse(),
    }
    return m
}
// CreateItemDevicesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemDevicesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemDevicesDeltaResponse(), nil
}
// ItemDevicesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemDevicesDeltaResponseable interface {
    ItemDevicesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
