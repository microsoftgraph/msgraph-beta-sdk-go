package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CallSettings struct {
    Entity
}
// NewCallSettings instantiates a new CallSettings and sets the default values.
func NewCallSettings()(*CallSettings) {
    m := &CallSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCallSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCallSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallSettings(), nil
}
// GetDelegates gets the delegates property value. The delegates property
// returns a []DelegationSettingsable when successful
func (m *CallSettings) GetDelegates()([]DelegationSettingsable) {
    val, err := m.GetBackingStore().Get("delegates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DelegationSettingsable)
    }
    return nil
}
// GetDelegators gets the delegators property value. The delegators property
// returns a []DelegationSettingsable when successful
func (m *CallSettings) GetDelegators()([]DelegationSettingsable) {
    val, err := m.GetBackingStore().Get("delegators")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DelegationSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CallSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["delegates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegationSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegationSettingsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DelegationSettingsable)
                }
            }
            m.SetDelegates(res)
        }
        return nil
    }
    res["delegators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegationSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegationSettingsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DelegationSettingsable)
                }
            }
            m.SetDelegators(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CallSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDelegates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDelegates()))
        for i, v := range m.GetDelegates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("delegates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDelegators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDelegators()))
        for i, v := range m.GetDelegators() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("delegators", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDelegates sets the delegates property value. The delegates property
func (m *CallSettings) SetDelegates(value []DelegationSettingsable)() {
    err := m.GetBackingStore().Set("delegates", value)
    if err != nil {
        panic(err)
    }
}
// SetDelegators sets the delegators property value. The delegators property
func (m *CallSettings) SetDelegators(value []DelegationSettingsable)() {
    err := m.GetBackingStore().Set("delegators", value)
    if err != nil {
        panic(err)
    }
}
type CallSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDelegates()([]DelegationSettingsable)
    GetDelegators()([]DelegationSettingsable)
    SetDelegates(value []DelegationSettingsable)()
    SetDelegators(value []DelegationSettingsable)()
}
