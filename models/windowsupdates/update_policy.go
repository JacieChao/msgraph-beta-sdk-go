package windowsupdates

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// UpdatePolicy 
type UpdatePolicy struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewUpdatePolicy instantiates a new updatePolicy and sets the default values.
func NewUpdatePolicy()(*UpdatePolicy) {
    m := &UpdatePolicy{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateUpdatePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdatePolicy(), nil
}
// GetAudience gets the audience property value. Specifies the audience to target.
func (m *UpdatePolicy) GetAudience()(DeploymentAudienceable) {
    val, err := m.GetBackingStore().Get("audience")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeploymentAudienceable)
    }
    return nil
}
// GetComplianceChangeRules gets the complianceChangeRules property value. Rules for governing the automatic creation of compliance changes.
func (m *UpdatePolicy) GetComplianceChangeRules()([]ComplianceChangeRuleable) {
    val, err := m.GetBackingStore().Get("complianceChangeRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ComplianceChangeRuleable)
    }
    return nil
}
// GetComplianceChanges gets the complianceChanges property value. Compliance changes like content approvals which result in the automatic creation of deployments using the audience and deploymentSettings of the policy.
func (m *UpdatePolicy) GetComplianceChanges()([]ComplianceChangeable) {
    val, err := m.GetBackingStore().Get("complianceChanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ComplianceChangeable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the update policy was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *UpdatePolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeploymentSettings gets the deploymentSettings property value. Settings for governing how to deploy content.
func (m *UpdatePolicy) GetDeploymentSettings()(DeploymentSettingsable) {
    val, err := m.GetBackingStore().Get("deploymentSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeploymentSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["audience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeploymentAudienceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudience(val.(DeploymentAudienceable))
        }
        return nil
    }
    res["complianceChangeRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceChangeRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceChangeRuleable, len(val))
            for i, v := range val {
                res[i] = v.(ComplianceChangeRuleable)
            }
            m.SetComplianceChangeRules(res)
        }
        return nil
    }
    res["complianceChanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceChangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceChangeable, len(val))
            for i, v := range val {
                res[i] = v.(ComplianceChangeable)
            }
            m.SetComplianceChanges(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deploymentSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeploymentSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentSettings(val.(DeploymentSettingsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UpdatePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("audience", m.GetAudience())
        if err != nil {
            return err
        }
    }
    if m.GetComplianceChangeRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComplianceChangeRules()))
        for i, v := range m.GetComplianceChangeRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("complianceChangeRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceChanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComplianceChanges()))
        for i, v := range m.GetComplianceChanges() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("complianceChanges", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentSettings", m.GetDeploymentSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudience sets the audience property value. Specifies the audience to target.
func (m *UpdatePolicy) SetAudience(value DeploymentAudienceable)() {
    err := m.GetBackingStore().Set("audience", value)
    if err != nil {
        panic(err)
    }
}
// SetComplianceChangeRules sets the complianceChangeRules property value. Rules for governing the automatic creation of compliance changes.
func (m *UpdatePolicy) SetComplianceChangeRules(value []ComplianceChangeRuleable)() {
    err := m.GetBackingStore().Set("complianceChangeRules", value)
    if err != nil {
        panic(err)
    }
}
// SetComplianceChanges sets the complianceChanges property value. Compliance changes like content approvals which result in the automatic creation of deployments using the audience and deploymentSettings of the policy.
func (m *UpdatePolicy) SetComplianceChanges(value []ComplianceChangeable)() {
    err := m.GetBackingStore().Set("complianceChanges", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the update policy was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *UpdatePolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentSettings sets the deploymentSettings property value. Settings for governing how to deploy content.
func (m *UpdatePolicy) SetDeploymentSettings(value DeploymentSettingsable)() {
    err := m.GetBackingStore().Set("deploymentSettings", value)
    if err != nil {
        panic(err)
    }
}
// UpdatePolicyable 
type UpdatePolicyable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudience()(DeploymentAudienceable)
    GetComplianceChangeRules()([]ComplianceChangeRuleable)
    GetComplianceChanges()([]ComplianceChangeable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentSettings()(DeploymentSettingsable)
    SetAudience(value DeploymentAudienceable)()
    SetComplianceChangeRules(value []ComplianceChangeRuleable)()
    SetComplianceChanges(value []ComplianceChangeable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentSettings(value DeploymentSettingsable)()
}
