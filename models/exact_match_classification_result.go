package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchClassificationResult 
type ExactMatchClassificationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The classification property
    classification []ExactMatchDetectedSensitiveContentable
    // The errors property
    errors []ClassificationErrorable
    // The OdataType property
    odataType *string
}
// NewExactMatchClassificationResult instantiates a new exactMatchClassificationResult and sets the default values.
func NewExactMatchClassificationResult()(*ExactMatchClassificationResult) {
    m := &ExactMatchClassificationResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExactMatchClassificationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchClassificationResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExactMatchClassificationResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExactMatchClassificationResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClassification gets the classification property value. The classification property
func (m *ExactMatchClassificationResult) GetClassification()([]ExactMatchDetectedSensitiveContentable) {
    return m.classification
}
// GetErrors gets the errors property value. The errors property
func (m *ExactMatchClassificationResult) GetErrors()([]ClassificationErrorable) {
    return m.errors
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchClassificationResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExactMatchDetectedSensitiveContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactMatchDetectedSensitiveContentable, len(val))
            for i, v := range val {
                res[i] = v.(ExactMatchDetectedSensitiveContentable)
            }
            m.SetClassification(res)
        }
        return nil
    }
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClassificationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClassificationErrorable, len(val))
            for i, v := range val {
                res[i] = v.(ClassificationErrorable)
            }
            m.SetErrors(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ExactMatchClassificationResult) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ExactMatchClassificationResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClassification() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClassification()))
        for i, v := range m.GetClassification() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("classification", cast)
        if err != nil {
            return err
        }
    }
    if m.GetErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetErrors()))
        for i, v := range m.GetErrors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("errors", cast)
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
func (m *ExactMatchClassificationResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClassification sets the classification property value. The classification property
func (m *ExactMatchClassificationResult) SetClassification(value []ExactMatchDetectedSensitiveContentable)() {
    m.classification = value
}
// SetErrors sets the errors property value. The errors property
func (m *ExactMatchClassificationResult) SetErrors(value []ClassificationErrorable)() {
    m.errors = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExactMatchClassificationResult) SetOdataType(value *string)() {
    m.odataType = value
}
