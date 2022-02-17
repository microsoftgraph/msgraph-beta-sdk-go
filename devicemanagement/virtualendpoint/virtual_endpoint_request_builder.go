package virtualendpoint

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i037f17d0e70c36a0446de251c289230e77400d9881a7f0256320f00dfe653b5f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/serviceplans"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i13cfcdf74101f5bb07fbf44e4337504d9adaa93a962b4588144742689f3e6c49 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/galleryimages"
    i1529c4b52ff4b5b9e3d21861baf5f70ff01e22c7ce170e59bf3ce66277506553 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs"
    i1696d9b0fb89b84debc07dd2077aff8e4f8ff2b028c14e5a9b3d5139c682a3f3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/deviceimages"
    i23a63a02c83d28d72ce351404f7169106e8b8114b7d9d0ca4052db2a624e6273 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/usersettings"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5b2233e79dc3147871fe57df3812d50c932074cfced41065bb12f77e2edec21e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/auditevents"
    i641f590b75b18bc875786f11f8c01028ef7a86284da618ffc6636c5f70071b7c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/geteffectivepermissions"
    ic615011cfc05f9e79dc505191df45c19871dc563dfe317f14f5eb7e21000fc23 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/onpremisesconnections"
    id4ffcadce3c6a0158b09f7c80fb182d7adea4c0b1e976947864d2f4a6310302e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/organizationsettings"
    idd9e805dd9b76881e3c848aa7cbb09327538706813b2808dfeb729d75afef079 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/supportedregions"
    ie9548ed1e88afc5117c59a732917ab8c23f48e727476fa485002e01d90898ce5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/snapshots"
    if4100d044fd5ea8af7d00890548b33d7f96329a5a6d63f806c3f7705cafc3c7b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/provisioningpolicies"
    i41c3478ec8e1cf11d4aa0d810af05da1184e70e2f0c4b1acd5144a495b556b79 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/serviceplans/item"
    i471da8905b96f5687f46d1c5dc0de1420e0c35b35374f400b7618e2937b4bf34 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/onpremisesconnections/item"
    i57a7497f8ee90067ee6f4aa2a9b6074c547470907dd198f9de96f123682ddda0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/auditevents/item"
    i6c86691f21aa3210591c58f278b27fee3b5fb3a866b020940f913e3f9a20a060 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/provisioningpolicies/item"
    i6cd22011575c099f23c4a4c1302ad841306148c11914cbf14eb33d52739607c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/snapshots/item"
    i91bd1ab7316bb059e22edb6c7fdf042441efd0af722f5ea856fc4be260402497 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/supportedregions/item"
    i940a1af81d34e1d5a5cfef93acb216fcbeba9c4d02700ce3000fd8f0588fd162 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/usersettings/item"
    ia3bcb5abfa0df15f63e2a6fa1f06bcbd43b05fc488a36754660b8c694246fe0b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/deviceimages/item"
    ia7d6e357762433a03086487f037283652587e5519e4f5513185ee9e97e7426e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/galleryimages/item"
    ic14d63ed61f1b60b506fcd81a146444ab96691619c2ac32f274b0793b622f23f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item"
)

// VirtualEndpointRequestBuilder builds and executes requests for operations under \deviceManagement\virtualEndpoint
type VirtualEndpointRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// VirtualEndpointRequestBuilderDeleteOptions options for Delete
type VirtualEndpointRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// VirtualEndpointRequestBuilderGetOptions options for Get
type VirtualEndpointRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *VirtualEndpointRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// VirtualEndpointRequestBuilderGetQueryParameters get virtualEndpoint from deviceManagement
type VirtualEndpointRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// VirtualEndpointRequestBuilderPatchOptions options for Patch
type VirtualEndpointRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.VirtualEndpoint;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *VirtualEndpointRequestBuilder) AuditEvents()(*i5b2233e79dc3147871fe57df3812d50c932074cfced41065bb12f77e2edec21e.AuditEventsRequestBuilder) {
    return i5b2233e79dc3147871fe57df3812d50c932074cfced41065bb12f77e2edec21e.NewAuditEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuditEventsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.auditEvents.item collection
func (m *VirtualEndpointRequestBuilder) AuditEventsById(id string)(*i57a7497f8ee90067ee6f4aa2a9b6074c547470907dd198f9de96f123682ddda0.CloudPcAuditEventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcAuditEvent_id"] = id
    }
    return i57a7497f8ee90067ee6f4aa2a9b6074c547470907dd198f9de96f123682ddda0.NewCloudPcAuditEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) CloudPCs()(*i1529c4b52ff4b5b9e3d21861baf5f70ff01e22c7ce170e59bf3ce66277506553.CloudPCsRequestBuilder) {
    return i1529c4b52ff4b5b9e3d21861baf5f70ff01e22c7ce170e59bf3ce66277506553.NewCloudPCsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.cloudPCs.item collection
func (m *VirtualEndpointRequestBuilder) CloudPCsById(id string)(*ic14d63ed61f1b60b506fcd81a146444ab96691619c2ac32f274b0793b622f23f.CloudPCRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPC_id"] = id
    }
    return ic14d63ed61f1b60b506fcd81a146444ab96691619c2ac32f274b0793b622f23f.NewCloudPCRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewVirtualEndpointRequestBuilderInternal instantiates a new VirtualEndpointRequestBuilder and sets the default values.
func NewVirtualEndpointRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*VirtualEndpointRequestBuilder) {
    m := &VirtualEndpointRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewVirtualEndpointRequestBuilder instantiates a new VirtualEndpointRequestBuilder and sets the default values.
func NewVirtualEndpointRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*VirtualEndpointRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property virtualEndpoint for deviceManagement
func (m *VirtualEndpointRequestBuilder) CreateDeleteRequestInformation(options *VirtualEndpointRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get virtualEndpoint from deviceManagement
func (m *VirtualEndpointRequestBuilder) CreateGetRequestInformation(options *VirtualEndpointRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property virtualEndpoint in deviceManagement
func (m *VirtualEndpointRequestBuilder) CreatePatchRequestInformation(options *VirtualEndpointRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property virtualEndpoint for deviceManagement
func (m *VirtualEndpointRequestBuilder) Delete(options *VirtualEndpointRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *VirtualEndpointRequestBuilder) DeviceImages()(*i1696d9b0fb89b84debc07dd2077aff8e4f8ff2b028c14e5a9b3d5139c682a3f3.DeviceImagesRequestBuilder) {
    return i1696d9b0fb89b84debc07dd2077aff8e4f8ff2b028c14e5a9b3d5139c682a3f3.NewDeviceImagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceImagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.deviceImages.item collection
func (m *VirtualEndpointRequestBuilder) DeviceImagesById(id string)(*ia3bcb5abfa0df15f63e2a6fa1f06bcbd43b05fc488a36754660b8c694246fe0b.CloudPcDeviceImageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcDeviceImage_id"] = id
    }
    return ia3bcb5abfa0df15f63e2a6fa1f06bcbd43b05fc488a36754660b8c694246fe0b.NewCloudPcDeviceImageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) GalleryImages()(*i13cfcdf74101f5bb07fbf44e4337504d9adaa93a962b4588144742689f3e6c49.GalleryImagesRequestBuilder) {
    return i13cfcdf74101f5bb07fbf44e4337504d9adaa93a962b4588144742689f3e6c49.NewGalleryImagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GalleryImagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.galleryImages.item collection
func (m *VirtualEndpointRequestBuilder) GalleryImagesById(id string)(*ia7d6e357762433a03086487f037283652587e5519e4f5513185ee9e97e7426e9.CloudPcGalleryImageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcGalleryImage_id"] = id
    }
    return ia7d6e357762433a03086487f037283652587e5519e4f5513185ee9e97e7426e9.NewCloudPcGalleryImageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get virtualEndpoint from deviceManagement
func (m *VirtualEndpointRequestBuilder) Get(options *VirtualEndpointRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.VirtualEndpoint, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewVirtualEndpoint() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.VirtualEndpoint), nil
}
// GetEffectivePermissions builds and executes requests for operations under \deviceManagement\virtualEndpoint\microsoft.graph.getEffectivePermissions()
func (m *VirtualEndpointRequestBuilder) GetEffectivePermissions()(*i641f590b75b18bc875786f11f8c01028ef7a86284da618ffc6636c5f70071b7c.GetEffectivePermissionsRequestBuilder) {
    return i641f590b75b18bc875786f11f8c01028ef7a86284da618ffc6636c5f70071b7c.NewGetEffectivePermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) OnPremisesConnections()(*ic615011cfc05f9e79dc505191df45c19871dc563dfe317f14f5eb7e21000fc23.OnPremisesConnectionsRequestBuilder) {
    return ic615011cfc05f9e79dc505191df45c19871dc563dfe317f14f5eb7e21000fc23.NewOnPremisesConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OnPremisesConnectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.onPremisesConnections.item collection
func (m *VirtualEndpointRequestBuilder) OnPremisesConnectionsById(id string)(*i471da8905b96f5687f46d1c5dc0de1420e0c35b35374f400b7618e2937b4bf34.CloudPcOnPremisesConnectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcOnPremisesConnection_id"] = id
    }
    return i471da8905b96f5687f46d1c5dc0de1420e0c35b35374f400b7618e2937b4bf34.NewCloudPcOnPremisesConnectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) OrganizationSettings()(*id4ffcadce3c6a0158b09f7c80fb182d7adea4c0b1e976947864d2f4a6310302e.OrganizationSettingsRequestBuilder) {
    return id4ffcadce3c6a0158b09f7c80fb182d7adea4c0b1e976947864d2f4a6310302e.NewOrganizationSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property virtualEndpoint in deviceManagement
func (m *VirtualEndpointRequestBuilder) Patch(options *VirtualEndpointRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *VirtualEndpointRequestBuilder) ProvisioningPolicies()(*if4100d044fd5ea8af7d00890548b33d7f96329a5a6d63f806c3f7705cafc3c7b.ProvisioningPoliciesRequestBuilder) {
    return if4100d044fd5ea8af7d00890548b33d7f96329a5a6d63f806c3f7705cafc3c7b.NewProvisioningPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisioningPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.provisioningPolicies.item collection
func (m *VirtualEndpointRequestBuilder) ProvisioningPoliciesById(id string)(*i6c86691f21aa3210591c58f278b27fee3b5fb3a866b020940f913e3f9a20a060.CloudPcProvisioningPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcProvisioningPolicy_id"] = id
    }
    return i6c86691f21aa3210591c58f278b27fee3b5fb3a866b020940f913e3f9a20a060.NewCloudPcProvisioningPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) ServicePlans()(*i037f17d0e70c36a0446de251c289230e77400d9881a7f0256320f00dfe653b5f.ServicePlansRequestBuilder) {
    return i037f17d0e70c36a0446de251c289230e77400d9881a7f0256320f00dfe653b5f.NewServicePlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePlansById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.servicePlans.item collection
func (m *VirtualEndpointRequestBuilder) ServicePlansById(id string)(*i41c3478ec8e1cf11d4aa0d810af05da1184e70e2f0c4b1acd5144a495b556b79.CloudPcServicePlanRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcServicePlan_id"] = id
    }
    return i41c3478ec8e1cf11d4aa0d810af05da1184e70e2f0c4b1acd5144a495b556b79.NewCloudPcServicePlanRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) Snapshots()(*ie9548ed1e88afc5117c59a732917ab8c23f48e727476fa485002e01d90898ce5.SnapshotsRequestBuilder) {
    return ie9548ed1e88afc5117c59a732917ab8c23f48e727476fa485002e01d90898ce5.NewSnapshotsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SnapshotsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.snapshots.item collection
func (m *VirtualEndpointRequestBuilder) SnapshotsById(id string)(*i6cd22011575c099f23c4a4c1302ad841306148c11914cbf14eb33d52739607c2.CloudPcSnapshotRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcSnapshot_id"] = id
    }
    return i6cd22011575c099f23c4a4c1302ad841306148c11914cbf14eb33d52739607c2.NewCloudPcSnapshotRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) SupportedRegions()(*idd9e805dd9b76881e3c848aa7cbb09327538706813b2808dfeb729d75afef079.SupportedRegionsRequestBuilder) {
    return idd9e805dd9b76881e3c848aa7cbb09327538706813b2808dfeb729d75afef079.NewSupportedRegionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SupportedRegionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.supportedRegions.item collection
func (m *VirtualEndpointRequestBuilder) SupportedRegionsById(id string)(*i91bd1ab7316bb059e22edb6c7fdf042441efd0af722f5ea856fc4be260402497.CloudPcSupportedRegionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcSupportedRegion_id"] = id
    }
    return i91bd1ab7316bb059e22edb6c7fdf042441efd0af722f5ea856fc4be260402497.NewCloudPcSupportedRegionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *VirtualEndpointRequestBuilder) UserSettings()(*i23a63a02c83d28d72ce351404f7169106e8b8114b7d9d0ca4052db2a624e6273.UserSettingsRequestBuilder) {
    return i23a63a02c83d28d72ce351404f7169106e8b8114b7d9d0ca4052db2a624e6273.NewUserSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.virtualEndpoint.userSettings.item collection
func (m *VirtualEndpointRequestBuilder) UserSettingsById(id string)(*i940a1af81d34e1d5a5cfef93acb216fcbeba9c4d02700ce3000fd8f0588fd162.CloudPcUserSettingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcUserSetting_id"] = id
    }
    return i940a1af81d34e1d5a5cfef93acb216fcbeba9c4d02700ce3000fd8f0588fd162.NewCloudPcUserSettingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
