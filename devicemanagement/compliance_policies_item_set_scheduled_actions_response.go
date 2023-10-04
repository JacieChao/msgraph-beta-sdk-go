package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CompliancePoliciesItemSetScheduledActionsResponse 
// Deprecated: This class is obsolete. Use setScheduledActionsPostResponse instead.
type CompliancePoliciesItemSetScheduledActionsResponse struct {
    CompliancePoliciesItemSetScheduledActionsPostResponse
}
// NewCompliancePoliciesItemSetScheduledActionsResponse instantiates a new CompliancePoliciesItemSetScheduledActionsResponse and sets the default values.
func NewCompliancePoliciesItemSetScheduledActionsResponse()(*CompliancePoliciesItemSetScheduledActionsResponse) {
    m := &CompliancePoliciesItemSetScheduledActionsResponse{
        CompliancePoliciesItemSetScheduledActionsPostResponse: *NewCompliancePoliciesItemSetScheduledActionsPostResponse(),
    }
    return m
}
// CreateCompliancePoliciesItemSetScheduledActionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCompliancePoliciesItemSetScheduledActionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompliancePoliciesItemSetScheduledActionsResponse(), nil
}
// CompliancePoliciesItemSetScheduledActionsResponseable 
// Deprecated: This class is obsolete. Use setScheduledActionsPostResponse instead.
type CompliancePoliciesItemSetScheduledActionsResponseable interface {
    CompliancePoliciesItemSetScheduledActionsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
