package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OfficeClientCheckinStatus struct {
    additionalData map[string]interface{};
    appliedPolicies []string;
    checkinDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    deviceName *string;
    devicePlatform *string;
    devicePlatformVersion *string;
    errorMessage *string;
    userId *string;
    userPrincipalName *string;
    wasSuccessful *bool;
}
func NewOfficeClientCheckinStatus()(*OfficeClientCheckinStatus) {
    m := &OfficeClientCheckinStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *OfficeClientCheckinStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *OfficeClientCheckinStatus) GetAppliedPolicies()([]string) {
    if m == nil {
        return nil
    } else {
        return m.appliedPolicies
    }
}
func (m *OfficeClientCheckinStatus) GetCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.checkinDateTime
    }
}
func (m *OfficeClientCheckinStatus) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *OfficeClientCheckinStatus) GetDevicePlatform()(*string) {
    if m == nil {
        return nil
    } else {
        return m.devicePlatform
    }
}
func (m *OfficeClientCheckinStatus) GetDevicePlatformVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.devicePlatformVersion
    }
}
func (m *OfficeClientCheckinStatus) GetErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorMessage
    }
}
func (m *OfficeClientCheckinStatus) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *OfficeClientCheckinStatus) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *OfficeClientCheckinStatus) GetWasSuccessful()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.wasSuccessful
    }
}
func (m *OfficeClientCheckinStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appliedPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAppliedPolicies(res)
        return nil
    }
    res["checkinDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCheckinDateTime(val)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["devicePlatform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDevicePlatform(val)
        return nil
    }
    res["devicePlatformVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDevicePlatformVersion(val)
        return nil
    }
    res["errorMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorMessage(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    res["wasSuccessful"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetWasSuccessful(val)
        return nil
    }
    return res
}
func (m *OfficeClientCheckinStatus) IsNil()(bool) {
    return m == nil
}
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
func (m *OfficeClientCheckinStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OfficeClientCheckinStatus) SetAppliedPolicies(value []string)() {
    m.appliedPolicies = value
}
func (m *OfficeClientCheckinStatus) SetCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.checkinDateTime = value
}
func (m *OfficeClientCheckinStatus) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *OfficeClientCheckinStatus) SetDevicePlatform(value *string)() {
    m.devicePlatform = value
}
func (m *OfficeClientCheckinStatus) SetDevicePlatformVersion(value *string)() {
    m.devicePlatformVersion = value
}
func (m *OfficeClientCheckinStatus) SetErrorMessage(value *string)() {
    m.errorMessage = value
}
func (m *OfficeClientCheckinStatus) SetUserId(value *string)() {
    m.userId = value
}
func (m *OfficeClientCheckinStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
func (m *OfficeClientCheckinStatus) SetWasSuccessful(value *bool)() {
    m.wasSuccessful = value
}
