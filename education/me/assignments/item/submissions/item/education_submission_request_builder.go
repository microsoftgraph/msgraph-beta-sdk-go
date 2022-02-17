package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0f5790d63826a04042952fe0a2e6ce4da66a243be29f0123091a6f2dd74240d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item/return_escaped"
    i22195e779b2240c1f83dca2de3490bd1d1e13fa927f487556ff7ecbfd7b62375 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item/submit"
    ibcc8e1dc67b4ce1248f639c6349e995cf30c1f60e0a8a0f6a9268116b9d0bcc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item/outcomes"
    ibfa40cb573543c5c9e3abd6e7b436a68e2c2eb18727d422001d93fe33eacd8fc "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item/setupresourcesfolder"
    icfec2a32428842d30df10f523da9e4fc526b8b7c836071866938ed6be3843d1f "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item/submittedresources"
    id44a147cea37ad0c66060fd79a1e7b15bfef922b94b8f38814b921a5ddf5959b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item/resources"
    ie320e3441c7e9ca041b3cf19bcb344f0d5b5f9f316e71bdad73de83948ecfd9d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item/reassign"
    ifbd6f6420c2c3a622cde0cdd389b1186926a6ad859deff11e8e961ce84edf272 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item/unsubmit"
    i1385c8c70a4568d699eec3bf9f510616b7cc7e4ff29e686c620650de2efd46c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item/outcomes/item"
    ic303354d0ebbaf25e55517d0e5a70f2d884d74d3f80686736f2592b8b57dc23f "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item/resources/item"
    ie785db4b38f9f4ba6fc3e6cfb4164fd07a4b957a6f8e1b2b297b1e8dfeb71a05 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item/submissions/item/submittedresources/item"
)

// EducationSubmissionRequestBuilder builds and executes requests for operations under \education\me\assignments\{educationAssignment-id}\submissions\{educationSubmission-id}
type EducationSubmissionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationSubmissionRequestBuilderDeleteOptions options for Delete
type EducationSubmissionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationSubmissionRequestBuilderGetOptions options for Get
type EducationSubmissionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationSubmissionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationSubmissionRequestBuilderGetQueryParameters once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
type EducationSubmissionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationSubmissionRequestBuilderPatchOptions options for Patch
type EducationSubmissionRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmission;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewEducationSubmissionRequestBuilderInternal instantiates a new EducationSubmissionRequestBuilder and sets the default values.
func NewEducationSubmissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSubmissionRequestBuilder) {
    m := &EducationSubmissionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment_id}/submissions/{educationSubmission_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationSubmissionRequestBuilder instantiates a new EducationSubmissionRequestBuilder and sets the default values.
func NewEducationSubmissionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSubmissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSubmissionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionRequestBuilder) CreateDeleteRequestInformation(options *EducationSubmissionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionRequestBuilder) CreateGetRequestInformation(options *EducationSubmissionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionRequestBuilder) CreatePatchRequestInformation(options *EducationSubmissionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionRequestBuilder) Delete(options *EducationSubmissionRequestBuilderDeleteOptions)(error) {
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
// Get once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionRequestBuilder) Get(options *EducationSubmissionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmission, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEducationSubmission() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmission), nil
}
func (m *EducationSubmissionRequestBuilder) Outcomes()(*ibcc8e1dc67b4ce1248f639c6349e995cf30c1f60e0a8a0f6a9268116b9d0bcc2.OutcomesRequestBuilder) {
    return ibcc8e1dc67b4ce1248f639c6349e995cf30c1f60e0a8a0f6a9268116b9d0bcc2.NewOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OutcomesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.assignments.item.submissions.item.outcomes.item collection
func (m *EducationSubmissionRequestBuilder) OutcomesById(id string)(*i1385c8c70a4568d699eec3bf9f510616b7cc7e4ff29e686c620650de2efd46c7.EducationOutcomeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationOutcome_id"] = id
    }
    return i1385c8c70a4568d699eec3bf9f510616b7cc7e4ff29e686c620650de2efd46c7.NewEducationOutcomeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionRequestBuilder) Patch(options *EducationSubmissionRequestBuilderPatchOptions)(error) {
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
func (m *EducationSubmissionRequestBuilder) Reassign()(*ie320e3441c7e9ca041b3cf19bcb344f0d5b5f9f316e71bdad73de83948ecfd9d.ReassignRequestBuilder) {
    return ie320e3441c7e9ca041b3cf19bcb344f0d5b5f9f316e71bdad73de83948ecfd9d.NewReassignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Resources()(*id44a147cea37ad0c66060fd79a1e7b15bfef922b94b8f38814b921a5ddf5959b.ResourcesRequestBuilder) {
    return id44a147cea37ad0c66060fd79a1e7b15bfef922b94b8f38814b921a5ddf5959b.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.assignments.item.submissions.item.resources.item collection
func (m *EducationSubmissionRequestBuilder) ResourcesById(id string)(*ic303354d0ebbaf25e55517d0e5a70f2d884d74d3f80686736f2592b8b57dc23f.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return ic303354d0ebbaf25e55517d0e5a70f2d884d74d3f80686736f2592b8b57dc23f.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Return_escaped()(*i0f5790d63826a04042952fe0a2e6ce4da66a243be29f0123091a6f2dd74240d6.ReturnRequestBuilder) {
    return i0f5790d63826a04042952fe0a2e6ce4da66a243be29f0123091a6f2dd74240d6.NewReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SetUpResourcesFolder()(*ibfa40cb573543c5c9e3abd6e7b436a68e2c2eb18727d422001d93fe33eacd8fc.SetUpResourcesFolderRequestBuilder) {
    return ibfa40cb573543c5c9e3abd6e7b436a68e2c2eb18727d422001d93fe33eacd8fc.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Submit()(*i22195e779b2240c1f83dca2de3490bd1d1e13fa927f487556ff7ecbfd7b62375.SubmitRequestBuilder) {
    return i22195e779b2240c1f83dca2de3490bd1d1e13fa927f487556ff7ecbfd7b62375.NewSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SubmittedResources()(*icfec2a32428842d30df10f523da9e4fc526b8b7c836071866938ed6be3843d1f.SubmittedResourcesRequestBuilder) {
    return icfec2a32428842d30df10f523da9e4fc526b8b7c836071866938ed6be3843d1f.NewSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.assignments.item.submissions.item.submittedResources.item collection
func (m *EducationSubmissionRequestBuilder) SubmittedResourcesById(id string)(*ie785db4b38f9f4ba6fc3e6cfb4164fd07a4b957a6f8e1b2b297b1e8dfeb71a05.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return ie785db4b38f9f4ba6fc3e6cfb4164fd07a4b957a6f8e1b2b297b1e8dfeb71a05.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Unsubmit()(*ifbd6f6420c2c3a622cde0cdd389b1186926a6ad859deff11e8e961ce84edf272.UnsubmitRequestBuilder) {
    return ifbd6f6420c2c3a622cde0cdd389b1186926a6ad859deff11e8e961ce84edf272.NewUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
