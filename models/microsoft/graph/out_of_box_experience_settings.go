package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OutOfBoxExperienceSettings 
type OutOfBoxExperienceSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // AAD join authentication type. Possible values are: singleUser, shared.
    deviceUsageType *WindowsDeviceUsageType;
    // If set to true, then the user can't start over with different account, on company sign-in
    hideEscapeLink *bool;
    // Show or hide EULA to user
    hideEULA *bool;
    // Show or hide privacy settings to user
    hidePrivacySettings *bool;
    // If set, then skip the keyboard selection page if Language and Region are set
    skipKeyboardSelectionPage *bool;
    // Type of user. Possible values are: administrator, standard.
    userType *WindowsUserType;
}
// NewOutOfBoxExperienceSettings instantiates a new outOfBoxExperienceSettings and sets the default values.
func NewOutOfBoxExperienceSettings()(*OutOfBoxExperienceSettings) {
    m := &OutOfBoxExperienceSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OutOfBoxExperienceSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceUsageType gets the deviceUsageType property value. AAD join authentication type. Possible values are: singleUser, shared.
func (m *OutOfBoxExperienceSettings) GetDeviceUsageType()(*WindowsDeviceUsageType) {
    if m == nil {
        return nil
    } else {
        return m.deviceUsageType
    }
}
// GetHideEscapeLink gets the hideEscapeLink property value. If set to true, then the user can't start over with different account, on company sign-in
func (m *OutOfBoxExperienceSettings) GetHideEscapeLink()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideEscapeLink
    }
}
// GetHideEULA gets the hideEULA property value. Show or hide EULA to user
func (m *OutOfBoxExperienceSettings) GetHideEULA()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideEULA
    }
}
// GetHidePrivacySettings gets the hidePrivacySettings property value. Show or hide privacy settings to user
func (m *OutOfBoxExperienceSettings) GetHidePrivacySettings()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidePrivacySettings
    }
}
// GetSkipKeyboardSelectionPage gets the skipKeyboardSelectionPage property value. If set, then skip the keyboard selection page if Language and Region are set
func (m *OutOfBoxExperienceSettings) GetSkipKeyboardSelectionPage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.skipKeyboardSelectionPage
    }
}
// GetUserType gets the userType property value. Type of user. Possible values are: administrator, standard.
func (m *OutOfBoxExperienceSettings) GetUserType()(*WindowsUserType) {
    if m == nil {
        return nil
    } else {
        return m.userType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutOfBoxExperienceSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceUsageType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeviceUsageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceUsageType(val.(*WindowsDeviceUsageType))
        }
        return nil
    }
    res["hideEscapeLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideEscapeLink(val)
        }
        return nil
    }
    res["hideEULA"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideEULA(val)
        }
        return nil
    }
    res["hidePrivacySettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHidePrivacySettings(val)
        }
        return nil
    }
    res["skipKeyboardSelectionPage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkipKeyboardSelectionPage(val)
        }
        return nil
    }
    res["userType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsUserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val.(*WindowsUserType))
        }
        return nil
    }
    return res
}
func (m *OutOfBoxExperienceSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OutOfBoxExperienceSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDeviceUsageType() != nil {
        cast := (*m.GetDeviceUsageType()).String()
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
        cast := (*m.GetUserType()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OutOfBoxExperienceSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceUsageType sets the deviceUsageType property value. AAD join authentication type. Possible values are: singleUser, shared.
func (m *OutOfBoxExperienceSettings) SetDeviceUsageType(value *WindowsDeviceUsageType)() {
    if m != nil {
        m.deviceUsageType = value
    }
}
// SetHideEscapeLink sets the hideEscapeLink property value. If set to true, then the user can't start over with different account, on company sign-in
func (m *OutOfBoxExperienceSettings) SetHideEscapeLink(value *bool)() {
    if m != nil {
        m.hideEscapeLink = value
    }
}
// SetHideEULA sets the hideEULA property value. Show or hide EULA to user
func (m *OutOfBoxExperienceSettings) SetHideEULA(value *bool)() {
    if m != nil {
        m.hideEULA = value
    }
}
// SetHidePrivacySettings sets the hidePrivacySettings property value. Show or hide privacy settings to user
func (m *OutOfBoxExperienceSettings) SetHidePrivacySettings(value *bool)() {
    if m != nil {
        m.hidePrivacySettings = value
    }
}
// SetSkipKeyboardSelectionPage sets the skipKeyboardSelectionPage property value. If set, then skip the keyboard selection page if Language and Region are set
func (m *OutOfBoxExperienceSettings) SetSkipKeyboardSelectionPage(value *bool)() {
    if m != nil {
        m.skipKeyboardSelectionPage = value
    }
}
// SetUserType sets the userType property value. Type of user. Possible values are: administrator, standard.
func (m *OutOfBoxExperienceSettings) SetUserType(value *WindowsUserType)() {
    if m != nil {
        m.userType = value
    }
}
