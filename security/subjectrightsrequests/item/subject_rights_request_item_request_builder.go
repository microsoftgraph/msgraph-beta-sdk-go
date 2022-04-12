package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3c6b3962f4f957ab38ad7c6292548c19324c840c5f4c6caf1d78354ceac67949 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/subjectrightsrequests/item/getfinalattachment"
    i85ceeb3fb58c5b0a45a32b65876d69fa3767299ea1e59ab70bb19f94fd16f3ba "github.com/microsoftgraph/msgraph-beta-sdk-go/security/subjectrightsrequests/item/notes"
    ia2013bbfcd0f09a5b744acfe4f28cd0f4a71f920f7dc684799f554684dec56af "github.com/microsoftgraph/msgraph-beta-sdk-go/security/subjectrightsrequests/item/getfinalreport"
    if314322e27ebba41b48454133ec3ca22da3b4aad16b2d97239c959afa4e445be "github.com/microsoftgraph/msgraph-beta-sdk-go/security/subjectrightsrequests/item/team"
    i8a7d93ed28ef349b2b993a1fc0c5f4e014bde5b2d9f556a4c3a79c6c310a82c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/subjectrightsrequests/item/notes/item"
)

// SubjectRightsRequestItemRequestBuilder provides operations to manage the subjectRightsRequests property of the microsoft.graph.security entity.
type SubjectRightsRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SubjectRightsRequestItemRequestBuilderDeleteOptions options for Delete
type SubjectRightsRequestItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SubjectRightsRequestItemRequestBuilderGetOptions options for Get
type SubjectRightsRequestItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SubjectRightsRequestItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SubjectRightsRequestItemRequestBuilderGetQueryParameters get subjectRightsRequests from security
type SubjectRightsRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SubjectRightsRequestItemRequestBuilderPatchOptions options for Patch
type SubjectRightsRequestItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectRightsRequestable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewSubjectRightsRequestItemRequestBuilderInternal instantiates a new SubjectRightsRequestItemRequestBuilder and sets the default values.
func NewSubjectRightsRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestItemRequestBuilder) {
    m := &SubjectRightsRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/subjectRightsRequests/{subjectRightsRequest%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSubjectRightsRequestItemRequestBuilder instantiates a new SubjectRightsRequestItemRequestBuilder and sets the default values.
func NewSubjectRightsRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubjectRightsRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubjectRightsRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property subjectRightsRequests for security
func (m *SubjectRightsRequestItemRequestBuilder) CreateDeleteRequestInformation(options *SubjectRightsRequestItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get subjectRightsRequests from security
func (m *SubjectRightsRequestItemRequestBuilder) CreateGetRequestInformation(options *SubjectRightsRequestItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property subjectRightsRequests in security
func (m *SubjectRightsRequestItemRequestBuilder) CreatePatchRequestInformation(options *SubjectRightsRequestItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property subjectRightsRequests for security
func (m *SubjectRightsRequestItemRequestBuilder) Delete(options *SubjectRightsRequestItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get subjectRightsRequests from security
func (m *SubjectRightsRequestItemRequestBuilder) Get(options *SubjectRightsRequestItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectRightsRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSubjectRightsRequestFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SubjectRightsRequestable), nil
}
// GetFinalAttachment provides operations to call the getFinalAttachment method.
func (m *SubjectRightsRequestItemRequestBuilder) GetFinalAttachment()(*i3c6b3962f4f957ab38ad7c6292548c19324c840c5f4c6caf1d78354ceac67949.GetFinalAttachmentRequestBuilder) {
    return i3c6b3962f4f957ab38ad7c6292548c19324c840c5f4c6caf1d78354ceac67949.NewGetFinalAttachmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFinalReport provides operations to call the getFinalReport method.
func (m *SubjectRightsRequestItemRequestBuilder) GetFinalReport()(*ia2013bbfcd0f09a5b744acfe4f28cd0f4a71f920f7dc684799f554684dec56af.GetFinalReportRequestBuilder) {
    return ia2013bbfcd0f09a5b744acfe4f28cd0f4a71f920f7dc684799f554684dec56af.NewGetFinalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Notes the notes property
func (m *SubjectRightsRequestItemRequestBuilder) Notes()(*i85ceeb3fb58c5b0a45a32b65876d69fa3767299ea1e59ab70bb19f94fd16f3ba.NotesRequestBuilder) {
    return i85ceeb3fb58c5b0a45a32b65876d69fa3767299ea1e59ab70bb19f94fd16f3ba.NewNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.security.subjectRightsRequests.item.notes.item collection
func (m *SubjectRightsRequestItemRequestBuilder) NotesById(id string)(*i8a7d93ed28ef349b2b993a1fc0c5f4e014bde5b2d9f556a4c3a79c6c310a82c0.AuthoredNoteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authoredNote%2Did"] = id
    }
    return i8a7d93ed28ef349b2b993a1fc0c5f4e014bde5b2d9f556a4c3a79c6c310a82c0.NewAuthoredNoteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property subjectRightsRequests in security
func (m *SubjectRightsRequestItemRequestBuilder) Patch(options *SubjectRightsRequestItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Team the team property
func (m *SubjectRightsRequestItemRequestBuilder) Team()(*if314322e27ebba41b48454133ec3ca22da3b4aad16b2d97239c959afa4e445be.TeamRequestBuilder) {
    return if314322e27ebba41b48454133ec3ca22da3b4aad16b2d97239c959afa4e445be.NewTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
