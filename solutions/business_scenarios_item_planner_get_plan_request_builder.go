package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BusinessScenariosItemPlannerGetPlanRequestBuilder provides operations to call the getPlan method.
type BusinessScenariosItemPlannerGetPlanRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BusinessScenariosItemPlannerGetPlanRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessScenariosItemPlannerGetPlanRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBusinessScenariosItemPlannerGetPlanRequestBuilderInternal instantiates a new GetPlanRequestBuilder and sets the default values.
func NewBusinessScenariosItemPlannerGetPlanRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessScenariosItemPlannerGetPlanRequestBuilder) {
    m := &BusinessScenariosItemPlannerGetPlanRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/businessScenarios/{businessScenario%2Did}/planner/getPlan", pathParameters),
    }
    return m
}
// NewBusinessScenariosItemPlannerGetPlanRequestBuilder instantiates a new GetPlanRequestBuilder and sets the default values.
func NewBusinessScenariosItemPlannerGetPlanRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessScenariosItemPlannerGetPlanRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBusinessScenariosItemPlannerGetPlanRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get information about the plannerPlan mapped to a given target. If a plannerPlan doesn't exist for the specified target at the time of the request, a new plan will be created for the businessScenario.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/businessscenarioplanner-getplan?view=graph-rest-1.0
func (m *BusinessScenariosItemPlannerGetPlanRequestBuilder) Post(ctx context.Context, body BusinessScenariosItemPlannerGetPlanPostRequestBodyable, requestConfiguration *BusinessScenariosItemPlannerGetPlanRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlanReferenceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessScenarioPlanReferenceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlanReferenceable), nil
}
// ToPostRequestInformation get information about the plannerPlan mapped to a given target. If a plannerPlan doesn't exist for the specified target at the time of the request, a new plan will be created for the businessScenario.
func (m *BusinessScenariosItemPlannerGetPlanRequestBuilder) ToPostRequestInformation(ctx context.Context, body BusinessScenariosItemPlannerGetPlanPostRequestBodyable, requestConfiguration *BusinessScenariosItemPlannerGetPlanRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *BusinessScenariosItemPlannerGetPlanRequestBuilder) WithUrl(rawUrl string)(*BusinessScenariosItemPlannerGetPlanRequestBuilder) {
    return NewBusinessScenariosItemPlannerGetPlanRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
