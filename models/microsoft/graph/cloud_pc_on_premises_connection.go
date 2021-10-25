package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CloudPcOnPremisesConnection struct {
    Entity
    adDomainName *string;
    adDomainPassword *string;
    adDomainUsername *string;
    displayName *string;
    healthCheckStatus *CloudPcOnPremisesConnectionStatus;
    healthCheckStatusDetails *CloudPcOnPremisesConnectionStatusDetails;
    inUse *bool;
    organizationalUnit *string;
    resourceGroupId *string;
    subnetId *string;
    subscriptionId *string;
    subscriptionName *string;
    virtualNetworkId *string;
}
func NewCloudPcOnPremisesConnection()(*CloudPcOnPremisesConnection) {
    m := &CloudPcOnPremisesConnection{
        Entity: *NewEntity(),
    }
    return m
}
func (m *CloudPcOnPremisesConnection) GetAdDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adDomainName
    }
}
func (m *CloudPcOnPremisesConnection) GetAdDomainPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adDomainPassword
    }
}
func (m *CloudPcOnPremisesConnection) GetAdDomainUsername()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adDomainUsername
    }
}
func (m *CloudPcOnPremisesConnection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatus()(*CloudPcOnPremisesConnectionStatus) {
    if m == nil {
        return nil
    } else {
        return m.healthCheckStatus
    }
}
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatusDetails()(*CloudPcOnPremisesConnectionStatusDetails) {
    if m == nil {
        return nil
    } else {
        return m.healthCheckStatusDetails
    }
}
func (m *CloudPcOnPremisesConnection) GetInUse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.inUse
    }
}
func (m *CloudPcOnPremisesConnection) GetOrganizationalUnit()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizationalUnit
    }
}
func (m *CloudPcOnPremisesConnection) GetResourceGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceGroupId
    }
}
func (m *CloudPcOnPremisesConnection) GetSubnetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subnetId
    }
}
func (m *CloudPcOnPremisesConnection) GetSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionId
    }
}
func (m *CloudPcOnPremisesConnection) GetSubscriptionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionName
    }
}
func (m *CloudPcOnPremisesConnection) GetVirtualNetworkId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.virtualNetworkId
    }
}
func (m *CloudPcOnPremisesConnection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["adDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdDomainName(val)
        return nil
    }
    res["adDomainPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdDomainPassword(val)
        return nil
    }
    res["adDomainUsername"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdDomainUsername(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["healthCheckStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionStatus)
        if err != nil {
            return err
        }
        cast := val.(CloudPcOnPremisesConnectionStatus)
        m.SetHealthCheckStatus(&cast)
        return nil
    }
    res["healthCheckStatusDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcOnPremisesConnectionStatusDetails() })
        if err != nil {
            return err
        }
        m.SetHealthCheckStatusDetails(val.(*CloudPcOnPremisesConnectionStatusDetails))
        return nil
    }
    res["inUse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetInUse(val)
        return nil
    }
    res["organizationalUnit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOrganizationalUnit(val)
        return nil
    }
    res["resourceGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceGroupId(val)
        return nil
    }
    res["subnetId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubnetId(val)
        return nil
    }
    res["subscriptionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubscriptionId(val)
        return nil
    }
    res["subscriptionName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubscriptionName(val)
        return nil
    }
    res["virtualNetworkId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVirtualNetworkId(val)
        return nil
    }
    return res
}
func (m *CloudPcOnPremisesConnection) IsNil()(bool) {
    return m == nil
}
func (m *CloudPcOnPremisesConnection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("adDomainName", m.GetAdDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("adDomainPassword", m.GetAdDomainPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("adDomainUsername", m.GetAdDomainUsername())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetHealthCheckStatus() != nil {
        cast := m.GetHealthCheckStatus().String()
        err = writer.WriteStringValue("healthCheckStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("healthCheckStatusDetails", m.GetHealthCheckStatusDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("inUse", m.GetInUse())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("organizationalUnit", m.GetOrganizationalUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceGroupId", m.GetResourceGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subnetId", m.GetSubnetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subscriptionName", m.GetSubscriptionName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("virtualNetworkId", m.GetVirtualNetworkId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *CloudPcOnPremisesConnection) SetAdDomainName(value *string)() {
    m.adDomainName = value
}
func (m *CloudPcOnPremisesConnection) SetAdDomainPassword(value *string)() {
    m.adDomainPassword = value
}
func (m *CloudPcOnPremisesConnection) SetAdDomainUsername(value *string)() {
    m.adDomainUsername = value
}
func (m *CloudPcOnPremisesConnection) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatus(value *CloudPcOnPremisesConnectionStatus)() {
    m.healthCheckStatus = value
}
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatusDetails(value *CloudPcOnPremisesConnectionStatusDetails)() {
    m.healthCheckStatusDetails = value
}
func (m *CloudPcOnPremisesConnection) SetInUse(value *bool)() {
    m.inUse = value
}
func (m *CloudPcOnPremisesConnection) SetOrganizationalUnit(value *string)() {
    m.organizationalUnit = value
}
func (m *CloudPcOnPremisesConnection) SetResourceGroupId(value *string)() {
    m.resourceGroupId = value
}
func (m *CloudPcOnPremisesConnection) SetSubnetId(value *string)() {
    m.subnetId = value
}
func (m *CloudPcOnPremisesConnection) SetSubscriptionId(value *string)() {
    m.subscriptionId = value
}
func (m *CloudPcOnPremisesConnection) SetSubscriptionName(value *string)() {
    m.subscriptionName = value
}
func (m *CloudPcOnPremisesConnection) SetVirtualNetworkId(value *string)() {
    m.virtualNetworkId = value
}
