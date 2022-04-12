package gettargetedusersanddevices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetTargetedUsersAndDevicesRequestBody provides operations to call the getTargetedUsersAndDevices method.
type GetTargetedUsersAndDevicesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceConfigurationIds property
    deviceConfigurationIds []string
}
// NewGetTargetedUsersAndDevicesRequestBody instantiates a new getTargetedUsersAndDevicesRequestBody and sets the default values.
func NewGetTargetedUsersAndDevicesRequestBody()(*GetTargetedUsersAndDevicesRequestBody) {
    m := &GetTargetedUsersAndDevicesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetTargetedUsersAndDevicesRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetTargetedUsersAndDevicesRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetTargetedUsersAndDevicesRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetTargetedUsersAndDevicesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceConfigurationIds gets the deviceConfigurationIds property value. The deviceConfigurationIds property
func (m *GetTargetedUsersAndDevicesRequestBody) GetDeviceConfigurationIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetTargetedUsersAndDevicesRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceConfigurationIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDeviceConfigurationIds(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GetTargetedUsersAndDevicesRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GetTargetedUsersAndDevicesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceConfigurationIds sets the deviceConfigurationIds property value. The deviceConfigurationIds property
func (m *GetTargetedUsersAndDevicesRequestBody) SetDeviceConfigurationIds(value []string)() {
    if m != nil {
        m.deviceConfigurationIds = value
    }
}
