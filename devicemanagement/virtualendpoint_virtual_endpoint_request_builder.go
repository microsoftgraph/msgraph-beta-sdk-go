package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointVirtualEndpointRequestBuilder provides operations to manage the virtualEndpoint property of the microsoft.graph.deviceManagement entity.
type VirtualendpointVirtualEndpointRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointVirtualEndpointRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointVirtualEndpointRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualendpointVirtualEndpointRequestBuilderGetQueryParameters get virtualEndpoint from deviceManagement
type VirtualendpointVirtualEndpointRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualendpointVirtualEndpointRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointVirtualEndpointRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointVirtualEndpointRequestBuilderGetQueryParameters
}
// VirtualendpointVirtualEndpointRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointVirtualEndpointRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointAuditeventsAuditEventsRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) AuditEvents()(*VirtualendpointAuditeventsAuditEventsRequestBuilder) {
    return NewVirtualendpointAuditeventsAuditEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BulkActions provides operations to manage the bulkActions property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointBulkactionsBulkActionsRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) BulkActions()(*VirtualendpointBulkactionsBulkActionsRequestBuilder) {
    return NewVirtualendpointBulkactionsBulkActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudPCs provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointCloudpcsCloudPCsRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) CloudPCs()(*VirtualendpointCloudpcsCloudPCsRequestBuilder) {
    return NewVirtualendpointCloudpcsCloudPCsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualendpointVirtualEndpointRequestBuilderInternal instantiates a new VirtualendpointVirtualEndpointRequestBuilder and sets the default values.
func NewVirtualendpointVirtualEndpointRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointVirtualEndpointRequestBuilder) {
    m := &VirtualendpointVirtualEndpointRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualendpointVirtualEndpointRequestBuilder instantiates a new VirtualendpointVirtualEndpointRequestBuilder and sets the default values.
func NewVirtualendpointVirtualEndpointRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointVirtualEndpointRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointVirtualEndpointRequestBuilderInternal(urlParams, requestAdapter)
}
// CrossCloudGovernmentOrganizationMapping provides operations to manage the crossCloudGovernmentOrganizationMapping property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) CrossCloudGovernmentOrganizationMapping()(*VirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilder) {
    return NewVirtualendpointCrosscloudgovernmentorganizationmappingCrossCloudGovernmentOrganizationMappingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property virtualEndpoint for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointVirtualEndpointRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualendpointVirtualEndpointRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *VirtualendpointDeviceimagesDeviceImagesRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) DeviceImages()(*VirtualendpointDeviceimagesDeviceImagesRequestBuilder) {
    return NewVirtualendpointDeviceimagesDeviceImagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExternalPartnerSettings provides operations to manage the externalPartnerSettings property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointExternalpartnersettingsExternalPartnerSettingsRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) ExternalPartnerSettings()(*VirtualendpointExternalpartnersettingsExternalPartnerSettingsRequestBuilder) {
    return NewVirtualendpointExternalpartnersettingsExternalPartnerSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FrontLineServicePlans provides operations to manage the frontLineServicePlans property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) FrontLineServicePlans()(*VirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilder) {
    return NewVirtualendpointFrontlineserviceplansFrontLineServicePlansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GalleryImages provides operations to manage the galleryImages property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointGalleryimagesGalleryImagesRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) GalleryImages()(*VirtualendpointGalleryimagesGalleryImagesRequestBuilder) {
    return NewVirtualendpointGalleryimagesGalleryImagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get virtualEndpoint from deviceManagement
// returns a VirtualEndpointable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointVirtualEndpointRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointVirtualEndpointRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, error) {
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
// returns a *VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) GetEffectivePermissions()(*VirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilder) {
    return NewVirtualendpointGeteffectivepermissionsGetEffectivePermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OnPremisesConnections provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) OnPremisesConnections()(*VirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilder) {
    return NewVirtualendpointOnpremisesconnectionsOnPremisesConnectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OrganizationSettings provides operations to manage the organizationSettings property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointOrganizationsettingsOrganizationSettingsRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) OrganizationSettings()(*VirtualendpointOrganizationsettingsOrganizationSettingsRequestBuilder) {
    return NewVirtualendpointOrganizationsettingsOrganizationSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property virtualEndpoint in deviceManagement
// returns a VirtualEndpointable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointVirtualEndpointRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, requestConfiguration *VirtualendpointVirtualEndpointRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, error) {
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
// returns a *VirtualendpointProvisioningpoliciesProvisioningPoliciesRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) ProvisioningPolicies()(*VirtualendpointProvisioningpoliciesProvisioningPoliciesRequestBuilder) {
    return NewVirtualendpointProvisioningpoliciesProvisioningPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reports provides operations to manage the reports property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointReportsRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) Reports()(*VirtualendpointReportsRequestBuilder) {
    return NewVirtualendpointReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveScopedPermissions provides operations to call the retrieveScopedPermissions method.
// returns a *VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) RetrieveScopedPermissions()(*VirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilder) {
    return NewVirtualendpointRetrievescopedpermissionsRetrieveScopedPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveTenantEncryptionSetting provides operations to call the retrieveTenantEncryptionSetting method.
// returns a *VirtualendpointRetrievetenantencryptionsettingRetrieveTenantEncryptionSettingRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) RetrieveTenantEncryptionSetting()(*VirtualendpointRetrievetenantencryptionsettingRetrieveTenantEncryptionSettingRequestBuilder) {
    return NewVirtualendpointRetrievetenantencryptionsettingRetrieveTenantEncryptionSettingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServicePlans provides operations to manage the servicePlans property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointServiceplansServicePlansRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) ServicePlans()(*VirtualendpointServiceplansServicePlansRequestBuilder) {
    return NewVirtualendpointServiceplansServicePlansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharedUseServicePlans provides operations to manage the sharedUseServicePlans property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointShareduseserviceplansSharedUseServicePlansRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) SharedUseServicePlans()(*VirtualendpointShareduseserviceplansSharedUseServicePlansRequestBuilder) {
    return NewVirtualendpointShareduseserviceplansSharedUseServicePlansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Snapshots provides operations to manage the snapshots property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointSnapshotsRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) Snapshots()(*VirtualendpointSnapshotsRequestBuilder) {
    return NewVirtualendpointSnapshotsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SupportedRegions provides operations to manage the supportedRegions property of the microsoft.graph.virtualEndpoint entity.
// returns a *VirtualendpointSupportedregionsSupportedRegionsRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) SupportedRegions()(*VirtualendpointSupportedregionsSupportedRegionsRequestBuilder) {
    return NewVirtualendpointSupportedregionsSupportedRegionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property virtualEndpoint for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointVirtualEndpointRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *VirtualendpointVirtualEndpointRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointVirtualEndpointRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *VirtualendpointVirtualEndpointRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, requestConfiguration *VirtualendpointVirtualEndpointRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointUsersettingsUserSettingsRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) UserSettings()(*VirtualendpointUsersettingsUserSettingsRequestBuilder) {
    return NewVirtualendpointUsersettingsUserSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointVirtualEndpointRequestBuilder when successful
func (m *VirtualendpointVirtualEndpointRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointVirtualEndpointRequestBuilder) {
    return NewVirtualendpointVirtualEndpointRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
