package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i08e66c5d7bd1abccfebd8fb93bd3824b29263d4a78a8e0619b6492ed8f3fad49 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/playprompt"
    i0c63fde205100ca16ab7a423f47cbdbde7db96e5bfd78a3f404e33371a9a304c "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/record"
    i18782668e274df0b14fede1b2d0b04b9dfad365a601b91cb1ae1236333912f77 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/recordresponse"
    i2ad07b1a9d626a23f6fb4c6980016bdfa91427506827a8183fe0ab85ccea7297 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/mute"
    i2e0863e08fc93d954484ce570f9308574c3c610b9cbf1e2c9d07876c640463ac "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/updaterecordingstatus"
    i3af510b639a3bfd33b6d77bb52f8d92c7f9e25367174bdfa6142903edc74995c "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/keepalive"
    i4cae3b1a69ba2254d74646c3ba02c22835c51e378675f8f86fa207a4bf13fbb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/unmute"
    i4ffb8669c93d0b11c0238e80369824011ea5e33bda9395560f9abe880eb8acc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/transfer"
    i51c2a09e4f41fd4c07a21a333066d7d9b5683f48919e40d555e43eeedc2d2b0f "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/redirect"
    i59f1af4a71937287f9a647acee47040de638eb6d16981d2dc3f1c22d63d3d210 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/subscribetotone"
    i6a956e9da3a6e7a50695a94bcd7dd16f04667e628762662e2a562dbb513808d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/contentsharingsessions"
    i6a96c1601820e25e2c11e319760d663fd67b6cdb677b5d599e7f4560698143e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/changescreensharingrole"
    i7b15e8c80b7aa50a0151a753f97ef7289d17791928250d62926b97f749d6c305 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/answer"
    i8659ea4008fd7459c7570b3b977a99568235e5600beaac36764abf7a09a1c6a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/operations"
    ib840d59e075c9f1a42737eb611c25c9b9760e7d651bf5aba643e84343070c957 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/addlargegalleryview"
    ic54ac4adcd492f82a003801c0ce35301e9f04c5676f80853066311daf9c569da "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/audioroutinggroups"
    id1e985c21267b30b9212f7e3d69af1960754fba433a529ab9f2819e92d7804e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/cancelmediaprocessing"
    idaf0ca56e2f68098afebd431de7c548ac38196c36423e844dd58942eb2a4ee85 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/reject"
    if78246d1256eb5bc7db2b2c1dcd3d9daaa6de508f121f6f0bb97fcc022fe09f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/participants"
    i151a479d1626e685450cad3f10b2ae807a5841c52441c261bfec932663c91e9f "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/audioroutinggroups/item"
    ie37e4f36986bce11c7045e81e2c2f02722145373c94ed229df4e1f757aa12a44 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/contentsharingsessions/item"
    ie3ddf2285119a257859e618e7d13130592dc1f6e886e271f46d05000453bffd2 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/participants/item"
    ifc54b1625a5ef155b182402b189053da9800796fb7507ed629f84919426791d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/calls/item/operations/item"
)

// CallItemRequestBuilder provides operations to manage the calls property of the microsoft.graph.cloudCommunications entity.
type CallItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CallItemRequestBuilderGetQueryParameters get calls from communications
type CallItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CallItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallItemRequestBuilderGetQueryParameters
}
// CallItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddLargeGalleryView the addLargeGalleryView property
func (m *CallItemRequestBuilder) AddLargeGalleryView()(*ib840d59e075c9f1a42737eb611c25c9b9760e7d651bf5aba643e84343070c957.AddLargeGalleryViewRequestBuilder) {
    return ib840d59e075c9f1a42737eb611c25c9b9760e7d651bf5aba643e84343070c957.NewAddLargeGalleryViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Answer the answer property
func (m *CallItemRequestBuilder) Answer()(*i7b15e8c80b7aa50a0151a753f97ef7289d17791928250d62926b97f749d6c305.AnswerRequestBuilder) {
    return i7b15e8c80b7aa50a0151a753f97ef7289d17791928250d62926b97f749d6c305.NewAnswerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroups the audioRoutingGroups property
func (m *CallItemRequestBuilder) AudioRoutingGroups()(*ic54ac4adcd492f82a003801c0ce35301e9f04c5676f80853066311daf9c569da.AudioRoutingGroupsRequestBuilder) {
    return ic54ac4adcd492f82a003801c0ce35301e9f04c5676f80853066311daf9c569da.NewAudioRoutingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.calls.item.audioRoutingGroups.item collection
func (m *CallItemRequestBuilder) AudioRoutingGroupsById(id string)(*i151a479d1626e685450cad3f10b2ae807a5841c52441c261bfec932663c91e9f.AudioRoutingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["audioRoutingGroup%2Did"] = id
    }
    return i151a479d1626e685450cad3f10b2ae807a5841c52441c261bfec932663c91e9f.NewAudioRoutingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CancelMediaProcessing the cancelMediaProcessing property
func (m *CallItemRequestBuilder) CancelMediaProcessing()(*id1e985c21267b30b9212f7e3d69af1960754fba433a529ab9f2819e92d7804e6.CancelMediaProcessingRequestBuilder) {
    return id1e985c21267b30b9212f7e3d69af1960754fba433a529ab9f2819e92d7804e6.NewCancelMediaProcessingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangeScreenSharingRole the changeScreenSharingRole property
func (m *CallItemRequestBuilder) ChangeScreenSharingRole()(*i6a96c1601820e25e2c11e319760d663fd67b6cdb677b5d599e7f4560698143e1.ChangeScreenSharingRoleRequestBuilder) {
    return i6a96c1601820e25e2c11e319760d663fd67b6cdb677b5d599e7f4560698143e1.NewChangeScreenSharingRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCallItemRequestBuilderInternal instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallItemRequestBuilder) {
    m := &CallItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCallItemRequestBuilder instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentSharingSessions the contentSharingSessions property
func (m *CallItemRequestBuilder) ContentSharingSessions()(*i6a956e9da3a6e7a50695a94bcd7dd16f04667e628762662e2a562dbb513808d4.ContentSharingSessionsRequestBuilder) {
    return i6a956e9da3a6e7a50695a94bcd7dd16f04667e628762662e2a562dbb513808d4.NewContentSharingSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentSharingSessionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.calls.item.contentSharingSessions.item collection
func (m *CallItemRequestBuilder) ContentSharingSessionsById(id string)(*ie37e4f36986bce11c7045e81e2c2f02722145373c94ed229df4e1f757aa12a44.ContentSharingSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentSharingSession%2Did"] = id
    }
    return ie37e4f36986bce11c7045e81e2c2f02722145373c94ed229df4e1f757aa12a44.NewContentSharingSessionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property calls for communications
func (m *CallItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property calls for communications
func (m *CallItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *CallItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get calls from communications
func (m *CallItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get calls from communications
func (m *CallItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *CallItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property calls in communications
func (m *CallItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property calls in communications
func (m *CallItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, requestConfiguration *CallItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property calls for communications
func (m *CallItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property calls for communications
func (m *CallItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *CallItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Get get calls from communications
func (m *CallItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get calls from communications
func (m *CallItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *CallItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCallFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable), nil
}
// KeepAlive the keepAlive property
func (m *CallItemRequestBuilder) KeepAlive()(*i3af510b639a3bfd33b6d77bb52f8d92c7f9e25367174bdfa6142903edc74995c.KeepAliveRequestBuilder) {
    return i3af510b639a3bfd33b6d77bb52f8d92c7f9e25367174bdfa6142903edc74995c.NewKeepAliveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Mute the mute property
func (m *CallItemRequestBuilder) Mute()(*i2ad07b1a9d626a23f6fb4c6980016bdfa91427506827a8183fe0ab85ccea7297.MuteRequestBuilder) {
    return i2ad07b1a9d626a23f6fb4c6980016bdfa91427506827a8183fe0ab85ccea7297.NewMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations the operations property
func (m *CallItemRequestBuilder) Operations()(*i8659ea4008fd7459c7570b3b977a99568235e5600beaac36764abf7a09a1c6a2.OperationsRequestBuilder) {
    return i8659ea4008fd7459c7570b3b977a99568235e5600beaac36764abf7a09a1c6a2.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.calls.item.operations.item collection
func (m *CallItemRequestBuilder) OperationsById(id string)(*ifc54b1625a5ef155b182402b189053da9800796fb7507ed629f84919426791d1.CommsOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["commsOperation%2Did"] = id
    }
    return ifc54b1625a5ef155b182402b189053da9800796fb7507ed629f84919426791d1.NewCommsOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Participants the participants property
func (m *CallItemRequestBuilder) Participants()(*if78246d1256eb5bc7db2b2c1dcd3d9daaa6de508f121f6f0bb97fcc022fe09f5.ParticipantsRequestBuilder) {
    return if78246d1256eb5bc7db2b2c1dcd3d9daaa6de508f121f6f0bb97fcc022fe09f5.NewParticipantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParticipantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.communications.calls.item.participants.item collection
func (m *CallItemRequestBuilder) ParticipantsById(id string)(*ie3ddf2285119a257859e618e7d13130592dc1f6e886e271f46d05000453bffd2.ParticipantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["participant%2Did"] = id
    }
    return ie3ddf2285119a257859e618e7d13130592dc1f6e886e271f46d05000453bffd2.NewParticipantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calls in communications
func (m *CallItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property calls in communications
func (m *CallItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, requestConfiguration *CallItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// PlayPrompt the playPrompt property
func (m *CallItemRequestBuilder) PlayPrompt()(*i08e66c5d7bd1abccfebd8fb93bd3824b29263d4a78a8e0619b6492ed8f3fad49.PlayPromptRequestBuilder) {
    return i08e66c5d7bd1abccfebd8fb93bd3824b29263d4a78a8e0619b6492ed8f3fad49.NewPlayPromptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Record the record property
func (m *CallItemRequestBuilder) Record()(*i0c63fde205100ca16ab7a423f47cbdbde7db96e5bfd78a3f404e33371a9a304c.RecordRequestBuilder) {
    return i0c63fde205100ca16ab7a423f47cbdbde7db96e5bfd78a3f404e33371a9a304c.NewRecordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecordResponse the recordResponse property
func (m *CallItemRequestBuilder) RecordResponse()(*i18782668e274df0b14fede1b2d0b04b9dfad365a601b91cb1ae1236333912f77.RecordResponseRequestBuilder) {
    return i18782668e274df0b14fede1b2d0b04b9dfad365a601b91cb1ae1236333912f77.NewRecordResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Redirect the redirect property
func (m *CallItemRequestBuilder) Redirect()(*i51c2a09e4f41fd4c07a21a333066d7d9b5683f48919e40d555e43eeedc2d2b0f.RedirectRequestBuilder) {
    return i51c2a09e4f41fd4c07a21a333066d7d9b5683f48919e40d555e43eeedc2d2b0f.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reject the reject property
func (m *CallItemRequestBuilder) Reject()(*idaf0ca56e2f68098afebd431de7c548ac38196c36423e844dd58942eb2a4ee85.RejectRequestBuilder) {
    return idaf0ca56e2f68098afebd431de7c548ac38196c36423e844dd58942eb2a4ee85.NewRejectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeToTone the subscribeToTone property
func (m *CallItemRequestBuilder) SubscribeToTone()(*i59f1af4a71937287f9a647acee47040de638eb6d16981d2dc3f1c22d63d3d210.SubscribeToToneRequestBuilder) {
    return i59f1af4a71937287f9a647acee47040de638eb6d16981d2dc3f1c22d63d3d210.NewSubscribeToToneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Transfer the transfer property
func (m *CallItemRequestBuilder) Transfer()(*i4ffb8669c93d0b11c0238e80369824011ea5e33bda9395560f9abe880eb8acc0.TransferRequestBuilder) {
    return i4ffb8669c93d0b11c0238e80369824011ea5e33bda9395560f9abe880eb8acc0.NewTransferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unmute the unmute property
func (m *CallItemRequestBuilder) Unmute()(*i4cae3b1a69ba2254d74646c3ba02c22835c51e378675f8f86fa207a4bf13fbb9.UnmuteRequestBuilder) {
    return i4cae3b1a69ba2254d74646c3ba02c22835c51e378675f8f86fa207a4bf13fbb9.NewUnmuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateRecordingStatus the updateRecordingStatus property
func (m *CallItemRequestBuilder) UpdateRecordingStatus()(*i2e0863e08fc93d954484ce570f9308574c3c610b9cbf1e2c9d07876c640463ac.UpdateRecordingStatusRequestBuilder) {
    return i2e0863e08fc93d954484ce570f9308574c3c610b9cbf1e2c9d07876c640463ac.NewUpdateRecordingStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
