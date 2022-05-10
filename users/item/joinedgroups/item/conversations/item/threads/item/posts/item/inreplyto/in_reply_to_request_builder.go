package inreplyto

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0038d0a050df8dc496ed99199d3248e4ac482e7909c0b5a813cc9160eb35755e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/reply"
    i025a9e398f6eb67238b6daed67ba88591dd164cabef6f51e64321138880c6e96 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/mentions"
    i08361f23da59f12c4a8a3e7e12267160d9b1791e09526195e53537eb339dd1f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/extensions"
    i244cec99d2422d570a86c53f0167ac1972170983da28f45d93a701f68ae37715 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties"
    i277327dd1c38baf4d918cef18c4d52627374e562e6ac79f7574b5aa53c426b64 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/multivalueextendedproperties"
    i6797e8349565f62fc3c75fdf16b5e498d43987f7592ebcd4132b2073d3e7fbb9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/attachments"
    icd21ed7e44aef73ed0aeb84d2dc70b5e3a4965eeb43e5bf0e34d8c9b6e090b1c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/forward"
    i02133c49300cf3ee1e949b726a13899e600751d85280ea190816bc55f4a77ca8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/attachments/item"
    i2cae1d7c03adfa36221292895c9fade732a7fba69a421aac8db3b9f9adce7aad "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties/item"
    ia9cf9d2ad4e0ab19d6aa372a00af6171d29bcab01a60ae974c85ec2021b36390 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/mentions/item"
    ibf9d2e2b8d030462cad748f045bad6031fbbad29746787843985415a9862cf4a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/extensions/item"
    idb1e129f2ae5456ae29ac203431db753df72a761d47bbfc0821a8cc84642c4f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/conversations/item/threads/item/posts/item/inreplyto/multivalueextendedproperties/item"
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
func (m *InReplyToRequestBuilder) Attachments()(*i6797e8349565f62fc3c75fdf16b5e498d43987f7592ebcd4132b2073d3e7fbb9.AttachmentsRequestBuilder) {
    return i6797e8349565f62fc3c75fdf16b5e498d43987f7592ebcd4132b2073d3e7fbb9.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.conversations.item.threads.item.posts.item.inReplyTo.attachments.item collection
func (m *InReplyToRequestBuilder) AttachmentsById(id string)(*i02133c49300cf3ee1e949b726a13899e600751d85280ea190816bc55f4a77ca8.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i02133c49300cf3ee1e949b726a13899e600751d85280ea190816bc55f4a77ca8.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewInReplyToRequestBuilderInternal instantiates a new InReplyToRequestBuilder and sets the default values.
func NewInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InReplyToRequestBuilder) {
    m := &InReplyToRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property inReplyTo for users
func (m *InReplyToRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property inReplyTo for users
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
// CreatePatchRequestInformation update the navigation property inReplyTo in users
func (m *InReplyToRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property inReplyTo in users
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
// Delete delete navigation property inReplyTo for users
func (m *InReplyToRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property inReplyTo for users
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
func (m *InReplyToRequestBuilder) Extensions()(*i08361f23da59f12c4a8a3e7e12267160d9b1791e09526195e53537eb339dd1f9.ExtensionsRequestBuilder) {
    return i08361f23da59f12c4a8a3e7e12267160d9b1791e09526195e53537eb339dd1f9.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.conversations.item.threads.item.posts.item.inReplyTo.extensions.item collection
func (m *InReplyToRequestBuilder) ExtensionsById(id string)(*ibf9d2e2b8d030462cad748f045bad6031fbbad29746787843985415a9862cf4a.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ibf9d2e2b8d030462cad748f045bad6031fbbad29746787843985415a9862cf4a.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *InReplyToRequestBuilder) Forward()(*icd21ed7e44aef73ed0aeb84d2dc70b5e3a4965eeb43e5bf0e34d8c9b6e090b1c.ForwardRequestBuilder) {
    return icd21ed7e44aef73ed0aeb84d2dc70b5e3a4965eeb43e5bf0e34d8c9b6e090b1c.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *InReplyToRequestBuilder) Mentions()(*i025a9e398f6eb67238b6daed67ba88591dd164cabef6f51e64321138880c6e96.MentionsRequestBuilder) {
    return i025a9e398f6eb67238b6daed67ba88591dd164cabef6f51e64321138880c6e96.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.conversations.item.threads.item.posts.item.inReplyTo.mentions.item collection
func (m *InReplyToRequestBuilder) MentionsById(id string)(*ia9cf9d2ad4e0ab19d6aa372a00af6171d29bcab01a60ae974c85ec2021b36390.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return ia9cf9d2ad4e0ab19d6aa372a00af6171d29bcab01a60ae974c85ec2021b36390.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *InReplyToRequestBuilder) MultiValueExtendedProperties()(*i277327dd1c38baf4d918cef18c4d52627374e562e6ac79f7574b5aa53c426b64.MultiValueExtendedPropertiesRequestBuilder) {
    return i277327dd1c38baf4d918cef18c4d52627374e562e6ac79f7574b5aa53c426b64.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.conversations.item.threads.item.posts.item.inReplyTo.multiValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) MultiValueExtendedPropertiesById(id string)(*idb1e129f2ae5456ae29ac203431db753df72a761d47bbfc0821a8cc84642c4f1.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return idb1e129f2ae5456ae29ac203431db753df72a761d47bbfc0821a8cc84642c4f1.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property inReplyTo in users
func (m *InReplyToRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property inReplyTo in users
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
func (m *InReplyToRequestBuilder) Reply()(*i0038d0a050df8dc496ed99199d3248e4ac482e7909c0b5a813cc9160eb35755e.ReplyRequestBuilder) {
    return i0038d0a050df8dc496ed99199d3248e4ac482e7909c0b5a813cc9160eb35755e.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *InReplyToRequestBuilder) SingleValueExtendedProperties()(*i244cec99d2422d570a86c53f0167ac1972170983da28f45d93a701f68ae37715.SingleValueExtendedPropertiesRequestBuilder) {
    return i244cec99d2422d570a86c53f0167ac1972170983da28f45d93a701f68ae37715.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.conversations.item.threads.item.posts.item.inReplyTo.singleValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i2cae1d7c03adfa36221292895c9fade732a7fba69a421aac8db3b9f9adce7aad.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i2cae1d7c03adfa36221292895c9fade732a7fba69a421aac8db3b9f9adce7aad.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
