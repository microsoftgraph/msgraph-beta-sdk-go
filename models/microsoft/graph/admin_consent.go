package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AdminConsent 
type AdminConsent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The admin consent state of sharing user and device data to Apple. Possible values are: notConfigured, granted, notGranted.
    shareAPNSData *AdminConsentState;
    // Gets or sets the admin consent for sharing User experience analytics data. Possible values are: notConfigured, granted, notGranted.
    shareUserExperienceAnalyticsData *AdminConsentState;
}
// NewAdminConsent instantiates a new adminConsent and sets the default values.
func NewAdminConsent()(*AdminConsent) {
    m := &AdminConsent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminConsent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetShareAPNSData gets the shareAPNSData property value. The admin consent state of sharing user and device data to Apple. Possible values are: notConfigured, granted, notGranted.
func (m *AdminConsent) GetShareAPNSData()(*AdminConsentState) {
    if m == nil {
        return nil
    } else {
        return m.shareAPNSData
    }
}
// GetShareUserExperienceAnalyticsData gets the shareUserExperienceAnalyticsData property value. Gets or sets the admin consent for sharing User experience analytics data. Possible values are: notConfigured, granted, notGranted.
func (m *AdminConsent) GetShareUserExperienceAnalyticsData()(*AdminConsentState) {
    if m == nil {
        return nil
    } else {
        return m.shareUserExperienceAnalyticsData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminConsent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["shareAPNSData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdminConsentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareAPNSData(val.(*AdminConsentState))
        }
        return nil
    }
    res["shareUserExperienceAnalyticsData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdminConsentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareUserExperienceAnalyticsData(val.(*AdminConsentState))
        }
        return nil
    }
    return res
}
func (m *AdminConsent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AdminConsent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetShareAPNSData() != nil {
        cast := (*m.GetShareAPNSData()).String()
        err := writer.WriteStringValue("shareAPNSData", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetShareUserExperienceAnalyticsData() != nil {
        cast := (*m.GetShareUserExperienceAnalyticsData()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AdminConsent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetShareAPNSData sets the shareAPNSData property value. The admin consent state of sharing user and device data to Apple. Possible values are: notConfigured, granted, notGranted.
func (m *AdminConsent) SetShareAPNSData(value *AdminConsentState)() {
    if m != nil {
        m.shareAPNSData = value
    }
}
// SetShareUserExperienceAnalyticsData sets the shareUserExperienceAnalyticsData property value. Gets or sets the admin consent for sharing User experience analytics data. Possible values are: notConfigured, granted, notGranted.
func (m *AdminConsent) SetShareUserExperienceAnalyticsData(value *AdminConsentState)() {
    if m != nil {
        m.shareUserExperienceAnalyticsData = value
    }
}
