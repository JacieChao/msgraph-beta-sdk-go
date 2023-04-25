package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeClientConfiguration 
type OfficeClientConfiguration struct {
    Entity
}
// NewOfficeClientConfiguration instantiates a new officeClientConfiguration and sets the default values.
func NewOfficeClientConfiguration()(*OfficeClientConfiguration) {
    m := &OfficeClientConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOfficeClientConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeClientConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsOfficeClientConfiguration":
                        return NewWindowsOfficeClientConfiguration(), nil
                    case "#microsoft.graph.windowsOfficeClientSecurityConfiguration":
                        return NewWindowsOfficeClientSecurityConfiguration(), nil
                }
            }
        }
    }
    return NewOfficeClientConfiguration(), nil
}
// GetAssignments gets the assignments property value. The assignments property
func (m *OfficeClientConfiguration) GetAssignments()([]OfficeClientConfigurationAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OfficeClientConfigurationAssignmentable)
    }
    return nil
}
// GetCheckinStatuses gets the checkinStatuses property value. The checkinStatuses property
func (m *OfficeClientConfiguration) GetCheckinStatuses()([]OfficeClientCheckinStatusable) {
    val, err := m.GetBackingStore().Get("checkinStatuses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OfficeClientCheckinStatusable)
    }
    return nil
}
// GetDescription gets the description property value. The description property
func (m *OfficeClientConfiguration) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *OfficeClientConfiguration) GetDisplayName()(*string) {
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
func (m *OfficeClientConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOfficeClientConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeClientConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(OfficeClientConfigurationAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["checkinStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOfficeClientCheckinStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeClientCheckinStatusable, len(val))
            for i, v := range val {
                res[i] = v.(OfficeClientCheckinStatusable)
            }
            m.SetCheckinStatuses(res)
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
    res["policyPayload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyPayload(val)
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
    res["userCheckinSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOfficeUserCheckinSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserCheckinSummary(val.(OfficeUserCheckinSummaryable))
        }
        return nil
    }
    res["userPreferencePayload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPreferencePayload(val)
        }
        return nil
    }
    return res
}
// GetPolicyPayload gets the policyPayload property value. The policyPayload property
func (m *OfficeClientConfiguration) GetPolicyPayload()([]byte) {
    val, err := m.GetBackingStore().Get("policyPayload")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetPriority gets the priority property value. The priority property
func (m *OfficeClientConfiguration) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUserCheckinSummary gets the userCheckinSummary property value. The userCheckinSummary property
func (m *OfficeClientConfiguration) GetUserCheckinSummary()(OfficeUserCheckinSummaryable) {
    val, err := m.GetBackingStore().Get("userCheckinSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OfficeUserCheckinSummaryable)
    }
    return nil
}
// GetUserPreferencePayload gets the userPreferencePayload property value. The userPreferencePayload property
func (m *OfficeClientConfiguration) GetUserPreferencePayload()([]byte) {
    val, err := m.GetBackingStore().Get("userPreferencePayload")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OfficeClientConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCheckinStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCheckinStatuses()))
        for i, v := range m.GetCheckinStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("checkinStatuses", cast)
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
    {
        err = writer.WriteByteArrayValue("policyPayload", m.GetPolicyPayload())
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
        err = writer.WriteObjectValue("userCheckinSummary", m.GetUserCheckinSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("userPreferencePayload", m.GetUserPreferencePayload())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The assignments property
func (m *OfficeClientConfiguration) SetAssignments(value []OfficeClientConfigurationAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetCheckinStatuses sets the checkinStatuses property value. The checkinStatuses property
func (m *OfficeClientConfiguration) SetCheckinStatuses(value []OfficeClientCheckinStatusable)() {
    err := m.GetBackingStore().Set("checkinStatuses", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *OfficeClientConfiguration) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *OfficeClientConfiguration) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyPayload sets the policyPayload property value. The policyPayload property
func (m *OfficeClientConfiguration) SetPolicyPayload(value []byte)() {
    err := m.GetBackingStore().Set("policyPayload", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. The priority property
func (m *OfficeClientConfiguration) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetUserCheckinSummary sets the userCheckinSummary property value. The userCheckinSummary property
func (m *OfficeClientConfiguration) SetUserCheckinSummary(value OfficeUserCheckinSummaryable)() {
    err := m.GetBackingStore().Set("userCheckinSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPreferencePayload sets the userPreferencePayload property value. The userPreferencePayload property
func (m *OfficeClientConfiguration) SetUserPreferencePayload(value []byte)() {
    err := m.GetBackingStore().Set("userPreferencePayload", value)
    if err != nil {
        panic(err)
    }
}
// OfficeClientConfigurationable 
type OfficeClientConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]OfficeClientConfigurationAssignmentable)
    GetCheckinStatuses()([]OfficeClientCheckinStatusable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetPolicyPayload()([]byte)
    GetPriority()(*int32)
    GetUserCheckinSummary()(OfficeUserCheckinSummaryable)
    GetUserPreferencePayload()([]byte)
    SetAssignments(value []OfficeClientConfigurationAssignmentable)()
    SetCheckinStatuses(value []OfficeClientCheckinStatusable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetPolicyPayload(value []byte)()
    SetPriority(value *int32)()
    SetUserCheckinSummary(value OfficeUserCheckinSummaryable)()
    SetUserPreferencePayload(value []byte)()
}
