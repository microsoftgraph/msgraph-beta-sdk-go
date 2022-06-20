package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i156960fb543e0ce4bf625d42526a73c4f7a0ec9672c001acf9d49cf44b8c2ff5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/mentions"
    i248c2e72a3592bf28dfbc2f71844da6e147b655d48ef2eb3f3ea62f42d90d966 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/singlevalueextendedproperties"
    i43d5bf7d266db50b50e04a682edc7fe514bafa2465355a94a007e6320d99d617 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/inreplyto"
    i752e70d2b7fc2704d4eef3b8d3f19794b3d3f44258045bb54edd2e65892d05bd "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/reply"
    i89db43ff66f447209002f5665fd5f78ed25f827659cff39d193e2f33ae220e2e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/multivalueextendedproperties"
    i9df977ca780dc0d140675ce87069d982d3151d901c7fd2a7e5823b3934fa25f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/attachments"
    ic609b04e6ef2ba4b655b061ece719f39685820230e434a7b13efd1a16d666761 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/extensions"
    iea164510b0c521876f9bf52ef4f4f7bcfb479964d6a7a35ad5ca843d1b5f2808 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/forward"
    i0f1e494ef605c5c995db7dbaa4086c416a19342429243811dc9e2320600ae90e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/extensions/item"
    i31e115848de504a674c4e13d6beadc9b2d83e879a5fdc82f93f5ddfb706ac343 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/attachments/item"
    i4120f55900fac33f51a526e959eeabc979bb8d2ba02378a3af4bf2b6fb858a02 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/multivalueextendedproperties/item"
    i767b4645e14dda89ea8a901343958f824c9cc0eb861a5ea7515de33e1d921593 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/mentions/item"
    ie83abea890f828e4525d3a6f7a17ab8967c2b4d3e746928994f05c75ce15a262 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/threads/item/posts/item/singlevalueextendedproperties/item"
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
// PostItemRequestBuilderGetQueryParameters get posts from groups
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
func (m *PostItemRequestBuilder) Attachments()(*i9df977ca780dc0d140675ce87069d982d3151d901c7fd2a7e5823b3934fa25f4.AttachmentsRequestBuilder) {
    return i9df977ca780dc0d140675ce87069d982d3151d901c7fd2a7e5823b3934fa25f4.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.attachments.item collection
func (m *PostItemRequestBuilder) AttachmentsById(id string)(*i31e115848de504a674c4e13d6beadc9b2d83e879a5fdc82f93f5ddfb706ac343.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i31e115848de504a674c4e13d6beadc9b2d83e879a5fdc82f93f5ddfb706ac343.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPostItemRequestBuilderInternal instantiates a new PostItemRequestBuilder and sets the default values.
func NewPostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostItemRequestBuilder) {
    m := &PostItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property posts for groups
func (m *PostItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property posts for groups
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
// CreateGetRequestInformation get posts from groups
func (m *PostItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get posts from groups
func (m *PostItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PostItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property posts in groups
func (m *PostItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property posts in groups
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
// Delete delete navigation property posts for groups
func (m *PostItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property posts for groups
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
func (m *PostItemRequestBuilder) Extensions()(*ic609b04e6ef2ba4b655b061ece719f39685820230e434a7b13efd1a16d666761.ExtensionsRequestBuilder) {
    return ic609b04e6ef2ba4b655b061ece719f39685820230e434a7b13efd1a16d666761.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.extensions.item collection
func (m *PostItemRequestBuilder) ExtensionsById(id string)(*i0f1e494ef605c5c995db7dbaa4086c416a19342429243811dc9e2320600ae90e.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i0f1e494ef605c5c995db7dbaa4086c416a19342429243811dc9e2320600ae90e.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *PostItemRequestBuilder) Forward()(*iea164510b0c521876f9bf52ef4f4f7bcfb479964d6a7a35ad5ca843d1b5f2808.ForwardRequestBuilder) {
    return iea164510b0c521876f9bf52ef4f4f7bcfb479964d6a7a35ad5ca843d1b5f2808.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get posts from groups
func (m *PostItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get posts from groups
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
func (m *PostItemRequestBuilder) InReplyTo()(*i43d5bf7d266db50b50e04a682edc7fe514bafa2465355a94a007e6320d99d617.InReplyToRequestBuilder) {
    return i43d5bf7d266db50b50e04a682edc7fe514bafa2465355a94a007e6320d99d617.NewInReplyToRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Mentions the mentions property
func (m *PostItemRequestBuilder) Mentions()(*i156960fb543e0ce4bf625d42526a73c4f7a0ec9672c001acf9d49cf44b8c2ff5.MentionsRequestBuilder) {
    return i156960fb543e0ce4bf625d42526a73c4f7a0ec9672c001acf9d49cf44b8c2ff5.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.mentions.item collection
func (m *PostItemRequestBuilder) MentionsById(id string)(*i767b4645e14dda89ea8a901343958f824c9cc0eb861a5ea7515de33e1d921593.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return i767b4645e14dda89ea8a901343958f824c9cc0eb861a5ea7515de33e1d921593.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *PostItemRequestBuilder) MultiValueExtendedProperties()(*i89db43ff66f447209002f5665fd5f78ed25f827659cff39d193e2f33ae220e2e.MultiValueExtendedPropertiesRequestBuilder) {
    return i89db43ff66f447209002f5665fd5f78ed25f827659cff39d193e2f33ae220e2e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.multiValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i4120f55900fac33f51a526e959eeabc979bb8d2ba02378a3af4bf2b6fb858a02.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i4120f55900fac33f51a526e959eeabc979bb8d2ba02378a3af4bf2b6fb858a02.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property posts in groups
func (m *PostItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property posts in groups
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
func (m *PostItemRequestBuilder) Reply()(*i752e70d2b7fc2704d4eef3b8d3f19794b3d3f44258045bb54edd2e65892d05bd.ReplyRequestBuilder) {
    return i752e70d2b7fc2704d4eef3b8d3f19794b3d3f44258045bb54edd2e65892d05bd.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *PostItemRequestBuilder) SingleValueExtendedProperties()(*i248c2e72a3592bf28dfbc2f71844da6e147b655d48ef2eb3f3ea62f42d90d966.SingleValueExtendedPropertiesRequestBuilder) {
    return i248c2e72a3592bf28dfbc2f71844da6e147b655d48ef2eb3f3ea62f42d90d966.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.threads.item.posts.item.singleValueExtendedProperties.item collection
func (m *PostItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ie83abea890f828e4525d3a6f7a17ab8967c2b4d3e746928994f05c75ce15a262.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ie83abea890f828e4525d3a6f7a17ab8967c2b4d3e746928994f05c75ce15a262.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
