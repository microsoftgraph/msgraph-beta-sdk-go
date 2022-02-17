package planner

import (
    i1f734a8e118043874e700fcde1d31c1b482783dc8c8d16b4b733c5c4468193a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/buckets"
    i20268753e7c3020410bc7b257943b38d7ddbd7dac155d5aa908638a484563a0b "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/tasks"
    i24498e00b94127cf52e32b44bb4a7c7b55b4996fcc8d8f7fc5588dc3f25081a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/plans"
    ib18d47314351fa2c87f6ae4b1626819d77c3e1980d8f164525db440491fe7b4e "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i49d7402ded8d0526c75784ca123eb70bb195f40f6fbf171738a0e3141a53d9fe "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/tasks/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i8cd70227fa2a2e73750c8e4e63440317eb3ed8404fbba0976dc45ebacacd40f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/buckets/item"
    i8de8f0901d178bd52a19035fe88edf4f26f76e989740d3eb1b20e83222def4c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/plans/item"
    ia1542c83c900b2622d1e688fb4b392f0b9a5aa87a65e4fd4f6b748c84899d0af "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters/item"
)

// PlannerRequestBuilder builds and executes requests for operations under \planner
type PlannerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
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
// PlannerRequestBuilderGetQueryParameters get planner
type PlannerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PlannerRequestBuilderPatchOptions options for Patch
type PlannerRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Planner;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PlannerRequestBuilder) Buckets()(*i1f734a8e118043874e700fcde1d31c1b482783dc8c8d16b4b733c5c4468193a7.BucketsRequestBuilder) {
    return i1f734a8e118043874e700fcde1d31c1b482783dc8c8d16b4b733c5c4468193a7.NewBucketsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BucketsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.planner.buckets.item collection
func (m *PlannerRequestBuilder) BucketsById(id string)(*i8cd70227fa2a2e73750c8e4e63440317eb3ed8404fbba0976dc45ebacacd40f5.PlannerBucketRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerBucket_id"] = id
    }
    return i8cd70227fa2a2e73750c8e4e63440317eb3ed8404fbba0976dc45ebacacd40f5.NewPlannerBucketRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPlannerRequestBuilderInternal instantiates a new PlannerRequestBuilder and sets the default values.
func NewPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRequestBuilder) {
    m := &PlannerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/planner{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPlannerRequestBuilder instantiates a new PlannerRequestBuilder and sets the default values.
func NewPlannerRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get planner
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
// CreatePatchRequestInformation update planner
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
// Get get planner
func (m *PlannerRequestBuilder) Get(options *PlannerRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Planner, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPlanner() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Planner), nil
}
// Patch update planner
func (m *PlannerRequestBuilder) Patch(options *PlannerRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *PlannerRequestBuilder) Plans()(*i24498e00b94127cf52e32b44bb4a7c7b55b4996fcc8d8f7fc5588dc3f25081a7.PlansRequestBuilder) {
    return i24498e00b94127cf52e32b44bb4a7c7b55b4996fcc8d8f7fc5588dc3f25081a7.NewPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PlansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.planner.plans.item collection
func (m *PlannerRequestBuilder) PlansById(id string)(*i8de8f0901d178bd52a19035fe88edf4f26f76e989740d3eb1b20e83222def4c6.PlannerPlanRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return i8de8f0901d178bd52a19035fe88edf4f26f76e989740d3eb1b20e83222def4c6.NewPlannerPlanRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PlannerRequestBuilder) Rosters()(*ib18d47314351fa2c87f6ae4b1626819d77c3e1980d8f164525db440491fe7b4e.RostersRequestBuilder) {
    return ib18d47314351fa2c87f6ae4b1626819d77c3e1980d8f164525db440491fe7b4e.NewRostersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RostersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.planner.rosters.item collection
func (m *PlannerRequestBuilder) RostersById(id string)(*ia1542c83c900b2622d1e688fb4b392f0b9a5aa87a65e4fd4f6b748c84899d0af.PlannerRosterRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerRoster_id"] = id
    }
    return ia1542c83c900b2622d1e688fb4b392f0b9a5aa87a65e4fd4f6b748c84899d0af.NewPlannerRosterRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PlannerRequestBuilder) Tasks()(*i20268753e7c3020410bc7b257943b38d7ddbd7dac155d5aa908638a484563a0b.TasksRequestBuilder) {
    return i20268753e7c3020410bc7b257943b38d7ddbd7dac155d5aa908638a484563a0b.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.planner.tasks.item collection
func (m *PlannerRequestBuilder) TasksById(id string)(*i49d7402ded8d0526c75784ca123eb70bb195f40f6fbf171738a0e3141a53d9fe.PlannerTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask_id"] = id
    }
    return i49d7402ded8d0526c75784ca123eb70bb195f40f6fbf171738a0e3141a53d9fe.NewPlannerTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
