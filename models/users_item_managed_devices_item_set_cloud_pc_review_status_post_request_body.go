package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody provides operations to call the setCloudPcReviewStatus method.
type UsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The reviewStatus property
    reviewStatus CloudPcReviewStatusable
}
// NewUsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody instantiates a new UsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody and sets the default values.
func NewUsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody()(*UsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody) {
    m := &UsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reviewStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcReviewStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewStatus(val.(CloudPcReviewStatusable))
        }
        return nil
    }
    return res
}
// GetReviewStatus gets the reviewStatus property value. The reviewStatus property
func (m *UsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody) GetReviewStatus()(CloudPcReviewStatusable) {
    return m.reviewStatus
}
// Serialize serializes information the current object
func (m *UsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetReviewStatus sets the reviewStatus property value. The reviewStatus property
func (m *UsersItemManagedDevicesItemSetCloudPcReviewStatusPostRequestBody) SetReviewStatus(value CloudPcReviewStatusable)() {
    m.reviewStatus = value
}
