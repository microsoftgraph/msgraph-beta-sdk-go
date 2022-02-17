package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i889297e93166312716592efbb05995455ae1882e388416fee06599864e54e462 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans/item/tasks"
    i97f4aa0504a94677b48d17585be69f3adfec35ef126990f5c0fa1deb971970ee "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans/item/buckets"
    i9800485461c637fd964b8b604adecc9eb52f29415b0fdcbce97286ef735a649d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans/item/details"
    i3466366278330596cef17ffa7a24be314ef4f15e7e9d83c16fc122c5d7bbec9d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans/item/buckets/item"
    ie21edb59ce06faac589655d693ce3b2317a4cec27df5b0fb71362a772a5e4fbd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/planner/plans/item/tasks/item"
)

// PlannerPlanRequestBuilder builds and executes requests for operations under \me\planner\plans\{plannerPlan-id}
type PlannerPlanRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PlannerPlanRequestBuilderDeleteOptions options for Delete
type PlannerPlanRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PlannerPlanRequestBuilderGetOptions options for Get
type PlannerPlanRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PlannerPlanRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PlannerPlanRequestBuilderGetQueryParameters read-only. Nullable. Returns the plannerTasks assigned to the user.
type PlannerPlanRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PlannerPlanRequestBuilderPatchOptions options for Patch
type PlannerPlanRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerPlan;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PlannerPlanRequestBuilder) Buckets()(*i97f4aa0504a94677b48d17585be69f3adfec35ef126990f5c0fa1deb971970ee.BucketsRequestBuilder) {
    return i97f4aa0504a94677b48d17585be69f3adfec35ef126990f5c0fa1deb971970ee.NewBucketsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BucketsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.planner.plans.item.buckets.item collection
func (m *PlannerPlanRequestBuilder) BucketsById(id string)(*i3466366278330596cef17ffa7a24be314ef4f15e7e9d83c16fc122c5d7bbec9d.PlannerBucketRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerBucket_id"] = id
    }
    return i3466366278330596cef17ffa7a24be314ef4f15e7e9d83c16fc122c5d7bbec9d.NewPlannerBucketRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPlannerPlanRequestBuilderInternal instantiates a new PlannerPlanRequestBuilder and sets the default values.
func NewPlannerPlanRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerPlanRequestBuilder) {
    m := &PlannerPlanRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/planner/plans/{plannerPlan_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPlannerPlanRequestBuilder instantiates a new PlannerPlanRequestBuilder and sets the default values.
func NewPlannerPlanRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerPlanRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerPlanRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerPlanRequestBuilder) CreateDeleteRequestInformation(options *PlannerPlanRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerPlanRequestBuilder) CreateGetRequestInformation(options *PlannerPlanRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerPlanRequestBuilder) CreatePatchRequestInformation(options *PlannerPlanRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerPlanRequestBuilder) Delete(options *PlannerPlanRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *PlannerPlanRequestBuilder) Details()(*i9800485461c637fd964b8b604adecc9eb52f29415b0fdcbce97286ef735a649d.DetailsRequestBuilder) {
    return i9800485461c637fd964b8b604adecc9eb52f29415b0fdcbce97286ef735a649d.NewDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerPlanRequestBuilder) Get(options *PlannerPlanRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerPlan, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPlannerPlan() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerPlan), nil
}
// Patch read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerPlanRequestBuilder) Patch(options *PlannerPlanRequestBuilderPatchOptions)(error) {
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
func (m *PlannerPlanRequestBuilder) Tasks()(*i889297e93166312716592efbb05995455ae1882e388416fee06599864e54e462.TasksRequestBuilder) {
    return i889297e93166312716592efbb05995455ae1882e388416fee06599864e54e462.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.planner.plans.item.tasks.item collection
func (m *PlannerPlanRequestBuilder) TasksById(id string)(*ie21edb59ce06faac589655d693ce3b2317a4cec27df5b0fb71362a772a5e4fbd.PlannerTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask_id"] = id
    }
    return ie21edb59ce06faac589655d693ce3b2317a4cec27df5b0fb71362a772a5e4fbd.NewPlannerTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
