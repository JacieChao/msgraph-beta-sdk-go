package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaDeploymentsItemCancelResponse 
// Deprecated: This class is obsolete. Use cancelPostResponse instead.
type ZebraFotaDeploymentsItemCancelResponse struct {
    ZebraFotaDeploymentsItemCancelPostResponse
}
// NewZebraFotaDeploymentsItemCancelResponse instantiates a new ZebraFotaDeploymentsItemCancelResponse and sets the default values.
func NewZebraFotaDeploymentsItemCancelResponse()(*ZebraFotaDeploymentsItemCancelResponse) {
    m := &ZebraFotaDeploymentsItemCancelResponse{
        ZebraFotaDeploymentsItemCancelPostResponse: *NewZebraFotaDeploymentsItemCancelPostResponse(),
    }
    return m
}
// CreateZebraFotaDeploymentsItemCancelResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateZebraFotaDeploymentsItemCancelResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaDeploymentsItemCancelResponse(), nil
}
// ZebraFotaDeploymentsItemCancelResponseable 
// Deprecated: This class is obsolete. Use cancelPostResponse instead.
type ZebraFotaDeploymentsItemCancelResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ZebraFotaDeploymentsItemCancelPostResponseable
}
