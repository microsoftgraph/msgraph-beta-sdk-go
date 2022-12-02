package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody provides operations to call the completeSignup method.
type DeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enterpriseToken property
    enterpriseToken *string
}
// NewDeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody instantiates a new DeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody and sets the default values.
func NewDeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody()(*DeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody) {
    m := &DeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnterpriseToken gets the enterpriseToken property value. The enterpriseToken property
func (m *DeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody) GetEnterpriseToken()(*string) {
    return m.enterpriseToken
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enterpriseToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseToken(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("enterpriseToken", m.GetEnterpriseToken())
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
func (m *DeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnterpriseToken sets the enterpriseToken property value. The enterpriseToken property
func (m *DeviceManagementAndroidForWorkSettingsCompleteSignupPostRequestBody) SetEnterpriseToken(value *string)() {
    m.enterpriseToken = value
}
