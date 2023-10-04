package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeploymentProfilesHasPayloadLinksResponse 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type WindowsAutopilotDeploymentProfilesHasPayloadLinksResponse struct {
    WindowsAutopilotDeploymentProfilesHasPayloadLinksPostResponse
}
// NewWindowsAutopilotDeploymentProfilesHasPayloadLinksResponse instantiates a new WindowsAutopilotDeploymentProfilesHasPayloadLinksResponse and sets the default values.
func NewWindowsAutopilotDeploymentProfilesHasPayloadLinksResponse()(*WindowsAutopilotDeploymentProfilesHasPayloadLinksResponse) {
    m := &WindowsAutopilotDeploymentProfilesHasPayloadLinksResponse{
        WindowsAutopilotDeploymentProfilesHasPayloadLinksPostResponse: *NewWindowsAutopilotDeploymentProfilesHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateWindowsAutopilotDeploymentProfilesHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeploymentProfilesHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeploymentProfilesHasPayloadLinksResponse(), nil
}
// WindowsAutopilotDeploymentProfilesHasPayloadLinksResponseable 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type WindowsAutopilotDeploymentProfilesHasPayloadLinksResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsAutopilotDeploymentProfilesHasPayloadLinksPostResponseable
}
