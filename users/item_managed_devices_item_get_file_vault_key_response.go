package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemManagedDevicesItemGetFileVaultKeyResponse 
// Deprecated: This class is obsolete. Use getFileVaultKeyGetResponse instead.
type ItemManagedDevicesItemGetFileVaultKeyResponse struct {
    ItemManagedDevicesItemGetFileVaultKeyGetResponse
}
// NewItemManagedDevicesItemGetFileVaultKeyResponse instantiates a new ItemManagedDevicesItemGetFileVaultKeyResponse and sets the default values.
func NewItemManagedDevicesItemGetFileVaultKeyResponse()(*ItemManagedDevicesItemGetFileVaultKeyResponse) {
    m := &ItemManagedDevicesItemGetFileVaultKeyResponse{
        ItemManagedDevicesItemGetFileVaultKeyGetResponse: *NewItemManagedDevicesItemGetFileVaultKeyGetResponse(),
    }
    return m
}
// CreateItemManagedDevicesItemGetFileVaultKeyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemManagedDevicesItemGetFileVaultKeyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemManagedDevicesItemGetFileVaultKeyResponse(), nil
}
// ItemManagedDevicesItemGetFileVaultKeyResponseable 
// Deprecated: This class is obsolete. Use getFileVaultKeyGetResponse instead.
type ItemManagedDevicesItemGetFileVaultKeyResponseable interface {
    ItemManagedDevicesItemGetFileVaultKeyGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
