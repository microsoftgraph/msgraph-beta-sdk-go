package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody 
type ComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The managedDeviceIds property
    managedDeviceIds []string
}
// NewComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody instantiates a new ComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody and sets the default values.
func NewComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody()(*ComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody) {
    m := &ComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody) GetManagedDeviceIds()([]string) {
    return m.managedDeviceIds
}
// Serialize serializes information the current object
func (m *ComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetManagedDeviceIds sets the managedDeviceIds property value. The managedDeviceIds property
func (m *ComanagedDevicesMicrosoftGraphBulkReprovisionCloudPcBulkReprovisionCloudPcPostRequestBody) SetManagedDeviceIds(value []string)() {
    m.managedDeviceIds = value
}
