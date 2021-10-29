package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DefaultUserRoleOverride struct {
    Entity
    // 
    isDefault *bool;
    // 
    rolePermissions []UnifiedRolePermission;
}
// Instantiates a new defaultUserRoleOverride and sets the default values.
func NewDefaultUserRoleOverride()(*DefaultUserRoleOverride) {
    m := &DefaultUserRoleOverride{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the isDefault property value. 
func (m *DefaultUserRoleOverride) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// Gets the rolePermissions property value. 
func (m *DefaultUserRoleOverride) GetRolePermissions()([]UnifiedRolePermission) {
    if m == nil {
        return nil
    } else {
        return m.rolePermissions
    }
}
// The deserialization information for the current model
func (m *DefaultUserRoleOverride) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["rolePermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRolePermission() })
        if err != nil {
            return err
        }
        res := make([]UnifiedRolePermission, len(val))
        for i, v := range val {
            res[i] = *(v.(*UnifiedRolePermission))
        }
        m.SetRolePermissions(res)
        return nil
    }
    return res
}
func (m *DefaultUserRoleOverride) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DefaultUserRoleOverride) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRolePermissions()))
        for i, v := range m.GetRolePermissions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("rolePermissions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the isDefault property value. 
// Parameters:
//  - value : Value to set for the isDefault property.
func (m *DefaultUserRoleOverride) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// Sets the rolePermissions property value. 
// Parameters:
//  - value : Value to set for the rolePermissions property.
func (m *DefaultUserRoleOverride) SetRolePermissions(value []UnifiedRolePermission)() {
    m.rolePermissions = value
}