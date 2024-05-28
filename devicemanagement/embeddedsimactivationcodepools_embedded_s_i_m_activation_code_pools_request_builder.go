package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
type EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderGetQueryParameters the embedded SIM activation code pools created by this account.
type EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderGetQueryParameters struct {
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
// EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderGetQueryParameters
}
// EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEmbeddedSIMActivationCodePoolId provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
// returns a *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder) ByEmbeddedSIMActivationCodePoolId(embeddedSIMActivationCodePoolId string)(*EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if embeddedSIMActivationCodePoolId != "" {
        urlTplParams["embeddedSIMActivationCodePool%2Did"] = embeddedSIMActivationCodePoolId
    }
    return NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderInternal instantiates a new EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder and sets the default values.
func NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder) {
    m := &EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/embeddedSIMActivationCodePools{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder instantiates a new EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder and sets the default values.
func NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EmbeddedsimactivationcodepoolsCountRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder) Count()(*EmbeddedsimactivationcodepoolsCountRequestBuilder) {
    return NewEmbeddedsimactivationcodepoolsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the embedded SIM activation code pools created by this account.
// returns a EmbeddedSIMActivationCodePoolCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder) Get(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMActivationCodePoolCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolCollectionResponseable), nil
}
// Post create new navigation property to embeddedSIMActivationCodePools for deviceManagement
// returns a EmbeddedSIMActivationCodePoolable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, requestConfiguration *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable), nil
}
// ToGetRequestInformation the embedded SIM activation code pools created by this account.
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to embeddedSIMActivationCodePools for deviceManagement
// returns a *RequestInformation when successful
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMActivationCodePoolable, requestConfiguration *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder when successful
func (m *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder) WithUrl(rawUrl string)(*EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder) {
    return NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
