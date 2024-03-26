package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EnrollmentCompletionPageConfiguration windows 10 Enrollment Status Page Configuration
type Windows10EnrollmentCompletionPageConfiguration struct {
    DeviceEnrollmentConfiguration
}
// NewWindows10EnrollmentCompletionPageConfiguration instantiates a new Windows10EnrollmentCompletionPageConfiguration and sets the default values.
func NewWindows10EnrollmentCompletionPageConfiguration()(*Windows10EnrollmentCompletionPageConfiguration) {
    m := &Windows10EnrollmentCompletionPageConfiguration{
        DeviceEnrollmentConfiguration: *NewDeviceEnrollmentConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10EnrollmentCompletionPageConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10EnrollmentCompletionPageConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindows10EnrollmentCompletionPageConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10EnrollmentCompletionPageConfiguration(), nil
}
// GetAllowDeviceResetOnInstallFailure gets the allowDeviceResetOnInstallFailure property value. When TRUE, allows device reset on installation failure. When false, reset is blocked. The default is false.
// returns a *bool when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetAllowDeviceResetOnInstallFailure()(*bool) {
    val, err := m.GetBackingStore().Get("allowDeviceResetOnInstallFailure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowDeviceUseOnInstallFailure gets the allowDeviceUseOnInstallFailure property value. When TRUE, allows the user to continue using the device on installation failure. When false, blocks the user on installation failure. The default is false.
// returns a *bool when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetAllowDeviceUseOnInstallFailure()(*bool) {
    val, err := m.GetBackingStore().Get("allowDeviceUseOnInstallFailure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowLogCollectionOnInstallFailure gets the allowLogCollectionOnInstallFailure property value. When TRUE, allows log collection on installation failure. When false, log collection is not allowed. The default is false.
// returns a *bool when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetAllowLogCollectionOnInstallFailure()(*bool) {
    val, err := m.GetBackingStore().Get("allowLogCollectionOnInstallFailure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowNonBlockingAppInstallation gets the allowNonBlockingAppInstallation property value. When TRUE, ESP (Enrollment Status Page) installs all required apps targeted during technician phase and ignores any failures for non-blocking apps. When FALSE, ESP fails on any error during app install. The default is false.
// returns a *bool when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetAllowNonBlockingAppInstallation()(*bool) {
    val, err := m.GetBackingStore().Get("allowNonBlockingAppInstallation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBlockDeviceSetupRetryByUser gets the blockDeviceSetupRetryByUser property value. When TRUE, blocks user from retrying the setup on installation failure. When false, user is allowed to retry. The default is false.
// returns a *bool when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetBlockDeviceSetupRetryByUser()(*bool) {
    val, err := m.GetBackingStore().Get("blockDeviceSetupRetryByUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCustomErrorMessage gets the customErrorMessage property value. The custom error message to show upon installation failure. Max length is 10000. example: 'Setup could not be completed. Please try again or contact your support person for help.'
// returns a *string when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetCustomErrorMessage()(*string) {
    val, err := m.GetBackingStore().Get("customErrorMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisableUserStatusTrackingAfterFirstUser gets the disableUserStatusTrackingAfterFirstUser property value. When TRUE, disables showing installation progress for first user post enrollment. When false, enables showing progress. The default is false.
// returns a *bool when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetDisableUserStatusTrackingAfterFirstUser()(*bool) {
    val, err := m.GetBackingStore().Get("disableUserStatusTrackingAfterFirstUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceEnrollmentConfiguration.GetFieldDeserializers()
    res["allowDeviceResetOnInstallFailure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeviceResetOnInstallFailure(val)
        }
        return nil
    }
    res["allowDeviceUseOnInstallFailure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDeviceUseOnInstallFailure(val)
        }
        return nil
    }
    res["allowLogCollectionOnInstallFailure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowLogCollectionOnInstallFailure(val)
        }
        return nil
    }
    res["allowNonBlockingAppInstallation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowNonBlockingAppInstallation(val)
        }
        return nil
    }
    res["blockDeviceSetupRetryByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockDeviceSetupRetryByUser(val)
        }
        return nil
    }
    res["customErrorMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomErrorMessage(val)
        }
        return nil
    }
    res["disableUserStatusTrackingAfterFirstUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableUserStatusTrackingAfterFirstUser(val)
        }
        return nil
    }
    res["installProgressTimeoutInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallProgressTimeoutInMinutes(val)
        }
        return nil
    }
    res["installQualityUpdates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallQualityUpdates(val)
        }
        return nil
    }
    res["selectedMobileAppIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSelectedMobileAppIds(res)
        }
        return nil
    }
    res["showInstallationProgress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowInstallationProgress(val)
        }
        return nil
    }
    res["trackInstallProgressForAutopilotOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrackInstallProgressForAutopilotOnly(val)
        }
        return nil
    }
    return res
}
// GetInstallProgressTimeoutInMinutes gets the installProgressTimeoutInMinutes property value. The installation progress timeout in minutes. Default is 60 minutes.
// returns a *int32 when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetInstallProgressTimeoutInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("installProgressTimeoutInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetInstallQualityUpdates gets the installQualityUpdates property value. Allows quality updates installation during OOBE
// returns a *bool when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetInstallQualityUpdates()(*bool) {
    val, err := m.GetBackingStore().Get("installQualityUpdates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSelectedMobileAppIds gets the selectedMobileAppIds property value. Selected applications to track the installation status. It is in the form of an array of GUIDs.
// returns a []string when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetSelectedMobileAppIds()([]string) {
    val, err := m.GetBackingStore().Get("selectedMobileAppIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetShowInstallationProgress gets the showInstallationProgress property value. When TRUE, shows installation progress to user. When false, hides installation progress. The default is false.
// returns a *bool when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetShowInstallationProgress()(*bool) {
    val, err := m.GetBackingStore().Get("showInstallationProgress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTrackInstallProgressForAutopilotOnly gets the trackInstallProgressForAutopilotOnly property value. When TRUE, installation progress is tracked for only Autopilot enrollment scenarios. When false, other scenarios are tracked as well. The default is false.
// returns a *bool when successful
func (m *Windows10EnrollmentCompletionPageConfiguration) GetTrackInstallProgressForAutopilotOnly()(*bool) {
    val, err := m.GetBackingStore().Get("trackInstallProgressForAutopilotOnly")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Windows10EnrollmentCompletionPageConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceEnrollmentConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowDeviceResetOnInstallFailure", m.GetAllowDeviceResetOnInstallFailure())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowDeviceUseOnInstallFailure", m.GetAllowDeviceUseOnInstallFailure())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowLogCollectionOnInstallFailure", m.GetAllowLogCollectionOnInstallFailure())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowNonBlockingAppInstallation", m.GetAllowNonBlockingAppInstallation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockDeviceSetupRetryByUser", m.GetBlockDeviceSetupRetryByUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customErrorMessage", m.GetCustomErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableUserStatusTrackingAfterFirstUser", m.GetDisableUserStatusTrackingAfterFirstUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("installProgressTimeoutInMinutes", m.GetInstallProgressTimeoutInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("installQualityUpdates", m.GetInstallQualityUpdates())
        if err != nil {
            return err
        }
    }
    if m.GetSelectedMobileAppIds() != nil {
        err = writer.WriteCollectionOfStringValues("selectedMobileAppIds", m.GetSelectedMobileAppIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showInstallationProgress", m.GetShowInstallationProgress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("trackInstallProgressForAutopilotOnly", m.GetTrackInstallProgressForAutopilotOnly())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowDeviceResetOnInstallFailure sets the allowDeviceResetOnInstallFailure property value. When TRUE, allows device reset on installation failure. When false, reset is blocked. The default is false.
func (m *Windows10EnrollmentCompletionPageConfiguration) SetAllowDeviceResetOnInstallFailure(value *bool)() {
    err := m.GetBackingStore().Set("allowDeviceResetOnInstallFailure", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowDeviceUseOnInstallFailure sets the allowDeviceUseOnInstallFailure property value. When TRUE, allows the user to continue using the device on installation failure. When false, blocks the user on installation failure. The default is false.
func (m *Windows10EnrollmentCompletionPageConfiguration) SetAllowDeviceUseOnInstallFailure(value *bool)() {
    err := m.GetBackingStore().Set("allowDeviceUseOnInstallFailure", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowLogCollectionOnInstallFailure sets the allowLogCollectionOnInstallFailure property value. When TRUE, allows log collection on installation failure. When false, log collection is not allowed. The default is false.
func (m *Windows10EnrollmentCompletionPageConfiguration) SetAllowLogCollectionOnInstallFailure(value *bool)() {
    err := m.GetBackingStore().Set("allowLogCollectionOnInstallFailure", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowNonBlockingAppInstallation sets the allowNonBlockingAppInstallation property value. When TRUE, ESP (Enrollment Status Page) installs all required apps targeted during technician phase and ignores any failures for non-blocking apps. When FALSE, ESP fails on any error during app install. The default is false.
func (m *Windows10EnrollmentCompletionPageConfiguration) SetAllowNonBlockingAppInstallation(value *bool)() {
    err := m.GetBackingStore().Set("allowNonBlockingAppInstallation", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockDeviceSetupRetryByUser sets the blockDeviceSetupRetryByUser property value. When TRUE, blocks user from retrying the setup on installation failure. When false, user is allowed to retry. The default is false.
func (m *Windows10EnrollmentCompletionPageConfiguration) SetBlockDeviceSetupRetryByUser(value *bool)() {
    err := m.GetBackingStore().Set("blockDeviceSetupRetryByUser", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomErrorMessage sets the customErrorMessage property value. The custom error message to show upon installation failure. Max length is 10000. example: 'Setup could not be completed. Please try again or contact your support person for help.'
func (m *Windows10EnrollmentCompletionPageConfiguration) SetCustomErrorMessage(value *string)() {
    err := m.GetBackingStore().Set("customErrorMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetDisableUserStatusTrackingAfterFirstUser sets the disableUserStatusTrackingAfterFirstUser property value. When TRUE, disables showing installation progress for first user post enrollment. When false, enables showing progress. The default is false.
func (m *Windows10EnrollmentCompletionPageConfiguration) SetDisableUserStatusTrackingAfterFirstUser(value *bool)() {
    err := m.GetBackingStore().Set("disableUserStatusTrackingAfterFirstUser", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallProgressTimeoutInMinutes sets the installProgressTimeoutInMinutes property value. The installation progress timeout in minutes. Default is 60 minutes.
func (m *Windows10EnrollmentCompletionPageConfiguration) SetInstallProgressTimeoutInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("installProgressTimeoutInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetInstallQualityUpdates sets the installQualityUpdates property value. Allows quality updates installation during OOBE
func (m *Windows10EnrollmentCompletionPageConfiguration) SetInstallQualityUpdates(value *bool)() {
    err := m.GetBackingStore().Set("installQualityUpdates", value)
    if err != nil {
        panic(err)
    }
}
// SetSelectedMobileAppIds sets the selectedMobileAppIds property value. Selected applications to track the installation status. It is in the form of an array of GUIDs.
func (m *Windows10EnrollmentCompletionPageConfiguration) SetSelectedMobileAppIds(value []string)() {
    err := m.GetBackingStore().Set("selectedMobileAppIds", value)
    if err != nil {
        panic(err)
    }
}
// SetShowInstallationProgress sets the showInstallationProgress property value. When TRUE, shows installation progress to user. When false, hides installation progress. The default is false.
func (m *Windows10EnrollmentCompletionPageConfiguration) SetShowInstallationProgress(value *bool)() {
    err := m.GetBackingStore().Set("showInstallationProgress", value)
    if err != nil {
        panic(err)
    }
}
// SetTrackInstallProgressForAutopilotOnly sets the trackInstallProgressForAutopilotOnly property value. When TRUE, installation progress is tracked for only Autopilot enrollment scenarios. When false, other scenarios are tracked as well. The default is false.
func (m *Windows10EnrollmentCompletionPageConfiguration) SetTrackInstallProgressForAutopilotOnly(value *bool)() {
    err := m.GetBackingStore().Set("trackInstallProgressForAutopilotOnly", value)
    if err != nil {
        panic(err)
    }
}
type Windows10EnrollmentCompletionPageConfigurationable interface {
    DeviceEnrollmentConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDeviceResetOnInstallFailure()(*bool)
    GetAllowDeviceUseOnInstallFailure()(*bool)
    GetAllowLogCollectionOnInstallFailure()(*bool)
    GetAllowNonBlockingAppInstallation()(*bool)
    GetBlockDeviceSetupRetryByUser()(*bool)
    GetCustomErrorMessage()(*string)
    GetDisableUserStatusTrackingAfterFirstUser()(*bool)
    GetInstallProgressTimeoutInMinutes()(*int32)
    GetInstallQualityUpdates()(*bool)
    GetSelectedMobileAppIds()([]string)
    GetShowInstallationProgress()(*bool)
    GetTrackInstallProgressForAutopilotOnly()(*bool)
    SetAllowDeviceResetOnInstallFailure(value *bool)()
    SetAllowDeviceUseOnInstallFailure(value *bool)()
    SetAllowLogCollectionOnInstallFailure(value *bool)()
    SetAllowNonBlockingAppInstallation(value *bool)()
    SetBlockDeviceSetupRetryByUser(value *bool)()
    SetCustomErrorMessage(value *string)()
    SetDisableUserStatusTrackingAfterFirstUser(value *bool)()
    SetInstallProgressTimeoutInMinutes(value *int32)()
    SetInstallQualityUpdates(value *bool)()
    SetSelectedMobileAppIds(value []string)()
    SetShowInstallationProgress(value *bool)()
    SetTrackInstallProgressForAutopilotOnly(value *bool)()
}
