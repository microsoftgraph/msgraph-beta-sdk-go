package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcOnPremisesConnection 
type CloudPcOnPremisesConnection struct {
    Entity
}
// NewCloudPcOnPremisesConnection instantiates a new cloudPcOnPremisesConnection and sets the default values.
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
    val, err := m.GetBackingStore().Get("adDomainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdDomainPassword gets the adDomainPassword property value. The password associated with adDomainUsername.
func (m *CloudPcOnPremisesConnection) GetAdDomainPassword()(*string) {
    val, err := m.GetBackingStore().Get("adDomainPassword")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdDomainUsername gets the adDomainUsername property value. The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com. Optional.
func (m *CloudPcOnPremisesConnection) GetAdDomainUsername()(*string) {
    val, err := m.GetBackingStore().Get("adDomainUsername")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAlternateResourceUrl gets the alternateResourceUrl property value. The interface URL of the partner service's resource that links to this Azure network connection. Returned only on $select.
func (m *CloudPcOnPremisesConnection) GetAlternateResourceUrl()(*string) {
    val, err := m.GetBackingStore().Get("alternateResourceUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectionType gets the connectionType property value. The connectionType property
func (m *CloudPcOnPremisesConnection) GetConnectionType()(*CloudPcOnPremisesConnection_connectionType) {
    val, err := m.GetBackingStore().Get("connectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcOnPremisesConnection_connectionType)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the Azure network connection.
func (m *CloudPcOnPremisesConnection) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    res["connectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnection_connectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val.(*CloudPcOnPremisesConnection_connectionType))
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
    res["healthCheckStatusDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcOnPremisesConnectionStatusDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthCheckStatusDetail(val.(CloudPcOnPremisesConnectionStatusDetailable))
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
    res["scopeIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetScopeIds(res)
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
        val, err := n.GetEnumValue(ParseCloudPcOnPremisesConnection_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*CloudPcOnPremisesConnection_type))
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
    res["virtualNetworkLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualNetworkLocation(val)
        }
        return nil
    }
    return res
}
// GetHealthCheckStatus gets the healthCheckStatus property value. The healthCheckStatus property
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatus()(*CloudPcOnPremisesConnectionStatus) {
    val, err := m.GetBackingStore().Get("healthCheckStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcOnPremisesConnectionStatus)
    }
    return nil
}
// GetHealthCheckStatusDetail gets the healthCheckStatusDetail property value. The healthCheckStatusDetail property
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatusDetail()(CloudPcOnPremisesConnectionStatusDetailable) {
    val, err := m.GetBackingStore().Get("healthCheckStatusDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcOnPremisesConnectionStatusDetailable)
    }
    return nil
}
// GetHealthCheckStatusDetails gets the healthCheckStatusDetails property value. The details of the connection's health checks and the corresponding results. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetails. Read-only.
func (m *CloudPcOnPremisesConnection) GetHealthCheckStatusDetails()(CloudPcOnPremisesConnectionStatusDetailsable) {
    val, err := m.GetBackingStore().Get("healthCheckStatusDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcOnPremisesConnectionStatusDetailsable)
    }
    return nil
}
// GetInUse gets the inUse property value. When true, the Azure network connection is in use. When false, the connection isn't in use. You can't delete a connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetails. Read-only.
func (m *CloudPcOnPremisesConnection) GetInUse()(*bool) {
    val, err := m.GetBackingStore().Get("inUse")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManagedBy gets the managedBy property value. The managedBy property
func (m *CloudPcOnPremisesConnection) GetManagedBy()(*CloudPcManagementService) {
    val, err := m.GetBackingStore().Get("managedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcManagementService)
    }
    return nil
}
// GetOrganizationalUnit gets the organizationalUnit property value. The organizational unit (OU) in which the computer account is created. If left null, the OU that’s configured as the default (a well-known computer object container) in your Active Directory domain (OU) is used. Optional.
func (m *CloudPcOnPremisesConnection) GetOrganizationalUnit()(*string) {
    val, err := m.GetBackingStore().Get("organizationalUnit")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResourceGroupId gets the resourceGroupId property value. The ID of the target resource group. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}.
func (m *CloudPcOnPremisesConnection) GetResourceGroupId()(*string) {
    val, err := m.GetBackingStore().Get("resourceGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetScopeIds gets the scopeIds property value. The scopeIds property
func (m *CloudPcOnPremisesConnection) GetScopeIds()([]string) {
    val, err := m.GetBackingStore().Get("scopeIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSubnetId gets the subnetId property value. The ID of the target subnet. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}.
func (m *CloudPcOnPremisesConnection) GetSubnetId()(*string) {
    val, err := m.GetBackingStore().Get("subnetId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubscriptionId gets the subscriptionId property value. The ID of the target Azure subscription that’s associated with your tenant.
func (m *CloudPcOnPremisesConnection) GetSubscriptionId()(*string) {
    val, err := m.GetBackingStore().Get("subscriptionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubscriptionName gets the subscriptionName property value. The name of the target Azure subscription. Read-only.
func (m *CloudPcOnPremisesConnection) GetSubscriptionName()(*string) {
    val, err := m.GetBackingStore().Get("subscriptionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTypeEscaped gets the type property value. Specifies how the provisioned Cloud PC is joined to Microsoft Entra ID. Default value is hybridAzureADJoin. Possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
func (m *CloudPcOnPremisesConnection) GetTypeEscaped()(*CloudPcOnPremisesConnection_type) {
    val, err := m.GetBackingStore().Get("typeEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcOnPremisesConnection_type)
    }
    return nil
}
// GetVirtualNetworkId gets the virtualNetworkId property value. The ID of the target virtual network. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}.
func (m *CloudPcOnPremisesConnection) GetVirtualNetworkId()(*string) {
    val, err := m.GetBackingStore().Get("virtualNetworkId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVirtualNetworkLocation gets the virtualNetworkLocation property value. Indicates resource location of the virtual target network. Read-only, computed value.
func (m *CloudPcOnPremisesConnection) GetVirtualNetworkLocation()(*string) {
    val, err := m.GetBackingStore().Get("virtualNetworkLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    if m.GetConnectionType() != nil {
        cast := (*m.GetConnectionType()).String()
        err = writer.WriteStringValue("connectionType", &cast)
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
        err = writer.WriteObjectValue("healthCheckStatusDetail", m.GetHealthCheckStatusDetail())
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
    if m.GetScopeIds() != nil {
        err = writer.WriteCollectionOfStringValues("scopeIds", m.GetScopeIds())
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
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
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
    {
        err = writer.WriteStringValue("virtualNetworkLocation", m.GetVirtualNetworkLocation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdDomainName sets the adDomainName property value. The fully qualified domain name (FQDN) of the Active Directory domain you want to join. Optional.
func (m *CloudPcOnPremisesConnection) SetAdDomainName(value *string)() {
    err := m.GetBackingStore().Set("adDomainName", value)
    if err != nil {
        panic(err)
    }
}
// SetAdDomainPassword sets the adDomainPassword property value. The password associated with adDomainUsername.
func (m *CloudPcOnPremisesConnection) SetAdDomainPassword(value *string)() {
    err := m.GetBackingStore().Set("adDomainPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetAdDomainUsername sets the adDomainUsername property value. The username of an Active Directory account (user or service account) that has permissions to create computer objects in Active Directory. Required format: admin@contoso.com. Optional.
func (m *CloudPcOnPremisesConnection) SetAdDomainUsername(value *string)() {
    err := m.GetBackingStore().Set("adDomainUsername", value)
    if err != nil {
        panic(err)
    }
}
// SetAlternateResourceUrl sets the alternateResourceUrl property value. The interface URL of the partner service's resource that links to this Azure network connection. Returned only on $select.
func (m *CloudPcOnPremisesConnection) SetAlternateResourceUrl(value *string)() {
    err := m.GetBackingStore().Set("alternateResourceUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionType sets the connectionType property value. The connectionType property
func (m *CloudPcOnPremisesConnection) SetConnectionType(value *CloudPcOnPremisesConnection_connectionType)() {
    err := m.GetBackingStore().Set("connectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the Azure network connection.
func (m *CloudPcOnPremisesConnection) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthCheckStatus sets the healthCheckStatus property value. The healthCheckStatus property
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatus(value *CloudPcOnPremisesConnectionStatus)() {
    err := m.GetBackingStore().Set("healthCheckStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthCheckStatusDetail sets the healthCheckStatusDetail property value. The healthCheckStatusDetail property
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatusDetail(value CloudPcOnPremisesConnectionStatusDetailable)() {
    err := m.GetBackingStore().Set("healthCheckStatusDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthCheckStatusDetails sets the healthCheckStatusDetails property value. The details of the connection's health checks and the corresponding results. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetails. Read-only.
func (m *CloudPcOnPremisesConnection) SetHealthCheckStatusDetails(value CloudPcOnPremisesConnectionStatusDetailsable)() {
    err := m.GetBackingStore().Set("healthCheckStatusDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetInUse sets the inUse property value. When true, the Azure network connection is in use. When false, the connection isn't in use. You can't delete a connection that’s in use. Returned only on $select. For an example that shows how to get the inUse property, see Example 2: Get the selected properties of an Azure network connection, including healthCheckStatusDetails. Read-only.
func (m *CloudPcOnPremisesConnection) SetInUse(value *bool)() {
    err := m.GetBackingStore().Set("inUse", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedBy sets the managedBy property value. The managedBy property
func (m *CloudPcOnPremisesConnection) SetManagedBy(value *CloudPcManagementService)() {
    err := m.GetBackingStore().Set("managedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationalUnit sets the organizationalUnit property value. The organizational unit (OU) in which the computer account is created. If left null, the OU that’s configured as the default (a well-known computer object container) in your Active Directory domain (OU) is used. Optional.
func (m *CloudPcOnPremisesConnection) SetOrganizationalUnit(value *string)() {
    err := m.GetBackingStore().Set("organizationalUnit", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceGroupId sets the resourceGroupId property value. The ID of the target resource group. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}.
func (m *CloudPcOnPremisesConnection) SetResourceGroupId(value *string)() {
    err := m.GetBackingStore().Set("resourceGroupId", value)
    if err != nil {
        panic(err)
    }
}
// SetScopeIds sets the scopeIds property value. The scopeIds property
func (m *CloudPcOnPremisesConnection) SetScopeIds(value []string)() {
    err := m.GetBackingStore().Set("scopeIds", value)
    if err != nil {
        panic(err)
    }
}
// SetSubnetId sets the subnetId property value. The ID of the target subnet. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkId}/subnets/{subnetName}.
func (m *CloudPcOnPremisesConnection) SetSubnetId(value *string)() {
    err := m.GetBackingStore().Set("subnetId", value)
    if err != nil {
        panic(err)
    }
}
// SetSubscriptionId sets the subscriptionId property value. The ID of the target Azure subscription that’s associated with your tenant.
func (m *CloudPcOnPremisesConnection) SetSubscriptionId(value *string)() {
    err := m.GetBackingStore().Set("subscriptionId", value)
    if err != nil {
        panic(err)
    }
}
// SetSubscriptionName sets the subscriptionName property value. The name of the target Azure subscription. Read-only.
func (m *CloudPcOnPremisesConnection) SetSubscriptionName(value *string)() {
    err := m.GetBackingStore().Set("subscriptionName", value)
    if err != nil {
        panic(err)
    }
}
// SetTypeEscaped sets the type property value. Specifies how the provisioned Cloud PC is joined to Microsoft Entra ID. Default value is hybridAzureADJoin. Possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
func (m *CloudPcOnPremisesConnection) SetTypeEscaped(value *CloudPcOnPremisesConnection_type)() {
    err := m.GetBackingStore().Set("typeEscaped", value)
    if err != nil {
        panic(err)
    }
}
// SetVirtualNetworkId sets the virtualNetworkId property value. The ID of the target virtual network. Required format: /subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}.
func (m *CloudPcOnPremisesConnection) SetVirtualNetworkId(value *string)() {
    err := m.GetBackingStore().Set("virtualNetworkId", value)
    if err != nil {
        panic(err)
    }
}
// SetVirtualNetworkLocation sets the virtualNetworkLocation property value. Indicates resource location of the virtual target network. Read-only, computed value.
func (m *CloudPcOnPremisesConnection) SetVirtualNetworkLocation(value *string)() {
    err := m.GetBackingStore().Set("virtualNetworkLocation", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcOnPremisesConnectionable 
type CloudPcOnPremisesConnectionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdDomainName()(*string)
    GetAdDomainPassword()(*string)
    GetAdDomainUsername()(*string)
    GetAlternateResourceUrl()(*string)
    GetConnectionType()(*CloudPcOnPremisesConnection_connectionType)
    GetDisplayName()(*string)
    GetHealthCheckStatus()(*CloudPcOnPremisesConnectionStatus)
    GetHealthCheckStatusDetail()(CloudPcOnPremisesConnectionStatusDetailable)
    GetHealthCheckStatusDetails()(CloudPcOnPremisesConnectionStatusDetailsable)
    GetInUse()(*bool)
    GetManagedBy()(*CloudPcManagementService)
    GetOrganizationalUnit()(*string)
    GetResourceGroupId()(*string)
    GetScopeIds()([]string)
    GetSubnetId()(*string)
    GetSubscriptionId()(*string)
    GetSubscriptionName()(*string)
    GetTypeEscaped()(*CloudPcOnPremisesConnection_type)
    GetVirtualNetworkId()(*string)
    GetVirtualNetworkLocation()(*string)
    SetAdDomainName(value *string)()
    SetAdDomainPassword(value *string)()
    SetAdDomainUsername(value *string)()
    SetAlternateResourceUrl(value *string)()
    SetConnectionType(value *CloudPcOnPremisesConnection_connectionType)()
    SetDisplayName(value *string)()
    SetHealthCheckStatus(value *CloudPcOnPremisesConnectionStatus)()
    SetHealthCheckStatusDetail(value CloudPcOnPremisesConnectionStatusDetailable)()
    SetHealthCheckStatusDetails(value CloudPcOnPremisesConnectionStatusDetailsable)()
    SetInUse(value *bool)()
    SetManagedBy(value *CloudPcManagementService)()
    SetOrganizationalUnit(value *string)()
    SetResourceGroupId(value *string)()
    SetScopeIds(value []string)()
    SetSubnetId(value *string)()
    SetSubscriptionId(value *string)()
    SetSubscriptionName(value *string)()
    SetTypeEscaped(value *CloudPcOnPremisesConnection_type)()
    SetVirtualNetworkId(value *string)()
    SetVirtualNetworkLocation(value *string)()
}
