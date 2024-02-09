package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WindowsFirewallRule a rule controlling traffic through the Windows Firewall.
type WindowsFirewallRule struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsFirewallRule instantiates a new WindowsFirewallRule and sets the default values.
func NewWindowsFirewallRule()(*WindowsFirewallRule) {
    m := &WindowsFirewallRule{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsFirewallRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsFirewallRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsFirewallRule(), nil
}
// GetAction gets the action property value. State Management Setting.
// returns a *StateManagementSetting when successful
func (m *WindowsFirewallRule) GetAction()(*StateManagementSetting) {
    val, err := m.GetBackingStore().Get("action")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*StateManagementSetting)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WindowsFirewallRule) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *WindowsFirewallRule) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDescription gets the description property value. The description of the rule.
// returns a *string when successful
func (m *WindowsFirewallRule) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the rule. Does not need to be unique.
// returns a *string when successful
func (m *WindowsFirewallRule) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEdgeTraversal gets the edgeTraversal property value. State Management Setting.
// returns a *StateManagementSetting when successful
func (m *WindowsFirewallRule) GetEdgeTraversal()(*StateManagementSetting) {
    val, err := m.GetBackingStore().Get("edgeTraversal")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*StateManagementSetting)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsFirewallRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStateManagementSetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*StateManagementSetting))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["edgeTraversal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStateManagementSetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeTraversal(val.(*StateManagementSetting))
        }
        return nil
    }
    res["filePath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilePath(val)
        }
        return nil
    }
    res["interfaceTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsFirewallRuleInterfaceTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInterfaceTypes(val.(*WindowsFirewallRuleInterfaceTypes))
        }
        return nil
    }
    res["localAddressRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetLocalAddressRanges(res)
        }
        return nil
    }
    res["localPortRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetLocalPortRanges(res)
        }
        return nil
    }
    res["localUserAuthorizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalUserAuthorizations(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["packageFamilyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageFamilyName(val)
        }
        return nil
    }
    res["profileTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsFirewallRuleNetworkProfileTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileTypes(val.(*WindowsFirewallRuleNetworkProfileTypes))
        }
        return nil
    }
    res["protocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtocol(val)
        }
        return nil
    }
    res["remoteAddressRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRemoteAddressRanges(res)
        }
        return nil
    }
    res["remotePortRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRemotePortRanges(res)
        }
        return nil
    }
    res["serviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceName(val)
        }
        return nil
    }
    res["trafficDirection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsFirewallRuleTrafficDirectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrafficDirection(val.(*WindowsFirewallRuleTrafficDirectionType))
        }
        return nil
    }
    return res
}
// GetFilePath gets the filePath property value. The full file path of an app that's affected by the firewall rule.
// returns a *string when successful
func (m *WindowsFirewallRule) GetFilePath()(*string) {
    val, err := m.GetBackingStore().Get("filePath")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInterfaceTypes gets the interfaceTypes property value. Flags representing firewall rule interface types.
// returns a *WindowsFirewallRuleInterfaceTypes when successful
func (m *WindowsFirewallRule) GetInterfaceTypes()(*WindowsFirewallRuleInterfaceTypes) {
    val, err := m.GetBackingStore().Get("interfaceTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsFirewallRuleInterfaceTypes)
    }
    return nil
}
// GetLocalAddressRanges gets the localAddressRanges property value. List of local addresses covered by the rule. Default is any address. Valid tokens include:'' indicates any local address. If present, this must be the only token included.A subnet can be specified using either the subnet mask or network prefix notation. If neither a subnet mask nor a network prefix is specified, the subnet mask defaults to 255.255.255.255.A valid IPv6 address.An IPv4 address range in the format of 'start address - end address' with no spaces included.An IPv6 address range in the format of 'start address - end address' with no spaces included.
// returns a []string when successful
func (m *WindowsFirewallRule) GetLocalAddressRanges()([]string) {
    val, err := m.GetBackingStore().Get("localAddressRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetLocalPortRanges gets the localPortRanges property value. List of local port ranges. For example, '100-120', '200', '300-320'. If not specified, the default is All.
// returns a []string when successful
func (m *WindowsFirewallRule) GetLocalPortRanges()([]string) {
    val, err := m.GetBackingStore().Get("localPortRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetLocalUserAuthorizations gets the localUserAuthorizations property value. Specifies the list of authorized local users for the app container. This is a string in Security Descriptor Definition Language (SDDL) format.
// returns a *string when successful
func (m *WindowsFirewallRule) GetLocalUserAuthorizations()(*string) {
    val, err := m.GetBackingStore().Get("localUserAuthorizations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WindowsFirewallRule) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPackageFamilyName gets the packageFamilyName property value. The package family name of a Microsoft Store application that's affected by the firewall rule.
// returns a *string when successful
func (m *WindowsFirewallRule) GetPackageFamilyName()(*string) {
    val, err := m.GetBackingStore().Get("packageFamilyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProfileTypes gets the profileTypes property value. Flags representing which network profile types apply to a firewall rule.
// returns a *WindowsFirewallRuleNetworkProfileTypes when successful
func (m *WindowsFirewallRule) GetProfileTypes()(*WindowsFirewallRuleNetworkProfileTypes) {
    val, err := m.GetBackingStore().Get("profileTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsFirewallRuleNetworkProfileTypes)
    }
    return nil
}
// GetProtocol gets the protocol property value. 0-255 number representing the IP protocol (TCP = 6, UDP = 17). If not specified, the default is All. Valid values 0 to 255
// returns a *int32 when successful
func (m *WindowsFirewallRule) GetProtocol()(*int32) {
    val, err := m.GetBackingStore().Get("protocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRemoteAddressRanges gets the remoteAddressRanges property value. List of tokens specifying the remote addresses covered by the rule. Tokens are case insensitive. Default is any address. Valid tokens include:'' indicates any remote address. If present, this must be the only token included.'Defaultgateway''DHCP''DNS''WINS''Intranet' (supported on Windows versions 1809+)'RmtIntranet' (supported on Windows versions 1809+)'Internet' (supported on Windows versions 1809+)'Ply2Renders' (supported on Windows versions 1809+)'LocalSubnet' indicates any local address on the local subnet.A subnet can be specified using either the subnet mask or network prefix notation. If neither a subnet mask nor a network prefix is specified, the subnet mask defaults to 255.255.255.255.A valid IPv6 address.An IPv4 address range in the format of 'start address - end address' with no spaces included.An IPv6 address range in the format of 'start address - end address' with no spaces included.
// returns a []string when successful
func (m *WindowsFirewallRule) GetRemoteAddressRanges()([]string) {
    val, err := m.GetBackingStore().Get("remoteAddressRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRemotePortRanges gets the remotePortRanges property value. List of remote port ranges. For example, '100-120', '200', '300-320'. If not specified, the default is All.
// returns a []string when successful
func (m *WindowsFirewallRule) GetRemotePortRanges()([]string) {
    val, err := m.GetBackingStore().Get("remotePortRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetServiceName gets the serviceName property value. The name used in cases when a service, not an application, is sending or receiving traffic.
// returns a *string when successful
func (m *WindowsFirewallRule) GetServiceName()(*string) {
    val, err := m.GetBackingStore().Get("serviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTrafficDirection gets the trafficDirection property value. Firewall rule traffic directions.
// returns a *WindowsFirewallRuleTrafficDirectionType when successful
func (m *WindowsFirewallRule) GetTrafficDirection()(*WindowsFirewallRuleTrafficDirectionType) {
    val, err := m.GetBackingStore().Get("trafficDirection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsFirewallRuleTrafficDirectionType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsFirewallRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEdgeTraversal() != nil {
        cast := (*m.GetEdgeTraversal()).String()
        err := writer.WriteStringValue("edgeTraversal", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filePath", m.GetFilePath())
        if err != nil {
            return err
        }
    }
    if m.GetInterfaceTypes() != nil {
        cast := (*m.GetInterfaceTypes()).String()
        err := writer.WriteStringValue("interfaceTypes", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalAddressRanges() != nil {
        err := writer.WriteCollectionOfStringValues("localAddressRanges", m.GetLocalAddressRanges())
        if err != nil {
            return err
        }
    }
    if m.GetLocalPortRanges() != nil {
        err := writer.WriteCollectionOfStringValues("localPortRanges", m.GetLocalPortRanges())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("localUserAuthorizations", m.GetLocalUserAuthorizations())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("packageFamilyName", m.GetPackageFamilyName())
        if err != nil {
            return err
        }
    }
    if m.GetProfileTypes() != nil {
        cast := (*m.GetProfileTypes()).String()
        err := writer.WriteStringValue("profileTypes", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("protocol", m.GetProtocol())
        if err != nil {
            return err
        }
    }
    if m.GetRemoteAddressRanges() != nil {
        err := writer.WriteCollectionOfStringValues("remoteAddressRanges", m.GetRemoteAddressRanges())
        if err != nil {
            return err
        }
    }
    if m.GetRemotePortRanges() != nil {
        err := writer.WriteCollectionOfStringValues("remotePortRanges", m.GetRemotePortRanges())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serviceName", m.GetServiceName())
        if err != nil {
            return err
        }
    }
    if m.GetTrafficDirection() != nil {
        cast := (*m.GetTrafficDirection()).String()
        err := writer.WriteStringValue("trafficDirection", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. State Management Setting.
func (m *WindowsFirewallRule) SetAction(value *StateManagementSetting)() {
    err := m.GetBackingStore().Set("action", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsFirewallRule) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WindowsFirewallRule) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDescription sets the description property value. The description of the rule.
func (m *WindowsFirewallRule) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the rule. Does not need to be unique.
func (m *WindowsFirewallRule) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEdgeTraversal sets the edgeTraversal property value. State Management Setting.
func (m *WindowsFirewallRule) SetEdgeTraversal(value *StateManagementSetting)() {
    err := m.GetBackingStore().Set("edgeTraversal", value)
    if err != nil {
        panic(err)
    }
}
// SetFilePath sets the filePath property value. The full file path of an app that's affected by the firewall rule.
func (m *WindowsFirewallRule) SetFilePath(value *string)() {
    err := m.GetBackingStore().Set("filePath", value)
    if err != nil {
        panic(err)
    }
}
// SetInterfaceTypes sets the interfaceTypes property value. Flags representing firewall rule interface types.
func (m *WindowsFirewallRule) SetInterfaceTypes(value *WindowsFirewallRuleInterfaceTypes)() {
    err := m.GetBackingStore().Set("interfaceTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalAddressRanges sets the localAddressRanges property value. List of local addresses covered by the rule. Default is any address. Valid tokens include:'' indicates any local address. If present, this must be the only token included.A subnet can be specified using either the subnet mask or network prefix notation. If neither a subnet mask nor a network prefix is specified, the subnet mask defaults to 255.255.255.255.A valid IPv6 address.An IPv4 address range in the format of 'start address - end address' with no spaces included.An IPv6 address range in the format of 'start address - end address' with no spaces included.
func (m *WindowsFirewallRule) SetLocalAddressRanges(value []string)() {
    err := m.GetBackingStore().Set("localAddressRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalPortRanges sets the localPortRanges property value. List of local port ranges. For example, '100-120', '200', '300-320'. If not specified, the default is All.
func (m *WindowsFirewallRule) SetLocalPortRanges(value []string)() {
    err := m.GetBackingStore().Set("localPortRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalUserAuthorizations sets the localUserAuthorizations property value. Specifies the list of authorized local users for the app container. This is a string in Security Descriptor Definition Language (SDDL) format.
func (m *WindowsFirewallRule) SetLocalUserAuthorizations(value *string)() {
    err := m.GetBackingStore().Set("localUserAuthorizations", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsFirewallRule) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPackageFamilyName sets the packageFamilyName property value. The package family name of a Microsoft Store application that's affected by the firewall rule.
func (m *WindowsFirewallRule) SetPackageFamilyName(value *string)() {
    err := m.GetBackingStore().Set("packageFamilyName", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileTypes sets the profileTypes property value. Flags representing which network profile types apply to a firewall rule.
func (m *WindowsFirewallRule) SetProfileTypes(value *WindowsFirewallRuleNetworkProfileTypes)() {
    err := m.GetBackingStore().Set("profileTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetProtocol sets the protocol property value. 0-255 number representing the IP protocol (TCP = 6, UDP = 17). If not specified, the default is All. Valid values 0 to 255
func (m *WindowsFirewallRule) SetProtocol(value *int32)() {
    err := m.GetBackingStore().Set("protocol", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteAddressRanges sets the remoteAddressRanges property value. List of tokens specifying the remote addresses covered by the rule. Tokens are case insensitive. Default is any address. Valid tokens include:'' indicates any remote address. If present, this must be the only token included.'Defaultgateway''DHCP''DNS''WINS''Intranet' (supported on Windows versions 1809+)'RmtIntranet' (supported on Windows versions 1809+)'Internet' (supported on Windows versions 1809+)'Ply2Renders' (supported on Windows versions 1809+)'LocalSubnet' indicates any local address on the local subnet.A subnet can be specified using either the subnet mask or network prefix notation. If neither a subnet mask nor a network prefix is specified, the subnet mask defaults to 255.255.255.255.A valid IPv6 address.An IPv4 address range in the format of 'start address - end address' with no spaces included.An IPv6 address range in the format of 'start address - end address' with no spaces included.
func (m *WindowsFirewallRule) SetRemoteAddressRanges(value []string)() {
    err := m.GetBackingStore().Set("remoteAddressRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetRemotePortRanges sets the remotePortRanges property value. List of remote port ranges. For example, '100-120', '200', '300-320'. If not specified, the default is All.
func (m *WindowsFirewallRule) SetRemotePortRanges(value []string)() {
    err := m.GetBackingStore().Set("remotePortRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetServiceName sets the serviceName property value. The name used in cases when a service, not an application, is sending or receiving traffic.
func (m *WindowsFirewallRule) SetServiceName(value *string)() {
    err := m.GetBackingStore().Set("serviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetTrafficDirection sets the trafficDirection property value. Firewall rule traffic directions.
func (m *WindowsFirewallRule) SetTrafficDirection(value *WindowsFirewallRuleTrafficDirectionType)() {
    err := m.GetBackingStore().Set("trafficDirection", value)
    if err != nil {
        panic(err)
    }
}
type WindowsFirewallRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*StateManagementSetting)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEdgeTraversal()(*StateManagementSetting)
    GetFilePath()(*string)
    GetInterfaceTypes()(*WindowsFirewallRuleInterfaceTypes)
    GetLocalAddressRanges()([]string)
    GetLocalPortRanges()([]string)
    GetLocalUserAuthorizations()(*string)
    GetOdataType()(*string)
    GetPackageFamilyName()(*string)
    GetProfileTypes()(*WindowsFirewallRuleNetworkProfileTypes)
    GetProtocol()(*int32)
    GetRemoteAddressRanges()([]string)
    GetRemotePortRanges()([]string)
    GetServiceName()(*string)
    GetTrafficDirection()(*WindowsFirewallRuleTrafficDirectionType)
    SetAction(value *StateManagementSetting)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEdgeTraversal(value *StateManagementSetting)()
    SetFilePath(value *string)()
    SetInterfaceTypes(value *WindowsFirewallRuleInterfaceTypes)()
    SetLocalAddressRanges(value []string)()
    SetLocalPortRanges(value []string)()
    SetLocalUserAuthorizations(value *string)()
    SetOdataType(value *string)()
    SetPackageFamilyName(value *string)()
    SetProfileTypes(value *WindowsFirewallRuleNetworkProfileTypes)()
    SetProtocol(value *int32)()
    SetRemoteAddressRanges(value []string)()
    SetRemotePortRanges(value []string)()
    SetServiceName(value *string)()
    SetTrafficDirection(value *WindowsFirewallRuleTrafficDirectionType)()
}
