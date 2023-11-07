package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\legalHolds\{legalHold-id}\unifiedGroupSources\{unifiedGroupSource-id}\group\serviceProvisioningErrors
type EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilderGetQueryParameters errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from a group object.  Supports $filter (eq, not, for isResolved and serviceInstance).
type EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilderGetQueryParameters struct {
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
// EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilderGetQueryParameters
}
// NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilderInternal instantiates a new ServiceProvisioningErrorsRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder) {
    m := &EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/legalHolds/{legalHold%2Did}/unifiedGroupSources/{unifiedGroupSource%2Did}/group/serviceProvisioningErrors{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder instantiates a new ServiceProvisioningErrorsRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder) Count()(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsCountRequestBuilder) {
    return NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from a group object.  Supports $filter (eq, not, for isResolved and serviceInstance).
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceProvisioningErrorCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServiceProvisioningErrorCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServiceProvisioningErrorCollectionResponseable), nil
}
// ToGetRequestInformation errors published by a federated service describing a non-transient, service-specific error regarding the properties or link from a group object.  Supports $filter (eq, not, for isResolved and serviceInstance).
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder) {
    return NewEdiscoveryCasesItemLegalHoldsItemUnifiedGroupSourcesItemGroupServiceProvisioningErrorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
