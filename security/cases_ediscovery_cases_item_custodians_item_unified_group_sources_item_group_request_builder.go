package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder provides operations to manage the group property of the microsoft.graph.security.unifiedGroupSource entity.
type CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilderGetQueryParameters represents a group.
type CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilderGetQueryParameters
}
// NewCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder) {
    m := &CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}/unifiedGroupSources/{unifiedGroupSource%2Did}/group{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder instantiates a new GroupRequestBuilder and sets the default values.
func NewCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get represents a group.
func (m *CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable), nil
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
func (m *CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder) ServiceProvisioningErrors()(*CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation represents a group.
func (m *CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder) {
    return NewCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
