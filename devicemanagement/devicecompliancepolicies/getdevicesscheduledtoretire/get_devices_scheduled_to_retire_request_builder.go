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
// GetDevicesScheduledToRetireRequestBuilderPostOptions options for Post
type GetDevicesScheduledToRetireRequestBuilderPostOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
// CreatePostRequestInformation invoke action getDevicesScheduledToRetire
func (m *GetDevicesScheduledToRetireRequestBuilder) CreatePostRequestInformation(options *GetDevicesScheduledToRetireRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
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
