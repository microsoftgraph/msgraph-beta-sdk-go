package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Goals 
type Goals struct {
    Entity
}
// NewGoals instantiates a new goals and sets the default values.
func NewGoals()(*Goals) {
    m := &Goals{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGoalsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGoalsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGoals(), nil
}
// GetExportJobs gets the exportJobs property value. Represents a collection of goals export jobs for Viva Goals.
func (m *Goals) GetExportJobs()([]GoalsExportJobable) {
    val, err := m.GetBackingStore().Get("exportJobs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GoalsExportJobable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Goals) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exportJobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGoalsExportJobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GoalsExportJobable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GoalsExportJobable)
                }
            }
            m.SetExportJobs(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Goals) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExportJobs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExportJobs()))
        for i, v := range m.GetExportJobs() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("exportJobs", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExportJobs sets the exportJobs property value. Represents a collection of goals export jobs for Viva Goals.
func (m *Goals) SetExportJobs(value []GoalsExportJobable)() {
    err := m.GetBackingStore().Set("exportJobs", value)
    if err != nil {
        panic(err)
    }
}
// Goalsable 
type Goalsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExportJobs()([]GoalsExportJobable)
    SetExportJobs(value []GoalsExportJobable)()
}
