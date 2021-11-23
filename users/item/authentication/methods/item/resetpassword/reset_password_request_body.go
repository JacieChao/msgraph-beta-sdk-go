package resetpassword

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// resetPasswordRequestBody 
type ResetPasswordRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    newPassword *string;
    // 
    requireChangeOnNextSignIn *bool;
}
// NewResetPasswordRequestBody instantiates a new resetPasswordRequestBody and sets the default values.
func NewResetPasswordRequestBody()(*ResetPasswordRequestBody) {
    m := &ResetPasswordRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResetPasswordRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetNewPassword gets the newPassword property value. 
func (m *ResetPasswordRequestBody) GetNewPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newPassword
    }
}
// GetRequireChangeOnNextSignIn gets the requireChangeOnNextSignIn property value. 
func (m *ResetPasswordRequestBody) GetRequireChangeOnNextSignIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requireChangeOnNextSignIn
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResetPasswordRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["newPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewPassword(val)
        }
        return nil
    }
    res["requireChangeOnNextSignIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireChangeOnNextSignIn(val)
        }
        return nil
    }
    return res
}
func (m *ResetPasswordRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ResetPasswordRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newPassword", m.GetNewPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("requireChangeOnNextSignIn", m.GetRequireChangeOnNextSignIn())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResetPasswordRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNewPassword sets the newPassword property value. 
func (m *ResetPasswordRequestBody) SetNewPassword(value *string)() {
    m.newPassword = value
}
// SetRequireChangeOnNextSignIn sets the requireChangeOnNextSignIn property value. 
func (m *ResetPasswordRequestBody) SetRequireChangeOnNextSignIn(value *bool)() {
    m.requireChangeOnNextSignIn = value
}
