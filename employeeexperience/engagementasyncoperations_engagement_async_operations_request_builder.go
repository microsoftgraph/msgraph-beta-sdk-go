package employeeexperience

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder provides operations to manage the engagementAsyncOperations property of the microsoft.graph.employeeExperience entity.
type EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EngagementasyncoperationsEngagementAsyncOperationsRequestBuilderGetQueryParameters get an engagementAsyncOperation to track a long-running operation request.
type EngagementasyncoperationsEngagementAsyncOperationsRequestBuilderGetQueryParameters struct {
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
// EngagementasyncoperationsEngagementAsyncOperationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EngagementasyncoperationsEngagementAsyncOperationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilderGetQueryParameters
}
// EngagementasyncoperationsEngagementAsyncOperationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EngagementasyncoperationsEngagementAsyncOperationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEngagementAsyncOperationId provides operations to manage the engagementAsyncOperations property of the microsoft.graph.employeeExperience entity.
// returns a *EngagementasyncoperationsEngagementAsyncOperationItemRequestBuilder when successful
func (m *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder) ByEngagementAsyncOperationId(engagementAsyncOperationId string)(*EngagementasyncoperationsEngagementAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if engagementAsyncOperationId != "" {
        urlTplParams["engagementAsyncOperation%2Did"] = engagementAsyncOperationId
    }
    return NewEngagementasyncoperationsEngagementAsyncOperationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEngagementasyncoperationsEngagementAsyncOperationsRequestBuilderInternal instantiates a new EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder and sets the default values.
func NewEngagementasyncoperationsEngagementAsyncOperationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder) {
    m := &EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/employeeExperience/engagementAsyncOperations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEngagementasyncoperationsEngagementAsyncOperationsRequestBuilder instantiates a new EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder and sets the default values.
func NewEngagementasyncoperationsEngagementAsyncOperationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEngagementasyncoperationsEngagementAsyncOperationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EngagementasyncoperationsCountRequestBuilder when successful
func (m *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder) Count()(*EngagementasyncoperationsCountRequestBuilder) {
    return NewEngagementasyncoperationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get an engagementAsyncOperation to track a long-running operation request.
// returns a EngagementAsyncOperationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder) Get(ctx context.Context, requestConfiguration *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EngagementAsyncOperationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEngagementAsyncOperationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EngagementAsyncOperationCollectionResponseable), nil
}
// Post create new navigation property to engagementAsyncOperations for employeeExperience
// returns a EngagementAsyncOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EngagementAsyncOperationable, requestConfiguration *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EngagementAsyncOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEngagementAsyncOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EngagementAsyncOperationable), nil
}
// ToGetRequestInformation get an engagementAsyncOperation to track a long-running operation request.
// returns a *RequestInformation when successful
func (m *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to engagementAsyncOperations for employeeExperience
// returns a *RequestInformation when successful
func (m *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EngagementAsyncOperationable, requestConfiguration *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder when successful
func (m *EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder) WithUrl(rawUrl string)(*EngagementasyncoperationsEngagementAsyncOperationsRequestBuilder) {
    return NewEngagementasyncoperationsEngagementAsyncOperationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
