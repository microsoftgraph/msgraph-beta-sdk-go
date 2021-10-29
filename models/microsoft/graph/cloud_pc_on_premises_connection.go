package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPcOnPremisesConnection struct {
    Entity
    // The fully qualified domain name (FQDN) of the Active Directory domain you want to join.
    adDomainName *string;
    // The password associated with adDomainUsername.
    adDomainPassword *string;
    // The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com.
    adDomainUsername *string;
    // The display name for the on-premises connection.
    displayName *string;
    // The status of the most recent health check done on the on-premises connection. For example, if status is 'passed', the on-premises connection has passed all checks run by the service. Possible values are: pending, running, passed, failed, unknownFutureValue. Read-only.
    healthCheckStatus *CloudPcOnPremisesConnectionStatus;
    // The details of the connection's health checks and the corresponding results. Returned only on $select.For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an on-premises connection, including healthCheckStatusDetails. Read-only.
    healthCheckStatusDetails *CloudPcOnPremisesConnectionStatusDetails;
    // When true, the on-premises connection is in use. When false, the connection is not in use. You cannot delete a connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an on-premises connection, including healthCheckStatusDetails. Read-only.
    inUse *bool;
    // The organizational unit (OU) in which the computer account is created. If left null, the OU that’s configured as the default (a well-known computer object container) in your Active Directory domain (OU) is used. Optional.
    organizationalUnit *string;
    // The ID of the target resource group. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}'.
    resourceGroupId *string;
    // The ID of the target subnet. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}'.
    subnetId *string;
    // The ID of the target Azure subscription that’s associated with your tenant.
    subscriptionId *string;
    // The name of the target Azure subscription. Read-only.
    subscriptionName *string;
    // 
    type_escaped *CloudPcOnPremisesConnectionType;
    // The ID of the target virtual network. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}'.
    virtualNetworkId *string;
}
// Instantiates a new cloudPcOnPremisesConnection and sets the default values.
func NewCloudPcOnPremisesConnection()(*CloudPcOnPremisesConnection) {
    m := &CloudPcOnPremisesConnection{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the adDomainName property value. The fully qualified domain name (FQDN) of the Active Directory domain you want to join.
func (m *CloudPcOnPremisesConnection) GetAdDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adDomainName
    }
}
// Gets the adDomainPassword property value. The password associated with adDomainUsername.
func (m *CloudPcOnPremisesConnection) GetAdDomainPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adDomainPassword
    }
}
// Gets the adDomainUsername property value. The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com.
func (m *CloudPcOnPremisesConnection) GetAdDomainUsername()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adDomainUsername
    }
}
// Gets the displayName property value. The display name for the on-premises connection.
func (m *CloudPcOnPremisesConnection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the healthCheckStatus property value. The status of the most recent health check done on the on-premises connection. For example, if status is 'passed', the on-premises connection has passed all checks run by the service. Possible values are: pending, running, passed, failed, unknownFutureValue. Read-only.
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatus()(*CloudPcOnPremisesConnectionStatus) {
    if m == nil {
        return nil
    } else {
        return m.healthCheckStatus
    }
}
// Gets the healthCheckStatusDetails property value. The details of the connection's health checks and the corresponding results. Returned only on $select.For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an on-premises connection, including healthCheckStatusDetails. Read-only.
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatusDetails()(*CloudPcOnPremisesConnectionStatusDetails) {
    if m == nil {
        return nil
    } else {
        return m.healthCheckStatusDetails
    }
}
// Gets the inUse property value. When true, the on-premises connection is in use. When false, the connection is not in use. You cannot delete a connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an on-premises connection, including healthCheckStatusDetails. Read-only.
func (m *CloudPcOnPremisesConnection) GetInUse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.inUse
    }
}
// Gets the organizationalUnit property value. The organizational unit (OU) in which the computer account is created. If left null, the OU that’s configured as the default (a well-known computer object container) in your Active Directory domain (OU) is used. Optional.
func (m *CloudPcOnPremisesConnection) GetOrganizationalUnit()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizationalUnit
    }
}
// Gets the resourceGroupId property value. The ID of the target resource group. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}'.
func (m *CloudPcOnPremisesConnection) GetResourceGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceGroupId
    }
}
// Gets the subnetId property value. The ID of the target subnet. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}'.
func (m *CloudPcOnPremisesConnection) GetSubnetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subnetId
    }
}
// Gets the subscriptionId property value. The ID of the target Azure subscription that’s associated with your tenant.
func (m *CloudPcOnPremisesConnection) GetSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionId
    }
}
// Gets the subscriptionName property value. The name of the target Azure subscription. Read-only.
func (m *CloudPcOnPremisesConnection) GetSubscriptionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionName
    }
}
// Gets the type_escaped property value. 
func (m *CloudPcOnPremisesConnection) GetType_escaped()(*CloudPcOnPremisesConnectionType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the virtualNetworkId property value. The ID of the target virtual network. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}'.
func (m *CloudPcOnPremisesConnection) GetVirtualNetworkId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.virtualNetworkId
    }
}
// The deserialization information for the current model
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionType)
        if err != nil {
            return err
        }
        cast := val.(CloudPcOnPremisesConnectionType)
        m.SetType_escaped(&cast)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err = writer.WriteStringValue("type_escaped", &cast)
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
// Sets the adDomainName property value. The fully qualified domain name (FQDN) of the Active Directory domain you want to join.
// Parameters:
//  - value : Value to set for the adDomainName property.
func (m *CloudPcOnPremisesConnection) SetAdDomainName(value *string)() {
    m.adDomainName = value
}
// Sets the adDomainPassword property value. The password associated with adDomainUsername.
// Parameters:
//  - value : Value to set for the adDomainPassword property.
func (m *CloudPcOnPremisesConnection) SetAdDomainPassword(value *string)() {
    m.adDomainPassword = value
}
// Sets the adDomainUsername property value. The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com.
// Parameters:
//  - value : Value to set for the adDomainUsername property.
func (m *CloudPcOnPremisesConnection) SetAdDomainUsername(value *string)() {
    m.adDomainUsername = value
}
// Sets the displayName property value. The display name for the on-premises connection.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CloudPcOnPremisesConnection) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the healthCheckStatus property value. The status of the most recent health check done on the on-premises connection. For example, if status is 'passed', the on-premises connection has passed all checks run by the service. Possible values are: pending, running, passed, failed, unknownFutureValue. Read-only.
// Parameters:
//  - value : Value to set for the healthCheckStatus property.
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatus(value *CloudPcOnPremisesConnectionStatus)() {
    m.healthCheckStatus = value
}
// Sets the healthCheckStatusDetails property value. The details of the connection's health checks and the corresponding results. Returned only on $select.For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an on-premises connection, including healthCheckStatusDetails. Read-only.
// Parameters:
//  - value : Value to set for the healthCheckStatusDetails property.
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatusDetails(value *CloudPcOnPremisesConnectionStatusDetails)() {
    m.healthCheckStatusDetails = value
}
// Sets the inUse property value. When true, the on-premises connection is in use. When false, the connection is not in use. You cannot delete a connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an on-premises connection, including healthCheckStatusDetails. Read-only.
// Parameters:
//  - value : Value to set for the inUse property.
func (m *CloudPcOnPremisesConnection) SetInUse(value *bool)() {
    m.inUse = value
}
// Sets the organizationalUnit property value. The organizational unit (OU) in which the computer account is created. If left null, the OU that’s configured as the default (a well-known computer object container) in your Active Directory domain (OU) is used. Optional.
// Parameters:
//  - value : Value to set for the organizationalUnit property.
func (m *CloudPcOnPremisesConnection) SetOrganizationalUnit(value *string)() {
    m.organizationalUnit = value
}
// Sets the resourceGroupId property value. The ID of the target resource group. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}'.
// Parameters:
//  - value : Value to set for the resourceGroupId property.
func (m *CloudPcOnPremisesConnection) SetResourceGroupId(value *string)() {
    m.resourceGroupId = value
}
// Sets the subnetId property value. The ID of the target subnet. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}'.
// Parameters:
//  - value : Value to set for the subnetId property.
func (m *CloudPcOnPremisesConnection) SetSubnetId(value *string)() {
    m.subnetId = value
}
// Sets the subscriptionId property value. The ID of the target Azure subscription that’s associated with your tenant.
// Parameters:
//  - value : Value to set for the subscriptionId property.
func (m *CloudPcOnPremisesConnection) SetSubscriptionId(value *string)() {
    m.subscriptionId = value
}
// Sets the subscriptionName property value. The name of the target Azure subscription. Read-only.
// Parameters:
//  - value : Value to set for the subscriptionName property.
func (m *CloudPcOnPremisesConnection) SetSubscriptionName(value *string)() {
    m.subscriptionName = value
}
// Sets the type_escaped property value. 
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *CloudPcOnPremisesConnection) SetType_escaped(value *CloudPcOnPremisesConnectionType)() {
    m.type_escaped = value
}
// Sets the virtualNetworkId property value. The ID of the target virtual network. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}'.
// Parameters:
//  - value : Value to set for the virtualNetworkId property.
func (m *CloudPcOnPremisesConnection) SetVirtualNetworkId(value *string)() {
    m.virtualNetworkId = value
}
