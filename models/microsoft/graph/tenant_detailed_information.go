package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TenantDetailedInformation struct {
    Entity
    // The city where the managed tenant is located. Optional. Read-only.
    city *string;
    // The code for the country where the managed tenant is located. Optional. Read-only.
    countryCode *string;
    // The name for the country where the managed tenant is located. Optional. Read-only.
    countryName *string;
    // The default domain name for the managed tenant. Optional. Read-only.
    defaultDomainName *string;
    // The display name for the managed tenant.
    displayName *string;
    // The business industry associated with the managed tenant. Optional. Read-only.
    industryName *string;
    // The region where the managed tenant is located. Optional. Read-only.
    region *string;
    // The business segment associated with the managed tenant. Optional. Read-only.
    segmentName *string;
    // The Azure Active Directory tenant identifier for the managed tenant.
    tenantId *string;
    // The vertical associated with the managed tenant. Optional. Read-only.
    verticalName *string;
}
// Instantiates a new tenantDetailedInformation and sets the default values.
func NewTenantDetailedInformation()(*TenantDetailedInformation) {
    m := &TenantDetailedInformation{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the city property value. The city where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// Gets the countryCode property value. The code for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetCountryCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryCode
    }
}
// Gets the countryName property value. The name for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetCountryName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryName
    }
}
// Gets the defaultDomainName property value. The default domain name for the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetDefaultDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDomainName
    }
}
// Gets the displayName property value. The display name for the managed tenant.
func (m *TenantDetailedInformation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the industryName property value. The business industry associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetIndustryName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.industryName
    }
}
// Gets the region property value. The region where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.region
    }
}
// Gets the segmentName property value. The business segment associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetSegmentName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.segmentName
    }
}
// Gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant.
func (m *TenantDetailedInformation) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Gets the verticalName property value. The vertical associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetVerticalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.verticalName
    }
}
// The deserialization information for the current model
func (m *TenantDetailedInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["countryCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryCode(val)
        }
        return nil
    }
    res["countryName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryName(val)
        }
        return nil
    }
    res["defaultDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDomainName(val)
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
    res["industryName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndustryName(val)
        }
        return nil
    }
    res["region"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["segmentName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSegmentName(val)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["verticalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerticalName(val)
        }
        return nil
    }
    return res
}
func (m *TenantDetailedInformation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TenantDetailedInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countryCode", m.GetCountryCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countryName", m.GetCountryName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultDomainName", m.GetDefaultDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("industryName", m.GetIndustryName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("segmentName", m.GetSegmentName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("verticalName", m.GetVerticalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the city property value. The city where the managed tenant is located. Optional. Read-only.
// Parameters:
//  - value : Value to set for the city property.
func (m *TenantDetailedInformation) SetCity(value *string)() {
    m.city = value
}
// Sets the countryCode property value. The code for the country where the managed tenant is located. Optional. Read-only.
// Parameters:
//  - value : Value to set for the countryCode property.
func (m *TenantDetailedInformation) SetCountryCode(value *string)() {
    m.countryCode = value
}
// Sets the countryName property value. The name for the country where the managed tenant is located. Optional. Read-only.
// Parameters:
//  - value : Value to set for the countryName property.
func (m *TenantDetailedInformation) SetCountryName(value *string)() {
    m.countryName = value
}
// Sets the defaultDomainName property value. The default domain name for the managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the defaultDomainName property.
func (m *TenantDetailedInformation) SetDefaultDomainName(value *string)() {
    m.defaultDomainName = value
}
// Sets the displayName property value. The display name for the managed tenant.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *TenantDetailedInformation) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the industryName property value. The business industry associated with the managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the industryName property.
func (m *TenantDetailedInformation) SetIndustryName(value *string)() {
    m.industryName = value
}
// Sets the region property value. The region where the managed tenant is located. Optional. Read-only.
// Parameters:
//  - value : Value to set for the region property.
func (m *TenantDetailedInformation) SetRegion(value *string)() {
    m.region = value
}
// Sets the segmentName property value. The business segment associated with the managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the segmentName property.
func (m *TenantDetailedInformation) SetSegmentName(value *string)() {
    m.segmentName = value
}
// Sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant.
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *TenantDetailedInformation) SetTenantId(value *string)() {
    m.tenantId = value
}
// Sets the verticalName property value. The vertical associated with the managed tenant. Optional. Read-only.
// Parameters:
//  - value : Value to set for the verticalName property.
func (m *TenantDetailedInformation) SetVerticalName(value *string)() {
    m.verticalName = value
}
