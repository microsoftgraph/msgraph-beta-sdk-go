package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppTroubleshootingEvent 
type MobileAppTroubleshootingEvent struct {
    DeviceManagementTroubleshootingEvent
}
// NewMobileAppTroubleshootingEvent instantiates a new MobileAppTroubleshootingEvent and sets the default values.
func NewMobileAppTroubleshootingEvent()(*MobileAppTroubleshootingEvent) {
    m := &MobileAppTroubleshootingEvent{
        DeviceManagementTroubleshootingEvent: *NewDeviceManagementTroubleshootingEvent(),
    }
    return m
}
// CreateMobileAppTroubleshootingEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppTroubleshootingEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppTroubleshootingEvent(), nil
}
// GetApplicationId gets the applicationId property value. Intune application identifier.
func (m *MobileAppTroubleshootingEvent) GetApplicationId()(*string) {
    val, err := m.GetBackingStore().Get("applicationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppLogCollectionRequests gets the appLogCollectionRequests property value. The collection property of AppLogUploadRequest.
func (m *MobileAppTroubleshootingEvent) GetAppLogCollectionRequests()([]AppLogCollectionRequestable) {
    val, err := m.GetBackingStore().Get("appLogCollectionRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppLogCollectionRequestable)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. Device identifier created or collected by Intune.
func (m *MobileAppTroubleshootingEvent) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppTroubleshootingEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementTroubleshootingEvent.GetFieldDeserializers()
    res["applicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
        }
        return nil
    }
    res["appLogCollectionRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppLogCollectionRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppLogCollectionRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppLogCollectionRequestable)
                }
            }
            m.SetAppLogCollectionRequests(res)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["history"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppTroubleshootingHistoryItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppTroubleshootingHistoryItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MobileAppTroubleshootingHistoryItemable)
                }
            }
            m.SetHistory(res)
        }
        return nil
    }
    res["managedDeviceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceIdentifier(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetHistory gets the history property value. Intune Mobile Application Troubleshooting History Item
func (m *MobileAppTroubleshootingEvent) GetHistory()([]MobileAppTroubleshootingHistoryItemable) {
    val, err := m.GetBackingStore().Get("history")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MobileAppTroubleshootingHistoryItemable)
    }
    return nil
}
// GetManagedDeviceIdentifier gets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppTroubleshootingEvent) GetManagedDeviceIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppTroubleshootingEvent) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobileAppTroubleshootingEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementTroubleshootingEvent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    if m.GetAppLogCollectionRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppLogCollectionRequests()))
        for i, v := range m.GetAppLogCollectionRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appLogCollectionRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    if m.GetHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHistory()))
        for i, v := range m.GetHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("history", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceIdentifier", m.GetManagedDeviceIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationId sets the applicationId property value. Intune application identifier.
func (m *MobileAppTroubleshootingEvent) SetApplicationId(value *string)() {
    err := m.GetBackingStore().Set("applicationId", value)
    if err != nil {
        panic(err)
    }
}
// SetAppLogCollectionRequests sets the appLogCollectionRequests property value. The collection property of AppLogUploadRequest.
func (m *MobileAppTroubleshootingEvent) SetAppLogCollectionRequests(value []AppLogCollectionRequestable)() {
    err := m.GetBackingStore().Set("appLogCollectionRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. Device identifier created or collected by Intune.
func (m *MobileAppTroubleshootingEvent) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetHistory sets the history property value. Intune Mobile Application Troubleshooting History Item
func (m *MobileAppTroubleshootingEvent) SetHistory(value []MobileAppTroubleshootingHistoryItemable)() {
    err := m.GetBackingStore().Set("history", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceIdentifier sets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppTroubleshootingEvent) SetManagedDeviceIdentifier(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppTroubleshootingEvent) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// MobileAppTroubleshootingEventable 
type MobileAppTroubleshootingEventable interface {
    DeviceManagementTroubleshootingEventable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationId()(*string)
    GetAppLogCollectionRequests()([]AppLogCollectionRequestable)
    GetDeviceId()(*string)
    GetHistory()([]MobileAppTroubleshootingHistoryItemable)
    GetManagedDeviceIdentifier()(*string)
    GetUserId()(*string)
    SetApplicationId(value *string)()
    SetAppLogCollectionRequests(value []AppLogCollectionRequestable)()
    SetDeviceId(value *string)()
    SetHistory(value []MobileAppTroubleshootingHistoryItemable)()
    SetManagedDeviceIdentifier(value *string)()
    SetUserId(value *string)()
}
