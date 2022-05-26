package serviceprincipal

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i066339b4e1e2d7b919da71fc6c51ca45a60746759daf9e404441e30faad703a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/checkmembergroups"
    i1f94bd1ec9ba9f8ce10416c3f8b76fb6da9dddd7ed173da40b5613d2b4aa52aa "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/getmemberobjects"
    i22118a679a710bc76e4e17e139f4262ac1dc74f52f290ea5be31d1e7ffa3c201 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/checkmemberobjects"
    i4cbadb93cdbf810c8551f7af31425d51c10f4daf3612aec0324280da87ef778a "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/restore"
    i54ff55c1232b7f7a77d695ed3fa0635e493179f9021338e0b05db9ed5c67a19f "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/updatepasswordsinglesignoncredentials"
    i7273daa3c244c1334395e3dc78e6aa7b63e405247ce1b153c808276999356234 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/deletepasswordsinglesignoncredentials"
    i7650ae7ee4e4c5e0ed3bbd79a74be7e7d5661ae8041fca71668918ef4bdfe92e "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/addtokensigningcertificate"
    i92fdcf12e4e6aec7316f7be9c10c4e3f760f4ac80f7c8e4b42f8d9a875b525f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/getpasswordsinglesignoncredentials"
    ia8bbe15a5cbcc08f1fb2f66c0a6243894cb4d2cebce992d12074d5d593e7ed2a "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/createpasswordsinglesignoncredentials"
    id747d8a4068cca2fc00c2a191d305a94cbeb5d7b4c49db6ed85e3584cfc7dba8 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/createdobjects/item/serviceprincipal/getmembergroups"
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
func (m *ServicePrincipalRequestBuilder) AddTokenSigningCertificate()(*i7650ae7ee4e4c5e0ed3bbd79a74be7e7d5661ae8041fca71668918ef4bdfe92e.AddTokenSigningCertificateRequestBuilder) {
    return i7650ae7ee4e4c5e0ed3bbd79a74be7e7d5661ae8041fca71668918ef4bdfe92e.NewAddTokenSigningCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups the checkMemberGroups property
func (m *ServicePrincipalRequestBuilder) CheckMemberGroups()(*i066339b4e1e2d7b919da71fc6c51ca45a60746759daf9e404441e30faad703a2.CheckMemberGroupsRequestBuilder) {
    return i066339b4e1e2d7b919da71fc6c51ca45a60746759daf9e404441e30faad703a2.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *ServicePrincipalRequestBuilder) CheckMemberObjects()(*i22118a679a710bc76e4e17e139f4262ac1dc74f52f290ea5be31d1e7ffa3c201.CheckMemberObjectsRequestBuilder) {
    return i22118a679a710bc76e4e17e139f4262ac1dc74f52f290ea5be31d1e7ffa3c201.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalRequestBuilderInternal instantiates a new ServicePrincipalRequestBuilder and sets the default values.
func NewServicePrincipalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalRequestBuilder) {
    m := &ServicePrincipalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/createdObjects/{directoryObject%2Did}/microsoft.graph.servicePrincipal{?%24select,%24expand}";
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
func (m *ServicePrincipalRequestBuilder) CreatePasswordSingleSignOnCredentials()(*ia8bbe15a5cbcc08f1fb2f66c0a6243894cb4d2cebce992d12074d5d593e7ed2a.CreatePasswordSingleSignOnCredentialsRequestBuilder) {
    return ia8bbe15a5cbcc08f1fb2f66c0a6243894cb4d2cebce992d12074d5d593e7ed2a.NewCreatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeletePasswordSingleSignOnCredentials the deletePasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) DeletePasswordSingleSignOnCredentials()(*i7273daa3c244c1334395e3dc78e6aa7b63e405247ce1b153c808276999356234.DeletePasswordSingleSignOnCredentialsRequestBuilder) {
    return i7273daa3c244c1334395e3dc78e6aa7b63e405247ce1b153c808276999356234.NewDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the item of type microsoft.graph.directoryObject as microsoft.graph.servicePrincipal
func (m *ServicePrincipalRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *ServicePrincipalRequestBuilder) GetMemberGroups()(*id747d8a4068cca2fc00c2a191d305a94cbeb5d7b4c49db6ed85e3584cfc7dba8.GetMemberGroupsRequestBuilder) {
    return id747d8a4068cca2fc00c2a191d305a94cbeb5d7b4c49db6ed85e3584cfc7dba8.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *ServicePrincipalRequestBuilder) GetMemberObjects()(*i1f94bd1ec9ba9f8ce10416c3f8b76fb6da9dddd7ed173da40b5613d2b4aa52aa.GetMemberObjectsRequestBuilder) {
    return i1f94bd1ec9ba9f8ce10416c3f8b76fb6da9dddd7ed173da40b5613d2b4aa52aa.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPasswordSingleSignOnCredentials the getPasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) GetPasswordSingleSignOnCredentials()(*i92fdcf12e4e6aec7316f7be9c10c4e3f760f4ac80f7c8e4b42f8d9a875b525f5.GetPasswordSingleSignOnCredentialsRequestBuilder) {
    return i92fdcf12e4e6aec7316f7be9c10c4e3f760f4ac80f7c8e4b42f8d9a875b525f5.NewGetPasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ServicePrincipalRequestBuilder) Restore()(*i4cbadb93cdbf810c8551f7af31425d51c10f4daf3612aec0324280da87ef778a.RestoreRequestBuilder) {
    return i4cbadb93cdbf810c8551f7af31425d51c10f4daf3612aec0324280da87ef778a.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdatePasswordSingleSignOnCredentials the updatePasswordSingleSignOnCredentials property
func (m *ServicePrincipalRequestBuilder) UpdatePasswordSingleSignOnCredentials()(*i54ff55c1232b7f7a77d695ed3fa0635e493179f9021338e0b05db9ed5c67a19f.UpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    return i54ff55c1232b7f7a77d695ed3fa0635e493179f9021338e0b05db9ed5c67a19f.NewUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
