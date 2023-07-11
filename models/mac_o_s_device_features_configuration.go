package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSDeviceFeaturesConfiguration macOS device features configuration profile.
type MacOSDeviceFeaturesConfiguration struct {
    AppleDeviceFeaturesConfigurationBase
    // The OdataType property
    OdataType *string
}
// NewMacOSDeviceFeaturesConfiguration instantiates a new macOSDeviceFeaturesConfiguration and sets the default values.
func NewMacOSDeviceFeaturesConfiguration()(*MacOSDeviceFeaturesConfiguration) {
    m := &MacOSDeviceFeaturesConfiguration{
        AppleDeviceFeaturesConfigurationBase: *NewAppleDeviceFeaturesConfigurationBase(),
    }
    odataTypeValue := "#microsoft.graph.macOSDeviceFeaturesConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSDeviceFeaturesConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSDeviceFeaturesConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSDeviceFeaturesConfiguration(), nil
}
// GetAdminShowHostInfo gets the adminShowHostInfo property value. Whether to show admin host information on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAdminShowHostInfo()(*bool) {
    val, err := m.GetBackingStore().Get("adminShowHostInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAppAssociatedDomains gets the appAssociatedDomains property value. Gets or sets a list that maps apps to their associated domains. Application identifiers must be unique. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetAppAssociatedDomains()([]MacOSAssociatedDomainsItemable) {
    val, err := m.GetBackingStore().Get("appAssociatedDomains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSAssociatedDomainsItemable)
    }
    return nil
}
// GetAssociatedDomains gets the associatedDomains property value. DEPRECATED: use appAssociatedDomains instead. Gets or sets a list that maps apps to their associated domains. The key should match the app's ID, and the value should be a string in the form of 'service:domain' where domain is a fully qualified hostname (e.g. webcredentials:example.com). This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetAssociatedDomains()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("associatedDomains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// GetAuthorizedUsersListHidden gets the authorizedUsersListHidden property value. Whether to show the name and password dialog or a list of users on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAuthorizedUsersListHidden()(*bool) {
    val, err := m.GetBackingStore().Get("authorizedUsersListHidden")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAuthorizedUsersListHideAdminUsers gets the authorizedUsersListHideAdminUsers property value. Whether to hide admin users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAuthorizedUsersListHideAdminUsers()(*bool) {
    val, err := m.GetBackingStore().Get("authorizedUsersListHideAdminUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAuthorizedUsersListHideLocalUsers gets the authorizedUsersListHideLocalUsers property value. Whether to show only network and system users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAuthorizedUsersListHideLocalUsers()(*bool) {
    val, err := m.GetBackingStore().Get("authorizedUsersListHideLocalUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAuthorizedUsersListHideMobileAccounts gets the authorizedUsersListHideMobileAccounts property value. Whether to hide mobile users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAuthorizedUsersListHideMobileAccounts()(*bool) {
    val, err := m.GetBackingStore().Get("authorizedUsersListHideMobileAccounts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAuthorizedUsersListIncludeNetworkUsers gets the authorizedUsersListIncludeNetworkUsers property value. Whether to show network users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAuthorizedUsersListIncludeNetworkUsers()(*bool) {
    val, err := m.GetBackingStore().Get("authorizedUsersListIncludeNetworkUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAuthorizedUsersListShowOtherManagedUsers gets the authorizedUsersListShowOtherManagedUsers property value. Whether to show other users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAuthorizedUsersListShowOtherManagedUsers()(*bool) {
    val, err := m.GetBackingStore().Get("authorizedUsersListShowOtherManagedUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAutoLaunchItems gets the autoLaunchItems property value. List of applications, files, folders, and other items to launch when the user logs in. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetAutoLaunchItems()([]MacOSLaunchItemable) {
    val, err := m.GetBackingStore().Get("autoLaunchItems")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSLaunchItemable)
    }
    return nil
}
// GetConsoleAccessDisabled gets the consoleAccessDisabled property value. Whether the Other user will disregard use of the console special user name.
func (m *MacOSDeviceFeaturesConfiguration) GetConsoleAccessDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("consoleAccessDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetContentCachingBlockDeletion gets the contentCachingBlockDeletion property value. Prevents content caches from purging content to free up disk space for other apps.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingBlockDeletion()(*bool) {
    val, err := m.GetBackingStore().Get("contentCachingBlockDeletion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetContentCachingClientListenRanges gets the contentCachingClientListenRanges property value. A list of custom IP ranges content caches will use to listen for clients. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingClientListenRanges()([]IpRangeable) {
    val, err := m.GetBackingStore().Get("contentCachingClientListenRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IpRangeable)
    }
    return nil
}
// GetContentCachingClientPolicy gets the contentCachingClientPolicy property value. Determines which clients a content cache will serve.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingClientPolicy()(*MacOSContentCachingClientPolicy) {
    val, err := m.GetBackingStore().Get("contentCachingClientPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSContentCachingClientPolicy)
    }
    return nil
}
// GetContentCachingDataPath gets the contentCachingDataPath property value. The path to the directory used to store cached content. The value must be (or end with) /Library/Application Support/Apple/AssetCache/Data
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingDataPath()(*string) {
    val, err := m.GetBackingStore().Get("contentCachingDataPath")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetContentCachingDisableConnectionSharing gets the contentCachingDisableConnectionSharing property value. Disables internet connection sharing.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingDisableConnectionSharing()(*bool) {
    val, err := m.GetBackingStore().Get("contentCachingDisableConnectionSharing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetContentCachingEnabled gets the contentCachingEnabled property value. Enables content caching and prevents it from being disabled by the user.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("contentCachingEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetContentCachingForceConnectionSharing gets the contentCachingForceConnectionSharing property value. Forces internet connection sharing. contentCachingDisableConnectionSharing overrides this setting.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingForceConnectionSharing()(*bool) {
    val, err := m.GetBackingStore().Get("contentCachingForceConnectionSharing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetContentCachingKeepAwake gets the contentCachingKeepAwake property value. Prevent the device from sleeping if content caching is enabled.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingKeepAwake()(*bool) {
    val, err := m.GetBackingStore().Get("contentCachingKeepAwake")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetContentCachingLogClientIdentities gets the contentCachingLogClientIdentities property value. Enables logging of IP addresses and ports of clients that request cached content.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingLogClientIdentities()(*bool) {
    val, err := m.GetBackingStore().Get("contentCachingLogClientIdentities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetContentCachingMaxSizeBytes gets the contentCachingMaxSizeBytes property value. The maximum number of bytes of disk space that will be used for the content cache. A value of 0 (default) indicates unlimited disk space.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingMaxSizeBytes()(*int64) {
    val, err := m.GetBackingStore().Get("contentCachingMaxSizeBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetContentCachingParents gets the contentCachingParents property value. A list of IP addresses representing parent content caches.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingParents()([]string) {
    val, err := m.GetBackingStore().Get("contentCachingParents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetContentCachingParentSelectionPolicy gets the contentCachingParentSelectionPolicy property value. Determines how content caches select a parent cache.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingParentSelectionPolicy()(*MacOSContentCachingParentSelectionPolicy) {
    val, err := m.GetBackingStore().Get("contentCachingParentSelectionPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSContentCachingParentSelectionPolicy)
    }
    return nil
}
// GetContentCachingPeerFilterRanges gets the contentCachingPeerFilterRanges property value. A list of custom IP ranges content caches will use to query for content from peers caches. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingPeerFilterRanges()([]IpRangeable) {
    val, err := m.GetBackingStore().Get("contentCachingPeerFilterRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IpRangeable)
    }
    return nil
}
// GetContentCachingPeerListenRanges gets the contentCachingPeerListenRanges property value. A list of custom IP ranges content caches will use to listen for peer caches. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingPeerListenRanges()([]IpRangeable) {
    val, err := m.GetBackingStore().Get("contentCachingPeerListenRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IpRangeable)
    }
    return nil
}
// GetContentCachingPeerPolicy gets the contentCachingPeerPolicy property value. Determines which content caches other content caches will peer with.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingPeerPolicy()(*MacOSContentCachingPeerPolicy) {
    val, err := m.GetBackingStore().Get("contentCachingPeerPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSContentCachingPeerPolicy)
    }
    return nil
}
// GetContentCachingPort gets the contentCachingPort property value. Sets the port used for content caching. If the value is 0, a random available port will be selected. Valid values 0 to 65535
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingPort()(*int32) {
    val, err := m.GetBackingStore().Get("contentCachingPort")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetContentCachingPublicRanges gets the contentCachingPublicRanges property value. A list of custom IP ranges that Apple's content caching service should use to match clients to content caches. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingPublicRanges()([]IpRangeable) {
    val, err := m.GetBackingStore().Get("contentCachingPublicRanges")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IpRangeable)
    }
    return nil
}
// GetContentCachingShowAlerts gets the contentCachingShowAlerts property value. Display content caching alerts as system notifications.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingShowAlerts()(*bool) {
    val, err := m.GetBackingStore().Get("contentCachingShowAlerts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetContentCachingType gets the contentCachingType property value. Indicates the type of content allowed to be cached by Apple's content caching service.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingType()(*MacOSContentCachingType) {
    val, err := m.GetBackingStore().Get("contentCachingType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSContentCachingType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSDeviceFeaturesConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppleDeviceFeaturesConfigurationBase.GetFieldDeserializers()
    res["adminShowHostInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminShowHostInfo(val)
        }
        return nil
    }
    res["appAssociatedDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSAssociatedDomainsItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSAssociatedDomainsItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MacOSAssociatedDomainsItemable)
                }
            }
            m.SetAppAssociatedDomains(res)
        }
        return nil
    }
    res["associatedDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAssociatedDomains(res)
        }
        return nil
    }
    res["authorizedUsersListHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedUsersListHidden(val)
        }
        return nil
    }
    res["authorizedUsersListHideAdminUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedUsersListHideAdminUsers(val)
        }
        return nil
    }
    res["authorizedUsersListHideLocalUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedUsersListHideLocalUsers(val)
        }
        return nil
    }
    res["authorizedUsersListHideMobileAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedUsersListHideMobileAccounts(val)
        }
        return nil
    }
    res["authorizedUsersListIncludeNetworkUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedUsersListIncludeNetworkUsers(val)
        }
        return nil
    }
    res["authorizedUsersListShowOtherManagedUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthorizedUsersListShowOtherManagedUsers(val)
        }
        return nil
    }
    res["autoLaunchItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSLaunchItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSLaunchItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MacOSLaunchItemable)
                }
            }
            m.SetAutoLaunchItems(res)
        }
        return nil
    }
    res["consoleAccessDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsoleAccessDisabled(val)
        }
        return nil
    }
    res["contentCachingBlockDeletion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingBlockDeletion(val)
        }
        return nil
    }
    res["contentCachingClientListenRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIpRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpRangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IpRangeable)
                }
            }
            m.SetContentCachingClientListenRanges(res)
        }
        return nil
    }
    res["contentCachingClientPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSContentCachingClientPolicy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingClientPolicy(val.(*MacOSContentCachingClientPolicy))
        }
        return nil
    }
    res["contentCachingDataPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingDataPath(val)
        }
        return nil
    }
    res["contentCachingDisableConnectionSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingDisableConnectionSharing(val)
        }
        return nil
    }
    res["contentCachingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingEnabled(val)
        }
        return nil
    }
    res["contentCachingForceConnectionSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingForceConnectionSharing(val)
        }
        return nil
    }
    res["contentCachingKeepAwake"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingKeepAwake(val)
        }
        return nil
    }
    res["contentCachingLogClientIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingLogClientIdentities(val)
        }
        return nil
    }
    res["contentCachingMaxSizeBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingMaxSizeBytes(val)
        }
        return nil
    }
    res["contentCachingParents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetContentCachingParents(res)
        }
        return nil
    }
    res["contentCachingParentSelectionPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSContentCachingParentSelectionPolicy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingParentSelectionPolicy(val.(*MacOSContentCachingParentSelectionPolicy))
        }
        return nil
    }
    res["contentCachingPeerFilterRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIpRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpRangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IpRangeable)
                }
            }
            m.SetContentCachingPeerFilterRanges(res)
        }
        return nil
    }
    res["contentCachingPeerListenRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIpRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpRangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IpRangeable)
                }
            }
            m.SetContentCachingPeerListenRanges(res)
        }
        return nil
    }
    res["contentCachingPeerPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSContentCachingPeerPolicy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingPeerPolicy(val.(*MacOSContentCachingPeerPolicy))
        }
        return nil
    }
    res["contentCachingPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingPort(val)
        }
        return nil
    }
    res["contentCachingPublicRanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIpRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IpRangeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IpRangeable)
                }
            }
            m.SetContentCachingPublicRanges(res)
        }
        return nil
    }
    res["contentCachingShowAlerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingShowAlerts(val)
        }
        return nil
    }
    res["contentCachingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSContentCachingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingType(val.(*MacOSContentCachingType))
        }
        return nil
    }
    res["loginWindowText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginWindowText(val)
        }
        return nil
    }
    res["logOutDisabledWhileLoggedIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogOutDisabledWhileLoggedIn(val)
        }
        return nil
    }
    res["macOSSingleSignOnExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMacOSSingleSignOnExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOSSingleSignOnExtension(val.(MacOSSingleSignOnExtensionable))
        }
        return nil
    }
    res["powerOffDisabledWhileLoggedIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPowerOffDisabledWhileLoggedIn(val)
        }
        return nil
    }
    res["restartDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartDisabled(val)
        }
        return nil
    }
    res["restartDisabledWhileLoggedIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartDisabledWhileLoggedIn(val)
        }
        return nil
    }
    res["screenLockDisableImmediate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenLockDisableImmediate(val)
        }
        return nil
    }
    res["shutDownDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShutDownDisabled(val)
        }
        return nil
    }
    res["shutDownDisabledWhileLoggedIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShutDownDisabledWhileLoggedIn(val)
        }
        return nil
    }
    res["singleSignOnExtension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSingleSignOnExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnExtension(val.(SingleSignOnExtensionable))
        }
        return nil
    }
    res["singleSignOnExtensionPkinitCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMacOSCertificateProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnExtensionPkinitCertificate(val.(MacOSCertificateProfileBaseable))
        }
        return nil
    }
    res["sleepDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSleepDisabled(val)
        }
        return nil
    }
    return res
}
// GetLoginWindowText gets the loginWindowText property value. Custom text to be displayed on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetLoginWindowText()(*string) {
    val, err := m.GetBackingStore().Get("loginWindowText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLogOutDisabledWhileLoggedIn gets the logOutDisabledWhileLoggedIn property value. Whether the Log Out menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) GetLogOutDisabledWhileLoggedIn()(*bool) {
    val, err := m.GetBackingStore().Get("logOutDisabledWhileLoggedIn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMacOSSingleSignOnExtension gets the macOSSingleSignOnExtension property value. Gets or sets a single sign-on extension profile.
func (m *MacOSDeviceFeaturesConfiguration) GetMacOSSingleSignOnExtension()(MacOSSingleSignOnExtensionable) {
    val, err := m.GetBackingStore().Get("macOSSingleSignOnExtension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MacOSSingleSignOnExtensionable)
    }
    return nil
}
// GetPowerOffDisabledWhileLoggedIn gets the powerOffDisabledWhileLoggedIn property value. Whether the Power Off menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) GetPowerOffDisabledWhileLoggedIn()(*bool) {
    val, err := m.GetBackingStore().Get("powerOffDisabledWhileLoggedIn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRestartDisabled gets the restartDisabled property value. Whether to hide the Restart button item on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetRestartDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("restartDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRestartDisabledWhileLoggedIn gets the restartDisabledWhileLoggedIn property value. Whether the Restart menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) GetRestartDisabledWhileLoggedIn()(*bool) {
    val, err := m.GetBackingStore().Get("restartDisabledWhileLoggedIn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetScreenLockDisableImmediate gets the screenLockDisableImmediate property value. Whether to disable the immediate screen lock functions.
func (m *MacOSDeviceFeaturesConfiguration) GetScreenLockDisableImmediate()(*bool) {
    val, err := m.GetBackingStore().Get("screenLockDisableImmediate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShutDownDisabled gets the shutDownDisabled property value. Whether to hide the Shut Down button item on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetShutDownDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("shutDownDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetShutDownDisabledWhileLoggedIn gets the shutDownDisabledWhileLoggedIn property value. Whether the Shut Down menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) GetShutDownDisabledWhileLoggedIn()(*bool) {
    val, err := m.GetBackingStore().Get("shutDownDisabledWhileLoggedIn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSingleSignOnExtension gets the singleSignOnExtension property value. Gets or sets a single sign-on extension profile. Deprecated: use MacOSSingleSignOnExtension instead.
func (m *MacOSDeviceFeaturesConfiguration) GetSingleSignOnExtension()(SingleSignOnExtensionable) {
    val, err := m.GetBackingStore().Get("singleSignOnExtension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SingleSignOnExtensionable)
    }
    return nil
}
// GetSingleSignOnExtensionPkinitCertificate gets the singleSignOnExtensionPkinitCertificate property value. PKINIT Certificate for the authentication with single sign-on extensions.
func (m *MacOSDeviceFeaturesConfiguration) GetSingleSignOnExtensionPkinitCertificate()(MacOSCertificateProfileBaseable) {
    val, err := m.GetBackingStore().Get("singleSignOnExtensionPkinitCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MacOSCertificateProfileBaseable)
    }
    return nil
}
// GetSleepDisabled gets the sleepDisabled property value. Whether to hide the Sleep menu item on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetSleepDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("sleepDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSDeviceFeaturesConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AppleDeviceFeaturesConfigurationBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("adminShowHostInfo", m.GetAdminShowHostInfo())
        if err != nil {
            return err
        }
    }
    if m.GetAppAssociatedDomains() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppAssociatedDomains()))
        for i, v := range m.GetAppAssociatedDomains() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appAssociatedDomains", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssociatedDomains() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssociatedDomains()))
        for i, v := range m.GetAssociatedDomains() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("associatedDomains", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("authorizedUsersListHidden", m.GetAuthorizedUsersListHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("authorizedUsersListHideAdminUsers", m.GetAuthorizedUsersListHideAdminUsers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("authorizedUsersListHideLocalUsers", m.GetAuthorizedUsersListHideLocalUsers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("authorizedUsersListHideMobileAccounts", m.GetAuthorizedUsersListHideMobileAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("authorizedUsersListIncludeNetworkUsers", m.GetAuthorizedUsersListIncludeNetworkUsers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("authorizedUsersListShowOtherManagedUsers", m.GetAuthorizedUsersListShowOtherManagedUsers())
        if err != nil {
            return err
        }
    }
    if m.GetAutoLaunchItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAutoLaunchItems()))
        for i, v := range m.GetAutoLaunchItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("autoLaunchItems", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("consoleAccessDisabled", m.GetConsoleAccessDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contentCachingBlockDeletion", m.GetContentCachingBlockDeletion())
        if err != nil {
            return err
        }
    }
    if m.GetContentCachingClientListenRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContentCachingClientListenRanges()))
        for i, v := range m.GetContentCachingClientListenRanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("contentCachingClientListenRanges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContentCachingClientPolicy() != nil {
        cast := (*m.GetContentCachingClientPolicy()).String()
        err = writer.WriteStringValue("contentCachingClientPolicy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentCachingDataPath", m.GetContentCachingDataPath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contentCachingDisableConnectionSharing", m.GetContentCachingDisableConnectionSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contentCachingEnabled", m.GetContentCachingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contentCachingForceConnectionSharing", m.GetContentCachingForceConnectionSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contentCachingKeepAwake", m.GetContentCachingKeepAwake())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contentCachingLogClientIdentities", m.GetContentCachingLogClientIdentities())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("contentCachingMaxSizeBytes", m.GetContentCachingMaxSizeBytes())
        if err != nil {
            return err
        }
    }
    if m.GetContentCachingParents() != nil {
        err = writer.WriteCollectionOfStringValues("contentCachingParents", m.GetContentCachingParents())
        if err != nil {
            return err
        }
    }
    if m.GetContentCachingParentSelectionPolicy() != nil {
        cast := (*m.GetContentCachingParentSelectionPolicy()).String()
        err = writer.WriteStringValue("contentCachingParentSelectionPolicy", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetContentCachingPeerFilterRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContentCachingPeerFilterRanges()))
        for i, v := range m.GetContentCachingPeerFilterRanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("contentCachingPeerFilterRanges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContentCachingPeerListenRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContentCachingPeerListenRanges()))
        for i, v := range m.GetContentCachingPeerListenRanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("contentCachingPeerListenRanges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContentCachingPeerPolicy() != nil {
        cast := (*m.GetContentCachingPeerPolicy()).String()
        err = writer.WriteStringValue("contentCachingPeerPolicy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("contentCachingPort", m.GetContentCachingPort())
        if err != nil {
            return err
        }
    }
    if m.GetContentCachingPublicRanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContentCachingPublicRanges()))
        for i, v := range m.GetContentCachingPublicRanges() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("contentCachingPublicRanges", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contentCachingShowAlerts", m.GetContentCachingShowAlerts())
        if err != nil {
            return err
        }
    }
    if m.GetContentCachingType() != nil {
        cast := (*m.GetContentCachingType()).String()
        err = writer.WriteStringValue("contentCachingType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("loginWindowText", m.GetLoginWindowText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("logOutDisabledWhileLoggedIn", m.GetLogOutDisabledWhileLoggedIn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("macOSSingleSignOnExtension", m.GetMacOSSingleSignOnExtension())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("powerOffDisabledWhileLoggedIn", m.GetPowerOffDisabledWhileLoggedIn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("restartDisabled", m.GetRestartDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("restartDisabledWhileLoggedIn", m.GetRestartDisabledWhileLoggedIn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("screenLockDisableImmediate", m.GetScreenLockDisableImmediate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("shutDownDisabled", m.GetShutDownDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("shutDownDisabledWhileLoggedIn", m.GetShutDownDisabledWhileLoggedIn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("singleSignOnExtension", m.GetSingleSignOnExtension())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("singleSignOnExtensionPkinitCertificate", m.GetSingleSignOnExtensionPkinitCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("sleepDisabled", m.GetSleepDisabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdminShowHostInfo sets the adminShowHostInfo property value. Whether to show admin host information on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAdminShowHostInfo(value *bool)() {
    err := m.GetBackingStore().Set("adminShowHostInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetAppAssociatedDomains sets the appAssociatedDomains property value. Gets or sets a list that maps apps to their associated domains. Application identifiers must be unique. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetAppAssociatedDomains(value []MacOSAssociatedDomainsItemable)() {
    err := m.GetBackingStore().Set("appAssociatedDomains", value)
    if err != nil {
        panic(err)
    }
}
// SetAssociatedDomains sets the associatedDomains property value. DEPRECATED: use appAssociatedDomains instead. Gets or sets a list that maps apps to their associated domains. The key should match the app's ID, and the value should be a string in the form of 'service:domain' where domain is a fully qualified hostname (e.g. webcredentials:example.com). This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetAssociatedDomains(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("associatedDomains", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthorizedUsersListHidden sets the authorizedUsersListHidden property value. Whether to show the name and password dialog or a list of users on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAuthorizedUsersListHidden(value *bool)() {
    err := m.GetBackingStore().Set("authorizedUsersListHidden", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthorizedUsersListHideAdminUsers sets the authorizedUsersListHideAdminUsers property value. Whether to hide admin users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAuthorizedUsersListHideAdminUsers(value *bool)() {
    err := m.GetBackingStore().Set("authorizedUsersListHideAdminUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthorizedUsersListHideLocalUsers sets the authorizedUsersListHideLocalUsers property value. Whether to show only network and system users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAuthorizedUsersListHideLocalUsers(value *bool)() {
    err := m.GetBackingStore().Set("authorizedUsersListHideLocalUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthorizedUsersListHideMobileAccounts sets the authorizedUsersListHideMobileAccounts property value. Whether to hide mobile users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAuthorizedUsersListHideMobileAccounts(value *bool)() {
    err := m.GetBackingStore().Set("authorizedUsersListHideMobileAccounts", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthorizedUsersListIncludeNetworkUsers sets the authorizedUsersListIncludeNetworkUsers property value. Whether to show network users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAuthorizedUsersListIncludeNetworkUsers(value *bool)() {
    err := m.GetBackingStore().Set("authorizedUsersListIncludeNetworkUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetAuthorizedUsersListShowOtherManagedUsers sets the authorizedUsersListShowOtherManagedUsers property value. Whether to show other users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAuthorizedUsersListShowOtherManagedUsers(value *bool)() {
    err := m.GetBackingStore().Set("authorizedUsersListShowOtherManagedUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetAutoLaunchItems sets the autoLaunchItems property value. List of applications, files, folders, and other items to launch when the user logs in. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetAutoLaunchItems(value []MacOSLaunchItemable)() {
    err := m.GetBackingStore().Set("autoLaunchItems", value)
    if err != nil {
        panic(err)
    }
}
// SetConsoleAccessDisabled sets the consoleAccessDisabled property value. Whether the Other user will disregard use of the console special user name.
func (m *MacOSDeviceFeaturesConfiguration) SetConsoleAccessDisabled(value *bool)() {
    err := m.GetBackingStore().Set("consoleAccessDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingBlockDeletion sets the contentCachingBlockDeletion property value. Prevents content caches from purging content to free up disk space for other apps.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingBlockDeletion(value *bool)() {
    err := m.GetBackingStore().Set("contentCachingBlockDeletion", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingClientListenRanges sets the contentCachingClientListenRanges property value. A list of custom IP ranges content caches will use to listen for clients. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingClientListenRanges(value []IpRangeable)() {
    err := m.GetBackingStore().Set("contentCachingClientListenRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingClientPolicy sets the contentCachingClientPolicy property value. Determines which clients a content cache will serve.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingClientPolicy(value *MacOSContentCachingClientPolicy)() {
    err := m.GetBackingStore().Set("contentCachingClientPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingDataPath sets the contentCachingDataPath property value. The path to the directory used to store cached content. The value must be (or end with) /Library/Application Support/Apple/AssetCache/Data
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingDataPath(value *string)() {
    err := m.GetBackingStore().Set("contentCachingDataPath", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingDisableConnectionSharing sets the contentCachingDisableConnectionSharing property value. Disables internet connection sharing.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingDisableConnectionSharing(value *bool)() {
    err := m.GetBackingStore().Set("contentCachingDisableConnectionSharing", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingEnabled sets the contentCachingEnabled property value. Enables content caching and prevents it from being disabled by the user.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingEnabled(value *bool)() {
    err := m.GetBackingStore().Set("contentCachingEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingForceConnectionSharing sets the contentCachingForceConnectionSharing property value. Forces internet connection sharing. contentCachingDisableConnectionSharing overrides this setting.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingForceConnectionSharing(value *bool)() {
    err := m.GetBackingStore().Set("contentCachingForceConnectionSharing", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingKeepAwake sets the contentCachingKeepAwake property value. Prevent the device from sleeping if content caching is enabled.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingKeepAwake(value *bool)() {
    err := m.GetBackingStore().Set("contentCachingKeepAwake", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingLogClientIdentities sets the contentCachingLogClientIdentities property value. Enables logging of IP addresses and ports of clients that request cached content.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingLogClientIdentities(value *bool)() {
    err := m.GetBackingStore().Set("contentCachingLogClientIdentities", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingMaxSizeBytes sets the contentCachingMaxSizeBytes property value. The maximum number of bytes of disk space that will be used for the content cache. A value of 0 (default) indicates unlimited disk space.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingMaxSizeBytes(value *int64)() {
    err := m.GetBackingStore().Set("contentCachingMaxSizeBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingParents sets the contentCachingParents property value. A list of IP addresses representing parent content caches.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingParents(value []string)() {
    err := m.GetBackingStore().Set("contentCachingParents", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingParentSelectionPolicy sets the contentCachingParentSelectionPolicy property value. Determines how content caches select a parent cache.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingParentSelectionPolicy(value *MacOSContentCachingParentSelectionPolicy)() {
    err := m.GetBackingStore().Set("contentCachingParentSelectionPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingPeerFilterRanges sets the contentCachingPeerFilterRanges property value. A list of custom IP ranges content caches will use to query for content from peers caches. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingPeerFilterRanges(value []IpRangeable)() {
    err := m.GetBackingStore().Set("contentCachingPeerFilterRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingPeerListenRanges sets the contentCachingPeerListenRanges property value. A list of custom IP ranges content caches will use to listen for peer caches. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingPeerListenRanges(value []IpRangeable)() {
    err := m.GetBackingStore().Set("contentCachingPeerListenRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingPeerPolicy sets the contentCachingPeerPolicy property value. Determines which content caches other content caches will peer with.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingPeerPolicy(value *MacOSContentCachingPeerPolicy)() {
    err := m.GetBackingStore().Set("contentCachingPeerPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingPort sets the contentCachingPort property value. Sets the port used for content caching. If the value is 0, a random available port will be selected. Valid values 0 to 65535
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingPort(value *int32)() {
    err := m.GetBackingStore().Set("contentCachingPort", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingPublicRanges sets the contentCachingPublicRanges property value. A list of custom IP ranges that Apple's content caching service should use to match clients to content caches. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingPublicRanges(value []IpRangeable)() {
    err := m.GetBackingStore().Set("contentCachingPublicRanges", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingShowAlerts sets the contentCachingShowAlerts property value. Display content caching alerts as system notifications.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingShowAlerts(value *bool)() {
    err := m.GetBackingStore().Set("contentCachingShowAlerts", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingType sets the contentCachingType property value. Indicates the type of content allowed to be cached by Apple's content caching service.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingType(value *MacOSContentCachingType)() {
    err := m.GetBackingStore().Set("contentCachingType", value)
    if err != nil {
        panic(err)
    }
}
// SetLoginWindowText sets the loginWindowText property value. Custom text to be displayed on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetLoginWindowText(value *string)() {
    err := m.GetBackingStore().Set("loginWindowText", value)
    if err != nil {
        panic(err)
    }
}
// SetLogOutDisabledWhileLoggedIn sets the logOutDisabledWhileLoggedIn property value. Whether the Log Out menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) SetLogOutDisabledWhileLoggedIn(value *bool)() {
    err := m.GetBackingStore().Set("logOutDisabledWhileLoggedIn", value)
    if err != nil {
        panic(err)
    }
}
// SetMacOSSingleSignOnExtension sets the macOSSingleSignOnExtension property value. Gets or sets a single sign-on extension profile.
func (m *MacOSDeviceFeaturesConfiguration) SetMacOSSingleSignOnExtension(value MacOSSingleSignOnExtensionable)() {
    err := m.GetBackingStore().Set("macOSSingleSignOnExtension", value)
    if err != nil {
        panic(err)
    }
}
// SetPowerOffDisabledWhileLoggedIn sets the powerOffDisabledWhileLoggedIn property value. Whether the Power Off menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) SetPowerOffDisabledWhileLoggedIn(value *bool)() {
    err := m.GetBackingStore().Set("powerOffDisabledWhileLoggedIn", value)
    if err != nil {
        panic(err)
    }
}
// SetRestartDisabled sets the restartDisabled property value. Whether to hide the Restart button item on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetRestartDisabled(value *bool)() {
    err := m.GetBackingStore().Set("restartDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetRestartDisabledWhileLoggedIn sets the restartDisabledWhileLoggedIn property value. Whether the Restart menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) SetRestartDisabledWhileLoggedIn(value *bool)() {
    err := m.GetBackingStore().Set("restartDisabledWhileLoggedIn", value)
    if err != nil {
        panic(err)
    }
}
// SetScreenLockDisableImmediate sets the screenLockDisableImmediate property value. Whether to disable the immediate screen lock functions.
func (m *MacOSDeviceFeaturesConfiguration) SetScreenLockDisableImmediate(value *bool)() {
    err := m.GetBackingStore().Set("screenLockDisableImmediate", value)
    if err != nil {
        panic(err)
    }
}
// SetShutDownDisabled sets the shutDownDisabled property value. Whether to hide the Shut Down button item on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetShutDownDisabled(value *bool)() {
    err := m.GetBackingStore().Set("shutDownDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetShutDownDisabledWhileLoggedIn sets the shutDownDisabledWhileLoggedIn property value. Whether the Shut Down menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) SetShutDownDisabledWhileLoggedIn(value *bool)() {
    err := m.GetBackingStore().Set("shutDownDisabledWhileLoggedIn", value)
    if err != nil {
        panic(err)
    }
}
// SetSingleSignOnExtension sets the singleSignOnExtension property value. Gets or sets a single sign-on extension profile. Deprecated: use MacOSSingleSignOnExtension instead.
func (m *MacOSDeviceFeaturesConfiguration) SetSingleSignOnExtension(value SingleSignOnExtensionable)() {
    err := m.GetBackingStore().Set("singleSignOnExtension", value)
    if err != nil {
        panic(err)
    }
}
// SetSingleSignOnExtensionPkinitCertificate sets the singleSignOnExtensionPkinitCertificate property value. PKINIT Certificate for the authentication with single sign-on extensions.
func (m *MacOSDeviceFeaturesConfiguration) SetSingleSignOnExtensionPkinitCertificate(value MacOSCertificateProfileBaseable)() {
    err := m.GetBackingStore().Set("singleSignOnExtensionPkinitCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetSleepDisabled sets the sleepDisabled property value. Whether to hide the Sleep menu item on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetSleepDisabled(value *bool)() {
    err := m.GetBackingStore().Set("sleepDisabled", value)
    if err != nil {
        panic(err)
    }
}
// MacOSDeviceFeaturesConfigurationable 
type MacOSDeviceFeaturesConfigurationable interface {
    AppleDeviceFeaturesConfigurationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminShowHostInfo()(*bool)
    GetAppAssociatedDomains()([]MacOSAssociatedDomainsItemable)
    GetAssociatedDomains()([]KeyValuePairable)
    GetAuthorizedUsersListHidden()(*bool)
    GetAuthorizedUsersListHideAdminUsers()(*bool)
    GetAuthorizedUsersListHideLocalUsers()(*bool)
    GetAuthorizedUsersListHideMobileAccounts()(*bool)
    GetAuthorizedUsersListIncludeNetworkUsers()(*bool)
    GetAuthorizedUsersListShowOtherManagedUsers()(*bool)
    GetAutoLaunchItems()([]MacOSLaunchItemable)
    GetConsoleAccessDisabled()(*bool)
    GetContentCachingBlockDeletion()(*bool)
    GetContentCachingClientListenRanges()([]IpRangeable)
    GetContentCachingClientPolicy()(*MacOSContentCachingClientPolicy)
    GetContentCachingDataPath()(*string)
    GetContentCachingDisableConnectionSharing()(*bool)
    GetContentCachingEnabled()(*bool)
    GetContentCachingForceConnectionSharing()(*bool)
    GetContentCachingKeepAwake()(*bool)
    GetContentCachingLogClientIdentities()(*bool)
    GetContentCachingMaxSizeBytes()(*int64)
    GetContentCachingParents()([]string)
    GetContentCachingParentSelectionPolicy()(*MacOSContentCachingParentSelectionPolicy)
    GetContentCachingPeerFilterRanges()([]IpRangeable)
    GetContentCachingPeerListenRanges()([]IpRangeable)
    GetContentCachingPeerPolicy()(*MacOSContentCachingPeerPolicy)
    GetContentCachingPort()(*int32)
    GetContentCachingPublicRanges()([]IpRangeable)
    GetContentCachingShowAlerts()(*bool)
    GetContentCachingType()(*MacOSContentCachingType)
    GetLoginWindowText()(*string)
    GetLogOutDisabledWhileLoggedIn()(*bool)
    GetMacOSSingleSignOnExtension()(MacOSSingleSignOnExtensionable)
    GetPowerOffDisabledWhileLoggedIn()(*bool)
    GetRestartDisabled()(*bool)
    GetRestartDisabledWhileLoggedIn()(*bool)
    GetScreenLockDisableImmediate()(*bool)
    GetShutDownDisabled()(*bool)
    GetShutDownDisabledWhileLoggedIn()(*bool)
    GetSingleSignOnExtension()(SingleSignOnExtensionable)
    GetSingleSignOnExtensionPkinitCertificate()(MacOSCertificateProfileBaseable)
    GetSleepDisabled()(*bool)
    SetAdminShowHostInfo(value *bool)()
    SetAppAssociatedDomains(value []MacOSAssociatedDomainsItemable)()
    SetAssociatedDomains(value []KeyValuePairable)()
    SetAuthorizedUsersListHidden(value *bool)()
    SetAuthorizedUsersListHideAdminUsers(value *bool)()
    SetAuthorizedUsersListHideLocalUsers(value *bool)()
    SetAuthorizedUsersListHideMobileAccounts(value *bool)()
    SetAuthorizedUsersListIncludeNetworkUsers(value *bool)()
    SetAuthorizedUsersListShowOtherManagedUsers(value *bool)()
    SetAutoLaunchItems(value []MacOSLaunchItemable)()
    SetConsoleAccessDisabled(value *bool)()
    SetContentCachingBlockDeletion(value *bool)()
    SetContentCachingClientListenRanges(value []IpRangeable)()
    SetContentCachingClientPolicy(value *MacOSContentCachingClientPolicy)()
    SetContentCachingDataPath(value *string)()
    SetContentCachingDisableConnectionSharing(value *bool)()
    SetContentCachingEnabled(value *bool)()
    SetContentCachingForceConnectionSharing(value *bool)()
    SetContentCachingKeepAwake(value *bool)()
    SetContentCachingLogClientIdentities(value *bool)()
    SetContentCachingMaxSizeBytes(value *int64)()
    SetContentCachingParents(value []string)()
    SetContentCachingParentSelectionPolicy(value *MacOSContentCachingParentSelectionPolicy)()
    SetContentCachingPeerFilterRanges(value []IpRangeable)()
    SetContentCachingPeerListenRanges(value []IpRangeable)()
    SetContentCachingPeerPolicy(value *MacOSContentCachingPeerPolicy)()
    SetContentCachingPort(value *int32)()
    SetContentCachingPublicRanges(value []IpRangeable)()
    SetContentCachingShowAlerts(value *bool)()
    SetContentCachingType(value *MacOSContentCachingType)()
    SetLoginWindowText(value *string)()
    SetLogOutDisabledWhileLoggedIn(value *bool)()
    SetMacOSSingleSignOnExtension(value MacOSSingleSignOnExtensionable)()
    SetPowerOffDisabledWhileLoggedIn(value *bool)()
    SetRestartDisabled(value *bool)()
    SetRestartDisabledWhileLoggedIn(value *bool)()
    SetScreenLockDisableImmediate(value *bool)()
    SetShutDownDisabled(value *bool)()
    SetShutDownDisabledWhileLoggedIn(value *bool)()
    SetSingleSignOnExtension(value SingleSignOnExtensionable)()
    SetSingleSignOnExtensionPkinitCertificate(value MacOSCertificateProfileBaseable)()
    SetSleepDisabled(value *bool)()
}
