package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosRedirectSingleSignOnExtension 
type IosRedirectSingleSignOnExtension struct {
    IosSingleSignOnExtension
    // Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
    configurations []KeyTypedValuePairable
    // Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
    extensionIdentifier *string
    // Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
    teamIdentifier *string
    // One or more URL prefixes of identity providers on whose behalf the app extension performs single sign-on. URLs must begin with http:// or https://. All URL prefixes must be unique for all profiles.
    urlPrefixes []string
}
// NewIosRedirectSingleSignOnExtension instantiates a new IosRedirectSingleSignOnExtension and sets the default values.
func NewIosRedirectSingleSignOnExtension()(*IosRedirectSingleSignOnExtension) {
    m := &IosRedirectSingleSignOnExtension{
        IosSingleSignOnExtension: *NewIosSingleSignOnExtension(),
    }
    odataTypeValue := "#microsoft.graph.iosRedirectSingleSignOnExtension";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosRedirectSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosRedirectSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosRedirectSingleSignOnExtension(), nil
}
// GetConfigurations gets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *IosRedirectSingleSignOnExtension) GetConfigurations()([]KeyTypedValuePairable) {
    return m.configurations
}
// GetExtensionIdentifier gets the extensionIdentifier property value. Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
func (m *IosRedirectSingleSignOnExtension) GetExtensionIdentifier()(*string) {
    return m.extensionIdentifier
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosRedirectSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosSingleSignOnExtension.GetFieldDeserializers()
    res["configurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyTypedValuePairFromDiscriminatorValue , m.SetConfigurations)
    res["extensionIdentifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExtensionIdentifier)
    res["teamIdentifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTeamIdentifier)
    res["urlPrefixes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetUrlPrefixes)
    return res
}
// GetTeamIdentifier gets the teamIdentifier property value. Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
func (m *IosRedirectSingleSignOnExtension) GetTeamIdentifier()(*string) {
    return m.teamIdentifier
}
// GetUrlPrefixes gets the urlPrefixes property value. One or more URL prefixes of identity providers on whose behalf the app extension performs single sign-on. URLs must begin with http:// or https://. All URL prefixes must be unique for all profiles.
func (m *IosRedirectSingleSignOnExtension) GetUrlPrefixes()([]string) {
    return m.urlPrefixes
}
// Serialize serializes information the current object
func (m *IosRedirectSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosSingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetConfigurations())
        err = writer.WriteCollectionOfObjectValues("configurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("extensionIdentifier", m.GetExtensionIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamIdentifier", m.GetTeamIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetUrlPrefixes() != nil {
        err = writer.WriteCollectionOfStringValues("urlPrefixes", m.GetUrlPrefixes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurations sets the configurations property value. Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
func (m *IosRedirectSingleSignOnExtension) SetConfigurations(value []KeyTypedValuePairable)() {
    m.configurations = value
}
// SetExtensionIdentifier sets the extensionIdentifier property value. Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
func (m *IosRedirectSingleSignOnExtension) SetExtensionIdentifier(value *string)() {
    m.extensionIdentifier = value
}
// SetTeamIdentifier sets the teamIdentifier property value. Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
func (m *IosRedirectSingleSignOnExtension) SetTeamIdentifier(value *string)() {
    m.teamIdentifier = value
}
// SetUrlPrefixes sets the urlPrefixes property value. One or more URL prefixes of identity providers on whose behalf the app extension performs single sign-on. URLs must begin with http:// or https://. All URL prefixes must be unique for all profiles.
func (m *IosRedirectSingleSignOnExtension) SetUrlPrefixes(value []string)() {
    m.urlPrefixes = value
}
