package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SolutionsBusinessScenariosItemPlannerRequestBuilder provides operations to manage the planner property of the microsoft.graph.businessScenario entity.
type SolutionsBusinessScenariosItemPlannerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SolutionsBusinessScenariosItemPlannerRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SolutionsBusinessScenariosItemPlannerRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SolutionsBusinessScenariosItemPlannerRequestBuilderGetQueryParameters get planner from solutions
type SolutionsBusinessScenariosItemPlannerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SolutionsBusinessScenariosItemPlannerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SolutionsBusinessScenariosItemPlannerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SolutionsBusinessScenariosItemPlannerRequestBuilderGetQueryParameters
}
// SolutionsBusinessScenariosItemPlannerRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SolutionsBusinessScenariosItemPlannerRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSolutionsBusinessScenariosItemPlannerRequestBuilderInternal instantiates a new PlannerRequestBuilder and sets the default values.
func NewSolutionsBusinessScenariosItemPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SolutionsBusinessScenariosItemPlannerRequestBuilder) {
    m := &SolutionsBusinessScenariosItemPlannerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/solutions/businessScenarios/{businessScenario%2Did}/planner{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSolutionsBusinessScenariosItemPlannerRequestBuilder instantiates a new PlannerRequestBuilder and sets the default values.
func NewSolutionsBusinessScenariosItemPlannerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SolutionsBusinessScenariosItemPlannerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSolutionsBusinessScenariosItemPlannerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property planner for solutions
func (m *SolutionsBusinessScenariosItemPlannerRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SolutionsBusinessScenariosItemPlannerRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get planner from solutions
func (m *SolutionsBusinessScenariosItemPlannerRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SolutionsBusinessScenariosItemPlannerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property planner in solutions
func (m *SolutionsBusinessScenariosItemPlannerRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable, requestConfiguration *SolutionsBusinessScenariosItemPlannerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property planner for solutions
func (m *SolutionsBusinessScenariosItemPlannerRequestBuilder) Delete(ctx context.Context, requestConfiguration *SolutionsBusinessScenariosItemPlannerRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get planner from solutions
func (m *SolutionsBusinessScenariosItemPlannerRequestBuilder) Get(ctx context.Context, requestConfiguration *SolutionsBusinessScenariosItemPlannerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessScenarioPlannerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable), nil
}
// GetPlan provides operations to call the getPlan method.
func (m *SolutionsBusinessScenariosItemPlannerRequestBuilder) GetPlan()(*SolutionsBusinessScenariosItemPlannerGetPlanRequestBuilder) {
    return NewSolutionsBusinessScenariosItemPlannerGetPlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property planner in solutions
func (m *SolutionsBusinessScenariosItemPlannerRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable, requestConfiguration *SolutionsBusinessScenariosItemPlannerRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessScenarioPlannerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable), nil
}
// PlanConfiguration provides operations to manage the planConfiguration property of the microsoft.graph.businessScenarioPlanner entity.
func (m *SolutionsBusinessScenariosItemPlannerRequestBuilder) PlanConfiguration()(*SolutionsBusinessScenariosItemPlannerPlanConfigurationRequestBuilder) {
    return NewSolutionsBusinessScenariosItemPlannerPlanConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskConfiguration provides operations to manage the taskConfiguration property of the microsoft.graph.businessScenarioPlanner entity.
func (m *SolutionsBusinessScenariosItemPlannerRequestBuilder) TaskConfiguration()(*SolutionsBusinessScenariosItemPlannerTaskConfigurationRequestBuilder) {
    return NewSolutionsBusinessScenariosItemPlannerTaskConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.businessScenarioPlanner entity.
func (m *SolutionsBusinessScenariosItemPlannerRequestBuilder) Tasks()(*SolutionsBusinessScenariosItemPlannerTasksRequestBuilder) {
    return NewSolutionsBusinessScenariosItemPlannerTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById provides operations to manage the tasks property of the microsoft.graph.businessScenarioPlanner entity.
func (m *SolutionsBusinessScenariosItemPlannerRequestBuilder) TasksById(id string)(*SolutionsBusinessScenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["businessScenarioTask%2Did"] = id
    }
    return NewSolutionsBusinessScenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
