package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntentsItemGetCustomizedSettingsResponse 
// Deprecated: This class is obsolete. Use getCustomizedSettingsGetResponse instead.
type IntentsItemGetCustomizedSettingsResponse struct {
    IntentsItemGetCustomizedSettingsGetResponse
}
// NewIntentsItemGetCustomizedSettingsResponse instantiates a new IntentsItemGetCustomizedSettingsResponse and sets the default values.
func NewIntentsItemGetCustomizedSettingsResponse()(*IntentsItemGetCustomizedSettingsResponse) {
    m := &IntentsItemGetCustomizedSettingsResponse{
        IntentsItemGetCustomizedSettingsGetResponse: *NewIntentsItemGetCustomizedSettingsGetResponse(),
    }
    return m
}
// CreateIntentsItemGetCustomizedSettingsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntentsItemGetCustomizedSettingsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntentsItemGetCustomizedSettingsResponse(), nil
}
// IntentsItemGetCustomizedSettingsResponseable 
// Deprecated: This class is obsolete. Use getCustomizedSettingsGetResponse instead.
type IntentsItemGetCustomizedSettingsResponseable interface {
    IntentsItemGetCustomizedSettingsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
