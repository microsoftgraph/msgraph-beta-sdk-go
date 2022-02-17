package getexpiringvpptokencountwithexpiringbeforedatetime

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder builds and executes requests for operations under \deviceManagement\depOnboardingSettings\microsoft.graph.getExpiringVppTokenCount(expiringBeforeDateTime='{expiringBeforeDateTime}')
type GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetOptions options for Get
type GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal instantiates a new GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder and sets the default values.
func NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, expiringBeforeDateTime *string)(*GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    m := &GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/microsoft.graph.getExpiringVppTokenCount(expiringBeforeDateTime='{expiringBeforeDateTime}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if expiringBeforeDateTime != nil {
        urlTplParams["expiringBeforeDateTime"] = *expiringBeforeDateTime
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder instantiates a new GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder and sets the default values.
func NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getExpiringVppTokenCount
func (m *GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) CreateGetRequestInformation(options *GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function getExpiringVppTokenCount
func (m *GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) Get(options *GetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetOptions)(*int32, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "int32", nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*int32), nil
}
