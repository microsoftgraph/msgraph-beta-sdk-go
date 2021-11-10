package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MobileAppTroubleshootingEvent struct {
    DeviceManagementTroubleshootingEvent
    // Intune application identifier.
    applicationId *string;
    // The collection property of AppLogUploadRequest.
    appLogCollectionRequests []AppLogCollectionRequest;
    // Intune Mobile Application Troubleshooting History Item
    history []MobileAppTroubleshootingHistoryItem;
    // Device identifier created or collected by Intune.
    managedDeviceIdentifier *string;
    // Identifier for the user that tried to enroll the device.
    userId *string;
}
// Instantiates a new mobileAppTroubleshootingEvent and sets the default values.
func NewMobileAppTroubleshootingEvent()(*MobileAppTroubleshootingEvent) {
    m := &MobileAppTroubleshootingEvent{
        DeviceManagementTroubleshootingEvent: *NewDeviceManagementTroubleshootingEvent(),
    }
    return m
}
// Gets the applicationId property value. Intune application identifier.
func (m *MobileAppTroubleshootingEvent) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
// Gets the appLogCollectionRequests property value. The collection property of AppLogUploadRequest.
func (m *MobileAppTroubleshootingEvent) GetAppLogCollectionRequests()([]AppLogCollectionRequest) {
    if m == nil {
        return nil
    } else {
        return m.appLogCollectionRequests
    }
}
// Gets the history property value. Intune Mobile Application Troubleshooting History Item
func (m *MobileAppTroubleshootingEvent) GetHistory()([]MobileAppTroubleshootingHistoryItem) {
    if m == nil {
        return nil
    } else {
        return m.history
    }
}
// Gets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppTroubleshootingEvent) GetManagedDeviceIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceIdentifier
    }
}
// Gets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppTroubleshootingEvent) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *MobileAppTroubleshootingEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DeviceManagementTroubleshootingEvent.GetFieldDeserializers()
    res["applicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
        }
        return nil
    }
    res["appLogCollectionRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppLogCollectionRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppLogCollectionRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppLogCollectionRequest))
            }
            m.SetAppLogCollectionRequests(res)
        }
        return nil
    }
    res["history"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppTroubleshootingHistoryItem() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppTroubleshootingHistoryItem, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileAppTroubleshootingHistoryItem))
            }
            m.SetHistory(res)
        }
        return nil
    }
    res["managedDeviceIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceIdentifier(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *MobileAppTroubleshootingEvent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MobileAppTroubleshootingEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppLogCollectionRequests()))
        for i, v := range m.GetAppLogCollectionRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appLogCollectionRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHistory()))
        for i, v := range m.GetHistory() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// Sets the applicationId property value. Intune application identifier.
// Parameters:
//  - value : Value to set for the applicationId property.
func (m *MobileAppTroubleshootingEvent) SetApplicationId(value *string)() {
    m.applicationId = value
}
// Sets the appLogCollectionRequests property value. The collection property of AppLogUploadRequest.
// Parameters:
//  - value : Value to set for the appLogCollectionRequests property.
func (m *MobileAppTroubleshootingEvent) SetAppLogCollectionRequests(value []AppLogCollectionRequest)() {
    m.appLogCollectionRequests = value
}
// Sets the history property value. Intune Mobile Application Troubleshooting History Item
// Parameters:
//  - value : Value to set for the history property.
func (m *MobileAppTroubleshootingEvent) SetHistory(value []MobileAppTroubleshootingHistoryItem)() {
    m.history = value
}
// Sets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
// Parameters:
//  - value : Value to set for the managedDeviceIdentifier property.
func (m *MobileAppTroubleshootingEvent) SetManagedDeviceIdentifier(value *string)() {
    m.managedDeviceIdentifier = value
}
// Sets the userId property value. Identifier for the user that tried to enroll the device.
// Parameters:
//  - value : Value to set for the userId property.
func (m *MobileAppTroubleshootingEvent) SetUserId(value *string)() {
    m.userId = value
}
