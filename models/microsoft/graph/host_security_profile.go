package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HostSecurityProfile 
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
// NewHostSecurityProfile instantiates a new hostSecurityProfile and sets the default values.
func NewHostSecurityProfile()(*HostSecurityProfile) {
    m := &HostSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. 
func (m *HostSecurityProfile) GetAzureSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureSubscriptionId
    }
}
// GetAzureTenantId gets the azureTenantId property value. 
func (m *HostSecurityProfile) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. 
func (m *HostSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.firstSeenDateTime
    }
}
// GetFqdn gets the fqdn property value. 
func (m *HostSecurityProfile) GetFqdn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fqdn
    }
}
// GetIsAzureAdJoined gets the isAzureAdJoined property value. 
func (m *HostSecurityProfile) GetIsAzureAdJoined()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAzureAdJoined
    }
}
// GetIsAzureAdRegistered gets the isAzureAdRegistered property value. 
func (m *HostSecurityProfile) GetIsAzureAdRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAzureAdRegistered
    }
}
// GetIsHybridAzureDomainJoined gets the isHybridAzureDomainJoined property value. 
func (m *HostSecurityProfile) GetIsHybridAzureDomainJoined()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHybridAzureDomainJoined
    }
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. 
func (m *HostSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSeenDateTime
    }
}
// GetLogonUsers gets the logonUsers property value. 
func (m *HostSecurityProfile) GetLogonUsers()([]LogonUser) {
    if m == nil {
        return nil
    } else {
        return m.logonUsers
    }
}
// GetNetBiosName gets the netBiosName property value. 
func (m *HostSecurityProfile) GetNetBiosName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.netBiosName
    }
}
// GetNetworkInterfaces gets the networkInterfaces property value. 
func (m *HostSecurityProfile) GetNetworkInterfaces()([]NetworkInterface) {
    if m == nil {
        return nil
    } else {
        return m.networkInterfaces
    }
}
// GetOs gets the os property value. 
func (m *HostSecurityProfile) GetOs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.os
    }
}
// GetOsVersion gets the osVersion property value. 
func (m *HostSecurityProfile) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetParentHost gets the parentHost property value. 
func (m *HostSecurityProfile) GetParentHost()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentHost
    }
}
// GetRelatedHostIds gets the relatedHostIds property value. 
func (m *HostSecurityProfile) GetRelatedHostIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.relatedHostIds
    }
}
// GetRiskScore gets the riskScore property value. 
func (m *HostSecurityProfile) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// GetTags gets the tags property value. 
func (m *HostSecurityProfile) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// GetVendorInformation gets the vendorInformation property value. 
func (m *HostSecurityProfile) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HostSecurityProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["fqdn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFqdn(val)
        }
        return nil
    }
    res["isAzureAdJoined"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAzureAdJoined(val)
        }
        return nil
    }
    res["isAzureAdRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAzureAdRegistered(val)
        }
        return nil
    }
    res["isHybridAzureDomainJoined"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHybridAzureDomainJoined(val)
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
    res["logonUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLogonUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LogonUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*LogonUser))
            }
            m.SetLogonUsers(res)
        }
        return nil
    }
    res["netBiosName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetBiosName(val)
        }
        return nil
    }
    res["networkInterfaces"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNetworkInterface() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NetworkInterface, len(val))
            for i, v := range val {
                res[i] = *(v.(*NetworkInterface))
            }
            m.SetNetworkInterfaces(res)
        }
        return nil
    }
    res["os"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOs(val)
        }
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["parentHost"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentHost(val)
        }
        return nil
    }
    res["relatedHostIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRelatedHostIds(res)
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
func (m *HostSecurityProfile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetLogonUsers() != nil {
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
    if m.GetNetworkInterfaces() != nil {
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
    if m.GetRelatedHostIds() != nil {
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
    if m.GetTags() != nil {
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
// SetAzureSubscriptionId sets the azureSubscriptionId property value. 
func (m *HostSecurityProfile) SetAzureSubscriptionId(value *string)() {
    if m != nil {
        m.azureSubscriptionId = value
    }
}
// SetAzureTenantId sets the azureTenantId property value. 
func (m *HostSecurityProfile) SetAzureTenantId(value *string)() {
    if m != nil {
        m.azureTenantId = value
    }
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. 
func (m *HostSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.firstSeenDateTime = value
    }
}
// SetFqdn sets the fqdn property value. 
func (m *HostSecurityProfile) SetFqdn(value *string)() {
    if m != nil {
        m.fqdn = value
    }
}
// SetIsAzureAdJoined sets the isAzureAdJoined property value. 
func (m *HostSecurityProfile) SetIsAzureAdJoined(value *bool)() {
    if m != nil {
        m.isAzureAdJoined = value
    }
}
// SetIsAzureAdRegistered sets the isAzureAdRegistered property value. 
func (m *HostSecurityProfile) SetIsAzureAdRegistered(value *bool)() {
    if m != nil {
        m.isAzureAdRegistered = value
    }
}
// SetIsHybridAzureDomainJoined sets the isHybridAzureDomainJoined property value. 
func (m *HostSecurityProfile) SetIsHybridAzureDomainJoined(value *bool)() {
    if m != nil {
        m.isHybridAzureDomainJoined = value
    }
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. 
func (m *HostSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSeenDateTime = value
    }
}
// SetLogonUsers sets the logonUsers property value. 
func (m *HostSecurityProfile) SetLogonUsers(value []LogonUser)() {
    if m != nil {
        m.logonUsers = value
    }
}
// SetNetBiosName sets the netBiosName property value. 
func (m *HostSecurityProfile) SetNetBiosName(value *string)() {
    if m != nil {
        m.netBiosName = value
    }
}
// SetNetworkInterfaces sets the networkInterfaces property value. 
func (m *HostSecurityProfile) SetNetworkInterfaces(value []NetworkInterface)() {
    if m != nil {
        m.networkInterfaces = value
    }
}
// SetOs sets the os property value. 
func (m *HostSecurityProfile) SetOs(value *string)() {
    if m != nil {
        m.os = value
    }
}
// SetOsVersion sets the osVersion property value. 
func (m *HostSecurityProfile) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetParentHost sets the parentHost property value. 
func (m *HostSecurityProfile) SetParentHost(value *string)() {
    if m != nil {
        m.parentHost = value
    }
}
// SetRelatedHostIds sets the relatedHostIds property value. 
func (m *HostSecurityProfile) SetRelatedHostIds(value []string)() {
    if m != nil {
        m.relatedHostIds = value
    }
}
// SetRiskScore sets the riskScore property value. 
func (m *HostSecurityProfile) SetRiskScore(value *string)() {
    if m != nil {
        m.riskScore = value
    }
}
// SetTags sets the tags property value. 
func (m *HostSecurityProfile) SetTags(value []string)() {
    if m != nil {
        m.tags = value
    }
}
// SetVendorInformation sets the vendorInformation property value. 
func (m *HostSecurityProfile) SetVendorInformation(value *SecurityVendorInformation)() {
    if m != nil {
        m.vendorInformation = value
    }
}
