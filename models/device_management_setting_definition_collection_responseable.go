package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingDefinitionCollectionResponseable 
type DeviceManagementSettingDefinitionCollectionResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNextLink()(*string)
    GetValue()([]DeviceManagementSettingDefinitionable)
    SetNextLink(value *string)()
    SetValue(value []DeviceManagementSettingDefinitionable)()
}