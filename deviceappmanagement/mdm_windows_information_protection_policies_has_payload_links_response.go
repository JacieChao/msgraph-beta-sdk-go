package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MdmWindowsInformationProtectionPoliciesHasPayloadLinksResponse 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type MdmWindowsInformationProtectionPoliciesHasPayloadLinksResponse struct {
    MdmWindowsInformationProtectionPoliciesHasPayloadLinksPostResponse
}
// NewMdmWindowsInformationProtectionPoliciesHasPayloadLinksResponse instantiates a new MdmWindowsInformationProtectionPoliciesHasPayloadLinksResponse and sets the default values.
func NewMdmWindowsInformationProtectionPoliciesHasPayloadLinksResponse()(*MdmWindowsInformationProtectionPoliciesHasPayloadLinksResponse) {
    m := &MdmWindowsInformationProtectionPoliciesHasPayloadLinksResponse{
        MdmWindowsInformationProtectionPoliciesHasPayloadLinksPostResponse: *NewMdmWindowsInformationProtectionPoliciesHasPayloadLinksPostResponse(),
    }
    return m
}
// CreateMdmWindowsInformationProtectionPoliciesHasPayloadLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMdmWindowsInformationProtectionPoliciesHasPayloadLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMdmWindowsInformationProtectionPoliciesHasPayloadLinksResponse(), nil
}
// MdmWindowsInformationProtectionPoliciesHasPayloadLinksResponseable 
// Deprecated: This class is obsolete. Use hasPayloadLinksPostResponse instead.
type MdmWindowsInformationProtectionPoliciesHasPayloadLinksResponseable interface {
    MdmWindowsInformationProtectionPoliciesHasPayloadLinksPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
