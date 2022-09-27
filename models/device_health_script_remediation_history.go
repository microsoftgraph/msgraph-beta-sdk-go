package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptRemediationHistory the number of devices remediated by a device health script on a given date with the last modified time.
type DeviceHealthScriptRemediationHistory struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The number of devices remediated by the device health script on the given date.
    historyData []DeviceHealthScriptRemediationHistoryDataable
    // The date on which the results history is calculated for the healthscript.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
}
// NewDeviceHealthScriptRemediationHistory instantiates a new deviceHealthScriptRemediationHistory and sets the default values.
func NewDeviceHealthScriptRemediationHistory()(*DeviceHealthScriptRemediationHistory) {
    m := &DeviceHealthScriptRemediationHistory{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.deviceHealthScriptRemediationHistory";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceHealthScriptRemediationHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptRemediationHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptRemediationHistory(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptRemediationHistory) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptRemediationHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["historyData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceHealthScriptRemediationHistoryDataFromDiscriminatorValue , m.SetHistoryData)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetHistoryData gets the historyData property value. The number of devices remediated by the device health script on the given date.
func (m *DeviceHealthScriptRemediationHistory) GetHistoryData()([]DeviceHealthScriptRemediationHistoryDataable) {
    return m.historyData
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date on which the results history is calculated for the healthscript.
func (m *DeviceHealthScriptRemediationHistory) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptRemediationHistory) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptRemediationHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetHistoryData() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetHistoryData())
        err := writer.WriteCollectionOfObjectValues("historyData", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
func (m *DeviceHealthScriptRemediationHistory) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetHistoryData sets the historyData property value. The number of devices remediated by the device health script on the given date.
func (m *DeviceHealthScriptRemediationHistory) SetHistoryData(value []DeviceHealthScriptRemediationHistoryDataable)() {
    m.historyData = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date on which the results history is calculated for the healthscript.
func (m *DeviceHealthScriptRemediationHistory) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptRemediationHistory) SetOdataType(value *string)() {
    m.odataType = value
}
