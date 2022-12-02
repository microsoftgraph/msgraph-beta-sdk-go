package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AppCallsCallItemRequestBuilder provides operations to manage the calls property of the microsoft.graph.commsApplication entity.
type AppCallsCallItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppCallsCallItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppCallsCallItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppCallsCallItemRequestBuilderGetQueryParameters get calls from app
type AppCallsCallItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppCallsCallItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppCallsCallItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppCallsCallItemRequestBuilderGetQueryParameters
}
// AppCallsCallItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppCallsCallItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddLargeGalleryView provides operations to call the addLargeGalleryView method.
func (m *AppCallsCallItemRequestBuilder) AddLargeGalleryView()(*AppCallsItemAddLargeGalleryViewRequestBuilder) {
    return NewAppCallsItemAddLargeGalleryViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Answer provides operations to call the answer method.
func (m *AppCallsCallItemRequestBuilder) Answer()(*AppCallsItemAnswerRequestBuilder) {
    return NewAppCallsItemAnswerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroups provides operations to manage the audioRoutingGroups property of the microsoft.graph.call entity.
func (m *AppCallsCallItemRequestBuilder) AudioRoutingGroups()(*AppCallsItemAudioRoutingGroupsRequestBuilder) {
    return NewAppCallsItemAudioRoutingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroupsById provides operations to manage the audioRoutingGroups property of the microsoft.graph.call entity.
func (m *AppCallsCallItemRequestBuilder) AudioRoutingGroupsById(id string)(*AppCallsItemAudioRoutingGroupsAudioRoutingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["audioRoutingGroup%2Did"] = id
    }
    return NewAppCallsItemAudioRoutingGroupsAudioRoutingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CancelMediaProcessing provides operations to call the cancelMediaProcessing method.
func (m *AppCallsCallItemRequestBuilder) CancelMediaProcessing()(*AppCallsItemCancelMediaProcessingRequestBuilder) {
    return NewAppCallsItemCancelMediaProcessingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangeScreenSharingRole provides operations to call the changeScreenSharingRole method.
func (m *AppCallsCallItemRequestBuilder) ChangeScreenSharingRole()(*AppCallsItemChangeScreenSharingRoleRequestBuilder) {
    return NewAppCallsItemChangeScreenSharingRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAppCallsCallItemRequestBuilderInternal instantiates a new CallItemRequestBuilder and sets the default values.
func NewAppCallsCallItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppCallsCallItemRequestBuilder) {
    m := &AppCallsCallItemRequestBuilder{
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
// NewAppCallsCallItemRequestBuilder instantiates a new CallItemRequestBuilder and sets the default values.
func NewAppCallsCallItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppCallsCallItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppCallsCallItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentSharingSessions provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
func (m *AppCallsCallItemRequestBuilder) ContentSharingSessions()(*AppCallsItemContentSharingSessionsRequestBuilder) {
    return NewAppCallsItemContentSharingSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentSharingSessionsById provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
func (m *AppCallsCallItemRequestBuilder) ContentSharingSessionsById(id string)(*AppCallsItemContentSharingSessionsContentSharingSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentSharingSession%2Did"] = id
    }
    return NewAppCallsItemContentSharingSessionsContentSharingSessionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property calls for app
func (m *AppCallsCallItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *AppCallsCallItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AppCallsCallItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AppCallsCallItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AppCallsCallItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, requestConfiguration *AppCallsCallItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *AppCallsCallItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppCallsCallItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *AppCallsCallItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppCallsCallItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, error) {
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
// KeepAlive provides operations to call the keepAlive method.
func (m *AppCallsCallItemRequestBuilder) KeepAlive()(*AppCallsItemKeepAliveRequestBuilder) {
    return NewAppCallsItemKeepAliveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Mute provides operations to call the mute method.
func (m *AppCallsCallItemRequestBuilder) Mute()(*AppCallsItemMuteRequestBuilder) {
    return NewAppCallsItemMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.call entity.
func (m *AppCallsCallItemRequestBuilder) Operations()(*AppCallsItemOperationsRequestBuilder) {
    return NewAppCallsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.call entity.
func (m *AppCallsCallItemRequestBuilder) OperationsById(id string)(*AppCallsItemOperationsCommsOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["commsOperation%2Did"] = id
    }
    return NewAppCallsItemOperationsCommsOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Participants provides operations to manage the participants property of the microsoft.graph.call entity.
func (m *AppCallsCallItemRequestBuilder) Participants()(*AppCallsItemParticipantsRequestBuilder) {
    return NewAppCallsItemParticipantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParticipantsById provides operations to manage the participants property of the microsoft.graph.call entity.
func (m *AppCallsCallItemRequestBuilder) ParticipantsById(id string)(*AppCallsItemParticipantsParticipantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["participant%2Did"] = id
    }
    return NewAppCallsItemParticipantsParticipantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calls in app
func (m *AppCallsCallItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, requestConfiguration *AppCallsCallItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, error) {
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
// PlayPrompt provides operations to call the playPrompt method.
func (m *AppCallsCallItemRequestBuilder) PlayPrompt()(*AppCallsItemPlayPromptRequestBuilder) {
    return NewAppCallsItemPlayPromptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Record provides operations to call the record method.
func (m *AppCallsCallItemRequestBuilder) Record()(*AppCallsItemRecordRequestBuilder) {
    return NewAppCallsItemRecordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecordResponse provides operations to call the recordResponse method.
func (m *AppCallsCallItemRequestBuilder) RecordResponse()(*AppCallsItemRecordResponseRequestBuilder) {
    return NewAppCallsItemRecordResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Redirect provides operations to call the redirect method.
func (m *AppCallsCallItemRequestBuilder) Redirect()(*AppCallsItemRedirectRequestBuilder) {
    return NewAppCallsItemRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reject provides operations to call the reject method.
func (m *AppCallsCallItemRequestBuilder) Reject()(*AppCallsItemRejectRequestBuilder) {
    return NewAppCallsItemRejectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeToTone provides operations to call the subscribeToTone method.
func (m *AppCallsCallItemRequestBuilder) SubscribeToTone()(*AppCallsItemSubscribeToToneRequestBuilder) {
    return NewAppCallsItemSubscribeToToneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Transfer provides operations to call the transfer method.
func (m *AppCallsCallItemRequestBuilder) Transfer()(*AppCallsItemTransferRequestBuilder) {
    return NewAppCallsItemTransferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unmute provides operations to call the unmute method.
func (m *AppCallsCallItemRequestBuilder) Unmute()(*AppCallsItemUnmuteRequestBuilder) {
    return NewAppCallsItemUnmuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateRecordingStatus provides operations to call the updateRecordingStatus method.
func (m *AppCallsCallItemRequestBuilder) UpdateRecordingStatus()(*AppCallsItemUpdateRecordingStatusRequestBuilder) {
    return NewAppCallsItemUpdateRecordingStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
