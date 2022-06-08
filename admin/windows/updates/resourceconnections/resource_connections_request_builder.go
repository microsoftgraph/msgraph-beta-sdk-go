package resourceconnections

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic5008f2ec9f3804779bca64c17296c9404282af98339cd88a74c5f55f7d7008d "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/windows/updates/resourceconnections/count"
)

// ResourceConnectionsRequestBuilder provides operations to manage the resourceConnections property of the microsoft.graph.windowsUpdates.updates entity.
type ResourceConnectionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ResourceConnectionsRequestBuilderGetQueryParameters get resourceConnections from admin
type ResourceConnectionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ResourceConnectionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceConnectionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ResourceConnectionsRequestBuilderGetQueryParameters
}
// ResourceConnectionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ResourceConnectionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewResourceConnectionsRequestBuilderInternal instantiates a new ResourceConnectionsRequestBuilder and sets the default values.
func NewResourceConnectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceConnectionsRequestBuilder) {
    m := &ResourceConnectionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/resourceConnections{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewResourceConnectionsRequestBuilder instantiates a new ResourceConnectionsRequestBuilder and sets the default values.
func NewResourceConnectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResourceConnectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceConnectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *ResourceConnectionsRequestBuilder) Count()(*ic5008f2ec9f3804779bca64c17296c9404282af98339cd88a74c5f55f7d7008d.CountRequestBuilder) {
    return ic5008f2ec9f3804779bca64c17296c9404282af98339cd88a74c5f55f7d7008d.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get resourceConnections from admin
func (m *ResourceConnectionsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get resourceConnections from admin
func (m *ResourceConnectionsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ResourceConnectionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to resourceConnections for admin
func (m *ResourceConnectionsRequestBuilder) CreatePostRequestInformation(body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ResourceConnectionable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to resourceConnections for admin
func (m *ResourceConnectionsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ResourceConnectionable, requestConfiguration *ResourceConnectionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get resourceConnections from admin
func (m *ResourceConnectionsRequestBuilder) Get()(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ResourceConnectionCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get resourceConnections from admin
func (m *ResourceConnectionsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ResourceConnectionsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ResourceConnectionCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateResourceConnectionCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ResourceConnectionCollectionResponseable), nil
}
// Post create new navigation property to resourceConnections for admin
func (m *ResourceConnectionsRequestBuilder) Post(body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ResourceConnectionable)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ResourceConnectionable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to resourceConnections for admin
func (m *ResourceConnectionsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ResourceConnectionable, requestConfiguration *ResourceConnectionsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ResourceConnectionable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateResourceConnectionFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ResourceConnectionable), nil
}
