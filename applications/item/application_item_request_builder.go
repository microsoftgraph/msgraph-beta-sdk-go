package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i04467bc34d30076c43d80ad9c9c6de922a43650392c42c6be56b6a35c122928a "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/addkey"
    i164dc366bcd5e26e795ff75a2fa13960d83e6396ec05782e6b76eb1f80c4b69b "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/removekey"
    i1a3a50fea4c09450bbaa46b5464d3f259d9942bd44610be4837c29a55cf268fa "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/addpassword"
    i2d2f49a5bb637ea5dd10ccf0e7c3776123a84e80ed6d9150c2fd2cc8ca9ed2f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/federatedidentitycredentials"
    i3d2f054c80bc83cd28a63a163f92036324daa39ab1bb99fd847444eaf48da698 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/connectorgroup"
    i40412afdaf06940a98614d093b1dd2f1ea12967583d8386057ef7f46b9fd39d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/restore"
    i421a4b8df71a68823647e3232383b123a09acda44ae135311cc12368c00e7d41 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/homerealmdiscoverypolicies"
    i44c509af08fd252f2be0cb5da863d7cca5d20ac2a9548cdd1f88301d9b2601d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/getmemberobjects"
    i4fb498b3f77ca56d96f4b992f35d175d96f70f18aeadf1bb78a8a7694b16564c "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization"
    i50081095d972d4a2a78924990332d6d59a05397bbfb0bf85eef836b409fff991 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/getmembergroups"
    i584715a25bcd7d9dc447f3cc78f525df77e5558f3196410f39f1a78f7b92ab5b "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/extensionproperties"
    i6ec319f3b5f1d756e35c94d5938dcc447f992675ff7619261013cfce47b5433f "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/appmanagementpolicies"
    i7e5220a4b31236710fca1c694c533878640ce68b3a6875a7a03705702b7dd46c "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/logo"
    i7f7ea3263ae77196be6ec02b1f28a4ae4b61762fbb52e65e9c9b247a2eea8db4 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/createdonbehalfof"
    i91ef18af4146c054c22e6b2fed52625f59106e7394aefb5393b0fc2f8702025b "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/tokenissuancepolicies"
    i968bec58f51f4a0e4b7589a024946f47a7f436310363a3495e774d74e852dc86 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/checkmembergroups"
    ia04b1ee345ac9f132a6ae36db8759224bc6e90208380aa9e9d72c710bc936d30 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/tokenlifetimepolicies"
    ia7925ed1ce8bdc9783d4f1c1fa9d9ef722a6794ba8b0df6950f7f8e94e406d49 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/checkmemberobjects"
    ice03e0056e3a2140c860a33ccc1cd964d23f1eddb120fb390eb2296af607367b "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/setverifiedpublisher"
    id95c889c9ed4179a13db497a7b0207f4c98d0896af3341816ebb4442d2408f90 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/owners"
    if7e9fe038cbc04c53e872ff5b374767b949c817ade18929784991f358eb803b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/unsetverifiedpublisher"
    ifd565415dad3a3abe5d2ac4a2be4b6c96c9ed1844b485ab5fbf1591e4d856097 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/removepassword"
    i1057b90c33df1436be53367c28118eb82679541523cc8fa16bc6c17493049c8e "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/owners/item"
    i1558b361e5c510a31b1b33eb613e141d243c7a3c850eaad8bb24ad236cf98b57 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/tokenissuancepolicies/item"
    i17d74174c0c203a9690a594729795841599510ec1474d89665b17043a9e520da "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/extensionproperties/item"
    i4f78489d134c25005e63ce787c48d42f78dc8cac5bcb1050d310001e92e929de "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/appmanagementpolicies/item"
    i5c358ed7d22542209549181254d6ed489daa984f427bcb9904dccf894d25c561 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/homerealmdiscoverypolicies/item"
    icb5d61d2d1b507de0b242478ab236165e1ff885491389b259718f3f2891d63f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/federatedidentitycredentials/item"
    ie900fefdee05771652101d59a999b01864158be539a034e360cbdbef37278abb "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/tokenlifetimepolicies/item"
)

// ApplicationItemRequestBuilder provides operations to manage the collection of application entities.
type ApplicationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ApplicationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplicationItemRequestBuilderGetQueryParameters get the properties and relationships of an application object.
type ApplicationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ApplicationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ApplicationItemRequestBuilderGetQueryParameters
}
// ApplicationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ApplicationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddKey the addKey property
func (m *ApplicationItemRequestBuilder) AddKey()(*i04467bc34d30076c43d80ad9c9c6de922a43650392c42c6be56b6a35c122928a.AddKeyRequestBuilder) {
    return i04467bc34d30076c43d80ad9c9c6de922a43650392c42c6be56b6a35c122928a.NewAddKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddPassword the addPassword property
func (m *ApplicationItemRequestBuilder) AddPassword()(*i1a3a50fea4c09450bbaa46b5464d3f259d9942bd44610be4837c29a55cf268fa.AddPasswordRequestBuilder) {
    return i1a3a50fea4c09450bbaa46b5464d3f259d9942bd44610be4837c29a55cf268fa.NewAddPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppManagementPolicies the appManagementPolicies property
func (m *ApplicationItemRequestBuilder) AppManagementPolicies()(*i6ec319f3b5f1d756e35c94d5938dcc447f992675ff7619261013cfce47b5433f.AppManagementPoliciesRequestBuilder) {
    return i6ec319f3b5f1d756e35c94d5938dcc447f992675ff7619261013cfce47b5433f.NewAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppManagementPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.appManagementPolicies.item collection
func (m *ApplicationItemRequestBuilder) AppManagementPoliciesById(id string)(*i4f78489d134c25005e63ce787c48d42f78dc8cac5bcb1050d310001e92e929de.AppManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appManagementPolicy%2Did"] = id
    }
    return i4f78489d134c25005e63ce787c48d42f78dc8cac5bcb1050d310001e92e929de.NewAppManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *ApplicationItemRequestBuilder) CheckMemberGroups()(*i968bec58f51f4a0e4b7589a024946f47a7f436310363a3495e774d74e852dc86.CheckMemberGroupsRequestBuilder) {
    return i968bec58f51f4a0e4b7589a024946f47a7f436310363a3495e774d74e852dc86.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *ApplicationItemRequestBuilder) CheckMemberObjects()(*ia7925ed1ce8bdc9783d4f1c1fa9d9ef722a6794ba8b0df6950f7f8e94e406d49.CheckMemberObjectsRequestBuilder) {
    return ia7925ed1ce8bdc9783d4f1c1fa9d9ef722a6794ba8b0df6950f7f8e94e406d49.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorGroup the connectorGroup property
func (m *ApplicationItemRequestBuilder) ConnectorGroup()(*i3d2f054c80bc83cd28a63a163f92036324daa39ab1bb99fd847444eaf48da698.ConnectorGroupRequestBuilder) {
    return i3d2f054c80bc83cd28a63a163f92036324daa39ab1bb99fd847444eaf48da698.NewConnectorGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewApplicationItemRequestBuilderInternal instantiates a new ApplicationItemRequestBuilder and sets the default values.
func NewApplicationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationItemRequestBuilder) {
    m := &ApplicationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApplicationItemRequestBuilder instantiates a new ApplicationItemRequestBuilder and sets the default values.
func NewApplicationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation deletes an application. When deleted, apps are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted.
func (m *ApplicationItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration deletes an application. When deleted, apps are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted.
func (m *ApplicationItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ApplicationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatedOnBehalfOf the createdOnBehalfOf property
func (m *ApplicationItemRequestBuilder) CreatedOnBehalfOf()(*i7f7ea3263ae77196be6ec02b1f28a4ae4b61762fbb52e65e9c9b247a2eea8db4.CreatedOnBehalfOfRequestBuilder) {
    return i7f7ea3263ae77196be6ec02b1f28a4ae4b61762fbb52e65e9c9b247a2eea8db4.NewCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get the properties and relationships of an application object.
func (m *ApplicationItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get the properties and relationships of an application object.
func (m *ApplicationItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ApplicationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of an application object.
func (m *ApplicationItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the properties of an application object.
func (m *ApplicationItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable, requestConfiguration *ApplicationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete deletes an application. When deleted, apps are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted.
func (m *ApplicationItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler deletes an application. When deleted, apps are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted.
func (m *ApplicationItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ApplicationItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// ExtensionProperties the extensionProperties property
func (m *ApplicationItemRequestBuilder) ExtensionProperties()(*i584715a25bcd7d9dc447f3cc78f525df77e5558f3196410f39f1a78f7b92ab5b.ExtensionPropertiesRequestBuilder) {
    return i584715a25bcd7d9dc447f3cc78f525df77e5558f3196410f39f1a78f7b92ab5b.NewExtensionPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.extensionProperties.item collection
func (m *ApplicationItemRequestBuilder) ExtensionPropertiesById(id string)(*i17d74174c0c203a9690a594729795841599510ec1474d89665b17043a9e520da.ExtensionPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extensionProperty%2Did"] = id
    }
    return i17d74174c0c203a9690a594729795841599510ec1474d89665b17043a9e520da.NewExtensionPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FederatedIdentityCredentials the federatedIdentityCredentials property
func (m *ApplicationItemRequestBuilder) FederatedIdentityCredentials()(*i2d2f49a5bb637ea5dd10ccf0e7c3776123a84e80ed6d9150c2fd2cc8ca9ed2f9.FederatedIdentityCredentialsRequestBuilder) {
    return i2d2f49a5bb637ea5dd10ccf0e7c3776123a84e80ed6d9150c2fd2cc8ca9ed2f9.NewFederatedIdentityCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FederatedIdentityCredentialsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.federatedIdentityCredentials.item collection
func (m *ApplicationItemRequestBuilder) FederatedIdentityCredentialsById(id string)(*icb5d61d2d1b507de0b242478ab236165e1ff885491389b259718f3f2891d63f2.FederatedIdentityCredentialItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["federatedIdentityCredential%2Did"] = id
    }
    return icb5d61d2d1b507de0b242478ab236165e1ff885491389b259718f3f2891d63f2.NewFederatedIdentityCredentialItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get the properties and relationships of an application object.
func (m *ApplicationItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *ApplicationItemRequestBuilder) GetMemberGroups()(*i50081095d972d4a2a78924990332d6d59a05397bbfb0bf85eef836b409fff991.GetMemberGroupsRequestBuilder) {
    return i50081095d972d4a2a78924990332d6d59a05397bbfb0bf85eef836b409fff991.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *ApplicationItemRequestBuilder) GetMemberObjects()(*i44c509af08fd252f2be0cb5da863d7cca5d20ac2a9548cdd1f88301d9b2601d5.GetMemberObjectsRequestBuilder) {
    return i44c509af08fd252f2be0cb5da863d7cca5d20ac2a9548cdd1f88301d9b2601d5.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get the properties and relationships of an application object.
func (m *ApplicationItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ApplicationItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApplicationFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable), nil
}
// HomeRealmDiscoveryPolicies the homeRealmDiscoveryPolicies property
func (m *ApplicationItemRequestBuilder) HomeRealmDiscoveryPolicies()(*i421a4b8df71a68823647e3232383b123a09acda44ae135311cc12368c00e7d41.HomeRealmDiscoveryPoliciesRequestBuilder) {
    return i421a4b8df71a68823647e3232383b123a09acda44ae135311cc12368c00e7d41.NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.homeRealmDiscoveryPolicies.item collection
func (m *ApplicationItemRequestBuilder) HomeRealmDiscoveryPoliciesById(id string)(*i5c358ed7d22542209549181254d6ed489daa984f427bcb9904dccf894d25c561.HomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["homeRealmDiscoveryPolicy%2Did"] = id
    }
    return i5c358ed7d22542209549181254d6ed489daa984f427bcb9904dccf894d25c561.NewHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Logo the logo property
func (m *ApplicationItemRequestBuilder) Logo()(*i7e5220a4b31236710fca1c694c533878640ce68b3a6875a7a03705702b7dd46c.LogoRequestBuilder) {
    return i7e5220a4b31236710fca1c694c533878640ce68b3a6875a7a03705702b7dd46c.NewLogoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Owners the owners property
func (m *ApplicationItemRequestBuilder) Owners()(*id95c889c9ed4179a13db497a7b0207f4c98d0896af3341816ebb4442d2408f90.OwnersRequestBuilder) {
    return id95c889c9ed4179a13db497a7b0207f4c98d0896af3341816ebb4442d2408f90.NewOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.owners.item collection
func (m *ApplicationItemRequestBuilder) OwnersById(id string)(*i1057b90c33df1436be53367c28118eb82679541523cc8fa16bc6c17493049c8e.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i1057b90c33df1436be53367c28118eb82679541523cc8fa16bc6c17493049c8e.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of an application object.
func (m *ApplicationItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the properties of an application object.
func (m *ApplicationItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Applicationable, requestConfiguration *ApplicationItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// RemoveKey the removeKey property
func (m *ApplicationItemRequestBuilder) RemoveKey()(*i164dc366bcd5e26e795ff75a2fa13960d83e6396ec05782e6b76eb1f80c4b69b.RemoveKeyRequestBuilder) {
    return i164dc366bcd5e26e795ff75a2fa13960d83e6396ec05782e6b76eb1f80c4b69b.NewRemoveKeyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemovePassword the removePassword property
func (m *ApplicationItemRequestBuilder) RemovePassword()(*ifd565415dad3a3abe5d2ac4a2be4b6c96c9ed1844b485ab5fbf1591e4d856097.RemovePasswordRequestBuilder) {
    return ifd565415dad3a3abe5d2ac4a2be4b6c96c9ed1844b485ab5fbf1591e4d856097.NewRemovePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *ApplicationItemRequestBuilder) Restore()(*i40412afdaf06940a98614d093b1dd2f1ea12967583d8386057ef7f46b9fd39d9.RestoreRequestBuilder) {
    return i40412afdaf06940a98614d093b1dd2f1ea12967583d8386057ef7f46b9fd39d9.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetVerifiedPublisher the setVerifiedPublisher property
func (m *ApplicationItemRequestBuilder) SetVerifiedPublisher()(*ice03e0056e3a2140c860a33ccc1cd964d23f1eddb120fb390eb2296af607367b.SetVerifiedPublisherRequestBuilder) {
    return ice03e0056e3a2140c860a33ccc1cd964d23f1eddb120fb390eb2296af607367b.NewSetVerifiedPublisherRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Synchronization the synchronization property
func (m *ApplicationItemRequestBuilder) Synchronization()(*i4fb498b3f77ca56d96f4b992f35d175d96f70f18aeadf1bb78a8a7694b16564c.SynchronizationRequestBuilder) {
    return i4fb498b3f77ca56d96f4b992f35d175d96f70f18aeadf1bb78a8a7694b16564c.NewSynchronizationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePolicies the tokenIssuancePolicies property
func (m *ApplicationItemRequestBuilder) TokenIssuancePolicies()(*i91ef18af4146c054c22e6b2fed52625f59106e7394aefb5393b0fc2f8702025b.TokenIssuancePoliciesRequestBuilder) {
    return i91ef18af4146c054c22e6b2fed52625f59106e7394aefb5393b0fc2f8702025b.NewTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.tokenIssuancePolicies.item collection
func (m *ApplicationItemRequestBuilder) TokenIssuancePoliciesById(id string)(*i1558b361e5c510a31b1b33eb613e141d243c7a3c850eaad8bb24ad236cf98b57.TokenIssuancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenIssuancePolicy%2Did"] = id
    }
    return i1558b361e5c510a31b1b33eb613e141d243c7a3c850eaad8bb24ad236cf98b57.NewTokenIssuancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TokenLifetimePolicies the tokenLifetimePolicies property
func (m *ApplicationItemRequestBuilder) TokenLifetimePolicies()(*ia04b1ee345ac9f132a6ae36db8759224bc6e90208380aa9e9d72c710bc936d30.TokenLifetimePoliciesRequestBuilder) {
    return ia04b1ee345ac9f132a6ae36db8759224bc6e90208380aa9e9d72c710bc936d30.NewTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenLifetimePoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.applications.item.tokenLifetimePolicies.item collection
func (m *ApplicationItemRequestBuilder) TokenLifetimePoliciesById(id string)(*ie900fefdee05771652101d59a999b01864158be539a034e360cbdbef37278abb.TokenLifetimePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenLifetimePolicy%2Did"] = id
    }
    return ie900fefdee05771652101d59a999b01864158be539a034e360cbdbef37278abb.NewTokenLifetimePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnsetVerifiedPublisher the unsetVerifiedPublisher property
func (m *ApplicationItemRequestBuilder) UnsetVerifiedPublisher()(*if7e9fe038cbc04c53e872ff5b374767b949c817ade18929784991f358eb803b0.UnsetVerifiedPublisherRequestBuilder) {
    return if7e9fe038cbc04c53e872ff5b374767b949c817ade18929784991f358eb803b0.NewUnsetVerifiedPublisherRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
