package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GovernanceRoleAssignmentRequestStatus struct {
    additionalData map[string]interface{};
    status *string;
    statusDetails []KeyValue;
    subStatus *string;
}
func NewGovernanceRoleAssignmentRequestStatus()(*GovernanceRoleAssignmentRequestStatus) {
    m := &GovernanceRoleAssignmentRequestStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *GovernanceRoleAssignmentRequestStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *GovernanceRoleAssignmentRequestStatus) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *GovernanceRoleAssignmentRequestStatus) GetStatusDetails()([]KeyValue) {
    if m == nil {
        return nil
    } else {
        return m.statusDetails
    }
}
func (m *GovernanceRoleAssignmentRequestStatus) GetSubStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subStatus
    }
}
func (m *GovernanceRoleAssignmentRequestStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    res["statusDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValue() })
        if err != nil {
            return err
        }
        res := make([]KeyValue, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValue))
        }
        m.SetStatusDetails(res)
        return nil
    }
    res["subStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubStatus(val)
        return nil
    }
    return res
}
func (m *GovernanceRoleAssignmentRequestStatus) IsNil()(bool) {
    return m == nil
}
func (m *GovernanceRoleAssignmentRequestStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStatusDetails()))
        for i, v := range m.GetStatusDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("statusDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subStatus", m.GetSubStatus())
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
func (m *GovernanceRoleAssignmentRequestStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *GovernanceRoleAssignmentRequestStatus) SetStatus(value *string)() {
    m.status = value
}
func (m *GovernanceRoleAssignmentRequestStatus) SetStatusDetails(value []KeyValue)() {
    m.statusDetails = value
}
func (m *GovernanceRoleAssignmentRequestStatus) SetSubStatus(value *string)() {
    m.subStatus = value
}