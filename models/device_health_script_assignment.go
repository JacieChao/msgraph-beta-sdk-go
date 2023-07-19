package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptAssignment contains properties used to assign a device management script to a group.
type DeviceHealthScriptAssignment struct {
    Entity
}
// NewDeviceHealthScriptAssignment instantiates a new deviceHealthScriptAssignment and sets the default values.
func NewDeviceHealthScriptAssignment()(*DeviceHealthScriptAssignment) {
    m := &DeviceHealthScriptAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceHealthScriptAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["runRemediationScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunRemediationScript(val)
        }
        return nil
    }
    res["runSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceHealthScriptRunScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunSchedule(val.(DeviceHealthScriptRunScheduleable))
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(DeviceAndAppManagementAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptAssignment) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRunRemediationScript gets the runRemediationScript property value. Determine whether we want to run detection script only or run both detection script and remediation script
func (m *DeviceHealthScriptAssignment) GetRunRemediationScript()(*bool) {
    val, err := m.GetBackingStore().Get("runRemediationScript")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRunSchedule gets the runSchedule property value. Script run schedule for the target group
func (m *DeviceHealthScriptAssignment) GetRunSchedule()(DeviceHealthScriptRunScheduleable) {
    val, err := m.GetBackingStore().Get("runSchedule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceHealthScriptRunScheduleable)
    }
    return nil
}
// GetTarget gets the target property value. The Azure Active Directory group we are targeting the script to
func (m *DeviceHealthScriptAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    val, err := m.GetBackingStore().Get("target")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceAndAppManagementAssignmentTargetable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("runRemediationScript", m.GetRunRemediationScript())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("runSchedule", m.GetRunSchedule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptAssignment) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRunRemediationScript sets the runRemediationScript property value. Determine whether we want to run detection script only or run both detection script and remediation script
func (m *DeviceHealthScriptAssignment) SetRunRemediationScript(value *bool)() {
    err := m.GetBackingStore().Set("runRemediationScript", value)
    if err != nil {
        panic(err)
    }
}
// SetRunSchedule sets the runSchedule property value. Script run schedule for the target group
func (m *DeviceHealthScriptAssignment) SetRunSchedule(value DeviceHealthScriptRunScheduleable)() {
    err := m.GetBackingStore().Set("runSchedule", value)
    if err != nil {
        panic(err)
    }
}
// SetTarget sets the target property value. The Azure Active Directory group we are targeting the script to
func (m *DeviceHealthScriptAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    err := m.GetBackingStore().Set("target", value)
    if err != nil {
        panic(err)
    }
}
// DeviceHealthScriptAssignmentable 
type DeviceHealthScriptAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetRunRemediationScript()(*bool)
    GetRunSchedule()(DeviceHealthScriptRunScheduleable)
    GetTarget()(DeviceAndAppManagementAssignmentTargetable)
    SetOdataType(value *string)()
    SetRunRemediationScript(value *bool)()
    SetRunSchedule(value DeviceHealthScriptRunScheduleable)()
    SetTarget(value DeviceAndAppManagementAssignmentTargetable)()
}
