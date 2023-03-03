package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallsCallItemRequestBuilder provides operations to manage the calls property of the microsoft.graph.cloudCommunications entity.
type CallsCallItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallsCallItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsCallItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CallsCallItemRequestBuilderGetQueryParameters get calls from communications
type CallsCallItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CallsCallItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsCallItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallsCallItemRequestBuilderGetQueryParameters
}
// CallsCallItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsCallItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddLargeGalleryView provides operations to call the addLargeGalleryView method.
func (m *CallsCallItemRequestBuilder) AddLargeGalleryView()(*CallsItemAddLargeGalleryViewRequestBuilder) {
    return NewCallsItemAddLargeGalleryViewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Answer provides operations to call the answer method.
func (m *CallsCallItemRequestBuilder) Answer()(*CallsItemAnswerRequestBuilder) {
    return NewCallsItemAnswerRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AudioRoutingGroups provides operations to manage the audioRoutingGroups property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) AudioRoutingGroups()(*CallsItemAudioRoutingGroupsRequestBuilder) {
    return NewCallsItemAudioRoutingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AudioRoutingGroupsById provides operations to manage the audioRoutingGroups property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) AudioRoutingGroupsById(id string)(*CallsItemAudioRoutingGroupsAudioRoutingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["audioRoutingGroup%2Did"] = id
    }
    return NewCallsItemAudioRoutingGroupsAudioRoutingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// CancelMediaProcessing provides operations to call the cancelMediaProcessing method.
func (m *CallsCallItemRequestBuilder) CancelMediaProcessing()(*CallsItemCancelMediaProcessingRequestBuilder) {
    return NewCallsItemCancelMediaProcessingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChangeScreenSharingRole provides operations to call the changeScreenSharingRole method.
func (m *CallsCallItemRequestBuilder) ChangeScreenSharingRole()(*CallsItemChangeScreenSharingRoleRequestBuilder) {
    return NewCallsItemChangeScreenSharingRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewCallsCallItemRequestBuilderInternal instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallsCallItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsCallItemRequestBuilder) {
    m := &CallsCallItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCallsCallItemRequestBuilder instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallsCallItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsCallItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsCallItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentSharingSessions provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) ContentSharingSessions()(*CallsItemContentSharingSessionsRequestBuilder) {
    return NewCallsItemContentSharingSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ContentSharingSessionsById provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) ContentSharingSessionsById(id string)(*CallsItemContentSharingSessionsContentSharingSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentSharingSession%2Did"] = id
    }
    return NewCallsItemContentSharingSessionsContentSharingSessionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Delete delete navigation property calls for communications
func (m *CallsCallItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CallsCallItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get calls from communications
func (m *CallsCallItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CallsCallItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCallFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable), nil
}
// KeepAlive provides operations to call the keepAlive method.
func (m *CallsCallItemRequestBuilder) KeepAlive()(*CallsItemKeepAliveRequestBuilder) {
    return NewCallsItemKeepAliveRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Mute provides operations to call the mute method.
func (m *CallsCallItemRequestBuilder) Mute()(*CallsItemMuteRequestBuilder) {
    return NewCallsItemMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) Operations()(*CallsItemOperationsRequestBuilder) {
    return NewCallsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) OperationsById(id string)(*CallsItemOperationsCommsOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["commsOperation%2Did"] = id
    }
    return NewCallsItemOperationsCommsOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Participants provides operations to manage the participants property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) Participants()(*CallsItemParticipantsRequestBuilder) {
    return NewCallsItemParticipantsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ParticipantsById provides operations to manage the participants property of the microsoft.graph.call entity.
func (m *CallsCallItemRequestBuilder) ParticipantsById(id string)(*CallsItemParticipantsParticipantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["participant%2Did"] = id
    }
    return NewCallsItemParticipantsParticipantItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Patch update the navigation property calls in communications
func (m *CallsCallItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, requestConfiguration *CallsCallItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCallFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable), nil
}
// PlayPrompt provides operations to call the playPrompt method.
func (m *CallsCallItemRequestBuilder) PlayPrompt()(*CallsItemPlayPromptRequestBuilder) {
    return NewCallsItemPlayPromptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Record provides operations to call the record method.
func (m *CallsCallItemRequestBuilder) Record()(*CallsItemRecordRequestBuilder) {
    return NewCallsItemRecordRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RecordResponse provides operations to call the recordResponse method.
func (m *CallsCallItemRequestBuilder) RecordResponse()(*CallsItemRecordResponseRequestBuilder) {
    return NewCallsItemRecordResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Redirect provides operations to call the redirect method.
func (m *CallsCallItemRequestBuilder) Redirect()(*CallsItemRedirectRequestBuilder) {
    return NewCallsItemRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Reject provides operations to call the reject method.
func (m *CallsCallItemRequestBuilder) Reject()(*CallsItemRejectRequestBuilder) {
    return NewCallsItemRejectRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubscribeToTone provides operations to call the subscribeToTone method.
func (m *CallsCallItemRequestBuilder) SubscribeToTone()(*CallsItemSubscribeToToneRequestBuilder) {
    return NewCallsItemSubscribeToToneRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property calls for communications
func (m *CallsCallItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CallsCallItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get calls from communications
func (m *CallsCallItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallsCallItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property calls in communications
func (m *CallsCallItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, requestConfiguration *CallsCallItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Transfer provides operations to call the transfer method.
func (m *CallsCallItemRequestBuilder) Transfer()(*CallsItemTransferRequestBuilder) {
    return NewCallsItemTransferRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Unmute provides operations to call the unmute method.
func (m *CallsCallItemRequestBuilder) Unmute()(*CallsItemUnmuteRequestBuilder) {
    return NewCallsItemUnmuteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UpdateRecordingStatus provides operations to call the updateRecordingStatus method.
func (m *CallsCallItemRequestBuilder) UpdateRecordingStatus()(*CallsItemUpdateRecordingStatusRequestBuilder) {
    return NewCallsItemUpdateRecordingStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
