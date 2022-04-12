package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HostSecurityProfile 
type HostSecurityProfile struct {
    Entity
    // The azureSubscriptionId property
    azureSubscriptionId *string
    // The azureTenantId property
    azureTenantId *string
    // The firstSeenDateTime property
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The fqdn property
    fqdn *string
    // The isAzureAdJoined property
    isAzureAdJoined *bool
    // The isAzureAdRegistered property
    isAzureAdRegistered *bool
    // The isHybridAzureDomainJoined property
    isHybridAzureDomainJoined *bool
    // The lastSeenDateTime property
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The logonUsers property
    logonUsers []LogonUserable
    // The netBiosName property
    netBiosName *string
    // The networkInterfaces property
    networkInterfaces []NetworkInterfaceable
    // The os property
    os *string
    // The osVersion property
    osVersion *string
    // The parentHost property
    parentHost *string
    // The relatedHostIds property
    relatedHostIds []string
    // The riskScore property
    riskScore *string
    // The tags property
    tags []string
    // The vendorInformation property
    vendorInformation SecurityVendorInformationable
}
// NewHostSecurityProfile instantiates a new hostSecurityProfile and sets the default values.
func NewHostSecurityProfile()(*HostSecurityProfile) {
    m := &HostSecurityProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateHostSecurityProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHostSecurityProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHostSecurityProfile(), nil
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *HostSecurityProfile) GetAzureSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureSubscriptionId
    }
}
// GetAzureTenantId gets the azureTenantId property value. The azureTenantId property
func (m *HostSecurityProfile) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HostSecurityProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureSubscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureSubscriptionId(val)
        }
        return nil
    }
    res["azureTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["firstSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstSeenDateTime(val)
        }
        return nil
    }
    res["fqdn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFqdn(val)
        }
        return nil
    }
    res["isAzureAdJoined"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAzureAdJoined(val)
        }
        return nil
    }
    res["isAzureAdRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAzureAdRegistered(val)
        }
        return nil
    }
    res["isHybridAzureDomainJoined"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHybridAzureDomainJoined(val)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
        }
        return nil
    }
    res["logonUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLogonUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LogonUserable, len(val))
            for i, v := range val {
                res[i] = v.(LogonUserable)
            }
            m.SetLogonUsers(res)
        }
        return nil
    }
    res["netBiosName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetBiosName(val)
        }
        return nil
    }
    res["networkInterfaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNetworkInterfaceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NetworkInterfaceable, len(val))
            for i, v := range val {
                res[i] = v.(NetworkInterfaceable)
            }
            m.SetNetworkInterfaces(res)
        }
        return nil
    }
    res["os"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOs(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["parentHost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentHost(val)
        }
        return nil
    }
    res["relatedHostIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["riskScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["vendorInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityVendorInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInformation(val.(SecurityVendorInformationable))
        }
        return nil
    }
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *HostSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.firstSeenDateTime
    }
}
// GetFqdn gets the fqdn property value. The fqdn property
func (m *HostSecurityProfile) GetFqdn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fqdn
    }
}
// GetIsAzureAdJoined gets the isAzureAdJoined property value. The isAzureAdJoined property
func (m *HostSecurityProfile) GetIsAzureAdJoined()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAzureAdJoined
    }
}
// GetIsAzureAdRegistered gets the isAzureAdRegistered property value. The isAzureAdRegistered property
func (m *HostSecurityProfile) GetIsAzureAdRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAzureAdRegistered
    }
}
// GetIsHybridAzureDomainJoined gets the isHybridAzureDomainJoined property value. The isHybridAzureDomainJoined property
func (m *HostSecurityProfile) GetIsHybridAzureDomainJoined()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHybridAzureDomainJoined
    }
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *HostSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSeenDateTime
    }
}
// GetLogonUsers gets the logonUsers property value. The logonUsers property
func (m *HostSecurityProfile) GetLogonUsers()([]LogonUserable) {
    if m == nil {
        return nil
    } else {
        return m.logonUsers
    }
}
// GetNetBiosName gets the netBiosName property value. The netBiosName property
func (m *HostSecurityProfile) GetNetBiosName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.netBiosName
    }
}
// GetNetworkInterfaces gets the networkInterfaces property value. The networkInterfaces property
func (m *HostSecurityProfile) GetNetworkInterfaces()([]NetworkInterfaceable) {
    if m == nil {
        return nil
    } else {
        return m.networkInterfaces
    }
}
// GetOs gets the os property value. The os property
func (m *HostSecurityProfile) GetOs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.os
    }
}
// GetOsVersion gets the osVersion property value. The osVersion property
func (m *HostSecurityProfile) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetParentHost gets the parentHost property value. The parentHost property
func (m *HostSecurityProfile) GetParentHost()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentHost
    }
}
// GetRelatedHostIds gets the relatedHostIds property value. The relatedHostIds property
func (m *HostSecurityProfile) GetRelatedHostIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.relatedHostIds
    }
}
// GetRiskScore gets the riskScore property value. The riskScore property
func (m *HostSecurityProfile) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// GetTags gets the tags property value. The tags property
func (m *HostSecurityProfile) GetTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tags
    }
}
// GetVendorInformation gets the vendorInformation property value. The vendorInformation property
func (m *HostSecurityProfile) GetVendorInformation()(SecurityVendorInformationable) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// Serialize serializes information the current object
func (m *HostSecurityProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLogonUsers()))
        for i, v := range m.GetLogonUsers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkInterfaces()))
        for i, v := range m.GetNetworkInterfaces() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
// SetAzureSubscriptionId sets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *HostSecurityProfile) SetAzureSubscriptionId(value *string)() {
    if m != nil {
        m.azureSubscriptionId = value
    }
}
// SetAzureTenantId sets the azureTenantId property value. The azureTenantId property
func (m *HostSecurityProfile) SetAzureTenantId(value *string)() {
    if m != nil {
        m.azureTenantId = value
    }
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *HostSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.firstSeenDateTime = value
    }
}
// SetFqdn sets the fqdn property value. The fqdn property
func (m *HostSecurityProfile) SetFqdn(value *string)() {
    if m != nil {
        m.fqdn = value
    }
}
// SetIsAzureAdJoined sets the isAzureAdJoined property value. The isAzureAdJoined property
func (m *HostSecurityProfile) SetIsAzureAdJoined(value *bool)() {
    if m != nil {
        m.isAzureAdJoined = value
    }
}
// SetIsAzureAdRegistered sets the isAzureAdRegistered property value. The isAzureAdRegistered property
func (m *HostSecurityProfile) SetIsAzureAdRegistered(value *bool)() {
    if m != nil {
        m.isAzureAdRegistered = value
    }
}
// SetIsHybridAzureDomainJoined sets the isHybridAzureDomainJoined property value. The isHybridAzureDomainJoined property
func (m *HostSecurityProfile) SetIsHybridAzureDomainJoined(value *bool)() {
    if m != nil {
        m.isHybridAzureDomainJoined = value
    }
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *HostSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSeenDateTime = value
    }
}
// SetLogonUsers sets the logonUsers property value. The logonUsers property
func (m *HostSecurityProfile) SetLogonUsers(value []LogonUserable)() {
    if m != nil {
        m.logonUsers = value
    }
}
// SetNetBiosName sets the netBiosName property value. The netBiosName property
func (m *HostSecurityProfile) SetNetBiosName(value *string)() {
    if m != nil {
        m.netBiosName = value
    }
}
// SetNetworkInterfaces sets the networkInterfaces property value. The networkInterfaces property
func (m *HostSecurityProfile) SetNetworkInterfaces(value []NetworkInterfaceable)() {
    if m != nil {
        m.networkInterfaces = value
    }
}
// SetOs sets the os property value. The os property
func (m *HostSecurityProfile) SetOs(value *string)() {
    if m != nil {
        m.os = value
    }
}
// SetOsVersion sets the osVersion property value. The osVersion property
func (m *HostSecurityProfile) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetParentHost sets the parentHost property value. The parentHost property
func (m *HostSecurityProfile) SetParentHost(value *string)() {
    if m != nil {
        m.parentHost = value
    }
}
// SetRelatedHostIds sets the relatedHostIds property value. The relatedHostIds property
func (m *HostSecurityProfile) SetRelatedHostIds(value []string)() {
    if m != nil {
        m.relatedHostIds = value
    }
}
// SetRiskScore sets the riskScore property value. The riskScore property
func (m *HostSecurityProfile) SetRiskScore(value *string)() {
    if m != nil {
        m.riskScore = value
    }
}
// SetTags sets the tags property value. The tags property
func (m *HostSecurityProfile) SetTags(value []string)() {
    if m != nil {
        m.tags = value
    }
}
// SetVendorInformation sets the vendorInformation property value. The vendorInformation property
func (m *HostSecurityProfile) SetVendorInformation(value SecurityVendorInformationable)() {
    if m != nil {
        m.vendorInformation = value
    }
}
