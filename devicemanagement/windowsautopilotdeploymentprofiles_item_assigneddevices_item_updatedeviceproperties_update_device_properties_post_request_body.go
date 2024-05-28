package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody instantiates a new WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody and sets the default values.
func NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody()(*WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) {
    m := &WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) GetAddressableUserName()(*string) {
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
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceAccountPassword gets the deviceAccountPassword property value. The deviceAccountPassword property
// returns a *string when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) GetDeviceAccountPassword()(*string) {
    val, err := m.GetBackingStore().Get("deviceAccountPassword")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceAccountUpn gets the deviceAccountUpn property value. The deviceAccountUpn property
// returns a *string when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) GetDeviceAccountUpn()(*string) {
    val, err := m.GetBackingStore().Get("deviceAccountUpn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceFriendlyName gets the deviceFriendlyName property value. The deviceFriendlyName property
// returns a *string when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) GetDeviceFriendlyName()(*string) {
    val, err := m.GetBackingStore().Get("deviceFriendlyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["deviceAccountPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAccountPassword(val)
        }
        return nil
    }
    res["deviceAccountUpn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAccountUpn(val)
        }
        return nil
    }
    res["deviceFriendlyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFriendlyName(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["groupTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupTag(val)
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
// GetGroupTag gets the groupTag property value. The groupTag property
// returns a *string when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) GetGroupTag()(*string) {
    val, err := m.GetBackingStore().Get("groupTag")
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
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) GetUserPrincipalName()(*string) {
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
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("addressableUserName", m.GetAddressableUserName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceAccountPassword", m.GetDeviceAccountPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceAccountUpn", m.GetDeviceAccountUpn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceFriendlyName", m.GetDeviceFriendlyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("groupTag", m.GetGroupTag())
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
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAddressableUserName sets the addressableUserName property value. The addressableUserName property
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) SetAddressableUserName(value *string)() {
    err := m.GetBackingStore().Set("addressableUserName", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceAccountPassword sets the deviceAccountPassword property value. The deviceAccountPassword property
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) SetDeviceAccountPassword(value *string)() {
    err := m.GetBackingStore().Set("deviceAccountPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceAccountUpn sets the deviceAccountUpn property value. The deviceAccountUpn property
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) SetDeviceAccountUpn(value *string)() {
    err := m.GetBackingStore().Set("deviceAccountUpn", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceFriendlyName sets the deviceFriendlyName property value. The deviceFriendlyName property
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) SetDeviceFriendlyName(value *string)() {
    err := m.GetBackingStore().Set("deviceFriendlyName", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupTag sets the groupTag property value. The groupTag property
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) SetGroupTag(value *string)() {
    err := m.GetBackingStore().Set("groupTag", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBody) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
type WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddressableUserName()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceAccountPassword()(*string)
    GetDeviceAccountUpn()(*string)
    GetDeviceFriendlyName()(*string)
    GetDisplayName()(*string)
    GetGroupTag()(*string)
    GetUserPrincipalName()(*string)
    SetAddressableUserName(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceAccountPassword(value *string)()
    SetDeviceAccountUpn(value *string)()
    SetDeviceFriendlyName(value *string)()
    SetDisplayName(value *string)()
    SetGroupTag(value *string)()
    SetUserPrincipalName(value *string)()
}
