package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder provides operations to call the getCloudPcRemoteActionResults method.
type ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters invoke function getCloudPcRemoteActionResults
type ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters
}
// NewComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal instantiates a new ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder and sets the default values.
func NewComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    m := &ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/getCloudPcRemoteActionResults(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder instantiates a new ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder and sets the default values.
func NewComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getCloudPcRemoteActionResults
// Deprecated: This method is obsolete. Use GetAsGetCloudPcRemoteActionResultsGetResponse instead.
// returns a ComanagedDevicesItemGetCloudPcRemoteActionResultsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(ComanagedDevicesItemGetCloudPcRemoteActionResultsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateComanagedDevicesItemGetCloudPcRemoteActionResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ComanagedDevicesItemGetCloudPcRemoteActionResultsResponseable), nil
}
// GetAsGetCloudPcRemoteActionResultsGetResponse invoke function getCloudPcRemoteActionResults
// returns a ComanagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) GetAsGetCloudPcRemoteActionResultsGetResponse(ctx context.Context, requestConfiguration *ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(ComanagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateComanagedDevicesItemGetCloudPcRemoteActionResultsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ComanagedDevicesItemGetCloudPcRemoteActionResultsGetResponseable), nil
}
// ToGetRequestInformation invoke function getCloudPcRemoteActionResults
// returns a *RequestInformation when successful
func (m *ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder when successful
func (m *ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    return NewComanagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
