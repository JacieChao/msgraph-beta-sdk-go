package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type EntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaResponse struct {
    EntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaGetResponse
}
// NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaResponse instantiates a new EntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaResponse and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaResponse()(*EntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaResponse) {
    m := &EntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaResponse{
        EntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaGetResponse: *NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaGetResponse(),
    }
    return m
}
// CreateEntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaResponse(), nil
}
// EntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type EntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaResponseable interface {
    EntitlementManagementConnectedOrganizationsItemExternalSponsorsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
