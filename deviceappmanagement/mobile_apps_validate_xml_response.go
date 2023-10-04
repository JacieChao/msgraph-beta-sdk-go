package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppsValidateXmlResponse 
// Deprecated: This class is obsolete. Use validateXmlPostResponse instead.
type MobileAppsValidateXmlResponse struct {
    MobileAppsValidateXmlPostResponse
}
// NewMobileAppsValidateXmlResponse instantiates a new MobileAppsValidateXmlResponse and sets the default values.
func NewMobileAppsValidateXmlResponse()(*MobileAppsValidateXmlResponse) {
    m := &MobileAppsValidateXmlResponse{
        MobileAppsValidateXmlPostResponse: *NewMobileAppsValidateXmlPostResponse(),
    }
    return m
}
// CreateMobileAppsValidateXmlResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppsValidateXmlResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppsValidateXmlResponse(), nil
}
// MobileAppsValidateXmlResponseable 
// Deprecated: This class is obsolete. Use validateXmlPostResponse instead.
type MobileAppsValidateXmlResponseable interface {
    MobileAppsValidateXmlPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
