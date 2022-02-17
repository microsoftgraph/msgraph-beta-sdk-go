package getmobileappcountwithstatus

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetMobileAppCountWithStatusRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps\microsoft.graph.getMobileAppCount(status='{status}')
type GetMobileAppCountWithStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetMobileAppCountWithStatusRequestBuilderGetOptions options for Get
type GetMobileAppCountWithStatusRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetMobileAppCountWithStatusRequestBuilderInternal instantiates a new GetMobileAppCountWithStatusRequestBuilder and sets the default values.
func NewGetMobileAppCountWithStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, status *string)(*GetMobileAppCountWithStatusRequestBuilder) {
    m := &GetMobileAppCountWithStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/microsoft.graph.getMobileAppCount(status='{status}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if status != nil {
        urlTplParams["status"] = *status
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetMobileAppCountWithStatusRequestBuilder instantiates a new GetMobileAppCountWithStatusRequestBuilder and sets the default values.
func NewGetMobileAppCountWithStatusRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetMobileAppCountWithStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetMobileAppCountWithStatusRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getMobileAppCount
func (m *GetMobileAppCountWithStatusRequestBuilder) CreateGetRequestInformation(options *GetMobileAppCountWithStatusRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getMobileAppCount
func (m *GetMobileAppCountWithStatusRequestBuilder) Get(options *GetMobileAppCountWithStatusRequestBuilderGetOptions)(*int64, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "int64", nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*int64), nil
}
