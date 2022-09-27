package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContinuousAccessEvaluationSessionControl 
type ContinuousAccessEvaluationSessionControl struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies continuous access evaluation settings. The possible values are: strictEnforcement, disabled, unknownFutureValue.
    mode *ContinuousAccessEvaluationMode
    // The OdataType property
    odataType *string
}
// NewContinuousAccessEvaluationSessionControl instantiates a new continuousAccessEvaluationSessionControl and sets the default values.
func NewContinuousAccessEvaluationSessionControl()(*ContinuousAccessEvaluationSessionControl) {
    m := &ContinuousAccessEvaluationSessionControl{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.continuousAccessEvaluationSessionControl";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateContinuousAccessEvaluationSessionControlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContinuousAccessEvaluationSessionControlFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContinuousAccessEvaluationSessionControl(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContinuousAccessEvaluationSessionControl) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContinuousAccessEvaluationSessionControl) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseContinuousAccessEvaluationMode , m.SetMode)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetMode gets the mode property value. Specifies continuous access evaluation settings. The possible values are: strictEnforcement, disabled, unknownFutureValue.
func (m *ContinuousAccessEvaluationSessionControl) GetMode()(*ContinuousAccessEvaluationMode) {
    return m.mode
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ContinuousAccessEvaluationSessionControl) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ContinuousAccessEvaluationSessionControl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMode() != nil {
        cast := (*m.GetMode()).String()
        err := writer.WriteStringValue("mode", &cast)
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
func (m *ContinuousAccessEvaluationSessionControl) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMode sets the mode property value. Specifies continuous access evaluation settings. The possible values are: strictEnforcement, disabled, unknownFutureValue.
func (m *ContinuousAccessEvaluationSessionControl) SetMode(value *ContinuousAccessEvaluationMode)() {
    m.mode = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ContinuousAccessEvaluationSessionControl) SetOdataType(value *string)() {
    m.odataType = value
}
