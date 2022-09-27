package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExtractSensitivityLabelsResult 
type ExtractSensitivityLabelsResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // List of sensitivity labels assigned to a file.
    labels []SensitivityLabelAssignmentable
    // The OdataType property
    odataType *string
}
// NewExtractSensitivityLabelsResult instantiates a new extractSensitivityLabelsResult and sets the default values.
func NewExtractSensitivityLabelsResult()(*ExtractSensitivityLabelsResult) {
    m := &ExtractSensitivityLabelsResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.extractSensitivityLabelsResult";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExtractSensitivityLabelsResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExtractSensitivityLabelsResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExtractSensitivityLabelsResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExtractSensitivityLabelsResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExtractSensitivityLabelsResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["labels"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSensitivityLabelAssignmentFromDiscriminatorValue , m.SetLabels)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetLabels gets the labels property value. List of sensitivity labels assigned to a file.
func (m *ExtractSensitivityLabelsResult) GetLabels()([]SensitivityLabelAssignmentable) {
    return m.labels
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ExtractSensitivityLabelsResult) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ExtractSensitivityLabelsResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLabels() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLabels())
        err := writer.WriteCollectionOfObjectValues("labels", cast)
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
func (m *ExtractSensitivityLabelsResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLabels sets the labels property value. List of sensitivity labels assigned to a file.
func (m *ExtractSensitivityLabelsResult) SetLabels(value []SensitivityLabelAssignmentable)() {
    m.labels = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExtractSensitivityLabelsResult) SetOdataType(value *string)() {
    m.odataType = value
}
