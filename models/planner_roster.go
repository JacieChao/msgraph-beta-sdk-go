package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PlannerRoster struct {
    Entity
}
// NewPlannerRoster instantiates a new PlannerRoster and sets the default values.
func NewPlannerRoster()(*PlannerRoster) {
    m := &PlannerRoster{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerRosterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePlannerRosterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerRoster(), nil
}
// GetAssignedSensitivityLabel gets the assignedSensitivityLabel property value. The assignedSensitivityLabel property
// returns a SensitivityLabelAssignmentable when successful
func (m *PlannerRoster) GetAssignedSensitivityLabel()(SensitivityLabelAssignmentable) {
    val, err := m.GetBackingStore().Get("assignedSensitivityLabel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SensitivityLabelAssignmentable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PlannerRoster) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedSensitivityLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSensitivityLabelAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedSensitivityLabel(val.(SensitivityLabelAssignmentable))
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerRosterMemberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerRosterMemberable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PlannerRosterMemberable)
                }
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["plans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlanable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PlannerPlanable)
                }
            }
            m.SetPlans(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. Retrieves the members of the plannerRoster.
// returns a []PlannerRosterMemberable when successful
func (m *PlannerRoster) GetMembers()([]PlannerRosterMemberable) {
    val, err := m.GetBackingStore().Get("members")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerRosterMemberable)
    }
    return nil
}
// GetPlans gets the plans property value. Retrieves the plans contained by the plannerRoster.
// returns a []PlannerPlanable when successful
func (m *PlannerRoster) GetPlans()([]PlannerPlanable) {
    val, err := m.GetBackingStore().Get("plans")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PlannerPlanable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PlannerRoster) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("assignedSensitivityLabel", m.GetAssignedSensitivityLabel())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPlans()))
        for i, v := range m.GetPlans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("plans", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedSensitivityLabel sets the assignedSensitivityLabel property value. The assignedSensitivityLabel property
func (m *PlannerRoster) SetAssignedSensitivityLabel(value SensitivityLabelAssignmentable)() {
    err := m.GetBackingStore().Set("assignedSensitivityLabel", value)
    if err != nil {
        panic(err)
    }
}
// SetMembers sets the members property value. Retrieves the members of the plannerRoster.
func (m *PlannerRoster) SetMembers(value []PlannerRosterMemberable)() {
    err := m.GetBackingStore().Set("members", value)
    if err != nil {
        panic(err)
    }
}
// SetPlans sets the plans property value. Retrieves the plans contained by the plannerRoster.
func (m *PlannerRoster) SetPlans(value []PlannerPlanable)() {
    err := m.GetBackingStore().Set("plans", value)
    if err != nil {
        panic(err)
    }
}
type PlannerRosterable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedSensitivityLabel()(SensitivityLabelAssignmentable)
    GetMembers()([]PlannerRosterMemberable)
    GetPlans()([]PlannerPlanable)
    SetAssignedSensitivityLabel(value SensitivityLabelAssignmentable)()
    SetMembers(value []PlannerRosterMemberable)()
    SetPlans(value []PlannerPlanable)()
}
