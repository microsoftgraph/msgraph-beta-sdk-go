package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateWindow 
type UpdateWindow struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // End of a time window during which agents can receive updates
    updateWindowEndTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // Start of a time window during which agents can receive updates
    updateWindowStartTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
}
// NewUpdateWindow instantiates a new updateWindow and sets the default values.
func NewUpdateWindow()(*UpdateWindow) {
    m := &UpdateWindow{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.updateWindow";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUpdateWindowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateWindowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateWindow(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateWindow) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateWindow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["updateWindowEndTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeOnlyValue(m.SetUpdateWindowEndTime)
    res["updateWindowStartTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeOnlyValue(m.SetUpdateWindowStartTime)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UpdateWindow) GetOdataType()(*string) {
    return m.odataType
}
// GetUpdateWindowEndTime gets the updateWindowEndTime property value. End of a time window during which agents can receive updates
func (m *UpdateWindow) GetUpdateWindowEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.updateWindowEndTime
}
// GetUpdateWindowStartTime gets the updateWindowStartTime property value. Start of a time window during which agents can receive updates
func (m *UpdateWindow) GetUpdateWindowStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.updateWindowStartTime
}
// Serialize serializes information the current object
func (m *UpdateWindow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("updateWindowEndTime", m.GetUpdateWindowEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("updateWindowStartTime", m.GetUpdateWindowStartTime())
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
func (m *UpdateWindow) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UpdateWindow) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUpdateWindowEndTime sets the updateWindowEndTime property value. End of a time window during which agents can receive updates
func (m *UpdateWindow) SetUpdateWindowEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.updateWindowEndTime = value
}
// SetUpdateWindowStartTime sets the updateWindowStartTime property value. Start of a time window during which agents can receive updates
func (m *UpdateWindow) SetUpdateWindowStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.updateWindowStartTime = value
}
