package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EncryptedAzureStorageAccountFinding struct {
    Finding
}
// NewEncryptedAzureStorageAccountFinding instantiates a new EncryptedAzureStorageAccountFinding and sets the default values.
func NewEncryptedAzureStorageAccountFinding()(*EncryptedAzureStorageAccountFinding) {
    m := &EncryptedAzureStorageAccountFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateEncryptedAzureStorageAccountFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEncryptedAzureStorageAccountFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEncryptedAzureStorageAccountFinding(), nil
}
// GetEncryptionManagedBy gets the encryptionManagedBy property value. The encryptionManagedBy property
// returns a *AzureEncryption when successful
func (m *EncryptedAzureStorageAccountFinding) GetEncryptionManagedBy()(*AzureEncryption) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EncryptedAzureStorageAccountFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
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
// returns a AuthorizationSystemResourceable when successful
func (m *EncryptedAzureStorageAccountFinding) GetStorageAccount()(AuthorizationSystemResourceable) {
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
func (m *EncryptedAzureStorageAccountFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Finding.Serialize(writer)
    if err != nil {
        return err
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
// SetEncryptionManagedBy sets the encryptionManagedBy property value. The encryptionManagedBy property
func (m *EncryptedAzureStorageAccountFinding) SetEncryptionManagedBy(value *AzureEncryption)() {
    err := m.GetBackingStore().Set("encryptionManagedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageAccount sets the storageAccount property value. The storageAccount property
func (m *EncryptedAzureStorageAccountFinding) SetStorageAccount(value AuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("storageAccount", value)
    if err != nil {
        panic(err)
    }
}
type EncryptedAzureStorageAccountFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncryptionManagedBy()(*AzureEncryption)
    GetStorageAccount()(AuthorizationSystemResourceable)
    SetEncryptionManagedBy(value *AzureEncryption)()
    SetStorageAccount(value AuthorizationSystemResourceable)()
}
