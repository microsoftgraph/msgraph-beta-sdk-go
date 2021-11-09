package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new officeClientCheckinStatus and sets the default values.
func NewOfficeClientCheckinStatus()(*OfficeClientCheckinStatus) {
    m := &OfficeClientCheckinStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeClientCheckinStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the appliedPolicies property value. List of policies delivered to the device as last checkin.
func (m *OfficeClientCheckinStatus) GetAppliedPolicies()([]string) {
    if m == nil {
        return nil
    } else {
        return m.appliedPolicies
    }
}
// Gets the checkinDateTime property value. Last device check-in time in UTC.
func (m *OfficeClientCheckinStatus) GetCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.checkinDateTime
    }
}
// Gets the deviceName property value. Device name trying to check-in.
func (m *OfficeClientCheckinStatus) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the devicePlatform property value. Device platform trying to check-in.
func (m *OfficeClientCheckinStatus) GetDevicePlatform()(*string) {
    if m == nil {
        return nil
    } else {
        return m.devicePlatform
    }
}
// Gets the devicePlatformVersion property value. Device platform version trying to check-in.
func (m *OfficeClientCheckinStatus) GetDevicePlatformVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.devicePlatformVersion
    }
}
// Gets the errorMessage property value. Error message if any associated for the last checkin.
func (m *OfficeClientCheckinStatus) GetErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorMessage
    }
}
// Gets the userId property value. User identifier using the device.
func (m *OfficeClientCheckinStatus) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the userPrincipalName property value. User principal name using the device.
func (m *OfficeClientCheckinStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the wasSuccessful property value. If the last checkin was successful.
func (m *OfficeClientCheckinStatus) GetWasSuccessful()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.wasSuccessful
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OfficeClientCheckinStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *OfficeClientCheckinStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the appliedPolicies property value. List of policies delivered to the device as last checkin.
// Parameters:
//  - value : Value to set for the appliedPolicies property.
func (m *OfficeClientCheckinStatus) SetAppliedPolicies(value []string)() {
    m.appliedPolicies = value
}
// Sets the checkinDateTime property value. Last device check-in time in UTC.
// Parameters:
//  - value : Value to set for the checkinDateTime property.
func (m *OfficeClientCheckinStatus) SetCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.checkinDateTime = value
}
// Sets the deviceName property value. Device name trying to check-in.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *OfficeClientCheckinStatus) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the devicePlatform property value. Device platform trying to check-in.
// Parameters:
//  - value : Value to set for the devicePlatform property.
func (m *OfficeClientCheckinStatus) SetDevicePlatform(value *string)() {
    m.devicePlatform = value
}
// Sets the devicePlatformVersion property value. Device platform version trying to check-in.
// Parameters:
//  - value : Value to set for the devicePlatformVersion property.
func (m *OfficeClientCheckinStatus) SetDevicePlatformVersion(value *string)() {
    m.devicePlatformVersion = value
}
// Sets the errorMessage property value. Error message if any associated for the last checkin.
// Parameters:
//  - value : Value to set for the errorMessage property.
func (m *OfficeClientCheckinStatus) SetErrorMessage(value *string)() {
    m.errorMessage = value
}
// Sets the userId property value. User identifier using the device.
// Parameters:
//  - value : Value to set for the userId property.
func (m *OfficeClientCheckinStatus) SetUserId(value *string)() {
    m.userId = value
}
// Sets the userPrincipalName property value. User principal name using the device.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *OfficeClientCheckinStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the wasSuccessful property value. If the last checkin was successful.
// Parameters:
//  - value : Value to set for the wasSuccessful property.
func (m *OfficeClientCheckinStatus) SetWasSuccessful(value *bool)() {
    m.wasSuccessful = value
}
