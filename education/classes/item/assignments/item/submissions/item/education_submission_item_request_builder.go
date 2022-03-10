package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i1928a13dc54c73d5c9d9c1da4d9b70aeed20716bc8d9c77536a95f1b2bc0a66c "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/resources"
    i1e3ae53018fa05de94186b37ef1966bcb4c6d4dce5ae955977fa68e30f2ca832 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/return_escaped"
    i43ba3ee4b1fe518a6d013e2f0b889058897248ca13cf8fd56f3a4e3f6dbcf726 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/outcomes"
    i572b649f92c4e714badac0d19fea01c4a6034be3c6051339c8788370fc5db763 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/reassign"
    i63c795cc15fa99fd56c532c984b57d606b0ce50ae0f0462a8f05f2c97403505a "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/submittedresources"
    i768987a34c53877b80bd9b75ccfde8cec3078a7b79600a3541157f92a02bd7ae "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/submit"
    iad0e46857ffe4aca71c347ec356f1f060d1acbfe420d89db1d61dd5f39ae6479 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/setupresourcesfolder"
    ie5fd4dbd2c6ed61da731d9780ed8f1ec6a8a6533c73bce16ca0f6e33a6b8cf4d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/unsubmit"
    i6d2213e16b46c89802a19dd596bafab2a2de02beb8b6102d680539d7de269026 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/resources/item"
    i96ff390efa3e1cf5f1c7b76015acbcac008ccfe65211d104227d3786192347c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/outcomes/item"
    ic22ac8d3bfbbeab72b02b700408b7d279cdca486676844964c112e35bc79e358 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/submittedresources/item"
)

// EducationSubmissionItemRequestBuilder provides operations to manage the submissions property of the microsoft.graph.educationAssignment entity.
type EducationSubmissionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationSubmissionItemRequestBuilderDeleteOptions options for Delete
type EducationSubmissionItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationSubmissionItemRequestBuilderGetOptions options for Get
type EducationSubmissionItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationSubmissionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationSubmissionItemRequestBuilderGetQueryParameters once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
type EducationSubmissionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationSubmissionItemRequestBuilderPatchOptions options for Patch
type EducationSubmissionItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmissionable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewEducationSubmissionItemRequestBuilderInternal instantiates a new EducationSubmissionItemRequestBuilder and sets the default values.
func NewEducationSubmissionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSubmissionItemRequestBuilder) {
    m := &EducationSubmissionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass_id}/assignments/{educationAssignment_id}/submissions/{educationSubmission_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationSubmissionItemRequestBuilder instantiates a new EducationSubmissionItemRequestBuilder and sets the default values.
func NewEducationSubmissionItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSubmissionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSubmissionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property submissions for education
func (m *EducationSubmissionItemRequestBuilder) CreateDeleteRequestInformation(options *EducationSubmissionItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationSubmissionItemRequestBuilder) CreateGetRequestInformation(options *EducationSubmissionItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property submissions in education
func (m *EducationSubmissionItemRequestBuilder) CreatePatchRequestInformation(options *EducationSubmissionItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property submissions for education
func (m *EducationSubmissionItemRequestBuilder) Delete(options *EducationSubmissionItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionItemRequestBuilder) Get(options *EducationSubmissionItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmissionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEducationSubmissionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmissionable), nil
}
func (m *EducationSubmissionItemRequestBuilder) Outcomes()(*i43ba3ee4b1fe518a6d013e2f0b889058897248ca13cf8fd56f3a4e3f6dbcf726.OutcomesRequestBuilder) {
    return i43ba3ee4b1fe518a6d013e2f0b889058897248ca13cf8fd56f3a4e3f6dbcf726.NewOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OutcomesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.classes.item.assignments.item.submissions.item.outcomes.item collection
func (m *EducationSubmissionItemRequestBuilder) OutcomesById(id string)(*i96ff390efa3e1cf5f1c7b76015acbcac008ccfe65211d104227d3786192347c4.EducationOutcomeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationOutcome_id"] = id
    }
    return i96ff390efa3e1cf5f1c7b76015acbcac008ccfe65211d104227d3786192347c4.NewEducationOutcomeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property submissions in education
func (m *EducationSubmissionItemRequestBuilder) Patch(options *EducationSubmissionItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationSubmissionItemRequestBuilder) Reassign()(*i572b649f92c4e714badac0d19fea01c4a6034be3c6051339c8788370fc5db763.ReassignRequestBuilder) {
    return i572b649f92c4e714badac0d19fea01c4a6034be3c6051339c8788370fc5db763.NewReassignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionItemRequestBuilder) Resources()(*i1928a13dc54c73d5c9d9c1da4d9b70aeed20716bc8d9c77536a95f1b2bc0a66c.ResourcesRequestBuilder) {
    return i1928a13dc54c73d5c9d9c1da4d9b70aeed20716bc8d9c77536a95f1b2bc0a66c.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.classes.item.assignments.item.submissions.item.resources.item collection
func (m *EducationSubmissionItemRequestBuilder) ResourcesById(id string)(*i6d2213e16b46c89802a19dd596bafab2a2de02beb8b6102d680539d7de269026.EducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return i6d2213e16b46c89802a19dd596bafab2a2de02beb8b6102d680539d7de269026.NewEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionItemRequestBuilder) Return_escaped()(*i1e3ae53018fa05de94186b37ef1966bcb4c6d4dce5ae955977fa68e30f2ca832.ReturnRequestBuilder) {
    return i1e3ae53018fa05de94186b37ef1966bcb4c6d4dce5ae955977fa68e30f2ca832.NewReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionItemRequestBuilder) SetUpResourcesFolder()(*iad0e46857ffe4aca71c347ec356f1f060d1acbfe420d89db1d61dd5f39ae6479.SetUpResourcesFolderRequestBuilder) {
    return iad0e46857ffe4aca71c347ec356f1f060d1acbfe420d89db1d61dd5f39ae6479.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionItemRequestBuilder) Submit()(*i768987a34c53877b80bd9b75ccfde8cec3078a7b79600a3541157f92a02bd7ae.SubmitRequestBuilder) {
    return i768987a34c53877b80bd9b75ccfde8cec3078a7b79600a3541157f92a02bd7ae.NewSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionItemRequestBuilder) SubmittedResources()(*i63c795cc15fa99fd56c532c984b57d606b0ce50ae0f0462a8f05f2c97403505a.SubmittedResourcesRequestBuilder) {
    return i63c795cc15fa99fd56c532c984b57d606b0ce50ae0f0462a8f05f2c97403505a.NewSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.classes.item.assignments.item.submissions.item.submittedResources.item collection
func (m *EducationSubmissionItemRequestBuilder) SubmittedResourcesById(id string)(*ic22ac8d3bfbbeab72b02b700408b7d279cdca486676844964c112e35bc79e358.EducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return ic22ac8d3bfbbeab72b02b700408b7d279cdca486676844964c112e35bc79e358.NewEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionItemRequestBuilder) Unsubmit()(*ie5fd4dbd2c6ed61da731d9780ed8f1ec6a8a6533c73bce16ca0f6e33a6b8cf4d.UnsubmitRequestBuilder) {
    return ie5fd4dbd2c6ed61da731d9780ed8f1ec6a8a6533c73bce16ca0f6e33a6b8cf4d.NewUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
