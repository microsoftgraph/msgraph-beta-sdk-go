// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallsCallItemRequestBuilder provides operations to manage the calls property of the microsoft.graph.commsApplication entity.
type CallsCallItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallsCallItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsCallItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CallsCallItemRequestBuilderGetQueryParameters get calls from app
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
// returns a *CallsItemAddLargeGalleryViewRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) AddLargeGalleryView()(*CallsItemAddLargeGalleryViewRequestBuilder) {
    return NewCallsItemAddLargeGalleryViewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Answer provides operations to call the answer method.
// returns a *CallsItemAnswerRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) Answer()(*CallsItemAnswerRequestBuilder) {
    return NewCallsItemAnswerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AudioRoutingGroups provides operations to manage the audioRoutingGroups property of the microsoft.graph.call entity.
// returns a *CallsItemAudioRoutingGroupsRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) AudioRoutingGroups()(*CallsItemAudioRoutingGroupsRequestBuilder) {
    return NewCallsItemAudioRoutingGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CancelMediaProcessing provides operations to call the cancelMediaProcessing method.
// returns a *CallsItemCancelMediaProcessingRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) CancelMediaProcessing()(*CallsItemCancelMediaProcessingRequestBuilder) {
    return NewCallsItemCancelMediaProcessingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChangeScreenSharingRole provides operations to call the changeScreenSharingRole method.
// returns a *CallsItemChangeScreenSharingRoleRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) ChangeScreenSharingRole()(*CallsItemChangeScreenSharingRoleRequestBuilder) {
    return NewCallsItemChangeScreenSharingRoleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCallsCallItemRequestBuilderInternal instantiates a new CallsCallItemRequestBuilder and sets the default values.
func NewCallsCallItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsCallItemRequestBuilder) {
    m := &CallsCallItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/calls/{call%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCallsCallItemRequestBuilder instantiates a new CallsCallItemRequestBuilder and sets the default values.
func NewCallsCallItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsCallItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsCallItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentSharingSessions provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
// returns a *CallsItemContentSharingSessionsRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) ContentSharingSessions()(*CallsItemContentSharingSessionsRequestBuilder) {
    return NewCallsItemContentSharingSessionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property calls for app
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallsCallItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CallsCallItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get calls from app
// returns a Callable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallsCallItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CallsCallItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCallFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable), nil
}
// KeepAlive provides operations to call the keepAlive method.
// returns a *CallsItemKeepAliveRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) KeepAlive()(*CallsItemKeepAliveRequestBuilder) {
    return NewCallsItemKeepAliveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Mute provides operations to call the mute method.
// returns a *CallsItemMuteRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) Mute()(*CallsItemMuteRequestBuilder) {
    return NewCallsItemMuteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.call entity.
// returns a *CallsItemOperationsRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) Operations()(*CallsItemOperationsRequestBuilder) {
    return NewCallsItemOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Participants provides operations to manage the participants property of the microsoft.graph.call entity.
// returns a *CallsItemParticipantsRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) Participants()(*CallsItemParticipantsRequestBuilder) {
    return NewCallsItemParticipantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property calls in app
// returns a Callable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallsCallItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, requestConfiguration *CallsCallItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCallFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable), nil
}
// PlayPrompt provides operations to call the playPrompt method.
// returns a *CallsItemPlayPromptRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) PlayPrompt()(*CallsItemPlayPromptRequestBuilder) {
    return NewCallsItemPlayPromptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Record provides operations to call the record method.
// returns a *CallsItemRecordRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) Record()(*CallsItemRecordRequestBuilder) {
    return NewCallsItemRecordRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RecordResponse provides operations to call the recordResponse method.
// returns a *CallsItemRecordResponseRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) RecordResponse()(*CallsItemRecordResponseRequestBuilder) {
    return NewCallsItemRecordResponseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Redirect provides operations to call the redirect method.
// returns a *CallsItemRedirectRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) Redirect()(*CallsItemRedirectRequestBuilder) {
    return NewCallsItemRedirectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reject provides operations to call the reject method.
// returns a *CallsItemRejectRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) Reject()(*CallsItemRejectRequestBuilder) {
    return NewCallsItemRejectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendDtmfTones provides operations to call the sendDtmfTones method.
// returns a *CallsItemSendDtmfTonesRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) SendDtmfTones()(*CallsItemSendDtmfTonesRequestBuilder) {
    return NewCallsItemSendDtmfTonesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StartRecording provides operations to call the startRecording method.
// returns a *CallsItemStartRecordingRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) StartRecording()(*CallsItemStartRecordingRequestBuilder) {
    return NewCallsItemStartRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StartTranscription provides operations to call the startTranscription method.
// returns a *CallsItemStartTranscriptionRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) StartTranscription()(*CallsItemStartTranscriptionRequestBuilder) {
    return NewCallsItemStartTranscriptionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StopRecording provides operations to call the stopRecording method.
// returns a *CallsItemStopRecordingRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) StopRecording()(*CallsItemStopRecordingRequestBuilder) {
    return NewCallsItemStopRecordingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// StopTranscription provides operations to call the stopTranscription method.
// returns a *CallsItemStopTranscriptionRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) StopTranscription()(*CallsItemStopTranscriptionRequestBuilder) {
    return NewCallsItemStopTranscriptionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SubscribeToTone provides operations to call the subscribeToTone method.
// returns a *CallsItemSubscribeToToneRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) SubscribeToTone()(*CallsItemSubscribeToToneRequestBuilder) {
    return NewCallsItemSubscribeToToneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property calls for app
// returns a *RequestInformation when successful
func (m *CallsCallItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CallsCallItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get calls from app
// returns a *RequestInformation when successful
func (m *CallsCallItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallsCallItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property calls in app
// returns a *RequestInformation when successful
func (m *CallsCallItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Callable, requestConfiguration *CallsCallItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Transfer provides operations to call the transfer method.
// returns a *CallsItemTransferRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) Transfer()(*CallsItemTransferRequestBuilder) {
    return NewCallsItemTransferRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Unmute provides operations to call the unmute method.
// returns a *CallsItemUnmuteRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) Unmute()(*CallsItemUnmuteRequestBuilder) {
    return NewCallsItemUnmuteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateRecordingStatus provides operations to call the updateRecordingStatus method.
// returns a *CallsItemUpdateRecordingStatusRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) UpdateRecordingStatus()(*CallsItemUpdateRecordingStatusRequestBuilder) {
    return NewCallsItemUpdateRecordingStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CallsCallItemRequestBuilder when successful
func (m *CallsCallItemRequestBuilder) WithUrl(rawUrl string)(*CallsCallItemRequestBuilder) {
    return NewCallsCallItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
