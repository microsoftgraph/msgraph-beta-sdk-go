package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPC struct {
    Entity
    // The Cloud PC display name.
    displayName *string;
    // The date and time when the grace period ends and reprovisioning/deprovisioning happens. Required only if status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    gracePeriodEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Name of the OS image that's on the Cloud PC.
    imageDisplayName *string;
    // The Cloud PC's last modified date and time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Cloud PC’s Intune device ID.
    managedDeviceId *string;
    // The Cloud PC’s Intune device name.
    managedDeviceName *string;
    // The on-premises connection that is applied during provisioning of Cloud PCs.
    onPremisesConnectionName *string;
    // The Cloud PC's provisioning policy ID.
    provisioningPolicyId *string;
    // The provisioning policy that is applied during provisioning of Cloud PCs.
    provisioningPolicyName *string;
    // The Cloud PC's service plan ID.
    servicePlanId *string;
    // The Cloud PC's service plan name.
    servicePlanName *string;
    // Status of the Cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod, deprovisioning, failed.
    status *CloudPcStatus;
    // The details of the Cloud PC status.
    statusDetails *CloudPcStatusDetails;
    // The user principal name (UPN) of the user assigned to the Cloud PC.
    userPrincipalName *string;
}
// Instantiates a new cloudPC and sets the default values.
func NewCloudPC()(*CloudPC) {
    m := &CloudPC{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The Cloud PC display name.
func (m *CloudPC) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the gracePeriodEndDateTime property value. The date and time when the grace period ends and reprovisioning/deprovisioning happens. Required only if status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) GetGracePeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.gracePeriodEndDateTime
    }
}
// Gets the imageDisplayName property value. Name of the OS image that's on the Cloud PC.
func (m *CloudPC) GetImageDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.imageDisplayName
    }
}
// Gets the lastModifiedDateTime property value. The Cloud PC's last modified date and time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the managedDeviceId property value. The Cloud PC’s Intune device ID.
func (m *CloudPC) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// Gets the managedDeviceName property value. The Cloud PC’s Intune device name.
func (m *CloudPC) GetManagedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceName
    }
}
// Gets the onPremisesConnectionName property value. The on-premises connection that is applied during provisioning of Cloud PCs.
func (m *CloudPC) GetOnPremisesConnectionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesConnectionName
    }
}
// Gets the provisioningPolicyId property value. The Cloud PC's provisioning policy ID.
func (m *CloudPC) GetProvisioningPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisioningPolicyId
    }
}
// Gets the provisioningPolicyName property value. The provisioning policy that is applied during provisioning of Cloud PCs.
func (m *CloudPC) GetProvisioningPolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisioningPolicyName
    }
}
// Gets the servicePlanId property value. The Cloud PC's service plan ID.
func (m *CloudPC) GetServicePlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanId
    }
}
// Gets the servicePlanName property value. The Cloud PC's service plan name.
func (m *CloudPC) GetServicePlanName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanName
    }
}
// Gets the status property value. Status of the Cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod, deprovisioning, failed.
func (m *CloudPC) GetStatus()(*CloudPcStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the statusDetails property value. The details of the Cloud PC status.
func (m *CloudPC) GetStatusDetails()(*CloudPcStatusDetails) {
    if m == nil {
        return nil
    } else {
        return m.statusDetails
    }
}
// Gets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the Cloud PC.
func (m *CloudPC) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *CloudPC) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["gracePeriodEndDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodEndDateTime(val)
        }
        return nil
    }
    res["imageDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageDisplayName(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
    res["onPremisesConnectionName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesConnectionName(val)
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
    res["provisioningPolicyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningPolicyName(val)
        }
        return nil
    }
    res["servicePlanId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanId(val)
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CloudPcStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["statusDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcStatusDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDetails(val.(*CloudPcStatusDetails))
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
func (m *CloudPC) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CloudPC) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("gracePeriodEndDateTime", m.GetGracePeriodEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("imageDisplayName", m.GetImageDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
        err = writer.WriteStringValue("onPremisesConnectionName", m.GetOnPremisesConnectionName())
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
        err = writer.WriteStringValue("provisioningPolicyName", m.GetProvisioningPolicyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePlanId", m.GetServicePlanId())
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
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("statusDetails", m.GetStatusDetails())
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
// Sets the displayName property value. The Cloud PC display name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CloudPC) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the gracePeriodEndDateTime property value. The date and time when the grace period ends and reprovisioning/deprovisioning happens. Required only if status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the gracePeriodEndDateTime property.
func (m *CloudPC) SetGracePeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.gracePeriodEndDateTime = value
}
// Sets the imageDisplayName property value. Name of the OS image that's on the Cloud PC.
// Parameters:
//  - value : Value to set for the imageDisplayName property.
func (m *CloudPC) SetImageDisplayName(value *string)() {
    m.imageDisplayName = value
}
// Sets the lastModifiedDateTime property value. The Cloud PC's last modified date and time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *CloudPC) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the managedDeviceId property value. The Cloud PC’s Intune device ID.
// Parameters:
//  - value : Value to set for the managedDeviceId property.
func (m *CloudPC) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// Sets the managedDeviceName property value. The Cloud PC’s Intune device name.
// Parameters:
//  - value : Value to set for the managedDeviceName property.
func (m *CloudPC) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// Sets the onPremisesConnectionName property value. The on-premises connection that is applied during provisioning of Cloud PCs.
// Parameters:
//  - value : Value to set for the onPremisesConnectionName property.
func (m *CloudPC) SetOnPremisesConnectionName(value *string)() {
    m.onPremisesConnectionName = value
}
// Sets the provisioningPolicyId property value. The Cloud PC's provisioning policy ID.
// Parameters:
//  - value : Value to set for the provisioningPolicyId property.
func (m *CloudPC) SetProvisioningPolicyId(value *string)() {
    m.provisioningPolicyId = value
}
// Sets the provisioningPolicyName property value. The provisioning policy that is applied during provisioning of Cloud PCs.
// Parameters:
//  - value : Value to set for the provisioningPolicyName property.
func (m *CloudPC) SetProvisioningPolicyName(value *string)() {
    m.provisioningPolicyName = value
}
// Sets the servicePlanId property value. The Cloud PC's service plan ID.
// Parameters:
//  - value : Value to set for the servicePlanId property.
func (m *CloudPC) SetServicePlanId(value *string)() {
    m.servicePlanId = value
}
// Sets the servicePlanName property value. The Cloud PC's service plan name.
// Parameters:
//  - value : Value to set for the servicePlanName property.
func (m *CloudPC) SetServicePlanName(value *string)() {
    m.servicePlanName = value
}
// Sets the status property value. Status of the Cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod, deprovisioning, failed.
// Parameters:
//  - value : Value to set for the status property.
func (m *CloudPC) SetStatus(value *CloudPcStatus)() {
    m.status = value
}
// Sets the statusDetails property value. The details of the Cloud PC status.
// Parameters:
//  - value : Value to set for the statusDetails property.
func (m *CloudPC) SetStatusDetails(value *CloudPcStatusDetails)() {
    m.statusDetails = value
}
// Sets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the Cloud PC.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *CloudPC) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
