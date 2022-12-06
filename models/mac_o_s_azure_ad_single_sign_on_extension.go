package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSAzureAdSingleSignOnExtension 
type MacOSAzureAdSingleSignOnExtension struct {
    MacOSSingleSignOnExtension
    // An optional list of additional bundle IDs allowed to use the AAD extension for single sign-on.
    bundleIdAccessControlList []string
    // Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
    configurations []KeyTypedValuePairable
    // Enables or disables shared device mode.
    enableSharedDeviceMode *bool
}
// NewMacOSAzureAdSingleSignOnExtension instantiates a new MacOSAzureAdSingleSignOnExtension and sets the default values.
func NewMacOSAzureAdSingleSignOnExtension()(*MacOSAzureAdSingleSignOnExtension) {
    m := &MacOSAzureAdSingleSignOnExtension{
        MacOSSingleSignOnExtension: *NewMacOSSingleSignOnExtension(),
    }
    odataTypeValue := "#microsoft.graph.macOSAzureAdSingleSignOnExtension";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSAzureAdSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSAzureAdSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSAzureAdSingleSignOnExtension(), nil
}
// GetBundleIdAccessControlList gets the bundleIdAccessControlList property value. An optional list of additional bundle IDs allowed to use the AAD extension for single sign-on.
func (m *MacOSAzureAdSingleSignOnExtension) GetBundleIdAccessControlList()([]string) {
    return m.bundleIdAccessControlList
}
// GetConfigurations gets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *MacOSAzureAdSingleSignOnExtension) GetConfigurations()([]KeyTypedValuePairable) {
    return m.configurations
}
// GetEnableSharedDeviceMode gets the enableSharedDeviceMode property value. Enables or disables shared device mode.
func (m *MacOSAzureAdSingleSignOnExtension) GetEnableSharedDeviceMode()(*bool) {
    return m.enableSharedDeviceMode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSAzureAdSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MacOSSingleSignOnExtension.GetFieldDeserializers()
    res["bundleIdAccessControlList"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetBundleIdAccessControlList)
    res["configurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyTypedValuePairFromDiscriminatorValue , m.SetConfigurations)
    res["enableSharedDeviceMode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnableSharedDeviceMode)
    return res
}
// Serialize serializes information the current object
func (m *MacOSAzureAdSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MacOSSingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBundleIdAccessControlList() != nil {
        err = writer.WriteCollectionOfStringValues("bundleIdAccessControlList", m.GetBundleIdAccessControlList())
        if err != nil {
            return err
        }
    }
    if m.GetConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetConfigurations())
        err = writer.WriteCollectionOfObjectValues("configurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableSharedDeviceMode", m.GetEnableSharedDeviceMode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBundleIdAccessControlList sets the bundleIdAccessControlList property value. An optional list of additional bundle IDs allowed to use the AAD extension for single sign-on.
func (m *MacOSAzureAdSingleSignOnExtension) SetBundleIdAccessControlList(value []string)() {
    m.bundleIdAccessControlList = value
}
// SetConfigurations sets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *MacOSAzureAdSingleSignOnExtension) SetConfigurations(value []KeyTypedValuePairable)() {
    m.configurations = value
}
// SetEnableSharedDeviceMode sets the enableSharedDeviceMode property value. Enables or disables shared device mode.
func (m *MacOSAzureAdSingleSignOnExtension) SetEnableSharedDeviceMode(value *bool)() {
    m.enableSharedDeviceMode = value
}
