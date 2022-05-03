package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property submissions for education
func (m *EducationSubmissionItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EducationSubmissionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property submissions in education
func (m *EducationSubmissionItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property submissions for education
func (m *EducationSubmissionItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *EducationSubmissionItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property submissions for education
func (m *EducationSubmissionItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *EducationSubmissionItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// GetWithResponseHandler once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionItemRequestBuilder) GetWithResponseHandler(requestConfiguration *EducationSubmissionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationSubmissionItemRequestBuilder) GetWithResponseHandler(requestConfiguration *EducationSubmissionItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationSubmissionFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable), nil
}
// Outcomes the outcomes property
func (m *EducationSubmissionItemRequestBuilder) Outcomes()(*ibcc8e1dc67b4ce1248f639c6349e995cf30c1f60e0a8a0f6a9268116b9d0bcc2.OutcomesRequestBuilder) {
    return ibcc8e1dc67b4ce1248f639c6349e995cf30c1f60e0a8a0f6a9268116b9d0bcc2.NewOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OutcomesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.assignments.item.submissions.item.outcomes.item collection
func (m *EducationSubmissionItemRequestBuilder) OutcomesById(id string)(*i1385c8c70a4568d699eec3bf9f510616b7cc7e4ff29e686c620650de2efd46c7.EducationOutcomeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationOutcome%2Did"] = id
    }
    return i1385c8c70a4568d699eec3bf9f510616b7cc7e4ff29e686c620650de2efd46c7.NewEducationOutcomeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property submissions in education
func (m *EducationSubmissionItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable, requestConfiguration *EducationSubmissionItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property submissions in education
func (m *EducationSubmissionItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationSubmissionable, requestConfiguration *EducationSubmissionItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Reassign the reassign property
func (m *EducationSubmissionItemRequestBuilder) Reassign()(*ie320e3441c7e9ca041b3cf19bcb344f0d5b5f9f316e71bdad73de83948ecfd9d.ReassignRequestBuilder) {
    return ie320e3441c7e9ca041b3cf19bcb344f0d5b5f9f316e71bdad73de83948ecfd9d.NewReassignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Resources the resources property
func (m *EducationSubmissionItemRequestBuilder) Resources()(*id44a147cea37ad0c66060fd79a1e7b15bfef922b94b8f38814b921a5ddf5959b.ResourcesRequestBuilder) {
    return id44a147cea37ad0c66060fd79a1e7b15bfef922b94b8f38814b921a5ddf5959b.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.assignments.item.submissions.item.resources.item collection
func (m *EducationSubmissionItemRequestBuilder) ResourcesById(id string)(*ic303354d0ebbaf25e55517d0e5a70f2d884d74d3f80686736f2592b8b57dc23f.EducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource%2Did"] = id
    }
    return ic303354d0ebbaf25e55517d0e5a70f2d884d74d3f80686736f2592b8b57dc23f.NewEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Return_escaped the return property
func (m *EducationSubmissionItemRequestBuilder) Return_escaped()(*i0f5790d63826a04042952fe0a2e6ce4da66a243be29f0123091a6f2dd74240d6.ReturnRequestBuilder) {
    return i0f5790d63826a04042952fe0a2e6ce4da66a243be29f0123091a6f2dd74240d6.NewReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUpResourcesFolder the setUpResourcesFolder property
func (m *EducationSubmissionItemRequestBuilder) SetUpResourcesFolder()(*ibfa40cb573543c5c9e3abd6e7b436a68e2c2eb18727d422001d93fe33eacd8fc.SetUpResourcesFolderRequestBuilder) {
    return ibfa40cb573543c5c9e3abd6e7b436a68e2c2eb18727d422001d93fe33eacd8fc.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Submit the submit property
func (m *EducationSubmissionItemRequestBuilder) Submit()(*i22195e779b2240c1f83dca2de3490bd1d1e13fa927f487556ff7ecbfd7b62375.SubmitRequestBuilder) {
    return i22195e779b2240c1f83dca2de3490bd1d1e13fa927f487556ff7ecbfd7b62375.NewSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResources the submittedResources property
func (m *EducationSubmissionItemRequestBuilder) SubmittedResources()(*icfec2a32428842d30df10f523da9e4fc526b8b7c836071866938ed6be3843d1f.SubmittedResourcesRequestBuilder) {
    return icfec2a32428842d30df10f523da9e4fc526b8b7c836071866938ed6be3843d1f.NewSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmittedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.me.assignments.item.submissions.item.submittedResources.item collection
func (m *EducationSubmissionItemRequestBuilder) SubmittedResourcesById(id string)(*ie785db4b38f9f4ba6fc3e6cfb4164fd07a4b957a6f8e1b2b297b1e8dfeb71a05.EducationSubmissionResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmissionResource%2Did"] = id
    }
    return ie785db4b38f9f4ba6fc3e6cfb4164fd07a4b957a6f8e1b2b297b1e8dfeb71a05.NewEducationSubmissionResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unsubmit the unsubmit property
func (m *EducationSubmissionItemRequestBuilder) Unsubmit()(*ifbd6f6420c2c3a622cde0cdd389b1186926a6ad859deff11e8e961ce84edf272.UnsubmitRequestBuilder) {
    return ifbd6f6420c2c3a622cde0cdd389b1186926a6ad859deff11e8e961ce84edf272.NewUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
