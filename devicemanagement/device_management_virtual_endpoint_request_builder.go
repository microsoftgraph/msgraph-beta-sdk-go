package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementVirtualEndpointRequestBuilder provides operations to manage the virtualEndpoint property of the microsoft.graph.deviceManagement entity.
type DeviceManagementVirtualEndpointRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementVirtualEndpointRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementVirtualEndpointRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementVirtualEndpointRequestBuilderGetQueryParameters get virtualEndpoint from deviceManagement
type DeviceManagementVirtualEndpointRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementVirtualEndpointRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementVirtualEndpointRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementVirtualEndpointRequestBuilderGetQueryParameters
}
// DeviceManagementVirtualEndpointRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementVirtualEndpointRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) AuditEvents()(*DeviceManagementVirtualEndpointAuditEventsRequestBuilder) {
    return NewDeviceManagementVirtualEndpointAuditEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuditEventsById provides operations to manage the auditEvents property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) AuditEventsById(id string)(*DeviceManagementVirtualEndpointAuditEventsCloudPcAuditEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcAuditEvent%2Did"] = id
    }
    return NewDeviceManagementVirtualEndpointAuditEventsCloudPcAuditEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPCs provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) CloudPCs()(*DeviceManagementVirtualEndpointCloudPCsRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCsById provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) CloudPCsById(id string)(*DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPC%2Did"] = id
    }
    return NewDeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementVirtualEndpointRequestBuilderInternal instantiates a new VirtualEndpointRequestBuilder and sets the default values.
func NewDeviceManagementVirtualEndpointRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementVirtualEndpointRequestBuilder) {
    m := &DeviceManagementVirtualEndpointRequestBuilder{
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
// NewDeviceManagementVirtualEndpointRequestBuilder instantiates a new VirtualEndpointRequestBuilder and sets the default values.
func NewDeviceManagementVirtualEndpointRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementVirtualEndpointRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementVirtualEndpointRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property virtualEndpoint for deviceManagement
func (m *DeviceManagementVirtualEndpointRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get virtualEndpoint from deviceManagement
func (m *DeviceManagementVirtualEndpointRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property virtualEndpoint in deviceManagement
func (m *DeviceManagementVirtualEndpointRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, requestConfiguration *DeviceManagementVirtualEndpointRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CrossCloudGovernmentOrganizationMapping provides operations to manage the crossCloudGovernmentOrganizationMapping property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) CrossCloudGovernmentOrganizationMapping()(*DeviceManagementVirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property virtualEndpoint for deviceManagement
func (m *DeviceManagementVirtualEndpointRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DeviceManagementVirtualEndpointRequestBuilder) DeviceImages()(*DeviceManagementVirtualEndpointDeviceImagesRequestBuilder) {
    return NewDeviceManagementVirtualEndpointDeviceImagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceImagesById provides operations to manage the deviceImages property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) DeviceImagesById(id string)(*DeviceManagementVirtualEndpointDeviceImagesCloudPcDeviceImageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcDeviceImage%2Did"] = id
    }
    return NewDeviceManagementVirtualEndpointDeviceImagesCloudPcDeviceImageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExternalPartnerSettings provides operations to manage the externalPartnerSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) ExternalPartnerSettings()(*DeviceManagementVirtualEndpointExternalPartnerSettingsRequestBuilder) {
    return NewDeviceManagementVirtualEndpointExternalPartnerSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalPartnerSettingsById provides operations to manage the externalPartnerSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) ExternalPartnerSettingsById(id string)(*DeviceManagementVirtualEndpointExternalPartnerSettingsCloudPcExternalPartnerSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcExternalPartnerSetting%2Did"] = id
    }
    return NewDeviceManagementVirtualEndpointExternalPartnerSettingsCloudPcExternalPartnerSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GalleryImages provides operations to manage the galleryImages property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) GalleryImages()(*DeviceManagementVirtualEndpointGalleryImagesRequestBuilder) {
    return NewDeviceManagementVirtualEndpointGalleryImagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GalleryImagesById provides operations to manage the galleryImages property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) GalleryImagesById(id string)(*DeviceManagementVirtualEndpointGalleryImagesCloudPcGalleryImageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcGalleryImage%2Did"] = id
    }
    return NewDeviceManagementVirtualEndpointGalleryImagesCloudPcGalleryImageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get virtualEndpoint from deviceManagement
func (m *DeviceManagementVirtualEndpointRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, error) {
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
func (m *DeviceManagementVirtualEndpointRequestBuilder) GetEffectivePermissions()(*DeviceManagementVirtualEndpointGetEffectivePermissionsRequestBuilder) {
    return NewDeviceManagementVirtualEndpointGetEffectivePermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnPremisesConnections provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) OnPremisesConnections()(*DeviceManagementVirtualEndpointOnPremisesConnectionsRequestBuilder) {
    return NewDeviceManagementVirtualEndpointOnPremisesConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnPremisesConnectionsById provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) OnPremisesConnectionsById(id string)(*DeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcOnPremisesConnection%2Did"] = id
    }
    return NewDeviceManagementVirtualEndpointOnPremisesConnectionsCloudPcOnPremisesConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OrganizationSettings provides operations to manage the organizationSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) OrganizationSettings()(*DeviceManagementVirtualEndpointOrganizationSettingsRequestBuilder) {
    return NewDeviceManagementVirtualEndpointOrganizationSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property virtualEndpoint in deviceManagement
func (m *DeviceManagementVirtualEndpointRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, requestConfiguration *DeviceManagementVirtualEndpointRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, error) {
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
func (m *DeviceManagementVirtualEndpointRequestBuilder) ProvisioningPolicies()(*DeviceManagementVirtualEndpointProvisioningPoliciesRequestBuilder) {
    return NewDeviceManagementVirtualEndpointProvisioningPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisioningPoliciesById provides operations to manage the provisioningPolicies property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) ProvisioningPoliciesById(id string)(*DeviceManagementVirtualEndpointProvisioningPoliciesCloudPcProvisioningPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcProvisioningPolicy%2Did"] = id
    }
    return NewDeviceManagementVirtualEndpointProvisioningPoliciesCloudPcProvisioningPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Reports provides operations to manage the reports property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) Reports()(*DeviceManagementVirtualEndpointReportsRequestBuilder) {
    return NewDeviceManagementVirtualEndpointReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePlans provides operations to manage the servicePlans property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) ServicePlans()(*DeviceManagementVirtualEndpointServicePlansRequestBuilder) {
    return NewDeviceManagementVirtualEndpointServicePlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePlansById provides operations to manage the servicePlans property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) ServicePlansById(id string)(*DeviceManagementVirtualEndpointServicePlansCloudPcServicePlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcServicePlan%2Did"] = id
    }
    return NewDeviceManagementVirtualEndpointServicePlansCloudPcServicePlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SharedUseServicePlans provides operations to manage the sharedUseServicePlans property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) SharedUseServicePlans()(*DeviceManagementVirtualEndpointSharedUseServicePlansRequestBuilder) {
    return NewDeviceManagementVirtualEndpointSharedUseServicePlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedUseServicePlansById provides operations to manage the sharedUseServicePlans property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) SharedUseServicePlansById(id string)(*DeviceManagementVirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcSharedUseServicePlan%2Did"] = id
    }
    return NewDeviceManagementVirtualEndpointSharedUseServicePlansCloudPcSharedUseServicePlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Snapshots provides operations to manage the snapshots property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) Snapshots()(*DeviceManagementVirtualEndpointSnapshotsRequestBuilder) {
    return NewDeviceManagementVirtualEndpointSnapshotsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SnapshotsById provides operations to manage the snapshots property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) SnapshotsById(id string)(*DeviceManagementVirtualEndpointSnapshotsCloudPcSnapshotItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcSnapshot%2Did"] = id
    }
    return NewDeviceManagementVirtualEndpointSnapshotsCloudPcSnapshotItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SupportedRegions provides operations to manage the supportedRegions property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) SupportedRegions()(*DeviceManagementVirtualEndpointSupportedRegionsRequestBuilder) {
    return NewDeviceManagementVirtualEndpointSupportedRegionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedRegionsById provides operations to manage the supportedRegions property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) SupportedRegionsById(id string)(*DeviceManagementVirtualEndpointSupportedRegionsCloudPcSupportedRegionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcSupportedRegion%2Did"] = id
    }
    return NewDeviceManagementVirtualEndpointSupportedRegionsCloudPcSupportedRegionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserSettings provides operations to manage the userSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) UserSettings()(*DeviceManagementVirtualEndpointUserSettingsRequestBuilder) {
    return NewDeviceManagementVirtualEndpointUserSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSettingsById provides operations to manage the userSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *DeviceManagementVirtualEndpointRequestBuilder) UserSettingsById(id string)(*DeviceManagementVirtualEndpointUserSettingsCloudPcUserSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcUserSetting%2Did"] = id
    }
    return NewDeviceManagementVirtualEndpointUserSettingsCloudPcUserSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
