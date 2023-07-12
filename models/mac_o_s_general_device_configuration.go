package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSGeneralDeviceConfiguration this topic provides descriptions of the declared methods, properties and relationships exposed by the macOSGeneralDeviceConfiguration resource.
type MacOSGeneralDeviceConfiguration struct {
    DeviceConfiguration
}
// NewMacOSGeneralDeviceConfiguration instantiates a new macOSGeneralDeviceConfiguration and sets the default values.
func NewMacOSGeneralDeviceConfiguration()(*MacOSGeneralDeviceConfiguration) {
    m := &MacOSGeneralDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSGeneralDeviceConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSGeneralDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSGeneralDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSGeneralDeviceConfiguration(), nil
}
// GetActivationLockWhenSupervisedAllowed gets the activationLockWhenSupervisedAllowed property value. When TRUE, activation lock is allowed when the devices is in the supervised mode. When FALSE, activation lock is not allowed. Default is false.
func (m *MacOSGeneralDeviceConfiguration) GetActivationLockWhenSupervisedAllowed()(*bool) {
    val, err := m.GetBackingStore().Get("activationLockWhenSupervisedAllowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAddingGameCenterFriendsBlocked gets the addingGameCenterFriendsBlocked property value. Yes prevents users from adding friends to Game Center. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) GetAddingGameCenterFriendsBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("addingGameCenterFriendsBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAirDropBlocked gets the airDropBlocked property value. Indicates whether or not to allow AirDrop.
func (m *MacOSGeneralDeviceConfiguration) GetAirDropBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("airDropBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAppleWatchBlockAutoUnlock gets the appleWatchBlockAutoUnlock property value. Indicates whether or to block users from unlocking their Mac with Apple Watch.
func (m *MacOSGeneralDeviceConfiguration) GetAppleWatchBlockAutoUnlock()(*bool) {
    val, err := m.GetBackingStore().Get("appleWatchBlockAutoUnlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCameraBlocked gets the cameraBlocked property value. Indicates whether or not to block the user from accessing the camera of the device.
func (m *MacOSGeneralDeviceConfiguration) GetCameraBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("cameraBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetClassroomAppBlockRemoteScreenObservation gets the classroomAppBlockRemoteScreenObservation property value. Indicates whether or not to allow remote screen observation by Classroom app. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) GetClassroomAppBlockRemoteScreenObservation()(*bool) {
    val, err := m.GetBackingStore().Get("classroomAppBlockRemoteScreenObservation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetClassroomAppForceUnpromptedScreenObservation gets the classroomAppForceUnpromptedScreenObservation property value. Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to view a student's screen without prompting. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) GetClassroomAppForceUnpromptedScreenObservation()(*bool) {
    val, err := m.GetBackingStore().Get("classroomAppForceUnpromptedScreenObservation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetClassroomForceAutomaticallyJoinClasses gets the classroomForceAutomaticallyJoinClasses property value. Indicates whether or not to automatically give permission to the teacher's requests, without prompting the student. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) GetClassroomForceAutomaticallyJoinClasses()(*bool) {
    val, err := m.GetBackingStore().Get("classroomForceAutomaticallyJoinClasses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetClassroomForceRequestPermissionToLeaveClasses gets the classroomForceRequestPermissionToLeaveClasses property value. Indicates whether a student enrolled in an unmanaged course via Classroom will be required to request permission from the teacher when attempting to leave the course. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) GetClassroomForceRequestPermissionToLeaveClasses()(*bool) {
    val, err := m.GetBackingStore().Get("classroomForceRequestPermissionToLeaveClasses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetClassroomForceUnpromptedAppAndDeviceLock gets the classroomForceUnpromptedAppAndDeviceLock property value. Indicates whether or not to allow the teacher to lock apps or the device without prompting the student. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) GetClassroomForceUnpromptedAppAndDeviceLock()(*bool) {
    val, err := m.GetBackingStore().Get("classroomForceUnpromptedAppAndDeviceLock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCompliantAppListType gets the compliantAppListType property value. Possible values of the compliance app list.
func (m *MacOSGeneralDeviceConfiguration) GetCompliantAppListType()(*AppListType) {
    val, err := m.GetBackingStore().Get("compliantAppListType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppListType)
    }
    return nil
}
// GetCompliantAppsList gets the compliantAppsList property value. List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
func (m *MacOSGeneralDeviceConfiguration) GetCompliantAppsList()([]AppListItemable) {
    val, err := m.GetBackingStore().Get("compliantAppsList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppListItemable)
    }
    return nil
}
// GetContentCachingBlocked gets the contentCachingBlocked property value. Indicates whether or not to allow content caching.
func (m *MacOSGeneralDeviceConfiguration) GetContentCachingBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("contentCachingBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDefinitionLookupBlocked gets the definitionLookupBlocked property value. Indicates whether or not to block definition lookup.
func (m *MacOSGeneralDeviceConfiguration) GetDefinitionLookupBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("definitionLookupBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEmailInDomainSuffixes gets the emailInDomainSuffixes property value. An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
func (m *MacOSGeneralDeviceConfiguration) GetEmailInDomainSuffixes()([]string) {
    val, err := m.GetBackingStore().Get("emailInDomainSuffixes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetEraseContentAndSettingsBlocked gets the eraseContentAndSettingsBlocked property value. TRUE disables the reset option on supervised devices. FALSE enables the reset option on supervised devices. Available for devices running macOS versions 12.0 and later.
func (m *MacOSGeneralDeviceConfiguration) GetEraseContentAndSettingsBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("eraseContentAndSettingsBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSGeneralDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["activationLockWhenSupervisedAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationLockWhenSupervisedAllowed(val)
        }
        return nil
    }
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
                if v != nil {
                    res[i] = v.(AppListItemable)
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
                if v != nil {
                    res[i] = v.(MacOSPrivacyAccessControlItemable)
                }
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
    val, err := m.GetBackingStore().Get("gameCenterBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudBlockActivityContinuation gets the iCloudBlockActivityContinuation property value. Indicates whether or not to block the user from continuing work that they started on a MacOS device on another iOS or MacOS device (MacOS 10.15 or later).
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockActivityContinuation()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudBlockActivityContinuation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudBlockAddressBook gets the iCloudBlockAddressBook property value. Indicates whether or not to block iCloud from syncing contacts.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockAddressBook()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudBlockAddressBook")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudBlockBookmarks gets the iCloudBlockBookmarks property value. Indicates whether or not to block iCloud from syncing bookmarks.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockBookmarks()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudBlockBookmarks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudBlockCalendar gets the iCloudBlockCalendar property value. Indicates whether or not to block iCloud from syncing calendars.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockCalendar()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudBlockCalendar")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudBlockDocumentSync gets the iCloudBlockDocumentSync property value. Indicates whether or not to block iCloud document sync.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockDocumentSync()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudBlockDocumentSync")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudBlockMail gets the iCloudBlockMail property value. Indicates whether or not to block iCloud from syncing mail.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockMail()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudBlockMail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudBlockNotes gets the iCloudBlockNotes property value. Indicates whether or not to block iCloud from syncing notes.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockNotes()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudBlockNotes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudBlockPhotoLibrary gets the iCloudBlockPhotoLibrary property value. Indicates whether or not to block iCloud Photo Library.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockPhotoLibrary()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudBlockPhotoLibrary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudBlockReminders gets the iCloudBlockReminders property value. Indicates whether or not to block iCloud from syncing reminders.
func (m *MacOSGeneralDeviceConfiguration) GetICloudBlockReminders()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudBlockReminders")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudDesktopAndDocumentsBlocked gets the iCloudDesktopAndDocumentsBlocked property value. When TRUE the synchronization of cloud desktop and documents is blocked. When FALSE, synchronization of the cloud desktop and documents are allowed. Available for devices running macOS 10.12.4 and later.
func (m *MacOSGeneralDeviceConfiguration) GetICloudDesktopAndDocumentsBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudDesktopAndDocumentsBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetICloudPrivateRelayBlocked gets the iCloudPrivateRelayBlocked property value. iCloud private relay is an iCloud+ service that prevents networks and servers from monitoring a person's activity across the internet. By blocking iCloud private relay, Apple will not encrypt the traffic leaving the device. Available for devices running macOS 12 and later.
func (m *MacOSGeneralDeviceConfiguration) GetICloudPrivateRelayBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("iCloudPrivateRelayBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetITunesBlockFileSharing gets the iTunesBlockFileSharing property value. Indicates whether or not to block files from being transferred using iTunes.
func (m *MacOSGeneralDeviceConfiguration) GetITunesBlockFileSharing()(*bool) {
    val, err := m.GetBackingStore().Get("iTunesBlockFileSharing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetITunesBlockMusicService gets the iTunesBlockMusicService property value. Indicates whether or not to block Music service and revert Music app to classic mode.
func (m *MacOSGeneralDeviceConfiguration) GetITunesBlockMusicService()(*bool) {
    val, err := m.GetBackingStore().Get("iTunesBlockMusicService")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKeyboardBlockDictation gets the keyboardBlockDictation property value. Indicates whether or not to block the user from using dictation input.
func (m *MacOSGeneralDeviceConfiguration) GetKeyboardBlockDictation()(*bool) {
    val, err := m.GetBackingStore().Get("keyboardBlockDictation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKeychainBlockCloudSync gets the keychainBlockCloudSync property value. Indicates whether or not iCloud keychain synchronization is blocked (macOS 10.12 and later).
func (m *MacOSGeneralDeviceConfiguration) GetKeychainBlockCloudSync()(*bool) {
    val, err := m.GetBackingStore().Get("keychainBlockCloudSync")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMultiplayerGamingBlocked gets the multiplayerGamingBlocked property value. TRUE prevents multiplayer gaming when using Game Center. FALSE allows multiplayer gaming when using Game Center. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) GetMultiplayerGamingBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("multiplayerGamingBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordBlockAirDropSharing gets the passwordBlockAirDropSharing property value. Indicates whether or not to block sharing passwords with the AirDrop passwords feature.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordBlockAirDropSharing()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockAirDropSharing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordBlockAutoFill gets the passwordBlockAutoFill property value. Indicates whether or not to block the AutoFill Passwords feature.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordBlockAutoFill()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockAutoFill")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordBlockFingerprintUnlock gets the passwordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordBlockFingerprintUnlock()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockFingerprintUnlock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordBlockModification gets the passwordBlockModification property value. Indicates whether or not to allow passcode modification.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordBlockModification()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockModification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordBlockProximityRequests gets the passwordBlockProximityRequests property value. Indicates whether or not to block requesting passwords from nearby devices.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordBlockProximityRequests()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockProximityRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordBlockSimple gets the passwordBlockSimple property value. Block simple passwords.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordBlockSimple()(*bool) {
    val, err := m.GetBackingStore().Get("passwordBlockSimple")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordExpirationDays()(*int32) {
    val, err := m.GetBackingStore().Get("passwordExpirationDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMaximumAttemptCount gets the passwordMaximumAttemptCount property value. The number of allowed failed attempts to enter the passcode at the device's lock screen. Valid values 2 to 11
func (m *MacOSGeneralDeviceConfiguration) GetPasswordMaximumAttemptCount()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMaximumAttemptCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumCharacterSetCount gets the passwordMinimumCharacterSetCount property value. Number of character sets a password must contain. Valid values 0 to 4
func (m *MacOSGeneralDeviceConfiguration) GetPasswordMinimumCharacterSetCount()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumCharacterSetCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum length of passwords.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordMinimumLength()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinimumLength")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinutesOfInactivityBeforeLock gets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity required before a password is required.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinutesOfInactivityBeforeLock")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity required before the screen times out.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinutesOfInactivityBeforeScreenTimeout")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordMinutesUntilFailedLoginReset gets the passwordMinutesUntilFailedLoginReset property value. The number of minutes before the login is reset after the maximum number of unsuccessful login attempts is reached.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordMinutesUntilFailedLoginReset()(*int32) {
    val, err := m.GetBackingStore().Get("passwordMinutesUntilFailedLoginReset")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordPreviousPasswordBlockCount()(*int32) {
    val, err := m.GetBackingStore().Get("passwordPreviousPasswordBlockCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPasswordRequired gets the passwordRequired property value. Whether or not to require a password.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordRequired()(*bool) {
    val, err := m.GetBackingStore().Get("passwordRequired")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Possible values of required passwords.
func (m *MacOSGeneralDeviceConfiguration) GetPasswordRequiredType()(*RequiredPasswordType) {
    val, err := m.GetBackingStore().Get("passwordRequiredType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RequiredPasswordType)
    }
    return nil
}
// GetPrivacyAccessControls gets the privacyAccessControls property value. List of privacy preference policy controls. This collection can contain a maximum of 10000 elements.
func (m *MacOSGeneralDeviceConfiguration) GetPrivacyAccessControls()([]MacOSPrivacyAccessControlItemable) {
    val, err := m.GetBackingStore().Get("privacyAccessControls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSPrivacyAccessControlItemable)
    }
    return nil
}
// GetSafariBlockAutofill gets the safariBlockAutofill property value. Indicates whether or not to block the user from using Auto fill in Safari.
func (m *MacOSGeneralDeviceConfiguration) GetSafariBlockAutofill()(*bool) {
    val, err := m.GetBackingStore().Get("safariBlockAutofill")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether or not to block the user from taking Screenshots.
func (m *MacOSGeneralDeviceConfiguration) GetScreenCaptureBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("screenCaptureBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSoftwareUpdateMajorOSDeferredInstallDelayInDays gets the softwareUpdateMajorOSDeferredInstallDelayInDays property value. Specify the number of days (1-90) to delay visibility of major OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) GetSoftwareUpdateMajorOSDeferredInstallDelayInDays()(*int32) {
    val, err := m.GetBackingStore().Get("softwareUpdateMajorOSDeferredInstallDelayInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSoftwareUpdateMinorOSDeferredInstallDelayInDays gets the softwareUpdateMinorOSDeferredInstallDelayInDays property value. Specify the number of days (1-90) to delay visibility of minor OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) GetSoftwareUpdateMinorOSDeferredInstallDelayInDays()(*int32) {
    val, err := m.GetBackingStore().Get("softwareUpdateMinorOSDeferredInstallDelayInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSoftwareUpdateNonOSDeferredInstallDelayInDays gets the softwareUpdateNonOSDeferredInstallDelayInDays property value. Specify the number of days (1-90) to delay visibility of non-OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) GetSoftwareUpdateNonOSDeferredInstallDelayInDays()(*int32) {
    val, err := m.GetBackingStore().Get("softwareUpdateNonOSDeferredInstallDelayInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSoftwareUpdatesEnforcedDelayInDays gets the softwareUpdatesEnforcedDelayInDays property value. Sets how many days a software update will be delyed for a supervised device. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) GetSoftwareUpdatesEnforcedDelayInDays()(*int32) {
    val, err := m.GetBackingStore().Get("softwareUpdatesEnforcedDelayInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSpotlightBlockInternetResults gets the spotlightBlockInternetResults property value. Indicates whether or not to block Spotlight from returning any results from an Internet search.
func (m *MacOSGeneralDeviceConfiguration) GetSpotlightBlockInternetResults()(*bool) {
    val, err := m.GetBackingStore().Get("spotlightBlockInternetResults")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTouchIdTimeoutInHours gets the touchIdTimeoutInHours property value. Maximum hours after which the user must enter their password to unlock the device instead of using Touch ID. Available for devices running macOS 12 and later. Valid values 0 to 2147483647
func (m *MacOSGeneralDeviceConfiguration) GetTouchIdTimeoutInHours()(*int32) {
    val, err := m.GetBackingStore().Get("touchIdTimeoutInHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUpdateDelayPolicy gets the updateDelayPolicy property value. Determines whether to delay OS and/or app updates for macOS. Possible values are: none, delayOSUpdateVisibility, delayAppUpdateVisibility, unknownFutureValue, delayMajorOsUpdateVisibility.
func (m *MacOSGeneralDeviceConfiguration) GetUpdateDelayPolicy()(*MacOSSoftwareUpdateDelayPolicy) {
    val, err := m.GetBackingStore().Get("updateDelayPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MacOSSoftwareUpdateDelayPolicy)
    }
    return nil
}
// GetWallpaperModificationBlocked gets the wallpaperModificationBlocked property value. TRUE prevents the wallpaper from being changed. FALSE allows the wallpaper to be changed. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) GetWallpaperModificationBlocked()(*bool) {
    val, err := m.GetBackingStore().Get("wallpaperModificationBlocked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSGeneralDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("activationLockWhenSupervisedAllowed", m.GetActivationLockWhenSupervisedAllowed())
        if err != nil {
            return err
        }
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
// SetActivationLockWhenSupervisedAllowed sets the activationLockWhenSupervisedAllowed property value. When TRUE, activation lock is allowed when the devices is in the supervised mode. When FALSE, activation lock is not allowed. Default is false.
func (m *MacOSGeneralDeviceConfiguration) SetActivationLockWhenSupervisedAllowed(value *bool)() {
    err := m.GetBackingStore().Set("activationLockWhenSupervisedAllowed", value)
    if err != nil {
        panic(err)
    }
}
// SetAddingGameCenterFriendsBlocked sets the addingGameCenterFriendsBlocked property value. Yes prevents users from adding friends to Game Center. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) SetAddingGameCenterFriendsBlocked(value *bool)() {
    err := m.GetBackingStore().Set("addingGameCenterFriendsBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetAirDropBlocked sets the airDropBlocked property value. Indicates whether or not to allow AirDrop.
func (m *MacOSGeneralDeviceConfiguration) SetAirDropBlocked(value *bool)() {
    err := m.GetBackingStore().Set("airDropBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetAppleWatchBlockAutoUnlock sets the appleWatchBlockAutoUnlock property value. Indicates whether or to block users from unlocking their Mac with Apple Watch.
func (m *MacOSGeneralDeviceConfiguration) SetAppleWatchBlockAutoUnlock(value *bool)() {
    err := m.GetBackingStore().Set("appleWatchBlockAutoUnlock", value)
    if err != nil {
        panic(err)
    }
}
// SetCameraBlocked sets the cameraBlocked property value. Indicates whether or not to block the user from accessing the camera of the device.
func (m *MacOSGeneralDeviceConfiguration) SetCameraBlocked(value *bool)() {
    err := m.GetBackingStore().Set("cameraBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetClassroomAppBlockRemoteScreenObservation sets the classroomAppBlockRemoteScreenObservation property value. Indicates whether or not to allow remote screen observation by Classroom app. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) SetClassroomAppBlockRemoteScreenObservation(value *bool)() {
    err := m.GetBackingStore().Set("classroomAppBlockRemoteScreenObservation", value)
    if err != nil {
        panic(err)
    }
}
// SetClassroomAppForceUnpromptedScreenObservation sets the classroomAppForceUnpromptedScreenObservation property value. Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to view a student's screen without prompting. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) SetClassroomAppForceUnpromptedScreenObservation(value *bool)() {
    err := m.GetBackingStore().Set("classroomAppForceUnpromptedScreenObservation", value)
    if err != nil {
        panic(err)
    }
}
// SetClassroomForceAutomaticallyJoinClasses sets the classroomForceAutomaticallyJoinClasses property value. Indicates whether or not to automatically give permission to the teacher's requests, without prompting the student. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) SetClassroomForceAutomaticallyJoinClasses(value *bool)() {
    err := m.GetBackingStore().Set("classroomForceAutomaticallyJoinClasses", value)
    if err != nil {
        panic(err)
    }
}
// SetClassroomForceRequestPermissionToLeaveClasses sets the classroomForceRequestPermissionToLeaveClasses property value. Indicates whether a student enrolled in an unmanaged course via Classroom will be required to request permission from the teacher when attempting to leave the course. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) SetClassroomForceRequestPermissionToLeaveClasses(value *bool)() {
    err := m.GetBackingStore().Set("classroomForceRequestPermissionToLeaveClasses", value)
    if err != nil {
        panic(err)
    }
}
// SetClassroomForceUnpromptedAppAndDeviceLock sets the classroomForceUnpromptedAppAndDeviceLock property value. Indicates whether or not to allow the teacher to lock apps or the device without prompting the student. Requires MDM enrollment via Apple School Manager or Apple Business Manager.
func (m *MacOSGeneralDeviceConfiguration) SetClassroomForceUnpromptedAppAndDeviceLock(value *bool)() {
    err := m.GetBackingStore().Set("classroomForceUnpromptedAppAndDeviceLock", value)
    if err != nil {
        panic(err)
    }
}
// SetCompliantAppListType sets the compliantAppListType property value. Possible values of the compliance app list.
func (m *MacOSGeneralDeviceConfiguration) SetCompliantAppListType(value *AppListType)() {
    err := m.GetBackingStore().Set("compliantAppListType", value)
    if err != nil {
        panic(err)
    }
}
// SetCompliantAppsList sets the compliantAppsList property value. List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
func (m *MacOSGeneralDeviceConfiguration) SetCompliantAppsList(value []AppListItemable)() {
    err := m.GetBackingStore().Set("compliantAppsList", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCachingBlocked sets the contentCachingBlocked property value. Indicates whether or not to allow content caching.
func (m *MacOSGeneralDeviceConfiguration) SetContentCachingBlocked(value *bool)() {
    err := m.GetBackingStore().Set("contentCachingBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetDefinitionLookupBlocked sets the definitionLookupBlocked property value. Indicates whether or not to block definition lookup.
func (m *MacOSGeneralDeviceConfiguration) SetDefinitionLookupBlocked(value *bool)() {
    err := m.GetBackingStore().Set("definitionLookupBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailInDomainSuffixes sets the emailInDomainSuffixes property value. An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
func (m *MacOSGeneralDeviceConfiguration) SetEmailInDomainSuffixes(value []string)() {
    err := m.GetBackingStore().Set("emailInDomainSuffixes", value)
    if err != nil {
        panic(err)
    }
}
// SetEraseContentAndSettingsBlocked sets the eraseContentAndSettingsBlocked property value. TRUE disables the reset option on supervised devices. FALSE enables the reset option on supervised devices. Available for devices running macOS versions 12.0 and later.
func (m *MacOSGeneralDeviceConfiguration) SetEraseContentAndSettingsBlocked(value *bool)() {
    err := m.GetBackingStore().Set("eraseContentAndSettingsBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetGameCenterBlocked sets the gameCenterBlocked property value. Yes disables Game Center, and the Game Center icon is removed from the Home screen. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) SetGameCenterBlocked(value *bool)() {
    err := m.GetBackingStore().Set("gameCenterBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudBlockActivityContinuation sets the iCloudBlockActivityContinuation property value. Indicates whether or not to block the user from continuing work that they started on a MacOS device on another iOS or MacOS device (MacOS 10.15 or later).
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockActivityContinuation(value *bool)() {
    err := m.GetBackingStore().Set("iCloudBlockActivityContinuation", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudBlockAddressBook sets the iCloudBlockAddressBook property value. Indicates whether or not to block iCloud from syncing contacts.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockAddressBook(value *bool)() {
    err := m.GetBackingStore().Set("iCloudBlockAddressBook", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudBlockBookmarks sets the iCloudBlockBookmarks property value. Indicates whether or not to block iCloud from syncing bookmarks.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockBookmarks(value *bool)() {
    err := m.GetBackingStore().Set("iCloudBlockBookmarks", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudBlockCalendar sets the iCloudBlockCalendar property value. Indicates whether or not to block iCloud from syncing calendars.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockCalendar(value *bool)() {
    err := m.GetBackingStore().Set("iCloudBlockCalendar", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudBlockDocumentSync sets the iCloudBlockDocumentSync property value. Indicates whether or not to block iCloud document sync.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockDocumentSync(value *bool)() {
    err := m.GetBackingStore().Set("iCloudBlockDocumentSync", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudBlockMail sets the iCloudBlockMail property value. Indicates whether or not to block iCloud from syncing mail.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockMail(value *bool)() {
    err := m.GetBackingStore().Set("iCloudBlockMail", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudBlockNotes sets the iCloudBlockNotes property value. Indicates whether or not to block iCloud from syncing notes.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockNotes(value *bool)() {
    err := m.GetBackingStore().Set("iCloudBlockNotes", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudBlockPhotoLibrary sets the iCloudBlockPhotoLibrary property value. Indicates whether or not to block iCloud Photo Library.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockPhotoLibrary(value *bool)() {
    err := m.GetBackingStore().Set("iCloudBlockPhotoLibrary", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudBlockReminders sets the iCloudBlockReminders property value. Indicates whether or not to block iCloud from syncing reminders.
func (m *MacOSGeneralDeviceConfiguration) SetICloudBlockReminders(value *bool)() {
    err := m.GetBackingStore().Set("iCloudBlockReminders", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudDesktopAndDocumentsBlocked sets the iCloudDesktopAndDocumentsBlocked property value. When TRUE the synchronization of cloud desktop and documents is blocked. When FALSE, synchronization of the cloud desktop and documents are allowed. Available for devices running macOS 10.12.4 and later.
func (m *MacOSGeneralDeviceConfiguration) SetICloudDesktopAndDocumentsBlocked(value *bool)() {
    err := m.GetBackingStore().Set("iCloudDesktopAndDocumentsBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetICloudPrivateRelayBlocked sets the iCloudPrivateRelayBlocked property value. iCloud private relay is an iCloud+ service that prevents networks and servers from monitoring a person's activity across the internet. By blocking iCloud private relay, Apple will not encrypt the traffic leaving the device. Available for devices running macOS 12 and later.
func (m *MacOSGeneralDeviceConfiguration) SetICloudPrivateRelayBlocked(value *bool)() {
    err := m.GetBackingStore().Set("iCloudPrivateRelayBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetITunesBlockFileSharing sets the iTunesBlockFileSharing property value. Indicates whether or not to block files from being transferred using iTunes.
func (m *MacOSGeneralDeviceConfiguration) SetITunesBlockFileSharing(value *bool)() {
    err := m.GetBackingStore().Set("iTunesBlockFileSharing", value)
    if err != nil {
        panic(err)
    }
}
// SetITunesBlockMusicService sets the iTunesBlockMusicService property value. Indicates whether or not to block Music service and revert Music app to classic mode.
func (m *MacOSGeneralDeviceConfiguration) SetITunesBlockMusicService(value *bool)() {
    err := m.GetBackingStore().Set("iTunesBlockMusicService", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyboardBlockDictation sets the keyboardBlockDictation property value. Indicates whether or not to block the user from using dictation input.
func (m *MacOSGeneralDeviceConfiguration) SetKeyboardBlockDictation(value *bool)() {
    err := m.GetBackingStore().Set("keyboardBlockDictation", value)
    if err != nil {
        panic(err)
    }
}
// SetKeychainBlockCloudSync sets the keychainBlockCloudSync property value. Indicates whether or not iCloud keychain synchronization is blocked (macOS 10.12 and later).
func (m *MacOSGeneralDeviceConfiguration) SetKeychainBlockCloudSync(value *bool)() {
    err := m.GetBackingStore().Set("keychainBlockCloudSync", value)
    if err != nil {
        panic(err)
    }
}
// SetMultiplayerGamingBlocked sets the multiplayerGamingBlocked property value. TRUE prevents multiplayer gaming when using Game Center. FALSE allows multiplayer gaming when using Game Center. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) SetMultiplayerGamingBlocked(value *bool)() {
    err := m.GetBackingStore().Set("multiplayerGamingBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockAirDropSharing sets the passwordBlockAirDropSharing property value. Indicates whether or not to block sharing passwords with the AirDrop passwords feature.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordBlockAirDropSharing(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockAirDropSharing", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockAutoFill sets the passwordBlockAutoFill property value. Indicates whether or not to block the AutoFill Passwords feature.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordBlockAutoFill(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockAutoFill", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockFingerprintUnlock sets the passwordBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordBlockFingerprintUnlock(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockFingerprintUnlock", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockModification sets the passwordBlockModification property value. Indicates whether or not to allow passcode modification.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordBlockModification(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockModification", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockProximityRequests sets the passwordBlockProximityRequests property value. Indicates whether or not to block requesting passwords from nearby devices.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordBlockProximityRequests(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockProximityRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordBlockSimple sets the passwordBlockSimple property value. Block simple passwords.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordBlockSimple(value *bool)() {
    err := m.GetBackingStore().Set("passwordBlockSimple", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordExpirationDays(value *int32)() {
    err := m.GetBackingStore().Set("passwordExpirationDays", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMaximumAttemptCount sets the passwordMaximumAttemptCount property value. The number of allowed failed attempts to enter the passcode at the device's lock screen. Valid values 2 to 11
func (m *MacOSGeneralDeviceConfiguration) SetPasswordMaximumAttemptCount(value *int32)() {
    err := m.GetBackingStore().Set("passwordMaximumAttemptCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumCharacterSetCount sets the passwordMinimumCharacterSetCount property value. Number of character sets a password must contain. Valid values 0 to 4
func (m *MacOSGeneralDeviceConfiguration) SetPasswordMinimumCharacterSetCount(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumCharacterSetCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum length of passwords.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordMinimumLength(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinimumLength", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity required before a password is required.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinutesOfInactivityBeforeLock", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity required before the screen times out.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinutesOfInactivityBeforeScreenTimeout", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordMinutesUntilFailedLoginReset sets the passwordMinutesUntilFailedLoginReset property value. The number of minutes before the login is reset after the maximum number of unsuccessful login attempts is reached.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordMinutesUntilFailedLoginReset(value *int32)() {
    err := m.GetBackingStore().Set("passwordMinutesUntilFailedLoginReset", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    err := m.GetBackingStore().Set("passwordPreviousPasswordBlockCount", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequired sets the passwordRequired property value. Whether or not to require a password.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordRequired(value *bool)() {
    err := m.GetBackingStore().Set("passwordRequired", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Possible values of required passwords.
func (m *MacOSGeneralDeviceConfiguration) SetPasswordRequiredType(value *RequiredPasswordType)() {
    err := m.GetBackingStore().Set("passwordRequiredType", value)
    if err != nil {
        panic(err)
    }
}
// SetPrivacyAccessControls sets the privacyAccessControls property value. List of privacy preference policy controls. This collection can contain a maximum of 10000 elements.
func (m *MacOSGeneralDeviceConfiguration) SetPrivacyAccessControls(value []MacOSPrivacyAccessControlItemable)() {
    err := m.GetBackingStore().Set("privacyAccessControls", value)
    if err != nil {
        panic(err)
    }
}
// SetSafariBlockAutofill sets the safariBlockAutofill property value. Indicates whether or not to block the user from using Auto fill in Safari.
func (m *MacOSGeneralDeviceConfiguration) SetSafariBlockAutofill(value *bool)() {
    err := m.GetBackingStore().Set("safariBlockAutofill", value)
    if err != nil {
        panic(err)
    }
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether or not to block the user from taking Screenshots.
func (m *MacOSGeneralDeviceConfiguration) SetScreenCaptureBlocked(value *bool)() {
    err := m.GetBackingStore().Set("screenCaptureBlocked", value)
    if err != nil {
        panic(err)
    }
}
// SetSoftwareUpdateMajorOSDeferredInstallDelayInDays sets the softwareUpdateMajorOSDeferredInstallDelayInDays property value. Specify the number of days (1-90) to delay visibility of major OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) SetSoftwareUpdateMajorOSDeferredInstallDelayInDays(value *int32)() {
    err := m.GetBackingStore().Set("softwareUpdateMajorOSDeferredInstallDelayInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetSoftwareUpdateMinorOSDeferredInstallDelayInDays sets the softwareUpdateMinorOSDeferredInstallDelayInDays property value. Specify the number of days (1-90) to delay visibility of minor OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) SetSoftwareUpdateMinorOSDeferredInstallDelayInDays(value *int32)() {
    err := m.GetBackingStore().Set("softwareUpdateMinorOSDeferredInstallDelayInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetSoftwareUpdateNonOSDeferredInstallDelayInDays sets the softwareUpdateNonOSDeferredInstallDelayInDays property value. Specify the number of days (1-90) to delay visibility of non-OS software updates. Available for devices running macOS versions 11.3 and later. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) SetSoftwareUpdateNonOSDeferredInstallDelayInDays(value *int32)() {
    err := m.GetBackingStore().Set("softwareUpdateNonOSDeferredInstallDelayInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetSoftwareUpdatesEnforcedDelayInDays sets the softwareUpdatesEnforcedDelayInDays property value. Sets how many days a software update will be delyed for a supervised device. Valid values 0 to 90
func (m *MacOSGeneralDeviceConfiguration) SetSoftwareUpdatesEnforcedDelayInDays(value *int32)() {
    err := m.GetBackingStore().Set("softwareUpdatesEnforcedDelayInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetSpotlightBlockInternetResults sets the spotlightBlockInternetResults property value. Indicates whether or not to block Spotlight from returning any results from an Internet search.
func (m *MacOSGeneralDeviceConfiguration) SetSpotlightBlockInternetResults(value *bool)() {
    err := m.GetBackingStore().Set("spotlightBlockInternetResults", value)
    if err != nil {
        panic(err)
    }
}
// SetTouchIdTimeoutInHours sets the touchIdTimeoutInHours property value. Maximum hours after which the user must enter their password to unlock the device instead of using Touch ID. Available for devices running macOS 12 and later. Valid values 0 to 2147483647
func (m *MacOSGeneralDeviceConfiguration) SetTouchIdTimeoutInHours(value *int32)() {
    err := m.GetBackingStore().Set("touchIdTimeoutInHours", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateDelayPolicy sets the updateDelayPolicy property value. Determines whether to delay OS and/or app updates for macOS. Possible values are: none, delayOSUpdateVisibility, delayAppUpdateVisibility, unknownFutureValue, delayMajorOsUpdateVisibility.
func (m *MacOSGeneralDeviceConfiguration) SetUpdateDelayPolicy(value *MacOSSoftwareUpdateDelayPolicy)() {
    err := m.GetBackingStore().Set("updateDelayPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetWallpaperModificationBlocked sets the wallpaperModificationBlocked property value. TRUE prevents the wallpaper from being changed. FALSE allows the wallpaper to be changed. Available for devices running macOS versions 10.13 and later.
func (m *MacOSGeneralDeviceConfiguration) SetWallpaperModificationBlocked(value *bool)() {
    err := m.GetBackingStore().Set("wallpaperModificationBlocked", value)
    if err != nil {
        panic(err)
    }
}
// MacOSGeneralDeviceConfigurationable 
type MacOSGeneralDeviceConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivationLockWhenSupervisedAllowed()(*bool)
    GetAddingGameCenterFriendsBlocked()(*bool)
    GetAirDropBlocked()(*bool)
    GetAppleWatchBlockAutoUnlock()(*bool)
    GetCameraBlocked()(*bool)
    GetClassroomAppBlockRemoteScreenObservation()(*bool)
    GetClassroomAppForceUnpromptedScreenObservation()(*bool)
    GetClassroomForceAutomaticallyJoinClasses()(*bool)
    GetClassroomForceRequestPermissionToLeaveClasses()(*bool)
    GetClassroomForceUnpromptedAppAndDeviceLock()(*bool)
    GetCompliantAppListType()(*AppListType)
    GetCompliantAppsList()([]AppListItemable)
    GetContentCachingBlocked()(*bool)
    GetDefinitionLookupBlocked()(*bool)
    GetEmailInDomainSuffixes()([]string)
    GetEraseContentAndSettingsBlocked()(*bool)
    GetGameCenterBlocked()(*bool)
    GetICloudBlockActivityContinuation()(*bool)
    GetICloudBlockAddressBook()(*bool)
    GetICloudBlockBookmarks()(*bool)
    GetICloudBlockCalendar()(*bool)
    GetICloudBlockDocumentSync()(*bool)
    GetICloudBlockMail()(*bool)
    GetICloudBlockNotes()(*bool)
    GetICloudBlockPhotoLibrary()(*bool)
    GetICloudBlockReminders()(*bool)
    GetICloudDesktopAndDocumentsBlocked()(*bool)
    GetICloudPrivateRelayBlocked()(*bool)
    GetITunesBlockFileSharing()(*bool)
    GetITunesBlockMusicService()(*bool)
    GetKeyboardBlockDictation()(*bool)
    GetKeychainBlockCloudSync()(*bool)
    GetMultiplayerGamingBlocked()(*bool)
    GetPasswordBlockAirDropSharing()(*bool)
    GetPasswordBlockAutoFill()(*bool)
    GetPasswordBlockFingerprintUnlock()(*bool)
    GetPasswordBlockModification()(*bool)
    GetPasswordBlockProximityRequests()(*bool)
    GetPasswordBlockSimple()(*bool)
    GetPasswordExpirationDays()(*int32)
    GetPasswordMaximumAttemptCount()(*int32)
    GetPasswordMinimumCharacterSetCount()(*int32)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinutesOfInactivityBeforeLock()(*int32)
    GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32)
    GetPasswordMinutesUntilFailedLoginReset()(*int32)
    GetPasswordPreviousPasswordBlockCount()(*int32)
    GetPasswordRequired()(*bool)
    GetPasswordRequiredType()(*RequiredPasswordType)
    GetPrivacyAccessControls()([]MacOSPrivacyAccessControlItemable)
    GetSafariBlockAutofill()(*bool)
    GetScreenCaptureBlocked()(*bool)
    GetSoftwareUpdateMajorOSDeferredInstallDelayInDays()(*int32)
    GetSoftwareUpdateMinorOSDeferredInstallDelayInDays()(*int32)
    GetSoftwareUpdateNonOSDeferredInstallDelayInDays()(*int32)
    GetSoftwareUpdatesEnforcedDelayInDays()(*int32)
    GetSpotlightBlockInternetResults()(*bool)
    GetTouchIdTimeoutInHours()(*int32)
    GetUpdateDelayPolicy()(*MacOSSoftwareUpdateDelayPolicy)
    GetWallpaperModificationBlocked()(*bool)
    SetActivationLockWhenSupervisedAllowed(value *bool)()
    SetAddingGameCenterFriendsBlocked(value *bool)()
    SetAirDropBlocked(value *bool)()
    SetAppleWatchBlockAutoUnlock(value *bool)()
    SetCameraBlocked(value *bool)()
    SetClassroomAppBlockRemoteScreenObservation(value *bool)()
    SetClassroomAppForceUnpromptedScreenObservation(value *bool)()
    SetClassroomForceAutomaticallyJoinClasses(value *bool)()
    SetClassroomForceRequestPermissionToLeaveClasses(value *bool)()
    SetClassroomForceUnpromptedAppAndDeviceLock(value *bool)()
    SetCompliantAppListType(value *AppListType)()
    SetCompliantAppsList(value []AppListItemable)()
    SetContentCachingBlocked(value *bool)()
    SetDefinitionLookupBlocked(value *bool)()
    SetEmailInDomainSuffixes(value []string)()
    SetEraseContentAndSettingsBlocked(value *bool)()
    SetGameCenterBlocked(value *bool)()
    SetICloudBlockActivityContinuation(value *bool)()
    SetICloudBlockAddressBook(value *bool)()
    SetICloudBlockBookmarks(value *bool)()
    SetICloudBlockCalendar(value *bool)()
    SetICloudBlockDocumentSync(value *bool)()
    SetICloudBlockMail(value *bool)()
    SetICloudBlockNotes(value *bool)()
    SetICloudBlockPhotoLibrary(value *bool)()
    SetICloudBlockReminders(value *bool)()
    SetICloudDesktopAndDocumentsBlocked(value *bool)()
    SetICloudPrivateRelayBlocked(value *bool)()
    SetITunesBlockFileSharing(value *bool)()
    SetITunesBlockMusicService(value *bool)()
    SetKeyboardBlockDictation(value *bool)()
    SetKeychainBlockCloudSync(value *bool)()
    SetMultiplayerGamingBlocked(value *bool)()
    SetPasswordBlockAirDropSharing(value *bool)()
    SetPasswordBlockAutoFill(value *bool)()
    SetPasswordBlockFingerprintUnlock(value *bool)()
    SetPasswordBlockModification(value *bool)()
    SetPasswordBlockProximityRequests(value *bool)()
    SetPasswordBlockSimple(value *bool)()
    SetPasswordExpirationDays(value *int32)()
    SetPasswordMaximumAttemptCount(value *int32)()
    SetPasswordMinimumCharacterSetCount(value *int32)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinutesOfInactivityBeforeLock(value *int32)()
    SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)()
    SetPasswordMinutesUntilFailedLoginReset(value *int32)()
    SetPasswordPreviousPasswordBlockCount(value *int32)()
    SetPasswordRequired(value *bool)()
    SetPasswordRequiredType(value *RequiredPasswordType)()
    SetPrivacyAccessControls(value []MacOSPrivacyAccessControlItemable)()
    SetSafariBlockAutofill(value *bool)()
    SetScreenCaptureBlocked(value *bool)()
    SetSoftwareUpdateMajorOSDeferredInstallDelayInDays(value *int32)()
    SetSoftwareUpdateMinorOSDeferredInstallDelayInDays(value *int32)()
    SetSoftwareUpdateNonOSDeferredInstallDelayInDays(value *int32)()
    SetSoftwareUpdatesEnforcedDelayInDays(value *int32)()
    SetSpotlightBlockInternetResults(value *bool)()
    SetTouchIdTimeoutInHours(value *int32)()
    SetUpdateDelayPolicy(value *MacOSSoftwareUpdateDelayPolicy)()
    SetWallpaperModificationBlocked(value *bool)()
}
