package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcProvisioningPolicyAssignment 
type CloudPcProvisioningPolicyAssignment struct {
    Entity
}
// NewCloudPcProvisioningPolicyAssignment instantiates a new cloudPcProvisioningPolicyAssignment and sets the default values.
func NewCloudPcProvisioningPolicyAssignment()(*CloudPcProvisioningPolicyAssignment) {
    m := &CloudPcProvisioningPolicyAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcProvisioningPolicyAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcProvisioningPolicyAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcProvisioningPolicyAssignment(), nil
}
// GetAssignedUsers gets the assignedUsers property value. The assignment targeted users for the provisioning policy. This list of users is computed based on assignments, licenses, group memberships, and policies. This property is read-only. Supports$expand.
func (m *CloudPcProvisioningPolicyAssignment) GetAssignedUsers()([]Userable) {
    val, err := m.GetBackingStore().Get("assignedUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Userable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcProvisioningPolicyAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetAssignedUsers(res)
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
        val, err := n.GetObjectValue(CreateCloudPcManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(CloudPcManagementAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcProvisioningPolicyAssignment) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTarget gets the target property value. The assignment target for the provisioning policy. Currently, the only target supported for this policy is a user group. For details, see cloudPcManagementGroupAssignmentTarget.
func (m *CloudPcProvisioningPolicyAssignment) GetTarget()(CloudPcManagementAssignmentTargetable) {
    val, err := m.GetBackingStore().Get("target")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcManagementAssignmentTargetable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcProvisioningPolicyAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignedUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedUsers()))
        for i, v := range m.GetAssignedUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignedUsers", cast)
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
// SetAssignedUsers sets the assignedUsers property value. The assignment targeted users for the provisioning policy. This list of users is computed based on assignments, licenses, group memberships, and policies. This property is read-only. Supports$expand.
func (m *CloudPcProvisioningPolicyAssignment) SetAssignedUsers(value []Userable)() {
    err := m.GetBackingStore().Set("assignedUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcProvisioningPolicyAssignment) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTarget sets the target property value. The assignment target for the provisioning policy. Currently, the only target supported for this policy is a user group. For details, see cloudPcManagementGroupAssignmentTarget.
func (m *CloudPcProvisioningPolicyAssignment) SetTarget(value CloudPcManagementAssignmentTargetable)() {
    err := m.GetBackingStore().Set("target", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcProvisioningPolicyAssignmentable 
type CloudPcProvisioningPolicyAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedUsers()([]Userable)
    GetOdataType()(*string)
    GetTarget()(CloudPcManagementAssignmentTargetable)
    SetAssignedUsers(value []Userable)()
    SetOdataType(value *string)()
    SetTarget(value CloudPcManagementAssignmentTargetable)()
}
