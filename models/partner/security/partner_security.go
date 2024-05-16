package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type PartnerSecurity struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewPartnerSecurity instantiates a new PartnerSecurity and sets the default values.
func NewPartnerSecurity()(*PartnerSecurity) {
    m := &PartnerSecurity{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreatePartnerSecurityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePartnerSecurityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPartnerSecurity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PartnerSecurity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["securityAlerts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePartnerSecurityAlertFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PartnerSecurityAlertable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PartnerSecurityAlertable)
                }
            }
            m.SetSecurityAlerts(res)
        }
        return nil
    }
    return res
}
// GetSecurityAlerts gets the securityAlerts property value. The security alerts or a vulnerability of a CSP partner's customer that the partner must be made aware of for further action.
// returns a []PartnerSecurityAlertable when successful
func (m *PartnerSecurity) GetSecurityAlerts()([]PartnerSecurityAlertable) {
    val, err := m.GetBackingStore().Get("securityAlerts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PartnerSecurityAlertable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PartnerSecurity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSecurityAlerts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecurityAlerts()))
        for i, v := range m.GetSecurityAlerts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("securityAlerts", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSecurityAlerts sets the securityAlerts property value. The security alerts or a vulnerability of a CSP partner's customer that the partner must be made aware of for further action.
func (m *PartnerSecurity) SetSecurityAlerts(value []PartnerSecurityAlertable)() {
    err := m.GetBackingStore().Set("securityAlerts", value)
    if err != nil {
        panic(err)
    }
}
type PartnerSecurityable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSecurityAlerts()([]PartnerSecurityAlertable)
    SetSecurityAlerts(value []PartnerSecurityAlertable)()
}
