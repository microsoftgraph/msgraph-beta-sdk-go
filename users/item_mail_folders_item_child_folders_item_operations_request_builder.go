package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder provides operations to manage the operations property of the microsoft.graph.mailFolder entity.
type ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemChildFoldersItemOperationsRequestBuilderGetQueryParameters get operations from users
type ItemMailFoldersItemChildFoldersItemOperationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemMailFoldersItemChildFoldersItemOperationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemChildFoldersItemOperationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailFoldersItemChildFoldersItemOperationsRequestBuilderGetQueryParameters
}
// ByMailFolderOperationId provides operations to manage the operations property of the microsoft.graph.mailFolder entity.
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI
// returns a *ItemMailFoldersItemChildFoldersItemOperationsMailFolderOperationItemRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder) ByMailFolderOperationId(mailFolderOperationId string)(*ItemMailFoldersItemChildFoldersItemOperationsMailFolderOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if mailFolderOperationId != "" {
        urlTplParams["mailFolderOperation%2Did"] = mailFolderOperationId
    }
    return NewItemMailFoldersItemChildFoldersItemOperationsMailFolderOperationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemMailFoldersItemChildFoldersItemOperationsRequestBuilderInternal instantiates a new ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemOperationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder) {
    m := &ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/operations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemChildFoldersItemOperationsRequestBuilder instantiates a new ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder and sets the default values.
func NewItemMailFoldersItemChildFoldersItemOperationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemChildFoldersItemOperationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemMailFoldersItemChildFoldersItemOperationsCountRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder) Count()(*ItemMailFoldersItemChildFoldersItemOperationsCountRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemOperationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get operations from users
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI
// returns a MailFolderOperationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemOperationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderOperationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailFolderOperationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderOperationCollectionResponseable), nil
}
// ToGetRequestInformation get operations from users
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI
// returns a *RequestInformation when successful
func (m *ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemChildFoldersItemOperationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI
// returns a *ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder when successful
func (m *ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder) WithUrl(rawUrl string)(*ItemMailFoldersItemChildFoldersItemOperationsRequestBuilder) {
    return NewItemMailFoldersItemChildFoldersItemOperationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
