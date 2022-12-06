package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i17d1574c3c57286814540657170c4b7233c1c6f26ec60bea16e08bf0678e59a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/userconfigurations"
    i541fc3399419d1ae2cab4368370a7227e59f2c04b7e54e4929b86ebe04698638 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages"
    i58c44598f2855313799f5f0ec80b7d31cf45e85c1903cdc2ee9bd6c5776786db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/move"
    i5bed6e35b48c78b26b58da0c35269f9fbe5e11549479f8aafb35b126f8e25ab5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/singlevalueextendedproperties"
    i88c04acf0aa04d581c2773f019a56cb1fdf513359c2fe21e75dcc60ebc80d511 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messagerules"
    ibeb8c4168d977b698b05801a1b10798f68c64e767d766ce8595176f72fc5093b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders"
    id07076013853ba5080f44b31c440adb7eef26be6bf39605515485ec5ece17190 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/multivalueextendedproperties"
    if3fb3684743655c918bd082004c113091359b0d53e481b54379b0fd767244c8d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/copy"
    i5093bb4836672beb5dd5657e003961ad167c9ffe6b1cfc7663eb7832d4a6de57 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/userconfigurations/item"
    i5e5f56aacb3cd9d3f5d914c12525e9de5b7e2db88b3753d1b0e65ad922a94eb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/singlevalueextendedproperties/item"
    i779f11aeb25854b5bc2350fde107c6812b983ae14fd525417dabc92b85de7d0a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/multivalueextendedproperties/item"
    i8ffc8cfef5f5058547bf175ad9868544c797b8dde19d8a6459fe8494963d0a24 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messagerules/item"
    ic4853dfe79219076841dce140d8e62c905b0804ee816e62386841f353c4b1148 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item"
    iede75dd69fcd9e5b626950014b98ed33d606c522a9728b4f3a33a890e6632856 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/childfolders/item"
)

// MailFolderItemRequestBuilder provides operations to manage the mailFolders property of the microsoft.graph.user entity.
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
// MailFolderItemRequestBuilderGetQueryParameters the user's mail folders. Read-only. Nullable.
type MailFolderItemRequestBuilderGetQueryParameters struct {
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
// ChildFolders provides operations to manage the childFolders property of the microsoft.graph.mailFolder entity.
func (m *MailFolderItemRequestBuilder) ChildFolders()(*ibeb8c4168d977b698b05801a1b10798f68c64e767d766ce8595176f72fc5093b.ChildFoldersRequestBuilder) {
    return ibeb8c4168d977b698b05801a1b10798f68c64e767d766ce8595176f72fc5093b.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildFoldersById provides operations to manage the childFolders property of the microsoft.graph.mailFolder entity.
func (m *MailFolderItemRequestBuilder) ChildFoldersById(id string)(*iede75dd69fcd9e5b626950014b98ed33d606c522a9728b4f3a33a890e6632856.MailFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder%2Did1"] = id
    }
    return iede75dd69fcd9e5b626950014b98ed33d606c522a9728b4f3a33a890e6632856.NewMailFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMailFolderItemRequestBuilderInternal instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFolderItemRequestBuilder) {
    m := &MailFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/mailFolders/{mailFolder%2Did}{?%24select}";
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
// Copy provides operations to call the copy method.
func (m *MailFolderItemRequestBuilder) Copy()(*if3fb3684743655c918bd082004c113091359b0d53e481b54379b0fd767244c8d.CopyRequestBuilder) {
    return if3fb3684743655c918bd082004c113091359b0d53e481b54379b0fd767244c8d.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property mailFolders for me
func (m *MailFolderItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MailFolderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MailFolderItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MailFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property mailFolders in me
func (m *MailFolderItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, requestConfiguration *MailFolderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property mailFolders for me
func (m *MailFolderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MailFolderItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *MailFolderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MailFolderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, error) {
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
func (m *MailFolderItemRequestBuilder) MessageRules()(*i88c04acf0aa04d581c2773f019a56cb1fdf513359c2fe21e75dcc60ebc80d511.MessageRulesRequestBuilder) {
    return i88c04acf0aa04d581c2773f019a56cb1fdf513359c2fe21e75dcc60ebc80d511.NewMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRulesById provides operations to manage the messageRules property of the microsoft.graph.mailFolder entity.
func (m *MailFolderItemRequestBuilder) MessageRulesById(id string)(*i8ffc8cfef5f5058547bf175ad9868544c797b8dde19d8a6459fe8494963d0a24.MessageRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRule%2Did"] = id
    }
    return i8ffc8cfef5f5058547bf175ad9868544c797b8dde19d8a6459fe8494963d0a24.NewMessageRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
func (m *MailFolderItemRequestBuilder) Messages()(*i541fc3399419d1ae2cab4368370a7227e59f2c04b7e54e4929b86ebe04698638.MessagesRequestBuilder) {
    return i541fc3399419d1ae2cab4368370a7227e59f2c04b7e54e4929b86ebe04698638.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.mailFolder entity.
func (m *MailFolderItemRequestBuilder) MessagesById(id string)(*ic4853dfe79219076841dce140d8e62c905b0804ee816e62386841f353c4b1148.MessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return ic4853dfe79219076841dce140d8e62c905b0804ee816e62386841f353c4b1148.NewMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Move provides operations to call the move method.
func (m *MailFolderItemRequestBuilder) Move()(*i58c44598f2855313799f5f0ec80b7d31cf45e85c1903cdc2ee9bd6c5776786db.MoveRequestBuilder) {
    return i58c44598f2855313799f5f0ec80b7d31cf45e85c1903cdc2ee9bd6c5776786db.NewMoveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.mailFolder entity.
func (m *MailFolderItemRequestBuilder) MultiValueExtendedProperties()(*id07076013853ba5080f44b31c440adb7eef26be6bf39605515485ec5ece17190.MultiValueExtendedPropertiesRequestBuilder) {
    return id07076013853ba5080f44b31c440adb7eef26be6bf39605515485ec5ece17190.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.mailFolder entity.
func (m *MailFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i779f11aeb25854b5bc2350fde107c6812b983ae14fd525417dabc92b85de7d0a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i779f11aeb25854b5bc2350fde107c6812b983ae14fd525417dabc92b85de7d0a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property mailFolders in me
func (m *MailFolderItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, requestConfiguration *MailFolderItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderable, error) {
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
func (m *MailFolderItemRequestBuilder) SingleValueExtendedProperties()(*i5bed6e35b48c78b26b58da0c35269f9fbe5e11549479f8aafb35b126f8e25ab5.SingleValueExtendedPropertiesRequestBuilder) {
    return i5bed6e35b48c78b26b58da0c35269f9fbe5e11549479f8aafb35b126f8e25ab5.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.mailFolder entity.
func (m *MailFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i5e5f56aacb3cd9d3f5d914c12525e9de5b7e2db88b3753d1b0e65ad922a94eb7.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i5e5f56aacb3cd9d3f5d914c12525e9de5b7e2db88b3753d1b0e65ad922a94eb7.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserConfigurations provides operations to manage the userConfigurations property of the microsoft.graph.mailFolder entity.
func (m *MailFolderItemRequestBuilder) UserConfigurations()(*i17d1574c3c57286814540657170c4b7233c1c6f26ec60bea16e08bf0678e59a0.UserConfigurationsRequestBuilder) {
    return i17d1574c3c57286814540657170c4b7233c1c6f26ec60bea16e08bf0678e59a0.NewUserConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserConfigurationsById provides operations to manage the userConfigurations property of the microsoft.graph.mailFolder entity.
func (m *MailFolderItemRequestBuilder) UserConfigurationsById(id string)(*i5093bb4836672beb5dd5657e003961ad167c9ffe6b1cfc7663eb7832d4a6de57.UserConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userConfiguration%2Did"] = id
    }
    return i5093bb4836672beb5dd5657e003961ad167c9ffe6b1cfc7663eb7832d4a6de57.NewUserConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
