package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcBulkCreateSnapshot struct {
    CloudPcBulkAction
}
// NewCloudPcBulkCreateSnapshot instantiates a new CloudPcBulkCreateSnapshot and sets the default values.
func NewCloudPcBulkCreateSnapshot()(*CloudPcBulkCreateSnapshot) {
    m := &CloudPcBulkCreateSnapshot{
        CloudPcBulkAction: *NewCloudPcBulkAction(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcBulkCreateSnapshot"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCloudPcBulkCreateSnapshotFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcBulkCreateSnapshotFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcBulkCreateSnapshot(), nil
}
// GetAccessTier gets the accessTier property value. Indicates the access tier of the blob file that the snapshot is copied to. Possible values are hot, cool, cold, archive, and unknownFutureValue. The default value is hot. Read-Only.
// returns a *CloudPcBlobAccessTier when successful
func (m *CloudPcBulkCreateSnapshot) GetAccessTier()(*CloudPcBlobAccessTier) {
    val, err := m.GetBackingStore().Get("accessTier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcBlobAccessTier)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcBulkCreateSnapshot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CloudPcBulkAction.GetFieldDeserializers()
    res["accessTier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcBlobAccessTier)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessTier(val.(*CloudPcBlobAccessTier))
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
    return res
}
// GetStorageAccountId gets the storageAccountId property value. The unique identifier for Secure Azure Storage Account, which receives the restore points (snapshots). The value can't be modified after it's created. For example, '/subscriptions/06199b73-30a1-4922-8734-93feca64cdf6/resourceGroups/res2627/providers/Microsoft.Storage/storageAccounts/sto1125'. Read-Only.
// returns a *string when successful
func (m *CloudPcBulkCreateSnapshot) GetStorageAccountId()(*string) {
    val, err := m.GetBackingStore().Get("storageAccountId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcBulkCreateSnapshot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CloudPcBulkAction.Serialize(writer)
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
        err = writer.WriteStringValue("storageAccountId", m.GetStorageAccountId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessTier sets the accessTier property value. Indicates the access tier of the blob file that the snapshot is copied to. Possible values are hot, cool, cold, archive, and unknownFutureValue. The default value is hot. Read-Only.
func (m *CloudPcBulkCreateSnapshot) SetAccessTier(value *CloudPcBlobAccessTier)() {
    err := m.GetBackingStore().Set("accessTier", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageAccountId sets the storageAccountId property value. The unique identifier for Secure Azure Storage Account, which receives the restore points (snapshots). The value can't be modified after it's created. For example, '/subscriptions/06199b73-30a1-4922-8734-93feca64cdf6/resourceGroups/res2627/providers/Microsoft.Storage/storageAccounts/sto1125'. Read-Only.
func (m *CloudPcBulkCreateSnapshot) SetStorageAccountId(value *string)() {
    err := m.GetBackingStore().Set("storageAccountId", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcBulkCreateSnapshotable interface {
    CloudPcBulkActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessTier()(*CloudPcBlobAccessTier)
    GetStorageAccountId()(*string)
    SetAccessTier(value *CloudPcBlobAccessTier)()
    SetStorageAccountId(value *string)()
}
