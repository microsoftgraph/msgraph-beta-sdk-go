package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FeatureUsageDetail provides operations to call the getAzureADUserFeatureUsage method.
type FeatureUsageDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    featureName *string;
    // 
    lastConfiguredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    lastUsedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    licenseAssigned *AzureADLicenseType;
    // 
    licenseRequired *AzureADLicenseType;
}
// NewFeatureUsageDetail instantiates a new featureUsageDetail and sets the default values.
func NewFeatureUsageDetail()(*FeatureUsageDetail) {
    m := &FeatureUsageDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFeatureUsageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFeatureUsageDetailFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewFeatureUsageDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FeatureUsageDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFeatureName gets the featureName property value. 
func (m *FeatureUsageDetail) GetFeatureName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.featureName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FeatureUsageDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["featureName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureName(val)
        }
        return nil
    }
    res["lastConfiguredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastConfiguredDateTime(val)
        }
        return nil
    }
    res["lastUsedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsedDateTime(val)
        }
        return nil
    }
    res["licenseAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAzureADLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseAssigned(val.(*AzureADLicenseType))
        }
        return nil
    }
    res["licenseRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAzureADLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseRequired(val.(*AzureADLicenseType))
        }
        return nil
    }
    return res
}
// GetLastConfiguredDateTime gets the lastConfiguredDateTime property value. 
func (m *FeatureUsageDetail) GetLastConfiguredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastConfiguredDateTime
    }
}
// GetLastUsedDateTime gets the lastUsedDateTime property value. 
func (m *FeatureUsageDetail) GetLastUsedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUsedDateTime
    }
}
// GetLicenseAssigned gets the licenseAssigned property value. 
func (m *FeatureUsageDetail) GetLicenseAssigned()(*AzureADLicenseType) {
    if m == nil {
        return nil
    } else {
        return m.licenseAssigned
    }
}
// GetLicenseRequired gets the licenseRequired property value. 
func (m *FeatureUsageDetail) GetLicenseRequired()(*AzureADLicenseType) {
    if m == nil {
        return nil
    } else {
        return m.licenseRequired
    }
}
func (m *FeatureUsageDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *FeatureUsageDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("featureName", m.GetFeatureName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastConfiguredDateTime", m.GetLastConfiguredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastUsedDateTime", m.GetLastUsedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLicenseAssigned() != nil {
        cast := (*m.GetLicenseAssigned()).String()
        err := writer.WriteStringValue("licenseAssigned", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLicenseRequired() != nil {
        cast := (*m.GetLicenseRequired()).String()
        err := writer.WriteStringValue("licenseRequired", &cast)
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
func (m *FeatureUsageDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFeatureName sets the featureName property value. 
func (m *FeatureUsageDetail) SetFeatureName(value *string)() {
    if m != nil {
        m.featureName = value
    }
}
// SetLastConfiguredDateTime sets the lastConfiguredDateTime property value. 
func (m *FeatureUsageDetail) SetLastConfiguredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastConfiguredDateTime = value
    }
}
// SetLastUsedDateTime sets the lastUsedDateTime property value. 
func (m *FeatureUsageDetail) SetLastUsedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUsedDateTime = value
    }
}
// SetLicenseAssigned sets the licenseAssigned property value. 
func (m *FeatureUsageDetail) SetLicenseAssigned(value *AzureADLicenseType)() {
    if m != nil {
        m.licenseAssigned = value
    }
}
// SetLicenseRequired sets the licenseRequired property value. 
func (m *FeatureUsageDetail) SetLicenseRequired(value *AzureADLicenseType)() {
    if m != nil {
        m.licenseRequired = value
    }
}
