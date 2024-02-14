package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BusinessScenariosItemPlannerRequestBuilder provides operations to manage the planner property of the microsoft.graph.businessScenario entity.
type BusinessScenariosItemPlannerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BusinessScenariosItemPlannerRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessScenariosItemPlannerRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BusinessScenariosItemPlannerRequestBuilderGetQueryParameters read the properties and relationships of a businessScenarioPlanner object.
type BusinessScenariosItemPlannerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BusinessScenariosItemPlannerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessScenariosItemPlannerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BusinessScenariosItemPlannerRequestBuilderGetQueryParameters
}
// BusinessScenariosItemPlannerRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessScenariosItemPlannerRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewBusinessScenariosItemPlannerRequestBuilderInternal instantiates a new BusinessScenariosItemPlannerRequestBuilder and sets the default values.
func NewBusinessScenariosItemPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessScenariosItemPlannerRequestBuilder) {
    m := &BusinessScenariosItemPlannerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/businessScenarios/{businessScenario%2Did}/planner{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBusinessScenariosItemPlannerRequestBuilder instantiates a new BusinessScenariosItemPlannerRequestBuilder and sets the default values.
func NewBusinessScenariosItemPlannerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessScenariosItemPlannerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBusinessScenariosItemPlannerRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property planner for solutions
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BusinessScenariosItemPlannerRequestBuilder) Delete(ctx context.Context, requestConfiguration *BusinessScenariosItemPlannerRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a businessScenarioPlanner object.
// returns a BusinessScenarioPlannerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/businessscenarioplanner-get?view=graph-rest-1.0
func (m *BusinessScenariosItemPlannerRequestBuilder) Get(ctx context.Context, requestConfiguration *BusinessScenariosItemPlannerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessScenarioPlannerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable), nil
}
// GetPlan provides operations to call the getPlan method.
// returns a *BusinessScenariosItemPlannerGetPlanRequestBuilder when successful
func (m *BusinessScenariosItemPlannerRequestBuilder) GetPlan()(*BusinessScenariosItemPlannerGetPlanRequestBuilder) {
    return NewBusinessScenariosItemPlannerGetPlanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property planner in solutions
// returns a BusinessScenarioPlannerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BusinessScenariosItemPlannerRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable, requestConfiguration *BusinessScenariosItemPlannerRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessScenarioPlannerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable), nil
}
// PlanConfiguration provides operations to manage the planConfiguration property of the microsoft.graph.businessScenarioPlanner entity.
// returns a *BusinessScenariosItemPlannerPlanConfigurationRequestBuilder when successful
func (m *BusinessScenariosItemPlannerRequestBuilder) PlanConfiguration()(*BusinessScenariosItemPlannerPlanConfigurationRequestBuilder) {
    return NewBusinessScenariosItemPlannerPlanConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaskConfiguration provides operations to manage the taskConfiguration property of the microsoft.graph.businessScenarioPlanner entity.
// returns a *BusinessScenariosItemPlannerTaskConfigurationRequestBuilder when successful
func (m *BusinessScenariosItemPlannerRequestBuilder) TaskConfiguration()(*BusinessScenariosItemPlannerTaskConfigurationRequestBuilder) {
    return NewBusinessScenariosItemPlannerTaskConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.businessScenarioPlanner entity.
// returns a *BusinessScenariosItemPlannerTasksRequestBuilder when successful
func (m *BusinessScenariosItemPlannerRequestBuilder) Tasks()(*BusinessScenariosItemPlannerTasksRequestBuilder) {
    return NewBusinessScenariosItemPlannerTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property planner for solutions
// returns a *RequestInformation when successful
func (m *BusinessScenariosItemPlannerRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *BusinessScenariosItemPlannerRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/solutions/businessScenarios/{businessScenario%2Did}/planner", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a businessScenarioPlanner object.
// returns a *RequestInformation when successful
func (m *BusinessScenariosItemPlannerRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BusinessScenariosItemPlannerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property planner in solutions
// returns a *RequestInformation when successful
func (m *BusinessScenariosItemPlannerRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable, requestConfiguration *BusinessScenariosItemPlannerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/solutions/businessScenarios/{businessScenario%2Did}/planner", m.BaseRequestBuilder.PathParameters)
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
// returns a *BusinessScenariosItemPlannerRequestBuilder when successful
func (m *BusinessScenariosItemPlannerRequestBuilder) WithUrl(rawUrl string)(*BusinessScenariosItemPlannerRequestBuilder) {
    return NewBusinessScenariosItemPlannerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
