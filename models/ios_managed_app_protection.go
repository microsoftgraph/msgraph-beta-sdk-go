package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosManagedAppProtection policy used to configure detailed management settings targeted to specific security groups and for a specified set of apps on an iOS device
type IosManagedAppProtection struct {
    TargetedManagedAppProtection
}
// NewIosManagedAppProtection instantiates a new IosManagedAppProtection and sets the default values.
func NewIosManagedAppProtection()(*IosManagedAppProtection) {
    m := &IosManagedAppProtection{
        TargetedManagedAppProtection: *NewTargetedManagedAppProtection(),
    }
    odataTypeValue := "#microsoft.graph.iosManagedAppProtection"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIosManagedAppProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosManagedAppProtection(), nil
}
// GetAllowedIosDeviceModels gets the allowedIosDeviceModels property value. Semicolon seperated list of device models allowed, as a string, for the managed app to work.
// returns a *string when successful
func (m *IosManagedAppProtection) GetAllowedIosDeviceModels()(*string) {
    val, err := m.GetBackingStore().Get("allowedIosDeviceModels")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAllowWidgetContentSync gets the allowWidgetContentSync property value. Indicates  if content sync for widgets is allowed for iOS on App Protection Policies
// returns a *bool when successful
func (m *IosManagedAppProtection) GetAllowWidgetContentSync()(*bool) {
    val, err := m.GetBackingStore().Get("allowWidgetContentSync")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAppActionIfAccountIsClockedOut gets the appActionIfAccountIsClockedOut property value. Defines a managed app behavior, either block or warn, if the user is clocked out (non-working time). Possible values are: block, wipe, warn.
// returns a *ManagedAppRemediationAction when successful
func (m *IosManagedAppProtection) GetAppActionIfAccountIsClockedOut()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfAccountIsClockedOut")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppActionIfIosDeviceModelNotAllowed gets the appActionIfIosDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
// returns a *ManagedAppRemediationAction when successful
func (m *IosManagedAppProtection) GetAppActionIfIosDeviceModelNotAllowed()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfIosDeviceModelNotAllowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppDataEncryptionType gets the appDataEncryptionType property value. Represents the level to which app data is encrypted for managed apps
// returns a *ManagedAppDataEncryptionType when successful
func (m *IosManagedAppProtection) GetAppDataEncryptionType()(*ManagedAppDataEncryptionType) {
    val, err := m.GetBackingStore().Get("appDataEncryptionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppDataEncryptionType)
    }
    return nil
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
// returns a []ManagedMobileAppable when successful
func (m *IosManagedAppProtection) GetApps()([]ManagedMobileAppable) {
    val, err := m.GetBackingStore().Get("apps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedMobileAppable)
    }
    return nil
}
// GetCustomBrowserProtocol gets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS.
// returns a *string when successful
func (m *IosManagedAppProtection) GetCustomBrowserProtocol()(*string) {
    val, err := m.GetBackingStore().Get("customBrowserProtocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomDialerAppProtocol gets the customDialerAppProtocol property value. Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
// returns a *string when successful
func (m *IosManagedAppProtection) GetCustomDialerAppProtocol()(*string) {
    val, err := m.GetBackingStore().Get("customDialerAppProtocol")
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
func (m *IosManagedAppProtection) GetDeployedAppCount()(*int32) {
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
func (m *IosManagedAppProtection) GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable) {
    val, err := m.GetBackingStore().Get("deploymentSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagedAppPolicyDeploymentSummaryable)
    }
    return nil
}
// GetDisableProtectionOfManagedOutboundOpenInData gets the disableProtectionOfManagedOutboundOpenInData property value. Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps.
// returns a *bool when successful
func (m *IosManagedAppProtection) GetDisableProtectionOfManagedOutboundOpenInData()(*bool) {
    val, err := m.GetBackingStore().Get("disableProtectionOfManagedOutboundOpenInData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetExemptedAppProtocols gets the exemptedAppProtocols property value. Apps in this list will be exempt from the policy and will be able to receive data from managed apps.
// returns a []KeyValuePairable when successful
func (m *IosManagedAppProtection) GetExemptedAppProtocols()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("exemptedAppProtocols")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// GetExemptedUniversalLinks gets the exemptedUniversalLinks property value. A list of custom urls that are allowed to invocate an unmanaged app
// returns a []string when successful
func (m *IosManagedAppProtection) GetExemptedUniversalLinks()([]string) {
    val, err := m.GetBackingStore().Get("exemptedUniversalLinks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFaceIdBlocked gets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
// returns a *bool when successful
func (m *IosManagedAppProtection) GetFaceIdBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("faceIdBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IosManagedAppProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TargetedManagedAppProtection.GetFieldDeserializers()
    res["allowedIosDeviceModels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedIosDeviceModels(val)
        }
        return nil
    }
    res["allowWidgetContentSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowWidgetContentSync(val)
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
    res["appActionIfIosDeviceModelNotAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfIosDeviceModelNotAllowed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appDataEncryptionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataEncryptionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDataEncryptionType(val.(*ManagedAppDataEncryptionType))
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
    res["customBrowserProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBrowserProtocol(val)
        }
        return nil
    }
    res["customDialerAppProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomDialerAppProtocol(val)
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
    res["disableProtectionOfManagedOutboundOpenInData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableProtectionOfManagedOutboundOpenInData(val)
        }
        return nil
    }
    res["exemptedAppProtocols"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetExemptedAppProtocols(res)
        }
        return nil
    }
    res["exemptedUniversalLinks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetExemptedUniversalLinks(res)
        }
        return nil
    }
    res["faceIdBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaceIdBlocked(val)
        }
        return nil
    }
    res["filterOpenInToOnlyManagedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilterOpenInToOnlyManagedApps(val)
        }
        return nil
    }
    res["managedUniversalLinks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetManagedUniversalLinks(res)
        }
        return nil
    }
    res["messagingRedirectAppUrlScheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessagingRedirectAppUrlScheme(val)
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
    res["minimumWarningSdkVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningSdkVersion(val)
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
    res["protectInboundDataFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectInboundDataFromUnknownSources(val)
        }
        return nil
    }
    res["thirdPartyKeyboardsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThirdPartyKeyboardsBlocked(val)
        }
        return nil
    }
    return res
}
// GetFilterOpenInToOnlyManagedApps gets the filterOpenInToOnlyManagedApps property value. Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False.
// returns a *bool when successful
func (m *IosManagedAppProtection) GetFilterOpenInToOnlyManagedApps()(*bool) {
    val, err := m.GetBackingStore().Get("filterOpenInToOnlyManagedApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManagedUniversalLinks gets the managedUniversalLinks property value. A list of custom urls that are allowed to invocate a managed app
// returns a []string when successful
func (m *IosManagedAppProtection) GetManagedUniversalLinks()([]string) {
    val, err := m.GetBackingStore().Get("managedUniversalLinks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetMessagingRedirectAppUrlScheme gets the messagingRedirectAppUrlScheme property value. When a specific app redirection is enforced by protectedMessagingRedirectAppType in an App Protection Policy, this value defines the app url redirect schemes which are allowed to be used.
// returns a *string when successful
func (m *IosManagedAppProtection) GetMessagingRedirectAppUrlScheme()(*string) {
    val, err := m.GetBackingStore().Get("messagingRedirectAppUrlScheme")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumRequiredSdkVersion gets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
// returns a *string when successful
func (m *IosManagedAppProtection) GetMinimumRequiredSdkVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumRequiredSdkVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWarningSdkVersion gets the minimumWarningSdkVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
// returns a *string when successful
func (m *IosManagedAppProtection) GetMinimumWarningSdkVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWarningSdkVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWipeSdkVersion gets the minimumWipeSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
// returns a *string when successful
func (m *IosManagedAppProtection) GetMinimumWipeSdkVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWipeSdkVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProtectInboundDataFromUnknownSources gets the protectInboundDataFromUnknownSources property value. Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps.
// returns a *bool when successful
func (m *IosManagedAppProtection) GetProtectInboundDataFromUnknownSources()(*bool) {
    val, err := m.GetBackingStore().Get("protectInboundDataFromUnknownSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetThirdPartyKeyboardsBlocked gets the thirdPartyKeyboardsBlocked property value. Defines if third party keyboards are allowed while accessing a managed app
// returns a *bool when successful
func (m *IosManagedAppProtection) GetThirdPartyKeyboardsBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("thirdPartyKeyboardsBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosManagedAppProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TargetedManagedAppProtection.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("allowedIosDeviceModels", m.GetAllowedIosDeviceModels())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowWidgetContentSync", m.GetAllowWidgetContentSync())
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
    if m.GetAppActionIfIosDeviceModelNotAllowed() != nil {
        cast := (*m.GetAppActionIfIosDeviceModelNotAllowed()).String()
        err = writer.WriteStringValue("appActionIfIosDeviceModelNotAllowed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppDataEncryptionType() != nil {
        cast := (*m.GetAppDataEncryptionType()).String()
        err = writer.WriteStringValue("appDataEncryptionType", &cast)
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
        err = writer.WriteStringValue("customBrowserProtocol", m.GetCustomBrowserProtocol())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customDialerAppProtocol", m.GetCustomDialerAppProtocol())
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
        err = writer.WriteBoolValue("disableProtectionOfManagedOutboundOpenInData", m.GetDisableProtectionOfManagedOutboundOpenInData())
        if err != nil {
            return err
        }
    }
    if m.GetExemptedAppProtocols() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExemptedAppProtocols()))
        for i, v := range m.GetExemptedAppProtocols() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("exemptedAppProtocols", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExemptedUniversalLinks() != nil {
        err = writer.WriteCollectionOfStringValues("exemptedUniversalLinks", m.GetExemptedUniversalLinks())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("faceIdBlocked", m.GetFaceIdBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("filterOpenInToOnlyManagedApps", m.GetFilterOpenInToOnlyManagedApps())
        if err != nil {
            return err
        }
    }
    if m.GetManagedUniversalLinks() != nil {
        err = writer.WriteCollectionOfStringValues("managedUniversalLinks", m.GetManagedUniversalLinks())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("messagingRedirectAppUrlScheme", m.GetMessagingRedirectAppUrlScheme())
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
        err = writer.WriteStringValue("minimumWarningSdkVersion", m.GetMinimumWarningSdkVersion())
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
    {
        err = writer.WriteBoolValue("protectInboundDataFromUnknownSources", m.GetProtectInboundDataFromUnknownSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("thirdPartyKeyboardsBlocked", m.GetThirdPartyKeyboardsBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedIosDeviceModels sets the allowedIosDeviceModels property value. Semicolon seperated list of device models allowed, as a string, for the managed app to work.
func (m *IosManagedAppProtection) SetAllowedIosDeviceModels(value *string)() {
    err := m.GetBackingStore().Set("allowedIosDeviceModels", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowWidgetContentSync sets the allowWidgetContentSync property value. Indicates  if content sync for widgets is allowed for iOS on App Protection Policies
func (m *IosManagedAppProtection) SetAllowWidgetContentSync(value *bool)() {
    err := m.GetBackingStore().Set("allowWidgetContentSync", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfAccountIsClockedOut sets the appActionIfAccountIsClockedOut property value. Defines a managed app behavior, either block or warn, if the user is clocked out (non-working time). Possible values are: block, wipe, warn.
func (m *IosManagedAppProtection) SetAppActionIfAccountIsClockedOut(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfAccountIsClockedOut", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfIosDeviceModelNotAllowed sets the appActionIfIosDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *IosManagedAppProtection) SetAppActionIfIosDeviceModelNotAllowed(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfIosDeviceModelNotAllowed", value)
    if err != nil {
        panic(err)
    }
}
// SetAppDataEncryptionType sets the appDataEncryptionType property value. Represents the level to which app data is encrypted for managed apps
func (m *IosManagedAppProtection) SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)() {
    err := m.GetBackingStore().Set("appDataEncryptionType", value)
    if err != nil {
        panic(err)
    }
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *IosManagedAppProtection) SetApps(value []ManagedMobileAppable)() {
    err := m.GetBackingStore().Set("apps", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomBrowserProtocol sets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS.
func (m *IosManagedAppProtection) SetCustomBrowserProtocol(value *string)() {
    err := m.GetBackingStore().Set("customBrowserProtocol", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomDialerAppProtocol sets the customDialerAppProtocol property value. Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
func (m *IosManagedAppProtection) SetCustomDialerAppProtocol(value *string)() {
    err := m.GetBackingStore().Set("customDialerAppProtocol", value)
    if err != nil {
        panic(err)
    }
}
// SetDeployedAppCount sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *IosManagedAppProtection) SetDeployedAppCount(value *int32)() {
    err := m.GetBackingStore().Set("deployedAppCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentSummary sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *IosManagedAppProtection) SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)() {
    err := m.GetBackingStore().Set("deploymentSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableProtectionOfManagedOutboundOpenInData sets the disableProtectionOfManagedOutboundOpenInData property value. Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps.
func (m *IosManagedAppProtection) SetDisableProtectionOfManagedOutboundOpenInData(value *bool)() {
    err := m.GetBackingStore().Set("disableProtectionOfManagedOutboundOpenInData", value)
    if err != nil {
        panic(err)
    }
}
// SetExemptedAppProtocols sets the exemptedAppProtocols property value. Apps in this list will be exempt from the policy and will be able to receive data from managed apps.
func (m *IosManagedAppProtection) SetExemptedAppProtocols(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("exemptedAppProtocols", value)
    if err != nil {
        panic(err)
    }
}
// SetExemptedUniversalLinks sets the exemptedUniversalLinks property value. A list of custom urls that are allowed to invocate an unmanaged app
func (m *IosManagedAppProtection) SetExemptedUniversalLinks(value []string)() {
    err := m.GetBackingStore().Set("exemptedUniversalLinks", value)
    if err != nil {
        panic(err)
    }
}
// SetFaceIdBlocked sets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
func (m *IosManagedAppProtection) SetFaceIdBlocked(value *bool)() {
    err := m.GetBackingStore().Set("faceIdBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetFilterOpenInToOnlyManagedApps sets the filterOpenInToOnlyManagedApps property value. Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False.
func (m *IosManagedAppProtection) SetFilterOpenInToOnlyManagedApps(value *bool)() {
    err := m.GetBackingStore().Set("filterOpenInToOnlyManagedApps", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedUniversalLinks sets the managedUniversalLinks property value. A list of custom urls that are allowed to invocate a managed app
func (m *IosManagedAppProtection) SetManagedUniversalLinks(value []string)() {
    err := m.GetBackingStore().Set("managedUniversalLinks", value)
    if err != nil {
        panic(err)
    }
}
// SetMessagingRedirectAppUrlScheme sets the messagingRedirectAppUrlScheme property value. When a specific app redirection is enforced by protectedMessagingRedirectAppType in an App Protection Policy, this value defines the app url redirect schemes which are allowed to be used.
func (m *IosManagedAppProtection) SetMessagingRedirectAppUrlScheme(value *string)() {
    err := m.GetBackingStore().Set("messagingRedirectAppUrlScheme", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumRequiredSdkVersion sets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) SetMinimumRequiredSdkVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumRequiredSdkVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWarningSdkVersion sets the minimumWarningSdkVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
func (m *IosManagedAppProtection) SetMinimumWarningSdkVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWarningSdkVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWipeSdkVersion sets the minimumWipeSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) SetMinimumWipeSdkVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWipeSdkVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetProtectInboundDataFromUnknownSources sets the protectInboundDataFromUnknownSources property value. Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps.
func (m *IosManagedAppProtection) SetProtectInboundDataFromUnknownSources(value *bool)() {
    err := m.GetBackingStore().Set("protectInboundDataFromUnknownSources", value)
    if err != nil {
        panic(err)
    }
}
// SetThirdPartyKeyboardsBlocked sets the thirdPartyKeyboardsBlocked property value. Defines if third party keyboards are allowed while accessing a managed app
func (m *IosManagedAppProtection) SetThirdPartyKeyboardsBlocked(value *bool)() {
    err := m.GetBackingStore().Set("thirdPartyKeyboardsBlocked", value)
    if err != nil {
        panic(err)
    }
}
type IosManagedAppProtectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TargetedManagedAppProtectionable
    GetAllowedIosDeviceModels()(*string)
    GetAllowWidgetContentSync()(*bool)
    GetAppActionIfAccountIsClockedOut()(*ManagedAppRemediationAction)
    GetAppActionIfIosDeviceModelNotAllowed()(*ManagedAppRemediationAction)
    GetAppDataEncryptionType()(*ManagedAppDataEncryptionType)
    GetApps()([]ManagedMobileAppable)
    GetCustomBrowserProtocol()(*string)
    GetCustomDialerAppProtocol()(*string)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetDisableProtectionOfManagedOutboundOpenInData()(*bool)
    GetExemptedAppProtocols()([]KeyValuePairable)
    GetExemptedUniversalLinks()([]string)
    GetFaceIdBlocked()(*bool)
    GetFilterOpenInToOnlyManagedApps()(*bool)
    GetManagedUniversalLinks()([]string)
    GetMessagingRedirectAppUrlScheme()(*string)
    GetMinimumRequiredSdkVersion()(*string)
    GetMinimumWarningSdkVersion()(*string)
    GetMinimumWipeSdkVersion()(*string)
    GetProtectInboundDataFromUnknownSources()(*bool)
    GetThirdPartyKeyboardsBlocked()(*bool)
    SetAllowedIosDeviceModels(value *string)()
    SetAllowWidgetContentSync(value *bool)()
    SetAppActionIfAccountIsClockedOut(value *ManagedAppRemediationAction)()
    SetAppActionIfIosDeviceModelNotAllowed(value *ManagedAppRemediationAction)()
    SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)()
    SetApps(value []ManagedMobileAppable)()
    SetCustomBrowserProtocol(value *string)()
    SetCustomDialerAppProtocol(value *string)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetDisableProtectionOfManagedOutboundOpenInData(value *bool)()
    SetExemptedAppProtocols(value []KeyValuePairable)()
    SetExemptedUniversalLinks(value []string)()
    SetFaceIdBlocked(value *bool)()
    SetFilterOpenInToOnlyManagedApps(value *bool)()
    SetManagedUniversalLinks(value []string)()
    SetMessagingRedirectAppUrlScheme(value *string)()
    SetMinimumRequiredSdkVersion(value *string)()
    SetMinimumWarningSdkVersion(value *string)()
    SetMinimumWipeSdkVersion(value *string)()
    SetProtectInboundDataFromUnknownSources(value *bool)()
    SetThirdPartyKeyboardsBlocked(value *bool)()
}
