package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCloudPCsItemRestorePostRequestBody provides operations to call the restore method.
type ItemCloudPCsItemRestorePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The cloudPcSnapshotId property
    cloudPcSnapshotId *string
}
// NewItemCloudPCsItemRestorePostRequestBody instantiates a new ItemCloudPCsItemRestorePostRequestBody and sets the default values.
func NewItemCloudPCsItemRestorePostRequestBody()(*ItemCloudPCsItemRestorePostRequestBody) {
    m := &ItemCloudPCsItemRestorePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemCloudPCsItemRestorePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCloudPCsItemRestorePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCloudPCsItemRestorePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCloudPCsItemRestorePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCloudPcSnapshotId gets the cloudPcSnapshotId property value. The cloudPcSnapshotId property
func (m *ItemCloudPCsItemRestorePostRequestBody) GetCloudPcSnapshotId()(*string) {
    return m.cloudPcSnapshotId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCloudPCsItemRestorePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudPcSnapshotId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcSnapshotId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemCloudPCsItemRestorePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cloudPcSnapshotId", m.GetCloudPcSnapshotId())
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
func (m *ItemCloudPCsItemRestorePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCloudPcSnapshotId sets the cloudPcSnapshotId property value. The cloudPcSnapshotId property
func (m *ItemCloudPCsItemRestorePostRequestBody) SetCloudPcSnapshotId(value *string)() {
    m.cloudPcSnapshotId = value
}
