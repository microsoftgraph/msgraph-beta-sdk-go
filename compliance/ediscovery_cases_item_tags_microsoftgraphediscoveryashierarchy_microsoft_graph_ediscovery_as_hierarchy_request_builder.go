package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder provides operations to call the asHierarchy method.
type EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetQueryParameters invoke function asHierarchy
type EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetQueryParameters struct {
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
// EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetQueryParameters
}
// NewEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderInternal instantiates a new EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) {
    m := &EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/tags/microsoft.graph.ediscovery.asHierarchy(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder instantiates a new EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function asHierarchy
// Deprecated: This method is obsolete. Use GetAsAsHierarchyGetResponse instead.
// returns a EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) Get(ctx context.Context, requestConfiguration *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetRequestConfiguration)(EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyResponseable), nil
}
// GetAsAsHierarchyGetResponse invoke function asHierarchy
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) GetAsAsHierarchyGetResponse(ctx context.Context, requestConfiguration *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetRequestConfiguration)(EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyAsHierarchyGetResponseable), nil
}
// ToGetRequestInformation invoke function asHierarchy
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder when successful
func (m *EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder) {
    return NewEdiscoveryCasesItemTagsMicrosoftgraphediscoveryashierarchyMicrosoftGraphEdiscoveryAsHierarchyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
