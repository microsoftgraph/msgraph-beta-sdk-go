package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HostSecurityProfile 
type HostSecurityProfile struct {
    Entity
    // The OdataType property
    OdataType *string
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
    val, err := m.GetBackingStore().Get("azureSubscriptionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureTenantId gets the azureTenantId property value. The azureTenantId property
func (m *HostSecurityProfile) GetAzureTenantId()(*string) {
    val, err := m.GetBackingStore().Get("azureTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
                if v != nil {
                    res[i] = v.(LogonUserable)
                }
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
                if v != nil {
                    res[i] = v.(NetworkInterfaceable)
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
    val, err := m.GetBackingStore().Get("firstSeenDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFqdn gets the fqdn property value. The fqdn property
func (m *HostSecurityProfile) GetFqdn()(*string) {
    val, err := m.GetBackingStore().Get("fqdn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsAzureAdJoined gets the isAzureAdJoined property value. The isAzureAdJoined property
func (m *HostSecurityProfile) GetIsAzureAdJoined()(*bool) {
    val, err := m.GetBackingStore().Get("isAzureAdJoined")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsAzureAdRegistered gets the isAzureAdRegistered property value. The isAzureAdRegistered property
func (m *HostSecurityProfile) GetIsAzureAdRegistered()(*bool) {
    val, err := m.GetBackingStore().Get("isAzureAdRegistered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsHybridAzureDomainJoined gets the isHybridAzureDomainJoined property value. The isHybridAzureDomainJoined property
func (m *HostSecurityProfile) GetIsHybridAzureDomainJoined()(*bool) {
    val, err := m.GetBackingStore().Get("isHybridAzureDomainJoined")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *HostSecurityProfile) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastSeenDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLogonUsers gets the logonUsers property value. The logonUsers property
func (m *HostSecurityProfile) GetLogonUsers()([]LogonUserable) {
    val, err := m.GetBackingStore().Get("logonUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]LogonUserable)
    }
    return nil
}
// GetNetBiosName gets the netBiosName property value. The netBiosName property
func (m *HostSecurityProfile) GetNetBiosName()(*string) {
    val, err := m.GetBackingStore().Get("netBiosName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNetworkInterfaces gets the networkInterfaces property value. The networkInterfaces property
func (m *HostSecurityProfile) GetNetworkInterfaces()([]NetworkInterfaceable) {
    val, err := m.GetBackingStore().Get("networkInterfaces")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]NetworkInterfaceable)
    }
    return nil
}
// GetOs gets the os property value. The os property
func (m *HostSecurityProfile) GetOs()(*string) {
    val, err := m.GetBackingStore().Get("os")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. The osVersion property
func (m *HostSecurityProfile) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetParentHost gets the parentHost property value. The parentHost property
func (m *HostSecurityProfile) GetParentHost()(*string) {
    val, err := m.GetBackingStore().Get("parentHost")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRelatedHostIds gets the relatedHostIds property value. The relatedHostIds property
func (m *HostSecurityProfile) GetRelatedHostIds()([]string) {
    val, err := m.GetBackingStore().Get("relatedHostIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRiskScore gets the riskScore property value. The riskScore property
func (m *HostSecurityProfile) GetRiskScore()(*string) {
    val, err := m.GetBackingStore().Get("riskScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTags gets the tags property value. The tags property
func (m *HostSecurityProfile) GetTags()([]string) {
    val, err := m.GetBackingStore().Get("tags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetVendorInformation gets the vendorInformation property value. The vendorInformation property
func (m *HostSecurityProfile) GetVendorInformation()(SecurityVendorInformationable) {
    val, err := m.GetBackingStore().Get("vendorInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SecurityVendorInformationable)
    }
    return nil
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    err := m.GetBackingStore().Set("azureSubscriptionId", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureTenantId sets the azureTenantId property value. The azureTenantId property
func (m *HostSecurityProfile) SetAzureTenantId(value *string)() {
    err := m.GetBackingStore().Set("azureTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *HostSecurityProfile) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("firstSeenDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFqdn sets the fqdn property value. The fqdn property
func (m *HostSecurityProfile) SetFqdn(value *string)() {
    err := m.GetBackingStore().Set("fqdn", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAzureAdJoined sets the isAzureAdJoined property value. The isAzureAdJoined property
func (m *HostSecurityProfile) SetIsAzureAdJoined(value *bool)() {
    err := m.GetBackingStore().Set("isAzureAdJoined", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAzureAdRegistered sets the isAzureAdRegistered property value. The isAzureAdRegistered property
func (m *HostSecurityProfile) SetIsAzureAdRegistered(value *bool)() {
    err := m.GetBackingStore().Set("isAzureAdRegistered", value)
    if err != nil {
        panic(err)
    }
}
// SetIsHybridAzureDomainJoined sets the isHybridAzureDomainJoined property value. The isHybridAzureDomainJoined property
func (m *HostSecurityProfile) SetIsHybridAzureDomainJoined(value *bool)() {
    err := m.GetBackingStore().Set("isHybridAzureDomainJoined", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. The lastSeenDateTime property
func (m *HostSecurityProfile) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastSeenDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLogonUsers sets the logonUsers property value. The logonUsers property
func (m *HostSecurityProfile) SetLogonUsers(value []LogonUserable)() {
    err := m.GetBackingStore().Set("logonUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetNetBiosName sets the netBiosName property value. The netBiosName property
func (m *HostSecurityProfile) SetNetBiosName(value *string)() {
    err := m.GetBackingStore().Set("netBiosName", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkInterfaces sets the networkInterfaces property value. The networkInterfaces property
func (m *HostSecurityProfile) SetNetworkInterfaces(value []NetworkInterfaceable)() {
    err := m.GetBackingStore().Set("networkInterfaces", value)
    if err != nil {
        panic(err)
    }
}
// SetOs sets the os property value. The os property
func (m *HostSecurityProfile) SetOs(value *string)() {
    err := m.GetBackingStore().Set("os", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. The osVersion property
func (m *HostSecurityProfile) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetParentHost sets the parentHost property value. The parentHost property
func (m *HostSecurityProfile) SetParentHost(value *string)() {
    err := m.GetBackingStore().Set("parentHost", value)
    if err != nil {
        panic(err)
    }
}
// SetRelatedHostIds sets the relatedHostIds property value. The relatedHostIds property
func (m *HostSecurityProfile) SetRelatedHostIds(value []string)() {
    err := m.GetBackingStore().Set("relatedHostIds", value)
    if err != nil {
        panic(err)
    }
}
// SetRiskScore sets the riskScore property value. The riskScore property
func (m *HostSecurityProfile) SetRiskScore(value *string)() {
    err := m.GetBackingStore().Set("riskScore", value)
    if err != nil {
        panic(err)
    }
}
// SetTags sets the tags property value. The tags property
func (m *HostSecurityProfile) SetTags(value []string)() {
    err := m.GetBackingStore().Set("tags", value)
    if err != nil {
        panic(err)
    }
}
// SetVendorInformation sets the vendorInformation property value. The vendorInformation property
func (m *HostSecurityProfile) SetVendorInformation(value SecurityVendorInformationable)() {
    err := m.GetBackingStore().Set("vendorInformation", value)
    if err != nil {
        panic(err)
    }
}
// HostSecurityProfileable 
type HostSecurityProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureSubscriptionId()(*string)
    GetAzureTenantId()(*string)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFqdn()(*string)
    GetIsAzureAdJoined()(*bool)
    GetIsAzureAdRegistered()(*bool)
    GetIsHybridAzureDomainJoined()(*bool)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLogonUsers()([]LogonUserable)
    GetNetBiosName()(*string)
    GetNetworkInterfaces()([]NetworkInterfaceable)
    GetOs()(*string)
    GetOsVersion()(*string)
    GetParentHost()(*string)
    GetRelatedHostIds()([]string)
    GetRiskScore()(*string)
    GetTags()([]string)
    GetVendorInformation()(SecurityVendorInformationable)
    SetAzureSubscriptionId(value *string)()
    SetAzureTenantId(value *string)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFqdn(value *string)()
    SetIsAzureAdJoined(value *bool)()
    SetIsAzureAdRegistered(value *bool)()
    SetIsHybridAzureDomainJoined(value *bool)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLogonUsers(value []LogonUserable)()
    SetNetBiosName(value *string)()
    SetNetworkInterfaces(value []NetworkInterfaceable)()
    SetOs(value *string)()
    SetOsVersion(value *string)()
    SetParentHost(value *string)()
    SetRelatedHostIds(value []string)()
    SetRiskScore(value *string)()
    SetTags(value []string)()
    SetVendorInformation(value SecurityVendorInformationable)()
}
