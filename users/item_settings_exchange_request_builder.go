package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSettingsExchangeRequestBuilder provides operations to manage the exchange property of the microsoft.graph.userSettings entity.
type ItemSettingsExchangeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSettingsExchangeRequestBuilderGetQueryParameters get a list of Exchange settings, including mailboxes that belong to a user. Currently, the mailbox types supported are the user's primary and in-place archive. To learn how to get a list of users in a tenant, see List users.
type ItemSettingsExchangeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSettingsExchangeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSettingsExchangeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSettingsExchangeRequestBuilderGetQueryParameters
}
// NewItemSettingsExchangeRequestBuilderInternal instantiates a new ItemSettingsExchangeRequestBuilder and sets the default values.
func NewItemSettingsExchangeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsExchangeRequestBuilder) {
    m := &ItemSettingsExchangeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/settings/exchange{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSettingsExchangeRequestBuilder instantiates a new ItemSettingsExchangeRequestBuilder and sets the default values.
func NewItemSettingsExchangeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSettingsExchangeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSettingsExchangeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a list of Exchange settings, including mailboxes that belong to a user. Currently, the mailbox types supported are the user's primary and in-place archive. To learn how to get a list of users in a tenant, see List users.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a ExchangeSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/usersettings-list-exchange?view=graph-rest-beta
func (m *ItemSettingsExchangeRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSettingsExchangeRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExchangeSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExchangeSettingsable), nil
}
// ToGetRequestInformation get a list of Exchange settings, including mailboxes that belong to a user. Currently, the mailbox types supported are the user's primary and in-place archive. To learn how to get a list of users in a tenant, see List users.
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *RequestInformation when successful
func (m *ItemSettingsExchangeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSettingsExchangeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: Private preview for Import Export APIs as of 2021-08/PrivatePreview:importExport
// returns a *ItemSettingsExchangeRequestBuilder when successful
func (m *ItemSettingsExchangeRequestBuilder) WithUrl(rawUrl string)(*ItemSettingsExchangeRequestBuilder) {
    return NewItemSettingsExchangeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
