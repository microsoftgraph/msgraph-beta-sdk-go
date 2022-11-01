package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DlpEvaluationInput 
type DlpEvaluationInput struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The currentLabel property
    currentLabel CurrentLabelable
    // The discoveredSensitiveTypes property
    discoveredSensitiveTypes []DiscoveredSensitiveTypeable
    // The OdataType property
    odataType *string
}
// NewDlpEvaluationInput instantiates a new dlpEvaluationInput and sets the default values.
func NewDlpEvaluationInput()(*DlpEvaluationInput) {
    m := &DlpEvaluationInput{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.dlpEvaluationInput";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDlpEvaluationInputFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDlpEvaluationInputFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.dlpEvaluationWindowsDevicesInput":
                        return NewDlpEvaluationWindowsDevicesInput(), nil
                }
            }
        }
    }
    return NewDlpEvaluationInput(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpEvaluationInput) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCurrentLabel gets the currentLabel property value. The currentLabel property
func (m *DlpEvaluationInput) GetCurrentLabel()(CurrentLabelable) {
    return m.currentLabel
}
// GetDiscoveredSensitiveTypes gets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *DlpEvaluationInput) GetDiscoveredSensitiveTypes()([]DiscoveredSensitiveTypeable) {
    return m.discoveredSensitiveTypes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DlpEvaluationInput) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentLabel"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCurrentLabelFromDiscriminatorValue , m.SetCurrentLabel)
    res["discoveredSensitiveTypes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDiscoveredSensitiveTypeFromDiscriminatorValue , m.SetDiscoveredSensitiveTypes)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DlpEvaluationInput) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DlpEvaluationInput) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("currentLabel", m.GetCurrentLabel())
        if err != nil {
            return err
        }
    }
    if m.GetDiscoveredSensitiveTypes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDiscoveredSensitiveTypes())
        err := writer.WriteCollectionOfObjectValues("discoveredSensitiveTypes", cast)
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
func (m *DlpEvaluationInput) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCurrentLabel sets the currentLabel property value. The currentLabel property
func (m *DlpEvaluationInput) SetCurrentLabel(value CurrentLabelable)() {
    m.currentLabel = value
}
// SetDiscoveredSensitiveTypes sets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *DlpEvaluationInput) SetDiscoveredSensitiveTypes(value []DiscoveredSensitiveTypeable)() {
    m.discoveredSensitiveTypes = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DlpEvaluationInput) SetOdataType(value *string)() {
    m.odataType = value
}
