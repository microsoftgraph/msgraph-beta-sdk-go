package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder provides operations to manage the transcripts property of the microsoft.graph.onlineMeeting entity.
type OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderGetQueryParameters retrieve a callTranscript object associated with a scheduled onlineMeeting. This API doesn't support getting call transcripts from channel meetings. Retrieving the transcript returns the metadata of the single transcript associated with the online meeting. Retrieving the content of the transcript returns the stream of text associated with the transcript.
type OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderGetQueryParameters
}
// OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderInternal instantiates a new OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder and sets the default values.
func NewOnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) {
    m := &OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/app/onlineMeetings/{onlineMeeting%2Did}/transcripts/{callTranscript%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewOnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder instantiates a new OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder and sets the default values.
func NewOnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the commsApplication entity.
// returns a *OnlineMeetingsItemTranscriptsItemContentRequestBuilder when successful
func (m *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) Content()(*OnlineMeetingsItemTranscriptsItemContentRequestBuilder) {
    return NewOnlineMeetingsItemTranscriptsItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property transcripts for app
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve a callTranscript object associated with a scheduled onlineMeeting. This API doesn't support getting call transcripts from channel meetings. Retrieving the transcript returns the metadata of the single transcript associated with the online meeting. Retrieving the content of the transcript returns the stream of text associated with the transcript.
// returns a CallTranscriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/calltranscript-get?view=graph-rest-1.0
func (m *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallTranscriptable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCallTranscriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallTranscriptable), nil
}
// MetadataContent provides operations to manage the media for the commsApplication entity.
// returns a *OnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder when successful
func (m *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) MetadataContent()(*OnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilder) {
    return NewOnlineMeetingsItemTranscriptsItemMetadataContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property transcripts in app
// returns a CallTranscriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallTranscriptable, requestConfiguration *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallTranscriptable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCallTranscriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallTranscriptable), nil
}
// ToDeleteRequestInformation delete navigation property transcripts for app
// returns a *RequestInformation when successful
func (m *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/app/onlineMeetings/{onlineMeeting%2Did}/transcripts/{callTranscript%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve a callTranscript object associated with a scheduled onlineMeeting. This API doesn't support getting call transcripts from channel meetings. Retrieving the transcript returns the metadata of the single transcript associated with the online meeting. Retrieving the content of the transcript returns the stream of text associated with the transcript.
// returns a *RequestInformation when successful
func (m *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property transcripts in app
// returns a *RequestInformation when successful
func (m *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallTranscriptable, requestConfiguration *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/app/onlineMeetings/{onlineMeeting%2Did}/transcripts/{callTranscript%2Did}", m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder when successful
func (m *OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) WithUrl(rawUrl string)(*OnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder) {
    return NewOnlineMeetingsItemTranscriptsCallTranscriptItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
