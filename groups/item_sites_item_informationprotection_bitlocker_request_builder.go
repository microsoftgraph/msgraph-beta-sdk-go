package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationprotectionBitlockerRequestBuilder provides operations to manage the bitlocker property of the microsoft.graph.informationProtection entity.
type ItemSitesItemInformationprotectionBitlockerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationprotectionBitlockerRequestBuilderGetQueryParameters get bitlocker from groups
type ItemSitesItemInformationprotectionBitlockerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemInformationprotectionBitlockerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationprotectionBitlockerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemInformationprotectionBitlockerRequestBuilderGetQueryParameters
}
// NewItemSitesItemInformationprotectionBitlockerRequestBuilderInternal instantiates a new ItemSitesItemInformationprotectionBitlockerRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionBitlockerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionBitlockerRequestBuilder) {
    m := &ItemSitesItemInformationprotectionBitlockerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/bitlocker{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationprotectionBitlockerRequestBuilder instantiates a new ItemSitesItemInformationprotectionBitlockerRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionBitlockerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionBitlockerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationprotectionBitlockerRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get bitlocker from groups
// returns a Bitlockerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationprotectionBitlockerRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemInformationprotectionBitlockerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Bitlockerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateBitlockerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Bitlockerable), nil
}
// RecoveryKeys provides operations to manage the recoveryKeys property of the microsoft.graph.bitlocker entity.
// returns a *ItemSitesItemInformationprotectionBitlockerRecoverykeysRecoveryKeysRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionBitlockerRequestBuilder) RecoveryKeys()(*ItemSitesItemInformationprotectionBitlockerRecoverykeysRecoveryKeysRequestBuilder) {
    return NewItemSitesItemInformationprotectionBitlockerRecoverykeysRecoveryKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get bitlocker from groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationprotectionBitlockerRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemInformationprotectionBitlockerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemInformationprotectionBitlockerRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionBitlockerRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationprotectionBitlockerRequestBuilder) {
    return NewItemSitesItemInformationprotectionBitlockerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
