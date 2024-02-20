package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemGetFileVaultKeyRequestBuilder provides operations to call the getFileVaultKey method.
type ItemManagedDevicesItemGetFileVaultKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesItemGetFileVaultKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemGetFileVaultKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesItemGetFileVaultKeyRequestBuilderInternal instantiates a new ItemManagedDevicesItemGetFileVaultKeyRequestBuilder and sets the default values.
func NewItemManagedDevicesItemGetFileVaultKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemGetFileVaultKeyRequestBuilder) {
    m := &ItemManagedDevicesItemGetFileVaultKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/getFileVaultKey()", pathParameters),
    }
    return m
}
// NewItemManagedDevicesItemGetFileVaultKeyRequestBuilder instantiates a new ItemManagedDevicesItemGetFileVaultKeyRequestBuilder and sets the default values.
func NewItemManagedDevicesItemGetFileVaultKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemGetFileVaultKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemGetFileVaultKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getFileVaultKey
// Deprecated: This method is obsolete. Use GetAsGetFileVaultKeyGetResponse instead.
// returns a ItemManagedDevicesItemGetFileVaultKeyResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesItemGetFileVaultKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManagedDevicesItemGetFileVaultKeyRequestBuilderGetRequestConfiguration)(ItemManagedDevicesItemGetFileVaultKeyResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManagedDevicesItemGetFileVaultKeyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManagedDevicesItemGetFileVaultKeyResponseable), nil
}
// GetAsGetFileVaultKeyGetResponse invoke function getFileVaultKey
// returns a ItemManagedDevicesItemGetFileVaultKeyGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesItemGetFileVaultKeyRequestBuilder) GetAsGetFileVaultKeyGetResponse(ctx context.Context, requestConfiguration *ItemManagedDevicesItemGetFileVaultKeyRequestBuilderGetRequestConfiguration)(ItemManagedDevicesItemGetFileVaultKeyGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManagedDevicesItemGetFileVaultKeyGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManagedDevicesItemGetFileVaultKeyGetResponseable), nil
}
// ToGetRequestInformation invoke function getFileVaultKey
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesItemGetFileVaultKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesItemGetFileVaultKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManagedDevicesItemGetFileVaultKeyRequestBuilder when successful
func (m *ItemManagedDevicesItemGetFileVaultKeyRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesItemGetFileVaultKeyRequestBuilder) {
    return NewItemManagedDevicesItemGetFileVaultKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
