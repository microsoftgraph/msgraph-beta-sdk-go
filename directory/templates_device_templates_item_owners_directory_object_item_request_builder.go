package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder provides operations to manage the owners property of the microsoft.graph.deviceTemplate entity.
type TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilderGetQueryParameters get owners from directory
type TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilderGetQueryParameters
}
// NewTemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder and sets the default values.
func NewTemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder) {
    m := &TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/templates/deviceTemplates/{deviceTemplate%2Did}/owners/{directoryObject%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder instantiates a new TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder and sets the default values.
func NewTemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get owners from directory
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// ToGetRequestInformation get owners from directory
// returns a *RequestInformation when successful
func (m *TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder when successful
func (m *TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder) WithUrl(rawUrl string)(*TemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder) {
    return NewTemplatesDeviceTemplatesItemOwnersDirectoryObjectItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
