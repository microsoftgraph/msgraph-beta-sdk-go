package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4f1275f4e01aec44554e85132450051d071ff03fdc809dc755449b8d0c8b8cce "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/federationconfiguration"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i7d77397ebc07bdc6c288b9583ecb10228944b85ebba0bd59a369bf31c29e9a00 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/forcedelete"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DomainItemRequestBuilderDeleteOptions options for Delete
type DomainItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DomainItemRequestBuilderGetOptions options for Get
type DomainItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DomainItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DomainItemRequestBuilderGetQueryParameters get entity from domains by key
type DomainItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DomainItemRequestBuilderPatchOptions options for Patch
type DomainItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Domainable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDomainItemRequestBuilderInternal instantiates a new DomainItemRequestBuilder and sets the default values.
func NewDomainItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DomainItemRequestBuilder) {
    m := &DomainItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/domains/{domain_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDomainItemRequestBuilder instantiates a new DomainItemRequestBuilder and sets the default values.
func NewDomainItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DomainItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDomainItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from domains
func (m *DomainItemRequestBuilder) CreateDeleteRequestInformation(options *DomainItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from domains by key
func (m *DomainItemRequestBuilder) CreateGetRequestInformation(options *DomainItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in domains
func (m *DomainItemRequestBuilder) CreatePatchRequestInformation(options *DomainItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from domains
func (m *DomainItemRequestBuilder) Delete(options *DomainItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["directoryObject_id"] = id
    }
    return i7c32a78edd2862f410595847b33ec0d3f055fc93d0995e9ea57b50500fa73d0b.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["internalDomainFederation_id"] = id
    }
    return i314e0da51b4eb8df87f88943217b580fcf890eeed085a0c274988376943a7736.NewInternalDomainFederationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DomainItemRequestBuilder) ForceDelete()(*i7d77397ebc07bdc6c288b9583ecb10228944b85ebba0bd59a369bf31c29e9a00.ForceDeleteRequestBuilder) {
    return i7d77397ebc07bdc6c288b9583ecb10228944b85ebba0bd59a369bf31c29e9a00.NewForceDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entity from domains by key
func (m *DomainItemRequestBuilder) Get(options *DomainItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Domainable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDomainFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Domainable), nil
}
// Patch update entity in domains
func (m *DomainItemRequestBuilder) Patch(options *DomainItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["domainDnsRecord_id"] = id
    }
    return i578fb01ecf2dcedf797bac3faf12e5a4474158a85b4339b93e4f98ea08e63e84.NewDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["sharedEmailDomainInvitation_id"] = id
    }
    return i517b2d0ed5fffe13304488e9ea824f1d7f1a07b611ca3244ccfaedcfa39a15bf.NewSharedEmailDomainInvitationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["domainDnsRecord_id"] = id
    }
    return id2974620d1311e128536a57b4b13ac7c35b1363dee58758242283afc01d5b273.NewDomainDnsRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DomainItemRequestBuilder) Verify()(*ib09a60b5bfd30cb1a07acb8495a20e06efe52643f73c5ed84f3dffc3a1e541d7.VerifyRequestBuilder) {
    return ib09a60b5bfd30cb1a07acb8495a20e06efe52643f73c5ed84f3dffc3a1e541d7.NewVerifyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
