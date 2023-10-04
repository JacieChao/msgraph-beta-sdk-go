package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevicesAppDiagnosticsWithUpnResponse 
// Deprecated: This class is obsolete. Use appDiagnosticsWithUpnGetResponse instead.
type ManagedDevicesAppDiagnosticsWithUpnResponse struct {
    ManagedDevicesAppDiagnosticsWithUpnGetResponse
}
// NewManagedDevicesAppDiagnosticsWithUpnResponse instantiates a new ManagedDevicesAppDiagnosticsWithUpnResponse and sets the default values.
func NewManagedDevicesAppDiagnosticsWithUpnResponse()(*ManagedDevicesAppDiagnosticsWithUpnResponse) {
    m := &ManagedDevicesAppDiagnosticsWithUpnResponse{
        ManagedDevicesAppDiagnosticsWithUpnGetResponse: *NewManagedDevicesAppDiagnosticsWithUpnGetResponse(),
    }
    return m
}
// CreateManagedDevicesAppDiagnosticsWithUpnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDevicesAppDiagnosticsWithUpnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesAppDiagnosticsWithUpnResponse(), nil
}
// ManagedDevicesAppDiagnosticsWithUpnResponseable 
// Deprecated: This class is obsolete. Use appDiagnosticsWithUpnGetResponse instead.
type ManagedDevicesAppDiagnosticsWithUpnResponseable interface {
    ManagedDevicesAppDiagnosticsWithUpnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
