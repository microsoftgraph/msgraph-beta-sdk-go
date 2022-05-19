package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExtractSensitivityLabelsResult 
type ExtractSensitivityLabelsResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // List of sensitivity labels assigned to a file.
    labels []SensitivityLabelAssignmentable
}
// NewExtractSensitivityLabelsResult instantiates a new ExtractSensitivityLabelsResult and sets the default values.
func NewExtractSensitivityLabelsResult()(*ExtractSensitivityLabelsResult) {
    m := &ExtractSensitivityLabelsResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExtractSensitivityLabelsResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExtractSensitivityLabelsResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExtractSensitivityLabelsResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExtractSensitivityLabelsResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExtractSensitivityLabelsResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitivityLabelAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabelAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(SensitivityLabelAssignmentable)
            }
            m.SetLabels(res)
        }
        return nil
    }
    return res
}
// GetLabels gets the labels property value. List of sensitivity labels assigned to a file.
func (m *ExtractSensitivityLabelsResult) GetLabels()([]SensitivityLabelAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.labels
    }
}
// Serialize serializes information the current object
func (m *ExtractSensitivityLabelsResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("labels", cast)
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
    if m != nil {
        m.additionalData = value
    }
}
// SetLabels sets the labels property value. List of sensitivity labels assigned to a file.
func (m *ExtractSensitivityLabelsResult) SetLabels(value []SensitivityLabelAssignmentable)() {
    if m != nil {
        m.labels = value
    }
}
