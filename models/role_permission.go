package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// RolePermission contains the set of ResourceActions determining the allowed and not allowed permissions for each role.
type RolePermission struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewRolePermission instantiates a new rolePermission and sets the default values.
func NewRolePermission()(*RolePermission) {
    m := &RolePermission{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRolePermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRolePermissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRolePermission(), nil
}
// GetActions gets the actions property value. Allowed Actions - Deprecated
func (m *RolePermission) GetActions()([]string) {
    val, err := m.GetBackingStore().Get("actions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RolePermission) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *RolePermission) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RolePermission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetActions(res)
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
    res["resourceActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceActionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ResourceActionable)
                }
            }
            m.SetResourceActions(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RolePermission) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceActions gets the resourceActions property value. Resource Actions each containing a set of allowed and not allowed permissions.
func (m *RolePermission) GetResourceActions()([]ResourceActionable) {
    val, err := m.GetBackingStore().Get("resourceActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ResourceActionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RolePermission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActions() != nil {
        err := writer.WriteCollectionOfStringValues("actions", m.GetActions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetResourceActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceActions()))
        for i, v := range m.GetResourceActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("resourceActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. Allowed Actions - Deprecated
func (m *RolePermission) SetActions(value []string)() {
    err := m.GetBackingStore().Set("actions", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RolePermission) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *RolePermission) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RolePermission) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceActions sets the resourceActions property value. Resource Actions each containing a set of allowed and not allowed permissions.
func (m *RolePermission) SetResourceActions(value []ResourceActionable)() {
    err := m.GetBackingStore().Set("resourceActions", value)
    if err != nil {
        panic(err)
    }
}
// RolePermissionable 
type RolePermissionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetResourceActions()([]ResourceActionable)
    SetActions(value []string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetResourceActions(value []ResourceActionable)()
}
