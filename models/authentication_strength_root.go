package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationStrengthRoot 
type AuthenticationStrengthRoot struct {
    Entity
    // The authenticationCombinations property
    authenticationCombinations []string
    // The authenticationMethodModes property
    authenticationMethodModes []AuthenticationMethodModeDetailable
    // The policies property
    policies []AuthenticationStrengthPolicyable
}
// NewAuthenticationStrengthRoot instantiates a new AuthenticationStrengthRoot and sets the default values.
func NewAuthenticationStrengthRoot()(*AuthenticationStrengthRoot) {
    m := &AuthenticationStrengthRoot{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.authenticationStrengthRoot";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationStrengthRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationStrengthRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationStrengthRoot(), nil
}
// GetAuthenticationCombinations gets the authenticationCombinations property value. The authenticationCombinations property
func (m *AuthenticationStrengthRoot) GetAuthenticationCombinations()([]string) {
    return m.authenticationCombinations
}
// GetAuthenticationMethodModes gets the authenticationMethodModes property value. The authenticationMethodModes property
func (m *AuthenticationStrengthRoot) GetAuthenticationMethodModes()([]AuthenticationMethodModeDetailable) {
    return m.authenticationMethodModes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationStrengthRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationCombinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAuthenticationCombinations(res)
        }
        return nil
    }
    res["authenticationMethodModes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationMethodModeDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodModeDetailable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationMethodModeDetailable)
            }
            m.SetAuthenticationMethodModes(res)
        }
        return nil
    }
    res["policies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationStrengthPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationStrengthPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationStrengthPolicyable)
            }
            m.SetPolicies(res)
        }
        return nil
    }
    return res
}
// GetPolicies gets the policies property value. The policies property
func (m *AuthenticationStrengthRoot) GetPolicies()([]AuthenticationStrengthPolicyable) {
    return m.policies
}
// Serialize serializes information the current object
func (m *AuthenticationStrengthRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthenticationCombinations() != nil {
        err = writer.WriteCollectionOfStringValues("authenticationCombinations", m.GetAuthenticationCombinations())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationMethodModes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthenticationMethodModes()))
        for i, v := range m.GetAuthenticationMethodModes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("authenticationMethodModes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPolicies()))
        for i, v := range m.GetPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("policies", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationCombinations sets the authenticationCombinations property value. The authenticationCombinations property
func (m *AuthenticationStrengthRoot) SetAuthenticationCombinations(value []string)() {
    m.authenticationCombinations = value
}
// SetAuthenticationMethodModes sets the authenticationMethodModes property value. The authenticationMethodModes property
func (m *AuthenticationStrengthRoot) SetAuthenticationMethodModes(value []AuthenticationMethodModeDetailable)() {
    m.authenticationMethodModes = value
}
// SetPolicies sets the policies property value. The policies property
func (m *AuthenticationStrengthRoot) SetPolicies(value []AuthenticationStrengthPolicyable)() {
    m.policies = value
}