package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder provides operations to manage the operations property of the microsoft.graph.mailFolder entity.
type ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilderGetQueryParameters the collection of long-running operations in the mailFolder.
type ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilderGetQueryParameters
}
// NewItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilderInternal instantiates a new ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder and sets the default values.
func NewItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder) {
    m := &ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/operations/{mailFolderOperation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder instantiates a new ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder and sets the default values.
func NewItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the collection of long-running operations in the mailFolder.
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI on 2024-04-29 and will be removed 2024-06-30
// returns a MailFolderOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderOperationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMailFolderOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailFolderOperationable), nil
}
// ToGetRequestInformation the collection of long-running operations in the mailFolder.
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI on 2024-04-29 and will be removed 2024-06-30
// returns a *RequestInformation when successful
func (m *ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI on 2024-04-29 and will be removed 2024-06-30
// returns a *ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder when successful
func (m *ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder) WithUrl(rawUrl string)(*ItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder) {
    return NewItemMailFoldersItemOperationsMailFolderOperationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
