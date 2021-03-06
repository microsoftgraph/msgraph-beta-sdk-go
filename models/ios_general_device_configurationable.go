package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosGeneralDeviceConfigurationable 
type IosGeneralDeviceConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountBlockModification()(*bool)
    GetActivationLockAllowWhenSupervised()(*bool)
    GetAirDropBlocked()(*bool)
    GetAirDropForceUnmanagedDropTarget()(*bool)
    GetAirPlayForcePairingPasswordForOutgoingRequests()(*bool)
    GetAirPrintBlockCredentialsStorage()(*bool)
    GetAirPrintBlocked()(*bool)
    GetAirPrintBlockiBeaconDiscovery()(*bool)
    GetAirPrintForceTrustedTLS()(*bool)
    GetAppClipsBlocked()(*bool)
    GetAppleNewsBlocked()(*bool)
    GetApplePersonalizedAdsBlocked()(*bool)
    GetAppleWatchBlockPairing()(*bool)
    GetAppleWatchForceWristDetection()(*bool)
    GetAppRemovalBlocked()(*bool)
    GetAppsSingleAppModeList()([]AppListItemable)
    GetAppStoreBlockAutomaticDownloads()(*bool)
    GetAppStoreBlocked()(*bool)
    GetAppStoreBlockInAppPurchases()(*bool)
    GetAppStoreBlockUIAppInstallation()(*bool)
    GetAppStoreRequirePassword()(*bool)
    GetAppsVisibilityList()([]AppListItemable)
    GetAppsVisibilityListType()(*AppListType)
    GetAutoFillForceAuthentication()(*bool)
    GetAutoUnlockBlocked()(*bool)
    GetBlockSystemAppRemoval()(*bool)
    GetBluetoothBlockModification()(*bool)
    GetCameraBlocked()(*bool)
    GetCellularBlockDataRoaming()(*bool)
    GetCellularBlockGlobalBackgroundFetchWhileRoaming()(*bool)
    GetCellularBlockPerAppDataModification()(*bool)
    GetCellularBlockPersonalHotspot()(*bool)
    GetCellularBlockPersonalHotspotModification()(*bool)
    GetCellularBlockPlanModification()(*bool)
    GetCellularBlockVoiceRoaming()(*bool)
    GetCertificatesBlockUntrustedTlsCertificates()(*bool)
    GetClassroomAppBlockRemoteScreenObservation()(*bool)
    GetClassroomAppForceUnpromptedScreenObservation()(*bool)
    GetClassroomForceAutomaticallyJoinClasses()(*bool)
    GetClassroomForceRequestPermissionToLeaveClasses()(*bool)
    GetClassroomForceUnpromptedAppAndDeviceLock()(*bool)
    GetCompliantAppListType()(*AppListType)
    GetCompliantAppsList()([]AppListItemable)
    GetConfigurationProfileBlockChanges()(*bool)
    GetContactsAllowManagedToUnmanagedWrite()(*bool)
    GetContactsAllowUnmanagedToManagedRead()(*bool)
    GetContinuousPathKeyboardBlocked()(*bool)
    GetDateAndTimeForceSetAutomatically()(*bool)
    GetDefinitionLookupBlocked()(*bool)
    GetDeviceBlockEnableRestrictions()(*bool)
    GetDeviceBlockEraseContentAndSettings()(*bool)
    GetDeviceBlockNameModification()(*bool)
    GetDiagnosticDataBlockSubmission()(*bool)
    GetDiagnosticDataBlockSubmissionModification()(*bool)
    GetDocumentsBlockManagedDocumentsInUnmanagedApps()(*bool)
    GetDocumentsBlockUnmanagedDocumentsInManagedApps()(*bool)
    GetEmailInDomainSuffixes()([]string)
    GetEnterpriseAppBlockTrust()(*bool)
    GetEnterpriseAppBlockTrustModification()(*bool)
    GetEnterpriseBookBlockBackup()(*bool)
    GetEnterpriseBookBlockMetadataSync()(*bool)
    GetEsimBlockModification()(*bool)
    GetFaceTimeBlocked()(*bool)
    GetFilesNetworkDriveAccessBlocked()(*bool)
    GetFilesUsbDriveAccessBlocked()(*bool)
    GetFindMyDeviceInFindMyAppBlocked()(*bool)
    GetFindMyFriendsBlocked()(*bool)
    GetFindMyFriendsInFindMyAppBlocked()(*bool)
    GetGameCenterBlocked()(*bool)
    GetGamingBlockGameCenterFriends()(*bool)
    GetGamingBlockMultiplayer()(*bool)
    GetHostPairingBlocked()(*bool)
    GetIBooksStoreBlocked()(*bool)
    GetIBooksStoreBlockErotica()(*bool)
    GetICloudBlockActivityContinuation()(*bool)
    GetICloudBlockBackup()(*bool)
    GetICloudBlockDocumentSync()(*bool)
    GetICloudBlockManagedAppsSync()(*bool)
    GetICloudBlockPhotoLibrary()(*bool)
    GetICloudBlockPhotoStreamSync()(*bool)
    GetICloudBlockSharedPhotoStream()(*bool)
    GetICloudPrivateRelayBlocked()(*bool)
    GetICloudRequireEncryptedBackup()(*bool)
    GetITunesBlocked()(*bool)
    GetITunesBlockExplicitContent()(*bool)
    GetITunesBlockMusicService()(*bool)
    GetITunesBlockRadio()(*bool)
    GetKeyboardBlockAutoCorrect()(*bool)
    GetKeyboardBlockDictation()(*bool)
    GetKeyboardBlockPredictive()(*bool)
    GetKeyboardBlockShortcuts()(*bool)
    GetKeyboardBlockSpellCheck()(*bool)
    GetKeychainBlockCloudSync()(*bool)
    GetKioskModeAllowAssistiveSpeak()(*bool)
    GetKioskModeAllowAssistiveTouchSettings()(*bool)
    GetKioskModeAllowAutoLock()(*bool)
    GetKioskModeAllowColorInversionSettings()(*bool)
    GetKioskModeAllowRingerSwitch()(*bool)
    GetKioskModeAllowScreenRotation()(*bool)
    GetKioskModeAllowSleepButton()(*bool)
    GetKioskModeAllowTouchscreen()(*bool)
    GetKioskModeAllowVoiceControlModification()(*bool)
    GetKioskModeAllowVoiceOverSettings()(*bool)
    GetKioskModeAllowVolumeButtons()(*bool)
    GetKioskModeAllowZoomSettings()(*bool)
    GetKioskModeAppStoreUrl()(*string)
    GetKioskModeAppType()(*IosKioskModeAppType)
    GetKioskModeBlockAutoLock()(*bool)
    GetKioskModeBlockRingerSwitch()(*bool)
    GetKioskModeBlockScreenRotation()(*bool)
    GetKioskModeBlockSleepButton()(*bool)
    GetKioskModeBlockTouchscreen()(*bool)
    GetKioskModeBlockVolumeButtons()(*bool)
    GetKioskModeBuiltInAppId()(*string)
    GetKioskModeEnableVoiceControl()(*bool)
    GetKioskModeManagedAppId()(*string)
    GetKioskModeRequireAssistiveTouch()(*bool)
    GetKioskModeRequireColorInversion()(*bool)
    GetKioskModeRequireMonoAudio()(*bool)
    GetKioskModeRequireVoiceOver()(*bool)
    GetKioskModeRequireZoom()(*bool)
    GetLockScreenBlockControlCenter()(*bool)
    GetLockScreenBlockNotificationView()(*bool)
    GetLockScreenBlockPassbook()(*bool)
    GetLockScreenBlockTodayView()(*bool)
    GetManagedPasteboardRequired()(*bool)
    GetMediaContentRatingApps()(*RatingAppsType)
    GetMediaContentRatingAustralia()(MediaContentRatingAustraliaable)
    GetMediaContentRatingCanada()(MediaContentRatingCanadaable)
    GetMediaContentRatingFrance()(MediaContentRatingFranceable)
    GetMediaContentRatingGermany()(MediaContentRatingGermanyable)
    GetMediaContentRatingIreland()(MediaContentRatingIrelandable)
    GetMediaContentRatingJapan()(MediaContentRatingJapanable)
    GetMediaContentRatingNewZealand()(MediaContentRatingNewZealandable)
    GetMediaContentRatingUnitedKingdom()(MediaContentRatingUnitedKingdomable)
    GetMediaContentRatingUnitedStates()(MediaContentRatingUnitedStatesable)
    GetMessagesBlocked()(*bool)
    GetNetworkUsageRules()([]IosNetworkUsageRuleable)
    GetNfcBlocked()(*bool)
    GetNotificationsBlockSettingsModification()(*bool)
    GetOnDeviceOnlyDictationForced()(*bool)
    GetOnDeviceOnlyTranslationForced()(*bool)
    GetPasscodeBlockFingerprintModification()(*bool)
    GetPasscodeBlockFingerprintUnlock()(*bool)
    GetPasscodeBlockModification()(*bool)
    GetPasscodeBlockSimple()(*bool)
    GetPasscodeExpirationDays()(*int32)
    GetPasscodeMinimumCharacterSetCount()(*int32)
    GetPasscodeMinimumLength()(*int32)
    GetPasscodeMinutesOfInactivityBeforeLock()(*int32)
    GetPasscodeMinutesOfInactivityBeforeScreenTimeout()(*int32)
    GetPasscodePreviousPasscodeBlockCount()(*int32)
    GetPasscodeRequired()(*bool)
    GetPasscodeRequiredType()(*RequiredPasswordType)
    GetPasscodeSignInFailureCountBeforeWipe()(*int32)
    GetPasswordBlockAirDropSharing()(*bool)
    GetPasswordBlockAutoFill()(*bool)
    GetPasswordBlockProximityRequests()(*bool)
    GetPkiBlockOTAUpdates()(*bool)
    GetPodcastsBlocked()(*bool)
    GetPrivacyForceLimitAdTracking()(*bool)
    GetProximityBlockSetupToNewDevice()(*bool)
    GetSafariBlockAutofill()(*bool)
    GetSafariBlocked()(*bool)
    GetSafariBlockJavaScript()(*bool)
    GetSafariBlockPopups()(*bool)
    GetSafariCookieSettings()(*WebBrowserCookieSettings)
    GetSafariManagedDomains()([]string)
    GetSafariPasswordAutoFillDomains()([]string)
    GetSafariRequireFraudWarning()(*bool)
    GetScreenCaptureBlocked()(*bool)
    GetSharedDeviceBlockTemporarySessions()(*bool)
    GetSiriBlocked()(*bool)
    GetSiriBlockedWhenLocked()(*bool)
    GetSiriBlockUserGeneratedContent()(*bool)
    GetSiriRequireProfanityFilter()(*bool)
    GetSoftwareUpdatesEnforcedDelayInDays()(*int32)
    GetSoftwareUpdatesForceDelayed()(*bool)
    GetSpotlightBlockInternetResults()(*bool)
    GetUnpairedExternalBootToRecoveryAllowed()(*bool)
    GetUsbRestrictedModeBlocked()(*bool)
    GetVoiceDialingBlocked()(*bool)
    GetVpnBlockCreation()(*bool)
    GetWallpaperBlockModification()(*bool)
    GetWiFiConnectOnlyToConfiguredNetworks()(*bool)
    GetWiFiConnectToAllowedNetworksOnlyForced()(*bool)
    GetWifiPowerOnForced()(*bool)
    SetAccountBlockModification(value *bool)()
    SetActivationLockAllowWhenSupervised(value *bool)()
    SetAirDropBlocked(value *bool)()
    SetAirDropForceUnmanagedDropTarget(value *bool)()
    SetAirPlayForcePairingPasswordForOutgoingRequests(value *bool)()
    SetAirPrintBlockCredentialsStorage(value *bool)()
    SetAirPrintBlocked(value *bool)()
    SetAirPrintBlockiBeaconDiscovery(value *bool)()
    SetAirPrintForceTrustedTLS(value *bool)()
    SetAppClipsBlocked(value *bool)()
    SetAppleNewsBlocked(value *bool)()
    SetApplePersonalizedAdsBlocked(value *bool)()
    SetAppleWatchBlockPairing(value *bool)()
    SetAppleWatchForceWristDetection(value *bool)()
    SetAppRemovalBlocked(value *bool)()
    SetAppsSingleAppModeList(value []AppListItemable)()
    SetAppStoreBlockAutomaticDownloads(value *bool)()
    SetAppStoreBlocked(value *bool)()
    SetAppStoreBlockInAppPurchases(value *bool)()
    SetAppStoreBlockUIAppInstallation(value *bool)()
    SetAppStoreRequirePassword(value *bool)()
    SetAppsVisibilityList(value []AppListItemable)()
    SetAppsVisibilityListType(value *AppListType)()
    SetAutoFillForceAuthentication(value *bool)()
    SetAutoUnlockBlocked(value *bool)()
    SetBlockSystemAppRemoval(value *bool)()
    SetBluetoothBlockModification(value *bool)()
    SetCameraBlocked(value *bool)()
    SetCellularBlockDataRoaming(value *bool)()
    SetCellularBlockGlobalBackgroundFetchWhileRoaming(value *bool)()
    SetCellularBlockPerAppDataModification(value *bool)()
    SetCellularBlockPersonalHotspot(value *bool)()
    SetCellularBlockPersonalHotspotModification(value *bool)()
    SetCellularBlockPlanModification(value *bool)()
    SetCellularBlockVoiceRoaming(value *bool)()
    SetCertificatesBlockUntrustedTlsCertificates(value *bool)()
    SetClassroomAppBlockRemoteScreenObservation(value *bool)()
    SetClassroomAppForceUnpromptedScreenObservation(value *bool)()
    SetClassroomForceAutomaticallyJoinClasses(value *bool)()
    SetClassroomForceRequestPermissionToLeaveClasses(value *bool)()
    SetClassroomForceUnpromptedAppAndDeviceLock(value *bool)()
    SetCompliantAppListType(value *AppListType)()
    SetCompliantAppsList(value []AppListItemable)()
    SetConfigurationProfileBlockChanges(value *bool)()
    SetContactsAllowManagedToUnmanagedWrite(value *bool)()
    SetContactsAllowUnmanagedToManagedRead(value *bool)()
    SetContinuousPathKeyboardBlocked(value *bool)()
    SetDateAndTimeForceSetAutomatically(value *bool)()
    SetDefinitionLookupBlocked(value *bool)()
    SetDeviceBlockEnableRestrictions(value *bool)()
    SetDeviceBlockEraseContentAndSettings(value *bool)()
    SetDeviceBlockNameModification(value *bool)()
    SetDiagnosticDataBlockSubmission(value *bool)()
    SetDiagnosticDataBlockSubmissionModification(value *bool)()
    SetDocumentsBlockManagedDocumentsInUnmanagedApps(value *bool)()
    SetDocumentsBlockUnmanagedDocumentsInManagedApps(value *bool)()
    SetEmailInDomainSuffixes(value []string)()
    SetEnterpriseAppBlockTrust(value *bool)()
    SetEnterpriseAppBlockTrustModification(value *bool)()
    SetEnterpriseBookBlockBackup(value *bool)()
    SetEnterpriseBookBlockMetadataSync(value *bool)()
    SetEsimBlockModification(value *bool)()
    SetFaceTimeBlocked(value *bool)()
    SetFilesNetworkDriveAccessBlocked(value *bool)()
    SetFilesUsbDriveAccessBlocked(value *bool)()
    SetFindMyDeviceInFindMyAppBlocked(value *bool)()
    SetFindMyFriendsBlocked(value *bool)()
    SetFindMyFriendsInFindMyAppBlocked(value *bool)()
    SetGameCenterBlocked(value *bool)()
    SetGamingBlockGameCenterFriends(value *bool)()
    SetGamingBlockMultiplayer(value *bool)()
    SetHostPairingBlocked(value *bool)()
    SetIBooksStoreBlocked(value *bool)()
    SetIBooksStoreBlockErotica(value *bool)()
    SetICloudBlockActivityContinuation(value *bool)()
    SetICloudBlockBackup(value *bool)()
    SetICloudBlockDocumentSync(value *bool)()
    SetICloudBlockManagedAppsSync(value *bool)()
    SetICloudBlockPhotoLibrary(value *bool)()
    SetICloudBlockPhotoStreamSync(value *bool)()
    SetICloudBlockSharedPhotoStream(value *bool)()
    SetICloudPrivateRelayBlocked(value *bool)()
    SetICloudRequireEncryptedBackup(value *bool)()
    SetITunesBlocked(value *bool)()
    SetITunesBlockExplicitContent(value *bool)()
    SetITunesBlockMusicService(value *bool)()
    SetITunesBlockRadio(value *bool)()
    SetKeyboardBlockAutoCorrect(value *bool)()
    SetKeyboardBlockDictation(value *bool)()
    SetKeyboardBlockPredictive(value *bool)()
    SetKeyboardBlockShortcuts(value *bool)()
    SetKeyboardBlockSpellCheck(value *bool)()
    SetKeychainBlockCloudSync(value *bool)()
    SetKioskModeAllowAssistiveSpeak(value *bool)()
    SetKioskModeAllowAssistiveTouchSettings(value *bool)()
    SetKioskModeAllowAutoLock(value *bool)()
    SetKioskModeAllowColorInversionSettings(value *bool)()
    SetKioskModeAllowRingerSwitch(value *bool)()
    SetKioskModeAllowScreenRotation(value *bool)()
    SetKioskModeAllowSleepButton(value *bool)()
    SetKioskModeAllowTouchscreen(value *bool)()
    SetKioskModeAllowVoiceControlModification(value *bool)()
    SetKioskModeAllowVoiceOverSettings(value *bool)()
    SetKioskModeAllowVolumeButtons(value *bool)()
    SetKioskModeAllowZoomSettings(value *bool)()
    SetKioskModeAppStoreUrl(value *string)()
    SetKioskModeAppType(value *IosKioskModeAppType)()
    SetKioskModeBlockAutoLock(value *bool)()
    SetKioskModeBlockRingerSwitch(value *bool)()
    SetKioskModeBlockScreenRotation(value *bool)()
    SetKioskModeBlockSleepButton(value *bool)()
    SetKioskModeBlockTouchscreen(value *bool)()
    SetKioskModeBlockVolumeButtons(value *bool)()
    SetKioskModeBuiltInAppId(value *string)()
    SetKioskModeEnableVoiceControl(value *bool)()
    SetKioskModeManagedAppId(value *string)()
    SetKioskModeRequireAssistiveTouch(value *bool)()
    SetKioskModeRequireColorInversion(value *bool)()
    SetKioskModeRequireMonoAudio(value *bool)()
    SetKioskModeRequireVoiceOver(value *bool)()
    SetKioskModeRequireZoom(value *bool)()
    SetLockScreenBlockControlCenter(value *bool)()
    SetLockScreenBlockNotificationView(value *bool)()
    SetLockScreenBlockPassbook(value *bool)()
    SetLockScreenBlockTodayView(value *bool)()
    SetManagedPasteboardRequired(value *bool)()
    SetMediaContentRatingApps(value *RatingAppsType)()
    SetMediaContentRatingAustralia(value MediaContentRatingAustraliaable)()
    SetMediaContentRatingCanada(value MediaContentRatingCanadaable)()
    SetMediaContentRatingFrance(value MediaContentRatingFranceable)()
    SetMediaContentRatingGermany(value MediaContentRatingGermanyable)()
    SetMediaContentRatingIreland(value MediaContentRatingIrelandable)()
    SetMediaContentRatingJapan(value MediaContentRatingJapanable)()
    SetMediaContentRatingNewZealand(value MediaContentRatingNewZealandable)()
    SetMediaContentRatingUnitedKingdom(value MediaContentRatingUnitedKingdomable)()
    SetMediaContentRatingUnitedStates(value MediaContentRatingUnitedStatesable)()
    SetMessagesBlocked(value *bool)()
    SetNetworkUsageRules(value []IosNetworkUsageRuleable)()
    SetNfcBlocked(value *bool)()
    SetNotificationsBlockSettingsModification(value *bool)()
    SetOnDeviceOnlyDictationForced(value *bool)()
    SetOnDeviceOnlyTranslationForced(value *bool)()
    SetPasscodeBlockFingerprintModification(value *bool)()
    SetPasscodeBlockFingerprintUnlock(value *bool)()
    SetPasscodeBlockModification(value *bool)()
    SetPasscodeBlockSimple(value *bool)()
    SetPasscodeExpirationDays(value *int32)()
    SetPasscodeMinimumCharacterSetCount(value *int32)()
    SetPasscodeMinimumLength(value *int32)()
    SetPasscodeMinutesOfInactivityBeforeLock(value *int32)()
    SetPasscodeMinutesOfInactivityBeforeScreenTimeout(value *int32)()
    SetPasscodePreviousPasscodeBlockCount(value *int32)()
    SetPasscodeRequired(value *bool)()
    SetPasscodeRequiredType(value *RequiredPasswordType)()
    SetPasscodeSignInFailureCountBeforeWipe(value *int32)()
    SetPasswordBlockAirDropSharing(value *bool)()
    SetPasswordBlockAutoFill(value *bool)()
    SetPasswordBlockProximityRequests(value *bool)()
    SetPkiBlockOTAUpdates(value *bool)()
    SetPodcastsBlocked(value *bool)()
    SetPrivacyForceLimitAdTracking(value *bool)()
    SetProximityBlockSetupToNewDevice(value *bool)()
    SetSafariBlockAutofill(value *bool)()
    SetSafariBlocked(value *bool)()
    SetSafariBlockJavaScript(value *bool)()
    SetSafariBlockPopups(value *bool)()
    SetSafariCookieSettings(value *WebBrowserCookieSettings)()
    SetSafariManagedDomains(value []string)()
    SetSafariPasswordAutoFillDomains(value []string)()
    SetSafariRequireFraudWarning(value *bool)()
    SetScreenCaptureBlocked(value *bool)()
    SetSharedDeviceBlockTemporarySessions(value *bool)()
    SetSiriBlocked(value *bool)()
    SetSiriBlockedWhenLocked(value *bool)()
    SetSiriBlockUserGeneratedContent(value *bool)()
    SetSiriRequireProfanityFilter(value *bool)()
    SetSoftwareUpdatesEnforcedDelayInDays(value *int32)()
    SetSoftwareUpdatesForceDelayed(value *bool)()
    SetSpotlightBlockInternetResults(value *bool)()
    SetUnpairedExternalBootToRecoveryAllowed(value *bool)()
    SetUsbRestrictedModeBlocked(value *bool)()
    SetVoiceDialingBlocked(value *bool)()
    SetVpnBlockCreation(value *bool)()
    SetWallpaperBlockModification(value *bool)()
    SetWiFiConnectOnlyToConfiguredNetworks(value *bool)()
    SetWiFiConnectToAllowedNetworksOnlyForced(value *bool)()
    SetWifiPowerOnForced(value *bool)()
}
