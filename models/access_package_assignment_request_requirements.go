package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AccessPackageAssignmentRequestRequirements 
type AccessPackageAssignmentRequestRequirements struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAccessPackageAssignmentRequestRequirements instantiates a new accessPackageAssignmentRequestRequirements and sets the default values.
func NewAccessPackageAssignmentRequestRequirements()(*AccessPackageAssignmentRequestRequirements) {
    m := &AccessPackageAssignmentRequestRequirements{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessPackageAssignmentRequestRequirementsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentRequestRequirementsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentRequestRequirements(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAssignmentRequestRequirements) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *AccessPackageAssignmentRequestRequirements) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExistingAnswers gets the existingAnswers property value. Answers that have already been provided.
func (m *AccessPackageAssignmentRequestRequirements) GetExistingAnswers()([]AccessPackageAnswerable) {
    val, err := m.GetBackingStore().Get("existingAnswers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageAnswerable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentRequestRequirements) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["existingAnswers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAnswerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAnswerable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAnswerable)
            }
            m.SetExistingAnswers(res)
        }
        return nil
    }
    res["isApprovalRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequired(val)
        }
        return nil
    }
    res["isApprovalRequiredForExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequiredForExtension(val)
        }
        return nil
    }
    res["isCustomAssignmentScheduleAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomAssignmentScheduleAllowed(val)
        }
        return nil
    }
    res["isRequestorJustificationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequestorJustificationRequired(val)
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
    res["policyDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyDescription(val)
        }
        return nil
    }
    res["policyDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyDisplayName(val)
        }
        return nil
    }
    res["policyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
        }
        return nil
    }
    res["questions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageQuestionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageQuestionable)
            }
            m.SetQuestions(res)
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(RequestScheduleable))
        }
        return nil
    }
    return res
}
// GetIsApprovalRequired gets the isApprovalRequired property value. Indicates whether a request must be approved by an approver.
func (m *AccessPackageAssignmentRequestRequirements) GetIsApprovalRequired()(*bool) {
    val, err := m.GetBackingStore().Get("isApprovalRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsApprovalRequiredForExtension gets the isApprovalRequiredForExtension property value. Indicates whether approval is required when a user tries to extend their access.
func (m *AccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForExtension()(*bool) {
    val, err := m.GetBackingStore().Get("isApprovalRequiredForExtension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsCustomAssignmentScheduleAllowed gets the isCustomAssignmentScheduleAllowed property value. Indicates whether the requestor is allowed to set a custom schedule.
func (m *AccessPackageAssignmentRequestRequirements) GetIsCustomAssignmentScheduleAllowed()(*bool) {
    val, err := m.GetBackingStore().Get("isCustomAssignmentScheduleAllowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsRequestorJustificationRequired gets the isRequestorJustificationRequired property value. Indicates whether a requestor must supply justification when submitting an assignment request.
func (m *AccessPackageAssignmentRequestRequirements) GetIsRequestorJustificationRequired()(*bool) {
    val, err := m.GetBackingStore().Get("isRequestorJustificationRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessPackageAssignmentRequestRequirements) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyDescription gets the policyDescription property value. The description of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) GetPolicyDescription()(*string) {
    val, err := m.GetBackingStore().Get("policyDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyDisplayName gets the policyDisplayName property value. The display name of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) GetPolicyDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("policyDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPolicyId gets the policyId property value. The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request.
func (m *AccessPackageAssignmentRequestRequirements) GetPolicyId()(*string) {
    val, err := m.GetBackingStore().Get("policyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetQuestions gets the questions property value. Questions that are configured on the policy. The questions can be required or optional; callers can determine whether a question is required or optional based on the isRequired property on accessPackageQuestion.
func (m *AccessPackageAssignmentRequestRequirements) GetQuestions()([]AccessPackageQuestionable) {
    val, err := m.GetBackingStore().Get("questions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageQuestionable)
    }
    return nil
}
// GetSchedule gets the schedule property value. Schedule restrictions enforced, if any.
func (m *AccessPackageAssignmentRequestRequirements) GetSchedule()(RequestScheduleable) {
    val, err := m.GetBackingStore().Get("schedule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RequestScheduleable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentRequestRequirements) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExistingAnswers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExistingAnswers()))
        for i, v := range m.GetExistingAnswers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("existingAnswers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequired", m.GetIsApprovalRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequiredForExtension", m.GetIsApprovalRequiredForExtension())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCustomAssignmentScheduleAllowed", m.GetIsCustomAssignmentScheduleAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRequestorJustificationRequired", m.GetIsRequestorJustificationRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyDescription", m.GetPolicyDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyDisplayName", m.GetPolicyDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyId", m.GetPolicyId())
        if err != nil {
            return err
        }
    }
    if m.GetQuestions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetQuestions()))
        for i, v := range m.GetQuestions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("questions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("schedule", m.GetSchedule())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessPackageAssignmentRequestRequirements) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *AccessPackageAssignmentRequestRequirements) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetExistingAnswers sets the existingAnswers property value. Answers that have already been provided.
func (m *AccessPackageAssignmentRequestRequirements) SetExistingAnswers(value []AccessPackageAnswerable)() {
    err := m.GetBackingStore().Set("existingAnswers", value)
    if err != nil {
        panic(err)
    }
}
// SetIsApprovalRequired sets the isApprovalRequired property value. Indicates whether a request must be approved by an approver.
func (m *AccessPackageAssignmentRequestRequirements) SetIsApprovalRequired(value *bool)() {
    err := m.GetBackingStore().Set("isApprovalRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetIsApprovalRequiredForExtension sets the isApprovalRequiredForExtension property value. Indicates whether approval is required when a user tries to extend their access.
func (m *AccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForExtension(value *bool)() {
    err := m.GetBackingStore().Set("isApprovalRequiredForExtension", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCustomAssignmentScheduleAllowed sets the isCustomAssignmentScheduleAllowed property value. Indicates whether the requestor is allowed to set a custom schedule.
func (m *AccessPackageAssignmentRequestRequirements) SetIsCustomAssignmentScheduleAllowed(value *bool)() {
    err := m.GetBackingStore().Set("isCustomAssignmentScheduleAllowed", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRequestorJustificationRequired sets the isRequestorJustificationRequired property value. Indicates whether a requestor must supply justification when submitting an assignment request.
func (m *AccessPackageAssignmentRequestRequirements) SetIsRequestorJustificationRequired(value *bool)() {
    err := m.GetBackingStore().Set("isRequestorJustificationRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageAssignmentRequestRequirements) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyDescription sets the policyDescription property value. The description of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) SetPolicyDescription(value *string)() {
    err := m.GetBackingStore().Set("policyDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyDisplayName sets the policyDisplayName property value. The display name of the policy that the user is trying to request access using.
func (m *AccessPackageAssignmentRequestRequirements) SetPolicyDisplayName(value *string)() {
    err := m.GetBackingStore().Set("policyDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyId sets the policyId property value. The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request.
func (m *AccessPackageAssignmentRequestRequirements) SetPolicyId(value *string)() {
    err := m.GetBackingStore().Set("policyId", value)
    if err != nil {
        panic(err)
    }
}
// SetQuestions sets the questions property value. Questions that are configured on the policy. The questions can be required or optional; callers can determine whether a question is required or optional based on the isRequired property on accessPackageQuestion.
func (m *AccessPackageAssignmentRequestRequirements) SetQuestions(value []AccessPackageQuestionable)() {
    err := m.GetBackingStore().Set("questions", value)
    if err != nil {
        panic(err)
    }
}
// SetSchedule sets the schedule property value. Schedule restrictions enforced, if any.
func (m *AccessPackageAssignmentRequestRequirements) SetSchedule(value RequestScheduleable)() {
    err := m.GetBackingStore().Set("schedule", value)
    if err != nil {
        panic(err)
    }
}
// AccessPackageAssignmentRequestRequirementsable 
type AccessPackageAssignmentRequestRequirementsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetExistingAnswers()([]AccessPackageAnswerable)
    GetIsApprovalRequired()(*bool)
    GetIsApprovalRequiredForExtension()(*bool)
    GetIsCustomAssignmentScheduleAllowed()(*bool)
    GetIsRequestorJustificationRequired()(*bool)
    GetOdataType()(*string)
    GetPolicyDescription()(*string)
    GetPolicyDisplayName()(*string)
    GetPolicyId()(*string)
    GetQuestions()([]AccessPackageQuestionable)
    GetSchedule()(RequestScheduleable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetExistingAnswers(value []AccessPackageAnswerable)()
    SetIsApprovalRequired(value *bool)()
    SetIsApprovalRequiredForExtension(value *bool)()
    SetIsCustomAssignmentScheduleAllowed(value *bool)()
    SetIsRequestorJustificationRequired(value *bool)()
    SetOdataType(value *string)()
    SetPolicyDescription(value *string)()
    SetPolicyDisplayName(value *string)()
    SetPolicyId(value *string)()
    SetQuestions(value []AccessPackageQuestionable)()
    SetSchedule(value RequestScheduleable)()
}
