package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody provides operations to call the windowsPrivacyAccessControls method.
type DeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The windowsPrivacyAccessControls property
    windowsPrivacyAccessControls []WindowsPrivacyDataAccessControlItemable
}
// NewDeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody instantiates a new DeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody and sets the default values.
func NewDeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody()(*DeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) {
    m := &DeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["windowsPrivacyAccessControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsPrivacyDataAccessControlItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsPrivacyDataAccessControlItemable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsPrivacyDataAccessControlItemable)
            }
            m.SetWindowsPrivacyAccessControls(res)
        }
        return nil
    }
    return res
}
// GetWindowsPrivacyAccessControls gets the windowsPrivacyAccessControls property value. The windowsPrivacyAccessControls property
func (m *DeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) GetWindowsPrivacyAccessControls()([]WindowsPrivacyDataAccessControlItemable) {
    return m.windowsPrivacyAccessControls
}
// Serialize serializes information the current object
func (m *DeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetWindowsPrivacyAccessControls() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsPrivacyAccessControls()))
        for i, v := range m.GetWindowsPrivacyAccessControls() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("windowsPrivacyAccessControls", cast)
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
func (m *DeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetWindowsPrivacyAccessControls sets the windowsPrivacyAccessControls property value. The windowsPrivacyAccessControls property
func (m *DeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBody) SetWindowsPrivacyAccessControls(value []WindowsPrivacyDataAccessControlItemable)() {
    m.windowsPrivacyAccessControls = value
}
