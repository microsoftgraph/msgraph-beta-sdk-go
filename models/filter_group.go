package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FilterGroup 
type FilterGroup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Filter clauses (conditions) of this group. All clauses in a group must be satisfied in order for the filter group to evaluate to true.
    clauses []FilterClauseable
    // Human-readable name of the filter group.
    name *string
    // The OdataType property
    odataType *string
}
// NewFilterGroup instantiates a new filterGroup and sets the default values.
func NewFilterGroup()(*FilterGroup) {
    m := &FilterGroup{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.filterGroup";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateFilterGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFilterGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilterGroup(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FilterGroup) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClauses gets the clauses property value. Filter clauses (conditions) of this group. All clauses in a group must be satisfied in order for the filter group to evaluate to true.
func (m *FilterGroup) GetClauses()([]FilterClauseable) {
    return m.clauses
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FilterGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clauses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateFilterClauseFromDiscriminatorValue , m.SetClauses)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetName gets the name property value. Human-readable name of the filter group.
func (m *FilterGroup) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *FilterGroup) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *FilterGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClauses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetClauses())
        err := writer.WriteCollectionOfObjectValues("clauses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *FilterGroup) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClauses sets the clauses property value. Filter clauses (conditions) of this group. All clauses in a group must be satisfied in order for the filter group to evaluate to true.
func (m *FilterGroup) SetClauses(value []FilterClauseable)() {
    m.clauses = value
}
// SetName sets the name property value. Human-readable name of the filter group.
func (m *FilterGroup) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *FilterGroup) SetOdataType(value *string)() {
    m.odataType = value
}
