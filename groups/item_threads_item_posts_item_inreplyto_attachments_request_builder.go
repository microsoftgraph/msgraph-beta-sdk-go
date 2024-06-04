package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder provides operations to manage the attachments property of the microsoft.graph.post entity.
type ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderGetQueryParameters the collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the post. Read-only. Nullable. Supports $expand.
type ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderGetQueryParameters
}
// ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAttachmentId provides operations to manage the attachments property of the microsoft.graph.post entity.
// returns a *ItemThreadsItemPostsItemInreplytoAttachmentsAttachmentItemRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) ByAttachmentId(attachmentId string)(*ItemThreadsItemPostsItemInreplytoAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if attachmentId != "" {
        urlTplParams["attachment%2Did"] = attachmentId
    }
    return NewItemThreadsItemPostsItemInreplytoAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderInternal instantiates a new ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder and sets the default values.
func NewItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) {
    m := &ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo/attachments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder instantiates a new ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder and sets the default values.
func NewItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemThreadsItemPostsItemInreplytoAttachmentsCountRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) Count()(*ItemThreadsItemPostsItemInreplytoAttachmentsCountRequestBuilder) {
    return NewItemThreadsItemPostsItemInreplytoAttachmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateUploadSession provides operations to call the createUploadSession method.
// returns a *ItemThreadsItemPostsItemInreplytoAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) CreateUploadSession()(*ItemThreadsItemPostsItemInreplytoAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilder) {
    return NewItemThreadsItemPostsItemInreplytoAttachmentsCreateuploadsessionCreateUploadSessionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the post. Read-only. Nullable. Supports $expand.
// returns a AttachmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttachmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttachmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttachmentCollectionResponseable), nil
}
// Post create new navigation property to attachments for groups
// returns a Attachmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, requestConfiguration *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttachmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable), nil
}
// ToGetRequestInformation the collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the post. Read-only. Nullable. Supports $expand.
// returns a *RequestInformation when successful
func (m *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to attachments for groups
// returns a *RequestInformation when successful
func (m *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Attachmentable, requestConfiguration *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder when successful
func (m *ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) WithUrl(rawUrl string)(*ItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder) {
    return NewItemThreadsItemPostsItemInreplytoAttachmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
