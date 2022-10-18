package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleDefinition 
type RoleDefinition struct {
    Entity
    // Description of the Role definition.
    description *string
    // Display Name of the Role definition.
    displayName *string
    // Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
    isBuiltIn *bool
    // Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
    isBuiltInRoleDefinition *bool
    // List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
    permissions []RolePermissionable
    // List of Role assignments for this role definition.
    roleAssignments []RoleAssignmentable
    // List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
    rolePermissions []RolePermissionable
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
}
// NewRoleDefinition instantiates a new roleDefinition and sets the default values.
func NewRoleDefinition()(*RoleDefinition) {
    m := &RoleDefinition{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.roleDefinition";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRoleDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceAndAppManagementRoleDefinition":
                        return NewDeviceAndAppManagementRoleDefinition(), nil
                }
            }
        }
    }
    return NewRoleDefinition(), nil
}
// GetDescription gets the description property value. Description of the Role definition.
func (m *RoleDefinition) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display Name of the Role definition.
func (m *RoleDefinition) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["isBuiltIn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsBuiltIn)
    res["isBuiltInRoleDefinition"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsBuiltInRoleDefinition)
    res["permissions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRolePermissionFromDiscriminatorValue , m.SetPermissions)
    res["roleAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRoleAssignmentFromDiscriminatorValue , m.SetRoleAssignments)
    res["rolePermissions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRolePermissionFromDiscriminatorValue , m.SetRolePermissions)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    return res
}
// GetIsBuiltIn gets the isBuiltIn property value. Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
func (m *RoleDefinition) GetIsBuiltIn()(*bool) {
    return m.isBuiltIn
}
// GetIsBuiltInRoleDefinition gets the isBuiltInRoleDefinition property value. Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
func (m *RoleDefinition) GetIsBuiltInRoleDefinition()(*bool) {
    return m.isBuiltInRoleDefinition
}
// GetPermissions gets the permissions property value. List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
func (m *RoleDefinition) GetPermissions()([]RolePermissionable) {
    return m.permissions
}
// GetRoleAssignments gets the roleAssignments property value. List of Role assignments for this role definition.
func (m *RoleDefinition) GetRoleAssignments()([]RoleAssignmentable) {
    return m.roleAssignments
}
// GetRolePermissions gets the rolePermissions property value. List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
func (m *RoleDefinition) GetRolePermissions()([]RolePermissionable) {
    return m.rolePermissions
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *RoleDefinition) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// Serialize serializes information the current object
func (m *RoleDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("isBuiltIn", m.GetIsBuiltIn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBuiltInRoleDefinition", m.GetIsBuiltInRoleDefinition())
        if err != nil {
            return err
        }
    }
    if m.GetPermissions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPermissions())
        err = writer.WriteCollectionOfObjectValues("permissions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleAssignments())
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRolePermissions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRolePermissions())
        err = writer.WriteCollectionOfObjectValues("rolePermissions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description of the Role definition.
func (m *RoleDefinition) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display Name of the Role definition.
func (m *RoleDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsBuiltIn sets the isBuiltIn property value. Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
func (m *RoleDefinition) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
// SetIsBuiltInRoleDefinition sets the isBuiltInRoleDefinition property value. Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
func (m *RoleDefinition) SetIsBuiltInRoleDefinition(value *bool)() {
    m.isBuiltInRoleDefinition = value
}
// SetPermissions sets the permissions property value. List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
func (m *RoleDefinition) SetPermissions(value []RolePermissionable)() {
    m.permissions = value
}
// SetRoleAssignments sets the roleAssignments property value. List of Role assignments for this role definition.
func (m *RoleDefinition) SetRoleAssignments(value []RoleAssignmentable)() {
    m.roleAssignments = value
}
// SetRolePermissions sets the rolePermissions property value. List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
func (m *RoleDefinition) SetRolePermissions(value []RolePermissionable)() {
    m.rolePermissions = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *RoleDefinition) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
