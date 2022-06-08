package serviceprincipal

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0f2de0f63d6d8458cacca4728c3e3ccfb4196169bcd1b56b58505b61755c4524 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/createpasswordsinglesignoncredentials"
    i27ca6a328f1c1efb1a1b5693b01513b9d52935958fe29af52d89b671dcabf2bf "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/restore"
    i2d44d1b4510c07ac99b627b0d74f44afbb28666abf0593d5d9b84993ec389b37 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/getpasswordsinglesignoncredentials"
    i4b96c191f9213c2032df0a0eab7ea47251670b0a2398dc4f437d25d3b6fd9885 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/deletepasswordsinglesignoncredentials"
    i74e90d5937c8543c88fd4f22bb84b87325163b22954182c4e71697e4521e57fa "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/getmembergroups"
    i7b5d28535d2b01d410fd7795a0a0e64ba2f0edc91ede15cc41749c9711149e5c "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/checkmembergroups"
    i87536da35200ad246fdff8de3f133c9998f02925e09d438d90bb4292a871e8c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/getmemberobjects"
    ia18bca6e8eb64967e1d5c72390db11b30b7fd9de504d303b0965c445c9273100 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/checkmemberobjects"
    ia4b4ef337d502e09bcb512fb2bfb7c4737711813aaf957b0609649801e61bad0 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/updatepasswordsinglesignoncredentials"
    ib8ce88e832bd90eed08326153b46fe424a619977a579bdc315c738b88ed90b20 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/owners/item/serviceprincipal/addtokensigningcertificate"
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
func (m *ServicePrincipalRequestBuilder) AddTokenSigningCertificate()(*ib8ce88e832bd90eed08326153b46fe424a619977a579bdc315c738b88ed90b20.AddTokenSigningCertificateRequestBuilder) {
    return ib8ce88e832bd90eed08326153b46fe424a619977a579bdc315c738b88ed90b20.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *ServicePrincipalRequestBuilder) CheckMemberGroups()(*i7b5d28535d2b01d410fd7795a0a0e64ba2f0edc91ede15cc41749c9711149e5c.CheckMemberGroupsRequestBuilder) {
    return i7b5d28535d2b01d410fd7795a0a0e64ba2f0edc91ede15cc41749c9711149e5c.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *ServicePrincipalRequestBuilder) CheckMemberObjects()(*ia18bca6e8eb64967e1d5c72390db11b30b7fd9de504d303b0965c445c9273100.CheckMemberObjectsRequestBuilder) {
    return ia18bca6e8eb64967e1d5c72390db11b30b7fd9de504d303b0965c445c9273100.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalRequestBuilderInternal instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    m := &ServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/owners/{directoryObject%2Did}/microsoft.graph.servicePrincipal{?%24select,%24expand}";
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
func (m *ServicePrincipalRequestBuilder) CreatePasswordSingleSignOnCredentials()(*i0f2de0f63d6d8458cacca4728c3e3ccfb4196169bcd1b56b58505b61755c4524.CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    return i0f2de0f63d6d8458cacca4728c3e3ccfb4196169bcd1b56b58505b61755c4524.NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeletePasswordSingleSignOnCredentials the deletePasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) DeletePasswordSingleSignOnCredentials()(*i4b96c191f9213c2032df0a0eab7ea47251670b0a2398dc4f437d25d3b6fd9885.DeletePasswordSingleSignOnCredentialsRequestBuilder) {
    return i4b96c191f9213c2032df0a0eab7ea47251670b0a2398dc4f437d25d3b6fd9885.NewDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *ServicePrincipalRequestBuilder) GetMemberGroups()(*i74e90d5937c8543c88fd4f22bb84b87325163b22954182c4e71697e4521e57fa.GetMemberGroupsRequestBuilder) {
    return i74e90d5937c8543c88fd4f22bb84b87325163b22954182c4e71697e4521e57fa.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *ServicePrincipalRequestBuilder) GetMemberObjects()(*i87536da35200ad246fdff8de3f133c9998f02925e09d438d90bb4292a871e8c3.GetMemberObjectsRequestBuilder) {
    return i87536da35200ad246fdff8de3f133c9998f02925e09d438d90bb4292a871e8c3.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPasswordSingleSignOnCredentials the getPasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) GetPasswordSingleSignOnCredentials()(*i2d44d1b4510c07ac99b627b0d74f44afbb28666abf0593d5d9b84993ec389b37.GetPasswordSingleSignOnCredentialsRequestBuilder) {
    return i2d44d1b4510c07ac99b627b0d74f44afbb28666abf0593d5d9b84993ec389b37.NewGetPasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ServicePrincipalRequestBuilder) Restore()(*i27ca6a328f1c1efb1a1b5693b01513b9d52935958fe29af52d89b671dcabf2bf.RestoreRequestBuilder) {
    return i27ca6a328f1c1efb1a1b5693b01513b9d52935958fe29af52d89b671dcabf2bf.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdatePasswordSingleSignOnCredentials the updatePasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) UpdatePasswordSingleSignOnCredentials()(*ia4b4ef337d502e09bcb512fb2bfb7c4737711813aaf957b0609649801e61bad0.UpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    return ia4b4ef337d502e09bcb512fb2bfb7c4737711813aaf957b0609649801e61bad0.NewUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
