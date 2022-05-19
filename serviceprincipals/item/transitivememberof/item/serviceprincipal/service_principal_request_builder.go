package serviceprincipal

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0f1dd1398fa7d39e1a2f116e64d3c53da063a155267643a9ce37291e6cb17e73 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/getmembergroups"
    i2fcd7974eaaae46c52df0854e59e276a7aa2b6835fe5ab533630a269d72afd71 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/checkmemberobjects"
    i444a0c0604bff1fc60998f92b80c3dc8f799b6539d8d65b8906519f8e3f187a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/createpasswordsinglesignoncredentials"
    i54c50f20cde92ff93320c7993fab782bf39807b3ea3bd7ed36d5b9bdeab39089 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/restore"
    i54dd1743202f8ae6d4bcead9e7426bd7202c13981fa502bcfdaf60c2e401f87d "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/getpasswordsinglesignoncredentials"
    i5c3b2b8f22f993d746797f5f2600f98af5ec67e516025030e91f7cfb5e863ddd "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/addtokensigningcertificate"
    i74aa8c6bcaa09c67a0b865864d458a2857692aa15ca54007b7ec4e5aaf240c16 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/deletepasswordsinglesignoncredentials"
    ib40825d856bee037764dbd96a35747da1e4671fa60eb762bc15285d9b8cf94ad "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/updatepasswordsinglesignoncredentials"
    ie4ffd78f88eeb9047d95768eb494f09a8112d79a15d9ee60e8bbc78ae05d6fc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/checkmembergroups"
    if29de8d6ebf7a3267ee8e8da256ebb5a966718025e5605b54fb85990cdba202d "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/transitivememberof/item/serviceprincipal/getmemberobjects"
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
func (m *ServicePrincipalRequestBuilder) AddTokenSigningCertificate()(*i5c3b2b8f22f993d746797f5f2600f98af5ec67e516025030e91f7cfb5e863ddd.AddTokenSigningCertificateRequestBuilder) {
    return i5c3b2b8f22f993d746797f5f2600f98af5ec67e516025030e91f7cfb5e863ddd.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *ServicePrincipalRequestBuilder) CheckMemberGroups()(*ie4ffd78f88eeb9047d95768eb494f09a8112d79a15d9ee60e8bbc78ae05d6fc2.CheckMemberGroupsRequestBuilder) {
    return ie4ffd78f88eeb9047d95768eb494f09a8112d79a15d9ee60e8bbc78ae05d6fc2.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *ServicePrincipalRequestBuilder) CheckMemberObjects()(*i2fcd7974eaaae46c52df0854e59e276a7aa2b6835fe5ab533630a269d72afd71.CheckMemberObjectsRequestBuilder) {
    return i2fcd7974eaaae46c52df0854e59e276a7aa2b6835fe5ab533630a269d72afd71.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalRequestBuilderInternal instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    m := &ServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/transitiveMemberOf/{directoryObject%2Did}/microsoft.graph.servicePrincipal{?%24select,%24expand}";
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
func (m *ServicePrincipalRequestBuilder) CreatePasswordSingleSignOnCredentials()(*i444a0c0604bff1fc60998f92b80c3dc8f799b6539d8d65b8906519f8e3f187a7.CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    return i444a0c0604bff1fc60998f92b80c3dc8f799b6539d8d65b8906519f8e3f187a7.NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeletePasswordSingleSignOnCredentials the deletePasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) DeletePasswordSingleSignOnCredentials()(*i74aa8c6bcaa09c67a0b865864d458a2857692aa15ca54007b7ec4e5aaf240c16.DeletePasswordSingleSignOnCredentialsRequestBuilder) {
    return i74aa8c6bcaa09c67a0b865864d458a2857692aa15ca54007b7ec4e5aaf240c16.NewDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *ServicePrincipalRequestBuilder) GetMemberGroups()(*i0f1dd1398fa7d39e1a2f116e64d3c53da063a155267643a9ce37291e6cb17e73.GetMemberGroupsRequestBuilder) {
    return i0f1dd1398fa7d39e1a2f116e64d3c53da063a155267643a9ce37291e6cb17e73.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *ServicePrincipalRequestBuilder) GetMemberObjects()(*if29de8d6ebf7a3267ee8e8da256ebb5a966718025e5605b54fb85990cdba202d.GetMemberObjectsRequestBuilder) {
    return if29de8d6ebf7a3267ee8e8da256ebb5a966718025e5605b54fb85990cdba202d.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPasswordSingleSignOnCredentials the getPasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) GetPasswordSingleSignOnCredentials()(*i54dd1743202f8ae6d4bcead9e7426bd7202c13981fa502bcfdaf60c2e401f87d.GetPasswordSingleSignOnCredentialsRequestBuilder) {
    return i54dd1743202f8ae6d4bcead9e7426bd7202c13981fa502bcfdaf60c2e401f87d.NewGetPasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ServicePrincipalRequestBuilder) Restore()(*i54c50f20cde92ff93320c7993fab782bf39807b3ea3bd7ed36d5b9bdeab39089.RestoreRequestBuilder) {
    return i54c50f20cde92ff93320c7993fab782bf39807b3ea3bd7ed36d5b9bdeab39089.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdatePasswordSingleSignOnCredentials the updatePasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) UpdatePasswordSingleSignOnCredentials()(*ib40825d856bee037764dbd96a35747da1e4671fa60eb762bc15285d9b8cf94ad.UpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    return ib40825d856bee037764dbd96a35747da1e4671fa60eb762bc15285d9b8cf94ad.NewUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
