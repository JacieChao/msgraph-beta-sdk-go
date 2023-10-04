package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagedDevicesAppDiagnosticsWithUpnResponse 
// Deprecated: This class is obsolete. Use appDiagnosticsWithUpnGetResponse instead.
type ComanagedDevicesAppDiagnosticsWithUpnResponse struct {
    ComanagedDevicesAppDiagnosticsWithUpnGetResponse
}
// NewComanagedDevicesAppDiagnosticsWithUpnResponse instantiates a new ComanagedDevicesAppDiagnosticsWithUpnResponse and sets the default values.
func NewComanagedDevicesAppDiagnosticsWithUpnResponse()(*ComanagedDevicesAppDiagnosticsWithUpnResponse) {
    m := &ComanagedDevicesAppDiagnosticsWithUpnResponse{
        ComanagedDevicesAppDiagnosticsWithUpnGetResponse: *NewComanagedDevicesAppDiagnosticsWithUpnGetResponse(),
    }
    return m
}
// CreateComanagedDevicesAppDiagnosticsWithUpnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesAppDiagnosticsWithUpnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesAppDiagnosticsWithUpnResponse(), nil
}
// ComanagedDevicesAppDiagnosticsWithUpnResponseable 
// Deprecated: This class is obsolete. Use appDiagnosticsWithUpnGetResponse instead.
type ComanagedDevicesAppDiagnosticsWithUpnResponseable interface {
    ComanagedDevicesAppDiagnosticsWithUpnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
