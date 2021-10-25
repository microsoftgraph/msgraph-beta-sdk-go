package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i7d77397ebc07bdc6c288b9583ecb10228944b85ebba0bd59a369bf31c29e9a00 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/forcedelete"
    iac35370406c87d750e8103c3720efe8ef0269dc234cd57664c9856322fb95dfc "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/sharedemaildomaininvitations"
    ib09a60b5bfd30cb1a07acb8495a20e06efe52643f73c5ed84f3dffc3a1e541d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/verify"
    ibf1ebf19d88af60c12e96615e8bf6cf8bd52bc3a35c310ec3209c033fa8a3202 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/verificationdnsrecords"
    icf8071f24e1196527a85c6af2bcf893b18ff56c62aeca03cd4e6206d6ddd5c00 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/serviceconfigurationrecords"
    id06f08aa20a5ea0585b12bac772b2701826427a0557268d532a009c909a63c61 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/domainnamereferences"
    i517b2d0ed5fffe13304488e9ea824f1d7f1a07b611ca3244ccfaedcfa39a15bf "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/sharedemaildomaininvitations/item"
    i578fb01ecf2dcedf797bac3faf12e5a4474158a85b4339b93e4f98ea08e63e84 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/serviceconfigurationrecords/item"
    id2974620d1311e128536a57b4b13ac7c35b1363dee58758242283afc01d5b273 "github.com/microsoftgraph/msgraph-beta-sdk-go/domains/item/verificationdnsrecords/item"
)

type DomainRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DomainRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func NewDomainRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DomainRequestBuilder) {
    m := &DomainRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/domains/{domain_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewDomainRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DomainRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDomainRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DomainRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DomainRequestBuilder) CreateGetRequestInformation(q func (value *DomainRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DomainRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DomainRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Domain, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DomainRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DomainRequestBuilder) DomainNameReferences()(*id06f08aa20a5ea0585b12bac772b2701826427a0557268d532a009c909a63c61.DomainNameReferencesRequestBuilder) {
    return id06f08aa20a5ea0585b12bac772b2701826427a0557268d532a009c909a63c61.NewDomainNameReferencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DomainRequestBuilder) ForceDelete()(*i7d77397ebc07bdc6c288b9583ecb10228944b85ebba0bd59a369bf31c29e9a00.ForceDeleteRequestBuilder) {
    return i7d77397ebc07bdc6c288b9583ecb10228944b85ebba0bd59a369bf31c29e9a00.NewForceDeleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DomainRequestBuilder) Get(q func (value *DomainRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Domain, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDomain() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Domain), nil
}
func (m *DomainRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Domain, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DomainRequestBuilder) ServiceConfigurationRecords()(*icf8071f24e1196527a85c6af2bcf893b18ff56c62aeca03cd4e6206d6ddd5c00.ServiceConfigurationRecordsRequestBuilder) {
    return icf8071f24e1196527a85c6af2bcf893b18ff56c62aeca03cd4e6206d6ddd5c00.NewServiceConfigurationRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DomainRequestBuilder) ServiceConfigurationRecordsById(id string)(*i578fb01ecf2dcedf797bac3faf12e5a4474158a85b4339b93e4f98ea08e63e84.DomainDnsRecordRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainDnsRecord_id"] = id
    }
    return i578fb01ecf2dcedf797bac3faf12e5a4474158a85b4339b93e4f98ea08e63e84.NewDomainDnsRecordRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DomainRequestBuilder) SharedEmailDomainInvitations()(*iac35370406c87d750e8103c3720efe8ef0269dc234cd57664c9856322fb95dfc.SharedEmailDomainInvitationsRequestBuilder) {
    return iac35370406c87d750e8103c3720efe8ef0269dc234cd57664c9856322fb95dfc.NewSharedEmailDomainInvitationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DomainRequestBuilder) SharedEmailDomainInvitationsById(id string)(*i517b2d0ed5fffe13304488e9ea824f1d7f1a07b611ca3244ccfaedcfa39a15bf.SharedEmailDomainInvitationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sharedEmailDomainInvitation_id"] = id
    }
    return i517b2d0ed5fffe13304488e9ea824f1d7f1a07b611ca3244ccfaedcfa39a15bf.NewSharedEmailDomainInvitationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DomainRequestBuilder) VerificationDnsRecords()(*ibf1ebf19d88af60c12e96615e8bf6cf8bd52bc3a35c310ec3209c033fa8a3202.VerificationDnsRecordsRequestBuilder) {
    return ibf1ebf19d88af60c12e96615e8bf6cf8bd52bc3a35c310ec3209c033fa8a3202.NewVerificationDnsRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DomainRequestBuilder) VerificationDnsRecordsById(id string)(*id2974620d1311e128536a57b4b13ac7c35b1363dee58758242283afc01d5b273.DomainDnsRecordRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainDnsRecord_id"] = id
    }
    return id2974620d1311e128536a57b4b13ac7c35b1363dee58758242283afc01d5b273.NewDomainDnsRecordRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DomainRequestBuilder) Verify()(*ib09a60b5bfd30cb1a07acb8495a20e06efe52643f73c5ed84f3dffc3a1e541d7.VerifyRequestBuilder) {
    return ib09a60b5bfd30cb1a07acb8495a20e06efe52643f73c5ed84f3dffc3a1e541d7.NewVerifyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
