package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedAppProtection 
type ManagedAppProtection struct {
    ManagedAppPolicy
    // Data storage locations where a user may store managed data.
    allowedDataIngestionLocations []ManagedAppDataIngestionLocation;
    // Data storage locations where a user may store managed data.
    allowedDataStorageLocations []ManagedAppDataStorageLocation;
    // Sources from which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
    allowedInboundDataTransferSources *ManagedAppDataTransferLevel;
    // Specify the number of characters that may be cut or copied from Org data and accounts to any application. This setting overrides the AllowedOutboundClipboardSharingLevel restriction. Default value of '0' means no exception is allowed.
    allowedOutboundClipboardSharingExceptionLength *int32;
    // The level to which the clipboard may be shared between apps on the managed device. Possible values are: allApps, managedAppsWithPasteIn, managedApps, blocked.
    allowedOutboundClipboardSharingLevel *ManagedAppClipboardSharingLevel;
    // Destinations to which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
    allowedOutboundDataTransferDestinations *ManagedAppDataTransferLevel;
    // Defines a managed app behavior, either block or wipe, when the device is either rooted or jailbroken, if DeviceComplianceRequired is set to true. Possible values are: block, wipe, warn.
    appActionIfDeviceComplianceRequired *ManagedAppRemediationAction;
    // Defines a managed app behavior, either block or wipe, based on maximum number of incorrect pin retry attempts. Possible values are: block, wipe, warn.
    appActionIfMaximumPinRetriesExceeded *ManagedAppRemediationAction;
    // If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Possible values are: block, wipe, warn.
    appActionIfUnableToAuthenticateUser *ManagedAppRemediationAction;
    // Indicates whether a user can bring data into org documents.
    blockDataIngestionIntoOrganizationDocuments *bool;
    // Indicates whether contacts can be synced to the user's device.
    contactSyncBlocked *bool;
    // Indicates whether the backup of a managed app's data is blocked.
    dataBackupBlocked *bool;
    // Indicates whether device compliance is required.
    deviceComplianceRequired *bool;
    // The classes of dialer apps that are allowed to click-to-open a phone number. Possible values are: allApps, managedApps, customApp, blocked.
    dialerRestrictionLevel *ManagedAppPhoneNumberRedirectLevel;
    // Indicates whether use of the app pin is required if the device pin is set.
    disableAppPinIfDevicePinIsSet *bool;
    // Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
    fingerprintBlocked *bool;
    // A grace period before blocking app access during off clock hours.
    gracePeriodToBlockAppsDuringOffClockHours *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // Indicates in which managed browser(s) that internet links should be opened. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. Possible values are: notConfigured, microsoftEdge.
    managedBrowser *ManagedBrowserType;
    // Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
    managedBrowserToOpenLinksRequired *bool;
    // Maximum allowed device threat level, as reported by the MTD app. Possible values are: notConfigured, secured, low, medium, high.
    maximumAllowedDeviceThreatLevel *ManagedAppDeviceThreatLevel;
    // Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
    maximumPinRetries *int32;
    // Versions bigger than the specified version will block the managed app from accessing company data.
    maximumRequiredOsVersion *string;
    // Versions bigger than the specified version will block the managed app from accessing company data.
    maximumWarningOsVersion *string;
    // Versions bigger than the specified version will block the managed app from accessing company data.
    maximumWipeOsVersion *string;
    // Minimum pin length required for an app-level pin if PinRequired is set to True
    minimumPinLength *int32;
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumRequiredAppVersion *string;
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumRequiredOsVersion *string;
    // Versions less than the specified version will result in warning message on the managed app.
    minimumWarningAppVersion *string;
    // Versions less than the specified version will result in warning message on the managed app from accessing company data.
    minimumWarningOsVersion *string;
    // Versions less than or equal to the specified version will wipe the managed app and the associated company data.
    minimumWipeAppVersion *string;
    // Versions less than or equal to the specified version will wipe the managed app and the associated company data.
    minimumWipeOsVersion *string;
    // Determines what action to take if the mobile threat defense threat threshold isn't met. Warn isn't a supported value for this property. Possible values are: block, wipe, warn.
    mobileThreatDefenseRemediationAction *ManagedAppRemediationAction;
    // Specify app notification restriction. Possible values are: allow, blockOrganizationalData, block.
    notificationRestriction *ManagedAppNotificationRestriction;
    // Indicates whether organizational credentials are required for app use.
    organizationalCredentialsRequired *bool;
    // TimePeriod before the all-level pin must be reset if PinRequired is set to True.
    periodBeforePinReset *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The period after which access is checked when the device is not connected to the internet.
    periodOfflineBeforeAccessCheck *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
    periodOfflineBeforeWipeIsEnforced *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The period after which access is checked when the device is connected to the internet.
    periodOnlineBeforeAccessCheck *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // Character set which may be used for an app-level pin if PinRequired is set to True. Possible values are: numeric, alphanumericAndSymbol.
    pinCharacterSet *ManagedAppPinCharacterSet;
    // Indicates whether an app-level pin is required.
    pinRequired *bool;
    // Timeout in minutes for an app pin instead of non biometrics passcode
    pinRequiredInsteadOfBiometricTimeout *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // Requires a pin to be unique from the number specified in this property.
    previousPinBlockCount *int32;
    // Indicates whether printing is allowed from managed apps.
    printBlocked *bool;
    // Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
    saveAsBlocked *bool;
    // Indicates whether simplePin is blocked.
    simplePinBlocked *bool;
}
// NewManagedAppProtection instantiates a new managedAppProtection and sets the default values.
func NewManagedAppProtection()(*ManagedAppProtection) {
    m := &ManagedAppProtection{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    return m
}
// GetAllowedDataIngestionLocations gets the allowedDataIngestionLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) GetAllowedDataIngestionLocations()([]ManagedAppDataIngestionLocation) {
    if m == nil {
        return nil
    } else {
        return m.allowedDataIngestionLocations
    }
}
// GetAllowedDataStorageLocations gets the allowedDataStorageLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) GetAllowedDataStorageLocations()([]ManagedAppDataStorageLocation) {
    if m == nil {
        return nil
    } else {
        return m.allowedDataStorageLocations
    }
}
// GetAllowedInboundDataTransferSources gets the allowedInboundDataTransferSources property value. Sources from which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
func (m *ManagedAppProtection) GetAllowedInboundDataTransferSources()(*ManagedAppDataTransferLevel) {
    if m == nil {
        return nil
    } else {
        return m.allowedInboundDataTransferSources
    }
}
// GetAllowedOutboundClipboardSharingExceptionLength gets the allowedOutboundClipboardSharingExceptionLength property value. Specify the number of characters that may be cut or copied from Org data and accounts to any application. This setting overrides the AllowedOutboundClipboardSharingLevel restriction. Default value of '0' means no exception is allowed.
func (m *ManagedAppProtection) GetAllowedOutboundClipboardSharingExceptionLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.allowedOutboundClipboardSharingExceptionLength
    }
}
// GetAllowedOutboundClipboardSharingLevel gets the allowedOutboundClipboardSharingLevel property value. The level to which the clipboard may be shared between apps on the managed device. Possible values are: allApps, managedAppsWithPasteIn, managedApps, blocked.
func (m *ManagedAppProtection) GetAllowedOutboundClipboardSharingLevel()(*ManagedAppClipboardSharingLevel) {
    if m == nil {
        return nil
    } else {
        return m.allowedOutboundClipboardSharingLevel
    }
}
// GetAllowedOutboundDataTransferDestinations gets the allowedOutboundDataTransferDestinations property value. Destinations to which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
func (m *ManagedAppProtection) GetAllowedOutboundDataTransferDestinations()(*ManagedAppDataTransferLevel) {
    if m == nil {
        return nil
    } else {
        return m.allowedOutboundDataTransferDestinations
    }
}
// GetAppActionIfDeviceComplianceRequired gets the appActionIfDeviceComplianceRequired property value. Defines a managed app behavior, either block or wipe, when the device is either rooted or jailbroken, if DeviceComplianceRequired is set to true. Possible values are: block, wipe, warn.
func (m *ManagedAppProtection) GetAppActionIfDeviceComplianceRequired()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfDeviceComplianceRequired
    }
}
// GetAppActionIfMaximumPinRetriesExceeded gets the appActionIfMaximumPinRetriesExceeded property value. Defines a managed app behavior, either block or wipe, based on maximum number of incorrect pin retry attempts. Possible values are: block, wipe, warn.
func (m *ManagedAppProtection) GetAppActionIfMaximumPinRetriesExceeded()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfMaximumPinRetriesExceeded
    }
}
// GetAppActionIfUnableToAuthenticateUser gets the appActionIfUnableToAuthenticateUser property value. If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Possible values are: block, wipe, warn.
func (m *ManagedAppProtection) GetAppActionIfUnableToAuthenticateUser()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.appActionIfUnableToAuthenticateUser
    }
}
// GetBlockDataIngestionIntoOrganizationDocuments gets the blockDataIngestionIntoOrganizationDocuments property value. Indicates whether a user can bring data into org documents.
func (m *ManagedAppProtection) GetBlockDataIngestionIntoOrganizationDocuments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockDataIngestionIntoOrganizationDocuments
    }
}
// GetContactSyncBlocked gets the contactSyncBlocked property value. Indicates whether contacts can be synced to the user's device.
func (m *ManagedAppProtection) GetContactSyncBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contactSyncBlocked
    }
}
// GetDataBackupBlocked gets the dataBackupBlocked property value. Indicates whether the backup of a managed app's data is blocked.
func (m *ManagedAppProtection) GetDataBackupBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dataBackupBlocked
    }
}
// GetDeviceComplianceRequired gets the deviceComplianceRequired property value. Indicates whether device compliance is required.
func (m *ManagedAppProtection) GetDeviceComplianceRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceRequired
    }
}
// GetDialerRestrictionLevel gets the dialerRestrictionLevel property value. The classes of dialer apps that are allowed to click-to-open a phone number. Possible values are: allApps, managedApps, customApp, blocked.
func (m *ManagedAppProtection) GetDialerRestrictionLevel()(*ManagedAppPhoneNumberRedirectLevel) {
    if m == nil {
        return nil
    } else {
        return m.dialerRestrictionLevel
    }
}
// GetDisableAppPinIfDevicePinIsSet gets the disableAppPinIfDevicePinIsSet property value. Indicates whether use of the app pin is required if the device pin is set.
func (m *ManagedAppProtection) GetDisableAppPinIfDevicePinIsSet()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableAppPinIfDevicePinIsSet
    }
}
// GetFingerprintBlocked gets the fingerprintBlocked property value. Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
func (m *ManagedAppProtection) GetFingerprintBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fingerprintBlocked
    }
}
// GetGracePeriodToBlockAppsDuringOffClockHours gets the gracePeriodToBlockAppsDuringOffClockHours property value. A grace period before blocking app access during off clock hours.
func (m *ManagedAppProtection) GetGracePeriodToBlockAppsDuringOffClockHours()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.gracePeriodToBlockAppsDuringOffClockHours
    }
}
// GetManagedBrowser gets the managedBrowser property value. Indicates in which managed browser(s) that internet links should be opened. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. Possible values are: notConfigured, microsoftEdge.
func (m *ManagedAppProtection) GetManagedBrowser()(*ManagedBrowserType) {
    if m == nil {
        return nil
    } else {
        return m.managedBrowser
    }
}
// GetManagedBrowserToOpenLinksRequired gets the managedBrowserToOpenLinksRequired property value. Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
func (m *ManagedAppProtection) GetManagedBrowserToOpenLinksRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.managedBrowserToOpenLinksRequired
    }
}
// GetMaximumAllowedDeviceThreatLevel gets the maximumAllowedDeviceThreatLevel property value. Maximum allowed device threat level, as reported by the MTD app. Possible values are: notConfigured, secured, low, medium, high.
func (m *ManagedAppProtection) GetMaximumAllowedDeviceThreatLevel()(*ManagedAppDeviceThreatLevel) {
    if m == nil {
        return nil
    } else {
        return m.maximumAllowedDeviceThreatLevel
    }
}
// GetMaximumPinRetries gets the maximumPinRetries property value. Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
func (m *ManagedAppProtection) GetMaximumPinRetries()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumPinRetries
    }
}
// GetMaximumRequiredOsVersion gets the maximumRequiredOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMaximumRequiredOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumRequiredOsVersion
    }
}
// GetMaximumWarningOsVersion gets the maximumWarningOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMaximumWarningOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumWarningOsVersion
    }
}
// GetMaximumWipeOsVersion gets the maximumWipeOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMaximumWipeOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maximumWipeOsVersion
    }
}
// GetMinimumPinLength gets the minimumPinLength property value. Minimum pin length required for an app-level pin if PinRequired is set to True
func (m *ManagedAppProtection) GetMinimumPinLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumPinLength
    }
}
// GetMinimumRequiredAppVersion gets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumRequiredAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredAppVersion
    }
}
// GetMinimumRequiredOsVersion gets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumRequiredOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredOsVersion
    }
}
// GetMinimumWarningAppVersion gets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app.
func (m *ManagedAppProtection) GetMinimumWarningAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningAppVersion
    }
}
// GetMinimumWarningOsVersion gets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumWarningOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningOsVersion
    }
}
// GetMinimumWipeAppVersion gets the minimumWipeAppVersion property value. Versions less than or equal to the specified version will wipe the managed app and the associated company data.
func (m *ManagedAppProtection) GetMinimumWipeAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWipeAppVersion
    }
}
// GetMinimumWipeOsVersion gets the minimumWipeOsVersion property value. Versions less than or equal to the specified version will wipe the managed app and the associated company data.
func (m *ManagedAppProtection) GetMinimumWipeOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWipeOsVersion
    }
}
// GetMobileThreatDefenseRemediationAction gets the mobileThreatDefenseRemediationAction property value. Determines what action to take if the mobile threat defense threat threshold isn't met. Warn isn't a supported value for this property. Possible values are: block, wipe, warn.
func (m *ManagedAppProtection) GetMobileThreatDefenseRemediationAction()(*ManagedAppRemediationAction) {
    if m == nil {
        return nil
    } else {
        return m.mobileThreatDefenseRemediationAction
    }
}
// GetNotificationRestriction gets the notificationRestriction property value. Specify app notification restriction. Possible values are: allow, blockOrganizationalData, block.
func (m *ManagedAppProtection) GetNotificationRestriction()(*ManagedAppNotificationRestriction) {
    if m == nil {
        return nil
    } else {
        return m.notificationRestriction
    }
}
// GetOrganizationalCredentialsRequired gets the organizationalCredentialsRequired property value. Indicates whether organizational credentials are required for app use.
func (m *ManagedAppProtection) GetOrganizationalCredentialsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.organizationalCredentialsRequired
    }
}
// GetPeriodBeforePinReset gets the periodBeforePinReset property value. TimePeriod before the all-level pin must be reset if PinRequired is set to True.
func (m *ManagedAppProtection) GetPeriodBeforePinReset()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.periodBeforePinReset
    }
}
// GetPeriodOfflineBeforeAccessCheck gets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet.
func (m *ManagedAppProtection) GetPeriodOfflineBeforeAccessCheck()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.periodOfflineBeforeAccessCheck
    }
}
// GetPeriodOfflineBeforeWipeIsEnforced gets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
func (m *ManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.periodOfflineBeforeWipeIsEnforced
    }
}
// GetPeriodOnlineBeforeAccessCheck gets the periodOnlineBeforeAccessCheck property value. The period after which access is checked when the device is connected to the internet.
func (m *ManagedAppProtection) GetPeriodOnlineBeforeAccessCheck()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.periodOnlineBeforeAccessCheck
    }
}
// GetPinCharacterSet gets the pinCharacterSet property value. Character set which may be used for an app-level pin if PinRequired is set to True. Possible values are: numeric, alphanumericAndSymbol.
func (m *ManagedAppProtection) GetPinCharacterSet()(*ManagedAppPinCharacterSet) {
    if m == nil {
        return nil
    } else {
        return m.pinCharacterSet
    }
}
// GetPinRequired gets the pinRequired property value. Indicates whether an app-level pin is required.
func (m *ManagedAppProtection) GetPinRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.pinRequired
    }
}
// GetPinRequiredInsteadOfBiometricTimeout gets the pinRequiredInsteadOfBiometricTimeout property value. Timeout in minutes for an app pin instead of non biometrics passcode
func (m *ManagedAppProtection) GetPinRequiredInsteadOfBiometricTimeout()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.pinRequiredInsteadOfBiometricTimeout
    }
}
// GetPreviousPinBlockCount gets the previousPinBlockCount property value. Requires a pin to be unique from the number specified in this property.
func (m *ManagedAppProtection) GetPreviousPinBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.previousPinBlockCount
    }
}
// GetPrintBlocked gets the printBlocked property value. Indicates whether printing is allowed from managed apps.
func (m *ManagedAppProtection) GetPrintBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.printBlocked
    }
}
// GetSaveAsBlocked gets the saveAsBlocked property value. Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
func (m *ManagedAppProtection) GetSaveAsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.saveAsBlocked
    }
}
// GetSimplePinBlocked gets the simplePinBlocked property value. Indicates whether simplePin is blocked.
func (m *ManagedAppProtection) GetSimplePinBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.simplePinBlocked
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagedAppPolicy.GetFieldDeserializers()
    res["allowedDataIngestionLocations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseManagedAppDataIngestionLocation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppDataIngestionLocation, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppDataIngestionLocation))
            }
            m.SetAllowedDataIngestionLocations(res)
        }
        return nil
    }
    res["allowedDataStorageLocations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseManagedAppDataStorageLocation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppDataStorageLocation, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppDataStorageLocation))
            }
            m.SetAllowedDataStorageLocations(res)
        }
        return nil
    }
    res["allowedInboundDataTransferSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedInboundDataTransferSources(val.(*ManagedAppDataTransferLevel))
        }
        return nil
    }
    res["allowedOutboundClipboardSharingExceptionLength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundClipboardSharingExceptionLength(val)
        }
        return nil
    }
    res["allowedOutboundClipboardSharingLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppClipboardSharingLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundClipboardSharingLevel(val.(*ManagedAppClipboardSharingLevel))
        }
        return nil
    }
    res["allowedOutboundDataTransferDestinations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundDataTransferDestinations(val.(*ManagedAppDataTransferLevel))
        }
        return nil
    }
    res["appActionIfDeviceComplianceRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfDeviceComplianceRequired(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfMaximumPinRetriesExceeded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfMaximumPinRetriesExceeded(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfUnableToAuthenticateUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfUnableToAuthenticateUser(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["blockDataIngestionIntoOrganizationDocuments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockDataIngestionIntoOrganizationDocuments(val)
        }
        return nil
    }
    res["contactSyncBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactSyncBlocked(val)
        }
        return nil
    }
    res["dataBackupBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataBackupBlocked(val)
        }
        return nil
    }
    res["deviceComplianceRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceRequired(val)
        }
        return nil
    }
    res["dialerRestrictionLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppPhoneNumberRedirectLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDialerRestrictionLevel(val.(*ManagedAppPhoneNumberRedirectLevel))
        }
        return nil
    }
    res["disableAppPinIfDevicePinIsSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAppPinIfDevicePinIsSet(val)
        }
        return nil
    }
    res["fingerprintBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprintBlocked(val)
        }
        return nil
    }
    res["gracePeriodToBlockAppsDuringOffClockHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodToBlockAppsDuringOffClockHours(val)
        }
        return nil
    }
    res["managedBrowser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedBrowserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBrowser(val.(*ManagedBrowserType))
        }
        return nil
    }
    res["managedBrowserToOpenLinksRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBrowserToOpenLinksRequired(val)
        }
        return nil
    }
    res["maximumAllowedDeviceThreatLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDeviceThreatLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAllowedDeviceThreatLevel(val.(*ManagedAppDeviceThreatLevel))
        }
        return nil
    }
    res["maximumPinRetries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumPinRetries(val)
        }
        return nil
    }
    res["maximumRequiredOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumRequiredOsVersion(val)
        }
        return nil
    }
    res["maximumWarningOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumWarningOsVersion(val)
        }
        return nil
    }
    res["maximumWipeOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumWipeOsVersion(val)
        }
        return nil
    }
    res["minimumPinLength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumPinLength(val)
        }
        return nil
    }
    res["minimumRequiredAppVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredAppVersion(val)
        }
        return nil
    }
    res["minimumRequiredOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredOsVersion(val)
        }
        return nil
    }
    res["minimumWarningAppVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningAppVersion(val)
        }
        return nil
    }
    res["minimumWarningOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningOsVersion(val)
        }
        return nil
    }
    res["minimumWipeAppVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipeAppVersion(val)
        }
        return nil
    }
    res["minimumWipeOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipeOsVersion(val)
        }
        return nil
    }
    res["mobileThreatDefenseRemediationAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileThreatDefenseRemediationAction(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["notificationRestriction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppNotificationRestriction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationRestriction(val.(*ManagedAppNotificationRestriction))
        }
        return nil
    }
    res["organizationalCredentialsRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalCredentialsRequired(val)
        }
        return nil
    }
    res["periodBeforePinReset"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodBeforePinReset(val)
        }
        return nil
    }
    res["periodOfflineBeforeAccessCheck"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodOfflineBeforeAccessCheck(val)
        }
        return nil
    }
    res["periodOfflineBeforeWipeIsEnforced"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodOfflineBeforeWipeIsEnforced(val)
        }
        return nil
    }
    res["periodOnlineBeforeAccessCheck"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodOnlineBeforeAccessCheck(val)
        }
        return nil
    }
    res["pinCharacterSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppPinCharacterSet)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinCharacterSet(val.(*ManagedAppPinCharacterSet))
        }
        return nil
    }
    res["pinRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinRequired(val)
        }
        return nil
    }
    res["pinRequiredInsteadOfBiometricTimeout"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinRequiredInsteadOfBiometricTimeout(val)
        }
        return nil
    }
    res["previousPinBlockCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousPinBlockCount(val)
        }
        return nil
    }
    res["printBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrintBlocked(val)
        }
        return nil
    }
    res["saveAsBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSaveAsBlocked(val)
        }
        return nil
    }
    res["simplePinBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimplePinBlocked(val)
        }
        return nil
    }
    return res
}
func (m *ManagedAppProtection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedAppProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagedAppPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedDataIngestionLocations() != nil {
        err = writer.WriteCollectionOfStringValues("allowedDataIngestionLocations", SerializeManagedAppDataIngestionLocation(m.GetAllowedDataIngestionLocations()))
        if err != nil {
            return err
        }
    }
    if m.GetAllowedDataStorageLocations() != nil {
        err = writer.WriteCollectionOfStringValues("allowedDataStorageLocations", SerializeManagedAppDataStorageLocation(m.GetAllowedDataStorageLocations()))
        if err != nil {
            return err
        }
    }
    if m.GetAllowedInboundDataTransferSources() != nil {
        cast := (*m.GetAllowedInboundDataTransferSources()).String()
        err = writer.WriteStringValue("allowedInboundDataTransferSources", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("allowedOutboundClipboardSharingExceptionLength", m.GetAllowedOutboundClipboardSharingExceptionLength())
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
    if m.GetAppActionIfDeviceComplianceRequired() != nil {
        cast := (*m.GetAppActionIfDeviceComplianceRequired()).String()
        err = writer.WriteStringValue("appActionIfDeviceComplianceRequired", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfMaximumPinRetriesExceeded() != nil {
        cast := (*m.GetAppActionIfMaximumPinRetriesExceeded()).String()
        err = writer.WriteStringValue("appActionIfMaximumPinRetriesExceeded", &cast)
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
    {
        err = writer.WriteBoolValue("blockDataIngestionIntoOrganizationDocuments", m.GetBlockDataIngestionIntoOrganizationDocuments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contactSyncBlocked", m.GetContactSyncBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dataBackupBlocked", m.GetDataBackupBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceComplianceRequired", m.GetDeviceComplianceRequired())
        if err != nil {
            return err
        }
    }
    if m.GetDialerRestrictionLevel() != nil {
        cast := (*m.GetDialerRestrictionLevel()).String()
        err = writer.WriteStringValue("dialerRestrictionLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableAppPinIfDevicePinIsSet", m.GetDisableAppPinIfDevicePinIsSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fingerprintBlocked", m.GetFingerprintBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("gracePeriodToBlockAppsDuringOffClockHours", m.GetGracePeriodToBlockAppsDuringOffClockHours())
        if err != nil {
            return err
        }
    }
    if m.GetManagedBrowser() != nil {
        cast := (*m.GetManagedBrowser()).String()
        err = writer.WriteStringValue("managedBrowser", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("managedBrowserToOpenLinksRequired", m.GetManagedBrowserToOpenLinksRequired())
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
        err = writer.WriteInt32Value("maximumPinRetries", m.GetMaximumPinRetries())
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
        err = writer.WriteInt32Value("minimumPinLength", m.GetMinimumPinLength())
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
    if m.GetMobileThreatDefenseRemediationAction() != nil {
        cast := (*m.GetMobileThreatDefenseRemediationAction()).String()
        err = writer.WriteStringValue("mobileThreatDefenseRemediationAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotificationRestriction() != nil {
        cast := (*m.GetNotificationRestriction()).String()
        err = writer.WriteStringValue("notificationRestriction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("organizationalCredentialsRequired", m.GetOrganizationalCredentialsRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("periodBeforePinReset", m.GetPeriodBeforePinReset())
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
        err = writer.WriteISODurationValue("periodOnlineBeforeAccessCheck", m.GetPeriodOnlineBeforeAccessCheck())
        if err != nil {
            return err
        }
    }
    if m.GetPinCharacterSet() != nil {
        cast := (*m.GetPinCharacterSet()).String()
        err = writer.WriteStringValue("pinCharacterSet", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("pinRequired", m.GetPinRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("pinRequiredInsteadOfBiometricTimeout", m.GetPinRequiredInsteadOfBiometricTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("previousPinBlockCount", m.GetPreviousPinBlockCount())
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
    {
        err = writer.WriteBoolValue("saveAsBlocked", m.GetSaveAsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("simplePinBlocked", m.GetSimplePinBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedDataIngestionLocations sets the allowedDataIngestionLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) SetAllowedDataIngestionLocations(value []ManagedAppDataIngestionLocation)() {
    if m != nil {
        m.allowedDataIngestionLocations = value
    }
}
// SetAllowedDataStorageLocations sets the allowedDataStorageLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) SetAllowedDataStorageLocations(value []ManagedAppDataStorageLocation)() {
    if m != nil {
        m.allowedDataStorageLocations = value
    }
}
// SetAllowedInboundDataTransferSources sets the allowedInboundDataTransferSources property value. Sources from which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
func (m *ManagedAppProtection) SetAllowedInboundDataTransferSources(value *ManagedAppDataTransferLevel)() {
    if m != nil {
        m.allowedInboundDataTransferSources = value
    }
}
// SetAllowedOutboundClipboardSharingExceptionLength sets the allowedOutboundClipboardSharingExceptionLength property value. Specify the number of characters that may be cut or copied from Org data and accounts to any application. This setting overrides the AllowedOutboundClipboardSharingLevel restriction. Default value of '0' means no exception is allowed.
func (m *ManagedAppProtection) SetAllowedOutboundClipboardSharingExceptionLength(value *int32)() {
    if m != nil {
        m.allowedOutboundClipboardSharingExceptionLength = value
    }
}
// SetAllowedOutboundClipboardSharingLevel sets the allowedOutboundClipboardSharingLevel property value. The level to which the clipboard may be shared between apps on the managed device. Possible values are: allApps, managedAppsWithPasteIn, managedApps, blocked.
func (m *ManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(value *ManagedAppClipboardSharingLevel)() {
    if m != nil {
        m.allowedOutboundClipboardSharingLevel = value
    }
}
// SetAllowedOutboundDataTransferDestinations sets the allowedOutboundDataTransferDestinations property value. Destinations to which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
func (m *ManagedAppProtection) SetAllowedOutboundDataTransferDestinations(value *ManagedAppDataTransferLevel)() {
    if m != nil {
        m.allowedOutboundDataTransferDestinations = value
    }
}
// SetAppActionIfDeviceComplianceRequired sets the appActionIfDeviceComplianceRequired property value. Defines a managed app behavior, either block or wipe, when the device is either rooted or jailbroken, if DeviceComplianceRequired is set to true. Possible values are: block, wipe, warn.
func (m *ManagedAppProtection) SetAppActionIfDeviceComplianceRequired(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfDeviceComplianceRequired = value
    }
}
// SetAppActionIfMaximumPinRetriesExceeded sets the appActionIfMaximumPinRetriesExceeded property value. Defines a managed app behavior, either block or wipe, based on maximum number of incorrect pin retry attempts. Possible values are: block, wipe, warn.
func (m *ManagedAppProtection) SetAppActionIfMaximumPinRetriesExceeded(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfMaximumPinRetriesExceeded = value
    }
}
// SetAppActionIfUnableToAuthenticateUser sets the appActionIfUnableToAuthenticateUser property value. If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Possible values are: block, wipe, warn.
func (m *ManagedAppProtection) SetAppActionIfUnableToAuthenticateUser(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.appActionIfUnableToAuthenticateUser = value
    }
}
// SetBlockDataIngestionIntoOrganizationDocuments sets the blockDataIngestionIntoOrganizationDocuments property value. Indicates whether a user can bring data into org documents.
func (m *ManagedAppProtection) SetBlockDataIngestionIntoOrganizationDocuments(value *bool)() {
    if m != nil {
        m.blockDataIngestionIntoOrganizationDocuments = value
    }
}
// SetContactSyncBlocked sets the contactSyncBlocked property value. Indicates whether contacts can be synced to the user's device.
func (m *ManagedAppProtection) SetContactSyncBlocked(value *bool)() {
    if m != nil {
        m.contactSyncBlocked = value
    }
}
// SetDataBackupBlocked sets the dataBackupBlocked property value. Indicates whether the backup of a managed app's data is blocked.
func (m *ManagedAppProtection) SetDataBackupBlocked(value *bool)() {
    if m != nil {
        m.dataBackupBlocked = value
    }
}
// SetDeviceComplianceRequired sets the deviceComplianceRequired property value. Indicates whether device compliance is required.
func (m *ManagedAppProtection) SetDeviceComplianceRequired(value *bool)() {
    if m != nil {
        m.deviceComplianceRequired = value
    }
}
// SetDialerRestrictionLevel sets the dialerRestrictionLevel property value. The classes of dialer apps that are allowed to click-to-open a phone number. Possible values are: allApps, managedApps, customApp, blocked.
func (m *ManagedAppProtection) SetDialerRestrictionLevel(value *ManagedAppPhoneNumberRedirectLevel)() {
    if m != nil {
        m.dialerRestrictionLevel = value
    }
}
// SetDisableAppPinIfDevicePinIsSet sets the disableAppPinIfDevicePinIsSet property value. Indicates whether use of the app pin is required if the device pin is set.
func (m *ManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(value *bool)() {
    if m != nil {
        m.disableAppPinIfDevicePinIsSet = value
    }
}
// SetFingerprintBlocked sets the fingerprintBlocked property value. Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
func (m *ManagedAppProtection) SetFingerprintBlocked(value *bool)() {
    if m != nil {
        m.fingerprintBlocked = value
    }
}
// SetGracePeriodToBlockAppsDuringOffClockHours sets the gracePeriodToBlockAppsDuringOffClockHours property value. A grace period before blocking app access during off clock hours.
func (m *ManagedAppProtection) SetGracePeriodToBlockAppsDuringOffClockHours(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.gracePeriodToBlockAppsDuringOffClockHours = value
    }
}
// SetManagedBrowser sets the managedBrowser property value. Indicates in which managed browser(s) that internet links should be opened. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. Possible values are: notConfigured, microsoftEdge.
func (m *ManagedAppProtection) SetManagedBrowser(value *ManagedBrowserType)() {
    if m != nil {
        m.managedBrowser = value
    }
}
// SetManagedBrowserToOpenLinksRequired sets the managedBrowserToOpenLinksRequired property value. Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
func (m *ManagedAppProtection) SetManagedBrowserToOpenLinksRequired(value *bool)() {
    if m != nil {
        m.managedBrowserToOpenLinksRequired = value
    }
}
// SetMaximumAllowedDeviceThreatLevel sets the maximumAllowedDeviceThreatLevel property value. Maximum allowed device threat level, as reported by the MTD app. Possible values are: notConfigured, secured, low, medium, high.
func (m *ManagedAppProtection) SetMaximumAllowedDeviceThreatLevel(value *ManagedAppDeviceThreatLevel)() {
    if m != nil {
        m.maximumAllowedDeviceThreatLevel = value
    }
}
// SetMaximumPinRetries sets the maximumPinRetries property value. Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
func (m *ManagedAppProtection) SetMaximumPinRetries(value *int32)() {
    if m != nil {
        m.maximumPinRetries = value
    }
}
// SetMaximumRequiredOsVersion sets the maximumRequiredOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMaximumRequiredOsVersion(value *string)() {
    if m != nil {
        m.maximumRequiredOsVersion = value
    }
}
// SetMaximumWarningOsVersion sets the maximumWarningOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMaximumWarningOsVersion(value *string)() {
    if m != nil {
        m.maximumWarningOsVersion = value
    }
}
// SetMaximumWipeOsVersion sets the maximumWipeOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMaximumWipeOsVersion(value *string)() {
    if m != nil {
        m.maximumWipeOsVersion = value
    }
}
// SetMinimumPinLength sets the minimumPinLength property value. Minimum pin length required for an app-level pin if PinRequired is set to True
func (m *ManagedAppProtection) SetMinimumPinLength(value *int32)() {
    if m != nil {
        m.minimumPinLength = value
    }
}
// SetMinimumRequiredAppVersion sets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMinimumRequiredAppVersion(value *string)() {
    if m != nil {
        m.minimumRequiredAppVersion = value
    }
}
// SetMinimumRequiredOsVersion sets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMinimumRequiredOsVersion(value *string)() {
    if m != nil {
        m.minimumRequiredOsVersion = value
    }
}
// SetMinimumWarningAppVersion sets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app.
func (m *ManagedAppProtection) SetMinimumWarningAppVersion(value *string)() {
    if m != nil {
        m.minimumWarningAppVersion = value
    }
}
// SetMinimumWarningOsVersion sets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
func (m *ManagedAppProtection) SetMinimumWarningOsVersion(value *string)() {
    if m != nil {
        m.minimumWarningOsVersion = value
    }
}
// SetMinimumWipeAppVersion sets the minimumWipeAppVersion property value. Versions less than or equal to the specified version will wipe the managed app and the associated company data.
func (m *ManagedAppProtection) SetMinimumWipeAppVersion(value *string)() {
    if m != nil {
        m.minimumWipeAppVersion = value
    }
}
// SetMinimumWipeOsVersion sets the minimumWipeOsVersion property value. Versions less than or equal to the specified version will wipe the managed app and the associated company data.
func (m *ManagedAppProtection) SetMinimumWipeOsVersion(value *string)() {
    if m != nil {
        m.minimumWipeOsVersion = value
    }
}
// SetMobileThreatDefenseRemediationAction sets the mobileThreatDefenseRemediationAction property value. Determines what action to take if the mobile threat defense threat threshold isn't met. Warn isn't a supported value for this property. Possible values are: block, wipe, warn.
func (m *ManagedAppProtection) SetMobileThreatDefenseRemediationAction(value *ManagedAppRemediationAction)() {
    if m != nil {
        m.mobileThreatDefenseRemediationAction = value
    }
}
// SetNotificationRestriction sets the notificationRestriction property value. Specify app notification restriction. Possible values are: allow, blockOrganizationalData, block.
func (m *ManagedAppProtection) SetNotificationRestriction(value *ManagedAppNotificationRestriction)() {
    if m != nil {
        m.notificationRestriction = value
    }
}
// SetOrganizationalCredentialsRequired sets the organizationalCredentialsRequired property value. Indicates whether organizational credentials are required for app use.
func (m *ManagedAppProtection) SetOrganizationalCredentialsRequired(value *bool)() {
    if m != nil {
        m.organizationalCredentialsRequired = value
    }
}
// SetPeriodBeforePinReset sets the periodBeforePinReset property value. TimePeriod before the all-level pin must be reset if PinRequired is set to True.
func (m *ManagedAppProtection) SetPeriodBeforePinReset(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.periodBeforePinReset = value
    }
}
// SetPeriodOfflineBeforeAccessCheck sets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet.
func (m *ManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.periodOfflineBeforeAccessCheck = value
    }
}
// SetPeriodOfflineBeforeWipeIsEnforced sets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
func (m *ManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.periodOfflineBeforeWipeIsEnforced = value
    }
}
// SetPeriodOnlineBeforeAccessCheck sets the periodOnlineBeforeAccessCheck property value. The period after which access is checked when the device is connected to the internet.
func (m *ManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.periodOnlineBeforeAccessCheck = value
    }
}
// SetPinCharacterSet sets the pinCharacterSet property value. Character set which may be used for an app-level pin if PinRequired is set to True. Possible values are: numeric, alphanumericAndSymbol.
func (m *ManagedAppProtection) SetPinCharacterSet(value *ManagedAppPinCharacterSet)() {
    if m != nil {
        m.pinCharacterSet = value
    }
}
// SetPinRequired sets the pinRequired property value. Indicates whether an app-level pin is required.
func (m *ManagedAppProtection) SetPinRequired(value *bool)() {
    if m != nil {
        m.pinRequired = value
    }
}
// SetPinRequiredInsteadOfBiometricTimeout sets the pinRequiredInsteadOfBiometricTimeout property value. Timeout in minutes for an app pin instead of non biometrics passcode
func (m *ManagedAppProtection) SetPinRequiredInsteadOfBiometricTimeout(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.pinRequiredInsteadOfBiometricTimeout = value
    }
}
// SetPreviousPinBlockCount sets the previousPinBlockCount property value. Requires a pin to be unique from the number specified in this property.
func (m *ManagedAppProtection) SetPreviousPinBlockCount(value *int32)() {
    if m != nil {
        m.previousPinBlockCount = value
    }
}
// SetPrintBlocked sets the printBlocked property value. Indicates whether printing is allowed from managed apps.
func (m *ManagedAppProtection) SetPrintBlocked(value *bool)() {
    if m != nil {
        m.printBlocked = value
    }
}
// SetSaveAsBlocked sets the saveAsBlocked property value. Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
func (m *ManagedAppProtection) SetSaveAsBlocked(value *bool)() {
    if m != nil {
        m.saveAsBlocked = value
    }
}
// SetSimplePinBlocked sets the simplePinBlocked property value. Indicates whether simplePin is blocked.
func (m *ManagedAppProtection) SetSimplePinBlocked(value *bool)() {
    if m != nil {
        m.simplePinBlocked = value
    }
}
