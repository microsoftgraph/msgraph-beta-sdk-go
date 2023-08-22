package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSWebClip contains properties and inherited properties for macOS web apps.
type MacOSWebClip struct {
    MobileApp
}
// NewMacOSWebClip instantiates a new macOSWebClip and sets the default values.
func NewMacOSWebClip()(*MacOSWebClip) {
    m := &MacOSWebClip{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.macOSWebClip"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSWebClipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSWebClipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSWebClip(), nil
}
// GetAppUrl gets the appUrl property value. The web app URL starting with http:// or https://, such as https://learn.microsoft.com/mem/.
func (m *MacOSWebClip) GetAppUrl()(*string) {
    val, err := m.GetBackingStore().Get("appUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSWebClip) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppUrl(val)
        }
        return nil
    }
    res["fullScreenEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullScreenEnabled(val)
        }
        return nil
    }
    res["preComposedIconEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreComposedIconEnabled(val)
        }
        return nil
    }
    return res
}
// GetFullScreenEnabled gets the fullScreenEnabled property value. Whether or not to open the web clip as a full-screen web app. Defaults to false. If TRUE, opens the web clip as a full-screen web app. If FALSE, the web clip opens inside of another app.
func (m *MacOSWebClip) GetFullScreenEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("fullScreenEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPreComposedIconEnabled gets the preComposedIconEnabled property value. Whether or not the icon for the app is precomosed. Defaults to false. If TRUE, prevents SpringBoard from adding 'shine' to the icon. If FALSE, SpringBoard can add 'shine'.
func (m *MacOSWebClip) GetPreComposedIconEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("preComposedIconEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MacOSWebClip) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appUrl", m.GetAppUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fullScreenEnabled", m.GetFullScreenEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("preComposedIconEnabled", m.GetPreComposedIconEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppUrl sets the appUrl property value. The web app URL starting with http:// or https://, such as https://learn.microsoft.com/mem/.
func (m *MacOSWebClip) SetAppUrl(value *string)() {
    err := m.GetBackingStore().Set("appUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetFullScreenEnabled sets the fullScreenEnabled property value. Whether or not to open the web clip as a full-screen web app. Defaults to false. If TRUE, opens the web clip as a full-screen web app. If FALSE, the web clip opens inside of another app.
func (m *MacOSWebClip) SetFullScreenEnabled(value *bool)() {
    err := m.GetBackingStore().Set("fullScreenEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetPreComposedIconEnabled sets the preComposedIconEnabled property value. Whether or not the icon for the app is precomosed. Defaults to false. If TRUE, prevents SpringBoard from adding 'shine' to the icon. If FALSE, SpringBoard can add 'shine'.
func (m *MacOSWebClip) SetPreComposedIconEnabled(value *bool)() {
    err := m.GetBackingStore().Set("preComposedIconEnabled", value)
    if err != nil {
        panic(err)
    }
}
// MacOSWebClipable 
type MacOSWebClipable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppUrl()(*string)
    GetFullScreenEnabled()(*bool)
    GetPreComposedIconEnabled()(*bool)
    SetAppUrl(value *string)()
    SetFullScreenEnabled(value *bool)()
    SetPreComposedIconEnabled(value *bool)()
}
