package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointRequestBuilder provides operations to manage the virtualEndpoint property of the microsoft.graph.deviceManagement entity.
type VirtualEndpointRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
// returns a *VirtualEndpointAuditEventsRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) AuditEvents()(*VirtualEndpointAuditEventsRequestBuilder) {
    return NewVirtualEndpointAuditEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BulkActions provides operations to manage the bulkActions property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointBulkActionsRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) BulkActions()(*VirtualEndpointBulkActionsRequestBuilder) {
    return NewVirtualEndpointBulkActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudPCs provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointCloudPCsRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) CloudPCs()(*VirtualEndpointCloudPCsRequestBuilder) {
    return NewVirtualEndpointCloudPCsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEndpointRequestBuilderInternal instantiates a new VirtualEndpointRequestBuilder and sets the default values.
func NewVirtualEndpointRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointRequestBuilder) {
    m := &VirtualEndpointRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualEndpointRequestBuilder instantiates a new VirtualEndpointRequestBuilder and sets the default values.
func NewVirtualEndpointRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointRequestBuilderInternal(urlParams, requestAdapter)
}
// CrossCloudGovernmentOrganizationMapping provides operations to manage the crossCloudGovernmentOrganizationMapping property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) CrossCloudGovernmentOrganizationMapping()(*VirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilder) {
    return NewVirtualEndpointCrossCloudGovernmentOrganizationMappingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property virtualEndpoint for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEndpointRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceImages provides operations to manage the deviceImages property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointDeviceImagesRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) DeviceImages()(*VirtualEndpointDeviceImagesRequestBuilder) {
    return NewVirtualEndpointDeviceImagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExternalPartnerSettings provides operations to manage the externalPartnerSettings property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointExternalPartnerSettingsRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) ExternalPartnerSettings()(*VirtualEndpointExternalPartnerSettingsRequestBuilder) {
    return NewVirtualEndpointExternalPartnerSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FrontLineServicePlans provides operations to manage the frontLineServicePlans property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointFrontLineServicePlansRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) FrontLineServicePlans()(*VirtualEndpointFrontLineServicePlansRequestBuilder) {
    return NewVirtualEndpointFrontLineServicePlansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GalleryImages provides operations to manage the galleryImages property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointGalleryImagesRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) GalleryImages()(*VirtualEndpointGalleryImagesRequestBuilder) {
    return NewVirtualEndpointGalleryImagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get virtualEndpoint from deviceManagement
// returns a VirtualEndpointable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEndpointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable), nil
}
// GetEffectivePermissions provides operations to call the getEffectivePermissions method.
// returns a *VirtualEndpointGetEffectivePermissionsRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) GetEffectivePermissions()(*VirtualEndpointGetEffectivePermissionsRequestBuilder) {
    return NewVirtualEndpointGetEffectivePermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OnPremisesConnections provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointOnPremisesConnectionsRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) OnPremisesConnections()(*VirtualEndpointOnPremisesConnectionsRequestBuilder) {
    return NewVirtualEndpointOnPremisesConnectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OrganizationSettings provides operations to manage the organizationSettings property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointOrganizationSettingsRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) OrganizationSettings()(*VirtualEndpointOrganizationSettingsRequestBuilder) {
    return NewVirtualEndpointOrganizationSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property virtualEndpoint in deviceManagement
// returns a VirtualEndpointable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, requestConfiguration *VirtualEndpointRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateVirtualEndpointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable), nil
}
// ProvisioningPolicies provides operations to manage the provisioningPolicies property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointProvisioningPoliciesRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) ProvisioningPolicies()(*VirtualEndpointProvisioningPoliciesRequestBuilder) {
    return NewVirtualEndpointProvisioningPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reports provides operations to manage the reports property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointReportsRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) Reports()(*VirtualEndpointReportsRequestBuilder) {
    return NewVirtualEndpointReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveScopedPermissions provides operations to call the retrieveScopedPermissions method.
// returns a *VirtualEndpointRetrieveScopedPermissionsRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) RetrieveScopedPermissions()(*VirtualEndpointRetrieveScopedPermissionsRequestBuilder) {
    return NewVirtualEndpointRetrieveScopedPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveTenantEncryptionSetting provides operations to call the retrieveTenantEncryptionSetting method.
// returns a *VirtualEndpointRetrieveTenantEncryptionSettingRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) RetrieveTenantEncryptionSetting()(*VirtualEndpointRetrieveTenantEncryptionSettingRequestBuilder) {
    return NewVirtualEndpointRetrieveTenantEncryptionSettingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServicePlans provides operations to manage the servicePlans property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointServicePlansRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) ServicePlans()(*VirtualEndpointServicePlansRequestBuilder) {
    return NewVirtualEndpointServicePlansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharedUseServicePlans provides operations to manage the sharedUseServicePlans property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointSharedUseServicePlansRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) SharedUseServicePlans()(*VirtualEndpointSharedUseServicePlansRequestBuilder) {
    return NewVirtualEndpointSharedUseServicePlansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Snapshots provides operations to manage the snapshots property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointSnapshotsRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) Snapshots()(*VirtualEndpointSnapshotsRequestBuilder) {
    return NewVirtualEndpointSnapshotsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SupportedRegions provides operations to manage the supportedRegions property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointSupportedRegionsRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) SupportedRegions()(*VirtualEndpointSupportedRegionsRequestBuilder) {
    return NewVirtualEndpointSupportedRegionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property virtualEndpoint for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualEndpointRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get virtualEndpoint from deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualEndpointRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property virtualEndpoint in deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualEndpointRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, requestConfiguration *VirtualEndpointRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// UserSettings provides operations to manage the userSettings property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualEndpointUserSettingsRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) UserSettings()(*VirtualEndpointUserSettingsRequestBuilder) {
    return NewVirtualEndpointUserSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEndpointRequestBuilder when successful
func (m *VirtualEndpointRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointRequestBuilder) {
    return NewVirtualEndpointRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
