package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DomainSecurityProfile struct {
    Entity
    // 
    activityGroupNames []string;
    // 
    azureSubscriptionId *string;
    // 
    azureTenantId *string;
    // 
    countHits *int32;
    // 
    countInOrg *int32;
    // 
    domainCategories []ReputationCategory;
    // 
    domainRegisteredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    name *string;
    // 
    registrant *DomainRegistrant;
    // 
    riskScore *string;
    // 
    tags []string;
    // 
    vendorInformation *SecurityVendorInformation;
}
// Instantiates a new domainSecurityProfile and sets the default values.
func NewDomainSecurityProfile()(*DomainSecurityProfile) {
    m := &DomainSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activityGroupNames property value. 
func (m *DomainSecurityProfile) GetActivityGroupNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.activityGroupNames
    }
}
// Gets the azureSubscriptionId property value. 
func (m *DomainSecurityProfile) GetAzureSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureSubscriptionId
    }
}
// Gets the azureTenantId property value. 
func (m *DomainSecurityProfile) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// Gets the countHits property value. 
func (m *DomainSecurityProfile) GetCountHits()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.countHits
    }
}
// Gets the countInOrg property value. 
func (m *DomainSecurityProfile) GetCountInOrg()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.countInOrg
    }
}
// Gets the domainCategories property value. 
func (m *DomainSecurityProfile) GetDomainCategories()([]ReputationCategory) {
    if m == nil {
        return nil
    } else {
        return m.domainCategories
    }
}
// Gets the domainRegisteredDateTime property value. 
func (m *DomainSecurityProfile) GetDomainRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.domainRegisteredDateTime
    }
}
// Gets the firstSeenDateTime property value. 
func (m *DomainSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.firstSeenDateTime
    }
}
// Gets the lastSeenDateTime property value. 
func (m *DomainSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSeenDateTime
    }
}
// Gets the name property value. 
func (m *DomainSecurityProfile) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the registrant property value. 
func (m *DomainSecurityProfile) GetRegistrant()(*DomainRegistrant) {
    if m == nil {
        return nil
    } else {
        return m.registrant
    }
}
// Gets the riskScore property value. 
func (m *DomainSecurityProfile) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// Gets the tags property value. 
func (m *DomainSecurityProfile) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// Gets the vendorInformation property value. 
func (m *DomainSecurityProfile) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// The deserialization information for the current model
func (m *DomainSecurityProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["countInOrg"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountInOrg(val)
        }
        return nil
    }
    res["domainCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewReputationCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ReputationCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*ReputationCategory))
            }
            m.SetDomainCategories(res)
        }
        return nil
    }
    res["domainRegisteredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainRegisteredDateTime(val)
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
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["registrant"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDomainRegistrant() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrant(val.(*DomainRegistrant))
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
func (m *DomainSecurityProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DomainSecurityProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("countInOrg", m.GetCountInOrg())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDomainCategories()))
        for i, v := range m.GetDomainCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("domainCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("domainRegisteredDateTime", m.GetDomainRegisteredDateTime())
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
        err = writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("registrant", m.GetRegistrant())
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
func (m *DomainSecurityProfile) SetActivityGroupNames(value []string)() {
    m.activityGroupNames = value
}
// Sets the azureSubscriptionId property value. 
// Parameters:
//  - value : Value to set for the azureSubscriptionId property.
func (m *DomainSecurityProfile) SetAzureSubscriptionId(value *string)() {
    m.azureSubscriptionId = value
}
// Sets the azureTenantId property value. 
// Parameters:
//  - value : Value to set for the azureTenantId property.
func (m *DomainSecurityProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// Sets the countHits property value. 
// Parameters:
//  - value : Value to set for the countHits property.
func (m *DomainSecurityProfile) SetCountHits(value *int32)() {
    m.countHits = value
}
// Sets the countInOrg property value. 
// Parameters:
//  - value : Value to set for the countInOrg property.
func (m *DomainSecurityProfile) SetCountInOrg(value *int32)() {
    m.countInOrg = value
}
// Sets the domainCategories property value. 
// Parameters:
//  - value : Value to set for the domainCategories property.
func (m *DomainSecurityProfile) SetDomainCategories(value []ReputationCategory)() {
    m.domainCategories = value
}
// Sets the domainRegisteredDateTime property value. 
// Parameters:
//  - value : Value to set for the domainRegisteredDateTime property.
func (m *DomainSecurityProfile) SetDomainRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.domainRegisteredDateTime = value
}
// Sets the firstSeenDateTime property value. 
// Parameters:
//  - value : Value to set for the firstSeenDateTime property.
func (m *DomainSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// Sets the lastSeenDateTime property value. 
// Parameters:
//  - value : Value to set for the lastSeenDateTime property.
func (m *DomainSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *DomainSecurityProfile) SetName(value *string)() {
    m.name = value
}
// Sets the registrant property value. 
// Parameters:
//  - value : Value to set for the registrant property.
func (m *DomainSecurityProfile) SetRegistrant(value *DomainRegistrant)() {
    m.registrant = value
}
// Sets the riskScore property value. 
// Parameters:
//  - value : Value to set for the riskScore property.
func (m *DomainSecurityProfile) SetRiskScore(value *string)() {
    m.riskScore = value
}
// Sets the tags property value. 
// Parameters:
//  - value : Value to set for the tags property.
func (m *DomainSecurityProfile) SetTags(value []string)() {
    m.tags = value
}
// Sets the vendorInformation property value. 
// Parameters:
//  - value : Value to set for the vendorInformation property.
func (m *DomainSecurityProfile) SetVendorInformation(value *SecurityVendorInformation)() {
    m.vendorInformation = value
}
