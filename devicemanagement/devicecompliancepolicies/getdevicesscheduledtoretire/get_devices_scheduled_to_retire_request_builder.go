package getdevicesscheduledtoretire

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetDevicesScheduledToRetireRequestBuilder provides operations to call the getDevicesScheduledToRetire method.
type GetDevicesScheduledToRetireRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetDevicesScheduledToRetireRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetDevicesScheduledToRetireRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetDevicesScheduledToRetireRequestBuilderInternal instantiates a new GetDevicesScheduledToRetireRequestBuilder and sets the default values.
func NewGetDevicesScheduledToRetireRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetDevicesScheduledToRetireRequestBuilder) {
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
func NewGetDevicesScheduledToRetireRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetDevicesScheduledToRetireRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetDevicesScheduledToRetireRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getDevicesScheduledToRetire
func (m *GetDevicesScheduledToRetireRequestBuilder) CreatePostRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getDevicesScheduledToRetire
func (m *GetDevicesScheduledToRetireRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *GetDevicesScheduledToRetireRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// PostWithResponseHandler invoke action getDevicesScheduledToRetire
func (m *GetDevicesScheduledToRetireRequestBuilder) PostWithResponseHandler(requestConfiguration *GetDevicesScheduledToRetireRequestBuilderPostRequestConfiguration)(GetDevicesScheduledToRetireResponseable, error) {
    return m.PostWithResponseHandler(requestConfiguration, nil);
}
// PostWithResponseHandler invoke action getDevicesScheduledToRetire
func (m *GetDevicesScheduledToRetireRequestBuilder) PostWithResponseHandler(requestConfiguration *GetDevicesScheduledToRetireRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetDevicesScheduledToRetireResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetDevicesScheduledToRetireResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetDevicesScheduledToRetireResponseable), nil
}
