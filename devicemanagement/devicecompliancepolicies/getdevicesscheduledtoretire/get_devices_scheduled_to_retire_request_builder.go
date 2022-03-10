package getdevicesscheduledtoretire

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetDevicesScheduledToRetireRequestBuilder provides operations to call the getDevicesScheduledToRetire method.
type GetDevicesScheduledToRetireRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetDevicesScheduledToRetireRequestBuilderPostOptions options for Post
type GetDevicesScheduledToRetireRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetDevicesScheduledToRetireRequestBuilderInternal instantiates a new GetDevicesScheduledToRetireRequestBuilder and sets the default values.
func NewGetDevicesScheduledToRetireRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetDevicesScheduledToRetireRequestBuilder) {
    m := &GetDevicesScheduledToRetireRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/microsoft.graph.getDevicesScheduledToRetire";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetDevicesScheduledToRetireRequestBuilder instantiates a new GetDevicesScheduledToRetireRequestBuilder and sets the default values.
func NewGetDevicesScheduledToRetireRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetDevicesScheduledToRetireRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetDevicesScheduledToRetireRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getDevicesScheduledToRetire
func (m *GetDevicesScheduledToRetireRequestBuilder) CreatePostRequestInformation(options *GetDevicesScheduledToRetireRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Post invoke action getDevicesScheduledToRetire
func (m *GetDevicesScheduledToRetireRequestBuilder) Post(options *GetDevicesScheduledToRetireRequestBuilderPostOptions)(GetDevicesScheduledToRetireResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetDevicesScheduledToRetireResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetDevicesScheduledToRetireResponseable), nil
}
