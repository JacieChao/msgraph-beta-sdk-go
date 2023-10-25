package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleDefinition 
type UnifiedRoleDefinition struct {
    Entity
}
// NewUnifiedRoleDefinition instantiates a new unifiedRoleDefinition and sets the default values.
func NewUnifiedRoleDefinition()(*UnifiedRoleDefinition) {
    m := &UnifiedRoleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleDefinition(), nil
}
// GetAllowedPrincipalTypes gets the allowedPrincipalTypes property value. Types of principals that can be assigned the role. Read-only. The possible values are: user, servicePrincipal, group, unknownFutureValue. This is a multi-valued enumeration that can contain up to three values as a comma-separated string. For example, user, group. Supports $filter (eq).
func (m *UnifiedRoleDefinition) GetAllowedPrincipalTypes()(*AllowedRolePrincipalTypes) {
    val, err := m.GetBackingStore().Get("allowedPrincipalTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AllowedRolePrincipalTypes)
    }
    return nil
}
// GetDescription gets the description property value. The description for the unifiedRoleDefinition. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the unifiedRoleDefinition. Read-only when isBuiltIn is true. Required.  Supports $filter (eq and startsWith).
func (m *UnifiedRoleDefinition) GetDisplayName()(*string) {
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
func (m *UnifiedRoleDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedPrincipalTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAllowedRolePrincipalTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedPrincipalTypes(val.(*AllowedRolePrincipalTypes))
        }
        return nil
    }
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
    res["inheritsPermissionsFrom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRoleDefinitionable)
                }
            }
            m.SetInheritsPermissionsFrom(res)
        }
        return nil
    }
    res["isBuiltIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBuiltIn(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["isPrivileged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPrivileged(val)
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetResourceScopes(res)
        }
        return nil
    }
    res["rolePermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRolePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRolePermissionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnifiedRolePermissionable)
                }
            }
            m.SetRolePermissions(res)
        }
        return nil
    }
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetInheritsPermissionsFrom gets the inheritsPermissionsFrom property value. Read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles support this attribute.
func (m *UnifiedRoleDefinition) GetInheritsPermissionsFrom()([]UnifiedRoleDefinitionable) {
    val, err := m.GetBackingStore().Get("inheritsPermissionsFrom")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRoleDefinitionable)
    }
    return nil
}
// GetIsBuiltIn gets the isBuiltIn property value. Flag indicating if the unifiedRoleDefinition is part of the default set included with the product or custom. Read-only.  Supports $filter (eq).
func (m *UnifiedRoleDefinition) GetIsBuiltIn()(*bool) {
    val, err := m.GetBackingStore().Get("isBuiltIn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsEnabled gets the isEnabled property value. Flag indicating if the role is enabled for assignment. If false the role is not available for assignment. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) GetIsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsPrivileged gets the isPrivileged property value. Flag indicating if the role is privileged. Microsoft Entra ID defines a role as privileged if it contains at least one sensitive resource action in the rolePermissions and allowedResourceActions objects. Applies only for actions in the microsoft.directory resource namespace. Read-only. Supports $filter (eq).
func (m *UnifiedRoleDefinition) GetIsPrivileged()(*bool) {
    val, err := m.GetBackingStore().Get("isPrivileged")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetResourceScopes gets the resourceScopes property value. List of scopes permissions granted by the role definition apply to. Currently only / is supported. Read-only when isBuiltIn is true. DO NOT USE. This will be deprecated soon. Attach scope to role assignment.
func (m *UnifiedRoleDefinition) GetResourceScopes()([]string) {
    val, err := m.GetBackingStore().Get("resourceScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRolePermissions gets the rolePermissions property value. List of permissions included in the role. Read-only when isBuiltIn is true. Required.
func (m *UnifiedRoleDefinition) GetRolePermissions()([]UnifiedRolePermissionable) {
    val, err := m.GetBackingStore().Get("rolePermissions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnifiedRolePermissionable)
    }
    return nil
}
// GetTemplateId gets the templateId property value. Custom template identifier that can be set when isBuiltIn is false. This identifier is typically used if one needs an identifier to be the same across different directories. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) GetTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("templateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVersion gets the version property value. Indicates the version of the unifiedRoleDefinition object. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) GetVersion()(*string) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UnifiedRoleDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedPrincipalTypes() != nil {
        cast := (*m.GetAllowedPrincipalTypes()).String()
        err = writer.WriteStringValue("allowedPrincipalTypes", &cast)
        if err != nil {
            return err
        }
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
    if m.GetInheritsPermissionsFrom() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInheritsPermissionsFrom()))
        for i, v := range m.GetInheritsPermissionsFrom() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("inheritsPermissionsFrom", cast)
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
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPrivileged", m.GetIsPrivileged())
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
    if m.GetRolePermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRolePermissions()))
        for i, v := range m.GetRolePermissions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("rolePermissions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedPrincipalTypes sets the allowedPrincipalTypes property value. Types of principals that can be assigned the role. Read-only. The possible values are: user, servicePrincipal, group, unknownFutureValue. This is a multi-valued enumeration that can contain up to three values as a comma-separated string. For example, user, group. Supports $filter (eq).
func (m *UnifiedRoleDefinition) SetAllowedPrincipalTypes(value *AllowedRolePrincipalTypes)() {
    err := m.GetBackingStore().Set("allowedPrincipalTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description for the unifiedRoleDefinition. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the unifiedRoleDefinition. Read-only when isBuiltIn is true. Required.  Supports $filter (eq and startsWith).
func (m *UnifiedRoleDefinition) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetInheritsPermissionsFrom sets the inheritsPermissionsFrom property value. Read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles support this attribute.
func (m *UnifiedRoleDefinition) SetInheritsPermissionsFrom(value []UnifiedRoleDefinitionable)() {
    err := m.GetBackingStore().Set("inheritsPermissionsFrom", value)
    if err != nil {
        panic(err)
    }
}
// SetIsBuiltIn sets the isBuiltIn property value. Flag indicating if the unifiedRoleDefinition is part of the default set included with the product or custom. Read-only.  Supports $filter (eq).
func (m *UnifiedRoleDefinition) SetIsBuiltIn(value *bool)() {
    err := m.GetBackingStore().Set("isBuiltIn", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEnabled sets the isEnabled property value. Flag indicating if the role is enabled for assignment. If false the role is not available for assignment. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) SetIsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPrivileged sets the isPrivileged property value. Flag indicating if the role is privileged. Microsoft Entra ID defines a role as privileged if it contains at least one sensitive resource action in the rolePermissions and allowedResourceActions objects. Applies only for actions in the microsoft.directory resource namespace. Read-only. Supports $filter (eq).
func (m *UnifiedRoleDefinition) SetIsPrivileged(value *bool)() {
    err := m.GetBackingStore().Set("isPrivileged", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceScopes sets the resourceScopes property value. List of scopes permissions granted by the role definition apply to. Currently only / is supported. Read-only when isBuiltIn is true. DO NOT USE. This will be deprecated soon. Attach scope to role assignment.
func (m *UnifiedRoleDefinition) SetResourceScopes(value []string)() {
    err := m.GetBackingStore().Set("resourceScopes", value)
    if err != nil {
        panic(err)
    }
}
// SetRolePermissions sets the rolePermissions property value. List of permissions included in the role. Read-only when isBuiltIn is true. Required.
func (m *UnifiedRoleDefinition) SetRolePermissions(value []UnifiedRolePermissionable)() {
    err := m.GetBackingStore().Set("rolePermissions", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateId sets the templateId property value. Custom template identifier that can be set when isBuiltIn is false. This identifier is typically used if one needs an identifier to be the same across different directories. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) SetTemplateId(value *string)() {
    err := m.GetBackingStore().Set("templateId", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. Indicates the version of the unifiedRoleDefinition object. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) SetVersion(value *string)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
// UnifiedRoleDefinitionable 
type UnifiedRoleDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedPrincipalTypes()(*AllowedRolePrincipalTypes)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetInheritsPermissionsFrom()([]UnifiedRoleDefinitionable)
    GetIsBuiltIn()(*bool)
    GetIsEnabled()(*bool)
    GetIsPrivileged()(*bool)
    GetResourceScopes()([]string)
    GetRolePermissions()([]UnifiedRolePermissionable)
    GetTemplateId()(*string)
    GetVersion()(*string)
    SetAllowedPrincipalTypes(value *AllowedRolePrincipalTypes)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetInheritsPermissionsFrom(value []UnifiedRoleDefinitionable)()
    SetIsBuiltIn(value *bool)()
    SetIsEnabled(value *bool)()
    SetIsPrivileged(value *bool)()
    SetResourceScopes(value []string)()
    SetRolePermissions(value []UnifiedRolePermissionable)()
    SetTemplateId(value *string)()
    SetVersion(value *string)()
}
