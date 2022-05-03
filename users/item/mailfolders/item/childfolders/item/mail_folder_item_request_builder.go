package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i22a437c72581391f7d60295ea3a325f20c96275ab58c0c076f5f8a9288d8635c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/userconfigurations"
    i2556ba62b29f22cfa1ea8d785052ceec3113d970bb34090668001a7dd6358a53 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/singlevalueextendedproperties"
    i8893b4e2134b4f9ba448f7727349679acd1d60ccabb7ba2ad781ca064c612c70 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/multivalueextendedproperties"
    i9a6b219b81c6a9c163bc68e05971dc867eec7e2647d98e9efe9e39d9f788f365 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/move"
    ib5a93e1e14111ee012b6d2eee6b7d3b4abae2b34a92bdfe1d4c02edb2eef9270 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/copy"
    ic92197b72d8c8d9f9b4fe062e365bfd8f00f25b30a83d073e7b81227222ec240 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messagerules"
    idf4564f3c29af0f214fe0b21d30ff67876698cd5f74e7549758d2689c80fb536 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages"
    i363408f76a0889b2eb68f9ea052da020bf3d6df66f1188822738c4244b8efba4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/multivalueextendedproperties/item"
    i8b2a3c272e505a39df1741a73775d5494b18cc948abe0b02d879bb555f03798f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item"
    i92322d1518d09bbe3c02da7962d86c5a3f120b8a42a9d104aba09e9015ff0d21 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/userconfigurations/item"
    id23d1f169fe2ff13104c0a108812d0de8c74f5ca8d786699efcd06fc4894d675 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/singlevalueextendedproperties/item"
    idf8eeb51bcdb15522f0082d0e6ec8b7f1f2a856d8dfaac8de2285f59cc339c0f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messagerules/item"
)

// MailFolderItemRequestBuilder provides operations to manage the childFolders property of the microsoft.graph.mailFolder entity.
type MailFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MailFolderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFolderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MailFolderItemRequestBuilderGetQueryParameters the collection of child folders in the mailFolder.
type MailFolderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MailFolderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFolderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MailFolderItemRequestBuilderGetQueryParameters
}
// MailFolderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFolderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMailFolderItemRequestBuilderInternal instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFolderItemRequestBuilder) {
    m := &MailFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMailFolderItemRequestBuilder instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMailFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Copy the copy property
func (m *MailFolderItemRequestBuilder) Copy()(*ib5a93e1e14111ee012b6d2eee6b7d3b4abae2b34a92bdfe1d4c02edb2eef9270.CopyRequestBuilder) {
    return ib5a93e1e14111ee012b6d2eee6b7d3b4abae2b34a92bdfe1d4c02edb2eef9270.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property childFolders for users
func (m *MailFolderItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property childFolders for users
func (m *MailFolderItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *MailFolderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration the collection of child folders in the mailFolder.
func (m *MailFolderItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of child folders in the mailFolder.
func (m *MailFolderItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MailFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property childFolders in users
func (m *MailFolderItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property childFolders in users
func (m *MailFolderItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, requestConfiguration *MailFolderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property childFolders for users
func (m *MailFolderItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *MailFolderItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property childFolders for users
func (m *MailFolderItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *MailFolderItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler the collection of child folders in the mailFolder.
func (m *MailFolderItemRequestBuilder) GetWithResponseHandler(requestConfiguration *MailFolderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the collection of child folders in the mailFolder.
func (m *MailFolderItemRequestBuilder) GetWithResponseHandler(requestConfiguration *MailFolderItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailFolderFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable), nil
}
// MessageRules the messageRules property
func (m *MailFolderItemRequestBuilder) MessageRules()(*ic92197b72d8c8d9f9b4fe062e365bfd8f00f25b30a83d073e7b81227222ec240.MessageRulesRequestBuilder) {
    return ic92197b72d8c8d9f9b4fe062e365bfd8f00f25b30a83d073e7b81227222ec240.NewMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRulesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.childFolders.item.messageRules.item collection
func (m *MailFolderItemRequestBuilder) MessageRulesById(id string)(*idf8eeb51bcdb15522f0082d0e6ec8b7f1f2a856d8dfaac8de2285f59cc339c0f.MessageRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRule%2Did"] = id
    }
    return idf8eeb51bcdb15522f0082d0e6ec8b7f1f2a856d8dfaac8de2285f59cc339c0f.NewMessageRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *MailFolderItemRequestBuilder) Messages()(*idf4564f3c29af0f214fe0b21d30ff67876698cd5f74e7549758d2689c80fb536.MessagesRequestBuilder) {
    return idf4564f3c29af0f214fe0b21d30ff67876698cd5f74e7549758d2689c80fb536.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.childFolders.item.messages.item collection
func (m *MailFolderItemRequestBuilder) MessagesById(id string)(*i8b2a3c272e505a39df1741a73775d5494b18cc948abe0b02d879bb555f03798f.MessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return i8b2a3c272e505a39df1741a73775d5494b18cc948abe0b02d879bb555f03798f.NewMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move the move property
func (m *MailFolderItemRequestBuilder) Move()(*i9a6b219b81c6a9c163bc68e05971dc867eec7e2647d98e9efe9e39d9f788f365.MoveRequestBuilder) {
    return i9a6b219b81c6a9c163bc68e05971dc867eec7e2647d98e9efe9e39d9f788f365.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *MailFolderItemRequestBuilder) MultiValueExtendedProperties()(*i8893b4e2134b4f9ba448f7727349679acd1d60ccabb7ba2ad781ca064c612c70.MultiValueExtendedPropertiesRequestBuilder) {
    return i8893b4e2134b4f9ba448f7727349679acd1d60ccabb7ba2ad781ca064c612c70.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.childFolders.item.multiValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i363408f76a0889b2eb68f9ea052da020bf3d6df66f1188822738c4244b8efba4.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i363408f76a0889b2eb68f9ea052da020bf3d6df66f1188822738c4244b8efba4.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property childFolders in users
func (m *MailFolderItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, requestConfiguration *MailFolderItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property childFolders in users
func (m *MailFolderItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, requestConfiguration *MailFolderItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *MailFolderItemRequestBuilder) SingleValueExtendedProperties()(*i2556ba62b29f22cfa1ea8d785052ceec3113d970bb34090668001a7dd6358a53.SingleValueExtendedPropertiesRequestBuilder) {
    return i2556ba62b29f22cfa1ea8d785052ceec3113d970bb34090668001a7dd6358a53.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.childFolders.item.singleValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*id23d1f169fe2ff13104c0a108812d0de8c74f5ca8d786699efcd06fc4894d675.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return id23d1f169fe2ff13104c0a108812d0de8c74f5ca8d786699efcd06fc4894d675.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserConfigurations the userConfigurations property
func (m *MailFolderItemRequestBuilder) UserConfigurations()(*i22a437c72581391f7d60295ea3a325f20c96275ab58c0c076f5f8a9288d8635c.UserConfigurationsRequestBuilder) {
    return i22a437c72581391f7d60295ea3a325f20c96275ab58c0c076f5f8a9288d8635c.NewUserConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.mailFolders.item.childFolders.item.userConfigurations.item collection
func (m *MailFolderItemRequestBuilder) UserConfigurationsById(id string)(*i92322d1518d09bbe3c02da7962d86c5a3f120b8a42a9d104aba09e9015ff0d21.UserConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userConfiguration%2Did"] = id
    }
    return i92322d1518d09bbe3c02da7962d86c5a3f120b8a42a9d104aba09e9015ff0d21.NewUserConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
