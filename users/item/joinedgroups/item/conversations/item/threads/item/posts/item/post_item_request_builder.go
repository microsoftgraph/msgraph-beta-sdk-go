package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2296c4df4c69facd8f70733a4445a4252bdbb05b1f267b8f56766dbb1a637ee1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/multivalueextendedproperties"
    i4239b8918b5c4cbf4765ccbe6679428109b6a5b2b76a1cfd452061913232d627 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto"
    ic3c754f8a7f3e73343870eab0bed80d1b24d51a9db00d2e802a2fd0f767c6deb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/forward"
    id465088205bcdf8ea6b51c829a964dedce007ceb93287121a6841c6530f29ea3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/mentions"
    if637eaef85087127f12c2411e01352d92aa284812ee2eabb26cb1fc7d3ff4da5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/extensions"
    if7184ad3e4dbd87ee02e456222ee430d3cb4f00e04ca859c471acae8e4071a38 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/singlevalueextendedproperties"
    if9a55737ae68c622a6ca0f4aba10214b1e1017b980ee6f452b170f9bd8b6fc34 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/reply"
    iff59b88fcc5760c7b1d69ad66b57e411500f0adacd16630691a153501e3e816e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/attachments"
    i074b06f5da6a474c92d706acd42b01c93e6918719601332c98ad63318a7e2b09 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/singlevalueextendedproperties/item"
    iab9944cb4f124c1b2e56293ab54a127f6f9603e62439995a6090da24a9a598c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/attachments/item"
    iae5eab0cdfe710ae7e5d0c337835e370bfb2ad92e83f231ab2cd9828d182f7ca "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/extensions/item"
    iae7ce3b52411bb187bd3fd984b72d80c5aefb53775748edd29efe8e556d9189a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/multivalueextendedproperties/item"
    if9c84d814e93c26e519731b5dff88f7db867d03da3d4c1c376944fcb8a00fa36 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/mentions/item"
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
func (m *PostItemRequestBuilder) Attachments()(*iff59b88fcc5760c7b1d69ad66b57e411500f0adacd16630691a153501e3e816e.AttachmentsRequestBuilder) {
    return iff59b88fcc5760c7b1d69ad66b57e411500f0adacd16630691a153501e3e816e.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.conversations.item.threads.item.posts.item.attachments.item collection
func (m *PostItemRequestBuilder) AttachmentsById(id string)(*iab9944cb4f124c1b2e56293ab54a127f6f9603e62439995a6090da24a9a598c7.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return iab9944cb4f124c1b2e56293ab54a127f6f9603e62439995a6090da24a9a598c7.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPostItemRequestBuilderInternal instantiates a new PostItemRequestBuilder and sets the default values.
func NewPostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostItemRequestBuilder) {
    m := &PostItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property posts for users
func (m *PostItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property posts for users
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
// CreatePatchRequestInformation update the navigation property posts in users
func (m *PostItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property posts in users
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
// Delete delete navigation property posts for users
func (m *PostItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property posts for users
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
func (m *PostItemRequestBuilder) Extensions()(*if637eaef85087127f12c2411e01352d92aa284812ee2eabb26cb1fc7d3ff4da5.ExtensionsRequestBuilder) {
    return if637eaef85087127f12c2411e01352d92aa284812ee2eabb26cb1fc7d3ff4da5.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.conversations.item.threads.item.posts.item.extensions.item collection
func (m *PostItemRequestBuilder) ExtensionsById(id string)(*iae5eab0cdfe710ae7e5d0c337835e370bfb2ad92e83f231ab2cd9828d182f7ca.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return iae5eab0cdfe710ae7e5d0c337835e370bfb2ad92e83f231ab2cd9828d182f7ca.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *PostItemRequestBuilder) Forward()(*ic3c754f8a7f3e73343870eab0bed80d1b24d51a9db00d2e802a2fd0f767c6deb.ForwardRequestBuilder) {
    return ic3c754f8a7f3e73343870eab0bed80d1b24d51a9db00d2e802a2fd0f767c6deb.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *PostItemRequestBuilder) InReplyTo()(*i4239b8918b5c4cbf4765ccbe6679428109b6a5b2b76a1cfd452061913232d627.InReplyToRequestBuilder) {
    return i4239b8918b5c4cbf4765ccbe6679428109b6a5b2b76a1cfd452061913232d627.NewInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Mentions the mentions property
func (m *PostItemRequestBuilder) Mentions()(*id465088205bcdf8ea6b51c829a964dedce007ceb93287121a6841c6530f29ea3.MentionsRequestBuilder) {
    return id465088205bcdf8ea6b51c829a964dedce007ceb93287121a6841c6530f29ea3.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.conversations.item.threads.item.posts.item.mentions.item collection
func (m *PostItemRequestBuilder) MentionsById(id string)(*if9c84d814e93c26e519731b5dff88f7db867d03da3d4c1c376944fcb8a00fa36.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return if9c84d814e93c26e519731b5dff88f7db867d03da3d4c1c376944fcb8a00fa36.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *PostItemRequestBuilder) MultiValueExtendedProperties()(*i2296c4df4c69facd8f70733a4445a4252bdbb05b1f267b8f56766dbb1a637ee1.MultiValueExtendedPropertiesRequestBuilder) {
    return i2296c4df4c69facd8f70733a4445a4252bdbb05b1f267b8f56766dbb1a637ee1.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.conversations.item.threads.item.posts.item.multiValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iae7ce3b52411bb187bd3fd984b72d80c5aefb53775748edd29efe8e556d9189a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iae7ce3b52411bb187bd3fd984b72d80c5aefb53775748edd29efe8e556d9189a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property posts in users
func (m *PostItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property posts in users
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
func (m *PostItemRequestBuilder) Reply()(*if9a55737ae68c622a6ca0f4aba10214b1e1017b980ee6f452b170f9bd8b6fc34.ReplyRequestBuilder) {
    return if9a55737ae68c622a6ca0f4aba10214b1e1017b980ee6f452b170f9bd8b6fc34.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *PostItemRequestBuilder) SingleValueExtendedProperties()(*if7184ad3e4dbd87ee02e456222ee430d3cb4f00e04ca859c471acae8e4071a38.SingleValueExtendedPropertiesRequestBuilder) {
    return if7184ad3e4dbd87ee02e456222ee430d3cb4f00e04ca859c471acae8e4071a38.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.conversations.item.threads.item.posts.item.singleValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i074b06f5da6a474c92d706acd42b01c93e6918719601332c98ad63318a7e2b09.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i074b06f5da6a474c92d706acd42b01c93e6918719601332c98ad63318a7e2b09.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
