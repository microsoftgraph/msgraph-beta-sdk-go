package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSGeneralDeviceConfiguration 
type MacOSGeneralDeviceConfiguration struct {
    DeviceConfiguration
    // Yes prevents users from adding friends to Game Center. Available for devices running macOS versions 10.13 and later.
    addingGameCenterFriendsBlocked *bool
    // Indicates whether or not to allow AirDrop.
    airDropBlocked *bool
    // Indicates whether or to block users from unlocking their Mac with Apple Watch.
    appleWatchBlockAutoUnlock *bool
    // Indicates whether or not to block the user from accessing the camera of the device.
    cameraBlocked *bool
    // Indicates whether or not to allow remote screen observation by Classroom app. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
    classroomAppBlockRemoteScreenObservation *bool
    // Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to view a student's screen without prompting. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
    classroomAppForceUnpromptedScreenObservation *bool
    // Indicates whether or not to automatically give permission to the teacher's requests, without prompting the student. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
    classroomForceAutomaticallyJoinClasses *bool
    // Indicates whether a student enrolled in an unmanaged course via Classroom will be required to request permission from the teacher when attempting to leave the course. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
    classroomForceRequestPermissionToLeaveClasses *bool
    // Indicates whether or not to allow the teacher to lock apps or the device without prompting the student. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
    classroomForceUnpromptedAppAndDeviceLock *bool
    // Possible values of the compliance app list.
    compliantAppListType *AppListType
    // List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
    compliantAppsList []AppListItemable
    // Indicates whether or not to allow content caching.
    contentCachingBlocked *bool
    // Indicates whether or not to block definition lookup.
    definitionLookupBlocked *bool
    // An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
    emailInDomainSuffixes []string
    // TRUE disables the reset option on supervised devices. FALSE enables the reset option on supervised devices. Available for devices running macOS versions 12.0 and later.
    eraseContentAndSettingsBlocked *bool
    // Yes disables Game Center, and the Game Center icon is removed from the Home screen. Available for devices running macOS versions 10.13 and later.
    gameCenterBlocked *bool
    // Indicates whether or not to block the user from continuing work that they started on a MacOS device on another iOS or MacOS device (MacOS 10.15 or later).
    iCloudBlockActivityContinuation *bool
    // Indicates whether or not to block iCloud from syncing contacts.
    iCloudBlockAddressBook *bool
    // Indicates whether or not to block iCloud from syncing bookmarks.
    iCloudBlockBookmarks *bool
    // Indicates whether or not to block iCloud from syncing calendars.
    iCloudBlockCalendar *bool
    // Indicates whether or not to block iCloud document sync.
    iCloudBlockDocumentSync *bool
    // Indicates whether or not to block iCloud from syncing mail.
    iCloudBlockMail *bool
    // Indicates whether or not to block iCloud from syncing notes.
    iCloudBlockNotes *bool
    // Indicates whether or not to block iCloud Photo Library.
    iCloudBlockPhotoLibrary *bool
    // Indicates whether or not to block iCloud from syncing reminders.
    iCloudBlockReminders *bool
    // When TRUE the synchronization of cloud desktop and documents is blocked. When FALSE, synchronization of the cloud desktop and documents are allowed. Available for devices running macOS 10.12.4 and later.
    iCloudDesktopAndDocumentsBlocked *bool
    // iCloud private relay is an iCloud+ service that prevents networks and servers from monitoring a person's activity across the internet. By blocking iCloud private relay, Apple will not encrypt the traffic leaving the device. Available for devices running macOS 12 and later.
    iCloudPrivateRelayBlocked *bool
    // Indicates whether or not to block files from being transferred using iTunes.
    iTunesBlockFileSharing *bool
    // Indicates whether or not to block Music service and revert Music app to classic mode.
    iTunesBlockMusicService *bool
    // Indicates whether or not to block the user from using dictation input.
    keyboardBlockDictation *bool
    // Indicates whether or not iCloud keychain synchronization is blocked (macOS 10.12 and later).
    keychainBlockCloudSync *bool
    // TRUE prevents multiplayer gaming when using Game Center. FALSE allows multiplayer gaming when using Game Center. Available for devices running macOS versions 10.13 and later.
    multiplayerGamingBlocked *bool
    // Indicates whether or not to block sharing passwords with the AirDrop passwords feature.
    passwordBlockAirDropSharing *bool
    // Indicates whether or not to block the AutoFill Passwords feature.
    passwordBlockAutoFill *bool
    // Indicates whether or not to block fingerprint unlock.
    passwordBlockFingerprintUnlock *bool
    // Indicates whether or not to allow passcode modification.
    passwordBlockModification *bool
    // Indicates whether or not to block requesting passwords from nearby devices.
    passwordBlockProximityRequests *bool
    // Block simple passwords.
    passwordBlockSimple *bool
    // Number of days before the password expires.
    passwordExpirationDays *int32
    // The number of allowed failed attempts to enter the passcode at the device's lock screen. Valid values 2 to 11
    passwordMaximumAttemptCount *int32
    // Number of character sets a password must contain. Valid values 0 to 4
    passwordMinimumCharacterSetCount *int32
    // Minimum length of passwords.
    passwordMinimumLength *int32
    // Minutes of inactivity required before a password is required.
    passwordMinutesOfInactivityBeforeLock *int32
    // Minutes of inactivity required before the screen times out.
    passwordMinutesOfInactivityBeforeScreenTimeout *int32
    // The number of minutes before the login is reset after the maximum number of unsuccessful login attempts is reached.
    passwordMinutesUntilFailedLoginReset *int32
    // Number of previous passwords to block.
    passwordPreviousPasswordBlockCount *int32
    // Whether or not to require a password.
    passwordRequired *bool
    // Possible values of required passwords.
    passwordRequiredType *RequiredPasswordType
    // List of privacy preference policy controls. This collection can contain a maximum of 10000 elements.
    privacyAccessControls []MacOSPrivacyAccessControlItemable
    // Indicates whether or not to block the user from using Auto fill in Safari.
    safariBlockAutofill *bool
    // Indicates whether or not to block the user from taking Screenshots.
    screenCaptureBlocked *bool
    // Specify the number of days (1-90) to delay visibility of major OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
    softwareUpdateMajorOSDeferredInstallDelayInDays *int32
    // Specify the number of days (1-90) to delay visibility of minor OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
    softwareUpdateMinorOSDeferredInstallDelayInDays *int32
    // Specify the number of days (1-90) to delay visibility of non-OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
    softwareUpdateNonOSDeferredInstallDelayInDays *int32
    // Sets how many days a software update will be delyed for a supervised device. Valid values 0 to 90
    softwareUpdatesEnforcedDelayInDays *int32
    // Indicates whether or not to block Spotlight from returning any results from an Internet search.
    spotlightBlockInternetResults *bool
    // Maximum hours after which the user must enter their password to unlock the device instead of using Touch ID. Available for devices running macOS 12 and later. Valid values 0 to 2147483647
    touchIdTimeoutInHours *int32
    // Determines whether to delay OS and/or app updates for macOS. Possible values are: none, delayOSUpdateVisibility, delayAppUpdateVisibility, unknownFutureValue, delayMajorOsUpdateVisibility.
    updateDelayPolicy *MacOSSoftwareUpdateDelayPolicy
    // TRUE prevents the wallpaper from being changed. FALSE allows the wallpaper to be changed. Available for devices running macOS versions 10.13 and later.
    wallpaperModificationBlocked *bool
}
// NewMacOSGeneralDeviceConfiguration instantiates a new MacOSGeneralDeviceConfiguration and sets the default values.
func NewMacOSGeneralDeviceConfiguration()(*MacOSGeneralDeviceConfiguration) {
    m := &MacOSGeneralDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    return m
}
// CreateMacOSGeneralDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSGeneralDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSGeneralDeviceConfiguration(), nil
}
// GetAddingGameCenterFriendsBlocked gets the addingGameCenterFriendsBlocked property value. Yes prevents users from adding friends to Game Center. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) GetAddingGameCenterFriendsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.addingGameCenterFriendsBlocked
    }
}
// GetAirDropBlocked gets the airDropBlocked property value. Indicates whether or not to allow AirDrop.
func (m *MacOSGeneralDeviceConfiguration) GetAirDropBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.airDropBlocked
    }
}
// GetAppleWatchBlockAutoUnlock gets the appleWatchBlockAutoUnlock property value. Indicates whether or to block users from unlocking their Mac with Apple Watch.
func (m *MacOSGeneralDeviceConfiguration) GetAppleWatchBlockAutoUnlock()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.appleWatchBlockAutoUnlock
    }
}
// GetCameraBlocked gets the cameraBlocked property value. Indicates whether or not to block the user from accessing the camera of the device.
func (m *MacOSGeneralDeviceConfiguration) GetCameraBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.cameraBlocked
    }
}
// GetClassroomAppBlockRemoteScreenObservation gets the classroomAppBlockRemoteScreenObservation property value. Indicates whether or not to allow remote screen observation by Classroom app. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) GetClassroomAppBlockRemoteScreenObservation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.classroomAppBlockRemoteScreenObservation
    }
}
// GetClassroomAppForceUnpromptedScreenObservation gets the classroomAppForceUnpromptedScreenObservation property value. Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to view a student's screen without prompting. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) GetClassroomAppForceUnpromptedScreenObservation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.classroomAppForceUnpromptedScreenObservation
    }
}
// GetClassroomForceAutomaticallyJoinClasses gets the classroomForceAutomaticallyJoinClasses property value. Indicates whether or not to automatically give permission to the teacher's requests, without prompting the student. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) GetClassroomForceAutomaticallyJoinClasses()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.classroomForceAutomaticallyJoinClasses
    }
}
// GetClassroomForceRequestPermissionToLeaveClasses gets the classroomForceRequestPermissionToLeaveClasses property value. Indicates whether a student enrolled in an unmanaged course via Classroom will be required to request permission from the teacher when attempting to leave the course. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) GetClassroomForceRequestPermissionToLeaveClasses()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.classroomForceRequestPermissionToLeaveClasses
    }
}
// GetClassroomForceUnpromptedAppAndDeviceLock gets the classroomForceUnpromptedAppAndDeviceLock property value. Indicates whether or not to allow the teacher to lock apps or the device without prompting the student. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) GetClassroomForceUnpromptedAppAndDeviceLock()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.classroomForceUnpromptedAppAndDeviceLock
    }
}
// GetCompliantAppListType gets the compliantAppListType property value. Possible values of the compliance app list.
func (m *MacOSGeneralDeviceConfiguration) GetCompliantAppListType()(*AppListType) {
    if m == nil {
        return nil
    } else {
        return m.compliantAppListType
    }
}
// GetCompliantAppsList gets the compliantAppsList property value. List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
func (m *MacOSGeneralDeviceConfiguration) GetCompliantAppsList()([]AppListItemable) {
    if m == nil {
        return nil
    } else {
        return m.compliantAppsList
    }
}
// GetContentCachingBlocked gets the contentCachingBlocked property value. Indicates whether or not to allow content caching.
func (m *MacOSGeneralDeviceConfiguration) GetContentCachingBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contentCachingBlocked
    }
}
// GetDefinitionLookupBlocked gets the definitionLookupBlocked property value. Indicates whether or not to block definition lookup.
func (m *MacOSGeneralDeviceConfiguration) GetDefinitionLookupBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.definitionLookupBlocked
    }
}
// GetEmailInDomainSuffixes gets the emailInDomainSuffixes property value. An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
func (m *MacOSGeneralDeviceConfiguration) GetEmailInDomainSuffixes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.emailInDomainSuffixes
    }
}
// GetEraseContentAndSettingsBlocked gets the eraseContentAndSettingsBlocked property value. TRUE disables the reset option on supervised devices. FALSE enables the reset option on supervised devices. Available for devices running macOS versions 12.0 and later.
func (m *MacOSGeneralDeviceConfiguration) GetEraseContentAndSettingsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.eraseContentAndSettingsBlocked
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSGeneralDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["addingGameCenterFriendsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddingGameCenterFriendsBlocked(val)
        }
        return nil
    }
    res["airDropBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAirDropBlocked(val)
        }
        return nil
    }
    res["appleWatchBlockAutoUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleWatchBlockAutoUnlock(val)
        }
        return nil
    }
    res["cameraBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameraBlocked(val)
        }
        return nil
    }
    res["classroomAppBlockRemoteScreenObservation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroomAppBlockRemoteScreenObservation(val)
        }
        return nil
    }
    res["classroomAppForceUnpromptedScreenObservation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroomAppForceUnpromptedScreenObservation(val)
        }
        return nil
    }
    res["classroomForceAutomaticallyJoinClasses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroomForceAutomaticallyJoinClasses(val)
        }
        return nil
    }
    res["classroomForceRequestPermissionToLeaveClasses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroomForceRequestPermissionToLeaveClasses(val)
        }
        return nil
    }
    res["classroomForceUnpromptedAppAndDeviceLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroomForceUnpromptedAppAndDeviceLock(val)
        }
        return nil
    }
    res["compliantAppListType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppListType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantAppListType(val.(*AppListType))
        }
        return nil
    }
    res["compliantAppsList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetCompliantAppsList(res)
        }
        return nil
    }
    res["contentCachingBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCachingBlocked(val)
        }
        return nil
    }
    res["definitionLookupBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionLookupBlocked(val)
        }
        return nil
    }
    res["emailInDomainSuffixes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEmailInDomainSuffixes(res)
        }
        return nil
    }
    res["eraseContentAndSettingsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEraseContentAndSettingsBlocked(val)
        }
        return nil
    }
    res["gameCenterBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGameCenterBlocked(val)
        }
        return nil
    }
    res["iCloudBlockActivityContinuation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockActivityContinuation(val)
        }
        return nil
    }
    res["iCloudBlockAddressBook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockAddressBook(val)
        }
        return nil
    }
    res["iCloudBlockBookmarks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockBookmarks(val)
        }
        return nil
    }
    res["iCloudBlockCalendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockCalendar(val)
        }
        return nil
    }
    res["iCloudBlockDocumentSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockDocumentSync(val)
        }
        return nil
    }
    res["iCloudBlockMail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockMail(val)
        }
        return nil
    }
    res["iCloudBlockNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockNotes(val)
        }
        return nil
    }
    res["iCloudBlockPhotoLibrary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockPhotoLibrary(val)
        }
        return nil
    }
    res["iCloudBlockReminders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockReminders(val)
        }
        return nil
    }
    res["iCloudDesktopAndDocumentsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudDesktopAndDocumentsBlocked(val)
        }
        return nil
    }
    res["iCloudPrivateRelayBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudPrivateRelayBlocked(val)
        }
        return nil
    }
    res["iTunesBlockFileSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetITunesBlockFileSharing(val)
        }
        return nil
    }
    res["iTunesBlockMusicService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetITunesBlockMusicService(val)
        }
        return nil
    }
    res["keyboardBlockDictation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardBlockDictation(val)
        }
        return nil
    }
    res["keychainBlockCloudSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeychainBlockCloudSync(val)
        }
        return nil
    }
    res["multiplayerGamingBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultiplayerGamingBlocked(val)
        }
        return nil
    }
    res["passwordBlockAirDropSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockAirDropSharing(val)
        }
        return nil
    }
    res["passwordBlockAutoFill"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockAutoFill(val)
        }
        return nil
    }
    res["passwordBlockFingerprintUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockFingerprintUnlock(val)
        }
        return nil
    }
    res["passwordBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockModification(val)
        }
        return nil
    }
    res["passwordBlockProximityRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockProximityRequests(val)
        }
        return nil
    }
    res["passwordBlockSimple"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockSimple(val)
        }
        return nil
    }
    res["passwordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordExpirationDays(val)
        }
        return nil
    }
    res["passwordMaximumAttemptCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMaximumAttemptCount(val)
        }
        return nil
    }
    res["passwordMinimumCharacterSetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumCharacterSetCount(val)
        }
        return nil
    }
    res["passwordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLength(val)
        }
        return nil
    }
    res["passwordMinutesOfInactivityBeforeLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeLock(val)
        }
        return nil
    }
    res["passwordMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeScreenTimeout(val)
        }
        return nil
    }
    res["passwordMinutesUntilFailedLoginReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesUntilFailedLoginReset(val)
        }
        return nil
    }
    res["passwordPreviousPasswordBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPreviousPasswordBlockCount(val)
        }
        return nil
    }
    res["passwordRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequired(val)
        }
        return nil
    }
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*RequiredPasswordType))
        }
        return nil
    }
    res["privacyAccessControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSPrivacyAccessControlItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSPrivacyAccessControlItemable, len(val))
            for i, v := range val {
                res[i] = v.(MacOSPrivacyAccessControlItemable)
            }
            m.SetPrivacyAccessControls(res)
        }
        return nil
    }
    res["safariBlockAutofill"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariBlockAutofill(val)
        }
        return nil
    }
    res["screenCaptureBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenCaptureBlocked(val)
        }
        return nil
    }
    res["softwareUpdateMajorOSDeferredInstallDelayInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdateMajorOSDeferredInstallDelayInDays(val)
        }
        return nil
    }
    res["softwareUpdateMinorOSDeferredInstallDelayInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdateMinorOSDeferredInstallDelayInDays(val)
        }
        return nil
    }
    res["softwareUpdateNonOSDeferredInstallDelayInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdateNonOSDeferredInstallDelayInDays(val)
        }
        return nil
    }
    res["softwareUpdatesEnforcedDelayInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdatesEnforcedDelayInDays(val)
        }
        return nil
    }
    res["spotlightBlockInternetResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpotlightBlockInternetResults(val)
        }
        return nil
    }
    res["touchIdTimeoutInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTouchIdTimeoutInHours(val)
        }
        return nil
    }
    res["updateDelayPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateDelayPolicy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateDelayPolicy(val.(*MacOSSoftwareUpdateDelayPolicy))
        }
        return nil
    }
    res["wallpaperModificationBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWallpaperModificationBlocked(val)
        }
        return nil
    }
    return res
}
// GetGameCenterBlocked gets the gameCenterBlocked property value. Yes disables Game Center, and the Game Center icon is removed from the Home screen. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) GetGameCenterBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.gameCenterBlocked
    }
}
// GetICloudBlockActivityContinuation gets the iCloudBlockActivityContinuation property value. Indicates whether or not to block the user from continuing work that they started on a MacOS device on another iOS or MacOS device (MacOS 10.15 or later).
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockActivityContinuation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockActivityContinuation
    }
}
// GetICloudBlockAddressBook gets the iCloudBlockAddressBook property value. Indicates whether or not to block iCloud from syncing contacts.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockAddressBook()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockAddressBook
    }
}
// GetICloudBlockBookmarks gets the iCloudBlockBookmarks property value. Indicates whether or not to block iCloud from syncing bookmarks.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockBookmarks()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockBookmarks
    }
}
// GetICloudBlockCalendar gets the iCloudBlockCalendar property value. Indicates whether or not to block iCloud from syncing calendars.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockCalendar()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockCalendar
    }
}
// GetICloudBlockDocumentSync gets the iCloudBlockDocumentSync property value. Indicates whether or not to block iCloud document sync.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockDocumentSync()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockDocumentSync
    }
}
// GetICloudBlockMail gets the iCloudBlockMail property value. Indicates whether or not to block iCloud from syncing mail.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockMail()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockMail
    }
}
// GetICloudBlockNotes gets the iCloudBlockNotes property value. Indicates whether or not to block iCloud from syncing notes.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockNotes()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockNotes
    }
}
// GetICloudBlockPhotoLibrary gets the iCloudBlockPhotoLibrary property value. Indicates whether or not to block iCloud Photo Library.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockPhotoLibrary()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockPhotoLibrary
    }
}
// GetICloudBlockReminders gets the iCloudBlockReminders property value. Indicates whether or not to block iCloud from syncing reminders.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockReminders()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudBlockReminders
    }
}
// GetICloudDesktopAndDocumentsBlocked gets the iCloudDesktopAndDocumentsBlocked property value. When TRUE the synchronization of cloud desktop and documents is blocked. When FALSE, synchronization of the cloud desktop and documents are allowed. Available for devices running macOS 10.12.4 and later.
func (m *MacOSGeneralDeviceConfiguration) GetICloudDesktopAndDocumentsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudDesktopAndDocumentsBlocked
    }
}
// GetICloudPrivateRelayBlocked gets the iCloudPrivateRelayBlocked property value. iCloud private relay is an iCloud+ service that prevents networks and servers from monitoring a person's activity across the internet. By blocking iCloud private relay, Apple will not encrypt the traffic leaving the device. Available for devices running macOS 12 and later.
func (m *MacOSGeneralDeviceConfiguration) GetICloudPrivateRelayBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iCloudPrivateRelayBlocked
    }
}
// GetITunesBlockFileSharing gets the iTunesBlockFileSharing property value. Indicates whether or not to block files from being transferred using iTunes.
func (m *MacOSGeneralDeviceConfiguration) GetITunesBlockFileSharing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iTunesBlockFileSharing
    }
}
// GetITunesBlockMusicService gets the iTunesBlockMusicService property value. Indicates whether or not to block Music service and revert Music app to classic mode.
func (m *MacOSGeneralDeviceConfiguration) GetITunesBlockMusicService()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.iTunesBlockMusicService
    }
}
// GetKeyboardBlockDictation gets the keyboardBlockDictation property value. Indicates whether or not to block the user from using dictation input.
func (m *MacOSGeneralDeviceConfiguration) GetKeyboardBlockDictation()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keyboardBlockDictation
    }
}
// GetKeychainBlockCloudSync gets the keychainBlockCloudSync property value. Indicates whether or not iCloud keychain synchronization is blocked (macOS 10.12 and later).
func (m *MacOSGeneralDeviceConfiguration) GetKeychainBlockCloudSync()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.keychainBlockCloudSync
    }
}
// GetMultiplayerGamingBlocked gets the multiplayerGamingBlocked property value. TRUE prevents multiplayer gaming when using Game Center. FALSE allows multiplayer gaming when using Game Center. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) GetMultiplayerGamingBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.multiplayerGamingBlocked
    }
}
// GetPasswordBlockAirDropSharing gets the passwordBlockAirDropSharing property value. Indicates whether or not to block sharing passwords with the AirDrop passwords feature.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordBlockAirDropSharing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordBlockAirDropSharing
    }
}
// GetPasswordBlockAutoFill gets the passwordBlockAutoFill property value. Indicates whether or not to block the AutoFill Passwords feature.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordBlockAutoFill()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordBlockAutoFill
    }
}
// GetPasswordBlockFingerprintUnlock gets the passwordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlock()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordBlockFingerprintUnlock
    }
}
// GetPasswordBlockModification gets the passwordBlockModification property value. Indicates whether or not to allow passcode modification.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordBlockModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordBlockModification
    }
}
// GetPasswordBlockProximityRequests gets the passwordBlockProximityRequests property value. Indicates whether or not to block requesting passwords from nearby devices.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordBlockProximityRequests()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordBlockProximityRequests
    }
}
// GetPasswordBlockSimple gets the passwordBlockSimple property value. Block simple passwords.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordBlockSimple()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordBlockSimple
    }
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordExpirationDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordExpirationDays
    }
}
// GetPasswordMaximumAttemptCount gets the passwordMaximumAttemptCount property value. The number of allowed failed attempts to enter the passcode at the device's lock screen. Valid values 2 to 11
func (m *MacOSGeneralDeviceConfiguration) GetPasswordMaximumAttemptCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMaximumAttemptCount
    }
}
// GetPasswordMinimumCharacterSetCount gets the passwordMinimumCharacterSetCount property value. Number of character sets a password must contain. Valid values 0 to 4
func (m *MacOSGeneralDeviceConfiguration) GetPasswordMinimumCharacterSetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMinimumCharacterSetCount
    }
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum length of passwords.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordMinimumLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMinimumLength
    }
}
// GetPasswordMinutesOfInactivityBeforeLock gets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity required before a password is required.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMinutesOfInactivityBeforeLock
    }
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity required before the screen times out.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMinutesOfInactivityBeforeScreenTimeout
    }
}
// GetPasswordMinutesUntilFailedLoginReset gets the passwordMinutesUntilFailedLoginReset property value. The number of minutes before the login is reset after the maximum number of unsuccessful login attempts is reached.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordMinutesUntilFailedLoginReset()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordMinutesUntilFailedLoginReset
    }
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.passwordPreviousPasswordBlockCount
    }
}
// GetPasswordRequired gets the passwordRequired property value. Whether or not to require a password.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordRequired
    }
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Possible values of required passwords.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordRequiredType()(*RequiredPasswordType) {
    if m == nil {
        return nil
    } else {
        return m.passwordRequiredType
    }
}
// GetPrivacyAccessControls gets the privacyAccessControls property value. List of privacy preference policy controls. This collection can contain a maximum of 10000 elements.
func (m *MacOSGeneralDeviceConfiguration) GetPrivacyAccessControls()([]MacOSPrivacyAccessControlItemable) {
    if m == nil {
        return nil
    } else {
        return m.privacyAccessControls
    }
}
// GetSafariBlockAutofill gets the safariBlockAutofill property value. Indicates whether or not to block the user from using Auto fill in Safari.
func (m *MacOSGeneralDeviceConfiguration) GetSafariBlockAutofill()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.safariBlockAutofill
    }
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether or not to block the user from taking Screenshots.
func (m *MacOSGeneralDeviceConfiguration) GetScreenCaptureBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.screenCaptureBlocked
    }
}
// GetSoftwareUpdateMajorOSDeferredInstallDelayInDays gets the softwareUpdateMajorOSDeferredInstallDelayInDays property value. Specify the number of days (1-90) to delay visibility of major OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) GetSoftwareUpdateMajorOSDeferredInstallDelayInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.softwareUpdateMajorOSDeferredInstallDelayInDays
    }
}
// GetSoftwareUpdateMinorOSDeferredInstallDelayInDays gets the softwareUpdateMinorOSDeferredInstallDelayInDays property value. Specify the number of days (1-90) to delay visibility of minor OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) GetSoftwareUpdateMinorOSDeferredInstallDelayInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.softwareUpdateMinorOSDeferredInstallDelayInDays
    }
}
// GetSoftwareUpdateNonOSDeferredInstallDelayInDays gets the softwareUpdateNonOSDeferredInstallDelayInDays property value. Specify the number of days (1-90) to delay visibility of non-OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) GetSoftwareUpdateNonOSDeferredInstallDelayInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.softwareUpdateNonOSDeferredInstallDelayInDays
    }
}
// GetSoftwareUpdatesEnforcedDelayInDays gets the softwareUpdatesEnforcedDelayInDays property value. Sets how many days a software update will be delyed for a supervised device. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) GetSoftwareUpdatesEnforcedDelayInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.softwareUpdatesEnforcedDelayInDays
    }
}
// GetSpotlightBlockInternetResults gets the spotlightBlockInternetResults property value. Indicates whether or not to block Spotlight from returning any results from an Internet search.
func (m *MacOSGeneralDeviceConfiguration) GetSpotlightBlockInternetResults()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.spotlightBlockInternetResults
    }
}
// GetTouchIdTimeoutInHours gets the touchIdTimeoutInHours property value. Maximum hours after which the user must enter their password to unlock the device instead of using Touch ID. Available for devices running macOS 12 and later. Valid values 0 to 2147483647
func (m *MacOSGeneralDeviceConfiguration) GetTouchIdTimeoutInHours()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.touchIdTimeoutInHours
    }
}
// GetUpdateDelayPolicy gets the updateDelayPolicy property value. Determines whether to delay OS and/or app updates for macOS. Possible values are: none, delayOSUpdateVisibility, delayAppUpdateVisibility, unknownFutureValue, delayMajorOsUpdateVisibility.
func (m *MacOSGeneralDeviceConfiguration) GetUpdateDelayPolicy()(*MacOSSoftwareUpdateDelayPolicy) {
    if m == nil {
        return nil
    } else {
        return m.updateDelayPolicy
    }
}
// GetWallpaperModificationBlocked gets the wallpaperModificationBlocked property value. TRUE prevents the wallpaper from being changed. FALSE allows the wallpaper to be changed. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) GetWallpaperModificationBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.wallpaperModificationBlocked
    }
}
// Serialize serializes information the current object
func (m *MacOSGeneralDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("addingGameCenterFriendsBlocked", m.GetAddingGameCenterFriendsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("airDropBlocked", m.GetAirDropBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appleWatchBlockAutoUnlock", m.GetAppleWatchBlockAutoUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cameraBlocked", m.GetCameraBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("classroomAppBlockRemoteScreenObservation", m.GetClassroomAppBlockRemoteScreenObservation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("classroomAppForceUnpromptedScreenObservation", m.GetClassroomAppForceUnpromptedScreenObservation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("classroomForceAutomaticallyJoinClasses", m.GetClassroomForceAutomaticallyJoinClasses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("classroomForceRequestPermissionToLeaveClasses", m.GetClassroomForceRequestPermissionToLeaveClasses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("classroomForceUnpromptedAppAndDeviceLock", m.GetClassroomForceUnpromptedAppAndDeviceLock())
        if err != nil {
            return err
        }
    }
    if m.GetCompliantAppListType() != nil {
        cast := (*m.GetCompliantAppListType()).String()
        err = writer.WriteStringValue("compliantAppListType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCompliantAppsList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompliantAppsList()))
        for i, v := range m.GetCompliantAppsList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("compliantAppsList", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contentCachingBlocked", m.GetContentCachingBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("definitionLookupBlocked", m.GetDefinitionLookupBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetEmailInDomainSuffixes() != nil {
        err = writer.WriteCollectionOfStringValues("emailInDomainSuffixes", m.GetEmailInDomainSuffixes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("eraseContentAndSettingsBlocked", m.GetEraseContentAndSettingsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("gameCenterBlocked", m.GetGameCenterBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockActivityContinuation", m.GetICloudBlockActivityContinuation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockAddressBook", m.GetICloudBlockAddressBook())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockBookmarks", m.GetICloudBlockBookmarks())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockCalendar", m.GetICloudBlockCalendar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockDocumentSync", m.GetICloudBlockDocumentSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockMail", m.GetICloudBlockMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockNotes", m.GetICloudBlockNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockPhotoLibrary", m.GetICloudBlockPhotoLibrary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockReminders", m.GetICloudBlockReminders())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudDesktopAndDocumentsBlocked", m.GetICloudDesktopAndDocumentsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudPrivateRelayBlocked", m.GetICloudPrivateRelayBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iTunesBlockFileSharing", m.GetITunesBlockFileSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iTunesBlockMusicService", m.GetITunesBlockMusicService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyboardBlockDictation", m.GetKeyboardBlockDictation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keychainBlockCloudSync", m.GetKeychainBlockCloudSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("multiplayerGamingBlocked", m.GetMultiplayerGamingBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockAirDropSharing", m.GetPasswordBlockAirDropSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockAutoFill", m.GetPasswordBlockAutoFill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockFingerprintUnlock", m.GetPasswordBlockFingerprintUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockModification", m.GetPasswordBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockProximityRequests", m.GetPasswordBlockProximityRequests())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockSimple", m.GetPasswordBlockSimple())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationDays", m.GetPasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMaximumAttemptCount", m.GetPasswordMaximumAttemptCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumCharacterSetCount", m.GetPasswordMinimumCharacterSetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeLock", m.GetPasswordMinutesOfInactivityBeforeLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeScreenTimeout", m.GetPasswordMinutesOfInactivityBeforeScreenTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesUntilFailedLoginReset", m.GetPasswordMinutesUntilFailedLoginReset())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordPreviousPasswordBlockCount", m.GetPasswordPreviousPasswordBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordRequired", m.GetPasswordRequired())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRequiredType() != nil {
        cast := (*m.GetPasswordRequiredType()).String()
        err = writer.WriteStringValue("passwordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPrivacyAccessControls() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPrivacyAccessControls()))
        for i, v := range m.GetPrivacyAccessControls() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("privacyAccessControls", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("safariBlockAutofill", m.GetSafariBlockAutofill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("screenCaptureBlocked", m.GetScreenCaptureBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("softwareUpdateMajorOSDeferredInstallDelayInDays", m.GetSoftwareUpdateMajorOSDeferredInstallDelayInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("softwareUpdateMinorOSDeferredInstallDelayInDays", m.GetSoftwareUpdateMinorOSDeferredInstallDelayInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("softwareUpdateNonOSDeferredInstallDelayInDays", m.GetSoftwareUpdateNonOSDeferredInstallDelayInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("softwareUpdatesEnforcedDelayInDays", m.GetSoftwareUpdatesEnforcedDelayInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("spotlightBlockInternetResults", m.GetSpotlightBlockInternetResults())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("touchIdTimeoutInHours", m.GetTouchIdTimeoutInHours())
        if err != nil {
            return err
        }
    }
    if m.GetUpdateDelayPolicy() != nil {
        cast := (*m.GetUpdateDelayPolicy()).String()
        err = writer.WriteStringValue("updateDelayPolicy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wallpaperModificationBlocked", m.GetWallpaperModificationBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddingGameCenterFriendsBlocked sets the addingGameCenterFriendsBlocked property value. Yes prevents users from adding friends to Game Center. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) SetAddingGameCenterFriendsBlocked(value *bool)() {
    if m != nil {
        m.addingGameCenterFriendsBlocked = value
    }
}
// SetAirDropBlocked sets the airDropBlocked property value. Indicates whether or not to allow AirDrop.
func (m *MacOSGeneralDeviceConfiguration) SetAirDropBlocked(value *bool)() {
    if m != nil {
        m.airDropBlocked = value
    }
}
// SetAppleWatchBlockAutoUnlock sets the appleWatchBlockAutoUnlock property value. Indicates whether or to block users from unlocking their Mac with Apple Watch.
func (m *MacOSGeneralDeviceConfiguration) SetAppleWatchBlockAutoUnlock(value *bool)() {
    if m != nil {
        m.appleWatchBlockAutoUnlock = value
    }
}
// SetCameraBlocked sets the cameraBlocked property value. Indicates whether or not to block the user from accessing the camera of the device.
func (m *MacOSGeneralDeviceConfiguration) SetCameraBlocked(value *bool)() {
    if m != nil {
        m.cameraBlocked = value
    }
}
// SetClassroomAppBlockRemoteScreenObservation sets the classroomAppBlockRemoteScreenObservation property value. Indicates whether or not to allow remote screen observation by Classroom app. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) SetClassroomAppBlockRemoteScreenObservation(value *bool)() {
    if m != nil {
        m.classroomAppBlockRemoteScreenObservation = value
    }
}
// SetClassroomAppForceUnpromptedScreenObservation sets the classroomAppForceUnpromptedScreenObservation property value. Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to view a student's screen without prompting. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) SetClassroomAppForceUnpromptedScreenObservation(value *bool)() {
    if m != nil {
        m.classroomAppForceUnpromptedScreenObservation = value
    }
}
// SetClassroomForceAutomaticallyJoinClasses sets the classroomForceAutomaticallyJoinClasses property value. Indicates whether or not to automatically give permission to the teacher's requests, without prompting the student. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) SetClassroomForceAutomaticallyJoinClasses(value *bool)() {
    if m != nil {
        m.classroomForceAutomaticallyJoinClasses = value
    }
}
// SetClassroomForceRequestPermissionToLeaveClasses sets the classroomForceRequestPermissionToLeaveClasses property value. Indicates whether a student enrolled in an unmanaged course via Classroom will be required to request permission from the teacher when attempting to leave the course. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) SetClassroomForceRequestPermissionToLeaveClasses(value *bool)() {
    if m != nil {
        m.classroomForceRequestPermissionToLeaveClasses = value
    }
}
// SetClassroomForceUnpromptedAppAndDeviceLock sets the classroomForceUnpromptedAppAndDeviceLock property value. Indicates whether or not to allow the teacher to lock apps or the device without prompting the student. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) SetClassroomForceUnpromptedAppAndDeviceLock(value *bool)() {
    if m != nil {
        m.classroomForceUnpromptedAppAndDeviceLock = value
    }
}
// SetCompliantAppListType sets the compliantAppListType property value. Possible values of the compliance app list.
func (m *MacOSGeneralDeviceConfiguration) SetCompliantAppListType(value *AppListType)() {
    if m != nil {
        m.compliantAppListType = value
    }
}
// SetCompliantAppsList sets the compliantAppsList property value. List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
func (m *MacOSGeneralDeviceConfiguration) SetCompliantAppsList(value []AppListItemable)() {
    if m != nil {
        m.compliantAppsList = value
    }
}
// SetContentCachingBlocked sets the contentCachingBlocked property value. Indicates whether or not to allow content caching.
func (m *MacOSGeneralDeviceConfiguration) SetContentCachingBlocked(value *bool)() {
    if m != nil {
        m.contentCachingBlocked = value
    }
}
// SetDefinitionLookupBlocked sets the definitionLookupBlocked property value. Indicates whether or not to block definition lookup.
func (m *MacOSGeneralDeviceConfiguration) SetDefinitionLookupBlocked(value *bool)() {
    if m != nil {
        m.definitionLookupBlocked = value
    }
}
// SetEmailInDomainSuffixes sets the emailInDomainSuffixes property value. An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
func (m *MacOSGeneralDeviceConfiguration) SetEmailInDomainSuffixes(value []string)() {
    if m != nil {
        m.emailInDomainSuffixes = value
    }
}
// SetEraseContentAndSettingsBlocked sets the eraseContentAndSettingsBlocked property value. TRUE disables the reset option on supervised devices. FALSE enables the reset option on supervised devices. Available for devices running macOS versions 12.0 and later.
func (m *MacOSGeneralDeviceConfiguration) SetEraseContentAndSettingsBlocked(value *bool)() {
    if m != nil {
        m.eraseContentAndSettingsBlocked = value
    }
}
// SetGameCenterBlocked sets the gameCenterBlocked property value. Yes disables Game Center, and the Game Center icon is removed from the Home screen. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) SetGameCenterBlocked(value *bool)() {
    if m != nil {
        m.gameCenterBlocked = value
    }
}
// SetICloudBlockActivityContinuation sets the iCloudBlockActivityContinuation property value. Indicates whether or not to block the user from continuing work that they started on a MacOS device on another iOS or MacOS device (MacOS 10.15 or later).
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockActivityContinuation(value *bool)() {
    if m != nil {
        m.iCloudBlockActivityContinuation = value
    }
}
// SetICloudBlockAddressBook sets the iCloudBlockAddressBook property value. Indicates whether or not to block iCloud from syncing contacts.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockAddressBook(value *bool)() {
    if m != nil {
        m.iCloudBlockAddressBook = value
    }
}
// SetICloudBlockBookmarks sets the iCloudBlockBookmarks property value. Indicates whether or not to block iCloud from syncing bookmarks.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockBookmarks(value *bool)() {
    if m != nil {
        m.iCloudBlockBookmarks = value
    }
}
// SetICloudBlockCalendar sets the iCloudBlockCalendar property value. Indicates whether or not to block iCloud from syncing calendars.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockCalendar(value *bool)() {
    if m != nil {
        m.iCloudBlockCalendar = value
    }
}
// SetICloudBlockDocumentSync sets the iCloudBlockDocumentSync property value. Indicates whether or not to block iCloud document sync.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockDocumentSync(value *bool)() {
    if m != nil {
        m.iCloudBlockDocumentSync = value
    }
}
// SetICloudBlockMail sets the iCloudBlockMail property value. Indicates whether or not to block iCloud from syncing mail.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockMail(value *bool)() {
    if m != nil {
        m.iCloudBlockMail = value
    }
}
// SetICloudBlockNotes sets the iCloudBlockNotes property value. Indicates whether or not to block iCloud from syncing notes.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockNotes(value *bool)() {
    if m != nil {
        m.iCloudBlockNotes = value
    }
}
// SetICloudBlockPhotoLibrary sets the iCloudBlockPhotoLibrary property value. Indicates whether or not to block iCloud Photo Library.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockPhotoLibrary(value *bool)() {
    if m != nil {
        m.iCloudBlockPhotoLibrary = value
    }
}
// SetICloudBlockReminders sets the iCloudBlockReminders property value. Indicates whether or not to block iCloud from syncing reminders.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockReminders(value *bool)() {
    if m != nil {
        m.iCloudBlockReminders = value
    }
}
// SetICloudDesktopAndDocumentsBlocked sets the iCloudDesktopAndDocumentsBlocked property value. When TRUE the synchronization of cloud desktop and documents is blocked. When FALSE, synchronization of the cloud desktop and documents are allowed. Available for devices running macOS 10.12.4 and later.
func (m *MacOSGeneralDeviceConfiguration) SetICloudDesktopAndDocumentsBlocked(value *bool)() {
    if m != nil {
        m.iCloudDesktopAndDocumentsBlocked = value
    }
}
// SetICloudPrivateRelayBlocked sets the iCloudPrivateRelayBlocked property value. iCloud private relay is an iCloud+ service that prevents networks and servers from monitoring a person's activity across the internet. By blocking iCloud private relay, Apple will not encrypt the traffic leaving the device. Available for devices running macOS 12 and later.
func (m *MacOSGeneralDeviceConfiguration) SetICloudPrivateRelayBlocked(value *bool)() {
    if m != nil {
        m.iCloudPrivateRelayBlocked = value
    }
}
// SetITunesBlockFileSharing sets the iTunesBlockFileSharing property value. Indicates whether or not to block files from being transferred using iTunes.
func (m *MacOSGeneralDeviceConfiguration) SetITunesBlockFileSharing(value *bool)() {
    if m != nil {
        m.iTunesBlockFileSharing = value
    }
}
// SetITunesBlockMusicService sets the iTunesBlockMusicService property value. Indicates whether or not to block Music service and revert Music app to classic mode.
func (m *MacOSGeneralDeviceConfiguration) SetITunesBlockMusicService(value *bool)() {
    if m != nil {
        m.iTunesBlockMusicService = value
    }
}
// SetKeyboardBlockDictation sets the keyboardBlockDictation property value. Indicates whether or not to block the user from using dictation input.
func (m *MacOSGeneralDeviceConfiguration) SetKeyboardBlockDictation(value *bool)() {
    if m != nil {
        m.keyboardBlockDictation = value
    }
}
// SetKeychainBlockCloudSync sets the keychainBlockCloudSync property value. Indicates whether or not iCloud keychain synchronization is blocked (macOS 10.12 and later).
func (m *MacOSGeneralDeviceConfiguration) SetKeychainBlockCloudSync(value *bool)() {
    if m != nil {
        m.keychainBlockCloudSync = value
    }
}
// SetMultiplayerGamingBlocked sets the multiplayerGamingBlocked property value. TRUE prevents multiplayer gaming when using Game Center. FALSE allows multiplayer gaming when using Game Center. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) SetMultiplayerGamingBlocked(value *bool)() {
    if m != nil {
        m.multiplayerGamingBlocked = value
    }
}
// SetPasswordBlockAirDropSharing sets the passwordBlockAirDropSharing property value. Indicates whether or not to block sharing passwords with the AirDrop passwords feature.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordBlockAirDropSharing(value *bool)() {
    if m != nil {
        m.passwordBlockAirDropSharing = value
    }
}
// SetPasswordBlockAutoFill sets the passwordBlockAutoFill property value. Indicates whether or not to block the AutoFill Passwords feature.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordBlockAutoFill(value *bool)() {
    if m != nil {
        m.passwordBlockAutoFill = value
    }
}
// SetPasswordBlockFingerprintUnlock sets the passwordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordBlockFingerprintUnlock(value *bool)() {
    if m != nil {
        m.passwordBlockFingerprintUnlock = value
    }
}
// SetPasswordBlockModification sets the passwordBlockModification property value. Indicates whether or not to allow passcode modification.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordBlockModification(value *bool)() {
    if m != nil {
        m.passwordBlockModification = value
    }
}
// SetPasswordBlockProximityRequests sets the passwordBlockProximityRequests property value. Indicates whether or not to block requesting passwords from nearby devices.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordBlockProximityRequests(value *bool)() {
    if m != nil {
        m.passwordBlockProximityRequests = value
    }
}
// SetPasswordBlockSimple sets the passwordBlockSimple property value. Block simple passwords.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordBlockSimple(value *bool)() {
    if m != nil {
        m.passwordBlockSimple = value
    }
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordExpirationDays(value *int32)() {
    if m != nil {
        m.passwordExpirationDays = value
    }
}
// SetPasswordMaximumAttemptCount sets the passwordMaximumAttemptCount property value. The number of allowed failed attempts to enter the passcode at the device's lock screen. Valid values 2 to 11
func (m *MacOSGeneralDeviceConfiguration) SetPasswordMaximumAttemptCount(value *int32)() {
    if m != nil {
        m.passwordMaximumAttemptCount = value
    }
}
// SetPasswordMinimumCharacterSetCount sets the passwordMinimumCharacterSetCount property value. Number of character sets a password must contain. Valid values 0 to 4
func (m *MacOSGeneralDeviceConfiguration) SetPasswordMinimumCharacterSetCount(value *int32)() {
    if m != nil {
        m.passwordMinimumCharacterSetCount = value
    }
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum length of passwords.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordMinimumLength(value *int32)() {
    if m != nil {
        m.passwordMinimumLength = value
    }
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity required before a password is required.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    if m != nil {
        m.passwordMinutesOfInactivityBeforeLock = value
    }
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity required before the screen times out.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    if m != nil {
        m.passwordMinutesOfInactivityBeforeScreenTimeout = value
    }
}
// SetPasswordMinutesUntilFailedLoginReset sets the passwordMinutesUntilFailedLoginReset property value. The number of minutes before the login is reset after the maximum number of unsuccessful login attempts is reached.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordMinutesUntilFailedLoginReset(value *int32)() {
    if m != nil {
        m.passwordMinutesUntilFailedLoginReset = value
    }
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    if m != nil {
        m.passwordPreviousPasswordBlockCount = value
    }
}
// SetPasswordRequired sets the passwordRequired property value. Whether or not to require a password.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordRequired(value *bool)() {
    if m != nil {
        m.passwordRequired = value
    }
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Possible values of required passwords.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordRequiredType(value *RequiredPasswordType)() {
    if m != nil {
        m.passwordRequiredType = value
    }
}
// SetPrivacyAccessControls sets the privacyAccessControls property value. List of privacy preference policy controls. This collection can contain a maximum of 10000 elements.
func (m *MacOSGeneralDeviceConfiguration) SetPrivacyAccessControls(value []MacOSPrivacyAccessControlItemable)() {
    if m != nil {
        m.privacyAccessControls = value
    }
}
// SetSafariBlockAutofill sets the safariBlockAutofill property value. Indicates whether or not to block the user from using Auto fill in Safari.
func (m *MacOSGeneralDeviceConfiguration) SetSafariBlockAutofill(value *bool)() {
    if m != nil {
        m.safariBlockAutofill = value
    }
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether or not to block the user from taking Screenshots.
func (m *MacOSGeneralDeviceConfiguration) SetScreenCaptureBlocked(value *bool)() {
    if m != nil {
        m.screenCaptureBlocked = value
    }
}
// SetSoftwareUpdateMajorOSDeferredInstallDelayInDays sets the softwareUpdateMajorOSDeferredInstallDelayInDays property value. Specify the number of days (1-90) to delay visibility of major OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) SetSoftwareUpdateMajorOSDeferredInstallDelayInDays(value *int32)() {
    if m != nil {
        m.softwareUpdateMajorOSDeferredInstallDelayInDays = value
    }
}
// SetSoftwareUpdateMinorOSDeferredInstallDelayInDays sets the softwareUpdateMinorOSDeferredInstallDelayInDays property value. Specify the number of days (1-90) to delay visibility of minor OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) SetSoftwareUpdateMinorOSDeferredInstallDelayInDays(value *int32)() {
    if m != nil {
        m.softwareUpdateMinorOSDeferredInstallDelayInDays = value
    }
}
// SetSoftwareUpdateNonOSDeferredInstallDelayInDays sets the softwareUpdateNonOSDeferredInstallDelayInDays property value. Specify the number of days (1-90) to delay visibility of non-OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) SetSoftwareUpdateNonOSDeferredInstallDelayInDays(value *int32)() {
    if m != nil {
        m.softwareUpdateNonOSDeferredInstallDelayInDays = value
    }
}
// SetSoftwareUpdatesEnforcedDelayInDays sets the softwareUpdatesEnforcedDelayInDays property value. Sets how many days a software update will be delyed for a supervised device. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) SetSoftwareUpdatesEnforcedDelayInDays(value *int32)() {
    if m != nil {
        m.softwareUpdatesEnforcedDelayInDays = value
    }
}
// SetSpotlightBlockInternetResults sets the spotlightBlockInternetResults property value. Indicates whether or not to block Spotlight from returning any results from an Internet search.
func (m *MacOSGeneralDeviceConfiguration) SetSpotlightBlockInternetResults(value *bool)() {
    if m != nil {
        m.spotlightBlockInternetResults = value
    }
}
// SetTouchIdTimeoutInHours sets the touchIdTimeoutInHours property value. Maximum hours after which the user must enter their password to unlock the device instead of using Touch ID. Available for devices running macOS 12 and later. Valid values 0 to 2147483647
func (m *MacOSGeneralDeviceConfiguration) SetTouchIdTimeoutInHours(value *int32)() {
    if m != nil {
        m.touchIdTimeoutInHours = value
    }
}
// SetUpdateDelayPolicy sets the updateDelayPolicy property value. Determines whether to delay OS and/or app updates for macOS. Possible values are: none, delayOSUpdateVisibility, delayAppUpdateVisibility, unknownFutureValue, delayMajorOsUpdateVisibility.
func (m *MacOSGeneralDeviceConfiguration) SetUpdateDelayPolicy(value *MacOSSoftwareUpdateDelayPolicy)() {
    if m != nil {
        m.updateDelayPolicy = value
    }
}
// SetWallpaperModificationBlocked sets the wallpaperModificationBlocked property value. TRUE prevents the wallpaper from being changed. FALSE allows the wallpaper to be changed. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) SetWallpaperModificationBlocked(value *bool)() {
    if m != nil {
        m.wallpaperModificationBlocked = value
    }
}
