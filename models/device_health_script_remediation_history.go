package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
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
}
// NewDeviceHealthScriptRemediationHistory instantiates a new DeviceHealthScriptRemediationHistory and sets the default values.
func NewDeviceHealthScriptRemediationHistory()(*DeviceHealthScriptRemediationHistory) {
    m := &DeviceHealthScriptRemediationHistory{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceHealthScriptRemediationHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptRemediationHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptRemediationHistory(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptRemediationHistory) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptRemediationHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["historyData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceHealthScriptRemediationHistoryDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptRemediationHistoryDataable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceHealthScriptRemediationHistoryDataable)
            }
            m.SetHistoryData(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
// GetHistoryData gets the historyData property value. The number of devices remediated by the device health script on the given date.
func (m *DeviceHealthScriptRemediationHistory) GetHistoryData()([]DeviceHealthScriptRemediationHistoryDataable) {
    if m == nil {
        return nil
    } else {
        return m.historyData
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date on which the results history is calculated for the healthscript.
func (m *DeviceHealthScriptRemediationHistory) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptRemediationHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetHistoryData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHistoryData()))
        for i, v := range m.GetHistoryData() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptRemediationHistory) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHistoryData sets the historyData property value. The number of devices remediated by the device health script on the given date.
func (m *DeviceHealthScriptRemediationHistory) SetHistoryData(value []DeviceHealthScriptRemediationHistoryDataable)() {
    if m != nil {
        m.historyData = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date on which the results history is calculated for the healthscript.
func (m *DeviceHealthScriptRemediationHistory) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
