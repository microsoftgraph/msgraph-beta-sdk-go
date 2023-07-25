package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcFrontLineServicePlan 
type CloudPcFrontLineServicePlan struct {
    Entity
}
// NewCloudPcFrontLineServicePlan instantiates a new cloudPcFrontLineServicePlan and sets the default values.
func NewCloudPcFrontLineServicePlan()(*CloudPcFrontLineServicePlan) {
    m := &CloudPcFrontLineServicePlan{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcFrontLineServicePlanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcFrontLineServicePlanFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcFrontLineServicePlan(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *CloudPcFrontLineServicePlan) GetDisplayName()(*string) {
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
func (m *CloudPcFrontLineServicePlan) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["totalCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    res["usedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the totalCount property value. The totalCount property
func (m *CloudPcFrontLineServicePlan) GetTotalCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUsedCount gets the usedCount property value. The usedCount property
func (m *CloudPcFrontLineServicePlan) GetUsedCount()(*int32) {
    val, err := m.GetBackingStore().Get("usedCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcFrontLineServicePlan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalCount", m.GetTotalCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usedCount", m.GetUsedCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CloudPcFrontLineServicePlan) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalCount sets the totalCount property value. The totalCount property
func (m *CloudPcFrontLineServicePlan) SetTotalCount(value *int32)() {
    err := m.GetBackingStore().Set("totalCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUsedCount sets the usedCount property value. The usedCount property
func (m *CloudPcFrontLineServicePlan) SetUsedCount(value *int32)() {
    err := m.GetBackingStore().Set("usedCount", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcFrontLineServicePlanable 
type CloudPcFrontLineServicePlanable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetTotalCount()(*int32)
    GetUsedCount()(*int32)
    SetDisplayName(value *string)()
    SetTotalCount(value *int32)()
    SetUsedCount(value *int32)()
}
