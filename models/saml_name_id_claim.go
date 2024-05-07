package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SamlNameIdClaim struct {
    CustomClaimBase
}
// NewSamlNameIdClaim instantiates a new SamlNameIdClaim and sets the default values.
func NewSamlNameIdClaim()(*SamlNameIdClaim) {
    m := &SamlNameIdClaim{
        CustomClaimBase: *NewCustomClaimBase(),
    }
    odataTypeValue := "#microsoft.graph.samlNameIdClaim"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSamlNameIdClaimFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSamlNameIdClaimFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSamlNameIdClaim(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SamlNameIdClaim) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomClaimBase.GetFieldDeserializers()
    res["nameIdFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSamlNameIDFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameIdFormat(val.(*SamlNameIDFormat))
        }
        return nil
    }
    res["serviceProviderNameQualifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceProviderNameQualifier(val)
        }
        return nil
    }
    return res
}
// GetNameIdFormat gets the nameIdFormat property value. The nameIdFormat property
// returns a *SamlNameIDFormat when successful
func (m *SamlNameIdClaim) GetNameIdFormat()(*SamlNameIDFormat) {
    val, err := m.GetBackingStore().Get("nameIdFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SamlNameIDFormat)
    }
    return nil
}
// GetServiceProviderNameQualifier gets the serviceProviderNameQualifier property value. The serviceProviderNameQualifier property
// returns a *string when successful
func (m *SamlNameIdClaim) GetServiceProviderNameQualifier()(*string) {
    val, err := m.GetBackingStore().Get("serviceProviderNameQualifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SamlNameIdClaim) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomClaimBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetNameIdFormat() != nil {
        cast := (*m.GetNameIdFormat()).String()
        err = writer.WriteStringValue("nameIdFormat", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceProviderNameQualifier", m.GetServiceProviderNameQualifier())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetNameIdFormat sets the nameIdFormat property value. The nameIdFormat property
func (m *SamlNameIdClaim) SetNameIdFormat(value *SamlNameIDFormat)() {
    err := m.GetBackingStore().Set("nameIdFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetServiceProviderNameQualifier sets the serviceProviderNameQualifier property value. The serviceProviderNameQualifier property
func (m *SamlNameIdClaim) SetServiceProviderNameQualifier(value *string)() {
    err := m.GetBackingStore().Set("serviceProviderNameQualifier", value)
    if err != nil {
        panic(err)
    }
}
type SamlNameIdClaimable interface {
    CustomClaimBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNameIdFormat()(*SamlNameIDFormat)
    GetServiceProviderNameQualifier()(*string)
    SetNameIdFormat(value *SamlNameIDFormat)()
    SetServiceProviderNameQualifier(value *string)()
}
