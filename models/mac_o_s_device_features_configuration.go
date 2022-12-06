package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSDeviceFeaturesConfiguration 
type MacOSDeviceFeaturesConfiguration struct {
    AppleDeviceFeaturesConfigurationBase
    // Whether to show admin host information on the login window.
    adminShowHostInfo *bool
    // Gets or sets a list that maps apps to their associated domains. Application identifiers must be unique. This collection can contain a maximum of 500 elements.
    appAssociatedDomains []MacOSAssociatedDomainsItemable
    // DEPRECATED: use appAssociatedDomains instead. Gets or sets a list that maps apps to their associated domains. The key should match the app's ID, and the value should be a string in the form of 'service:domain' where domain is a fully qualified hostname (e.g. webcredentials:example.com). This collection can contain a maximum of 500 elements.
    associatedDomains []KeyValuePairable
    // Whether to show the name and password dialog or a list of users on the login window.
    authorizedUsersListHidden *bool
    // Whether to hide admin users in the authorized users list on the login window.
    authorizedUsersListHideAdminUsers *bool
    // Whether to show only network and system users in the authorized users list on the login window.
    authorizedUsersListHideLocalUsers *bool
    // Whether to hide mobile users in the authorized users list on the login window.
    authorizedUsersListHideMobileAccounts *bool
    // Whether to show network users in the authorized users list on the login window.
    authorizedUsersListIncludeNetworkUsers *bool
    // Whether to show other users in the authorized users list on the login window.
    authorizedUsersListShowOtherManagedUsers *bool
    // List of applications, files, folders, and other items to launch when the user logs in. This collection can contain a maximum of 500 elements.
    autoLaunchItems []MacOSLaunchItemable
    // Whether the Other user will disregard use of the console special user name.
    consoleAccessDisabled *bool
    // Prevents content caches from purging content to free up disk space for other apps.
    contentCachingBlockDeletion *bool
    // A list of custom IP ranges content caches will use to listen for clients. This collection can contain a maximum of 500 elements.
    contentCachingClientListenRanges []IpRangeable
    // Determines which clients a content cache will serve.
    contentCachingClientPolicy *MacOSContentCachingClientPolicy
    // The path to the directory used to store cached content. The value must be (or end with) /Library/Application Support/Apple/AssetCache/Data
    contentCachingDataPath *string
    // Disables internet connection sharing.
    contentCachingDisableConnectionSharing *bool
    // Enables content caching and prevents it from being disabled by the user.
    contentCachingEnabled *bool
    // Forces internet connection sharing. contentCachingDisableConnectionSharing overrides this setting.
    contentCachingForceConnectionSharing *bool
    // Prevent the device from sleeping if content caching is enabled.
    contentCachingKeepAwake *bool
    // Enables logging of IP addresses and ports of clients that request cached content.
    contentCachingLogClientIdentities *bool
    // The maximum number of bytes of disk space that will be used for the content cache. A value of 0 (default) indicates unlimited disk space.
    contentCachingMaxSizeBytes *int64
    // A list of IP addresses representing parent content caches.
    contentCachingParents []string
    // Determines how content caches select a parent cache.
    contentCachingParentSelectionPolicy *MacOSContentCachingParentSelectionPolicy
    // A list of custom IP ranges content caches will use to query for content from peers caches. This collection can contain a maximum of 500 elements.
    contentCachingPeerFilterRanges []IpRangeable
    // A list of custom IP ranges content caches will use to listen for peer caches. This collection can contain a maximum of 500 elements.
    contentCachingPeerListenRanges []IpRangeable
    // Determines which content caches other content caches will peer with.
    contentCachingPeerPolicy *MacOSContentCachingPeerPolicy
    // Sets the port used for content caching. If the value is 0, a random available port will be selected. Valid values 0 to 65535
    contentCachingPort *int32
    // A list of custom IP ranges that Apple's content caching service should use to match clients to content caches. This collection can contain a maximum of 500 elements.
    contentCachingPublicRanges []IpRangeable
    // Display content caching alerts as system notifications.
    contentCachingShowAlerts *bool
    // Indicates the type of content allowed to be cached by Apple's content caching service.
    contentCachingType *MacOSContentCachingType
    // Custom text to be displayed on the login window.
    loginWindowText *string
    // Whether the Log Out menu item on the login window will be disabled while the user is logged in.
    logOutDisabledWhileLoggedIn *bool
    // Gets or sets a single sign-on extension profile.
    macOSSingleSignOnExtension MacOSSingleSignOnExtensionable
    // Whether the Power Off menu item on the login window will be disabled while the user is logged in.
    powerOffDisabledWhileLoggedIn *bool
    // Whether to hide the Restart button item on the login window.
    restartDisabled *bool
    // Whether the Restart menu item on the login window will be disabled while the user is logged in.
    restartDisabledWhileLoggedIn *bool
    // Whether to disable the immediate screen lock functions.
    screenLockDisableImmediate *bool
    // Whether to hide the Shut Down button item on the login window.
    shutDownDisabled *bool
    // Whether the Shut Down menu item on the login window will be disabled while the user is logged in.
    shutDownDisabledWhileLoggedIn *bool
    // Gets or sets a single sign-on extension profile. Deprecated: use MacOSSingleSignOnExtension instead.
    singleSignOnExtension SingleSignOnExtensionable
    // PKINIT Certificate for the authentication with single sign-on extensions.
    singleSignOnExtensionPkinitCertificate MacOSCertificateProfileBaseable
    // Whether to hide the Sleep menu item on the login window.
    sleepDisabled *bool
}
// NewMacOSDeviceFeaturesConfiguration instantiates a new MacOSDeviceFeaturesConfiguration and sets the default values.
func NewMacOSDeviceFeaturesConfiguration()(*MacOSDeviceFeaturesConfiguration) {
    m := &MacOSDeviceFeaturesConfiguration{
        AppleDeviceFeaturesConfigurationBase: *NewAppleDeviceFeaturesConfigurationBase(),
    }
    odataTypeValue := "#microsoft.graph.macOSDeviceFeaturesConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSDeviceFeaturesConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSDeviceFeaturesConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSDeviceFeaturesConfiguration(), nil
}
// GetAdminShowHostInfo gets the adminShowHostInfo property value. Whether to show admin host information on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAdminShowHostInfo()(*bool) {
    return m.adminShowHostInfo
}
// GetAppAssociatedDomains gets the appAssociatedDomains property value. Gets or sets a list that maps apps to their associated domains. Application identifiers must be unique. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetAppAssociatedDomains()([]MacOSAssociatedDomainsItemable) {
    return m.appAssociatedDomains
}
// GetAssociatedDomains gets the associatedDomains property value. DEPRECATED: use appAssociatedDomains instead. Gets or sets a list that maps apps to their associated domains. The key should match the app's ID, and the value should be a string in the form of 'service:domain' where domain is a fully qualified hostname (e.g. webcredentials:example.com). This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetAssociatedDomains()([]KeyValuePairable) {
    return m.associatedDomains
}
// GetAuthorizedUsersListHidden gets the authorizedUsersListHidden property value. Whether to show the name and password dialog or a list of users on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAuthorizedUsersListHidden()(*bool) {
    return m.authorizedUsersListHidden
}
// GetAuthorizedUsersListHideAdminUsers gets the authorizedUsersListHideAdminUsers property value. Whether to hide admin users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAuthorizedUsersListHideAdminUsers()(*bool) {
    return m.authorizedUsersListHideAdminUsers
}
// GetAuthorizedUsersListHideLocalUsers gets the authorizedUsersListHideLocalUsers property value. Whether to show only network and system users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAuthorizedUsersListHideLocalUsers()(*bool) {
    return m.authorizedUsersListHideLocalUsers
}
// GetAuthorizedUsersListHideMobileAccounts gets the authorizedUsersListHideMobileAccounts property value. Whether to hide mobile users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAuthorizedUsersListHideMobileAccounts()(*bool) {
    return m.authorizedUsersListHideMobileAccounts
}
// GetAuthorizedUsersListIncludeNetworkUsers gets the authorizedUsersListIncludeNetworkUsers property value. Whether to show network users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAuthorizedUsersListIncludeNetworkUsers()(*bool) {
    return m.authorizedUsersListIncludeNetworkUsers
}
// GetAuthorizedUsersListShowOtherManagedUsers gets the authorizedUsersListShowOtherManagedUsers property value. Whether to show other users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetAuthorizedUsersListShowOtherManagedUsers()(*bool) {
    return m.authorizedUsersListShowOtherManagedUsers
}
// GetAutoLaunchItems gets the autoLaunchItems property value. List of applications, files, folders, and other items to launch when the user logs in. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetAutoLaunchItems()([]MacOSLaunchItemable) {
    return m.autoLaunchItems
}
// GetConsoleAccessDisabled gets the consoleAccessDisabled property value. Whether the Other user will disregard use of the console special user name.
func (m *MacOSDeviceFeaturesConfiguration) GetConsoleAccessDisabled()(*bool) {
    return m.consoleAccessDisabled
}
// GetContentCachingBlockDeletion gets the contentCachingBlockDeletion property value. Prevents content caches from purging content to free up disk space for other apps.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingBlockDeletion()(*bool) {
    return m.contentCachingBlockDeletion
}
// GetContentCachingClientListenRanges gets the contentCachingClientListenRanges property value. A list of custom IP ranges content caches will use to listen for clients. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingClientListenRanges()([]IpRangeable) {
    return m.contentCachingClientListenRanges
}
// GetContentCachingClientPolicy gets the contentCachingClientPolicy property value. Determines which clients a content cache will serve.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingClientPolicy()(*MacOSContentCachingClientPolicy) {
    return m.contentCachingClientPolicy
}
// GetContentCachingDataPath gets the contentCachingDataPath property value. The path to the directory used to store cached content. The value must be (or end with) /Library/Application Support/Apple/AssetCache/Data
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingDataPath()(*string) {
    return m.contentCachingDataPath
}
// GetContentCachingDisableConnectionSharing gets the contentCachingDisableConnectionSharing property value. Disables internet connection sharing.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingDisableConnectionSharing()(*bool) {
    return m.contentCachingDisableConnectionSharing
}
// GetContentCachingEnabled gets the contentCachingEnabled property value. Enables content caching and prevents it from being disabled by the user.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingEnabled()(*bool) {
    return m.contentCachingEnabled
}
// GetContentCachingForceConnectionSharing gets the contentCachingForceConnectionSharing property value. Forces internet connection sharing. contentCachingDisableConnectionSharing overrides this setting.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingForceConnectionSharing()(*bool) {
    return m.contentCachingForceConnectionSharing
}
// GetContentCachingKeepAwake gets the contentCachingKeepAwake property value. Prevent the device from sleeping if content caching is enabled.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingKeepAwake()(*bool) {
    return m.contentCachingKeepAwake
}
// GetContentCachingLogClientIdentities gets the contentCachingLogClientIdentities property value. Enables logging of IP addresses and ports of clients that request cached content.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingLogClientIdentities()(*bool) {
    return m.contentCachingLogClientIdentities
}
// GetContentCachingMaxSizeBytes gets the contentCachingMaxSizeBytes property value. The maximum number of bytes of disk space that will be used for the content cache. A value of 0 (default) indicates unlimited disk space.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingMaxSizeBytes()(*int64) {
    return m.contentCachingMaxSizeBytes
}
// GetContentCachingParents gets the contentCachingParents property value. A list of IP addresses representing parent content caches.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingParents()([]string) {
    return m.contentCachingParents
}
// GetContentCachingParentSelectionPolicy gets the contentCachingParentSelectionPolicy property value. Determines how content caches select a parent cache.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingParentSelectionPolicy()(*MacOSContentCachingParentSelectionPolicy) {
    return m.contentCachingParentSelectionPolicy
}
// GetContentCachingPeerFilterRanges gets the contentCachingPeerFilterRanges property value. A list of custom IP ranges content caches will use to query for content from peers caches. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingPeerFilterRanges()([]IpRangeable) {
    return m.contentCachingPeerFilterRanges
}
// GetContentCachingPeerListenRanges gets the contentCachingPeerListenRanges property value. A list of custom IP ranges content caches will use to listen for peer caches. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingPeerListenRanges()([]IpRangeable) {
    return m.contentCachingPeerListenRanges
}
// GetContentCachingPeerPolicy gets the contentCachingPeerPolicy property value. Determines which content caches other content caches will peer with.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingPeerPolicy()(*MacOSContentCachingPeerPolicy) {
    return m.contentCachingPeerPolicy
}
// GetContentCachingPort gets the contentCachingPort property value. Sets the port used for content caching. If the value is 0, a random available port will be selected. Valid values 0 to 65535
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingPort()(*int32) {
    return m.contentCachingPort
}
// GetContentCachingPublicRanges gets the contentCachingPublicRanges property value. A list of custom IP ranges that Apple's content caching service should use to match clients to content caches. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingPublicRanges()([]IpRangeable) {
    return m.contentCachingPublicRanges
}
// GetContentCachingShowAlerts gets the contentCachingShowAlerts property value. Display content caching alerts as system notifications.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingShowAlerts()(*bool) {
    return m.contentCachingShowAlerts
}
// GetContentCachingType gets the contentCachingType property value. Indicates the type of content allowed to be cached by Apple's content caching service.
func (m *MacOSDeviceFeaturesConfiguration) GetContentCachingType()(*MacOSContentCachingType) {
    return m.contentCachingType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSDeviceFeaturesConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AppleDeviceFeaturesConfigurationBase.GetFieldDeserializers()
    res["adminShowHostInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAdminShowHostInfo)
    res["appAssociatedDomains"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMacOSAssociatedDomainsItemFromDiscriminatorValue , m.SetAppAssociatedDomains)
    res["associatedDomains"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetAssociatedDomains)
    res["authorizedUsersListHidden"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAuthorizedUsersListHidden)
    res["authorizedUsersListHideAdminUsers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAuthorizedUsersListHideAdminUsers)
    res["authorizedUsersListHideLocalUsers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAuthorizedUsersListHideLocalUsers)
    res["authorizedUsersListHideMobileAccounts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAuthorizedUsersListHideMobileAccounts)
    res["authorizedUsersListIncludeNetworkUsers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAuthorizedUsersListIncludeNetworkUsers)
    res["authorizedUsersListShowOtherManagedUsers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAuthorizedUsersListShowOtherManagedUsers)
    res["autoLaunchItems"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMacOSLaunchItemFromDiscriminatorValue , m.SetAutoLaunchItems)
    res["consoleAccessDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetConsoleAccessDisabled)
    res["contentCachingBlockDeletion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetContentCachingBlockDeletion)
    res["contentCachingClientListenRanges"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIpRangeFromDiscriminatorValue , m.SetContentCachingClientListenRanges)
    res["contentCachingClientPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMacOSContentCachingClientPolicy , m.SetContentCachingClientPolicy)
    res["contentCachingDataPath"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContentCachingDataPath)
    res["contentCachingDisableConnectionSharing"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetContentCachingDisableConnectionSharing)
    res["contentCachingEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetContentCachingEnabled)
    res["contentCachingForceConnectionSharing"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetContentCachingForceConnectionSharing)
    res["contentCachingKeepAwake"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetContentCachingKeepAwake)
    res["contentCachingLogClientIdentities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetContentCachingLogClientIdentities)
    res["contentCachingMaxSizeBytes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetContentCachingMaxSizeBytes)
    res["contentCachingParents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetContentCachingParents)
    res["contentCachingParentSelectionPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMacOSContentCachingParentSelectionPolicy , m.SetContentCachingParentSelectionPolicy)
    res["contentCachingPeerFilterRanges"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIpRangeFromDiscriminatorValue , m.SetContentCachingPeerFilterRanges)
    res["contentCachingPeerListenRanges"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIpRangeFromDiscriminatorValue , m.SetContentCachingPeerListenRanges)
    res["contentCachingPeerPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMacOSContentCachingPeerPolicy , m.SetContentCachingPeerPolicy)
    res["contentCachingPort"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetContentCachingPort)
    res["contentCachingPublicRanges"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIpRangeFromDiscriminatorValue , m.SetContentCachingPublicRanges)
    res["contentCachingShowAlerts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetContentCachingShowAlerts)
    res["contentCachingType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMacOSContentCachingType , m.SetContentCachingType)
    res["loginWindowText"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLoginWindowText)
    res["logOutDisabledWhileLoggedIn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetLogOutDisabledWhileLoggedIn)
    res["macOSSingleSignOnExtension"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMacOSSingleSignOnExtensionFromDiscriminatorValue , m.SetMacOSSingleSignOnExtension)
    res["powerOffDisabledWhileLoggedIn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPowerOffDisabledWhileLoggedIn)
    res["restartDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRestartDisabled)
    res["restartDisabledWhileLoggedIn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRestartDisabledWhileLoggedIn)
    res["screenLockDisableImmediate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetScreenLockDisableImmediate)
    res["shutDownDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetShutDownDisabled)
    res["shutDownDisabledWhileLoggedIn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetShutDownDisabledWhileLoggedIn)
    res["singleSignOnExtension"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSingleSignOnExtensionFromDiscriminatorValue , m.SetSingleSignOnExtension)
    res["singleSignOnExtensionPkinitCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMacOSCertificateProfileBaseFromDiscriminatorValue , m.SetSingleSignOnExtensionPkinitCertificate)
    res["sleepDisabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSleepDisabled)
    return res
}
// GetLoginWindowText gets the loginWindowText property value. Custom text to be displayed on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetLoginWindowText()(*string) {
    return m.loginWindowText
}
// GetLogOutDisabledWhileLoggedIn gets the logOutDisabledWhileLoggedIn property value. Whether the Log Out menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) GetLogOutDisabledWhileLoggedIn()(*bool) {
    return m.logOutDisabledWhileLoggedIn
}
// GetMacOSSingleSignOnExtension gets the macOSSingleSignOnExtension property value. Gets or sets a single sign-on extension profile.
func (m *MacOSDeviceFeaturesConfiguration) GetMacOSSingleSignOnExtension()(MacOSSingleSignOnExtensionable) {
    return m.macOSSingleSignOnExtension
}
// GetPowerOffDisabledWhileLoggedIn gets the powerOffDisabledWhileLoggedIn property value. Whether the Power Off menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) GetPowerOffDisabledWhileLoggedIn()(*bool) {
    return m.powerOffDisabledWhileLoggedIn
}
// GetRestartDisabled gets the restartDisabled property value. Whether to hide the Restart button item on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetRestartDisabled()(*bool) {
    return m.restartDisabled
}
// GetRestartDisabledWhileLoggedIn gets the restartDisabledWhileLoggedIn property value. Whether the Restart menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) GetRestartDisabledWhileLoggedIn()(*bool) {
    return m.restartDisabledWhileLoggedIn
}
// GetScreenLockDisableImmediate gets the screenLockDisableImmediate property value. Whether to disable the immediate screen lock functions.
func (m *MacOSDeviceFeaturesConfiguration) GetScreenLockDisableImmediate()(*bool) {
    return m.screenLockDisableImmediate
}
// GetShutDownDisabled gets the shutDownDisabled property value. Whether to hide the Shut Down button item on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetShutDownDisabled()(*bool) {
    return m.shutDownDisabled
}
// GetShutDownDisabledWhileLoggedIn gets the shutDownDisabledWhileLoggedIn property value. Whether the Shut Down menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) GetShutDownDisabledWhileLoggedIn()(*bool) {
    return m.shutDownDisabledWhileLoggedIn
}
// GetSingleSignOnExtension gets the singleSignOnExtension property value. Gets or sets a single sign-on extension profile. Deprecated: use MacOSSingleSignOnExtension instead.
func (m *MacOSDeviceFeaturesConfiguration) GetSingleSignOnExtension()(SingleSignOnExtensionable) {
    return m.singleSignOnExtension
}
// GetSingleSignOnExtensionPkinitCertificate gets the singleSignOnExtensionPkinitCertificate property value. PKINIT Certificate for the authentication with single sign-on extensions.
func (m *MacOSDeviceFeaturesConfiguration) GetSingleSignOnExtensionPkinitCertificate()(MacOSCertificateProfileBaseable) {
    return m.singleSignOnExtensionPkinitCertificate
}
// GetSleepDisabled gets the sleepDisabled property value. Whether to hide the Sleep menu item on the login window.
func (m *MacOSDeviceFeaturesConfiguration) GetSleepDisabled()(*bool) {
    return m.sleepDisabled
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppAssociatedDomains())
        err = writer.WriteCollectionOfObjectValues("appAssociatedDomains", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssociatedDomains() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssociatedDomains())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAutoLaunchItems())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetContentCachingClientListenRanges())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetContentCachingPeerFilterRanges())
        err = writer.WriteCollectionOfObjectValues("contentCachingPeerFilterRanges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContentCachingPeerListenRanges() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetContentCachingPeerListenRanges())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetContentCachingPublicRanges())
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
    m.adminShowHostInfo = value
}
// SetAppAssociatedDomains sets the appAssociatedDomains property value. Gets or sets a list that maps apps to their associated domains. Application identifiers must be unique. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetAppAssociatedDomains(value []MacOSAssociatedDomainsItemable)() {
    m.appAssociatedDomains = value
}
// SetAssociatedDomains sets the associatedDomains property value. DEPRECATED: use appAssociatedDomains instead. Gets or sets a list that maps apps to their associated domains. The key should match the app's ID, and the value should be a string in the form of 'service:domain' where domain is a fully qualified hostname (e.g. webcredentials:example.com). This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetAssociatedDomains(value []KeyValuePairable)() {
    m.associatedDomains = value
}
// SetAuthorizedUsersListHidden sets the authorizedUsersListHidden property value. Whether to show the name and password dialog or a list of users on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAuthorizedUsersListHidden(value *bool)() {
    m.authorizedUsersListHidden = value
}
// SetAuthorizedUsersListHideAdminUsers sets the authorizedUsersListHideAdminUsers property value. Whether to hide admin users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAuthorizedUsersListHideAdminUsers(value *bool)() {
    m.authorizedUsersListHideAdminUsers = value
}
// SetAuthorizedUsersListHideLocalUsers sets the authorizedUsersListHideLocalUsers property value. Whether to show only network and system users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAuthorizedUsersListHideLocalUsers(value *bool)() {
    m.authorizedUsersListHideLocalUsers = value
}
// SetAuthorizedUsersListHideMobileAccounts sets the authorizedUsersListHideMobileAccounts property value. Whether to hide mobile users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAuthorizedUsersListHideMobileAccounts(value *bool)() {
    m.authorizedUsersListHideMobileAccounts = value
}
// SetAuthorizedUsersListIncludeNetworkUsers sets the authorizedUsersListIncludeNetworkUsers property value. Whether to show network users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAuthorizedUsersListIncludeNetworkUsers(value *bool)() {
    m.authorizedUsersListIncludeNetworkUsers = value
}
// SetAuthorizedUsersListShowOtherManagedUsers sets the authorizedUsersListShowOtherManagedUsers property value. Whether to show other users in the authorized users list on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetAuthorizedUsersListShowOtherManagedUsers(value *bool)() {
    m.authorizedUsersListShowOtherManagedUsers = value
}
// SetAutoLaunchItems sets the autoLaunchItems property value. List of applications, files, folders, and other items to launch when the user logs in. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetAutoLaunchItems(value []MacOSLaunchItemable)() {
    m.autoLaunchItems = value
}
// SetConsoleAccessDisabled sets the consoleAccessDisabled property value. Whether the Other user will disregard use of the console special user name.
func (m *MacOSDeviceFeaturesConfiguration) SetConsoleAccessDisabled(value *bool)() {
    m.consoleAccessDisabled = value
}
// SetContentCachingBlockDeletion sets the contentCachingBlockDeletion property value. Prevents content caches from purging content to free up disk space for other apps.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingBlockDeletion(value *bool)() {
    m.contentCachingBlockDeletion = value
}
// SetContentCachingClientListenRanges sets the contentCachingClientListenRanges property value. A list of custom IP ranges content caches will use to listen for clients. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingClientListenRanges(value []IpRangeable)() {
    m.contentCachingClientListenRanges = value
}
// SetContentCachingClientPolicy sets the contentCachingClientPolicy property value. Determines which clients a content cache will serve.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingClientPolicy(value *MacOSContentCachingClientPolicy)() {
    m.contentCachingClientPolicy = value
}
// SetContentCachingDataPath sets the contentCachingDataPath property value. The path to the directory used to store cached content. The value must be (or end with) /Library/Application Support/Apple/AssetCache/Data
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingDataPath(value *string)() {
    m.contentCachingDataPath = value
}
// SetContentCachingDisableConnectionSharing sets the contentCachingDisableConnectionSharing property value. Disables internet connection sharing.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingDisableConnectionSharing(value *bool)() {
    m.contentCachingDisableConnectionSharing = value
}
// SetContentCachingEnabled sets the contentCachingEnabled property value. Enables content caching and prevents it from being disabled by the user.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingEnabled(value *bool)() {
    m.contentCachingEnabled = value
}
// SetContentCachingForceConnectionSharing sets the contentCachingForceConnectionSharing property value. Forces internet connection sharing. contentCachingDisableConnectionSharing overrides this setting.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingForceConnectionSharing(value *bool)() {
    m.contentCachingForceConnectionSharing = value
}
// SetContentCachingKeepAwake sets the contentCachingKeepAwake property value. Prevent the device from sleeping if content caching is enabled.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingKeepAwake(value *bool)() {
    m.contentCachingKeepAwake = value
}
// SetContentCachingLogClientIdentities sets the contentCachingLogClientIdentities property value. Enables logging of IP addresses and ports of clients that request cached content.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingLogClientIdentities(value *bool)() {
    m.contentCachingLogClientIdentities = value
}
// SetContentCachingMaxSizeBytes sets the contentCachingMaxSizeBytes property value. The maximum number of bytes of disk space that will be used for the content cache. A value of 0 (default) indicates unlimited disk space.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingMaxSizeBytes(value *int64)() {
    m.contentCachingMaxSizeBytes = value
}
// SetContentCachingParents sets the contentCachingParents property value. A list of IP addresses representing parent content caches.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingParents(value []string)() {
    m.contentCachingParents = value
}
// SetContentCachingParentSelectionPolicy sets the contentCachingParentSelectionPolicy property value. Determines how content caches select a parent cache.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingParentSelectionPolicy(value *MacOSContentCachingParentSelectionPolicy)() {
    m.contentCachingParentSelectionPolicy = value
}
// SetContentCachingPeerFilterRanges sets the contentCachingPeerFilterRanges property value. A list of custom IP ranges content caches will use to query for content from peers caches. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingPeerFilterRanges(value []IpRangeable)() {
    m.contentCachingPeerFilterRanges = value
}
// SetContentCachingPeerListenRanges sets the contentCachingPeerListenRanges property value. A list of custom IP ranges content caches will use to listen for peer caches. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingPeerListenRanges(value []IpRangeable)() {
    m.contentCachingPeerListenRanges = value
}
// SetContentCachingPeerPolicy sets the contentCachingPeerPolicy property value. Determines which content caches other content caches will peer with.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingPeerPolicy(value *MacOSContentCachingPeerPolicy)() {
    m.contentCachingPeerPolicy = value
}
// SetContentCachingPort sets the contentCachingPort property value. Sets the port used for content caching. If the value is 0, a random available port will be selected. Valid values 0 to 65535
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingPort(value *int32)() {
    m.contentCachingPort = value
}
// SetContentCachingPublicRanges sets the contentCachingPublicRanges property value. A list of custom IP ranges that Apple's content caching service should use to match clients to content caches. This collection can contain a maximum of 500 elements.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingPublicRanges(value []IpRangeable)() {
    m.contentCachingPublicRanges = value
}
// SetContentCachingShowAlerts sets the contentCachingShowAlerts property value. Display content caching alerts as system notifications.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingShowAlerts(value *bool)() {
    m.contentCachingShowAlerts = value
}
// SetContentCachingType sets the contentCachingType property value. Indicates the type of content allowed to be cached by Apple's content caching service.
func (m *MacOSDeviceFeaturesConfiguration) SetContentCachingType(value *MacOSContentCachingType)() {
    m.contentCachingType = value
}
// SetLoginWindowText sets the loginWindowText property value. Custom text to be displayed on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetLoginWindowText(value *string)() {
    m.loginWindowText = value
}
// SetLogOutDisabledWhileLoggedIn sets the logOutDisabledWhileLoggedIn property value. Whether the Log Out menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) SetLogOutDisabledWhileLoggedIn(value *bool)() {
    m.logOutDisabledWhileLoggedIn = value
}
// SetMacOSSingleSignOnExtension sets the macOSSingleSignOnExtension property value. Gets or sets a single sign-on extension profile.
func (m *MacOSDeviceFeaturesConfiguration) SetMacOSSingleSignOnExtension(value MacOSSingleSignOnExtensionable)() {
    m.macOSSingleSignOnExtension = value
}
// SetPowerOffDisabledWhileLoggedIn sets the powerOffDisabledWhileLoggedIn property value. Whether the Power Off menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) SetPowerOffDisabledWhileLoggedIn(value *bool)() {
    m.powerOffDisabledWhileLoggedIn = value
}
// SetRestartDisabled sets the restartDisabled property value. Whether to hide the Restart button item on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetRestartDisabled(value *bool)() {
    m.restartDisabled = value
}
// SetRestartDisabledWhileLoggedIn sets the restartDisabledWhileLoggedIn property value. Whether the Restart menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) SetRestartDisabledWhileLoggedIn(value *bool)() {
    m.restartDisabledWhileLoggedIn = value
}
// SetScreenLockDisableImmediate sets the screenLockDisableImmediate property value. Whether to disable the immediate screen lock functions.
func (m *MacOSDeviceFeaturesConfiguration) SetScreenLockDisableImmediate(value *bool)() {
    m.screenLockDisableImmediate = value
}
// SetShutDownDisabled sets the shutDownDisabled property value. Whether to hide the Shut Down button item on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetShutDownDisabled(value *bool)() {
    m.shutDownDisabled = value
}
// SetShutDownDisabledWhileLoggedIn sets the shutDownDisabledWhileLoggedIn property value. Whether the Shut Down menu item on the login window will be disabled while the user is logged in.
func (m *MacOSDeviceFeaturesConfiguration) SetShutDownDisabledWhileLoggedIn(value *bool)() {
    m.shutDownDisabledWhileLoggedIn = value
}
// SetSingleSignOnExtension sets the singleSignOnExtension property value. Gets or sets a single sign-on extension profile. Deprecated: use MacOSSingleSignOnExtension instead.
func (m *MacOSDeviceFeaturesConfiguration) SetSingleSignOnExtension(value SingleSignOnExtensionable)() {
    m.singleSignOnExtension = value
}
// SetSingleSignOnExtensionPkinitCertificate sets the singleSignOnExtensionPkinitCertificate property value. PKINIT Certificate for the authentication with single sign-on extensions.
func (m *MacOSDeviceFeaturesConfiguration) SetSingleSignOnExtensionPkinitCertificate(value MacOSCertificateProfileBaseable)() {
    m.singleSignOnExtensionPkinitCertificate = value
}
// SetSleepDisabled sets the sleepDisabled property value. Whether to hide the Sleep menu item on the login window.
func (m *MacOSDeviceFeaturesConfiguration) SetSleepDisabled(value *bool)() {
    m.sleepDisabled = value
}
