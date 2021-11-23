package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// roleAssignment 
type RoleAssignment struct {
    Entity
    // Description of the Role Assignment.
    description *string;
    // The display or friendly name of the role Assignment.
    displayName *string;
    // List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
    resourceScopes []string;
    // Role definition this assignment is part of.
    roleDefinition *RoleDefinition;
    // List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
    scopeMembers []string;
    // Specifies the type of scope for a Role Assignment. Default type 'ResourceScope' allows assignment of ResourceScopes. For 'AllDevices', 'AllLicensedUsers', and 'AllDevicesAndLicensedUsers', the ResourceScopes property should be left empty. Possible values are: resourceScope, allDevices, allLicensedUsers, allDevicesAndLicensedUsers.
    scopeType *RoleAssignmentScopeType;
}
// NewRoleAssignment instantiates a new roleAssignment and sets the default values.
func NewRoleAssignment()(*RoleAssignment) {
    m := &RoleAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetDescription gets the description property value. Description of the Role Assignment.
func (m *RoleAssignment) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display or friendly name of the role Assignment.
func (m *RoleAssignment) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetResourceScopes gets the resourceScopes property value. List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
func (m *RoleAssignment) GetResourceScopes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.resourceScopes
    }
}
// GetRoleDefinition gets the roleDefinition property value. Role definition this assignment is part of.
func (m *RoleAssignment) GetRoleDefinition()(*RoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinition
    }
}
// GetScopeMembers gets the scopeMembers property value. List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
func (m *RoleAssignment) GetScopeMembers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.scopeMembers
    }
}
// GetScopeType gets the scopeType property value. Specifies the type of scope for a Role Assignment. Default type 'ResourceScope' allows assignment of ResourceScopes. For 'AllDevices', 'AllLicensedUsers', and 'AllDevicesAndLicensedUsers', the ResourceScopes property should be left empty. Possible values are: resourceScope, allDevices, allLicensedUsers, allDevicesAndLicensedUsers.
func (m *RoleAssignment) GetScopeType()(*RoleAssignmentScopeType) {
    if m == nil {
        return nil
    } else {
        return m.scopeType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["resourceScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["roleDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinition(val.(*RoleDefinition))
        }
        return nil
    }
    res["scopeMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["scopeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoleAssignmentScopeType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RoleAssignmentScopeType)
            m.SetScopeType(&cast)
        }
        return nil
    }
    return res
}
func (m *RoleAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RoleAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
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
    {
        err = writer.WriteCollectionOfStringValues("scopeMembers", m.GetScopeMembers())
        if err != nil {
            return err
        }
    }
    if m.GetScopeType() != nil {
        cast := m.GetScopeType().String()
        err = writer.WriteStringValue("scopeType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description of the Role Assignment.
func (m *RoleAssignment) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display or friendly name of the role Assignment.
func (m *RoleAssignment) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetResourceScopes sets the resourceScopes property value. List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
func (m *RoleAssignment) SetResourceScopes(value []string)() {
    m.resourceScopes = value
}
// SetRoleDefinition sets the roleDefinition property value. Role definition this assignment is part of.
func (m *RoleAssignment) SetRoleDefinition(value *RoleDefinition)() {
    m.roleDefinition = value
}
// SetScopeMembers sets the scopeMembers property value. List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
func (m *RoleAssignment) SetScopeMembers(value []string)() {
    m.scopeMembers = value
}
// SetScopeType sets the scopeType property value. Specifies the type of scope for a Role Assignment. Default type 'ResourceScope' allows assignment of ResourceScopes. For 'AllDevices', 'AllLicensedUsers', and 'AllDevicesAndLicensedUsers', the ResourceScopes property should be left empty. Possible values are: resourceScope, allDevices, allLicensedUsers, allDevicesAndLicensedUsers.
func (m *RoleAssignment) SetScopeType(value *RoleAssignmentScopeType)() {
    m.scopeType = value
}
