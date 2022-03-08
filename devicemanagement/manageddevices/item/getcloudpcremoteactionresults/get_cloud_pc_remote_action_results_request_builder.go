package getcloudpcremoteactionresults

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetCloudPcRemoteActionResultsRequestBuilder provides operations to call the getCloudPcRemoteActionResults method.
type GetCloudPcRemoteActionResultsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetCloudPcRemoteActionResultsRequestBuilderGetOptions options for Get
type GetCloudPcRemoteActionResultsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetCloudPcRemoteActionResultsRequestBuilderInternal instantiates a new GetCloudPcRemoteActionResultsRequestBuilder and sets the default values.
func NewGetCloudPcRemoteActionResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetCloudPcRemoteActionResultsRequestBuilder) {
    m := &GetCloudPcRemoteActionResultsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managedDevices/{managedDevice_id}/microsoft.graph.getCloudPcRemoteActionResults()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetCloudPcRemoteActionResultsRequestBuilder instantiates a new GetCloudPcRemoteActionResultsRequestBuilder and sets the default values.
func NewGetCloudPcRemoteActionResultsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetCloudPcRemoteActionResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetCloudPcRemoteActionResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getCloudPcRemoteActionResults
func (m *GetCloudPcRemoteActionResultsRequestBuilder) CreateGetRequestInformation(options *GetCloudPcRemoteActionResultsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getCloudPcRemoteActionResults
func (m *GetCloudPcRemoteActionResultsRequestBuilder) Get(options *GetCloudPcRemoteActionResultsRequestBuilderGetOptions)(GetCloudPcRemoteActionResultsResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetCloudPcRemoteActionResultsResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetCloudPcRemoteActionResultsResponseable), nil
}
