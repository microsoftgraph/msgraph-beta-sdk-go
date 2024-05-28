package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder provides operations to manage the acceptedFor property of the microsoft.graph.managedTenants.managementTemplateStepVersion entity.
type ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilderGetQueryParameters get acceptedFor from tenantRelationships
type ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilderGetQueryParameters
}
// NewManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilderInternal instantiates a new ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder) {
    m := &ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepVersions/{managementTemplateStepVersion%2Did}/acceptedFor{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder instantiates a new ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get acceptedFor from tenantRelationships
// returns a ManagementTemplateStepable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepable, error) {
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
// ToGetRequestInformation get acceptedFor from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder) {
    return NewManagedtenantsManagementtemplatestepversionsItemAcceptedforAcceptedForRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
