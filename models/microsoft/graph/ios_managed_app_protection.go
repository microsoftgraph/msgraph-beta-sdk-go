package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IosManagedAppProtection 
type IosManagedAppProtection struct {
    TargetedManagedAppProtection
    // Semicolon seperated list of device models allowed, as a string, for the managed app to work.
    allowedIosDeviceModels *string;
    // Defines a managed app behavior, either block or wipe, if the specified device model is not allowed. Possible values are: block, wipe, warn.
    appActionIfIosDeviceModelNotAllowed *ManagedAppRemediationAction;
    // Type of encryption which should be used for data in a managed app. Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked.
    appDataEncryptionType *ManagedAppDataEncryptionType;
    // List of apps to which the policy is deployed.
    apps []ManagedMobileApp;
    // A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
    customBrowserProtocol *string;
    // Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
    customDialerAppProtocol *string;
    // Count of apps to which the current policy is deployed.
    deployedAppCount *int32;
    // Navigation property to deployment summary of the configuration.
    deploymentSummary *ManagedAppPolicyDeploymentSummary;
    // Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps.
    disableProtectionOfManagedOutboundOpenInData *bool;
    // Apps in this list will be exempt from the policy and will be able to receive data from managed apps.
    exemptedAppProtocols []KeyValuePair;
    // A list of custom urls that are allowed to invocate an unmanaged app
    exemptedUniversalLinks []string;
    // Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
    faceIdBlocked *bool;
    // Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False.
    filterOpenInToOnlyManagedApps *bool;
    // A list of custom urls that are allowed to invocate a managed app
    managedUniversalLinks []string;
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumRequiredSdkVersion *string;
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumWipeSdkVersion *string;
    // Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps.
    protectInboundDataFromUnknownSources *bool;
    // Defines if third party keyboards are allowed while accessing a managed app
    thirdPartyKeyboardsBlocked *bool;
}
// NewIosManagedAppProtection instantiates a new iosManagedAppProtection and sets the default values.
func NewIosManagedAppProtection()(*IosManagedAppProtection) {
    m := &IosManagedAppProtection{
        TargetedManagedAppProtection: *NewTargetedManagedAppProtection(),
    }
    return m
}
// GetAllowedIosDeviceModels gets the allowedIosDeviceModels property value. Semicolon seperated list of device models allowed, as a string, for the managed app to work.
func (m *IosManagedAppProtection) GetAllowedIosDeviceModels()(*string) {
    if m == nil {
        return nil
    } else {
        return m.allowedIosDeviceModels
    }
}
// GetAppActionIfIosDeviceModelNotAllowed gets the appActionIfIosDeviceModelNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device model is not allowed. Possible values are: block, wipe, warn.
func (m *IosManagedAppProtection) GetAppActionIfIosDeviceModelNotAllowed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfIosDeviceModelNotAllowed
    }
}
// GetAppDataEncryptionType gets the appDataEncryptionType property value. Type of encryption which should be used for data in a managed app. Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked.
func (m *IosManagedAppProtection) GetAppDataEncryptionType()(*ManagedAppDataEncryptionType) {
    if m == nil {
        return nil
    } else {
        return m.appDataEncryptionType
    }
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
func (m *IosManagedAppProtection) GetApps()([]ManagedMobileApp) {
    if m == nil {
        return nil
    } else {
        return m.apps
    }
}
// GetCustomBrowserProtocol gets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *IosManagedAppProtection) GetCustomBrowserProtocol()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customBrowserProtocol
    }
}
// GetCustomDialerAppProtocol gets the customDialerAppProtocol property value. Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
func (m *IosManagedAppProtection) GetCustomDialerAppProtocol()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customDialerAppProtocol
    }
}
// GetDeployedAppCount gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *IosManagedAppProtection) GetDeployedAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deployedAppCount
    }
}
// GetDeploymentSummary gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *IosManagedAppProtection) GetDeploymentSummary()(*ManagedAppPolicyDeploymentSummary) {
    if m == nil {
        return nil
    } else {
        return m.deploymentSummary
    }
}
// GetDisableProtectionOfManagedOutboundOpenInData gets the disableProtectionOfManagedOutboundOpenInData property value. Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps.
func (m *IosManagedAppProtection) GetDisableProtectionOfManagedOutboundOpenInData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableProtectionOfManagedOutboundOpenInData
    }
}
// GetExemptedAppProtocols gets the exemptedAppProtocols property value. Apps in this list will be exempt from the policy and will be able to receive data from managed apps.
func (m *IosManagedAppProtection) GetExemptedAppProtocols()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.exemptedAppProtocols
    }
}
// GetExemptedUniversalLinks gets the exemptedUniversalLinks property value. A list of custom urls that are allowed to invocate an unmanaged app
func (m *IosManagedAppProtection) GetExemptedUniversalLinks()([]string) {
    if m == nil {
        return nil
    } else {
        return m.exemptedUniversalLinks
    }
}
// GetFaceIdBlocked gets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
func (m *IosManagedAppProtection) GetFaceIdBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.faceIdBlocked
    }
}
// GetFilterOpenInToOnlyManagedApps gets the filterOpenInToOnlyManagedApps property value. Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False.
func (m *IosManagedAppProtection) GetFilterOpenInToOnlyManagedApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.filterOpenInToOnlyManagedApps
    }
}
// GetManagedUniversalLinks gets the managedUniversalLinks property value. A list of custom urls that are allowed to invocate a managed app
func (m *IosManagedAppProtection) GetManagedUniversalLinks()([]string) {
    if m == nil {
        return nil
    } else {
        return m.managedUniversalLinks
    }
}
// GetMinimumRequiredSdkVersion gets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) GetMinimumRequiredSdkVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredSdkVersion
    }
}
// GetMinimumWipeSdkVersion gets the minimumWipeSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) GetMinimumWipeSdkVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWipeSdkVersion
    }
}
// GetProtectInboundDataFromUnknownSources gets the protectInboundDataFromUnknownSources property value. Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps.
func (m *IosManagedAppProtection) GetProtectInboundDataFromUnknownSources()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.protectInboundDataFromUnknownSources
    }
}
// GetThirdPartyKeyboardsBlocked gets the thirdPartyKeyboardsBlocked property value. Defines if third party keyboards are allowed while accessing a managed app
func (m *IosManagedAppProtection) GetThirdPartyKeyboardsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.thirdPartyKeyboardsBlocked
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosManagedAppProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.TargetedManagedAppProtection.GetFieldDeserializers()
    res["allowedIosDeviceModels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedIosDeviceModels(val)
        }
        return nil
    }
    res["appActionIfIosDeviceModelNotAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfIosDeviceModelNotAllowed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appDataEncryptionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataEncryptionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDataEncryptionType(val.(*ManagedAppDataEncryptionType))
        }
        return nil
    }
    res["apps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedMobileApp() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedMobileApp, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedMobileApp))
            }
            m.SetApps(res)
        }
        return nil
    }
    res["customBrowserProtocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBrowserProtocol(val)
        }
        return nil
    }
    res["customDialerAppProtocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomDialerAppProtocol(val)
        }
        return nil
    }
    res["deployedAppCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedAppCount(val)
        }
        return nil
    }
    res["deploymentSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicyDeploymentSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentSummary(val.(*ManagedAppPolicyDeploymentSummary))
        }
        return nil
    }
    res["disableProtectionOfManagedOutboundOpenInData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableProtectionOfManagedOutboundOpenInData(val)
        }
        return nil
    }
    res["exemptedAppProtocols"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValuePair))
            }
            m.SetExemptedAppProtocols(res)
        }
        return nil
    }
    res["exemptedUniversalLinks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExemptedUniversalLinks(res)
        }
        return nil
    }
    res["faceIdBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaceIdBlocked(val)
        }
        return nil
    }
    res["filterOpenInToOnlyManagedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilterOpenInToOnlyManagedApps(val)
        }
        return nil
    }
    res["managedUniversalLinks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetManagedUniversalLinks(res)
        }
        return nil
    }
    res["minimumRequiredSdkVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredSdkVersion(val)
        }
        return nil
    }
    res["minimumWipeSdkVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipeSdkVersion(val)
        }
        return nil
    }
    res["protectInboundDataFromUnknownSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectInboundDataFromUnknownSources(val)
        }
        return nil
    }
    res["thirdPartyKeyboardsBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *IosManagedAppProtection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IosManagedAppProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExemptedAppProtocols()))
        for i, v := range m.GetExemptedAppProtocols() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        err = writer.WriteStringValue("minimumRequiredSdkVersion", m.GetMinimumRequiredSdkVersion())
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
    if m != nil {
        m.allowedIosDeviceModels = value
    }
}
// SetAppActionIfIosDeviceModelNotAllowed sets the appActionIfIosDeviceModelNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device model is not allowed. Possible values are: block, wipe, warn.
func (m *IosManagedAppProtection) SetAppActionIfIosDeviceModelNotAllowed(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfIosDeviceModelNotAllowed = value
    }
}
// SetAppDataEncryptionType sets the appDataEncryptionType property value. Type of encryption which should be used for data in a managed app. Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked.
func (m *IosManagedAppProtection) SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)() {
    if m != nil {
        m.appDataEncryptionType = value
    }
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *IosManagedAppProtection) SetApps(value []ManagedMobileApp)() {
    if m != nil {
        m.apps = value
    }
}
// SetCustomBrowserProtocol sets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *IosManagedAppProtection) SetCustomBrowserProtocol(value *string)() {
    if m != nil {
        m.customBrowserProtocol = value
    }
}
// SetCustomDialerAppProtocol sets the customDialerAppProtocol property value. Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
func (m *IosManagedAppProtection) SetCustomDialerAppProtocol(value *string)() {
    if m != nil {
        m.customDialerAppProtocol = value
    }
}
// SetDeployedAppCount sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *IosManagedAppProtection) SetDeployedAppCount(value *int32)() {
    if m != nil {
        m.deployedAppCount = value
    }
}
// SetDeploymentSummary sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *IosManagedAppProtection) SetDeploymentSummary(value *ManagedAppPolicyDeploymentSummary)() {
    if m != nil {
        m.deploymentSummary = value
    }
}
// SetDisableProtectionOfManagedOutboundOpenInData sets the disableProtectionOfManagedOutboundOpenInData property value. Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps.
func (m *IosManagedAppProtection) SetDisableProtectionOfManagedOutboundOpenInData(value *bool)() {
    if m != nil {
        m.disableProtectionOfManagedOutboundOpenInData = value
    }
}
// SetExemptedAppProtocols sets the exemptedAppProtocols property value. Apps in this list will be exempt from the policy and will be able to receive data from managed apps.
func (m *IosManagedAppProtection) SetExemptedAppProtocols(value []KeyValuePair)() {
    if m != nil {
        m.exemptedAppProtocols = value
    }
}
// SetExemptedUniversalLinks sets the exemptedUniversalLinks property value. A list of custom urls that are allowed to invocate an unmanaged app
func (m *IosManagedAppProtection) SetExemptedUniversalLinks(value []string)() {
    if m != nil {
        m.exemptedUniversalLinks = value
    }
}
// SetFaceIdBlocked sets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
func (m *IosManagedAppProtection) SetFaceIdBlocked(value *bool)() {
    if m != nil {
        m.faceIdBlocked = value
    }
}
// SetFilterOpenInToOnlyManagedApps sets the filterOpenInToOnlyManagedApps property value. Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False.
func (m *IosManagedAppProtection) SetFilterOpenInToOnlyManagedApps(value *bool)() {
    if m != nil {
        m.filterOpenInToOnlyManagedApps = value
    }
}
// SetManagedUniversalLinks sets the managedUniversalLinks property value. A list of custom urls that are allowed to invocate a managed app
func (m *IosManagedAppProtection) SetManagedUniversalLinks(value []string)() {
    if m != nil {
        m.managedUniversalLinks = value
    }
}
// SetMinimumRequiredSdkVersion sets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) SetMinimumRequiredSdkVersion(value *string)() {
    if m != nil {
        m.minimumRequiredSdkVersion = value
    }
}
// SetMinimumWipeSdkVersion sets the minimumWipeSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) SetMinimumWipeSdkVersion(value *string)() {
    if m != nil {
        m.minimumWipeSdkVersion = value
    }
}
// SetProtectInboundDataFromUnknownSources sets the protectInboundDataFromUnknownSources property value. Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps.
func (m *IosManagedAppProtection) SetProtectInboundDataFromUnknownSources(value *bool)() {
    if m != nil {
        m.protectInboundDataFromUnknownSources = value
    }
}
// SetThirdPartyKeyboardsBlocked sets the thirdPartyKeyboardsBlocked property value. Defines if third party keyboards are allowed while accessing a managed app
func (m *IosManagedAppProtection) SetThirdPartyKeyboardsBlocked(value *bool)() {
    if m != nil {
        m.thirdPartyKeyboardsBlocked = value
    }
}
