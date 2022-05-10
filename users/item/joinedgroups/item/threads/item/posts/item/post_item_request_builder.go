package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1b634c5b41dcf85f5e9fcab387f5a37ea92d21af1c9f33cf3dafc63aa731a6ad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/multivalueextendedproperties"
    i32f96cc8db998666c81311dd97e5127bc81f436ac6819513dbdfc27a26e36ec3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/forward"
    i56197a7dcc96e7790418b7c8081f8e792eb80b29a5b5a73ed8aeb66d9766392a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/extensions"
    ia2e594660755771e9c93cf29b8142035aeb85f34af1a701a36c27145add4cb09 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto"
    ib63e4451217aab116f052baeeec7c39f30fc5720147aca3123de4804a7c64390 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/reply"
    iccb34ac87973b1bd02ebf44edd7ec665146fd4cf6e8ed1010f2a14aa6f578108 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/singlevalueextendedproperties"
    iebf2c6866b766ea1b5b32e8efed2e6c3355078ba88b502a675619693c4e04bd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/attachments"
    if08cd96be736958cc06a8f0357a162b2cf88586b9ff09cf82d9ec4190dc669a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/mentions"
    i24fd336a65deb229e8cbe5d0e9c425e44a3e006f26dcce24a90b2a7a900d4d8c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/extensions/item"
    i7c13459049fc760596ee7cd1eaeff38d1eb2802e048cbc3a4d78611084d0f460 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/multivalueextendedproperties/item"
    i8c8e082540c8a30b0086e9ceb05de81c809c19eba5ebbd06aefc57ad1f67ebd2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/singlevalueextendedproperties/item"
    i8ca54a365f9a798793ce65f16e69bbb361017f829af8111d8f53bd8c1a9394fb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/attachments/item"
    ic93232bda06677f14f7f216964da04f65245a40062ef5802a10da92c2951751d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/mentions/item"
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
func (m *PostItemRequestBuilder) Attachments()(*iebf2c6866b766ea1b5b32e8efed2e6c3355078ba88b502a675619693c4e04bd3.AttachmentsRequestBuilder) {
    return iebf2c6866b766ea1b5b32e8efed2e6c3355078ba88b502a675619693c4e04bd3.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.threads.item.posts.item.attachments.item collection
func (m *PostItemRequestBuilder) AttachmentsById(id string)(*i8ca54a365f9a798793ce65f16e69bbb361017f829af8111d8f53bd8c1a9394fb.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i8ca54a365f9a798793ce65f16e69bbb361017f829af8111d8f53bd8c1a9394fb.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPostItemRequestBuilderInternal instantiates a new PostItemRequestBuilder and sets the default values.
func NewPostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostItemRequestBuilder) {
    m := &PostItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}{?%24select,%24expand}";
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
func (m *PostItemRequestBuilder) Extensions()(*i56197a7dcc96e7790418b7c8081f8e792eb80b29a5b5a73ed8aeb66d9766392a.ExtensionsRequestBuilder) {
    return i56197a7dcc96e7790418b7c8081f8e792eb80b29a5b5a73ed8aeb66d9766392a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.threads.item.posts.item.extensions.item collection
func (m *PostItemRequestBuilder) ExtensionsById(id string)(*i24fd336a65deb229e8cbe5d0e9c425e44a3e006f26dcce24a90b2a7a900d4d8c.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i24fd336a65deb229e8cbe5d0e9c425e44a3e006f26dcce24a90b2a7a900d4d8c.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *PostItemRequestBuilder) Forward()(*i32f96cc8db998666c81311dd97e5127bc81f436ac6819513dbdfc27a26e36ec3.ForwardRequestBuilder) {
    return i32f96cc8db998666c81311dd97e5127bc81f436ac6819513dbdfc27a26e36ec3.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *PostItemRequestBuilder) InReplyTo()(*ia2e594660755771e9c93cf29b8142035aeb85f34af1a701a36c27145add4cb09.InReplyToRequestBuilder) {
    return ia2e594660755771e9c93cf29b8142035aeb85f34af1a701a36c27145add4cb09.NewInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Mentions the mentions property
func (m *PostItemRequestBuilder) Mentions()(*if08cd96be736958cc06a8f0357a162b2cf88586b9ff09cf82d9ec4190dc669a8.MentionsRequestBuilder) {
    return if08cd96be736958cc06a8f0357a162b2cf88586b9ff09cf82d9ec4190dc669a8.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.threads.item.posts.item.mentions.item collection
func (m *PostItemRequestBuilder) MentionsById(id string)(*ic93232bda06677f14f7f216964da04f65245a40062ef5802a10da92c2951751d.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return ic93232bda06677f14f7f216964da04f65245a40062ef5802a10da92c2951751d.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *PostItemRequestBuilder) MultiValueExtendedProperties()(*i1b634c5b41dcf85f5e9fcab387f5a37ea92d21af1c9f33cf3dafc63aa731a6ad.MultiValueExtendedPropertiesRequestBuilder) {
    return i1b634c5b41dcf85f5e9fcab387f5a37ea92d21af1c9f33cf3dafc63aa731a6ad.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.threads.item.posts.item.multiValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i7c13459049fc760596ee7cd1eaeff38d1eb2802e048cbc3a4d78611084d0f460.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i7c13459049fc760596ee7cd1eaeff38d1eb2802e048cbc3a4d78611084d0f460.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *PostItemRequestBuilder) Reply()(*ib63e4451217aab116f052baeeec7c39f30fc5720147aca3123de4804a7c64390.ReplyRequestBuilder) {
    return ib63e4451217aab116f052baeeec7c39f30fc5720147aca3123de4804a7c64390.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *PostItemRequestBuilder) SingleValueExtendedProperties()(*iccb34ac87973b1bd02ebf44edd7ec665146fd4cf6e8ed1010f2a14aa6f578108.SingleValueExtendedPropertiesRequestBuilder) {
    return iccb34ac87973b1bd02ebf44edd7ec665146fd4cf6e8ed1010f2a14aa6f578108.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.threads.item.posts.item.singleValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i8c8e082540c8a30b0086e9ceb05de81c809c19eba5ebbd06aefc57ad1f67ebd2.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i8c8e082540c8a30b0086e9ceb05de81c809c19eba5ebbd06aefc57ad1f67ebd2.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
