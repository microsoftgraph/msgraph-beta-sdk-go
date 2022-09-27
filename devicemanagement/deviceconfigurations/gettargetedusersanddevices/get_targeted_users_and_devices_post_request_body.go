package gettargetedusersanddevices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetTargetedUsersAndDevicesPostRequestBody provides operations to call the getTargetedUsersAndDevices method.
type GetTargetedUsersAndDevicesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceConfigurationIds property
    deviceConfigurationIds []string
}
// NewGetTargetedUsersAndDevicesPostRequestBody instantiates a new getTargetedUsersAndDevicesPostRequestBody and sets the default values.
func NewGetTargetedUsersAndDevicesPostRequestBody()(*GetTargetedUsersAndDevicesPostRequestBody) {
    m := &GetTargetedUsersAndDevicesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetTargetedUsersAndDevicesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetTargetedUsersAndDevicesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetTargetedUsersAndDevicesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetTargetedUsersAndDevicesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceConfigurationIds gets the deviceConfigurationIds property value. The deviceConfigurationIds property
func (m *GetTargetedUsersAndDevicesPostRequestBody) GetDeviceConfigurationIds()([]string) {
    return m.deviceConfigurationIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetTargetedUsersAndDevicesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceConfigurationIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDeviceConfigurationIds)
    return res
}
// Serialize serializes information the current object
func (m *GetTargetedUsersAndDevicesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceConfigurationIds() != nil {
        err := writer.WriteCollectionOfStringValues("deviceConfigurationIds", m.GetDeviceConfigurationIds())
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
func (m *GetTargetedUsersAndDevicesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceConfigurationIds sets the deviceConfigurationIds property value. The deviceConfigurationIds property
func (m *GetTargetedUsersAndDevicesPostRequestBody) SetDeviceConfigurationIds(value []string)() {
    m.deviceConfigurationIds = value
}
