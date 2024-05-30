package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder provides operations to manage the managementTemplateStepTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderGetQueryParameters get managementTemplateStepTenantSummaries from tenantRelationships
type ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderGetQueryParameters struct {
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
// ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderGetQueryParameters
}
// ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagementTemplateStepTenantSummaryId provides operations to manage the managementTemplateStepTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummaryItemRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder) ByManagementTemplateStepTenantSummaryId(managementTemplateStepTenantSummaryId string)(*ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managementTemplateStepTenantSummaryId != "" {
        urlTplParams["managementTemplateStepTenantSummary%2Did"] = managementTemplateStepTenantSummaryId
    }
    return NewManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummaryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderInternal instantiates a new ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder) {
    m := &ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepTenantSummaries{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder instantiates a new ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedtenantsManagementtemplatesteptenantsummariesCountRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder) Count()(*ManagedtenantsManagementtemplatesteptenantsummariesCountRequestBuilder) {
    return NewManagedtenantsManagementtemplatesteptenantsummariesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get managementTemplateStepTenantSummaries from tenantRelationships
// returns a ManagementTemplateStepTenantSummaryCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepTenantSummaryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepTenantSummaryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepTenantSummaryCollectionResponseable), nil
}
// Post create new navigation property to managementTemplateStepTenantSummaries for tenantRelationships
// returns a ManagementTemplateStepTenantSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepTenantSummaryable, requestConfiguration *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepTenantSummaryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepTenantSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepTenantSummaryable), nil
}
// ToGetRequestInformation get managementTemplateStepTenantSummaries from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to managementTemplateStepTenantSummaries for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepTenantSummaryable, requestConfiguration *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder) {
    return NewManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
