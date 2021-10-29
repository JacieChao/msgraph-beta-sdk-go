package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApprovalWorkflowProvider struct {
    Entity
    businessFlows []BusinessFlow;
    businessFlowsWithRequestsAwaitingMyDecision []BusinessFlow;
    displayName *string;
    policyTemplates []GovernancePolicyTemplate;
}
func NewApprovalWorkflowProvider()(*ApprovalWorkflowProvider) {
    m := &ApprovalWorkflowProvider{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ApprovalWorkflowProvider) GetBusinessFlows()([]BusinessFlow) {
    if m == nil {
        return nil
    } else {
        return m.businessFlows
    }
}
func (m *ApprovalWorkflowProvider) GetBusinessFlowsWithRequestsAwaitingMyDecision()([]BusinessFlow) {
    if m == nil {
        return nil
    } else {
        return m.businessFlowsWithRequestsAwaitingMyDecision
    }
}
func (m *ApprovalWorkflowProvider) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *ApprovalWorkflowProvider) GetPolicyTemplates()([]GovernancePolicyTemplate) {
    if m == nil {
        return nil
    } else {
        return m.policyTemplates
    }
}
func (m *ApprovalWorkflowProvider) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["businessFlows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBusinessFlow() })
        if err != nil {
            return err
        }
        res := make([]BusinessFlow, len(val))
        for i, v := range val {
            res[i] = *(v.(*BusinessFlow))
        }
        m.SetBusinessFlows(res)
        return nil
    }
    res["businessFlowsWithRequestsAwaitingMyDecision"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBusinessFlow() })
        if err != nil {
            return err
        }
        res := make([]BusinessFlow, len(val))
        for i, v := range val {
            res[i] = *(v.(*BusinessFlow))
        }
        m.SetBusinessFlowsWithRequestsAwaitingMyDecision(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["policyTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernancePolicyTemplate() })
        if err != nil {
            return err
        }
        res := make([]GovernancePolicyTemplate, len(val))
        for i, v := range val {
            res[i] = *(v.(*GovernancePolicyTemplate))
        }
        m.SetPolicyTemplates(res)
        return nil
    }
    return res
}
func (m *ApprovalWorkflowProvider) IsNil()(bool) {
    return m == nil
}
func (m *ApprovalWorkflowProvider) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBusinessFlows()))
        for i, v := range m.GetBusinessFlows() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("businessFlows", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBusinessFlowsWithRequestsAwaitingMyDecision()))
        for i, v := range m.GetBusinessFlowsWithRequestsAwaitingMyDecision() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("businessFlowsWithRequestsAwaitingMyDecision", cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPolicyTemplates()))
        for i, v := range m.GetPolicyTemplates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("policyTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ApprovalWorkflowProvider) SetBusinessFlows(value []BusinessFlow)() {
    m.businessFlows = value
}
func (m *ApprovalWorkflowProvider) SetBusinessFlowsWithRequestsAwaitingMyDecision(value []BusinessFlow)() {
    m.businessFlowsWithRequestsAwaitingMyDecision = value
}
func (m *ApprovalWorkflowProvider) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *ApprovalWorkflowProvider) SetPolicyTemplates(value []GovernancePolicyTemplate)() {
    m.policyTemplates = value
}