package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcOnPremisesConnection 
type CloudPcOnPremisesConnection struct {
    Entity
    // The fully qualified domain name (FQDN) of the Active Directory domain you want to join. Optional.
    adDomainName *string
    // The password associated with adDomainUsername.
    adDomainPassword *string
    // The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com. Optional.
    adDomainUsername *string
    // The interface URL of the partner service's resource that links to this Azure network connection. Returned only on $select.
    alternateResourceUrl *string
    // The display name for the Azure network connection.
    displayName *string
    // The status of the most recent health check done on the Azure network connection. For example, if status is passed, the Azure network connection has passed all checks run by the service. Possible values are: pending, running, passed, failed, unknownFutureValue. Read-only.
    healthCheckStatus *CloudPcOnPremisesConnectionStatus
    // The details of the connection's health checks and the corresponding results. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetails. Read-only.
    healthCheckStatusDetails CloudPcOnPremisesConnectionStatusDetailsable
    // When true, the Azure network connection is in use. When false, the connection is not in use. You cannot delete a connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetails. Read-only.
    inUse *bool
    // Specifies which services manage the Azure network connection. Possible values are: windows365, devBox, unknownFutureValue. Read-only.
    managedBy *CloudPcManagementService
    // The organizational unit (OU) in which the computer account is created. If left null, the OU that’s configured as the default (a well-known computer object container) in your Active Directory domain (OU) is used. Optional.
    organizationalUnit *string
    // The ID of the target resource group. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}.
    resourceGroupId *string
    // The ID of the target subnet. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}.
    subnetId *string
    // The ID of the target Azure subscription that’s associated with your tenant.
    subscriptionId *string
    // The name of the target Azure subscription. Read-only.
    subscriptionName *string
    // Specifies how the provisioned Cloud PC will be joined to Azure Active Directory. Default value is hybridAzureADJoin. Possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
    type_escaped *CloudPcOnPremisesConnectionType
    // The ID of the target virtual network. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}.
    virtualNetworkId *string
}
// NewCloudPcOnPremisesConnection instantiates a new CloudPcOnPremisesConnection and sets the default values.
func NewCloudPcOnPremisesConnection()(*CloudPcOnPremisesConnection) {
    m := &CloudPcOnPremisesConnection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcOnPremisesConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcOnPremisesConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcOnPremisesConnection(), nil
}
// GetAdDomainName gets the adDomainName property value. The fully qualified domain name (FQDN) of the Active Directory domain you want to join. Optional.
func (m *CloudPcOnPremisesConnection) GetAdDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adDomainName
    }
}
// GetAdDomainPassword gets the adDomainPassword property value. The password associated with adDomainUsername.
func (m *CloudPcOnPremisesConnection) GetAdDomainPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adDomainPassword
    }
}
// GetAdDomainUsername gets the adDomainUsername property value. The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com. Optional.
func (m *CloudPcOnPremisesConnection) GetAdDomainUsername()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adDomainUsername
    }
}
// GetAlternateResourceUrl gets the alternateResourceUrl property value. The interface URL of the partner service's resource that links to this Azure network connection. Returned only on $select.
func (m *CloudPcOnPremisesConnection) GetAlternateResourceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternateResourceUrl
    }
}
// GetDisplayName gets the displayName property value. The display name for the Azure network connection.
func (m *CloudPcOnPremisesConnection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcOnPremisesConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["adDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdDomainName(val)
        }
        return nil
    }
    res["adDomainPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdDomainPassword(val)
        }
        return nil
    }
    res["adDomainUsername"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdDomainUsername(val)
        }
        return nil
    }
    res["alternateResourceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateResourceUrl(val)
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
    res["healthCheckStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthCheckStatus(val.(*CloudPcOnPremisesConnectionStatus))
        }
        return nil
    }
    res["healthCheckStatusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcOnPremisesConnectionStatusDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthCheckStatusDetails(val.(CloudPcOnPremisesConnectionStatusDetailsable))
        }
        return nil
    }
    res["inUse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInUse(val)
        }
        return nil
    }
    res["managedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcManagementService)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBy(val.(*CloudPcManagementService))
        }
        return nil
    }
    res["organizationalUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalUnit(val)
        }
        return nil
    }
    res["resourceGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceGroupId(val)
        }
        return nil
    }
    res["subnetId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetId(val)
        }
        return nil
    }
    res["subscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionId(val)
        }
        return nil
    }
    res["subscriptionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionName(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*CloudPcOnPremisesConnectionType))
        }
        return nil
    }
    res["virtualNetworkId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualNetworkId(val)
        }
        return nil
    }
    return res
}
// GetHealthCheckStatus gets the healthCheckStatus property value. The status of the most recent health check done on the Azure network connection. For example, if status is passed, the Azure network connection has passed all checks run by the service. Possible values are: pending, running, passed, failed, unknownFutureValue. Read-only.
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatus()(*CloudPcOnPremisesConnectionStatus) {
    if m == nil {
        return nil
    } else {
        return m.healthCheckStatus
    }
}
// GetHealthCheckStatusDetails gets the healthCheckStatusDetails property value. The details of the connection's health checks and the corresponding results. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetails. Read-only.
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatusDetails()(CloudPcOnPremisesConnectionStatusDetailsable) {
    if m == nil {
        return nil
    } else {
        return m.healthCheckStatusDetails
    }
}
// GetInUse gets the inUse property value. When true, the Azure network connection is in use. When false, the connection is not in use. You cannot delete a connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetails. Read-only.
func (m *CloudPcOnPremisesConnection) GetInUse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.inUse
    }
}
// GetManagedBy gets the managedBy property value. Specifies which services manage the Azure network connection. Possible values are: windows365, devBox, unknownFutureValue. Read-only.
func (m *CloudPcOnPremisesConnection) GetManagedBy()(*CloudPcManagementService) {
    if m == nil {
        return nil
    } else {
        return m.managedBy
    }
}
// GetOrganizationalUnit gets the organizationalUnit property value. The organizational unit (OU) in which the computer account is created. If left null, the OU that’s configured as the default (a well-known computer object container) in your Active Directory domain (OU) is used. Optional.
func (m *CloudPcOnPremisesConnection) GetOrganizationalUnit()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizationalUnit
    }
}
// GetResourceGroupId gets the resourceGroupId property value. The ID of the target resource group. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}.
func (m *CloudPcOnPremisesConnection) GetResourceGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceGroupId
    }
}
// GetSubnetId gets the subnetId property value. The ID of the target subnet. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}.
func (m *CloudPcOnPremisesConnection) GetSubnetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subnetId
    }
}
// GetSubscriptionId gets the subscriptionId property value. The ID of the target Azure subscription that’s associated with your tenant.
func (m *CloudPcOnPremisesConnection) GetSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionId
    }
}
// GetSubscriptionName gets the subscriptionName property value. The name of the target Azure subscription. Read-only.
func (m *CloudPcOnPremisesConnection) GetSubscriptionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionName
    }
}
// GetType gets the type property value. Specifies how the provisioned Cloud PC will be joined to Azure Active Directory. Default value is hybridAzureADJoin. Possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
func (m *CloudPcOnPremisesConnection) GetType()(*CloudPcOnPremisesConnectionType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetVirtualNetworkId gets the virtualNetworkId property value. The ID of the target virtual network. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}.
func (m *CloudPcOnPremisesConnection) GetVirtualNetworkId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.virtualNetworkId
    }
}
// Serialize serializes information the current object
func (m *CloudPcOnPremisesConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("alternateResourceUrl", m.GetAlternateResourceUrl())
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
        cast := (*m.GetHealthCheckStatus()).String()
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
    if m.GetManagedBy() != nil {
        cast := (*m.GetManagedBy()).String()
        err = writer.WriteStringValue("managedBy", &cast)
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
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
// SetAdDomainName sets the adDomainName property value. The fully qualified domain name (FQDN) of the Active Directory domain you want to join. Optional.
func (m *CloudPcOnPremisesConnection) SetAdDomainName(value *string)() {
    if m != nil {
        m.adDomainName = value
    }
}
// SetAdDomainPassword sets the adDomainPassword property value. The password associated with adDomainUsername.
func (m *CloudPcOnPremisesConnection) SetAdDomainPassword(value *string)() {
    if m != nil {
        m.adDomainPassword = value
    }
}
// SetAdDomainUsername sets the adDomainUsername property value. The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com. Optional.
func (m *CloudPcOnPremisesConnection) SetAdDomainUsername(value *string)() {
    if m != nil {
        m.adDomainUsername = value
    }
}
// SetAlternateResourceUrl sets the alternateResourceUrl property value. The interface URL of the partner service's resource that links to this Azure network connection. Returned only on $select.
func (m *CloudPcOnPremisesConnection) SetAlternateResourceUrl(value *string)() {
    if m != nil {
        m.alternateResourceUrl = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the Azure network connection.
func (m *CloudPcOnPremisesConnection) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetHealthCheckStatus sets the healthCheckStatus property value. The status of the most recent health check done on the Azure network connection. For example, if status is passed, the Azure network connection has passed all checks run by the service. Possible values are: pending, running, passed, failed, unknownFutureValue. Read-only.
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatus(value *CloudPcOnPremisesConnectionStatus)() {
    if m != nil {
        m.healthCheckStatus = value
    }
}
// SetHealthCheckStatusDetails sets the healthCheckStatusDetails property value. The details of the connection's health checks and the corresponding results. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetails. Read-only.
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatusDetails(value CloudPcOnPremisesConnectionStatusDetailsable)() {
    if m != nil {
        m.healthCheckStatusDetails = value
    }
}
// SetInUse sets the inUse property value. When true, the Azure network connection is in use. When false, the connection is not in use. You cannot delete a connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetails. Read-only.
func (m *CloudPcOnPremisesConnection) SetInUse(value *bool)() {
    if m != nil {
        m.inUse = value
    }
}
// SetManagedBy sets the managedBy property value. Specifies which services manage the Azure network connection. Possible values are: windows365, devBox, unknownFutureValue. Read-only.
func (m *CloudPcOnPremisesConnection) SetManagedBy(value *CloudPcManagementService)() {
    if m != nil {
        m.managedBy = value
    }
}
// SetOrganizationalUnit sets the organizationalUnit property value. The organizational unit (OU) in which the computer account is created. If left null, the OU that’s configured as the default (a well-known computer object container) in your Active Directory domain (OU) is used. Optional.
func (m *CloudPcOnPremisesConnection) SetOrganizationalUnit(value *string)() {
    if m != nil {
        m.organizationalUnit = value
    }
}
// SetResourceGroupId sets the resourceGroupId property value. The ID of the target resource group. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}.
func (m *CloudPcOnPremisesConnection) SetResourceGroupId(value *string)() {
    if m != nil {
        m.resourceGroupId = value
    }
}
// SetSubnetId sets the subnetId property value. The ID of the target subnet. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}.
func (m *CloudPcOnPremisesConnection) SetSubnetId(value *string)() {
    if m != nil {
        m.subnetId = value
    }
}
// SetSubscriptionId sets the subscriptionId property value. The ID of the target Azure subscription that’s associated with your tenant.
func (m *CloudPcOnPremisesConnection) SetSubscriptionId(value *string)() {
    if m != nil {
        m.subscriptionId = value
    }
}
// SetSubscriptionName sets the subscriptionName property value. The name of the target Azure subscription. Read-only.
func (m *CloudPcOnPremisesConnection) SetSubscriptionName(value *string)() {
    if m != nil {
        m.subscriptionName = value
    }
}
// SetType sets the type property value. Specifies how the provisioned Cloud PC will be joined to Azure Active Directory. Default value is hybridAzureADJoin. Possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
func (m *CloudPcOnPremisesConnection) SetType(value *CloudPcOnPremisesConnectionType)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetVirtualNetworkId sets the virtualNetworkId property value. The ID of the target virtual network. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}.
func (m *CloudPcOnPremisesConnection) SetVirtualNetworkId(value *string)() {
    if m != nil {
        m.virtualNetworkId = value
    }
}
