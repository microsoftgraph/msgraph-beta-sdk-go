package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder provides operations to manage the planConfiguration property of the microsoft.graph.businessScenarioPlanner entity.
type BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderGetQueryParameters read the properties and relationships of a plannerPlanConfiguration object.
type BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderGetQueryParameters
}
// BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderInternal instantiates a new BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder and sets the default values.
func NewBusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder) {
    m := &BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/businessScenarios/{businessScenario%2Did}/planner/planConfiguration{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder instantiates a new BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder and sets the default values.
func NewBusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property planConfiguration for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder) Delete(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a plannerPlanConfiguration object.
// returns a PlannerPlanConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/plannerplanconfiguration-get?view=graph-rest-beta
func (m *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder) Get(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationable), nil
}
// Localizations provides operations to manage the localizations property of the microsoft.graph.plannerPlanConfiguration entity.
// returns a *BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder) Localizations()(*BusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilder) {
    return NewBusinessscenariosItemPlannerPlanconfigurationLocalizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a plannerPlanConfiguration object for a businessScenario.
// returns a PlannerPlanConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/plannerplanconfiguration-update?view=graph-rest-beta
func (m *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationable, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerPlanConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property planConfiguration for solutions
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a plannerPlanConfiguration object.
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a plannerPlanConfiguration object for a businessScenario.
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerPlanConfigurationable, requestConfiguration *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder when successful
func (m *BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder) WithUrl(rawUrl string)(*BusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder) {
    return NewBusinessscenariosItemPlannerPlanconfigurationPlanConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
