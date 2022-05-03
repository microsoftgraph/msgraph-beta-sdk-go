package deviceappmanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i002d4b08a58acc3cf19dcedd322215597fc792aab29c13c978987157740bb4c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps"
    i102dee350049990db6017a7e5aaec602f046d2eab032b9cfe8fdd20bb889fe71 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/androidmanagedappprotections"
    i1d8bd406d3b1507c2172d70e5990518e80c89c738e5da31e1c4a628d3bff7d8f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappcategories"
    i20583d0a177898187a71dda1d89153f4ea0e3952fb9f921dbe151384949dfd62 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2409240b9ff00edd267b9ce9480582a136074b1a1ef0f1a320e40469879b5b4a "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappstatuses"
    i24941860ea69ec24f14182d511152b68a5da0429ab75a0d6a1e6ccecfd1a155c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/defaultmanagedappprotections"
    i2c0e155b4c97677199c972087f9bc6d38e155b8f14c9d3cfa0eb0c056759608c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/deviceappmanagementtasks"
    i2dad5aeb46d764dd36e179abb8522de62e89f2a53fdcd4b051505fecc5c0f7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/iosmanagedappprotections"
    i39a576817b6f726c83eda563ed7f79852514888b74a8942a87e4fb08eb489a74 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations"
    i3e0f6a6236335b92506dbec1a57163f0d6904827a60904e0159bf0e01304e5c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations"
    i47f8613c686153638e221e0ff110d8a5a6e072e9cf93b668c444d83a782cf6fe "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/enterprisecodesigningcertificates"
    i4ae1191713b2bcf5293d5d9a7f57d7875e1fb8395d4f1ac48c09144198cf2584 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations"
    i506595e5c982d9cbef552a980467df6c4caffc29ce38c31544b7a4a9e91b4ef9 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebookcategories"
    i55154c3923d16b5ed36ced8585e111c41354882664d34667a6e3b97700313735 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/symanteccodesigningcertificate"
    i55b67fda8074ac8f07c7e2339a212683a0fb8f21e6b18b534d0e1c1d0d87cf54 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations"
    i56f77a9c5e1a7ca279d87b52f582d52adbbc4de96f0fcef383fa9ded66e0a4c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/sideloadingkeys"
    i63940f883f70d79c9f036190af3786b785aa38e88077dee0da7d8bdd443ffb95 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsinformationprotectiondeviceregistrations"
    i76773835ad298178375c61fd2bb3db2a00f6d1b79d366e5235fc02cd5a71d5ef "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mdmwindowsinformationprotectionpolicies"
    i869aad9b4b95674254fbc32234cd41bc975fb5a270d41358d0b38c0a515d0e76 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsinformationprotectionwipeactions"
    i8b9340bb6f8a5fae484feebe60a0c77ea793398c1470ee9294c37011a384a585 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies"
    i9b677b07864604eb0a1dc7960a42ddf4914647e3fa83a6d2cfea3e12250a4704 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsmanagementapp"
    ib1493941e0d94862b2a07650ec9a7a40216e313891ffb7e2b502884c5d0796df "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedapppolicies"
    icabc5c2bd51f5daedc915efd55ca836aaaa2a4f1e8745b5ea4f3e393db1797f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/policysets"
    ie802a5a736b1aceb4dbbe40ec170346b4ae36c2165343b37c69767334cdf1f05 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsinformationprotectionpolicies"
    iee0ff8dcd5df6303aaa9e4b0eb5d5437722e178fc10f16c2284bd8702f364eaa "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/vpptokens"
    ief11f22801272e064478cc3e45b3a0f1e63de4a918dab77a8cb1f3924109be70 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/syncmicrosoftstoreforbusinessapps"
    i111d2625f1ad2501f178f2146156ce46548595d9f1f6169f1d43762d4d373d9d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/deviceappmanagementtasks/item"
    i15b8057a97c25aa14e9a4d6fc9def02dce42569f753f7a62f61b392239a5d612 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappstatuses/item"
    i1fa4345424f4e8032c54aadba42924b3c9de89ca46689dc3b1b829ffe729b8f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/iosmanagedappprotections/item"
    i29437e4fbb36224c3fa3e19cdd35ffd707faff0d9b46f5b91bb301fa83c2fc02 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item"
    i31dc45267ed832743f873d9cbf6d2e8ce83262adc3ce83b4698d646470d1061d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedapppolicies/item"
    i37f5ce10204a635850269ccbe65874caffc5e6d79af25b669118274ddd1b167e "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item"
    i3b0fefd7a08f40697f53b6e4e31c5e2c9753528a64f9b232b76ca8f31a4e944e "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/androidmanagedappprotections/item"
    i3b995f6d41023e3b43346c95c0f59fe84cb19f9a05efc8999cbbd88c8e143b9c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/vpptokens/item"
    i40679b07c9d7f966f341b0ef7533a4761cd0fc2dd66e199f9bb9bef68c64cdeb "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mdmwindowsinformationprotectionpolicies/item"
    i492539e1cc50966c62570b7b76059bb701558d85482e3ca95c396dc36c10607a "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/policysets/item"
    i587ff2c64043a8d6a7ba7b62c47e4c87fddaaad70ec6796eb827a7be948a9033 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/defaultmanagedappprotections/item"
    i6b77a717f9ad30ec5ef2cbce04e89b2244f69b8e725f39fb3990beeed77e175d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/sideloadingkeys/item"
    i6d212574cda4bdc3baad6eb00d5c9eaf6d61f31d25ad4cdc63ed678de8d2f959 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebookcategories/item"
    i6ed3d58de1a55f05afd474b93dfc98801a4e594f2d674957216b11e72cef3f3b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/ioslobappprovisioningconfigurations/item"
    i77e081996fdbcfeb6b1bd69d57141d661fd7bda147a8b41a9cfbbeecda840090 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedappregistrations/item"
    i926032b256e689b9a8930cece406118a2156139b79c1b064550fc7bba6a804b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item"
    i9c05ace01a97f06826939943c032786e845da3515dd142da04b21d1ed556d243 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsinformationprotectiondeviceregistrations/item"
    ib4dcb1c692a4cab02e05d87e0707aa4831b962f34cad5157cb939f82a9345efb "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsinformationprotectionwipeactions/item"
    ib84c840989f9ab439ec3ef07f903c5bac878d6066f60bc53dac1f18abcc7ad13 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsinformationprotectionpolicies/item"
    ic25fb9fffd8c02c79bfc42b3e12ad25f3fbaf117526021e4526f4cd527116b00 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/enterprisecodesigningcertificates/item"
    iceecf955eeeaa6e59a7ff5c62bcd7ca93ca56e0d58f15cbe8d764abcf525d4fc "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item"
    iebd0ee945a2658314caea3437be3d7434b12969418fd5a37ff33dec3769a17de "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappcategories/item"
    if980bb6134295700a6ecd4621e106b4faf366915f3f6b4a9decbfa4a90b923df "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileappconfigurations/item"
)

// DeviceAppManagementRequestBuilder provides operations to manage the deviceAppManagement singleton.
type DeviceAppManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementRequestBuilderGetQueryParameters get deviceAppManagement
type DeviceAppManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementRequestBuilderGetQueryParameters
}
// DeviceAppManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AndroidManagedAppProtections the androidManagedAppProtections property
func (m *DeviceAppManagementRequestBuilder) AndroidManagedAppProtections()(*i102dee350049990db6017a7e5aaec602f046d2eab032b9cfe8fdd20bb889fe71.AndroidManagedAppProtectionsRequestBuilder) {
    return i102dee350049990db6017a7e5aaec602f046d2eab032b9cfe8fdd20bb889fe71.NewAndroidManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidManagedAppProtectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.androidManagedAppProtections.item collection
func (m *DeviceAppManagementRequestBuilder) AndroidManagedAppProtectionsById(id string)(*i3b0fefd7a08f40697f53b6e4e31c5e2c9753528a64f9b232b76ca8f31a4e944e.AndroidManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["androidManagedAppProtection%2Did"] = id
    }
    return i3b0fefd7a08f40697f53b6e4e31c5e2c9753528a64f9b232b76ca8f31a4e944e.NewAndroidManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceAppManagementRequestBuilderInternal instantiates a new DeviceAppManagementRequestBuilder and sets the default values.
func NewDeviceAppManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementRequestBuilder) {
    m := &DeviceAppManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementRequestBuilder instantiates a new DeviceAppManagementRequestBuilder and sets the default values.
func NewDeviceAppManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceAppManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable, requestConfiguration *DeviceAppManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DefaultManagedAppProtections the defaultManagedAppProtections property
func (m *DeviceAppManagementRequestBuilder) DefaultManagedAppProtections()(*i24941860ea69ec24f14182d511152b68a5da0429ab75a0d6a1e6ccecfd1a155c.DefaultManagedAppProtectionsRequestBuilder) {
    return i24941860ea69ec24f14182d511152b68a5da0429ab75a0d6a1e6ccecfd1a155c.NewDefaultManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultManagedAppProtectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.defaultManagedAppProtections.item collection
func (m *DeviceAppManagementRequestBuilder) DefaultManagedAppProtectionsById(id string)(*i587ff2c64043a8d6a7ba7b62c47e4c87fddaaad70ec6796eb827a7be948a9033.DefaultManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["defaultManagedAppProtection%2Did"] = id
    }
    return i587ff2c64043a8d6a7ba7b62c47e4c87fddaaad70ec6796eb827a7be948a9033.NewDefaultManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceAppManagementTasks the deviceAppManagementTasks property
func (m *DeviceAppManagementRequestBuilder) DeviceAppManagementTasks()(*i2c0e155b4c97677199c972087f9bc6d38e155b8f14c9d3cfa0eb0c056759608c.DeviceAppManagementTasksRequestBuilder) {
    return i2c0e155b4c97677199c972087f9bc6d38e155b8f14c9d3cfa0eb0c056759608c.NewDeviceAppManagementTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceAppManagementTasksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.deviceAppManagementTasks.item collection
func (m *DeviceAppManagementRequestBuilder) DeviceAppManagementTasksById(id string)(*i111d2625f1ad2501f178f2146156ce46548595d9f1f6169f1d43762d4d373d9d.DeviceAppManagementTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceAppManagementTask%2Did"] = id
    }
    return i111d2625f1ad2501f178f2146156ce46548595d9f1f6169f1d43762d4d373d9d.NewDeviceAppManagementTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// EnterpriseCodeSigningCertificates the enterpriseCodeSigningCertificates property
func (m *DeviceAppManagementRequestBuilder) EnterpriseCodeSigningCertificates()(*i47f8613c686153638e221e0ff110d8a5a6e072e9cf93b668c444d83a782cf6fe.EnterpriseCodeSigningCertificatesRequestBuilder) {
    return i47f8613c686153638e221e0ff110d8a5a6e072e9cf93b668c444d83a782cf6fe.NewEnterpriseCodeSigningCertificatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnterpriseCodeSigningCertificatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.enterpriseCodeSigningCertificates.item collection
func (m *DeviceAppManagementRequestBuilder) EnterpriseCodeSigningCertificatesById(id string)(*ic25fb9fffd8c02c79bfc42b3e12ad25f3fbaf117526021e4526f4cd527116b00.EnterpriseCodeSigningCertificateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["enterpriseCodeSigningCertificate%2Did"] = id
    }
    return ic25fb9fffd8c02c79bfc42b3e12ad25f3fbaf117526021e4526f4cd527116b00.NewEnterpriseCodeSigningCertificateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DeviceAppManagementRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAppManagementFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable), nil
}
// IosLobAppProvisioningConfigurations the iosLobAppProvisioningConfigurations property
func (m *DeviceAppManagementRequestBuilder) IosLobAppProvisioningConfigurations()(*i4ae1191713b2bcf5293d5d9a7f57d7875e1fb8395d4f1ac48c09144198cf2584.IosLobAppProvisioningConfigurationsRequestBuilder) {
    return i4ae1191713b2bcf5293d5d9a7f57d7875e1fb8395d4f1ac48c09144198cf2584.NewIosLobAppProvisioningConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IosLobAppProvisioningConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosLobAppProvisioningConfigurations.item collection
func (m *DeviceAppManagementRequestBuilder) IosLobAppProvisioningConfigurationsById(id string)(*i6ed3d58de1a55f05afd474b93dfc98801a4e594f2d674957216b11e72cef3f3b.IosLobAppProvisioningConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosLobAppProvisioningConfiguration%2Did"] = id
    }
    return i6ed3d58de1a55f05afd474b93dfc98801a4e594f2d674957216b11e72cef3f3b.NewIosLobAppProvisioningConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IosManagedAppProtections the iosManagedAppProtections property
func (m *DeviceAppManagementRequestBuilder) IosManagedAppProtections()(*i2dad5aeb46d764dd36e179abb8522de62e89f2a53fdcd4b051505fecc5c0f7bc.IosManagedAppProtectionsRequestBuilder) {
    return i2dad5aeb46d764dd36e179abb8522de62e89f2a53fdcd4b051505fecc5c0f7bc.NewIosManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IosManagedAppProtectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.iosManagedAppProtections.item collection
func (m *DeviceAppManagementRequestBuilder) IosManagedAppProtectionsById(id string)(*i1fa4345424f4e8032c54aadba42924b3c9de89ca46689dc3b1b829ffe729b8f6.IosManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosManagedAppProtection%2Did"] = id
    }
    return i1fa4345424f4e8032c54aadba42924b3c9de89ca46689dc3b1b829ffe729b8f6.NewIosManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppPolicies the managedAppPolicies property
func (m *DeviceAppManagementRequestBuilder) ManagedAppPolicies()(*ib1493941e0d94862b2a07650ec9a7a40216e313891ffb7e2b502884c5d0796df.ManagedAppPoliciesRequestBuilder) {
    return ib1493941e0d94862b2a07650ec9a7a40216e313891ffb7e2b502884c5d0796df.NewManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedAppPolicies.item collection
func (m *DeviceAppManagementRequestBuilder) ManagedAppPoliciesById(id string)(*i31dc45267ed832743f873d9cbf6d2e8ce83262adc3ce83b4698d646470d1061d.ManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy%2Did"] = id
    }
    return i31dc45267ed832743f873d9cbf6d2e8ce83262adc3ce83b4698d646470d1061d.NewManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppRegistrations the managedAppRegistrations property
func (m *DeviceAppManagementRequestBuilder) ManagedAppRegistrations()(*i3e0f6a6236335b92506dbec1a57163f0d6904827a60904e0159bf0e01304e5c1.ManagedAppRegistrationsRequestBuilder) {
    return i3e0f6a6236335b92506dbec1a57163f0d6904827a60904e0159bf0e01304e5c1.NewManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppRegistrationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedAppRegistrations.item collection
func (m *DeviceAppManagementRequestBuilder) ManagedAppRegistrationsById(id string)(*i77e081996fdbcfeb6b1bd69d57141d661fd7bda147a8b41a9cfbbeecda840090.ManagedAppRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppRegistration%2Did"] = id
    }
    return i77e081996fdbcfeb6b1bd69d57141d661fd7bda147a8b41a9cfbbeecda840090.NewManagedAppRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppStatuses the managedAppStatuses property
func (m *DeviceAppManagementRequestBuilder) ManagedAppStatuses()(*i2409240b9ff00edd267b9ce9480582a136074b1a1ef0f1a320e40469879b5b4a.ManagedAppStatusesRequestBuilder) {
    return i2409240b9ff00edd267b9ce9480582a136074b1a1ef0f1a320e40469879b5b4a.NewManagedAppStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedAppStatuses.item collection
func (m *DeviceAppManagementRequestBuilder) ManagedAppStatusesById(id string)(*i15b8057a97c25aa14e9a4d6fc9def02dce42569f753f7a62f61b392239a5d612.ManagedAppStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppStatus%2Did"] = id
    }
    return i15b8057a97c25aa14e9a4d6fc9def02dce42569f753f7a62f61b392239a5d612.NewManagedAppStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedEBookCategories the managedEBookCategories property
func (m *DeviceAppManagementRequestBuilder) ManagedEBookCategories()(*i506595e5c982d9cbef552a980467df6c4caffc29ce38c31544b7a4a9e91b4ef9.ManagedEBookCategoriesRequestBuilder) {
    return i506595e5c982d9cbef552a980467df6c4caffc29ce38c31544b7a4a9e91b4ef9.NewManagedEBookCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedEBookCategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedEBookCategories.item collection
func (m *DeviceAppManagementRequestBuilder) ManagedEBookCategoriesById(id string)(*i6d212574cda4bdc3baad6eb00d5c9eaf6d61f31d25ad4cdc63ed678de8d2f959.ManagedEBookCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBookCategory%2Did"] = id
    }
    return i6d212574cda4bdc3baad6eb00d5c9eaf6d61f31d25ad4cdc63ed678de8d2f959.NewManagedEBookCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedEBooks the managedEBooks property
func (m *DeviceAppManagementRequestBuilder) ManagedEBooks()(*i20583d0a177898187a71dda1d89153f4ea0e3952fb9f921dbe151384949dfd62.ManagedEBooksRequestBuilder) {
    return i20583d0a177898187a71dda1d89153f4ea0e3952fb9f921dbe151384949dfd62.NewManagedEBooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedEBooksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedEBooks.item collection
func (m *DeviceAppManagementRequestBuilder) ManagedEBooksById(id string)(*i926032b256e689b9a8930cece406118a2156139b79c1b064550fc7bba6a804b8.ManagedEBookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBook%2Did"] = id
    }
    return i926032b256e689b9a8930cece406118a2156139b79c1b064550fc7bba6a804b8.NewManagedEBookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MdmWindowsInformationProtectionPolicies the mdmWindowsInformationProtectionPolicies property
func (m *DeviceAppManagementRequestBuilder) MdmWindowsInformationProtectionPolicies()(*i76773835ad298178375c61fd2bb3db2a00f6d1b79d366e5235fc02cd5a71d5ef.MdmWindowsInformationProtectionPoliciesRequestBuilder) {
    return i76773835ad298178375c61fd2bb3db2a00f6d1b79d366e5235fc02cd5a71d5ef.NewMdmWindowsInformationProtectionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MdmWindowsInformationProtectionPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mdmWindowsInformationProtectionPolicies.item collection
func (m *DeviceAppManagementRequestBuilder) MdmWindowsInformationProtectionPoliciesById(id string)(*i40679b07c9d7f966f341b0ef7533a4761cd0fc2dd66e199f9bb9bef68c64cdeb.MdmWindowsInformationProtectionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mdmWindowsInformationProtectionPolicy%2Did"] = id
    }
    return i40679b07c9d7f966f341b0ef7533a4761cd0fc2dd66e199f9bb9bef68c64cdeb.NewMdmWindowsInformationProtectionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppCategories the mobileAppCategories property
func (m *DeviceAppManagementRequestBuilder) MobileAppCategories()(*i1d8bd406d3b1507c2172d70e5990518e80c89c738e5da31e1c4a628d3bff7d8f.MobileAppCategoriesRequestBuilder) {
    return i1d8bd406d3b1507c2172d70e5990518e80c89c738e5da31e1c4a628d3bff7d8f.NewMobileAppCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppCategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileAppCategories.item collection
func (m *DeviceAppManagementRequestBuilder) MobileAppCategoriesById(id string)(*iebd0ee945a2658314caea3437be3d7434b12969418fd5a37ff33dec3769a17de.MobileAppCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppCategory%2Did"] = id
    }
    return iebd0ee945a2658314caea3437be3d7434b12969418fd5a37ff33dec3769a17de.NewMobileAppCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppConfigurations the mobileAppConfigurations property
func (m *DeviceAppManagementRequestBuilder) MobileAppConfigurations()(*i55b67fda8074ac8f07c7e2339a212683a0fb8f21e6b18b534d0e1c1d0d87cf54.MobileAppConfigurationsRequestBuilder) {
    return i55b67fda8074ac8f07c7e2339a212683a0fb8f21e6b18b534d0e1c1d0d87cf54.NewMobileAppConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileAppConfigurations.item collection
func (m *DeviceAppManagementRequestBuilder) MobileAppConfigurationsById(id string)(*if980bb6134295700a6ecd4621e106b4faf366915f3f6b4a9decbfa4a90b923df.ManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfiguration%2Did"] = id
    }
    return if980bb6134295700a6ecd4621e106b4faf366915f3f6b4a9decbfa4a90b923df.NewManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileApps the mobileApps property
func (m *DeviceAppManagementRequestBuilder) MobileApps()(*i002d4b08a58acc3cf19dcedd322215597fc792aab29c13c978987157740bb4c9.MobileAppsRequestBuilder) {
    return i002d4b08a58acc3cf19dcedd322215597fc792aab29c13c978987157740bb4c9.NewMobileAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileApps.item collection
func (m *DeviceAppManagementRequestBuilder) MobileAppsById(id string)(*i29437e4fbb36224c3fa3e19cdd35ffd707faff0d9b46f5b91bb301fa83c2fc02.MobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileApp%2Did"] = id
    }
    return i29437e4fbb36224c3fa3e19cdd35ffd707faff0d9b46f5b91bb301fa83c2fc02.NewMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAppManagementable, requestConfiguration *DeviceAppManagementRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// PolicySets the policySets property
func (m *DeviceAppManagementRequestBuilder) PolicySets()(*icabc5c2bd51f5daedc915efd55ca836aaaa2a4f1e8745b5ea4f3e393db1797f1.PolicySetsRequestBuilder) {
    return icabc5c2bd51f5daedc915efd55ca836aaaa2a4f1e8745b5ea4f3e393db1797f1.NewPolicySetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PolicySetsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.policySets.item collection
func (m *DeviceAppManagementRequestBuilder) PolicySetsById(id string)(*i492539e1cc50966c62570b7b76059bb701558d85482e3ca95c396dc36c10607a.PolicySetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["policySet%2Did"] = id
    }
    return i492539e1cc50966c62570b7b76059bb701558d85482e3ca95c396dc36c10607a.NewPolicySetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SideLoadingKeys the sideLoadingKeys property
func (m *DeviceAppManagementRequestBuilder) SideLoadingKeys()(*i56f77a9c5e1a7ca279d87b52f582d52adbbc4de96f0fcef383fa9ded66e0a4c1.SideLoadingKeysRequestBuilder) {
    return i56f77a9c5e1a7ca279d87b52f582d52adbbc4de96f0fcef383fa9ded66e0a4c1.NewSideLoadingKeysRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SideLoadingKeysById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.sideLoadingKeys.item collection
func (m *DeviceAppManagementRequestBuilder) SideLoadingKeysById(id string)(*i6b77a717f9ad30ec5ef2cbce04e89b2244f69b8e725f39fb3990beeed77e175d.SideLoadingKeyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sideLoadingKey%2Did"] = id
    }
    return i6b77a717f9ad30ec5ef2cbce04e89b2244f69b8e725f39fb3990beeed77e175d.NewSideLoadingKeyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SymantecCodeSigningCertificate the symantecCodeSigningCertificate property
func (m *DeviceAppManagementRequestBuilder) SymantecCodeSigningCertificate()(*i55154c3923d16b5ed36ced8585e111c41354882664d34667a6e3b97700313735.SymantecCodeSigningCertificateRequestBuilder) {
    return i55154c3923d16b5ed36ced8585e111c41354882664d34667a6e3b97700313735.NewSymantecCodeSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SyncMicrosoftStoreForBusinessApps the syncMicrosoftStoreForBusinessApps property
func (m *DeviceAppManagementRequestBuilder) SyncMicrosoftStoreForBusinessApps()(*ief11f22801272e064478cc3e45b3a0f1e63de4a918dab77a8cb1f3924109be70.SyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    return ief11f22801272e064478cc3e45b3a0f1e63de4a918dab77a8cb1f3924109be70.NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetedManagedAppConfigurations the targetedManagedAppConfigurations property
func (m *DeviceAppManagementRequestBuilder) TargetedManagedAppConfigurations()(*i39a576817b6f726c83eda563ed7f79852514888b74a8942a87e4fb08eb489a74.TargetedManagedAppConfigurationsRequestBuilder) {
    return i39a576817b6f726c83eda563ed7f79852514888b74a8942a87e4fb08eb489a74.NewTargetedManagedAppConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetedManagedAppConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item collection
func (m *DeviceAppManagementRequestBuilder) TargetedManagedAppConfigurationsById(id string)(*iceecf955eeeaa6e59a7ff5c62bcd7ca93ca56e0d58f15cbe8d764abcf525d4fc.TargetedManagedAppConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppConfiguration%2Did"] = id
    }
    return iceecf955eeeaa6e59a7ff5c62bcd7ca93ca56e0d58f15cbe8d764abcf525d4fc.NewTargetedManagedAppConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VppTokens the vppTokens property
func (m *DeviceAppManagementRequestBuilder) VppTokens()(*iee0ff8dcd5df6303aaa9e4b0eb5d5437722e178fc10f16c2284bd8702f364eaa.VppTokensRequestBuilder) {
    return iee0ff8dcd5df6303aaa9e4b0eb5d5437722e178fc10f16c2284bd8702f364eaa.NewVppTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VppTokensById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.vppTokens.item collection
func (m *DeviceAppManagementRequestBuilder) VppTokensById(id string)(*i3b995f6d41023e3b43346c95c0f59fe84cb19f9a05efc8999cbbd88c8e143b9c.VppTokenItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["vppToken%2Did"] = id
    }
    return i3b995f6d41023e3b43346c95c0f59fe84cb19f9a05efc8999cbbd88c8e143b9c.NewVppTokenItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WdacSupplementalPolicies the wdacSupplementalPolicies property
func (m *DeviceAppManagementRequestBuilder) WdacSupplementalPolicies()(*i8b9340bb6f8a5fae484feebe60a0c77ea793398c1470ee9294c37011a384a585.WdacSupplementalPoliciesRequestBuilder) {
    return i8b9340bb6f8a5fae484feebe60a0c77ea793398c1470ee9294c37011a384a585.NewWdacSupplementalPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WdacSupplementalPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.wdacSupplementalPolicies.item collection
func (m *DeviceAppManagementRequestBuilder) WdacSupplementalPoliciesById(id string)(*i37f5ce10204a635850269ccbe65874caffc5e6d79af25b669118274ddd1b167e.WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDefenderApplicationControlSupplementalPolicy%2Did"] = id
    }
    return i37f5ce10204a635850269ccbe65874caffc5e6d79af25b669118274ddd1b167e.NewWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionDeviceRegistrations the windowsInformationProtectionDeviceRegistrations property
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionDeviceRegistrations()(*i63940f883f70d79c9f036190af3786b785aa38e88077dee0da7d8bdd443ffb95.WindowsInformationProtectionDeviceRegistrationsRequestBuilder) {
    return i63940f883f70d79c9f036190af3786b785aa38e88077dee0da7d8bdd443ffb95.NewWindowsInformationProtectionDeviceRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionDeviceRegistrationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.windowsInformationProtectionDeviceRegistrations.item collection
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionDeviceRegistrationsById(id string)(*i9c05ace01a97f06826939943c032786e845da3515dd142da04b21d1ed556d243.WindowsInformationProtectionDeviceRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionDeviceRegistration%2Did"] = id
    }
    return i9c05ace01a97f06826939943c032786e845da3515dd142da04b21d1ed556d243.NewWindowsInformationProtectionDeviceRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionPolicies the windowsInformationProtectionPolicies property
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionPolicies()(*ie802a5a736b1aceb4dbbe40ec170346b4ae36c2165343b37c69767334cdf1f05.WindowsInformationProtectionPoliciesRequestBuilder) {
    return ie802a5a736b1aceb4dbbe40ec170346b4ae36c2165343b37c69767334cdf1f05.NewWindowsInformationProtectionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.windowsInformationProtectionPolicies.item collection
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionPoliciesById(id string)(*ib84c840989f9ab439ec3ef07f903c5bac878d6066f60bc53dac1f18abcc7ad13.WindowsInformationProtectionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionPolicy%2Did"] = id
    }
    return ib84c840989f9ab439ec3ef07f903c5bac878d6066f60bc53dac1f18abcc7ad13.NewWindowsInformationProtectionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionWipeActions the windowsInformationProtectionWipeActions property
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionWipeActions()(*i869aad9b4b95674254fbc32234cd41bc975fb5a270d41358d0b38c0a515d0e76.WindowsInformationProtectionWipeActionsRequestBuilder) {
    return i869aad9b4b95674254fbc32234cd41bc975fb5a270d41358d0b38c0a515d0e76.NewWindowsInformationProtectionWipeActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionWipeActionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.windowsInformationProtectionWipeActions.item collection
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionWipeActionsById(id string)(*ib4dcb1c692a4cab02e05d87e0707aa4831b962f34cad5157cb939f82a9345efb.WindowsInformationProtectionWipeActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionWipeAction%2Did"] = id
    }
    return ib4dcb1c692a4cab02e05d87e0707aa4831b962f34cad5157cb939f82a9345efb.NewWindowsInformationProtectionWipeActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsManagementApp the windowsManagementApp property
func (m *DeviceAppManagementRequestBuilder) WindowsManagementApp()(*i9b677b07864604eb0a1dc7960a42ddf4914647e3fa83a6d2cfea3e12250a4704.WindowsManagementAppRequestBuilder) {
    return i9b677b07864604eb0a1dc7960a42ddf4914647e3fa83a6d2cfea3e12250a4704.NewWindowsManagementAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
