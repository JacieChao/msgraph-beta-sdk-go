package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaConnectorConnectResponse 
// Deprecated: This class is obsolete. Use connectPostResponse instead.
type ZebraFotaConnectorConnectResponse struct {
    ZebraFotaConnectorConnectPostResponse
}
// NewZebraFotaConnectorConnectResponse instantiates a new ZebraFotaConnectorConnectResponse and sets the default values.
func NewZebraFotaConnectorConnectResponse()(*ZebraFotaConnectorConnectResponse) {
    m := &ZebraFotaConnectorConnectResponse{
        ZebraFotaConnectorConnectPostResponse: *NewZebraFotaConnectorConnectPostResponse(),
    }
    return m
}
// CreateZebraFotaConnectorConnectResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateZebraFotaConnectorConnectResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaConnectorConnectResponse(), nil
}
// ZebraFotaConnectorConnectResponseable 
// Deprecated: This class is obsolete. Use connectPostResponse instead.
type ZebraFotaConnectorConnectResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ZebraFotaConnectorConnectPostResponseable
}
