package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AzureADUserFeatureUsage provides operations to call the getAzureADUserFeatureUsage method.
type AzureADUserFeatureUsage struct {
    Entity
    // 
    featureUsageDetails []FeatureUsageDetailable;
    // 
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    licenseAssigned *AzureADLicenseType;
    // 
    licenseRecommended *AzureADLicenseType;
    // 
    userDisplayName *string;
    // 
    userId *string;
    // 
    userPrincipalName *string;
}
// NewAzureADUserFeatureUsage instantiates a new azureADUserFeatureUsage and sets the default values.
func NewAzureADUserFeatureUsage()(*AzureADUserFeatureUsage) {
    m := &AzureADUserFeatureUsage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAzureADUserFeatureUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureADUserFeatureUsageFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAzureADUserFeatureUsage(), nil
}
// GetFeatureUsageDetails gets the featureUsageDetails property value. 
func (m *AzureADUserFeatureUsage) GetFeatureUsageDetails()([]FeatureUsageDetailable) {
    if m == nil {
        return nil
    } else {
        return m.featureUsageDetails
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureADUserFeatureUsage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["featureUsageDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFeatureUsageDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FeatureUsageDetailable, len(val))
            for i, v := range val {
                res[i] = v.(FeatureUsageDetailable)
            }
            m.SetFeatureUsageDetails(res)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
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
    res["licenseRecommended"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAzureADLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseRecommended(val.(*AzureADLicenseType))
        }
        return nil
    }
    res["userDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. 
func (m *AzureADUserFeatureUsage) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// GetLicenseAssigned gets the licenseAssigned property value. 
func (m *AzureADUserFeatureUsage) GetLicenseAssigned()(*AzureADLicenseType) {
    if m == nil {
        return nil
    } else {
        return m.licenseAssigned
    }
}
// GetLicenseRecommended gets the licenseRecommended property value. 
func (m *AzureADUserFeatureUsage) GetLicenseRecommended()(*AzureADLicenseType) {
    if m == nil {
        return nil
    } else {
        return m.licenseRecommended
    }
}
// GetUserDisplayName gets the userDisplayName property value. 
func (m *AzureADUserFeatureUsage) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// GetUserId gets the userId property value. 
func (m *AzureADUserFeatureUsage) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. 
func (m *AzureADUserFeatureUsage) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *AzureADUserFeatureUsage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AzureADUserFeatureUsage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetFeatureUsageDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFeatureUsageDetails()))
        for i, v := range m.GetFeatureUsageDetails() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("featureUsageDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLicenseAssigned() != nil {
        cast := (*m.GetLicenseAssigned()).String()
        err = writer.WriteStringValue("licenseAssigned", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLicenseRecommended() != nil {
        cast := (*m.GetLicenseRecommended()).String()
        err = writer.WriteStringValue("licenseRecommended", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFeatureUsageDetails sets the featureUsageDetails property value. 
func (m *AzureADUserFeatureUsage) SetFeatureUsageDetails(value []FeatureUsageDetailable)() {
    if m != nil {
        m.featureUsageDetails = value
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. 
func (m *AzureADUserFeatureUsage) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdatedDateTime = value
    }
}
// SetLicenseAssigned sets the licenseAssigned property value. 
func (m *AzureADUserFeatureUsage) SetLicenseAssigned(value *AzureADLicenseType)() {
    if m != nil {
        m.licenseAssigned = value
    }
}
// SetLicenseRecommended sets the licenseRecommended property value. 
func (m *AzureADUserFeatureUsage) SetLicenseRecommended(value *AzureADLicenseType)() {
    if m != nil {
        m.licenseRecommended = value
    }
}
// SetUserDisplayName sets the userDisplayName property value. 
func (m *AzureADUserFeatureUsage) SetUserDisplayName(value *string)() {
    if m != nil {
        m.userDisplayName = value
    }
}
// SetUserId sets the userId property value. 
func (m *AzureADUserFeatureUsage) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. 
func (m *AzureADUserFeatureUsage) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
