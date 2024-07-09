package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AzureADAuthentication struct {
    Entity
}
// NewAzureADAuthentication instantiates a new AzureADAuthentication and sets the default values.
func NewAzureADAuthentication()(*AzureADAuthentication) {
    m := &AzureADAuthentication{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAzureADAuthenticationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAzureADAuthenticationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureADAuthentication(), nil
}
// GetAttainments gets the attainments property value. SLA data for a Microsoft Entra tenant for a calendar month.
// returns a []ServiceLevelAgreementAttainmentable when successful
func (m *AzureADAuthentication) GetAttainments()([]ServiceLevelAgreementAttainmentable) {
    val, err := m.GetBackingStore().Get("attainments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ServiceLevelAgreementAttainmentable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AzureADAuthentication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attainments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceLevelAgreementAttainmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceLevelAgreementAttainmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ServiceLevelAgreementAttainmentable)
                }
            }
            m.SetAttainments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AzureADAuthentication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttainments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttainments()))
        for i, v := range m.GetAttainments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("attainments", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttainments sets the attainments property value. SLA data for a Microsoft Entra tenant for a calendar month.
func (m *AzureADAuthentication) SetAttainments(value []ServiceLevelAgreementAttainmentable)() {
    err := m.GetBackingStore().Set("attainments", value)
    if err != nil {
        panic(err)
    }
}
type AzureADAuthenticationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttainments()([]ServiceLevelAgreementAttainmentable)
    SetAttainments(value []ServiceLevelAgreementAttainmentable)()
}
