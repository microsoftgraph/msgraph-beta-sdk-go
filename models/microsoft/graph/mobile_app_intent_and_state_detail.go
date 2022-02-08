package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileAppIntentAndStateDetail 
type MobileAppIntentAndStateDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // MobieApp identifier.
    applicationId *string;
    // The admin provided or imported title of the app.
    displayName *string;
    // Human readable version of the application
    displayVersion *string;
    // The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
    installState *ResultantAppState;
    // Mobile App Intent. Possible values are: available, notAvailable, requiredInstall, requiredUninstall, requiredAndAvailableInstall, availableInstallWithoutEnrollment, exclude.
    mobileAppIntent *MobileAppIntent;
    // The supported platforms for the app.
    supportedDeviceTypes []MobileAppSupportedDeviceType;
}
// NewMobileAppIntentAndStateDetail instantiates a new mobileAppIntentAndStateDetail and sets the default values.
func NewMobileAppIntentAndStateDetail()(*MobileAppIntentAndStateDetail) {
    m := &MobileAppIntentAndStateDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppIntentAndStateDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationId gets the applicationId property value. MobieApp identifier.
func (m *MobileAppIntentAndStateDetail) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
// GetDisplayName gets the displayName property value. The admin provided or imported title of the app.
func (m *MobileAppIntentAndStateDetail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDisplayVersion gets the displayVersion property value. Human readable version of the application
func (m *MobileAppIntentAndStateDetail) GetDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayVersion
    }
}
// GetInstallState gets the installState property value. The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
func (m *MobileAppIntentAndStateDetail) GetInstallState()(*ResultantAppState) {
    if m == nil {
        return nil
    } else {
        return m.installState
    }
}
// GetMobileAppIntent gets the mobileAppIntent property value. Mobile App Intent. Possible values are: available, notAvailable, requiredInstall, requiredUninstall, requiredAndAvailableInstall, availableInstallWithoutEnrollment, exclude.
func (m *MobileAppIntentAndStateDetail) GetMobileAppIntent()(*MobileAppIntent) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppIntent
    }
}
// GetSupportedDeviceTypes gets the supportedDeviceTypes property value. The supported platforms for the app.
func (m *MobileAppIntentAndStateDetail) GetSupportedDeviceTypes()([]MobileAppSupportedDeviceType) {
    if m == nil {
        return nil
    } else {
        return m.supportedDeviceTypes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppIntentAndStateDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["displayVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayVersion(val)
        }
        return nil
    }
    res["installState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallState(val.(*ResultantAppState))
        }
        return nil
    }
    res["mobileAppIntent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileAppIntent(val.(*MobileAppIntent))
        }
        return nil
    }
    res["supportedDeviceTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppSupportedDeviceType() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppSupportedDeviceType, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileAppSupportedDeviceType))
            }
            m.SetSupportedDeviceTypes(res)
        }
        return nil
    }
    return res
}
func (m *MobileAppIntentAndStateDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MobileAppIntentAndStateDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayVersion", m.GetDisplayVersion())
        if err != nil {
            return err
        }
    }
    if m.GetInstallState() != nil {
        cast := (*m.GetInstallState()).String()
        err := writer.WriteStringValue("installState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppIntent() != nil {
        cast := (*m.GetMobileAppIntent()).String()
        err := writer.WriteStringValue("mobileAppIntent", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSupportedDeviceTypes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSupportedDeviceTypes()))
        for i, v := range m.GetSupportedDeviceTypes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("supportedDeviceTypes", cast)
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
func (m *MobileAppIntentAndStateDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationId sets the applicationId property value. MobieApp identifier.
func (m *MobileAppIntentAndStateDetail) SetApplicationId(value *string)() {
    if m != nil {
        m.applicationId = value
    }
}
// SetDisplayName sets the displayName property value. The admin provided or imported title of the app.
func (m *MobileAppIntentAndStateDetail) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDisplayVersion sets the displayVersion property value. Human readable version of the application
func (m *MobileAppIntentAndStateDetail) SetDisplayVersion(value *string)() {
    if m != nil {
        m.displayVersion = value
    }
}
// SetInstallState sets the installState property value. The install state of the app. Possible values are: installed, failed, notInstalled, uninstallFailed, pendingInstall, unknown, notApplicable.
func (m *MobileAppIntentAndStateDetail) SetInstallState(value *ResultantAppState)() {
    if m != nil {
        m.installState = value
    }
}
// SetMobileAppIntent sets the mobileAppIntent property value. Mobile App Intent. Possible values are: available, notAvailable, requiredInstall, requiredUninstall, requiredAndAvailableInstall, availableInstallWithoutEnrollment, exclude.
func (m *MobileAppIntentAndStateDetail) SetMobileAppIntent(value *MobileAppIntent)() {
    if m != nil {
        m.mobileAppIntent = value
    }
}
// SetSupportedDeviceTypes sets the supportedDeviceTypes property value. The supported platforms for the app.
func (m *MobileAppIntentAndStateDetail) SetSupportedDeviceTypes(value []MobileAppSupportedDeviceType)() {
    if m != nil {
        m.supportedDeviceTypes = value
    }
}
