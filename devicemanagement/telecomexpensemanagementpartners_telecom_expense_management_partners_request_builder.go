package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
type TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetQueryParameters the telecom expense management partners.
type TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetQueryParameters struct {
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
// TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetQueryParameters
}
// TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTelecomExpenseManagementPartnerId provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
// returns a *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnerItemRequestBuilder when successful
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) ByTelecomExpenseManagementPartnerId(telecomExpenseManagementPartnerId string)(*TelecomexpensemanagementpartnersTelecomExpenseManagementPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if telecomExpenseManagementPartnerId != "" {
        urlTplParams["telecomExpenseManagementPartner%2Did"] = telecomExpenseManagementPartnerId
    }
    return NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderInternal instantiates a new TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder and sets the default values.
func NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) {
    m := &TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/telecomExpenseManagementPartners{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder instantiates a new TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder and sets the default values.
func NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TelecomexpensemanagementpartnersCountRequestBuilder when successful
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) Count()(*TelecomexpensemanagementpartnersCountRequestBuilder) {
    return NewTelecomexpensemanagementpartnersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the telecom expense management partners.
// returns a TelecomExpenseManagementPartnerCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) Get(ctx context.Context, requestConfiguration *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TelecomExpenseManagementPartnerCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTelecomExpenseManagementPartnerCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TelecomExpenseManagementPartnerCollectionResponseable), nil
}
// Post create new navigation property to telecomExpenseManagementPartners for deviceManagement
// returns a TelecomExpenseManagementPartnerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TelecomExpenseManagementPartnerable, requestConfiguration *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TelecomExpenseManagementPartnerable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTelecomExpenseManagementPartnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TelecomExpenseManagementPartnerable), nil
}
// ToGetRequestInformation the telecom expense management partners.
// returns a *RequestInformation when successful
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to telecomExpenseManagementPartners for deviceManagement
// returns a *RequestInformation when successful
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TelecomExpenseManagementPartnerable, requestConfiguration *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder when successful
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) WithUrl(rawUrl string)(*TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) {
    return NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
