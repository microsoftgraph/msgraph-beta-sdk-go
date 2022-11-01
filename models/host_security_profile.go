package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HostSecurityProfile provides operations to manage the collection of activityStatistics entities.
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
    odataTypeValue := "#microsoft.graph.hostSecurityProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateHostSecurityProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHostSecurityProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHostSecurityProfile(), nil
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *HostSecurityProfile) GetAzureSubscriptionId()(*string) {
    return m.azureSubscriptionId
}
// GetAzureTenantId gets the azureTenantId property value. The azureTenantId property
func (m *HostSecurityProfile) GetAzureTenantId()(*string) {
    return m.azureTenantId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HostSecurityProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureSubscriptionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureSubscriptionId)
    res["azureTenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureTenantId)
    res["firstSeenDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetFirstSeenDateTime)
    res["fqdn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFqdn)
    res["isAzureAdJoined"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAzureAdJoined)
    res["isAzureAdRegistered"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAzureAdRegistered)
    res["isHybridAzureDomainJoined"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsHybridAzureDomainJoined)
    res["lastSeenDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastSeenDateTime)
    res["logonUsers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLogonUserFromDiscriminatorValue , m.SetLogonUsers)
    res["netBiosName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNetBiosName)
    res["networkInterfaces"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateNetworkInterfaceFromDiscriminatorValue , m.SetNetworkInterfaces)
    res["os"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOs)
    res["osVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOsVersion)
    res["parentHost"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetParentHost)
    res["relatedHostIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRelatedHostIds)
    res["riskScore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRiskScore)
    res["tags"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTags)
    res["vendorInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSecurityVendorInformationFromDiscriminatorValue , m.SetVendorInformation)
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *HostSecurityProfile) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetFqdn gets the fqdn property value. The fqdn property
func (m *HostSecurityProfile) GetFqdn()(*string) {
    return m.fqdn
}
// GetIsAzureAdJoined gets the isAzureAdJoined property value. The isAzureAdJoined property
func (m *HostSecurityProfile) GetIsAzureAdJoined()(*bool) {
    return m.isAzureAdJoined
}
// GetIsAzureAdRegistered gets the isAzureAdRegistered property value. The isAzureAdRegistered property
func (m *HostSecurityProfile) GetIsAzureAdRegistered()(*bool) {
    return m.isAzureAdRegistered
}
// GetIsHybridAzureDomainJoined gets the isHybridAzureDomainJoined property value. The isHybridAzureDomainJoined property
func (m *HostSecurityProfile) GetIsHybridAzureDomainJoined()(*bool) {
    return m.isHybridAzureDomainJoined
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *HostSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSeenDateTime
}
// GetLogonUsers gets the logonUsers property value. The logonUsers property
func (m *HostSecurityProfile) GetLogonUsers()([]LogonUserable) {
    return m.logonUsers
}
// GetNetBiosName gets the netBiosName property value. The netBiosName property
func (m *HostSecurityProfile) GetNetBiosName()(*string) {
    return m.netBiosName
}
// GetNetworkInterfaces gets the networkInterfaces property value. The networkInterfaces property
func (m *HostSecurityProfile) GetNetworkInterfaces()([]NetworkInterfaceable) {
    return m.networkInterfaces
}
// GetOs gets the os property value. The os property
func (m *HostSecurityProfile) GetOs()(*string) {
    return m.os
}
// GetOsVersion gets the osVersion property value. The osVersion property
func (m *HostSecurityProfile) GetOsVersion()(*string) {
    return m.osVersion
}
// GetParentHost gets the parentHost property value. The parentHost property
func (m *HostSecurityProfile) GetParentHost()(*string) {
    return m.parentHost
}
// GetRelatedHostIds gets the relatedHostIds property value. The relatedHostIds property
func (m *HostSecurityProfile) GetRelatedHostIds()([]string) {
    return m.relatedHostIds
}
// GetRiskScore gets the riskScore property value. The riskScore property
func (m *HostSecurityProfile) GetRiskScore()(*string) {
    return m.riskScore
}
// GetTags gets the tags property value. The tags property
func (m *HostSecurityProfile) GetTags()([]string) {
    return m.tags
}
// GetVendorInformation gets the vendorInformation property value. The vendorInformation property
func (m *HostSecurityProfile) GetVendorInformation()(SecurityVendorInformationable) {
    return m.vendorInformation
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLogonUsers())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetNetworkInterfaces())
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
    m.azureSubscriptionId = value
}
// SetAzureTenantId sets the azureTenantId property value. The azureTenantId property
func (m *HostSecurityProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *HostSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetFqdn sets the fqdn property value. The fqdn property
func (m *HostSecurityProfile) SetFqdn(value *string)() {
    m.fqdn = value
}
// SetIsAzureAdJoined sets the isAzureAdJoined property value. The isAzureAdJoined property
func (m *HostSecurityProfile) SetIsAzureAdJoined(value *bool)() {
    m.isAzureAdJoined = value
}
// SetIsAzureAdRegistered sets the isAzureAdRegistered property value. The isAzureAdRegistered property
func (m *HostSecurityProfile) SetIsAzureAdRegistered(value *bool)() {
    m.isAzureAdRegistered = value
}
// SetIsHybridAzureDomainJoined sets the isHybridAzureDomainJoined property value. The isHybridAzureDomainJoined property
func (m *HostSecurityProfile) SetIsHybridAzureDomainJoined(value *bool)() {
    m.isHybridAzureDomainJoined = value
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *HostSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
// SetLogonUsers sets the logonUsers property value. The logonUsers property
func (m *HostSecurityProfile) SetLogonUsers(value []LogonUserable)() {
    m.logonUsers = value
}
// SetNetBiosName sets the netBiosName property value. The netBiosName property
func (m *HostSecurityProfile) SetNetBiosName(value *string)() {
    m.netBiosName = value
}
// SetNetworkInterfaces sets the networkInterfaces property value. The networkInterfaces property
func (m *HostSecurityProfile) SetNetworkInterfaces(value []NetworkInterfaceable)() {
    m.networkInterfaces = value
}
// SetOs sets the os property value. The os property
func (m *HostSecurityProfile) SetOs(value *string)() {
    m.os = value
}
// SetOsVersion sets the osVersion property value. The osVersion property
func (m *HostSecurityProfile) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetParentHost sets the parentHost property value. The parentHost property
func (m *HostSecurityProfile) SetParentHost(value *string)() {
    m.parentHost = value
}
// SetRelatedHostIds sets the relatedHostIds property value. The relatedHostIds property
func (m *HostSecurityProfile) SetRelatedHostIds(value []string)() {
    m.relatedHostIds = value
}
// SetRiskScore sets the riskScore property value. The riskScore property
func (m *HostSecurityProfile) SetRiskScore(value *string)() {
    m.riskScore = value
}
// SetTags sets the tags property value. The tags property
func (m *HostSecurityProfile) SetTags(value []string)() {
    m.tags = value
}
// SetVendorInformation sets the vendorInformation property value. The vendorInformation property
func (m *HostSecurityProfile) SetVendorInformation(value SecurityVendorInformationable)() {
    m.vendorInformation = value
}
