package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInformationProtectionBitlockerRecoveryKeysRequestBuilder provides operations to manage the recoveryKeys property of the microsoft.graph.bitlocker entity.
type ItemInformationProtectionBitlockerRecoveryKeysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetQueryParameters get a list of the bitlockerRecoveryKey objects and their properties.  This operation does not return the key property. For information about how to read the key property, see Get bitlockerRecoveryKey.
type ItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetQueryParameters struct {
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
// ItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetQueryParameters
}
// ByBitlockerRecoveryKeyId provides operations to manage the recoveryKeys property of the microsoft.graph.bitlocker entity.
func (m *ItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) ByBitlockerRecoveryKeyId(bitlockerRecoveryKeyId string)(*ItemInformationProtectionBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if bitlockerRecoveryKeyId != "" {
        urlTplParams["bitlockerRecoveryKey%2Did"] = bitlockerRecoveryKeyId
    }
    return NewItemInformationProtectionBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInformationProtectionBitlockerRecoveryKeysRequestBuilderInternal instantiates a new RecoveryKeysRequestBuilder and sets the default values.
func NewItemInformationProtectionBitlockerRecoveryKeysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) {
    m := &ItemInformationProtectionBitlockerRecoveryKeysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/informationProtection/bitlocker/recoveryKeys{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemInformationProtectionBitlockerRecoveryKeysRequestBuilder instantiates a new RecoveryKeysRequestBuilder and sets the default values.
func NewItemInformationProtectionBitlockerRecoveryKeysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationProtectionBitlockerRecoveryKeysRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) Count()(*ItemInformationProtectionBitlockerRecoveryKeysCountRequestBuilder) {
    return NewItemInformationProtectionBitlockerRecoveryKeysCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the bitlockerRecoveryKey objects and their properties.  This operation does not return the key property. For information about how to read the key property, see Get bitlockerRecoveryKey.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/bitlocker-list-recoverykeys?view=graph-rest-1.0
func (m *ItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BitlockerRecoveryKeyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBitlockerRecoveryKeyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BitlockerRecoveryKeyCollectionResponseable), nil
}
// ToGetRequestInformation get a list of the bitlockerRecoveryKey objects and their properties.  This operation does not return the key property. For information about how to read the key property, see Get bitlockerRecoveryKey.
func (m *ItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInformationProtectionBitlockerRecoveryKeysRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) WithUrl(rawUrl string)(*ItemInformationProtectionBitlockerRecoveryKeysRequestBuilder) {
    return NewItemInformationProtectionBitlockerRecoveryKeysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
