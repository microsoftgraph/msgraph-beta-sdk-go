package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPlannerFavoritePlansRequestBuilder provides operations to manage the favoritePlans property of the microsoft.graph.plannerUser entity.
type ItemPlannerFavoritePlansRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPlannerFavoritePlansRequestBuilderGetQueryParameters read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
type ItemPlannerFavoritePlansRequestBuilderGetQueryParameters struct {
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
// ItemPlannerFavoritePlansRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPlannerFavoritePlansRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPlannerFavoritePlansRequestBuilderGetQueryParameters
}
// ByPlannerPlanId provides operations to manage the favoritePlans property of the microsoft.graph.plannerUser entity.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemPlannerFavoritePlansPlannerPlanItemRequestBuilder when successful
func (m *ItemPlannerFavoritePlansRequestBuilder) ByPlannerPlanId(plannerPlanId string)(*ItemPlannerFavoritePlansPlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if plannerPlanId != "" {
        urlTplParams["plannerPlan%2Did"] = plannerPlanId
    }
    return NewItemPlannerFavoritePlansPlannerPlanItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPlannerFavoritePlansRequestBuilderInternal instantiates a new ItemPlannerFavoritePlansRequestBuilder and sets the default values.
func NewItemPlannerFavoritePlansRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPlannerFavoritePlansRequestBuilder) {
    m := &ItemPlannerFavoritePlansRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/planner/favoritePlans{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemPlannerFavoritePlansRequestBuilder instantiates a new ItemPlannerFavoritePlansRequestBuilder and sets the default values.
func NewItemPlannerFavoritePlansRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPlannerFavoritePlansRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPlannerFavoritePlansRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemPlannerFavoritePlansCountRequestBuilder when successful
func (m *ItemPlannerFavoritePlansRequestBuilder) Count()(*ItemPlannerFavoritePlansCountRequestBuilder) {
    return NewItemPlannerFavoritePlansCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a PlannerPlanCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPlannerFavoritePlansRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPlannerFavoritePlansRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanCollectionResponseable), nil
}
// ToGetRequestInformation read-only. Nullable. Returns the plannerPlans that the user marked as favorites.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemPlannerFavoritePlansRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPlannerFavoritePlansRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemPlannerFavoritePlansRequestBuilder when successful
func (m *ItemPlannerFavoritePlansRequestBuilder) WithUrl(rawUrl string)(*ItemPlannerFavoritePlansRequestBuilder) {
    return NewItemPlannerFavoritePlansRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
