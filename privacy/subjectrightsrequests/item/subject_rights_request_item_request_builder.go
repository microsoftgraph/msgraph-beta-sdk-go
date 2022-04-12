package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i06e33bb086cbb5f9333cc1c8bb09529f00623800d8f148e19c13fab399be0e21 "github.com/microsoftgraph/msgraph-beta-sdk-go/privacy/subjectrightsrequests/item/notes"
    i2bf59a4d763ab3355a4526b9ff2ddcae384f1210dd464ff5940ff3a054d5f26e "github.com/microsoftgraph/msgraph-beta-sdk-go/privacy/subjectrightsrequests/item/getfinalattachment"
    i958489264196dd08f36dd719f907da913cf1a6b619228f3d6bc8da52997d8449 "github.com/microsoftgraph/msgraph-beta-sdk-go/privacy/subjectrightsrequests/item/getfinalreport"
    ia790ce0fe221c2e579ab27030d34081e5415f38f47af2263081eb369b2561ccd "github.com/microsoftgraph/msgraph-beta-sdk-go/privacy/subjectrightsrequests/item/team"
    ibe7a5a69eee3e39fd9fd52c318c683abd357fbe438629376f053fbf8dafee6d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/privacy/subjectrightsrequests/item/notes/item"
)

// SubjectRightsRequestItemRequestBuilder provides operations to manage the subjectRightsRequests property of the microsoft.graph.privacy entity.
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
// SubjectRightsRequestItemRequestBuilderGetQueryParameters get subjectRightsRequests from privacy
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
    m.urlTemplate = "{+baseurl}/privacy/subjectRightsRequests/{subjectRightsRequest%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property subjectRightsRequests for privacy
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
// CreateGetRequestInformation get subjectRightsRequests from privacy
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
// CreatePatchRequestInformation update the navigation property subjectRightsRequests in privacy
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
// Delete delete navigation property subjectRightsRequests for privacy
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
// Get get subjectRightsRequests from privacy
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
func (m *SubjectRightsRequestItemRequestBuilder) GetFinalAttachment()(*i2bf59a4d763ab3355a4526b9ff2ddcae384f1210dd464ff5940ff3a054d5f26e.GetFinalAttachmentRequestBuilder) {
    return i2bf59a4d763ab3355a4526b9ff2ddcae384f1210dd464ff5940ff3a054d5f26e.NewGetFinalAttachmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetFinalReport provides operations to call the getFinalReport method.
func (m *SubjectRightsRequestItemRequestBuilder) GetFinalReport()(*i958489264196dd08f36dd719f907da913cf1a6b619228f3d6bc8da52997d8449.GetFinalReportRequestBuilder) {
    return i958489264196dd08f36dd719f907da913cf1a6b619228f3d6bc8da52997d8449.NewGetFinalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Notes the notes property
func (m *SubjectRightsRequestItemRequestBuilder) Notes()(*i06e33bb086cbb5f9333cc1c8bb09529f00623800d8f148e19c13fab399be0e21.NotesRequestBuilder) {
    return i06e33bb086cbb5f9333cc1c8bb09529f00623800d8f148e19c13fab399be0e21.NewNotesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privacy.subjectRightsRequests.item.notes.item collection
func (m *SubjectRightsRequestItemRequestBuilder) NotesById(id string)(*ibe7a5a69eee3e39fd9fd52c318c683abd357fbe438629376f053fbf8dafee6d9.AuthoredNoteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authoredNote%2Did"] = id
    }
    return ibe7a5a69eee3e39fd9fd52c318c683abd357fbe438629376f053fbf8dafee6d9.NewAuthoredNoteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property subjectRightsRequests in privacy
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
func (m *SubjectRightsRequestItemRequestBuilder) Team()(*ia790ce0fe221c2e579ab27030d34081e5415f38f47af2263081eb369b2561ccd.TeamRequestBuilder) {
    return ia790ce0fe221c2e579ab27030d34081e5415f38f47af2263081eb369b2561ccd.NewTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
