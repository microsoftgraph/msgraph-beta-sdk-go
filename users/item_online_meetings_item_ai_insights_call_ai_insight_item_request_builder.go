package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder provides operations to manage the aiInsights property of the microsoft.graph.onlineMeeting entity.
type ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderGetQueryParameters get aiInsights from users
type ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderGetQueryParameters
}
// ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderInternal instantiates a new ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder and sets the default values.
func NewItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder) {
    m := &ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/aiInsights/{callAiInsight%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder instantiates a new ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder and sets the default values.
func NewItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property aiInsights for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get aiInsights from users
// returns a CallAiInsightable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallAiInsightable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCallAiInsightFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallAiInsightable), nil
}
// Patch update the navigation property aiInsights in users
// returns a CallAiInsightable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallAiInsightable, requestConfiguration *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallAiInsightable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCallAiInsightFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallAiInsightable), nil
}
// ToDeleteRequestInformation delete navigation property aiInsights for users
// returns a *RequestInformation when successful
func (m *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get aiInsights from users
// returns a *RequestInformation when successful
func (m *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property aiInsights in users
// returns a *RequestInformation when successful
func (m *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallAiInsightable, requestConfiguration *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder when successful
func (m *ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder) WithUrl(rawUrl string)(*ItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder) {
    return NewItemOnlineMeetingsItemAiInsightsCallAiInsightItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
