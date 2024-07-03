package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder provides operations to manage the managementTemplateSteps property of the microsoft.graph.managedTenants.managementTemplate entity.
type ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilderGetQueryParameters get managementTemplateSteps from tenantRelationships
type ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilderGetQueryParameters struct {
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
// ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilderGetQueryParameters
}
// ByManagementTemplateStepId provides operations to manage the managementTemplateSteps property of the microsoft.graph.managedTenants.managementTemplate entity.
// returns a *ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder when successful
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder) ByManagementTemplateStepId(managementTemplateStepId string)(*ManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managementTemplateStepId != "" {
        urlTplParams["managementTemplateStep%2Did"] = managementTemplateStepId
    }
    return NewManagedTenantsManagementTemplatesItemManagementTemplateStepsManagementTemplateStepItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilderInternal instantiates a new ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder) {
    m := &ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplates/{managementTemplate%2Did}/managementTemplateSteps{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder instantiates a new ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedTenantsManagementTemplatesItemManagementTemplateStepsCountRequestBuilder when successful
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder) Count()(*ManagedTenantsManagementTemplatesItemManagementTemplateStepsCountRequestBuilder) {
    return NewManagedTenantsManagementTemplatesItemManagementTemplateStepsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get managementTemplateSteps from tenantRelationships
// returns a ManagementTemplateStepCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepCollectionResponseable), nil
}
// ToGetRequestInformation get managementTemplateSteps from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder when successful
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder) {
    return NewManagedTenantsManagementTemplatesItemManagementTemplateStepsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
