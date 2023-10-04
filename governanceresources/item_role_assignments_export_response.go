package governanceresources

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemRoleAssignmentsExportResponse 
// Deprecated: This class is obsolete. Use exportGetResponse instead.
type ItemRoleAssignmentsExportResponse struct {
    ItemRoleAssignmentsExportGetResponse
}
// NewItemRoleAssignmentsExportResponse instantiates a new ItemRoleAssignmentsExportResponse and sets the default values.
func NewItemRoleAssignmentsExportResponse()(*ItemRoleAssignmentsExportResponse) {
    m := &ItemRoleAssignmentsExportResponse{
        ItemRoleAssignmentsExportGetResponse: *NewItemRoleAssignmentsExportGetResponse(),
    }
    return m
}
// CreateItemRoleAssignmentsExportResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemRoleAssignmentsExportResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemRoleAssignmentsExportResponse(), nil
}
// ItemRoleAssignmentsExportResponseable 
// Deprecated: This class is obsolete. Use exportGetResponse instead.
type ItemRoleAssignmentsExportResponseable interface {
    ItemRoleAssignmentsExportGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
