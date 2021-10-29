package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InformationProtection struct {
    Entity
    bitlocker *Bitlocker;
    dataLossPreventionPolicies []DataLossPreventionPolicy;
    policy *InformationProtectionPolicy;
    sensitivityLabels []SensitivityLabel;
    sensitivityPolicySettings *SensitivityPolicySettings;
    threatAssessmentRequests []ThreatAssessmentRequest;
}
func NewInformationProtection()(*InformationProtection) {
    m := &InformationProtection{
        Entity: *NewEntity(),
    }
    return m
}
func (m *InformationProtection) GetBitlocker()(*Bitlocker) {
    if m == nil {
        return nil
    } else {
        return m.bitlocker
    }
}
func (m *InformationProtection) GetDataLossPreventionPolicies()([]DataLossPreventionPolicy) {
    if m == nil {
        return nil
    } else {
        return m.dataLossPreventionPolicies
    }
}
func (m *InformationProtection) GetPolicy()(*InformationProtectionPolicy) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
func (m *InformationProtection) GetSensitivityLabels()([]SensitivityLabel) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityLabels
    }
}
func (m *InformationProtection) GetSensitivityPolicySettings()(*SensitivityPolicySettings) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityPolicySettings
    }
}
func (m *InformationProtection) GetThreatAssessmentRequests()([]ThreatAssessmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.threatAssessmentRequests
    }
}
func (m *InformationProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bitlocker"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBitlocker() })
        if err != nil {
            return err
        }
        m.SetBitlocker(val.(*Bitlocker))
        return nil
    }
    res["dataLossPreventionPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDataLossPreventionPolicy() })
        if err != nil {
            return err
        }
        res := make([]DataLossPreventionPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*DataLossPreventionPolicy))
        }
        m.SetDataLossPreventionPolicies(res)
        return nil
    }
    res["policy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInformationProtectionPolicy() })
        if err != nil {
            return err
        }
        m.SetPolicy(val.(*InformationProtectionPolicy))
        return nil
    }
    res["sensitivityLabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitivityLabel() })
        if err != nil {
            return err
        }
        res := make([]SensitivityLabel, len(val))
        for i, v := range val {
            res[i] = *(v.(*SensitivityLabel))
        }
        m.SetSensitivityLabels(res)
        return nil
    }
    res["sensitivityPolicySettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitivityPolicySettings() })
        if err != nil {
            return err
        }
        m.SetSensitivityPolicySettings(val.(*SensitivityPolicySettings))
        return nil
    }
    res["threatAssessmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThreatAssessmentRequest() })
        if err != nil {
            return err
        }
        res := make([]ThreatAssessmentRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*ThreatAssessmentRequest))
        }
        m.SetThreatAssessmentRequests(res)
        return nil
    }
    return res
}
func (m *InformationProtection) IsNil()(bool) {
    return m == nil
}
func (m *InformationProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("bitlocker", m.GetBitlocker())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDataLossPreventionPolicies()))
        for i, v := range m.GetDataLossPreventionPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("dataLossPreventionPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSensitivityLabels()))
        for i, v := range m.GetSensitivityLabels() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sensitivityLabels", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sensitivityPolicySettings", m.GetSensitivityPolicySettings())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetThreatAssessmentRequests()))
        for i, v := range m.GetThreatAssessmentRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("threatAssessmentRequests", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *InformationProtection) SetBitlocker(value *Bitlocker)() {
    m.bitlocker = value
}
func (m *InformationProtection) SetDataLossPreventionPolicies(value []DataLossPreventionPolicy)() {
    m.dataLossPreventionPolicies = value
}
func (m *InformationProtection) SetPolicy(value *InformationProtectionPolicy)() {
    m.policy = value
}
func (m *InformationProtection) SetSensitivityLabels(value []SensitivityLabel)() {
    m.sensitivityLabels = value
}
func (m *InformationProtection) SetSensitivityPolicySettings(value *SensitivityPolicySettings)() {
    m.sensitivityPolicySettings = value
}
func (m *InformationProtection) SetThreatAssessmentRequests(value []ThreatAssessmentRequest)() {
    m.threatAssessmentRequests = value
}