package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerGeneralDeviceConfiguration this topic provides descriptions of the declared methods, properties and relationships exposed by the androidDeviceOwnerGeneralDeviceConfiguration resource.
type AndroidDeviceOwnerGeneralDeviceConfiguration struct {
    DeviceConfiguration
}
// NewAndroidDeviceOwnerGeneralDeviceConfiguration instantiates a new androidDeviceOwnerGeneralDeviceConfiguration and sets the default values.
func NewAndroidDeviceOwnerGeneralDeviceConfiguration()(*AndroidDeviceOwnerGeneralDeviceConfiguration) {
    m := &AndroidDeviceOwnerGeneralDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerGeneralDeviceConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidDeviceOwnerGeneralDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerGeneralDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerGeneralDeviceConfiguration(), nil
}
// GetAccountsBlockModification gets the accountsBlockModification property value. Indicates whether or not adding or removing accounts is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetAccountsBlockModification()(*bool) {
    val, err := m.GetBackingStore().Get("accountsBlockModification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAppsAllowInstallFromUnknownSources gets the appsAllowInstallFromUnknownSources property value. Indicates whether or not the user is allowed to enable to unknown sources setting.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetAppsAllowInstallFromUnknownSources()(*bool) {
    val, err := m.GetBackingStore().Get("appsAllowInstallFromUnknownSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAppsAutoUpdatePolicy gets the appsAutoUpdatePolicy property value. Indicates the value of the app auto update policy. Possible values are: notConfigured, userChoice, never, wiFiOnly, always.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetAppsAutoUpdatePolicy()(*AndroidDeviceOwnerAppAutoUpdatePolicyType) {
    val, err := m.GetBackingStore().Get("appsAutoUpdatePolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerAppAutoUpdatePolicyType)
    }
    return nil
}
// GetAppsDefaultPermissionPolicy gets the appsDefaultPermissionPolicy property value. Indicates the permission policy for requests for runtime permissions if one is not defined for the app specifically. Possible values are: deviceDefault, prompt, autoGrant, autoDeny.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetAppsDefaultPermissionPolicy()(*AndroidDeviceOwnerDefaultAppPermissionPolicyType) {
    val, err := m.GetBackingStore().Get("appsDefaultPermissionPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerDefaultAppPermissionPolicyType)
    }
    return nil
}
// GetAppsRecommendSkippingFirstUseHints gets the appsRecommendSkippingFirstUseHints property value. Whether or not to recommend all apps skip any first-time-use hints they may have added.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetAppsRecommendSkippingFirstUseHints()(*bool) {
    val, err := m.GetBackingStore().Get("appsRecommendSkippingFirstUseHints")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAzureAdSharedDeviceDataClearApps gets the azureAdSharedDeviceDataClearApps property value. A list of managed apps that will have their data cleared during a global sign-out in AAD shared device mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetAzureAdSharedDeviceDataClearApps()([]AppListItemable) {
    val, err := m.GetBackingStore().Get("azureAdSharedDeviceDataClearApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppListItemable)
    }
    return nil
}
// GetBluetoothBlockConfiguration gets the bluetoothBlockConfiguration property value. Indicates whether or not to block a user from configuring bluetooth.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetBluetoothBlockConfiguration()(*bool) {
    val, err := m.GetBackingStore().Get("bluetoothBlockConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBluetoothBlockContactSharing gets the bluetoothBlockContactSharing property value. Indicates whether or not to block a user from sharing contacts via bluetooth.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetBluetoothBlockContactSharing()(*bool) {
    val, err := m.GetBackingStore().Get("bluetoothBlockContactSharing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCameraBlocked gets the cameraBlocked property value. Indicates whether or not to disable the use of the camera.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetCameraBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("cameraBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCellularBlockWiFiTethering gets the cellularBlockWiFiTethering property value. Indicates whether or not to block Wi-Fi tethering.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetCellularBlockWiFiTethering()(*bool) {
    val, err := m.GetBackingStore().Get("cellularBlockWiFiTethering")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCertificateCredentialConfigurationDisabled gets the certificateCredentialConfigurationDisabled property value. Indicates whether or not to block users from any certificate credential configuration.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetCertificateCredentialConfigurationDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("certificateCredentialConfigurationDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCrossProfilePoliciesAllowCopyPaste gets the crossProfilePoliciesAllowCopyPaste property value. Indicates whether or not text copied from one profile (personal or work) can be pasted in the other.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetCrossProfilePoliciesAllowCopyPaste()(*bool) {
    val, err := m.GetBackingStore().Get("crossProfilePoliciesAllowCopyPaste")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCrossProfilePoliciesAllowDataSharing gets the crossProfilePoliciesAllowDataSharing property value. Indicates whether data from one profile (personal or work) can be shared with apps in the other profile. Possible values are: notConfigured, crossProfileDataSharingBlocked, dataSharingFromWorkToPersonalBlocked, crossProfileDataSharingAllowed, unkownFutureValue.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetCrossProfilePoliciesAllowDataSharing()(*AndroidDeviceOwnerCrossProfileDataSharing) {
    val, err := m.GetBackingStore().Get("crossProfilePoliciesAllowDataSharing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerCrossProfileDataSharing)
    }
    return nil
}
// GetCrossProfilePoliciesShowWorkContactsInPersonalProfile gets the crossProfilePoliciesShowWorkContactsInPersonalProfile property value. Indicates whether or not contacts stored in work profile are shown in personal profile contact searches/incoming calls.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetCrossProfilePoliciesShowWorkContactsInPersonalProfile()(*bool) {
    val, err := m.GetBackingStore().Get("crossProfilePoliciesShowWorkContactsInPersonalProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDataRoamingBlocked gets the dataRoamingBlocked property value. Indicates whether or not to block a user from data roaming.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetDataRoamingBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("dataRoamingBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDateTimeConfigurationBlocked gets the dateTimeConfigurationBlocked property value. Indicates whether or not to block the user from manually changing the date or time on the device
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetDateTimeConfigurationBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("dateTimeConfigurationBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDetailedHelpText gets the detailedHelpText property value. Represents the customized detailed help text provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetDetailedHelpText()(AndroidDeviceOwnerUserFacingMessageable) {
    val, err := m.GetBackingStore().Get("detailedHelpText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidDeviceOwnerUserFacingMessageable)
    }
    return nil
}
// GetDeviceOwnerLockScreenMessage gets the deviceOwnerLockScreenMessage property value. Represents the customized lock screen message provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetDeviceOwnerLockScreenMessage()(AndroidDeviceOwnerUserFacingMessageable) {
    val, err := m.GetBackingStore().Get("deviceOwnerLockScreenMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidDeviceOwnerUserFacingMessageable)
    }
    return nil
}
// GetEnrollmentProfile gets the enrollmentProfile property value. Android Device Owner Enrollment Profile types.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetEnrollmentProfile()(*AndroidDeviceOwnerEnrollmentProfileType) {
    val, err := m.GetBackingStore().Get("enrollmentProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerEnrollmentProfileType)
    }
    return nil
}
// GetFactoryResetBlocked gets the factoryResetBlocked property value. Indicates whether or not the factory reset option in settings is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetFactoryResetBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("factoryResetBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFactoryResetDeviceAdministratorEmails gets the factoryResetDeviceAdministratorEmails property value. List of Google account emails that will be required to authenticate after a device is factory reset before it can be set up.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetFactoryResetDeviceAdministratorEmails()([]string) {
    val, err := m.GetBackingStore().Get("factoryResetDeviceAdministratorEmails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["accountsBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountsBlockModification(val)
        }
        return nil
    }
    res["appsAllowInstallFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsAllowInstallFromUnknownSources(val)
        }
        return nil
    }
    res["appsAutoUpdatePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerAppAutoUpdatePolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsAutoUpdatePolicy(val.(*AndroidDeviceOwnerAppAutoUpdatePolicyType))
        }
        return nil
    }
    res["appsDefaultPermissionPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerDefaultAppPermissionPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsDefaultPermissionPolicy(val.(*AndroidDeviceOwnerDefaultAppPermissionPolicyType))
        }
        return nil
    }
    res["appsRecommendSkippingFirstUseHints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsRecommendSkippingFirstUseHints(val)
        }
        return nil
    }
    res["azureAdSharedDeviceDataClearApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppListItemable)
                }
            }
            m.SetAzureAdSharedDeviceDataClearApps(res)
        }
        return nil
    }
    res["bluetoothBlockConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetoothBlockConfiguration(val)
        }
        return nil
    }
    res["bluetoothBlockContactSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetoothBlockContactSharing(val)
        }
        return nil
    }
    res["cameraBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameraBlocked(val)
        }
        return nil
    }
    res["cellularBlockWiFiTethering"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockWiFiTethering(val)
        }
        return nil
    }
    res["certificateCredentialConfigurationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateCredentialConfigurationDisabled(val)
        }
        return nil
    }
    res["crossProfilePoliciesAllowCopyPaste"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossProfilePoliciesAllowCopyPaste(val)
        }
        return nil
    }
    res["crossProfilePoliciesAllowDataSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerCrossProfileDataSharing)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossProfilePoliciesAllowDataSharing(val.(*AndroidDeviceOwnerCrossProfileDataSharing))
        }
        return nil
    }
    res["crossProfilePoliciesShowWorkContactsInPersonalProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossProfilePoliciesShowWorkContactsInPersonalProfile(val)
        }
        return nil
    }
    res["dataRoamingBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataRoamingBlocked(val)
        }
        return nil
    }
    res["dateTimeConfigurationBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTimeConfigurationBlocked(val)
        }
        return nil
    }
    res["detailedHelpText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerUserFacingMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetailedHelpText(val.(AndroidDeviceOwnerUserFacingMessageable))
        }
        return nil
    }
    res["deviceOwnerLockScreenMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerUserFacingMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOwnerLockScreenMessage(val.(AndroidDeviceOwnerUserFacingMessageable))
        }
        return nil
    }
    res["enrollmentProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerEnrollmentProfileType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentProfile(val.(*AndroidDeviceOwnerEnrollmentProfileType))
        }
        return nil
    }
    res["factoryResetBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFactoryResetBlocked(val)
        }
        return nil
    }
    res["factoryResetDeviceAdministratorEmails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetFactoryResetDeviceAdministratorEmails(res)
        }
        return nil
    }
    res["globalProxy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerGlobalProxyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGlobalProxy(val.(AndroidDeviceOwnerGlobalProxyable))
        }
        return nil
    }
    res["googleAccountsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGoogleAccountsBlocked(val)
        }
        return nil
    }
    res["kioskCustomizationDeviceSettingsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskCustomizationDeviceSettingsBlocked(val)
        }
        return nil
    }
    res["kioskCustomizationPowerButtonActionsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskCustomizationPowerButtonActionsBlocked(val)
        }
        return nil
    }
    res["kioskCustomizationStatusBar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerKioskCustomizationStatusBar)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskCustomizationStatusBar(val.(*AndroidDeviceOwnerKioskCustomizationStatusBar))
        }
        return nil
    }
    res["kioskCustomizationSystemErrorWarnings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskCustomizationSystemErrorWarnings(val)
        }
        return nil
    }
    res["kioskCustomizationSystemNavigation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerKioskCustomizationSystemNavigation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskCustomizationSystemNavigation(val.(*AndroidDeviceOwnerKioskCustomizationSystemNavigation))
        }
        return nil
    }
    res["kioskModeAppOrderEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAppOrderEnabled(val)
        }
        return nil
    }
    res["kioskModeAppPositions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerKioskModeAppPositionItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerKioskModeAppPositionItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidDeviceOwnerKioskModeAppPositionItemable)
                }
            }
            m.SetKioskModeAppPositions(res)
        }
        return nil
    }
    res["kioskModeApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppListItemable)
                }
            }
            m.SetKioskModeApps(res)
        }
        return nil
    }
    res["kioskModeAppsInFolderOrderedByName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAppsInFolderOrderedByName(val)
        }
        return nil
    }
    res["kioskModeBluetoothConfigurationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeBluetoothConfigurationEnabled(val)
        }
        return nil
    }
    res["kioskModeDebugMenuEasyAccessEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeDebugMenuEasyAccessEnabled(val)
        }
        return nil
    }
    res["kioskModeExitCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeExitCode(val)
        }
        return nil
    }
    res["kioskModeFlashlightConfigurationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeFlashlightConfigurationEnabled(val)
        }
        return nil
    }
    res["kioskModeFolderIcon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerKioskModeFolderIcon)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeFolderIcon(val.(*AndroidDeviceOwnerKioskModeFolderIcon))
        }
        return nil
    }
    res["kioskModeGridHeight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeGridHeight(val)
        }
        return nil
    }
    res["kioskModeGridWidth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeGridWidth(val)
        }
        return nil
    }
    res["kioskModeIconSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerKioskModeIconSize)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeIconSize(val.(*AndroidDeviceOwnerKioskModeIconSize))
        }
        return nil
    }
    res["kioskModeLockHomeScreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeLockHomeScreen(val)
        }
        return nil
    }
    res["kioskModeManagedFolders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerKioskModeManagedFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerKioskModeManagedFolderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidDeviceOwnerKioskModeManagedFolderable)
                }
            }
            m.SetKioskModeManagedFolders(res)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenAutoSignout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenAutoSignout(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenPinComplexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKioskModeManagedHomeScreenPinComplexity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenPinComplexity(val.(*KioskModeManagedHomeScreenPinComplexity))
        }
        return nil
    }
    res["kioskModeManagedHomeScreenPinRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenPinRequired(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenPinRequiredToResume"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenPinRequiredToResume(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenSignInBackground"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenSignInBackground(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenSignInBrandingLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenSignInBrandingLogo(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenSignInEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenSignInEnabled(val)
        }
        return nil
    }
    res["kioskModeManagedSettingsEntryDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedSettingsEntryDisabled(val)
        }
        return nil
    }
    res["kioskModeMediaVolumeConfigurationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeMediaVolumeConfigurationEnabled(val)
        }
        return nil
    }
    res["kioskModeScreenOrientation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerKioskModeScreenOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeScreenOrientation(val.(*AndroidDeviceOwnerKioskModeScreenOrientation))
        }
        return nil
    }
    res["kioskModeScreenSaverConfigurationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeScreenSaverConfigurationEnabled(val)
        }
        return nil
    }
    res["kioskModeScreenSaverDetectMediaDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeScreenSaverDetectMediaDisabled(val)
        }
        return nil
    }
    res["kioskModeScreenSaverDisplayTimeInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeScreenSaverDisplayTimeInSeconds(val)
        }
        return nil
    }
    res["kioskModeScreenSaverImageUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeScreenSaverImageUrl(val)
        }
        return nil
    }
    res["kioskModeScreenSaverStartDelayInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeScreenSaverStartDelayInSeconds(val)
        }
        return nil
    }
    res["kioskModeShowAppNotificationBadge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeShowAppNotificationBadge(val)
        }
        return nil
    }
    res["kioskModeShowDeviceInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeShowDeviceInfo(val)
        }
        return nil
    }
    res["kioskModeUseManagedHomeScreenApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKioskModeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeUseManagedHomeScreenApp(val.(*KioskModeType))
        }
        return nil
    }
    res["kioskModeVirtualHomeButtonEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeVirtualHomeButtonEnabled(val)
        }
        return nil
    }
    res["kioskModeVirtualHomeButtonType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerVirtualHomeButtonType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeVirtualHomeButtonType(val.(*AndroidDeviceOwnerVirtualHomeButtonType))
        }
        return nil
    }
    res["kioskModeWallpaperUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeWallpaperUrl(val)
        }
        return nil
    }
    res["kioskModeWifiAllowedSsids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetKioskModeWifiAllowedSsids(res)
        }
        return nil
    }
    res["kioskModeWiFiConfigurationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeWiFiConfigurationEnabled(val)
        }
        return nil
    }
    res["locateDeviceLostModeEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocateDeviceLostModeEnabled(val)
        }
        return nil
    }
    res["locateDeviceUserlessDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocateDeviceUserlessDisabled(val)
        }
        return nil
    }
    res["microphoneForceMute"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophoneForceMute(val)
        }
        return nil
    }
    res["microsoftLauncherConfigurationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherConfigurationEnabled(val)
        }
        return nil
    }
    res["microsoftLauncherCustomWallpaperAllowUserModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherCustomWallpaperAllowUserModification(val)
        }
        return nil
    }
    res["microsoftLauncherCustomWallpaperEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherCustomWallpaperEnabled(val)
        }
        return nil
    }
    res["microsoftLauncherCustomWallpaperImageUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherCustomWallpaperImageUrl(val)
        }
        return nil
    }
    res["microsoftLauncherDockPresenceAllowUserModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherDockPresenceAllowUserModification(val)
        }
        return nil
    }
    res["microsoftLauncherDockPresenceConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftLauncherDockPresence)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherDockPresenceConfiguration(val.(*MicrosoftLauncherDockPresence))
        }
        return nil
    }
    res["microsoftLauncherFeedAllowUserModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherFeedAllowUserModification(val)
        }
        return nil
    }
    res["microsoftLauncherFeedEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherFeedEnabled(val)
        }
        return nil
    }
    res["microsoftLauncherSearchBarPlacementConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftLauncherSearchBarPlacement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherSearchBarPlacementConfiguration(val.(*MicrosoftLauncherSearchBarPlacement))
        }
        return nil
    }
    res["networkEscapeHatchAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkEscapeHatchAllowed(val)
        }
        return nil
    }
    res["nfcBlockOutgoingBeam"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNfcBlockOutgoingBeam(val)
        }
        return nil
    }
    res["passwordBlockKeyguard"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockKeyguard(val)
        }
        return nil
    }
    res["passwordBlockKeyguardFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAndroidKeyguardFeature)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidKeyguardFeature, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*AndroidKeyguardFeature))
                }
            }
            m.SetPasswordBlockKeyguardFeatures(res)
        }
        return nil
    }
    res["passwordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordExpirationDays(val)
        }
        return nil
    }
    res["passwordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLength(val)
        }
        return nil
    }
    res["passwordMinimumLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLetterCharacters(val)
        }
        return nil
    }
    res["passwordMinimumLowerCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLowerCaseCharacters(val)
        }
        return nil
    }
    res["passwordMinimumNonLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumNonLetterCharacters(val)
        }
        return nil
    }
    res["passwordMinimumNumericCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumNumericCharacters(val)
        }
        return nil
    }
    res["passwordMinimumSymbolCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumSymbolCharacters(val)
        }
        return nil
    }
    res["passwordMinimumUpperCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumUpperCaseCharacters(val)
        }
        return nil
    }
    res["passwordMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeScreenTimeout(val)
        }
        return nil
    }
    res["passwordPreviousPasswordCountToBlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPreviousPasswordCountToBlock(val)
        }
        return nil
    }
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*AndroidDeviceOwnerRequiredPasswordType))
        }
        return nil
    }
    res["passwordRequireUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerRequiredPasswordUnlock)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequireUnlock(val.(*AndroidDeviceOwnerRequiredPasswordUnlock))
        }
        return nil
    }
    res["passwordSignInFailureCountBeforeFactoryReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordSignInFailureCountBeforeFactoryReset(val)
        }
        return nil
    }
    res["personalProfileAppsAllowInstallFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalProfileAppsAllowInstallFromUnknownSources(val)
        }
        return nil
    }
    res["personalProfileCameraBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalProfileCameraBlocked(val)
        }
        return nil
    }
    res["personalProfilePersonalApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppListItemable)
                }
            }
            m.SetPersonalProfilePersonalApplications(res)
        }
        return nil
    }
    res["personalProfilePlayStoreMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePersonalProfilePersonalPlayStoreMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalProfilePlayStoreMode(val.(*PersonalProfilePersonalPlayStoreMode))
        }
        return nil
    }
    res["personalProfileScreenCaptureBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalProfileScreenCaptureBlocked(val)
        }
        return nil
    }
    res["playStoreMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerPlayStoreMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlayStoreMode(val.(*AndroidDeviceOwnerPlayStoreMode))
        }
        return nil
    }
    res["screenCaptureBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenCaptureBlocked(val)
        }
        return nil
    }
    res["securityCommonCriteriaModeEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityCommonCriteriaModeEnabled(val)
        }
        return nil
    }
    res["securityDeveloperSettingsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityDeveloperSettingsEnabled(val)
        }
        return nil
    }
    res["securityRequireVerifyApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireVerifyApps(val)
        }
        return nil
    }
    res["shortHelpText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerUserFacingMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortHelpText(val.(AndroidDeviceOwnerUserFacingMessageable))
        }
        return nil
    }
    res["statusBarBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusBarBlocked(val)
        }
        return nil
    }
    res["stayOnModes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAndroidDeviceOwnerBatteryPluggedMode)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerBatteryPluggedMode, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*AndroidDeviceOwnerBatteryPluggedMode))
                }
            }
            m.SetStayOnModes(res)
        }
        return nil
    }
    res["storageAllowUsb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageAllowUsb(val)
        }
        return nil
    }
    res["storageBlockExternalMedia"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageBlockExternalMedia(val)
        }
        return nil
    }
    res["storageBlockUsbFileTransfer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageBlockUsbFileTransfer(val)
        }
        return nil
    }
    res["systemUpdateFreezePeriods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerSystemUpdateFreezePeriodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerSystemUpdateFreezePeriodable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidDeviceOwnerSystemUpdateFreezePeriodable)
                }
            }
            m.SetSystemUpdateFreezePeriods(res)
        }
        return nil
    }
    res["systemUpdateInstallType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerSystemUpdateInstallType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemUpdateInstallType(val.(*AndroidDeviceOwnerSystemUpdateInstallType))
        }
        return nil
    }
    res["systemUpdateWindowEndMinutesAfterMidnight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemUpdateWindowEndMinutesAfterMidnight(val)
        }
        return nil
    }
    res["systemUpdateWindowStartMinutesAfterMidnight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemUpdateWindowStartMinutesAfterMidnight(val)
        }
        return nil
    }
    res["systemWindowsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemWindowsBlocked(val)
        }
        return nil
    }
    res["usersBlockAdd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersBlockAdd(val)
        }
        return nil
    }
    res["usersBlockRemove"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersBlockRemove(val)
        }
        return nil
    }
    res["volumeBlockAdjustment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVolumeBlockAdjustment(val)
        }
        return nil
    }
    res["vpnAlwaysOnLockdownMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnAlwaysOnLockdownMode(val)
        }
        return nil
    }
    res["vpnAlwaysOnPackageIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnAlwaysOnPackageIdentifier(val)
        }
        return nil
    }
    res["wifiBlockEditConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiBlockEditConfigurations(val)
        }
        return nil
    }
    res["wifiBlockEditPolicyDefinedConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiBlockEditPolicyDefinedConfigurations(val)
        }
        return nil
    }
    res["workProfilePasswordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordExpirationDays(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumLength(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumLetterCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumLowerCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumLowerCaseCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumNonLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumNonLetterCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumNumericCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumNumericCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumSymbolCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumSymbolCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumUpperCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumUpperCaseCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordPreviousPasswordCountToBlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordPreviousPasswordCountToBlock(val)
        }
        return nil
    }
    res["workProfilePasswordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordRequiredType(val.(*AndroidDeviceOwnerRequiredPasswordType))
        }
        return nil
    }
    res["workProfilePasswordRequireUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerRequiredPasswordUnlock)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordRequireUnlock(val.(*AndroidDeviceOwnerRequiredPasswordUnlock))
        }
        return nil
    }
    res["workProfilePasswordSignInFailureCountBeforeFactoryReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(val)
        }
        return nil
    }
    return res
}
// GetGlobalProxy gets the globalProxy property value. Proxy is set up directly with host, port and excluded hosts.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetGlobalProxy()(AndroidDeviceOwnerGlobalProxyable) {
    val, err := m.GetBackingStore().Get("globalProxy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidDeviceOwnerGlobalProxyable)
    }
    return nil
}
// GetGoogleAccountsBlocked gets the googleAccountsBlocked property value. Indicates whether or not google accounts will be blocked.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetGoogleAccountsBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("googleAccountsBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskCustomizationDeviceSettingsBlocked gets the kioskCustomizationDeviceSettingsBlocked property value. Indicates whether a user can access the device's Settings app while in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskCustomizationDeviceSettingsBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("kioskCustomizationDeviceSettingsBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskCustomizationPowerButtonActionsBlocked gets the kioskCustomizationPowerButtonActionsBlocked property value. Whether the power menu is shown when a user long presses the Power button of a device in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskCustomizationPowerButtonActionsBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("kioskCustomizationPowerButtonActionsBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskCustomizationStatusBar gets the kioskCustomizationStatusBar property value. Indicates whether system info and notifications are disabled in Kiosk Mode. Possible values are: notConfigured, notificationsAndSystemInfoEnabled, systemInfoOnly.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskCustomizationStatusBar()(*AndroidDeviceOwnerKioskCustomizationStatusBar) {
    val, err := m.GetBackingStore().Get("kioskCustomizationStatusBar")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerKioskCustomizationStatusBar)
    }
    return nil
}
// GetKioskCustomizationSystemErrorWarnings gets the kioskCustomizationSystemErrorWarnings property value. Indicates whether system error dialogs for crashed or unresponsive apps are shown in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskCustomizationSystemErrorWarnings()(*bool) {
    val, err := m.GetBackingStore().Get("kioskCustomizationSystemErrorWarnings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskCustomizationSystemNavigation gets the kioskCustomizationSystemNavigation property value. Indicates which navigation features are enabled in Kiosk Mode. Possible values are: notConfigured, navigationEnabled, homeButtonOnly.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskCustomizationSystemNavigation()(*AndroidDeviceOwnerKioskCustomizationSystemNavigation) {
    val, err := m.GetBackingStore().Get("kioskCustomizationSystemNavigation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerKioskCustomizationSystemNavigation)
    }
    return nil
}
// GetKioskModeAppOrderEnabled gets the kioskModeAppOrderEnabled property value. Whether or not to enable app ordering in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeAppOrderEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeAppOrderEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeAppPositions gets the kioskModeAppPositions property value. The ordering of items on Kiosk Mode Managed Home Screen. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeAppPositions()([]AndroidDeviceOwnerKioskModeAppPositionItemable) {
    val, err := m.GetBackingStore().Get("kioskModeAppPositions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidDeviceOwnerKioskModeAppPositionItemable)
    }
    return nil
}
// GetKioskModeApps gets the kioskModeApps property value. A list of managed apps that will be shown when the device is in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeApps()([]AppListItemable) {
    val, err := m.GetBackingStore().Get("kioskModeApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppListItemable)
    }
    return nil
}
// GetKioskModeAppsInFolderOrderedByName gets the kioskModeAppsInFolderOrderedByName property value. Whether or not to alphabetize applications within a folder in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeAppsInFolderOrderedByName()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeAppsInFolderOrderedByName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeBluetoothConfigurationEnabled gets the kioskModeBluetoothConfigurationEnabled property value. Whether or not to allow a user to configure Bluetooth settings in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeBluetoothConfigurationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeBluetoothConfigurationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeDebugMenuEasyAccessEnabled gets the kioskModeDebugMenuEasyAccessEnabled property value. Whether or not to allow a user to easy access to the debug menu in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeDebugMenuEasyAccessEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeDebugMenuEasyAccessEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeExitCode gets the kioskModeExitCode property value. Exit code to allow a user to escape from Kiosk Mode when the device is in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeExitCode()(*string) {
    val, err := m.GetBackingStore().Get("kioskModeExitCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKioskModeFlashlightConfigurationEnabled gets the kioskModeFlashlightConfigurationEnabled property value. Whether or not to allow a user to use the flashlight in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeFlashlightConfigurationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeFlashlightConfigurationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeFolderIcon gets the kioskModeFolderIcon property value. Folder icon configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, darkSquare, darkCircle, lightSquare, lightCircle.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeFolderIcon()(*AndroidDeviceOwnerKioskModeFolderIcon) {
    val, err := m.GetBackingStore().Get("kioskModeFolderIcon")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerKioskModeFolderIcon)
    }
    return nil
}
// GetKioskModeGridHeight gets the kioskModeGridHeight property value. Number of rows for Managed Home Screen grid with app ordering enabled in Kiosk Mode. Valid values 1 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeGridHeight()(*int32) {
    val, err := m.GetBackingStore().Get("kioskModeGridHeight")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetKioskModeGridWidth gets the kioskModeGridWidth property value. Number of columns for Managed Home Screen grid with app ordering enabled in Kiosk Mode. Valid values 1 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeGridWidth()(*int32) {
    val, err := m.GetBackingStore().Get("kioskModeGridWidth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetKioskModeIconSize gets the kioskModeIconSize property value. Icon size configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, smallest, small, regular, large, largest.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeIconSize()(*AndroidDeviceOwnerKioskModeIconSize) {
    val, err := m.GetBackingStore().Get("kioskModeIconSize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerKioskModeIconSize)
    }
    return nil
}
// GetKioskModeLockHomeScreen gets the kioskModeLockHomeScreen property value. Whether or not to lock home screen to the end user in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeLockHomeScreen()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeLockHomeScreen")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeManagedFolders gets the kioskModeManagedFolders property value. A list of managed folders for a device in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedFolders()([]AndroidDeviceOwnerKioskModeManagedFolderable) {
    val, err := m.GetBackingStore().Get("kioskModeManagedFolders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidDeviceOwnerKioskModeManagedFolderable)
    }
    return nil
}
// GetKioskModeManagedHomeScreenAutoSignout gets the kioskModeManagedHomeScreenAutoSignout property value. Whether or not to automatically sign-out of MHS and Shared device mode applications after inactive for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenAutoSignout()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeManagedHomeScreenAutoSignout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds gets the kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds property value. Number of seconds to give user notice before automatically signing them out for Managed Home Screen. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds gets the kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds property value. Number of seconds device is inactive before automatically signing user out for Managed Home Screen. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetKioskModeManagedHomeScreenPinComplexity gets the kioskModeManagedHomeScreenPinComplexity property value. Complexity of PIN for sign-in session for Managed Home Screen. Possible values are: notConfigured, simple, complex.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenPinComplexity()(*KioskModeManagedHomeScreenPinComplexity) {
    val, err := m.GetBackingStore().Get("kioskModeManagedHomeScreenPinComplexity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*KioskModeManagedHomeScreenPinComplexity)
    }
    return nil
}
// GetKioskModeManagedHomeScreenPinRequired gets the kioskModeManagedHomeScreenPinRequired property value. Whether or not require user to set a PIN for sign-in session for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenPinRequired()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeManagedHomeScreenPinRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeManagedHomeScreenPinRequiredToResume gets the kioskModeManagedHomeScreenPinRequiredToResume property value. Whether or not required user to enter session PIN if screensaver has appeared for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenPinRequiredToResume()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeManagedHomeScreenPinRequiredToResume")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeManagedHomeScreenSignInBackground gets the kioskModeManagedHomeScreenSignInBackground property value. Custom URL background for sign-in screen for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenSignInBackground()(*string) {
    val, err := m.GetBackingStore().Get("kioskModeManagedHomeScreenSignInBackground")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKioskModeManagedHomeScreenSignInBrandingLogo gets the kioskModeManagedHomeScreenSignInBrandingLogo property value. Custom URL branding logo for sign-in screen and session pin page for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenSignInBrandingLogo()(*string) {
    val, err := m.GetBackingStore().Get("kioskModeManagedHomeScreenSignInBrandingLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKioskModeManagedHomeScreenSignInEnabled gets the kioskModeManagedHomeScreenSignInEnabled property value. Whether or not show sign-in screen for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenSignInEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeManagedHomeScreenSignInEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeManagedSettingsEntryDisabled gets the kioskModeManagedSettingsEntryDisabled property value. Whether or not to display the Managed Settings entry point on the managed home screen in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedSettingsEntryDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeManagedSettingsEntryDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeMediaVolumeConfigurationEnabled gets the kioskModeMediaVolumeConfigurationEnabled property value. Whether or not to allow a user to change the media volume in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeMediaVolumeConfigurationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeMediaVolumeConfigurationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeScreenOrientation gets the kioskModeScreenOrientation property value. Screen orientation configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, portrait, landscape, autoRotate.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeScreenOrientation()(*AndroidDeviceOwnerKioskModeScreenOrientation) {
    val, err := m.GetBackingStore().Get("kioskModeScreenOrientation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerKioskModeScreenOrientation)
    }
    return nil
}
// GetKioskModeScreenSaverConfigurationEnabled gets the kioskModeScreenSaverConfigurationEnabled property value. Whether or not to enable screen saver mode or not in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeScreenSaverConfigurationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeScreenSaverConfigurationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeScreenSaverDetectMediaDisabled gets the kioskModeScreenSaverDetectMediaDisabled property value. Whether or not the device screen should show the screen saver if audio/video is playing in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeScreenSaverDetectMediaDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeScreenSaverDetectMediaDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeScreenSaverDisplayTimeInSeconds gets the kioskModeScreenSaverDisplayTimeInSeconds property value. The number of seconds that the device will display the screen saver for in Kiosk Mode. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeScreenSaverDisplayTimeInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("kioskModeScreenSaverDisplayTimeInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetKioskModeScreenSaverImageUrl gets the kioskModeScreenSaverImageUrl property value. URL for an image that will be the device's screen saver in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeScreenSaverImageUrl()(*string) {
    val, err := m.GetBackingStore().Get("kioskModeScreenSaverImageUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKioskModeScreenSaverStartDelayInSeconds gets the kioskModeScreenSaverStartDelayInSeconds property value. The number of seconds the device needs to be inactive for before the screen saver is shown in Kiosk Mode. Valid values 1 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeScreenSaverStartDelayInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("kioskModeScreenSaverStartDelayInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetKioskModeShowAppNotificationBadge gets the kioskModeShowAppNotificationBadge property value. Whether or not to display application notification badges in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeShowAppNotificationBadge()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeShowAppNotificationBadge")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeShowDeviceInfo gets the kioskModeShowDeviceInfo property value. Whether or not to allow a user to access basic device information.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeShowDeviceInfo()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeShowDeviceInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeUseManagedHomeScreenApp gets the kioskModeUseManagedHomeScreenApp property value. Whether or not to use single app kiosk mode or multi-app kiosk mode. Possible values are: notConfigured, singleAppMode, multiAppMode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeUseManagedHomeScreenApp()(*KioskModeType) {
    val, err := m.GetBackingStore().Get("kioskModeUseManagedHomeScreenApp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*KioskModeType)
    }
    return nil
}
// GetKioskModeVirtualHomeButtonEnabled gets the kioskModeVirtualHomeButtonEnabled property value. Whether or not to display a virtual home button when the device is in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeVirtualHomeButtonEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeVirtualHomeButtonEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskModeVirtualHomeButtonType gets the kioskModeVirtualHomeButtonType property value. Indicates whether the virtual home button is a swipe up home button or a floating home button. Possible values are: notConfigured, swipeUp, floating.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeVirtualHomeButtonType()(*AndroidDeviceOwnerVirtualHomeButtonType) {
    val, err := m.GetBackingStore().Get("kioskModeVirtualHomeButtonType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerVirtualHomeButtonType)
    }
    return nil
}
// GetKioskModeWallpaperUrl gets the kioskModeWallpaperUrl property value. URL to a publicly accessible image to use for the wallpaper when the device is in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeWallpaperUrl()(*string) {
    val, err := m.GetBackingStore().Get("kioskModeWallpaperUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKioskModeWifiAllowedSsids gets the kioskModeWifiAllowedSsids property value. The restricted set of WIFI SSIDs available for the user to configure in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeWifiAllowedSsids()([]string) {
    val, err := m.GetBackingStore().Get("kioskModeWifiAllowedSsids")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetKioskModeWiFiConfigurationEnabled gets the kioskModeWiFiConfigurationEnabled property value. Whether or not to allow a user to configure Wi-Fi settings in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeWiFiConfigurationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("kioskModeWiFiConfigurationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocateDeviceLostModeEnabled gets the locateDeviceLostModeEnabled property value. Indicates whether or not LocateDevice for devices with lost mode (COBO, COPE) is enabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetLocateDeviceLostModeEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("locateDeviceLostModeEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocateDeviceUserlessDisabled gets the locateDeviceUserlessDisabled property value. Indicates whether or not LocateDevice for userless (COSU) devices is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetLocateDeviceUserlessDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("locateDeviceUserlessDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMicrophoneForceMute gets the microphoneForceMute property value. Indicates whether or not to block unmuting the microphone on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrophoneForceMute()(*bool) {
    val, err := m.GetBackingStore().Get("microphoneForceMute")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMicrosoftLauncherConfigurationEnabled gets the microsoftLauncherConfigurationEnabled property value. Indicates whether or not to you want configure Microsoft Launcher.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherConfigurationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("microsoftLauncherConfigurationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMicrosoftLauncherCustomWallpaperAllowUserModification gets the microsoftLauncherCustomWallpaperAllowUserModification property value. Indicates whether or not the user can modify the wallpaper to personalize their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherCustomWallpaperAllowUserModification()(*bool) {
    val, err := m.GetBackingStore().Get("microsoftLauncherCustomWallpaperAllowUserModification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMicrosoftLauncherCustomWallpaperEnabled gets the microsoftLauncherCustomWallpaperEnabled property value. Indicates whether or not to configure the wallpaper on the targeted devices.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherCustomWallpaperEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("microsoftLauncherCustomWallpaperEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMicrosoftLauncherCustomWallpaperImageUrl gets the microsoftLauncherCustomWallpaperImageUrl property value. Indicates the URL for the image file to use as the wallpaper on the targeted devices.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherCustomWallpaperImageUrl()(*string) {
    val, err := m.GetBackingStore().Get("microsoftLauncherCustomWallpaperImageUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMicrosoftLauncherDockPresenceAllowUserModification gets the microsoftLauncherDockPresenceAllowUserModification property value. Indicates whether or not the user can modify the device dock configuration on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherDockPresenceAllowUserModification()(*bool) {
    val, err := m.GetBackingStore().Get("microsoftLauncherDockPresenceAllowUserModification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMicrosoftLauncherDockPresenceConfiguration gets the microsoftLauncherDockPresenceConfiguration property value. Indicates whether or not you want to configure the device dock. Possible values are: notConfigured, show, hide, disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherDockPresenceConfiguration()(*MicrosoftLauncherDockPresence) {
    val, err := m.GetBackingStore().Get("microsoftLauncherDockPresenceConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MicrosoftLauncherDockPresence)
    }
    return nil
}
// GetMicrosoftLauncherFeedAllowUserModification gets the microsoftLauncherFeedAllowUserModification property value. Indicates whether or not the user can modify the launcher feed on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherFeedAllowUserModification()(*bool) {
    val, err := m.GetBackingStore().Get("microsoftLauncherFeedAllowUserModification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMicrosoftLauncherFeedEnabled gets the microsoftLauncherFeedEnabled property value. Indicates whether or not you want to enable the launcher feed on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherFeedEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("microsoftLauncherFeedEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMicrosoftLauncherSearchBarPlacementConfiguration gets the microsoftLauncherSearchBarPlacementConfiguration property value. Indicates the search bar placement configuration on the device. Possible values are: notConfigured, top, bottom, hide.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherSearchBarPlacementConfiguration()(*MicrosoftLauncherSearchBarPlacement) {
    val, err := m.GetBackingStore().Get("microsoftLauncherSearchBarPlacementConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MicrosoftLauncherSearchBarPlacement)
    }
    return nil
}
// GetNetworkEscapeHatchAllowed gets the networkEscapeHatchAllowed property value. Indicates whether or not the device will allow connecting to a temporary network connection at boot time.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetNetworkEscapeHatchAllowed()(*bool) {
    val, err := m.GetBackingStore().Get("networkEscapeHatchAllowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetNfcBlockOutgoingBeam gets the nfcBlockOutgoingBeam property value. Indicates whether or not to block NFC outgoing beam.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetNfcBlockOutgoingBeam()(*bool) {
    val, err := m.GetBackingStore().Get("nfcBlockOutgoingBeam")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordBlockKeyguard gets the passwordBlockKeyguard property value. Indicates whether or not the keyguard is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordBlockKeyguard()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockKeyguard")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordBlockKeyguardFeatures gets the passwordBlockKeyguardFeatures property value. List of device keyguard features to block. This collection can contain a maximum of 11 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordBlockKeyguardFeatures()([]AndroidKeyguardFeature) {
    val, err := m.GetBackingStore().Get("passwordBlockKeyguardFeatures")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidKeyguardFeature)
    }
    return nil
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Indicates the amount of time that a password can be set for before it expires and a new password will be required. Valid values 1 to 365
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordExpirationDays()(*int32) {
    val, err := m.GetBackingStore().Get("passwordExpirationDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Indicates the minimum length of the password required on the device. Valid values 4 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumLength()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumLetterCharacters gets the passwordMinimumLetterCharacters property value. Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumLetterCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumLetterCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumLowerCaseCharacters gets the passwordMinimumLowerCaseCharacters property value. Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumLowerCaseCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumLowerCaseCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumNonLetterCharacters gets the passwordMinimumNonLetterCharacters property value. Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumNonLetterCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumNonLetterCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumNumericCharacters gets the passwordMinimumNumericCharacters property value. Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumNumericCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumNumericCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumSymbolCharacters gets the passwordMinimumSymbolCharacters property value. Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumSymbolCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumSymbolCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumUpperCaseCharacters gets the passwordMinimumUpperCaseCharacters property value. Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumUpperCaseCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumUpperCaseCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinutesOfInactivityBeforeScreenTimeout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordPreviousPasswordCountToBlock gets the passwordPreviousPasswordCountToBlock property value. Indicates the length of password history, where the user will not be able to enter a new password that is the same as any password in the history. Valid values 0 to 24
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordPreviousPasswordCountToBlock()(*int32) {
    val, err := m.GetBackingStore().Get("passwordPreviousPasswordCountToBlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Indicates the minimum password quality required on the device. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType) {
    val, err := m.GetBackingStore().Get("passwordRequiredType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerRequiredPasswordType)
    }
    return nil
}
// GetPasswordRequireUnlock gets the passwordRequireUnlock property value. Indicates the timeout period after which a device must be unlocked using a form of strong authentication. Possible values are: deviceDefault, daily, unkownFutureValue.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordRequireUnlock()(*AndroidDeviceOwnerRequiredPasswordUnlock) {
    val, err := m.GetBackingStore().Get("passwordRequireUnlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerRequiredPasswordUnlock)
    }
    return nil
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    val, err := m.GetBackingStore().Get("passwordSignInFailureCountBeforeFactoryReset")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPersonalProfileAppsAllowInstallFromUnknownSources gets the personalProfileAppsAllowInstallFromUnknownSources property value. Indicates whether the user can install apps from unknown sources on the personal profile.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPersonalProfileAppsAllowInstallFromUnknownSources()(*bool) {
    val, err := m.GetBackingStore().Get("personalProfileAppsAllowInstallFromUnknownSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPersonalProfileCameraBlocked gets the personalProfileCameraBlocked property value. Indicates whether to disable the use of the camera on the personal profile.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPersonalProfileCameraBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("personalProfileCameraBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPersonalProfilePersonalApplications gets the personalProfilePersonalApplications property value. Policy applied to applications in the personal profile. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPersonalProfilePersonalApplications()([]AppListItemable) {
    val, err := m.GetBackingStore().Get("personalProfilePersonalApplications")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppListItemable)
    }
    return nil
}
// GetPersonalProfilePlayStoreMode gets the personalProfilePlayStoreMode property value. Used together with PersonalProfilePersonalApplications to control how apps in the personal profile are allowed or blocked. Possible values are: notConfigured, blockedApps, allowedApps.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPersonalProfilePlayStoreMode()(*PersonalProfilePersonalPlayStoreMode) {
    val, err := m.GetBackingStore().Get("personalProfilePlayStoreMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PersonalProfilePersonalPlayStoreMode)
    }
    return nil
}
// GetPersonalProfileScreenCaptureBlocked gets the personalProfileScreenCaptureBlocked property value. Indicates whether to disable the capability to take screenshots on the personal profile.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPersonalProfileScreenCaptureBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("personalProfileScreenCaptureBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPlayStoreMode gets the playStoreMode property value. Indicates the Play Store mode of the device. Possible values are: notConfigured, allowList, blockList.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPlayStoreMode()(*AndroidDeviceOwnerPlayStoreMode) {
    val, err := m.GetBackingStore().Get("playStoreMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerPlayStoreMode)
    }
    return nil
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether or not to disable the capability to take screenshots.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetScreenCaptureBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("screenCaptureBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityCommonCriteriaModeEnabled gets the securityCommonCriteriaModeEnabled property value. Represents the security common criteria mode enabled provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSecurityCommonCriteriaModeEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("securityCommonCriteriaModeEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityDeveloperSettingsEnabled gets the securityDeveloperSettingsEnabled property value. Indicates whether or not the user is allowed to access developer settings like developer options and safe boot on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSecurityDeveloperSettingsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("securityDeveloperSettingsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecurityRequireVerifyApps gets the securityRequireVerifyApps property value. Indicates whether or not verify apps is required.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSecurityRequireVerifyApps()(*bool) {
    val, err := m.GetBackingStore().Get("securityRequireVerifyApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShortHelpText gets the shortHelpText property value. Represents the customized short help text provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetShortHelpText()(AndroidDeviceOwnerUserFacingMessageable) {
    val, err := m.GetBackingStore().Get("shortHelpText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidDeviceOwnerUserFacingMessageable)
    }
    return nil
}
// GetStatusBarBlocked gets the statusBarBlocked property value. Indicates whether or the status bar is disabled, including notifications, quick settings and other screen overlays.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetStatusBarBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("statusBarBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStayOnModes gets the stayOnModes property value. List of modes in which the device's display will stay powered-on. This collection can contain a maximum of 4 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetStayOnModes()([]AndroidDeviceOwnerBatteryPluggedMode) {
    val, err := m.GetBackingStore().Get("stayOnModes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidDeviceOwnerBatteryPluggedMode)
    }
    return nil
}
// GetStorageAllowUsb gets the storageAllowUsb property value. Indicates whether or not to allow USB mass storage.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetStorageAllowUsb()(*bool) {
    val, err := m.GetBackingStore().Get("storageAllowUsb")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStorageBlockExternalMedia gets the storageBlockExternalMedia property value. Indicates whether or not to block external media.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetStorageBlockExternalMedia()(*bool) {
    val, err := m.GetBackingStore().Get("storageBlockExternalMedia")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetStorageBlockUsbFileTransfer gets the storageBlockUsbFileTransfer property value. Indicates whether or not to block USB file transfer.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetStorageBlockUsbFileTransfer()(*bool) {
    val, err := m.GetBackingStore().Get("storageBlockUsbFileTransfer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSystemUpdateFreezePeriods gets the systemUpdateFreezePeriods property value. Indicates the annually repeating time periods during which system updates are postponed. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSystemUpdateFreezePeriods()([]AndroidDeviceOwnerSystemUpdateFreezePeriodable) {
    val, err := m.GetBackingStore().Get("systemUpdateFreezePeriods")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidDeviceOwnerSystemUpdateFreezePeriodable)
    }
    return nil
}
// GetSystemUpdateInstallType gets the systemUpdateInstallType property value. The type of system update configuration. Possible values are: deviceDefault, postpone, windowed, automatic.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSystemUpdateInstallType()(*AndroidDeviceOwnerSystemUpdateInstallType) {
    val, err := m.GetBackingStore().Get("systemUpdateInstallType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerSystemUpdateInstallType)
    }
    return nil
}
// GetSystemUpdateWindowEndMinutesAfterMidnight gets the systemUpdateWindowEndMinutesAfterMidnight property value. Indicates the number of minutes after midnight that the system update window ends. Valid values 0 to 1440
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSystemUpdateWindowEndMinutesAfterMidnight()(*int32) {
    val, err := m.GetBackingStore().Get("systemUpdateWindowEndMinutesAfterMidnight")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSystemUpdateWindowStartMinutesAfterMidnight gets the systemUpdateWindowStartMinutesAfterMidnight property value. Indicates the number of minutes after midnight that the system update window starts. Valid values 0 to 1440
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSystemUpdateWindowStartMinutesAfterMidnight()(*int32) {
    val, err := m.GetBackingStore().Get("systemUpdateWindowStartMinutesAfterMidnight")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSystemWindowsBlocked gets the systemWindowsBlocked property value. Whether or not to block Android system prompt windows, like toasts, phone activities, and system alerts.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSystemWindowsBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("systemWindowsBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUsersBlockAdd gets the usersBlockAdd property value. Indicates whether or not adding users and profiles is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetUsersBlockAdd()(*bool) {
    val, err := m.GetBackingStore().Get("usersBlockAdd")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUsersBlockRemove gets the usersBlockRemove property value. Indicates whether or not to disable removing other users from the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetUsersBlockRemove()(*bool) {
    val, err := m.GetBackingStore().Get("usersBlockRemove")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetVolumeBlockAdjustment gets the volumeBlockAdjustment property value. Indicates whether or not adjusting the master volume is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetVolumeBlockAdjustment()(*bool) {
    val, err := m.GetBackingStore().Get("volumeBlockAdjustment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetVpnAlwaysOnLockdownMode gets the vpnAlwaysOnLockdownMode property value. If an always on VPN package name is specified, whether or not to lock network traffic when that VPN is disconnected.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetVpnAlwaysOnLockdownMode()(*bool) {
    val, err := m.GetBackingStore().Get("vpnAlwaysOnLockdownMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetVpnAlwaysOnPackageIdentifier gets the vpnAlwaysOnPackageIdentifier property value. Android app package name for app that will handle an always-on VPN connection.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetVpnAlwaysOnPackageIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("vpnAlwaysOnPackageIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWifiBlockEditConfigurations gets the wifiBlockEditConfigurations property value. Indicates whether or not to block the user from editing the wifi connection settings.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWifiBlockEditConfigurations()(*bool) {
    val, err := m.GetBackingStore().Get("wifiBlockEditConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWifiBlockEditPolicyDefinedConfigurations gets the wifiBlockEditPolicyDefinedConfigurations property value. Indicates whether or not to block the user from editing just the networks defined by the policy.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWifiBlockEditPolicyDefinedConfigurations()(*bool) {
    val, err := m.GetBackingStore().Get("wifiBlockEditPolicyDefinedConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWorkProfilePasswordExpirationDays gets the workProfilePasswordExpirationDays property value. Indicates the number of days that a work profile password can be set before it expires and a new password will be required. Valid values 1 to 365
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordExpirationDays()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordExpirationDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinimumLength gets the workProfilePasswordMinimumLength property value. Indicates the minimum length of the work profile password. Valid values 4 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLength()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinimumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinimumLetterCharacters gets the workProfilePasswordMinimumLetterCharacters property value. Indicates the minimum number of letter characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLetterCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinimumLetterCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinimumLowerCaseCharacters gets the workProfilePasswordMinimumLowerCaseCharacters property value. Indicates the minimum number of lower-case characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLowerCaseCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinimumLowerCaseCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinimumNonLetterCharacters gets the workProfilePasswordMinimumNonLetterCharacters property value. Indicates the minimum number of non-letter characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumNonLetterCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinimumNonLetterCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinimumNumericCharacters gets the workProfilePasswordMinimumNumericCharacters property value. Indicates the minimum number of numeric characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumNumericCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinimumNumericCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinimumSymbolCharacters gets the workProfilePasswordMinimumSymbolCharacters property value. Indicates the minimum number of symbol characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumSymbolCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinimumSymbolCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordMinimumUpperCaseCharacters gets the workProfilePasswordMinimumUpperCaseCharacters property value. Indicates the minimum number of upper-case letter characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumUpperCaseCharacters()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordMinimumUpperCaseCharacters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordPreviousPasswordCountToBlock gets the workProfilePasswordPreviousPasswordCountToBlock property value. Indicates the length of the work profile password history, where the user will not be able to enter a new password that is the same as any password in the history. Valid values 0 to 24
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordPreviousPasswordCountToBlock()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordPreviousPasswordCountToBlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWorkProfilePasswordRequiredType gets the workProfilePasswordRequiredType property value. Indicates the minimum password quality required on the work profile password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType) {
    val, err := m.GetBackingStore().Get("workProfilePasswordRequiredType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerRequiredPasswordType)
    }
    return nil
}
// GetWorkProfilePasswordRequireUnlock gets the workProfilePasswordRequireUnlock property value. Indicates the timeout period after which a work profile must be unlocked using a form of strong authentication. Possible values are: deviceDefault, daily, unkownFutureValue.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordRequireUnlock()(*AndroidDeviceOwnerRequiredPasswordUnlock) {
    val, err := m.GetBackingStore().Get("workProfilePasswordRequireUnlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerRequiredPasswordUnlock)
    }
    return nil
}
// GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset gets the workProfilePasswordSignInFailureCountBeforeFactoryReset property value. Indicates the number of times a user can enter an incorrect work profile password before the device is wiped. Valid values 4 to 11
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    val, err := m.GetBackingStore().Get("workProfilePasswordSignInFailureCountBeforeFactoryReset")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountsBlockModification", m.GetAccountsBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appsAllowInstallFromUnknownSources", m.GetAppsAllowInstallFromUnknownSources())
        if err != nil {
            return err
        }
    }
    if m.GetAppsAutoUpdatePolicy() != nil {
        cast := (*m.GetAppsAutoUpdatePolicy()).String()
        err = writer.WriteStringValue("appsAutoUpdatePolicy", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppsDefaultPermissionPolicy() != nil {
        cast := (*m.GetAppsDefaultPermissionPolicy()).String()
        err = writer.WriteStringValue("appsDefaultPermissionPolicy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appsRecommendSkippingFirstUseHints", m.GetAppsRecommendSkippingFirstUseHints())
        if err != nil {
            return err
        }
    }
    if m.GetAzureAdSharedDeviceDataClearApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAzureAdSharedDeviceDataClearApps()))
        for i, v := range m.GetAzureAdSharedDeviceDataClearApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("azureAdSharedDeviceDataClearApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bluetoothBlockConfiguration", m.GetBluetoothBlockConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bluetoothBlockContactSharing", m.GetBluetoothBlockContactSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cameraBlocked", m.GetCameraBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockWiFiTethering", m.GetCellularBlockWiFiTethering())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("certificateCredentialConfigurationDisabled", m.GetCertificateCredentialConfigurationDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("crossProfilePoliciesAllowCopyPaste", m.GetCrossProfilePoliciesAllowCopyPaste())
        if err != nil {
            return err
        }
    }
    if m.GetCrossProfilePoliciesAllowDataSharing() != nil {
        cast := (*m.GetCrossProfilePoliciesAllowDataSharing()).String()
        err = writer.WriteStringValue("crossProfilePoliciesAllowDataSharing", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("crossProfilePoliciesShowWorkContactsInPersonalProfile", m.GetCrossProfilePoliciesShowWorkContactsInPersonalProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dataRoamingBlocked", m.GetDataRoamingBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dateTimeConfigurationBlocked", m.GetDateTimeConfigurationBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("detailedHelpText", m.GetDetailedHelpText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceOwnerLockScreenMessage", m.GetDeviceOwnerLockScreenMessage())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentProfile() != nil {
        cast := (*m.GetEnrollmentProfile()).String()
        err = writer.WriteStringValue("enrollmentProfile", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("factoryResetBlocked", m.GetFactoryResetBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetFactoryResetDeviceAdministratorEmails() != nil {
        err = writer.WriteCollectionOfStringValues("factoryResetDeviceAdministratorEmails", m.GetFactoryResetDeviceAdministratorEmails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("globalProxy", m.GetGlobalProxy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("googleAccountsBlocked", m.GetGoogleAccountsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskCustomizationDeviceSettingsBlocked", m.GetKioskCustomizationDeviceSettingsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskCustomizationPowerButtonActionsBlocked", m.GetKioskCustomizationPowerButtonActionsBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetKioskCustomizationStatusBar() != nil {
        cast := (*m.GetKioskCustomizationStatusBar()).String()
        err = writer.WriteStringValue("kioskCustomizationStatusBar", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskCustomizationSystemErrorWarnings", m.GetKioskCustomizationSystemErrorWarnings())
        if err != nil {
            return err
        }
    }
    if m.GetKioskCustomizationSystemNavigation() != nil {
        cast := (*m.GetKioskCustomizationSystemNavigation()).String()
        err = writer.WriteStringValue("kioskCustomizationSystemNavigation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAppOrderEnabled", m.GetKioskModeAppOrderEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeAppPositions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKioskModeAppPositions()))
        for i, v := range m.GetKioskModeAppPositions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("kioskModeAppPositions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKioskModeApps()))
        for i, v := range m.GetKioskModeApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("kioskModeApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAppsInFolderOrderedByName", m.GetKioskModeAppsInFolderOrderedByName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeBluetoothConfigurationEnabled", m.GetKioskModeBluetoothConfigurationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeDebugMenuEasyAccessEnabled", m.GetKioskModeDebugMenuEasyAccessEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeExitCode", m.GetKioskModeExitCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeFlashlightConfigurationEnabled", m.GetKioskModeFlashlightConfigurationEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeFolderIcon() != nil {
        cast := (*m.GetKioskModeFolderIcon()).String()
        err = writer.WriteStringValue("kioskModeFolderIcon", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskModeGridHeight", m.GetKioskModeGridHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskModeGridWidth", m.GetKioskModeGridWidth())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeIconSize() != nil {
        cast := (*m.GetKioskModeIconSize()).String()
        err = writer.WriteStringValue("kioskModeIconSize", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeLockHomeScreen", m.GetKioskModeLockHomeScreen())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeManagedFolders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKioskModeManagedFolders()))
        for i, v := range m.GetKioskModeManagedFolders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("kioskModeManagedFolders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeManagedHomeScreenAutoSignout", m.GetKioskModeManagedHomeScreenAutoSignout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds", m.GetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds", m.GetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeManagedHomeScreenPinComplexity() != nil {
        cast := (*m.GetKioskModeManagedHomeScreenPinComplexity()).String()
        err = writer.WriteStringValue("kioskModeManagedHomeScreenPinComplexity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeManagedHomeScreenPinRequired", m.GetKioskModeManagedHomeScreenPinRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeManagedHomeScreenPinRequiredToResume", m.GetKioskModeManagedHomeScreenPinRequiredToResume())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeManagedHomeScreenSignInBackground", m.GetKioskModeManagedHomeScreenSignInBackground())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeManagedHomeScreenSignInBrandingLogo", m.GetKioskModeManagedHomeScreenSignInBrandingLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeManagedHomeScreenSignInEnabled", m.GetKioskModeManagedHomeScreenSignInEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeManagedSettingsEntryDisabled", m.GetKioskModeManagedSettingsEntryDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeMediaVolumeConfigurationEnabled", m.GetKioskModeMediaVolumeConfigurationEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeScreenOrientation() != nil {
        cast := (*m.GetKioskModeScreenOrientation()).String()
        err = writer.WriteStringValue("kioskModeScreenOrientation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeScreenSaverConfigurationEnabled", m.GetKioskModeScreenSaverConfigurationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeScreenSaverDetectMediaDisabled", m.GetKioskModeScreenSaverDetectMediaDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskModeScreenSaverDisplayTimeInSeconds", m.GetKioskModeScreenSaverDisplayTimeInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeScreenSaverImageUrl", m.GetKioskModeScreenSaverImageUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskModeScreenSaverStartDelayInSeconds", m.GetKioskModeScreenSaverStartDelayInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeShowAppNotificationBadge", m.GetKioskModeShowAppNotificationBadge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeShowDeviceInfo", m.GetKioskModeShowDeviceInfo())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeUseManagedHomeScreenApp() != nil {
        cast := (*m.GetKioskModeUseManagedHomeScreenApp()).String()
        err = writer.WriteStringValue("kioskModeUseManagedHomeScreenApp", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeVirtualHomeButtonEnabled", m.GetKioskModeVirtualHomeButtonEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeVirtualHomeButtonType() != nil {
        cast := (*m.GetKioskModeVirtualHomeButtonType()).String()
        err = writer.WriteStringValue("kioskModeVirtualHomeButtonType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeWallpaperUrl", m.GetKioskModeWallpaperUrl())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeWifiAllowedSsids() != nil {
        err = writer.WriteCollectionOfStringValues("kioskModeWifiAllowedSsids", m.GetKioskModeWifiAllowedSsids())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeWiFiConfigurationEnabled", m.GetKioskModeWiFiConfigurationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("locateDeviceLostModeEnabled", m.GetLocateDeviceLostModeEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("locateDeviceUserlessDisabled", m.GetLocateDeviceUserlessDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microphoneForceMute", m.GetMicrophoneForceMute())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftLauncherConfigurationEnabled", m.GetMicrosoftLauncherConfigurationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftLauncherCustomWallpaperAllowUserModification", m.GetMicrosoftLauncherCustomWallpaperAllowUserModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftLauncherCustomWallpaperEnabled", m.GetMicrosoftLauncherCustomWallpaperEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("microsoftLauncherCustomWallpaperImageUrl", m.GetMicrosoftLauncherCustomWallpaperImageUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftLauncherDockPresenceAllowUserModification", m.GetMicrosoftLauncherDockPresenceAllowUserModification())
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftLauncherDockPresenceConfiguration() != nil {
        cast := (*m.GetMicrosoftLauncherDockPresenceConfiguration()).String()
        err = writer.WriteStringValue("microsoftLauncherDockPresenceConfiguration", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftLauncherFeedAllowUserModification", m.GetMicrosoftLauncherFeedAllowUserModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftLauncherFeedEnabled", m.GetMicrosoftLauncherFeedEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftLauncherSearchBarPlacementConfiguration() != nil {
        cast := (*m.GetMicrosoftLauncherSearchBarPlacementConfiguration()).String()
        err = writer.WriteStringValue("microsoftLauncherSearchBarPlacementConfiguration", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("networkEscapeHatchAllowed", m.GetNetworkEscapeHatchAllowed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("nfcBlockOutgoingBeam", m.GetNfcBlockOutgoingBeam())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockKeyguard", m.GetPasswordBlockKeyguard())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordBlockKeyguardFeatures() != nil {
        err = writer.WriteCollectionOfStringValues("passwordBlockKeyguardFeatures", SerializeAndroidKeyguardFeature(m.GetPasswordBlockKeyguardFeatures()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationDays", m.GetPasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLetterCharacters", m.GetPasswordMinimumLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLowerCaseCharacters", m.GetPasswordMinimumLowerCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumNonLetterCharacters", m.GetPasswordMinimumNonLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumNumericCharacters", m.GetPasswordMinimumNumericCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumSymbolCharacters", m.GetPasswordMinimumSymbolCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumUpperCaseCharacters", m.GetPasswordMinimumUpperCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeScreenTimeout", m.GetPasswordMinutesOfInactivityBeforeScreenTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordPreviousPasswordCountToBlock", m.GetPasswordPreviousPasswordCountToBlock())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRequiredType() != nil {
        cast := (*m.GetPasswordRequiredType()).String()
        err = writer.WriteStringValue("passwordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRequireUnlock() != nil {
        cast := (*m.GetPasswordRequireUnlock()).String()
        err = writer.WriteStringValue("passwordRequireUnlock", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordSignInFailureCountBeforeFactoryReset", m.GetPasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("personalProfileAppsAllowInstallFromUnknownSources", m.GetPersonalProfileAppsAllowInstallFromUnknownSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("personalProfileCameraBlocked", m.GetPersonalProfileCameraBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetPersonalProfilePersonalApplications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPersonalProfilePersonalApplications()))
        for i, v := range m.GetPersonalProfilePersonalApplications() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("personalProfilePersonalApplications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPersonalProfilePlayStoreMode() != nil {
        cast := (*m.GetPersonalProfilePlayStoreMode()).String()
        err = writer.WriteStringValue("personalProfilePlayStoreMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("personalProfileScreenCaptureBlocked", m.GetPersonalProfileScreenCaptureBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetPlayStoreMode() != nil {
        cast := (*m.GetPlayStoreMode()).String()
        err = writer.WriteStringValue("playStoreMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("screenCaptureBlocked", m.GetScreenCaptureBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityCommonCriteriaModeEnabled", m.GetSecurityCommonCriteriaModeEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityDeveloperSettingsEnabled", m.GetSecurityDeveloperSettingsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireVerifyApps", m.GetSecurityRequireVerifyApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shortHelpText", m.GetShortHelpText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("statusBarBlocked", m.GetStatusBarBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetStayOnModes() != nil {
        err = writer.WriteCollectionOfStringValues("stayOnModes", SerializeAndroidDeviceOwnerBatteryPluggedMode(m.GetStayOnModes()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageAllowUsb", m.GetStorageAllowUsb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageBlockExternalMedia", m.GetStorageBlockExternalMedia())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageBlockUsbFileTransfer", m.GetStorageBlockUsbFileTransfer())
        if err != nil {
            return err
        }
    }
    if m.GetSystemUpdateFreezePeriods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSystemUpdateFreezePeriods()))
        for i, v := range m.GetSystemUpdateFreezePeriods() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("systemUpdateFreezePeriods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemUpdateInstallType() != nil {
        cast := (*m.GetSystemUpdateInstallType()).String()
        err = writer.WriteStringValue("systemUpdateInstallType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("systemUpdateWindowEndMinutesAfterMidnight", m.GetSystemUpdateWindowEndMinutesAfterMidnight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("systemUpdateWindowStartMinutesAfterMidnight", m.GetSystemUpdateWindowStartMinutesAfterMidnight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("systemWindowsBlocked", m.GetSystemWindowsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usersBlockAdd", m.GetUsersBlockAdd())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usersBlockRemove", m.GetUsersBlockRemove())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("volumeBlockAdjustment", m.GetVolumeBlockAdjustment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("vpnAlwaysOnLockdownMode", m.GetVpnAlwaysOnLockdownMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vpnAlwaysOnPackageIdentifier", m.GetVpnAlwaysOnPackageIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wifiBlockEditConfigurations", m.GetWifiBlockEditConfigurations())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wifiBlockEditPolicyDefinedConfigurations", m.GetWifiBlockEditPolicyDefinedConfigurations())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordExpirationDays", m.GetWorkProfilePasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumLength", m.GetWorkProfilePasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumLetterCharacters", m.GetWorkProfilePasswordMinimumLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumLowerCaseCharacters", m.GetWorkProfilePasswordMinimumLowerCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumNonLetterCharacters", m.GetWorkProfilePasswordMinimumNonLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumNumericCharacters", m.GetWorkProfilePasswordMinimumNumericCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumSymbolCharacters", m.GetWorkProfilePasswordMinimumSymbolCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumUpperCaseCharacters", m.GetWorkProfilePasswordMinimumUpperCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordPreviousPasswordCountToBlock", m.GetWorkProfilePasswordPreviousPasswordCountToBlock())
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfilePasswordRequiredType() != nil {
        cast := (*m.GetWorkProfilePasswordRequiredType()).String()
        err = writer.WriteStringValue("workProfilePasswordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfilePasswordRequireUnlock() != nil {
        cast := (*m.GetWorkProfilePasswordRequireUnlock()).String()
        err = writer.WriteStringValue("workProfilePasswordRequireUnlock", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordSignInFailureCountBeforeFactoryReset", m.GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountsBlockModification sets the accountsBlockModification property value. Indicates whether or not adding or removing accounts is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetAccountsBlockModification(value *bool)() {
    err := m.GetBackingStore().Set("accountsBlockModification", value)
    if err != nil {
        panic(err)
    }
}
// SetAppsAllowInstallFromUnknownSources sets the appsAllowInstallFromUnknownSources property value. Indicates whether or not the user is allowed to enable to unknown sources setting.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetAppsAllowInstallFromUnknownSources(value *bool)() {
    err := m.GetBackingStore().Set("appsAllowInstallFromUnknownSources", value)
    if err != nil {
        panic(err)
    }
}
// SetAppsAutoUpdatePolicy sets the appsAutoUpdatePolicy property value. Indicates the value of the app auto update policy. Possible values are: notConfigured, userChoice, never, wiFiOnly, always.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetAppsAutoUpdatePolicy(value *AndroidDeviceOwnerAppAutoUpdatePolicyType)() {
    err := m.GetBackingStore().Set("appsAutoUpdatePolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetAppsDefaultPermissionPolicy sets the appsDefaultPermissionPolicy property value. Indicates the permission policy for requests for runtime permissions if one is not defined for the app specifically. Possible values are: deviceDefault, prompt, autoGrant, autoDeny.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetAppsDefaultPermissionPolicy(value *AndroidDeviceOwnerDefaultAppPermissionPolicyType)() {
    err := m.GetBackingStore().Set("appsDefaultPermissionPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetAppsRecommendSkippingFirstUseHints sets the appsRecommendSkippingFirstUseHints property value. Whether or not to recommend all apps skip any first-time-use hints they may have added.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetAppsRecommendSkippingFirstUseHints(value *bool)() {
    err := m.GetBackingStore().Set("appsRecommendSkippingFirstUseHints", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureAdSharedDeviceDataClearApps sets the azureAdSharedDeviceDataClearApps property value. A list of managed apps that will have their data cleared during a global sign-out in AAD shared device mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetAzureAdSharedDeviceDataClearApps(value []AppListItemable)() {
    err := m.GetBackingStore().Set("azureAdSharedDeviceDataClearApps", value)
    if err != nil {
        panic(err)
    }
}
// SetBluetoothBlockConfiguration sets the bluetoothBlockConfiguration property value. Indicates whether or not to block a user from configuring bluetooth.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetBluetoothBlockConfiguration(value *bool)() {
    err := m.GetBackingStore().Set("bluetoothBlockConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetBluetoothBlockContactSharing sets the bluetoothBlockContactSharing property value. Indicates whether or not to block a user from sharing contacts via bluetooth.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetBluetoothBlockContactSharing(value *bool)() {
    err := m.GetBackingStore().Set("bluetoothBlockContactSharing", value)
    if err != nil {
        panic(err)
    }
}
// SetCameraBlocked sets the cameraBlocked property value. Indicates whether or not to disable the use of the camera.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetCameraBlocked(value *bool)() {
    err := m.GetBackingStore().Set("cameraBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetCellularBlockWiFiTethering sets the cellularBlockWiFiTethering property value. Indicates whether or not to block Wi-Fi tethering.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetCellularBlockWiFiTethering(value *bool)() {
    err := m.GetBackingStore().Set("cellularBlockWiFiTethering", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateCredentialConfigurationDisabled sets the certificateCredentialConfigurationDisabled property value. Indicates whether or not to block users from any certificate credential configuration.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetCertificateCredentialConfigurationDisabled(value *bool)() {
    err := m.GetBackingStore().Set("certificateCredentialConfigurationDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetCrossProfilePoliciesAllowCopyPaste sets the crossProfilePoliciesAllowCopyPaste property value. Indicates whether or not text copied from one profile (personal or work) can be pasted in the other.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetCrossProfilePoliciesAllowCopyPaste(value *bool)() {
    err := m.GetBackingStore().Set("crossProfilePoliciesAllowCopyPaste", value)
    if err != nil {
        panic(err)
    }
}
// SetCrossProfilePoliciesAllowDataSharing sets the crossProfilePoliciesAllowDataSharing property value. Indicates whether data from one profile (personal or work) can be shared with apps in the other profile. Possible values are: notConfigured, crossProfileDataSharingBlocked, dataSharingFromWorkToPersonalBlocked, crossProfileDataSharingAllowed, unkownFutureValue.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetCrossProfilePoliciesAllowDataSharing(value *AndroidDeviceOwnerCrossProfileDataSharing)() {
    err := m.GetBackingStore().Set("crossProfilePoliciesAllowDataSharing", value)
    if err != nil {
        panic(err)
    }
}
// SetCrossProfilePoliciesShowWorkContactsInPersonalProfile sets the crossProfilePoliciesShowWorkContactsInPersonalProfile property value. Indicates whether or not contacts stored in work profile are shown in personal profile contact searches/incoming calls.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetCrossProfilePoliciesShowWorkContactsInPersonalProfile(value *bool)() {
    err := m.GetBackingStore().Set("crossProfilePoliciesShowWorkContactsInPersonalProfile", value)
    if err != nil {
        panic(err)
    }
}
// SetDataRoamingBlocked sets the dataRoamingBlocked property value. Indicates whether or not to block a user from data roaming.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetDataRoamingBlocked(value *bool)() {
    err := m.GetBackingStore().Set("dataRoamingBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetDateTimeConfigurationBlocked sets the dateTimeConfigurationBlocked property value. Indicates whether or not to block the user from manually changing the date or time on the device
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetDateTimeConfigurationBlocked(value *bool)() {
    err := m.GetBackingStore().Set("dateTimeConfigurationBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetDetailedHelpText sets the detailedHelpText property value. Represents the customized detailed help text provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetDetailedHelpText(value AndroidDeviceOwnerUserFacingMessageable)() {
    err := m.GetBackingStore().Set("detailedHelpText", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceOwnerLockScreenMessage sets the deviceOwnerLockScreenMessage property value. Represents the customized lock screen message provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetDeviceOwnerLockScreenMessage(value AndroidDeviceOwnerUserFacingMessageable)() {
    err := m.GetBackingStore().Set("deviceOwnerLockScreenMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentProfile sets the enrollmentProfile property value. Android Device Owner Enrollment Profile types.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetEnrollmentProfile(value *AndroidDeviceOwnerEnrollmentProfileType)() {
    err := m.GetBackingStore().Set("enrollmentProfile", value)
    if err != nil {
        panic(err)
    }
}
// SetFactoryResetBlocked sets the factoryResetBlocked property value. Indicates whether or not the factory reset option in settings is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetFactoryResetBlocked(value *bool)() {
    err := m.GetBackingStore().Set("factoryResetBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetFactoryResetDeviceAdministratorEmails sets the factoryResetDeviceAdministratorEmails property value. List of Google account emails that will be required to authenticate after a device is factory reset before it can be set up.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetFactoryResetDeviceAdministratorEmails(value []string)() {
    err := m.GetBackingStore().Set("factoryResetDeviceAdministratorEmails", value)
    if err != nil {
        panic(err)
    }
}
// SetGlobalProxy sets the globalProxy property value. Proxy is set up directly with host, port and excluded hosts.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetGlobalProxy(value AndroidDeviceOwnerGlobalProxyable)() {
    err := m.GetBackingStore().Set("globalProxy", value)
    if err != nil {
        panic(err)
    }
}
// SetGoogleAccountsBlocked sets the googleAccountsBlocked property value. Indicates whether or not google accounts will be blocked.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetGoogleAccountsBlocked(value *bool)() {
    err := m.GetBackingStore().Set("googleAccountsBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskCustomizationDeviceSettingsBlocked sets the kioskCustomizationDeviceSettingsBlocked property value. Indicates whether a user can access the device's Settings app while in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskCustomizationDeviceSettingsBlocked(value *bool)() {
    err := m.GetBackingStore().Set("kioskCustomizationDeviceSettingsBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskCustomizationPowerButtonActionsBlocked sets the kioskCustomizationPowerButtonActionsBlocked property value. Whether the power menu is shown when a user long presses the Power button of a device in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskCustomizationPowerButtonActionsBlocked(value *bool)() {
    err := m.GetBackingStore().Set("kioskCustomizationPowerButtonActionsBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskCustomizationStatusBar sets the kioskCustomizationStatusBar property value. Indicates whether system info and notifications are disabled in Kiosk Mode. Possible values are: notConfigured, notificationsAndSystemInfoEnabled, systemInfoOnly.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskCustomizationStatusBar(value *AndroidDeviceOwnerKioskCustomizationStatusBar)() {
    err := m.GetBackingStore().Set("kioskCustomizationStatusBar", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskCustomizationSystemErrorWarnings sets the kioskCustomizationSystemErrorWarnings property value. Indicates whether system error dialogs for crashed or unresponsive apps are shown in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskCustomizationSystemErrorWarnings(value *bool)() {
    err := m.GetBackingStore().Set("kioskCustomizationSystemErrorWarnings", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskCustomizationSystemNavigation sets the kioskCustomizationSystemNavigation property value. Indicates which navigation features are enabled in Kiosk Mode. Possible values are: notConfigured, navigationEnabled, homeButtonOnly.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskCustomizationSystemNavigation(value *AndroidDeviceOwnerKioskCustomizationSystemNavigation)() {
    err := m.GetBackingStore().Set("kioskCustomizationSystemNavigation", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeAppOrderEnabled sets the kioskModeAppOrderEnabled property value. Whether or not to enable app ordering in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeAppOrderEnabled(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeAppOrderEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeAppPositions sets the kioskModeAppPositions property value. The ordering of items on Kiosk Mode Managed Home Screen. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeAppPositions(value []AndroidDeviceOwnerKioskModeAppPositionItemable)() {
    err := m.GetBackingStore().Set("kioskModeAppPositions", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeApps sets the kioskModeApps property value. A list of managed apps that will be shown when the device is in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeApps(value []AppListItemable)() {
    err := m.GetBackingStore().Set("kioskModeApps", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeAppsInFolderOrderedByName sets the kioskModeAppsInFolderOrderedByName property value. Whether or not to alphabetize applications within a folder in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeAppsInFolderOrderedByName(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeAppsInFolderOrderedByName", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeBluetoothConfigurationEnabled sets the kioskModeBluetoothConfigurationEnabled property value. Whether or not to allow a user to configure Bluetooth settings in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeBluetoothConfigurationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeBluetoothConfigurationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeDebugMenuEasyAccessEnabled sets the kioskModeDebugMenuEasyAccessEnabled property value. Whether or not to allow a user to easy access to the debug menu in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeDebugMenuEasyAccessEnabled(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeDebugMenuEasyAccessEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeExitCode sets the kioskModeExitCode property value. Exit code to allow a user to escape from Kiosk Mode when the device is in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeExitCode(value *string)() {
    err := m.GetBackingStore().Set("kioskModeExitCode", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeFlashlightConfigurationEnabled sets the kioskModeFlashlightConfigurationEnabled property value. Whether or not to allow a user to use the flashlight in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeFlashlightConfigurationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeFlashlightConfigurationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeFolderIcon sets the kioskModeFolderIcon property value. Folder icon configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, darkSquare, darkCircle, lightSquare, lightCircle.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeFolderIcon(value *AndroidDeviceOwnerKioskModeFolderIcon)() {
    err := m.GetBackingStore().Set("kioskModeFolderIcon", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeGridHeight sets the kioskModeGridHeight property value. Number of rows for Managed Home Screen grid with app ordering enabled in Kiosk Mode. Valid values 1 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeGridHeight(value *int32)() {
    err := m.GetBackingStore().Set("kioskModeGridHeight", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeGridWidth sets the kioskModeGridWidth property value. Number of columns for Managed Home Screen grid with app ordering enabled in Kiosk Mode. Valid values 1 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeGridWidth(value *int32)() {
    err := m.GetBackingStore().Set("kioskModeGridWidth", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeIconSize sets the kioskModeIconSize property value. Icon size configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, smallest, small, regular, large, largest.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeIconSize(value *AndroidDeviceOwnerKioskModeIconSize)() {
    err := m.GetBackingStore().Set("kioskModeIconSize", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeLockHomeScreen sets the kioskModeLockHomeScreen property value. Whether or not to lock home screen to the end user in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeLockHomeScreen(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeLockHomeScreen", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeManagedFolders sets the kioskModeManagedFolders property value. A list of managed folders for a device in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedFolders(value []AndroidDeviceOwnerKioskModeManagedFolderable)() {
    err := m.GetBackingStore().Set("kioskModeManagedFolders", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeManagedHomeScreenAutoSignout sets the kioskModeManagedHomeScreenAutoSignout property value. Whether or not to automatically sign-out of MHS and Shared device mode applications after inactive for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenAutoSignout(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeManagedHomeScreenAutoSignout", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds sets the kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds property value. Number of seconds to give user notice before automatically signing them out for Managed Home Screen. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds sets the kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds property value. Number of seconds device is inactive before automatically signing user out for Managed Home Screen. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeManagedHomeScreenPinComplexity sets the kioskModeManagedHomeScreenPinComplexity property value. Complexity of PIN for sign-in session for Managed Home Screen. Possible values are: notConfigured, simple, complex.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenPinComplexity(value *KioskModeManagedHomeScreenPinComplexity)() {
    err := m.GetBackingStore().Set("kioskModeManagedHomeScreenPinComplexity", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeManagedHomeScreenPinRequired sets the kioskModeManagedHomeScreenPinRequired property value. Whether or not require user to set a PIN for sign-in session for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenPinRequired(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeManagedHomeScreenPinRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeManagedHomeScreenPinRequiredToResume sets the kioskModeManagedHomeScreenPinRequiredToResume property value. Whether or not required user to enter session PIN if screensaver has appeared for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenPinRequiredToResume(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeManagedHomeScreenPinRequiredToResume", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeManagedHomeScreenSignInBackground sets the kioskModeManagedHomeScreenSignInBackground property value. Custom URL background for sign-in screen for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenSignInBackground(value *string)() {
    err := m.GetBackingStore().Set("kioskModeManagedHomeScreenSignInBackground", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeManagedHomeScreenSignInBrandingLogo sets the kioskModeManagedHomeScreenSignInBrandingLogo property value. Custom URL branding logo for sign-in screen and session pin page for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenSignInBrandingLogo(value *string)() {
    err := m.GetBackingStore().Set("kioskModeManagedHomeScreenSignInBrandingLogo", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeManagedHomeScreenSignInEnabled sets the kioskModeManagedHomeScreenSignInEnabled property value. Whether or not show sign-in screen for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenSignInEnabled(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeManagedHomeScreenSignInEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeManagedSettingsEntryDisabled sets the kioskModeManagedSettingsEntryDisabled property value. Whether or not to display the Managed Settings entry point on the managed home screen in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedSettingsEntryDisabled(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeManagedSettingsEntryDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeMediaVolumeConfigurationEnabled sets the kioskModeMediaVolumeConfigurationEnabled property value. Whether or not to allow a user to change the media volume in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeMediaVolumeConfigurationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeMediaVolumeConfigurationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeScreenOrientation sets the kioskModeScreenOrientation property value. Screen orientation configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, portrait, landscape, autoRotate.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeScreenOrientation(value *AndroidDeviceOwnerKioskModeScreenOrientation)() {
    err := m.GetBackingStore().Set("kioskModeScreenOrientation", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeScreenSaverConfigurationEnabled sets the kioskModeScreenSaverConfigurationEnabled property value. Whether or not to enable screen saver mode or not in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeScreenSaverConfigurationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeScreenSaverConfigurationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeScreenSaverDetectMediaDisabled sets the kioskModeScreenSaverDetectMediaDisabled property value. Whether or not the device screen should show the screen saver if audio/video is playing in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeScreenSaverDetectMediaDisabled(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeScreenSaverDetectMediaDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeScreenSaverDisplayTimeInSeconds sets the kioskModeScreenSaverDisplayTimeInSeconds property value. The number of seconds that the device will display the screen saver for in Kiosk Mode. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeScreenSaverDisplayTimeInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("kioskModeScreenSaverDisplayTimeInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeScreenSaverImageUrl sets the kioskModeScreenSaverImageUrl property value. URL for an image that will be the device's screen saver in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeScreenSaverImageUrl(value *string)() {
    err := m.GetBackingStore().Set("kioskModeScreenSaverImageUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeScreenSaverStartDelayInSeconds sets the kioskModeScreenSaverStartDelayInSeconds property value. The number of seconds the device needs to be inactive for before the screen saver is shown in Kiosk Mode. Valid values 1 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeScreenSaverStartDelayInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("kioskModeScreenSaverStartDelayInSeconds", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeShowAppNotificationBadge sets the kioskModeShowAppNotificationBadge property value. Whether or not to display application notification badges in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeShowAppNotificationBadge(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeShowAppNotificationBadge", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeShowDeviceInfo sets the kioskModeShowDeviceInfo property value. Whether or not to allow a user to access basic device information.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeShowDeviceInfo(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeShowDeviceInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeUseManagedHomeScreenApp sets the kioskModeUseManagedHomeScreenApp property value. Whether or not to use single app kiosk mode or multi-app kiosk mode. Possible values are: notConfigured, singleAppMode, multiAppMode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeUseManagedHomeScreenApp(value *KioskModeType)() {
    err := m.GetBackingStore().Set("kioskModeUseManagedHomeScreenApp", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeVirtualHomeButtonEnabled sets the kioskModeVirtualHomeButtonEnabled property value. Whether or not to display a virtual home button when the device is in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeVirtualHomeButtonEnabled(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeVirtualHomeButtonEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeVirtualHomeButtonType sets the kioskModeVirtualHomeButtonType property value. Indicates whether the virtual home button is a swipe up home button or a floating home button. Possible values are: notConfigured, swipeUp, floating.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeVirtualHomeButtonType(value *AndroidDeviceOwnerVirtualHomeButtonType)() {
    err := m.GetBackingStore().Set("kioskModeVirtualHomeButtonType", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeWallpaperUrl sets the kioskModeWallpaperUrl property value. URL to a publicly accessible image to use for the wallpaper when the device is in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeWallpaperUrl(value *string)() {
    err := m.GetBackingStore().Set("kioskModeWallpaperUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeWifiAllowedSsids sets the kioskModeWifiAllowedSsids property value. The restricted set of WIFI SSIDs available for the user to configure in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeWifiAllowedSsids(value []string)() {
    err := m.GetBackingStore().Set("kioskModeWifiAllowedSsids", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskModeWiFiConfigurationEnabled sets the kioskModeWiFiConfigurationEnabled property value. Whether or not to allow a user to configure Wi-Fi settings in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeWiFiConfigurationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("kioskModeWiFiConfigurationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetLocateDeviceLostModeEnabled sets the locateDeviceLostModeEnabled property value. Indicates whether or not LocateDevice for devices with lost mode (COBO, COPE) is enabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetLocateDeviceLostModeEnabled(value *bool)() {
    err := m.GetBackingStore().Set("locateDeviceLostModeEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetLocateDeviceUserlessDisabled sets the locateDeviceUserlessDisabled property value. Indicates whether or not LocateDevice for userless (COSU) devices is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetLocateDeviceUserlessDisabled(value *bool)() {
    err := m.GetBackingStore().Set("locateDeviceUserlessDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrophoneForceMute sets the microphoneForceMute property value. Indicates whether or not to block unmuting the microphone on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrophoneForceMute(value *bool)() {
    err := m.GetBackingStore().Set("microphoneForceMute", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftLauncherConfigurationEnabled sets the microsoftLauncherConfigurationEnabled property value. Indicates whether or not to you want configure Microsoft Launcher.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherConfigurationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("microsoftLauncherConfigurationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftLauncherCustomWallpaperAllowUserModification sets the microsoftLauncherCustomWallpaperAllowUserModification property value. Indicates whether or not the user can modify the wallpaper to personalize their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherCustomWallpaperAllowUserModification(value *bool)() {
    err := m.GetBackingStore().Set("microsoftLauncherCustomWallpaperAllowUserModification", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftLauncherCustomWallpaperEnabled sets the microsoftLauncherCustomWallpaperEnabled property value. Indicates whether or not to configure the wallpaper on the targeted devices.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherCustomWallpaperEnabled(value *bool)() {
    err := m.GetBackingStore().Set("microsoftLauncherCustomWallpaperEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftLauncherCustomWallpaperImageUrl sets the microsoftLauncherCustomWallpaperImageUrl property value. Indicates the URL for the image file to use as the wallpaper on the targeted devices.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherCustomWallpaperImageUrl(value *string)() {
    err := m.GetBackingStore().Set("microsoftLauncherCustomWallpaperImageUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftLauncherDockPresenceAllowUserModification sets the microsoftLauncherDockPresenceAllowUserModification property value. Indicates whether or not the user can modify the device dock configuration on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherDockPresenceAllowUserModification(value *bool)() {
    err := m.GetBackingStore().Set("microsoftLauncherDockPresenceAllowUserModification", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftLauncherDockPresenceConfiguration sets the microsoftLauncherDockPresenceConfiguration property value. Indicates whether or not you want to configure the device dock. Possible values are: notConfigured, show, hide, disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherDockPresenceConfiguration(value *MicrosoftLauncherDockPresence)() {
    err := m.GetBackingStore().Set("microsoftLauncherDockPresenceConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftLauncherFeedAllowUserModification sets the microsoftLauncherFeedAllowUserModification property value. Indicates whether or not the user can modify the launcher feed on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherFeedAllowUserModification(value *bool)() {
    err := m.GetBackingStore().Set("microsoftLauncherFeedAllowUserModification", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftLauncherFeedEnabled sets the microsoftLauncherFeedEnabled property value. Indicates whether or not you want to enable the launcher feed on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherFeedEnabled(value *bool)() {
    err := m.GetBackingStore().Set("microsoftLauncherFeedEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftLauncherSearchBarPlacementConfiguration sets the microsoftLauncherSearchBarPlacementConfiguration property value. Indicates the search bar placement configuration on the device. Possible values are: notConfigured, top, bottom, hide.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherSearchBarPlacementConfiguration(value *MicrosoftLauncherSearchBarPlacement)() {
    err := m.GetBackingStore().Set("microsoftLauncherSearchBarPlacementConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkEscapeHatchAllowed sets the networkEscapeHatchAllowed property value. Indicates whether or not the device will allow connecting to a temporary network connection at boot time.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetNetworkEscapeHatchAllowed(value *bool)() {
    err := m.GetBackingStore().Set("networkEscapeHatchAllowed", value)
    if err != nil {
        panic(err)
    }
}
// SetNfcBlockOutgoingBeam sets the nfcBlockOutgoingBeam property value. Indicates whether or not to block NFC outgoing beam.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetNfcBlockOutgoingBeam(value *bool)() {
    err := m.GetBackingStore().Set("nfcBlockOutgoingBeam", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockKeyguard sets the passwordBlockKeyguard property value. Indicates whether or not the keyguard is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordBlockKeyguard(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockKeyguard", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockKeyguardFeatures sets the passwordBlockKeyguardFeatures property value. List of device keyguard features to block. This collection can contain a maximum of 11 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordBlockKeyguardFeatures(value []AndroidKeyguardFeature)() {
    err := m.GetBackingStore().Set("passwordBlockKeyguardFeatures", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Indicates the amount of time that a password can be set for before it expires and a new password will be required. Valid values 1 to 365
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordExpirationDays(value *int32)() {
    err := m.GetBackingStore().Set("passwordExpirationDays", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Indicates the minimum length of the password required on the device. Valid values 4 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumLetterCharacters sets the passwordMinimumLetterCharacters property value. Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumLetterCharacters(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumLetterCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumLowerCaseCharacters sets the passwordMinimumLowerCaseCharacters property value. Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumLowerCaseCharacters(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumLowerCaseCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumNonLetterCharacters sets the passwordMinimumNonLetterCharacters property value. Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumNonLetterCharacters(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumNonLetterCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumNumericCharacters sets the passwordMinimumNumericCharacters property value. Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumNumericCharacters(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumNumericCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumSymbolCharacters sets the passwordMinimumSymbolCharacters property value. Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumSymbolCharacters(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumSymbolCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumUpperCaseCharacters sets the passwordMinimumUpperCaseCharacters property value. Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumUpperCaseCharacters(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumUpperCaseCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinutesOfInactivityBeforeScreenTimeout", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordPreviousPasswordCountToBlock sets the passwordPreviousPasswordCountToBlock property value. Indicates the length of password history, where the user will not be able to enter a new password that is the same as any password in the history. Valid values 0 to 24
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordPreviousPasswordCountToBlock(value *int32)() {
    err := m.GetBackingStore().Set("passwordPreviousPasswordCountToBlock", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Indicates the minimum password quality required on the device. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)() {
    err := m.GetBackingStore().Set("passwordRequiredType", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequireUnlock sets the passwordRequireUnlock property value. Indicates the timeout period after which a device must be unlocked using a form of strong authentication. Possible values are: deviceDefault, daily, unkownFutureValue.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordRequireUnlock(value *AndroidDeviceOwnerRequiredPasswordUnlock)() {
    err := m.GetBackingStore().Set("passwordRequireUnlock", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    err := m.GetBackingStore().Set("passwordSignInFailureCountBeforeFactoryReset", value)
    if err != nil {
        panic(err)
    }
}
// SetPersonalProfileAppsAllowInstallFromUnknownSources sets the personalProfileAppsAllowInstallFromUnknownSources property value. Indicates whether the user can install apps from unknown sources on the personal profile.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPersonalProfileAppsAllowInstallFromUnknownSources(value *bool)() {
    err := m.GetBackingStore().Set("personalProfileAppsAllowInstallFromUnknownSources", value)
    if err != nil {
        panic(err)
    }
}
// SetPersonalProfileCameraBlocked sets the personalProfileCameraBlocked property value. Indicates whether to disable the use of the camera on the personal profile.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPersonalProfileCameraBlocked(value *bool)() {
    err := m.GetBackingStore().Set("personalProfileCameraBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetPersonalProfilePersonalApplications sets the personalProfilePersonalApplications property value. Policy applied to applications in the personal profile. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPersonalProfilePersonalApplications(value []AppListItemable)() {
    err := m.GetBackingStore().Set("personalProfilePersonalApplications", value)
    if err != nil {
        panic(err)
    }
}
// SetPersonalProfilePlayStoreMode sets the personalProfilePlayStoreMode property value. Used together with PersonalProfilePersonalApplications to control how apps in the personal profile are allowed or blocked. Possible values are: notConfigured, blockedApps, allowedApps.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPersonalProfilePlayStoreMode(value *PersonalProfilePersonalPlayStoreMode)() {
    err := m.GetBackingStore().Set("personalProfilePlayStoreMode", value)
    if err != nil {
        panic(err)
    }
}
// SetPersonalProfileScreenCaptureBlocked sets the personalProfileScreenCaptureBlocked property value. Indicates whether to disable the capability to take screenshots on the personal profile.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPersonalProfileScreenCaptureBlocked(value *bool)() {
    err := m.GetBackingStore().Set("personalProfileScreenCaptureBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetPlayStoreMode sets the playStoreMode property value. Indicates the Play Store mode of the device. Possible values are: notConfigured, allowList, blockList.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPlayStoreMode(value *AndroidDeviceOwnerPlayStoreMode)() {
    err := m.GetBackingStore().Set("playStoreMode", value)
    if err != nil {
        panic(err)
    }
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether or not to disable the capability to take screenshots.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetScreenCaptureBlocked(value *bool)() {
    err := m.GetBackingStore().Set("screenCaptureBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityCommonCriteriaModeEnabled sets the securityCommonCriteriaModeEnabled property value. Represents the security common criteria mode enabled provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSecurityCommonCriteriaModeEnabled(value *bool)() {
    err := m.GetBackingStore().Set("securityCommonCriteriaModeEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityDeveloperSettingsEnabled sets the securityDeveloperSettingsEnabled property value. Indicates whether or not the user is allowed to access developer settings like developer options and safe boot on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSecurityDeveloperSettingsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("securityDeveloperSettingsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityRequireVerifyApps sets the securityRequireVerifyApps property value. Indicates whether or not verify apps is required.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSecurityRequireVerifyApps(value *bool)() {
    err := m.GetBackingStore().Set("securityRequireVerifyApps", value)
    if err != nil {
        panic(err)
    }
}
// SetShortHelpText sets the shortHelpText property value. Represents the customized short help text provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetShortHelpText(value AndroidDeviceOwnerUserFacingMessageable)() {
    err := m.GetBackingStore().Set("shortHelpText", value)
    if err != nil {
        panic(err)
    }
}
// SetStatusBarBlocked sets the statusBarBlocked property value. Indicates whether or the status bar is disabled, including notifications, quick settings and other screen overlays.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetStatusBarBlocked(value *bool)() {
    err := m.GetBackingStore().Set("statusBarBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetStayOnModes sets the stayOnModes property value. List of modes in which the device's display will stay powered-on. This collection can contain a maximum of 4 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetStayOnModes(value []AndroidDeviceOwnerBatteryPluggedMode)() {
    err := m.GetBackingStore().Set("stayOnModes", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageAllowUsb sets the storageAllowUsb property value. Indicates whether or not to allow USB mass storage.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetStorageAllowUsb(value *bool)() {
    err := m.GetBackingStore().Set("storageAllowUsb", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageBlockExternalMedia sets the storageBlockExternalMedia property value. Indicates whether or not to block external media.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetStorageBlockExternalMedia(value *bool)() {
    err := m.GetBackingStore().Set("storageBlockExternalMedia", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageBlockUsbFileTransfer sets the storageBlockUsbFileTransfer property value. Indicates whether or not to block USB file transfer.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetStorageBlockUsbFileTransfer(value *bool)() {
    err := m.GetBackingStore().Set("storageBlockUsbFileTransfer", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemUpdateFreezePeriods sets the systemUpdateFreezePeriods property value. Indicates the annually repeating time periods during which system updates are postponed. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSystemUpdateFreezePeriods(value []AndroidDeviceOwnerSystemUpdateFreezePeriodable)() {
    err := m.GetBackingStore().Set("systemUpdateFreezePeriods", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemUpdateInstallType sets the systemUpdateInstallType property value. The type of system update configuration. Possible values are: deviceDefault, postpone, windowed, automatic.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSystemUpdateInstallType(value *AndroidDeviceOwnerSystemUpdateInstallType)() {
    err := m.GetBackingStore().Set("systemUpdateInstallType", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemUpdateWindowEndMinutesAfterMidnight sets the systemUpdateWindowEndMinutesAfterMidnight property value. Indicates the number of minutes after midnight that the system update window ends. Valid values 0 to 1440
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSystemUpdateWindowEndMinutesAfterMidnight(value *int32)() {
    err := m.GetBackingStore().Set("systemUpdateWindowEndMinutesAfterMidnight", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemUpdateWindowStartMinutesAfterMidnight sets the systemUpdateWindowStartMinutesAfterMidnight property value. Indicates the number of minutes after midnight that the system update window starts. Valid values 0 to 1440
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSystemUpdateWindowStartMinutesAfterMidnight(value *int32)() {
    err := m.GetBackingStore().Set("systemUpdateWindowStartMinutesAfterMidnight", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemWindowsBlocked sets the systemWindowsBlocked property value. Whether or not to block Android system prompt windows, like toasts, phone activities, and system alerts.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSystemWindowsBlocked(value *bool)() {
    err := m.GetBackingStore().Set("systemWindowsBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetUsersBlockAdd sets the usersBlockAdd property value. Indicates whether or not adding users and profiles is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetUsersBlockAdd(value *bool)() {
    err := m.GetBackingStore().Set("usersBlockAdd", value)
    if err != nil {
        panic(err)
    }
}
// SetUsersBlockRemove sets the usersBlockRemove property value. Indicates whether or not to disable removing other users from the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetUsersBlockRemove(value *bool)() {
    err := m.GetBackingStore().Set("usersBlockRemove", value)
    if err != nil {
        panic(err)
    }
}
// SetVolumeBlockAdjustment sets the volumeBlockAdjustment property value. Indicates whether or not adjusting the master volume is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetVolumeBlockAdjustment(value *bool)() {
    err := m.GetBackingStore().Set("volumeBlockAdjustment", value)
    if err != nil {
        panic(err)
    }
}
// SetVpnAlwaysOnLockdownMode sets the vpnAlwaysOnLockdownMode property value. If an always on VPN package name is specified, whether or not to lock network traffic when that VPN is disconnected.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetVpnAlwaysOnLockdownMode(value *bool)() {
    err := m.GetBackingStore().Set("vpnAlwaysOnLockdownMode", value)
    if err != nil {
        panic(err)
    }
}
// SetVpnAlwaysOnPackageIdentifier sets the vpnAlwaysOnPackageIdentifier property value. Android app package name for app that will handle an always-on VPN connection.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetVpnAlwaysOnPackageIdentifier(value *string)() {
    err := m.GetBackingStore().Set("vpnAlwaysOnPackageIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetWifiBlockEditConfigurations sets the wifiBlockEditConfigurations property value. Indicates whether or not to block the user from editing the wifi connection settings.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWifiBlockEditConfigurations(value *bool)() {
    err := m.GetBackingStore().Set("wifiBlockEditConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetWifiBlockEditPolicyDefinedConfigurations sets the wifiBlockEditPolicyDefinedConfigurations property value. Indicates whether or not to block the user from editing just the networks defined by the policy.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWifiBlockEditPolicyDefinedConfigurations(value *bool)() {
    err := m.GetBackingStore().Set("wifiBlockEditPolicyDefinedConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordExpirationDays sets the workProfilePasswordExpirationDays property value. Indicates the number of days that a work profile password can be set before it expires and a new password will be required. Valid values 1 to 365
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordExpirationDays(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordExpirationDays", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinimumLength sets the workProfilePasswordMinimumLength property value. Indicates the minimum length of the work profile password. Valid values 4 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinimumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinimumLetterCharacters sets the workProfilePasswordMinimumLetterCharacters property value. Indicates the minimum number of letter characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLetterCharacters(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinimumLetterCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinimumLowerCaseCharacters sets the workProfilePasswordMinimumLowerCaseCharacters property value. Indicates the minimum number of lower-case characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLowerCaseCharacters(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinimumLowerCaseCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinimumNonLetterCharacters sets the workProfilePasswordMinimumNonLetterCharacters property value. Indicates the minimum number of non-letter characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumNonLetterCharacters(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinimumNonLetterCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinimumNumericCharacters sets the workProfilePasswordMinimumNumericCharacters property value. Indicates the minimum number of numeric characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumNumericCharacters(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinimumNumericCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinimumSymbolCharacters sets the workProfilePasswordMinimumSymbolCharacters property value. Indicates the minimum number of symbol characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumSymbolCharacters(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinimumSymbolCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordMinimumUpperCaseCharacters sets the workProfilePasswordMinimumUpperCaseCharacters property value. Indicates the minimum number of upper-case letter characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumUpperCaseCharacters(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordMinimumUpperCaseCharacters", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordPreviousPasswordCountToBlock sets the workProfilePasswordPreviousPasswordCountToBlock property value. Indicates the length of the work profile password history, where the user will not be able to enter a new password that is the same as any password in the history. Valid values 0 to 24
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordPreviousPasswordCountToBlock(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordPreviousPasswordCountToBlock", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordRequiredType sets the workProfilePasswordRequiredType property value. Indicates the minimum password quality required on the work profile password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)() {
    err := m.GetBackingStore().Set("workProfilePasswordRequiredType", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordRequireUnlock sets the workProfilePasswordRequireUnlock property value. Indicates the timeout period after which a work profile must be unlocked using a form of strong authentication. Possible values are: deviceDefault, daily, unkownFutureValue.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordRequireUnlock(value *AndroidDeviceOwnerRequiredPasswordUnlock)() {
    err := m.GetBackingStore().Set("workProfilePasswordRequireUnlock", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset sets the workProfilePasswordSignInFailureCountBeforeFactoryReset property value. Indicates the number of times a user can enter an incorrect work profile password before the device is wiped. Valid values 4 to 11
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    err := m.GetBackingStore().Set("workProfilePasswordSignInFailureCountBeforeFactoryReset", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerGeneralDeviceConfigurationable 
type AndroidDeviceOwnerGeneralDeviceConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountsBlockModification()(*bool)
    GetAppsAllowInstallFromUnknownSources()(*bool)
    GetAppsAutoUpdatePolicy()(*AndroidDeviceOwnerAppAutoUpdatePolicyType)
    GetAppsDefaultPermissionPolicy()(*AndroidDeviceOwnerDefaultAppPermissionPolicyType)
    GetAppsRecommendSkippingFirstUseHints()(*bool)
    GetAzureAdSharedDeviceDataClearApps()([]AppListItemable)
    GetBluetoothBlockConfiguration()(*bool)
    GetBluetoothBlockContactSharing()(*bool)
    GetCameraBlocked()(*bool)
    GetCellularBlockWiFiTethering()(*bool)
    GetCertificateCredentialConfigurationDisabled()(*bool)
    GetCrossProfilePoliciesAllowCopyPaste()(*bool)
    GetCrossProfilePoliciesAllowDataSharing()(*AndroidDeviceOwnerCrossProfileDataSharing)
    GetCrossProfilePoliciesShowWorkContactsInPersonalProfile()(*bool)
    GetDataRoamingBlocked()(*bool)
    GetDateTimeConfigurationBlocked()(*bool)
    GetDetailedHelpText()(AndroidDeviceOwnerUserFacingMessageable)
    GetDeviceOwnerLockScreenMessage()(AndroidDeviceOwnerUserFacingMessageable)
    GetEnrollmentProfile()(*AndroidDeviceOwnerEnrollmentProfileType)
    GetFactoryResetBlocked()(*bool)
    GetFactoryResetDeviceAdministratorEmails()([]string)
    GetGlobalProxy()(AndroidDeviceOwnerGlobalProxyable)
    GetGoogleAccountsBlocked()(*bool)
    GetKioskCustomizationDeviceSettingsBlocked()(*bool)
    GetKioskCustomizationPowerButtonActionsBlocked()(*bool)
    GetKioskCustomizationStatusBar()(*AndroidDeviceOwnerKioskCustomizationStatusBar)
    GetKioskCustomizationSystemErrorWarnings()(*bool)
    GetKioskCustomizationSystemNavigation()(*AndroidDeviceOwnerKioskCustomizationSystemNavigation)
    GetKioskModeAppOrderEnabled()(*bool)
    GetKioskModeAppPositions()([]AndroidDeviceOwnerKioskModeAppPositionItemable)
    GetKioskModeApps()([]AppListItemable)
    GetKioskModeAppsInFolderOrderedByName()(*bool)
    GetKioskModeBluetoothConfigurationEnabled()(*bool)
    GetKioskModeDebugMenuEasyAccessEnabled()(*bool)
    GetKioskModeExitCode()(*string)
    GetKioskModeFlashlightConfigurationEnabled()(*bool)
    GetKioskModeFolderIcon()(*AndroidDeviceOwnerKioskModeFolderIcon)
    GetKioskModeGridHeight()(*int32)
    GetKioskModeGridWidth()(*int32)
    GetKioskModeIconSize()(*AndroidDeviceOwnerKioskModeIconSize)
    GetKioskModeLockHomeScreen()(*bool)
    GetKioskModeManagedFolders()([]AndroidDeviceOwnerKioskModeManagedFolderable)
    GetKioskModeManagedHomeScreenAutoSignout()(*bool)
    GetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds()(*int32)
    GetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds()(*int32)
    GetKioskModeManagedHomeScreenPinComplexity()(*KioskModeManagedHomeScreenPinComplexity)
    GetKioskModeManagedHomeScreenPinRequired()(*bool)
    GetKioskModeManagedHomeScreenPinRequiredToResume()(*bool)
    GetKioskModeManagedHomeScreenSignInBackground()(*string)
    GetKioskModeManagedHomeScreenSignInBrandingLogo()(*string)
    GetKioskModeManagedHomeScreenSignInEnabled()(*bool)
    GetKioskModeManagedSettingsEntryDisabled()(*bool)
    GetKioskModeMediaVolumeConfigurationEnabled()(*bool)
    GetKioskModeScreenOrientation()(*AndroidDeviceOwnerKioskModeScreenOrientation)
    GetKioskModeScreenSaverConfigurationEnabled()(*bool)
    GetKioskModeScreenSaverDetectMediaDisabled()(*bool)
    GetKioskModeScreenSaverDisplayTimeInSeconds()(*int32)
    GetKioskModeScreenSaverImageUrl()(*string)
    GetKioskModeScreenSaverStartDelayInSeconds()(*int32)
    GetKioskModeShowAppNotificationBadge()(*bool)
    GetKioskModeShowDeviceInfo()(*bool)
    GetKioskModeUseManagedHomeScreenApp()(*KioskModeType)
    GetKioskModeVirtualHomeButtonEnabled()(*bool)
    GetKioskModeVirtualHomeButtonType()(*AndroidDeviceOwnerVirtualHomeButtonType)
    GetKioskModeWallpaperUrl()(*string)
    GetKioskModeWifiAllowedSsids()([]string)
    GetKioskModeWiFiConfigurationEnabled()(*bool)
    GetLocateDeviceLostModeEnabled()(*bool)
    GetLocateDeviceUserlessDisabled()(*bool)
    GetMicrophoneForceMute()(*bool)
    GetMicrosoftLauncherConfigurationEnabled()(*bool)
    GetMicrosoftLauncherCustomWallpaperAllowUserModification()(*bool)
    GetMicrosoftLauncherCustomWallpaperEnabled()(*bool)
    GetMicrosoftLauncherCustomWallpaperImageUrl()(*string)
    GetMicrosoftLauncherDockPresenceAllowUserModification()(*bool)
    GetMicrosoftLauncherDockPresenceConfiguration()(*MicrosoftLauncherDockPresence)
    GetMicrosoftLauncherFeedAllowUserModification()(*bool)
    GetMicrosoftLauncherFeedEnabled()(*bool)
    GetMicrosoftLauncherSearchBarPlacementConfiguration()(*MicrosoftLauncherSearchBarPlacement)
    GetNetworkEscapeHatchAllowed()(*bool)
    GetNfcBlockOutgoingBeam()(*bool)
    GetPasswordBlockKeyguard()(*bool)
    GetPasswordBlockKeyguardFeatures()([]AndroidKeyguardFeature)
    GetPasswordExpirationDays()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinimumLetterCharacters()(*int32)
    GetPasswordMinimumLowerCaseCharacters()(*int32)
    GetPasswordMinimumNonLetterCharacters()(*int32)
    GetPasswordMinimumNumericCharacters()(*int32)
    GetPasswordMinimumSymbolCharacters()(*int32)
    GetPasswordMinimumUpperCaseCharacters()(*int32)
    GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32)
    GetPasswordPreviousPasswordCountToBlock()(*int32)
    GetPasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType)
    GetPasswordRequireUnlock()(*AndroidDeviceOwnerRequiredPasswordUnlock)
    GetPasswordSignInFailureCountBeforeFactoryReset()(*int32)
    GetPersonalProfileAppsAllowInstallFromUnknownSources()(*bool)
    GetPersonalProfileCameraBlocked()(*bool)
    GetPersonalProfilePersonalApplications()([]AppListItemable)
    GetPersonalProfilePlayStoreMode()(*PersonalProfilePersonalPlayStoreMode)
    GetPersonalProfileScreenCaptureBlocked()(*bool)
    GetPlayStoreMode()(*AndroidDeviceOwnerPlayStoreMode)
    GetScreenCaptureBlocked()(*bool)
    GetSecurityCommonCriteriaModeEnabled()(*bool)
    GetSecurityDeveloperSettingsEnabled()(*bool)
    GetSecurityRequireVerifyApps()(*bool)
    GetShortHelpText()(AndroidDeviceOwnerUserFacingMessageable)
    GetStatusBarBlocked()(*bool)
    GetStayOnModes()([]AndroidDeviceOwnerBatteryPluggedMode)
    GetStorageAllowUsb()(*bool)
    GetStorageBlockExternalMedia()(*bool)
    GetStorageBlockUsbFileTransfer()(*bool)
    GetSystemUpdateFreezePeriods()([]AndroidDeviceOwnerSystemUpdateFreezePeriodable)
    GetSystemUpdateInstallType()(*AndroidDeviceOwnerSystemUpdateInstallType)
    GetSystemUpdateWindowEndMinutesAfterMidnight()(*int32)
    GetSystemUpdateWindowStartMinutesAfterMidnight()(*int32)
    GetSystemWindowsBlocked()(*bool)
    GetUsersBlockAdd()(*bool)
    GetUsersBlockRemove()(*bool)
    GetVolumeBlockAdjustment()(*bool)
    GetVpnAlwaysOnLockdownMode()(*bool)
    GetVpnAlwaysOnPackageIdentifier()(*string)
    GetWifiBlockEditConfigurations()(*bool)
    GetWifiBlockEditPolicyDefinedConfigurations()(*bool)
    GetWorkProfilePasswordExpirationDays()(*int32)
    GetWorkProfilePasswordMinimumLength()(*int32)
    GetWorkProfilePasswordMinimumLetterCharacters()(*int32)
    GetWorkProfilePasswordMinimumLowerCaseCharacters()(*int32)
    GetWorkProfilePasswordMinimumNonLetterCharacters()(*int32)
    GetWorkProfilePasswordMinimumNumericCharacters()(*int32)
    GetWorkProfilePasswordMinimumSymbolCharacters()(*int32)
    GetWorkProfilePasswordMinimumUpperCaseCharacters()(*int32)
    GetWorkProfilePasswordPreviousPasswordCountToBlock()(*int32)
    GetWorkProfilePasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType)
    GetWorkProfilePasswordRequireUnlock()(*AndroidDeviceOwnerRequiredPasswordUnlock)
    GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset()(*int32)
    SetAccountsBlockModification(value *bool)()
    SetAppsAllowInstallFromUnknownSources(value *bool)()
    SetAppsAutoUpdatePolicy(value *AndroidDeviceOwnerAppAutoUpdatePolicyType)()
    SetAppsDefaultPermissionPolicy(value *AndroidDeviceOwnerDefaultAppPermissionPolicyType)()
    SetAppsRecommendSkippingFirstUseHints(value *bool)()
    SetAzureAdSharedDeviceDataClearApps(value []AppListItemable)()
    SetBluetoothBlockConfiguration(value *bool)()
    SetBluetoothBlockContactSharing(value *bool)()
    SetCameraBlocked(value *bool)()
    SetCellularBlockWiFiTethering(value *bool)()
    SetCertificateCredentialConfigurationDisabled(value *bool)()
    SetCrossProfilePoliciesAllowCopyPaste(value *bool)()
    SetCrossProfilePoliciesAllowDataSharing(value *AndroidDeviceOwnerCrossProfileDataSharing)()
    SetCrossProfilePoliciesShowWorkContactsInPersonalProfile(value *bool)()
    SetDataRoamingBlocked(value *bool)()
    SetDateTimeConfigurationBlocked(value *bool)()
    SetDetailedHelpText(value AndroidDeviceOwnerUserFacingMessageable)()
    SetDeviceOwnerLockScreenMessage(value AndroidDeviceOwnerUserFacingMessageable)()
    SetEnrollmentProfile(value *AndroidDeviceOwnerEnrollmentProfileType)()
    SetFactoryResetBlocked(value *bool)()
    SetFactoryResetDeviceAdministratorEmails(value []string)()
    SetGlobalProxy(value AndroidDeviceOwnerGlobalProxyable)()
    SetGoogleAccountsBlocked(value *bool)()
    SetKioskCustomizationDeviceSettingsBlocked(value *bool)()
    SetKioskCustomizationPowerButtonActionsBlocked(value *bool)()
    SetKioskCustomizationStatusBar(value *AndroidDeviceOwnerKioskCustomizationStatusBar)()
    SetKioskCustomizationSystemErrorWarnings(value *bool)()
    SetKioskCustomizationSystemNavigation(value *AndroidDeviceOwnerKioskCustomizationSystemNavigation)()
    SetKioskModeAppOrderEnabled(value *bool)()
    SetKioskModeAppPositions(value []AndroidDeviceOwnerKioskModeAppPositionItemable)()
    SetKioskModeApps(value []AppListItemable)()
    SetKioskModeAppsInFolderOrderedByName(value *bool)()
    SetKioskModeBluetoothConfigurationEnabled(value *bool)()
    SetKioskModeDebugMenuEasyAccessEnabled(value *bool)()
    SetKioskModeExitCode(value *string)()
    SetKioskModeFlashlightConfigurationEnabled(value *bool)()
    SetKioskModeFolderIcon(value *AndroidDeviceOwnerKioskModeFolderIcon)()
    SetKioskModeGridHeight(value *int32)()
    SetKioskModeGridWidth(value *int32)()
    SetKioskModeIconSize(value *AndroidDeviceOwnerKioskModeIconSize)()
    SetKioskModeLockHomeScreen(value *bool)()
    SetKioskModeManagedFolders(value []AndroidDeviceOwnerKioskModeManagedFolderable)()
    SetKioskModeManagedHomeScreenAutoSignout(value *bool)()
    SetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds(value *int32)()
    SetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds(value *int32)()
    SetKioskModeManagedHomeScreenPinComplexity(value *KioskModeManagedHomeScreenPinComplexity)()
    SetKioskModeManagedHomeScreenPinRequired(value *bool)()
    SetKioskModeManagedHomeScreenPinRequiredToResume(value *bool)()
    SetKioskModeManagedHomeScreenSignInBackground(value *string)()
    SetKioskModeManagedHomeScreenSignInBrandingLogo(value *string)()
    SetKioskModeManagedHomeScreenSignInEnabled(value *bool)()
    SetKioskModeManagedSettingsEntryDisabled(value *bool)()
    SetKioskModeMediaVolumeConfigurationEnabled(value *bool)()
    SetKioskModeScreenOrientation(value *AndroidDeviceOwnerKioskModeScreenOrientation)()
    SetKioskModeScreenSaverConfigurationEnabled(value *bool)()
    SetKioskModeScreenSaverDetectMediaDisabled(value *bool)()
    SetKioskModeScreenSaverDisplayTimeInSeconds(value *int32)()
    SetKioskModeScreenSaverImageUrl(value *string)()
    SetKioskModeScreenSaverStartDelayInSeconds(value *int32)()
    SetKioskModeShowAppNotificationBadge(value *bool)()
    SetKioskModeShowDeviceInfo(value *bool)()
    SetKioskModeUseManagedHomeScreenApp(value *KioskModeType)()
    SetKioskModeVirtualHomeButtonEnabled(value *bool)()
    SetKioskModeVirtualHomeButtonType(value *AndroidDeviceOwnerVirtualHomeButtonType)()
    SetKioskModeWallpaperUrl(value *string)()
    SetKioskModeWifiAllowedSsids(value []string)()
    SetKioskModeWiFiConfigurationEnabled(value *bool)()
    SetLocateDeviceLostModeEnabled(value *bool)()
    SetLocateDeviceUserlessDisabled(value *bool)()
    SetMicrophoneForceMute(value *bool)()
    SetMicrosoftLauncherConfigurationEnabled(value *bool)()
    SetMicrosoftLauncherCustomWallpaperAllowUserModification(value *bool)()
    SetMicrosoftLauncherCustomWallpaperEnabled(value *bool)()
    SetMicrosoftLauncherCustomWallpaperImageUrl(value *string)()
    SetMicrosoftLauncherDockPresenceAllowUserModification(value *bool)()
    SetMicrosoftLauncherDockPresenceConfiguration(value *MicrosoftLauncherDockPresence)()
    SetMicrosoftLauncherFeedAllowUserModification(value *bool)()
    SetMicrosoftLauncherFeedEnabled(value *bool)()
    SetMicrosoftLauncherSearchBarPlacementConfiguration(value *MicrosoftLauncherSearchBarPlacement)()
    SetNetworkEscapeHatchAllowed(value *bool)()
    SetNfcBlockOutgoingBeam(value *bool)()
    SetPasswordBlockKeyguard(value *bool)()
    SetPasswordBlockKeyguardFeatures(value []AndroidKeyguardFeature)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinimumLetterCharacters(value *int32)()
    SetPasswordMinimumLowerCaseCharacters(value *int32)()
    SetPasswordMinimumNonLetterCharacters(value *int32)()
    SetPasswordMinimumNumericCharacters(value *int32)()
    SetPasswordMinimumSymbolCharacters(value *int32)()
    SetPasswordMinimumUpperCaseCharacters(value *int32)()
    SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)()
    SetPasswordPreviousPasswordCountToBlock(value *int32)()
    SetPasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)()
    SetPasswordRequireUnlock(value *AndroidDeviceOwnerRequiredPasswordUnlock)()
    SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)()
    SetPersonalProfileAppsAllowInstallFromUnknownSources(value *bool)()
    SetPersonalProfileCameraBlocked(value *bool)()
    SetPersonalProfilePersonalApplications(value []AppListItemable)()
    SetPersonalProfilePlayStoreMode(value *PersonalProfilePersonalPlayStoreMode)()
    SetPersonalProfileScreenCaptureBlocked(value *bool)()
    SetPlayStoreMode(value *AndroidDeviceOwnerPlayStoreMode)()
    SetScreenCaptureBlocked(value *bool)()
    SetSecurityCommonCriteriaModeEnabled(value *bool)()
    SetSecurityDeveloperSettingsEnabled(value *bool)()
    SetSecurityRequireVerifyApps(value *bool)()
    SetShortHelpText(value AndroidDeviceOwnerUserFacingMessageable)()
    SetStatusBarBlocked(value *bool)()
    SetStayOnModes(value []AndroidDeviceOwnerBatteryPluggedMode)()
    SetStorageAllowUsb(value *bool)()
    SetStorageBlockExternalMedia(value *bool)()
    SetStorageBlockUsbFileTransfer(value *bool)()
    SetSystemUpdateFreezePeriods(value []AndroidDeviceOwnerSystemUpdateFreezePeriodable)()
    SetSystemUpdateInstallType(value *AndroidDeviceOwnerSystemUpdateInstallType)()
    SetSystemUpdateWindowEndMinutesAfterMidnight(value *int32)()
    SetSystemUpdateWindowStartMinutesAfterMidnight(value *int32)()
    SetSystemWindowsBlocked(value *bool)()
    SetUsersBlockAdd(value *bool)()
    SetUsersBlockRemove(value *bool)()
    SetVolumeBlockAdjustment(value *bool)()
    SetVpnAlwaysOnLockdownMode(value *bool)()
    SetVpnAlwaysOnPackageIdentifier(value *string)()
    SetWifiBlockEditConfigurations(value *bool)()
    SetWifiBlockEditPolicyDefinedConfigurations(value *bool)()
    SetWorkProfilePasswordExpirationDays(value *int32)()
    SetWorkProfilePasswordMinimumLength(value *int32)()
    SetWorkProfilePasswordMinimumLetterCharacters(value *int32)()
    SetWorkProfilePasswordMinimumLowerCaseCharacters(value *int32)()
    SetWorkProfilePasswordMinimumNonLetterCharacters(value *int32)()
    SetWorkProfilePasswordMinimumNumericCharacters(value *int32)()
    SetWorkProfilePasswordMinimumSymbolCharacters(value *int32)()
    SetWorkProfilePasswordMinimumUpperCaseCharacters(value *int32)()
    SetWorkProfilePasswordPreviousPasswordCountToBlock(value *int32)()
    SetWorkProfilePasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)()
    SetWorkProfilePasswordRequireUnlock(value *AndroidDeviceOwnerRequiredPasswordUnlock)()
    SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(value *int32)()
}
