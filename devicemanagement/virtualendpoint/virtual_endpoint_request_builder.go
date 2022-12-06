package virtualendpoint

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i037f17d0e70c36a0446de251c289230e77400d9881a7f0256320f00dfe653b5f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/serviceplans"
    i13cfcdf74101f5bb07fbf44e4337504d9adaa93a962b4588144742689f3e6c49 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/galleryimages"
    i148fde72b8eda8f73da4ff581d59b20a648d5f31db55fbc98cebfc9fba002281 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/shareduseserviceplans"
    i1529c4b52ff4b5b9e3d21861baf5f70ff01e22c7ce170e59bf3ce66277506553 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs"
    i1696d9b0fb89b84debc07dd2077aff8e4f8ff2b028c14e5a9b3d5139c682a3f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/deviceimages"
    i23a63a02c83d28d72ce351404f7169106e8b8114b7d9d0ca4052db2a624e6273 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/usersettings"
    i5b2233e79dc3147871fe57df3812d50c932074cfced41065bb12f77e2edec21e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/auditevents"
    i641f590b75b18bc875786f11f8c01028ef7a86284da618ffc6636c5f70071b7c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/geteffectivepermissions"
    i7797f67c8b714660406fd3191cfbb22f837ba6075cc7a38889ec3a97de75f7d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/reports"
    i8d3de0aa309e67af3900939f40971eff6bb1f2ef4f7cece48df22e29e1c1a98e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/externalpartnersettings"
    ib721ddefcd2a7684894eea8d4d1e9352a78b6805ab7b3f6f328db56575d44444 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/crosscloudgovernmentorganizationmapping"
    ic615011cfc05f9e79dc505191df45c19871dc563dfe317f14f5eb7e21000fc23 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/onpremisesconnections"
    id4ffcadce3c6a0158b09f7c80fb182d7adea4c0b1e976947864d2f4a6310302e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/organizationsettings"
    idd9e805dd9b76881e3c848aa7cbb09327538706813b2808dfeb729d75afef079 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/supportedregions"
    ie9548ed1e88afc5117c59a732917ab8c23f48e727476fa485002e01d90898ce5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/snapshots"
    if4100d044fd5ea8af7d00890548b33d7f96329a5a6d63f806c3f7705cafc3c7b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/provisioningpolicies"
    i41c3478ec8e1cf11d4aa0d810af05da1184e70e2f0c4b1acd5144a495b556b79 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/serviceplans/item"
    i471da8905b96f5687f46d1c5dc0de1420e0c35b35374f400b7618e2937b4bf34 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/onpremisesconnections/item"
    i4a263989074f26231ca14dc310ebe5fe63e611c6c584ae036d74fff02c90ea2d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/shareduseserviceplans/item"
    i57a7497f8ee90067ee6f4aa2a9b6074c547470907dd198f9de96f123682ddda0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/auditevents/item"
    i6c86691f21aa3210591c58f278b27fee3b5fb3a866b020940f913e3f9a20a060 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/provisioningpolicies/item"
    i6cd22011575c099f23c4a4c1302ad841306148c11914cbf14eb33d52739607c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/snapshots/item"
    i91bd1ab7316bb059e22edb6c7fdf042441efd0af722f5ea856fc4be260402497 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/supportedregions/item"
    i940a1af81d34e1d5a5cfef93acb216fcbeba9c4d02700ce3000fd8f0588fd162 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/usersettings/item"
    ia3bcb5abfa0df15f63e2a6fa1f06bcbd43b05fc488a36754660b8c694246fe0b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/deviceimages/item"
    ia7d6e357762433a03086487f037283652587e5519e4f5513185ee9e97e7426e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/galleryimages/item"
    ic14d63ed61f1b60b506fcd81a146444ab96691619c2ac32f274b0793b622f23f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item"
    ief38e332ba0a8567dc33c08d7a4374b2469b756d157cd9dda5252926cfc5109d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/externalpartnersettings/item"
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
    Headers map[string]string
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
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointRequestBuilderGetQueryParameters
}
// VirtualEndpointRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) AuditEvents()(*i5b2233e79dc3147871fe57df3812d50c932074cfced41065bb12f77e2edec21e.AuditEventsRequestBuilder) {
    return i5b2233e79dc3147871fe57df3812d50c932074cfced41065bb12f77e2edec21e.NewAuditEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuditEventsById provides operations to manage the auditEvents property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) AuditEventsById(id string)(*i57a7497f8ee90067ee6f4aa2a9b6074c547470907dd198f9de96f123682ddda0.CloudPcAuditEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcAuditEvent%2Did"] = id
    }
    return i57a7497f8ee90067ee6f4aa2a9b6074c547470907dd198f9de96f123682ddda0.NewCloudPcAuditEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPCs provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) CloudPCs()(*i1529c4b52ff4b5b9e3d21861baf5f70ff01e22c7ce170e59bf3ce66277506553.CloudPCsRequestBuilder) {
    return i1529c4b52ff4b5b9e3d21861baf5f70ff01e22c7ce170e59bf3ce66277506553.NewCloudPCsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCsById provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) CloudPCsById(id string)(*ic14d63ed61f1b60b506fcd81a146444ab96691619c2ac32f274b0793b622f23f.CloudPCItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPC%2Did"] = id
    }
    return ic14d63ed61f1b60b506fcd81a146444ab96691619c2ac32f274b0793b622f23f.NewCloudPCItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
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
func (m *VirtualEndpointRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.VirtualEndpointable, requestConfiguration *VirtualEndpointRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *VirtualEndpointRequestBuilder) CrossCloudGovernmentOrganizationMapping()(*ib721ddefcd2a7684894eea8d4d1e9352a78b6805ab7b3f6f328db56575d44444.CrossCloudGovernmentOrganizationMappingRequestBuilder) {
    return ib721ddefcd2a7684894eea8d4d1e9352a78b6805ab7b3f6f328db56575d44444.NewCrossCloudGovernmentOrganizationMappingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *VirtualEndpointRequestBuilder) DeviceImages()(*i1696d9b0fb89b84debc07dd2077aff8e4f8ff2b028c14e5a9b3d5139c682a3f3.DeviceImagesRequestBuilder) {
    return i1696d9b0fb89b84debc07dd2077aff8e4f8ff2b028c14e5a9b3d5139c682a3f3.NewDeviceImagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceImagesById provides operations to manage the deviceImages property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) DeviceImagesById(id string)(*ia3bcb5abfa0df15f63e2a6fa1f06bcbd43b05fc488a36754660b8c694246fe0b.CloudPcDeviceImageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcDeviceImage%2Did"] = id
    }
    return ia3bcb5abfa0df15f63e2a6fa1f06bcbd43b05fc488a36754660b8c694246fe0b.NewCloudPcDeviceImageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExternalPartnerSettings provides operations to manage the externalPartnerSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) ExternalPartnerSettings()(*i8d3de0aa309e67af3900939f40971eff6bb1f2ef4f7cece48df22e29e1c1a98e.ExternalPartnerSettingsRequestBuilder) {
    return i8d3de0aa309e67af3900939f40971eff6bb1f2ef4f7cece48df22e29e1c1a98e.NewExternalPartnerSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalPartnerSettingsById provides operations to manage the externalPartnerSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) ExternalPartnerSettingsById(id string)(*ief38e332ba0a8567dc33c08d7a4374b2469b756d157cd9dda5252926cfc5109d.CloudPcExternalPartnerSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcExternalPartnerSetting%2Did"] = id
    }
    return ief38e332ba0a8567dc33c08d7a4374b2469b756d157cd9dda5252926cfc5109d.NewCloudPcExternalPartnerSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GalleryImages provides operations to manage the galleryImages property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) GalleryImages()(*i13cfcdf74101f5bb07fbf44e4337504d9adaa93a962b4588144742689f3e6c49.GalleryImagesRequestBuilder) {
    return i13cfcdf74101f5bb07fbf44e4337504d9adaa93a962b4588144742689f3e6c49.NewGalleryImagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GalleryImagesById provides operations to manage the galleryImages property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) GalleryImagesById(id string)(*ia7d6e357762433a03086487f037283652587e5519e4f5513185ee9e97e7426e9.CloudPcGalleryImageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcGalleryImage%2Did"] = id
    }
    return ia7d6e357762433a03086487f037283652587e5519e4f5513185ee9e97e7426e9.NewCloudPcGalleryImageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *VirtualEndpointRequestBuilder) GetEffectivePermissions()(*i641f590b75b18bc875786f11f8c01028ef7a86284da618ffc6636c5f70071b7c.GetEffectivePermissionsRequestBuilder) {
    return i641f590b75b18bc875786f11f8c01028ef7a86284da618ffc6636c5f70071b7c.NewGetEffectivePermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnPremisesConnections provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) OnPremisesConnections()(*ic615011cfc05f9e79dc505191df45c19871dc563dfe317f14f5eb7e21000fc23.OnPremisesConnectionsRequestBuilder) {
    return ic615011cfc05f9e79dc505191df45c19871dc563dfe317f14f5eb7e21000fc23.NewOnPremisesConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnPremisesConnectionsById provides operations to manage the onPremisesConnections property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) OnPremisesConnectionsById(id string)(*i471da8905b96f5687f46d1c5dc0de1420e0c35b35374f400b7618e2937b4bf34.CloudPcOnPremisesConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcOnPremisesConnection%2Did"] = id
    }
    return i471da8905b96f5687f46d1c5dc0de1420e0c35b35374f400b7618e2937b4bf34.NewCloudPcOnPremisesConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OrganizationSettings provides operations to manage the organizationSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) OrganizationSettings()(*id4ffcadce3c6a0158b09f7c80fb182d7adea4c0b1e976947864d2f4a6310302e.OrganizationSettingsRequestBuilder) {
    return id4ffcadce3c6a0158b09f7c80fb182d7adea4c0b1e976947864d2f4a6310302e.NewOrganizationSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *VirtualEndpointRequestBuilder) ProvisioningPolicies()(*if4100d044fd5ea8af7d00890548b33d7f96329a5a6d63f806c3f7705cafc3c7b.ProvisioningPoliciesRequestBuilder) {
    return if4100d044fd5ea8af7d00890548b33d7f96329a5a6d63f806c3f7705cafc3c7b.NewProvisioningPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisioningPoliciesById provides operations to manage the provisioningPolicies property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) ProvisioningPoliciesById(id string)(*i6c86691f21aa3210591c58f278b27fee3b5fb3a866b020940f913e3f9a20a060.CloudPcProvisioningPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcProvisioningPolicy%2Did"] = id
    }
    return i6c86691f21aa3210591c58f278b27fee3b5fb3a866b020940f913e3f9a20a060.NewCloudPcProvisioningPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Reports provides operations to manage the reports property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) Reports()(*i7797f67c8b714660406fd3191cfbb22f837ba6075cc7a38889ec3a97de75f7d4.ReportsRequestBuilder) {
    return i7797f67c8b714660406fd3191cfbb22f837ba6075cc7a38889ec3a97de75f7d4.NewReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePlans provides operations to manage the servicePlans property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) ServicePlans()(*i037f17d0e70c36a0446de251c289230e77400d9881a7f0256320f00dfe653b5f.ServicePlansRequestBuilder) {
    return i037f17d0e70c36a0446de251c289230e77400d9881a7f0256320f00dfe653b5f.NewServicePlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePlansById provides operations to manage the servicePlans property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) ServicePlansById(id string)(*i41c3478ec8e1cf11d4aa0d810af05da1184e70e2f0c4b1acd5144a495b556b79.CloudPcServicePlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcServicePlan%2Did"] = id
    }
    return i41c3478ec8e1cf11d4aa0d810af05da1184e70e2f0c4b1acd5144a495b556b79.NewCloudPcServicePlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SharedUseServicePlans provides operations to manage the sharedUseServicePlans property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) SharedUseServicePlans()(*i148fde72b8eda8f73da4ff581d59b20a648d5f31db55fbc98cebfc9fba002281.SharedUseServicePlansRequestBuilder) {
    return i148fde72b8eda8f73da4ff581d59b20a648d5f31db55fbc98cebfc9fba002281.NewSharedUseServicePlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedUseServicePlansById provides operations to manage the sharedUseServicePlans property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) SharedUseServicePlansById(id string)(*i4a263989074f26231ca14dc310ebe5fe63e611c6c584ae036d74fff02c90ea2d.CloudPcSharedUseServicePlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcSharedUseServicePlan%2Did"] = id
    }
    return i4a263989074f26231ca14dc310ebe5fe63e611c6c584ae036d74fff02c90ea2d.NewCloudPcSharedUseServicePlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Snapshots provides operations to manage the snapshots property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) Snapshots()(*ie9548ed1e88afc5117c59a732917ab8c23f48e727476fa485002e01d90898ce5.SnapshotsRequestBuilder) {
    return ie9548ed1e88afc5117c59a732917ab8c23f48e727476fa485002e01d90898ce5.NewSnapshotsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SnapshotsById provides operations to manage the snapshots property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) SnapshotsById(id string)(*i6cd22011575c099f23c4a4c1302ad841306148c11914cbf14eb33d52739607c2.CloudPcSnapshotItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcSnapshot%2Did"] = id
    }
    return i6cd22011575c099f23c4a4c1302ad841306148c11914cbf14eb33d52739607c2.NewCloudPcSnapshotItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SupportedRegions provides operations to manage the supportedRegions property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) SupportedRegions()(*idd9e805dd9b76881e3c848aa7cbb09327538706813b2808dfeb729d75afef079.SupportedRegionsRequestBuilder) {
    return idd9e805dd9b76881e3c848aa7cbb09327538706813b2808dfeb729d75afef079.NewSupportedRegionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedRegionsById provides operations to manage the supportedRegions property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) SupportedRegionsById(id string)(*i91bd1ab7316bb059e22edb6c7fdf042441efd0af722f5ea856fc4be260402497.CloudPcSupportedRegionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcSupportedRegion%2Did"] = id
    }
    return i91bd1ab7316bb059e22edb6c7fdf042441efd0af722f5ea856fc4be260402497.NewCloudPcSupportedRegionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserSettings provides operations to manage the userSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) UserSettings()(*i23a63a02c83d28d72ce351404f7169106e8b8114b7d9d0ca4052db2a624e6273.UserSettingsRequestBuilder) {
    return i23a63a02c83d28d72ce351404f7169106e8b8114b7d9d0ca4052db2a624e6273.NewUserSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSettingsById provides operations to manage the userSettings property of the microsoft.graph.virtualEndpoint entity.
func (m *VirtualEndpointRequestBuilder) UserSettingsById(id string)(*i940a1af81d34e1d5a5cfef93acb216fcbeba9c4d02700ce3000fd8f0588fd162.CloudPcUserSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcUserSetting%2Did"] = id
    }
    return i940a1af81d34e1d5a5cfef93acb216fcbeba9c4d02700ce3000fd8f0588fd162.NewCloudPcUserSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
