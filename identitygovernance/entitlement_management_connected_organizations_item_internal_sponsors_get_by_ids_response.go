package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponse 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponse struct {
    EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostResponse
}
// NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponse instantiates a new EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponse and sets the default values.
func NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponse()(*EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponse) {
    m := &EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponse{
        EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostResponse: *NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostResponse(),
    }
    return m
}
// CreateEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponse(), nil
}
// EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponseable 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsResponseable interface {
    EntitlementManagementConnectedOrganizationsItemInternalSponsorsGetByIdsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
