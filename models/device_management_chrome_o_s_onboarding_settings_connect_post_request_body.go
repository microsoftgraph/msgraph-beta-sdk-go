package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody provides operations to call the connect method.
type DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ownerAccessToken property
    ownerAccessToken *string
    // The ownerUserPrincipalName property
    ownerUserPrincipalName *string
}
// NewDeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody instantiates a new DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody and sets the default values.
func NewDeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody()(*DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody) {
    m := &DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementChromeOSOnboardingSettingsConnectPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementChromeOSOnboardingSettingsConnectPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ownerAccessToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerAccessToken(val)
        }
        return nil
    }
    res["ownerUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetOwnerAccessToken gets the ownerAccessToken property value. The ownerAccessToken property
func (m *DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody) GetOwnerAccessToken()(*string) {
    return m.ownerAccessToken
}
// GetOwnerUserPrincipalName gets the ownerUserPrincipalName property value. The ownerUserPrincipalName property
func (m *DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody) GetOwnerUserPrincipalName()(*string) {
    return m.ownerUserPrincipalName
}
// Serialize serializes information the current object
func (m *DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ownerAccessToken", m.GetOwnerAccessToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ownerUserPrincipalName", m.GetOwnerUserPrincipalName())
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
func (m *DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOwnerAccessToken sets the ownerAccessToken property value. The ownerAccessToken property
func (m *DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody) SetOwnerAccessToken(value *string)() {
    m.ownerAccessToken = value
}
// SetOwnerUserPrincipalName sets the ownerUserPrincipalName property value. The ownerUserPrincipalName property
func (m *DeviceManagementChromeOSOnboardingSettingsConnectPostRequestBody) SetOwnerUserPrincipalName(value *string)() {
    m.ownerUserPrincipalName = value
}
