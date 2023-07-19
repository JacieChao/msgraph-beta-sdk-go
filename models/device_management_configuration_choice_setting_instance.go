package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingInstance setting instance within policy
type DeviceManagementConfigurationChoiceSettingInstance struct {
    DeviceManagementConfigurationSettingInstance
}
// NewDeviceManagementConfigurationChoiceSettingInstance instantiates a new deviceManagementConfigurationChoiceSettingInstance and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingInstance()(*DeviceManagementConfigurationChoiceSettingInstance) {
    m := &DeviceManagementConfigurationChoiceSettingInstance{
        DeviceManagementConfigurationSettingInstance: *NewDeviceManagementConfigurationSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationChoiceSettingInstance"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingInstance(), nil
}
// GetChoiceSettingValue gets the choiceSettingValue property value. The choiceSettingValue property
func (m *DeviceManagementConfigurationChoiceSettingInstance) GetChoiceSettingValue()(DeviceManagementConfigurationChoiceSettingValueable) {
    val, err := m.GetBackingStore().Get("choiceSettingValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationChoiceSettingValueable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstance.GetFieldDeserializers()
    res["choiceSettingValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationChoiceSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChoiceSettingValue(val.(DeviceManagementConfigurationChoiceSettingValueable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationChoiceSettingInstance) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("choiceSettingValue", m.GetChoiceSettingValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChoiceSettingValue sets the choiceSettingValue property value. The choiceSettingValue property
func (m *DeviceManagementConfigurationChoiceSettingInstance) SetChoiceSettingValue(value DeviceManagementConfigurationChoiceSettingValueable)() {
    err := m.GetBackingStore().Set("choiceSettingValue", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationChoiceSettingInstance) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationChoiceSettingInstanceable 
type DeviceManagementConfigurationChoiceSettingInstanceable interface {
    DeviceManagementConfigurationSettingInstanceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChoiceSettingValue()(DeviceManagementConfigurationChoiceSettingValueable)
    GetOdataType()(*string)
    SetChoiceSettingValue(value DeviceManagementConfigurationChoiceSettingValueable)()
    SetOdataType(value *string)()
}
