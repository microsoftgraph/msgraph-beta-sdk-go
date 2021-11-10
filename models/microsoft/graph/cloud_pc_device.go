package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPcDevice struct {
    Entity
    // The status of the cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod, deprovisioning, failed. Required. Read-only.
    cloudPcStatus *string;
    // The display name for the cloud PC. Required. Read-only.
    displayName *string;
    // Date and time the entity was last updated in the multi-tenant management platform. Required. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The managed device identifier for the cloud PC. Optional. Read-only.
    managedDeviceId *string;
    // The managed device display name for the cloud PC. Optional. Read-only.
    managedDeviceName *string;
    // The provisioning policy identifier for the cloud PC. Required. Read-only.
    provisioningPolicyId *string;
    // The service plan name for the cloud PC. Required. Read-only.
    servicePlanName *string;
    // The display name for the managed tenant. Required. Read-only.
    tenantDisplayName *string;
    // The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
    tenantId *string;
    // The user principal name (UPN) of the user assigned to the cloud PC. Required. Read-only.
    userPrincipalName *string;
}
// Instantiates a new cloudPcDevice and sets the default values.
func NewCloudPcDevice()(*CloudPcDevice) {
    m := &CloudPcDevice{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the cloudPcStatus property value. The status of the cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod, deprovisioning, failed. Required. Read-only.
func (m *CloudPcDevice) GetCloudPcStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcStatus
    }
}
// Gets the displayName property value. The display name for the cloud PC. Required. Read-only.
func (m *CloudPcDevice) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Required. Read-only.
func (m *CloudPcDevice) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
// Gets the managedDeviceId property value. The managed device identifier for the cloud PC. Optional. Read-only.
func (m *CloudPcDevice) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// Gets the managedDeviceName property value. The managed device display name for the cloud PC. Optional. Read-only.
func (m *CloudPcDevice) GetManagedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceName
    }
}
// Gets the provisioningPolicyId property value. The provisioning policy identifier for the cloud PC. Required. Read-only.
func (m *CloudPcDevice) GetProvisioningPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisioningPolicyId
    }
}
// Gets the servicePlanName property value. The service plan name for the cloud PC. Required. Read-only.
func (m *CloudPcDevice) GetServicePlanName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanName
    }
}
// Gets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *CloudPcDevice) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
// Gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *CloudPcDevice) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Gets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the cloud PC. Required. Read-only.
func (m *CloudPcDevice) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *CloudPcDevice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cloudPcStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcStatus(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastRefreshedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshedDateTime(val)
        }
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["managedDeviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["provisioningPolicyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningPolicyId(val)
        }
        return nil
    }
    res["servicePlanName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanName(val)
        }
        return nil
    }
    res["tenantDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDisplayName(val)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *CloudPcDevice) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CloudPcDevice) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("cloudPcStatus", m.GetCloudPcStatus())
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
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provisioningPolicyId", m.GetProvisioningPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePlanName", m.GetServicePlanName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDisplayName", m.GetTenantDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the cloudPcStatus property value. The status of the cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod, deprovisioning, failed. Required. Read-only.
// Parameters:
//  - value : Value to set for the cloudPcStatus property.
func (m *CloudPcDevice) SetCloudPcStatus(value *string)() {
    m.cloudPcStatus = value
}
// Sets the displayName property value. The display name for the cloud PC. Required. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CloudPcDevice) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Required. Read-only.
// Parameters:
//  - value : Value to set for the lastRefreshedDateTime property.
func (m *CloudPcDevice) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
// Sets the managedDeviceId property value. The managed device identifier for the cloud PC. Optional. Read-only.
// Parameters:
//  - value : Value to set for the managedDeviceId property.
func (m *CloudPcDevice) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// Sets the managedDeviceName property value. The managed device display name for the cloud PC. Optional. Read-only.
// Parameters:
//  - value : Value to set for the managedDeviceName property.
func (m *CloudPcDevice) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// Sets the provisioningPolicyId property value. The provisioning policy identifier for the cloud PC. Required. Read-only.
// Parameters:
//  - value : Value to set for the provisioningPolicyId property.
func (m *CloudPcDevice) SetProvisioningPolicyId(value *string)() {
    m.provisioningPolicyId = value
}
// Sets the servicePlanName property value. The service plan name for the cloud PC. Required. Read-only.
// Parameters:
//  - value : Value to set for the servicePlanName property.
func (m *CloudPcDevice) SetServicePlanName(value *string)() {
    m.servicePlanName = value
}
// Sets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
// Parameters:
//  - value : Value to set for the tenantDisplayName property.
func (m *CloudPcDevice) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// Sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *CloudPcDevice) SetTenantId(value *string)() {
    m.tenantId = value
}
// Sets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the cloud PC. Required. Read-only.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *CloudPcDevice) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
