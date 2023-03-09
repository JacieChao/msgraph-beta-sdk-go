package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserRegistrationFeatureSummary 
type UserRegistrationFeatureSummary struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserRegistrationFeatureSummary instantiates a new UserRegistrationFeatureSummary and sets the default values.
func NewUserRegistrationFeatureSummary()(*UserRegistrationFeatureSummary) {
    m := &UserRegistrationFeatureSummary{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserRegistrationFeatureSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserRegistrationFeatureSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserRegistrationFeatureSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserRegistrationFeatureSummary) GetAdditionalData()(map[string]any) {
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
func (m *UserRegistrationFeatureSummary) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserRegistrationFeatureSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["totalUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUserCount(val)
        }
        return nil
    }
    res["userRegistrationFeatureCounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserRegistrationFeatureCountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserRegistrationFeatureCountable, len(val))
            for i, v := range val {
                res[i] = v.(UserRegistrationFeatureCountable)
            }
            m.SetUserRegistrationFeatureCounts(res)
        }
        return nil
    }
    res["userRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncludedUserRoles)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRoles(val.(*IncludedUserRoles))
        }
        return nil
    }
    res["userTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIncludedUserTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserTypes(val.(*IncludedUserTypes))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserRegistrationFeatureSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalUserCount gets the totalUserCount property value. Total number of users accounts, excluding those that are blocked
func (m *UserRegistrationFeatureSummary) GetTotalUserCount()(*int64) {
    val, err := m.GetBackingStore().Get("totalUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetUserRegistrationFeatureCounts gets the userRegistrationFeatureCounts property value. Number of users registered or capable for Multi-Factor Authentication, Self-Service Password Reset and Passwordless Authentication.
func (m *UserRegistrationFeatureSummary) GetUserRegistrationFeatureCounts()([]UserRegistrationFeatureCountable) {
    val, err := m.GetBackingStore().Get("userRegistrationFeatureCounts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserRegistrationFeatureCountable)
    }
    return nil
}
// GetUserRoles gets the userRoles property value. User role type. Possible values are: all, privilegedAdmin, admin, user.
func (m *UserRegistrationFeatureSummary) GetUserRoles()(*IncludedUserRoles) {
    val, err := m.GetBackingStore().Get("userRoles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IncludedUserRoles)
    }
    return nil
}
// GetUserTypes gets the userTypes property value. User type. Possible values are: all, member, guest.
func (m *UserRegistrationFeatureSummary) GetUserTypes()(*IncludedUserTypes) {
    val, err := m.GetBackingStore().Get("userTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IncludedUserTypes)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserRegistrationFeatureSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalUserCount", m.GetTotalUserCount())
        if err != nil {
            return err
        }
    }
    if m.GetUserRegistrationFeatureCounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserRegistrationFeatureCounts()))
        for i, v := range m.GetUserRegistrationFeatureCounts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("userRegistrationFeatureCounts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserRoles() != nil {
        cast := (*m.GetUserRoles()).String()
        err := writer.WriteStringValue("userRoles", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserTypes() != nil {
        cast := (*m.GetUserTypes()).String()
        err := writer.WriteStringValue("userTypes", &cast)
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
func (m *UserRegistrationFeatureSummary) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *UserRegistrationFeatureSummary) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserRegistrationFeatureSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalUserCount sets the totalUserCount property value. Total number of users accounts, excluding those that are blocked
func (m *UserRegistrationFeatureSummary) SetTotalUserCount(value *int64)() {
    err := m.GetBackingStore().Set("totalUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRegistrationFeatureCounts sets the userRegistrationFeatureCounts property value. Number of users registered or capable for Multi-Factor Authentication, Self-Service Password Reset and Passwordless Authentication.
func (m *UserRegistrationFeatureSummary) SetUserRegistrationFeatureCounts(value []UserRegistrationFeatureCountable)() {
    err := m.GetBackingStore().Set("userRegistrationFeatureCounts", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRoles sets the userRoles property value. User role type. Possible values are: all, privilegedAdmin, admin, user.
func (m *UserRegistrationFeatureSummary) SetUserRoles(value *IncludedUserRoles)() {
    err := m.GetBackingStore().Set("userRoles", value)
    if err != nil {
        panic(err)
    }
}
// SetUserTypes sets the userTypes property value. User type. Possible values are: all, member, guest.
func (m *UserRegistrationFeatureSummary) SetUserTypes(value *IncludedUserTypes)() {
    err := m.GetBackingStore().Set("userTypes", value)
    if err != nil {
        panic(err)
    }
}
// UserRegistrationFeatureSummaryable 
type UserRegistrationFeatureSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetTotalUserCount()(*int64)
    GetUserRegistrationFeatureCounts()([]UserRegistrationFeatureCountable)
    GetUserRoles()(*IncludedUserRoles)
    GetUserTypes()(*IncludedUserTypes)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetTotalUserCount(value *int64)()
    SetUserRegistrationFeatureCounts(value []UserRegistrationFeatureCountable)()
    SetUserRoles(value *IncludedUserRoles)()
    SetUserTypes(value *IncludedUserTypes)()
}
