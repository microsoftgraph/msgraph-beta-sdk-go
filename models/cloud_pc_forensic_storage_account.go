package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcForensicStorageAccount struct {
    Entity
}
// NewCloudPcForensicStorageAccount instantiates a new CloudPcForensicStorageAccount and sets the default values.
func NewCloudPcForensicStorageAccount()(*CloudPcForensicStorageAccount) {
    m := &CloudPcForensicStorageAccount{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcForensicStorageAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcForensicStorageAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcForensicStorageAccount(), nil
}
// GetAccessTier gets the accessTier property value. The access tier of the storage account. Possible values are hot, cool, premium, cold, and unknownFutureValue. Default value is hot. Read-only.
// returns a *CloudPcStorageAccountAccessTier when successful
func (m *CloudPcForensicStorageAccount) GetAccessTier()(*CloudPcStorageAccountAccessTier) {
    val, err := m.GetBackingStore().Get("accessTier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcStorageAccountAccessTier)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcForensicStorageAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessTier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcStorageAccountAccessTier)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessTier(val.(*CloudPcStorageAccountAccessTier))
        }
        return nil
    }
    res["immutableStorage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImmutableStorage(val)
        }
        return nil
    }
    res["storageAccountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageAccountId(val)
        }
        return nil
    }
    res["storageAccountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageAccountName(val)
        }
        return nil
    }
    return res
}
// GetImmutableStorage gets the immutableStorage property value. Indicates whether immutability policies are configured for the storage account. When true, the storage account only accepts hot as the snapshot access tier. When false, the storage account accepts all valid access tiers. Read-Only.
// returns a *bool when successful
func (m *CloudPcForensicStorageAccount) GetImmutableStorage()(*bool) {
    val, err := m.GetBackingStore().Get("immutableStorage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStorageAccountId gets the storageAccountId property value. The ID of the storage account. Read-only.
// returns a *string when successful
func (m *CloudPcForensicStorageAccount) GetStorageAccountId()(*string) {
    val, err := m.GetBackingStore().Get("storageAccountId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStorageAccountName gets the storageAccountName property value. The name of the storage account. Read-only.
// returns a *string when successful
func (m *CloudPcForensicStorageAccount) GetStorageAccountName()(*string) {
    val, err := m.GetBackingStore().Get("storageAccountName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcForensicStorageAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessTier() != nil {
        cast := (*m.GetAccessTier()).String()
        err = writer.WriteStringValue("accessTier", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("immutableStorage", m.GetImmutableStorage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("storageAccountId", m.GetStorageAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("storageAccountName", m.GetStorageAccountName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessTier sets the accessTier property value. The access tier of the storage account. Possible values are hot, cool, premium, cold, and unknownFutureValue. Default value is hot. Read-only.
func (m *CloudPcForensicStorageAccount) SetAccessTier(value *CloudPcStorageAccountAccessTier)() {
    err := m.GetBackingStore().Set("accessTier", value)
    if err != nil {
        panic(err)
    }
}
// SetImmutableStorage sets the immutableStorage property value. Indicates whether immutability policies are configured for the storage account. When true, the storage account only accepts hot as the snapshot access tier. When false, the storage account accepts all valid access tiers. Read-Only.
func (m *CloudPcForensicStorageAccount) SetImmutableStorage(value *bool)() {
    err := m.GetBackingStore().Set("immutableStorage", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageAccountId sets the storageAccountId property value. The ID of the storage account. Read-only.
func (m *CloudPcForensicStorageAccount) SetStorageAccountId(value *string)() {
    err := m.GetBackingStore().Set("storageAccountId", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageAccountName sets the storageAccountName property value. The name of the storage account. Read-only.
func (m *CloudPcForensicStorageAccount) SetStorageAccountName(value *string)() {
    err := m.GetBackingStore().Set("storageAccountName", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcForensicStorageAccountable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessTier()(*CloudPcStorageAccountAccessTier)
    GetImmutableStorage()(*bool)
    GetStorageAccountId()(*string)
    GetStorageAccountName()(*string)
    SetAccessTier(value *CloudPcStorageAccountAccessTier)()
    SetImmutableStorage(value *bool)()
    SetStorageAccountId(value *string)()
    SetStorageAccountName(value *string)()
}
