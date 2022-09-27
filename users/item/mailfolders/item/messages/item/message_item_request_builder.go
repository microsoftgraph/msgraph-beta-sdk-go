package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1388a5b7afebd5c1d7ad7238e86dfa05a923b4d8e35dc4265272c52df5b7ab94 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/move"
    i3c0595e344f6094e852e3be4813c86bfd5e14109889db8d73b847fbb96d50e80 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/createreply"
    i58c0a5d3e1252e77598ba1ea581032b7c028f21319ac81bacba8d36aaae8d7c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/forward"
    i779693a4ed700a6242a8750d06dd0c544dee12d3a8a544cfcfd6a28155e1ed26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/extensions"
    i80e9f30d67f4b0e96cf2e956e67ef07179485a391069b5616f4763663728fad5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/copy"
    i884f837848c6f0c6423949ee4d2944c09197989edfb9edbb71b2db93ccab29ba "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/multivalueextendedproperties"
    i8d70e83d3fc5a5d914de83826ade66aa9b91ec1bd829919d84e5e87c18949a9e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/singlevalueextendedproperties"
    ia41ad1f56fdb6d2218a17ab1c1042e82ae808d6e3db1a17d6944323659bb604a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/value"
    iaaf573b089683872c986a56013f03c41df00859342d210ce296d5f9be8c32fa3 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/send"
    ib48014d797abbb97d18c2fa3c3e893092e5b77db3345b6160da6647afde9d16d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/createreplyall"
    icca11b370ecbb3da8483d6ba09a0c88648c6a354801f5c04b32d16f550a2c1b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/replyall"
    id21d4bcffeb94fbdc758f4d604a94503e400e91d464e2ea5b2878b43f841f829 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/mentions"
    id9eb0bcd2e9ba6811167e83aa7fba832e7bdcccbf038736c6356977466d233b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/reply"
    idb1ae05cc149928c525695427be582f6e0764cb0ec6bd45bd0e9b145823d8389 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/createforward"
    ie25dbd4775466b7a05037c48ef048e5a9634efd19de463e40384afaa61eb7bcc "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/attachments"
    ife8a43d0f9a712fa0dd58ac70aab173d38980c2d93aa4229d9760c4c0907335b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/unsubscribe"
    i41df3c0d8cbd67958e95ec7d8e5e8caeba6577f510da216d6243755d4d47e60b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/extensions/item"
    i758243302df29cce2fe9041a95e536b614e1bce0fa181ec51735b45c28450aac "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/attachments/item"
    i8914b25cf327deb26ccf2e46fc7a562348efa5e5885764ce98f231989db617ab "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/multivalueextendedproperties/item"
    i9617bb1755b63307a9fbb614af96ed336375ccfbf08d3da6a1dd89acbde80d5b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/singlevalueextendedproperties/item"
    idb97761d489a095537daad3218e84b76808b5d8fc51cba00b233be4c0e0f9e9f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/messages/item/mentions/item"
)

// MessageItemRequestBuilder provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
type MessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MessageItemRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type MessageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MessageItemRequestBuilderGetQueryParameters
}
// MessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments the attachments property
func (m *MessageItemRequestBuilder) Attachments()(*ie25dbd4775466b7a05037c48ef048e5a9634efd19de463e40384afaa61eb7bcc.AttachmentsRequestBuilder) {
    return ie25dbd4775466b7a05037c48ef048e5a9634efd19de463e40384afaa61eb7bcc.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*i758243302df29cce2fe9041a95e536b614e1bce0fa181ec51735b45c28450aac.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i758243302df29cce2fe9041a95e536b614e1bce0fa181ec51735b45c28450aac.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the Content property
func (m *MessageItemRequestBuilder) Content()(*ia41ad1f56fdb6d2218a17ab1c1042e82ae808d6e3db1a17d6944323659bb604a.ContentRequestBuilder) {
    return ia41ad1f56fdb6d2218a17ab1c1042e82ae808d6e3db1a17d6944323659bb604a.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *MessageItemRequestBuilder) Copy()(*i80e9f30d67f4b0e96cf2e956e67ef07179485a391069b5616f4763663728fad5.CopyRequestBuilder) {
    return i80e9f30d67f4b0e96cf2e956e67ef07179485a391069b5616f4763663728fad5.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for users
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property messages for users
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *MessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateForward the createForward property
func (m *MessageItemRequestBuilder) CreateForward()(*idb1ae05cc149928c525695427be582f6e0764cb0ec6bd45bd0e9b145823d8389.CreateForwardRequestBuilder) {
    return idb1ae05cc149928c525695427be582f6e0764cb0ec6bd45bd0e9b145823d8389.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property messages in users
func (m *MessageItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property messages in users
func (m *MessageItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *MessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateReply the createReply property
func (m *MessageItemRequestBuilder) CreateReply()(*i3c0595e344f6094e852e3be4813c86bfd5e14109889db8d73b847fbb96d50e80.CreateReplyRequestBuilder) {
    return i3c0595e344f6094e852e3be4813c86bfd5e14109889db8d73b847fbb96d50e80.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateReplyAll the createReplyAll property
func (m *MessageItemRequestBuilder) CreateReplyAll()(*ib48014d797abbb97d18c2fa3c3e893092e5b77db3345b6160da6647afde9d16d.CreateReplyAllRequestBuilder) {
    return ib48014d797abbb97d18c2fa3c3e893092e5b77db3345b6160da6647afde9d16d.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property messages for users
func (m *MessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MessageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions the extensions property
func (m *MessageItemRequestBuilder) Extensions()(*i779693a4ed700a6242a8750d06dd0c544dee12d3a8a544cfcfd6a28155e1ed26.ExtensionsRequestBuilder) {
    return i779693a4ed700a6242a8750d06dd0c544dee12d3a8a544cfcfd6a28155e1ed26.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*i41df3c0d8cbd67958e95ec7d8e5e8caeba6577f510da216d6243755d4d47e60b.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i41df3c0d8cbd67958e95ec7d8e5e8caeba6577f510da216d6243755d4d47e60b.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *MessageItemRequestBuilder) Forward()(*i58c0a5d3e1252e77598ba1ea581032b7c028f21319ac81bacba8d36aaae8d7c9.ForwardRequestBuilder) {
    return i58c0a5d3e1252e77598ba1ea581032b7c028f21319ac81bacba8d36aaae8d7c9.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MessageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable), nil
}
// Mentions the mentions property
func (m *MessageItemRequestBuilder) Mentions()(*id21d4bcffeb94fbdc758f4d604a94503e400e91d464e2ea5b2878b43f841f829.MentionsRequestBuilder) {
    return id21d4bcffeb94fbdc758f4d604a94503e400e91d464e2ea5b2878b43f841f829.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.messages.item.mentions.item collection
func (m *MessageItemRequestBuilder) MentionsById(id string)(*idb97761d489a095537daad3218e84b76808b5d8fc51cba00b233be4c0e0f9e9f.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return idb97761d489a095537daad3218e84b76808b5d8fc51cba00b233be4c0e0f9e9f.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move the move property
func (m *MessageItemRequestBuilder) Move()(*i1388a5b7afebd5c1d7ad7238e86dfa05a923b4d8e35dc4265272c52df5b7ab94.MoveRequestBuilder) {
    return i1388a5b7afebd5c1d7ad7238e86dfa05a923b4d8e35dc4265272c52df5b7ab94.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*i884f837848c6f0c6423949ee4d2944c09197989edfb9edbb71b2db93ccab29ba.MultiValueExtendedPropertiesRequestBuilder) {
    return i884f837848c6f0c6423949ee4d2944c09197989edfb9edbb71b2db93ccab29ba.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i8914b25cf327deb26ccf2e46fc7a562348efa5e5885764ce98f231989db617ab.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i8914b25cf327deb26ccf2e46fc7a562348efa5e5885764ce98f231989db617ab.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in users
func (m *MessageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *MessageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable), nil
}
// Reply the reply property
func (m *MessageItemRequestBuilder) Reply()(*id9eb0bcd2e9ba6811167e83aa7fba832e7bdcccbf038736c6356977466d233b9.ReplyRequestBuilder) {
    return id9eb0bcd2e9ba6811167e83aa7fba832e7bdcccbf038736c6356977466d233b9.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReplyAll the replyAll property
func (m *MessageItemRequestBuilder) ReplyAll()(*icca11b370ecbb3da8483d6ba09a0c88648c6a354801f5c04b32d16f550a2c1b9.ReplyAllRequestBuilder) {
    return icca11b370ecbb3da8483d6ba09a0c88648c6a354801f5c04b32d16f550a2c1b9.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Send the send property
func (m *MessageItemRequestBuilder) Send()(*iaaf573b089683872c986a56013f03c41df00859342d210ce296d5f9be8c32fa3.SendRequestBuilder) {
    return iaaf573b089683872c986a56013f03c41df00859342d210ce296d5f9be8c32fa3.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*i8d70e83d3fc5a5d914de83826ade66aa9b91ec1bd829919d84e5e87c18949a9e.SingleValueExtendedPropertiesRequestBuilder) {
    return i8d70e83d3fc5a5d914de83826ade66aa9b91ec1bd829919d84e5e87c18949a9e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i9617bb1755b63307a9fbb614af96ed336375ccfbf08d3da6a1dd89acbde80d5b.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i9617bb1755b63307a9fbb614af96ed336375ccfbf08d3da6a1dd89acbde80d5b.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unsubscribe the unsubscribe property
func (m *MessageItemRequestBuilder) Unsubscribe()(*ife8a43d0f9a712fa0dd58ac70aab173d38980c2d93aa4229d9760c4c0907335b.UnsubscribeRequestBuilder) {
    return ife8a43d0f9a712fa0dd58ac70aab173d38980c2d93aa4229d9760c4c0907335b.NewUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
