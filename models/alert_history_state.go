package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AlertHistoryState 
type AlertHistoryState struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAlertHistoryState instantiates a new alertHistoryState and sets the default values.
func NewAlertHistoryState()(*AlertHistoryState) {
    m := &AlertHistoryState{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAlertHistoryStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertHistoryStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertHistoryState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertHistoryState) GetAdditionalData()(map[string]any) {
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
// GetAppId gets the appId property value. The Application ID of the calling application that submitted an update (PATCH) to the alert. The appId should be extracted from the auth token and not entered manually by the calling application.
func (m *AlertHistoryState) GetAppId()(*string) {
    val, err := m.GetBackingStore().Get("appId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignedTo gets the assignedTo property value. UPN of user the alert was assigned to (note: alert.assignedTo only stores the last value/UPN).
func (m *AlertHistoryState) GetAssignedTo()(*string) {
    val, err := m.GetBackingStore().Get("assignedTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *AlertHistoryState) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetComments gets the comments property value. Comment entered by signed-in user.
func (m *AlertHistoryState) GetComments()([]string) {
    val, err := m.GetBackingStore().Get("comments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFeedback gets the feedback property value. Analyst feedback on the alert in this update. Possible values are: unknown, truePositive, falsePositive, benignPositive.
func (m *AlertHistoryState) GetFeedback()(*AlertFeedback) {
    val, err := m.GetBackingStore().Get("feedback")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AlertFeedback)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertHistoryState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["assignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val)
        }
        return nil
    }
    res["comments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetComments(res)
        }
        return nil
    }
    res["feedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertFeedback)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedback(val.(*AlertFeedback))
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AlertStatus))
        }
        return nil
    }
    res["updatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedDateTime(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AlertHistoryState) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. Alert status value (if updated). Possible values are: unknown, newAlert, inProgress, resolved, dismissed.
func (m *AlertHistoryState) GetStatus()(*AlertStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AlertStatus)
    }
    return nil
}
// GetUpdatedDateTime gets the updatedDateTime property value. Date and time of the alert update. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AlertHistoryState) GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("updatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetUser gets the user property value. UPN of the signed-in user that updated the alert (taken from the bearer token - if in user/delegated auth mode).
func (m *AlertHistoryState) GetUser()(*string) {
    val, err := m.GetBackingStore().Get("user")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AlertHistoryState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    if m.GetComments() != nil {
        err := writer.WriteCollectionOfStringValues("comments", m.GetComments())
        if err != nil {
            return err
        }
    }
    if m.GetFeedback() != nil {
        cast := (*m.GetFeedback()).String()
        err := writer.WriteStringValue("feedback", &cast)
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedDateTime", m.GetUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user", m.GetUser())
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
func (m *AlertHistoryState) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAppId sets the appId property value. The Application ID of the calling application that submitted an update (PATCH) to the alert. The appId should be extracted from the auth token and not entered manually by the calling application.
func (m *AlertHistoryState) SetAppId(value *string)() {
    err := m.GetBackingStore().Set("appId", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignedTo sets the assignedTo property value. UPN of user the alert was assigned to (note: alert.assignedTo only stores the last value/UPN).
func (m *AlertHistoryState) SetAssignedTo(value *string)() {
    err := m.GetBackingStore().Set("assignedTo", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *AlertHistoryState) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetComments sets the comments property value. Comment entered by signed-in user.
func (m *AlertHistoryState) SetComments(value []string)() {
    err := m.GetBackingStore().Set("comments", value)
    if err != nil {
        panic(err)
    }
}
// SetFeedback sets the feedback property value. Analyst feedback on the alert in this update. Possible values are: unknown, truePositive, falsePositive, benignPositive.
func (m *AlertHistoryState) SetFeedback(value *AlertFeedback)() {
    err := m.GetBackingStore().Set("feedback", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AlertHistoryState) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Alert status value (if updated). Possible values are: unknown, newAlert, inProgress, resolved, dismissed.
func (m *AlertHistoryState) SetStatus(value *AlertStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdatedDateTime sets the updatedDateTime property value. Date and time of the alert update. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AlertHistoryState) SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("updatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetUser sets the user property value. UPN of the signed-in user that updated the alert (taken from the bearer token - if in user/delegated auth mode).
func (m *AlertHistoryState) SetUser(value *string)() {
    err := m.GetBackingStore().Set("user", value)
    if err != nil {
        panic(err)
    }
}
// AlertHistoryStateable 
type AlertHistoryStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetAssignedTo()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetComments()([]string)
    GetFeedback()(*AlertFeedback)
    GetOdataType()(*string)
    GetStatus()(*AlertStatus)
    GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUser()(*string)
    SetAppId(value *string)()
    SetAssignedTo(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetComments(value []string)()
    SetFeedback(value *AlertFeedback)()
    SetOdataType(value *string)()
    SetStatus(value *AlertStatus)()
    SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUser(value *string)()
}
