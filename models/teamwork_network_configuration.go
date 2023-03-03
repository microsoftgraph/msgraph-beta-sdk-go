package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TeamworkNetworkConfiguration 
type TeamworkNetworkConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkNetworkConfiguration instantiates a new teamworkNetworkConfiguration and sets the default values.
func NewTeamworkNetworkConfiguration()(*TeamworkNetworkConfiguration) {
    m := &TeamworkNetworkConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkNetworkConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkNetworkConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkNetworkConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkNetworkConfiguration) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *TeamworkNetworkConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDefaultGateway gets the defaultGateway property value. The default gateway is the path used to pass information when the destination is unknown to the device.
func (m *TeamworkNetworkConfiguration) GetDefaultGateway()(*string) {
    val, err := m.GetBackingStore().Get("defaultGateway")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDomainName gets the domainName property value. The network domain of the device, for example, contoso.com.
func (m *TeamworkNetworkConfiguration) GetDomainName()(*string) {
    val, err := m.GetBackingStore().Get("domainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkNetworkConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultGateway"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultGateway(val)
        }
        return nil
    }
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["hostName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostName(val)
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["isDhcpEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDhcpEnabled(val)
        }
        return nil
    }
    res["isPCPortEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPCPortEnabled(val)
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
    res["primaryDns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryDns(val)
        }
        return nil
    }
    res["secondaryDns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryDns(val)
        }
        return nil
    }
    res["subnetMask"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetMask(val)
        }
        return nil
    }
    return res
}
// GetHostName gets the hostName property value. The device name on a network.
func (m *TeamworkNetworkConfiguration) GetHostName()(*string) {
    val, err := m.GetBackingStore().Get("hostName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIpAddress gets the ipAddress property value. The IP address is a numerical label that uniquely identifies every device connected to the internet.
func (m *TeamworkNetworkConfiguration) GetIpAddress()(*string) {
    val, err := m.GetBackingStore().Get("ipAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsDhcpEnabled gets the isDhcpEnabled property value. True if DHCP is enabled.
func (m *TeamworkNetworkConfiguration) GetIsDhcpEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isDhcpEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsPCPortEnabled gets the isPCPortEnabled property value. True if the PC port is enabled.
func (m *TeamworkNetworkConfiguration) GetIsPCPortEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isPCPortEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkNetworkConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrimaryDns gets the primaryDns property value. A primary DNS is the first point of contact for a device that translates the hostname into an IP address.
func (m *TeamworkNetworkConfiguration) GetPrimaryDns()(*string) {
    val, err := m.GetBackingStore().Get("primaryDns")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSecondaryDns gets the secondaryDns property value. A secondary DNS is used when the primary DNS is not available.
func (m *TeamworkNetworkConfiguration) GetSecondaryDns()(*string) {
    val, err := m.GetBackingStore().Get("secondaryDns")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubnetMask gets the subnetMask property value. A subnet mask is a number that distinguishes the network address and the host address within an IP address.
func (m *TeamworkNetworkConfiguration) GetSubnetMask()(*string) {
    val, err := m.GetBackingStore().Get("subnetMask")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
func (m *TeamworkNetworkConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *TeamworkNetworkConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDefaultGateway sets the defaultGateway property value. The default gateway is the path used to pass information when the destination is unknown to the device.
func (m *TeamworkNetworkConfiguration) SetDefaultGateway(value *string)() {
    err := m.GetBackingStore().Set("defaultGateway", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainName sets the domainName property value. The network domain of the device, for example, contoso.com.
func (m *TeamworkNetworkConfiguration) SetDomainName(value *string)() {
    err := m.GetBackingStore().Set("domainName", value)
    if err != nil {
        panic(err)
    }
}
// SetHostName sets the hostName property value. The device name on a network.
func (m *TeamworkNetworkConfiguration) SetHostName(value *string)() {
    err := m.GetBackingStore().Set("hostName", value)
    if err != nil {
        panic(err)
    }
}
// SetIpAddress sets the ipAddress property value. The IP address is a numerical label that uniquely identifies every device connected to the internet.
func (m *TeamworkNetworkConfiguration) SetIpAddress(value *string)() {
    err := m.GetBackingStore().Set("ipAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDhcpEnabled sets the isDhcpEnabled property value. True if DHCP is enabled.
func (m *TeamworkNetworkConfiguration) SetIsDhcpEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isDhcpEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsPCPortEnabled sets the isPCPortEnabled property value. True if the PC port is enabled.
func (m *TeamworkNetworkConfiguration) SetIsPCPortEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isPCPortEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkNetworkConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPrimaryDns sets the primaryDns property value. A primary DNS is the first point of contact for a device that translates the hostname into an IP address.
func (m *TeamworkNetworkConfiguration) SetPrimaryDns(value *string)() {
    err := m.GetBackingStore().Set("primaryDns", value)
    if err != nil {
        panic(err)
    }
}
// SetSecondaryDns sets the secondaryDns property value. A secondary DNS is used when the primary DNS is not available.
func (m *TeamworkNetworkConfiguration) SetSecondaryDns(value *string)() {
    err := m.GetBackingStore().Set("secondaryDns", value)
    if err != nil {
        panic(err)
    }
}
// SetSubnetMask sets the subnetMask property value. A subnet mask is a number that distinguishes the network address and the host address within an IP address.
func (m *TeamworkNetworkConfiguration) SetSubnetMask(value *string)() {
    err := m.GetBackingStore().Set("subnetMask", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkNetworkConfigurationable 
type TeamworkNetworkConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDefaultGateway()(*string)
    GetDomainName()(*string)
    GetHostName()(*string)
    GetIpAddress()(*string)
    GetIsDhcpEnabled()(*bool)
    GetIsPCPortEnabled()(*bool)
    GetOdataType()(*string)
    GetPrimaryDns()(*string)
    GetSecondaryDns()(*string)
    GetSubnetMask()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDefaultGateway(value *string)()
    SetDomainName(value *string)()
    SetHostName(value *string)()
    SetIpAddress(value *string)()
    SetIsDhcpEnabled(value *bool)()
    SetIsPCPortEnabled(value *bool)()
    SetOdataType(value *string)()
    SetPrimaryDns(value *string)()
    SetSecondaryDns(value *string)()
    SetSubnetMask(value *string)()
}
