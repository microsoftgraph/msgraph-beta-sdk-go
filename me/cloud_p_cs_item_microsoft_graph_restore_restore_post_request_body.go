package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody 
type CloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The cloudPcSnapshotId property
    cloudPcSnapshotId *string
}
// NewCloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody instantiates a new CloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody and sets the default values.
func NewCloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody()(*CloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody) {
    m := &CloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPCsItemMicrosoftGraphRestoreRestorePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPCsItemMicrosoftGraphRestoreRestorePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCloudPcSnapshotId gets the cloudPcSnapshotId property value. The cloudPcSnapshotId property
func (m *CloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody) GetCloudPcSnapshotId()(*string) {
    return m.cloudPcSnapshotId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *CloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCloudPcSnapshotId sets the cloudPcSnapshotId property value. The cloudPcSnapshotId property
func (m *CloudPCsItemMicrosoftGraphRestoreRestorePostRequestBody) SetCloudPcSnapshotId(value *string)() {
    m.cloudPcSnapshotId = value
}
