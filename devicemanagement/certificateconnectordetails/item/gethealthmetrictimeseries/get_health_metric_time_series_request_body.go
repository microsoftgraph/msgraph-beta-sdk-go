package gethealthmetrictimeseries

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetHealthMetricTimeSeriesRequestBody provides operations to call the getHealthMetricTimeSeries method.
type GetHealthMetricTimeSeriesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    timeSeries i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeSeriesParameterable;
}
// NewGetHealthMetricTimeSeriesRequestBody instantiates a new getHealthMetricTimeSeriesRequestBody and sets the default values.
func NewGetHealthMetricTimeSeriesRequestBody()(*GetHealthMetricTimeSeriesRequestBody) {
    m := &GetHealthMetricTimeSeriesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetHealthMetricTimeSeriesRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetHealthMetricTimeSeriesRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGetHealthMetricTimeSeriesRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetHealthMetricTimeSeriesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetHealthMetricTimeSeriesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["timeSeries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateTimeSeriesParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeSeries(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeSeriesParameterable))
        }
        return nil
    }
    return res
}
// GetTimeSeries gets the timeSeries property value. 
func (m *GetHealthMetricTimeSeriesRequestBody) GetTimeSeries()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeSeriesParameterable) {
    if m == nil {
        return nil
    } else {
        return m.timeSeries
    }
}
// Serialize serializes information the current object
func (m *GetHealthMetricTimeSeriesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("timeSeries", m.GetTimeSeries())
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
func (m *GetHealthMetricTimeSeriesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTimeSeries sets the timeSeries property value. 
func (m *GetHealthMetricTimeSeriesRequestBody) SetTimeSeries(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeSeriesParameterable)() {
    if m != nil {
        m.timeSeries = value
    }
}
