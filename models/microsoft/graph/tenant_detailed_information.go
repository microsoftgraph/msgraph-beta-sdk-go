package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TenantDetailedInformation struct {
    Entity
    city *string;
    countryCode *string;
    countryName *string;
    defaultDomainName *string;
    displayName *string;
    industryName *string;
    region *string;
    segmentName *string;
    tenantId *string;
    verticalName *string;
}
func NewTenantDetailedInformation()(*TenantDetailedInformation) {
    m := &TenantDetailedInformation{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TenantDetailedInformation) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
func (m *TenantDetailedInformation) GetCountryCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryCode
    }
}
func (m *TenantDetailedInformation) GetCountryName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryName
    }
}
func (m *TenantDetailedInformation) GetDefaultDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDomainName
    }
}
func (m *TenantDetailedInformation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *TenantDetailedInformation) GetIndustryName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.industryName
    }
}
func (m *TenantDetailedInformation) GetRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.region
    }
}
func (m *TenantDetailedInformation) GetSegmentName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.segmentName
    }
}
func (m *TenantDetailedInformation) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
func (m *TenantDetailedInformation) GetVerticalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.verticalName
    }
}
func (m *TenantDetailedInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["city"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCity(val)
        return nil
    }
    res["countryCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCountryCode(val)
        return nil
    }
    res["countryName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCountryName(val)
        return nil
    }
    res["defaultDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultDomainName(val)
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
    res["industryName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIndustryName(val)
        return nil
    }
    res["region"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRegion(val)
        return nil
    }
    res["segmentName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSegmentName(val)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    res["verticalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVerticalName(val)
        return nil
    }
    return res
}
func (m *TenantDetailedInformation) IsNil()(bool) {
    return m == nil
}
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
func (m *TenantDetailedInformation) SetCity(value *string)() {
    m.city = value
}
func (m *TenantDetailedInformation) SetCountryCode(value *string)() {
    m.countryCode = value
}
func (m *TenantDetailedInformation) SetCountryName(value *string)() {
    m.countryName = value
}
func (m *TenantDetailedInformation) SetDefaultDomainName(value *string)() {
    m.defaultDomainName = value
}
func (m *TenantDetailedInformation) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *TenantDetailedInformation) SetIndustryName(value *string)() {
    m.industryName = value
}
func (m *TenantDetailedInformation) SetRegion(value *string)() {
    m.region = value
}
func (m *TenantDetailedInformation) SetSegmentName(value *string)() {
    m.segmentName = value
}
func (m *TenantDetailedInformation) SetTenantId(value *string)() {
    m.tenantId = value
}
func (m *TenantDetailedInformation) SetVerticalName(value *string)() {
    m.verticalName = value
}
