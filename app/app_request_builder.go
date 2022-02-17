package app

import (
    i950c82ce9946e255745879c43558e6119ec2c2d7f73ff599a8825a675691d160 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings"
    ibcd7ebbfadd792f3a3f918631f2d9771cd65cc9b81894406f89550a45009a90f "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0f59cb615d0a8fdb53a9db31bee9b4e4081ab7f69d9a6bbba5839e7066dfb903 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ibe96689de55a51f0c56d2a0e70225b89a766a6c4cdc8d088380eb648b999dd3f "github.com/microsoftgraph/msgraph-beta-sdk-go/app/onlinemeetings/item"
)

// AppRequestBuilder builds and executes requests for operations under \app
type AppRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AppRequestBuilderGetOptions options for Get
type AppRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AppRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AppRequestBuilderGetQueryParameters get app
type AppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AppRequestBuilderPatchOptions options for Patch
type AppRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CommsApplication;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AppRequestBuilder) Calls()(*ibcd7ebbfadd792f3a3f918631f2d9771cd65cc9b81894406f89550a45009a90f.CallsRequestBuilder) {
    return ibcd7ebbfadd792f3a3f918631f2d9771cd65cc9b81894406f89550a45009a90f.NewCallsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CallsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.calls.item collection
func (m *AppRequestBuilder) CallsById(id string)(*i0f59cb615d0a8fdb53a9db31bee9b4e4081ab7f69d9a6bbba5839e7066dfb903.CallRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["call_id"] = id
    }
    return i0f59cb615d0a8fdb53a9db31bee9b4e4081ab7f69d9a6bbba5839e7066dfb903.NewCallRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAppRequestBuilderInternal instantiates a new AppRequestBuilder and sets the default values.
func NewAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppRequestBuilder) {
    m := &AppRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAppRequestBuilder instantiates a new AppRequestBuilder and sets the default values.
func NewAppRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get app
func (m *AppRequestBuilder) CreateGetRequestInformation(options *AppRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
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
// CreatePatchRequestInformation update app
func (m *AppRequestBuilder) CreatePatchRequestInformation(options *AppRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Get get app
func (m *AppRequestBuilder) Get(options *AppRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CommsApplication, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCommsApplication() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CommsApplication), nil
}
func (m *AppRequestBuilder) OnlineMeetings()(*i950c82ce9946e255745879c43558e6119ec2c2d7f73ff599a8825a675691d160.OnlineMeetingsRequestBuilder) {
    return i950c82ce9946e255745879c43558e6119ec2c2d7f73ff599a8825a675691d160.NewOnlineMeetingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnlineMeetingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.onlineMeetings.item collection
func (m *AppRequestBuilder) OnlineMeetingsById(id string)(*ibe96689de55a51f0c56d2a0e70225b89a766a6c4cdc8d088380eb648b999dd3f.OnlineMeetingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onlineMeeting_id"] = id
    }
    return ibe96689de55a51f0c56d2a0e70225b89a766a6c4cdc8d088380eb648b999dd3f.NewOnlineMeetingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update app
func (m *AppRequestBuilder) Patch(options *AppRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
