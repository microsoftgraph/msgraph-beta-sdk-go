package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallsItemParticipantsItemMuteRequestBuilder provides operations to call the mute method.
type CallsItemParticipantsItemMuteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallsItemParticipantsItemMuteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemParticipantsItemMuteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemParticipantsItemMuteRequestBuilderInternal instantiates a new MuteRequestBuilder and sets the default values.
func NewCallsItemParticipantsItemMuteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemParticipantsItemMuteRequestBuilder) {
    m := &CallsItemParticipantsItemMuteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/calls/{call%2Did}/participants/{participant%2Did}/mute", pathParameters),
    }
    return m
}
// NewCallsItemParticipantsItemMuteRequestBuilder instantiates a new MuteRequestBuilder and sets the default values.
func NewCallsItemParticipantsItemMuteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemParticipantsItemMuteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemParticipantsItemMuteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post mute a specific participant in the call. This is a server mute, meaning that the server will drop all audio packets for this participant, even if the participant continues to stream audio. For more information about how to handle mute operations, see muteParticipantOperation. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/participant-mute?view=graph-rest-1.0
func (m *CallsItemParticipantsItemMuteRequestBuilder) Post(ctx context.Context, body CallsItemParticipantsItemMutePostRequestBodyable, requestConfiguration *CallsItemParticipantsItemMuteRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MuteParticipantOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMuteParticipantOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MuteParticipantOperationable), nil
}
// ToPostRequestInformation mute a specific participant in the call. This is a server mute, meaning that the server will drop all audio packets for this participant, even if the participant continues to stream audio. For more information about how to handle mute operations, see muteParticipantOperation. This API is available in the following national cloud deployments.
func (m *CallsItemParticipantsItemMuteRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemParticipantsItemMutePostRequestBodyable, requestConfiguration *CallsItemParticipantsItemMuteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CallsItemParticipantsItemMuteRequestBuilder) WithUrl(rawUrl string)(*CallsItemParticipantsItemMuteRequestBuilder) {
    return NewCallsItemParticipantsItemMuteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
