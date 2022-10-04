package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i032e80f26f63621eac368a0e78e5d567ee153d1bcbbb457d6b581d74b6b0dc6a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/contentsharingsessions"
    i0918551da47a228d24ebdfaf6df77c471471893c9c40b23a04e4c7649b85b45d "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/reject"
    i1b9466d0caef82626f32a61ccf84f8b293d289edb108038c604aef2d18197865 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/audioroutinggroups"
    i1c38d0d034bef05f35da04b0f4b5dcfbc13afbe00ddb552eb78087b9de98ecee "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/unmute"
    i1cd79905fc528b166f37b3757de3739a99d140dd293c565a482a0ccfcc9151d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/participants"
    i1dcc77e9871abeb934d55bc57c72c4ba35a7f03ea906a51875d763fa8fa5249c "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/cancelmediaprocessing"
    i22def5bc84ea78118759eb49e711824623c72f98f1ad008c2795e4e7b8843287 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/playprompt"
    i2b1770e190ed76efc5e53b7e27770c0cec2b19f628732fc7141dd7c047d710ce "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/updaterecordingstatus"
    i540daaa499f2a1b3aa80d06d790d11b8789cb164f83994ac8ef2666aa88225bf "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/recordresponse"
    i6f1572cf2923d7161e932806b4455848280c38ad3685ab81f1a3301526b5e3c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/operations"
    i739635c91b9aae9a71009f21d3aabbd6c3cdbd84bf891707ccb7ad34ab9056a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/record"
    i748993866d3c287419aae52c6b71e09769323e7c5f317345f259e426bc3d9af2 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/subscribetotone"
    i874a580791117a3b7d274c118410fe2d9225cceaca52e78051f9cb4b1a453662 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/redirect"
    i8ba5b1d1710189d6fbeadcc3c0d5040fd10a40b9f9cb675742fafb85918a74e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/mute"
    ia617450ea6a6c55b69e3406d78900114c4a839be32a64fbf47520d45899eab0b "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/addlargegalleryview"
    iaf15ceb43d09ee52ea76cf06a8a2800ffa9e17af30581ca61badb5604676ace8 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/keepalive"
    icbe1c65482be5ba77560b0854bb72eb2fce823cdda74fb44a06b72364da1f6a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/changescreensharingrole"
    iea07e291a46b990a979a038612645323fe7180f50eb887d4daf5d74572112a37 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/answer"
    if587e80b75e5d8ffe7d9fddabdc3d886c5ce62ad5059d5c8a6464a4bcfa59b5e "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/transfer"
    i0e77db2894bae546ce3028844c9be92c1f453524b9969cb6a1c5fbc30a81428b "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/contentsharingsessions/item"
    i39deddda18f57c5a3a9987261e9e10fe2b86924c965453300c89efb467115e2a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/operations/item"
    i76c60465809f227922ef886abb4be5e5ba936f1aee25121d8be66ca149691e37 "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/audioroutinggroups/item"
    idfae2eb2f859c83384878a7565f78eb2c09250cf7b3892104836a28c6368c37a "github.com/microsoftgraph/msgraph-beta-sdk-go/app/calls/item/participants/item"
)

// CallItemRequestBuilder provides operations to manage the calls property of the microsoft.graph.commsApplication entity.
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
// CallItemRequestBuilderGetQueryParameters get calls from app
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
func (m *CallItemRequestBuilder) AddLargeGalleryView()(*ia617450ea6a6c55b69e3406d78900114c4a839be32a64fbf47520d45899eab0b.AddLargeGalleryViewRequestBuilder) {
    return ia617450ea6a6c55b69e3406d78900114c4a839be32a64fbf47520d45899eab0b.NewAddLargeGalleryViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Answer the answer property
func (m *CallItemRequestBuilder) Answer()(*iea07e291a46b990a979a038612645323fe7180f50eb887d4daf5d74572112a37.AnswerRequestBuilder) {
    return iea07e291a46b990a979a038612645323fe7180f50eb887d4daf5d74572112a37.NewAnswerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroups the audioRoutingGroups property
func (m *CallItemRequestBuilder) AudioRoutingGroups()(*i1b9466d0caef82626f32a61ccf84f8b293d289edb108038c604aef2d18197865.AudioRoutingGroupsRequestBuilder) {
    return i1b9466d0caef82626f32a61ccf84f8b293d289edb108038c604aef2d18197865.NewAudioRoutingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.calls.item.audioRoutingGroups.item collection
func (m *CallItemRequestBuilder) AudioRoutingGroupsById(id string)(*i76c60465809f227922ef886abb4be5e5ba936f1aee25121d8be66ca149691e37.AudioRoutingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["audioRoutingGroup%2Did"] = id
    }
    return i76c60465809f227922ef886abb4be5e5ba936f1aee25121d8be66ca149691e37.NewAudioRoutingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CancelMediaProcessing the cancelMediaProcessing property
func (m *CallItemRequestBuilder) CancelMediaProcessing()(*i1dcc77e9871abeb934d55bc57c72c4ba35a7f03ea906a51875d763fa8fa5249c.CancelMediaProcessingRequestBuilder) {
    return i1dcc77e9871abeb934d55bc57c72c4ba35a7f03ea906a51875d763fa8fa5249c.NewCancelMediaProcessingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangeScreenSharingRole the changeScreenSharingRole property
func (m *CallItemRequestBuilder) ChangeScreenSharingRole()(*icbe1c65482be5ba77560b0854bb72eb2fce823cdda74fb44a06b72364da1f6a2.ChangeScreenSharingRoleRequestBuilder) {
    return icbe1c65482be5ba77560b0854bb72eb2fce823cdda74fb44a06b72364da1f6a2.NewChangeScreenSharingRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCallItemRequestBuilderInternal instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallItemRequestBuilder) {
    m := &CallItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app/calls/{call%2Did}{?%24select,%24expand}";
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
func (m *CallItemRequestBuilder) ContentSharingSessions()(*i032e80f26f63621eac368a0e78e5d567ee153d1bcbbb457d6b581d74b6b0dc6a.ContentSharingSessionsRequestBuilder) {
    return i032e80f26f63621eac368a0e78e5d567ee153d1bcbbb457d6b581d74b6b0dc6a.NewContentSharingSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentSharingSessionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.calls.item.contentSharingSessions.item collection
func (m *CallItemRequestBuilder) ContentSharingSessionsById(id string)(*i0e77db2894bae546ce3028844c9be92c1f453524b9969cb6a1c5fbc30a81428b.ContentSharingSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentSharingSession%2Did"] = id
    }
    return i0e77db2894bae546ce3028844c9be92c1f453524b9969cb6a1c5fbc30a81428b.NewContentSharingSessionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property calls for app
func (m *CallItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *CallItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get calls from app
func (m *CallItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *CallItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property calls in app
func (m *CallItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, requestConfiguration *CallItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property calls for app
func (m *CallItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CallItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get get calls from app
func (m *CallItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CallItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCallFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable), nil
}
// KeepAlive the keepAlive property
func (m *CallItemRequestBuilder) KeepAlive()(*iaf15ceb43d09ee52ea76cf06a8a2800ffa9e17af30581ca61badb5604676ace8.KeepAliveRequestBuilder) {
    return iaf15ceb43d09ee52ea76cf06a8a2800ffa9e17af30581ca61badb5604676ace8.NewKeepAliveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Mute the mute property
func (m *CallItemRequestBuilder) Mute()(*i8ba5b1d1710189d6fbeadcc3c0d5040fd10a40b9f9cb675742fafb85918a74e0.MuteRequestBuilder) {
    return i8ba5b1d1710189d6fbeadcc3c0d5040fd10a40b9f9cb675742fafb85918a74e0.NewMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations the operations property
func (m *CallItemRequestBuilder) Operations()(*i6f1572cf2923d7161e932806b4455848280c38ad3685ab81f1a3301526b5e3c1.OperationsRequestBuilder) {
    return i6f1572cf2923d7161e932806b4455848280c38ad3685ab81f1a3301526b5e3c1.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.calls.item.operations.item collection
func (m *CallItemRequestBuilder) OperationsById(id string)(*i39deddda18f57c5a3a9987261e9e10fe2b86924c965453300c89efb467115e2a.CommsOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["commsOperation%2Did"] = id
    }
    return i39deddda18f57c5a3a9987261e9e10fe2b86924c965453300c89efb467115e2a.NewCommsOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Participants the participants property
func (m *CallItemRequestBuilder) Participants()(*i1cd79905fc528b166f37b3757de3739a99d140dd293c565a482a0ccfcc9151d6.ParticipantsRequestBuilder) {
    return i1cd79905fc528b166f37b3757de3739a99d140dd293c565a482a0ccfcc9151d6.NewParticipantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParticipantsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.app.calls.item.participants.item collection
func (m *CallItemRequestBuilder) ParticipantsById(id string)(*idfae2eb2f859c83384878a7565f78eb2c09250cf7b3892104836a28c6368c37a.ParticipantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["participant%2Did"] = id
    }
    return idfae2eb2f859c83384878a7565f78eb2c09250cf7b3892104836a28c6368c37a.NewParticipantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calls in app
func (m *CallItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, requestConfiguration *CallItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCallFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable), nil
}
// PlayPrompt the playPrompt property
func (m *CallItemRequestBuilder) PlayPrompt()(*i22def5bc84ea78118759eb49e711824623c72f98f1ad008c2795e4e7b8843287.PlayPromptRequestBuilder) {
    return i22def5bc84ea78118759eb49e711824623c72f98f1ad008c2795e4e7b8843287.NewPlayPromptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Record the record property
func (m *CallItemRequestBuilder) Record()(*i739635c91b9aae9a71009f21d3aabbd6c3cdbd84bf891707ccb7ad34ab9056a6.RecordRequestBuilder) {
    return i739635c91b9aae9a71009f21d3aabbd6c3cdbd84bf891707ccb7ad34ab9056a6.NewRecordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecordResponse the recordResponse property
func (m *CallItemRequestBuilder) RecordResponse()(*i540daaa499f2a1b3aa80d06d790d11b8789cb164f83994ac8ef2666aa88225bf.RecordResponseRequestBuilder) {
    return i540daaa499f2a1b3aa80d06d790d11b8789cb164f83994ac8ef2666aa88225bf.NewRecordResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Redirect the redirect property
func (m *CallItemRequestBuilder) Redirect()(*i874a580791117a3b7d274c118410fe2d9225cceaca52e78051f9cb4b1a453662.RedirectRequestBuilder) {
    return i874a580791117a3b7d274c118410fe2d9225cceaca52e78051f9cb4b1a453662.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reject the reject property
func (m *CallItemRequestBuilder) Reject()(*i0918551da47a228d24ebdfaf6df77c471471893c9c40b23a04e4c7649b85b45d.RejectRequestBuilder) {
    return i0918551da47a228d24ebdfaf6df77c471471893c9c40b23a04e4c7649b85b45d.NewRejectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeToTone the subscribeToTone property
func (m *CallItemRequestBuilder) SubscribeToTone()(*i748993866d3c287419aae52c6b71e09769323e7c5f317345f259e426bc3d9af2.SubscribeToToneRequestBuilder) {
    return i748993866d3c287419aae52c6b71e09769323e7c5f317345f259e426bc3d9af2.NewSubscribeToToneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Transfer the transfer property
func (m *CallItemRequestBuilder) Transfer()(*if587e80b75e5d8ffe7d9fddabdc3d886c5ce62ad5059d5c8a6464a4bcfa59b5e.TransferRequestBuilder) {
    return if587e80b75e5d8ffe7d9fddabdc3d886c5ce62ad5059d5c8a6464a4bcfa59b5e.NewTransferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unmute the unmute property
func (m *CallItemRequestBuilder) Unmute()(*i1c38d0d034bef05f35da04b0f4b5dcfbc13afbe00ddb552eb78087b9de98ecee.UnmuteRequestBuilder) {
    return i1c38d0d034bef05f35da04b0f4b5dcfbc13afbe00ddb552eb78087b9de98ecee.NewUnmuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateRecordingStatus the updateRecordingStatus property
func (m *CallItemRequestBuilder) UpdateRecordingStatus()(*i2b1770e190ed76efc5e53b7e27770c0cec2b19f628732fc7141dd7c047d710ce.UpdateRecordingStatusRequestBuilder) {
    return i2b1770e190ed76efc5e53b7e27770c0cec2b19f628732fc7141dd7c047d710ce.NewUpdateRecordingStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
