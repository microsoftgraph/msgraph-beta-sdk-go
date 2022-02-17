package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i03a5515c9c388939cec0c62975b5200f7cc84414e6ff4962c81d48008588c4ca "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/submit"
    i164979e9c6e40c75579e19bae874710bbe54a4757fa82f177ce989d7159b8115 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/submittedresources"
    i29e598c0e3132502e75ebc5f97d630c0d489e1a3695a279765b136eb63851edb "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/unsubmit"
    i620d352f6d923470da79c3942b8c8aaaf958d8aaa2e3d4eb26af40a92833870f "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/resources"
    i6ba46d6835c0a47a27d09f0253c579859dcd40873b53932a91d1e8d864ea4af4 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/outcomes"
    i7dd6fe92e931b3075c9a3419ab371f02a77ec29b4fe853adf3e476b895fc4af4 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/return_escaped"
    id4a8c5d09c4fd819ff61848c117c7c2839a0383c023c9c4a1c49e3ab17c1adb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/reassign"
    if15e4fd1817bfddf93807ec99a8a0cbce18f2c72a43161efa316e4ff2ad07585 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/setupresourcesfolder"
    idcb18db22001d0c8560f4e04f592d2ab741e41fdccd32e2bbd2826405e8162b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/outcomes/item"
    ided68156635ca8883976581f28a920199cc5242c12a0b15c16f989f304e5824b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/resources/item"
    ie8d02aac9f22a617e7db20c68df056fad4bb1fb1425faaa25cf4ef67594043ec "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/submissions/item/submittedresources/item"
)

// EducationSubmissionRequestBuilder builds and executes requests for operations under \education\users\{educationUser-id}\assignments\{educationAssignment-id}\submissions\{educationSubmission-id}
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
    m.urlTemplate = "{+baseurl}/education/users/{educationUser_id}/assignments/{educationAssignment_id}/submissions/{educationSubmission_id}{?select,expand}";
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
func (m *EducationSubmissionRequestBuilder) Outcomes()(*i6ba46d6835c0a47a27d09f0253c579859dcd40873b53932a91d1e8d864ea4af4.OutcomesRequestBuilder) {
    return i6ba46d6835c0a47a27d09f0253c579859dcd40873b53932a91d1e8d864ea4af4.NewOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OutcomesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item.assignments.item.submissions.item.outcomes.item collection
func (m *EducationSubmissionRequestBuilder) OutcomesById(id string)(*idcb18db22001d0c8560f4e04f592d2ab741e41fdccd32e2bbd2826405e8162b6.EducationOutcomeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationOutcome_id"] = id
    }
    return idcb18db22001d0c8560f4e04f592d2ab741e41fdccd32e2bbd2826405e8162b6.NewEducationOutcomeRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EducationSubmissionRequestBuilder) Reassign()(*id4a8c5d09c4fd819ff61848c117c7c2839a0383c023c9c4a1c49e3ab17c1adb7.ReassignRequestBuilder) {
    return id4a8c5d09c4fd819ff61848c117c7c2839a0383c023c9c4a1c49e3ab17c1adb7.NewReassignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Resources()(*i620d352f6d923470da79c3942b8c8aaaf958d8aaa2e3d4eb26af40a92833870f.ResourcesRequestBuilder) {
    return i620d352f6d923470da79c3942b8c8aaaf958d8aaa2e3d4eb26af40a92833870f.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item.assignments.item.submissions.item.resources.item collection
func (m *EducationSubmissionRequestBuilder) ResourcesById(id string)(*ided68156635ca8883976581f28a920199cc5242c12a0b15c16f989f304e5824b.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return ided68156635ca8883976581f28a920199cc5242c12a0b15c16f989f304e5824b.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Return_escaped()(*i7dd6fe92e931b3075c9a3419ab371f02a77ec29b4fe853adf3e476b895fc4af4.ReturnRequestBuilder) {
    return i7dd6fe92e931b3075c9a3419ab371f02a77ec29b4fe853adf3e476b895fc4af4.NewReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SetUpResourcesFolder()(*if15e4fd1817bfddf93807ec99a8a0cbce18f2c72a43161efa316e4ff2ad07585.SetUpResourcesFolderRequestBuilder) {
    return if15e4fd1817bfddf93807ec99a8a0cbce18f2c72a43161efa316e4ff2ad07585.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Submit()(*i03a5515c9c388939cec0c62975b5200f7cc84414e6ff4962c81d48008588c4ca.SubmitRequestBuilder) {
    return i03a5515c9c388939cec0c62975b5200f7cc84414e6ff4962c81d48008588c4ca.NewSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SubmittedResources()(*i164979e9c6e40c75579e19bae874710bbe54a4757fa82f177ce989d7159b8115.SubmittedResourcesRequestBuilder) {
    return i164979e9c6e40c75579e19bae874710bbe54a4757fa82f177ce989d7159b8115.NewSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item.assignments.item.submissions.item.submittedResources.item collection
func (m *EducationSubmissionRequestBuilder) SubmittedResourcesById(id string)(*ie8d02aac9f22a617e7db20c68df056fad4bb1fb1425faaa25cf4ef67594043ec.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return ie8d02aac9f22a617e7db20c68df056fad4bb1fb1425faaa25cf4ef67594043ec.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Unsubmit()(*i29e598c0e3132502e75ebc5f97d630c0d489e1a3695a279765b136eb63851edb.UnsubmitRequestBuilder) {
    return i29e598c0e3132502e75ebc5f97d630c0d489e1a3695a279765b136eb63851edb.NewUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
