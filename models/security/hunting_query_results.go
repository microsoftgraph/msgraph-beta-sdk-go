package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HuntingQueryResults 
type HuntingQueryResults struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The results property
    results []HuntingRowResultable
    // The schema property
    schema []SinglePropertySchemaable
}
// NewHuntingQueryResults instantiates a new huntingQueryResults and sets the default values.
func NewHuntingQueryResults()(*HuntingQueryResults) {
    m := &HuntingQueryResults{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.security.huntingQueryResults";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateHuntingQueryResultsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHuntingQueryResultsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHuntingQueryResults(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HuntingQueryResults) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HuntingQueryResults) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["results"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateHuntingRowResultFromDiscriminatorValue , m.SetResults)
    res["schema"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSinglePropertySchemaFromDiscriminatorValue , m.SetSchema)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *HuntingQueryResults) GetOdataType()(*string) {
    return m.odataType
}
// GetResults gets the results property value. The results property
func (m *HuntingQueryResults) GetResults()([]HuntingRowResultable) {
    return m.results
}
// GetSchema gets the schema property value. The schema property
func (m *HuntingQueryResults) GetSchema()([]SinglePropertySchemaable) {
    return m.schema
}
// Serialize serializes information the current object
func (m *HuntingQueryResults) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetResults() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetResults())
        err := writer.WriteCollectionOfObjectValues("results", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSchema() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSchema())
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
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *HuntingQueryResults) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResults sets the results property value. The results property
func (m *HuntingQueryResults) SetResults(value []HuntingRowResultable)() {
    m.results = value
}
// SetSchema sets the schema property value. The schema property
func (m *HuntingQueryResults) SetSchema(value []SinglePropertySchemaable)() {
    m.schema = value
}
