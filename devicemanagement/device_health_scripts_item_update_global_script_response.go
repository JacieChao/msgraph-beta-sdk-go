package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptsItemUpdateGlobalScriptResponse 
// Deprecated: This class is obsolete. Use updateGlobalScriptPostResponse instead.
type DeviceHealthScriptsItemUpdateGlobalScriptResponse struct {
    DeviceHealthScriptsItemUpdateGlobalScriptPostResponse
}
// NewDeviceHealthScriptsItemUpdateGlobalScriptResponse instantiates a new DeviceHealthScriptsItemUpdateGlobalScriptResponse and sets the default values.
func NewDeviceHealthScriptsItemUpdateGlobalScriptResponse()(*DeviceHealthScriptsItemUpdateGlobalScriptResponse) {
    m := &DeviceHealthScriptsItemUpdateGlobalScriptResponse{
        DeviceHealthScriptsItemUpdateGlobalScriptPostResponse: *NewDeviceHealthScriptsItemUpdateGlobalScriptPostResponse(),
    }
    return m
}
// CreateDeviceHealthScriptsItemUpdateGlobalScriptResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptsItemUpdateGlobalScriptResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptsItemUpdateGlobalScriptResponse(), nil
}
// DeviceHealthScriptsItemUpdateGlobalScriptResponseable 
// Deprecated: This class is obsolete. Use updateGlobalScriptPostResponse instead.
type DeviceHealthScriptsItemUpdateGlobalScriptResponseable interface {
    DeviceHealthScriptsItemUpdateGlobalScriptPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
