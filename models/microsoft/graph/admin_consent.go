package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AdminConsent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The admin consent state of sharing user and device data to Apple. Possible values are: notConfigured, granted, notGranted.
    shareAPNSData *AdminConsentState;
    // Gets or sets the admin consent for sharing User experience analytics data. Possible values are: notConfigured, granted, notGranted.
    shareUserExperienceAnalyticsData *AdminConsentState;
}
// Instantiates a new adminConsent and sets the default values.
func NewAdminConsent()(*AdminConsent) {
    m := &AdminConsent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminConsent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the shareAPNSData property value. The admin consent state of sharing user and device data to Apple. Possible values are: notConfigured, granted, notGranted.
func (m *AdminConsent) GetShareAPNSData()(*AdminConsentState) {
    if m == nil {
        return nil
    } else {
        return m.shareAPNSData
    }
}
// Gets the shareUserExperienceAnalyticsData property value. Gets or sets the admin consent for sharing User experience analytics data. Possible values are: notConfigured, granted, notGranted.
func (m *AdminConsent) GetShareUserExperienceAnalyticsData()(*AdminConsentState) {
    if m == nil {
        return nil
    } else {
        return m.shareUserExperienceAnalyticsData
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AdminConsent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the shareAPNSData property value. The admin consent state of sharing user and device data to Apple. Possible values are: notConfigured, granted, notGranted.
// Parameters:
//  - value : Value to set for the shareAPNSData property.
func (m *AdminConsent) SetShareAPNSData(value *AdminConsentState)() {
    m.shareAPNSData = value
}
// Sets the shareUserExperienceAnalyticsData property value. Gets or sets the admin consent for sharing User experience analytics data. Possible values are: notConfigured, granted, notGranted.
// Parameters:
//  - value : Value to set for the shareUserExperienceAnalyticsData property.
func (m *AdminConsent) SetShareUserExperienceAnalyticsData(value *AdminConsentState)() {
    m.shareUserExperienceAnalyticsData = value
}
