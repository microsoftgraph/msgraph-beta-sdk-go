package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PermissionsAnalytics struct {
    Entity
}
// NewPermissionsAnalytics instantiates a new PermissionsAnalytics and sets the default values.
func NewPermissionsAnalytics()(*PermissionsAnalytics) {
    m := &PermissionsAnalytics{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePermissionsAnalyticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePermissionsAnalyticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPermissionsAnalytics(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PermissionsAnalytics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["findings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFindingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Findingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Findingable)
                }
            }
            m.SetFindings(res)
        }
        return nil
    }
    res["permissionsCreepIndexDistributions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionsCreepIndexDistributionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionsCreepIndexDistributionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PermissionsCreepIndexDistributionable)
                }
            }
            m.SetPermissionsCreepIndexDistributions(res)
        }
        return nil
    }
    return res
}
// GetFindings gets the findings property value. The output of the permissions usage data analysis performed by Permissions Management to assess risk with identities and resources.
// returns a []Findingable when successful
func (m *PermissionsAnalytics) GetFindings()([]Findingable) {
    val, err := m.GetBackingStore().Get("findings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Findingable)
    }
    return nil
}
// GetPermissionsCreepIndexDistributions gets the permissionsCreepIndexDistributions property value. Represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
// returns a []PermissionsCreepIndexDistributionable when successful
func (m *PermissionsAnalytics) GetPermissionsCreepIndexDistributions()([]PermissionsCreepIndexDistributionable) {
    val, err := m.GetBackingStore().Get("permissionsCreepIndexDistributions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PermissionsCreepIndexDistributionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PermissionsAnalytics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetFindings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFindings()))
        for i, v := range m.GetFindings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("findings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPermissionsCreepIndexDistributions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPermissionsCreepIndexDistributions()))
        for i, v := range m.GetPermissionsCreepIndexDistributions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("permissionsCreepIndexDistributions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFindings sets the findings property value. The output of the permissions usage data analysis performed by Permissions Management to assess risk with identities and resources.
func (m *PermissionsAnalytics) SetFindings(value []Findingable)() {
    err := m.GetBackingStore().Set("findings", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionsCreepIndexDistributions sets the permissionsCreepIndexDistributions property value. Represents the Permissions Creep Index (PCI) for the authorization system. PCI distribution chart shows the classification of human and nonhuman identities based on the PCI score in three buckets (low, medium, high).
func (m *PermissionsAnalytics) SetPermissionsCreepIndexDistributions(value []PermissionsCreepIndexDistributionable)() {
    err := m.GetBackingStore().Set("permissionsCreepIndexDistributions", value)
    if err != nil {
        panic(err)
    }
}
type PermissionsAnalyticsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFindings()([]Findingable)
    GetPermissionsCreepIndexDistributions()([]PermissionsCreepIndexDistributionable)
    SetFindings(value []Findingable)()
    SetPermissionsCreepIndexDistributions(value []PermissionsCreepIndexDistributionable)()
}
