package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementAccessPackageCatalogsSearchResponse 
// Deprecated: This class is obsolete. Use SearchGetResponse instead.
type EntitlementManagementAccessPackageCatalogsSearchResponse struct {
    EntitlementManagementAccessPackageCatalogsSearchGetResponse
}
// NewEntitlementManagementAccessPackageCatalogsSearchResponse instantiates a new EntitlementManagementAccessPackageCatalogsSearchResponse and sets the default values.
func NewEntitlementManagementAccessPackageCatalogsSearchResponse()(*EntitlementManagementAccessPackageCatalogsSearchResponse) {
    m := &EntitlementManagementAccessPackageCatalogsSearchResponse{
        EntitlementManagementAccessPackageCatalogsSearchGetResponse: *NewEntitlementManagementAccessPackageCatalogsSearchGetResponse(),
    }
    return m
}
// CreateEntitlementManagementAccessPackageCatalogsSearchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementAccessPackageCatalogsSearchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementAccessPackageCatalogsSearchResponse(), nil
}
// EntitlementManagementAccessPackageCatalogsSearchResponseable 
// Deprecated: This class is obsolete. Use SearchGetResponse instead.
type EntitlementManagementAccessPackageCatalogsSearchResponseable interface {
    EntitlementManagementAccessPackageCatalogsSearchGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
