package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationListener 
type AuthenticationListener struct {
    Entity
}
// NewAuthenticationListener instantiates a new authenticationListener and sets the default values.
func NewAuthenticationListener()(*AuthenticationListener) {
    m := &AuthenticationListener{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationListenerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationListenerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.invokeUserFlowListener":
                        return NewInvokeUserFlowListener(), nil
                }
            }
        }
    }
    return NewAuthenticationListener(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationListener) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["sourceFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationSourceFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceFilter(val.(AuthenticationSourceFilterable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationListener) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPriority gets the priority property value. The priority of the listener. Determines the order of evaluation when an event has multiple listeners. The priority is evaluated from low to high.
func (m *AuthenticationListener) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSourceFilter gets the sourceFilter property value. Filter based on the source of the authentication that is used to determine whether the listener is evaluated. This is currently limited to evaluations based on application the user is authenticating to.
func (m *AuthenticationListener) GetSourceFilter()(AuthenticationSourceFilterable) {
    val, err := m.GetBackingStore().Get("sourceFilter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthenticationSourceFilterable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AuthenticationListener) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceFilter", m.GetSourceFilter())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationListener) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. The priority of the listener. Determines the order of evaluation when an event has multiple listeners. The priority is evaluated from low to high.
func (m *AuthenticationListener) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceFilter sets the sourceFilter property value. Filter based on the source of the authentication that is used to determine whether the listener is evaluated. This is currently limited to evaluations based on application the user is authenticating to.
func (m *AuthenticationListener) SetSourceFilter(value AuthenticationSourceFilterable)() {
    err := m.GetBackingStore().Set("sourceFilter", value)
    if err != nil {
        panic(err)
    }
}
// AuthenticationListenerable 
type AuthenticationListenerable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPriority()(*int32)
    GetSourceFilter()(AuthenticationSourceFilterable)
    SetOdataType(value *string)()
    SetPriority(value *int32)()
    SetSourceFilter(value AuthenticationSourceFilterable)()
}
