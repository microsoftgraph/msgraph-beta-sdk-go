package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContinuousAccessEvaluationSessionControl 
type ContinuousAccessEvaluationSessionControl struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies continuous access evaluation settings. The possible values are: strictEnforcement, disabled, unknownFutureValue.
    mode *ContinuousAccessEvaluationMode;
}
// NewContinuousAccessEvaluationSessionControl instantiates a new continuousAccessEvaluationSessionControl and sets the default values.
func NewContinuousAccessEvaluationSessionControl()(*ContinuousAccessEvaluationSessionControl) {
    m := &ContinuousAccessEvaluationSessionControl{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateContinuousAccessEvaluationSessionControlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContinuousAccessEvaluationSessionControlFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContinuousAccessEvaluationSessionControl(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContinuousAccessEvaluationSessionControl) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContinuousAccessEvaluationSessionControl) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mode"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContinuousAccessEvaluationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMode(val.(*ContinuousAccessEvaluationMode))
        }
        return nil
    }
    return res
}
// GetMode gets the mode property value. Specifies continuous access evaluation settings. The possible values are: strictEnforcement, disabled, unknownFutureValue.
func (m *ContinuousAccessEvaluationSessionControl) GetMode()(*ContinuousAccessEvaluationMode) {
    if m == nil {
        return nil
    } else {
        return m.mode
    }
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContinuousAccessEvaluationSessionControl) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMode sets the mode property value. Specifies continuous access evaluation settings. The possible values are: strictEnforcement, disabled, unknownFutureValue.
func (m *ContinuousAccessEvaluationSessionControl) SetMode(value *ContinuousAccessEvaluationMode)() {
    if m != nil {
        m.mode = value
    }
}
