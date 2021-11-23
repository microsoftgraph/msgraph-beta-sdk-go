package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPC 
type CloudPC struct {
    Entity
    // 
    aadDeviceId *string;
    // The Cloud PC display name.
    displayName *string;
    // The date and time when the grace period ends and reprovisioning/deprovisioning happens. Required only if status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    gracePeriodEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Name of the OS image that's on the Cloud PC.
    imageDisplayName *string;
    // 
    lastLoginResult *CloudPcLoginResult;
    // The Cloud PC's last modified date and time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    lastRemoteActionResult *CloudPcRemoteActionResult;
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
    // 
    servicePlanType *CloudPcServicePlanType;
    // Status of the Cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod, deprovisioning, failed.
    status *CloudPcStatus;
    // The details of the Cloud PC status.
    statusDetails *CloudPcStatusDetails;
    // The user principal name (UPN) of the user assigned to the Cloud PC.
    userPrincipalName *string;
}
// NewCloudPC instantiates a new cloudPC and sets the default values.
func NewCloudPC()(*CloudPC) {
    m := &CloudPC{
        Entity: *NewEntity(),
    }
    return m
}
// GetAadDeviceId gets the aadDeviceId property value. 
func (m *CloudPC) GetAadDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.aadDeviceId
    }
}
// GetDisplayName gets the displayName property value. The Cloud PC display name.
func (m *CloudPC) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetGracePeriodEndDateTime gets the gracePeriodEndDateTime property value. The date and time when the grace period ends and reprovisioning/deprovisioning happens. Required only if status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) GetGracePeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.gracePeriodEndDateTime
    }
}
// GetImageDisplayName gets the imageDisplayName property value. Name of the OS image that's on the Cloud PC.
func (m *CloudPC) GetImageDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.imageDisplayName
    }
}
// GetLastLoginResult gets the lastLoginResult property value. 
func (m *CloudPC) GetLastLoginResult()(*CloudPcLoginResult) {
    if m == nil {
        return nil
    } else {
        return m.lastLoginResult
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The Cloud PC's last modified date and time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLastRemoteActionResult gets the lastRemoteActionResult property value. 
func (m *CloudPC) GetLastRemoteActionResult()(*CloudPcRemoteActionResult) {
    if m == nil {
        return nil
    } else {
        return m.lastRemoteActionResult
    }
}
// GetManagedDeviceId gets the managedDeviceId property value. The Cloud PC’s Intune device ID.
func (m *CloudPC) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// GetManagedDeviceName gets the managedDeviceName property value. The Cloud PC’s Intune device name.
func (m *CloudPC) GetManagedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceName
    }
}
// GetOnPremisesConnectionName gets the onPremisesConnectionName property value. The on-premises connection that is applied during provisioning of Cloud PCs.
func (m *CloudPC) GetOnPremisesConnectionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesConnectionName
    }
}
// GetProvisioningPolicyId gets the provisioningPolicyId property value. The Cloud PC's provisioning policy ID.
func (m *CloudPC) GetProvisioningPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisioningPolicyId
    }
}
// GetProvisioningPolicyName gets the provisioningPolicyName property value. The provisioning policy that is applied during provisioning of Cloud PCs.
func (m *CloudPC) GetProvisioningPolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisioningPolicyName
    }
}
// GetServicePlanId gets the servicePlanId property value. The Cloud PC's service plan ID.
func (m *CloudPC) GetServicePlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanId
    }
}
// GetServicePlanName gets the servicePlanName property value. The Cloud PC's service plan name.
func (m *CloudPC) GetServicePlanName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanName
    }
}
// GetServicePlanType gets the servicePlanType property value. 
func (m *CloudPC) GetServicePlanType()(*CloudPcServicePlanType) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanType
    }
}
// GetStatus gets the status property value. Status of the Cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod, deprovisioning, failed.
func (m *CloudPC) GetStatus()(*CloudPcStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetStatusDetails gets the statusDetails property value. The details of the Cloud PC status.
func (m *CloudPC) GetStatusDetails()(*CloudPcStatusDetails) {
    if m == nil {
        return nil
    } else {
        return m.statusDetails
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the Cloud PC.
func (m *CloudPC) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPC) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aadDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAadDeviceId(val)
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
    res["lastLoginResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcLoginResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastLoginResult(val.(*CloudPcLoginResult))
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
    res["lastRemoteActionResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcRemoteActionResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRemoteActionResult(val.(*CloudPcRemoteActionResult))
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
    res["servicePlanType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcServicePlanType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CloudPcServicePlanType)
            m.SetServicePlanType(&cast)
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
// Serialize serializes information the current object
func (m *CloudPC) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("aadDeviceId", m.GetAadDeviceId())
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
        err = writer.WriteObjectValue("lastLoginResult", m.GetLastLoginResult())
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
        err = writer.WriteObjectValue("lastRemoteActionResult", m.GetLastRemoteActionResult())
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
    if m.GetServicePlanType() != nil {
        cast := m.GetServicePlanType().String()
        err = writer.WriteStringValue("servicePlanType", &cast)
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
// SetAadDeviceId sets the aadDeviceId property value. 
func (m *CloudPC) SetAadDeviceId(value *string)() {
    m.aadDeviceId = value
}
// SetDisplayName sets the displayName property value. The Cloud PC display name.
func (m *CloudPC) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGracePeriodEndDateTime sets the gracePeriodEndDateTime property value. The date and time when the grace period ends and reprovisioning/deprovisioning happens. Required only if status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) SetGracePeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.gracePeriodEndDateTime = value
}
// SetImageDisplayName sets the imageDisplayName property value. Name of the OS image that's on the Cloud PC.
func (m *CloudPC) SetImageDisplayName(value *string)() {
    m.imageDisplayName = value
}
// SetLastLoginResult sets the lastLoginResult property value. 
func (m *CloudPC) SetLastLoginResult(value *CloudPcLoginResult)() {
    m.lastLoginResult = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The Cloud PC's last modified date and time. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLastRemoteActionResult sets the lastRemoteActionResult property value. 
func (m *CloudPC) SetLastRemoteActionResult(value *CloudPcRemoteActionResult)() {
    m.lastRemoteActionResult = value
}
// SetManagedDeviceId sets the managedDeviceId property value. The Cloud PC’s Intune device ID.
func (m *CloudPC) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetManagedDeviceName sets the managedDeviceName property value. The Cloud PC’s Intune device name.
func (m *CloudPC) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// SetOnPremisesConnectionName sets the onPremisesConnectionName property value. The on-premises connection that is applied during provisioning of Cloud PCs.
func (m *CloudPC) SetOnPremisesConnectionName(value *string)() {
    m.onPremisesConnectionName = value
}
// SetProvisioningPolicyId sets the provisioningPolicyId property value. The Cloud PC's provisioning policy ID.
func (m *CloudPC) SetProvisioningPolicyId(value *string)() {
    m.provisioningPolicyId = value
}
// SetProvisioningPolicyName sets the provisioningPolicyName property value. The provisioning policy that is applied during provisioning of Cloud PCs.
func (m *CloudPC) SetProvisioningPolicyName(value *string)() {
    m.provisioningPolicyName = value
}
// SetServicePlanId sets the servicePlanId property value. The Cloud PC's service plan ID.
func (m *CloudPC) SetServicePlanId(value *string)() {
    m.servicePlanId = value
}
// SetServicePlanName sets the servicePlanName property value. The Cloud PC's service plan name.
func (m *CloudPC) SetServicePlanName(value *string)() {
    m.servicePlanName = value
}
// SetServicePlanType sets the servicePlanType property value. 
func (m *CloudPC) SetServicePlanType(value *CloudPcServicePlanType)() {
    m.servicePlanType = value
}
// SetStatus sets the status property value. Status of the Cloud PC. Possible values are: notProvisioned, provisioning, provisioned, upgrading, inGracePeriod, deprovisioning, failed.
func (m *CloudPC) SetStatus(value *CloudPcStatus)() {
    m.status = value
}
// SetStatusDetails sets the statusDetails property value. The details of the Cloud PC status.
func (m *CloudPC) SetStatusDetails(value *CloudPcStatusDetails)() {
    m.statusDetails = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the Cloud PC.
func (m *CloudPC) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
