package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder provides operations to manage the templateStep property of the microsoft.graph.managedTenants.managementTemplateStepVersion entity.
type ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilderGetQueryParameters get templateStep from tenantRelationships
type ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilderGetQueryParameters
}
// NewManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilderInternal instantiates a new ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder) {
    m := &ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepVersions/{managementTemplateStepVersion%2Did}/templateStep{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder instantiates a new ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get templateStep from tenantRelationships
// returns a ManagementTemplateStepable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepable), nil
}
// ToGetRequestInformation get templateStep from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder) {
    return NewManagedtenantsManagementtemplatestepversionsItemTemplatestepTemplateStepRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
