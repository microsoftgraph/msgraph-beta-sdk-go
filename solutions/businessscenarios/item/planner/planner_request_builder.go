package planner

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3464f37fce63fbe71f0b87126d102910e4679d50598f60818312015c40b9b5e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/solutions/businessscenarios/item/planner/tasks"
    i426181369e3a496020a9c54ae2fb39dda9d20f3b4832db5cbfe5801b85d68c11 "github.com/microsoftgraph/msgraph-beta-sdk-go/solutions/businessscenarios/item/planner/planconfiguration"
    i5d5d1b11f81265adc48072ba27d19afad2cecdd8709f7464c2ac38b256e06d92 "github.com/microsoftgraph/msgraph-beta-sdk-go/solutions/businessscenarios/item/planner/getplan"
    ia1585e2fb942650082bfed4b4a1e8acf6272f563e20741c03a31ab1dcc46bc61 "github.com/microsoftgraph/msgraph-beta-sdk-go/solutions/businessscenarios/item/planner/taskconfiguration"
    ic65a2ecb53e89b91403ad96e65204e0f68441814d1f79853790c1dea9bcfc5e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/solutions/businessscenarios/item/planner/tasks/item"
)

// PlannerRequestBuilder provides operations to manage the planner property of the microsoft.graph.businessScenario entity.
type PlannerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PlannerRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlannerRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PlannerRequestBuilderGetQueryParameters get planner from solutions
type PlannerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PlannerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlannerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PlannerRequestBuilderGetQueryParameters
}
// PlannerRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PlannerRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPlannerRequestBuilderInternal instantiates a new PlannerRequestBuilder and sets the default values.
func NewPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerRequestBuilder) {
    m := &PlannerRequestBuilder{
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
// NewPlannerRequestBuilder instantiates a new PlannerRequestBuilder and sets the default values.
func NewPlannerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property planner for solutions
func (m *PlannerRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PlannerRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PlannerRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PlannerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PlannerRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable, requestConfiguration *PlannerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *PlannerRequestBuilder) Delete(ctx context.Context, requestConfiguration *PlannerRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *PlannerRequestBuilder) Get(ctx context.Context, requestConfiguration *PlannerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable, error) {
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
func (m *PlannerRequestBuilder) GetPlan()(*i5d5d1b11f81265adc48072ba27d19afad2cecdd8709f7464c2ac38b256e06d92.GetPlanRequestBuilder) {
    return i5d5d1b11f81265adc48072ba27d19afad2cecdd8709f7464c2ac38b256e06d92.NewGetPlanRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property planner in solutions
func (m *PlannerRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable, requestConfiguration *PlannerRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioPlannerable, error) {
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
func (m *PlannerRequestBuilder) PlanConfiguration()(*i426181369e3a496020a9c54ae2fb39dda9d20f3b4832db5cbfe5801b85d68c11.PlanConfigurationRequestBuilder) {
    return i426181369e3a496020a9c54ae2fb39dda9d20f3b4832db5cbfe5801b85d68c11.NewPlanConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskConfiguration provides operations to manage the taskConfiguration property of the microsoft.graph.businessScenarioPlanner entity.
func (m *PlannerRequestBuilder) TaskConfiguration()(*ia1585e2fb942650082bfed4b4a1e8acf6272f563e20741c03a31ab1dcc46bc61.TaskConfigurationRequestBuilder) {
    return ia1585e2fb942650082bfed4b4a1e8acf6272f563e20741c03a31ab1dcc46bc61.NewTaskConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.businessScenarioPlanner entity.
func (m *PlannerRequestBuilder) Tasks()(*i3464f37fce63fbe71f0b87126d102910e4679d50598f60818312015c40b9b5e9.TasksRequestBuilder) {
    return i3464f37fce63fbe71f0b87126d102910e4679d50598f60818312015c40b9b5e9.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById provides operations to manage the tasks property of the microsoft.graph.businessScenarioPlanner entity.
func (m *PlannerRequestBuilder) TasksById(id string)(*ic65a2ecb53e89b91403ad96e65204e0f68441814d1f79853790c1dea9bcfc5e5.BusinessScenarioTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["businessScenarioTask%2Did"] = id
    }
    return ic65a2ecb53e89b91403ad96e65204e0f68441814d1f79853790c1dea9bcfc5e5.NewBusinessScenarioTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
