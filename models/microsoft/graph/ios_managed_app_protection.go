package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new iosManagedAppProtection and sets the default values.
func NewIosManagedAppProtection()(*IosManagedAppProtection) {
    m := &IosManagedAppProtection{
        TargetedManagedAppProtection: *NewTargetedManagedAppProtection(),
    }
    return m
}
// Gets the allowedIosDeviceModels property value. Semicolon seperated list of device models allowed, as a string, for the managed app to work.
func (m *IosManagedAppProtection) GetAllowedIosDeviceModels()(*string) {
    if m == nil {
        return nil
    } else {
        return m.allowedIosDeviceModels
    }
}
// Gets the appActionIfIosDeviceModelNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device model is not allowed. Possible values are: block, wipe, warn.
func (m *IosManagedAppProtection) GetAppActionIfIosDeviceModelNotAllowed()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfIosDeviceModelNotAllowed
    }
}
// Gets the appDataEncryptionType property value. Type of encryption which should be used for data in a managed app. Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked.
func (m *IosManagedAppProtection) GetAppDataEncryptionType()(*ManagedAppDataEncryptionType) {
    if m == nil {
        return nil
    } else {
        return m.appDataEncryptionType
    }
}
// Gets the apps property value. List of apps to which the policy is deployed.
func (m *IosManagedAppProtection) GetApps()([]ManagedMobileApp) {
    if m == nil {
        return nil
    } else {
        return m.apps
    }
}
// Gets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *IosManagedAppProtection) GetCustomBrowserProtocol()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customBrowserProtocol
    }
}
// Gets the customDialerAppProtocol property value. Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
func (m *IosManagedAppProtection) GetCustomDialerAppProtocol()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customDialerAppProtocol
    }
}
// Gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *IosManagedAppProtection) GetDeployedAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deployedAppCount
    }
}
// Gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *IosManagedAppProtection) GetDeploymentSummary()(*ManagedAppPolicyDeploymentSummary) {
    if m == nil {
        return nil
    } else {
        return m.deploymentSummary
    }
}
// Gets the disableProtectionOfManagedOutboundOpenInData property value. Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps.
func (m *IosManagedAppProtection) GetDisableProtectionOfManagedOutboundOpenInData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableProtectionOfManagedOutboundOpenInData
    }
}
// Gets the exemptedAppProtocols property value. Apps in this list will be exempt from the policy and will be able to receive data from managed apps.
func (m *IosManagedAppProtection) GetExemptedAppProtocols()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.exemptedAppProtocols
    }
}
// Gets the exemptedUniversalLinks property value. A list of custom urls that are allowed to invocate an unmanaged app
func (m *IosManagedAppProtection) GetExemptedUniversalLinks()([]string) {
    if m == nil {
        return nil
    } else {
        return m.exemptedUniversalLinks
    }
}
// Gets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
func (m *IosManagedAppProtection) GetFaceIdBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.faceIdBlocked
    }
}
// Gets the filterOpenInToOnlyManagedApps property value. Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False.
func (m *IosManagedAppProtection) GetFilterOpenInToOnlyManagedApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.filterOpenInToOnlyManagedApps
    }
}
// Gets the managedUniversalLinks property value. A list of custom urls that are allowed to invocate a managed app
func (m *IosManagedAppProtection) GetManagedUniversalLinks()([]string) {
    if m == nil {
        return nil
    } else {
        return m.managedUniversalLinks
    }
}
// Gets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) GetMinimumRequiredSdkVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredSdkVersion
    }
}
// Gets the minimumWipeSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) GetMinimumWipeSdkVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWipeSdkVersion
    }
}
// Gets the protectInboundDataFromUnknownSources property value. Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps.
func (m *IosManagedAppProtection) GetProtectInboundDataFromUnknownSources()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.protectInboundDataFromUnknownSources
    }
}
// Gets the thirdPartyKeyboardsBlocked property value. Defines if third party keyboards are allowed while accessing a managed app
func (m *IosManagedAppProtection) GetThirdPartyKeyboardsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.thirdPartyKeyboardsBlocked
    }
}
// The deserialization information for the current model
func (m *IosManagedAppProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.TargetedManagedAppProtection.GetFieldDeserializers()
    res["allowedIosDeviceModels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAllowedIosDeviceModels(val)
        return nil
    }
    res["appActionIfIosDeviceModelNotAllowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppRemediationAction)
        m.SetAppActionIfIosDeviceModelNotAllowed(&cast)
        return nil
    }
    res["appDataEncryptionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataEncryptionType)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppDataEncryptionType)
        m.SetAppDataEncryptionType(&cast)
        return nil
    }
    res["apps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedMobileApp() })
        if err != nil {
            return err
        }
        res := make([]ManagedMobileApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedMobileApp))
        }
        m.SetApps(res)
        return nil
    }
    res["customBrowserProtocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomBrowserProtocol(val)
        return nil
    }
    res["customDialerAppProtocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomDialerAppProtocol(val)
        return nil
    }
    res["deployedAppCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeployedAppCount(val)
        return nil
    }
    res["deploymentSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicyDeploymentSummary() })
        if err != nil {
            return err
        }
        m.SetDeploymentSummary(val.(*ManagedAppPolicyDeploymentSummary))
        return nil
    }
    res["disableProtectionOfManagedOutboundOpenInData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDisableProtectionOfManagedOutboundOpenInData(val)
        return nil
    }
    res["exemptedAppProtocols"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValuePair))
        }
        m.SetExemptedAppProtocols(res)
        return nil
    }
    res["exemptedUniversalLinks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetExemptedUniversalLinks(res)
        return nil
    }
    res["faceIdBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFaceIdBlocked(val)
        return nil
    }
    res["filterOpenInToOnlyManagedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFilterOpenInToOnlyManagedApps(val)
        return nil
    }
    res["managedUniversalLinks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetManagedUniversalLinks(res)
        return nil
    }
    res["minimumRequiredSdkVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumRequiredSdkVersion(val)
        return nil
    }
    res["minimumWipeSdkVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumWipeSdkVersion(val)
        return nil
    }
    res["protectInboundDataFromUnknownSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetProtectInboundDataFromUnknownSources(val)
        return nil
    }
    res["thirdPartyKeyboardsBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetThirdPartyKeyboardsBlocked(val)
        return nil
    }
    return res
}
func (m *IosManagedAppProtection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        cast := m.GetAppActionIfIosDeviceModelNotAllowed().String()
        err = writer.WriteStringValue("appActionIfIosDeviceModelNotAllowed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppDataEncryptionType() != nil {
        cast := m.GetAppDataEncryptionType().String()
        err = writer.WriteStringValue("appDataEncryptionType", &cast)
        if err != nil {
            return err
        }
    }
    {
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
    {
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
    {
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
    {
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
// Sets the allowedIosDeviceModels property value. Semicolon seperated list of device models allowed, as a string, for the managed app to work.
// Parameters:
//  - value : Value to set for the allowedIosDeviceModels property.
func (m *IosManagedAppProtection) SetAllowedIosDeviceModels(value *string)() {
    m.allowedIosDeviceModels = value
}
// Sets the appActionIfIosDeviceModelNotAllowed property value. Defines a managed app behavior, either block or wipe, if the specified device model is not allowed. Possible values are: block, wipe, warn.
// Parameters:
//  - value : Value to set for the appActionIfIosDeviceModelNotAllowed property.
func (m *IosManagedAppProtection) SetAppActionIfIosDeviceModelNotAllowed(value *ManagedAppRemediationAction)() {
    m.appActionIfIosDeviceModelNotAllowed = value
}
// Sets the appDataEncryptionType property value. Type of encryption which should be used for data in a managed app. Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked.
// Parameters:
//  - value : Value to set for the appDataEncryptionType property.
func (m *IosManagedAppProtection) SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)() {
    m.appDataEncryptionType = value
}
// Sets the apps property value. List of apps to which the policy is deployed.
// Parameters:
//  - value : Value to set for the apps property.
func (m *IosManagedAppProtection) SetApps(value []ManagedMobileApp)() {
    m.apps = value
}
// Sets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
// Parameters:
//  - value : Value to set for the customBrowserProtocol property.
func (m *IosManagedAppProtection) SetCustomBrowserProtocol(value *string)() {
    m.customBrowserProtocol = value
}
// Sets the customDialerAppProtocol property value. Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
// Parameters:
//  - value : Value to set for the customDialerAppProtocol property.
func (m *IosManagedAppProtection) SetCustomDialerAppProtocol(value *string)() {
    m.customDialerAppProtocol = value
}
// Sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
// Parameters:
//  - value : Value to set for the deployedAppCount property.
func (m *IosManagedAppProtection) SetDeployedAppCount(value *int32)() {
    m.deployedAppCount = value
}
// Sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
// Parameters:
//  - value : Value to set for the deploymentSummary property.
func (m *IosManagedAppProtection) SetDeploymentSummary(value *ManagedAppPolicyDeploymentSummary)() {
    m.deploymentSummary = value
}
// Sets the disableProtectionOfManagedOutboundOpenInData property value. Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps.
// Parameters:
//  - value : Value to set for the disableProtectionOfManagedOutboundOpenInData property.
func (m *IosManagedAppProtection) SetDisableProtectionOfManagedOutboundOpenInData(value *bool)() {
    m.disableProtectionOfManagedOutboundOpenInData = value
}
// Sets the exemptedAppProtocols property value. Apps in this list will be exempt from the policy and will be able to receive data from managed apps.
// Parameters:
//  - value : Value to set for the exemptedAppProtocols property.
func (m *IosManagedAppProtection) SetExemptedAppProtocols(value []KeyValuePair)() {
    m.exemptedAppProtocols = value
}
// Sets the exemptedUniversalLinks property value. A list of custom urls that are allowed to invocate an unmanaged app
// Parameters:
//  - value : Value to set for the exemptedUniversalLinks property.
func (m *IosManagedAppProtection) SetExemptedUniversalLinks(value []string)() {
    m.exemptedUniversalLinks = value
}
// Sets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
// Parameters:
//  - value : Value to set for the faceIdBlocked property.
func (m *IosManagedAppProtection) SetFaceIdBlocked(value *bool)() {
    m.faceIdBlocked = value
}
// Sets the filterOpenInToOnlyManagedApps property value. Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False.
// Parameters:
//  - value : Value to set for the filterOpenInToOnlyManagedApps property.
func (m *IosManagedAppProtection) SetFilterOpenInToOnlyManagedApps(value *bool)() {
    m.filterOpenInToOnlyManagedApps = value
}
// Sets the managedUniversalLinks property value. A list of custom urls that are allowed to invocate a managed app
// Parameters:
//  - value : Value to set for the managedUniversalLinks property.
func (m *IosManagedAppProtection) SetManagedUniversalLinks(value []string)() {
    m.managedUniversalLinks = value
}
// Sets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
// Parameters:
//  - value : Value to set for the minimumRequiredSdkVersion property.
func (m *IosManagedAppProtection) SetMinimumRequiredSdkVersion(value *string)() {
    m.minimumRequiredSdkVersion = value
}
// Sets the minimumWipeSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
// Parameters:
//  - value : Value to set for the minimumWipeSdkVersion property.
func (m *IosManagedAppProtection) SetMinimumWipeSdkVersion(value *string)() {
    m.minimumWipeSdkVersion = value
}
// Sets the protectInboundDataFromUnknownSources property value. Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps.
// Parameters:
//  - value : Value to set for the protectInboundDataFromUnknownSources property.
func (m *IosManagedAppProtection) SetProtectInboundDataFromUnknownSources(value *bool)() {
    m.protectInboundDataFromUnknownSources = value
}
// Sets the thirdPartyKeyboardsBlocked property value. Defines if third party keyboards are allowed while accessing a managed app
// Parameters:
//  - value : Value to set for the thirdPartyKeyboardsBlocked property.
func (m *IosManagedAppProtection) SetThirdPartyKeyboardsBlocked(value *bool)() {
    m.thirdPartyKeyboardsBlocked = value
}
