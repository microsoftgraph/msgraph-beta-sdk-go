package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2243091c8fa68ba9e7a97bce329cf8353df512034105a6cdb60022d8a3462d3f "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters/item/plans"
    i9ad7de99f1b757c4ff16ccef4d347a843f114b88ad7c9579c76ffa1dd8f293f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters/item/members"
    ie53712fa7eb3bcb49fdf48153b4145c0918f55b3d8a9851bfbf8ce2045244c3c "github.com/microsoftgraph/msgraph-beta-sdk-go/planner/rosters/item/members/item"
)

// PlannerRosterRequestBuilder builds and executes requests for operations under \planner\rosters\{plannerRoster-id}
type PlannerRosterRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PlannerRosterRequestBuilderDeleteOptions options for Delete
type PlannerRosterRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PlannerRosterRequestBuilderGetOptions options for Get
type PlannerRosterRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PlannerRosterRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PlannerRosterRequestBuilderGetQueryParameters read-only. Nullable. Returns a collection of the specified rosters
type PlannerRosterRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PlannerRosterRequestBuilderPatchOptions options for Patch
type PlannerRosterRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerRoster;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewPlannerRosterRequestBuilderInternal instantiates a new PlannerRosterRequestBuilder and sets the default values.
func NewPlannerRosterRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRosterRequestBuilder) {
    m := &PlannerRosterRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/planner/rosters/{plannerRoster_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPlannerRosterRequestBuilder instantiates a new PlannerRosterRequestBuilder and sets the default values.
func NewPlannerRosterRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRosterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerRosterRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation read-only. Nullable. Returns a collection of the specified rosters
func (m *PlannerRosterRequestBuilder) CreateDeleteRequestInformation(options *PlannerRosterRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable. Returns a collection of the specified rosters
func (m *PlannerRosterRequestBuilder) CreateGetRequestInformation(options *PlannerRosterRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation read-only. Nullable. Returns a collection of the specified rosters
func (m *PlannerRosterRequestBuilder) CreatePatchRequestInformation(options *PlannerRosterRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete read-only. Nullable. Returns a collection of the specified rosters
func (m *PlannerRosterRequestBuilder) Delete(options *PlannerRosterRequestBuilderDeleteOptions)(error) {
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
// Get read-only. Nullable. Returns a collection of the specified rosters
func (m *PlannerRosterRequestBuilder) Get(options *PlannerRosterRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerRoster, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPlannerRoster() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PlannerRoster), nil
}
func (m *PlannerRosterRequestBuilder) Members()(*i9ad7de99f1b757c4ff16ccef4d347a843f114b88ad7c9579c76ffa1dd8f293f9.MembersRequestBuilder) {
    return i9ad7de99f1b757c4ff16ccef4d347a843f114b88ad7c9579c76ffa1dd8f293f9.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.planner.rosters.item.members.item collection
func (m *PlannerRosterRequestBuilder) MembersById(id string)(*ie53712fa7eb3bcb49fdf48153b4145c0918f55b3d8a9851bfbf8ce2045244c3c.PlannerRosterMemberRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerRosterMember_id"] = id
    }
    return ie53712fa7eb3bcb49fdf48153b4145c0918f55b3d8a9851bfbf8ce2045244c3c.NewPlannerRosterMemberRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch read-only. Nullable. Returns a collection of the specified rosters
func (m *PlannerRosterRequestBuilder) Patch(options *PlannerRosterRequestBuilderPatchOptions)(error) {
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
func (m *PlannerRosterRequestBuilder) Plans()(*i2243091c8fa68ba9e7a97bce329cf8353df512034105a6cdb60022d8a3462d3f.PlansRequestBuilder) {
    return i2243091c8fa68ba9e7a97bce329cf8353df512034105a6cdb60022d8a3462d3f.NewPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
