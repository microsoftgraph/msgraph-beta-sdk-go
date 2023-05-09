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
// NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderInternal instantiates a new MicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) {
    m := &ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepVersions/{managementTemplateStepVersion%2Did}/deployments/{managementTemplateStepDeployment%2Did}/microsoft.graph.managedTenants.changeDeploymentStatus", pathParameters),
    }
    return m
}
// NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder instantiates a new MicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action changeDeploymentStatus
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) Post(ctx context.Context, body ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBodyable, requestConfiguration *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepDeploymentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBodyable, requestConfiguration *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
