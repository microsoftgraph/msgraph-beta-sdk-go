package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TiIndicatorsRequestBuilder provides operations to manage the tiIndicators property of the microsoft.graph.security entity.
type TiIndicatorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TiIndicatorsRequestBuilderGetQueryParameters get tiIndicators from security
type TiIndicatorsRequestBuilderGetQueryParameters struct {
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
// TiIndicatorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TiIndicatorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TiIndicatorsRequestBuilderGetQueryParameters
}
// TiIndicatorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TiIndicatorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTiIndicatorId provides operations to manage the tiIndicators property of the microsoft.graph.security entity.
// returns a *TiIndicatorsTiIndicatorItemRequestBuilder when successful
func (m *TiIndicatorsRequestBuilder) ByTiIndicatorId(tiIndicatorId string)(*TiIndicatorsTiIndicatorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if tiIndicatorId != "" {
        urlTplParams["tiIndicator%2Did"] = tiIndicatorId
    }
    return NewTiIndicatorsTiIndicatorItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTiIndicatorsRequestBuilderInternal instantiates a new TiIndicatorsRequestBuilder and sets the default values.
func NewTiIndicatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiIndicatorsRequestBuilder) {
    m := &TiIndicatorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/tiIndicators{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTiIndicatorsRequestBuilder instantiates a new TiIndicatorsRequestBuilder and sets the default values.
func NewTiIndicatorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiIndicatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTiIndicatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TiIndicatorsCountRequestBuilder when successful
func (m *TiIndicatorsRequestBuilder) Count()(*TiIndicatorsCountRequestBuilder) {
    return NewTiIndicatorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeleteTiIndicators provides operations to call the deleteTiIndicators method.
// returns a *TiIndicatorsDeleteTiIndicatorsRequestBuilder when successful
func (m *TiIndicatorsRequestBuilder) DeleteTiIndicators()(*TiIndicatorsDeleteTiIndicatorsRequestBuilder) {
    return NewTiIndicatorsDeleteTiIndicatorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeleteTiIndicatorsByExternalId provides operations to call the deleteTiIndicatorsByExternalId method.
// returns a *TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder when successful
func (m *TiIndicatorsRequestBuilder) DeleteTiIndicatorsByExternalId()(*TiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilder) {
    return NewTiIndicatorsDeleteTiIndicatorsByExternalIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get tiIndicators from security
// returns a TiIndicatorCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TiIndicatorsRequestBuilder) Get(ctx context.Context, requestConfiguration *TiIndicatorsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTiIndicatorCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorCollectionResponseable), nil
}
// Post create new navigation property to tiIndicators for security
// returns a TiIndicatorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TiIndicatorsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorable, requestConfiguration *TiIndicatorsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTiIndicatorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorable), nil
}
// SubmitTiIndicators provides operations to call the submitTiIndicators method.
// returns a *TiIndicatorsSubmitTiIndicatorsRequestBuilder when successful
func (m *TiIndicatorsRequestBuilder) SubmitTiIndicators()(*TiIndicatorsSubmitTiIndicatorsRequestBuilder) {
    return NewTiIndicatorsSubmitTiIndicatorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get tiIndicators from security
// returns a *RequestInformation when successful
func (m *TiIndicatorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TiIndicatorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to tiIndicators for security
// returns a *RequestInformation when successful
func (m *TiIndicatorsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorable, requestConfiguration *TiIndicatorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdateTiIndicators provides operations to call the updateTiIndicators method.
// returns a *TiIndicatorsUpdateTiIndicatorsRequestBuilder when successful
func (m *TiIndicatorsRequestBuilder) UpdateTiIndicators()(*TiIndicatorsUpdateTiIndicatorsRequestBuilder) {
    return NewTiIndicatorsUpdateTiIndicatorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TiIndicatorsRequestBuilder when successful
func (m *TiIndicatorsRequestBuilder) WithUrl(rawUrl string)(*TiIndicatorsRequestBuilder) {
    return NewTiIndicatorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
