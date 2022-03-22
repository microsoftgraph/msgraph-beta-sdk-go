package gethealthmetrics

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetHealthMetricsRequestBody provides operations to call the getHealthMetrics method.
type GetHealthMetricsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    metricNames []string;
}
// NewGetHealthMetricsRequestBody instantiates a new getHealthMetricsRequestBody and sets the default values.
func NewGetHealthMetricsRequestBody()(*GetHealthMetricsRequestBody) {
    m := &GetHealthMetricsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetHealthMetricsRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetHealthMetricsRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGetHealthMetricsRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetHealthMetricsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetHealthMetricsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["metricNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMetricNames(res)
        }
        return nil
    }
    return res
}
// GetMetricNames gets the metricNames property value. 
func (m *GetHealthMetricsRequestBody) GetMetricNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.metricNames
    }
}
// Serialize serializes information the current object
func (m *GetHealthMetricsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetMetricNames() != nil {
        err := writer.WriteCollectionOfStringValues("metricNames", m.GetMetricNames())
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
func (m *GetHealthMetricsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMetricNames sets the metricNames property value. 
func (m *GetHealthMetricsRequestBody) SetMetricNames(value []string)() {
    if m != nil {
        m.metricNames = value
    }
}
