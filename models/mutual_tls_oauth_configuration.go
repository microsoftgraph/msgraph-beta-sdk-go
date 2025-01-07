package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MutualTlsOauthConfiguration struct {
    TrustedCertificateAuthorityBase
}
// NewMutualTlsOauthConfiguration instantiates a new MutualTlsOauthConfiguration and sets the default values.
func NewMutualTlsOauthConfiguration()(*MutualTlsOauthConfiguration) {
    m := &MutualTlsOauthConfiguration{
        TrustedCertificateAuthorityBase: *NewTrustedCertificateAuthorityBase(),
    }
    odataTypeValue := "#microsoft.graph.mutualTlsOauthConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMutualTlsOauthConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMutualTlsOauthConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMutualTlsOauthConfiguration(), nil
}
// GetDisplayName gets the displayName property value. Friendly name. Supports $filter (eq, in).
// returns a *string when successful
func (m *MutualTlsOauthConfiguration) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MutualTlsOauthConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TrustedCertificateAuthorityBase.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["tlsClientAuthParameter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTlsClientRegistrationMetadata)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTlsClientAuthParameter(val.(*TlsClientRegistrationMetadata))
        }
        return nil
    }
    return res
}
// GetTlsClientAuthParameter gets the tlsClientAuthParameter property value. The tlsClientAuthParameter property
// returns a *TlsClientRegistrationMetadata when successful
func (m *MutualTlsOauthConfiguration) GetTlsClientAuthParameter()(*TlsClientRegistrationMetadata) {
    val, err := m.GetBackingStore().Get("tlsClientAuthParameter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TlsClientRegistrationMetadata)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MutualTlsOauthConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TrustedCertificateAuthorityBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetTlsClientAuthParameter() != nil {
        cast := (*m.GetTlsClientAuthParameter()).String()
        err = writer.WriteStringValue("tlsClientAuthParameter", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Friendly name. Supports $filter (eq, in).
func (m *MutualTlsOauthConfiguration) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTlsClientAuthParameter sets the tlsClientAuthParameter property value. The tlsClientAuthParameter property
func (m *MutualTlsOauthConfiguration) SetTlsClientAuthParameter(value *TlsClientRegistrationMetadata)() {
    err := m.GetBackingStore().Set("tlsClientAuthParameter", value)
    if err != nil {
        panic(err)
    }
}
type MutualTlsOauthConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TrustedCertificateAuthorityBaseable
    GetDisplayName()(*string)
    GetTlsClientAuthParameter()(*TlsClientRegistrationMetadata)
    SetDisplayName(value *string)()
    SetTlsClientAuthParameter(value *TlsClientRegistrationMetadata)()
}
