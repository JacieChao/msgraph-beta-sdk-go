package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus struct {
    Entity
    deploymentStatus *WindowsDefenderApplicationControlSupplementalPolicyStatuses;
    deviceId *string;
    deviceName *string;
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    osDescription *string;
    osVersion *string;
    policy *WindowsDefenderApplicationControlSupplementalPolicy;
    policyVersion *string;
    userName *string;
    userPrincipalName *string;
}
func NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus()(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetDeploymentStatus()(*WindowsDefenderApplicationControlSupplementalPolicyStatuses) {
    if m == nil {
        return nil
    } else {
        return m.deploymentStatus
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetOsDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osDescription
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetPolicy()(*WindowsDefenderApplicationControlSupplementalPolicy) {
    if m == nil {
        return nil
    } else {
        return m.policy
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetPolicyVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyVersion
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deploymentStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDefenderApplicationControlSupplementalPolicyStatuses)
        if err != nil {
            return err
        }
        cast := val.(WindowsDefenderApplicationControlSupplementalPolicyStatuses)
        m.SetDeploymentStatus(&cast)
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSyncDateTime(val)
        return nil
    }
    res["osDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsDescription(val)
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsVersion(val)
        return nil
    }
    res["policy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDefenderApplicationControlSupplementalPolicy() })
        if err != nil {
            return err
        }
        m.SetPolicy(val.(*WindowsDefenderApplicationControlSupplementalPolicy))
        return nil
    }
    res["policyVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyVersion(val)
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserName(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) IsNil()(bool) {
    return m == nil
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeploymentStatus() != nil {
        cast := m.GetDeploymentStatus().String()
        err = writer.WriteStringValue("deploymentStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osDescription", m.GetOsDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
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
        err = writer.WriteStringValue("policyVersion", m.GetPolicyVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetDeploymentStatus(value *WindowsDefenderApplicationControlSupplementalPolicyStatuses)() {
    m.deploymentStatus = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetDeviceId(value *string)() {
    m.deviceId = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetOsDescription(value *string)() {
    m.osDescription = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetOsVersion(value *string)() {
    m.osVersion = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetPolicy(value *WindowsDefenderApplicationControlSupplementalPolicy)() {
    m.policy = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetPolicyVersion(value *string)() {
    m.policyVersion = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetUserName(value *string)() {
    m.userName = value
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}