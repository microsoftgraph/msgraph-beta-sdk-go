package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OutOfBoxExperienceSettings struct {
    additionalData map[string]interface{};
    deviceUsageType *WindowsDeviceUsageType;
    hideEscapeLink *bool;
    hideEULA *bool;
    hidePrivacySettings *bool;
    skipKeyboardSelectionPage *bool;
    userType *WindowsUserType;
}
func NewOutOfBoxExperienceSettings()(*OutOfBoxExperienceSettings) {
    m := &OutOfBoxExperienceSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *OutOfBoxExperienceSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *OutOfBoxExperienceSettings) GetDeviceUsageType()(*WindowsDeviceUsageType) {
    if m == nil {
        return nil
    } else {
        return m.deviceUsageType
    }
}
func (m *OutOfBoxExperienceSettings) GetHideEscapeLink()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideEscapeLink
    }
}
func (m *OutOfBoxExperienceSettings) GetHideEULA()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideEULA
    }
}
func (m *OutOfBoxExperienceSettings) GetHidePrivacySettings()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidePrivacySettings
    }
}
func (m *OutOfBoxExperienceSettings) GetSkipKeyboardSelectionPage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.skipKeyboardSelectionPage
    }
}
func (m *OutOfBoxExperienceSettings) GetUserType()(*WindowsUserType) {
    if m == nil {
        return nil
    } else {
        return m.userType
    }
}
func (m *OutOfBoxExperienceSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceUsageType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeviceUsageType)
        if err != nil {
            return err
        }
        cast := val.(WindowsDeviceUsageType)
        m.SetDeviceUsageType(&cast)
        return nil
    }
    res["hideEscapeLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHideEscapeLink(val)
        return nil
    }
    res["hideEULA"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHideEULA(val)
        return nil
    }
    res["hidePrivacySettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHidePrivacySettings(val)
        return nil
    }
    res["skipKeyboardSelectionPage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSkipKeyboardSelectionPage(val)
        return nil
    }
    res["userType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsUserType)
        if err != nil {
            return err
        }
        cast := val.(WindowsUserType)
        m.SetUserType(&cast)
        return nil
    }
    return res
}
func (m *OutOfBoxExperienceSettings) IsNil()(bool) {
    return m == nil
}
func (m *OutOfBoxExperienceSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDeviceUsageType() != nil {
        cast := m.GetDeviceUsageType().String()
        err := writer.WriteStringValue("deviceUsageType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideEscapeLink", m.GetHideEscapeLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideEULA", m.GetHideEULA())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hidePrivacySettings", m.GetHidePrivacySettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("skipKeyboardSelectionPage", m.GetSkipKeyboardSelectionPage())
        if err != nil {
            return err
        }
    }
    if m.GetUserType() != nil {
        cast := m.GetUserType().String()
        err := writer.WriteStringValue("userType", &cast)
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
func (m *OutOfBoxExperienceSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OutOfBoxExperienceSettings) SetDeviceUsageType(value *WindowsDeviceUsageType)() {
    m.deviceUsageType = value
}
func (m *OutOfBoxExperienceSettings) SetHideEscapeLink(value *bool)() {
    m.hideEscapeLink = value
}
func (m *OutOfBoxExperienceSettings) SetHideEULA(value *bool)() {
    m.hideEULA = value
}
func (m *OutOfBoxExperienceSettings) SetHidePrivacySettings(value *bool)() {
    m.hidePrivacySettings = value
}
func (m *OutOfBoxExperienceSettings) SetSkipKeyboardSelectionPage(value *bool)() {
    m.skipKeyboardSelectionPage = value
}
func (m *OutOfBoxExperienceSettings) SetUserType(value *WindowsUserType)() {
    m.userType = value
}
