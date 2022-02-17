package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1902f3a0ccc4b4d090900e03e4ad1c30b5e8404be560e8d2ae81ccf93c63bc9b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/setupresourcesfolder"
    i34f2686ac9e572a0aab5c6d3a460d4aab368203139a21a9e6676179cbc76f29d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions"
    i7aec1a2ebd08a150442dd3f2dab446f096282c891d32dee633746b4a74177b6d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/resources"
    i9f6aff0f0fb85711382cc277517e1973a86d4d579f3d7e2f1c175fad1b973d28 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/categories"
    iaaa8edd814077594dcfe08fabae894b6b945f70a45186a958cdb96a4f3c66924 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/rubric"
    iba118f4097921de3e02dbaeb597a1cf1cc65f08ed80fe4ffdd011bbdeb936ed2 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/publish"
    i5da31198b4698c9632f11b1d4cb2162d37d7be431577136dcb6670fb9876d353 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item"
    i7406f0fbcb81f97517ec46d1a396eee4f12972305c2bd1f4a72ae53516d8d8b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/resources/item"
    ib56c03b6b9d132ee7fc45579e6690dcf8e4f40924a1bccf9043eee4b1c8322d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/categories/item"
)

// EducationAssignmentRequestBuilder builds and executes requests for operations under \education\me\assignments\{educationAssignment-id}
type EducationAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationAssignmentRequestBuilderDeleteOptions options for Delete
type EducationAssignmentRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationAssignmentRequestBuilderGetOptions options for Get
type EducationAssignmentRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationAssignmentRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationAssignmentRequestBuilderGetQueryParameters assignments belonging to the user.
type EducationAssignmentRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationAssignmentRequestBuilderPatchOptions options for Patch
type EducationAssignmentRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationAssignment;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EducationAssignmentRequestBuilder) Categories()(*i9f6aff0f0fb85711382cc277517e1973a86d4d579f3d7e2f1c175fad1b973d28.CategoriesRequestBuilder) {
    return i9f6aff0f0fb85711382cc277517e1973a86d4d579f3d7e2f1c175fad1b973d28.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.assignments.item.categories.item collection
func (m *EducationAssignmentRequestBuilder) CategoriesById(id string)(*ib56c03b6b9d132ee7fc45579e6690dcf8e4f40924a1bccf9043eee4b1c8322d8.EducationCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory_id"] = id
    }
    return ib56c03b6b9d132ee7fc45579e6690dcf8e4f40924a1bccf9043eee4b1c8322d8.NewEducationCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationAssignmentRequestBuilderInternal instantiates a new EducationAssignmentRequestBuilder and sets the default values.
func NewEducationAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    m := &EducationAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationAssignmentRequestBuilder instantiates a new EducationAssignmentRequestBuilder and sets the default values.
func NewEducationAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation assignments belonging to the user.
func (m *EducationAssignmentRequestBuilder) CreateDeleteRequestInformation(options *EducationAssignmentRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation assignments belonging to the user.
func (m *EducationAssignmentRequestBuilder) CreateGetRequestInformation(options *EducationAssignmentRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation assignments belonging to the user.
func (m *EducationAssignmentRequestBuilder) CreatePatchRequestInformation(options *EducationAssignmentRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete assignments belonging to the user.
func (m *EducationAssignmentRequestBuilder) Delete(options *EducationAssignmentRequestBuilderDeleteOptions)(error) {
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
// Get assignments belonging to the user.
func (m *EducationAssignmentRequestBuilder) Get(options *EducationAssignmentRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEducationAssignment() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationAssignment), nil
}
// Patch assignments belonging to the user.
func (m *EducationAssignmentRequestBuilder) Patch(options *EducationAssignmentRequestBuilderPatchOptions)(error) {
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
func (m *EducationAssignmentRequestBuilder) Publish()(*iba118f4097921de3e02dbaeb597a1cf1cc65f08ed80fe4ffdd011bbdeb936ed2.PublishRequestBuilder) {
    return iba118f4097921de3e02dbaeb597a1cf1cc65f08ed80fe4ffdd011bbdeb936ed2.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Resources()(*i7aec1a2ebd08a150442dd3f2dab446f096282c891d32dee633746b4a74177b6d.ResourcesRequestBuilder) {
    return i7aec1a2ebd08a150442dd3f2dab446f096282c891d32dee633746b4a74177b6d.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.assignments.item.resources.item collection
func (m *EducationAssignmentRequestBuilder) ResourcesById(id string)(*i7406f0fbcb81f97517ec46d1a396eee4f12972305c2bd1f4a72ae53516d8d8b3.EducationAssignmentResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource_id"] = id
    }
    return i7406f0fbcb81f97517ec46d1a396eee4f12972305c2bd1f4a72ae53516d8d8b3.NewEducationAssignmentResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Rubric()(*iaaa8edd814077594dcfe08fabae894b6b945f70a45186a958cdb96a4f3c66924.RubricRequestBuilder) {
    return iaaa8edd814077594dcfe08fabae894b6b945f70a45186a958cdb96a4f3c66924.NewRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) SetUpResourcesFolder()(*i1902f3a0ccc4b4d090900e03e4ad1c30b5e8404be560e8d2ae81ccf93c63bc9b.SetUpResourcesFolderRequestBuilder) {
    return i1902f3a0ccc4b4d090900e03e4ad1c30b5e8404be560e8d2ae81ccf93c63bc9b.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Submissions()(*i34f2686ac9e572a0aab5c6d3a460d4aab368203139a21a9e6676179cbc76f29d.SubmissionsRequestBuilder) {
    return i34f2686ac9e572a0aab5c6d3a460d4aab368203139a21a9e6676179cbc76f29d.NewSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.assignments.item.submissions.item collection
func (m *EducationAssignmentRequestBuilder) SubmissionsById(id string)(*i5da31198b4698c9632f11b1d4cb2162d37d7be431577136dcb6670fb9876d353.EducationSubmissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission_id"] = id
    }
    return i5da31198b4698c9632f11b1d4cb2162d37d7be431577136dcb6670fb9876d353.NewEducationSubmissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
