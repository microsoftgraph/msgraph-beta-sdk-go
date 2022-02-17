package gettopmobileappswithstatuswithcount

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetTopMobileAppsWithStatusWithCountRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps\microsoft.graph.getTopMobileApps(status='{status}',count={count})
type GetTopMobileAppsWithStatusWithCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetTopMobileAppsWithStatusWithCountRequestBuilderGetOptions options for Get
type GetTopMobileAppsWithStatusWithCountRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetTopMobileAppsWithStatusWithCountRequestBuilderInternal instantiates a new GetTopMobileAppsWithStatusWithCountRequestBuilder and sets the default values.
func NewGetTopMobileAppsWithStatusWithCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, status *string, count *int64)(*GetTopMobileAppsWithStatusWithCountRequestBuilder) {
    m := &GetTopMobileAppsWithStatusWithCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/microsoft.graph.getTopMobileApps(status='{status}',count={count})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if status != nil {
        urlTplParams["status"] = *status
    }
    if count != nil {
        urlTplParams["count"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(*count, 10)
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetTopMobileAppsWithStatusWithCountRequestBuilder instantiates a new GetTopMobileAppsWithStatusWithCountRequestBuilder and sets the default values.
func NewGetTopMobileAppsWithStatusWithCountRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetTopMobileAppsWithStatusWithCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTopMobileAppsWithStatusWithCountRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// CreateGetRequestInformation invoke function getTopMobileApps
func (m *GetTopMobileAppsWithStatusWithCountRequestBuilder) CreateGetRequestInformation(options *GetTopMobileAppsWithStatusWithCountRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getTopMobileApps
func (m *GetTopMobileAppsWithStatusWithCountRequestBuilder) Get(options *GetTopMobileAppsWithStatusWithCountRequestBuilderGetOptions)([]GetTopMobileAppsWithStatusWithCount, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetTopMobileAppsWithStatusWithCount() }, nil, nil)
    if err != nil {
        return nil, err
    }
    val := make([]GetTopMobileAppsWithStatusWithCount, len(res))
    for i, v := range res {
        val[i] = *(v.(*GetTopMobileAppsWithStatusWithCount))
    }
    return val, nil
}
