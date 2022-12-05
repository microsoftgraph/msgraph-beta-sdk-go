package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody provides operations to call the uploadDepToken method.
type DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appleId property
    appleId *string
    // The depToken property
    depToken *string
}
// NewDeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody instantiates a new DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody and sets the default values.
func NewDeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody()(*DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody) {
    m := &DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppleId gets the appleId property value. The appleId property
func (m *DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody) GetAppleId()(*string) {
    return m.appleId
}
// GetDepToken gets the depToken property value. The depToken property
func (m *DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody) GetDepToken()(*string) {
    return m.depToken
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleId(val)
        }
        return nil
    }
    res["depToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepToken(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appleId", m.GetAppleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("depToken", m.GetDepToken())
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
func (m *DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppleId sets the appleId property value. The appleId property
func (m *DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody) SetAppleId(value *string)() {
    m.appleId = value
}
// SetDepToken sets the depToken property value. The depToken property
func (m *DeviceManagementDepOnboardingSettingsItemUploadDepTokenPostRequestBody) SetDepToken(value *string)() {
    m.depToken = value
}
