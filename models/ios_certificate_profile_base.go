package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosCertificateProfileBase 
type IosCertificateProfileBase struct {
    IosCertificateProfile
}
// NewIosCertificateProfileBase instantiates a new IosCertificateProfileBase and sets the default values.
func NewIosCertificateProfileBase()(*IosCertificateProfileBase) {
    m := &IosCertificateProfileBase{
        IosCertificateProfile: *NewIosCertificateProfile(),
    }
    odataTypeValue := "#microsoft.graph.iosCertificateProfileBase"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosCertificateProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosCertificateProfileBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.iosPkcsCertificateProfile":
                        return NewIosPkcsCertificateProfile(), nil
                    case "#microsoft.graph.iosScepCertificateProfile":
                        return NewIosScepCertificateProfile(), nil
                }
            }
        }
    }
    return NewIosCertificateProfileBase(), nil
}
// GetCertificateValidityPeriodScale gets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *IosCertificateProfileBase) GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale) {
    val, err := m.GetBackingStore().Get("certificateValidityPeriodScale")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CertificateValidityPeriodScale)
    }
    return nil
}
// GetCertificateValidityPeriodValue gets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *IosCertificateProfileBase) GetCertificateValidityPeriodValue()(*int32) {
    val, err := m.GetBackingStore().Get("certificateValidityPeriodValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosCertificateProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosCertificateProfile.GetFieldDeserializers()
    res["certificateValidityPeriodScale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateValidityPeriodScale)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateValidityPeriodScale(val.(*CertificateValidityPeriodScale))
        }
        return nil
    }
    res["certificateValidityPeriodValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateValidityPeriodValue(val)
        }
        return nil
    }
    res["renewalThresholdPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenewalThresholdPercentage(val)
        }
        return nil
    }
    res["subjectAlternativeNameType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectAlternativeNameType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectAlternativeNameType(val.(*SubjectAlternativeNameType))
        }
        return nil
    }
    res["subjectNameFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppleSubjectNameFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectNameFormat(val.(*AppleSubjectNameFormat))
        }
        return nil
    }
    return res
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *IosCertificateProfileBase) GetRenewalThresholdPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("renewalThresholdPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSubjectAlternativeNameType gets the subjectAlternativeNameType property value. Certificate Subject Alternative Name type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
func (m *IosCertificateProfileBase) GetSubjectAlternativeNameType()(*SubjectAlternativeNameType) {
    val, err := m.GetBackingStore().Get("subjectAlternativeNameType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubjectAlternativeNameType)
    }
    return nil
}
// GetSubjectNameFormat gets the subjectNameFormat property value. Subject Name Format Options for Apple devices.
func (m *IosCertificateProfileBase) GetSubjectNameFormat()(*AppleSubjectNameFormat) {
    val, err := m.GetBackingStore().Get("subjectNameFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppleSubjectNameFormat)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosCertificateProfileBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosCertificateProfile.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateValidityPeriodScale() != nil {
        cast := (*m.GetCertificateValidityPeriodScale()).String()
        err = writer.WriteStringValue("certificateValidityPeriodScale", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("certificateValidityPeriodValue", m.GetCertificateValidityPeriodValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("renewalThresholdPercentage", m.GetRenewalThresholdPercentage())
        if err != nil {
            return err
        }
    }
    if m.GetSubjectAlternativeNameType() != nil {
        cast := (*m.GetSubjectAlternativeNameType()).String()
        err = writer.WriteStringValue("subjectAlternativeNameType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubjectNameFormat() != nil {
        cast := (*m.GetSubjectNameFormat()).String()
        err = writer.WriteStringValue("subjectNameFormat", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateValidityPeriodScale sets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *IosCertificateProfileBase) SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)() {
    err := m.GetBackingStore().Set("certificateValidityPeriodScale", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateValidityPeriodValue sets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *IosCertificateProfileBase) SetCertificateValidityPeriodValue(value *int32)() {
    err := m.GetBackingStore().Set("certificateValidityPeriodValue", value)
    if err != nil {
        panic(err)
    }
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *IosCertificateProfileBase) SetRenewalThresholdPercentage(value *int32)() {
    err := m.GetBackingStore().Set("renewalThresholdPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectAlternativeNameType sets the subjectAlternativeNameType property value. Certificate Subject Alternative Name type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
func (m *IosCertificateProfileBase) SetSubjectAlternativeNameType(value *SubjectAlternativeNameType)() {
    err := m.GetBackingStore().Set("subjectAlternativeNameType", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectNameFormat sets the subjectNameFormat property value. Subject Name Format Options for Apple devices.
func (m *IosCertificateProfileBase) SetSubjectNameFormat(value *AppleSubjectNameFormat)() {
    err := m.GetBackingStore().Set("subjectNameFormat", value)
    if err != nil {
        panic(err)
    }
}
// IosCertificateProfileBaseable 
type IosCertificateProfileBaseable interface {
    IosCertificateProfileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale)
    GetCertificateValidityPeriodValue()(*int32)
    GetRenewalThresholdPercentage()(*int32)
    GetSubjectAlternativeNameType()(*SubjectAlternativeNameType)
    GetSubjectNameFormat()(*AppleSubjectNameFormat)
    SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)()
    SetCertificateValidityPeriodValue(value *int32)()
    SetRenewalThresholdPercentage(value *int32)()
    SetSubjectAlternativeNameType(value *SubjectAlternativeNameType)()
    SetSubjectNameFormat(value *AppleSubjectNameFormat)()
}
