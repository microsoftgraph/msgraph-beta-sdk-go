package serviceprincipal

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i34519f13643d1db296e62d1204eb9ea24f407c36dad4fbf4a0f217e8498dfc4e "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/checkmemberobjects"
    i4b36e321909cbb2c1e762a4da263937846e9b7ca339300c68921ac07c41e84a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/checkmembergroups"
    i507b1dae249b705d2e374563b1734359657950464a8cf1a49a6bf968acca72f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/getmembergroups"
    i5a5db61bd427dbafe593e648227ecac4de7f13c4043fa9011aeb0a0438805d90 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/restore"
    i5f7200fbb1388951d9399e13732517e4afff32e7bf982450313a51a717418999 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/deletepasswordsinglesignoncredentials"
    i6e717dbdf8ae8e48390ec12be43ae53598a847a8b59857b332841cf22fef1f5b "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/getpasswordsinglesignoncredentials"
    i8848f7f85f17bdffe5ce13e3beec52d89edcc1ad18455b744fb5f3756b8cce50 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/createpasswordsinglesignoncredentials"
    id2d48713926bc516c87ea989844d9ecd93f073d28f17eb8514754bc5be5147e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/updatepasswordsinglesignoncredentials"
    ie149e21bf28c887eebae4da51997a0de352b0b15f423ecd34b8426a6fdf91ee3 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/getmemberobjects"
    ife6334bed96d6d5399ea7a6053672804b2747176611846fb3243757d416934e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/memberof/item/serviceprincipal/addtokensigningcertificate"
)

// ServicePrincipalRequestBuilder casts the previous resource to servicePrincipal.
type ServicePrincipalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ServicePrincipalRequestBuilderGetQueryParameters get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
type ServicePrincipalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ServicePrincipalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServicePrincipalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServicePrincipalRequestBuilderGetQueryParameters
}
// AddTokenSigningCertificate the addTokenSigningCertificate property
func (m *ServicePrincipalRequestBuilder) AddTokenSigningCertificate()(*ife6334bed96d6d5399ea7a6053672804b2747176611846fb3243757d416934e6.AddTokenSigningCertificateRequestBuilder) {
    return ife6334bed96d6d5399ea7a6053672804b2747176611846fb3243757d416934e6.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *ServicePrincipalRequestBuilder) CheckMemberGroups()(*i4b36e321909cbb2c1e762a4da263937846e9b7ca339300c68921ac07c41e84a3.CheckMemberGroupsRequestBuilder) {
    return i4b36e321909cbb2c1e762a4da263937846e9b7ca339300c68921ac07c41e84a3.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *ServicePrincipalRequestBuilder) CheckMemberObjects()(*i34519f13643d1db296e62d1204eb9ea24f407c36dad4fbf4a0f217e8498dfc4e.CheckMemberObjectsRequestBuilder) {
    return i34519f13643d1db296e62d1204eb9ea24f407c36dad4fbf4a0f217e8498dfc4e.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalRequestBuilderInternal instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    m := &ServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/memberOf/{directoryObject%2Did}/microsoft.graph.servicePrincipal{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalRequestBuilder instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ServicePrincipalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePasswordSingleSignOnCredentials the createPasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) CreatePasswordSingleSignOnCredentials()(*i8848f7f85f17bdffe5ce13e3beec52d89edcc1ad18455b744fb5f3756b8cce50.CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    return i8848f7f85f17bdffe5ce13e3beec52d89edcc1ad18455b744fb5f3756b8cce50.NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeletePasswordSingleSignOnCredentials the deletePasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) DeletePasswordSingleSignOnCredentials()(*i5f7200fbb1388951d9399e13732517e4afff32e7bf982450313a51a717418999.DeletePasswordSingleSignOnCredentialsRequestBuilder) {
    return i5f7200fbb1388951d9399e13732517e4afff32e7bf982450313a51a717418999.NewDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *ServicePrincipalRequestBuilder) GetMemberGroups()(*i507b1dae249b705d2e374563b1734359657950464a8cf1a49a6bf968acca72f9.GetMemberGroupsRequestBuilder) {
    return i507b1dae249b705d2e374563b1734359657950464a8cf1a49a6bf968acca72f9.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *ServicePrincipalRequestBuilder) GetMemberObjects()(*ie149e21bf28c887eebae4da51997a0de352b0b15f423ecd34b8426a6fdf91ee3.GetMemberObjectsRequestBuilder) {
    return ie149e21bf28c887eebae4da51997a0de352b0b15f423ecd34b8426a6fdf91ee3.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPasswordSingleSignOnCredentials the getPasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) GetPasswordSingleSignOnCredentials()(*i6e717dbdf8ae8e48390ec12be43ae53598a847a8b59857b332841cf22fef1f5b.GetPasswordSingleSignOnCredentialsRequestBuilder) {
    return i6e717dbdf8ae8e48390ec12be43ae53598a847a8b59857b332841cf22fef1f5b.NewGetPasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ServicePrincipalRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable), nil
}
// Restore the restore property
func (m *ServicePrincipalRequestBuilder) Restore()(*i5a5db61bd427dbafe593e648227ecac4de7f13c4043fa9011aeb0a0438805d90.RestoreRequestBuilder) {
    return i5a5db61bd427dbafe593e648227ecac4de7f13c4043fa9011aeb0a0438805d90.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdatePasswordSingleSignOnCredentials the updatePasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) UpdatePasswordSingleSignOnCredentials()(*id2d48713926bc516c87ea989844d9ecd93f073d28f17eb8514754bc5be5147e2.UpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    return id2d48713926bc516c87ea989844d9ecd93f073d28f17eb8514754bc5be5147e2.NewUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
