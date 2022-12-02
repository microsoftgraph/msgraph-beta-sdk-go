package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody provides operations to call the completeSignup method.
type DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enterpriseToken property
    enterpriseToken *string
}
// NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody instantiates a new DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody and sets the default values.
func NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody()(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody) {
    m := &DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnterpriseToken gets the enterpriseToken property value. The enterpriseToken property
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody) GetEnterpriseToken()(*string) {
    return m.enterpriseToken
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnterpriseToken sets the enterpriseToken property value. The enterpriseToken property
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCompleteSignupPostRequestBody) SetEnterpriseToken(value *string)() {
    m.enterpriseToken = value
}
