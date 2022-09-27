package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskConfiguration 
type WindowsKioskConfiguration struct {
    DeviceConfiguration
    // Enable public browsing kiosk mode for the Microsoft Edge browser. The Default is false.
    edgeKioskEnablePublicBrowsing *bool
    // Specify URLs that the kiosk browser is allowed to navigate to
    kioskBrowserBlockedUrlExceptions []string
    // Specify URLs that the kiosk browsers should not navigate to
    kioskBrowserBlockedURLs []string
    // Specify the default URL the browser should navigate to on launch.
    kioskBrowserDefaultUrl *string
    // Enable the kiosk browser's end session button. By default, the end session button is disabled.
    kioskBrowserEnableEndSessionButton *bool
    // Enable the kiosk browser's home button. By default, the home button is disabled.
    kioskBrowserEnableHomeButton *bool
    // Enable the kiosk browser's navigation buttons(forward/back). By default, the navigation buttons are disabled.
    kioskBrowserEnableNavigationButtons *bool
    // Specify the number of minutes the session is idle until the kiosk browser restarts in a fresh state.  Valid values are 1-1440. Valid values 1 to 1440
    kioskBrowserRestartOnIdleTimeInMinutes *int32
    // This policy setting allows to define a list of Kiosk profiles for a Kiosk configuration. This collection can contain a maximum of 3 elements.
    kioskProfiles []WindowsKioskProfileable
    // force update schedule for Kiosk devices.
    windowsKioskForceUpdateSchedule WindowsKioskForceUpdateScheduleable
}
// NewWindowsKioskConfiguration instantiates a new WindowsKioskConfiguration and sets the default values.
func NewWindowsKioskConfiguration()(*WindowsKioskConfiguration) {
    m := &WindowsKioskConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsKioskConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskConfiguration(), nil
}
// GetEdgeKioskEnablePublicBrowsing gets the edgeKioskEnablePublicBrowsing property value. Enable public browsing kiosk mode for the Microsoft Edge browser. The Default is false.
func (m *WindowsKioskConfiguration) GetEdgeKioskEnablePublicBrowsing()(*bool) {
    return m.edgeKioskEnablePublicBrowsing
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["edgeKioskEnablePublicBrowsing"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEdgeKioskEnablePublicBrowsing)
    res["kioskBrowserBlockedUrlExceptions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetKioskBrowserBlockedUrlExceptions)
    res["kioskBrowserBlockedURLs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetKioskBrowserBlockedURLs)
    res["kioskBrowserDefaultUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetKioskBrowserDefaultUrl)
    res["kioskBrowserEnableEndSessionButton"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetKioskBrowserEnableEndSessionButton)
    res["kioskBrowserEnableHomeButton"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetKioskBrowserEnableHomeButton)
    res["kioskBrowserEnableNavigationButtons"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetKioskBrowserEnableNavigationButtons)
    res["kioskBrowserRestartOnIdleTimeInMinutes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetKioskBrowserRestartOnIdleTimeInMinutes)
    res["kioskProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsKioskProfileFromDiscriminatorValue , m.SetKioskProfiles)
    res["windowsKioskForceUpdateSchedule"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindowsKioskForceUpdateScheduleFromDiscriminatorValue , m.SetWindowsKioskForceUpdateSchedule)
    return res
}
// GetKioskBrowserBlockedUrlExceptions gets the kioskBrowserBlockedUrlExceptions property value. Specify URLs that the kiosk browser is allowed to navigate to
func (m *WindowsKioskConfiguration) GetKioskBrowserBlockedUrlExceptions()([]string) {
    return m.kioskBrowserBlockedUrlExceptions
}
// GetKioskBrowserBlockedURLs gets the kioskBrowserBlockedURLs property value. Specify URLs that the kiosk browsers should not navigate to
func (m *WindowsKioskConfiguration) GetKioskBrowserBlockedURLs()([]string) {
    return m.kioskBrowserBlockedURLs
}
// GetKioskBrowserDefaultUrl gets the kioskBrowserDefaultUrl property value. Specify the default URL the browser should navigate to on launch.
func (m *WindowsKioskConfiguration) GetKioskBrowserDefaultUrl()(*string) {
    return m.kioskBrowserDefaultUrl
}
// GetKioskBrowserEnableEndSessionButton gets the kioskBrowserEnableEndSessionButton property value. Enable the kiosk browser's end session button. By default, the end session button is disabled.
func (m *WindowsKioskConfiguration) GetKioskBrowserEnableEndSessionButton()(*bool) {
    return m.kioskBrowserEnableEndSessionButton
}
// GetKioskBrowserEnableHomeButton gets the kioskBrowserEnableHomeButton property value. Enable the kiosk browser's home button. By default, the home button is disabled.
func (m *WindowsKioskConfiguration) GetKioskBrowserEnableHomeButton()(*bool) {
    return m.kioskBrowserEnableHomeButton
}
// GetKioskBrowserEnableNavigationButtons gets the kioskBrowserEnableNavigationButtons property value. Enable the kiosk browser's navigation buttons(forward/back). By default, the navigation buttons are disabled.
func (m *WindowsKioskConfiguration) GetKioskBrowserEnableNavigationButtons()(*bool) {
    return m.kioskBrowserEnableNavigationButtons
}
// GetKioskBrowserRestartOnIdleTimeInMinutes gets the kioskBrowserRestartOnIdleTimeInMinutes property value. Specify the number of minutes the session is idle until the kiosk browser restarts in a fresh state.  Valid values are 1-1440. Valid values 1 to 1440
func (m *WindowsKioskConfiguration) GetKioskBrowserRestartOnIdleTimeInMinutes()(*int32) {
    return m.kioskBrowserRestartOnIdleTimeInMinutes
}
// GetKioskProfiles gets the kioskProfiles property value. This policy setting allows to define a list of Kiosk profiles for a Kiosk configuration. This collection can contain a maximum of 3 elements.
func (m *WindowsKioskConfiguration) GetKioskProfiles()([]WindowsKioskProfileable) {
    return m.kioskProfiles
}
// GetWindowsKioskForceUpdateSchedule gets the windowsKioskForceUpdateSchedule property value. force update schedule for Kiosk devices.
func (m *WindowsKioskConfiguration) GetWindowsKioskForceUpdateSchedule()(WindowsKioskForceUpdateScheduleable) {
    return m.windowsKioskForceUpdateSchedule
}
// Serialize serializes information the current object
func (m *WindowsKioskConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("edgeKioskEnablePublicBrowsing", m.GetEdgeKioskEnablePublicBrowsing())
        if err != nil {
            return err
        }
    }
    if m.GetKioskBrowserBlockedUrlExceptions() != nil {
        err = writer.WriteCollectionOfStringValues("kioskBrowserBlockedUrlExceptions", m.GetKioskBrowserBlockedUrlExceptions())
        if err != nil {
            return err
        }
    }
    if m.GetKioskBrowserBlockedURLs() != nil {
        err = writer.WriteCollectionOfStringValues("kioskBrowserBlockedURLs", m.GetKioskBrowserBlockedURLs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskBrowserDefaultUrl", m.GetKioskBrowserDefaultUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskBrowserEnableEndSessionButton", m.GetKioskBrowserEnableEndSessionButton())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskBrowserEnableHomeButton", m.GetKioskBrowserEnableHomeButton())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskBrowserEnableNavigationButtons", m.GetKioskBrowserEnableNavigationButtons())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskBrowserRestartOnIdleTimeInMinutes", m.GetKioskBrowserRestartOnIdleTimeInMinutes())
        if err != nil {
            return err
        }
    }
    if m.GetKioskProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetKioskProfiles())
        err = writer.WriteCollectionOfObjectValues("kioskProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsKioskForceUpdateSchedule", m.GetWindowsKioskForceUpdateSchedule())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEdgeKioskEnablePublicBrowsing sets the edgeKioskEnablePublicBrowsing property value. Enable public browsing kiosk mode for the Microsoft Edge browser. The Default is false.
func (m *WindowsKioskConfiguration) SetEdgeKioskEnablePublicBrowsing(value *bool)() {
    m.edgeKioskEnablePublicBrowsing = value
}
// SetKioskBrowserBlockedUrlExceptions sets the kioskBrowserBlockedUrlExceptions property value. Specify URLs that the kiosk browser is allowed to navigate to
func (m *WindowsKioskConfiguration) SetKioskBrowserBlockedUrlExceptions(value []string)() {
    m.kioskBrowserBlockedUrlExceptions = value
}
// SetKioskBrowserBlockedURLs sets the kioskBrowserBlockedURLs property value. Specify URLs that the kiosk browsers should not navigate to
func (m *WindowsKioskConfiguration) SetKioskBrowserBlockedURLs(value []string)() {
    m.kioskBrowserBlockedURLs = value
}
// SetKioskBrowserDefaultUrl sets the kioskBrowserDefaultUrl property value. Specify the default URL the browser should navigate to on launch.
func (m *WindowsKioskConfiguration) SetKioskBrowserDefaultUrl(value *string)() {
    m.kioskBrowserDefaultUrl = value
}
// SetKioskBrowserEnableEndSessionButton sets the kioskBrowserEnableEndSessionButton property value. Enable the kiosk browser's end session button. By default, the end session button is disabled.
func (m *WindowsKioskConfiguration) SetKioskBrowserEnableEndSessionButton(value *bool)() {
    m.kioskBrowserEnableEndSessionButton = value
}
// SetKioskBrowserEnableHomeButton sets the kioskBrowserEnableHomeButton property value. Enable the kiosk browser's home button. By default, the home button is disabled.
func (m *WindowsKioskConfiguration) SetKioskBrowserEnableHomeButton(value *bool)() {
    m.kioskBrowserEnableHomeButton = value
}
// SetKioskBrowserEnableNavigationButtons sets the kioskBrowserEnableNavigationButtons property value. Enable the kiosk browser's navigation buttons(forward/back). By default, the navigation buttons are disabled.
func (m *WindowsKioskConfiguration) SetKioskBrowserEnableNavigationButtons(value *bool)() {
    m.kioskBrowserEnableNavigationButtons = value
}
// SetKioskBrowserRestartOnIdleTimeInMinutes sets the kioskBrowserRestartOnIdleTimeInMinutes property value. Specify the number of minutes the session is idle until the kiosk browser restarts in a fresh state.  Valid values are 1-1440. Valid values 1 to 1440
func (m *WindowsKioskConfiguration) SetKioskBrowserRestartOnIdleTimeInMinutes(value *int32)() {
    m.kioskBrowserRestartOnIdleTimeInMinutes = value
}
// SetKioskProfiles sets the kioskProfiles property value. This policy setting allows to define a list of Kiosk profiles for a Kiosk configuration. This collection can contain a maximum of 3 elements.
func (m *WindowsKioskConfiguration) SetKioskProfiles(value []WindowsKioskProfileable)() {
    m.kioskProfiles = value
}
// SetWindowsKioskForceUpdateSchedule sets the windowsKioskForceUpdateSchedule property value. force update schedule for Kiosk devices.
func (m *WindowsKioskConfiguration) SetWindowsKioskForceUpdateSchedule(value WindowsKioskForceUpdateScheduleable)() {
    m.windowsKioskForceUpdateSchedule = value
}
