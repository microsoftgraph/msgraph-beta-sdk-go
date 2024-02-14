package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder provides operations to call the changeDeploymentStatus method.
type ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderInternal instantiates a new ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) {
    m := &ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepVersions/{managementTemplateStepVersion%2Did}/deployments/{managementTemplateStepDeployment%2Did}/microsoft.graph.managedTenants.changeDeploymentStatus", pathParameters),
    }
    return m
}
// NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder instantiates a new ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action changeDeploymentStatus
// returns a ManagementTemplateStepDeploymentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) Post(ctx context.Context, body ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBodyable, requestConfiguration *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepDeploymentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepDeploymentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepDeploymentable), nil
}
// ToPostRequestInformation invoke action changeDeploymentStatus
// returns a *RequestInformation when successful
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBodyable, requestConfiguration *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder when successful
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) {
    return NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
