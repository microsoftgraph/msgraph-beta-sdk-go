package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkNetworkConfiguration 
type TeamworkNetworkConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    defaultGateway *string;
    // 
    domainName *string;
    // 
    hostName *string;
    // 
    ipAddress *string;
    // 
    isDhcpEnabled *bool;
    // 
    isPCPortEnabled *bool;
    // 
    primaryDns *string;
    // 
    secondaryDns *string;
    // 
    subnetMask *string;
}
// NewTeamworkNetworkConfiguration instantiates a new teamworkNetworkConfiguration and sets the default values.
func NewTeamworkNetworkConfiguration()(*TeamworkNetworkConfiguration) {
    m := &TeamworkNetworkConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkNetworkConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultGateway gets the defaultGateway property value. 
func (m *TeamworkNetworkConfiguration) GetDefaultGateway()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultGateway
    }
}
// GetDomainName gets the domainName property value. 
func (m *TeamworkNetworkConfiguration) GetDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainName
    }
}
// GetHostName gets the hostName property value. 
func (m *TeamworkNetworkConfiguration) GetHostName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hostName
    }
}
// GetIpAddress gets the ipAddress property value. 
func (m *TeamworkNetworkConfiguration) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// GetIsDhcpEnabled gets the isDhcpEnabled property value. 
func (m *TeamworkNetworkConfiguration) GetIsDhcpEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDhcpEnabled
    }
}
// GetIsPCPortEnabled gets the isPCPortEnabled property value. 
func (m *TeamworkNetworkConfiguration) GetIsPCPortEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPCPortEnabled
    }
}
// GetPrimaryDns gets the primaryDns property value. 
func (m *TeamworkNetworkConfiguration) GetPrimaryDns()(*string) {
    if m == nil {
        return nil
    } else {
        return m.primaryDns
    }
}
// GetSecondaryDns gets the secondaryDns property value. 
func (m *TeamworkNetworkConfiguration) GetSecondaryDns()(*string) {
    if m == nil {
        return nil
    } else {
        return m.secondaryDns
    }
}
// GetSubnetMask gets the subnetMask property value. 
func (m *TeamworkNetworkConfiguration) GetSubnetMask()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subnetMask
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkNetworkConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["defaultGateway"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultGateway(val)
        }
        return nil
    }
    res["domainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
        }
        return nil
    }
    res["hostName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostName(val)
        }
        return nil
    }
    res["ipAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["isDhcpEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDhcpEnabled(val)
        }
        return nil
    }
    res["isPCPortEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPCPortEnabled(val)
        }
        return nil
    }
    res["primaryDns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryDns(val)
        }
        return nil
    }
    res["secondaryDns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecondaryDns(val)
        }
        return nil
    }
    res["subnetMask"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *TeamworkNetworkConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkNetworkConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.additionalData = value
    }
}
// SetDefaultGateway sets the defaultGateway property value. 
func (m *TeamworkNetworkConfiguration) SetDefaultGateway(value *string)() {
    if m != nil {
        m.defaultGateway = value
    }
}
// SetDomainName sets the domainName property value. 
func (m *TeamworkNetworkConfiguration) SetDomainName(value *string)() {
    if m != nil {
        m.domainName = value
    }
}
// SetHostName sets the hostName property value. 
func (m *TeamworkNetworkConfiguration) SetHostName(value *string)() {
    if m != nil {
        m.hostName = value
    }
}
// SetIpAddress sets the ipAddress property value. 
func (m *TeamworkNetworkConfiguration) SetIpAddress(value *string)() {
    if m != nil {
        m.ipAddress = value
    }
}
// SetIsDhcpEnabled sets the isDhcpEnabled property value. 
func (m *TeamworkNetworkConfiguration) SetIsDhcpEnabled(value *bool)() {
    if m != nil {
        m.isDhcpEnabled = value
    }
}
// SetIsPCPortEnabled sets the isPCPortEnabled property value. 
func (m *TeamworkNetworkConfiguration) SetIsPCPortEnabled(value *bool)() {
    if m != nil {
        m.isPCPortEnabled = value
    }
}
// SetPrimaryDns sets the primaryDns property value. 
func (m *TeamworkNetworkConfiguration) SetPrimaryDns(value *string)() {
    if m != nil {
        m.primaryDns = value
    }
}
// SetSecondaryDns sets the secondaryDns property value. 
func (m *TeamworkNetworkConfiguration) SetSecondaryDns(value *string)() {
    if m != nil {
        m.secondaryDns = value
    }
}
// SetSubnetMask sets the subnetMask property value. 
func (m *TeamworkNetworkConfiguration) SetSubnetMask(value *string)() {
    if m != nil {
        m.subnetMask = value
    }
}
