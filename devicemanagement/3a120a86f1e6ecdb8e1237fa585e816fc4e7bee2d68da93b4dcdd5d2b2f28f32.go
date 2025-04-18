// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody instantiates a new WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody and sets the default values.
func NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody()(*WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) {
    m := &WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAddressableUserName gets the addressableUserName property value. The addressableUserName property
// returns a *string when successful
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) GetAddressableUserName()(*string) {
    val, err := m.GetBackingStore().Get("addressableUserName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["resourceAccountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAccountName(val)
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
// GetResourceAccountName gets the resourceAccountName property value. The resourceAccountName property
// returns a *string when successful
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) GetResourceAccountName()(*string) {
    val, err := m.GetBackingStore().Get("resourceAccountName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
// returns a *string when successful
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("addressableUserName", m.GetAddressableUserName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceAccountName", m.GetResourceAccountName())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAddressableUserName sets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) SetAddressableUserName(value *string)() {
    err := m.GetBackingStore().Set("addressableUserName", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetResourceAccountName sets the resourceAccountName property value. The resourceAccountName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) SetResourceAccountName(value *string)() {
    err := m.GetBackingStore().Set("resourceAccountName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBody) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
type WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddressableUserName()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetResourceAccountName()(*string)
    GetUserPrincipalName()(*string)
    SetAddressableUserName(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetResourceAccountName(value *string)()
    SetUserPrincipalName(value *string)()
}
