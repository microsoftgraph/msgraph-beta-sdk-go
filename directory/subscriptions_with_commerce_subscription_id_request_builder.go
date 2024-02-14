package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SubscriptionsWithCommerceSubscriptionIdRequestBuilder provides operations to manage the subscriptions property of the microsoft.graph.directory entity.
type SubscriptionsWithCommerceSubscriptionIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SubscriptionsWithCommerceSubscriptionIdRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubscriptionsWithCommerceSubscriptionIdRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SubscriptionsWithCommerceSubscriptionIdRequestBuilderGetQueryParameters get a specific commercial subscription that an organization has acquired.
type SubscriptionsWithCommerceSubscriptionIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SubscriptionsWithCommerceSubscriptionIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubscriptionsWithCommerceSubscriptionIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SubscriptionsWithCommerceSubscriptionIdRequestBuilderGetQueryParameters
}
// SubscriptionsWithCommerceSubscriptionIdRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubscriptionsWithCommerceSubscriptionIdRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSubscriptionsWithCommerceSubscriptionIdRequestBuilderInternal instantiates a new SubscriptionsWithCommerceSubscriptionIdRequestBuilder and sets the default values.
func NewSubscriptionsWithCommerceSubscriptionIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, commerceSubscriptionId *string)(*SubscriptionsWithCommerceSubscriptionIdRequestBuilder) {
    m := &SubscriptionsWithCommerceSubscriptionIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/subscriptions(commerceSubscriptionId='{commerceSubscriptionId}'){?%24expand,%24select}", pathParameters),
    }
    if commerceSubscriptionId != nil {
        m.BaseRequestBuilder.PathParameters["commerceSubscriptionId"] = *commerceSubscriptionId
    }
    return m
}
// NewSubscriptionsWithCommerceSubscriptionIdRequestBuilder instantiates a new SubscriptionsWithCommerceSubscriptionIdRequestBuilder and sets the default values.
func NewSubscriptionsWithCommerceSubscriptionIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubscriptionsWithCommerceSubscriptionIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubscriptionsWithCommerceSubscriptionIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property subscriptions for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SubscriptionsWithCommerceSubscriptionIdRequestBuilder) Delete(ctx context.Context, requestConfiguration *SubscriptionsWithCommerceSubscriptionIdRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get a specific commercial subscription that an organization has acquired.
// returns a CompanySubscriptionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/companysubscription-get?view=graph-rest-1.0
func (m *SubscriptionsWithCommerceSubscriptionIdRequestBuilder) Get(ctx context.Context, requestConfiguration *SubscriptionsWithCommerceSubscriptionIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CompanySubscriptionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCompanySubscriptionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CompanySubscriptionable), nil
}
// Patch update the navigation property subscriptions in directory
// returns a CompanySubscriptionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SubscriptionsWithCommerceSubscriptionIdRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CompanySubscriptionable, requestConfiguration *SubscriptionsWithCommerceSubscriptionIdRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CompanySubscriptionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCompanySubscriptionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CompanySubscriptionable), nil
}
// ToDeleteRequestInformation delete navigation property subscriptions for directory
// returns a *RequestInformation when successful
func (m *SubscriptionsWithCommerceSubscriptionIdRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SubscriptionsWithCommerceSubscriptionIdRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/directory/subscriptions(commerceSubscriptionId='{commerceSubscriptionId}')", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get a specific commercial subscription that an organization has acquired.
// returns a *RequestInformation when successful
func (m *SubscriptionsWithCommerceSubscriptionIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubscriptionsWithCommerceSubscriptionIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property subscriptions in directory
// returns a *RequestInformation when successful
func (m *SubscriptionsWithCommerceSubscriptionIdRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CompanySubscriptionable, requestConfiguration *SubscriptionsWithCommerceSubscriptionIdRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/directory/subscriptions(commerceSubscriptionId='{commerceSubscriptionId}')", m.BaseRequestBuilder.PathParameters)
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
// returns a *SubscriptionsWithCommerceSubscriptionIdRequestBuilder when successful
func (m *SubscriptionsWithCommerceSubscriptionIdRequestBuilder) WithUrl(rawUrl string)(*SubscriptionsWithCommerceSubscriptionIdRequestBuilder) {
    return NewSubscriptionsWithCommerceSubscriptionIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
