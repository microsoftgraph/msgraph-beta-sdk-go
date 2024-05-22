package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomClaim struct {
    CustomClaimBase
}
// NewCustomClaim instantiates a new CustomClaim and sets the default values.
func NewCustomClaim()(*CustomClaim) {
    m := &CustomClaim{
        CustomClaimBase: *NewCustomClaimBase(),
    }
    odataTypeValue := "#microsoft.graph.customClaim"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomClaimFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomClaimFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomClaim(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomClaim) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomClaimBase.GetFieldDeserializers()
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["namespace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNamespace(val)
        }
        return nil
    }
    res["samlAttributeNameFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSamlAttributeNameFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSamlAttributeNameFormat(val.(*SamlAttributeNameFormat))
        }
        return nil
    }
    res["tokenFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseTokenFormat)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenFormat, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*TokenFormat))
                }
            }
            m.SetTokenFormat(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the claim to be emitted.
// returns a *string when successful
func (m *CustomClaim) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNamespace gets the namespace property value. An optional namespace to be included as part of the claim name.
// returns a *string when successful
func (m *CustomClaim) GetNamespace()(*string) {
    val, err := m.GetBackingStore().Get("namespace")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSamlAttributeNameFormat gets the samlAttributeNameFormat property value. If specified, it sets the nameFormat attribute associated with the claim in the SAML response. The possible values are: unspecified, uri, basic, unknownFutureValue.
// returns a *SamlAttributeNameFormat when successful
func (m *CustomClaim) GetSamlAttributeNameFormat()(*SamlAttributeNameFormat) {
    val, err := m.GetBackingStore().Get("samlAttributeNameFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SamlAttributeNameFormat)
    }
    return nil
}
// GetTokenFormat gets the tokenFormat property value. List of token formats for which this claim should be emitted. The possible values are: saml,jwt, unknownFutureValue
// returns a []TokenFormat when successful
func (m *CustomClaim) GetTokenFormat()([]TokenFormat) {
    val, err := m.GetBackingStore().Get("tokenFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TokenFormat)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomClaim) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomClaimBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("namespace", m.GetNamespace())
        if err != nil {
            return err
        }
    }
    if m.GetSamlAttributeNameFormat() != nil {
        cast := (*m.GetSamlAttributeNameFormat()).String()
        err = writer.WriteStringValue("samlAttributeNameFormat", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenFormat() != nil {
        err = writer.WriteCollectionOfStringValues("tokenFormat", SerializeTokenFormat(m.GetTokenFormat()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The name of the claim to be emitted.
func (m *CustomClaim) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetNamespace sets the namespace property value. An optional namespace to be included as part of the claim name.
func (m *CustomClaim) SetNamespace(value *string)() {
    err := m.GetBackingStore().Set("namespace", value)
    if err != nil {
        panic(err)
    }
}
// SetSamlAttributeNameFormat sets the samlAttributeNameFormat property value. If specified, it sets the nameFormat attribute associated with the claim in the SAML response. The possible values are: unspecified, uri, basic, unknownFutureValue.
func (m *CustomClaim) SetSamlAttributeNameFormat(value *SamlAttributeNameFormat)() {
    err := m.GetBackingStore().Set("samlAttributeNameFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenFormat sets the tokenFormat property value. List of token formats for which this claim should be emitted. The possible values are: saml,jwt, unknownFutureValue
func (m *CustomClaim) SetTokenFormat(value []TokenFormat)() {
    err := m.GetBackingStore().Set("tokenFormat", value)
    if err != nil {
        panic(err)
    }
}
type CustomClaimable interface {
    CustomClaimBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetNamespace()(*string)
    GetSamlAttributeNameFormat()(*SamlAttributeNameFormat)
    GetTokenFormat()([]TokenFormat)
    SetName(value *string)()
    SetNamespace(value *string)()
    SetSamlAttributeNameFormat(value *SamlAttributeNameFormat)()
    SetTokenFormat(value []TokenFormat)()
}
