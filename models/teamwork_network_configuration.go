package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkNetworkConfiguration 
type TeamworkNetworkConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The default gateway is the path used to pass information when the destination is unknown to the device.
    defaultGateway *string
    // The network domain of the device, for example, contoso.com.
    domainName *string
    // The device name on a network.
    hostName *string
    // The IP address is a numerical label that uniquely identifies every device connected to the internet.
    ipAddress *string
    // True if DHCP is enabled.
    isDhcpEnabled *bool
    // True if the PC port is enabled.
    isPCPortEnabled *bool
    // The OdataType property
    odataType *string
    // A primary DNS is the first point of contact for a device that translates the hostname into an IP address.
    primaryDns *string
    // A secondary DNS is used when the primary DNS is not available.
    secondaryDns *string
    // A subnet mask is a number that distinguishes the network address and the host address within an IP address.
    subnetMask *string
}
// NewTeamworkNetworkConfiguration instantiates a new teamworkNetworkConfiguration and sets the default values.
func NewTeamworkNetworkConfiguration()(*TeamworkNetworkConfiguration) {
    m := &TeamworkNetworkConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.teamworkNetworkConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTeamworkNetworkConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkNetworkConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkNetworkConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkNetworkConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDefaultGateway gets the defaultGateway property value. The default gateway is the path used to pass information when the destination is unknown to the device.
func (m *TeamworkNetworkConfiguration) GetDefaultGateway()(*string) {
    return m.defaultGateway
}
// GetDomainName gets the domainName property value. The network domain of the device, for example, contoso.com.
func (m *TeamworkNetworkConfiguration) GetDomainName()(*string) {
    return m.domainName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkNetworkConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultGateway"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultGateway)
    res["domainName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDomainName)
    res["hostName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHostName)
    res["ipAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIpAddress)
    res["isDhcpEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsDhcpEnabled)
    res["isPCPortEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsPCPortEnabled)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["primaryDns"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPrimaryDns)
    res["secondaryDns"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSecondaryDns)
    res["subnetMask"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubnetMask)
    return res
}
// GetHostName gets the hostName property value. The device name on a network.
func (m *TeamworkNetworkConfiguration) GetHostName()(*string) {
    return m.hostName
}
// GetIpAddress gets the ipAddress property value. The IP address is a numerical label that uniquely identifies every device connected to the internet.
func (m *TeamworkNetworkConfiguration) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetIsDhcpEnabled gets the isDhcpEnabled property value. True if DHCP is enabled.
func (m *TeamworkNetworkConfiguration) GetIsDhcpEnabled()(*bool) {
    return m.isDhcpEnabled
}
// GetIsPCPortEnabled gets the isPCPortEnabled property value. True if the PC port is enabled.
func (m *TeamworkNetworkConfiguration) GetIsPCPortEnabled()(*bool) {
    return m.isPCPortEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkNetworkConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetPrimaryDns gets the primaryDns property value. A primary DNS is the first point of contact for a device that translates the hostname into an IP address.
func (m *TeamworkNetworkConfiguration) GetPrimaryDns()(*string) {
    return m.primaryDns
}
// GetSecondaryDns gets the secondaryDns property value. A secondary DNS is used when the primary DNS is not available.
func (m *TeamworkNetworkConfiguration) GetSecondaryDns()(*string) {
    return m.secondaryDns
}
// GetSubnetMask gets the subnetMask property value. A subnet mask is a number that distinguishes the network address and the host address within an IP address.
func (m *TeamworkNetworkConfiguration) GetSubnetMask()(*string) {
    return m.subnetMask
}
// Serialize serializes information the current object
func (m *TeamworkNetworkConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultGateway", m.GetDefaultGateway())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hostName", m.GetHostName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDhcpEnabled", m.GetIsDhcpEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPCPortEnabled", m.GetIsPCPortEnabled())
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
        err := writer.WriteStringValue("primaryDns", m.GetPrimaryDns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secondaryDns", m.GetSecondaryDns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnetMask", m.GetSubnetMask())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkNetworkConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDefaultGateway sets the defaultGateway property value. The default gateway is the path used to pass information when the destination is unknown to the device.
func (m *TeamworkNetworkConfiguration) SetDefaultGateway(value *string)() {
    m.defaultGateway = value
}
// SetDomainName sets the domainName property value. The network domain of the device, for example, contoso.com.
func (m *TeamworkNetworkConfiguration) SetDomainName(value *string)() {
    m.domainName = value
}
// SetHostName sets the hostName property value. The device name on a network.
func (m *TeamworkNetworkConfiguration) SetHostName(value *string)() {
    m.hostName = value
}
// SetIpAddress sets the ipAddress property value. The IP address is a numerical label that uniquely identifies every device connected to the internet.
func (m *TeamworkNetworkConfiguration) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetIsDhcpEnabled sets the isDhcpEnabled property value. True if DHCP is enabled.
func (m *TeamworkNetworkConfiguration) SetIsDhcpEnabled(value *bool)() {
    m.isDhcpEnabled = value
}
// SetIsPCPortEnabled sets the isPCPortEnabled property value. True if the PC port is enabled.
func (m *TeamworkNetworkConfiguration) SetIsPCPortEnabled(value *bool)() {
    m.isPCPortEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkNetworkConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPrimaryDns sets the primaryDns property value. A primary DNS is the first point of contact for a device that translates the hostname into an IP address.
func (m *TeamworkNetworkConfiguration) SetPrimaryDns(value *string)() {
    m.primaryDns = value
}
// SetSecondaryDns sets the secondaryDns property value. A secondary DNS is used when the primary DNS is not available.
func (m *TeamworkNetworkConfiguration) SetSecondaryDns(value *string)() {
    m.secondaryDns = value
}
// SetSubnetMask sets the subnetMask property value. A subnet mask is a number that distinguishes the network address and the host address within an IP address.
func (m *TeamworkNetworkConfiguration) SetSubnetMask(value *string)() {
    m.subnetMask = value
}
