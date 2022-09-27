package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActionStep 
type ActionStep struct {
    // The actionUrl property
    actionUrl ActionUrlable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The stepNumber property
    stepNumber *int64
    // The text property
    text *string
}
// NewActionStep instantiates a new actionStep and sets the default values.
func NewActionStep()(*ActionStep) {
    m := &ActionStep{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.actionStep";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateActionStepFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActionStepFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionStep(), nil
}
// GetActionUrl gets the actionUrl property value. The actionUrl property
func (m *ActionStep) GetActionUrl()(ActionUrlable) {
    return m.actionUrl
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActionStep) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActionStep) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateActionUrlFromDiscriminatorValue , m.SetActionUrl)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["stepNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetStepNumber)
    res["text"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetText)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ActionStep) GetOdataType()(*string) {
    return m.odataType
}
// GetStepNumber gets the stepNumber property value. The stepNumber property
func (m *ActionStep) GetStepNumber()(*int64) {
    return m.stepNumber
}
// GetText gets the text property value. The text property
func (m *ActionStep) GetText()(*string) {
    return m.text
}
// Serialize serializes information the current object
func (m *ActionStep) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("actionUrl", m.GetActionUrl())
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
        err := writer.WriteInt64Value("stepNumber", m.GetStepNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
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
// SetActionUrl sets the actionUrl property value. The actionUrl property
func (m *ActionStep) SetActionUrl(value ActionUrlable)() {
    m.actionUrl = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActionStep) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ActionStep) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStepNumber sets the stepNumber property value. The stepNumber property
func (m *ActionStep) SetStepNumber(value *int64)() {
    m.stepNumber = value
}
// SetText sets the text property value. The text property
func (m *ActionStep) SetText(value *string)() {
    m.text = value
}
