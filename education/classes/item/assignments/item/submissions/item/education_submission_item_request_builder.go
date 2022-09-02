package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EducationSubmissionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationSubmissionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EducationSubmissionItemRequestBuilderGetQueryParameters once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
type EducationSubmissionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EducationSubmissionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationSubmissionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EducationSubmissionItemRequestBuilderGetQueryParameters
}
// EducationSubmissionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationSubmissionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEducationSubmissionItemRequestBuilderInternal instantiates a new EducationSubmissionItemRequestBuilder and sets the default values.
func NewEducationSubmissionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationSubmissionItemRequestBuilder) {
    m := &EducationSubmissionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationSubmissionItemRequestBuilder instantiates a new EducationSubmissionItemRequestBuilder and sets the default values.
func NewEducationSubmissionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationSubmissionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSubmissionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property submissions for education
func (m *EducationSubmissionItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property submissions for education
func (m *EducationSubmissionItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EducationSubmissionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EducationSubmissionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property submissions in education
func (m *EducationSubmissionItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property submissions in education
func (m *EducationSubmissionItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable, requestConfiguration *EducationSubmissionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property submissions for education
func (m *EducationSubmissionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EducationSubmissionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EducationSubmissionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationSubmissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable), nil
}
// Outcomes the outcomes property
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
        urlTplParams["educationOutcome%2Did"] = id
    }
    return i96ff390efa3e1cf5f1c7b76015acbcac008ccfe65211d104227d3786192347c4.NewEducationOutcomeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property submissions in education
func (m *EducationSubmissionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable, requestConfiguration *EducationSubmissionItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Reassign the reassign property
func (m *EducationSubmissionItemRequestBuilder) Reassign()(*i572b649f92c4e714badac0d19fea01c4a6034be3c6051339c8788370fc5db763.ReassignRequestBuilder) {
    return i572b649f92c4e714badac0d19fea01c4a6034be3c6051339c8788370fc5db763.NewReassignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Resources the resources property
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
        urlTplParams["educationSubmissionResource%2Did"] = id
    }
    return i6d2213e16b46c89802a19dd596bafab2a2de02beb8b6102d680539d7de269026.NewEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Return_escaped the return property
func (m *EducationSubmissionItemRequestBuilder) Return_escaped()(*i1e3ae53018fa05de94186b37ef1966bcb4c6d4dce5ae955977fa68e30f2ca832.ReturnRequestBuilder) {
    return i1e3ae53018fa05de94186b37ef1966bcb4c6d4dce5ae955977fa68e30f2ca832.NewReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUpResourcesFolder the setUpResourcesFolder property
func (m *EducationSubmissionItemRequestBuilder) SetUpResourcesFolder()(*iad0e46857ffe4aca71c347ec356f1f060d1acbfe420d89db1d61dd5f39ae6479.SetUpResourcesFolderRequestBuilder) {
    return iad0e46857ffe4aca71c347ec356f1f060d1acbfe420d89db1d61dd5f39ae6479.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Submit the submit property
func (m *EducationSubmissionItemRequestBuilder) Submit()(*i768987a34c53877b80bd9b75ccfde8cec3078a7b79600a3541157f92a02bd7ae.SubmitRequestBuilder) {
    return i768987a34c53877b80bd9b75ccfde8cec3078a7b79600a3541157f92a02bd7ae.NewSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResources the submittedResources property
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
        urlTplParams["educationSubmissionResource%2Did"] = id
    }
    return ic22ac8d3bfbbeab72b02b700408b7d279cdca486676844964c112e35bc79e358.NewEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unsubmit the unsubmit property
func (m *EducationSubmissionItemRequestBuilder) Unsubmit()(*ie5fd4dbd2c6ed61da731d9780ed8f1ec6a8a6533c73bce16ca0f6e33a6b8cf4d.UnsubmitRequestBuilder) {
    return ie5fd4dbd2c6ed61da731d9780ed8f1ec6a8a6533c73bce16ca0f6e33a6b8cf4d.NewUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
