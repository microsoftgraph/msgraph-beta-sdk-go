package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchLookupJob 
type ExactMatchLookupJob struct {
    ExactMatchJobBase
}
// NewExactMatchLookupJob instantiates a new exactMatchLookupJob and sets the default values.
func NewExactMatchLookupJob()(*ExactMatchLookupJob) {
    m := &ExactMatchLookupJob{
        ExactMatchJobBase: *NewExactMatchJobBase(),
    }
    odataTypeValue := "#microsoft.graph.exactMatchLookupJob"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateExactMatchLookupJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchLookupJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactMatchLookupJob(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchLookupJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExactMatchJobBase.GetFieldDeserializers()
    res["matchingRows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLookupResultRowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LookupResultRowable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LookupResultRowable)
                }
            }
            m.SetMatchingRows(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    return res
}
// GetMatchingRows gets the matchingRows property value. The matchingRows property
func (m *ExactMatchLookupJob) GetMatchingRows()([]LookupResultRowable) {
    val, err := m.GetBackingStore().Get("matchingRows")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]LookupResultRowable)
    }
    return nil
}
// GetState gets the state property value. The state property
func (m *ExactMatchLookupJob) GetState()(*string) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ExactMatchLookupJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExactMatchJobBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMatchingRows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMatchingRows()))
        for i, v := range m.GetMatchingRows() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("matchingRows", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMatchingRows sets the matchingRows property value. The matchingRows property
func (m *ExactMatchLookupJob) SetMatchingRows(value []LookupResultRowable)() {
    err := m.GetBackingStore().Set("matchingRows", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The state property
func (m *ExactMatchLookupJob) SetState(value *string)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// ExactMatchLookupJobable 
type ExactMatchLookupJobable interface {
    ExactMatchJobBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMatchingRows()([]LookupResultRowable)
    GetState()(*string)
    SetMatchingRows(value []LookupResultRowable)()
    SetState(value *string)()
}
