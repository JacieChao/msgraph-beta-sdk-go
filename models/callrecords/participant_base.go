package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type ParticipantBase struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewParticipantBase instantiates a new ParticipantBase and sets the default values.
func NewParticipantBase()(*ParticipantBase) {
    m := &ParticipantBase{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateParticipantBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateParticipantBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.callRecords.organizer":
                        return NewOrganizer(), nil
                    case "#microsoft.graph.callRecords.participant":
                        return NewParticipant(), nil
                }
            }
        }
    }
    return NewParticipantBase(), nil
}
// GetAdministrativeUnitInfos gets the administrativeUnitInfos property value. List of administrativeUnitInfo of the call participant.
// returns a []AdministrativeUnitInfoable when successful
func (m *ParticipantBase) GetAdministrativeUnitInfos()([]AdministrativeUnitInfoable) {
    val, err := m.GetBackingStore().Get("administrativeUnitInfos")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AdministrativeUnitInfoable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ParticipantBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["administrativeUnitInfos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAdministrativeUnitInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AdministrativeUnitInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AdministrativeUnitInfoable)
                }
            }
            m.SetAdministrativeUnitInfos(res)
        }
        return nil
    }
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCommunicationsIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CommunicationsIdentitySetable))
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. The identity of the call participant.
// returns a CommunicationsIdentitySetable when successful
func (m *ParticipantBase) GetIdentity()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CommunicationsIdentitySetable) {
    val, err := m.GetBackingStore().Get("identity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CommunicationsIdentitySetable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ParticipantBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdministrativeUnitInfos() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdministrativeUnitInfos()))
        for i, v := range m.GetAdministrativeUnitInfos() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("administrativeUnitInfos", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdministrativeUnitInfos sets the administrativeUnitInfos property value. List of administrativeUnitInfo of the call participant.
func (m *ParticipantBase) SetAdministrativeUnitInfos(value []AdministrativeUnitInfoable)() {
    err := m.GetBackingStore().Set("administrativeUnitInfos", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentity sets the identity property value. The identity of the call participant.
func (m *ParticipantBase) SetIdentity(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CommunicationsIdentitySetable)() {
    err := m.GetBackingStore().Set("identity", value)
    if err != nil {
        panic(err)
    }
}
type ParticipantBaseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdministrativeUnitInfos()([]AdministrativeUnitInfoable)
    GetIdentity()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CommunicationsIdentitySetable)
    SetAdministrativeUnitInfos(value []AdministrativeUnitInfoable)()
    SetIdentity(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CommunicationsIdentitySetable)()
}
