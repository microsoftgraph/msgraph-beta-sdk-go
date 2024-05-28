package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder provides operations to manage the group property of the microsoft.graph.ediscovery.unifiedGroupSource entity.
type EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilderGetQueryParameters the group associated with the unifiedGroupSource.
type EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilderGetQueryParameters
}
// NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilderInternal instantiates a new EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder) {
    m := &EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/legalHolds/{legalHold%2Did}/unifiedGroupSources/{unifiedGroupSource%2Did}/group{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder instantiates a new EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the group associated with the unifiedGroupSource.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a Groupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder) ServiceProvisioningErrors()(*EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the group associated with the unifiedGroupSource.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder when successful
func (m *EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder) {
    return NewEdiscoveryCasesItemLegalholdsItemUnifiedgroupsourcesItemGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
