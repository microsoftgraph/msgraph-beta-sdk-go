package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SmsAuthenticationMethodConfiguration 
type SmsAuthenticationMethodConfiguration struct {
    AuthenticationMethodConfiguration
    // A collection of users or groups who are enabled to use the authentication method.
    includeTargets []SmsAuthenticationMethodTargetable
}
// NewSmsAuthenticationMethodConfiguration instantiates a new SmsAuthenticationMethodConfiguration and sets the default values.
func NewSmsAuthenticationMethodConfiguration()(*SmsAuthenticationMethodConfiguration) {
    m := &SmsAuthenticationMethodConfiguration{
        AuthenticationMethodConfiguration: *NewAuthenticationMethodConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.smsAuthenticationMethodConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSmsAuthenticationMethodConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSmsAuthenticationMethodConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSmsAuthenticationMethodConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SmsAuthenticationMethodConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethodConfiguration.GetFieldDeserializers()
    res["includeTargets"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSmsAuthenticationMethodTargetFromDiscriminatorValue , m.SetIncludeTargets)
    return res
}
// GetIncludeTargets gets the includeTargets property value. A collection of users or groups who are enabled to use the authentication method.
func (m *SmsAuthenticationMethodConfiguration) GetIncludeTargets()([]SmsAuthenticationMethodTargetable) {
    return m.includeTargets
}
// Serialize serializes information the current object
func (m *SmsAuthenticationMethodConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethodConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIncludeTargets() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetIncludeTargets())
        err = writer.WriteCollectionOfObjectValues("includeTargets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIncludeTargets sets the includeTargets property value. A collection of users or groups who are enabled to use the authentication method.
func (m *SmsAuthenticationMethodConfiguration) SetIncludeTargets(value []SmsAuthenticationMethodTargetable)() {
    m.includeTargets = value
}
