package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnterpriseCodeSigningCertificate 
type EnterpriseCodeSigningCertificate struct {
    Entity
}
// NewEnterpriseCodeSigningCertificate instantiates a new enterpriseCodeSigningCertificate and sets the default values.
func NewEnterpriseCodeSigningCertificate()(*EnterpriseCodeSigningCertificate) {
    m := &EnterpriseCodeSigningCertificate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEnterpriseCodeSigningCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnterpriseCodeSigningCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseCodeSigningCertificate(), nil
}
// GetContent gets the content property value. The Windows Enterprise Code-Signing Certificate in the raw data format.
func (m *EnterpriseCodeSigningCertificate) GetContent()([]byte) {
    val, err := m.GetBackingStore().Get("content")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The Cert Expiration Date.
func (m *EnterpriseCodeSigningCertificate) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnterpriseCodeSigningCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["issuerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CertificateStatus))
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["subjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectName(val)
        }
        return nil
    }
    res["uploadDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadDateTime(val)
        }
        return nil
    }
    return res
}
// GetIssuer gets the issuer property value. The Issuer value for the cert.
func (m *EnterpriseCodeSigningCertificate) GetIssuer()(*string) {
    val, err := m.GetBackingStore().Get("issuer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIssuerName gets the issuerName property value. The Issuer Name for the cert.
func (m *EnterpriseCodeSigningCertificate) GetIssuerName()(*string) {
    val, err := m.GetBackingStore().Get("issuerName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *EnterpriseCodeSigningCertificate) GetStatus()(*CertificateStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CertificateStatus)
    }
    return nil
}
// GetSubject gets the subject property value. The Subject Value for the cert.
func (m *EnterpriseCodeSigningCertificate) GetSubject()(*string) {
    val, err := m.GetBackingStore().Get("subject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubjectName gets the subjectName property value. The Subject Name for the cert.
func (m *EnterpriseCodeSigningCertificate) GetSubjectName()(*string) {
    val, err := m.GetBackingStore().Get("subjectName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUploadDateTime gets the uploadDateTime property value. The date time of CodeSigning Cert when it is uploaded.
func (m *EnterpriseCodeSigningCertificate) GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("uploadDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EnterpriseCodeSigningCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuerName", m.GetIssuerName())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectName", m.GetSubjectName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("uploadDateTime", m.GetUploadDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The Windows Enterprise Code-Signing Certificate in the raw data format.
func (m *EnterpriseCodeSigningCertificate) SetContent(value []byte)() {
    err := m.GetBackingStore().Set("content", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The Cert Expiration Date.
func (m *EnterpriseCodeSigningCertificate) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIssuer sets the issuer property value. The Issuer value for the cert.
func (m *EnterpriseCodeSigningCertificate) SetIssuer(value *string)() {
    err := m.GetBackingStore().Set("issuer", value)
    if err != nil {
        panic(err)
    }
}
// SetIssuerName sets the issuerName property value. The Issuer Name for the cert.
func (m *EnterpriseCodeSigningCertificate) SetIssuerName(value *string)() {
    err := m.GetBackingStore().Set("issuerName", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *EnterpriseCodeSigningCertificate) SetStatus(value *CertificateStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetSubject sets the subject property value. The Subject Value for the cert.
func (m *EnterpriseCodeSigningCertificate) SetSubject(value *string)() {
    err := m.GetBackingStore().Set("subject", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectName sets the subjectName property value. The Subject Name for the cert.
func (m *EnterpriseCodeSigningCertificate) SetSubjectName(value *string)() {
    err := m.GetBackingStore().Set("subjectName", value)
    if err != nil {
        panic(err)
    }
}
// SetUploadDateTime sets the uploadDateTime property value. The date time of CodeSigning Cert when it is uploaded.
func (m *EnterpriseCodeSigningCertificate) SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("uploadDateTime", value)
    if err != nil {
        panic(err)
    }
}
// EnterpriseCodeSigningCertificateable 
type EnterpriseCodeSigningCertificateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIssuer()(*string)
    GetIssuerName()(*string)
    GetStatus()(*CertificateStatus)
    GetSubject()(*string)
    GetSubjectName()(*string)
    GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetContent(value []byte)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIssuer(value *string)()
    SetIssuerName(value *string)()
    SetStatus(value *CertificateStatus)()
    SetSubject(value *string)()
    SetSubjectName(value *string)()
    SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
