package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdateForBusinessConfiguration 
type WindowsUpdateForBusinessConfiguration struct {
    DeviceConfiguration
    // Allow eligible Windows 10 devices to upgrade to the latest version of Windows 11.
    allowWindows11Upgrade *bool
    // Possible values for automatic update mode.
    automaticUpdateMode *AutomaticUpdateMode
    // Auto restart required notification dismissal method
    autoRestartNotificationDismissal *AutoRestartNotificationDismissalMethod
    // Which branch devices will receive their updates from
    businessReadyUpdatesOnly *WindowsUpdateType
    // Number of days before feature updates are installed automatically with valid range from 0 to 30 days
    deadlineForFeatureUpdatesInDays *int32
    // Number of days before quality updates are installed automatically with valid range from 0 to 30 days
    deadlineForQualityUpdatesInDays *int32
    // Number of days after deadline  until restarts occur automatically with valid range from 0 to 7 days
    deadlineGracePeriodInDays *int32
    // Delivery optimization mode for peer distribution
    deliveryOptimizationMode *WindowsDeliveryOptimizationMode
    // Windows update for business configuration device states. This collection can contain a maximum of 500 elements.
    deviceUpdateStates []WindowsUpdateStateable
    // Exclude Windows update Drivers
    driversExcluded *bool
    // Deadline in days before automatically scheduling and executing a pending restart outside of active hours, with valid range from 2 to 30 days
    engagedRestartDeadlineInDays *int32
    // Number of days a user can snooze Engaged Restart reminder notifications with valid range from 1 to 3 days
    engagedRestartSnoozeScheduleInDays *int32
    // Number of days before transitioning from Auto Restarts scheduled outside of active hours to Engaged Restart, which requires the user to schedule, with valid range from 0 to 30 days
    engagedRestartTransitionScheduleInDays *int32
    // Defer Feature Updates by these many days
    featureUpdatesDeferralPeriodInDays *int32
    // Pause Feature Updates
    featureUpdatesPaused *bool
    // Feature Updates Pause Expiry datetime
    featureUpdatesPauseExpiryDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Feature Updates Pause start date. This property is read-only.
    featureUpdatesPauseStartDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Feature Updates Rollback Start datetime
    featureUpdatesRollbackStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of days after a Feature Update for which a rollback is valid
    featureUpdatesRollbackWindowInDays *int32
    // Specifies whether to rollback Feature Updates on the next device check in
    featureUpdatesWillBeRolledBack *bool
    // Installation schedule
    installationSchedule WindowsUpdateInstallScheduleTypeable
    // Allow Microsoft Update Service
    microsoftUpdateServiceAllowed *bool
    // Specifies if the device should wait until deadline for rebooting outside of active hours
    postponeRebootUntilAfterDeadline *bool
    // Possible values for pre-release features.
    prereleaseFeatures *PrereleaseFeatures
    // Defer Quality Updates by these many days
    qualityUpdatesDeferralPeriodInDays *int32
    // Pause Quality Updates
    qualityUpdatesPaused *bool
    // Quality Updates Pause Expiry datetime
    qualityUpdatesPauseExpiryDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Quality Updates Pause start date. This property is read-only.
    qualityUpdatesPauseStartDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Quality Updates Rollback Start datetime
    qualityUpdatesRollbackStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies whether to rollback Quality Updates on the next device check in
    qualityUpdatesWillBeRolledBack *bool
    // Specify the period for auto-restart imminent warning notifications. Supported values: 15, 30 or 60 (minutes).
    scheduleImminentRestartWarningInMinutes *int32
    // Specify the period for auto-restart warning reminder notifications. Supported values: 2, 4, 8, 12 or 24 (hours).
    scheduleRestartWarningInHours *int32
    // Set to skip all check before restart: Battery level = 40%, User presence, Display Needed, Presentation mode, Full screen mode, phone call state, game mode etc.
    skipChecksBeforeRestart *bool
    // Windows Update Notification Display Options
    updateNotificationLevel *WindowsUpdateNotificationDisplayOption
    // Scheduled the update installation on the weeks of the month. Possible values are: userDefined, firstWeek, secondWeek, thirdWeek, fourthWeek, everyWeek.
    updateWeeks *WindowsUpdateForBusinessUpdateWeeks
    // Possible values of a property
    userPauseAccess *Enablement
    // Possible values of a property
    userWindowsUpdateScanAccess *Enablement
}
// NewWindowsUpdateForBusinessConfiguration instantiates a new WindowsUpdateForBusinessConfiguration and sets the default values.
func NewWindowsUpdateForBusinessConfiguration()(*WindowsUpdateForBusinessConfiguration) {
    m := &WindowsUpdateForBusinessConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdateForBusinessConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsUpdateForBusinessConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdateForBusinessConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdateForBusinessConfiguration(), nil
}
// GetAllowWindows11Upgrade gets the allowWindows11Upgrade property value. Allow eligible Windows 10 devices to upgrade to the latest version of Windows 11.
func (m *WindowsUpdateForBusinessConfiguration) GetAllowWindows11Upgrade()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowWindows11Upgrade
    }
}
// GetAutomaticUpdateMode gets the automaticUpdateMode property value. Possible values for automatic update mode.
func (m *WindowsUpdateForBusinessConfiguration) GetAutomaticUpdateMode()(*AutomaticUpdateMode) {
    if m == nil {
        return nil
    } else {
        return m.automaticUpdateMode
    }
}
// GetAutoRestartNotificationDismissal gets the autoRestartNotificationDismissal property value. Auto restart required notification dismissal method
func (m *WindowsUpdateForBusinessConfiguration) GetAutoRestartNotificationDismissal()(*AutoRestartNotificationDismissalMethod) {
    if m == nil {
        return nil
    } else {
        return m.autoRestartNotificationDismissal
    }
}
// GetBusinessReadyUpdatesOnly gets the businessReadyUpdatesOnly property value. Which branch devices will receive their updates from
func (m *WindowsUpdateForBusinessConfiguration) GetBusinessReadyUpdatesOnly()(*WindowsUpdateType) {
    if m == nil {
        return nil
    } else {
        return m.businessReadyUpdatesOnly
    }
}
// GetDeadlineForFeatureUpdatesInDays gets the deadlineForFeatureUpdatesInDays property value. Number of days before feature updates are installed automatically with valid range from 0 to 30 days
func (m *WindowsUpdateForBusinessConfiguration) GetDeadlineForFeatureUpdatesInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deadlineForFeatureUpdatesInDays
    }
}
// GetDeadlineForQualityUpdatesInDays gets the deadlineForQualityUpdatesInDays property value. Number of days before quality updates are installed automatically with valid range from 0 to 30 days
func (m *WindowsUpdateForBusinessConfiguration) GetDeadlineForQualityUpdatesInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deadlineForQualityUpdatesInDays
    }
}
// GetDeadlineGracePeriodInDays gets the deadlineGracePeriodInDays property value. Number of days after deadline  until restarts occur automatically with valid range from 0 to 7 days
func (m *WindowsUpdateForBusinessConfiguration) GetDeadlineGracePeriodInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deadlineGracePeriodInDays
    }
}
// GetDeliveryOptimizationMode gets the deliveryOptimizationMode property value. Delivery optimization mode for peer distribution
func (m *WindowsUpdateForBusinessConfiguration) GetDeliveryOptimizationMode()(*WindowsDeliveryOptimizationMode) {
    if m == nil {
        return nil
    } else {
        return m.deliveryOptimizationMode
    }
}
// GetDeviceUpdateStates gets the deviceUpdateStates property value. Windows update for business configuration device states. This collection can contain a maximum of 500 elements.
func (m *WindowsUpdateForBusinessConfiguration) GetDeviceUpdateStates()([]WindowsUpdateStateable) {
    if m == nil {
        return nil
    } else {
        return m.deviceUpdateStates
    }
}
// GetDriversExcluded gets the driversExcluded property value. Exclude Windows update Drivers
func (m *WindowsUpdateForBusinessConfiguration) GetDriversExcluded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.driversExcluded
    }
}
// GetEngagedRestartDeadlineInDays gets the engagedRestartDeadlineInDays property value. Deadline in days before automatically scheduling and executing a pending restart outside of active hours, with valid range from 2 to 30 days
func (m *WindowsUpdateForBusinessConfiguration) GetEngagedRestartDeadlineInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.engagedRestartDeadlineInDays
    }
}
// GetEngagedRestartSnoozeScheduleInDays gets the engagedRestartSnoozeScheduleInDays property value. Number of days a user can snooze Engaged Restart reminder notifications with valid range from 1 to 3 days
func (m *WindowsUpdateForBusinessConfiguration) GetEngagedRestartSnoozeScheduleInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.engagedRestartSnoozeScheduleInDays
    }
}
// GetEngagedRestartTransitionScheduleInDays gets the engagedRestartTransitionScheduleInDays property value. Number of days before transitioning from Auto Restarts scheduled outside of active hours to Engaged Restart, which requires the user to schedule, with valid range from 0 to 30 days
func (m *WindowsUpdateForBusinessConfiguration) GetEngagedRestartTransitionScheduleInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.engagedRestartTransitionScheduleInDays
    }
}
// GetFeatureUpdatesDeferralPeriodInDays gets the featureUpdatesDeferralPeriodInDays property value. Defer Feature Updates by these many days
func (m *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesDeferralPeriodInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.featureUpdatesDeferralPeriodInDays
    }
}
// GetFeatureUpdatesPaused gets the featureUpdatesPaused property value. Pause Feature Updates
func (m *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPaused()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.featureUpdatesPaused
    }
}
// GetFeatureUpdatesPauseExpiryDateTime gets the featureUpdatesPauseExpiryDateTime property value. Feature Updates Pause Expiry datetime
func (m *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPauseExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.featureUpdatesPauseExpiryDateTime
    }
}
// GetFeatureUpdatesPauseStartDate gets the featureUpdatesPauseStartDate property value. Feature Updates Pause start date. This property is read-only.
func (m *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesPauseStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.featureUpdatesPauseStartDate
    }
}
// GetFeatureUpdatesRollbackStartDateTime gets the featureUpdatesRollbackStartDateTime property value. Feature Updates Rollback Start datetime
func (m *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesRollbackStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.featureUpdatesRollbackStartDateTime
    }
}
// GetFeatureUpdatesRollbackWindowInDays gets the featureUpdatesRollbackWindowInDays property value. The number of days after a Feature Update for which a rollback is valid
func (m *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesRollbackWindowInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.featureUpdatesRollbackWindowInDays
    }
}
// GetFeatureUpdatesWillBeRolledBack gets the featureUpdatesWillBeRolledBack property value. Specifies whether to rollback Feature Updates on the next device check in
func (m *WindowsUpdateForBusinessConfiguration) GetFeatureUpdatesWillBeRolledBack()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.featureUpdatesWillBeRolledBack
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdateForBusinessConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["allowWindows11Upgrade"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowWindows11Upgrade(val)
        }
        return nil
    }
    res["automaticUpdateMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAutomaticUpdateMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticUpdateMode(val.(*AutomaticUpdateMode))
        }
        return nil
    }
    res["autoRestartNotificationDismissal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAutoRestartNotificationDismissalMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoRestartNotificationDismissal(val.(*AutoRestartNotificationDismissalMethod))
        }
        return nil
    }
    res["businessReadyUpdatesOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsUpdateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessReadyUpdatesOnly(val.(*WindowsUpdateType))
        }
        return nil
    }
    res["deadlineForFeatureUpdatesInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeadlineForFeatureUpdatesInDays(val)
        }
        return nil
    }
    res["deadlineForQualityUpdatesInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeadlineForQualityUpdatesInDays(val)
        }
        return nil
    }
    res["deadlineGracePeriodInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeadlineGracePeriodInDays(val)
        }
        return nil
    }
    res["deliveryOptimizationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeliveryOptimizationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryOptimizationMode(val.(*WindowsDeliveryOptimizationMode))
        }
        return nil
    }
    res["deviceUpdateStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsUpdateStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsUpdateStateable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsUpdateStateable)
            }
            m.SetDeviceUpdateStates(res)
        }
        return nil
    }
    res["driversExcluded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriversExcluded(val)
        }
        return nil
    }
    res["engagedRestartDeadlineInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEngagedRestartDeadlineInDays(val)
        }
        return nil
    }
    res["engagedRestartSnoozeScheduleInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEngagedRestartSnoozeScheduleInDays(val)
        }
        return nil
    }
    res["engagedRestartTransitionScheduleInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEngagedRestartTransitionScheduleInDays(val)
        }
        return nil
    }
    res["featureUpdatesDeferralPeriodInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdatesDeferralPeriodInDays(val)
        }
        return nil
    }
    res["featureUpdatesPaused"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdatesPaused(val)
        }
        return nil
    }
    res["featureUpdatesPauseExpiryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdatesPauseExpiryDateTime(val)
        }
        return nil
    }
    res["featureUpdatesPauseStartDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdatesPauseStartDate(val)
        }
        return nil
    }
    res["featureUpdatesRollbackStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdatesRollbackStartDateTime(val)
        }
        return nil
    }
    res["featureUpdatesRollbackWindowInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdatesRollbackWindowInDays(val)
        }
        return nil
    }
    res["featureUpdatesWillBeRolledBack"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdatesWillBeRolledBack(val)
        }
        return nil
    }
    res["installationSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsUpdateInstallScheduleTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallationSchedule(val.(WindowsUpdateInstallScheduleTypeable))
        }
        return nil
    }
    res["microsoftUpdateServiceAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftUpdateServiceAllowed(val)
        }
        return nil
    }
    res["postponeRebootUntilAfterDeadline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostponeRebootUntilAfterDeadline(val)
        }
        return nil
    }
    res["prereleaseFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrereleaseFeatures)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrereleaseFeatures(val.(*PrereleaseFeatures))
        }
        return nil
    }
    res["qualityUpdatesDeferralPeriodInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdatesDeferralPeriodInDays(val)
        }
        return nil
    }
    res["qualityUpdatesPaused"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdatesPaused(val)
        }
        return nil
    }
    res["qualityUpdatesPauseExpiryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdatesPauseExpiryDateTime(val)
        }
        return nil
    }
    res["qualityUpdatesPauseStartDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdatesPauseStartDate(val)
        }
        return nil
    }
    res["qualityUpdatesRollbackStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdatesRollbackStartDateTime(val)
        }
        return nil
    }
    res["qualityUpdatesWillBeRolledBack"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdatesWillBeRolledBack(val)
        }
        return nil
    }
    res["scheduleImminentRestartWarningInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleImminentRestartWarningInMinutes(val)
        }
        return nil
    }
    res["scheduleRestartWarningInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleRestartWarningInHours(val)
        }
        return nil
    }
    res["skipChecksBeforeRestart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkipChecksBeforeRestart(val)
        }
        return nil
    }
    res["updateNotificationLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsUpdateNotificationDisplayOption)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateNotificationLevel(val.(*WindowsUpdateNotificationDisplayOption))
        }
        return nil
    }
    res["updateWeeks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsUpdateForBusinessUpdateWeeks)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateWeeks(val.(*WindowsUpdateForBusinessUpdateWeeks))
        }
        return nil
    }
    res["userPauseAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPauseAccess(val.(*Enablement))
        }
        return nil
    }
    res["userWindowsUpdateScanAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserWindowsUpdateScanAccess(val.(*Enablement))
        }
        return nil
    }
    return res
}
// GetInstallationSchedule gets the installationSchedule property value. Installation schedule
func (m *WindowsUpdateForBusinessConfiguration) GetInstallationSchedule()(WindowsUpdateInstallScheduleTypeable) {
    if m == nil {
        return nil
    } else {
        return m.installationSchedule
    }
}
// GetMicrosoftUpdateServiceAllowed gets the microsoftUpdateServiceAllowed property value. Allow Microsoft Update Service
func (m *WindowsUpdateForBusinessConfiguration) GetMicrosoftUpdateServiceAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.microsoftUpdateServiceAllowed
    }
}
// GetPostponeRebootUntilAfterDeadline gets the postponeRebootUntilAfterDeadline property value. Specifies if the device should wait until deadline for rebooting outside of active hours
func (m *WindowsUpdateForBusinessConfiguration) GetPostponeRebootUntilAfterDeadline()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.postponeRebootUntilAfterDeadline
    }
}
// GetPrereleaseFeatures gets the prereleaseFeatures property value. Possible values for pre-release features.
func (m *WindowsUpdateForBusinessConfiguration) GetPrereleaseFeatures()(*PrereleaseFeatures) {
    if m == nil {
        return nil
    } else {
        return m.prereleaseFeatures
    }
}
// GetQualityUpdatesDeferralPeriodInDays gets the qualityUpdatesDeferralPeriodInDays property value. Defer Quality Updates by these many days
func (m *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesDeferralPeriodInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.qualityUpdatesDeferralPeriodInDays
    }
}
// GetQualityUpdatesPaused gets the qualityUpdatesPaused property value. Pause Quality Updates
func (m *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesPaused()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.qualityUpdatesPaused
    }
}
// GetQualityUpdatesPauseExpiryDateTime gets the qualityUpdatesPauseExpiryDateTime property value. Quality Updates Pause Expiry datetime
func (m *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesPauseExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.qualityUpdatesPauseExpiryDateTime
    }
}
// GetQualityUpdatesPauseStartDate gets the qualityUpdatesPauseStartDate property value. Quality Updates Pause start date. This property is read-only.
func (m *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesPauseStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.qualityUpdatesPauseStartDate
    }
}
// GetQualityUpdatesRollbackStartDateTime gets the qualityUpdatesRollbackStartDateTime property value. Quality Updates Rollback Start datetime
func (m *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesRollbackStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.qualityUpdatesRollbackStartDateTime
    }
}
// GetQualityUpdatesWillBeRolledBack gets the qualityUpdatesWillBeRolledBack property value. Specifies whether to rollback Quality Updates on the next device check in
func (m *WindowsUpdateForBusinessConfiguration) GetQualityUpdatesWillBeRolledBack()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.qualityUpdatesWillBeRolledBack
    }
}
// GetScheduleImminentRestartWarningInMinutes gets the scheduleImminentRestartWarningInMinutes property value. Specify the period for auto-restart imminent warning notifications. Supported values: 15, 30 or 60 (minutes).
func (m *WindowsUpdateForBusinessConfiguration) GetScheduleImminentRestartWarningInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.scheduleImminentRestartWarningInMinutes
    }
}
// GetScheduleRestartWarningInHours gets the scheduleRestartWarningInHours property value. Specify the period for auto-restart warning reminder notifications. Supported values: 2, 4, 8, 12 or 24 (hours).
func (m *WindowsUpdateForBusinessConfiguration) GetScheduleRestartWarningInHours()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.scheduleRestartWarningInHours
    }
}
// GetSkipChecksBeforeRestart gets the skipChecksBeforeRestart property value. Set to skip all check before restart: Battery level = 40%, User presence, Display Needed, Presentation mode, Full screen mode, phone call state, game mode etc.
func (m *WindowsUpdateForBusinessConfiguration) GetSkipChecksBeforeRestart()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.skipChecksBeforeRestart
    }
}
// GetUpdateNotificationLevel gets the updateNotificationLevel property value. Windows Update Notification Display Options
func (m *WindowsUpdateForBusinessConfiguration) GetUpdateNotificationLevel()(*WindowsUpdateNotificationDisplayOption) {
    if m == nil {
        return nil
    } else {
        return m.updateNotificationLevel
    }
}
// GetUpdateWeeks gets the updateWeeks property value. Scheduled the update installation on the weeks of the month. Possible values are: userDefined, firstWeek, secondWeek, thirdWeek, fourthWeek, everyWeek.
func (m *WindowsUpdateForBusinessConfiguration) GetUpdateWeeks()(*WindowsUpdateForBusinessUpdateWeeks) {
    if m == nil {
        return nil
    } else {
        return m.updateWeeks
    }
}
// GetUserPauseAccess gets the userPauseAccess property value. Possible values of a property
func (m *WindowsUpdateForBusinessConfiguration) GetUserPauseAccess()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.userPauseAccess
    }
}
// GetUserWindowsUpdateScanAccess gets the userWindowsUpdateScanAccess property value. Possible values of a property
func (m *WindowsUpdateForBusinessConfiguration) GetUserWindowsUpdateScanAccess()(*Enablement) {
    if m == nil {
        return nil
    } else {
        return m.userWindowsUpdateScanAccess
    }
}
// Serialize serializes information the current object
func (m *WindowsUpdateForBusinessConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowWindows11Upgrade", m.GetAllowWindows11Upgrade())
        if err != nil {
            return err
        }
    }
    if m.GetAutomaticUpdateMode() != nil {
        cast := (*m.GetAutomaticUpdateMode()).String()
        err = writer.WriteStringValue("automaticUpdateMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAutoRestartNotificationDismissal() != nil {
        cast := (*m.GetAutoRestartNotificationDismissal()).String()
        err = writer.WriteStringValue("autoRestartNotificationDismissal", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetBusinessReadyUpdatesOnly() != nil {
        cast := (*m.GetBusinessReadyUpdatesOnly()).String()
        err = writer.WriteStringValue("businessReadyUpdatesOnly", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deadlineForFeatureUpdatesInDays", m.GetDeadlineForFeatureUpdatesInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deadlineForQualityUpdatesInDays", m.GetDeadlineForQualityUpdatesInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deadlineGracePeriodInDays", m.GetDeadlineGracePeriodInDays())
        if err != nil {
            return err
        }
    }
    if m.GetDeliveryOptimizationMode() != nil {
        cast := (*m.GetDeliveryOptimizationMode()).String()
        err = writer.WriteStringValue("deliveryOptimizationMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceUpdateStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceUpdateStates()))
        for i, v := range m.GetDeviceUpdateStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceUpdateStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("driversExcluded", m.GetDriversExcluded())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("engagedRestartDeadlineInDays", m.GetEngagedRestartDeadlineInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("engagedRestartSnoozeScheduleInDays", m.GetEngagedRestartSnoozeScheduleInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("engagedRestartTransitionScheduleInDays", m.GetEngagedRestartTransitionScheduleInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("featureUpdatesDeferralPeriodInDays", m.GetFeatureUpdatesDeferralPeriodInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("featureUpdatesPaused", m.GetFeatureUpdatesPaused())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("featureUpdatesPauseExpiryDateTime", m.GetFeatureUpdatesPauseExpiryDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("featureUpdatesPauseStartDate", m.GetFeatureUpdatesPauseStartDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("featureUpdatesRollbackStartDateTime", m.GetFeatureUpdatesRollbackStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("featureUpdatesRollbackWindowInDays", m.GetFeatureUpdatesRollbackWindowInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("featureUpdatesWillBeRolledBack", m.GetFeatureUpdatesWillBeRolledBack())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("installationSchedule", m.GetInstallationSchedule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftUpdateServiceAllowed", m.GetMicrosoftUpdateServiceAllowed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("postponeRebootUntilAfterDeadline", m.GetPostponeRebootUntilAfterDeadline())
        if err != nil {
            return err
        }
    }
    if m.GetPrereleaseFeatures() != nil {
        cast := (*m.GetPrereleaseFeatures()).String()
        err = writer.WriteStringValue("prereleaseFeatures", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("qualityUpdatesDeferralPeriodInDays", m.GetQualityUpdatesDeferralPeriodInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("qualityUpdatesPaused", m.GetQualityUpdatesPaused())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("qualityUpdatesPauseExpiryDateTime", m.GetQualityUpdatesPauseExpiryDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("qualityUpdatesPauseStartDate", m.GetQualityUpdatesPauseStartDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("qualityUpdatesRollbackStartDateTime", m.GetQualityUpdatesRollbackStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("qualityUpdatesWillBeRolledBack", m.GetQualityUpdatesWillBeRolledBack())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("scheduleImminentRestartWarningInMinutes", m.GetScheduleImminentRestartWarningInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("scheduleRestartWarningInHours", m.GetScheduleRestartWarningInHours())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("skipChecksBeforeRestart", m.GetSkipChecksBeforeRestart())
        if err != nil {
            return err
        }
    }
    if m.GetUpdateNotificationLevel() != nil {
        cast := (*m.GetUpdateNotificationLevel()).String()
        err = writer.WriteStringValue("updateNotificationLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdateWeeks() != nil {
        cast := (*m.GetUpdateWeeks()).String()
        err = writer.WriteStringValue("updateWeeks", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserPauseAccess() != nil {
        cast := (*m.GetUserPauseAccess()).String()
        err = writer.WriteStringValue("userPauseAccess", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserWindowsUpdateScanAccess() != nil {
        cast := (*m.GetUserWindowsUpdateScanAccess()).String()
        err = writer.WriteStringValue("userWindowsUpdateScanAccess", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowWindows11Upgrade sets the allowWindows11Upgrade property value. Allow eligible Windows 10 devices to upgrade to the latest version of Windows 11.
func (m *WindowsUpdateForBusinessConfiguration) SetAllowWindows11Upgrade(value *bool)() {
    if m != nil {
        m.allowWindows11Upgrade = value
    }
}
// SetAutomaticUpdateMode sets the automaticUpdateMode property value. Possible values for automatic update mode.
func (m *WindowsUpdateForBusinessConfiguration) SetAutomaticUpdateMode(value *AutomaticUpdateMode)() {
    if m != nil {
        m.automaticUpdateMode = value
    }
}
// SetAutoRestartNotificationDismissal sets the autoRestartNotificationDismissal property value. Auto restart required notification dismissal method
func (m *WindowsUpdateForBusinessConfiguration) SetAutoRestartNotificationDismissal(value *AutoRestartNotificationDismissalMethod)() {
    if m != nil {
        m.autoRestartNotificationDismissal = value
    }
}
// SetBusinessReadyUpdatesOnly sets the businessReadyUpdatesOnly property value. Which branch devices will receive their updates from
func (m *WindowsUpdateForBusinessConfiguration) SetBusinessReadyUpdatesOnly(value *WindowsUpdateType)() {
    if m != nil {
        m.businessReadyUpdatesOnly = value
    }
}
// SetDeadlineForFeatureUpdatesInDays sets the deadlineForFeatureUpdatesInDays property value. Number of days before feature updates are installed automatically with valid range from 0 to 30 days
func (m *WindowsUpdateForBusinessConfiguration) SetDeadlineForFeatureUpdatesInDays(value *int32)() {
    if m != nil {
        m.deadlineForFeatureUpdatesInDays = value
    }
}
// SetDeadlineForQualityUpdatesInDays sets the deadlineForQualityUpdatesInDays property value. Number of days before quality updates are installed automatically with valid range from 0 to 30 days
func (m *WindowsUpdateForBusinessConfiguration) SetDeadlineForQualityUpdatesInDays(value *int32)() {
    if m != nil {
        m.deadlineForQualityUpdatesInDays = value
    }
}
// SetDeadlineGracePeriodInDays sets the deadlineGracePeriodInDays property value. Number of days after deadline  until restarts occur automatically with valid range from 0 to 7 days
func (m *WindowsUpdateForBusinessConfiguration) SetDeadlineGracePeriodInDays(value *int32)() {
    if m != nil {
        m.deadlineGracePeriodInDays = value
    }
}
// SetDeliveryOptimizationMode sets the deliveryOptimizationMode property value. Delivery optimization mode for peer distribution
func (m *WindowsUpdateForBusinessConfiguration) SetDeliveryOptimizationMode(value *WindowsDeliveryOptimizationMode)() {
    if m != nil {
        m.deliveryOptimizationMode = value
    }
}
// SetDeviceUpdateStates sets the deviceUpdateStates property value. Windows update for business configuration device states. This collection can contain a maximum of 500 elements.
func (m *WindowsUpdateForBusinessConfiguration) SetDeviceUpdateStates(value []WindowsUpdateStateable)() {
    if m != nil {
        m.deviceUpdateStates = value
    }
}
// SetDriversExcluded sets the driversExcluded property value. Exclude Windows update Drivers
func (m *WindowsUpdateForBusinessConfiguration) SetDriversExcluded(value *bool)() {
    if m != nil {
        m.driversExcluded = value
    }
}
// SetEngagedRestartDeadlineInDays sets the engagedRestartDeadlineInDays property value. Deadline in days before automatically scheduling and executing a pending restart outside of active hours, with valid range from 2 to 30 days
func (m *WindowsUpdateForBusinessConfiguration) SetEngagedRestartDeadlineInDays(value *int32)() {
    if m != nil {
        m.engagedRestartDeadlineInDays = value
    }
}
// SetEngagedRestartSnoozeScheduleInDays sets the engagedRestartSnoozeScheduleInDays property value. Number of days a user can snooze Engaged Restart reminder notifications with valid range from 1 to 3 days
func (m *WindowsUpdateForBusinessConfiguration) SetEngagedRestartSnoozeScheduleInDays(value *int32)() {
    if m != nil {
        m.engagedRestartSnoozeScheduleInDays = value
    }
}
// SetEngagedRestartTransitionScheduleInDays sets the engagedRestartTransitionScheduleInDays property value. Number of days before transitioning from Auto Restarts scheduled outside of active hours to Engaged Restart, which requires the user to schedule, with valid range from 0 to 30 days
func (m *WindowsUpdateForBusinessConfiguration) SetEngagedRestartTransitionScheduleInDays(value *int32)() {
    if m != nil {
        m.engagedRestartTransitionScheduleInDays = value
    }
}
// SetFeatureUpdatesDeferralPeriodInDays sets the featureUpdatesDeferralPeriodInDays property value. Defer Feature Updates by these many days
func (m *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesDeferralPeriodInDays(value *int32)() {
    if m != nil {
        m.featureUpdatesDeferralPeriodInDays = value
    }
}
// SetFeatureUpdatesPaused sets the featureUpdatesPaused property value. Pause Feature Updates
func (m *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesPaused(value *bool)() {
    if m != nil {
        m.featureUpdatesPaused = value
    }
}
// SetFeatureUpdatesPauseExpiryDateTime sets the featureUpdatesPauseExpiryDateTime property value. Feature Updates Pause Expiry datetime
func (m *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesPauseExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.featureUpdatesPauseExpiryDateTime = value
    }
}
// SetFeatureUpdatesPauseStartDate sets the featureUpdatesPauseStartDate property value. Feature Updates Pause start date. This property is read-only.
func (m *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesPauseStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.featureUpdatesPauseStartDate = value
    }
}
// SetFeatureUpdatesRollbackStartDateTime sets the featureUpdatesRollbackStartDateTime property value. Feature Updates Rollback Start datetime
func (m *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesRollbackStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.featureUpdatesRollbackStartDateTime = value
    }
}
// SetFeatureUpdatesRollbackWindowInDays sets the featureUpdatesRollbackWindowInDays property value. The number of days after a Feature Update for which a rollback is valid
func (m *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesRollbackWindowInDays(value *int32)() {
    if m != nil {
        m.featureUpdatesRollbackWindowInDays = value
    }
}
// SetFeatureUpdatesWillBeRolledBack sets the featureUpdatesWillBeRolledBack property value. Specifies whether to rollback Feature Updates on the next device check in
func (m *WindowsUpdateForBusinessConfiguration) SetFeatureUpdatesWillBeRolledBack(value *bool)() {
    if m != nil {
        m.featureUpdatesWillBeRolledBack = value
    }
}
// SetInstallationSchedule sets the installationSchedule property value. Installation schedule
func (m *WindowsUpdateForBusinessConfiguration) SetInstallationSchedule(value WindowsUpdateInstallScheduleTypeable)() {
    if m != nil {
        m.installationSchedule = value
    }
}
// SetMicrosoftUpdateServiceAllowed sets the microsoftUpdateServiceAllowed property value. Allow Microsoft Update Service
func (m *WindowsUpdateForBusinessConfiguration) SetMicrosoftUpdateServiceAllowed(value *bool)() {
    if m != nil {
        m.microsoftUpdateServiceAllowed = value
    }
}
// SetPostponeRebootUntilAfterDeadline sets the postponeRebootUntilAfterDeadline property value. Specifies if the device should wait until deadline for rebooting outside of active hours
func (m *WindowsUpdateForBusinessConfiguration) SetPostponeRebootUntilAfterDeadline(value *bool)() {
    if m != nil {
        m.postponeRebootUntilAfterDeadline = value
    }
}
// SetPrereleaseFeatures sets the prereleaseFeatures property value. Possible values for pre-release features.
func (m *WindowsUpdateForBusinessConfiguration) SetPrereleaseFeatures(value *PrereleaseFeatures)() {
    if m != nil {
        m.prereleaseFeatures = value
    }
}
// SetQualityUpdatesDeferralPeriodInDays sets the qualityUpdatesDeferralPeriodInDays property value. Defer Quality Updates by these many days
func (m *WindowsUpdateForBusinessConfiguration) SetQualityUpdatesDeferralPeriodInDays(value *int32)() {
    if m != nil {
        m.qualityUpdatesDeferralPeriodInDays = value
    }
}
// SetQualityUpdatesPaused sets the qualityUpdatesPaused property value. Pause Quality Updates
func (m *WindowsUpdateForBusinessConfiguration) SetQualityUpdatesPaused(value *bool)() {
    if m != nil {
        m.qualityUpdatesPaused = value
    }
}
// SetQualityUpdatesPauseExpiryDateTime sets the qualityUpdatesPauseExpiryDateTime property value. Quality Updates Pause Expiry datetime
func (m *WindowsUpdateForBusinessConfiguration) SetQualityUpdatesPauseExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.qualityUpdatesPauseExpiryDateTime = value
    }
}
// SetQualityUpdatesPauseStartDate sets the qualityUpdatesPauseStartDate property value. Quality Updates Pause start date. This property is read-only.
func (m *WindowsUpdateForBusinessConfiguration) SetQualityUpdatesPauseStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    if m != nil {
        m.qualityUpdatesPauseStartDate = value
    }
}
// SetQualityUpdatesRollbackStartDateTime sets the qualityUpdatesRollbackStartDateTime property value. Quality Updates Rollback Start datetime
func (m *WindowsUpdateForBusinessConfiguration) SetQualityUpdatesRollbackStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.qualityUpdatesRollbackStartDateTime = value
    }
}
// SetQualityUpdatesWillBeRolledBack sets the qualityUpdatesWillBeRolledBack property value. Specifies whether to rollback Quality Updates on the next device check in
func (m *WindowsUpdateForBusinessConfiguration) SetQualityUpdatesWillBeRolledBack(value *bool)() {
    if m != nil {
        m.qualityUpdatesWillBeRolledBack = value
    }
}
// SetScheduleImminentRestartWarningInMinutes sets the scheduleImminentRestartWarningInMinutes property value. Specify the period for auto-restart imminent warning notifications. Supported values: 15, 30 or 60 (minutes).
func (m *WindowsUpdateForBusinessConfiguration) SetScheduleImminentRestartWarningInMinutes(value *int32)() {
    if m != nil {
        m.scheduleImminentRestartWarningInMinutes = value
    }
}
// SetScheduleRestartWarningInHours sets the scheduleRestartWarningInHours property value. Specify the period for auto-restart warning reminder notifications. Supported values: 2, 4, 8, 12 or 24 (hours).
func (m *WindowsUpdateForBusinessConfiguration) SetScheduleRestartWarningInHours(value *int32)() {
    if m != nil {
        m.scheduleRestartWarningInHours = value
    }
}
// SetSkipChecksBeforeRestart sets the skipChecksBeforeRestart property value. Set to skip all check before restart: Battery level = 40%, User presence, Display Needed, Presentation mode, Full screen mode, phone call state, game mode etc.
func (m *WindowsUpdateForBusinessConfiguration) SetSkipChecksBeforeRestart(value *bool)() {
    if m != nil {
        m.skipChecksBeforeRestart = value
    }
}
// SetUpdateNotificationLevel sets the updateNotificationLevel property value. Windows Update Notification Display Options
func (m *WindowsUpdateForBusinessConfiguration) SetUpdateNotificationLevel(value *WindowsUpdateNotificationDisplayOption)() {
    if m != nil {
        m.updateNotificationLevel = value
    }
}
// SetUpdateWeeks sets the updateWeeks property value. Scheduled the update installation on the weeks of the month. Possible values are: userDefined, firstWeek, secondWeek, thirdWeek, fourthWeek, everyWeek.
func (m *WindowsUpdateForBusinessConfiguration) SetUpdateWeeks(value *WindowsUpdateForBusinessUpdateWeeks)() {
    if m != nil {
        m.updateWeeks = value
    }
}
// SetUserPauseAccess sets the userPauseAccess property value. Possible values of a property
func (m *WindowsUpdateForBusinessConfiguration) SetUserPauseAccess(value *Enablement)() {
    if m != nil {
        m.userPauseAccess = value
    }
}
// SetUserWindowsUpdateScanAccess sets the userWindowsUpdateScanAccess property value. Possible values of a property
func (m *WindowsUpdateForBusinessConfiguration) SetUserWindowsUpdateScanAccess(value *Enablement)() {
    if m != nil {
        m.userWindowsUpdateScanAccess = value
    }
}
