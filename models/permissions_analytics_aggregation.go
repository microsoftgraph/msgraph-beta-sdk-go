package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PermissionsAnalyticsAggregation struct {
    Entity
}
// NewPermissionsAnalyticsAggregation instantiates a new PermissionsAnalyticsAggregation and sets the default values.
func NewPermissionsAnalyticsAggregation()(*PermissionsAnalyticsAggregation) {
    m := &PermissionsAnalyticsAggregation{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePermissionsAnalyticsAggregationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePermissionsAnalyticsAggregationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPermissionsAnalyticsAggregation(), nil
}
// GetAws gets the aws property value. The aws property
// returns a PermissionsAnalyticsable when successful
func (m *PermissionsAnalyticsAggregation) GetAws()(PermissionsAnalyticsable) {
    val, err := m.GetBackingStore().Get("aws")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsAnalyticsable)
    }
    return nil
}
// GetAzure gets the azure property value. The azure property
// returns a PermissionsAnalyticsable when successful
func (m *PermissionsAnalyticsAggregation) GetAzure()(PermissionsAnalyticsable) {
    val, err := m.GetBackingStore().Get("azure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsAnalyticsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PermissionsAnalyticsAggregation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aws"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionsAnalyticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAws(val.(PermissionsAnalyticsable))
        }
        return nil
    }
    res["azure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionsAnalyticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzure(val.(PermissionsAnalyticsable))
        }
        return nil
    }
    res["gcp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionsAnalyticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGcp(val.(PermissionsAnalyticsable))
        }
        return nil
    }
    return res
}
// GetGcp gets the gcp property value. The gcp property
// returns a PermissionsAnalyticsable when successful
func (m *PermissionsAnalyticsAggregation) GetGcp()(PermissionsAnalyticsable) {
    val, err := m.GetBackingStore().Get("gcp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsAnalyticsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PermissionsAnalyticsAggregation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("aws", m.GetAws())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("azure", m.GetAzure())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("gcp", m.GetGcp())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAws sets the aws property value. The aws property
func (m *PermissionsAnalyticsAggregation) SetAws(value PermissionsAnalyticsable)() {
    err := m.GetBackingStore().Set("aws", value)
    if err != nil {
        panic(err)
    }
}
// SetAzure sets the azure property value. The azure property
func (m *PermissionsAnalyticsAggregation) SetAzure(value PermissionsAnalyticsable)() {
    err := m.GetBackingStore().Set("azure", value)
    if err != nil {
        panic(err)
    }
}
// SetGcp sets the gcp property value. The gcp property
func (m *PermissionsAnalyticsAggregation) SetGcp(value PermissionsAnalyticsable)() {
    err := m.GetBackingStore().Set("gcp", value)
    if err != nil {
        panic(err)
    }
}
type PermissionsAnalyticsAggregationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAws()(PermissionsAnalyticsable)
    GetAzure()(PermissionsAnalyticsable)
    GetGcp()(PermissionsAnalyticsable)
    SetAws(value PermissionsAnalyticsable)()
    SetAzure(value PermissionsAnalyticsable)()
    SetGcp(value PermissionsAnalyticsable)()
}
