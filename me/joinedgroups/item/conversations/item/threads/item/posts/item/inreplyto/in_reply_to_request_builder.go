package inreplyto

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1bd03a03ccc5778ebe032e21624ee67d30b2d6812c5d1e65d751634b7beb3550 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/multivalueextendedproperties"
    i4287f7c4e39227d67f6ae6a66f92d0a207824b158481ad3b74763c98eafe2499 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/mentions"
    i7bdaedd8c06fd6eb08140ba00ff2b3e75d4ac761f03ee08dcc84fb3bd2f07554 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/extensions"
    i8fb630c4d4791244d93f714abcdce04fa46ba93354dd8365befade74ab3516b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties"
    ia575107bcdb2821519ed668914bef00be8edada06e303aaa3ebaa604fe3bdd86 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/reply"
    icb235ecf85f0d70eeab6fc65ea383e345e27f3e84ec1f707841250b6f2dd8fb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/attachments"
    ie5f61e956a484d5d44d60d3ccc1bc9964a8485036eac89e004560f1f142fecdf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/forward"
    i0d068d4e30e861c15c0b8cba431cc827cb813375f1668e900a1cc742c7f01df2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties/item"
    i51a2d71c23614dc629806d158dafe0c4b88aad218dbab3db5b0d4a64307e6027 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/multivalueextendedproperties/item"
    i8019c278f8d32a1bd4f75865d088dd1f92d656500e4fcef753f4a74067b96472 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/extensions/item"
    i969a1b13cf01993a2c4684fe835a44bf90b723eae6c3cd2e7971f401dde1baba "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/attachments/item"
    ib0247a8a5e5c874b050245b97d35b1b380862e9f7d847134e5402df8487cfbe0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/mentions/item"
)

// InReplyToRequestBuilder provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
type InReplyToRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InReplyToRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InReplyToRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InReplyToRequestBuilderGetQueryParameters the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
type InReplyToRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// InReplyToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InReplyToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InReplyToRequestBuilderGetQueryParameters
}
// InReplyToRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InReplyToRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments the attachments property
func (m *InReplyToRequestBuilder) Attachments()(*icb235ecf85f0d70eeab6fc65ea383e345e27f3e84ec1f707841250b6f2dd8fb7.AttachmentsRequestBuilder) {
    return icb235ecf85f0d70eeab6fc65ea383e345e27f3e84ec1f707841250b6f2dd8fb7.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.conversations.item.threads.item.posts.item.inReplyTo.attachments.item collection
func (m *InReplyToRequestBuilder) AttachmentsById(id string)(*i969a1b13cf01993a2c4684fe835a44bf90b723eae6c3cd2e7971f401dde1baba.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i969a1b13cf01993a2c4684fe835a44bf90b723eae6c3cd2e7971f401dde1baba.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewInReplyToRequestBuilderInternal instantiates a new InReplyToRequestBuilder and sets the default values.
func NewInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InReplyToRequestBuilder) {
    m := &InReplyToRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInReplyToRequestBuilder instantiates a new InReplyToRequestBuilder and sets the default values.
func NewInReplyToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InReplyToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInReplyToRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property inReplyTo for me
func (m *InReplyToRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property inReplyTo for me
func (m *InReplyToRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *InReplyToRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
func (m *InReplyToRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
func (m *InReplyToRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *InReplyToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property inReplyTo in me
func (m *InReplyToRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property inReplyTo in me
func (m *InReplyToRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, requestConfiguration *InReplyToRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property inReplyTo for me
func (m *InReplyToRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property inReplyTo for me
func (m *InReplyToRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *InReplyToRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *InReplyToRequestBuilder) Extensions()(*i7bdaedd8c06fd6eb08140ba00ff2b3e75d4ac761f03ee08dcc84fb3bd2f07554.ExtensionsRequestBuilder) {
    return i7bdaedd8c06fd6eb08140ba00ff2b3e75d4ac761f03ee08dcc84fb3bd2f07554.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.conversations.item.threads.item.posts.item.inReplyTo.extensions.item collection
func (m *InReplyToRequestBuilder) ExtensionsById(id string)(*i8019c278f8d32a1bd4f75865d088dd1f92d656500e4fcef753f4a74067b96472.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i8019c278f8d32a1bd4f75865d088dd1f92d656500e4fcef753f4a74067b96472.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *InReplyToRequestBuilder) Forward()(*ie5f61e956a484d5d44d60d3ccc1bc9964a8485036eac89e004560f1f142fecdf.ForwardRequestBuilder) {
    return ie5f61e956a484d5d44d60d3ccc1bc9964a8485036eac89e004560f1f142fecdf.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
func (m *InReplyToRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
func (m *InReplyToRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *InReplyToRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, error) {
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
// Mentions the mentions property
func (m *InReplyToRequestBuilder) Mentions()(*i4287f7c4e39227d67f6ae6a66f92d0a207824b158481ad3b74763c98eafe2499.MentionsRequestBuilder) {
    return i4287f7c4e39227d67f6ae6a66f92d0a207824b158481ad3b74763c98eafe2499.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.conversations.item.threads.item.posts.item.inReplyTo.mentions.item collection
func (m *InReplyToRequestBuilder) MentionsById(id string)(*ib0247a8a5e5c874b050245b97d35b1b380862e9f7d847134e5402df8487cfbe0.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return ib0247a8a5e5c874b050245b97d35b1b380862e9f7d847134e5402df8487cfbe0.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *InReplyToRequestBuilder) MultiValueExtendedProperties()(*i1bd03a03ccc5778ebe032e21624ee67d30b2d6812c5d1e65d751634b7beb3550.MultiValueExtendedPropertiesRequestBuilder) {
    return i1bd03a03ccc5778ebe032e21624ee67d30b2d6812c5d1e65d751634b7beb3550.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.conversations.item.threads.item.posts.item.inReplyTo.multiValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i51a2d71c23614dc629806d158dafe0c4b88aad218dbab3db5b0d4a64307e6027.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i51a2d71c23614dc629806d158dafe0c4b88aad218dbab3db5b0d4a64307e6027.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property inReplyTo in me
func (m *InReplyToRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property inReplyTo in me
func (m *InReplyToRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, requestConfiguration *InReplyToRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *InReplyToRequestBuilder) Reply()(*ia575107bcdb2821519ed668914bef00be8edada06e303aaa3ebaa604fe3bdd86.ReplyRequestBuilder) {
    return ia575107bcdb2821519ed668914bef00be8edada06e303aaa3ebaa604fe3bdd86.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *InReplyToRequestBuilder) SingleValueExtendedProperties()(*i8fb630c4d4791244d93f714abcdce04fa46ba93354dd8365befade74ab3516b4.SingleValueExtendedPropertiesRequestBuilder) {
    return i8fb630c4d4791244d93f714abcdce04fa46ba93354dd8365befade74ab3516b4.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.conversations.item.threads.item.posts.item.inReplyTo.singleValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i0d068d4e30e861c15c0b8cba431cc827cb813375f1668e900a1cc742c7f01df2.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i0d068d4e30e861c15c0b8cba431cc827cb813375f1668e900a1cc742c7f01df2.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
