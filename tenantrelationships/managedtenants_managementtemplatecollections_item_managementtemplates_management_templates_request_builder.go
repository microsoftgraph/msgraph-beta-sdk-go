package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managementTemplateCollection entity.
type ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilderGetQueryParameters get managementTemplates from tenantRelationships
type ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilderGetQueryParameters struct {
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
// ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilderGetQueryParameters
}
// ByManagementTemplateId provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managementTemplateCollection entity.
// returns a *ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplateItemRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder) ByManagementTemplateId(managementTemplateId string)(*ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managementTemplateId != "" {
        urlTplParams["managementTemplate%2Did"] = managementTemplateId
    }
    return NewManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilderInternal instantiates a new ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder) {
    m := &ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplateCollections/{managementTemplateCollection%2Did}/managementTemplates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder instantiates a new ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesCountRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder) Count()(*ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesCountRequestBuilder) {
    return NewManagedtenantsManagementtemplatecollectionsItemManagementtemplatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get managementTemplates from tenantRelationships
// returns a ManagementTemplateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionResponseable), nil
}
// ToGetRequestInformation get managementTemplates from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder) {
    return NewManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
