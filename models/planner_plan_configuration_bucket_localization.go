package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanConfigurationBucketLocalization 
type PlannerPlanConfigurationBucketLocalization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The externalBucketId property
    externalBucketId *string
    // The name property
    name *string
    // The OdataType property
    odataType *string
}
// NewPlannerPlanConfigurationBucketLocalization instantiates a new plannerPlanConfigurationBucketLocalization and sets the default values.
func NewPlannerPlanConfigurationBucketLocalization()(*PlannerPlanConfigurationBucketLocalization) {
    m := &PlannerPlanConfigurationBucketLocalization{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlannerPlanConfigurationBucketLocalizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanConfigurationBucketLocalizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerPlanConfigurationBucketLocalization(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerPlanConfigurationBucketLocalization) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExternalBucketId gets the externalBucketId property value. The externalBucketId property
func (m *PlannerPlanConfigurationBucketLocalization) GetExternalBucketId()(*string) {
    return m.externalBucketId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanConfigurationBucketLocalization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalBucketId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalBucketId(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
// GetName gets the name property value. The name property
func (m *PlannerPlanConfigurationBucketLocalization) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerPlanConfigurationBucketLocalization) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *PlannerPlanConfigurationBucketLocalization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("externalBucketId", m.GetExternalBucketId())
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
func (m *PlannerPlanConfigurationBucketLocalization) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExternalBucketId sets the externalBucketId property value. The externalBucketId property
func (m *PlannerPlanConfigurationBucketLocalization) SetExternalBucketId(value *string)() {
    m.externalBucketId = value
}
// SetName sets the name property value. The name property
func (m *PlannerPlanConfigurationBucketLocalization) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerPlanConfigurationBucketLocalization) SetOdataType(value *string)() {
    m.odataType = value
}
