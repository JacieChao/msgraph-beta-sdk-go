package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrivilegedRole struct {
    Entity
    assignments []PrivilegedRoleAssignment;
    name *string;
    settings *PrivilegedRoleSettings;
    summary *PrivilegedRoleSummary;
}
func NewPrivilegedRole()(*PrivilegedRole) {
    m := &PrivilegedRole{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PrivilegedRole) GetAssignments()([]PrivilegedRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
func (m *PrivilegedRole) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *PrivilegedRole) GetSettings()(*PrivilegedRoleSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *PrivilegedRole) GetSummary()(*PrivilegedRoleSummary) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
func (m *PrivilegedRole) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRoleAssignment() })
        if err != nil {
            return err
        }
        res := make([]PrivilegedRoleAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrivilegedRoleAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRoleSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*PrivilegedRoleSettings))
        return nil
    }
    res["summary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRoleSummary() })
        if err != nil {
            return err
        }
        m.SetSummary(val.(*PrivilegedRoleSummary))
        return nil
    }
    return res
}
func (m *PrivilegedRole) IsNil()(bool) {
    return m == nil
}
func (m *PrivilegedRole) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("summary", m.GetSummary())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PrivilegedRole) SetAssignments(value []PrivilegedRoleAssignment)() {
    m.assignments = value
}
func (m *PrivilegedRole) SetName(value *string)() {
    m.name = value
}
func (m *PrivilegedRole) SetSettings(value *PrivilegedRoleSettings)() {
    m.settings = value
}
func (m *PrivilegedRole) SetSummary(value *PrivilegedRoleSummary)() {
    m.summary = value
}