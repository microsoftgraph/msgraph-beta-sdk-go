package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilder provides operations to call the changeDeploymentStatus method.
type ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilderInternal instantiates a new ChangeDeploymentStatusRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilder) {
    m := &ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepVersions/{managementTemplateStepVersion%2Did}/deployments/{managementTemplateStepDeployment%2Did}/microsoft.graph.managedTenants.changeDeploymentStatus";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilder instantiates a new ChangeDeploymentStatusRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action changeDeploymentStatus
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilder) Post(ctx context.Context, body ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBodyable, requestConfiguration *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepDeploymentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepDeploymentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepDeploymentable), nil
}
// ToPostRequestInformation invoke action changeDeploymentStatus
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBodyable, requestConfiguration *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
