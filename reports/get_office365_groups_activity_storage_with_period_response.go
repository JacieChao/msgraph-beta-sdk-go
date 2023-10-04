package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetOffice365GroupsActivityStorageWithPeriodResponse 
// Deprecated: This class is obsolete. Use getOffice365GroupsActivityStorageWithPeriodGetResponse instead.
type GetOffice365GroupsActivityStorageWithPeriodResponse struct {
    GetOffice365GroupsActivityStorageWithPeriodGetResponse
}
// NewGetOffice365GroupsActivityStorageWithPeriodResponse instantiates a new GetOffice365GroupsActivityStorageWithPeriodResponse and sets the default values.
func NewGetOffice365GroupsActivityStorageWithPeriodResponse()(*GetOffice365GroupsActivityStorageWithPeriodResponse) {
    m := &GetOffice365GroupsActivityStorageWithPeriodResponse{
        GetOffice365GroupsActivityStorageWithPeriodGetResponse: *NewGetOffice365GroupsActivityStorageWithPeriodGetResponse(),
    }
    return m
}
// CreateGetOffice365GroupsActivityStorageWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetOffice365GroupsActivityStorageWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetOffice365GroupsActivityStorageWithPeriodResponse(), nil
}
// GetOffice365GroupsActivityStorageWithPeriodResponseable 
// Deprecated: This class is obsolete. Use getOffice365GroupsActivityStorageWithPeriodGetResponse instead.
type GetOffice365GroupsActivityStorageWithPeriodResponseable interface {
    GetOffice365GroupsActivityStorageWithPeriodGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
