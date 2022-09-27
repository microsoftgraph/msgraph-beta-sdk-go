package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosHomeScreenApp 
type IosHomeScreenApp struct {
    IosHomeScreenItem
    // BundleID of the app if isWebClip is false or the URL of a web clip if isWebClip is true.
    bundleID *string
    // When true, the bundle ID will be handled as a URL for a web clip.
    isWebClip *bool
}
// NewIosHomeScreenApp instantiates a new IosHomeScreenApp and sets the default values.
func NewIosHomeScreenApp()(*IosHomeScreenApp) {
    m := &IosHomeScreenApp{
        IosHomeScreenItem: *NewIosHomeScreenItem(),
    }
    odataTypeValue := "#microsoft.graph.iosHomeScreenApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosHomeScreenAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosHomeScreenAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosHomeScreenApp(), nil
}
// GetBundleID gets the bundleID property value. BundleID of the app if isWebClip is false or the URL of a web clip if isWebClip is true.
func (m *IosHomeScreenApp) GetBundleID()(*string) {
    return m.bundleID
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosHomeScreenApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosHomeScreenItem.GetFieldDeserializers()
    res["bundleID"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBundleID)
    res["isWebClip"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsWebClip)
    return res
}
// GetIsWebClip gets the isWebClip property value. When true, the bundle ID will be handled as a URL for a web clip.
func (m *IosHomeScreenApp) GetIsWebClip()(*bool) {
    return m.isWebClip
}
// Serialize serializes information the current object
func (m *IosHomeScreenApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosHomeScreenItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("bundleID", m.GetBundleID())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isWebClip", m.GetIsWebClip())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBundleID sets the bundleID property value. BundleID of the app if isWebClip is false or the URL of a web clip if isWebClip is true.
func (m *IosHomeScreenApp) SetBundleID(value *string)() {
    m.bundleID = value
}
// SetIsWebClip sets the isWebClip property value. When true, the bundle ID will be handled as a URL for a web clip.
func (m *IosHomeScreenApp) SetIsWebClip(value *bool)() {
    m.isWebClip = value
}
