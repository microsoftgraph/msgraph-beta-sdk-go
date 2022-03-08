package planner

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i276465e5f4c886d6a61356d326a2f9f84555d806812828064e6703be56a5db80 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/recentplans"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
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
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PlannerRequestBuilderDeleteOptions options for Delete
type PlannerRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PlannerRequestBuilderGetOptions options for Get
type PlannerRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PlannerRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
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
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerUserable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
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
func NewPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRequestBuilder) {
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
func NewPlannerRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property planner for me
func (m *PlannerRequestBuilder) CreateDeleteRequestInformation(options *PlannerRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation entry-point to the Planner resource that might exist for a user. Read-only.
func (m *PlannerRequestBuilder) CreateGetRequestInformation(options *PlannerRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property planner in me
func (m *PlannerRequestBuilder) CreatePatchRequestInformation(options *PlannerRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
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
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
func (m *PlannerRequestBuilder) Get(options *PlannerRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerUserable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreatePlannerUserFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerUserable), nil
}
// Patch update the navigation property planner in me
func (m *PlannerRequestBuilder) Patch(options *PlannerRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
