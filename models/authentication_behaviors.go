package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AuthenticationBehaviors 
type AuthenticationBehaviors struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAuthenticationBehaviors instantiates a new authenticationBehaviors and sets the default values.
func NewAuthenticationBehaviors()(*AuthenticationBehaviors) {
    m := &AuthenticationBehaviors{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationBehaviorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationBehaviorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationBehaviors(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationBehaviors) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *AuthenticationBehaviors) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBlockAzureADGraphAccess gets the blockAzureADGraphAccess property value. The blockAzureADGraphAccess property
func (m *AuthenticationBehaviors) GetBlockAzureADGraphAccess()(*bool) {
    val, err := m.GetBackingStore().Get("blockAzureADGraphAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationBehaviors) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blockAzureADGraphAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockAzureADGraphAccess(val)
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
    res["removeUnverifiedEmailClaim"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveUnverifiedEmailClaim(val)
        }
        return nil
    }
    res["requireClientServicePrincipal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireClientServicePrincipal(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationBehaviors) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemoveUnverifiedEmailClaim gets the removeUnverifiedEmailClaim property value. Removes the email claim from tokens sent to an application when the email address's domain can't be verified.
func (m *AuthenticationBehaviors) GetRemoveUnverifiedEmailClaim()(*bool) {
    val, err := m.GetBackingStore().Get("removeUnverifiedEmailClaim")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRequireClientServicePrincipal gets the requireClientServicePrincipal property value. Requires multitenant applications to have a service principal in the resource tenant as part of authorization checks before they're granted access tokens. This property is only modifiable for multi-tenant resource applications that rely on access from clients without a service principal and had this behavior as set to false by Microsoft. Tenant administrators should respond to security advisories sent through Azure Health Service events and the Microsoft 365 message center.
func (m *AuthenticationBehaviors) GetRequireClientServicePrincipal()(*bool) {
    val, err := m.GetBackingStore().Get("requireClientServicePrincipal")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuthenticationBehaviors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("blockAzureADGraphAccess", m.GetBlockAzureADGraphAccess())
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
    {
        err := writer.WriteBoolValue("removeUnverifiedEmailClaim", m.GetRemoveUnverifiedEmailClaim())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("requireClientServicePrincipal", m.GetRequireClientServicePrincipal())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationBehaviors) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AuthenticationBehaviors) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBlockAzureADGraphAccess sets the blockAzureADGraphAccess property value. The blockAzureADGraphAccess property
func (m *AuthenticationBehaviors) SetBlockAzureADGraphAccess(value *bool)() {
    err := m.GetBackingStore().Set("blockAzureADGraphAccess", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationBehaviors) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoveUnverifiedEmailClaim sets the removeUnverifiedEmailClaim property value. Removes the email claim from tokens sent to an application when the email address's domain can't be verified.
func (m *AuthenticationBehaviors) SetRemoveUnverifiedEmailClaim(value *bool)() {
    err := m.GetBackingStore().Set("removeUnverifiedEmailClaim", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireClientServicePrincipal sets the requireClientServicePrincipal property value. Requires multitenant applications to have a service principal in the resource tenant as part of authorization checks before they're granted access tokens. This property is only modifiable for multi-tenant resource applications that rely on access from clients without a service principal and had this behavior as set to false by Microsoft. Tenant administrators should respond to security advisories sent through Azure Health Service events and the Microsoft 365 message center.
func (m *AuthenticationBehaviors) SetRequireClientServicePrincipal(value *bool)() {
    err := m.GetBackingStore().Set("requireClientServicePrincipal", value)
    if err != nil {
        panic(err)
    }
}
// AuthenticationBehaviorsable 
type AuthenticationBehaviorsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBlockAzureADGraphAccess()(*bool)
    GetOdataType()(*string)
    GetRemoveUnverifiedEmailClaim()(*bool)
    GetRequireClientServicePrincipal()(*bool)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBlockAzureADGraphAccess(value *bool)()
    SetOdataType(value *string)()
    SetRemoveUnverifiedEmailClaim(value *bool)()
    SetRequireClientServicePrincipal(value *bool)()
}
