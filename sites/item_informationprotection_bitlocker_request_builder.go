package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemInformationprotectionBitlockerRequestBuilder provides operations to manage the bitlocker property of the microsoft.graph.informationProtection entity.
type ItemInformationprotectionBitlockerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInformationprotectionBitlockerRequestBuilderGetQueryParameters get bitlocker from sites
type ItemInformationprotectionBitlockerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemInformationprotectionBitlockerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInformationprotectionBitlockerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInformationprotectionBitlockerRequestBuilderGetQueryParameters
}
// NewItemInformationprotectionBitlockerRequestBuilderInternal instantiates a new ItemInformationprotectionBitlockerRequestBuilder and sets the default values.
func NewItemInformationprotectionBitlockerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionBitlockerRequestBuilder) {
    m := &ItemInformationprotectionBitlockerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/informationProtection/bitlocker{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemInformationprotectionBitlockerRequestBuilder instantiates a new ItemInformationprotectionBitlockerRequestBuilder and sets the default values.
func NewItemInformationprotectionBitlockerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInformationprotectionBitlockerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInformationprotectionBitlockerRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get bitlocker from sites
// returns a Bitlockerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInformationprotectionBitlockerRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInformationprotectionBitlockerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Bitlockerable, error) {
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
// returns a *ItemInformationprotectionBitlockerRecoverykeysRecoveryKeysRequestBuilder when successful
func (m *ItemInformationprotectionBitlockerRequestBuilder) RecoveryKeys()(*ItemInformationprotectionBitlockerRecoverykeysRecoveryKeysRequestBuilder) {
    return NewItemInformationprotectionBitlockerRecoverykeysRecoveryKeysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get bitlocker from sites
// returns a *RequestInformation when successful
func (m *ItemInformationprotectionBitlockerRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInformationprotectionBitlockerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemInformationprotectionBitlockerRequestBuilder when successful
func (m *ItemInformationprotectionBitlockerRequestBuilder) WithUrl(rawUrl string)(*ItemInformationprotectionBitlockerRequestBuilder) {
    return NewItemInformationprotectionBitlockerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
