package tenantadmin

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Settings 
type Settings struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Collection of trusted domain GUIDs for the OneDrive sync app.
    allowedDomainGuidsForSyncApp []string
    // Collection of managed paths available for site creation. Read-only.
    availableManagedPathsForSiteCreation []string
    // The number of days for preserving a deleted user's OneDrive.
    deletedUserPersonalSiteRetentionPeriodInDays *int32
    // Collection of file extensions not uploaded by the OneDrive sync app.
    excludedFileExtensionsForSyncApp []string
    // Specifies the idle session sign-out policies for the tenant.
    idleSessionSignOut IdleSessionSignOutable
    // Specifies the image tagging option for the tenant. Possible values are: disabled, basic, enhanced.
    imageTaggingOption *ImageTaggingChoice
    // Indicates whether comments are allowed on modern site pages in SharePoint.
    isCommentingOnSitePagesEnabled *bool
    // Indicates whether push notifications are enabled for OneDrive events.
    isFileActivityNotificationEnabled *bool
    // Indicates whether legacy authentication protocols are enabled for the tenant.
    isLegacyAuthProtocolsEnabled *bool
    // Indicates whetherif Fluid Framework is allowed on SharePoint sites.
    isLoopEnabled *bool
    // Indicates whether files can be synced using the OneDrive sync app for Mac.
    isMacSyncAppEnabled *bool
    // Indicates whether guests must sign in using the same account to which sharing invitations are sent.
    isRequireAcceptingUserToMatchInvitedUserEnabled *bool
    // Indicates whether guests are allowed to reshare files, folders, and sites they don't own.
    isResharingByExternalUsersEnabled *bool
    // Indicates whether mobile push notifications are enabled for SharePoint.
    isSharePointMobileNotificationEnabled *bool
    // Indicates whether the newsfeed is allowed on the modern site pages in SharePoint.
    isSharePointNewsfeedEnabled *bool
    // Indicates whether users are allowed to create sites.
    isSiteCreationEnabled *bool
    // Indicates whether the UI commands for creating sites are shown.
    isSiteCreationUIEnabled *bool
    // Indicates whether creating new modern pages is allowed on SharePoint sites.
    isSitePagesCreationEnabled *bool
    // Indicates whether site storage space is automatically managed or if specific storage limits are set per site.
    isSitesStorageLimitAutomatic *bool
    // Indicates whether the sync button in OneDrive is hidden.
    isSyncButtonHiddenOnPersonalSite *bool
    // Indicates whether users are allowed to sync files only on PCs joined to specific domains.
    isUnmanagedSyncAppForTenantRestricted *bool
    // The default OneDrive storage limit for all new and existing users who are assigned a qualifying license. Measured in megabytes (MB).
    personalSiteDefaultStorageLimitInMB *int64
    // Collection of email domains that are allowed for sharing outside the organization.
    sharingAllowedDomainList []string
    // Collection of email domains that are blocked for sharing outside the organization.
    sharingBlockedDomainList []string
    // Sharing capability for the tenant. Possible values are: disabled, externalUserSharingOnly, externalUserAndGuestSharing, existingExternalUserSharingOnly.
    sharingCapability *SharingCapabilities
    // Specifies the external sharing mode for domains. Possible values are: none, allowList, blockList.
    sharingDomainRestrictionMode *SharingDomainRestrictionMode
    // The value of the team site managed path. This is the path under which new team sites will be created.
    siteCreationDefaultManagedPath *string
    // The default storage quota for a new site upon creation. Measured in megabytes (MB).
    siteCreationDefaultStorageLimitInMB *int32
    // The default timezone of a tenant for newly created sites.
    tenantDefaultTimezone *string
}
// NewSettings instantiates a new settings and sets the default values.
func NewSettings()(*Settings) {
    m := &Settings{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.tenantAdmin.settings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSettings(), nil
}
// GetAllowedDomainGuidsForSyncApp gets the allowedDomainGuidsForSyncApp property value. Collection of trusted domain GUIDs for the OneDrive sync app.
func (m *Settings) GetAllowedDomainGuidsForSyncApp()([]string) {
    return m.allowedDomainGuidsForSyncApp
}
// GetAvailableManagedPathsForSiteCreation gets the availableManagedPathsForSiteCreation property value. Collection of managed paths available for site creation. Read-only.
func (m *Settings) GetAvailableManagedPathsForSiteCreation()([]string) {
    return m.availableManagedPathsForSiteCreation
}
// GetDeletedUserPersonalSiteRetentionPeriodInDays gets the deletedUserPersonalSiteRetentionPeriodInDays property value. The number of days for preserving a deleted user's OneDrive.
func (m *Settings) GetDeletedUserPersonalSiteRetentionPeriodInDays()(*int32) {
    return m.deletedUserPersonalSiteRetentionPeriodInDays
}
// GetExcludedFileExtensionsForSyncApp gets the excludedFileExtensionsForSyncApp property value. Collection of file extensions not uploaded by the OneDrive sync app.
func (m *Settings) GetExcludedFileExtensionsForSyncApp()([]string) {
    return m.excludedFileExtensionsForSyncApp
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Settings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedDomainGuidsForSyncApp"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAllowedDomainGuidsForSyncApp)
    res["availableManagedPathsForSiteCreation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetAvailableManagedPathsForSiteCreation)
    res["deletedUserPersonalSiteRetentionPeriodInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDeletedUserPersonalSiteRetentionPeriodInDays)
    res["excludedFileExtensionsForSyncApp"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetExcludedFileExtensionsForSyncApp)
    res["idleSessionSignOut"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdleSessionSignOutFromDiscriminatorValue , m.SetIdleSessionSignOut)
    res["imageTaggingOption"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseImageTaggingChoice , m.SetImageTaggingOption)
    res["isCommentingOnSitePagesEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsCommentingOnSitePagesEnabled)
    res["isFileActivityNotificationEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsFileActivityNotificationEnabled)
    res["isLegacyAuthProtocolsEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsLegacyAuthProtocolsEnabled)
    res["isLoopEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsLoopEnabled)
    res["isMacSyncAppEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsMacSyncAppEnabled)
    res["isRequireAcceptingUserToMatchInvitedUserEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsRequireAcceptingUserToMatchInvitedUserEnabled)
    res["isResharingByExternalUsersEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsResharingByExternalUsersEnabled)
    res["isSharePointMobileNotificationEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSharePointMobileNotificationEnabled)
    res["isSharePointNewsfeedEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSharePointNewsfeedEnabled)
    res["isSiteCreationEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSiteCreationEnabled)
    res["isSiteCreationUIEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSiteCreationUIEnabled)
    res["isSitePagesCreationEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSitePagesCreationEnabled)
    res["isSitesStorageLimitAutomatic"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSitesStorageLimitAutomatic)
    res["isSyncButtonHiddenOnPersonalSite"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSyncButtonHiddenOnPersonalSite)
    res["isUnmanagedSyncAppForTenantRestricted"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsUnmanagedSyncAppForTenantRestricted)
    res["personalSiteDefaultStorageLimitInMB"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetPersonalSiteDefaultStorageLimitInMB)
    res["sharingAllowedDomainList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSharingAllowedDomainList)
    res["sharingBlockedDomainList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSharingBlockedDomainList)
    res["sharingCapability"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSharingCapabilities , m.SetSharingCapability)
    res["sharingDomainRestrictionMode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSharingDomainRestrictionMode , m.SetSharingDomainRestrictionMode)
    res["siteCreationDefaultManagedPath"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSiteCreationDefaultManagedPath)
    res["siteCreationDefaultStorageLimitInMB"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSiteCreationDefaultStorageLimitInMB)
    res["tenantDefaultTimezone"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantDefaultTimezone)
    return res
}
// GetIdleSessionSignOut gets the idleSessionSignOut property value. Specifies the idle session sign-out policies for the tenant.
func (m *Settings) GetIdleSessionSignOut()(IdleSessionSignOutable) {
    return m.idleSessionSignOut
}
// GetImageTaggingOption gets the imageTaggingOption property value. Specifies the image tagging option for the tenant. Possible values are: disabled, basic, enhanced.
func (m *Settings) GetImageTaggingOption()(*ImageTaggingChoice) {
    return m.imageTaggingOption
}
// GetIsCommentingOnSitePagesEnabled gets the isCommentingOnSitePagesEnabled property value. Indicates whether comments are allowed on modern site pages in SharePoint.
func (m *Settings) GetIsCommentingOnSitePagesEnabled()(*bool) {
    return m.isCommentingOnSitePagesEnabled
}
// GetIsFileActivityNotificationEnabled gets the isFileActivityNotificationEnabled property value. Indicates whether push notifications are enabled for OneDrive events.
func (m *Settings) GetIsFileActivityNotificationEnabled()(*bool) {
    return m.isFileActivityNotificationEnabled
}
// GetIsLegacyAuthProtocolsEnabled gets the isLegacyAuthProtocolsEnabled property value. Indicates whether legacy authentication protocols are enabled for the tenant.
func (m *Settings) GetIsLegacyAuthProtocolsEnabled()(*bool) {
    return m.isLegacyAuthProtocolsEnabled
}
// GetIsLoopEnabled gets the isLoopEnabled property value. Indicates whetherif Fluid Framework is allowed on SharePoint sites.
func (m *Settings) GetIsLoopEnabled()(*bool) {
    return m.isLoopEnabled
}
// GetIsMacSyncAppEnabled gets the isMacSyncAppEnabled property value. Indicates whether files can be synced using the OneDrive sync app for Mac.
func (m *Settings) GetIsMacSyncAppEnabled()(*bool) {
    return m.isMacSyncAppEnabled
}
// GetIsRequireAcceptingUserToMatchInvitedUserEnabled gets the isRequireAcceptingUserToMatchInvitedUserEnabled property value. Indicates whether guests must sign in using the same account to which sharing invitations are sent.
func (m *Settings) GetIsRequireAcceptingUserToMatchInvitedUserEnabled()(*bool) {
    return m.isRequireAcceptingUserToMatchInvitedUserEnabled
}
// GetIsResharingByExternalUsersEnabled gets the isResharingByExternalUsersEnabled property value. Indicates whether guests are allowed to reshare files, folders, and sites they don't own.
func (m *Settings) GetIsResharingByExternalUsersEnabled()(*bool) {
    return m.isResharingByExternalUsersEnabled
}
// GetIsSharePointMobileNotificationEnabled gets the isSharePointMobileNotificationEnabled property value. Indicates whether mobile push notifications are enabled for SharePoint.
func (m *Settings) GetIsSharePointMobileNotificationEnabled()(*bool) {
    return m.isSharePointMobileNotificationEnabled
}
// GetIsSharePointNewsfeedEnabled gets the isSharePointNewsfeedEnabled property value. Indicates whether the newsfeed is allowed on the modern site pages in SharePoint.
func (m *Settings) GetIsSharePointNewsfeedEnabled()(*bool) {
    return m.isSharePointNewsfeedEnabled
}
// GetIsSiteCreationEnabled gets the isSiteCreationEnabled property value. Indicates whether users are allowed to create sites.
func (m *Settings) GetIsSiteCreationEnabled()(*bool) {
    return m.isSiteCreationEnabled
}
// GetIsSiteCreationUIEnabled gets the isSiteCreationUIEnabled property value. Indicates whether the UI commands for creating sites are shown.
func (m *Settings) GetIsSiteCreationUIEnabled()(*bool) {
    return m.isSiteCreationUIEnabled
}
// GetIsSitePagesCreationEnabled gets the isSitePagesCreationEnabled property value. Indicates whether creating new modern pages is allowed on SharePoint sites.
func (m *Settings) GetIsSitePagesCreationEnabled()(*bool) {
    return m.isSitePagesCreationEnabled
}
// GetIsSitesStorageLimitAutomatic gets the isSitesStorageLimitAutomatic property value. Indicates whether site storage space is automatically managed or if specific storage limits are set per site.
func (m *Settings) GetIsSitesStorageLimitAutomatic()(*bool) {
    return m.isSitesStorageLimitAutomatic
}
// GetIsSyncButtonHiddenOnPersonalSite gets the isSyncButtonHiddenOnPersonalSite property value. Indicates whether the sync button in OneDrive is hidden.
func (m *Settings) GetIsSyncButtonHiddenOnPersonalSite()(*bool) {
    return m.isSyncButtonHiddenOnPersonalSite
}
// GetIsUnmanagedSyncAppForTenantRestricted gets the isUnmanagedSyncAppForTenantRestricted property value. Indicates whether users are allowed to sync files only on PCs joined to specific domains.
func (m *Settings) GetIsUnmanagedSyncAppForTenantRestricted()(*bool) {
    return m.isUnmanagedSyncAppForTenantRestricted
}
// GetPersonalSiteDefaultStorageLimitInMB gets the personalSiteDefaultStorageLimitInMB property value. The default OneDrive storage limit for all new and existing users who are assigned a qualifying license. Measured in megabytes (MB).
func (m *Settings) GetPersonalSiteDefaultStorageLimitInMB()(*int64) {
    return m.personalSiteDefaultStorageLimitInMB
}
// GetSharingAllowedDomainList gets the sharingAllowedDomainList property value. Collection of email domains that are allowed for sharing outside the organization.
func (m *Settings) GetSharingAllowedDomainList()([]string) {
    return m.sharingAllowedDomainList
}
// GetSharingBlockedDomainList gets the sharingBlockedDomainList property value. Collection of email domains that are blocked for sharing outside the organization.
func (m *Settings) GetSharingBlockedDomainList()([]string) {
    return m.sharingBlockedDomainList
}
// GetSharingCapability gets the sharingCapability property value. Sharing capability for the tenant. Possible values are: disabled, externalUserSharingOnly, externalUserAndGuestSharing, existingExternalUserSharingOnly.
func (m *Settings) GetSharingCapability()(*SharingCapabilities) {
    return m.sharingCapability
}
// GetSharingDomainRestrictionMode gets the sharingDomainRestrictionMode property value. Specifies the external sharing mode for domains. Possible values are: none, allowList, blockList.
func (m *Settings) GetSharingDomainRestrictionMode()(*SharingDomainRestrictionMode) {
    return m.sharingDomainRestrictionMode
}
// GetSiteCreationDefaultManagedPath gets the siteCreationDefaultManagedPath property value. The value of the team site managed path. This is the path under which new team sites will be created.
func (m *Settings) GetSiteCreationDefaultManagedPath()(*string) {
    return m.siteCreationDefaultManagedPath
}
// GetSiteCreationDefaultStorageLimitInMB gets the siteCreationDefaultStorageLimitInMB property value. The default storage quota for a new site upon creation. Measured in megabytes (MB).
func (m *Settings) GetSiteCreationDefaultStorageLimitInMB()(*int32) {
    return m.siteCreationDefaultStorageLimitInMB
}
// GetTenantDefaultTimezone gets the tenantDefaultTimezone property value. The default timezone of a tenant for newly created sites.
func (m *Settings) GetTenantDefaultTimezone()(*string) {
    return m.tenantDefaultTimezone
}
// Serialize serializes information the current object
func (m *Settings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedDomainGuidsForSyncApp() != nil {
        err = writer.WriteCollectionOfStringValues("allowedDomainGuidsForSyncApp", m.GetAllowedDomainGuidsForSyncApp())
        if err != nil {
            return err
        }
    }
    if m.GetAvailableManagedPathsForSiteCreation() != nil {
        err = writer.WriteCollectionOfStringValues("availableManagedPathsForSiteCreation", m.GetAvailableManagedPathsForSiteCreation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deletedUserPersonalSiteRetentionPeriodInDays", m.GetDeletedUserPersonalSiteRetentionPeriodInDays())
        if err != nil {
            return err
        }
    }
    if m.GetExcludedFileExtensionsForSyncApp() != nil {
        err = writer.WriteCollectionOfStringValues("excludedFileExtensionsForSyncApp", m.GetExcludedFileExtensionsForSyncApp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("idleSessionSignOut", m.GetIdleSessionSignOut())
        if err != nil {
            return err
        }
    }
    if m.GetImageTaggingOption() != nil {
        cast := (*m.GetImageTaggingOption()).String()
        err = writer.WriteStringValue("imageTaggingOption", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCommentingOnSitePagesEnabled", m.GetIsCommentingOnSitePagesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFileActivityNotificationEnabled", m.GetIsFileActivityNotificationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLegacyAuthProtocolsEnabled", m.GetIsLegacyAuthProtocolsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLoopEnabled", m.GetIsLoopEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMacSyncAppEnabled", m.GetIsMacSyncAppEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRequireAcceptingUserToMatchInvitedUserEnabled", m.GetIsRequireAcceptingUserToMatchInvitedUserEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isResharingByExternalUsersEnabled", m.GetIsResharingByExternalUsersEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSharePointMobileNotificationEnabled", m.GetIsSharePointMobileNotificationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSharePointNewsfeedEnabled", m.GetIsSharePointNewsfeedEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSiteCreationEnabled", m.GetIsSiteCreationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSiteCreationUIEnabled", m.GetIsSiteCreationUIEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSitePagesCreationEnabled", m.GetIsSitePagesCreationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSitesStorageLimitAutomatic", m.GetIsSitesStorageLimitAutomatic())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSyncButtonHiddenOnPersonalSite", m.GetIsSyncButtonHiddenOnPersonalSite())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isUnmanagedSyncAppForTenantRestricted", m.GetIsUnmanagedSyncAppForTenantRestricted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("personalSiteDefaultStorageLimitInMB", m.GetPersonalSiteDefaultStorageLimitInMB())
        if err != nil {
            return err
        }
    }
    if m.GetSharingAllowedDomainList() != nil {
        err = writer.WriteCollectionOfStringValues("sharingAllowedDomainList", m.GetSharingAllowedDomainList())
        if err != nil {
            return err
        }
    }
    if m.GetSharingBlockedDomainList() != nil {
        err = writer.WriteCollectionOfStringValues("sharingBlockedDomainList", m.GetSharingBlockedDomainList())
        if err != nil {
            return err
        }
    }
    if m.GetSharingCapability() != nil {
        cast := (*m.GetSharingCapability()).String()
        err = writer.WriteStringValue("sharingCapability", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSharingDomainRestrictionMode() != nil {
        cast := (*m.GetSharingDomainRestrictionMode()).String()
        err = writer.WriteStringValue("sharingDomainRestrictionMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("siteCreationDefaultManagedPath", m.GetSiteCreationDefaultManagedPath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("siteCreationDefaultStorageLimitInMB", m.GetSiteCreationDefaultStorageLimitInMB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDefaultTimezone", m.GetTenantDefaultTimezone())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedDomainGuidsForSyncApp sets the allowedDomainGuidsForSyncApp property value. Collection of trusted domain GUIDs for the OneDrive sync app.
func (m *Settings) SetAllowedDomainGuidsForSyncApp(value []string)() {
    m.allowedDomainGuidsForSyncApp = value
}
// SetAvailableManagedPathsForSiteCreation sets the availableManagedPathsForSiteCreation property value. Collection of managed paths available for site creation. Read-only.
func (m *Settings) SetAvailableManagedPathsForSiteCreation(value []string)() {
    m.availableManagedPathsForSiteCreation = value
}
// SetDeletedUserPersonalSiteRetentionPeriodInDays sets the deletedUserPersonalSiteRetentionPeriodInDays property value. The number of days for preserving a deleted user's OneDrive.
func (m *Settings) SetDeletedUserPersonalSiteRetentionPeriodInDays(value *int32)() {
    m.deletedUserPersonalSiteRetentionPeriodInDays = value
}
// SetExcludedFileExtensionsForSyncApp sets the excludedFileExtensionsForSyncApp property value. Collection of file extensions not uploaded by the OneDrive sync app.
func (m *Settings) SetExcludedFileExtensionsForSyncApp(value []string)() {
    m.excludedFileExtensionsForSyncApp = value
}
// SetIdleSessionSignOut sets the idleSessionSignOut property value. Specifies the idle session sign-out policies for the tenant.
func (m *Settings) SetIdleSessionSignOut(value IdleSessionSignOutable)() {
    m.idleSessionSignOut = value
}
// SetImageTaggingOption sets the imageTaggingOption property value. Specifies the image tagging option for the tenant. Possible values are: disabled, basic, enhanced.
func (m *Settings) SetImageTaggingOption(value *ImageTaggingChoice)() {
    m.imageTaggingOption = value
}
// SetIsCommentingOnSitePagesEnabled sets the isCommentingOnSitePagesEnabled property value. Indicates whether comments are allowed on modern site pages in SharePoint.
func (m *Settings) SetIsCommentingOnSitePagesEnabled(value *bool)() {
    m.isCommentingOnSitePagesEnabled = value
}
// SetIsFileActivityNotificationEnabled sets the isFileActivityNotificationEnabled property value. Indicates whether push notifications are enabled for OneDrive events.
func (m *Settings) SetIsFileActivityNotificationEnabled(value *bool)() {
    m.isFileActivityNotificationEnabled = value
}
// SetIsLegacyAuthProtocolsEnabled sets the isLegacyAuthProtocolsEnabled property value. Indicates whether legacy authentication protocols are enabled for the tenant.
func (m *Settings) SetIsLegacyAuthProtocolsEnabled(value *bool)() {
    m.isLegacyAuthProtocolsEnabled = value
}
// SetIsLoopEnabled sets the isLoopEnabled property value. Indicates whetherif Fluid Framework is allowed on SharePoint sites.
func (m *Settings) SetIsLoopEnabled(value *bool)() {
    m.isLoopEnabled = value
}
// SetIsMacSyncAppEnabled sets the isMacSyncAppEnabled property value. Indicates whether files can be synced using the OneDrive sync app for Mac.
func (m *Settings) SetIsMacSyncAppEnabled(value *bool)() {
    m.isMacSyncAppEnabled = value
}
// SetIsRequireAcceptingUserToMatchInvitedUserEnabled sets the isRequireAcceptingUserToMatchInvitedUserEnabled property value. Indicates whether guests must sign in using the same account to which sharing invitations are sent.
func (m *Settings) SetIsRequireAcceptingUserToMatchInvitedUserEnabled(value *bool)() {
    m.isRequireAcceptingUserToMatchInvitedUserEnabled = value
}
// SetIsResharingByExternalUsersEnabled sets the isResharingByExternalUsersEnabled property value. Indicates whether guests are allowed to reshare files, folders, and sites they don't own.
func (m *Settings) SetIsResharingByExternalUsersEnabled(value *bool)() {
    m.isResharingByExternalUsersEnabled = value
}
// SetIsSharePointMobileNotificationEnabled sets the isSharePointMobileNotificationEnabled property value. Indicates whether mobile push notifications are enabled for SharePoint.
func (m *Settings) SetIsSharePointMobileNotificationEnabled(value *bool)() {
    m.isSharePointMobileNotificationEnabled = value
}
// SetIsSharePointNewsfeedEnabled sets the isSharePointNewsfeedEnabled property value. Indicates whether the newsfeed is allowed on the modern site pages in SharePoint.
func (m *Settings) SetIsSharePointNewsfeedEnabled(value *bool)() {
    m.isSharePointNewsfeedEnabled = value
}
// SetIsSiteCreationEnabled sets the isSiteCreationEnabled property value. Indicates whether users are allowed to create sites.
func (m *Settings) SetIsSiteCreationEnabled(value *bool)() {
    m.isSiteCreationEnabled = value
}
// SetIsSiteCreationUIEnabled sets the isSiteCreationUIEnabled property value. Indicates whether the UI commands for creating sites are shown.
func (m *Settings) SetIsSiteCreationUIEnabled(value *bool)() {
    m.isSiteCreationUIEnabled = value
}
// SetIsSitePagesCreationEnabled sets the isSitePagesCreationEnabled property value. Indicates whether creating new modern pages is allowed on SharePoint sites.
func (m *Settings) SetIsSitePagesCreationEnabled(value *bool)() {
    m.isSitePagesCreationEnabled = value
}
// SetIsSitesStorageLimitAutomatic sets the isSitesStorageLimitAutomatic property value. Indicates whether site storage space is automatically managed or if specific storage limits are set per site.
func (m *Settings) SetIsSitesStorageLimitAutomatic(value *bool)() {
    m.isSitesStorageLimitAutomatic = value
}
// SetIsSyncButtonHiddenOnPersonalSite sets the isSyncButtonHiddenOnPersonalSite property value. Indicates whether the sync button in OneDrive is hidden.
func (m *Settings) SetIsSyncButtonHiddenOnPersonalSite(value *bool)() {
    m.isSyncButtonHiddenOnPersonalSite = value
}
// SetIsUnmanagedSyncAppForTenantRestricted sets the isUnmanagedSyncAppForTenantRestricted property value. Indicates whether users are allowed to sync files only on PCs joined to specific domains.
func (m *Settings) SetIsUnmanagedSyncAppForTenantRestricted(value *bool)() {
    m.isUnmanagedSyncAppForTenantRestricted = value
}
// SetPersonalSiteDefaultStorageLimitInMB sets the personalSiteDefaultStorageLimitInMB property value. The default OneDrive storage limit for all new and existing users who are assigned a qualifying license. Measured in megabytes (MB).
func (m *Settings) SetPersonalSiteDefaultStorageLimitInMB(value *int64)() {
    m.personalSiteDefaultStorageLimitInMB = value
}
// SetSharingAllowedDomainList sets the sharingAllowedDomainList property value. Collection of email domains that are allowed for sharing outside the organization.
func (m *Settings) SetSharingAllowedDomainList(value []string)() {
    m.sharingAllowedDomainList = value
}
// SetSharingBlockedDomainList sets the sharingBlockedDomainList property value. Collection of email domains that are blocked for sharing outside the organization.
func (m *Settings) SetSharingBlockedDomainList(value []string)() {
    m.sharingBlockedDomainList = value
}
// SetSharingCapability sets the sharingCapability property value. Sharing capability for the tenant. Possible values are: disabled, externalUserSharingOnly, externalUserAndGuestSharing, existingExternalUserSharingOnly.
func (m *Settings) SetSharingCapability(value *SharingCapabilities)() {
    m.sharingCapability = value
}
// SetSharingDomainRestrictionMode sets the sharingDomainRestrictionMode property value. Specifies the external sharing mode for domains. Possible values are: none, allowList, blockList.
func (m *Settings) SetSharingDomainRestrictionMode(value *SharingDomainRestrictionMode)() {
    m.sharingDomainRestrictionMode = value
}
// SetSiteCreationDefaultManagedPath sets the siteCreationDefaultManagedPath property value. The value of the team site managed path. This is the path under which new team sites will be created.
func (m *Settings) SetSiteCreationDefaultManagedPath(value *string)() {
    m.siteCreationDefaultManagedPath = value
}
// SetSiteCreationDefaultStorageLimitInMB sets the siteCreationDefaultStorageLimitInMB property value. The default storage quota for a new site upon creation. Measured in megabytes (MB).
func (m *Settings) SetSiteCreationDefaultStorageLimitInMB(value *int32)() {
    m.siteCreationDefaultStorageLimitInMB = value
}
// SetTenantDefaultTimezone sets the tenantDefaultTimezone property value. The default timezone of a tenant for newly created sites.
func (m *Settings) SetTenantDefaultTimezone(value *string)() {
    m.tenantDefaultTimezone = value
}
