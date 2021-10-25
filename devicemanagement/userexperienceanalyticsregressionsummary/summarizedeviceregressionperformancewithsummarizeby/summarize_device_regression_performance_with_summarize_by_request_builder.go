package summarizedeviceregressionperformancewithsummarizeby

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SummarizeDeviceRegressionPerformanceWithSummarizeByResponse struct {
    additionalData map[string]interface{};
    userExperienceAnalyticsRegressionSummary *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsRegressionSummary;
}
func NewSummarizeDeviceRegressionPerformanceWithSummarizeByResponse()(*SummarizeDeviceRegressionPerformanceWithSummarizeByResponse) {
    m := &SummarizeDeviceRegressionPerformanceWithSummarizeByResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByResponse) GetUserExperienceAnalyticsRegressionSummary()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsRegressionSummary) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsRegressionSummary
    }
}
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["userExperienceAnalyticsRegressionSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsRegressionSummary() })
        if err != nil {
            return err
        }
        m.SetUserExperienceAnalyticsRegressionSummary(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsRegressionSummary))
        return nil
    }
    return res
}
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByResponse) IsNil()(bool) {
    return m == nil
}
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("userExperienceAnalyticsRegressionSummary", m.GetUserExperienceAnalyticsRegressionSummary())
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
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByResponse) SetUserExperienceAnalyticsRegressionSummary(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsRegressionSummary)() {
    m.userExperienceAnalyticsRegressionSummary = value
}
func NewSummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, summarizeBy *string)(*SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder) {
    m := &SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/userExperienceAnalyticsRegressionSummary/microsoft.graph.summarizeDeviceRegressionPerformance(summarizeBy={summarizeBy})";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    if summarizeBy != nil {
        urlTplParams["summarizeBy"] = *summarizeBy
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewSummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilderInternal(urlParams, requestAdapter, nil)
}
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder) CreateGetRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *SummarizeDeviceRegressionPerformanceWithSummarizeByRequestBuilder) Get(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*SummarizeDeviceRegressionPerformanceWithSummarizeByResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSummarizeDeviceRegressionPerformanceWithSummarizeByResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*SummarizeDeviceRegressionPerformanceWithSummarizeByResponse), nil
}
