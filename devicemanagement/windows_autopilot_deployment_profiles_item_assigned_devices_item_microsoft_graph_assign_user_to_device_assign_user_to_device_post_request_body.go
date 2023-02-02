package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody 
type WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The addressableUserName property
    addressableUserName *string
    // The userPrincipalName property
    userPrincipalName *string
}
// NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody instantiates a new WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody and sets the default values.
func NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody()(*WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) {
    m := &WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddressableUserName gets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) GetAddressableUserName()(*string) {
    return m.addressableUserName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addressableUserName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressableUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("addressableUserName", m.GetAddressableUserName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddressableUserName sets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) SetAddressableUserName(value *string)() {
    m.addressableUserName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
