package planner

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i276465e5f4c886d6a61356d326a2f9f84555d806812828064e6703be56a5db80 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/recentplans"
    ia8b4b566a909edf1d2a6703ac30b7e3b8deeee1d18d9a70abe370e3f4c39de38 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/tasks"
    ie9512dfe2baec3e9b57eeec21a36064d1ca4f6c60bcf3659f5ae7cc90af786fd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/rosterplans"
    iec266a670c9004b3548436ef84f6d344262593ebfb8a5931109f8ae25d2ed114 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans"
    if951d99d6353794a9a240c80019dac3c903ab51eb2b16b3ffd092788641abe57 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/favoriteplans"
    ifbcaf5cdca7ba00b1339240d3736a4bb960b1974633dd806eb7b6b176c1ff1a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/all"
    i2181c3e27d21cc20f2f5d53a362d14d7eb43be3d6148c8d7221204191dcdb3a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/rosterplans/item"
    i29b52f5f3d78189485704d59dabe88f314c749e3fa9a830a26c5c232be71a81d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/recentplans/item"
    ia18da0e36a14ee15225f84a02a84b46591fb9f7c1fe1ea17a3e0bd47b603f5de "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/all/item"
    id3160a97184de0aa3787e4c5c3fab9e0c4a4affdab8f37380379b78945853ba9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/favoriteplans/item"
    ie8e0af54c5710406728d2f0cf561480ba34965894549bc0ef8a5a97e6516e83b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans/item"
    if89caed7c6cf9e56d6d5335f6c6908c367d55a0c198a5eef8530de7a51ed03f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/tasks/item"
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
func (m *PlannerRequestBuilder) All()(*ifbcaf5cdca7ba00b1339240d3736a4bb960b1974633dd806eb7b6b176c1ff1a4.AllRequestBuilder) {
    return ifbcaf5cdca7ba00b1339240d3736a4bb960b1974633dd806eb7b6b176c1ff1a4.NewAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.planner.all.item collection
func (m *PlannerRequestBuilder) AllById(id string)(*ia18da0e36a14ee15225f84a02a84b46591fb9f7c1fe1ea17a3e0bd47b603f5de.PlannerDeltaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerDelta_id"] = id
    }
    return ia18da0e36a14ee15225f84a02a84b46591fb9f7c1fe1ea17a3e0bd47b603f5de.NewPlannerDeltaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPlannerRequestBuilderInternal instantiates a new PlannerRequestBuilder and sets the default values.
func NewPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerRequestBuilder) {
    m := &PlannerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/planner{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property planner for me
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
// CreatePatchRequestInformation update the navigation property planner in me
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
// Delete delete navigation property planner for me
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
func (m *PlannerRequestBuilder) FavoritePlans()(*if951d99d6353794a9a240c80019dac3c903ab51eb2b16b3ffd092788641abe57.FavoritePlansRequestBuilder) {
    return if951d99d6353794a9a240c80019dac3c903ab51eb2b16b3ffd092788641abe57.NewFavoritePlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FavoritePlansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.planner.favoritePlans.item collection
func (m *PlannerRequestBuilder) FavoritePlansById(id string)(*id3160a97184de0aa3787e4c5c3fab9e0c4a4affdab8f37380379b78945853ba9.PlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return id3160a97184de0aa3787e4c5c3fab9e0c4a4affdab8f37380379b78945853ba9.NewPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
// Patch update the navigation property planner in me
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
func (m *PlannerRequestBuilder) Plans()(*iec266a670c9004b3548436ef84f6d344262593ebfb8a5931109f8ae25d2ed114.PlansRequestBuilder) {
    return iec266a670c9004b3548436ef84f6d344262593ebfb8a5931109f8ae25d2ed114.NewPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PlansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.planner.plans.item collection
func (m *PlannerRequestBuilder) PlansById(id string)(*ie8e0af54c5710406728d2f0cf561480ba34965894549bc0ef8a5a97e6516e83b.PlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return ie8e0af54c5710406728d2f0cf561480ba34965894549bc0ef8a5a97e6516e83b.NewPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RecentPlans the recentPlans property
func (m *PlannerRequestBuilder) RecentPlans()(*i276465e5f4c886d6a61356d326a2f9f84555d806812828064e6703be56a5db80.RecentPlansRequestBuilder) {
    return i276465e5f4c886d6a61356d326a2f9f84555d806812828064e6703be56a5db80.NewRecentPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecentPlansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.planner.recentPlans.item collection
func (m *PlannerRequestBuilder) RecentPlansById(id string)(*i29b52f5f3d78189485704d59dabe88f314c749e3fa9a830a26c5c232be71a81d.PlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return i29b52f5f3d78189485704d59dabe88f314c749e3fa9a830a26c5c232be71a81d.NewPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RosterPlans the rosterPlans property
func (m *PlannerRequestBuilder) RosterPlans()(*ie9512dfe2baec3e9b57eeec21a36064d1ca4f6c60bcf3659f5ae7cc90af786fd.RosterPlansRequestBuilder) {
    return ie9512dfe2baec3e9b57eeec21a36064d1ca4f6c60bcf3659f5ae7cc90af786fd.NewRosterPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RosterPlansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.planner.rosterPlans.item collection
func (m *PlannerRequestBuilder) RosterPlansById(id string)(*i2181c3e27d21cc20f2f5d53a362d14d7eb43be3d6148c8d7221204191dcdb3a7.PlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return i2181c3e27d21cc20f2f5d53a362d14d7eb43be3d6148c8d7221204191dcdb3a7.NewPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tasks the tasks property
func (m *PlannerRequestBuilder) Tasks()(*ia8b4b566a909edf1d2a6703ac30b7e3b8deeee1d18d9a70abe370e3f4c39de38.TasksRequestBuilder) {
    return ia8b4b566a909edf1d2a6703ac30b7e3b8deeee1d18d9a70abe370e3f4c39de38.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.planner.tasks.item collection
func (m *PlannerRequestBuilder) TasksById(id string)(*if89caed7c6cf9e56d6d5335f6c6908c367d55a0c198a5eef8530de7a51ed03f9.PlannerTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask_id"] = id
    }
    return if89caed7c6cf9e56d6d5335f6c6908c367d55a0c198a5eef8530de7a51ed03f9.NewPlannerTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
