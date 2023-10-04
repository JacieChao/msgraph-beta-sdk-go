package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TiIndicatorsSubmitTiIndicatorsResponse 
// Deprecated: This class is obsolete. Use submitTiIndicatorsPostResponse instead.
type TiIndicatorsSubmitTiIndicatorsResponse struct {
    TiIndicatorsSubmitTiIndicatorsPostResponse
}
// NewTiIndicatorsSubmitTiIndicatorsResponse instantiates a new TiIndicatorsSubmitTiIndicatorsResponse and sets the default values.
func NewTiIndicatorsSubmitTiIndicatorsResponse()(*TiIndicatorsSubmitTiIndicatorsResponse) {
    m := &TiIndicatorsSubmitTiIndicatorsResponse{
        TiIndicatorsSubmitTiIndicatorsPostResponse: *NewTiIndicatorsSubmitTiIndicatorsPostResponse(),
    }
    return m
}
// CreateTiIndicatorsSubmitTiIndicatorsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTiIndicatorsSubmitTiIndicatorsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTiIndicatorsSubmitTiIndicatorsResponse(), nil
}
// TiIndicatorsSubmitTiIndicatorsResponseable 
// Deprecated: This class is obsolete. Use submitTiIndicatorsPostResponse instead.
type TiIndicatorsSubmitTiIndicatorsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TiIndicatorsSubmitTiIndicatorsPostResponseable
}
