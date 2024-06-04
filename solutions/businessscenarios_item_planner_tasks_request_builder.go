package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BusinessscenariosItemPlannerTasksRequestBuilder provides operations to manage the tasks property of the microsoft.graph.businessScenarioPlanner entity.
type BusinessscenariosItemPlannerTasksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BusinessscenariosItemPlannerTasksRequestBuilderGetQueryParameters get a list of the businessScenarioTask objects and their properties.
type BusinessscenariosItemPlannerTasksRequestBuilderGetQueryParameters struct {
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
// BusinessscenariosItemPlannerTasksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerTasksRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BusinessscenariosItemPlannerTasksRequestBuilderGetQueryParameters
}
// BusinessscenariosItemPlannerTasksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessscenariosItemPlannerTasksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByBusinessScenarioTaskId provides operations to manage the tasks property of the microsoft.graph.businessScenarioPlanner entity.
// returns a *BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder when successful
func (m *BusinessscenariosItemPlannerTasksRequestBuilder) ByBusinessScenarioTaskId(businessScenarioTaskId string)(*BusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if businessScenarioTaskId != "" {
        urlTplParams["businessScenarioTask%2Did"] = businessScenarioTaskId
    }
    return NewBusinessscenariosItemPlannerTasksBusinessScenarioTaskItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBusinessscenariosItemPlannerTasksRequestBuilderInternal instantiates a new BusinessscenariosItemPlannerTasksRequestBuilder and sets the default values.
func NewBusinessscenariosItemPlannerTasksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessscenariosItemPlannerTasksRequestBuilder) {
    m := &BusinessscenariosItemPlannerTasksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/businessScenarios/{businessScenario%2Did}/planner/tasks{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBusinessscenariosItemPlannerTasksRequestBuilder instantiates a new BusinessscenariosItemPlannerTasksRequestBuilder and sets the default values.
func NewBusinessscenariosItemPlannerTasksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BusinessscenariosItemPlannerTasksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBusinessscenariosItemPlannerTasksRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BusinessscenariosItemPlannerTasksCountRequestBuilder when successful
func (m *BusinessscenariosItemPlannerTasksRequestBuilder) Count()(*BusinessscenariosItemPlannerTasksCountRequestBuilder) {
    return NewBusinessscenariosItemPlannerTasksCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the businessScenarioTask objects and their properties.
// returns a BusinessScenarioTaskCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/businessscenarioplanner-list-tasks?view=graph-rest-beta
func (m *BusinessscenariosItemPlannerTasksRequestBuilder) Get(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerTasksRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessScenarioTaskCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskCollectionResponseable), nil
}
// Post create a new businessScenarioTask object.
// returns a BusinessScenarioTaskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/businessscenarioplanner-post-tasks?view=graph-rest-beta
func (m *BusinessscenariosItemPlannerTasksRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskable, requestConfiguration *BusinessscenariosItemPlannerTasksRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBusinessScenarioTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskable), nil
}
// ToGetRequestInformation get a list of the businessScenarioTask objects and their properties.
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerTasksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BusinessscenariosItemPlannerTasksRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new businessScenarioTask object.
// returns a *RequestInformation when successful
func (m *BusinessscenariosItemPlannerTasksRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BusinessScenarioTaskable, requestConfiguration *BusinessscenariosItemPlannerTasksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BusinessscenariosItemPlannerTasksRequestBuilder when successful
func (m *BusinessscenariosItemPlannerTasksRequestBuilder) WithUrl(rawUrl string)(*BusinessscenariosItemPlannerTasksRequestBuilder) {
    return NewBusinessscenariosItemPlannerTasksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
