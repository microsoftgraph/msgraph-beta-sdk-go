package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsManagedAppProtection policy used to configure detailed management settings targeted to specific security groups and for a specified set of apps on a Windows device
type WindowsManagedAppProtection struct {
    ManagedAppPolicy
}
// NewWindowsManagedAppProtection instantiates a new WindowsManagedAppProtection and sets the default values.
func NewWindowsManagedAppProtection()(*WindowsManagedAppProtection) {
    m := &WindowsManagedAppProtection{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    odataTypeValue := "#microsoft.graph.windowsManagedAppProtection"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsManagedAppProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsManagedAppProtection(), nil
}
// GetAllowedInboundDataTransferSources gets the allowedInboundDataTransferSources property value. Data can be transferred from/to these classes of apps
// returns a *WindowsManagedAppDataTransferLevel when successful
func (m *WindowsManagedAppProtection) GetAllowedInboundDataTransferSources()(*WindowsManagedAppDataTransferLevel) {
    val, err := m.GetBackingStore().Get("allowedInboundDataTransferSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsManagedAppDataTransferLevel)
    }
    return nil
}
// GetAllowedOutboundClipboardSharingLevel gets the allowedOutboundClipboardSharingLevel property value. Represents the level to which the device's clipboard may be shared between apps
// returns a *WindowsManagedAppClipboardSharingLevel when successful
func (m *WindowsManagedAppProtection) GetAllowedOutboundClipboardSharingLevel()(*WindowsManagedAppClipboardSharingLevel) {
    val, err := m.GetBackingStore().Get("allowedOutboundClipboardSharingLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsManagedAppClipboardSharingLevel)
    }
    return nil
}
// GetAllowedOutboundDataTransferDestinations gets the allowedOutboundDataTransferDestinations property value. Data can be transferred from/to these classes of apps
// returns a *WindowsManagedAppDataTransferLevel when successful
func (m *WindowsManagedAppProtection) GetAllowedOutboundDataTransferDestinations()(*WindowsManagedAppDataTransferLevel) {
    val, err := m.GetBackingStore().Get("allowedOutboundDataTransferDestinations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsManagedAppDataTransferLevel)
    }
    return nil
}
// GetAppActionIfUnableToAuthenticateUser gets the appActionIfUnableToAuthenticateUser property value. If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Some possible values are block or wipe. If this property is not set, no action will be taken. Possible values are: block, wipe, warn.
// returns a *ManagedAppRemediationAction when successful
func (m *WindowsManagedAppProtection) GetAppActionIfUnableToAuthenticateUser()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfUnableToAuthenticateUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
// returns a []ManagedMobileAppable when successful
func (m *WindowsManagedAppProtection) GetApps()([]ManagedMobileAppable) {
    val, err := m.GetBackingStore().Get("apps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedMobileAppable)
    }
    return nil
}
// GetAssignments gets the assignments property value. Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
// returns a []TargetedManagedAppPolicyAssignmentable when successful
func (m *WindowsManagedAppProtection) GetAssignments()([]TargetedManagedAppPolicyAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TargetedManagedAppPolicyAssignmentable)
    }
    return nil
}
// GetDeployedAppCount gets the deployedAppCount property value. Indicates the total number of applications for which the current policy is deployed.
// returns a *int32 when successful
func (m *WindowsManagedAppProtection) GetDeployedAppCount()(*int32) {
    val, err := m.GetBackingStore().Get("deployedAppCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsManagedAppProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedAppPolicy.GetFieldDeserializers()
    res["allowedInboundDataTransferSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedInboundDataTransferSources(val.(*WindowsManagedAppDataTransferLevel))
        }
        return nil
    }
    res["allowedOutboundClipboardSharingLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsManagedAppClipboardSharingLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundClipboardSharingLevel(val.(*WindowsManagedAppClipboardSharingLevel))
        }
        return nil
    }
    res["allowedOutboundDataTransferDestinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundDataTransferDestinations(val.(*WindowsManagedAppDataTransferLevel))
        }
        return nil
    }
    res["appActionIfUnableToAuthenticateUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfUnableToAuthenticateUser(val.(*ManagedAppRemediationAction))
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
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTargetedManagedAppPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TargetedManagedAppPolicyAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TargetedManagedAppPolicyAssignmentable)
                }
            }
            m.SetAssignments(res)
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
    res["isAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssigned(val)
        }
        return nil
    }
    res["maximumAllowedDeviceThreatLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDeviceThreatLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAllowedDeviceThreatLevel(val.(*ManagedAppDeviceThreatLevel))
        }
        return nil
    }
    res["maximumRequiredOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumRequiredOsVersion(val)
        }
        return nil
    }
    res["maximumWarningOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumWarningOsVersion(val)
        }
        return nil
    }
    res["maximumWipeOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumWipeOsVersion(val)
        }
        return nil
    }
    res["minimumRequiredAppVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredAppVersion(val)
        }
        return nil
    }
    res["minimumRequiredOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredOsVersion(val)
        }
        return nil
    }
    res["minimumRequiredSdkVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredSdkVersion(val)
        }
        return nil
    }
    res["minimumWarningAppVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningAppVersion(val)
        }
        return nil
    }
    res["minimumWarningOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningOsVersion(val)
        }
        return nil
    }
    res["minimumWipeAppVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipeAppVersion(val)
        }
        return nil
    }
    res["minimumWipeOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipeOsVersion(val)
        }
        return nil
    }
    res["minimumWipeSdkVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipeSdkVersion(val)
        }
        return nil
    }
    res["mobileThreatDefenseRemediationAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileThreatDefenseRemediationAction(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["periodOfflineBeforeAccessCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodOfflineBeforeAccessCheck(val)
        }
        return nil
    }
    res["periodOfflineBeforeWipeIsEnforced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodOfflineBeforeWipeIsEnforced(val)
        }
        return nil
    }
    res["printBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrintBlocked(val)
        }
        return nil
    }
    return res
}
// GetIsAssigned gets the isAssigned property value. When TRUE, indicates that the policy is deployed to some inclusion groups. When FALSE, indicates that the policy is not deployed to any inclusion groups. Default value is FALSE.
// returns a *bool when successful
func (m *WindowsManagedAppProtection) GetIsAssigned()(*bool) {
    val, err := m.GetBackingStore().Get("isAssigned")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMaximumAllowedDeviceThreatLevel gets the maximumAllowedDeviceThreatLevel property value. The maxium threat level allowed for an app to be compliant.
// returns a *ManagedAppDeviceThreatLevel when successful
func (m *WindowsManagedAppProtection) GetMaximumAllowedDeviceThreatLevel()(*ManagedAppDeviceThreatLevel) {
    val, err := m.GetBackingStore().Get("maximumAllowedDeviceThreatLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppDeviceThreatLevel)
    }
    return nil
}
// GetMaximumRequiredOsVersion gets the maximumRequiredOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
// returns a *string when successful
func (m *WindowsManagedAppProtection) GetMaximumRequiredOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("maximumRequiredOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMaximumWarningOsVersion gets the maximumWarningOsVersion property value. Versions bigger than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
// returns a *string when successful
func (m *WindowsManagedAppProtection) GetMaximumWarningOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("maximumWarningOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMaximumWipeOsVersion gets the maximumWipeOsVersion property value. Versions bigger than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
// returns a *string when successful
func (m *WindowsManagedAppProtection) GetMaximumWipeOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("maximumWipeOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumRequiredAppVersion gets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
// returns a *string when successful
func (m *WindowsManagedAppProtection) GetMinimumRequiredAppVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumRequiredAppVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumRequiredOsVersion gets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
// returns a *string when successful
func (m *WindowsManagedAppProtection) GetMinimumRequiredOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumRequiredOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumRequiredSdkVersion gets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
// returns a *string when successful
func (m *WindowsManagedAppProtection) GetMinimumRequiredSdkVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumRequiredSdkVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWarningAppVersion gets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
// returns a *string when successful
func (m *WindowsManagedAppProtection) GetMinimumWarningAppVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWarningAppVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWarningOsVersion gets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
// returns a *string when successful
func (m *WindowsManagedAppProtection) GetMinimumWarningOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWarningOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWipeAppVersion gets the minimumWipeAppVersion property value. Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
// returns a *string when successful
func (m *WindowsManagedAppProtection) GetMinimumWipeAppVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWipeAppVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWipeOsVersion gets the minimumWipeOsVersion property value. Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
// returns a *string when successful
func (m *WindowsManagedAppProtection) GetMinimumWipeOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWipeOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWipeSdkVersion gets the minimumWipeSdkVersion property value. Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
// returns a *string when successful
func (m *WindowsManagedAppProtection) GetMinimumWipeSdkVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWipeSdkVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMobileThreatDefenseRemediationAction gets the mobileThreatDefenseRemediationAction property value. An admin initiated action to be applied on a managed app.
// returns a *ManagedAppRemediationAction when successful
func (m *WindowsManagedAppProtection) GetMobileThreatDefenseRemediationAction()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("mobileThreatDefenseRemediationAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetPeriodOfflineBeforeAccessCheck gets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet. For example, PT5M indicates that the interval is 5 minutes in duration. A timespan value of PT0S indicates that access will be blocked immediately when the device is not connected to the internet.
// returns a *ISODuration when successful
func (m *WindowsManagedAppProtection) GetPeriodOfflineBeforeAccessCheck()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("periodOfflineBeforeAccessCheck")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetPeriodOfflineBeforeWipeIsEnforced gets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped. For example, P5D indicates that the interval is 5 days in duration. A timespan value of PT0S indicates that managed data will never be wiped when the device is not connected to the internet.
// returns a *ISODuration when successful
func (m *WindowsManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("periodOfflineBeforeWipeIsEnforced")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetPrintBlocked gets the printBlocked property value. When TRUE, indicates that printing is blocked from managed apps. When FALSE, indicates that printing is allowed from managed apps. Default value is FALSE.
// returns a *bool when successful
func (m *WindowsManagedAppProtection) GetPrintBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("printBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsManagedAppProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedAppPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedInboundDataTransferSources() != nil {
        cast := (*m.GetAllowedInboundDataTransferSources()).String()
        err = writer.WriteStringValue("allowedInboundDataTransferSources", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowedOutboundClipboardSharingLevel() != nil {
        cast := (*m.GetAllowedOutboundClipboardSharingLevel()).String()
        err = writer.WriteStringValue("allowedOutboundClipboardSharingLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowedOutboundDataTransferDestinations() != nil {
        cast := (*m.GetAllowedOutboundDataTransferDestinations()).String()
        err = writer.WriteStringValue("allowedOutboundDataTransferDestinations", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfUnableToAuthenticateUser() != nil {
        cast := (*m.GetAppActionIfUnableToAuthenticateUser()).String()
        err = writer.WriteStringValue("appActionIfUnableToAuthenticateUser", &cast)
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
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
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
        err = writer.WriteBoolValue("isAssigned", m.GetIsAssigned())
        if err != nil {
            return err
        }
    }
    if m.GetMaximumAllowedDeviceThreatLevel() != nil {
        cast := (*m.GetMaximumAllowedDeviceThreatLevel()).String()
        err = writer.WriteStringValue("maximumAllowedDeviceThreatLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("maximumRequiredOsVersion", m.GetMaximumRequiredOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("maximumWarningOsVersion", m.GetMaximumWarningOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("maximumWipeOsVersion", m.GetMaximumWipeOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredAppVersion", m.GetMinimumRequiredAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredOsVersion", m.GetMinimumRequiredOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredSdkVersion", m.GetMinimumRequiredSdkVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWarningAppVersion", m.GetMinimumWarningAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWarningOsVersion", m.GetMinimumWarningOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWipeAppVersion", m.GetMinimumWipeAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWipeOsVersion", m.GetMinimumWipeOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWipeSdkVersion", m.GetMinimumWipeSdkVersion())
        if err != nil {
            return err
        }
    }
    if m.GetMobileThreatDefenseRemediationAction() != nil {
        cast := (*m.GetMobileThreatDefenseRemediationAction()).String()
        err = writer.WriteStringValue("mobileThreatDefenseRemediationAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("periodOfflineBeforeAccessCheck", m.GetPeriodOfflineBeforeAccessCheck())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("periodOfflineBeforeWipeIsEnforced", m.GetPeriodOfflineBeforeWipeIsEnforced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("printBlocked", m.GetPrintBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedInboundDataTransferSources sets the allowedInboundDataTransferSources property value. Data can be transferred from/to these classes of apps
func (m *WindowsManagedAppProtection) SetAllowedInboundDataTransferSources(value *WindowsManagedAppDataTransferLevel)() {
    err := m.GetBackingStore().Set("allowedInboundDataTransferSources", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedOutboundClipboardSharingLevel sets the allowedOutboundClipboardSharingLevel property value. Represents the level to which the device's clipboard may be shared between apps
func (m *WindowsManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(value *WindowsManagedAppClipboardSharingLevel)() {
    err := m.GetBackingStore().Set("allowedOutboundClipboardSharingLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedOutboundDataTransferDestinations sets the allowedOutboundDataTransferDestinations property value. Data can be transferred from/to these classes of apps
func (m *WindowsManagedAppProtection) SetAllowedOutboundDataTransferDestinations(value *WindowsManagedAppDataTransferLevel)() {
    err := m.GetBackingStore().Set("allowedOutboundDataTransferDestinations", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfUnableToAuthenticateUser sets the appActionIfUnableToAuthenticateUser property value. If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Some possible values are block or wipe. If this property is not set, no action will be taken. Possible values are: block, wipe, warn.
func (m *WindowsManagedAppProtection) SetAppActionIfUnableToAuthenticateUser(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfUnableToAuthenticateUser", value)
    if err != nil {
        panic(err)
    }
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *WindowsManagedAppProtection) SetApps(value []ManagedMobileAppable)() {
    err := m.GetBackingStore().Set("apps", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignments sets the assignments property value. Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
func (m *WindowsManagedAppProtection) SetAssignments(value []TargetedManagedAppPolicyAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetDeployedAppCount sets the deployedAppCount property value. Indicates the total number of applications for which the current policy is deployed.
func (m *WindowsManagedAppProtection) SetDeployedAppCount(value *int32)() {
    err := m.GetBackingStore().Set("deployedAppCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAssigned sets the isAssigned property value. When TRUE, indicates that the policy is deployed to some inclusion groups. When FALSE, indicates that the policy is not deployed to any inclusion groups. Default value is FALSE.
func (m *WindowsManagedAppProtection) SetIsAssigned(value *bool)() {
    err := m.GetBackingStore().Set("isAssigned", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumAllowedDeviceThreatLevel sets the maximumAllowedDeviceThreatLevel property value. The maxium threat level allowed for an app to be compliant.
func (m *WindowsManagedAppProtection) SetMaximumAllowedDeviceThreatLevel(value *ManagedAppDeviceThreatLevel)() {
    err := m.GetBackingStore().Set("maximumAllowedDeviceThreatLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumRequiredOsVersion sets the maximumRequiredOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMaximumRequiredOsVersion(value *string)() {
    err := m.GetBackingStore().Set("maximumRequiredOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumWarningOsVersion sets the maximumWarningOsVersion property value. Versions bigger than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMaximumWarningOsVersion(value *string)() {
    err := m.GetBackingStore().Set("maximumWarningOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumWipeOsVersion sets the maximumWipeOsVersion property value. Versions bigger than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMaximumWipeOsVersion(value *string)() {
    err := m.GetBackingStore().Set("maximumWipeOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumRequiredAppVersion sets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumRequiredAppVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumRequiredAppVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumRequiredOsVersion sets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumRequiredOsVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumRequiredOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumRequiredSdkVersion sets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumRequiredSdkVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumRequiredSdkVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWarningAppVersion sets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumWarningAppVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWarningAppVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWarningOsVersion sets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumWarningOsVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWarningOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWipeAppVersion sets the minimumWipeAppVersion property value. Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumWipeAppVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWipeAppVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWipeOsVersion sets the minimumWipeOsVersion property value. Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumWipeOsVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWipeOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWipeSdkVersion sets the minimumWipeSdkVersion property value. Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumWipeSdkVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWipeSdkVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileThreatDefenseRemediationAction sets the mobileThreatDefenseRemediationAction property value. An admin initiated action to be applied on a managed app.
func (m *WindowsManagedAppProtection) SetMobileThreatDefenseRemediationAction(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("mobileThreatDefenseRemediationAction", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriodOfflineBeforeAccessCheck sets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet. For example, PT5M indicates that the interval is 5 minutes in duration. A timespan value of PT0S indicates that access will be blocked immediately when the device is not connected to the internet.
func (m *WindowsManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("periodOfflineBeforeAccessCheck", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriodOfflineBeforeWipeIsEnforced sets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped. For example, P5D indicates that the interval is 5 days in duration. A timespan value of PT0S indicates that managed data will never be wiped when the device is not connected to the internet.
func (m *WindowsManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("periodOfflineBeforeWipeIsEnforced", value)
    if err != nil {
        panic(err)
    }
}
// SetPrintBlocked sets the printBlocked property value. When TRUE, indicates that printing is blocked from managed apps. When FALSE, indicates that printing is allowed from managed apps. Default value is FALSE.
func (m *WindowsManagedAppProtection) SetPrintBlocked(value *bool)() {
    err := m.GetBackingStore().Set("printBlocked", value)
    if err != nil {
        panic(err)
    }
}
type WindowsManagedAppProtectionable interface {
    ManagedAppPolicyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedInboundDataTransferSources()(*WindowsManagedAppDataTransferLevel)
    GetAllowedOutboundClipboardSharingLevel()(*WindowsManagedAppClipboardSharingLevel)
    GetAllowedOutboundDataTransferDestinations()(*WindowsManagedAppDataTransferLevel)
    GetAppActionIfUnableToAuthenticateUser()(*ManagedAppRemediationAction)
    GetApps()([]ManagedMobileAppable)
    GetAssignments()([]TargetedManagedAppPolicyAssignmentable)
    GetDeployedAppCount()(*int32)
    GetIsAssigned()(*bool)
    GetMaximumAllowedDeviceThreatLevel()(*ManagedAppDeviceThreatLevel)
    GetMaximumRequiredOsVersion()(*string)
    GetMaximumWarningOsVersion()(*string)
    GetMaximumWipeOsVersion()(*string)
    GetMinimumRequiredAppVersion()(*string)
    GetMinimumRequiredOsVersion()(*string)
    GetMinimumRequiredSdkVersion()(*string)
    GetMinimumWarningAppVersion()(*string)
    GetMinimumWarningOsVersion()(*string)
    GetMinimumWipeAppVersion()(*string)
    GetMinimumWipeOsVersion()(*string)
    GetMinimumWipeSdkVersion()(*string)
    GetMobileThreatDefenseRemediationAction()(*ManagedAppRemediationAction)
    GetPeriodOfflineBeforeAccessCheck()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPeriodOfflineBeforeWipeIsEnforced()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPrintBlocked()(*bool)
    SetAllowedInboundDataTransferSources(value *WindowsManagedAppDataTransferLevel)()
    SetAllowedOutboundClipboardSharingLevel(value *WindowsManagedAppClipboardSharingLevel)()
    SetAllowedOutboundDataTransferDestinations(value *WindowsManagedAppDataTransferLevel)()
    SetAppActionIfUnableToAuthenticateUser(value *ManagedAppRemediationAction)()
    SetApps(value []ManagedMobileAppable)()
    SetAssignments(value []TargetedManagedAppPolicyAssignmentable)()
    SetDeployedAppCount(value *int32)()
    SetIsAssigned(value *bool)()
    SetMaximumAllowedDeviceThreatLevel(value *ManagedAppDeviceThreatLevel)()
    SetMaximumRequiredOsVersion(value *string)()
    SetMaximumWarningOsVersion(value *string)()
    SetMaximumWipeOsVersion(value *string)()
    SetMinimumRequiredAppVersion(value *string)()
    SetMinimumRequiredOsVersion(value *string)()
    SetMinimumRequiredSdkVersion(value *string)()
    SetMinimumWarningAppVersion(value *string)()
    SetMinimumWarningOsVersion(value *string)()
    SetMinimumWipeAppVersion(value *string)()
    SetMinimumWipeOsVersion(value *string)()
    SetMinimumWipeSdkVersion(value *string)()
    SetMobileThreatDefenseRemediationAction(value *ManagedAppRemediationAction)()
    SetPeriodOfflineBeforeAccessCheck(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPeriodOfflineBeforeWipeIsEnforced(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPrintBlocked(value *bool)()
}
