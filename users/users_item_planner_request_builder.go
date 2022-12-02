package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemPlannerRequestBuilder provides operations to manage the planner property of the microsoft.graph.user entity.
type UsersItemPlannerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemPlannerRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPlannerRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemPlannerRequestBuilderGetQueryParameters retrieve the properties and relationships of a plannerUser object. The returned properties include the user's favorite plans and recently viewed plans. 
type UsersItemPlannerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemPlannerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPlannerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemPlannerRequestBuilderGetQueryParameters
}
// UsersItemPlannerRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemPlannerRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// All provides operations to manage the all property of the microsoft.graph.plannerUser entity.
func (m *UsersItemPlannerRequestBuilder) All()(*UsersItemPlannerAllRequestBuilder) {
    return NewUsersItemPlannerAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllById provides operations to manage the all property of the microsoft.graph.plannerUser entity.
func (m *UsersItemPlannerRequestBuilder) AllById(id string)(*UsersItemPlannerAllPlannerDeltaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerDelta%2Did"] = id
    }
    return NewUsersItemPlannerAllPlannerDeltaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUsersItemPlannerRequestBuilderInternal instantiates a new PlannerRequestBuilder and sets the default values.
func NewUsersItemPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemPlannerRequestBuilder) {
    m := &UsersItemPlannerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/planner{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemPlannerRequestBuilder instantiates a new PlannerRequestBuilder and sets the default values.
func NewUsersItemPlannerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemPlannerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemPlannerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property planner for users
func (m *UsersItemPlannerRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemPlannerRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation retrieve the properties and relationships of a plannerUser object. The returned properties include the user's favorite plans and recently viewed plans. 
func (m *UsersItemPlannerRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemPlannerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of a plannerUser object. You can use this operation to add or remove plans from a user's favorite plans list, and to indicate which plans the user has recently viewed.
func (m *UsersItemPlannerRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerUserable, requestConfiguration *UsersItemPlannerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property planner for users
func (m *UsersItemPlannerRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemPlannerRequestBuilderDeleteRequestConfiguration)(error) {
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
// FavoritePlans provides operations to manage the favoritePlans property of the microsoft.graph.plannerUser entity.
func (m *UsersItemPlannerRequestBuilder) FavoritePlans()(*UsersItemPlannerFavoritePlansRequestBuilder) {
    return NewUsersItemPlannerFavoritePlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FavoritePlansById provides operations to manage the favoritePlans property of the microsoft.graph.plannerUser entity.
func (m *UsersItemPlannerRequestBuilder) FavoritePlansById(id string)(*UsersItemPlannerFavoritePlansPlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan%2Did"] = id
    }
    return NewUsersItemPlannerFavoritePlansPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get retrieve the properties and relationships of a plannerUser object. The returned properties include the user's favorite plans and recently viewed plans. 
func (m *UsersItemPlannerRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemPlannerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerUserable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerUserable), nil
}
// Patch update the properties of a plannerUser object. You can use this operation to add or remove plans from a user's favorite plans list, and to indicate which plans the user has recently viewed.
func (m *UsersItemPlannerRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerUserable, requestConfiguration *UsersItemPlannerRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerUserable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerUserable), nil
}
// Plans provides operations to manage the plans property of the microsoft.graph.plannerUser entity.
func (m *UsersItemPlannerRequestBuilder) Plans()(*UsersItemPlannerPlansRequestBuilder) {
    return NewUsersItemPlannerPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PlansById provides operations to manage the plans property of the microsoft.graph.plannerUser entity.
func (m *UsersItemPlannerRequestBuilder) PlansById(id string)(*UsersItemPlannerPlansPlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan%2Did"] = id
    }
    return NewUsersItemPlannerPlansPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RecentPlans provides operations to manage the recentPlans property of the microsoft.graph.plannerUser entity.
func (m *UsersItemPlannerRequestBuilder) RecentPlans()(*UsersItemPlannerRecentPlansRequestBuilder) {
    return NewUsersItemPlannerRecentPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecentPlansById provides operations to manage the recentPlans property of the microsoft.graph.plannerUser entity.
func (m *UsersItemPlannerRequestBuilder) RecentPlansById(id string)(*UsersItemPlannerRecentPlansPlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan%2Did"] = id
    }
    return NewUsersItemPlannerRecentPlansPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RosterPlans provides operations to manage the rosterPlans property of the microsoft.graph.plannerUser entity.
func (m *UsersItemPlannerRequestBuilder) RosterPlans()(*UsersItemPlannerRosterPlansRequestBuilder) {
    return NewUsersItemPlannerRosterPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RosterPlansById provides operations to manage the rosterPlans property of the microsoft.graph.plannerUser entity.
func (m *UsersItemPlannerRequestBuilder) RosterPlansById(id string)(*UsersItemPlannerRosterPlansPlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan%2Did"] = id
    }
    return NewUsersItemPlannerRosterPlansPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.plannerUser entity.
func (m *UsersItemPlannerRequestBuilder) Tasks()(*UsersItemPlannerTasksRequestBuilder) {
    return NewUsersItemPlannerTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById provides operations to manage the tasks property of the microsoft.graph.plannerUser entity.
func (m *UsersItemPlannerRequestBuilder) TasksById(id string)(*UsersItemPlannerTasksPlannerTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask%2Did"] = id
    }
    return NewUsersItemPlannerTasksPlannerTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
