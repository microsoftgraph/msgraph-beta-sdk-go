package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsFirewallRule a rule controlling traffic through the Windows Firewall.
type WindowsFirewallRule struct {
    // The action the rule enforces. If not specified, the default is Allowed. Possible values are: notConfigured, blocked, allowed.
    action *StateManagementSetting
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The description of the rule.
    description *string
    // The display name of the rule. Does not need to be unique.
    displayName *string
    // Indicates whether edge traversal is enabled or disabled for this rule. The EdgeTraversal setting indicates that specific inbound traffic is allowed to tunnel through NATs and other edge devices using the Teredo tunneling technology. In order for this setting to work correctly, the application or service with the inbound firewall rule needs to support IPv6. The primary application of this setting allows listeners on the host to be globally addressable through a Teredo IPv6 address. New rules have the EdgeTraversal property disabled by default. Possible values are: notConfigured, blocked, allowed.
    edgeTraversal *StateManagementSetting
    // The full file path of an app that's affected by the firewall rule.
    filePath *string
    // The interface types of the rule. Possible values are: notConfigured, remoteAccess, wireless, lan.
    interfaceTypes *WindowsFirewallRuleInterfaceTypes
    // List of local addresses covered by the rule. Default is any address. Valid tokens include:'' indicates any local address. If present, this must be the only token included.A subnet can be specified using either the subnet mask or network prefix notation. If neither a subnet mask nor a network prefix is specified, the subnet mask defaults to 255.255.255.255.A valid IPv6 address.An IPv4 address range in the format of 'start address - end address' with no spaces included.An IPv6 address range in the format of 'start address - end address' with no spaces included.
    localAddressRanges []string
    // List of local port ranges. For example, '100-120', '200', '300-320'. If not specified, the default is All.
    localPortRanges []string
    // Specifies the list of authorized local users for the app container. This is a string in Security Descriptor Definition Language (SDDL) format.
    localUserAuthorizations *string
    // The package family name of a Microsoft Store application that's affected by the firewall rule.
    packageFamilyName *string
    // Specifies the profiles to which the rule belongs. If not specified, the default is All. Possible values are: notConfigured, domain, private, public.
    profileTypes *WindowsFirewallRuleNetworkProfileTypes
    // 0-255 number representing the IP protocol (TCP = 6, UDP = 17). If not specified, the default is All. Valid values 0 to 255
    protocol *int32
    // List of tokens specifying the remote addresses covered by the rule. Tokens are case insensitive. Default is any address. Valid tokens include:'' indicates any remote address. If present, this must be the only token included.'Defaultgateway''DHCP''DNS''WINS''Intranet' (supported on Windows versions 1809+)'RmtIntranet' (supported on Windows versions 1809+)'Internet' (supported on Windows versions 1809+)'Ply2Renders' (supported on Windows versions 1809+)'LocalSubnet' indicates any local address on the local subnet.A subnet can be specified using either the subnet mask or network prefix notation. If neither a subnet mask nor a network prefix is specified, the subnet mask defaults to 255.255.255.255.A valid IPv6 address.An IPv4 address range in the format of 'start address - end address' with no spaces included.An IPv6 address range in the format of 'start address - end address' with no spaces included.
    remoteAddressRanges []string
    // List of remote port ranges. For example, '100-120', '200', '300-320'. If not specified, the default is All.
    remotePortRanges []string
    // The name used in cases when a service, not an application, is sending or receiving traffic.
    serviceName *string
    // The traffic direction that the rule is enabled for. If not specified, the default is Out. Possible values are: notConfigured, out, in.
    trafficDirection *WindowsFirewallRuleTrafficDirectionType
}
// NewWindowsFirewallRule instantiates a new windowsFirewallRule and sets the default values.
func NewWindowsFirewallRule()(*WindowsFirewallRule) {
    m := &WindowsFirewallRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsFirewallRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsFirewallRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsFirewallRule(), nil
}
// GetAction gets the action property value. The action the rule enforces. If not specified, the default is Allowed. Possible values are: notConfigured, blocked, allowed.
func (m *WindowsFirewallRule) GetAction()(*StateManagementSetting) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsFirewallRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. The description of the rule.
func (m *WindowsFirewallRule) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name of the rule. Does not need to be unique.
func (m *WindowsFirewallRule) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEdgeTraversal gets the edgeTraversal property value. Indicates whether edge traversal is enabled or disabled for this rule. The EdgeTraversal setting indicates that specific inbound traffic is allowed to tunnel through NATs and other edge devices using the Teredo tunneling technology. In order for this setting to work correctly, the application or service with the inbound firewall rule needs to support IPv6. The primary application of this setting allows listeners on the host to be globally addressable through a Teredo IPv6 address. New rules have the EdgeTraversal property disabled by default. Possible values are: notConfigured, blocked, allowed.
func (m *WindowsFirewallRule) GetEdgeTraversal()(*StateManagementSetting) {
    if m == nil {
        return nil
    } else {
        return m.edgeTraversal
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
                res[i] = *(v.(*string))
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
                res[i] = *(v.(*string))
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
                res[i] = *(v.(*string))
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
                res[i] = *(v.(*string))
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
func (m *WindowsFirewallRule) GetFilePath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filePath
    }
}
// GetInterfaceTypes gets the interfaceTypes property value. The interface types of the rule. Possible values are: notConfigured, remoteAccess, wireless, lan.
func (m *WindowsFirewallRule) GetInterfaceTypes()(*WindowsFirewallRuleInterfaceTypes) {
    if m == nil {
        return nil
    } else {
        return m.interfaceTypes
    }
}
// GetLocalAddressRanges gets the localAddressRanges property value. List of local addresses covered by the rule. Default is any address. Valid tokens include:'' indicates any local address. If present, this must be the only token included.A subnet can be specified using either the subnet mask or network prefix notation. If neither a subnet mask nor a network prefix is specified, the subnet mask defaults to 255.255.255.255.A valid IPv6 address.An IPv4 address range in the format of 'start address - end address' with no spaces included.An IPv6 address range in the format of 'start address - end address' with no spaces included.
func (m *WindowsFirewallRule) GetLocalAddressRanges()([]string) {
    if m == nil {
        return nil
    } else {
        return m.localAddressRanges
    }
}
// GetLocalPortRanges gets the localPortRanges property value. List of local port ranges. For example, '100-120', '200', '300-320'. If not specified, the default is All.
func (m *WindowsFirewallRule) GetLocalPortRanges()([]string) {
    if m == nil {
        return nil
    } else {
        return m.localPortRanges
    }
}
// GetLocalUserAuthorizations gets the localUserAuthorizations property value. Specifies the list of authorized local users for the app container. This is a string in Security Descriptor Definition Language (SDDL) format.
func (m *WindowsFirewallRule) GetLocalUserAuthorizations()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localUserAuthorizations
    }
}
// GetPackageFamilyName gets the packageFamilyName property value. The package family name of a Microsoft Store application that's affected by the firewall rule.
func (m *WindowsFirewallRule) GetPackageFamilyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.packageFamilyName
    }
}
// GetProfileTypes gets the profileTypes property value. Specifies the profiles to which the rule belongs. If not specified, the default is All. Possible values are: notConfigured, domain, private, public.
func (m *WindowsFirewallRule) GetProfileTypes()(*WindowsFirewallRuleNetworkProfileTypes) {
    if m == nil {
        return nil
    } else {
        return m.profileTypes
    }
}
// GetProtocol gets the protocol property value. 0-255 number representing the IP protocol (TCP = 6, UDP = 17). If not specified, the default is All. Valid values 0 to 255
func (m *WindowsFirewallRule) GetProtocol()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.protocol
    }
}
// GetRemoteAddressRanges gets the remoteAddressRanges property value. List of tokens specifying the remote addresses covered by the rule. Tokens are case insensitive. Default is any address. Valid tokens include:'' indicates any remote address. If present, this must be the only token included.'Defaultgateway''DHCP''DNS''WINS''Intranet' (supported on Windows versions 1809+)'RmtIntranet' (supported on Windows versions 1809+)'Internet' (supported on Windows versions 1809+)'Ply2Renders' (supported on Windows versions 1809+)'LocalSubnet' indicates any local address on the local subnet.A subnet can be specified using either the subnet mask or network prefix notation. If neither a subnet mask nor a network prefix is specified, the subnet mask defaults to 255.255.255.255.A valid IPv6 address.An IPv4 address range in the format of 'start address - end address' with no spaces included.An IPv6 address range in the format of 'start address - end address' with no spaces included.
func (m *WindowsFirewallRule) GetRemoteAddressRanges()([]string) {
    if m == nil {
        return nil
    } else {
        return m.remoteAddressRanges
    }
}
// GetRemotePortRanges gets the remotePortRanges property value. List of remote port ranges. For example, '100-120', '200', '300-320'. If not specified, the default is All.
func (m *WindowsFirewallRule) GetRemotePortRanges()([]string) {
    if m == nil {
        return nil
    } else {
        return m.remotePortRanges
    }
}
// GetServiceName gets the serviceName property value. The name used in cases when a service, not an application, is sending or receiving traffic.
func (m *WindowsFirewallRule) GetServiceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceName
    }
}
// GetTrafficDirection gets the trafficDirection property value. The traffic direction that the rule is enabled for. If not specified, the default is Out. Possible values are: notConfigured, out, in.
func (m *WindowsFirewallRule) GetTrafficDirection()(*WindowsFirewallRuleTrafficDirectionType) {
    if m == nil {
        return nil
    } else {
        return m.trafficDirection
    }
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
// SetAction sets the action property value. The action the rule enforces. If not specified, the default is Allowed. Possible values are: notConfigured, blocked, allowed.
func (m *WindowsFirewallRule) SetAction(value *StateManagementSetting)() {
    if m != nil {
        m.action = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsFirewallRule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. The description of the rule.
func (m *WindowsFirewallRule) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the rule. Does not need to be unique.
func (m *WindowsFirewallRule) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEdgeTraversal sets the edgeTraversal property value. Indicates whether edge traversal is enabled or disabled for this rule. The EdgeTraversal setting indicates that specific inbound traffic is allowed to tunnel through NATs and other edge devices using the Teredo tunneling technology. In order for this setting to work correctly, the application or service with the inbound firewall rule needs to support IPv6. The primary application of this setting allows listeners on the host to be globally addressable through a Teredo IPv6 address. New rules have the EdgeTraversal property disabled by default. Possible values are: notConfigured, blocked, allowed.
func (m *WindowsFirewallRule) SetEdgeTraversal(value *StateManagementSetting)() {
    if m != nil {
        m.edgeTraversal = value
    }
}
// SetFilePath sets the filePath property value. The full file path of an app that's affected by the firewall rule.
func (m *WindowsFirewallRule) SetFilePath(value *string)() {
    if m != nil {
        m.filePath = value
    }
}
// SetInterfaceTypes sets the interfaceTypes property value. The interface types of the rule. Possible values are: notConfigured, remoteAccess, wireless, lan.
func (m *WindowsFirewallRule) SetInterfaceTypes(value *WindowsFirewallRuleInterfaceTypes)() {
    if m != nil {
        m.interfaceTypes = value
    }
}
// SetLocalAddressRanges sets the localAddressRanges property value. List of local addresses covered by the rule. Default is any address. Valid tokens include:'' indicates any local address. If present, this must be the only token included.A subnet can be specified using either the subnet mask or network prefix notation. If neither a subnet mask nor a network prefix is specified, the subnet mask defaults to 255.255.255.255.A valid IPv6 address.An IPv4 address range in the format of 'start address - end address' with no spaces included.An IPv6 address range in the format of 'start address - end address' with no spaces included.
func (m *WindowsFirewallRule) SetLocalAddressRanges(value []string)() {
    if m != nil {
        m.localAddressRanges = value
    }
}
// SetLocalPortRanges sets the localPortRanges property value. List of local port ranges. For example, '100-120', '200', '300-320'. If not specified, the default is All.
func (m *WindowsFirewallRule) SetLocalPortRanges(value []string)() {
    if m != nil {
        m.localPortRanges = value
    }
}
// SetLocalUserAuthorizations sets the localUserAuthorizations property value. Specifies the list of authorized local users for the app container. This is a string in Security Descriptor Definition Language (SDDL) format.
func (m *WindowsFirewallRule) SetLocalUserAuthorizations(value *string)() {
    if m != nil {
        m.localUserAuthorizations = value
    }
}
// SetPackageFamilyName sets the packageFamilyName property value. The package family name of a Microsoft Store application that's affected by the firewall rule.
func (m *WindowsFirewallRule) SetPackageFamilyName(value *string)() {
    if m != nil {
        m.packageFamilyName = value
    }
}
// SetProfileTypes sets the profileTypes property value. Specifies the profiles to which the rule belongs. If not specified, the default is All. Possible values are: notConfigured, domain, private, public.
func (m *WindowsFirewallRule) SetProfileTypes(value *WindowsFirewallRuleNetworkProfileTypes)() {
    if m != nil {
        m.profileTypes = value
    }
}
// SetProtocol sets the protocol property value. 0-255 number representing the IP protocol (TCP = 6, UDP = 17). If not specified, the default is All. Valid values 0 to 255
func (m *WindowsFirewallRule) SetProtocol(value *int32)() {
    if m != nil {
        m.protocol = value
    }
}
// SetRemoteAddressRanges sets the remoteAddressRanges property value. List of tokens specifying the remote addresses covered by the rule. Tokens are case insensitive. Default is any address. Valid tokens include:'' indicates any remote address. If present, this must be the only token included.'Defaultgateway''DHCP''DNS''WINS''Intranet' (supported on Windows versions 1809+)'RmtIntranet' (supported on Windows versions 1809+)'Internet' (supported on Windows versions 1809+)'Ply2Renders' (supported on Windows versions 1809+)'LocalSubnet' indicates any local address on the local subnet.A subnet can be specified using either the subnet mask or network prefix notation. If neither a subnet mask nor a network prefix is specified, the subnet mask defaults to 255.255.255.255.A valid IPv6 address.An IPv4 address range in the format of 'start address - end address' with no spaces included.An IPv6 address range in the format of 'start address - end address' with no spaces included.
func (m *WindowsFirewallRule) SetRemoteAddressRanges(value []string)() {
    if m != nil {
        m.remoteAddressRanges = value
    }
}
// SetRemotePortRanges sets the remotePortRanges property value. List of remote port ranges. For example, '100-120', '200', '300-320'. If not specified, the default is All.
func (m *WindowsFirewallRule) SetRemotePortRanges(value []string)() {
    if m != nil {
        m.remotePortRanges = value
    }
}
// SetServiceName sets the serviceName property value. The name used in cases when a service, not an application, is sending or receiving traffic.
func (m *WindowsFirewallRule) SetServiceName(value *string)() {
    if m != nil {
        m.serviceName = value
    }
}
// SetTrafficDirection sets the trafficDirection property value. The traffic direction that the rule is enabled for. If not specified, the default is Out. Possible values are: notConfigured, out, in.
func (m *WindowsFirewallRule) SetTrafficDirection(value *WindowsFirewallRuleTrafficDirectionType)() {
    if m != nil {
        m.trafficDirection = value
    }
}
