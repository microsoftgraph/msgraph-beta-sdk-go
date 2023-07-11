package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPC 
type CloudPC struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewCloudPC instantiates a new cloudPC and sets the default values.
func NewCloudPC()(*CloudPC) {
    m := &CloudPC{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPCFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPCFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPC(), nil
}
// GetAadDeviceId gets the aadDeviceId property value. The Azure Active Directory (Azure AD) device ID of the Cloud PC.
func (m *CloudPC) GetAadDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("aadDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectivityResult gets the connectivityResult property value. The connectivity health check result of a Cloud PC, including the updated timestamp and whether the Cloud PC can be connected.
func (m *CloudPC) GetConnectivityResult()(CloudPcConnectivityResultable) {
    val, err := m.GetBackingStore().Get("connectivityResult")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcConnectivityResultable)
    }
    return nil
}
// GetDiskEncryptionState gets the diskEncryptionState property value. The disk encryption applied to the Cloud PC. Possible values: notAvailable, notEncrypted, encryptedUsingPlatformManagedKey, encryptedUsingCustomerManagedKey, and unknownFutureValue.
func (m *CloudPC) GetDiskEncryptionState()(*CloudPcDiskEncryptionState) {
    val, err := m.GetBackingStore().Get("diskEncryptionState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcDiskEncryptionState)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the Cloud PC.
func (m *CloudPC) GetDisplayName()(*string) {
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
func (m *CloudPC) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aadDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAadDeviceId(val)
        }
        return nil
    }
    res["connectivityResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcConnectivityResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectivityResult(val.(CloudPcConnectivityResultable))
        }
        return nil
    }
    res["diskEncryptionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDiskEncryptionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiskEncryptionState(val.(*CloudPcDiskEncryptionState))
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
    res["gracePeriodEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodEndDateTime(val)
        }
        return nil
    }
    res["imageDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageDisplayName(val)
        }
        return nil
    }
    res["lastLoginResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcLoginResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastLoginResult(val.(CloudPcLoginResultable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["lastRemoteActionResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcRemoteActionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRemoteActionResult(val.(CloudPcRemoteActionResultable))
        }
        return nil
    }
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["managedDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["onPremisesConnectionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesConnectionName(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOperatingSystem)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val.(*CloudPcOperatingSystem))
        }
        return nil
    }
    res["partnerAgentInstallResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcPartnerAgentInstallResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcPartnerAgentInstallResultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcPartnerAgentInstallResultable)
                }
            }
            m.SetPartnerAgentInstallResults(res)
        }
        return nil
    }
    res["powerState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcPowerState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPowerState(val.(*CloudPcPowerState))
        }
        return nil
    }
    res["provisioningPolicyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningPolicyId(val)
        }
        return nil
    }
    res["provisioningPolicyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningPolicyName(val)
        }
        return nil
    }
    res["provisioningType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcProvisioningType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningType(val.(*CloudPcProvisioningType))
        }
        return nil
    }
    res["servicePlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanId(val)
        }
        return nil
    }
    res["servicePlanName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanName(val)
        }
        return nil
    }
    res["servicePlanType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcServicePlanType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanType(val.(*CloudPcServicePlanType))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CloudPcStatus))
        }
        return nil
    }
    res["statusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcStatusDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDetails(val.(CloudPcStatusDetailsable))
        }
        return nil
    }
    res["userAccountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcUserAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccountType(val.(*CloudPcUserAccountType))
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetGracePeriodEndDateTime gets the gracePeriodEndDateTime property value. The date and time when the grace period ends and reprovisioning/deprovisioning happens. Required only if the status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) GetGracePeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("gracePeriodEndDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetImageDisplayName gets the imageDisplayName property value. Name of the OS image that's on the Cloud PC.
func (m *CloudPC) GetImageDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("imageDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastLoginResult gets the lastLoginResult property value. The last login result of the Cloud PC. For example, { 'time': '2014-01-01T00:00:00Z'}.
func (m *CloudPC) GetLastLoginResult()(CloudPcLoginResultable) {
    val, err := m.GetBackingStore().Get("lastLoginResult")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcLoginResultable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The last modified date and time of the Cloud PC. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastRemoteActionResult gets the lastRemoteActionResult property value. The last remote action result of the enterprise Cloud PCs. The supported remote actions are: Reboot, Rename, Reprovision, Restore, and Troubleshoot.
func (m *CloudPC) GetLastRemoteActionResult()(CloudPcRemoteActionResultable) {
    val, err := m.GetBackingStore().Get("lastRemoteActionResult")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcRemoteActionResultable)
    }
    return nil
}
// GetManagedDeviceId gets the managedDeviceId property value. The Intune device ID of the Cloud PC.
func (m *CloudPC) GetManagedDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagedDeviceName gets the managedDeviceName property value. The Intune device name of the Cloud PC.
func (m *CloudPC) GetManagedDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnPremisesConnectionName gets the onPremisesConnectionName property value. The Azure network connection that is applied during the provisioning of Cloud PCs.
func (m *CloudPC) GetOnPremisesConnectionName()(*string) {
    val, err := m.GetBackingStore().Get("onPremisesConnectionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. The version of the operating system (OS) to provision on Cloud PCs. Possible values are: windows10, windows11, and unknownFutureValue.
func (m *CloudPC) GetOsVersion()(*CloudPcOperatingSystem) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcOperatingSystem)
    }
    return nil
}
// GetPartnerAgentInstallResults gets the partnerAgentInstallResults property value. The results of every partner agent's installation status on Cloud PC.
func (m *CloudPC) GetPartnerAgentInstallResults()([]CloudPcPartnerAgentInstallResultable) {
    val, err := m.GetBackingStore().Get("partnerAgentInstallResults")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcPartnerAgentInstallResultable)
    }
    return nil
}
// GetPowerState gets the powerState property value. The power state of a Cloud PC. The possible values are: running, poweredOff and unknown. This property only supports shift work Cloud PCs.
func (m *CloudPC) GetPowerState()(*CloudPcPowerState) {
    val, err := m.GetBackingStore().Get("powerState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcPowerState)
    }
    return nil
}
// GetProvisioningPolicyId gets the provisioningPolicyId property value. The provisioning policy ID of the Cloud PC.
func (m *CloudPC) GetProvisioningPolicyId()(*string) {
    val, err := m.GetBackingStore().Get("provisioningPolicyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProvisioningPolicyName gets the provisioningPolicyName property value. The provisioning policy that is applied during the provisioning of Cloud PCs.
func (m *CloudPC) GetProvisioningPolicyName()(*string) {
    val, err := m.GetBackingStore().Get("provisioningPolicyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProvisioningType gets the provisioningType property value. The type of licenses to be used when provisioning Cloud PCs using this policy. Possible values are: dedicated, shared, unknownFutureValue. Default value is dedicated.
func (m *CloudPC) GetProvisioningType()(*CloudPcProvisioningType) {
    val, err := m.GetBackingStore().Get("provisioningType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcProvisioningType)
    }
    return nil
}
// GetServicePlanId gets the servicePlanId property value. The service plan ID of the Cloud PC.
func (m *CloudPC) GetServicePlanId()(*string) {
    val, err := m.GetBackingStore().Get("servicePlanId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServicePlanName gets the servicePlanName property value. The service plan name of the Cloud PC.
func (m *CloudPC) GetServicePlanName()(*string) {
    val, err := m.GetBackingStore().Get("servicePlanName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetServicePlanType gets the servicePlanType property value. The service plan type of the Cloud PC.
func (m *CloudPC) GetServicePlanType()(*CloudPcServicePlanType) {
    val, err := m.GetBackingStore().Get("servicePlanType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcServicePlanType)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *CloudPC) GetStatus()(*CloudPcStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcStatus)
    }
    return nil
}
// GetStatusDetails gets the statusDetails property value. The details of the Cloud PC status.
func (m *CloudPC) GetStatusDetails()(CloudPcStatusDetailsable) {
    val, err := m.GetBackingStore().Get("statusDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcStatusDetailsable)
    }
    return nil
}
// GetUserAccountType gets the userAccountType property value. The account type of the user on provisioned Cloud PCs. Possible values are: standardUser, administrator, and unknownFutureValue.
func (m *CloudPC) GetUserAccountType()(*CloudPcUserAccountType) {
    val, err := m.GetBackingStore().Get("userAccountType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcUserAccountType)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the Cloud PC.
func (m *CloudPC) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPC) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("connectivityResult", m.GetConnectivityResult())
        if err != nil {
            return err
        }
    }
    if m.GetDiskEncryptionState() != nil {
        cast := (*m.GetDiskEncryptionState()).String()
        err = writer.WriteStringValue("diskEncryptionState", &cast)
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
    if m.GetOsVersion() != nil {
        cast := (*m.GetOsVersion()).String()
        err = writer.WriteStringValue("osVersion", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPartnerAgentInstallResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPartnerAgentInstallResults()))
        for i, v := range m.GetPartnerAgentInstallResults() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("partnerAgentInstallResults", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPowerState() != nil {
        cast := (*m.GetPowerState()).String()
        err = writer.WriteStringValue("powerState", &cast)
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
    if m.GetProvisioningType() != nil {
        cast := (*m.GetProvisioningType()).String()
        err = writer.WriteStringValue("provisioningType", &cast)
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
        cast := (*m.GetServicePlanType()).String()
        err = writer.WriteStringValue("servicePlanType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
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
    if m.GetUserAccountType() != nil {
        cast := (*m.GetUserAccountType()).String()
        err = writer.WriteStringValue("userAccountType", &cast)
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
// SetAadDeviceId sets the aadDeviceId property value. The Azure Active Directory (Azure AD) device ID of the Cloud PC.
func (m *CloudPC) SetAadDeviceId(value *string)() {
    err := m.GetBackingStore().Set("aadDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectivityResult sets the connectivityResult property value. The connectivity health check result of a Cloud PC, including the updated timestamp and whether the Cloud PC can be connected.
func (m *CloudPC) SetConnectivityResult(value CloudPcConnectivityResultable)() {
    err := m.GetBackingStore().Set("connectivityResult", value)
    if err != nil {
        panic(err)
    }
}
// SetDiskEncryptionState sets the diskEncryptionState property value. The disk encryption applied to the Cloud PC. Possible values: notAvailable, notEncrypted, encryptedUsingPlatformManagedKey, encryptedUsingCustomerManagedKey, and unknownFutureValue.
func (m *CloudPC) SetDiskEncryptionState(value *CloudPcDiskEncryptionState)() {
    err := m.GetBackingStore().Set("diskEncryptionState", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the Cloud PC.
func (m *CloudPC) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetGracePeriodEndDateTime sets the gracePeriodEndDateTime property value. The date and time when the grace period ends and reprovisioning/deprovisioning happens. Required only if the status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) SetGracePeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("gracePeriodEndDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetImageDisplayName sets the imageDisplayName property value. Name of the OS image that's on the Cloud PC.
func (m *CloudPC) SetImageDisplayName(value *string)() {
    err := m.GetBackingStore().Set("imageDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastLoginResult sets the lastLoginResult property value. The last login result of the Cloud PC. For example, { 'time': '2014-01-01T00:00:00Z'}.
func (m *CloudPC) SetLastLoginResult(value CloudPcLoginResultable)() {
    err := m.GetBackingStore().Set("lastLoginResult", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The last modified date and time of the Cloud PC. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastRemoteActionResult sets the lastRemoteActionResult property value. The last remote action result of the enterprise Cloud PCs. The supported remote actions are: Reboot, Rename, Reprovision, Restore, and Troubleshoot.
func (m *CloudPC) SetLastRemoteActionResult(value CloudPcRemoteActionResultable)() {
    err := m.GetBackingStore().Set("lastRemoteActionResult", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. The Intune device ID of the Cloud PC.
func (m *CloudPC) SetManagedDeviceId(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceName sets the managedDeviceName property value. The Intune device name of the Cloud PC.
func (m *CloudPC) SetManagedDeviceName(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetOnPremisesConnectionName sets the onPremisesConnectionName property value. The Azure network connection that is applied during the provisioning of Cloud PCs.
func (m *CloudPC) SetOnPremisesConnectionName(value *string)() {
    err := m.GetBackingStore().Set("onPremisesConnectionName", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. The version of the operating system (OS) to provision on Cloud PCs. Possible values are: windows10, windows11, and unknownFutureValue.
func (m *CloudPC) SetOsVersion(value *CloudPcOperatingSystem)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetPartnerAgentInstallResults sets the partnerAgentInstallResults property value. The results of every partner agent's installation status on Cloud PC.
func (m *CloudPC) SetPartnerAgentInstallResults(value []CloudPcPartnerAgentInstallResultable)() {
    err := m.GetBackingStore().Set("partnerAgentInstallResults", value)
    if err != nil {
        panic(err)
    }
}
// SetPowerState sets the powerState property value. The power state of a Cloud PC. The possible values are: running, poweredOff and unknown. This property only supports shift work Cloud PCs.
func (m *CloudPC) SetPowerState(value *CloudPcPowerState)() {
    err := m.GetBackingStore().Set("powerState", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisioningPolicyId sets the provisioningPolicyId property value. The provisioning policy ID of the Cloud PC.
func (m *CloudPC) SetProvisioningPolicyId(value *string)() {
    err := m.GetBackingStore().Set("provisioningPolicyId", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisioningPolicyName sets the provisioningPolicyName property value. The provisioning policy that is applied during the provisioning of Cloud PCs.
func (m *CloudPC) SetProvisioningPolicyName(value *string)() {
    err := m.GetBackingStore().Set("provisioningPolicyName", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisioningType sets the provisioningType property value. The type of licenses to be used when provisioning Cloud PCs using this policy. Possible values are: dedicated, shared, unknownFutureValue. Default value is dedicated.
func (m *CloudPC) SetProvisioningType(value *CloudPcProvisioningType)() {
    err := m.GetBackingStore().Set("provisioningType", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePlanId sets the servicePlanId property value. The service plan ID of the Cloud PC.
func (m *CloudPC) SetServicePlanId(value *string)() {
    err := m.GetBackingStore().Set("servicePlanId", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePlanName sets the servicePlanName property value. The service plan name of the Cloud PC.
func (m *CloudPC) SetServicePlanName(value *string)() {
    err := m.GetBackingStore().Set("servicePlanName", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePlanType sets the servicePlanType property value. The service plan type of the Cloud PC.
func (m *CloudPC) SetServicePlanType(value *CloudPcServicePlanType)() {
    err := m.GetBackingStore().Set("servicePlanType", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *CloudPC) SetStatus(value *CloudPcStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetStatusDetails sets the statusDetails property value. The details of the Cloud PC status.
func (m *CloudPC) SetStatusDetails(value CloudPcStatusDetailsable)() {
    err := m.GetBackingStore().Set("statusDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetUserAccountType sets the userAccountType property value. The account type of the user on provisioned Cloud PCs. Possible values are: standardUser, administrator, and unknownFutureValue.
func (m *CloudPC) SetUserAccountType(value *CloudPcUserAccountType)() {
    err := m.GetBackingStore().Set("userAccountType", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the Cloud PC.
func (m *CloudPC) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// CloudPCable 
type CloudPCable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAadDeviceId()(*string)
    GetConnectivityResult()(CloudPcConnectivityResultable)
    GetDiskEncryptionState()(*CloudPcDiskEncryptionState)
    GetDisplayName()(*string)
    GetGracePeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetImageDisplayName()(*string)
    GetLastLoginResult()(CloudPcLoginResultable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastRemoteActionResult()(CloudPcRemoteActionResultable)
    GetManagedDeviceId()(*string)
    GetManagedDeviceName()(*string)
    GetOnPremisesConnectionName()(*string)
    GetOsVersion()(*CloudPcOperatingSystem)
    GetPartnerAgentInstallResults()([]CloudPcPartnerAgentInstallResultable)
    GetPowerState()(*CloudPcPowerState)
    GetProvisioningPolicyId()(*string)
    GetProvisioningPolicyName()(*string)
    GetProvisioningType()(*CloudPcProvisioningType)
    GetServicePlanId()(*string)
    GetServicePlanName()(*string)
    GetServicePlanType()(*CloudPcServicePlanType)
    GetStatus()(*CloudPcStatus)
    GetStatusDetails()(CloudPcStatusDetailsable)
    GetUserAccountType()(*CloudPcUserAccountType)
    GetUserPrincipalName()(*string)
    SetAadDeviceId(value *string)()
    SetConnectivityResult(value CloudPcConnectivityResultable)()
    SetDiskEncryptionState(value *CloudPcDiskEncryptionState)()
    SetDisplayName(value *string)()
    SetGracePeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetImageDisplayName(value *string)()
    SetLastLoginResult(value CloudPcLoginResultable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastRemoteActionResult(value CloudPcRemoteActionResultable)()
    SetManagedDeviceId(value *string)()
    SetManagedDeviceName(value *string)()
    SetOnPremisesConnectionName(value *string)()
    SetOsVersion(value *CloudPcOperatingSystem)()
    SetPartnerAgentInstallResults(value []CloudPcPartnerAgentInstallResultable)()
    SetPowerState(value *CloudPcPowerState)()
    SetProvisioningPolicyId(value *string)()
    SetProvisioningPolicyName(value *string)()
    SetProvisioningType(value *CloudPcProvisioningType)()
    SetServicePlanId(value *string)()
    SetServicePlanName(value *string)()
    SetServicePlanType(value *CloudPcServicePlanType)()
    SetStatus(value *CloudPcStatus)()
    SetStatusDetails(value CloudPcStatusDetailsable)()
    SetUserAccountType(value *CloudPcUserAccountType)()
    SetUserPrincipalName(value *string)()
}
