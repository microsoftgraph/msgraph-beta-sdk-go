package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppProtection policy used to configure detailed management settings for a specified set of apps
type ManagedAppProtection struct {
    ManagedAppPolicy
}
// NewManagedAppProtection instantiates a new managedAppProtection and sets the default values.
func NewManagedAppProtection()(*ManagedAppProtection) {
    m := &ManagedAppProtection{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    odataTypeValue := "#microsoft.graph.managedAppProtection"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.androidManagedAppProtection":
                        return NewAndroidManagedAppProtection(), nil
                    case "#microsoft.graph.defaultManagedAppProtection":
                        return NewDefaultManagedAppProtection(), nil
                    case "#microsoft.graph.iosManagedAppProtection":
                        return NewIosManagedAppProtection(), nil
                    case "#microsoft.graph.targetedManagedAppProtection":
                        return NewTargetedManagedAppProtection(), nil
                }
            }
        }
    }
    return NewManagedAppProtection(), nil
}
// GetAllowedDataIngestionLocations gets the allowedDataIngestionLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) GetAllowedDataIngestionLocations()([]ManagedAppDataIngestionLocation) {
    val, err := m.GetBackingStore().Get("allowedDataIngestionLocations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedAppDataIngestionLocation)
    }
    return nil
}
// GetAllowedDataStorageLocations gets the allowedDataStorageLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) GetAllowedDataStorageLocations()([]ManagedAppDataStorageLocation) {
    val, err := m.GetBackingStore().Get("allowedDataStorageLocations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedAppDataStorageLocation)
    }
    return nil
}
// GetAllowedInboundDataTransferSources gets the allowedInboundDataTransferSources property value. Data can be transferred from/to these classes of apps
func (m *ManagedAppProtection) GetAllowedInboundDataTransferSources()(*ManagedAppDataTransferLevel) {
    val, err := m.GetBackingStore().Get("allowedInboundDataTransferSources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppDataTransferLevel)
    }
    return nil
}
// GetAllowedOutboundClipboardSharingExceptionLength gets the allowedOutboundClipboardSharingExceptionLength property value. Specify the number of characters that may be cut or copied from Org data and accounts to any application. This setting overrides the AllowedOutboundClipboardSharingLevel restriction. Default value of '0' means no exception is allowed.
func (m *ManagedAppProtection) GetAllowedOutboundClipboardSharingExceptionLength()(*int32) {
    val, err := m.GetBackingStore().Get("allowedOutboundClipboardSharingExceptionLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAllowedOutboundClipboardSharingLevel gets the allowedOutboundClipboardSharingLevel property value. Represents the level to which the device's clipboard may be shared between apps
func (m *ManagedAppProtection) GetAllowedOutboundClipboardSharingLevel()(*ManagedAppClipboardSharingLevel) {
    val, err := m.GetBackingStore().Get("allowedOutboundClipboardSharingLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppClipboardSharingLevel)
    }
    return nil
}
// GetAllowedOutboundDataTransferDestinations gets the allowedOutboundDataTransferDestinations property value. Data can be transferred from/to these classes of apps
func (m *ManagedAppProtection) GetAllowedOutboundDataTransferDestinations()(*ManagedAppDataTransferLevel) {
    val, err := m.GetBackingStore().Get("allowedOutboundDataTransferDestinations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppDataTransferLevel)
    }
    return nil
}
// GetAppActionIfDeviceComplianceRequired gets the appActionIfDeviceComplianceRequired property value. An admin initiated action to be applied on a managed app.
func (m *ManagedAppProtection) GetAppActionIfDeviceComplianceRequired()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfDeviceComplianceRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppActionIfMaximumPinRetriesExceeded gets the appActionIfMaximumPinRetriesExceeded property value. An admin initiated action to be applied on a managed app.
func (m *ManagedAppProtection) GetAppActionIfMaximumPinRetriesExceeded()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfMaximumPinRetriesExceeded")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetAppActionIfUnableToAuthenticateUser gets the appActionIfUnableToAuthenticateUser property value. If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Possible values are: block, wipe, warn.
func (m *ManagedAppProtection) GetAppActionIfUnableToAuthenticateUser()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("appActionIfUnableToAuthenticateUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetBlockDataIngestionIntoOrganizationDocuments gets the blockDataIngestionIntoOrganizationDocuments property value. Indicates whether a user can bring data into org documents.
func (m *ManagedAppProtection) GetBlockDataIngestionIntoOrganizationDocuments()(*bool) {
    val, err := m.GetBackingStore().Get("blockDataIngestionIntoOrganizationDocuments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetContactSyncBlocked gets the contactSyncBlocked property value. Indicates whether contacts can be synced to the user's device.
func (m *ManagedAppProtection) GetContactSyncBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("contactSyncBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDataBackupBlocked gets the dataBackupBlocked property value. Indicates whether the backup of a managed app's data is blocked.
func (m *ManagedAppProtection) GetDataBackupBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("dataBackupBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeviceComplianceRequired gets the deviceComplianceRequired property value. Indicates whether device compliance is required.
func (m *ManagedAppProtection) GetDeviceComplianceRequired()(*bool) {
    val, err := m.GetBackingStore().Get("deviceComplianceRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDialerRestrictionLevel gets the dialerRestrictionLevel property value. The classes of apps that are allowed to click-to-open a phone number, for making phone calls or sending text messages.
func (m *ManagedAppProtection) GetDialerRestrictionLevel()(*ManagedAppPhoneNumberRedirectLevel) {
    val, err := m.GetBackingStore().Get("dialerRestrictionLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppPhoneNumberRedirectLevel)
    }
    return nil
}
// GetDisableAppPinIfDevicePinIsSet gets the disableAppPinIfDevicePinIsSet property value. Indicates whether use of the app pin is required if the device pin is set.
func (m *ManagedAppProtection) GetDisableAppPinIfDevicePinIsSet()(*bool) {
    val, err := m.GetBackingStore().Get("disableAppPinIfDevicePinIsSet")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedAppPolicy.GetFieldDeserializers()
    res["allowedDataIngestionLocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseManagedAppDataIngestionLocation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppDataIngestionLocation, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*ManagedAppDataIngestionLocation))
                }
            }
            m.SetAllowedDataIngestionLocations(res)
        }
        return nil
    }
    res["allowedDataStorageLocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseManagedAppDataStorageLocation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppDataStorageLocation, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*ManagedAppDataStorageLocation))
                }
            }
            m.SetAllowedDataStorageLocations(res)
        }
        return nil
    }
    res["allowedInboundDataTransferSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedInboundDataTransferSources(val.(*ManagedAppDataTransferLevel))
        }
        return nil
    }
    res["allowedOutboundClipboardSharingExceptionLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundClipboardSharingExceptionLength(val)
        }
        return nil
    }
    res["allowedOutboundClipboardSharingLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppClipboardSharingLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundClipboardSharingLevel(val.(*ManagedAppClipboardSharingLevel))
        }
        return nil
    }
    res["allowedOutboundDataTransferDestinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundDataTransferDestinations(val.(*ManagedAppDataTransferLevel))
        }
        return nil
    }
    res["appActionIfDeviceComplianceRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfDeviceComplianceRequired(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfMaximumPinRetriesExceeded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfMaximumPinRetriesExceeded(val.(*ManagedAppRemediationAction))
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
    res["blockDataIngestionIntoOrganizationDocuments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockDataIngestionIntoOrganizationDocuments(val)
        }
        return nil
    }
    res["contactSyncBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactSyncBlocked(val)
        }
        return nil
    }
    res["dataBackupBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataBackupBlocked(val)
        }
        return nil
    }
    res["deviceComplianceRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceRequired(val)
        }
        return nil
    }
    res["dialerRestrictionLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppPhoneNumberRedirectLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDialerRestrictionLevel(val.(*ManagedAppPhoneNumberRedirectLevel))
        }
        return nil
    }
    res["disableAppPinIfDevicePinIsSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAppPinIfDevicePinIsSet(val)
        }
        return nil
    }
    res["fingerprintBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprintBlocked(val)
        }
        return nil
    }
    res["gracePeriodToBlockAppsDuringOffClockHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodToBlockAppsDuringOffClockHours(val)
        }
        return nil
    }
    res["managedBrowser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedBrowserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBrowser(val.(*ManagedBrowserType))
        }
        return nil
    }
    res["managedBrowserToOpenLinksRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBrowserToOpenLinksRequired(val)
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
    res["maximumPinRetries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumPinRetries(val)
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
    res["minimumPinLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumPinLength(val)
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
    res["mobileThreatDefensePartnerPriority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileThreatDefensePartnerPriority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileThreatDefensePartnerPriority(val.(*MobileThreatDefensePartnerPriority))
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
    res["notificationRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppNotificationRestriction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationRestriction(val.(*ManagedAppNotificationRestriction))
        }
        return nil
    }
    res["organizationalCredentialsRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalCredentialsRequired(val)
        }
        return nil
    }
    res["periodBeforePinReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodBeforePinReset(val)
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
    res["periodOnlineBeforeAccessCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodOnlineBeforeAccessCheck(val)
        }
        return nil
    }
    res["pinCharacterSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppPinCharacterSet)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinCharacterSet(val.(*ManagedAppPinCharacterSet))
        }
        return nil
    }
    res["pinRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinRequired(val)
        }
        return nil
    }
    res["pinRequiredInsteadOfBiometricTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinRequiredInsteadOfBiometricTimeout(val)
        }
        return nil
    }
    res["previousPinBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousPinBlockCount(val)
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
    res["saveAsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSaveAsBlocked(val)
        }
        return nil
    }
    res["simplePinBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetFingerprintBlocked gets the fingerprintBlocked property value. Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
func (m *ManagedAppProtection) GetFingerprintBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("fingerprintBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetGracePeriodToBlockAppsDuringOffClockHours gets the gracePeriodToBlockAppsDuringOffClockHours property value. A grace period before blocking app access during off clock hours.
func (m *ManagedAppProtection) GetGracePeriodToBlockAppsDuringOffClockHours()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("gracePeriodToBlockAppsDuringOffClockHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetManagedBrowser gets the managedBrowser property value. Type of managed browser
func (m *ManagedAppProtection) GetManagedBrowser()(*ManagedBrowserType) {
    val, err := m.GetBackingStore().Get("managedBrowser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedBrowserType)
    }
    return nil
}
// GetManagedBrowserToOpenLinksRequired gets the managedBrowserToOpenLinksRequired property value. Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
func (m *ManagedAppProtection) GetManagedBrowserToOpenLinksRequired()(*bool) {
    val, err := m.GetBackingStore().Get("managedBrowserToOpenLinksRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMaximumAllowedDeviceThreatLevel gets the maximumAllowedDeviceThreatLevel property value. The maxium threat level allowed for an app to be compliant.
func (m *ManagedAppProtection) GetMaximumAllowedDeviceThreatLevel()(*ManagedAppDeviceThreatLevel) {
    val, err := m.GetBackingStore().Get("maximumAllowedDeviceThreatLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppDeviceThreatLevel)
    }
    return nil
}
// GetMaximumPinRetries gets the maximumPinRetries property value. Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
func (m *ManagedAppProtection) GetMaximumPinRetries()(*int32) {
    val, err := m.GetBackingStore().Get("maximumPinRetries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMaximumRequiredOsVersion gets the maximumRequiredOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMaximumRequiredOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("maximumRequiredOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMaximumWarningOsVersion gets the maximumWarningOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMaximumWarningOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("maximumWarningOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMaximumWipeOsVersion gets the maximumWipeOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMaximumWipeOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("maximumWipeOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumPinLength gets the minimumPinLength property value. Minimum pin length required for an app-level pin if PinRequired is set to True
func (m *ManagedAppProtection) GetMinimumPinLength()(*int32) {
    val, err := m.GetBackingStore().Get("minimumPinLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMinimumRequiredAppVersion gets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumRequiredAppVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumRequiredAppVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumRequiredOsVersion gets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumRequiredOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumRequiredOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWarningAppVersion gets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app.
func (m *ManagedAppProtection) GetMinimumWarningAppVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWarningAppVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWarningOsVersion gets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumWarningOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWarningOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWipeAppVersion gets the minimumWipeAppVersion property value. Versions less than or equal to the specified version will wipe the managed app and the associated company data.
func (m *ManagedAppProtection) GetMinimumWipeAppVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWipeAppVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumWipeOsVersion gets the minimumWipeOsVersion property value. Versions less than or equal to the specified version will wipe the managed app and the associated company data.
func (m *ManagedAppProtection) GetMinimumWipeOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("minimumWipeOsVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMobileThreatDefensePartnerPriority gets the mobileThreatDefensePartnerPriority property value. Indicates how to prioritize which Mobile Threat Defense (MTD) partner is enabled for a given platform, when more than one is enabled. An app can only be actively using a single Mobile Threat Defense partner. When NULL, Microsoft Defender will be given preference. Otherwise setting the value to defenderOverThirdPartyPartner or thirdPartyPartnerOverDefender will make explicit which partner to prioritize. Possible values are: null, defenderOverThirdPartyPartner, thirdPartyPartnerOverDefender and unknownFutureValue. Default value is null. Possible values are: defenderOverThirdPartyPartner, thirdPartyPartnerOverDefender, unknownFutureValue.
func (m *ManagedAppProtection) GetMobileThreatDefensePartnerPriority()(*MobileThreatDefensePartnerPriority) {
    val, err := m.GetBackingStore().Get("mobileThreatDefensePartnerPriority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MobileThreatDefensePartnerPriority)
    }
    return nil
}
// GetMobileThreatDefenseRemediationAction gets the mobileThreatDefenseRemediationAction property value. An admin initiated action to be applied on a managed app.
func (m *ManagedAppProtection) GetMobileThreatDefenseRemediationAction()(*ManagedAppRemediationAction) {
    val, err := m.GetBackingStore().Get("mobileThreatDefenseRemediationAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppRemediationAction)
    }
    return nil
}
// GetNotificationRestriction gets the notificationRestriction property value. Restrict managed app notification
func (m *ManagedAppProtection) GetNotificationRestriction()(*ManagedAppNotificationRestriction) {
    val, err := m.GetBackingStore().Get("notificationRestriction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppNotificationRestriction)
    }
    return nil
}
// GetOrganizationalCredentialsRequired gets the organizationalCredentialsRequired property value. Indicates whether organizational credentials are required for app use.
func (m *ManagedAppProtection) GetOrganizationalCredentialsRequired()(*bool) {
    val, err := m.GetBackingStore().Get("organizationalCredentialsRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPeriodBeforePinReset gets the periodBeforePinReset property value. TimePeriod before the all-level pin must be reset if PinRequired is set to True.
func (m *ManagedAppProtection) GetPeriodBeforePinReset()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("periodBeforePinReset")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetPeriodOfflineBeforeAccessCheck gets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet.
func (m *ManagedAppProtection) GetPeriodOfflineBeforeAccessCheck()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("periodOfflineBeforeAccessCheck")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetPeriodOfflineBeforeWipeIsEnforced gets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
func (m *ManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("periodOfflineBeforeWipeIsEnforced")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetPeriodOnlineBeforeAccessCheck gets the periodOnlineBeforeAccessCheck property value. The period after which access is checked when the device is connected to the internet.
func (m *ManagedAppProtection) GetPeriodOnlineBeforeAccessCheck()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("periodOnlineBeforeAccessCheck")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetPinCharacterSet gets the pinCharacterSet property value. Character set which is to be used for a user's app PIN
func (m *ManagedAppProtection) GetPinCharacterSet()(*ManagedAppPinCharacterSet) {
    val, err := m.GetBackingStore().Get("pinCharacterSet")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedAppPinCharacterSet)
    }
    return nil
}
// GetPinRequired gets the pinRequired property value. Indicates whether an app-level pin is required.
func (m *ManagedAppProtection) GetPinRequired()(*bool) {
    val, err := m.GetBackingStore().Get("pinRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPinRequiredInsteadOfBiometricTimeout gets the pinRequiredInsteadOfBiometricTimeout property value. Timeout in minutes for an app pin instead of non biometrics passcode
func (m *ManagedAppProtection) GetPinRequiredInsteadOfBiometricTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("pinRequiredInsteadOfBiometricTimeout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetPreviousPinBlockCount gets the previousPinBlockCount property value. Requires a pin to be unique from the number specified in this property.
func (m *ManagedAppProtection) GetPreviousPinBlockCount()(*int32) {
    val, err := m.GetBackingStore().Get("previousPinBlockCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPrintBlocked gets the printBlocked property value. Indicates whether printing is allowed from managed apps.
func (m *ManagedAppProtection) GetPrintBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("printBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSaveAsBlocked gets the saveAsBlocked property value. Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
func (m *ManagedAppProtection) GetSaveAsBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("saveAsBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSimplePinBlocked gets the simplePinBlocked property value. Indicates whether simplePin is blocked.
func (m *ManagedAppProtection) GetSimplePinBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("simplePinBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedAppProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetMobileThreatDefensePartnerPriority() != nil {
        cast := (*m.GetMobileThreatDefensePartnerPriority()).String()
        err = writer.WriteStringValue("mobileThreatDefensePartnerPriority", &cast)
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
    err := m.GetBackingStore().Set("allowedDataIngestionLocations", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedDataStorageLocations sets the allowedDataStorageLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) SetAllowedDataStorageLocations(value []ManagedAppDataStorageLocation)() {
    err := m.GetBackingStore().Set("allowedDataStorageLocations", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedInboundDataTransferSources sets the allowedInboundDataTransferSources property value. Data can be transferred from/to these classes of apps
func (m *ManagedAppProtection) SetAllowedInboundDataTransferSources(value *ManagedAppDataTransferLevel)() {
    err := m.GetBackingStore().Set("allowedInboundDataTransferSources", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedOutboundClipboardSharingExceptionLength sets the allowedOutboundClipboardSharingExceptionLength property value. Specify the number of characters that may be cut or copied from Org data and accounts to any application. This setting overrides the AllowedOutboundClipboardSharingLevel restriction. Default value of '0' means no exception is allowed.
func (m *ManagedAppProtection) SetAllowedOutboundClipboardSharingExceptionLength(value *int32)() {
    err := m.GetBackingStore().Set("allowedOutboundClipboardSharingExceptionLength", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedOutboundClipboardSharingLevel sets the allowedOutboundClipboardSharingLevel property value. Represents the level to which the device's clipboard may be shared between apps
func (m *ManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(value *ManagedAppClipboardSharingLevel)() {
    err := m.GetBackingStore().Set("allowedOutboundClipboardSharingLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedOutboundDataTransferDestinations sets the allowedOutboundDataTransferDestinations property value. Data can be transferred from/to these classes of apps
func (m *ManagedAppProtection) SetAllowedOutboundDataTransferDestinations(value *ManagedAppDataTransferLevel)() {
    err := m.GetBackingStore().Set("allowedOutboundDataTransferDestinations", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfDeviceComplianceRequired sets the appActionIfDeviceComplianceRequired property value. An admin initiated action to be applied on a managed app.
func (m *ManagedAppProtection) SetAppActionIfDeviceComplianceRequired(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfDeviceComplianceRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfMaximumPinRetriesExceeded sets the appActionIfMaximumPinRetriesExceeded property value. An admin initiated action to be applied on a managed app.
func (m *ManagedAppProtection) SetAppActionIfMaximumPinRetriesExceeded(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfMaximumPinRetriesExceeded", value)
    if err != nil {
        panic(err)
    }
}
// SetAppActionIfUnableToAuthenticateUser sets the appActionIfUnableToAuthenticateUser property value. If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Possible values are: block, wipe, warn.
func (m *ManagedAppProtection) SetAppActionIfUnableToAuthenticateUser(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("appActionIfUnableToAuthenticateUser", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockDataIngestionIntoOrganizationDocuments sets the blockDataIngestionIntoOrganizationDocuments property value. Indicates whether a user can bring data into org documents.
func (m *ManagedAppProtection) SetBlockDataIngestionIntoOrganizationDocuments(value *bool)() {
    err := m.GetBackingStore().Set("blockDataIngestionIntoOrganizationDocuments", value)
    if err != nil {
        panic(err)
    }
}
// SetContactSyncBlocked sets the contactSyncBlocked property value. Indicates whether contacts can be synced to the user's device.
func (m *ManagedAppProtection) SetContactSyncBlocked(value *bool)() {
    err := m.GetBackingStore().Set("contactSyncBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetDataBackupBlocked sets the dataBackupBlocked property value. Indicates whether the backup of a managed app's data is blocked.
func (m *ManagedAppProtection) SetDataBackupBlocked(value *bool)() {
    err := m.GetBackingStore().Set("dataBackupBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceComplianceRequired sets the deviceComplianceRequired property value. Indicates whether device compliance is required.
func (m *ManagedAppProtection) SetDeviceComplianceRequired(value *bool)() {
    err := m.GetBackingStore().Set("deviceComplianceRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetDialerRestrictionLevel sets the dialerRestrictionLevel property value. The classes of apps that are allowed to click-to-open a phone number, for making phone calls or sending text messages.
func (m *ManagedAppProtection) SetDialerRestrictionLevel(value *ManagedAppPhoneNumberRedirectLevel)() {
    err := m.GetBackingStore().Set("dialerRestrictionLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableAppPinIfDevicePinIsSet sets the disableAppPinIfDevicePinIsSet property value. Indicates whether use of the app pin is required if the device pin is set.
func (m *ManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(value *bool)() {
    err := m.GetBackingStore().Set("disableAppPinIfDevicePinIsSet", value)
    if err != nil {
        panic(err)
    }
}
// SetFingerprintBlocked sets the fingerprintBlocked property value. Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
func (m *ManagedAppProtection) SetFingerprintBlocked(value *bool)() {
    err := m.GetBackingStore().Set("fingerprintBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetGracePeriodToBlockAppsDuringOffClockHours sets the gracePeriodToBlockAppsDuringOffClockHours property value. A grace period before blocking app access during off clock hours.
func (m *ManagedAppProtection) SetGracePeriodToBlockAppsDuringOffClockHours(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("gracePeriodToBlockAppsDuringOffClockHours", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedBrowser sets the managedBrowser property value. Type of managed browser
func (m *ManagedAppProtection) SetManagedBrowser(value *ManagedBrowserType)() {
    err := m.GetBackingStore().Set("managedBrowser", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedBrowserToOpenLinksRequired sets the managedBrowserToOpenLinksRequired property value. Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
func (m *ManagedAppProtection) SetManagedBrowserToOpenLinksRequired(value *bool)() {
    err := m.GetBackingStore().Set("managedBrowserToOpenLinksRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumAllowedDeviceThreatLevel sets the maximumAllowedDeviceThreatLevel property value. The maxium threat level allowed for an app to be compliant.
func (m *ManagedAppProtection) SetMaximumAllowedDeviceThreatLevel(value *ManagedAppDeviceThreatLevel)() {
    err := m.GetBackingStore().Set("maximumAllowedDeviceThreatLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumPinRetries sets the maximumPinRetries property value. Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
func (m *ManagedAppProtection) SetMaximumPinRetries(value *int32)() {
    err := m.GetBackingStore().Set("maximumPinRetries", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumRequiredOsVersion sets the maximumRequiredOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMaximumRequiredOsVersion(value *string)() {
    err := m.GetBackingStore().Set("maximumRequiredOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumWarningOsVersion sets the maximumWarningOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMaximumWarningOsVersion(value *string)() {
    err := m.GetBackingStore().Set("maximumWarningOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumWipeOsVersion sets the maximumWipeOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMaximumWipeOsVersion(value *string)() {
    err := m.GetBackingStore().Set("maximumWipeOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumPinLength sets the minimumPinLength property value. Minimum pin length required for an app-level pin if PinRequired is set to True
func (m *ManagedAppProtection) SetMinimumPinLength(value *int32)() {
    err := m.GetBackingStore().Set("minimumPinLength", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumRequiredAppVersion sets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMinimumRequiredAppVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumRequiredAppVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumRequiredOsVersion sets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMinimumRequiredOsVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumRequiredOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWarningAppVersion sets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app.
func (m *ManagedAppProtection) SetMinimumWarningAppVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWarningAppVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWarningOsVersion sets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
func (m *ManagedAppProtection) SetMinimumWarningOsVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWarningOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWipeAppVersion sets the minimumWipeAppVersion property value. Versions less than or equal to the specified version will wipe the managed app and the associated company data.
func (m *ManagedAppProtection) SetMinimumWipeAppVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWipeAppVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumWipeOsVersion sets the minimumWipeOsVersion property value. Versions less than or equal to the specified version will wipe the managed app and the associated company data.
func (m *ManagedAppProtection) SetMinimumWipeOsVersion(value *string)() {
    err := m.GetBackingStore().Set("minimumWipeOsVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileThreatDefensePartnerPriority sets the mobileThreatDefensePartnerPriority property value. Indicates how to prioritize which Mobile Threat Defense (MTD) partner is enabled for a given platform, when more than one is enabled. An app can only be actively using a single Mobile Threat Defense partner. When NULL, Microsoft Defender will be given preference. Otherwise setting the value to defenderOverThirdPartyPartner or thirdPartyPartnerOverDefender will make explicit which partner to prioritize. Possible values are: null, defenderOverThirdPartyPartner, thirdPartyPartnerOverDefender and unknownFutureValue. Default value is null. Possible values are: defenderOverThirdPartyPartner, thirdPartyPartnerOverDefender, unknownFutureValue.
func (m *ManagedAppProtection) SetMobileThreatDefensePartnerPriority(value *MobileThreatDefensePartnerPriority)() {
    err := m.GetBackingStore().Set("mobileThreatDefensePartnerPriority", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileThreatDefenseRemediationAction sets the mobileThreatDefenseRemediationAction property value. An admin initiated action to be applied on a managed app.
func (m *ManagedAppProtection) SetMobileThreatDefenseRemediationAction(value *ManagedAppRemediationAction)() {
    err := m.GetBackingStore().Set("mobileThreatDefenseRemediationAction", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationRestriction sets the notificationRestriction property value. Restrict managed app notification
func (m *ManagedAppProtection) SetNotificationRestriction(value *ManagedAppNotificationRestriction)() {
    err := m.GetBackingStore().Set("notificationRestriction", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationalCredentialsRequired sets the organizationalCredentialsRequired property value. Indicates whether organizational credentials are required for app use.
func (m *ManagedAppProtection) SetOrganizationalCredentialsRequired(value *bool)() {
    err := m.GetBackingStore().Set("organizationalCredentialsRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriodBeforePinReset sets the periodBeforePinReset property value. TimePeriod before the all-level pin must be reset if PinRequired is set to True.
func (m *ManagedAppProtection) SetPeriodBeforePinReset(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("periodBeforePinReset", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriodOfflineBeforeAccessCheck sets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet.
func (m *ManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("periodOfflineBeforeAccessCheck", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriodOfflineBeforeWipeIsEnforced sets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
func (m *ManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("periodOfflineBeforeWipeIsEnforced", value)
    if err != nil {
        panic(err)
    }
}
// SetPeriodOnlineBeforeAccessCheck sets the periodOnlineBeforeAccessCheck property value. The period after which access is checked when the device is connected to the internet.
func (m *ManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("periodOnlineBeforeAccessCheck", value)
    if err != nil {
        panic(err)
    }
}
// SetPinCharacterSet sets the pinCharacterSet property value. Character set which is to be used for a user's app PIN
func (m *ManagedAppProtection) SetPinCharacterSet(value *ManagedAppPinCharacterSet)() {
    err := m.GetBackingStore().Set("pinCharacterSet", value)
    if err != nil {
        panic(err)
    }
}
// SetPinRequired sets the pinRequired property value. Indicates whether an app-level pin is required.
func (m *ManagedAppProtection) SetPinRequired(value *bool)() {
    err := m.GetBackingStore().Set("pinRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetPinRequiredInsteadOfBiometricTimeout sets the pinRequiredInsteadOfBiometricTimeout property value. Timeout in minutes for an app pin instead of non biometrics passcode
func (m *ManagedAppProtection) SetPinRequiredInsteadOfBiometricTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("pinRequiredInsteadOfBiometricTimeout", value)
    if err != nil {
        panic(err)
    }
}
// SetPreviousPinBlockCount sets the previousPinBlockCount property value. Requires a pin to be unique from the number specified in this property.
func (m *ManagedAppProtection) SetPreviousPinBlockCount(value *int32)() {
    err := m.GetBackingStore().Set("previousPinBlockCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPrintBlocked sets the printBlocked property value. Indicates whether printing is allowed from managed apps.
func (m *ManagedAppProtection) SetPrintBlocked(value *bool)() {
    err := m.GetBackingStore().Set("printBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetSaveAsBlocked sets the saveAsBlocked property value. Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
func (m *ManagedAppProtection) SetSaveAsBlocked(value *bool)() {
    err := m.GetBackingStore().Set("saveAsBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetSimplePinBlocked sets the simplePinBlocked property value. Indicates whether simplePin is blocked.
func (m *ManagedAppProtection) SetSimplePinBlocked(value *bool)() {
    err := m.GetBackingStore().Set("simplePinBlocked", value)
    if err != nil {
        panic(err)
    }
}
// ManagedAppProtectionable 
type ManagedAppProtectionable interface {
    ManagedAppPolicyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedDataIngestionLocations()([]ManagedAppDataIngestionLocation)
    GetAllowedDataStorageLocations()([]ManagedAppDataStorageLocation)
    GetAllowedInboundDataTransferSources()(*ManagedAppDataTransferLevel)
    GetAllowedOutboundClipboardSharingExceptionLength()(*int32)
    GetAllowedOutboundClipboardSharingLevel()(*ManagedAppClipboardSharingLevel)
    GetAllowedOutboundDataTransferDestinations()(*ManagedAppDataTransferLevel)
    GetAppActionIfDeviceComplianceRequired()(*ManagedAppRemediationAction)
    GetAppActionIfMaximumPinRetriesExceeded()(*ManagedAppRemediationAction)
    GetAppActionIfUnableToAuthenticateUser()(*ManagedAppRemediationAction)
    GetBlockDataIngestionIntoOrganizationDocuments()(*bool)
    GetContactSyncBlocked()(*bool)
    GetDataBackupBlocked()(*bool)
    GetDeviceComplianceRequired()(*bool)
    GetDialerRestrictionLevel()(*ManagedAppPhoneNumberRedirectLevel)
    GetDisableAppPinIfDevicePinIsSet()(*bool)
    GetFingerprintBlocked()(*bool)
    GetGracePeriodToBlockAppsDuringOffClockHours()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetManagedBrowser()(*ManagedBrowserType)
    GetManagedBrowserToOpenLinksRequired()(*bool)
    GetMaximumAllowedDeviceThreatLevel()(*ManagedAppDeviceThreatLevel)
    GetMaximumPinRetries()(*int32)
    GetMaximumRequiredOsVersion()(*string)
    GetMaximumWarningOsVersion()(*string)
    GetMaximumWipeOsVersion()(*string)
    GetMinimumPinLength()(*int32)
    GetMinimumRequiredAppVersion()(*string)
    GetMinimumRequiredOsVersion()(*string)
    GetMinimumWarningAppVersion()(*string)
    GetMinimumWarningOsVersion()(*string)
    GetMinimumWipeAppVersion()(*string)
    GetMinimumWipeOsVersion()(*string)
    GetMobileThreatDefensePartnerPriority()(*MobileThreatDefensePartnerPriority)
    GetMobileThreatDefenseRemediationAction()(*ManagedAppRemediationAction)
    GetNotificationRestriction()(*ManagedAppNotificationRestriction)
    GetOrganizationalCredentialsRequired()(*bool)
    GetPeriodBeforePinReset()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPeriodOfflineBeforeAccessCheck()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPeriodOfflineBeforeWipeIsEnforced()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPeriodOnlineBeforeAccessCheck()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPinCharacterSet()(*ManagedAppPinCharacterSet)
    GetPinRequired()(*bool)
    GetPinRequiredInsteadOfBiometricTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPreviousPinBlockCount()(*int32)
    GetPrintBlocked()(*bool)
    GetSaveAsBlocked()(*bool)
    GetSimplePinBlocked()(*bool)
    SetAllowedDataIngestionLocations(value []ManagedAppDataIngestionLocation)()
    SetAllowedDataStorageLocations(value []ManagedAppDataStorageLocation)()
    SetAllowedInboundDataTransferSources(value *ManagedAppDataTransferLevel)()
    SetAllowedOutboundClipboardSharingExceptionLength(value *int32)()
    SetAllowedOutboundClipboardSharingLevel(value *ManagedAppClipboardSharingLevel)()
    SetAllowedOutboundDataTransferDestinations(value *ManagedAppDataTransferLevel)()
    SetAppActionIfDeviceComplianceRequired(value *ManagedAppRemediationAction)()
    SetAppActionIfMaximumPinRetriesExceeded(value *ManagedAppRemediationAction)()
    SetAppActionIfUnableToAuthenticateUser(value *ManagedAppRemediationAction)()
    SetBlockDataIngestionIntoOrganizationDocuments(value *bool)()
    SetContactSyncBlocked(value *bool)()
    SetDataBackupBlocked(value *bool)()
    SetDeviceComplianceRequired(value *bool)()
    SetDialerRestrictionLevel(value *ManagedAppPhoneNumberRedirectLevel)()
    SetDisableAppPinIfDevicePinIsSet(value *bool)()
    SetFingerprintBlocked(value *bool)()
    SetGracePeriodToBlockAppsDuringOffClockHours(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetManagedBrowser(value *ManagedBrowserType)()
    SetManagedBrowserToOpenLinksRequired(value *bool)()
    SetMaximumAllowedDeviceThreatLevel(value *ManagedAppDeviceThreatLevel)()
    SetMaximumPinRetries(value *int32)()
    SetMaximumRequiredOsVersion(value *string)()
    SetMaximumWarningOsVersion(value *string)()
    SetMaximumWipeOsVersion(value *string)()
    SetMinimumPinLength(value *int32)()
    SetMinimumRequiredAppVersion(value *string)()
    SetMinimumRequiredOsVersion(value *string)()
    SetMinimumWarningAppVersion(value *string)()
    SetMinimumWarningOsVersion(value *string)()
    SetMinimumWipeAppVersion(value *string)()
    SetMinimumWipeOsVersion(value *string)()
    SetMobileThreatDefensePartnerPriority(value *MobileThreatDefensePartnerPriority)()
    SetMobileThreatDefenseRemediationAction(value *ManagedAppRemediationAction)()
    SetNotificationRestriction(value *ManagedAppNotificationRestriction)()
    SetOrganizationalCredentialsRequired(value *bool)()
    SetPeriodBeforePinReset(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPeriodOfflineBeforeAccessCheck(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPeriodOfflineBeforeWipeIsEnforced(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPeriodOnlineBeforeAccessCheck(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPinCharacterSet(value *ManagedAppPinCharacterSet)()
    SetPinRequired(value *bool)()
    SetPinRequiredInsteadOfBiometricTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPreviousPinBlockCount(value *int32)()
    SetPrintBlocked(value *bool)()
    SetSaveAsBlocked(value *bool)()
    SetSimplePinBlocked(value *bool)()
}
