package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i05ccbbf4c4b0d0591be23572d5c88285d13f488480ce9e147056e16e74647839 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/plans/item/tasks/item/progresstaskboardformat"
    i40fcec495cfdeb183cf916256a9c69db16c82a441a3ddaa64ef988641a44602f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/plans/item/tasks/item/assignedtotaskboardformat"
    i80130588d58e1a703e0a62e4536c4fc962d53849d49502392a890c56328102f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/plans/item/tasks/item/buckettaskboardformat"
    iff8792866487e10b2b51f1d551bd01106c36d7482105172419d099c515bc14f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/planner/plans/item/tasks/item/details"
)

// PlannerTaskRequestBuilder builds and executes requests for operations under \users\{user-id}\planner\plans\{plannerPlan-id}\tasks\{plannerTask-id}
type PlannerTaskRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PlannerTaskRequestBuilderDeleteOptions options for Delete
type PlannerTaskRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PlannerTaskRequestBuilderGetOptions options for Get
type PlannerTaskRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PlannerTaskRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PlannerTaskRequestBuilderGetQueryParameters read-only. Nullable. Collection of tasks in the plan.
type PlannerTaskRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PlannerTaskRequestBuilderPatchOptions options for Patch
type PlannerTaskRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerTask;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PlannerTaskRequestBuilder) AssignedToTaskBoardFormat()(*i40fcec495cfdeb183cf916256a9c69db16c82a441a3ddaa64ef988641a44602f.AssignedToTaskBoardFormatRequestBuilder) {
    return i40fcec495cfdeb183cf916256a9c69db16c82a441a3ddaa64ef988641a44602f.NewAssignedToTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerTaskRequestBuilder) BucketTaskBoardFormat()(*i80130588d58e1a703e0a62e4536c4fc962d53849d49502392a890c56328102f6.BucketTaskBoardFormatRequestBuilder) {
    return i80130588d58e1a703e0a62e4536c4fc962d53849d49502392a890c56328102f6.NewBucketTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPlannerTaskRequestBuilderInternal instantiates a new PlannerTaskRequestBuilder and sets the default values.
func NewPlannerTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerTaskRequestBuilder) {
    m := &PlannerTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/planner/plans/{plannerPlan_id}/tasks/{plannerTask_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPlannerTaskRequestBuilder instantiates a new PlannerTaskRequestBuilder and sets the default values.
func NewPlannerTaskRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerTaskRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation read-only. Nullable. Collection of tasks in the plan.
func (m *PlannerTaskRequestBuilder) CreateDeleteRequestInformation(options *PlannerTaskRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable. Collection of tasks in the plan.
func (m *PlannerTaskRequestBuilder) CreateGetRequestInformation(options *PlannerTaskRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation read-only. Nullable. Collection of tasks in the plan.
func (m *PlannerTaskRequestBuilder) CreatePatchRequestInformation(options *PlannerTaskRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete read-only. Nullable. Collection of tasks in the plan.
func (m *PlannerTaskRequestBuilder) Delete(options *PlannerTaskRequestBuilderDeleteOptions)(error) {
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
func (m *PlannerTaskRequestBuilder) Details()(*iff8792866487e10b2b51f1d551bd01106c36d7482105172419d099c515bc14f5.DetailsRequestBuilder) {
    return iff8792866487e10b2b51f1d551bd01106c36d7482105172419d099c515bc14f5.NewDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read-only. Nullable. Collection of tasks in the plan.
func (m *PlannerTaskRequestBuilder) Get(options *PlannerTaskRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerTask, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPlannerTask() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerTask), nil
}
// Patch read-only. Nullable. Collection of tasks in the plan.
func (m *PlannerTaskRequestBuilder) Patch(options *PlannerTaskRequestBuilderPatchOptions)(error) {
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
func (m *PlannerTaskRequestBuilder) ProgressTaskBoardFormat()(*i05ccbbf4c4b0d0591be23572d5c88285d13f488480ce9e147056e16e74647839.ProgressTaskBoardFormatRequestBuilder) {
    return i05ccbbf4c4b0d0591be23572d5c88285d13f488480ce9e147056e16e74647839.NewProgressTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
