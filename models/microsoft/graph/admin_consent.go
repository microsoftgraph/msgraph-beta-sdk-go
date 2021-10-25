package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AdminConsent struct {
    additionalData map[string]interface{};
    shareAPNSData *AdminConsentState;
    shareUserExperienceAnalyticsData *AdminConsentState;
}
func NewAdminConsent()(*AdminConsent) {
    m := &AdminConsent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AdminConsent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AdminConsent) GetShareAPNSData()(*AdminConsentState) {
    if m == nil {
        return nil
    } else {
        return m.shareAPNSData
    }
}
func (m *AdminConsent) GetShareUserExperienceAnalyticsData()(*AdminConsentState) {
    if m == nil {
        return nil
    } else {
        return m.shareUserExperienceAnalyticsData
    }
}
func (m *AdminConsent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["shareAPNSData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdminConsentState)
        if err != nil {
            return err
        }
        cast := val.(AdminConsentState)
        m.SetShareAPNSData(&cast)
        return nil
    }
    res["shareUserExperienceAnalyticsData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdminConsentState)
        if err != nil {
            return err
        }
        cast := val.(AdminConsentState)
        m.SetShareUserExperienceAnalyticsData(&cast)
        return nil
    }
    return res
}
func (m *AdminConsent) IsNil()(bool) {
    return m == nil
}
func (m *AdminConsent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetShareAPNSData() != nil {
        cast := m.GetShareAPNSData().String()
        err := writer.WriteStringValue("shareAPNSData", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetShareUserExperienceAnalyticsData() != nil {
        cast := m.GetShareUserExperienceAnalyticsData().String()
        err := writer.WriteStringValue("shareUserExperienceAnalyticsData", &cast)
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
func (m *AdminConsent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AdminConsent) SetShareAPNSData(value *AdminConsentState)() {
    m.shareAPNSData = value
}
func (m *AdminConsent) SetShareUserExperienceAnalyticsData(value *AdminConsentState)() {
    m.shareUserExperienceAnalyticsData = value
}
