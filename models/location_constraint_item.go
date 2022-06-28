package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LocationConstraintItem 
type LocationConstraintItem struct {
    Location
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If set to true and the specified resource is busy, findMeetingTimes looks for another resource that is free. If set to false and the specified resource is busy, findMeetingTimes returns the resource best ranked in the user's cache without checking if it's free. Default is true.
    resolveAvailability *bool
}
// NewLocationConstraintItem instantiates a new LocationConstraintItem and sets the default values.
func NewLocationConstraintItem()(*LocationConstraintItem) {
    m := &LocationConstraintItem{
        Location: *NewLocation(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLocationConstraintItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocationConstraintItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocationConstraintItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LocationConstraintItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LocationConstraintItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Location.GetFieldDeserializers()
    res["resolveAvailability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolveAvailability(val)
        }
        return nil
    }
    return res
}
// GetResolveAvailability gets the resolveAvailability property value. If set to true and the specified resource is busy, findMeetingTimes looks for another resource that is free. If set to false and the specified resource is busy, findMeetingTimes returns the resource best ranked in the user's cache without checking if it's free. Default is true.
func (m *LocationConstraintItem) GetResolveAvailability()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.resolveAvailability
    }
}
// Serialize serializes information the current object
func (m *LocationConstraintItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Location.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("resolveAvailability", m.GetResolveAvailability())
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
func (m *LocationConstraintItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetResolveAvailability sets the resolveAvailability property value. If set to true and the specified resource is busy, findMeetingTimes looks for another resource that is free. If set to false and the specified resource is busy, findMeetingTimes returns the resource best ranked in the user's cache without checking if it's free. Default is true.
func (m *LocationConstraintItem) SetResolveAvailability(value *bool)() {
    if m != nil {
        m.resolveAvailability = value
    }
}
