package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i084cca04694d7243d666a4e58b6e392d82016a6a5d4aa04c201f29832305b2e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/mentions"
    i44baed2c17fa38869d67d057e4c0d223d25110d66a2a95345a0bdb48877c0f3a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/singlevalueextendedproperties"
    i731c269ef34ed90b509eddde0670a48b0520449542c2f5bf3d8fbb64284f955d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/forward"
    ia03c6902718dab8649fcf59dff44ab40283c06a322d7c371356cd2abf839aecf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/attachments"
    ib2d9f36420ff40a0de6040cb63fbd2964c4989f2b92dc280958fbc73d398799e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto"
    ide057bbb00756c6bd6a940eaf4cf1a97f2e7f0490c7f289bd49a5ccb22ffc45c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/extensions"
    ieb577962bb6d6117e6c2ea2a89352bc91f7e5d1f6a54dc95871cd113e314ea23 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/reply"
    ifa6728321f318f720e802b218856ed51ac05b7cfab5379556456e9135fc37b28 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/multivalueextendedproperties"
    i144cb64e4df1ff1e41bbb0995b5cc766e18663a5e22f97b1c64f55bd7d480e24 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/multivalueextendedproperties/item"
    i172317af95293eeaa84b69051a0b2811497bcf97abf5fdee59e6a97308cea3b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/attachments/item"
    i40afb8409c3cd306f5f1d742239c52e860ba7e5d16e55a25d7ada66956341593 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/singlevalueextendedproperties/item"
    i54e5c839e9742864d08e23f2eaafb83784a9bb1c81668ee263dc82a5c550d289 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/extensions/item"
    ia6c595c11a0789c6b1b1e5caf8c774003cdf92cbdf69235cba53e03fe539741a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/mentions/item"
)

// PostItemRequestBuilder provides operations to manage the posts property of the microsoft.graph.conversationThread entity.
type PostItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PostItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PostItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PostItemRequestBuilderGetQueryParameters read-only. Nullable.
type PostItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PostItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PostItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PostItemRequestBuilderGetQueryParameters
}
// PostItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PostItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments the attachments property
func (m *PostItemRequestBuilder) Attachments()(*ia03c6902718dab8649fcf59dff44ab40283c06a322d7c371356cd2abf839aecf.AttachmentsRequestBuilder) {
    return ia03c6902718dab8649fcf59dff44ab40283c06a322d7c371356cd2abf839aecf.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.conversations.item.threads.item.posts.item.attachments.item collection
func (m *PostItemRequestBuilder) AttachmentsById(id string)(*i172317af95293eeaa84b69051a0b2811497bcf97abf5fdee59e6a97308cea3b8.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i172317af95293eeaa84b69051a0b2811497bcf97abf5fdee59e6a97308cea3b8.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPostItemRequestBuilderInternal instantiates a new PostItemRequestBuilder and sets the default values.
func NewPostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostItemRequestBuilder) {
    m := &PostItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPostItemRequestBuilder instantiates a new PostItemRequestBuilder and sets the default values.
func NewPostItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPostItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property posts for me
func (m *PostItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property posts for me
func (m *PostItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *PostItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation read-only. Nullable.
func (m *PostItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration read-only. Nullable.
func (m *PostItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PostItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property posts in me
func (m *PostItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property posts in me
func (m *PostItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, requestConfiguration *PostItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property posts for me
func (m *PostItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property posts for me
func (m *PostItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *PostItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Extensions the extensions property
func (m *PostItemRequestBuilder) Extensions()(*ide057bbb00756c6bd6a940eaf4cf1a97f2e7f0490c7f289bd49a5ccb22ffc45c.ExtensionsRequestBuilder) {
    return ide057bbb00756c6bd6a940eaf4cf1a97f2e7f0490c7f289bd49a5ccb22ffc45c.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.conversations.item.threads.item.posts.item.extensions.item collection
func (m *PostItemRequestBuilder) ExtensionsById(id string)(*i54e5c839e9742864d08e23f2eaafb83784a9bb1c81668ee263dc82a5c550d289.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i54e5c839e9742864d08e23f2eaafb83784a9bb1c81668ee263dc82a5c550d289.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *PostItemRequestBuilder) Forward()(*i731c269ef34ed90b509eddde0670a48b0520449542c2f5bf3d8fbb64284f955d.ForwardRequestBuilder) {
    return i731c269ef34ed90b509eddde0670a48b0520449542c2f5bf3d8fbb64284f955d.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read-only. Nullable.
func (m *PostItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler read-only. Nullable.
func (m *PostItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *PostItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePostFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable), nil
}
// InReplyTo the inReplyTo property
func (m *PostItemRequestBuilder) InReplyTo()(*ib2d9f36420ff40a0de6040cb63fbd2964c4989f2b92dc280958fbc73d398799e.InReplyToRequestBuilder) {
    return ib2d9f36420ff40a0de6040cb63fbd2964c4989f2b92dc280958fbc73d398799e.NewInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Mentions the mentions property
func (m *PostItemRequestBuilder) Mentions()(*i084cca04694d7243d666a4e58b6e392d82016a6a5d4aa04c201f29832305b2e9.MentionsRequestBuilder) {
    return i084cca04694d7243d666a4e58b6e392d82016a6a5d4aa04c201f29832305b2e9.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.conversations.item.threads.item.posts.item.mentions.item collection
func (m *PostItemRequestBuilder) MentionsById(id string)(*ia6c595c11a0789c6b1b1e5caf8c774003cdf92cbdf69235cba53e03fe539741a.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return ia6c595c11a0789c6b1b1e5caf8c774003cdf92cbdf69235cba53e03fe539741a.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *PostItemRequestBuilder) MultiValueExtendedProperties()(*ifa6728321f318f720e802b218856ed51ac05b7cfab5379556456e9135fc37b28.MultiValueExtendedPropertiesRequestBuilder) {
    return ifa6728321f318f720e802b218856ed51ac05b7cfab5379556456e9135fc37b28.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.conversations.item.threads.item.posts.item.multiValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i144cb64e4df1ff1e41bbb0995b5cc766e18663a5e22f97b1c64f55bd7d480e24.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i144cb64e4df1ff1e41bbb0995b5cc766e18663a5e22f97b1c64f55bd7d480e24.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property posts in me
func (m *PostItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property posts in me
func (m *PostItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, requestConfiguration *PostItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Reply the reply property
func (m *PostItemRequestBuilder) Reply()(*ieb577962bb6d6117e6c2ea2a89352bc91f7e5d1f6a54dc95871cd113e314ea23.ReplyRequestBuilder) {
    return ieb577962bb6d6117e6c2ea2a89352bc91f7e5d1f6a54dc95871cd113e314ea23.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *PostItemRequestBuilder) SingleValueExtendedProperties()(*i44baed2c17fa38869d67d057e4c0d223d25110d66a2a95345a0bdb48877c0f3a.SingleValueExtendedPropertiesRequestBuilder) {
    return i44baed2c17fa38869d67d057e4c0d223d25110d66a2a95345a0bdb48877c0f3a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.conversations.item.threads.item.posts.item.singleValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i40afb8409c3cd306f5f1d742239c52e860ba7e5d16e55a25d7ada66956341593.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i40afb8409c3cd306f5f1d742239c52e860ba7e5d16e55a25d7ada66956341593.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
