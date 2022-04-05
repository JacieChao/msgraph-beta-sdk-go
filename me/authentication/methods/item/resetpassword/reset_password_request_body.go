package resetpassword

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResetPasswordRequestBody provides operations to call the resetPassword method.
type ResetPasswordRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The newPassword property
    newPassword *string;
    // The requireChangeOnNextSignIn property
    requireChangeOnNextSignIn *bool;
}
// NewResetPasswordRequestBody instantiates a new resetPasswordRequestBody and sets the default values.
func NewResetPasswordRequestBody()(*ResetPasswordRequestBody) {
    m := &ResetPasswordRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateResetPasswordRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResetPasswordRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResetPasswordRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResetPasswordRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResetPasswordRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newPassword"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewPassword(val)
        }
        return nil
    }
    res["requireChangeOnNextSignIn"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetNewPassword gets the newPassword property value. The newPassword property
func (m *ResetPasswordRequestBody) GetNewPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newPassword
    }
}
// GetRequireChangeOnNextSignIn gets the requireChangeOnNextSignIn property value. The requireChangeOnNextSignIn property
func (m *ResetPasswordRequestBody) GetRequireChangeOnNextSignIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requireChangeOnNextSignIn
    }
}
// Serialize serializes information the current object
func (m *ResetPasswordRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResetPasswordRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNewPassword sets the newPassword property value. The newPassword property
func (m *ResetPasswordRequestBody) SetNewPassword(value *string)() {
    if m != nil {
        m.newPassword = value
    }
}
// SetRequireChangeOnNextSignIn sets the requireChangeOnNextSignIn property value. The requireChangeOnNextSignIn property
func (m *ResetPasswordRequestBody) SetRequireChangeOnNextSignIn(value *bool)() {
    if m != nil {
        m.requireChangeOnNextSignIn = value
    }
}
