package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ForwardingProfile 
type ForwardingProfile struct {
    Profile
}
// NewForwardingProfile instantiates a new forwardingProfile and sets the default values.
func NewForwardingProfile()(*ForwardingProfile) {
    m := &ForwardingProfile{
        Profile: *NewProfile(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.forwardingProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateForwardingProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateForwardingProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewForwardingProfile(), nil
}
// GetAssociations gets the associations property value. Specifies the users, groups, devices, and branch locations whose traffic is associated with the given traffic forwarding profile.
func (m *ForwardingProfile) GetAssociations()([]Associationable) {
    val, err := m.GetBackingStore().Get("associations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Associationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ForwardingProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Profile.GetFieldDeserializers()
    res["associations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssociationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Associationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Associationable)
                }
            }
            m.SetAssociations(res)
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
    res["trafficForwardingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrafficForwardingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrafficForwardingType(val.(*TrafficForwardingType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ForwardingProfile) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPriority gets the priority property value. Profile priority.
func (m *ForwardingProfile) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTrafficForwardingType gets the trafficForwardingType property value. The trafficForwardingType property
func (m *ForwardingProfile) GetTrafficForwardingType()(*TrafficForwardingType) {
    val, err := m.GetBackingStore().Get("trafficForwardingType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TrafficForwardingType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ForwardingProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Profile.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssociations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssociations()))
        for i, v := range m.GetAssociations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("associations", cast)
        if err != nil {
            return err
        }
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
    if m.GetTrafficForwardingType() != nil {
        cast := (*m.GetTrafficForwardingType()).String()
        err = writer.WriteStringValue("trafficForwardingType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssociations sets the associations property value. Specifies the users, groups, devices, and branch locations whose traffic is associated with the given traffic forwarding profile.
func (m *ForwardingProfile) SetAssociations(value []Associationable)() {
    err := m.GetBackingStore().Set("associations", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ForwardingProfile) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. Profile priority.
func (m *ForwardingProfile) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetTrafficForwardingType sets the trafficForwardingType property value. The trafficForwardingType property
func (m *ForwardingProfile) SetTrafficForwardingType(value *TrafficForwardingType)() {
    err := m.GetBackingStore().Set("trafficForwardingType", value)
    if err != nil {
        panic(err)
    }
}
// ForwardingProfileable 
type ForwardingProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Profileable
    GetAssociations()([]Associationable)
    GetOdataType()(*string)
    GetPriority()(*int32)
    GetTrafficForwardingType()(*TrafficForwardingType)
    SetAssociations(value []Associationable)()
    SetOdataType(value *string)()
    SetPriority(value *int32)()
    SetTrafficForwardingType(value *TrafficForwardingType)()
}
