package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchDataStore 
type ExactMatchDataStore struct {
    ExactMatchDataStoreBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The sessions property
    sessions []ExactMatchSessionable
}
// NewExactMatchDataStore instantiates a new ExactMatchDataStore and sets the default values.
func NewExactMatchDataStore()(*ExactMatchDataStore) {
    m := &ExactMatchDataStore{
        ExactMatchDataStoreBase: *NewExactMatchDataStoreBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExactMatchDataStoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchDataStoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactMatchDataStore(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExactMatchDataStore) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchDataStore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExactMatchDataStoreBase.GetFieldDeserializers()
    res["sessions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExactMatchSessionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactMatchSessionable, len(val))
            for i, v := range val {
                res[i] = v.(ExactMatchSessionable)
            }
            m.SetSessions(res)
        }
        return nil
    }
    return res
}
// GetSessions gets the sessions property value. The sessions property
func (m *ExactMatchDataStore) GetSessions()([]ExactMatchSessionable) {
    if m == nil {
        return nil
    } else {
        return m.sessions
    }
}
// Serialize serializes information the current object
func (m *ExactMatchDataStore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExactMatchDataStoreBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSessions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSessions()))
        for i, v := range m.GetSessions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sessions", cast)
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
func (m *ExactMatchDataStore) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSessions sets the sessions property value. The sessions property
func (m *ExactMatchDataStore) SetSessions(value []ExactMatchSessionable)() {
    if m != nil {
        m.sessions = value
    }
}
