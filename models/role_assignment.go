package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleAssignment 
type RoleAssignment struct {
    Entity
}
// NewRoleAssignment instantiates a new RoleAssignment and sets the default values.
func NewRoleAssignment()(*RoleAssignment) {
    m := &RoleAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRoleAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceAndAppManagementRoleAssignment":
                        return NewDeviceAndAppManagementRoleAssignment(), nil
                }
            }
        }
    }
    return NewRoleAssignment(), nil
}
// GetDescription gets the description property value. Description of the Role Assignment.
func (m *RoleAssignment) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display or friendly name of the role Assignment.
func (m *RoleAssignment) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["resourceScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetResourceScopes(res)
        }
        return nil
    }
    res["roleDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinition(val.(RoleDefinitionable))
        }
        return nil
    }
    res["scopeMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetScopeMembers(res)
        }
        return nil
    }
    res["scopeType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoleAssignmentScopeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScopeType(val.(*RoleAssignmentScopeType))
        }
        return nil
    }
    return res
}
// GetResourceScopes gets the resourceScopes property value. List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
func (m *RoleAssignment) GetResourceScopes()([]string) {
    val, err := m.GetBackingStore().Get("resourceScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRoleDefinition gets the roleDefinition property value. Role definition this assignment is part of.
func (m *RoleAssignment) GetRoleDefinition()(RoleDefinitionable) {
    val, err := m.GetBackingStore().Get("roleDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RoleDefinitionable)
    }
    return nil
}
// GetScopeMembers gets the scopeMembers property value. List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
func (m *RoleAssignment) GetScopeMembers()([]string) {
    val, err := m.GetBackingStore().Get("scopeMembers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetScopeType gets the scopeType property value. Specifies the type of scope for a Role Assignment.
func (m *RoleAssignment) GetScopeType()(*RoleAssignmentScopeType) {
    val, err := m.GetBackingStore().Get("scopeType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RoleAssignmentScopeType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RoleAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetResourceScopes() != nil {
        err = writer.WriteCollectionOfStringValues("resourceScopes", m.GetResourceScopes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleDefinition", m.GetRoleDefinition())
        if err != nil {
            return err
        }
    }
    if m.GetScopeMembers() != nil {
        err = writer.WriteCollectionOfStringValues("scopeMembers", m.GetScopeMembers())
        if err != nil {
            return err
        }
    }
    if m.GetScopeType() != nil {
        cast := (*m.GetScopeType()).String()
        err = writer.WriteStringValue("scopeType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description of the Role Assignment.
func (m *RoleAssignment) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display or friendly name of the role Assignment.
func (m *RoleAssignment) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceScopes sets the resourceScopes property value. List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
func (m *RoleAssignment) SetResourceScopes(value []string)() {
    err := m.GetBackingStore().Set("resourceScopes", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinition sets the roleDefinition property value. Role definition this assignment is part of.
func (m *RoleAssignment) SetRoleDefinition(value RoleDefinitionable)() {
    err := m.GetBackingStore().Set("roleDefinition", value)
    if err != nil {
        panic(err)
    }
}
// SetScopeMembers sets the scopeMembers property value. List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
func (m *RoleAssignment) SetScopeMembers(value []string)() {
    err := m.GetBackingStore().Set("scopeMembers", value)
    if err != nil {
        panic(err)
    }
}
// SetScopeType sets the scopeType property value. Specifies the type of scope for a Role Assignment.
func (m *RoleAssignment) SetScopeType(value *RoleAssignmentScopeType)() {
    err := m.GetBackingStore().Set("scopeType", value)
    if err != nil {
        panic(err)
    }
}
// RoleAssignmentable 
type RoleAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetResourceScopes()([]string)
    GetRoleDefinition()(RoleDefinitionable)
    GetScopeMembers()([]string)
    GetScopeType()(*RoleAssignmentScopeType)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetResourceScopes(value []string)()
    SetRoleDefinition(value RoleDefinitionable)()
    SetScopeMembers(value []string)()
    SetScopeType(value *RoleAssignmentScopeType)()
}
