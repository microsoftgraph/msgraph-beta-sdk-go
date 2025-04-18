// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomClaimsPolicy struct {
    Entity
}
// NewCustomClaimsPolicy instantiates a new CustomClaimsPolicy and sets the default values.
func NewCustomClaimsPolicy()(*CustomClaimsPolicy) {
    m := &CustomClaimsPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCustomClaimsPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomClaimsPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomClaimsPolicy(), nil
}
// GetAudienceOverride gets the audienceOverride property value. If specified, it overrides the content of the audience claim for WS-Federation and SAML2 protocols. A custom signing key must be used for audienceOverride to be applied, otherwise, the audienceOverride value is ignored. The value provided must be in the format of an absolute URI.
// returns a *string when successful
func (m *CustomClaimsPolicy) GetAudienceOverride()(*string) {
    val, err := m.GetBackingStore().Get("audienceOverride")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetClaims gets the claims property value. Defines which claims are present in the tokens affected by the policy, in addition to the basic claim and the core claim set. Inherited from customclaimbase.
// returns a []CustomClaimBaseable when successful
func (m *CustomClaimsPolicy) GetClaims()([]CustomClaimBaseable) {
    val, err := m.GetBackingStore().Get("claims")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomClaimBaseable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomClaimsPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["audienceOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudienceOverride(val)
        }
        return nil
    }
    res["claims"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomClaimBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomClaimBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CustomClaimBaseable)
                }
            }
            m.SetClaims(res)
        }
        return nil
    }
    res["includeApplicationIdInIssuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeApplicationIdInIssuer(val)
        }
        return nil
    }
    res["includeBasicClaimSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeBasicClaimSet(val)
        }
        return nil
    }
    return res
}
// GetIncludeApplicationIdInIssuer gets the includeApplicationIdInIssuer property value. Indicates whether the application ID is added to the claim. It is relevant only for SAML2.0 and if a custom signing key is used. the default value is true. Optional.
// returns a *bool when successful
func (m *CustomClaimsPolicy) GetIncludeApplicationIdInIssuer()(*bool) {
    val, err := m.GetBackingStore().Get("includeApplicationIdInIssuer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIncludeBasicClaimSet gets the includeBasicClaimSet property value. Determines whether the basic claim set is included in tokens affected by this policy. If set to true, all claims in the basic claim set are emitted in tokens affected by the policy. By default the basic claim set isn't in the tokens unless they're explicitly configured in this policy.
// returns a *bool when successful
func (m *CustomClaimsPolicy) GetIncludeBasicClaimSet()(*bool) {
    val, err := m.GetBackingStore().Get("includeBasicClaimSet")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CustomClaimsPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("audienceOverride", m.GetAudienceOverride())
        if err != nil {
            return err
        }
    }
    if m.GetClaims() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClaims()))
        for i, v := range m.GetClaims() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("claims", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeApplicationIdInIssuer", m.GetIncludeApplicationIdInIssuer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeBasicClaimSet", m.GetIncludeBasicClaimSet())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudienceOverride sets the audienceOverride property value. If specified, it overrides the content of the audience claim for WS-Federation and SAML2 protocols. A custom signing key must be used for audienceOverride to be applied, otherwise, the audienceOverride value is ignored. The value provided must be in the format of an absolute URI.
func (m *CustomClaimsPolicy) SetAudienceOverride(value *string)() {
    err := m.GetBackingStore().Set("audienceOverride", value)
    if err != nil {
        panic(err)
    }
}
// SetClaims sets the claims property value. Defines which claims are present in the tokens affected by the policy, in addition to the basic claim and the core claim set. Inherited from customclaimbase.
func (m *CustomClaimsPolicy) SetClaims(value []CustomClaimBaseable)() {
    err := m.GetBackingStore().Set("claims", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeApplicationIdInIssuer sets the includeApplicationIdInIssuer property value. Indicates whether the application ID is added to the claim. It is relevant only for SAML2.0 and if a custom signing key is used. the default value is true. Optional.
func (m *CustomClaimsPolicy) SetIncludeApplicationIdInIssuer(value *bool)() {
    err := m.GetBackingStore().Set("includeApplicationIdInIssuer", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludeBasicClaimSet sets the includeBasicClaimSet property value. Determines whether the basic claim set is included in tokens affected by this policy. If set to true, all claims in the basic claim set are emitted in tokens affected by the policy. By default the basic claim set isn't in the tokens unless they're explicitly configured in this policy.
func (m *CustomClaimsPolicy) SetIncludeBasicClaimSet(value *bool)() {
    err := m.GetBackingStore().Set("includeBasicClaimSet", value)
    if err != nil {
        panic(err)
    }
}
type CustomClaimsPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudienceOverride()(*string)
    GetClaims()([]CustomClaimBaseable)
    GetIncludeApplicationIdInIssuer()(*bool)
    GetIncludeBasicClaimSet()(*bool)
    SetAudienceOverride(value *string)()
    SetClaims(value []CustomClaimBaseable)()
    SetIncludeApplicationIdInIssuer(value *bool)()
    SetIncludeBasicClaimSet(value *bool)()
}
