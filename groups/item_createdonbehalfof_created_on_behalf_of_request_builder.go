package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder provides operations to manage the createdOnBehalfOf property of the microsoft.graph.group entity.
type ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilderGetQueryParameters the user (or application) that created the group. Note: This isn't set if the user is an administrator. Read-only.
type ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilderGetQueryParameters
}
// NewItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilderInternal instantiates a new ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder and sets the default values.
func NewItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder) {
    m := &ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/createdOnBehalfOf{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder instantiates a new ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder and sets the default values.
func NewItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the user (or application) that created the group. Note: This isn't set if the user is an administrator. Read-only.
// returns a DirectoryObjectable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
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
// ToGetRequestInformation the user (or application) that created the group. Note: This isn't set if the user is an administrator. Read-only.
// returns a *RequestInformation when successful
func (m *ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder when successful
func (m *ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder) WithUrl(rawUrl string)(*ItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder) {
    return NewItemCreatedonbehalfofCreatedOnBehalfOfRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
