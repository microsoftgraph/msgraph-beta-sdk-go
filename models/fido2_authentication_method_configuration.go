package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Fido2AuthenticationMethodConfiguration 
type Fido2AuthenticationMethodConfiguration struct {
    AuthenticationMethodConfiguration
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A collection of users or groups who are enabled to use the authentication method.
    includeTargets []AuthenticationMethodTargetable
    // Determines whether attestation must be enforced for FIDO2 security key registration.
    isAttestationEnforced *bool
    // Determines if users can register new FIDO2 security keys.
    isSelfServiceRegistrationAllowed *bool
    // Controls whether key restrictions are enforced on FIDO2 security keys, either allowing or disallowing certain key types as defined by Authenticator Attestation GUID (AAGUID), an identifier that indicates the type (e.g. make and model) of the authenticator.
    keyRestrictions Fido2KeyRestrictionsable
}
// NewFido2AuthenticationMethodConfiguration instantiates a new Fido2AuthenticationMethodConfiguration and sets the default values.
func NewFido2AuthenticationMethodConfiguration()(*Fido2AuthenticationMethodConfiguration) {
    m := &Fido2AuthenticationMethodConfiguration{
        AuthenticationMethodConfiguration: *NewAuthenticationMethodConfiguration(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFido2AuthenticationMethodConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFido2AuthenticationMethodConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFido2AuthenticationMethodConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Fido2AuthenticationMethodConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Fido2AuthenticationMethodConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethodConfiguration.GetFieldDeserializers()
    res["includeTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationMethodTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodTargetable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationMethodTargetable)
            }
            m.SetIncludeTargets(res)
        }
        return nil
    }
    res["isAttestationEnforced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAttestationEnforced(val)
        }
        return nil
    }
    res["isSelfServiceRegistrationAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSelfServiceRegistrationAllowed(val)
        }
        return nil
    }
    res["keyRestrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFido2KeyRestrictionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyRestrictions(val.(Fido2KeyRestrictionsable))
        }
        return nil
    }
    return res
}
// GetIncludeTargets gets the includeTargets property value. A collection of users or groups who are enabled to use the authentication method.
func (m *Fido2AuthenticationMethodConfiguration) GetIncludeTargets()([]AuthenticationMethodTargetable) {
    if m == nil {
        return nil
    } else {
        return m.includeTargets
    }
}
// GetIsAttestationEnforced gets the isAttestationEnforced property value. Determines whether attestation must be enforced for FIDO2 security key registration.
func (m *Fido2AuthenticationMethodConfiguration) GetIsAttestationEnforced()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAttestationEnforced
    }
}
// GetIsSelfServiceRegistrationAllowed gets the isSelfServiceRegistrationAllowed property value. Determines if users can register new FIDO2 security keys.
func (m *Fido2AuthenticationMethodConfiguration) GetIsSelfServiceRegistrationAllowed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSelfServiceRegistrationAllowed
    }
}
// GetKeyRestrictions gets the keyRestrictions property value. Controls whether key restrictions are enforced on FIDO2 security keys, either allowing or disallowing certain key types as defined by Authenticator Attestation GUID (AAGUID), an identifier that indicates the type (e.g. make and model) of the authenticator.
func (m *Fido2AuthenticationMethodConfiguration) GetKeyRestrictions()(Fido2KeyRestrictionsable) {
    if m == nil {
        return nil
    } else {
        return m.keyRestrictions
    }
}
// Serialize serializes information the current object
func (m *Fido2AuthenticationMethodConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethodConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIncludeTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludeTargets()))
        for i, v := range m.GetIncludeTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("includeTargets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAttestationEnforced", m.GetIsAttestationEnforced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSelfServiceRegistrationAllowed", m.GetIsSelfServiceRegistrationAllowed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("keyRestrictions", m.GetKeyRestrictions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Fido2AuthenticationMethodConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIncludeTargets sets the includeTargets property value. A collection of users or groups who are enabled to use the authentication method.
func (m *Fido2AuthenticationMethodConfiguration) SetIncludeTargets(value []AuthenticationMethodTargetable)() {
    if m != nil {
        m.includeTargets = value
    }
}
// SetIsAttestationEnforced sets the isAttestationEnforced property value. Determines whether attestation must be enforced for FIDO2 security key registration.
func (m *Fido2AuthenticationMethodConfiguration) SetIsAttestationEnforced(value *bool)() {
    if m != nil {
        m.isAttestationEnforced = value
    }
}
// SetIsSelfServiceRegistrationAllowed sets the isSelfServiceRegistrationAllowed property value. Determines if users can register new FIDO2 security keys.
func (m *Fido2AuthenticationMethodConfiguration) SetIsSelfServiceRegistrationAllowed(value *bool)() {
    if m != nil {
        m.isSelfServiceRegistrationAllowed = value
    }
}
// SetKeyRestrictions sets the keyRestrictions property value. Controls whether key restrictions are enforced on FIDO2 security keys, either allowing or disallowing certain key types as defined by Authenticator Attestation GUID (AAGUID), an identifier that indicates the type (e.g. make and model) of the authenticator.
func (m *Fido2AuthenticationMethodConfiguration) SetKeyRestrictions(value Fido2KeyRestrictionsable)() {
    if m != nil {
        m.keyRestrictions = value
    }
}
