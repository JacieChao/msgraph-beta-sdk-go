package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Fqdn 
type Fqdn struct {
    RuleDestination
    // The OdataType property
    OdataType *string
}
// NewFqdn instantiates a new fqdn and sets the default values.
func NewFqdn()(*Fqdn) {
    m := &Fqdn{
        RuleDestination: *NewRuleDestination(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.fqdn"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFqdnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFqdnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFqdn(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Fqdn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RuleDestination.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. Defines the FQDN used in a destination for a rule.
func (m *Fqdn) GetValue()(*string) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Fqdn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RuleDestination.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. Defines the FQDN used in a destination for a rule.
func (m *Fqdn) SetValue(value *string)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// Fqdnable 
type Fqdnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RuleDestinationable
    GetValue()(*string)
    SetValue(value *string)()
}
