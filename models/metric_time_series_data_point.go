package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MetricTimeSeriesDataPoint metric Time series data point
type MetricTimeSeriesDataPoint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Time of the metric time series data point
    dateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // Value of the metric time series data point
    value *int64
}
// NewMetricTimeSeriesDataPoint instantiates a new metricTimeSeriesDataPoint and sets the default values.
func NewMetricTimeSeriesDataPoint()(*MetricTimeSeriesDataPoint) {
    m := &MetricTimeSeriesDataPoint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.metricTimeSeriesDataPoint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMetricTimeSeriesDataPointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMetricTimeSeriesDataPointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMetricTimeSeriesDataPoint(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MetricTimeSeriesDataPoint) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDateTime gets the dateTime property value. Time of the metric time series data point
func (m *MetricTimeSeriesDataPoint) GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MetricTimeSeriesDataPoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetDateTime)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetValue)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MetricTimeSeriesDataPoint) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. Value of the metric time series data point
func (m *MetricTimeSeriesDataPoint) GetValue()(*int64) {
    return m.value
}
// Serialize serializes information the current object
func (m *MetricTimeSeriesDataPoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("dateTime", m.GetDateTime())
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
        err := writer.WriteInt64Value("value", m.GetValue())
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
func (m *MetricTimeSeriesDataPoint) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDateTime sets the dateTime property value. Time of the metric time series data point
func (m *MetricTimeSeriesDataPoint) SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MetricTimeSeriesDataPoint) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. Value of the metric time series data point
func (m *MetricTimeSeriesDataPoint) SetValue(value *int64)() {
    m.value = value
}
