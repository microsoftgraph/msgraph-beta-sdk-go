package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody provides operations to call the createDeviceLogCollectionRequest method.
type UsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The templateType property
    templateType DeviceLogCollectionRequestable
}
// NewUsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody instantiates a new UsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody and sets the default values.
func NewUsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody()(*UsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody) {
    m := &UsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["templateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceLogCollectionRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateType(val.(DeviceLogCollectionRequestable))
        }
        return nil
    }
    return res
}
// GetTemplateType gets the templateType property value. The templateType property
func (m *UsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody) GetTemplateType()(DeviceLogCollectionRequestable) {
    return m.templateType
}
// Serialize serializes information the current object
func (m *UsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("templateType", m.GetTemplateType())
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
func (m *UsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetTemplateType sets the templateType property value. The templateType property
func (m *UsersItemManagedDevicesItemCreateDeviceLogCollectionRequestPostRequestBody) SetTemplateType(value DeviceLogCollectionRequestable)() {
    m.templateType = value
}
