package serviceprincipal

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i15596b5959ce8f1a3b5b12a70b0006d4f554321bf51b256c9f43bc98d8f3e4f9 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/updatepasswordsinglesignoncredentials"
    i3c43da359dbe8a2abb70814ae9b2991d4d713c27585f6007a471a1b2d9ec2d43 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/addtokensigningcertificate"
    i5cd500be322da5700eeabb78364ead49aa3d781792d2a7e36fc3cd20f448a13b "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/getmemberobjects"
    i7a1063174de0e0b51c19e0119e13e56f33a42c1aceba382f22aaab553c7372cd "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/getmembergroups"
    i7b767cbc543654e76ba6b9732cb00f0036703057128af1b71e9970f3e830638f "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/deletepasswordsinglesignoncredentials"
    i8a8f7e38602926a73ff51f91ea4e722ab3b2e6f14e2aeeed6f4b9900fd176f01 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/checkmemberobjects"
    i8b83d5843aaf6001b192f695a1c90fc9bc28be65a57c1058a39e9ff3d20bc830 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/getpasswordsinglesignoncredentials"
    iacfc28e87a7bf6d3b3e82b86d10df9ecad52034efc28f3ce0e6df7d30e0addfd "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/checkmembergroups"
    iaedc166b8e666331e93f6958f3efd6d428986b0f1ace2302319a3b79ae16906d "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/createpasswordsinglesignoncredentials"
    iba83c93ca38410f8f22da28bd25ca78c352956c757646a220583a8ecce2fbd3e "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/ownedobjects/item/serviceprincipal/restore"
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
func (m *ServicePrincipalRequestBuilder) AddTokenSigningCertificate()(*i3c43da359dbe8a2abb70814ae9b2991d4d713c27585f6007a471a1b2d9ec2d43.AddTokenSigningCertificateRequestBuilder) {
    return i3c43da359dbe8a2abb70814ae9b2991d4d713c27585f6007a471a1b2d9ec2d43.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *ServicePrincipalRequestBuilder) CheckMemberGroups()(*iacfc28e87a7bf6d3b3e82b86d10df9ecad52034efc28f3ce0e6df7d30e0addfd.CheckMemberGroupsRequestBuilder) {
    return iacfc28e87a7bf6d3b3e82b86d10df9ecad52034efc28f3ce0e6df7d30e0addfd.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *ServicePrincipalRequestBuilder) CheckMemberObjects()(*i8a8f7e38602926a73ff51f91ea4e722ab3b2e6f14e2aeeed6f4b9900fd176f01.CheckMemberObjectsRequestBuilder) {
    return i8a8f7e38602926a73ff51f91ea4e722ab3b2e6f14e2aeeed6f4b9900fd176f01.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalRequestBuilderInternal instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    m := &ServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/ownedObjects/{directoryObject%2Did}/microsoft.graph.servicePrincipal{?%24select,%24expand}";
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
// CreatePasswordSingleSignOnCredentials the createPasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) CreatePasswordSingleSignOnCredentials()(*iaedc166b8e666331e93f6958f3efd6d428986b0f1ace2302319a3b79ae16906d.CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    return iaedc166b8e666331e93f6958f3efd6d428986b0f1ace2302319a3b79ae16906d.NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeletePasswordSingleSignOnCredentials the deletePasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) DeletePasswordSingleSignOnCredentials()(*i7b767cbc543654e76ba6b9732cb00f0036703057128af1b71e9970f3e830638f.DeletePasswordSingleSignOnCredentialsRequestBuilder) {
    return i7b767cbc543654e76ba6b9732cb00f0036703057128af1b71e9970f3e830638f.NewDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *ServicePrincipalRequestBuilder) GetMemberGroups()(*i7a1063174de0e0b51c19e0119e13e56f33a42c1aceba382f22aaab553c7372cd.GetMemberGroupsRequestBuilder) {
    return i7a1063174de0e0b51c19e0119e13e56f33a42c1aceba382f22aaab553c7372cd.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *ServicePrincipalRequestBuilder) GetMemberObjects()(*i5cd500be322da5700eeabb78364ead49aa3d781792d2a7e36fc3cd20f448a13b.GetMemberObjectsRequestBuilder) {
    return i5cd500be322da5700eeabb78364ead49aa3d781792d2a7e36fc3cd20f448a13b.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPasswordSingleSignOnCredentials the getPasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) GetPasswordSingleSignOnCredentials()(*i8b83d5843aaf6001b192f695a1c90fc9bc28be65a57c1058a39e9ff3d20bc830.GetPasswordSingleSignOnCredentialsRequestBuilder) {
    return i8b83d5843aaf6001b192f695a1c90fc9bc28be65a57c1058a39e9ff3d20bc830.NewGetPasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ServicePrincipalRequestBuilder) Restore()(*iba83c93ca38410f8f22da28bd25ca78c352956c757646a220583a8ecce2fbd3e.RestoreRequestBuilder) {
    return iba83c93ca38410f8f22da28bd25ca78c352956c757646a220583a8ecce2fbd3e.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdatePasswordSingleSignOnCredentials the updatePasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) UpdatePasswordSingleSignOnCredentials()(*i15596b5959ce8f1a3b5b12a70b0006d4f554321bf51b256c9f43bc98d8f3e4f9.UpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    return i15596b5959ce8f1a3b5b12a70b0006d4f554321bf51b256c9f43bc98d8f3e4f9.NewUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
