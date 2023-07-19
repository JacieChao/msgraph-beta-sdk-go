package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidFotaDeploymentAssignmentTarget the AAD Group we are deploying firmware updates to
type AndroidFotaDeploymentAssignmentTarget struct {
    DeviceAndAppManagementAssignmentTarget
    // The OdataType property
    OdataType *string
}
// NewAndroidFotaDeploymentAssignmentTarget instantiates a new androidFotaDeploymentAssignmentTarget and sets the default values.
func NewAndroidFotaDeploymentAssignmentTarget()(*AndroidFotaDeploymentAssignmentTarget) {
    m := &AndroidFotaDeploymentAssignmentTarget{
        DeviceAndAppManagementAssignmentTarget: *NewDeviceAndAppManagementAssignmentTarget(),
    }
    odataTypeValue := "#microsoft.graph.androidFotaDeploymentAssignmentTarget"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidFotaDeploymentAssignmentTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidFotaDeploymentAssignmentTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidFotaDeploymentAssignmentTarget(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidFotaDeploymentAssignmentTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceAndAppManagementAssignmentTarget.GetFieldDeserializers()
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. AAD Group Id.
func (m *AndroidFotaDeploymentAssignmentTarget) GetGroupId()(*string) {
    val, err := m.GetBackingStore().Get("groupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidFotaDeploymentAssignmentTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceAndAppManagementAssignmentTarget.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupId sets the groupId property value. AAD Group Id.
func (m *AndroidFotaDeploymentAssignmentTarget) SetGroupId(value *string)() {
    err := m.GetBackingStore().Set("groupId", value)
    if err != nil {
        panic(err)
    }
}
// AndroidFotaDeploymentAssignmentTargetable 
type AndroidFotaDeploymentAssignmentTargetable interface {
    DeviceAndAppManagementAssignmentTargetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupId()(*string)
    SetGroupId(value *string)()
}
