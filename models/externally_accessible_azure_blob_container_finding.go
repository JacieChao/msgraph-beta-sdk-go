package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternallyAccessibleAzureBlobContainerFinding 
type ExternallyAccessibleAzureBlobContainerFinding struct {
    Finding
}
// NewExternallyAccessibleAzureBlobContainerFinding instantiates a new externallyAccessibleAzureBlobContainerFinding and sets the default values.
func NewExternallyAccessibleAzureBlobContainerFinding()(*ExternallyAccessibleAzureBlobContainerFinding) {
    m := &ExternallyAccessibleAzureBlobContainerFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateExternallyAccessibleAzureBlobContainerFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternallyAccessibleAzureBlobContainerFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternallyAccessibleAzureBlobContainerFinding(), nil
}
// GetAccessibility gets the accessibility property value. The accessibility property
func (m *ExternallyAccessibleAzureBlobContainerFinding) GetAccessibility()(*AzureAccessType) {
    val, err := m.GetBackingStore().Get("accessibility")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AzureAccessType)
    }
    return nil
}
// GetEncryptionManagedBy gets the encryptionManagedBy property value. The encryptionManagedBy property
func (m *ExternallyAccessibleAzureBlobContainerFinding) GetEncryptionManagedBy()(*AzureEncryption) {
    val, err := m.GetBackingStore().Get("encryptionManagedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AzureEncryption)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternallyAccessibleAzureBlobContainerFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
    res["accessibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAzureAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessibility(val.(*AzureAccessType))
        }
        return nil
    }
    res["encryptionManagedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAzureEncryption)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionManagedBy(val.(*AzureEncryption))
        }
        return nil
    }
    res["storageAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageAccount(val.(AuthorizationSystemResourceable))
        }
        return nil
    }
    return res
}
// GetStorageAccount gets the storageAccount property value. The storageAccount property
func (m *ExternallyAccessibleAzureBlobContainerFinding) GetStorageAccount()(AuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("storageAccount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemResourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExternallyAccessibleAzureBlobContainerFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Finding.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessibility() != nil {
        cast := (*m.GetAccessibility()).String()
        err = writer.WriteStringValue("accessibility", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEncryptionManagedBy() != nil {
        cast := (*m.GetEncryptionManagedBy()).String()
        err = writer.WriteStringValue("encryptionManagedBy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("storageAccount", m.GetStorageAccount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessibility sets the accessibility property value. The accessibility property
func (m *ExternallyAccessibleAzureBlobContainerFinding) SetAccessibility(value *AzureAccessType)() {
    err := m.GetBackingStore().Set("accessibility", value)
    if err != nil {
        panic(err)
    }
}
// SetEncryptionManagedBy sets the encryptionManagedBy property value. The encryptionManagedBy property
func (m *ExternallyAccessibleAzureBlobContainerFinding) SetEncryptionManagedBy(value *AzureEncryption)() {
    err := m.GetBackingStore().Set("encryptionManagedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageAccount sets the storageAccount property value. The storageAccount property
func (m *ExternallyAccessibleAzureBlobContainerFinding) SetStorageAccount(value AuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("storageAccount", value)
    if err != nil {
        panic(err)
    }
}
// ExternallyAccessibleAzureBlobContainerFindingable 
type ExternallyAccessibleAzureBlobContainerFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessibility()(*AzureAccessType)
    GetEncryptionManagedBy()(*AzureEncryption)
    GetStorageAccount()(AuthorizationSystemResourceable)
    SetAccessibility(value *AzureAccessType)()
    SetEncryptionManagedBy(value *AzureEncryption)()
    SetStorageAccount(value AuthorizationSystemResourceable)()
}
