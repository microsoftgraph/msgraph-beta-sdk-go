package inreplyto

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i092660cb4daa3bade3280c0067d7f0b23d9c6d7e8158cb5e15a25cea50946748 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item/posts/item/inreplyto/attachments"
    i2e51a3dd8c4b451920117227064045d71545429f6402b2ff7900727b463a5c30 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item/posts/item/inreplyto/extensions"
    i53678a70e2c4b223ac2e3e392b5c40033b1da53839d0427e86f7d585d0ebfb55 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item/posts/item/inreplyto/reply"
    i76418a6a1e054c6ebbc2ed59e9f2b05bead0c4da7818797ff898e1baf8263a0e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties"
    ia364bdcd2a6e57d976ceab30a22121eeb96285be413950638489b84d9da2f835 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item/posts/item/inreplyto/forward"
    iea86c71442c5060f5244a974448a5eacccdf9ec509aa2175f51f9b2083cc1850 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item/posts/item/inreplyto/mentions"
    ifa6585f9afa943bdb7f997bd1a203350a67b1ac396bf44ab1ca366815ab5a36a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item/posts/item/inreplyto/multivalueextendedproperties"
    i4588ed803157be90c33174b17e4448bc4ea9f583c22efbeb0fbadc4ff92b3e11 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item/posts/item/inreplyto/singlevalueextendedproperties/item"
    i5681248c974836e23df7548d912f16baa7d72d089dea83c7841df40c310705e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item/posts/item/inreplyto/attachments/item"
    i94e555578ecfe650cfe6231c326cfc2a65359cad6cb1cea24403cce8da23247e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item/posts/item/inreplyto/multivalueextendedproperties/item"
    i9ac963d7e18d560692ba9b8b639d83a1a7f7a362a417fc5f4bacae515656c3a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item/posts/item/inreplyto/extensions/item"
    ie8745c8454c4bd85999f77b11332a060a6163d987c8a0d80b8287acc9d61df40 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/threads/item/posts/item/inreplyto/mentions/item"
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
func (m *InReplyToRequestBuilder) Attachments()(*i092660cb4daa3bade3280c0067d7f0b23d9c6d7e8158cb5e15a25cea50946748.AttachmentsRequestBuilder) {
    return i092660cb4daa3bade3280c0067d7f0b23d9c6d7e8158cb5e15a25cea50946748.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.threads.item.posts.item.inReplyTo.attachments.item collection
func (m *InReplyToRequestBuilder) AttachmentsById(id string)(*i5681248c974836e23df7548d912f16baa7d72d089dea83c7841df40c310705e9.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i5681248c974836e23df7548d912f16baa7d72d089dea83c7841df40c310705e9.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewInReplyToRequestBuilderInternal instantiates a new InReplyToRequestBuilder and sets the default values.
func NewInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InReplyToRequestBuilder) {
    m := &InReplyToRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo{?%24select,%24expand}";
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
func (m *InReplyToRequestBuilder) Extensions()(*i2e51a3dd8c4b451920117227064045d71545429f6402b2ff7900727b463a5c30.ExtensionsRequestBuilder) {
    return i2e51a3dd8c4b451920117227064045d71545429f6402b2ff7900727b463a5c30.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.threads.item.posts.item.inReplyTo.extensions.item collection
func (m *InReplyToRequestBuilder) ExtensionsById(id string)(*i9ac963d7e18d560692ba9b8b639d83a1a7f7a362a417fc5f4bacae515656c3a5.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i9ac963d7e18d560692ba9b8b639d83a1a7f7a362a417fc5f4bacae515656c3a5.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *InReplyToRequestBuilder) Forward()(*ia364bdcd2a6e57d976ceab30a22121eeb96285be413950638489b84d9da2f835.ForwardRequestBuilder) {
    return ia364bdcd2a6e57d976ceab30a22121eeb96285be413950638489b84d9da2f835.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *InReplyToRequestBuilder) Mentions()(*iea86c71442c5060f5244a974448a5eacccdf9ec509aa2175f51f9b2083cc1850.MentionsRequestBuilder) {
    return iea86c71442c5060f5244a974448a5eacccdf9ec509aa2175f51f9b2083cc1850.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.threads.item.posts.item.inReplyTo.mentions.item collection
func (m *InReplyToRequestBuilder) MentionsById(id string)(*ie8745c8454c4bd85999f77b11332a060a6163d987c8a0d80b8287acc9d61df40.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return ie8745c8454c4bd85999f77b11332a060a6163d987c8a0d80b8287acc9d61df40.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *InReplyToRequestBuilder) MultiValueExtendedProperties()(*ifa6585f9afa943bdb7f997bd1a203350a67b1ac396bf44ab1ca366815ab5a36a.MultiValueExtendedPropertiesRequestBuilder) {
    return ifa6585f9afa943bdb7f997bd1a203350a67b1ac396bf44ab1ca366815ab5a36a.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.threads.item.posts.item.inReplyTo.multiValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i94e555578ecfe650cfe6231c326cfc2a65359cad6cb1cea24403cce8da23247e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i94e555578ecfe650cfe6231c326cfc2a65359cad6cb1cea24403cce8da23247e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *InReplyToRequestBuilder) Reply()(*i53678a70e2c4b223ac2e3e392b5c40033b1da53839d0427e86f7d585d0ebfb55.ReplyRequestBuilder) {
    return i53678a70e2c4b223ac2e3e392b5c40033b1da53839d0427e86f7d585d0ebfb55.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *InReplyToRequestBuilder) SingleValueExtendedProperties()(*i76418a6a1e054c6ebbc2ed59e9f2b05bead0c4da7818797ff898e1baf8263a0e.SingleValueExtendedPropertiesRequestBuilder) {
    return i76418a6a1e054c6ebbc2ed59e9f2b05bead0c4da7818797ff898e1baf8263a0e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.threads.item.posts.item.inReplyTo.singleValueExtendedProperties.item collection
func (m *InReplyToRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4588ed803157be90c33174b17e4448bc4ea9f583c22efbeb0fbadc4ff92b3e11.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i4588ed803157be90c33174b17e4448bc4ea9f583c22efbeb0fbadc4ff92b3e11.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
