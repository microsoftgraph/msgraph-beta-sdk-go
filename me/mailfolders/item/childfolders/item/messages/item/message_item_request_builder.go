package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i18f0f3d80ebc26c14d0b298a950e35a82ff4632a691418331bc5c8ae9f7e1749 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/attachments"
    i25ba706e121e0ece5ebf98516c94795df9564057f989070249206d4cfb86fd93 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/mentions"
    i38bd3b59f483648cf474b3f78ca224a4030f1ccdc95d8b3eca57afa7733c4cdc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/createreplyall"
    i97f19227de75a13e8bce2a9f2f0fe826731ce9d14c85ded6e3178da92ad9f8e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/move"
    i98dfc01c9609a1f92d8d40b05cc6915fe1b5f42504125809b5cbadd5bac1be4a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/extensions"
    ia085204d95108fdfba3f30179e84fe9d56d2a41f0196b3faaccac5933897f4d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/multivalueextendedproperties"
    ia2c3cee1b082982c4046236c2e1484b1d12b01e7d202b5ad24986f82a79d1367 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/value"
    iaff702774cc3e61c13d7e163970e1629db453df271cbb59488849ed2a3ae0ec7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/createforward"
    ib47eb4bd511bd586fac1efd76227784d7f819cd4b070f8b6ebf528e44d879c62 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/copy"
    ic04bbf54cf6a3374e1844545b8f1fbb2c3e149b51fc1b499324271cb325f7593 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/forward"
    ic2eb00d9cef4107c8d86f41b9ef42043e0038c1234cf6e8fb334c17ec24dd276 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/createreply"
    ic57b0d8e565c4fbad1c8bb7680c7f1bebe14d763e641b8aaef0cf7ed6e29b08a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/send"
    idb7f86c95ef4ba5d0a0f2d5e20325af74620e389e5d247f6967baca5d122864a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/reply"
    ide503fc3f46ba166f51a715b9dc0b7e9a83eaf22bc09b1d3634024b436aad454 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/singlevalueextendedproperties"
    ie673455ab447de7fc89be141da7ff6bb95f9d92378eaefe3914fe3f9e387c4e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/unsubscribe"
    ie9c259a2102e27c149931d444b88050c501e3097576971fd37ec324913adb6b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/replyall"
    i12ca9076148c4a8f67d38f9a192b8c7b772df3687a3cc8863d5e5956f3e8e836 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/attachments/item"
    i402e59797825c6487d460e7a202a15eff3f3d7da3857b91b14cc5ac2341d77cb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/multivalueextendedproperties/item"
    i8fd7cb6657cd70065c95e991e7c57ff53240dfa4ec292476f12a358369c42f03 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/singlevalueextendedproperties/item"
    ic6ebbe1487ad64bdb56672d7a43c44d420745ba51711b450fa500eb99d41f83b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/extensions/item"
    ie8c7eefceaadf180d29aab9898dd5e226c4bc43a528d2ec1beaac3b40d5d3747 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item/messages/item/mentions/item"
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
func (m *MessageItemRequestBuilder) Attachments()(*i18f0f3d80ebc26c14d0b298a950e35a82ff4632a691418331bc5c8ae9f7e1749.AttachmentsRequestBuilder) {
    return i18f0f3d80ebc26c14d0b298a950e35a82ff4632a691418331bc5c8ae9f7e1749.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.childFolders.item.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*i12ca9076148c4a8f67d38f9a192b8c7b772df3687a3cc8863d5e5956f3e8e836.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return i12ca9076148c4a8f67d38f9a192b8c7b772df3687a3cc8863d5e5956f3e8e836.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}{?%24select,%24expand}";
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
func (m *MessageItemRequestBuilder) Content()(*ia2c3cee1b082982c4046236c2e1484b1d12b01e7d202b5ad24986f82a79d1367.ContentRequestBuilder) {
    return ia2c3cee1b082982c4046236c2e1484b1d12b01e7d202b5ad24986f82a79d1367.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy the copy property
func (m *MessageItemRequestBuilder) Copy()(*ib47eb4bd511bd586fac1efd76227784d7f819cd4b070f8b6ebf528e44d879c62.CopyRequestBuilder) {
    return ib47eb4bd511bd586fac1efd76227784d7f819cd4b070f8b6ebf528e44d879c62.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for me
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property messages for me
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
func (m *MessageItemRequestBuilder) CreateForward()(*iaff702774cc3e61c13d7e163970e1629db453df271cbb59488849ed2a3ae0ec7.CreateForwardRequestBuilder) {
    return iaff702774cc3e61c13d7e163970e1629db453df271cbb59488849ed2a3ae0ec7.NewCreateForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
// CreatePatchRequestInformation update the navigation property messages in me
func (m *MessageItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property messages in me
func (m *MessageItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *MessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateReply the createReply property
func (m *MessageItemRequestBuilder) CreateReply()(*ic2eb00d9cef4107c8d86f41b9ef42043e0038c1234cf6e8fb334c17ec24dd276.CreateReplyRequestBuilder) {
    return ic2eb00d9cef4107c8d86f41b9ef42043e0038c1234cf6e8fb334c17ec24dd276.NewCreateReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateReplyAll the createReplyAll property
func (m *MessageItemRequestBuilder) CreateReplyAll()(*i38bd3b59f483648cf474b3f78ca224a4030f1ccdc95d8b3eca57afa7733c4cdc.CreateReplyAllRequestBuilder) {
    return i38bd3b59f483648cf474b3f78ca224a4030f1ccdc95d8b3eca57afa7733c4cdc.NewCreateReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property messages for me
func (m *MessageItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property messages for me
func (m *MessageItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *MessageItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *MessageItemRequestBuilder) Extensions()(*i98dfc01c9609a1f92d8d40b05cc6915fe1b5f42504125809b5cbadd5bac1be4a.ExtensionsRequestBuilder) {
    return i98dfc01c9609a1f92d8d40b05cc6915fe1b5f42504125809b5cbadd5bac1be4a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.childFolders.item.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*ic6ebbe1487ad64bdb56672d7a43c44d420745ba51711b450fa500eb99d41f83b.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return ic6ebbe1487ad64bdb56672d7a43c44d420745ba51711b450fa500eb99d41f83b.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *MessageItemRequestBuilder) Forward()(*ic04bbf54cf6a3374e1844545b8f1fbb2c3e149b51fc1b499324271cb325f7593.ForwardRequestBuilder) {
    return ic04bbf54cf6a3374e1844545b8f1fbb2c3e149b51fc1b499324271cb325f7593.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the collection of messages in the mailFolder.
func (m *MessageItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *MessageItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMessageFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable), nil
}
// Mentions the mentions property
func (m *MessageItemRequestBuilder) Mentions()(*i25ba706e121e0ece5ebf98516c94795df9564057f989070249206d4cfb86fd93.MentionsRequestBuilder) {
    return i25ba706e121e0ece5ebf98516c94795df9564057f989070249206d4cfb86fd93.NewMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.childFolders.item.messages.item.mentions.item collection
func (m *MessageItemRequestBuilder) MentionsById(id string)(*ie8c7eefceaadf180d29aab9898dd5e226c4bc43a528d2ec1beaac3b40d5d3747.MentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return ie8c7eefceaadf180d29aab9898dd5e226c4bc43a528d2ec1beaac3b40d5d3747.NewMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move the move property
func (m *MessageItemRequestBuilder) Move()(*i97f19227de75a13e8bce2a9f2f0fe826731ce9d14c85ded6e3178da92ad9f8e1.MoveRequestBuilder) {
    return i97f19227de75a13e8bce2a9f2f0fe826731ce9d14c85ded6e3178da92ad9f8e1.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*ia085204d95108fdfba3f30179e84fe9d56d2a41f0196b3faaccac5933897f4d8.MultiValueExtendedPropertiesRequestBuilder) {
    return ia085204d95108fdfba3f30179e84fe9d56d2a41f0196b3faaccac5933897f4d8.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.childFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i402e59797825c6487d460e7a202a15eff3f3d7da3857b91b14cc5ac2341d77cb.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i402e59797825c6487d460e7a202a15eff3f3d7da3857b91b14cc5ac2341d77cb.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in me
func (m *MessageItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property messages in me
func (m *MessageItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Messageable, requestConfiguration *MessageItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *MessageItemRequestBuilder) Reply()(*idb7f86c95ef4ba5d0a0f2d5e20325af74620e389e5d247f6967baca5d122864a.ReplyRequestBuilder) {
    return idb7f86c95ef4ba5d0a0f2d5e20325af74620e389e5d247f6967baca5d122864a.NewReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReplyAll the replyAll property
func (m *MessageItemRequestBuilder) ReplyAll()(*ie9c259a2102e27c149931d444b88050c501e3097576971fd37ec324913adb6b7.ReplyAllRequestBuilder) {
    return ie9c259a2102e27c149931d444b88050c501e3097576971fd37ec324913adb6b7.NewReplyAllRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Send the send property
func (m *MessageItemRequestBuilder) Send()(*ic57b0d8e565c4fbad1c8bb7680c7f1bebe14d763e641b8aaef0cf7ed6e29b08a.SendRequestBuilder) {
    return ic57b0d8e565c4fbad1c8bb7680c7f1bebe14d763e641b8aaef0cf7ed6e29b08a.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*ide503fc3f46ba166f51a715b9dc0b7e9a83eaf22bc09b1d3634024b436aad454.SingleValueExtendedPropertiesRequestBuilder) {
    return ide503fc3f46ba166f51a715b9dc0b7e9a83eaf22bc09b1d3634024b436aad454.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.mailFolders.item.childFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i8fd7cb6657cd70065c95e991e7c57ff53240dfa4ec292476f12a358369c42f03.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i8fd7cb6657cd70065c95e991e7c57ff53240dfa4ec292476f12a358369c42f03.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Unsubscribe the unsubscribe property
func (m *MessageItemRequestBuilder) Unsubscribe()(*ie673455ab447de7fc89be141da7ff6bb95f9d92378eaefe3914fe3f9e387c4e3.UnsubscribeRequestBuilder) {
    return ie673455ab447de7fc89be141da7ff6bb95f9d92378eaefe3914fe3f9e387c4e3.NewUnsubscribeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
