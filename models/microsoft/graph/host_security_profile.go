package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type HostSecurityProfile struct {
    Entity
    // 
    azureSubscriptionId *string;
    // 
    azureTenantId *string;
    // 
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    fqdn *string;
    // 
    isAzureAdJoined *bool;
    // 
    isAzureAdRegistered *bool;
    // 
    isHybridAzureDomainJoined *bool;
    // 
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    logonUsers []LogonUser;
    // 
    netBiosName *string;
    // 
    networkInterfaces []NetworkInterface;
    // 
    os *string;
    // 
    osVersion *string;
    // 
    parentHost *string;
    // 
    relatedHostIds []string;
    // 
    riskScore *string;
    // 
    tags []string;
    // 
    vendorInformation *SecurityVendorInformation;
}
// Instantiates a new hostSecurityProfile and sets the default values.
func NewHostSecurityProfile()(*HostSecurityProfile) {
    m := &HostSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the azureSubscriptionId property value. 
func (m *HostSecurityProfile) GetAzureSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureSubscriptionId
    }
}
// Gets the azureTenantId property value. 
func (m *HostSecurityProfile) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// Gets the firstSeenDateTime property value. 
func (m *HostSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.firstSeenDateTime
    }
}
// Gets the fqdn property value. 
func (m *HostSecurityProfile) GetFqdn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fqdn
    }
}
// Gets the isAzureAdJoined property value. 
func (m *HostSecurityProfile) GetIsAzureAdJoined()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAzureAdJoined
    }
}
// Gets the isAzureAdRegistered property value. 
func (m *HostSecurityProfile) GetIsAzureAdRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAzureAdRegistered
    }
}
// Gets the isHybridAzureDomainJoined property value. 
func (m *HostSecurityProfile) GetIsHybridAzureDomainJoined()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHybridAzureDomainJoined
    }
}
// Gets the lastSeenDateTime property value. 
func (m *HostSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSeenDateTime
    }
}
// Gets the logonUsers property value. 
func (m *HostSecurityProfile) GetLogonUsers()([]LogonUser) {
    if m == nil {
        return nil
    } else {
        return m.logonUsers
    }
}
// Gets the netBiosName property value. 
func (m *HostSecurityProfile) GetNetBiosName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.netBiosName
    }
}
// Gets the networkInterfaces property value. 
func (m *HostSecurityProfile) GetNetworkInterfaces()([]NetworkInterface) {
    if m == nil {
        return nil
    } else {
        return m.networkInterfaces
    }
}
// Gets the os property value. 
func (m *HostSecurityProfile) GetOs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.os
    }
}
// Gets the osVersion property value. 
func (m *HostSecurityProfile) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// Gets the parentHost property value. 
func (m *HostSecurityProfile) GetParentHost()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentHost
    }
}
// Gets the relatedHostIds property value. 
func (m *HostSecurityProfile) GetRelatedHostIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.relatedHostIds
    }
}
// Gets the riskScore property value. 
func (m *HostSecurityProfile) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// Gets the tags property value. 
func (m *HostSecurityProfile) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// Gets the vendorInformation property value. 
func (m *HostSecurityProfile) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// The deserialization information for the current model
func (m *HostSecurityProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureSubscriptionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureSubscriptionId(val)
        return nil
    }
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureTenantId(val)
        return nil
    }
    res["firstSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetFirstSeenDateTime(val)
        return nil
    }
    res["fqdn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFqdn(val)
        return nil
    }
    res["isAzureAdJoined"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAzureAdJoined(val)
        return nil
    }
    res["isAzureAdRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAzureAdRegistered(val)
        return nil
    }
    res["isHybridAzureDomainJoined"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsHybridAzureDomainJoined(val)
        return nil
    }
    res["lastSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSeenDateTime(val)
        return nil
    }
    res["logonUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLogonUser() })
        if err != nil {
            return err
        }
        res := make([]LogonUser, len(val))
        for i, v := range val {
            res[i] = *(v.(*LogonUser))
        }
        m.SetLogonUsers(res)
        return nil
    }
    res["netBiosName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetBiosName(val)
        return nil
    }
    res["networkInterfaces"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNetworkInterface() })
        if err != nil {
            return err
        }
        res := make([]NetworkInterface, len(val))
        for i, v := range val {
            res[i] = *(v.(*NetworkInterface))
        }
        m.SetNetworkInterfaces(res)
        return nil
    }
    res["os"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOs(val)
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsVersion(val)
        return nil
    }
    res["parentHost"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParentHost(val)
        return nil
    }
    res["relatedHostIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRelatedHostIds(res)
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRiskScore(val)
        return nil
    }
    res["tags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetTags(res)
        return nil
    }
    res["vendorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityVendorInformation() })
        if err != nil {
            return err
        }
        m.SetVendorInformation(val.(*SecurityVendorInformation))
        return nil
    }
    return res
}
func (m *HostSecurityProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *HostSecurityProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteTimeValue("firstSeenDateTime", m.GetFirstSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fqdn", m.GetFqdn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAzureAdJoined", m.GetIsAzureAdJoined())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAzureAdRegistered", m.GetIsAzureAdRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isHybridAzureDomainJoined", m.GetIsHybridAzureDomainJoined())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLogonUsers()))
        for i, v := range m.GetLogonUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("logonUsers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("netBiosName", m.GetNetBiosName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNetworkInterfaces()))
        for i, v := range m.GetNetworkInterfaces() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("networkInterfaces", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("os", m.GetOs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentHost", m.GetParentHost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("relatedHostIds", m.GetRelatedHostIds())
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
// Sets the azureSubscriptionId property value. 
// Parameters:
//  - value : Value to set for the azureSubscriptionId property.
func (m *HostSecurityProfile) SetAzureSubscriptionId(value *string)() {
    m.azureSubscriptionId = value
}
// Sets the azureTenantId property value. 
// Parameters:
//  - value : Value to set for the azureTenantId property.
func (m *HostSecurityProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// Sets the firstSeenDateTime property value. 
// Parameters:
//  - value : Value to set for the firstSeenDateTime property.
func (m *HostSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// Sets the fqdn property value. 
// Parameters:
//  - value : Value to set for the fqdn property.
func (m *HostSecurityProfile) SetFqdn(value *string)() {
    m.fqdn = value
}
// Sets the isAzureAdJoined property value. 
// Parameters:
//  - value : Value to set for the isAzureAdJoined property.
func (m *HostSecurityProfile) SetIsAzureAdJoined(value *bool)() {
    m.isAzureAdJoined = value
}
// Sets the isAzureAdRegistered property value. 
// Parameters:
//  - value : Value to set for the isAzureAdRegistered property.
func (m *HostSecurityProfile) SetIsAzureAdRegistered(value *bool)() {
    m.isAzureAdRegistered = value
}
// Sets the isHybridAzureDomainJoined property value. 
// Parameters:
//  - value : Value to set for the isHybridAzureDomainJoined property.
func (m *HostSecurityProfile) SetIsHybridAzureDomainJoined(value *bool)() {
    m.isHybridAzureDomainJoined = value
}
// Sets the lastSeenDateTime property value. 
// Parameters:
//  - value : Value to set for the lastSeenDateTime property.
func (m *HostSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// Sets the logonUsers property value. 
// Parameters:
//  - value : Value to set for the logonUsers property.
func (m *HostSecurityProfile) SetLogonUsers(value []LogonUser)() {
    m.logonUsers = value
}
// Sets the netBiosName property value. 
// Parameters:
//  - value : Value to set for the netBiosName property.
func (m *HostSecurityProfile) SetNetBiosName(value *string)() {
    m.netBiosName = value
}
// Sets the networkInterfaces property value. 
// Parameters:
//  - value : Value to set for the networkInterfaces property.
func (m *HostSecurityProfile) SetNetworkInterfaces(value []NetworkInterface)() {
    m.networkInterfaces = value
}
// Sets the os property value. 
// Parameters:
//  - value : Value to set for the os property.
func (m *HostSecurityProfile) SetOs(value *string)() {
    m.os = value
}
// Sets the osVersion property value. 
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *HostSecurityProfile) SetOsVersion(value *string)() {
    m.osVersion = value
}
// Sets the parentHost property value. 
// Parameters:
//  - value : Value to set for the parentHost property.
func (m *HostSecurityProfile) SetParentHost(value *string)() {
    m.parentHost = value
}
// Sets the relatedHostIds property value. 
// Parameters:
//  - value : Value to set for the relatedHostIds property.
func (m *HostSecurityProfile) SetRelatedHostIds(value []string)() {
    m.relatedHostIds = value
}
// Sets the riskScore property value. 
// Parameters:
//  - value : Value to set for the riskScore property.
func (m *HostSecurityProfile) SetRiskScore(value *string)() {
    m.riskScore = value
}
// Sets the tags property value. 
// Parameters:
//  - value : Value to set for the tags property.
func (m *HostSecurityProfile) SetTags(value []string)() {
    m.tags = value
}
// Sets the vendorInformation property value. 
// Parameters:
//  - value : Value to set for the vendorInformation property.
func (m *HostSecurityProfile) SetVendorInformation(value *SecurityVendorInformation)() {
    m.vendorInformation = value
}
