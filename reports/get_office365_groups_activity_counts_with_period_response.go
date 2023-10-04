package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetOffice365GroupsActivityCountsWithPeriodResponse 
// Deprecated: This class is obsolete. Use getOffice365GroupsActivityCountsWithPeriodGetResponse instead.
type GetOffice365GroupsActivityCountsWithPeriodResponse struct {
    GetOffice365GroupsActivityCountsWithPeriodGetResponse
}
// NewGetOffice365GroupsActivityCountsWithPeriodResponse instantiates a new GetOffice365GroupsActivityCountsWithPeriodResponse and sets the default values.
func NewGetOffice365GroupsActivityCountsWithPeriodResponse()(*GetOffice365GroupsActivityCountsWithPeriodResponse) {
    m := &GetOffice365GroupsActivityCountsWithPeriodResponse{
        GetOffice365GroupsActivityCountsWithPeriodGetResponse: *NewGetOffice365GroupsActivityCountsWithPeriodGetResponse(),
    }
    return m
}
// CreateGetOffice365GroupsActivityCountsWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetOffice365GroupsActivityCountsWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetOffice365GroupsActivityCountsWithPeriodResponse(), nil
}
// GetOffice365GroupsActivityCountsWithPeriodResponseable 
// Deprecated: This class is obsolete. Use getOffice365GroupsActivityCountsWithPeriodGetResponse instead.
type GetOffice365GroupsActivityCountsWithPeriodResponseable interface {
    GetOffice365GroupsActivityCountsWithPeriodGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
