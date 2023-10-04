package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AadSource 
type AadSource struct {
    AuthorizationSystemIdentitySource
}
// NewAadSource instantiates a new aadSource and sets the default values.
func NewAadSource()(*AadSource) {
    m := &AadSource{
        AuthorizationSystemIdentitySource: *NewAuthorizationSystemIdentitySource(),
    }
    odataTypeValue := "#microsoft.graph.aadSource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAadSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAadSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAadSource(), nil
}
// GetDomain gets the domain property value. The domain property
func (m *AadSource) GetDomain()(*string) {
    val, err := m.GetBackingStore().Get("domain")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AadSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizationSystemIdentitySource.GetFieldDeserializers()
    res["domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AadSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthorizationSystemIdentitySource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("domain", m.GetDomain())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDomain sets the domain property value. The domain property
func (m *AadSource) SetDomain(value *string)() {
    err := m.GetBackingStore().Set("domain", value)
    if err != nil {
        panic(err)
    }
}
// AadSourceable 
type AadSourceable interface {
    AuthorizationSystemIdentitySourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDomain()(*string)
    SetDomain(value *string)()
}
