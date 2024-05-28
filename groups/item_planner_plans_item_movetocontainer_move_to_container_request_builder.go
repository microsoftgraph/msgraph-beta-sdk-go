package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder provides operations to call the moveToContainer method.
type ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilderInternal instantiates a new ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder and sets the default values.
func NewItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder) {
    m := &ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/planner/plans/{plannerPlan%2Did}/moveToContainer", pathParameters),
    }
    return m
}
// NewItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder instantiates a new ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder and sets the default values.
func NewItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilderInternal(urlParams, requestAdapter)
}
// Post move a planner plan object from one planner plan container to another. Planner plans can only be moved from a user container to a group container.
// returns a PlannerPlanable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/plannerplan-movetocontainer?view=graph-rest-beta
func (m *ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder) Post(ctx context.Context, body ItemPlannerPlansItemMovetocontainerMoveToContainerPostRequestBodyable, requestConfiguration *ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanable), nil
}
// ToPostRequestInformation move a planner plan object from one planner plan container to another. Planner plans can only be moved from a user container to a group container.
// returns a *RequestInformation when successful
func (m *ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPlannerPlansItemMovetocontainerMoveToContainerPostRequestBodyable, requestConfiguration *ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder when successful
func (m *ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder) WithUrl(rawUrl string)(*ItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder) {
    return NewItemPlannerPlansItemMovetocontainerMoveToContainerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
