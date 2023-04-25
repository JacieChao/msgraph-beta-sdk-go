package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessGroup 
type PrivilegedAccessGroup struct {
    Entity
}
// NewPrivilegedAccessGroup instantiates a new privilegedAccessGroup and sets the default values.
func NewPrivilegedAccessGroup()(*PrivilegedAccessGroup) {
    m := &PrivilegedAccessGroup{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedAccessGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessGroup(), nil
}
// GetAssignmentApprovals gets the assignmentApprovals property value. The assignmentApprovals property
func (m *PrivilegedAccessGroup) GetAssignmentApprovals()([]Approvalable) {
    val, err := m.GetBackingStore().Get("assignmentApprovals")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Approvalable)
    }
    return nil
}
// GetAssignmentScheduleInstances gets the assignmentScheduleInstances property value. The assignmentScheduleInstances property
func (m *PrivilegedAccessGroup) GetAssignmentScheduleInstances()([]PrivilegedAccessGroupAssignmentScheduleInstanceable) {
    val, err := m.GetBackingStore().Get("assignmentScheduleInstances")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrivilegedAccessGroupAssignmentScheduleInstanceable)
    }
    return nil
}
// GetAssignmentScheduleRequests gets the assignmentScheduleRequests property value. The assignmentScheduleRequests property
func (m *PrivilegedAccessGroup) GetAssignmentScheduleRequests()([]PrivilegedAccessGroupAssignmentScheduleRequestable) {
    val, err := m.GetBackingStore().Get("assignmentScheduleRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrivilegedAccessGroupAssignmentScheduleRequestable)
    }
    return nil
}
// GetAssignmentSchedules gets the assignmentSchedules property value. The assignmentSchedules property
func (m *PrivilegedAccessGroup) GetAssignmentSchedules()([]PrivilegedAccessGroupAssignmentScheduleable) {
    val, err := m.GetBackingStore().Get("assignmentSchedules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrivilegedAccessGroupAssignmentScheduleable)
    }
    return nil
}
// GetEligibilityScheduleInstances gets the eligibilityScheduleInstances property value. The eligibilityScheduleInstances property
func (m *PrivilegedAccessGroup) GetEligibilityScheduleInstances()([]PrivilegedAccessGroupEligibilityScheduleInstanceable) {
    val, err := m.GetBackingStore().Get("eligibilityScheduleInstances")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrivilegedAccessGroupEligibilityScheduleInstanceable)
    }
    return nil
}
// GetEligibilityScheduleRequests gets the eligibilityScheduleRequests property value. The eligibilityScheduleRequests property
func (m *PrivilegedAccessGroup) GetEligibilityScheduleRequests()([]PrivilegedAccessGroupEligibilityScheduleRequestable) {
    val, err := m.GetBackingStore().Get("eligibilityScheduleRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrivilegedAccessGroupEligibilityScheduleRequestable)
    }
    return nil
}
// GetEligibilitySchedules gets the eligibilitySchedules property value. The eligibilitySchedules property
func (m *PrivilegedAccessGroup) GetEligibilitySchedules()([]PrivilegedAccessGroupEligibilityScheduleable) {
    val, err := m.GetBackingStore().Get("eligibilitySchedules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrivilegedAccessGroupEligibilityScheduleable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccessGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentApprovals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Approvalable, len(val))
            for i, v := range val {
                res[i] = v.(Approvalable)
            }
            m.SetAssignmentApprovals(res)
        }
        return nil
    }
    res["assignmentScheduleInstances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegedAccessGroupAssignmentScheduleInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegedAccessGroupAssignmentScheduleInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(PrivilegedAccessGroupAssignmentScheduleInstanceable)
            }
            m.SetAssignmentScheduleInstances(res)
        }
        return nil
    }
    res["assignmentScheduleRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegedAccessGroupAssignmentScheduleRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegedAccessGroupAssignmentScheduleRequestable, len(val))
            for i, v := range val {
                res[i] = v.(PrivilegedAccessGroupAssignmentScheduleRequestable)
            }
            m.SetAssignmentScheduleRequests(res)
        }
        return nil
    }
    res["assignmentSchedules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegedAccessGroupAssignmentScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegedAccessGroupAssignmentScheduleable, len(val))
            for i, v := range val {
                res[i] = v.(PrivilegedAccessGroupAssignmentScheduleable)
            }
            m.SetAssignmentSchedules(res)
        }
        return nil
    }
    res["eligibilityScheduleInstances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegedAccessGroupEligibilityScheduleInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegedAccessGroupEligibilityScheduleInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(PrivilegedAccessGroupEligibilityScheduleInstanceable)
            }
            m.SetEligibilityScheduleInstances(res)
        }
        return nil
    }
    res["eligibilityScheduleRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegedAccessGroupEligibilityScheduleRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegedAccessGroupEligibilityScheduleRequestable, len(val))
            for i, v := range val {
                res[i] = v.(PrivilegedAccessGroupEligibilityScheduleRequestable)
            }
            m.SetEligibilityScheduleRequests(res)
        }
        return nil
    }
    res["eligibilitySchedules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegedAccessGroupEligibilityScheduleable, len(val))
            for i, v := range val {
                res[i] = v.(PrivilegedAccessGroupEligibilityScheduleable)
            }
            m.SetEligibilitySchedules(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PrivilegedAccessGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignmentApprovals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentApprovals()))
        for i, v := range m.GetAssignmentApprovals() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentApprovals", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentScheduleInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentScheduleInstances()))
        for i, v := range m.GetAssignmentScheduleInstances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentScheduleInstances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentScheduleRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentScheduleRequests()))
        for i, v := range m.GetAssignmentScheduleRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentScheduleRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentSchedules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentSchedules()))
        for i, v := range m.GetAssignmentSchedules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentSchedules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEligibilityScheduleInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEligibilityScheduleInstances()))
        for i, v := range m.GetEligibilityScheduleInstances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("eligibilityScheduleInstances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEligibilityScheduleRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEligibilityScheduleRequests()))
        for i, v := range m.GetEligibilityScheduleRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("eligibilityScheduleRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEligibilitySchedules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEligibilitySchedules()))
        for i, v := range m.GetEligibilitySchedules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("eligibilitySchedules", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignmentApprovals sets the assignmentApprovals property value. The assignmentApprovals property
func (m *PrivilegedAccessGroup) SetAssignmentApprovals(value []Approvalable)() {
    err := m.GetBackingStore().Set("assignmentApprovals", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentScheduleInstances sets the assignmentScheduleInstances property value. The assignmentScheduleInstances property
func (m *PrivilegedAccessGroup) SetAssignmentScheduleInstances(value []PrivilegedAccessGroupAssignmentScheduleInstanceable)() {
    err := m.GetBackingStore().Set("assignmentScheduleInstances", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentScheduleRequests sets the assignmentScheduleRequests property value. The assignmentScheduleRequests property
func (m *PrivilegedAccessGroup) SetAssignmentScheduleRequests(value []PrivilegedAccessGroupAssignmentScheduleRequestable)() {
    err := m.GetBackingStore().Set("assignmentScheduleRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentSchedules sets the assignmentSchedules property value. The assignmentSchedules property
func (m *PrivilegedAccessGroup) SetAssignmentSchedules(value []PrivilegedAccessGroupAssignmentScheduleable)() {
    err := m.GetBackingStore().Set("assignmentSchedules", value)
    if err != nil {
        panic(err)
    }
}
// SetEligibilityScheduleInstances sets the eligibilityScheduleInstances property value. The eligibilityScheduleInstances property
func (m *PrivilegedAccessGroup) SetEligibilityScheduleInstances(value []PrivilegedAccessGroupEligibilityScheduleInstanceable)() {
    err := m.GetBackingStore().Set("eligibilityScheduleInstances", value)
    if err != nil {
        panic(err)
    }
}
// SetEligibilityScheduleRequests sets the eligibilityScheduleRequests property value. The eligibilityScheduleRequests property
func (m *PrivilegedAccessGroup) SetEligibilityScheduleRequests(value []PrivilegedAccessGroupEligibilityScheduleRequestable)() {
    err := m.GetBackingStore().Set("eligibilityScheduleRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetEligibilitySchedules sets the eligibilitySchedules property value. The eligibilitySchedules property
func (m *PrivilegedAccessGroup) SetEligibilitySchedules(value []PrivilegedAccessGroupEligibilityScheduleable)() {
    err := m.GetBackingStore().Set("eligibilitySchedules", value)
    if err != nil {
        panic(err)
    }
}
// PrivilegedAccessGroupable 
type PrivilegedAccessGroupable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentApprovals()([]Approvalable)
    GetAssignmentScheduleInstances()([]PrivilegedAccessGroupAssignmentScheduleInstanceable)
    GetAssignmentScheduleRequests()([]PrivilegedAccessGroupAssignmentScheduleRequestable)
    GetAssignmentSchedules()([]PrivilegedAccessGroupAssignmentScheduleable)
    GetEligibilityScheduleInstances()([]PrivilegedAccessGroupEligibilityScheduleInstanceable)
    GetEligibilityScheduleRequests()([]PrivilegedAccessGroupEligibilityScheduleRequestable)
    GetEligibilitySchedules()([]PrivilegedAccessGroupEligibilityScheduleable)
    SetAssignmentApprovals(value []Approvalable)()
    SetAssignmentScheduleInstances(value []PrivilegedAccessGroupAssignmentScheduleInstanceable)()
    SetAssignmentScheduleRequests(value []PrivilegedAccessGroupAssignmentScheduleRequestable)()
    SetAssignmentSchedules(value []PrivilegedAccessGroupAssignmentScheduleable)()
    SetEligibilityScheduleInstances(value []PrivilegedAccessGroupEligibilityScheduleInstanceable)()
    SetEligibilityScheduleRequests(value []PrivilegedAccessGroupEligibilityScheduleRequestable)()
    SetEligibilitySchedules(value []PrivilegedAccessGroupEligibilityScheduleable)()
}
