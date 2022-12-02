package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemMailFoldersMailFolderItemRequestBuilder provides operations to manage the mailFolders property of the microsoft.graph.user entity.
type UsersItemMailFoldersMailFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemMailFoldersMailFolderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemMailFoldersMailFolderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemMailFoldersMailFolderItemRequestBuilderGetQueryParameters the user's mail folders. Read-only. Nullable.
type UsersItemMailFoldersMailFolderItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemMailFoldersMailFolderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemMailFoldersMailFolderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemMailFoldersMailFolderItemRequestBuilderGetQueryParameters
}
// UsersItemMailFoldersMailFolderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemMailFoldersMailFolderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChildFolders provides operations to manage the childFolders property of the microsoft.graph.mailFolder entity.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) ChildFolders()(*UsersItemMailFoldersItemChildFoldersRequestBuilder) {
    return NewUsersItemMailFoldersItemChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildFoldersById provides operations to manage the childFolders property of the microsoft.graph.mailFolder entity.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) ChildFoldersById(id string)(*UsersItemMailFoldersItemChildFoldersMailFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder%2Did1"] = id
    }
    return NewUsersItemMailFoldersItemChildFoldersMailFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUsersItemMailFoldersMailFolderItemRequestBuilderInternal instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewUsersItemMailFoldersMailFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemMailFoldersMailFolderItemRequestBuilder) {
    m := &UsersItemMailFoldersMailFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemMailFoldersMailFolderItemRequestBuilder instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewUsersItemMailFoldersMailFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemMailFoldersMailFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemMailFoldersMailFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Copy provides operations to call the copy method.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) Copy()(*UsersItemMailFoldersItemCopyRequestBuilder) {
    return NewUsersItemMailFoldersItemCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property mailFolders for users
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemMailFoldersMailFolderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the user's mail folders. Read-only. Nullable.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemMailFoldersMailFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property mailFolders in users
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, requestConfiguration *UsersItemMailFoldersMailFolderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property mailFolders for users
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemMailFoldersMailFolderItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the user's mail folders. Read-only. Nullable.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemMailFoldersMailFolderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, error) {
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
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) MessageRules()(*UsersItemMailFoldersItemMessageRulesRequestBuilder) {
    return NewUsersItemMailFoldersItemMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRulesById provides operations to manage the messageRules property of the microsoft.graph.mailFolder entity.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) MessageRulesById(id string)(*UsersItemMailFoldersItemMessageRulesMessageRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRule%2Did"] = id
    }
    return NewUsersItemMailFoldersItemMessageRulesMessageRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) Messages()(*UsersItemMailFoldersItemMessagesRequestBuilder) {
    return NewUsersItemMailFoldersItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) MessagesById(id string)(*UsersItemMailFoldersItemMessagesMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return NewUsersItemMailFoldersItemMessagesMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move provides operations to call the move method.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) Move()(*UsersItemMailFoldersItemMoveRequestBuilder) {
    return NewUsersItemMailFoldersItemMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.mailFolder entity.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) MultiValueExtendedProperties()(*UsersItemMailFoldersItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemMailFoldersItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.mailFolder entity.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*UsersItemMailFoldersItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemMailFoldersItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property mailFolders in users
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, requestConfiguration *UsersItemMailFoldersMailFolderItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, error) {
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
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) SingleValueExtendedProperties()(*UsersItemMailFoldersItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewUsersItemMailFoldersItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.mailFolder entity.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*UsersItemMailFoldersItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewUsersItemMailFoldersItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserConfigurations provides operations to manage the userConfigurations property of the microsoft.graph.mailFolder entity.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) UserConfigurations()(*UsersItemMailFoldersItemUserConfigurationsRequestBuilder) {
    return NewUsersItemMailFoldersItemUserConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserConfigurationsById provides operations to manage the userConfigurations property of the microsoft.graph.mailFolder entity.
func (m *UsersItemMailFoldersMailFolderItemRequestBuilder) UserConfigurationsById(id string)(*UsersItemMailFoldersItemUserConfigurationsUserConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userConfiguration%2Did"] = id
    }
    return NewUsersItemMailFoldersItemUserConfigurationsUserConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
