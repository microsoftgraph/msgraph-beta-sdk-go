package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MobileAppIntentAndStateDetail struct {
    additionalData map[string]interface{};
    applicationId *string;
    displayName *string;
    displayVersion *string;
    installState *ResultantAppState;
    mobileAppIntent *MobileAppIntent;
    supportedDeviceTypes []MobileAppSupportedDeviceType;
}
func NewMobileAppIntentAndStateDetail()(*MobileAppIntentAndStateDetail) {
    m := &MobileAppIntentAndStateDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MobileAppIntentAndStateDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MobileAppIntentAndStateDetail) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
func (m *MobileAppIntentAndStateDetail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *MobileAppIntentAndStateDetail) GetDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayVersion
    }
}
func (m *MobileAppIntentAndStateDetail) GetInstallState()(*ResultantAppState) {
    if m == nil {
        return nil
    } else {
        return m.installState
    }
}
func (m *MobileAppIntentAndStateDetail) GetMobileAppIntent()(*MobileAppIntent) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppIntent
    }
}
func (m *MobileAppIntentAndStateDetail) GetSupportedDeviceTypes()([]MobileAppSupportedDeviceType) {
    if m == nil {
        return nil
    } else {
        return m.supportedDeviceTypes
    }
}
func (m *MobileAppIntentAndStateDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationId(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["displayVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayVersion(val)
        return nil
    }
    res["installState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseResultantAppState)
        if err != nil {
            return err
        }
        cast := val.(ResultantAppState)
        m.SetInstallState(&cast)
        return nil
    }
    res["mobileAppIntent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppIntent)
        if err != nil {
            return err
        }
        cast := val.(MobileAppIntent)
        m.SetMobileAppIntent(&cast)
        return nil
    }
    res["supportedDeviceTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppSupportedDeviceType() })
        if err != nil {
            return err
        }
        res := make([]MobileAppSupportedDeviceType, len(val))
        for i, v := range val {
            res[i] = *(v.(*MobileAppSupportedDeviceType))
        }
        m.SetSupportedDeviceTypes(res)
        return nil
    }
    return res
}
func (m *MobileAppIntentAndStateDetail) IsNil()(bool) {
    return m == nil
}
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
        cast := m.GetInstallState().String()
        err := writer.WriteStringValue("installState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppIntent() != nil {
        cast := m.GetMobileAppIntent().String()
        err := writer.WriteStringValue("mobileAppIntent", &cast)
        if err != nil {
            return err
        }
    }
    {
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
func (m *MobileAppIntentAndStateDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MobileAppIntentAndStateDetail) SetApplicationId(value *string)() {
    m.applicationId = value
}
func (m *MobileAppIntentAndStateDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *MobileAppIntentAndStateDetail) SetDisplayVersion(value *string)() {
    m.displayVersion = value
}
func (m *MobileAppIntentAndStateDetail) SetInstallState(value *ResultantAppState)() {
    m.installState = value
}
func (m *MobileAppIntentAndStateDetail) SetMobileAppIntent(value *MobileAppIntent)() {
    m.mobileAppIntent = value
}
func (m *MobileAppIntentAndStateDetail) SetSupportedDeviceTypes(value []MobileAppSupportedDeviceType)() {
    m.supportedDeviceTypes = value
}
