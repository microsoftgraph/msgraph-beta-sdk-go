package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder provides operations to call the getFileVaultKey method.
type ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderInternal instantiates a new ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder and sets the default values.
func NewItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) {
    m := &ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/getFileVaultKey()", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder instantiates a new ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder and sets the default values.
func NewItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getFileVaultKey
// Deprecated: This method is obsolete. Use GetAsGetFileVaultKeyGetResponse instead.
// returns a ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration)(ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseable), nil
}
// GetAsGetFileVaultKeyGetResponse invoke function getFileVaultKey
// returns a ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) GetAsGetFileVaultKeyGetResponse(ctx context.Context, requestConfiguration *ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration)(ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable), nil
}
// ToGetRequestInformation invoke function getFileVaultKey
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder when successful
func (m *ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) {
    return NewItemManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
