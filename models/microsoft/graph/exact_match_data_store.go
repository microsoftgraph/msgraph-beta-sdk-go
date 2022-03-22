package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExactMatchDataStore 
type ExactMatchDataStore struct {
    ExactMatchDataStoreBase
    // 
    sessions []ExactMatchSessionable;
}
// NewExactMatchDataStore instantiates a new exactMatchDataStore and sets the default values.
func NewExactMatchDataStore()(*ExactMatchDataStore) {
    m := &ExactMatchDataStore{
        ExactMatchDataStoreBase: *NewExactMatchDataStoreBase(),
    }
    return m
}
// CreateExactMatchDataStoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchDataStoreFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewExactMatchDataStore(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchDataStore) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ExactMatchDataStoreBase.GetFieldDeserializers()
    res["sessions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetSessions gets the sessions property value. 
func (m *ExactMatchDataStore) GetSessions()([]ExactMatchSessionable) {
    if m == nil {
        return nil
    } else {
        return m.sessions
    }
}
// Serialize serializes information the current object
func (m *ExactMatchDataStore) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ExactMatchDataStoreBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSessions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSessions()))
        for i, v := range m.GetSessions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sessions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSessions sets the sessions property value. 
func (m *ExactMatchDataStore) SetSessions(value []ExactMatchSessionable)() {
    if m != nil {
        m.sessions = value
    }
}
