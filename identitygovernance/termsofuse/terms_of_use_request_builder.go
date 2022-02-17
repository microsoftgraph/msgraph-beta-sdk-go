package termsofuse

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i7a12547468a4d42ab8f3d3d3cd4e300f925a0ca2e7691adcd36d1a01331d01b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/termsofuse/agreements"
    idbdd55c4f874f351f01786070df721944ce5b3467ee40a9186d8e1d8254c5038 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/termsofuse/agreementacceptances"
    i4c625eec080c8046d067dfb89b7686880f5fffd058fe784176bdef86f3d98004 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/termsofuse/agreements/item"
    ib81ff97ec65f62c7e3bf45944864f15e6a426c2a95656377c94aedf6454d4d96 "github.com/microsoftgraph/msgraph-beta-sdk-go/identitygovernance/termsofuse/agreementacceptances/item"
)

// TermsOfUseRequestBuilder builds and executes requests for operations under \identityGovernance\termsOfUse
type TermsOfUseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TermsOfUseRequestBuilderDeleteOptions options for Delete
type TermsOfUseRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermsOfUseRequestBuilderGetOptions options for Get
type TermsOfUseRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TermsOfUseRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermsOfUseRequestBuilderGetQueryParameters get termsOfUse from identityGovernance
type TermsOfUseRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TermsOfUseRequestBuilderPatchOptions options for Patch
type TermsOfUseRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TermsOfUseContainer;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TermsOfUseRequestBuilder) AgreementAcceptances()(*idbdd55c4f874f351f01786070df721944ce5b3467ee40a9186d8e1d8254c5038.AgreementAcceptancesRequestBuilder) {
    return idbdd55c4f874f351f01786070df721944ce5b3467ee40a9186d8e1d8254c5038.NewAgreementAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementAcceptancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.termsOfUse.agreementAcceptances.item collection
func (m *TermsOfUseRequestBuilder) AgreementAcceptancesById(id string)(*ib81ff97ec65f62c7e3bf45944864f15e6a426c2a95656377c94aedf6454d4d96.AgreementAcceptanceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance_id"] = id
    }
    return ib81ff97ec65f62c7e3bf45944864f15e6a426c2a95656377c94aedf6454d4d96.NewAgreementAcceptanceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TermsOfUseRequestBuilder) Agreements()(*i7a12547468a4d42ab8f3d3d3cd4e300f925a0ca2e7691adcd36d1a01331d01b4.AgreementsRequestBuilder) {
    return i7a12547468a4d42ab8f3d3d3cd4e300f925a0ca2e7691adcd36d1a01331d01b4.NewAgreementsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgreementsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityGovernance.termsOfUse.agreements.item collection
func (m *TermsOfUseRequestBuilder) AgreementsById(id string)(*i4c625eec080c8046d067dfb89b7686880f5fffd058fe784176bdef86f3d98004.AgreementRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreement_id"] = id
    }
    return i4c625eec080c8046d067dfb89b7686880f5fffd058fe784176bdef86f3d98004.NewAgreementRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTermsOfUseRequestBuilderInternal instantiates a new TermsOfUseRequestBuilder and sets the default values.
func NewTermsOfUseRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermsOfUseRequestBuilder) {
    m := &TermsOfUseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/termsOfUse{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTermsOfUseRequestBuilder instantiates a new TermsOfUseRequestBuilder and sets the default values.
func NewTermsOfUseRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermsOfUseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsOfUseRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property termsOfUse for identityGovernance
func (m *TermsOfUseRequestBuilder) CreateDeleteRequestInformation(options *TermsOfUseRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get termsOfUse from identityGovernance
func (m *TermsOfUseRequestBuilder) CreateGetRequestInformation(options *TermsOfUseRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property termsOfUse in identityGovernance
func (m *TermsOfUseRequestBuilder) CreatePatchRequestInformation(options *TermsOfUseRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property termsOfUse for identityGovernance
func (m *TermsOfUseRequestBuilder) Delete(options *TermsOfUseRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get termsOfUse from identityGovernance
func (m *TermsOfUseRequestBuilder) Get(options *TermsOfUseRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TermsOfUseContainer, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTermsOfUseContainer() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TermsOfUseContainer), nil
}
// Patch update the navigation property termsOfUse in identityGovernance
func (m *TermsOfUseRequestBuilder) Patch(options *TermsOfUseRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
