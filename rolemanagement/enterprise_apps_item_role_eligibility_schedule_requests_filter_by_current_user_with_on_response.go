package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type EnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse struct {
    EnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponse
}
// NewEnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse instantiates a new EnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse and sets the default values.
func NewEnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse()(*EnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse) {
    m := &EnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse{
        EnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponse: *NewEnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreateEnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponse(), nil
}
// EnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseable 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type EnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnResponseable interface {
    EnterpriseAppsItemRoleEligibilityScheduleRequestsFilterByCurrentUserWithOnGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
