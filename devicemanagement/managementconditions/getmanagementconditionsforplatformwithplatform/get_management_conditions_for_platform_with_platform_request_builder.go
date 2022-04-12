package getmanagementconditionsforplatformwithplatform

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetManagementConditionsForPlatformWithPlatformRequestBuilder provides operations to call the getManagementConditionsForPlatform method.
type GetManagementConditionsForPlatformWithPlatformRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetManagementConditionsForPlatformWithPlatformRequestBuilderGetOptions options for Get
type GetManagementConditionsForPlatformWithPlatformRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetManagementConditionsForPlatformWithPlatformRequestBuilderInternal instantiates a new GetManagementConditionsForPlatformWithPlatformRequestBuilder and sets the default values.
func NewGetManagementConditionsForPlatformWithPlatformRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, platform *string)(*GetManagementConditionsForPlatformWithPlatformRequestBuilder) {
    m := &GetManagementConditionsForPlatformWithPlatformRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/managementConditions/microsoft.graph.getManagementConditionsForPlatform(platform='{platform}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if platform != nil {
        urlTplParams[""] = *platform
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetManagementConditionsForPlatformWithPlatformRequestBuilder instantiates a new GetManagementConditionsForPlatformWithPlatformRequestBuilder and sets the default values.
func NewGetManagementConditionsForPlatformWithPlatformRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetManagementConditionsForPlatformWithPlatformRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetManagementConditionsForPlatformWithPlatformRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getManagementConditionsForPlatform
func (m *GetManagementConditionsForPlatformWithPlatformRequestBuilder) CreateGetRequestInformation(options *GetManagementConditionsForPlatformWithPlatformRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
// Get invoke function getManagementConditionsForPlatform
func (m *GetManagementConditionsForPlatformWithPlatformRequestBuilder) Get(options *GetManagementConditionsForPlatformWithPlatformRequestBuilderGetOptions)(GetManagementConditionsForPlatformWithPlatformResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetManagementConditionsForPlatformWithPlatformResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetManagementConditionsForPlatformWithPlatformResponseable), nil
}
