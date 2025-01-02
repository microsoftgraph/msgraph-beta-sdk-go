package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPC struct {
    Entity
}
// NewCloudPC instantiates a new CloudPC and sets the default values.
func NewCloudPC()(*CloudPC) {
    m := &CloudPC{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPCFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPCFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPC(), nil
}
// GetAadDeviceId gets the aadDeviceId property value. The Microsoft Entra device ID of the Cloud PC.
// returns a *string when successful
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
// GetAllotmentDisplayName gets the allotmentDisplayName property value. The allotment name divides tenant licenses into smaller batches or groups that help restrict the number of licenses available for use in a specific assignment. When the provisioningType is dedicated, the allotment name is null. Read-only.
// returns a *string when successful
func (m *CloudPC) GetAllotmentDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("allotmentDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectionSetting gets the connectionSetting property value. The connection setting of the Cloud PC. Possible values: enableSingleSignOn. Read Only.
// returns a CloudPcConnectionSettingable when successful
func (m *CloudPC) GetConnectionSetting()(CloudPcConnectionSettingable) {
    val, err := m.GetBackingStore().Get("connectionSetting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcConnectionSettingable)
    }
    return nil
}
// GetConnectionSettings gets the connectionSettings property value. The connectionSettings property
// returns a CloudPcConnectionSettingsable when successful
func (m *CloudPC) GetConnectionSettings()(CloudPcConnectionSettingsable) {
    val, err := m.GetBackingStore().Get("connectionSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcConnectionSettingsable)
    }
    return nil
}
// GetConnectivityResult gets the connectivityResult property value. The connectivity health check result of a Cloud PC, including the updated timestamp and whether the Cloud PC can be connected.
// returns a CloudPcConnectivityResultable when successful
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
// GetDeviceRegionName gets the deviceRegionName property value. The name of the geographical region where the Cloud PC is currently provisioned. For example, westus3, eastus2, and southeastasia. Read-only.
// returns a *string when successful
func (m *CloudPC) GetDeviceRegionName()(*string) {
    val, err := m.GetBackingStore().Get("deviceRegionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisasterRecoveryCapability gets the disasterRecoveryCapability property value. The disaster recovery status of the Cloud PC, including the primary region, secondary region, and capability type. The default value is null that indicates that the disaster recovery setting is disabled. To receive a response with the disasterRecoveryCapability property, $select and $filter it by disasterRecoveryCapability/{subProperty} in the request URL. For more information, see Example 4: List Cloud PCs filtered by disaster recovery capability type. Read-only.
// returns a CloudPcDisasterRecoveryCapabilityable when successful
func (m *CloudPC) GetDisasterRecoveryCapability()(CloudPcDisasterRecoveryCapabilityable) {
    val, err := m.GetBackingStore().Get("disasterRecoveryCapability")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcDisasterRecoveryCapabilityable)
    }
    return nil
}
// GetDiskEncryptionState gets the diskEncryptionState property value. The disk encryption applied to the Cloud PC. Possible values: notAvailable, notEncrypted, encryptedUsingPlatformManagedKey, encryptedUsingCustomerManagedKey, and unknownFutureValue.
// returns a *CloudPcDiskEncryptionState when successful
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
// returns a *string when successful
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
    res["allotmentDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllotmentDisplayName(val)
        }
        return nil
    }
    res["connectionSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcConnectionSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionSetting(val.(CloudPcConnectionSettingable))
        }
        return nil
    }
    res["connectionSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcConnectionSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionSettings(val.(CloudPcConnectionSettingsable))
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
    res["deviceRegionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceRegionName(val)
        }
        return nil
    }
    res["disasterRecoveryCapability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcDisasterRecoveryCapabilityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisasterRecoveryCapability(val.(CloudPcDisasterRecoveryCapabilityable))
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
    res["frontlineCloudPcAvailability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFrontlineCloudPcAvailability)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrontlineCloudPcAvailability(val.(*FrontlineCloudPcAvailability))
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
    res["statusDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcStatusDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDetail(val.(CloudPcStatusDetailable))
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
// GetFrontlineCloudPcAvailability gets the frontlineCloudPcAvailability property value. The frontlineCloudPcAvailability property
// returns a *FrontlineCloudPcAvailability when successful
func (m *CloudPC) GetFrontlineCloudPcAvailability()(*FrontlineCloudPcAvailability) {
    val, err := m.GetBackingStore().Get("frontlineCloudPcAvailability")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FrontlineCloudPcAvailability)
    }
    return nil
}
// GetGracePeriodEndDateTime gets the gracePeriodEndDateTime property value. The date and time when the grace period ends and reprovisioning or deprovisioning happens. Required only if the status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
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
// returns a *string when successful
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
// returns a CloudPcLoginResultable when successful
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
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The last modified date and time of the Cloud PC. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
// returns a *Time when successful
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
// GetLastRemoteActionResult gets the lastRemoteActionResult property value. The last remote action result of the enterprise Cloud PCs. The supported remote actions are: Reboot, Rename, Reprovision, Restore, Troubleshoot.
// returns a CloudPcRemoteActionResultable when successful
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
// returns a *string when successful
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
// returns a *string when successful
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
// returns a *string when successful
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
// GetOsVersion gets the osVersion property value. The version of the operating system (OS) to provision on Cloud PCs. Possible values are: windows10, windows11, unknownFutureValue.
// returns a *CloudPcOperatingSystem when successful
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
// returns a []CloudPcPartnerAgentInstallResultable when successful
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
// GetPowerState gets the powerState property value. The power state of a Cloud PC. The possible values are: running, poweredOff, unknown. This property only supports shift work Cloud PCs.
// returns a *CloudPcPowerState when successful
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
// returns a *string when successful
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
// returns a *string when successful
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
// GetProvisioningType gets the provisioningType property value. The type of licenses to be used when provisioning Cloud PCs using this policy. Possible values are: dedicated, shared, unknownFutureValue,sharedByUser, sharedByEntraGroup. You must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: sharedByUser, sharedByEntraGroup. The default value is dedicated. CAUTION: The shared member is deprecated and will stop returning on April 30, 2027； in the future, use the sharedByUser member.
// returns a *CloudPcProvisioningType when successful
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
// GetScopeIds gets the scopeIds property value. The scopeIds property
// returns a []string when successful
func (m *CloudPC) GetScopeIds()([]string) {
    val, err := m.GetBackingStore().Get("scopeIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetServicePlanId gets the servicePlanId property value. The service plan ID of the Cloud PC.
// returns a *string when successful
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
// returns a *string when successful
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
// returns a *CloudPcServicePlanType when successful
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
// returns a *CloudPcStatus when successful
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
// GetStatusDetail gets the statusDetail property value. Indicates the detailed status associated with Cloud PC, including error/warning code, error/warning message, additionalInformation. For example, { 'code': 'internalServerError', 'message': 'There was an error during the Cloud PC upgrade. Please contact support.', 'additionalInformation': null }.
// returns a CloudPcStatusDetailable when successful
func (m *CloudPC) GetStatusDetail()(CloudPcStatusDetailable) {
    val, err := m.GetBackingStore().Get("statusDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcStatusDetailable)
    }
    return nil
}
// GetStatusDetails gets the statusDetails property value. The details of the Cloud PC status. For example, { 'code': 'internalServerError', 'message': 'There was an error during the Cloud PC upgrade. Please contact support.', 'additionalInformation': null }. This property is deprecated and will no longer be supported effective August 31, 2024. Use statusDetail instead.
// returns a CloudPcStatusDetailsable when successful
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
// GetUserAccountType gets the userAccountType property value. The account type of the user on provisioned Cloud PCs. Possible values are: standardUser, administrator, unknownFutureValue.
// returns a *CloudPcUserAccountType when successful
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
// returns a *string when successful
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
        err = writer.WriteStringValue("allotmentDisplayName", m.GetAllotmentDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("connectionSetting", m.GetConnectionSetting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("connectionSettings", m.GetConnectionSettings())
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
    {
        err = writer.WriteStringValue("deviceRegionName", m.GetDeviceRegionName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("disasterRecoveryCapability", m.GetDisasterRecoveryCapability())
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
    if m.GetFrontlineCloudPcAvailability() != nil {
        cast := (*m.GetFrontlineCloudPcAvailability()).String()
        err = writer.WriteStringValue("frontlineCloudPcAvailability", &cast)
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
    if m.GetScopeIds() != nil {
        err = writer.WriteCollectionOfStringValues("scopeIds", m.GetScopeIds())
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
        err = writer.WriteObjectValue("statusDetail", m.GetStatusDetail())
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
// SetAadDeviceId sets the aadDeviceId property value. The Microsoft Entra device ID of the Cloud PC.
func (m *CloudPC) SetAadDeviceId(value *string)() {
    err := m.GetBackingStore().Set("aadDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetAllotmentDisplayName sets the allotmentDisplayName property value. The allotment name divides tenant licenses into smaller batches or groups that help restrict the number of licenses available for use in a specific assignment. When the provisioningType is dedicated, the allotment name is null. Read-only.
func (m *CloudPC) SetAllotmentDisplayName(value *string)() {
    err := m.GetBackingStore().Set("allotmentDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionSetting sets the connectionSetting property value. The connection setting of the Cloud PC. Possible values: enableSingleSignOn. Read Only.
func (m *CloudPC) SetConnectionSetting(value CloudPcConnectionSettingable)() {
    err := m.GetBackingStore().Set("connectionSetting", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionSettings sets the connectionSettings property value. The connectionSettings property
func (m *CloudPC) SetConnectionSettings(value CloudPcConnectionSettingsable)() {
    err := m.GetBackingStore().Set("connectionSettings", value)
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
// SetDeviceRegionName sets the deviceRegionName property value. The name of the geographical region where the Cloud PC is currently provisioned. For example, westus3, eastus2, and southeastasia. Read-only.
func (m *CloudPC) SetDeviceRegionName(value *string)() {
    err := m.GetBackingStore().Set("deviceRegionName", value)
    if err != nil {
        panic(err)
    }
}
// SetDisasterRecoveryCapability sets the disasterRecoveryCapability property value. The disaster recovery status of the Cloud PC, including the primary region, secondary region, and capability type. The default value is null that indicates that the disaster recovery setting is disabled. To receive a response with the disasterRecoveryCapability property, $select and $filter it by disasterRecoveryCapability/{subProperty} in the request URL. For more information, see Example 4: List Cloud PCs filtered by disaster recovery capability type. Read-only.
func (m *CloudPC) SetDisasterRecoveryCapability(value CloudPcDisasterRecoveryCapabilityable)() {
    err := m.GetBackingStore().Set("disasterRecoveryCapability", value)
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
// SetFrontlineCloudPcAvailability sets the frontlineCloudPcAvailability property value. The frontlineCloudPcAvailability property
func (m *CloudPC) SetFrontlineCloudPcAvailability(value *FrontlineCloudPcAvailability)() {
    err := m.GetBackingStore().Set("frontlineCloudPcAvailability", value)
    if err != nil {
        panic(err)
    }
}
// SetGracePeriodEndDateTime sets the gracePeriodEndDateTime property value. The date and time when the grace period ends and reprovisioning or deprovisioning happens. Required only if the status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
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
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The last modified date and time of the Cloud PC. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014, is 2014-01-01T00:00:00Z.
func (m *CloudPC) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastRemoteActionResult sets the lastRemoteActionResult property value. The last remote action result of the enterprise Cloud PCs. The supported remote actions are: Reboot, Rename, Reprovision, Restore, Troubleshoot.
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
// SetOsVersion sets the osVersion property value. The version of the operating system (OS) to provision on Cloud PCs. Possible values are: windows10, windows11, unknownFutureValue.
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
// SetPowerState sets the powerState property value. The power state of a Cloud PC. The possible values are: running, poweredOff, unknown. This property only supports shift work Cloud PCs.
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
// SetProvisioningType sets the provisioningType property value. The type of licenses to be used when provisioning Cloud PCs using this policy. Possible values are: dedicated, shared, unknownFutureValue,sharedByUser, sharedByEntraGroup. You must use the Prefer: include-unknown-enum-members request header to get the following values from this evolvable enum: sharedByUser, sharedByEntraGroup. The default value is dedicated. CAUTION: The shared member is deprecated and will stop returning on April 30, 2027； in the future, use the sharedByUser member.
func (m *CloudPC) SetProvisioningType(value *CloudPcProvisioningType)() {
    err := m.GetBackingStore().Set("provisioningType", value)
    if err != nil {
        panic(err)
    }
}
// SetScopeIds sets the scopeIds property value. The scopeIds property
func (m *CloudPC) SetScopeIds(value []string)() {
    err := m.GetBackingStore().Set("scopeIds", value)
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
// SetStatusDetail sets the statusDetail property value. Indicates the detailed status associated with Cloud PC, including error/warning code, error/warning message, additionalInformation. For example, { 'code': 'internalServerError', 'message': 'There was an error during the Cloud PC upgrade. Please contact support.', 'additionalInformation': null }.
func (m *CloudPC) SetStatusDetail(value CloudPcStatusDetailable)() {
    err := m.GetBackingStore().Set("statusDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetStatusDetails sets the statusDetails property value. The details of the Cloud PC status. For example, { 'code': 'internalServerError', 'message': 'There was an error during the Cloud PC upgrade. Please contact support.', 'additionalInformation': null }. This property is deprecated and will no longer be supported effective August 31, 2024. Use statusDetail instead.
func (m *CloudPC) SetStatusDetails(value CloudPcStatusDetailsable)() {
    err := m.GetBackingStore().Set("statusDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetUserAccountType sets the userAccountType property value. The account type of the user on provisioned Cloud PCs. Possible values are: standardUser, administrator, unknownFutureValue.
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
type CloudPCable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAadDeviceId()(*string)
    GetAllotmentDisplayName()(*string)
    GetConnectionSetting()(CloudPcConnectionSettingable)
    GetConnectionSettings()(CloudPcConnectionSettingsable)
    GetConnectivityResult()(CloudPcConnectivityResultable)
    GetDeviceRegionName()(*string)
    GetDisasterRecoveryCapability()(CloudPcDisasterRecoveryCapabilityable)
    GetDiskEncryptionState()(*CloudPcDiskEncryptionState)
    GetDisplayName()(*string)
    GetFrontlineCloudPcAvailability()(*FrontlineCloudPcAvailability)
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
    GetScopeIds()([]string)
    GetServicePlanId()(*string)
    GetServicePlanName()(*string)
    GetServicePlanType()(*CloudPcServicePlanType)
    GetStatus()(*CloudPcStatus)
    GetStatusDetail()(CloudPcStatusDetailable)
    GetStatusDetails()(CloudPcStatusDetailsable)
    GetUserAccountType()(*CloudPcUserAccountType)
    GetUserPrincipalName()(*string)
    SetAadDeviceId(value *string)()
    SetAllotmentDisplayName(value *string)()
    SetConnectionSetting(value CloudPcConnectionSettingable)()
    SetConnectionSettings(value CloudPcConnectionSettingsable)()
    SetConnectivityResult(value CloudPcConnectivityResultable)()
    SetDeviceRegionName(value *string)()
    SetDisasterRecoveryCapability(value CloudPcDisasterRecoveryCapabilityable)()
    SetDiskEncryptionState(value *CloudPcDiskEncryptionState)()
    SetDisplayName(value *string)()
    SetFrontlineCloudPcAvailability(value *FrontlineCloudPcAvailability)()
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
    SetScopeIds(value []string)()
    SetServicePlanId(value *string)()
    SetServicePlanName(value *string)()
    SetServicePlanType(value *CloudPcServicePlanType)()
    SetStatus(value *CloudPcStatus)()
    SetStatusDetail(value CloudPcStatusDetailable)()
    SetStatusDetails(value CloudPcStatusDetailsable)()
    SetUserAccountType(value *CloudPcUserAccountType)()
    SetUserPrincipalName(value *string)()
}
