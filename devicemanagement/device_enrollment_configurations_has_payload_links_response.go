package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentConfigurationsHasPayloadLinksResponse 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type DeviceEnrollmentConfigurationsHasPayloadLinksResponse struct {
    DeviceEnrollmentConfigurationsHasPayloadLinksPostResponse
}
// NewDeviceEnrollmentConfigurationsHasPayloadLinksResponse instantiates a new DeviceEnrollmentConfigurationsHasPayloadLinksResponse and sets the default values.
func NewDeviceEnrollmentConfigurationsHasPayloadLinksResponse()(*DeviceEnrollmentConfigurationsHasPayloadLinksResponse) {
    m := &DeviceEnrollmentConfigurationsHasPayloadLinksResponse{
        DeviceEnrollmentConfigurationsHasPayloadLinksPostResponse: *NewDeviceEnrollmentConfigurationsHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateDeviceEnrollmentConfigurationsHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentConfigurationsHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentConfigurationsHasPayloadLinksResponse(), nil
}
// DeviceEnrollmentConfigurationsHasPayloadLinksResponseable 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type DeviceEnrollmentConfigurationsHasPayloadLinksResponseable interface {
    DeviceEnrollmentConfigurationsHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
