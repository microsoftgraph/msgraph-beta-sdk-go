package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody 
type ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The managedDeviceIds property
    managedDeviceIds []string
    // The reviewStatus property
    reviewStatus ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable
}
// NewComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody instantiates a new ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody and sets the default values.
func NewComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody()(*ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody) {
    m := &ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["reviewStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcReviewStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewStatus(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable))
        }
        return nil
    }
    return res
}
// GetManagedDeviceIds gets the managedDeviceIds property value. The managedDeviceIds property
func (m *ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody) GetManagedDeviceIds()([]string) {
    return m.managedDeviceIds
}
// GetReviewStatus gets the reviewStatus property value. The reviewStatus property
func (m *ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody) GetReviewStatus()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable) {
    return m.reviewStatus
}
// Serialize serializes information the current object
func (m *ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetManagedDeviceIds() != nil {
        err := writer.WriteCollectionOfStringValues("managedDeviceIds", m.GetManagedDeviceIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reviewStatus", m.GetReviewStatus())
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
func (m *ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetManagedDeviceIds sets the managedDeviceIds property value. The managedDeviceIds property
func (m *ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody) SetManagedDeviceIds(value []string)() {
    m.managedDeviceIds = value
}
// SetReviewStatus sets the reviewStatus property value. The reviewStatus property
func (m *ComanagedDevicesMicrosoftGraphBulkSetCloudPcReviewStatusBulkSetCloudPcReviewStatusPostRequestBody) SetReviewStatus(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcReviewStatusable)() {
    m.reviewStatus = value
}
