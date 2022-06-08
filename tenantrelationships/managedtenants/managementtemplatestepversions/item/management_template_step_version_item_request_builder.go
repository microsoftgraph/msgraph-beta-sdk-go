package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
    i37f8fd44c4b8d4efac150369caf9889b3af38351154a940eccfc4e1b186fd7fe "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatestepversions/item/templatestep"
    idc9c9c72187536762e55249e42165f88e95d633f6ef070b6e7d829abb6ab1f0e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatestepversions/item/deployments"
    if87b436cc3a02d65f692ca138fd98075e7ec8cbf48477638ecdd006f2b0afb25 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatestepversions/item/acceptedfor"
    i0ba3e5fdbb0b3586ca7fde15d259d4e5860a38c09cc28ffd1bc9850dfeacf5ac "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatestepversions/item/deployments/item"
)

// ManagementTemplateStepVersionItemRequestBuilder provides operations to manage the managementTemplateStepVersions property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagementTemplateStepVersionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagementTemplateStepVersionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagementTemplateStepVersionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagementTemplateStepVersionItemRequestBuilderGetQueryParameters get managementTemplateStepVersions from tenantRelationships
type ManagementTemplateStepVersionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagementTemplateStepVersionItemRequestBuilderGetQueryParameters
}
// ManagementTemplateStepVersionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagementTemplateStepVersionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptedFor the acceptedFor property
func (m *ManagementTemplateStepVersionItemRequestBuilder) AcceptedFor()(*if87b436cc3a02d65f692ca138fd98075e7ec8cbf48477638ecdd006f2b0afb25.AcceptedForRequestBuilder) {
    return if87b436cc3a02d65f692ca138fd98075e7ec8cbf48477638ecdd006f2b0afb25.NewAcceptedForRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagementTemplateStepVersionItemRequestBuilderInternal instantiates a new ManagementTemplateStepVersionItemRequestBuilder and sets the default values.
func NewManagementTemplateStepVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagementTemplateStepVersionItemRequestBuilder) {
    m := &ManagementTemplateStepVersionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepVersions/{managementTemplateStepVersion%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagementTemplateStepVersionItemRequestBuilder instantiates a new ManagementTemplateStepVersionItemRequestBuilder and sets the default values.
func NewManagementTemplateStepVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagementTemplateStepVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagementTemplateStepVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managementTemplateStepVersions for tenantRelationships
func (m *ManagementTemplateStepVersionItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property managementTemplateStepVersions for tenantRelationships
func (m *ManagementTemplateStepVersionItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ManagementTemplateStepVersionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get managementTemplateStepVersions from tenantRelationships
func (m *ManagementTemplateStepVersionItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get managementTemplateStepVersions from tenantRelationships
func (m *ManagementTemplateStepVersionItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property managementTemplateStepVersions in tenantRelationships
func (m *ManagementTemplateStepVersionItemRequestBuilder) CreatePatchRequestInformation(body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property managementTemplateStepVersions in tenantRelationships
func (m *ManagementTemplateStepVersionItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, requestConfiguration *ManagementTemplateStepVersionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property managementTemplateStepVersions for tenantRelationships
func (m *ManagementTemplateStepVersionItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property managementTemplateStepVersions for tenantRelationships
func (m *ManagementTemplateStepVersionItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ManagementTemplateStepVersionItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Deployments the deployments property
func (m *ManagementTemplateStepVersionItemRequestBuilder) Deployments()(*idc9c9c72187536762e55249e42165f88e95d633f6ef070b6e7d829abb6ab1f0e.DeploymentsRequestBuilder) {
    return idc9c9c72187536762e55249e42165f88e95d633f6ef070b6e7d829abb6ab1f0e.NewDeploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeploymentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.managedTenants.managementTemplateStepVersions.item.deployments.item collection
func (m *ManagementTemplateStepVersionItemRequestBuilder) DeploymentsById(id string)(*i0ba3e5fdbb0b3586ca7fde15d259d4e5860a38c09cc28ffd1bc9850dfeacf5ac.ManagementTemplateStepDeploymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStepDeployment%2Did"] = id
    }
    return i0ba3e5fdbb0b3586ca7fde15d259d4e5860a38c09cc28ffd1bc9850dfeacf5ac.NewManagementTemplateStepDeploymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get managementTemplateStepVersions from tenantRelationships
func (m *ManagementTemplateStepVersionItemRequestBuilder) Get()(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get managementTemplateStepVersions from tenantRelationships
func (m *ManagementTemplateStepVersionItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepVersionFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable), nil
}
// Patch update the navigation property managementTemplateStepVersions in tenantRelationships
func (m *ManagementTemplateStepVersionItemRequestBuilder) Patch(body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property managementTemplateStepVersions in tenantRelationships
func (m *ManagementTemplateStepVersionItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, requestConfiguration *ManagementTemplateStepVersionItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// TemplateStep the templateStep property
func (m *ManagementTemplateStepVersionItemRequestBuilder) TemplateStep()(*i37f8fd44c4b8d4efac150369caf9889b3af38351154a940eccfc4e1b186fd7fe.TemplateStepRequestBuilder) {
    return i37f8fd44c4b8d4efac150369caf9889b3af38351154a940eccfc4e1b186fd7fe.NewTemplateStepRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
