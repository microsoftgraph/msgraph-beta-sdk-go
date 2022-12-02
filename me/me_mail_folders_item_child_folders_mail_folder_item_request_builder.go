package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder provides operations to manage the childFolders property of the microsoft.graph.mailFolder entity.
type MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderGetQueryParameters the collection of child folders in the mailFolder.
type MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderGetQueryParameters
}
// MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeMailFoldersItemChildFoldersMailFolderItemRequestBuilderInternal instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMeMailFoldersItemChildFoldersMailFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) {
    m := &MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeMailFoldersItemChildFoldersMailFolderItemRequestBuilder instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMeMailFoldersItemChildFoldersMailFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeMailFoldersItemChildFoldersMailFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Copy provides operations to call the copy method.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) Copy()(*MeMailFoldersItemChildFoldersItemCopyRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property childFolders for me
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of child folders in the mailFolder.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property childFolders in me
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, requestConfiguration *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property childFolders for me
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get the collection of child folders in the mailFolder.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailFolderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable), nil
}
// MessageRules provides operations to manage the messageRules property of the microsoft.graph.mailFolder entity.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) MessageRules()(*MeMailFoldersItemChildFoldersItemMessageRulesRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRulesById provides operations to manage the messageRules property of the microsoft.graph.mailFolder entity.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) MessageRulesById(id string)(*MeMailFoldersItemChildFoldersItemMessageRulesMessageRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRule%2Did"] = id
    }
    return NewMeMailFoldersItemChildFoldersItemMessageRulesMessageRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) Messages()(*MeMailFoldersItemChildFoldersItemMessagesRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) MessagesById(id string)(*MeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return NewMeMailFoldersItemChildFoldersItemMessagesMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move provides operations to call the move method.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) Move()(*MeMailFoldersItemChildFoldersItemMoveRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.mailFolder entity.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) MultiValueExtendedProperties()(*MeMailFoldersItemChildFoldersItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.mailFolder entity.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*MeMailFoldersItemChildFoldersItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMeMailFoldersItemChildFoldersItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property childFolders in me
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, requestConfiguration *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailFolderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable), nil
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.mailFolder entity.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) SingleValueExtendedProperties()(*MeMailFoldersItemChildFoldersItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.mailFolder entity.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*MeMailFoldersItemChildFoldersItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewMeMailFoldersItemChildFoldersItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserConfigurations provides operations to manage the userConfigurations property of the microsoft.graph.mailFolder entity.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) UserConfigurations()(*MeMailFoldersItemChildFoldersItemUserConfigurationsRequestBuilder) {
    return NewMeMailFoldersItemChildFoldersItemUserConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserConfigurationsById provides operations to manage the userConfigurations property of the microsoft.graph.mailFolder entity.
func (m *MeMailFoldersItemChildFoldersMailFolderItemRequestBuilder) UserConfigurationsById(id string)(*MeMailFoldersItemChildFoldersItemUserConfigurationsUserConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userConfiguration%2Did"] = id
    }
    return NewMeMailFoldersItemChildFoldersItemUserConfigurationsUserConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
