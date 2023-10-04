package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyMigrationReportsItemUpdateScopeTagsResponse 
// Deprecated: This class is obsolete. Use updateScopeTagsPostResponse instead.
type GroupPolicyMigrationReportsItemUpdateScopeTagsResponse struct {
    GroupPolicyMigrationReportsItemUpdateScopeTagsPostResponse
}
// NewGroupPolicyMigrationReportsItemUpdateScopeTagsResponse instantiates a new GroupPolicyMigrationReportsItemUpdateScopeTagsResponse and sets the default values.
func NewGroupPolicyMigrationReportsItemUpdateScopeTagsResponse()(*GroupPolicyMigrationReportsItemUpdateScopeTagsResponse) {
    m := &GroupPolicyMigrationReportsItemUpdateScopeTagsResponse{
        GroupPolicyMigrationReportsItemUpdateScopeTagsPostResponse: *NewGroupPolicyMigrationReportsItemUpdateScopeTagsPostResponse(),
    }
    return m
}
// CreateGroupPolicyMigrationReportsItemUpdateScopeTagsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyMigrationReportsItemUpdateScopeTagsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyMigrationReportsItemUpdateScopeTagsResponse(), nil
}
// GroupPolicyMigrationReportsItemUpdateScopeTagsResponseable 
// Deprecated: This class is obsolete. Use updateScopeTagsPostResponse instead.
type GroupPolicyMigrationReportsItemUpdateScopeTagsResponseable interface {
    GroupPolicyMigrationReportsItemUpdateScopeTagsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
