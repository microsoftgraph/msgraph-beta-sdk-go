package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationOAuth2ClientCredentialsConnectionSettings 
type EducationSynchronizationOAuth2ClientCredentialsConnectionSettings struct {
    EducationSynchronizationConnectionSettings
    // The scope of the access request (see RFC6749).
    scope *string
    // The URL to get access tokens for the data provider.
    tokenUrl *string
}
// NewEducationSynchronizationOAuth2ClientCredentialsConnectionSettings instantiates a new EducationSynchronizationOAuth2ClientCredentialsConnectionSettings and sets the default values.
func NewEducationSynchronizationOAuth2ClientCredentialsConnectionSettings()(*EducationSynchronizationOAuth2ClientCredentialsConnectionSettings) {
    m := &EducationSynchronizationOAuth2ClientCredentialsConnectionSettings{
        EducationSynchronizationConnectionSettings: *NewEducationSynchronizationConnectionSettings(),
    }
    odataTypeValue := "#microsoft.graph.educationSynchronizationOAuth2ClientCredentialsConnectionSettings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEducationSynchronizationOAuth2ClientCredentialsConnectionSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationOAuth2ClientCredentialsConnectionSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationOAuth2ClientCredentialsConnectionSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationOAuth2ClientCredentialsConnectionSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationSynchronizationConnectionSettings.GetFieldDeserializers()
    res["scope"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetScope)
    res["tokenUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTokenUrl)
    return res
}
// GetScope gets the scope property value. The scope of the access request (see RFC6749).
func (m *EducationSynchronizationOAuth2ClientCredentialsConnectionSettings) GetScope()(*string) {
    return m.scope
}
// GetTokenUrl gets the tokenUrl property value. The URL to get access tokens for the data provider.
func (m *EducationSynchronizationOAuth2ClientCredentialsConnectionSettings) GetTokenUrl()(*string) {
    return m.tokenUrl
}
// Serialize serializes information the current object
func (m *EducationSynchronizationOAuth2ClientCredentialsConnectionSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationSynchronizationConnectionSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tokenUrl", m.GetTokenUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetScope sets the scope property value. The scope of the access request (see RFC6749).
func (m *EducationSynchronizationOAuth2ClientCredentialsConnectionSettings) SetScope(value *string)() {
    m.scope = value
}
// SetTokenUrl sets the tokenUrl property value. The URL to get access tokens for the data provider.
func (m *EducationSynchronizationOAuth2ClientCredentialsConnectionSettings) SetTokenUrl(value *string)() {
    m.tokenUrl = value
}
