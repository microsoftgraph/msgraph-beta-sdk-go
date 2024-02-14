package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedAppProtection policy used to configure detailed management settings targeted to specific security groups and for a specified set of apps on an Android device
type AndroidManagedAppProtection struct {
    TargetedManagedAppProtection
}
// NewAndroidManagedAppProtection instantiates a new AndroidManagedAppProtection and sets the default values.
func NewAndroidManagedAppProtection()(*AndroidManagedAppProtection) {
    m := &AndroidManagedAppProtection{
        TargetedManagedAppProtection: *NewTargetedManagedAppProtection(),
    }
    odataTypeValue := "#microsoft.graph.androidManagedAppProtection"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidManagedAppProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedAppProtection(), nil
}
// GetAllowedAndroidDeviceManufacturers gets the allowedAndroidDeviceManufacturers property value. Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work.
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetAllowedAndroidDeviceManufacturers()(*string) {
    val, err := m.GetBackingStore().Get("allowedAndroidDeviceManufacturers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAllowedAndroidDeviceModels gets the allowedAndroidDeviceModels property value. List of device models allowed, as a string, for the managed app to work.
// returns a []string when successful
func (m *AndroidManagedAppProtection) GetAllowedAndroidDeviceModels()([]string) {
    val, err := m.GetBackingStore().Get("allowedAndroidDeviceModels")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetAppActionIfAccountIsClockedOut gets the appActionIfAccountIsClockedOut property value. Defines a managed app behavior, either block or warn, if the user is clocked out (non-working time).
// returns a *ManagedAppRemediationAction when successful
func (m *AndroidManagedAppProtection) GetAppActionIfAccountIsClockedOut()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfAccountIsClockedOut")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppActionIfAndroidDeviceManufacturerNotAllowed gets the appActionIfAndroidDeviceManufacturerNotAllowed property value. An admin initiated action to be applied on a managed app.
// returns a *ManagedAppRemediationAction when successful
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidDeviceManufacturerNotAllowed()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfAndroidDeviceManufacturerNotAllowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppActionIfAndroidDeviceModelNotAllowed gets the appActionIfAndroidDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
// returns a *ManagedAppRemediationAction when successful
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidDeviceModelNotAllowed()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfAndroidDeviceModelNotAllowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppActionIfAndroidSafetyNetAppsVerificationFailed gets the appActionIfAndroidSafetyNetAppsVerificationFailed property value. An admin initiated action to be applied on a managed app.
// returns a *ManagedAppRemediationAction when successful
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidSafetyNetAppsVerificationFailed()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfAndroidSafetyNetAppsVerificationFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppActionIfAndroidSafetyNetDeviceAttestationFailed gets the appActionIfAndroidSafetyNetDeviceAttestationFailed property value. An admin initiated action to be applied on a managed app.
// returns a *ManagedAppRemediationAction when successful
func (m *AndroidManagedAppProtection) GetAppActionIfAndroidSafetyNetDeviceAttestationFailed()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfAndroidSafetyNetDeviceAttestationFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppActionIfDeviceLockNotSet gets the appActionIfDeviceLockNotSet property value. An admin initiated action to be applied on a managed app.
// returns a *ManagedAppRemediationAction when successful
func (m *AndroidManagedAppProtection) GetAppActionIfDeviceLockNotSet()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfDeviceLockNotSet")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppActionIfDevicePasscodeComplexityLessThanHigh gets the appActionIfDevicePasscodeComplexityLessThanHigh property value. If the device does not have a passcode of high complexity or higher, trigger the stored action.
// returns a *ManagedAppRemediationAction when successful
func (m *AndroidManagedAppProtection) GetAppActionIfDevicePasscodeComplexityLessThanHigh()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfDevicePasscodeComplexityLessThanHigh")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppActionIfDevicePasscodeComplexityLessThanLow gets the appActionIfDevicePasscodeComplexityLessThanLow property value. If the device does not have a passcode of low complexity or higher, trigger the stored action.
// returns a *ManagedAppRemediationAction when successful
func (m *AndroidManagedAppProtection) GetAppActionIfDevicePasscodeComplexityLessThanLow()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfDevicePasscodeComplexityLessThanLow")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppActionIfDevicePasscodeComplexityLessThanMedium gets the appActionIfDevicePasscodeComplexityLessThanMedium property value. If the device does not have a passcode of medium complexity or higher, trigger the stored action.
// returns a *ManagedAppRemediationAction when successful
func (m *AndroidManagedAppProtection) GetAppActionIfDevicePasscodeComplexityLessThanMedium()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfDevicePasscodeComplexityLessThanMedium")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppActionIfSamsungKnoxAttestationRequired gets the appActionIfSamsungKnoxAttestationRequired property value. Defines the behavior of a managed app when Samsung Knox Attestation is required. Possible values are null, warn, block & wipe. If the admin does not set this action, the default is null, which indicates this setting is not configured.
// returns a *ManagedAppRemediationAction when successful
func (m *AndroidManagedAppProtection) GetAppActionIfSamsungKnoxAttestationRequired()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfSamsungKnoxAttestationRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetApprovedKeyboards gets the approvedKeyboards property value. If Keyboard Restriction is enabled, only keyboards in this approved list will be allowed. A key should be Android package id for a keyboard and value should be a friendly name
// returns a []KeyValuePairable when successful
func (m *AndroidManagedAppProtection) GetApprovedKeyboards()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("approvedKeyboards")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
// returns a []ManagedMobileAppable when successful
func (m *AndroidManagedAppProtection) GetApps()([]ManagedMobileAppable) {
    val, err := m.GetBackingStore().Get("apps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedMobileAppable)
    }
    return nil
}
// GetBiometricAuthenticationBlocked gets the biometricAuthenticationBlocked property value. Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True.
// returns a *bool when successful
func (m *AndroidManagedAppProtection) GetBiometricAuthenticationBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("biometricAuthenticationBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBlockAfterCompanyPortalUpdateDeferralInDays gets the blockAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
// returns a *int32 when successful
func (m *AndroidManagedAppProtection) GetBlockAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    val, err := m.GetBackingStore().Get("blockAfterCompanyPortalUpdateDeferralInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetConnectToVpnOnLaunch gets the connectToVpnOnLaunch property value. Whether the app should connect to the configured VPN on launch.
// returns a *bool when successful
func (m *AndroidManagedAppProtection) GetConnectToVpnOnLaunch()(*bool) {
    val, err := m.GetBackingStore().Get("connectToVpnOnLaunch")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCustomBrowserDisplayName gets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android.
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetCustomBrowserDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("customBrowserDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomBrowserPackageId gets the customBrowserPackageId property value. Unique identifier of a custom browser to open weblink on Android.
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetCustomBrowserPackageId()(*string) {
    val, err := m.GetBackingStore().Get("customBrowserPackageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomDialerAppDisplayName gets the customDialerAppDisplayName property value. Friendly name of a custom dialer app to click-to-open a phone number on Android.
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetCustomDialerAppDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("customDialerAppDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomDialerAppPackageId gets the customDialerAppPackageId property value. PackageId of a custom dialer app to click-to-open a phone number on Android.
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetCustomDialerAppPackageId()(*string) {
    val, err := m.GetBackingStore().Get("customDialerAppPackageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeployedAppCount gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
// returns a *int32 when successful
func (m *AndroidManagedAppProtection) GetDeployedAppCount()(*int32) {
    val, err := m.GetBackingStore().Get("deployedAppCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDeploymentSummary gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
// returns a ManagedAppPolicyDeploymentSummaryable when successful
func (m *AndroidManagedAppProtection) GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable) {
    val, err := m.GetBackingStore().Get("deploymentSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagedAppPolicyDeploymentSummaryable)
    }
    return nil
}
// GetDeviceLockRequired gets the deviceLockRequired property value. Defines if any kind of lock must be required on android device
// returns a *bool when successful
func (m *AndroidManagedAppProtection) GetDeviceLockRequired()(*bool) {
    val, err := m.GetBackingStore().Get("deviceLockRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDisableAppEncryptionIfDeviceEncryptionIsEnabled gets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled
// returns a *bool when successful
func (m *AndroidManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("disableAppEncryptionIfDeviceEncryptionIsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEncryptAppData gets the encryptAppData property value. Indicates whether application data for managed apps should be encrypted
// returns a *bool when successful
func (m *AndroidManagedAppProtection) GetEncryptAppData()(*bool) {
    val, err := m.GetBackingStore().Get("encryptAppData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetExemptedAppPackages gets the exemptedAppPackages property value. App packages in this list will be exempt from the policy and will be able to receive data from managed apps.
// returns a []KeyValuePairable when successful
func (m *AndroidManagedAppProtection) GetExemptedAppPackages()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("exemptedAppPackages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AndroidManagedAppProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TargetedManagedAppProtection.GetFieldDeserializers()
    res["allowedAndroidDeviceManufacturers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedAndroidDeviceManufacturers(val)
        }
        return nil
    }
    res["allowedAndroidDeviceModels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAllowedAndroidDeviceModels(res)
        }
        return nil
    }
    res["appActionIfAccountIsClockedOut"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAccountIsClockedOut(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfAndroidDeviceManufacturerNotAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAndroidDeviceManufacturerNotAllowed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfAndroidDeviceModelNotAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAndroidDeviceModelNotAllowed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfAndroidSafetyNetAppsVerificationFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAndroidSafetyNetAppsVerificationFailed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfAndroidSafetyNetDeviceAttestationFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfDeviceLockNotSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfDeviceLockNotSet(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfDevicePasscodeComplexityLessThanHigh"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfDevicePasscodeComplexityLessThanHigh(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfDevicePasscodeComplexityLessThanLow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfDevicePasscodeComplexityLessThanLow(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfDevicePasscodeComplexityLessThanMedium"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfDevicePasscodeComplexityLessThanMedium(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfSamsungKnoxAttestationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfSamsungKnoxAttestationRequired(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["approvedKeyboards"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyValuePairable)
                }
            }
            m.SetApprovedKeyboards(res)
        }
        return nil
    }
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedMobileAppable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedMobileAppable)
                }
            }
            m.SetApps(res)
        }
        return nil
    }
    res["biometricAuthenticationBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBiometricAuthenticationBlocked(val)
        }
        return nil
    }
    res["blockAfterCompanyPortalUpdateDeferralInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockAfterCompanyPortalUpdateDeferralInDays(val)
        }
        return nil
    }
    res["connectToVpnOnLaunch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectToVpnOnLaunch(val)
        }
        return nil
    }
    res["customBrowserDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBrowserDisplayName(val)
        }
        return nil
    }
    res["customBrowserPackageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBrowserPackageId(val)
        }
        return nil
    }
    res["customDialerAppDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomDialerAppDisplayName(val)
        }
        return nil
    }
    res["customDialerAppPackageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomDialerAppPackageId(val)
        }
        return nil
    }
    res["deployedAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedAppCount(val)
        }
        return nil
    }
    res["deploymentSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentSummary(val.(ManagedAppPolicyDeploymentSummaryable))
        }
        return nil
    }
    res["deviceLockRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLockRequired(val)
        }
        return nil
    }
    res["disableAppEncryptionIfDeviceEncryptionIsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(val)
        }
        return nil
    }
    res["encryptAppData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptAppData(val)
        }
        return nil
    }
    res["exemptedAppPackages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyValuePairable)
                }
            }
            m.SetExemptedAppPackages(res)
        }
        return nil
    }
    res["fingerprintAndBiometricEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprintAndBiometricEnabled(val)
        }
        return nil
    }
    res["keyboardsRestricted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardsRestricted(val)
        }
        return nil
    }
    res["messagingRedirectAppDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessagingRedirectAppDisplayName(val)
        }
        return nil
    }
    res["messagingRedirectAppPackageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessagingRedirectAppPackageId(val)
        }
        return nil
    }
    res["minimumRequiredCompanyPortalVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredCompanyPortalVersion(val)
        }
        return nil
    }
    res["minimumRequiredPatchVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredPatchVersion(val)
        }
        return nil
    }
    res["minimumWarningCompanyPortalVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningCompanyPortalVersion(val)
        }
        return nil
    }
    res["minimumWarningPatchVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningPatchVersion(val)
        }
        return nil
    }
    res["minimumWipeCompanyPortalVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipeCompanyPortalVersion(val)
        }
        return nil
    }
    res["minimumWipePatchVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipePatchVersion(val)
        }
        return nil
    }
    res["requireClass3Biometrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireClass3Biometrics(val)
        }
        return nil
    }
    res["requiredAndroidSafetyNetAppsVerificationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedAppSafetyNetAppsVerificationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredAndroidSafetyNetAppsVerificationType(val.(*AndroidManagedAppSafetyNetAppsVerificationType))
        }
        return nil
    }
    res["requiredAndroidSafetyNetDeviceAttestationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedAppSafetyNetDeviceAttestationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredAndroidSafetyNetDeviceAttestationType(val.(*AndroidManagedAppSafetyNetDeviceAttestationType))
        }
        return nil
    }
    res["requiredAndroidSafetyNetEvaluationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidManagedAppSafetyNetEvaluationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredAndroidSafetyNetEvaluationType(val.(*AndroidManagedAppSafetyNetEvaluationType))
        }
        return nil
    }
    res["requirePinAfterBiometricChange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequirePinAfterBiometricChange(val)
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
    res["warnAfterCompanyPortalUpdateDeferralInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWarnAfterCompanyPortalUpdateDeferralInDays(val)
        }
        return nil
    }
    res["wipeAfterCompanyPortalUpdateDeferralInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWipeAfterCompanyPortalUpdateDeferralInDays(val)
        }
        return nil
    }
    return res
}
// GetFingerprintAndBiometricEnabled gets the fingerprintAndBiometricEnabled property value. If null, this setting will be ignored. If false both fingerprints and biometrics will not be enabled. If true, both fingerprints and biometrics will be enabled.
// returns a *bool when successful
func (m *AndroidManagedAppProtection) GetFingerprintAndBiometricEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("fingerprintAndBiometricEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKeyboardsRestricted gets the keyboardsRestricted property value. Indicates if keyboard restriction is enabled. If enabled list of approved keyboards must be provided as well.
// returns a *bool when successful
func (m *AndroidManagedAppProtection) GetKeyboardsRestricted()(*bool) {
    val, err := m.GetBackingStore().Get("keyboardsRestricted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMessagingRedirectAppDisplayName gets the messagingRedirectAppDisplayName property value. When a specific app redirection is enforced by protectedMessagingRedirectAppType in an App Protection Policy, this value defines the app name which is allowed to be used.
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetMessagingRedirectAppDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("messagingRedirectAppDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMessagingRedirectAppPackageId gets the messagingRedirectAppPackageId property value. When a specific app redirection is enforced by protectedMessagingRedirectAppType in an App Protection Policy, this value defines the app package id which is allowed to be used.
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetMessagingRedirectAppPackageId()(*string) {
    val, err := m.GetBackingStore().Get("messagingRedirectAppPackageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumRequiredCompanyPortalVersion gets the minimumRequiredCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or app access will be blocked
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetMinimumRequiredCompanyPortalVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumRequiredCompanyPortalVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumRequiredPatchVersion gets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app.
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetMinimumRequiredPatchVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumRequiredPatchVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWarningCompanyPortalVersion gets the minimumWarningCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the user will receive a warning
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetMinimumWarningCompanyPortalVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWarningCompanyPortalVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWarningPatchVersion gets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app.
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetMinimumWarningPatchVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWarningPatchVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWipeCompanyPortalVersion gets the minimumWipeCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetMinimumWipeCompanyPortalVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWipeCompanyPortalVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWipePatchVersion gets the minimumWipePatchVersion property value. Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data.
// returns a *string when successful
func (m *AndroidManagedAppProtection) GetMinimumWipePatchVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWipePatchVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequireClass3Biometrics gets the requireClass3Biometrics property value. Require user to apply Class 3 Biometrics on their Android device.
// returns a *bool when successful
func (m *AndroidManagedAppProtection) GetRequireClass3Biometrics()(*bool) {
    val, err := m.GetBackingStore().Get("requireClass3Biometrics")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRequiredAndroidSafetyNetAppsVerificationType gets the requiredAndroidSafetyNetAppsVerificationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
// returns a *AndroidManagedAppSafetyNetAppsVerificationType when successful
func (m *AndroidManagedAppProtection) GetRequiredAndroidSafetyNetAppsVerificationType()(*AndroidManagedAppSafetyNetAppsVerificationType) {
    val, err := m.GetBackingStore().Get("requiredAndroidSafetyNetAppsVerificationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidManagedAppSafetyNetAppsVerificationType)
    }
    return nil
}
// GetRequiredAndroidSafetyNetDeviceAttestationType gets the requiredAndroidSafetyNetDeviceAttestationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
// returns a *AndroidManagedAppSafetyNetDeviceAttestationType when successful
func (m *AndroidManagedAppProtection) GetRequiredAndroidSafetyNetDeviceAttestationType()(*AndroidManagedAppSafetyNetDeviceAttestationType) {
    val, err := m.GetBackingStore().Get("requiredAndroidSafetyNetDeviceAttestationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidManagedAppSafetyNetDeviceAttestationType)
    }
    return nil
}
// GetRequiredAndroidSafetyNetEvaluationType gets the requiredAndroidSafetyNetEvaluationType property value. An admin enforced Android SafetyNet evaluation type requirement on a managed app.
// returns a *AndroidManagedAppSafetyNetEvaluationType when successful
func (m *AndroidManagedAppProtection) GetRequiredAndroidSafetyNetEvaluationType()(*AndroidManagedAppSafetyNetEvaluationType) {
    val, err := m.GetBackingStore().Get("requiredAndroidSafetyNetEvaluationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidManagedAppSafetyNetEvaluationType)
    }
    return nil
}
// GetRequirePinAfterBiometricChange gets the requirePinAfterBiometricChange property value. A PIN prompt will override biometric prompts if class 3 biometrics are updated on the device.
// returns a *bool when successful
func (m *AndroidManagedAppProtection) GetRequirePinAfterBiometricChange()(*bool) {
    val, err := m.GetBackingStore().Get("requirePinAfterBiometricChange")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether a managed user can take screen captures of managed apps
// returns a *bool when successful
func (m *AndroidManagedAppProtection) GetScreenCaptureBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("screenCaptureBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWarnAfterCompanyPortalUpdateDeferralInDays gets the warnAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
// returns a *int32 when successful
func (m *AndroidManagedAppProtection) GetWarnAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    val, err := m.GetBackingStore().Get("warnAfterCompanyPortalUpdateDeferralInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWipeAfterCompanyPortalUpdateDeferralInDays gets the wipeAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
// returns a *int32 when successful
func (m *AndroidManagedAppProtection) GetWipeAfterCompanyPortalUpdateDeferralInDays()(*int32) {
    val, err := m.GetBackingStore().Get("wipeAfterCompanyPortalUpdateDeferralInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidManagedAppProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TargetedManagedAppProtection.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("allowedAndroidDeviceManufacturers", m.GetAllowedAndroidDeviceManufacturers())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedAndroidDeviceModels() != nil {
        err = writer.WriteCollectionOfStringValues("allowedAndroidDeviceModels", m.GetAllowedAndroidDeviceModels())
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAccountIsClockedOut() != nil {
        cast := (*m.GetAppActionIfAccountIsClockedOut()).String()
        err = writer.WriteStringValue("appActionIfAccountIsClockedOut", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAndroidDeviceManufacturerNotAllowed() != nil {
        cast := (*m.GetAppActionIfAndroidDeviceManufacturerNotAllowed()).String()
        err = writer.WriteStringValue("appActionIfAndroidDeviceManufacturerNotAllowed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAndroidDeviceModelNotAllowed() != nil {
        cast := (*m.GetAppActionIfAndroidDeviceModelNotAllowed()).String()
        err = writer.WriteStringValue("appActionIfAndroidDeviceModelNotAllowed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAndroidSafetyNetAppsVerificationFailed() != nil {
        cast := (*m.GetAppActionIfAndroidSafetyNetAppsVerificationFailed()).String()
        err = writer.WriteStringValue("appActionIfAndroidSafetyNetAppsVerificationFailed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfAndroidSafetyNetDeviceAttestationFailed() != nil {
        cast := (*m.GetAppActionIfAndroidSafetyNetDeviceAttestationFailed()).String()
        err = writer.WriteStringValue("appActionIfAndroidSafetyNetDeviceAttestationFailed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfDeviceLockNotSet() != nil {
        cast := (*m.GetAppActionIfDeviceLockNotSet()).String()
        err = writer.WriteStringValue("appActionIfDeviceLockNotSet", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfDevicePasscodeComplexityLessThanHigh() != nil {
        cast := (*m.GetAppActionIfDevicePasscodeComplexityLessThanHigh()).String()
        err = writer.WriteStringValue("appActionIfDevicePasscodeComplexityLessThanHigh", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfDevicePasscodeComplexityLessThanLow() != nil {
        cast := (*m.GetAppActionIfDevicePasscodeComplexityLessThanLow()).String()
        err = writer.WriteStringValue("appActionIfDevicePasscodeComplexityLessThanLow", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfDevicePasscodeComplexityLessThanMedium() != nil {
        cast := (*m.GetAppActionIfDevicePasscodeComplexityLessThanMedium()).String()
        err = writer.WriteStringValue("appActionIfDevicePasscodeComplexityLessThanMedium", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfSamsungKnoxAttestationRequired() != nil {
        cast := (*m.GetAppActionIfSamsungKnoxAttestationRequired()).String()
        err = writer.WriteStringValue("appActionIfSamsungKnoxAttestationRequired", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApprovedKeyboards() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApprovedKeyboards()))
        for i, v := range m.GetApprovedKeyboards() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("approvedKeyboards", cast)
        if err != nil {
            return err
        }
    }
    if m.GetApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("apps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("biometricAuthenticationBlocked", m.GetBiometricAuthenticationBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("blockAfterCompanyPortalUpdateDeferralInDays", m.GetBlockAfterCompanyPortalUpdateDeferralInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("connectToVpnOnLaunch", m.GetConnectToVpnOnLaunch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customBrowserDisplayName", m.GetCustomBrowserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customBrowserPackageId", m.GetCustomBrowserPackageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customDialerAppDisplayName", m.GetCustomDialerAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customDialerAppPackageId", m.GetCustomDialerAppPackageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deployedAppCount", m.GetDeployedAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentSummary", m.GetDeploymentSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceLockRequired", m.GetDeviceLockRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableAppEncryptionIfDeviceEncryptionIsEnabled", m.GetDisableAppEncryptionIfDeviceEncryptionIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("encryptAppData", m.GetEncryptAppData())
        if err != nil {
            return err
        }
    }
    if m.GetExemptedAppPackages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExemptedAppPackages()))
        for i, v := range m.GetExemptedAppPackages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("exemptedAppPackages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fingerprintAndBiometricEnabled", m.GetFingerprintAndBiometricEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyboardsRestricted", m.GetKeyboardsRestricted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("messagingRedirectAppDisplayName", m.GetMessagingRedirectAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("messagingRedirectAppPackageId", m.GetMessagingRedirectAppPackageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredCompanyPortalVersion", m.GetMinimumRequiredCompanyPortalVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredPatchVersion", m.GetMinimumRequiredPatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWarningCompanyPortalVersion", m.GetMinimumWarningCompanyPortalVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWarningPatchVersion", m.GetMinimumWarningPatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWipeCompanyPortalVersion", m.GetMinimumWipeCompanyPortalVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWipePatchVersion", m.GetMinimumWipePatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireClass3Biometrics", m.GetRequireClass3Biometrics())
        if err != nil {
            return err
        }
    }
    if m.GetRequiredAndroidSafetyNetAppsVerificationType() != nil {
        cast := (*m.GetRequiredAndroidSafetyNetAppsVerificationType()).String()
        err = writer.WriteStringValue("requiredAndroidSafetyNetAppsVerificationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequiredAndroidSafetyNetDeviceAttestationType() != nil {
        cast := (*m.GetRequiredAndroidSafetyNetDeviceAttestationType()).String()
        err = writer.WriteStringValue("requiredAndroidSafetyNetDeviceAttestationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequiredAndroidSafetyNetEvaluationType() != nil {
        cast := (*m.GetRequiredAndroidSafetyNetEvaluationType()).String()
        err = writer.WriteStringValue("requiredAndroidSafetyNetEvaluationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requirePinAfterBiometricChange", m.GetRequirePinAfterBiometricChange())
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
        err = writer.WriteInt32Value("warnAfterCompanyPortalUpdateDeferralInDays", m.GetWarnAfterCompanyPortalUpdateDeferralInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("wipeAfterCompanyPortalUpdateDeferralInDays", m.GetWipeAfterCompanyPortalUpdateDeferralInDays())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedAndroidDeviceManufacturers sets the allowedAndroidDeviceManufacturers property value. Semicolon seperated list of device manufacturers allowed, as a string, for the managed app to work.
func (m *AndroidManagedAppProtection) SetAllowedAndroidDeviceManufacturers(value *string)() {
    err := m.GetBackingStore().Set("allowedAndroidDeviceManufacturers", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedAndroidDeviceModels sets the allowedAndroidDeviceModels property value. List of device models allowed, as a string, for the managed app to work.
func (m *AndroidManagedAppProtection) SetAllowedAndroidDeviceModels(value []string)() {
    err := m.GetBackingStore().Set("allowedAndroidDeviceModels", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfAccountIsClockedOut sets the appActionIfAccountIsClockedOut property value. Defines a managed app behavior, either block or warn, if the user is clocked out (non-working time).
func (m *AndroidManagedAppProtection) SetAppActionIfAccountIsClockedOut(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfAccountIsClockedOut", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfAndroidDeviceManufacturerNotAllowed sets the appActionIfAndroidDeviceManufacturerNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidDeviceManufacturerNotAllowed(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfAndroidDeviceManufacturerNotAllowed", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfAndroidDeviceModelNotAllowed sets the appActionIfAndroidDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidDeviceModelNotAllowed(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfAndroidDeviceModelNotAllowed", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfAndroidSafetyNetAppsVerificationFailed sets the appActionIfAndroidSafetyNetAppsVerificationFailed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidSafetyNetAppsVerificationFailed(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfAndroidSafetyNetAppsVerificationFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfAndroidSafetyNetDeviceAttestationFailed sets the appActionIfAndroidSafetyNetDeviceAttestationFailed property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfAndroidSafetyNetDeviceAttestationFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfDeviceLockNotSet sets the appActionIfDeviceLockNotSet property value. An admin initiated action to be applied on a managed app.
func (m *AndroidManagedAppProtection) SetAppActionIfDeviceLockNotSet(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfDeviceLockNotSet", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfDevicePasscodeComplexityLessThanHigh sets the appActionIfDevicePasscodeComplexityLessThanHigh property value. If the device does not have a passcode of high complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) SetAppActionIfDevicePasscodeComplexityLessThanHigh(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfDevicePasscodeComplexityLessThanHigh", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfDevicePasscodeComplexityLessThanLow sets the appActionIfDevicePasscodeComplexityLessThanLow property value. If the device does not have a passcode of low complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) SetAppActionIfDevicePasscodeComplexityLessThanLow(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfDevicePasscodeComplexityLessThanLow", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfDevicePasscodeComplexityLessThanMedium sets the appActionIfDevicePasscodeComplexityLessThanMedium property value. If the device does not have a passcode of medium complexity or higher, trigger the stored action.
func (m *AndroidManagedAppProtection) SetAppActionIfDevicePasscodeComplexityLessThanMedium(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfDevicePasscodeComplexityLessThanMedium", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfSamsungKnoxAttestationRequired sets the appActionIfSamsungKnoxAttestationRequired property value. Defines the behavior of a managed app when Samsung Knox Attestation is required. Possible values are null, warn, block & wipe. If the admin does not set this action, the default is null, which indicates this setting is not configured.
func (m *AndroidManagedAppProtection) SetAppActionIfSamsungKnoxAttestationRequired(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfSamsungKnoxAttestationRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetApprovedKeyboards sets the approvedKeyboards property value. If Keyboard Restriction is enabled, only keyboards in this approved list will be allowed. A key should be Android package id for a keyboard and value should be a friendly name
func (m *AndroidManagedAppProtection) SetApprovedKeyboards(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("approvedKeyboards", value)
    if err != nil {
        panic(err)
    }
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *AndroidManagedAppProtection) SetApps(value []ManagedMobileAppable)() {
    err := m.GetBackingStore().Set("apps", value)
    if err != nil {
        panic(err)
    }
}
// SetBiometricAuthenticationBlocked sets the biometricAuthenticationBlocked property value. Indicates whether use of the biometric authentication is allowed in place of a pin if PinRequired is set to True.
func (m *AndroidManagedAppProtection) SetBiometricAuthenticationBlocked(value *bool)() {
    err := m.GetBackingStore().Set("biometricAuthenticationBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockAfterCompanyPortalUpdateDeferralInDays sets the blockAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or app access will be blocked.
func (m *AndroidManagedAppProtection) SetBlockAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    err := m.GetBackingStore().Set("blockAfterCompanyPortalUpdateDeferralInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectToVpnOnLaunch sets the connectToVpnOnLaunch property value. Whether the app should connect to the configured VPN on launch.
func (m *AndroidManagedAppProtection) SetConnectToVpnOnLaunch(value *bool)() {
    err := m.GetBackingStore().Set("connectToVpnOnLaunch", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomBrowserDisplayName sets the customBrowserDisplayName property value. Friendly name of the preferred custom browser to open weblink on Android.
func (m *AndroidManagedAppProtection) SetCustomBrowserDisplayName(value *string)() {
    err := m.GetBackingStore().Set("customBrowserDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomBrowserPackageId sets the customBrowserPackageId property value. Unique identifier of a custom browser to open weblink on Android.
func (m *AndroidManagedAppProtection) SetCustomBrowserPackageId(value *string)() {
    err := m.GetBackingStore().Set("customBrowserPackageId", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomDialerAppDisplayName sets the customDialerAppDisplayName property value. Friendly name of a custom dialer app to click-to-open a phone number on Android.
func (m *AndroidManagedAppProtection) SetCustomDialerAppDisplayName(value *string)() {
    err := m.GetBackingStore().Set("customDialerAppDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomDialerAppPackageId sets the customDialerAppPackageId property value. PackageId of a custom dialer app to click-to-open a phone number on Android.
func (m *AndroidManagedAppProtection) SetCustomDialerAppPackageId(value *string)() {
    err := m.GetBackingStore().Set("customDialerAppPackageId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeployedAppCount sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *AndroidManagedAppProtection) SetDeployedAppCount(value *int32)() {
    err := m.GetBackingStore().Set("deployedAppCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentSummary sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *AndroidManagedAppProtection) SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)() {
    err := m.GetBackingStore().Set("deploymentSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceLockRequired sets the deviceLockRequired property value. Defines if any kind of lock must be required on android device
func (m *AndroidManagedAppProtection) SetDeviceLockRequired(value *bool)() {
    err := m.GetBackingStore().Set("deviceLockRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableAppEncryptionIfDeviceEncryptionIsEnabled sets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled
func (m *AndroidManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("disableAppEncryptionIfDeviceEncryptionIsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetEncryptAppData sets the encryptAppData property value. Indicates whether application data for managed apps should be encrypted
func (m *AndroidManagedAppProtection) SetEncryptAppData(value *bool)() {
    err := m.GetBackingStore().Set("encryptAppData", value)
    if err != nil {
        panic(err)
    }
}
// SetExemptedAppPackages sets the exemptedAppPackages property value. App packages in this list will be exempt from the policy and will be able to receive data from managed apps.
func (m *AndroidManagedAppProtection) SetExemptedAppPackages(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("exemptedAppPackages", value)
    if err != nil {
        panic(err)
    }
}
// SetFingerprintAndBiometricEnabled sets the fingerprintAndBiometricEnabled property value. If null, this setting will be ignored. If false both fingerprints and biometrics will not be enabled. If true, both fingerprints and biometrics will be enabled.
func (m *AndroidManagedAppProtection) SetFingerprintAndBiometricEnabled(value *bool)() {
    err := m.GetBackingStore().Set("fingerprintAndBiometricEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyboardsRestricted sets the keyboardsRestricted property value. Indicates if keyboard restriction is enabled. If enabled list of approved keyboards must be provided as well.
func (m *AndroidManagedAppProtection) SetKeyboardsRestricted(value *bool)() {
    err := m.GetBackingStore().Set("keyboardsRestricted", value)
    if err != nil {
        panic(err)
    }
}
// SetMessagingRedirectAppDisplayName sets the messagingRedirectAppDisplayName property value. When a specific app redirection is enforced by protectedMessagingRedirectAppType in an App Protection Policy, this value defines the app name which is allowed to be used.
func (m *AndroidManagedAppProtection) SetMessagingRedirectAppDisplayName(value *string)() {
    err := m.GetBackingStore().Set("messagingRedirectAppDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetMessagingRedirectAppPackageId sets the messagingRedirectAppPackageId property value. When a specific app redirection is enforced by protectedMessagingRedirectAppType in an App Protection Policy, this value defines the app package id which is allowed to be used.
func (m *AndroidManagedAppProtection) SetMessagingRedirectAppPackageId(value *string)() {
    err := m.GetBackingStore().Set("messagingRedirectAppPackageId", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumRequiredCompanyPortalVersion sets the minimumRequiredCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or app access will be blocked
func (m *AndroidManagedAppProtection) SetMinimumRequiredCompanyPortalVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumRequiredCompanyPortalVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumRequiredPatchVersion sets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app.
func (m *AndroidManagedAppProtection) SetMinimumRequiredPatchVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumRequiredPatchVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWarningCompanyPortalVersion sets the minimumWarningCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the user will receive a warning
func (m *AndroidManagedAppProtection) SetMinimumWarningCompanyPortalVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWarningCompanyPortalVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWarningPatchVersion sets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app.
func (m *AndroidManagedAppProtection) SetMinimumWarningPatchVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWarningPatchVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWipeCompanyPortalVersion sets the minimumWipeCompanyPortalVersion property value. Minimum version of the Company portal that must be installed on the device or the company data on the app will be wiped
func (m *AndroidManagedAppProtection) SetMinimumWipeCompanyPortalVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWipeCompanyPortalVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWipePatchVersion sets the minimumWipePatchVersion property value. Android security patch level  less than or equal to the specified value will wipe the managed app and the associated company data.
func (m *AndroidManagedAppProtection) SetMinimumWipePatchVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWipePatchVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireClass3Biometrics sets the requireClass3Biometrics property value. Require user to apply Class 3 Biometrics on their Android device.
func (m *AndroidManagedAppProtection) SetRequireClass3Biometrics(value *bool)() {
    err := m.GetBackingStore().Set("requireClass3Biometrics", value)
    if err != nil {
        panic(err)
    }
}
// SetRequiredAndroidSafetyNetAppsVerificationType sets the requiredAndroidSafetyNetAppsVerificationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *AndroidManagedAppProtection) SetRequiredAndroidSafetyNetAppsVerificationType(value *AndroidManagedAppSafetyNetAppsVerificationType)() {
    err := m.GetBackingStore().Set("requiredAndroidSafetyNetAppsVerificationType", value)
    if err != nil {
        panic(err)
    }
}
// SetRequiredAndroidSafetyNetDeviceAttestationType sets the requiredAndroidSafetyNetDeviceAttestationType property value. An admin enforced Android SafetyNet Device Attestation requirement on a managed app.
func (m *AndroidManagedAppProtection) SetRequiredAndroidSafetyNetDeviceAttestationType(value *AndroidManagedAppSafetyNetDeviceAttestationType)() {
    err := m.GetBackingStore().Set("requiredAndroidSafetyNetDeviceAttestationType", value)
    if err != nil {
        panic(err)
    }
}
// SetRequiredAndroidSafetyNetEvaluationType sets the requiredAndroidSafetyNetEvaluationType property value. An admin enforced Android SafetyNet evaluation type requirement on a managed app.
func (m *AndroidManagedAppProtection) SetRequiredAndroidSafetyNetEvaluationType(value *AndroidManagedAppSafetyNetEvaluationType)() {
    err := m.GetBackingStore().Set("requiredAndroidSafetyNetEvaluationType", value)
    if err != nil {
        panic(err)
    }
}
// SetRequirePinAfterBiometricChange sets the requirePinAfterBiometricChange property value. A PIN prompt will override biometric prompts if class 3 biometrics are updated on the device.
func (m *AndroidManagedAppProtection) SetRequirePinAfterBiometricChange(value *bool)() {
    err := m.GetBackingStore().Set("requirePinAfterBiometricChange", value)
    if err != nil {
        panic(err)
    }
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether a managed user can take screen captures of managed apps
func (m *AndroidManagedAppProtection) SetScreenCaptureBlocked(value *bool)() {
    err := m.GetBackingStore().Set("screenCaptureBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetWarnAfterCompanyPortalUpdateDeferralInDays sets the warnAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the user will receive the warning
func (m *AndroidManagedAppProtection) SetWarnAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    err := m.GetBackingStore().Set("warnAfterCompanyPortalUpdateDeferralInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetWipeAfterCompanyPortalUpdateDeferralInDays sets the wipeAfterCompanyPortalUpdateDeferralInDays property value. Maximum number of days Company Portal update can be deferred on the device or the company data on the app will be wiped
func (m *AndroidManagedAppProtection) SetWipeAfterCompanyPortalUpdateDeferralInDays(value *int32)() {
    err := m.GetBackingStore().Set("wipeAfterCompanyPortalUpdateDeferralInDays", value)
    if err != nil {
        panic(err)
    }
}
type AndroidManagedAppProtectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TargetedManagedAppProtectionable
    GetAllowedAndroidDeviceManufacturers()(*string)
    GetAllowedAndroidDeviceModels()([]string)
    GetAppActionIfAccountIsClockedOut()(*ManagedAppRemediationAction)
    GetAppActionIfAndroidDeviceManufacturerNotAllowed()(*ManagedAppRemediationAction)
    GetAppActionIfAndroidDeviceModelNotAllowed()(*ManagedAppRemediationAction)
    GetAppActionIfAndroidSafetyNetAppsVerificationFailed()(*ManagedAppRemediationAction)
    GetAppActionIfAndroidSafetyNetDeviceAttestationFailed()(*ManagedAppRemediationAction)
    GetAppActionIfDeviceLockNotSet()(*ManagedAppRemediationAction)
    GetAppActionIfDevicePasscodeComplexityLessThanHigh()(*ManagedAppRemediationAction)
    GetAppActionIfDevicePasscodeComplexityLessThanLow()(*ManagedAppRemediationAction)
    GetAppActionIfDevicePasscodeComplexityLessThanMedium()(*ManagedAppRemediationAction)
    GetAppActionIfSamsungKnoxAttestationRequired()(*ManagedAppRemediationAction)
    GetApprovedKeyboards()([]KeyValuePairable)
    GetApps()([]ManagedMobileAppable)
    GetBiometricAuthenticationBlocked()(*bool)
    GetBlockAfterCompanyPortalUpdateDeferralInDays()(*int32)
    GetConnectToVpnOnLaunch()(*bool)
    GetCustomBrowserDisplayName()(*string)
    GetCustomBrowserPackageId()(*string)
    GetCustomDialerAppDisplayName()(*string)
    GetCustomDialerAppPackageId()(*string)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetDeviceLockRequired()(*bool)
    GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool)
    GetEncryptAppData()(*bool)
    GetExemptedAppPackages()([]KeyValuePairable)
    GetFingerprintAndBiometricEnabled()(*bool)
    GetKeyboardsRestricted()(*bool)
    GetMessagingRedirectAppDisplayName()(*string)
    GetMessagingRedirectAppPackageId()(*string)
    GetMinimumRequiredCompanyPortalVersion()(*string)
    GetMinimumRequiredPatchVersion()(*string)
    GetMinimumWarningCompanyPortalVersion()(*string)
    GetMinimumWarningPatchVersion()(*string)
    GetMinimumWipeCompanyPortalVersion()(*string)
    GetMinimumWipePatchVersion()(*string)
    GetRequireClass3Biometrics()(*bool)
    GetRequiredAndroidSafetyNetAppsVerificationType()(*AndroidManagedAppSafetyNetAppsVerificationType)
    GetRequiredAndroidSafetyNetDeviceAttestationType()(*AndroidManagedAppSafetyNetDeviceAttestationType)
    GetRequiredAndroidSafetyNetEvaluationType()(*AndroidManagedAppSafetyNetEvaluationType)
    GetRequirePinAfterBiometricChange()(*bool)
    GetScreenCaptureBlocked()(*bool)
    GetWarnAfterCompanyPortalUpdateDeferralInDays()(*int32)
    GetWipeAfterCompanyPortalUpdateDeferralInDays()(*int32)
    SetAllowedAndroidDeviceManufacturers(value *string)()
    SetAllowedAndroidDeviceModels(value []string)()
    SetAppActionIfAccountIsClockedOut(value *ManagedAppRemediationAction)()
    SetAppActionIfAndroidDeviceManufacturerNotAllowed(value *ManagedAppRemediationAction)()
    SetAppActionIfAndroidDeviceModelNotAllowed(value *ManagedAppRemediationAction)()
    SetAppActionIfAndroidSafetyNetAppsVerificationFailed(value *ManagedAppRemediationAction)()
    SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(value *ManagedAppRemediationAction)()
    SetAppActionIfDeviceLockNotSet(value *ManagedAppRemediationAction)()
    SetAppActionIfDevicePasscodeComplexityLessThanHigh(value *ManagedAppRemediationAction)()
    SetAppActionIfDevicePasscodeComplexityLessThanLow(value *ManagedAppRemediationAction)()
    SetAppActionIfDevicePasscodeComplexityLessThanMedium(value *ManagedAppRemediationAction)()
    SetAppActionIfSamsungKnoxAttestationRequired(value *ManagedAppRemediationAction)()
    SetApprovedKeyboards(value []KeyValuePairable)()
    SetApps(value []ManagedMobileAppable)()
    SetBiometricAuthenticationBlocked(value *bool)()
    SetBlockAfterCompanyPortalUpdateDeferralInDays(value *int32)()
    SetConnectToVpnOnLaunch(value *bool)()
    SetCustomBrowserDisplayName(value *string)()
    SetCustomBrowserPackageId(value *string)()
    SetCustomDialerAppDisplayName(value *string)()
    SetCustomDialerAppPackageId(value *string)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetDeviceLockRequired(value *bool)()
    SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)()
    SetEncryptAppData(value *bool)()
    SetExemptedAppPackages(value []KeyValuePairable)()
    SetFingerprintAndBiometricEnabled(value *bool)()
    SetKeyboardsRestricted(value *bool)()
    SetMessagingRedirectAppDisplayName(value *string)()
    SetMessagingRedirectAppPackageId(value *string)()
    SetMinimumRequiredCompanyPortalVersion(value *string)()
    SetMinimumRequiredPatchVersion(value *string)()
    SetMinimumWarningCompanyPortalVersion(value *string)()
    SetMinimumWarningPatchVersion(value *string)()
    SetMinimumWipeCompanyPortalVersion(value *string)()
    SetMinimumWipePatchVersion(value *string)()
    SetRequireClass3Biometrics(value *bool)()
    SetRequiredAndroidSafetyNetAppsVerificationType(value *AndroidManagedAppSafetyNetAppsVerificationType)()
    SetRequiredAndroidSafetyNetDeviceAttestationType(value *AndroidManagedAppSafetyNetDeviceAttestationType)()
    SetRequiredAndroidSafetyNetEvaluationType(value *AndroidManagedAppSafetyNetEvaluationType)()
    SetRequirePinAfterBiometricChange(value *bool)()
    SetScreenCaptureBlocked(value *bool)()
    SetWarnAfterCompanyPortalUpdateDeferralInDays(value *int32)()
    SetWipeAfterCompanyPortalUpdateDeferralInDays(value *int32)()
}
