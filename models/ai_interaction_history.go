package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AiInteractionHistory struct {
    Entity
}
// NewAiInteractionHistory instantiates a new AiInteractionHistory and sets the default values.
func NewAiInteractionHistory()(*AiInteractionHistory) {
    m := &AiInteractionHistory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAiInteractionHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAiInteractionHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAiInteractionHistory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AiInteractionHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["interactions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAiInteractionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AiInteractionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AiInteractionable)
                }
            }
            m.SetInteractions(res)
        }
        return nil
    }
    return res
}
// GetInteractions gets the interactions property value. The interactions property
// returns a []AiInteractionable when successful
func (m *AiInteractionHistory) GetInteractions()([]AiInteractionable) {
    val, err := m.GetBackingStore().Get("interactions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AiInteractionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AiInteractionHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetInteractions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInteractions()))
        for i, v := range m.GetInteractions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("interactions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInteractions sets the interactions property value. The interactions property
func (m *AiInteractionHistory) SetInteractions(value []AiInteractionable)() {
    err := m.GetBackingStore().Set("interactions", value)
    if err != nil {
        panic(err)
    }
}
type AiInteractionHistoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInteractions()([]AiInteractionable)
    SetInteractions(value []AiInteractionable)()
}
