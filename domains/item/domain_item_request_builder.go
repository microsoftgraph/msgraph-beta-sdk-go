package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4f1275f4e01aec44554e85132450051d071ff03fdc809dc755449b8d0c8b8cce "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/federationconfiguration"
    i7d77397ebc07bdc6c288b9583ecb10228944b85ebba0bd59a369bf31c29e9a00 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/forcedelete"
    i91dcbb36516a1f64b3eb5efbaaf9ee0b1fd80919b325ceacd3295464b98a51db "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/promote"
    iac35370406c87d750e8103c3720efe8ef0269dc234cd57664c9856322fb95dfc "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/sharedemaildomaininvitations"
    ib09a60b5bfd30cb1a07acb8495a20e06efe52643f73c5ed84f3dffc3a1e541d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/verify"
    ibf1ebf19d88af60c12e96615e8bf6cf8bd52bc3a35c310ec3209c033fa8a3202 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/verificationdnsrecords"
    icf8071f24e1196527a85c6af2bcf893b18ff56c62aeca03cd4e6206d6ddd5c00 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/serviceconfigurationrecords"
    id06f08aa20a5ea0585b12bac772b2701826427a0557268d532a009c909a63c61 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/domainnamereferences"
    i314e0da51b4eb8df87f88943217b580fcf890eeed085a0c274988376943a7736 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/federationconfiguration/item"
    i517b2d0ed5fffe13304488e9ea824f1d7f1a07b611ca3244ccfaedcfa39a15bf "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/sharedemaildomaininvitations/item"
    i578fb01ecf2dcedf797bac3faf12e5a4474158a85b4339b93e4f98ea08e63e84 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/serviceconfigurationrecords/item"
    i7c32a78edd2862f410595847b33ec0d3f055fc93d0995e9ea57b50500fa73d0b "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/domainnamereferences/item"
    id2974620d1311e128536a57b4b13ac7c35b1363dee58758242283afc01d5b273 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/verificationdnsrecords/item"
)

// DomainItemRequestBuilder provides operations to manage the collection of domain entities.
type DomainItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DomainItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DomainItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DomainItemRequestBuilderGetQueryParameters retrieve the properties and relationships of domain object.
type DomainItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DomainItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DomainItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DomainItemRequestBuilderGetQueryParameters
}
// DomainItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DomainItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDomainItemRequestBuilderInternal instantiates a new DomainItemRequestBuilder and sets the default values.
func NewDomainItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DomainItemRequestBuilder) {
    m := &DomainItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/domains/{domain%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDomainItemRequestBuilder instantiates a new DomainItemRequestBuilder and sets the default values.
func NewDomainItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DomainItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDomainItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation deletes a domain from a tenant.
func (m *DomainItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration deletes a domain from a tenant.
func (m *DomainItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DomainItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation retrieve the properties and relationships of domain object.
func (m *DomainItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration retrieve the properties and relationships of domain object.
func (m *DomainItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DomainItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of domain object.
func (m *DomainItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Domainable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the properties of domain object.
func (m *DomainItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Domainable, requestConfiguration *DomainItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete deletes a domain from a tenant.
func (m *DomainItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DomainItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// DomainNameReferences the domainNameReferences property
func (m *DomainItemRequestBuilder) DomainNameReferences()(*id06f08aa20a5ea0585b12bac772b2701826427a0557268d532a009c909a63c61.DomainNameReferencesRequestBuilder) {
    return id06f08aa20a5ea0585b12bac772b2701826427a0557268d532a009c909a63c61.NewDomainNameReferencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainNameReferencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.domains.item.domainNameReferences.item collection
func (m *DomainItemRequestBuilder) DomainNameReferencesById(id string)(*i7c32a78edd2862f410595847b33ec0d3f055fc93d0995e9ea57b50500fa73d0b.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return i7c32a78edd2862f410595847b33ec0d3f055fc93d0995e9ea57b50500fa73d0b.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FederationConfiguration the federationConfiguration property
func (m *DomainItemRequestBuilder) FederationConfiguration()(*i4f1275f4e01aec44554e85132450051d071ff03fdc809dc755449b8d0c8b8cce.FederationConfigurationRequestBuilder) {
    return i4f1275f4e01aec44554e85132450051d071ff03fdc809dc755449b8d0c8b8cce.NewFederationConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FederationConfigurationById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.domains.item.federationConfiguration.item collection
func (m *DomainItemRequestBuilder) FederationConfigurationById(id string)(*i314e0da51b4eb8df87f88943217b580fcf890eeed085a0c274988376943a7736.InternalDomainFederationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["internalDomainFederation%2Did"] = id
    }
    return i314e0da51b4eb8df87f88943217b580fcf890eeed085a0c274988376943a7736.NewInternalDomainFederationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ForceDelete the forceDelete property
func (m *DomainItemRequestBuilder) ForceDelete()(*i7d77397ebc07bdc6c288b9583ecb10228944b85ebba0bd59a369bf31c29e9a00.ForceDeleteRequestBuilder) {
    return i7d77397ebc07bdc6c288b9583ecb10228944b85ebba0bd59a369bf31c29e9a00.NewForceDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve the properties and relationships of domain object.
func (m *DomainItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DomainItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Domainable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDomainFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Domainable), nil
}
// Patch update the properties of domain object.
func (m *DomainItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Domainable, requestConfiguration *DomainItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Promote the promote property
func (m *DomainItemRequestBuilder) Promote()(*i91dcbb36516a1f64b3eb5efbaaf9ee0b1fd80919b325ceacd3295464b98a51db.PromoteRequestBuilder) {
    return i91dcbb36516a1f64b3eb5efbaaf9ee0b1fd80919b325ceacd3295464b98a51db.NewPromoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServiceConfigurationRecords the serviceConfigurationRecords property
func (m *DomainItemRequestBuilder) ServiceConfigurationRecords()(*icf8071f24e1196527a85c6af2bcf893b18ff56c62aeca03cd4e6206d6ddd5c00.ServiceConfigurationRecordsRequestBuilder) {
    return icf8071f24e1196527a85c6af2bcf893b18ff56c62aeca03cd4e6206d6ddd5c00.NewServiceConfigurationRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServiceConfigurationRecordsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.domains.item.serviceConfigurationRecords.item collection
func (m *DomainItemRequestBuilder) ServiceConfigurationRecordsById(id string)(*i578fb01ecf2dcedf797bac3faf12e5a4474158a85b4339b93e4f98ea08e63e84.DomainDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainDnsRecord%2Did"] = id
    }
    return i578fb01ecf2dcedf797bac3faf12e5a4474158a85b4339b93e4f98ea08e63e84.NewDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SharedEmailDomainInvitations the sharedEmailDomainInvitations property
func (m *DomainItemRequestBuilder) SharedEmailDomainInvitations()(*iac35370406c87d750e8103c3720efe8ef0269dc234cd57664c9856322fb95dfc.SharedEmailDomainInvitationsRequestBuilder) {
    return iac35370406c87d750e8103c3720efe8ef0269dc234cd57664c9856322fb95dfc.NewSharedEmailDomainInvitationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SharedEmailDomainInvitationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.domains.item.sharedEmailDomainInvitations.item collection
func (m *DomainItemRequestBuilder) SharedEmailDomainInvitationsById(id string)(*i517b2d0ed5fffe13304488e9ea824f1d7f1a07b611ca3244ccfaedcfa39a15bf.SharedEmailDomainInvitationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedEmailDomainInvitation%2Did"] = id
    }
    return i517b2d0ed5fffe13304488e9ea824f1d7f1a07b611ca3244ccfaedcfa39a15bf.NewSharedEmailDomainInvitationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VerificationDnsRecords the verificationDnsRecords property
func (m *DomainItemRequestBuilder) VerificationDnsRecords()(*ibf1ebf19d88af60c12e96615e8bf6cf8bd52bc3a35c310ec3209c033fa8a3202.VerificationDnsRecordsRequestBuilder) {
    return ibf1ebf19d88af60c12e96615e8bf6cf8bd52bc3a35c310ec3209c033fa8a3202.NewVerificationDnsRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VerificationDnsRecordsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.domains.item.verificationDnsRecords.item collection
func (m *DomainItemRequestBuilder) VerificationDnsRecordsById(id string)(*id2974620d1311e128536a57b4b13ac7c35b1363dee58758242283afc01d5b273.DomainDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainDnsRecord%2Did"] = id
    }
    return id2974620d1311e128536a57b4b13ac7c35b1363dee58758242283afc01d5b273.NewDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Verify the verify property
func (m *DomainItemRequestBuilder) Verify()(*ib09a60b5bfd30cb1a07acb8495a20e06efe52643f73c5ed84f3dffc3a1e541d7.VerifyRequestBuilder) {
    return ib09a60b5bfd30cb1a07acb8495a20e06efe52643f73c5ed84f3dffc3a1e541d7.NewVerifyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
