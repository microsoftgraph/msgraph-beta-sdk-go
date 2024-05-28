package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder provides operations to call the getFileVaultKey method.
type ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderInternal instantiates a new ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder and sets the default values.
func NewComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) {
    m := &ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/getFileVaultKey()", pathParameters),
    }
    return m
}
// NewComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder instantiates a new ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder and sets the default values.
func NewComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getFileVaultKey
// Deprecated: This method is obsolete. Use GetAsGetFileVaultKeyGetResponse instead.
// returns a ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration)(ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyResponseable), nil
}
// GetAsGetFileVaultKeyGetResponse invoke function getFileVaultKey
// returns a ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) GetAsGetFileVaultKeyGetResponse(ctx context.Context, requestConfiguration *ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration)(ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateComanageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyGetResponseable), nil
}
// ToGetRequestInformation invoke function getFileVaultKey
// returns a *RequestInformation when successful
func (m *ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder when successful
func (m *ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) WithUrl(rawUrl string)(*ComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder) {
    return NewComanageddevicesItemGetfilevaultkeyGetFileVaultKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
