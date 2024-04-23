package informationprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// BitlockerRecoveryKeysRequestBuilder provides operations to manage the recoveryKeys property of the microsoft.graph.bitlocker entity.
type BitlockerRecoveryKeysRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BitlockerRecoveryKeysRequestBuilderGetQueryParameters the recovery keys associated with the bitlocker entity.
type BitlockerRecoveryKeysRequestBuilderGetQueryParameters struct {
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
// BitlockerRecoveryKeysRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BitlockerRecoveryKeysRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BitlockerRecoveryKeysRequestBuilderGetQueryParameters
}
// ByBitlockerRecoveryKeyId provides operations to manage the recoveryKeys property of the microsoft.graph.bitlocker entity.
// returns a *BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder when successful
func (m *BitlockerRecoveryKeysRequestBuilder) ByBitlockerRecoveryKeyId(bitlockerRecoveryKeyId string)(*BitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if bitlockerRecoveryKeyId != "" {
        urlTplParams["bitlockerRecoveryKey%2Did"] = bitlockerRecoveryKeyId
    }
    return NewBitlockerRecoveryKeysBitlockerRecoveryKeyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewBitlockerRecoveryKeysRequestBuilderInternal instantiates a new BitlockerRecoveryKeysRequestBuilder and sets the default values.
func NewBitlockerRecoveryKeysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRecoveryKeysRequestBuilder) {
    m := &BitlockerRecoveryKeysRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/informationProtection/bitlocker/recoveryKeys{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewBitlockerRecoveryKeysRequestBuilder instantiates a new BitlockerRecoveryKeysRequestBuilder and sets the default values.
func NewBitlockerRecoveryKeysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BitlockerRecoveryKeysRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBitlockerRecoveryKeysRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *BitlockerRecoveryKeysCountRequestBuilder when successful
func (m *BitlockerRecoveryKeysRequestBuilder) Count()(*BitlockerRecoveryKeysCountRequestBuilder) {
    return NewBitlockerRecoveryKeysCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the recovery keys associated with the bitlocker entity.
// returns a BitlockerRecoveryKeyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *BitlockerRecoveryKeysRequestBuilder) Get(ctx context.Context, requestConfiguration *BitlockerRecoveryKeysRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BitlockerRecoveryKeyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation the recovery keys associated with the bitlocker entity.
// returns a *RequestInformation when successful
func (m *BitlockerRecoveryKeysRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BitlockerRecoveryKeysRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *BitlockerRecoveryKeysRequestBuilder when successful
func (m *BitlockerRecoveryKeysRequestBuilder) WithUrl(rawUrl string)(*BitlockerRecoveryKeysRequestBuilder) {
    return NewBitlockerRecoveryKeysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
