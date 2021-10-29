package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new outOfBoxExperienceSettings and sets the default values.
func NewOutOfBoxExperienceSettings()(*OutOfBoxExperienceSettings) {
    m := &OutOfBoxExperienceSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OutOfBoxExperienceSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the deviceUsageType property value. AAD join authentication type. Possible values are: singleUser, shared.
func (m *OutOfBoxExperienceSettings) GetDeviceUsageType()(*WindowsDeviceUsageType) {
    if m == nil {
        return nil
    } else {
        return m.deviceUsageType
    }
}
// Gets the hideEscapeLink property value. If set to true, then the user can't start over with different account, on company sign-in
func (m *OutOfBoxExperienceSettings) GetHideEscapeLink()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideEscapeLink
    }
}
// Gets the hideEULA property value. Show or hide EULA to user
func (m *OutOfBoxExperienceSettings) GetHideEULA()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideEULA
    }
}
// Gets the hidePrivacySettings property value. Show or hide privacy settings to user
func (m *OutOfBoxExperienceSettings) GetHidePrivacySettings()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidePrivacySettings
    }
}
// Gets the skipKeyboardSelectionPage property value. If set, then skip the keyboard selection page if Language and Region are set
func (m *OutOfBoxExperienceSettings) GetSkipKeyboardSelectionPage()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.skipKeyboardSelectionPage
    }
}
// Gets the userType property value. Type of user. Possible values are: administrator, standard.
func (m *OutOfBoxExperienceSettings) GetUserType()(*WindowsUserType) {
    if m == nil {
        return nil
    } else {
        return m.userType
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *OutOfBoxExperienceSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the deviceUsageType property value. AAD join authentication type. Possible values are: singleUser, shared.
// Parameters:
//  - value : Value to set for the deviceUsageType property.
func (m *OutOfBoxExperienceSettings) SetDeviceUsageType(value *WindowsDeviceUsageType)() {
    m.deviceUsageType = value
}
// Sets the hideEscapeLink property value. If set to true, then the user can't start over with different account, on company sign-in
// Parameters:
//  - value : Value to set for the hideEscapeLink property.
func (m *OutOfBoxExperienceSettings) SetHideEscapeLink(value *bool)() {
    m.hideEscapeLink = value
}
// Sets the hideEULA property value. Show or hide EULA to user
// Parameters:
//  - value : Value to set for the hideEULA property.
func (m *OutOfBoxExperienceSettings) SetHideEULA(value *bool)() {
    m.hideEULA = value
}
// Sets the hidePrivacySettings property value. Show or hide privacy settings to user
// Parameters:
//  - value : Value to set for the hidePrivacySettings property.
func (m *OutOfBoxExperienceSettings) SetHidePrivacySettings(value *bool)() {
    m.hidePrivacySettings = value
}
// Sets the skipKeyboardSelectionPage property value. If set, then skip the keyboard selection page if Language and Region are set
// Parameters:
//  - value : Value to set for the skipKeyboardSelectionPage property.
func (m *OutOfBoxExperienceSettings) SetSkipKeyboardSelectionPage(value *bool)() {
    m.skipKeyboardSelectionPage = value
}
// Sets the userType property value. Type of user. Possible values are: administrator, standard.
// Parameters:
//  - value : Value to set for the userType property.
func (m *OutOfBoxExperienceSettings) SetUserType(value *WindowsUserType)() {
    m.userType = value
}
