package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppTroubleshootingEvent provides operations to manage the deviceManagement singleton.
type MobileAppTroubleshootingEvent struct {
    DeviceManagementTroubleshootingEvent
    // Intune application identifier.
    applicationId *string;
    // The collection property of AppLogUploadRequest.
    appLogCollectionRequests []AppLogCollectionRequestable;
    // Intune Mobile Application Troubleshooting History Item
    history []MobileAppTroubleshootingHistoryItemable;
    // Device identifier created or collected by Intune.
    managedDeviceIdentifier *string;
    // Identifier for the user that tried to enroll the device.
    userId *string;
}
// NewMobileAppTroubleshootingEvent instantiates a new mobileAppTroubleshootingEvent and sets the default values.
func NewMobileAppTroubleshootingEvent()(*MobileAppTroubleshootingEvent) {
    m := &MobileAppTroubleshootingEvent{
        DeviceManagementTroubleshootingEvent: *NewDeviceManagementTroubleshootingEvent(),
    }
    return m
}
// CreateMobileAppTroubleshootingEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppTroubleshootingEventFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMobileAppTroubleshootingEvent(), nil
}
// GetApplicationId gets the applicationId property value. Intune application identifier.
func (m *MobileAppTroubleshootingEvent) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
// GetAppLogCollectionRequests gets the appLogCollectionRequests property value. The collection property of AppLogUploadRequest.
func (m *MobileAppTroubleshootingEvent) GetAppLogCollectionRequests()([]AppLogCollectionRequestable) {
    if m == nil {
        return nil
    } else {
        return m.appLogCollectionRequests
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
        val, err := n.GetCollectionOfObjectValues(CreateAppLogCollectionRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppLogCollectionRequestable, len(val))
            for i, v := range val {
                res[i] = v.(AppLogCollectionRequestable)
            }
            m.SetAppLogCollectionRequests(res)
        }
        return nil
    }
    res["history"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppTroubleshootingHistoryItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppTroubleshootingHistoryItemable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppTroubleshootingHistoryItemable)
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
// GetHistory gets the history property value. Intune Mobile Application Troubleshooting History Item
func (m *MobileAppTroubleshootingEvent) GetHistory()([]MobileAppTroubleshootingHistoryItemable) {
    if m == nil {
        return nil
    } else {
        return m.history
    }
}
// GetManagedDeviceIdentifier gets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppTroubleshootingEvent) GetManagedDeviceIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceIdentifier
    }
}
// GetUserId gets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppTroubleshootingEvent) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *MobileAppTroubleshootingEvent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetAppLogCollectionRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppLogCollectionRequests()))
        for i, v := range m.GetAppLogCollectionRequests() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appLogCollectionRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHistory() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHistory()))
        for i, v := range m.GetHistory() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m != nil {
        m.applicationId = value
    }
}
// SetAppLogCollectionRequests sets the appLogCollectionRequests property value. The collection property of AppLogUploadRequest.
func (m *MobileAppTroubleshootingEvent) SetAppLogCollectionRequests(value []AppLogCollectionRequestable)() {
    if m != nil {
        m.appLogCollectionRequests = value
    }
}
// SetHistory sets the history property value. Intune Mobile Application Troubleshooting History Item
func (m *MobileAppTroubleshootingEvent) SetHistory(value []MobileAppTroubleshootingHistoryItemable)() {
    if m != nil {
        m.history = value
    }
}
// SetManagedDeviceIdentifier sets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppTroubleshootingEvent) SetManagedDeviceIdentifier(value *string)() {
    if m != nil {
        m.managedDeviceIdentifier = value
    }
}
// SetUserId sets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppTroubleshootingEvent) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
