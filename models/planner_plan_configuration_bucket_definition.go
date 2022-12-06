package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanConfigurationBucketDefinition 
type PlannerPlanConfigurationBucketDefinition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The externalBucketId property
    externalBucketId *string
    // The OdataType property
    odataType *string
}
// NewPlannerPlanConfigurationBucketDefinition instantiates a new plannerPlanConfigurationBucketDefinition and sets the default values.
func NewPlannerPlanConfigurationBucketDefinition()(*PlannerPlanConfigurationBucketDefinition) {
    m := &PlannerPlanConfigurationBucketDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlannerPlanConfigurationBucketDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanConfigurationBucketDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlanConfigurationBucketDefinition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerPlanConfigurationBucketDefinition) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExternalBucketId gets the externalBucketId property value. The externalBucketId property
func (m *PlannerPlanConfigurationBucketDefinition) GetExternalBucketId()(*string) {
    return m.externalBucketId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanConfigurationBucketDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalBucketId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalBucketId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerPlanConfigurationBucketDefinition) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *PlannerPlanConfigurationBucketDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("externalBucketId", m.GetExternalBucketId())
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
func (m *PlannerPlanConfigurationBucketDefinition) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExternalBucketId sets the externalBucketId property value. The externalBucketId property
func (m *PlannerPlanConfigurationBucketDefinition) SetExternalBucketId(value *string)() {
    m.externalBucketId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerPlanConfigurationBucketDefinition) SetOdataType(value *string)() {
    m.odataType = value
}
