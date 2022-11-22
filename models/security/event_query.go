package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EventQuery 
type EventQuery struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The query property
    query *string
    // The queryType property
    queryType *QueryType
}
// NewEventQuery instantiates a new eventQuery and sets the default values.
func NewEventQuery()(*EventQuery) {
    m := &EventQuery{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEventQueryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEventQueryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEventQuery(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EventQuery) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EventQuery) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["query"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetQuery)
    res["queryType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseQueryType , m.SetQueryType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EventQuery) GetOdataType()(*string) {
    return m.odataType
}
// GetQuery gets the query property value. The query property
func (m *EventQuery) GetQuery()(*string) {
    return m.query
}
// GetQueryType gets the queryType property value. The queryType property
func (m *EventQuery) GetQueryType()(*QueryType) {
    return m.queryType
}
// Serialize serializes information the current object
func (m *EventQuery) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    if m.GetQueryType() != nil {
        cast := (*m.GetQueryType()).String()
        err := writer.WriteStringValue("queryType", &cast)
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
func (m *EventQuery) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EventQuery) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQuery sets the query property value. The query property
func (m *EventQuery) SetQuery(value *string)() {
    m.query = value
}
// SetQueryType sets the queryType property value. The queryType property
func (m *EventQuery) SetQueryType(value *QueryType)() {
    m.queryType = value
}
