package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerFieldRules 
type PlannerFieldRules struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The defaultRules property
    defaultRules []string
    // The OdataType property
    odataType *string
    // The overrides property
    overrides []PlannerRuleOverrideable
}
// NewPlannerFieldRules instantiates a new plannerFieldRules and sets the default values.
func NewPlannerFieldRules()(*PlannerFieldRules) {
    m := &PlannerFieldRules{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlannerFieldRulesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerFieldRulesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerFieldRules(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerFieldRules) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDefaultRules gets the defaultRules property value. The defaultRules property
func (m *PlannerFieldRules) GetDefaultRules()([]string) {
    return m.defaultRules
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerFieldRules) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultRules"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDefaultRules)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["overrides"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePlannerRuleOverrideFromDiscriminatorValue , m.SetOverrides)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerFieldRules) GetOdataType()(*string) {
    return m.odataType
}
// GetOverrides gets the overrides property value. The overrides property
func (m *PlannerFieldRules) GetOverrides()([]PlannerRuleOverrideable) {
    return m.overrides
}
// Serialize serializes information the current object
func (m *PlannerFieldRules) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDefaultRules() != nil {
        err := writer.WriteCollectionOfStringValues("defaultRules", m.GetDefaultRules())
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
    if m.GetOverrides() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOverrides())
        err := writer.WriteCollectionOfObjectValues("overrides", cast)
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
func (m *PlannerFieldRules) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDefaultRules sets the defaultRules property value. The defaultRules property
func (m *PlannerFieldRules) SetDefaultRules(value []string)() {
    m.defaultRules = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerFieldRules) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOverrides sets the overrides property value. The overrides property
func (m *PlannerFieldRules) SetOverrides(value []PlannerRuleOverrideable)() {
    m.overrides = value
}
