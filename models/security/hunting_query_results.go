package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HuntingQueryResults 
type HuntingQueryResults struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The results property
    results []HuntingRowResultable
    // The schema property
    schema []SinglePropertySchemaable
}
// NewHuntingQueryResults instantiates a new HuntingQueryResults and sets the default values.
func NewHuntingQueryResults()(*HuntingQueryResults) {
    m := &HuntingQueryResults{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateHuntingQueryResultsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHuntingQueryResultsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHuntingQueryResults(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HuntingQueryResults) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HuntingQueryResults) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["results"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHuntingRowResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HuntingRowResultable, len(val))
            for i, v := range val {
                res[i] = v.(HuntingRowResultable)
            }
            m.SetResults(res)
        }
        return nil
    }
    res["schema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSinglePropertySchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SinglePropertySchemaable, len(val))
            for i, v := range val {
                res[i] = v.(SinglePropertySchemaable)
            }
            m.SetSchema(res)
        }
        return nil
    }
    return res
}
// GetResults gets the results property value. The results property
func (m *HuntingQueryResults) GetResults()([]HuntingRowResultable) {
    if m == nil {
        return nil
    } else {
        return m.results
    }
}
// GetSchema gets the schema property value. The schema property
func (m *HuntingQueryResults) GetSchema()([]SinglePropertySchemaable) {
    if m == nil {
        return nil
    } else {
        return m.schema
    }
}
// Serialize serializes information the current object
func (m *HuntingQueryResults) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResults()))
        for i, v := range m.GetResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("results", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSchema() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSchema()))
        for i, v := range m.GetSchema() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("schema", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HuntingQueryResults) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetResults sets the results property value. The results property
func (m *HuntingQueryResults) SetResults(value []HuntingRowResultable)() {
    if m != nil {
        m.results = value
    }
}
// SetSchema sets the schema property value. The schema property
func (m *HuntingQueryResults) SetSchema(value []SinglePropertySchemaable)() {
    if m != nil {
        m.schema = value
    }
}
