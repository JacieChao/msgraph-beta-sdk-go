package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HasPayloadLinkResultItemable 
type HasPayloadLinkResultItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetError()(*string)
    GetHasLink()(*bool)
    GetOdataType()(*string)
    GetPayloadId()(*string)
    GetSources()([]DeviceAndAppManagementAssignmentSource)
    SetError(value *string)()
    SetHasLink(value *bool)()
    SetOdataType(value *string)()
    SetPayloadId(value *string)()
    SetSources(value []DeviceAndAppManagementAssignmentSource)()
}
