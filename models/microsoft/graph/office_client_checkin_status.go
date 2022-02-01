package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OfficeClientCheckinStatus 
type OfficeClientCheckinStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of policies delivered to the device as last checkin.
    appliedPolicies []string;
    // Last device check-in time in UTC.
    checkinDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Device name trying to check-in.
    deviceName *string;
    // Device platform trying to check-in.
    devicePlatform *string;
    // Device platform version trying to check-in.
    devicePlatformVersion *string;
    // Error message if any associated for the last checkin.
    errorMessage *string;
    // User identifier using the device.
    userId *string;
    // User principal name using the device.
    userPrincipalName *string;
    // If the last checkin was successful.
    wasSuccessful *bool;
}
// NewOfficeClientCheckinStatus instantiates a new officeClientCheckinStatus and sets the default values.
func NewOfficeClientCheckinStatus()(*OfficeClientCheckinStatus) {
    m := &OfficeClientCheckinStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeClientCheckinStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppliedPolicies gets the appliedPolicies property value. List of policies delivered to the device as last checkin.
func (m *OfficeClientCheckinStatus) GetAppliedPolicies()([]string) {
    if m == nil {
        return nil
    } else {
        return m.appliedPolicies
    }
}
// GetCheckinDateTime gets the checkinDateTime property value. Last device check-in time in UTC.
func (m *OfficeClientCheckinStatus) GetCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.checkinDateTime
    }
}
// GetDeviceName gets the deviceName property value. Device name trying to check-in.
func (m *OfficeClientCheckinStatus) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetDevicePlatform gets the devicePlatform property value. Device platform trying to check-in.
func (m *OfficeClientCheckinStatus) GetDevicePlatform()(*string) {
    if m == nil {
        return nil
    } else {
        return m.devicePlatform
    }
}
// GetDevicePlatformVersion gets the devicePlatformVersion property value. Device platform version trying to check-in.
func (m *OfficeClientCheckinStatus) GetDevicePlatformVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.devicePlatformVersion
    }
}
// GetErrorMessage gets the errorMessage property value. Error message if any associated for the last checkin.
func (m *OfficeClientCheckinStatus) GetErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorMessage
    }
}
// GetUserId gets the userId property value. User identifier using the device.
func (m *OfficeClientCheckinStatus) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User principal name using the device.
func (m *OfficeClientCheckinStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetWasSuccessful gets the wasSuccessful property value. If the last checkin was successful.
func (m *OfficeClientCheckinStatus) GetWasSuccessful()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.wasSuccessful
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeClientCheckinStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appliedPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAppliedPolicies(res)
        }
        return nil
    }
    res["checkinDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckinDateTime(val)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["devicePlatform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePlatform(val)
        }
        return nil
    }
    res["devicePlatformVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePlatformVersion(val)
        }
        return nil
    }
    res["errorMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorMessage(val)
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
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["wasSuccessful"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWasSuccessful(val)
        }
        return nil
    }
    return res
}
func (m *OfficeClientCheckinStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OfficeClientCheckinStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAppliedPolicies() != nil {
        err := writer.WriteCollectionOfStringValues("appliedPolicies", m.GetAppliedPolicies())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("checkinDateTime", m.GetCheckinDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("devicePlatform", m.GetDevicePlatform())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("devicePlatformVersion", m.GetDevicePlatformVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("errorMessage", m.GetErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
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
        err := writer.WriteBoolValue("wasSuccessful", m.GetWasSuccessful())
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
func (m *OfficeClientCheckinStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppliedPolicies sets the appliedPolicies property value. List of policies delivered to the device as last checkin.
func (m *OfficeClientCheckinStatus) SetAppliedPolicies(value []string)() {
    if m != nil {
        m.appliedPolicies = value
    }
}
// SetCheckinDateTime sets the checkinDateTime property value. Last device check-in time in UTC.
func (m *OfficeClientCheckinStatus) SetCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.checkinDateTime = value
    }
}
// SetDeviceName sets the deviceName property value. Device name trying to check-in.
func (m *OfficeClientCheckinStatus) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetDevicePlatform sets the devicePlatform property value. Device platform trying to check-in.
func (m *OfficeClientCheckinStatus) SetDevicePlatform(value *string)() {
    if m != nil {
        m.devicePlatform = value
    }
}
// SetDevicePlatformVersion sets the devicePlatformVersion property value. Device platform version trying to check-in.
func (m *OfficeClientCheckinStatus) SetDevicePlatformVersion(value *string)() {
    if m != nil {
        m.devicePlatformVersion = value
    }
}
// SetErrorMessage sets the errorMessage property value. Error message if any associated for the last checkin.
func (m *OfficeClientCheckinStatus) SetErrorMessage(value *string)() {
    if m != nil {
        m.errorMessage = value
    }
}
// SetUserId sets the userId property value. User identifier using the device.
func (m *OfficeClientCheckinStatus) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name using the device.
func (m *OfficeClientCheckinStatus) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetWasSuccessful sets the wasSuccessful property value. If the last checkin was successful.
func (m *OfficeClientCheckinStatus) SetWasSuccessful(value *bool)() {
    if m != nil {
        m.wasSuccessful = value
    }
}
