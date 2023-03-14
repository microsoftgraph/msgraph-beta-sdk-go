package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder provides operations to manage the subjectRightsRequests property of the microsoft.graph.security entity.
type SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetQueryParameters get subjectRightsRequests from security
type SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetQueryParameters
}
// SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Approvers provides operations to manage the approvers property of the microsoft.graph.subjectRightsRequest entity.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) Approvers()(*SubjectRightsRequestsItemApproversRequestBuilder) {
    return NewSubjectRightsRequestsItemApproversRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ApproversById provides operations to manage the approvers property of the microsoft.graph.subjectRightsRequest entity.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) ApproversById(id string)(*SubjectRightsRequestsItemApproversUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return NewSubjectRightsRequestsItemApproversUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Collaborators provides operations to manage the collaborators property of the microsoft.graph.subjectRightsRequest entity.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) Collaborators()(*SubjectRightsRequestsItemCollaboratorsRequestBuilder) {
    return NewSubjectRightsRequestsItemCollaboratorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CollaboratorsById provides operations to manage the collaborators property of the microsoft.graph.subjectRightsRequest entity.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) CollaboratorsById(id string)(*SubjectRightsRequestsItemCollaboratorsUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return NewSubjectRightsRequestsItemCollaboratorsUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewSubjectRightsRequestsSubjectRightsRequestItemRequestBuilderInternal instantiates a new SubjectRightsRequestItemRequestBuilder and sets the default values.
func NewSubjectRightsRequestsSubjectRightsRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) {
    m := &SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/subjectRightsRequests/{subjectRightsRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSubjectRightsRequestsSubjectRightsRequestItemRequestBuilder instantiates a new SubjectRightsRequestItemRequestBuilder and sets the default values.
func NewSubjectRightsRequestsSubjectRightsRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectRightsRequestsSubjectRightsRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property subjectRightsRequests for security
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get subjectRightsRequests from security
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectRightsRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSubjectRightsRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectRightsRequestable), nil
}
// GetFinalAttachment provides operations to call the getFinalAttachment method.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) GetFinalAttachment()(*SubjectRightsRequestsItemGetFinalAttachmentRequestBuilder) {
    return NewSubjectRightsRequestsItemGetFinalAttachmentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GetFinalReport provides operations to call the getFinalReport method.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) GetFinalReport()(*SubjectRightsRequestsItemGetFinalReportRequestBuilder) {
    return NewSubjectRightsRequestsItemGetFinalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Notes provides operations to manage the notes property of the microsoft.graph.subjectRightsRequest entity.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) Notes()(*SubjectRightsRequestsItemNotesRequestBuilder) {
    return NewSubjectRightsRequestsItemNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NotesById provides operations to manage the notes property of the microsoft.graph.subjectRightsRequest entity.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) NotesById(id string)(*SubjectRightsRequestsItemNotesAuthoredNoteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authoredNote%2Did"] = id
    }
    return NewSubjectRightsRequestsItemNotesAuthoredNoteItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property subjectRightsRequests in security
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectRightsRequestable, requestConfiguration *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectRightsRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSubjectRightsRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectRightsRequestable), nil
}
// Team provides operations to manage the team property of the microsoft.graph.subjectRightsRequest entity.
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) Team()(*SubjectRightsRequestsItemTeamRequestBuilder) {
    return NewSubjectRightsRequestsItemTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property subjectRightsRequests for security
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get subjectRightsRequests from security
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property subjectRightsRequests in security
func (m *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectRightsRequestable, requestConfiguration *SubjectRightsRequestsSubjectRightsRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
