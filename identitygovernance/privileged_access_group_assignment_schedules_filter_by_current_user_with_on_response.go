package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnResponse 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type PrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnResponse struct {
    PrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnGetResponse
}
// NewPrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnResponse instantiates a new PrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnResponse and sets the default values.
func NewPrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnResponse()(*PrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnResponse) {
    m := &PrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnResponse{
        PrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnGetResponse: *NewPrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnGetResponse(),
    }
    return m
}
// CreatePrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnResponse(), nil
}
// PrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnResponseable 
// Deprecated: This class is obsolete. Use filterByCurrentUserWithOnGetResponse instead.
type PrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PrivilegedAccessGroupAssignmentSchedulesFilterByCurrentUserWithOnGetResponseable
}
