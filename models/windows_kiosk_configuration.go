package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskConfiguration 
type WindowsKioskConfiguration struct {
    DeviceConfiguration
}
// NewWindowsKioskConfiguration instantiates a new WindowsKioskConfiguration and sets the default values.
func NewWindowsKioskConfiguration()(*WindowsKioskConfiguration) {
    m := &WindowsKioskConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsKioskConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsKioskConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskConfiguration(), nil
}
// GetEdgeKioskEnablePublicBrowsing gets the edgeKioskEnablePublicBrowsing property value. Enable public browsing kiosk mode for the Microsoft Edge browser. The Default is false.
func (m *WindowsKioskConfiguration) GetEdgeKioskEnablePublicBrowsing()(*bool) {
    val, err := m.GetBackingStore().Get("edgeKioskEnablePublicBrowsing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["edgeKioskEnablePublicBrowsing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeKioskEnablePublicBrowsing(val)
        }
        return nil
    }
    res["kioskBrowserBlockedUrlExceptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetKioskBrowserBlockedUrlExceptions(res)
        }
        return nil
    }
    res["kioskBrowserBlockedURLs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetKioskBrowserBlockedURLs(res)
        }
        return nil
    }
    res["kioskBrowserDefaultUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskBrowserDefaultUrl(val)
        }
        return nil
    }
    res["kioskBrowserEnableEndSessionButton"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskBrowserEnableEndSessionButton(val)
        }
        return nil
    }
    res["kioskBrowserEnableHomeButton"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskBrowserEnableHomeButton(val)
        }
        return nil
    }
    res["kioskBrowserEnableNavigationButtons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskBrowserEnableNavigationButtons(val)
        }
        return nil
    }
    res["kioskBrowserRestartOnIdleTimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskBrowserRestartOnIdleTimeInMinutes(val)
        }
        return nil
    }
    res["kioskProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsKioskProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsKioskProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsKioskProfileable)
                }
            }
            m.SetKioskProfiles(res)
        }
        return nil
    }
    res["windowsKioskForceUpdateSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsKioskForceUpdateScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsKioskForceUpdateSchedule(val.(WindowsKioskForceUpdateScheduleable))
        }
        return nil
    }
    return res
}
// GetKioskBrowserBlockedUrlExceptions gets the kioskBrowserBlockedUrlExceptions property value. Specify URLs that the kiosk browser is allowed to navigate to
func (m *WindowsKioskConfiguration) GetKioskBrowserBlockedUrlExceptions()([]string) {
    val, err := m.GetBackingStore().Get("kioskBrowserBlockedUrlExceptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetKioskBrowserBlockedURLs gets the kioskBrowserBlockedURLs property value. Specify URLs that the kiosk browsers should not navigate to
func (m *WindowsKioskConfiguration) GetKioskBrowserBlockedURLs()([]string) {
    val, err := m.GetBackingStore().Get("kioskBrowserBlockedURLs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetKioskBrowserDefaultUrl gets the kioskBrowserDefaultUrl property value. Specify the default URL the browser should navigate to on launch.
func (m *WindowsKioskConfiguration) GetKioskBrowserDefaultUrl()(*string) {
    val, err := m.GetBackingStore().Get("kioskBrowserDefaultUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKioskBrowserEnableEndSessionButton gets the kioskBrowserEnableEndSessionButton property value. Enable the kiosk browser's end session button. By default, the end session button is disabled.
func (m *WindowsKioskConfiguration) GetKioskBrowserEnableEndSessionButton()(*bool) {
    val, err := m.GetBackingStore().Get("kioskBrowserEnableEndSessionButton")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskBrowserEnableHomeButton gets the kioskBrowserEnableHomeButton property value. Enable the kiosk browser's home button. By default, the home button is disabled.
func (m *WindowsKioskConfiguration) GetKioskBrowserEnableHomeButton()(*bool) {
    val, err := m.GetBackingStore().Get("kioskBrowserEnableHomeButton")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskBrowserEnableNavigationButtons gets the kioskBrowserEnableNavigationButtons property value. Enable the kiosk browser's navigation buttons(forward/back). By default, the navigation buttons are disabled.
func (m *WindowsKioskConfiguration) GetKioskBrowserEnableNavigationButtons()(*bool) {
    val, err := m.GetBackingStore().Get("kioskBrowserEnableNavigationButtons")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKioskBrowserRestartOnIdleTimeInMinutes gets the kioskBrowserRestartOnIdleTimeInMinutes property value. Specify the number of minutes the session is idle until the kiosk browser restarts in a fresh state.  Valid values are 1-1440. Valid values 1 to 1440
func (m *WindowsKioskConfiguration) GetKioskBrowserRestartOnIdleTimeInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("kioskBrowserRestartOnIdleTimeInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetKioskProfiles gets the kioskProfiles property value. This policy setting allows to define a list of Kiosk profiles for a Kiosk configuration. This collection can contain a maximum of 3 elements.
func (m *WindowsKioskConfiguration) GetKioskProfiles()([]WindowsKioskProfileable) {
    val, err := m.GetBackingStore().Get("kioskProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsKioskProfileable)
    }
    return nil
}
// GetWindowsKioskForceUpdateSchedule gets the windowsKioskForceUpdateSchedule property value. force update schedule for Kiosk devices.
func (m *WindowsKioskConfiguration) GetWindowsKioskForceUpdateSchedule()(WindowsKioskForceUpdateScheduleable) {
    val, err := m.GetBackingStore().Get("windowsKioskForceUpdateSchedule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsKioskForceUpdateScheduleable)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKioskProfiles()))
        for i, v := range m.GetKioskProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    err := m.GetBackingStore().Set("edgeKioskEnablePublicBrowsing", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskBrowserBlockedUrlExceptions sets the kioskBrowserBlockedUrlExceptions property value. Specify URLs that the kiosk browser is allowed to navigate to
func (m *WindowsKioskConfiguration) SetKioskBrowserBlockedUrlExceptions(value []string)() {
    err := m.GetBackingStore().Set("kioskBrowserBlockedUrlExceptions", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskBrowserBlockedURLs sets the kioskBrowserBlockedURLs property value. Specify URLs that the kiosk browsers should not navigate to
func (m *WindowsKioskConfiguration) SetKioskBrowserBlockedURLs(value []string)() {
    err := m.GetBackingStore().Set("kioskBrowserBlockedURLs", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskBrowserDefaultUrl sets the kioskBrowserDefaultUrl property value. Specify the default URL the browser should navigate to on launch.
func (m *WindowsKioskConfiguration) SetKioskBrowserDefaultUrl(value *string)() {
    err := m.GetBackingStore().Set("kioskBrowserDefaultUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskBrowserEnableEndSessionButton sets the kioskBrowserEnableEndSessionButton property value. Enable the kiosk browser's end session button. By default, the end session button is disabled.
func (m *WindowsKioskConfiguration) SetKioskBrowserEnableEndSessionButton(value *bool)() {
    err := m.GetBackingStore().Set("kioskBrowserEnableEndSessionButton", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskBrowserEnableHomeButton sets the kioskBrowserEnableHomeButton property value. Enable the kiosk browser's home button. By default, the home button is disabled.
func (m *WindowsKioskConfiguration) SetKioskBrowserEnableHomeButton(value *bool)() {
    err := m.GetBackingStore().Set("kioskBrowserEnableHomeButton", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskBrowserEnableNavigationButtons sets the kioskBrowserEnableNavigationButtons property value. Enable the kiosk browser's navigation buttons(forward/back). By default, the navigation buttons are disabled.
func (m *WindowsKioskConfiguration) SetKioskBrowserEnableNavigationButtons(value *bool)() {
    err := m.GetBackingStore().Set("kioskBrowserEnableNavigationButtons", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskBrowserRestartOnIdleTimeInMinutes sets the kioskBrowserRestartOnIdleTimeInMinutes property value. Specify the number of minutes the session is idle until the kiosk browser restarts in a fresh state.  Valid values are 1-1440. Valid values 1 to 1440
func (m *WindowsKioskConfiguration) SetKioskBrowserRestartOnIdleTimeInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("kioskBrowserRestartOnIdleTimeInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetKioskProfiles sets the kioskProfiles property value. This policy setting allows to define a list of Kiosk profiles for a Kiosk configuration. This collection can contain a maximum of 3 elements.
func (m *WindowsKioskConfiguration) SetKioskProfiles(value []WindowsKioskProfileable)() {
    err := m.GetBackingStore().Set("kioskProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsKioskForceUpdateSchedule sets the windowsKioskForceUpdateSchedule property value. force update schedule for Kiosk devices.
func (m *WindowsKioskConfiguration) SetWindowsKioskForceUpdateSchedule(value WindowsKioskForceUpdateScheduleable)() {
    err := m.GetBackingStore().Set("windowsKioskForceUpdateSchedule", value)
    if err != nil {
        panic(err)
    }
}
// WindowsKioskConfigurationable 
type WindowsKioskConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEdgeKioskEnablePublicBrowsing()(*bool)
    GetKioskBrowserBlockedUrlExceptions()([]string)
    GetKioskBrowserBlockedURLs()([]string)
    GetKioskBrowserDefaultUrl()(*string)
    GetKioskBrowserEnableEndSessionButton()(*bool)
    GetKioskBrowserEnableHomeButton()(*bool)
    GetKioskBrowserEnableNavigationButtons()(*bool)
    GetKioskBrowserRestartOnIdleTimeInMinutes()(*int32)
    GetKioskProfiles()([]WindowsKioskProfileable)
    GetWindowsKioskForceUpdateSchedule()(WindowsKioskForceUpdateScheduleable)
    SetEdgeKioskEnablePublicBrowsing(value *bool)()
    SetKioskBrowserBlockedUrlExceptions(value []string)()
    SetKioskBrowserBlockedURLs(value []string)()
    SetKioskBrowserDefaultUrl(value *string)()
    SetKioskBrowserEnableEndSessionButton(value *bool)()
    SetKioskBrowserEnableHomeButton(value *bool)()
    SetKioskBrowserEnableNavigationButtons(value *bool)()
    SetKioskBrowserRestartOnIdleTimeInMinutes(value *int32)()
    SetKioskProfiles(value []WindowsKioskProfileable)()
    SetWindowsKioskForceUpdateSchedule(value WindowsKioskForceUpdateScheduleable)()
}
