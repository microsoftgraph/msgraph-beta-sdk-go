package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsPrivacyDataAccessControlItem struct {
    Entity
    // This indicates an access level for the privacy data category to which the specified application will be given to. Possible values are: notConfigured, forceAllow, forceDeny, userInControl.
    accessLevel *WindowsPrivacyDataAccessLevel;
    // The Package Family Name of a Windows app. When set, the access level applies to the specified application.
    appDisplayName *string;
    // The Package Family Name of a Windows app. When set, the access level applies to the specified application.
    appPackageFamilyName *string;
    // This indicates a privacy data category to which the specific access control will apply. Possible values are: notConfigured, accountInfo, appsRunInBackground, calendar, callHistory, camera, contacts, diagnosticsInfo, email, location, messaging, microphone, motion, notifications, phone, radios, tasks, syncWithDevices, trustedDevices.
    dataCategory *WindowsPrivacyDataCategory;
}
// Instantiates a new windowsPrivacyDataAccessControlItem and sets the default values.
func NewWindowsPrivacyDataAccessControlItem()(*WindowsPrivacyDataAccessControlItem) {
    m := &WindowsPrivacyDataAccessControlItem{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessLevel property value. This indicates an access level for the privacy data category to which the specified application will be given to. Possible values are: notConfigured, forceAllow, forceDeny, userInControl.
func (m *WindowsPrivacyDataAccessControlItem) GetAccessLevel()(*WindowsPrivacyDataAccessLevel) {
    if m == nil {
        return nil
    } else {
        return m.accessLevel
    }
}
// Gets the appDisplayName property value. The Package Family Name of a Windows app. When set, the access level applies to the specified application.
func (m *WindowsPrivacyDataAccessControlItem) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// Gets the appPackageFamilyName property value. The Package Family Name of a Windows app. When set, the access level applies to the specified application.
func (m *WindowsPrivacyDataAccessControlItem) GetAppPackageFamilyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appPackageFamilyName
    }
}
// Gets the dataCategory property value. This indicates a privacy data category to which the specific access control will apply. Possible values are: notConfigured, accountInfo, appsRunInBackground, calendar, callHistory, camera, contacts, diagnosticsInfo, email, location, messaging, microphone, motion, notifications, phone, radios, tasks, syncWithDevices, trustedDevices.
func (m *WindowsPrivacyDataAccessControlItem) GetDataCategory()(*WindowsPrivacyDataCategory) {
    if m == nil {
        return nil
    } else {
        return m.dataCategory
    }
}
// The deserialization information for the current model
func (m *WindowsPrivacyDataAccessControlItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsPrivacyDataAccessLevel)
        if err != nil {
            return err
        }
        cast := val.(WindowsPrivacyDataAccessLevel)
        m.SetAccessLevel(&cast)
        return nil
    }
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppDisplayName(val)
        return nil
    }
    res["appPackageFamilyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppPackageFamilyName(val)
        return nil
    }
    res["dataCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsPrivacyDataCategory)
        if err != nil {
            return err
        }
        cast := val.(WindowsPrivacyDataCategory)
        m.SetDataCategory(&cast)
        return nil
    }
    return res
}
func (m *WindowsPrivacyDataAccessControlItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsPrivacyDataAccessControlItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessLevel() != nil {
        cast := m.GetAccessLevel().String()
        err = writer.WriteStringValue("accessLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appPackageFamilyName", m.GetAppPackageFamilyName())
        if err != nil {
            return err
        }
    }
    if m.GetDataCategory() != nil {
        cast := m.GetDataCategory().String()
        err = writer.WriteStringValue("dataCategory", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the accessLevel property value. This indicates an access level for the privacy data category to which the specified application will be given to. Possible values are: notConfigured, forceAllow, forceDeny, userInControl.
// Parameters:
//  - value : Value to set for the accessLevel property.
func (m *WindowsPrivacyDataAccessControlItem) SetAccessLevel(value *WindowsPrivacyDataAccessLevel)() {
    m.accessLevel = value
}
// Sets the appDisplayName property value. The Package Family Name of a Windows app. When set, the access level applies to the specified application.
// Parameters:
//  - value : Value to set for the appDisplayName property.
func (m *WindowsPrivacyDataAccessControlItem) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// Sets the appPackageFamilyName property value. The Package Family Name of a Windows app. When set, the access level applies to the specified application.
// Parameters:
//  - value : Value to set for the appPackageFamilyName property.
func (m *WindowsPrivacyDataAccessControlItem) SetAppPackageFamilyName(value *string)() {
    m.appPackageFamilyName = value
}
// Sets the dataCategory property value. This indicates a privacy data category to which the specific access control will apply. Possible values are: notConfigured, accountInfo, appsRunInBackground, calendar, callHistory, camera, contacts, diagnosticsInfo, email, location, messaging, microphone, motion, notifications, phone, radios, tasks, syncWithDevices, trustedDevices.
// Parameters:
//  - value : Value to set for the dataCategory property.
func (m *WindowsPrivacyDataAccessControlItem) SetDataCategory(value *WindowsPrivacyDataCategory)() {
    m.dataCategory = value
}
