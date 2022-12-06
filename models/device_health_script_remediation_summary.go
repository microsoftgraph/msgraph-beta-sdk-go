package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptRemediationSummary the number of device health scripts deployed and the number of devices the scripts remediated.
type DeviceHealthScriptRemediationSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The number of devices remediated by device health scripts.
    remediatedDeviceCount *int32
    // The number of device health scripts deployed.
    scriptCount *int32
}
// NewDeviceHealthScriptRemediationSummary instantiates a new deviceHealthScriptRemediationSummary and sets the default values.
func NewDeviceHealthScriptRemediationSummary()(*DeviceHealthScriptRemediationSummary) {
    m := &DeviceHealthScriptRemediationSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceHealthScriptRemediationSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptRemediationSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptRemediationSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptRemediationSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptRemediationSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["remediatedDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRemediatedDeviceCount)
    res["scriptCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetScriptCount)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptRemediationSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetRemediatedDeviceCount gets the remediatedDeviceCount property value. The number of devices remediated by device health scripts.
func (m *DeviceHealthScriptRemediationSummary) GetRemediatedDeviceCount()(*int32) {
    return m.remediatedDeviceCount
}
// GetScriptCount gets the scriptCount property value. The number of device health scripts deployed.
func (m *DeviceHealthScriptRemediationSummary) GetScriptCount()(*int32) {
    return m.scriptCount
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptRemediationSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("remediatedDeviceCount", m.GetRemediatedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("scriptCount", m.GetScriptCount())
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
func (m *DeviceHealthScriptRemediationSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptRemediationSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRemediatedDeviceCount sets the remediatedDeviceCount property value. The number of devices remediated by device health scripts.
func (m *DeviceHealthScriptRemediationSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
// SetScriptCount sets the scriptCount property value. The number of device health scripts deployed.
func (m *DeviceHealthScriptRemediationSummary) SetScriptCount(value *int32)() {
    m.scriptCount = value
}
