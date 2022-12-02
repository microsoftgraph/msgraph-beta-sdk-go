package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeManagedDevicesBulkReprovisionCloudPcPostRequestBody provides operations to call the bulkReprovisionCloudPc method.
type MeManagedDevicesBulkReprovisionCloudPcPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The managedDeviceIds property
    managedDeviceIds []string
}
// NewMeManagedDevicesBulkReprovisionCloudPcPostRequestBody instantiates a new MeManagedDevicesBulkReprovisionCloudPcPostRequestBody and sets the default values.
func NewMeManagedDevicesBulkReprovisionCloudPcPostRequestBody()(*MeManagedDevicesBulkReprovisionCloudPcPostRequestBody) {
    m := &MeManagedDevicesBulkReprovisionCloudPcPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeManagedDevicesBulkReprovisionCloudPcPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeManagedDevicesBulkReprovisionCloudPcPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeManagedDevicesBulkReprovisionCloudPcPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeManagedDevicesBulkReprovisionCloudPcPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeManagedDevicesBulkReprovisionCloudPcPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managedDeviceIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetManagedDeviceIds(res)
        }
        return nil
    }
    return res
}
// GetManagedDeviceIds gets the managedDeviceIds property value. The managedDeviceIds property
func (m *MeManagedDevicesBulkReprovisionCloudPcPostRequestBody) GetManagedDeviceIds()([]string) {
    return m.managedDeviceIds
}
// Serialize serializes information the current object
func (m *MeManagedDevicesBulkReprovisionCloudPcPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetManagedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("managedDeviceIds", m.GetManagedDeviceIds())
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
func (m *MeManagedDevicesBulkReprovisionCloudPcPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetManagedDeviceIds sets the managedDeviceIds property value. The managedDeviceIds property
func (m *MeManagedDevicesBulkReprovisionCloudPcPostRequestBody) SetManagedDeviceIds(value []string)() {
    m.managedDeviceIds = value
}
