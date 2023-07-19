package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedEBookAssignment contains properties used to assign a eBook to a group.
type ManagedEBookAssignment struct {
    Entity
}
// NewManagedEBookAssignment instantiates a new managedEBookAssignment and sets the default values.
func NewManagedEBookAssignment()(*ManagedEBookAssignment) {
    m := &ManagedEBookAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedEBookAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedEBookAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.iosVppEBookAssignment":
                        return NewIosVppEBookAssignment(), nil
                }
            }
        }
    }
    return NewManagedEBookAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedEBookAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["installIntent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInstallIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallIntent(val.(*InstallIntent))
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
// GetInstallIntent gets the installIntent property value. Possible values for the install intent chosen by the admin.
func (m *ManagedEBookAssignment) GetInstallIntent()(*InstallIntent) {
    val, err := m.GetBackingStore().Get("installIntent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*InstallIntent)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ManagedEBookAssignment) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTarget gets the target property value. The assignment target for eBook.
func (m *ManagedEBookAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
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
func (m *ManagedEBookAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetInstallIntent() != nil {
        cast := (*m.GetInstallIntent()).String()
        err = writer.WriteStringValue("installIntent", &cast)
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
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInstallIntent sets the installIntent property value. Possible values for the install intent chosen by the admin.
func (m *ManagedEBookAssignment) SetInstallIntent(value *InstallIntent)() {
    err := m.GetBackingStore().Set("installIntent", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagedEBookAssignment) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTarget sets the target property value. The assignment target for eBook.
func (m *ManagedEBookAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    err := m.GetBackingStore().Set("target", value)
    if err != nil {
        panic(err)
    }
}
// ManagedEBookAssignmentable 
type ManagedEBookAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInstallIntent()(*InstallIntent)
    GetOdataType()(*string)
    GetTarget()(DeviceAndAppManagementAssignmentTargetable)
    SetInstallIntent(value *InstallIntent)()
    SetOdataType(value *string)()
    SetTarget(value DeviceAndAppManagementAssignmentTargetable)()
}
