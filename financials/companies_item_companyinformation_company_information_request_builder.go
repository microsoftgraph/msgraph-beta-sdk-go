package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemCompanyinformationCompanyInformationRequestBuilder provides operations to manage the companyInformation property of the microsoft.graph.company entity.
type CompaniesItemCompanyinformationCompanyInformationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemCompanyinformationCompanyInformationRequestBuilderGetQueryParameters get companyInformation from financials
type CompaniesItemCompanyinformationCompanyInformationRequestBuilderGetQueryParameters struct {
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
// CompaniesItemCompanyinformationCompanyInformationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCompanyinformationCompanyInformationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemCompanyinformationCompanyInformationRequestBuilderGetQueryParameters
}
// ByCompanyInformationId provides operations to manage the companyInformation property of the microsoft.graph.company entity.
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *CompaniesItemCompanyinformationCompanyInformationItemRequestBuilder when successful
func (m *CompaniesItemCompanyinformationCompanyInformationRequestBuilder) ByCompanyInformationId(companyInformationId string)(*CompaniesItemCompanyinformationCompanyInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if companyInformationId != "" {
        urlTplParams["companyInformation%2Did"] = companyInformationId
    }
    return NewCompaniesItemCompanyinformationCompanyInformationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByCompanyInformationIdGuid provides operations to manage the companyInformation property of the microsoft.graph.company entity.
// returns a *CompaniesItemCompanyinformationCompanyInformationItemRequestBuilder when successful
func (m *CompaniesItemCompanyinformationCompanyInformationRequestBuilder) ByCompanyInformationIdGuid(companyInformationId i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*CompaniesItemCompanyinformationCompanyInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["companyInformation%2Did"] = companyInformationId.String()
    return NewCompaniesItemCompanyinformationCompanyInformationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemCompanyinformationCompanyInformationRequestBuilderInternal instantiates a new CompaniesItemCompanyinformationCompanyInformationRequestBuilder and sets the default values.
func NewCompaniesItemCompanyinformationCompanyInformationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCompanyinformationCompanyInformationRequestBuilder) {
    m := &CompaniesItemCompanyinformationCompanyInformationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/companyInformation{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompaniesItemCompanyinformationCompanyInformationRequestBuilder instantiates a new CompaniesItemCompanyinformationCompanyInformationRequestBuilder and sets the default values.
func NewCompaniesItemCompanyinformationCompanyInformationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCompanyinformationCompanyInformationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemCompanyinformationCompanyInformationRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompaniesItemCompanyinformationCountRequestBuilder when successful
func (m *CompaniesItemCompanyinformationCompanyInformationRequestBuilder) Count()(*CompaniesItemCompanyinformationCountRequestBuilder) {
    return NewCompaniesItemCompanyinformationCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get companyInformation from financials
// returns a CompanyInformationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemCompanyinformationCompanyInformationRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemCompanyinformationCompanyInformationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CompanyInformationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCompanyInformationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CompanyInformationCollectionResponseable), nil
}
// ToGetRequestInformation get companyInformation from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemCompanyinformationCompanyInformationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemCompanyinformationCompanyInformationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemCompanyinformationCompanyInformationRequestBuilder when successful
func (m *CompaniesItemCompanyinformationCompanyInformationRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemCompanyinformationCompanyInformationRequestBuilder) {
    return NewCompaniesItemCompanyinformationCompanyInformationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
