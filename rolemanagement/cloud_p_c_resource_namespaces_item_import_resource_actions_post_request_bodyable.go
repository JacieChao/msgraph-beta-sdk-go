package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPCResourceNamespacesItemImportResourceActionsPostRequestBodyable 
type CloudPCResourceNamespacesItemImportResourceActionsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormat()(*string)
    GetOverwriteResourceNamespace()(*bool)
    GetValue()(*string)
    SetFormat(value *string)()
    SetOverwriteResourceNamespace(value *bool)()
    SetValue(value *string)()
}