package inreplyto

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i11884222b7f1721a0d5988a838912b17201f50c5f1578f9dbce60e170f1d2812 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto/forward"
    i7cd9e5f5887bdaca430afa6827a0512f0288b09806388629df313bd0c5948913 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto/reply"
    i7f2fb942d7a70c1d3f8ad0b9ff556a0ebaf3314d9888b63dc13a032a08896628 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto/mentions"
    ia0a022af5c1c027e2107030e0582bae953c652e119a4a42fd6a5f96b2bff61a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto/attachments"
    iad11e364ddc2fcadea49885e039910c0507ba4c63f13284d968ca535552490c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties"
    ib5183fa5bf23e8de28236ae8310da9ef29a0d67408eed770a63259ee3ff279e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto/extensions"
    id5b297d21d647790eb37f6df323d0a08d082898cd53d01b9c10be20ccd937ef6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto/multivalueextendedproperties"
    i2e96bfa51c0e00e14c2a0fac2db331f71a20c1358f99240cc31ee2503baeb8d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto/attachments/item"
    i8069df313b8061e055df51ec24d5f27727cdc442133a8914e5196a21de231529 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto/extensions/item"
    id2a38be2afe235b613df9f86b65b8fbac2034c72309115a5d61487581ed9785c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto/multivalueextendedproperties/item"
    idfc285b8d8dcfd393f3d6cc3edb04fdb5ce39dab739b19a73805f986ccbbf03f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto/mentions/item"
    ie24445f7665aca660ce0d4f174e141b09fbe4cc44150588c0af44a19c4ef162d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties/item"
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
func (m *InReplyToRequestBuilder) Attachments()(*ia0a022af5c1c027e2107030e0582bae953c652e119a4a42fd6a5f96b2bff61a1.AttachmentsRequestBuilder) {
    return ia0a022af5c1c027e2107030e0582bae953c652e119a4a42fd6a5f96b2bff61a1.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.threads.item.posts.item.inReplyTo.attachments.item collection
func (m *InReplyToRequestBuilder) AttachmentsById(id string)(*i2e96bfa51c0e00e14c2a0fac2db331f71a20c1358f99240cc31ee2503baeb8d8.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i2e96bfa51c0e00e14c2a0fac2db331f71a20c1358f99240cc31ee2503baeb8d8.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewInReplyToRequestBuilderInternal instantiates a new InReplyToRequestBuilder and sets the default values.
func NewInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InReplyToRequestBuilder) {
    m := &InReplyToRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups/{group%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo{?%24select,%24expand}";
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
func (m *InReplyToRequestBuilder) Extensions()(*ib5183fa5bf23e8de28236ae8310da9ef29a0d67408eed770a63259ee3ff279e0.ExtensionsRequestBuilder) {
    return ib5183fa5bf23e8de28236ae8310da9ef29a0d67408eed770a63259ee3ff279e0.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.threads.item.posts.item.inReplyTo.extensions.item collection
func (m *InReplyToRequestBuilder) ExtensionsById(id string)(*i8069df313b8061e055df51ec24d5f27727cdc442133a8914e5196a21de231529.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i8069df313b8061e055df51ec24d5f27727cdc442133a8914e5196a21de231529.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *InReplyToRequestBuilder) Forward()(*i11884222b7f1721a0d5988a838912b17201f50c5f1578f9dbce60e170f1d2812.ForwardRequestBuilder) {
    return i11884222b7f1721a0d5988a838912b17201f50c5f1578f9dbce60e170f1d2812.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *InReplyToRequestBuilder) Mentions()(*i7f2fb942d7a70c1d3f8ad0b9ff556a0ebaf3314d9888b63dc13a032a08896628.MentionsRequestBuilder) {
    return i7f2fb942d7a70c1d3f8ad0b9ff556a0ebaf3314d9888b63dc13a032a08896628.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.threads.item.posts.item.inReplyTo.mentions.item collection
func (m *InReplyToRequestBuilder) MentionsById(id string)(*idfc285b8d8dcfd393f3d6cc3edb04fdb5ce39dab739b19a73805f986ccbbf03f.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return idfc285b8d8dcfd393f3d6cc3edb04fdb5ce39dab739b19a73805f986ccbbf03f.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *InReplyToRequestBuilder) MultiValueExtendedProperties()(*id5b297d21d647790eb37f6df323d0a08d082898cd53d01b9c10be20ccd937ef6.MultiValueExtendedPropertiesRequestBuilder) {
    return id5b297d21d647790eb37f6df323d0a08d082898cd53d01b9c10be20ccd937ef6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.threads.item.posts.item.inReplyTo.multiValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) MultiValueExtendedPropertiesById(id string)(*id2a38be2afe235b613df9f86b65b8fbac2034c72309115a5d61487581ed9785c.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return id2a38be2afe235b613df9f86b65b8fbac2034c72309115a5d61487581ed9785c.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *InReplyToRequestBuilder) Reply()(*i7cd9e5f5887bdaca430afa6827a0512f0288b09806388629df313bd0c5948913.ReplyRequestBuilder) {
    return i7cd9e5f5887bdaca430afa6827a0512f0288b09806388629df313bd0c5948913.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *InReplyToRequestBuilder) SingleValueExtendedProperties()(*iad11e364ddc2fcadea49885e039910c0507ba4c63f13284d968ca535552490c9.SingleValueExtendedPropertiesRequestBuilder) {
    return iad11e364ddc2fcadea49885e039910c0507ba4c63f13284d968ca535552490c9.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.joinedGroups.item.threads.item.posts.item.inReplyTo.singleValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ie24445f7665aca660ce0d4f174e141b09fbe4cc44150588c0af44a19c4ef162d.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return ie24445f7665aca660ce0d4f174e141b09fbe4cc44150588c0af44a19c4ef162d.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
