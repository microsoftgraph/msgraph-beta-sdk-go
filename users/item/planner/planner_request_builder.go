package planner

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i089c3d3abe24c51d75b3535dafe2c0f6501ace7e1bbade3c5349577c564a3642 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/tasks"
    i4e26d6e5e4e60c41cf7e0126190c23e4d4b6ac317c46b74e7a26c14b6e1ee02d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/plans"
    i9f3bb673ec53be741be4ac1dff78058797b2e863d80e66433c720e99f9bffdad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/recentplans"
    ibf58c71a094fcdfd1bc9c2e0c6c4a7f17629a13a38945c331ac2090bdd864963 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/all"
    id046e247c84498d4c7ad4bea9771219a60d6b5fb6fbd69f7131a8e807b851816 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/rosterplans"
    iff81dee8c57e8cdf388de7660fa9d05ee375305f050c5322e7392b5cfa49e6b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/favoriteplans"
    i723887d2b7bb1d5b453d57c98fe83989e90f0efdf4225427a652314064664806 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/rosterplans/item"
    i75445737f77febd8e8ab6f9ab4a51048a800000b3b33d3582945e735d9732587 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/tasks/item"
    i86017907c6216aa0164fd179d5714657602a4a6e42c45ca14763c542c6cbd85e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/recentplans/item"
    i86f40f45c3011d5e68ca73d248500d11e56081a192959e8099e5577b8ee31efd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/plans/item"
    ib8e4b0e29fa0bab4a12714aa2c4ba3540af7adea774b8f3bac0f193c1c83baa2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/favoriteplans/item"
    iea71824a70bb4972dcddc0b056992a9f6c83b550f5c1191520345be2898bd0f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/all/item"
)

// PlannerRequestBuilder provides operations to manage the planner property of the microsoft.graph.user entity.
type PlannerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PlannerRequestBuilderDeleteOptions options for Delete
type PlannerRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PlannerRequestBuilderGetOptions options for Get
type PlannerRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *PlannerRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PlannerRequestBuilderGetQueryParameters entry-point to the Planner resource that might exist for a user. Read-only.
type PlannerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PlannerRequestBuilderPatchOptions options for Patch
type PlannerRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerUserable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// All the all property
func (m *PlannerRequestBuilder) All()(*ibf58c71a094fcdfd1bc9c2e0c6c4a7f17629a13a38945c331ac2090bdd864963.AllRequestBuilder) {
    return ibf58c71a094fcdfd1bc9c2e0c6c4a7f17629a13a38945c331ac2090bdd864963.NewAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.planner.all.item collection
func (m *PlannerRequestBuilder) AllById(id string)(*iea71824a70bb4972dcddc0b056992a9f6c83b550f5c1191520345be2898bd0f2.PlannerDeltaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerDelta_id"] = id
    }
    return iea71824a70bb4972dcddc0b056992a9f6c83b550f5c1191520345be2898bd0f2.NewPlannerDeltaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPlannerRequestBuilderInternal instantiates a new PlannerRequestBuilder and sets the default values.
func NewPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerRequestBuilder) {
    m := &PlannerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/planner{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property planner for users
func (m *PlannerRequestBuilder) CreateDeleteRequestInformation(options *PlannerRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation entry-point to the Planner resource that might exist for a user. Read-only.
func (m *PlannerRequestBuilder) CreateGetRequestInformation(options *PlannerRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property planner in users
func (m *PlannerRequestBuilder) CreatePatchRequestInformation(options *PlannerRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property planner for users
func (m *PlannerRequestBuilder) Delete(options *PlannerRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// FavoritePlans the favoritePlans property
func (m *PlannerRequestBuilder) FavoritePlans()(*iff81dee8c57e8cdf388de7660fa9d05ee375305f050c5322e7392b5cfa49e6b6.FavoritePlansRequestBuilder) {
    return iff81dee8c57e8cdf388de7660fa9d05ee375305f050c5322e7392b5cfa49e6b6.NewFavoritePlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FavoritePlansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.planner.favoritePlans.item collection
func (m *PlannerRequestBuilder) FavoritePlansById(id string)(*ib8e4b0e29fa0bab4a12714aa2c4ba3540af7adea774b8f3bac0f193c1c83baa2.PlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return ib8e4b0e29fa0bab4a12714aa2c4ba3540af7adea774b8f3bac0f193c1c83baa2.NewPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get entry-point to the Planner resource that might exist for a user. Read-only.
func (m *PlannerRequestBuilder) Get(options *PlannerRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerUserable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePlannerUserFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PlannerUserable), nil
}
// Patch update the navigation property planner in users
func (m *PlannerRequestBuilder) Patch(options *PlannerRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Plans the plans property
func (m *PlannerRequestBuilder) Plans()(*i4e26d6e5e4e60c41cf7e0126190c23e4d4b6ac317c46b74e7a26c14b6e1ee02d.PlansRequestBuilder) {
    return i4e26d6e5e4e60c41cf7e0126190c23e4d4b6ac317c46b74e7a26c14b6e1ee02d.NewPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PlansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.planner.plans.item collection
func (m *PlannerRequestBuilder) PlansById(id string)(*i86f40f45c3011d5e68ca73d248500d11e56081a192959e8099e5577b8ee31efd.PlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return i86f40f45c3011d5e68ca73d248500d11e56081a192959e8099e5577b8ee31efd.NewPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RecentPlans the recentPlans property
func (m *PlannerRequestBuilder) RecentPlans()(*i9f3bb673ec53be741be4ac1dff78058797b2e863d80e66433c720e99f9bffdad.RecentPlansRequestBuilder) {
    return i9f3bb673ec53be741be4ac1dff78058797b2e863d80e66433c720e99f9bffdad.NewRecentPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecentPlansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.planner.recentPlans.item collection
func (m *PlannerRequestBuilder) RecentPlansById(id string)(*i86017907c6216aa0164fd179d5714657602a4a6e42c45ca14763c542c6cbd85e.PlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return i86017907c6216aa0164fd179d5714657602a4a6e42c45ca14763c542c6cbd85e.NewPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RosterPlans the rosterPlans property
func (m *PlannerRequestBuilder) RosterPlans()(*id046e247c84498d4c7ad4bea9771219a60d6b5fb6fbd69f7131a8e807b851816.RosterPlansRequestBuilder) {
    return id046e247c84498d4c7ad4bea9771219a60d6b5fb6fbd69f7131a8e807b851816.NewRosterPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RosterPlansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.planner.rosterPlans.item collection
func (m *PlannerRequestBuilder) RosterPlansById(id string)(*i723887d2b7bb1d5b453d57c98fe83989e90f0efdf4225427a652314064664806.PlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return i723887d2b7bb1d5b453d57c98fe83989e90f0efdf4225427a652314064664806.NewPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tasks the tasks property
func (m *PlannerRequestBuilder) Tasks()(*i089c3d3abe24c51d75b3535dafe2c0f6501ace7e1bbade3c5349577c564a3642.TasksRequestBuilder) {
    return i089c3d3abe24c51d75b3535dafe2c0f6501ace7e1bbade3c5349577c564a3642.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.planner.tasks.item collection
func (m *PlannerRequestBuilder) TasksById(id string)(*i75445737f77febd8e8ab6f9ab4a51048a800000b3b33d3582945e735d9732587.PlannerTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask_id"] = id
    }
    return i75445737f77febd8e8ab6f9ab4a51048a800000b3b33d3582945e735d9732587.NewPlannerTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
