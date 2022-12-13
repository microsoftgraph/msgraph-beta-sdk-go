package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointRequestBuilder provides operations to manage the virtualEndpoint property of the microsoft.graph.deviceManagement entity.
type VirtualEndpointRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// VirtualEndpointRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEndpointRequestBuilderGetQueryParameters get virtualEndpoint from deviceManagement
type VirtualEndpointRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEndpointRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointRequestBuilderGetQueryParameters
}
// VirtualEndpointRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) AuditEvents()(*VirtualEndpointAuditEventsRequestBuilder) {
    return NewVirtualEndpointAuditEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuditEventsById provides operations to manage the auditEvents property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) AuditEventsById(id string)(*VirtualEndpointAuditEventsCloudPcAuditEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcAuditEvent%2Did"] = id
    }
    return NewVirtualEndpointAuditEventsCloudPcAuditEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPCs provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) CloudPCs()(*VirtualEndpointCloudPCsRequestBuilder) {
    return NewVirtualEndpointCloudPCsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCsById provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) CloudPCsById(id string)(*VirtualEndpointCloudPCsCloudPCItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPC%2Did"] = id
    }
    return NewVirtualEndpointCloudPCsCloudPCItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewVirtualEndpointRequestBuilderInternal instantiates a new VirtualEndpointRequestBuilder and sets the default values.
func NewVirtualEndpointRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointRequestBuilder) {
    m := &VirtualEndpointRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewVirtualEndpointRequestBuilder instantiates a new VirtualEndpointRequestBuilder and sets the default values.
func NewVirtualEndpointRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property virtualEndpoint for deviceManagement
func (m *VirtualEndpointRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get virtualEndpoint from deviceManagement
func (m *VirtualEndpointRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property virtualEndpoint in deviceManagement
func (m *VirtualEndpointRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, requestConfiguration *VirtualEndpointRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CrossCloudGovernmentOrganizationMapping provides operations to manage the crossCloudGovernmentOrganizationMapping property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) CrossCloudGovernmentOrganizationMapping()(*VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) {
    return NewVirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property virtualEndpoint for deviceManagement
func (m *VirtualEndpointRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEndpointRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceImages provides operations to manage the deviceImages property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) DeviceImages()(*VirtualEndpointDeviceImagesRequestBuilder) {
    return NewVirtualEndpointDeviceImagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceImagesById provides operations to manage the deviceImages property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) DeviceImagesById(id string)(*VirtualEndpointDeviceImagesCloudPcDeviceImageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcDeviceImage%2Did"] = id
    }
    return NewVirtualEndpointDeviceImagesCloudPcDeviceImageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExternalPartnerSettings provides operations to manage the externalPartnerSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) ExternalPartnerSettings()(*VirtualEndpointExternalPartnerSettingsRequestBuilder) {
    return NewVirtualEndpointExternalPartnerSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalPartnerSettingsById provides operations to manage the externalPartnerSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) ExternalPartnerSettingsById(id string)(*VirtualEndpointExternalPartnerSettingsCloudPcExternalPartnerSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcExternalPartnerSetting%2Did"] = id
    }
    return NewVirtualEndpointExternalPartnerSettingsCloudPcExternalPartnerSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GalleryImages provides operations to manage the galleryImages property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) GalleryImages()(*VirtualEndpointGalleryImagesRequestBuilder) {
    return NewVirtualEndpointGalleryImagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GalleryImagesById provides operations to manage the galleryImages property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) GalleryImagesById(id string)(*VirtualEndpointGalleryImagesCloudPcGalleryImageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcGalleryImage%2Did"] = id
    }
    return NewVirtualEndpointGalleryImagesCloudPcGalleryImageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get virtualEndpoint from deviceManagement
func (m *VirtualEndpointRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEndpointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable), nil
}
// GetEffectivePermissions provides operations to call the getEffectivePermissions method.
func (m *VirtualEndpointRequestBuilder) GetEffectivePermissions()(*VirtualEndpointGetEffectivePermissionsRequestBuilder) {
    return NewVirtualEndpointGetEffectivePermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnPremisesConnections provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) OnPremisesConnections()(*VirtualEndpointOnPremisesConnectionsRequestBuilder) {
    return NewVirtualEndpointOnPremisesConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnPremisesConnectionsById provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) OnPremisesConnectionsById(id string)(*VirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcOnPremisesConnection%2Did"] = id
    }
    return NewVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OrganizationSettings provides operations to manage the organizationSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) OrganizationSettings()(*VirtualEndpointOrganizationSettingsRequestBuilder) {
    return NewVirtualEndpointOrganizationSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property virtualEndpoint in deviceManagement
func (m *VirtualEndpointRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, requestConfiguration *VirtualEndpointRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEndpointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable), nil
}
// ProvisioningPolicies provides operations to manage the provisioningPolicies property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) ProvisioningPolicies()(*VirtualEndpointProvisioningPoliciesRequestBuilder) {
    return NewVirtualEndpointProvisioningPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisioningPoliciesById provides operations to manage the provisioningPolicies property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) ProvisioningPoliciesById(id string)(*VirtualEndpointProvisioningPoliciesCloudPcProvisioningPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcProvisioningPolicy%2Did"] = id
    }
    return NewVirtualEndpointProvisioningPoliciesCloudPcProvisioningPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Reports provides operations to manage the reports property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) Reports()(*VirtualEndpointReportsRequestBuilder) {
    return NewVirtualEndpointReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePlans provides operations to manage the servicePlans property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) ServicePlans()(*VirtualEndpointServicePlansRequestBuilder) {
    return NewVirtualEndpointServicePlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePlansById provides operations to manage the servicePlans property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) ServicePlansById(id string)(*VirtualEndpointServicePlansCloudPcServicePlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcServicePlan%2Did"] = id
    }
    return NewVirtualEndpointServicePlansCloudPcServicePlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SharedUseServicePlans provides operations to manage the sharedUseServicePlans property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) SharedUseServicePlans()(*VirtualEndpointSharedUseServicePlansRequestBuilder) {
    return NewVirtualEndpointSharedUseServicePlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedUseServicePlansById provides operations to manage the sharedUseServicePlans property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) SharedUseServicePlansById(id string)(*VirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcSharedUseServicePlan%2Did"] = id
    }
    return NewVirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Snapshots provides operations to manage the snapshots property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) Snapshots()(*VirtualEndpointSnapshotsRequestBuilder) {
    return NewVirtualEndpointSnapshotsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SnapshotsById provides operations to manage the snapshots property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) SnapshotsById(id string)(*VirtualEndpointSnapshotsCloudPcSnapshotItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcSnapshot%2Did"] = id
    }
    return NewVirtualEndpointSnapshotsCloudPcSnapshotItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SupportedRegions provides operations to manage the supportedRegions property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) SupportedRegions()(*VirtualEndpointSupportedRegionsRequestBuilder) {
    return NewVirtualEndpointSupportedRegionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedRegionsById provides operations to manage the supportedRegions property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) SupportedRegionsById(id string)(*VirtualEndpointSupportedRegionsCloudPcSupportedRegionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcSupportedRegion%2Did"] = id
    }
    return NewVirtualEndpointSupportedRegionsCloudPcSupportedRegionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserSettings provides operations to manage the userSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) UserSettings()(*VirtualEndpointUserSettingsRequestBuilder) {
    return NewVirtualEndpointUserSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSettingsById provides operations to manage the userSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) UserSettingsById(id string)(*VirtualEndpointUserSettingsCloudPcUserSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcUserSetting%2Did"] = id
    }
    return NewVirtualEndpointUserSettingsCloudPcUserSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
