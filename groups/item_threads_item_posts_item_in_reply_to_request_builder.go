package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemThreadsItemPostsItemInReplyToRequestBuilder provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
type ItemThreadsItemPostsItemInReplyToRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
type ItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.post entity.
// returns a *ItemThreadsItemPostsItemInReplyToAttachmentsRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) Attachments()(*ItemThreadsItemPostsItemInReplyToAttachmentsRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemThreadsItemPostsItemInReplyToRequestBuilderInternal instantiates a new ItemThreadsItemPostsItemInReplyToRequestBuilder and sets the default values.
func NewItemThreadsItemPostsItemInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsItemInReplyToRequestBuilder) {
    m := &ItemThreadsItemPostsItemInReplyToRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemThreadsItemPostsItemInReplyToRequestBuilder instantiates a new ItemThreadsItemPostsItemInReplyToRequestBuilder and sets the default values.
func NewItemThreadsItemPostsItemInReplyToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsItemInReplyToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemThreadsItemPostsItemInReplyToRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.post entity.
// returns a *ItemThreadsItemPostsItemInReplyToExtensionsRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) Extensions()(*ItemThreadsItemPostsItemInReplyToExtensionsRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
// returns a *ItemThreadsItemPostsItemInReplyToForwardRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) Forward()(*ItemThreadsItemPostsItemInReplyToForwardRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
// returns a Postable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePostFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable), nil
}
// Mentions provides operations to manage the mentions property of the microsoft.graph.post entity.
// returns a *ItemThreadsItemPostsItemInReplyToMentionsRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) Mentions()(*ItemThreadsItemPostsItemInReplyToMentionsRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToMentionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reply provides operations to call the reply method.
// returns a *ItemThreadsItemPostsItemInReplyToReplyRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) Reply()(*ItemThreadsItemPostsItemInReplyToReplyRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToReplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
// returns a *RequestInformation when successful
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemThreadsItemPostsItemInReplyToRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) WithUrl(rawUrl string)(*ItemThreadsItemPostsItemInReplyToRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
