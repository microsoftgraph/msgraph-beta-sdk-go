package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5b10b16160956c55311bec6d730fa2056fd37872e249abeca309071996a8fc0b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/resources"
    i8050f70ff077aa7a00ee30e1a3c86c0f0615a85e77ec2d73dc4e1bafa3f4e16f "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/setupresourcesfolder"
    ib73813a5d39b89f70827c0fa1818289fccec256ca0ea65df5135c2bf448aa519 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/rubric"
    ic194e429a42abea4fc9457129ec19b14e4af236b99092c2a55f726379f0ecc39 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions"
    idf7d52d5e40c4a2f0a279dbefcf41f586915062fc10338040bd70c8bde7211d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/publish"
    ie3e3ed607d0779e161141593c134093354a4e6852436b9197a40b9ac9a01eb0e "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/categories"
    i105c230efb20cac97a293b5469f4f607fe6cb947c167d5ee426b70bce49e9df5 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item"
    ibe13b36acf757c0c212481dd71384ef998a32ff3577f6c574e9f1785f7bb86f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/resources/item"
    iee68e3ee77c1c1596d86d9cdb46e4cba567ac32c664bddc07320057cef108c0c "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/categories/item"
)

// EducationAssignmentRequestBuilder builds and executes requests for operations under \education\users\{educationUser-id}\assignments\{educationAssignment-id}
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
func (m *EducationAssignmentRequestBuilder) Categories()(*ie3e3ed607d0779e161141593c134093354a4e6852436b9197a40b9ac9a01eb0e.CategoriesRequestBuilder) {
    return ie3e3ed607d0779e161141593c134093354a4e6852436b9197a40b9ac9a01eb0e.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item.assignments.item.categories.item collection
func (m *EducationAssignmentRequestBuilder) CategoriesById(id string)(*iee68e3ee77c1c1596d86d9cdb46e4cba567ac32c664bddc07320057cef108c0c.EducationCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory_id"] = id
    }
    return iee68e3ee77c1c1596d86d9cdb46e4cba567ac32c664bddc07320057cef108c0c.NewEducationCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationAssignmentRequestBuilderInternal instantiates a new EducationAssignmentRequestBuilder and sets the default values.
func NewEducationAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    m := &EducationAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser_id}/assignments/{educationAssignment_id}{?select,expand}";
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
func (m *EducationAssignmentRequestBuilder) Publish()(*idf7d52d5e40c4a2f0a279dbefcf41f586915062fc10338040bd70c8bde7211d7.PublishRequestBuilder) {
    return idf7d52d5e40c4a2f0a279dbefcf41f586915062fc10338040bd70c8bde7211d7.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Resources()(*i5b10b16160956c55311bec6d730fa2056fd37872e249abeca309071996a8fc0b.ResourcesRequestBuilder) {
    return i5b10b16160956c55311bec6d730fa2056fd37872e249abeca309071996a8fc0b.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item.assignments.item.resources.item collection
func (m *EducationAssignmentRequestBuilder) ResourcesById(id string)(*ibe13b36acf757c0c212481dd71384ef998a32ff3577f6c574e9f1785f7bb86f2.EducationAssignmentResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource_id"] = id
    }
    return ibe13b36acf757c0c212481dd71384ef998a32ff3577f6c574e9f1785f7bb86f2.NewEducationAssignmentResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Rubric()(*ib73813a5d39b89f70827c0fa1818289fccec256ca0ea65df5135c2bf448aa519.RubricRequestBuilder) {
    return ib73813a5d39b89f70827c0fa1818289fccec256ca0ea65df5135c2bf448aa519.NewRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) SetUpResourcesFolder()(*i8050f70ff077aa7a00ee30e1a3c86c0f0615a85e77ec2d73dc4e1bafa3f4e16f.SetUpResourcesFolderRequestBuilder) {
    return i8050f70ff077aa7a00ee30e1a3c86c0f0615a85e77ec2d73dc4e1bafa3f4e16f.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Submissions()(*ic194e429a42abea4fc9457129ec19b14e4af236b99092c2a55f726379f0ecc39.SubmissionsRequestBuilder) {
    return ic194e429a42abea4fc9457129ec19b14e4af236b99092c2a55f726379f0ecc39.NewSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item.assignments.item.submissions.item collection
func (m *EducationAssignmentRequestBuilder) SubmissionsById(id string)(*i105c230efb20cac97a293b5469f4f607fe6cb947c167d5ee426b70bce49e9df5.EducationSubmissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission_id"] = id
    }
    return i105c230efb20cac97a293b5469f4f607fe6cb947c167d5ee426b70bce49e9df5.NewEducationSubmissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
