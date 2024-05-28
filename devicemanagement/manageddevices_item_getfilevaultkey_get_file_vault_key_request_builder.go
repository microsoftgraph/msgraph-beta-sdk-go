package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder provides operations to call the getFileVaultKey method.
type ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderInternal instantiates a new ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder and sets the default values.
func NewManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) {
    m := &ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/getFileVaultKey()", pathParameters),
    }
    return m
}
// NewManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder instantiates a new ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder and sets the default values.
func NewManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getFileVaultKey
// Deprecated: This method is obsolete. Use GetAsGetFileVaultKeyGetResponse instead.
// returns a ManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration)(ManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseable), nil
}
// GetAsGetFileVaultKeyGetResponse invoke function getFileVaultKey
// returns a ManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) GetAsGetFileVaultKeyGetResponse(ctx context.Context, requestConfiguration *ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration)(ManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable), nil
}
// ToGetRequestInformation invoke function getFileVaultKey
// returns a *RequestInformation when successful
func (m *ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder when successful
func (m *ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) {
    return NewManageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
