package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IpSecurityProfile struct {
    Entity
    // 
    activityGroupNames []string;
    // 
    address *string;
    // 
    azureSubscriptionId *string;
    // 
    azureTenantId *string;
    // 
    countHits *int32;
    // 
    countHosts *int32;
    // 
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    ipCategories []IpCategory;
    // 
    ipReferenceData []IpReferenceData;
    // 
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    riskScore *string;
    // 
    tags []string;
    // 
    vendorInformation *SecurityVendorInformation;
}
// Instantiates a new ipSecurityProfile and sets the default values.
func NewIpSecurityProfile()(*IpSecurityProfile) {
    m := &IpSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activityGroupNames property value. 
func (m *IpSecurityProfile) GetActivityGroupNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.activityGroupNames
    }
}
// Gets the address property value. 
func (m *IpSecurityProfile) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// Gets the azureSubscriptionId property value. 
func (m *IpSecurityProfile) GetAzureSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureSubscriptionId
    }
}
// Gets the azureTenantId property value. 
func (m *IpSecurityProfile) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// Gets the countHits property value. 
func (m *IpSecurityProfile) GetCountHits()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.countHits
    }
}
// Gets the countHosts property value. 
func (m *IpSecurityProfile) GetCountHosts()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.countHosts
    }
}
// Gets the firstSeenDateTime property value. 
func (m *IpSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.firstSeenDateTime
    }
}
// Gets the ipCategories property value. 
func (m *IpSecurityProfile) GetIpCategories()([]IpCategory) {
    if m == nil {
        return nil
    } else {
        return m.ipCategories
    }
}
// Gets the ipReferenceData property value. 
func (m *IpSecurityProfile) GetIpReferenceData()([]IpReferenceData) {
    if m == nil {
        return nil
    } else {
        return m.ipReferenceData
    }
}
// Gets the lastSeenDateTime property value. 
func (m *IpSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSeenDateTime
    }
}
// Gets the riskScore property value. 
func (m *IpSecurityProfile) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// Gets the tags property value. 
func (m *IpSecurityProfile) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// Gets the vendorInformation property value. 
func (m *IpSecurityProfile) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// The deserialization information for the current model
func (m *IpSecurityProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityGroupNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetActivityGroupNames(res)
        }
        return nil
    }
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["azureSubscriptionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureSubscriptionId(val)
        }
        return nil
    }
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["countHits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountHits(val)
        }
        return nil
    }
    res["countHosts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountHosts(val)
        }
        return nil
    }
    res["firstSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstSeenDateTime(val)
        }
        return nil
    }
    res["ipCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIpCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*IpCategory))
            }
            m.SetIpCategories(res)
        }
        return nil
    }
    res["ipReferenceData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIpReferenceData() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpReferenceData, len(val))
            for i, v := range val {
                res[i] = *(v.(*IpReferenceData))
            }
            m.SetIpReferenceData(res)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
        }
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    res["vendorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityVendorInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInformation(val.(*SecurityVendorInformation))
        }
        return nil
    }
    return res
}
func (m *IpSecurityProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IpSecurityProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("activityGroupNames", m.GetActivityGroupNames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureSubscriptionId", m.GetAzureSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("countHits", m.GetCountHits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("countHosts", m.GetCountHosts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("firstSeenDateTime", m.GetFirstSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIpCategories()))
        for i, v := range m.GetIpCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("ipCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIpReferenceData()))
        for i, v := range m.GetIpReferenceData() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("ipReferenceData", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("riskScore", m.GetRiskScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the activityGroupNames property value. 
// Parameters:
//  - value : Value to set for the activityGroupNames property.
func (m *IpSecurityProfile) SetActivityGroupNames(value []string)() {
    m.activityGroupNames = value
}
// Sets the address property value. 
// Parameters:
//  - value : Value to set for the address property.
func (m *IpSecurityProfile) SetAddress(value *string)() {
    m.address = value
}
// Sets the azureSubscriptionId property value. 
// Parameters:
//  - value : Value to set for the azureSubscriptionId property.
func (m *IpSecurityProfile) SetAzureSubscriptionId(value *string)() {
    m.azureSubscriptionId = value
}
// Sets the azureTenantId property value. 
// Parameters:
//  - value : Value to set for the azureTenantId property.
func (m *IpSecurityProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// Sets the countHits property value. 
// Parameters:
//  - value : Value to set for the countHits property.
func (m *IpSecurityProfile) SetCountHits(value *int32)() {
    m.countHits = value
}
// Sets the countHosts property value. 
// Parameters:
//  - value : Value to set for the countHosts property.
func (m *IpSecurityProfile) SetCountHosts(value *int32)() {
    m.countHosts = value
}
// Sets the firstSeenDateTime property value. 
// Parameters:
//  - value : Value to set for the firstSeenDateTime property.
func (m *IpSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// Sets the ipCategories property value. 
// Parameters:
//  - value : Value to set for the ipCategories property.
func (m *IpSecurityProfile) SetIpCategories(value []IpCategory)() {
    m.ipCategories = value
}
// Sets the ipReferenceData property value. 
// Parameters:
//  - value : Value to set for the ipReferenceData property.
func (m *IpSecurityProfile) SetIpReferenceData(value []IpReferenceData)() {
    m.ipReferenceData = value
}
// Sets the lastSeenDateTime property value. 
// Parameters:
//  - value : Value to set for the lastSeenDateTime property.
func (m *IpSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// Sets the riskScore property value. 
// Parameters:
//  - value : Value to set for the riskScore property.
func (m *IpSecurityProfile) SetRiskScore(value *string)() {
    m.riskScore = value
}
// Sets the tags property value. 
// Parameters:
//  - value : Value to set for the tags property.
func (m *IpSecurityProfile) SetTags(value []string)() {
    m.tags = value
}
// Sets the vendorInformation property value. 
// Parameters:
//  - value : Value to set for the vendorInformation property.
func (m *IpSecurityProfile) SetVendorInformation(value *SecurityVendorInformation)() {
    m.vendorInformation = value
}
