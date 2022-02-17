package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i327fcd640227f503f8954e6b42ff3020c935058951cc9cd59e3d13ef561d4a35 "github.com/microsoftgraph/msgraph-beta-sdk-go/agreements/item/acceptances"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i7c390e1de4aa07a80ae54615ecee39712acdc4aa46152235bc7e5a120d0c2b3e "github.com/microsoftgraph/msgraph-beta-sdk-go/agreements/item/files"
    ieb15810f180e10f80c2cdb1c4fd1740becd78855ed63b329c016adce90654749 "github.com/microsoftgraph/msgraph-beta-sdk-go/agreements/item/file"
    i3bee643e3932996f745cbce66aa7b1cb97137f56a746664bf5147711700962e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/agreements/item/acceptances/item"
    id2a53a68b053462b5a0a928c19838a8d68e478bd5246d11c35fbb140518f97d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/agreements/item/files/item"
)

// AgreementRequestBuilder builds and executes requests for operations under \agreements\{agreement-id}
type AgreementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AgreementRequestBuilderDeleteOptions options for Delete
type AgreementRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AgreementRequestBuilderGetOptions options for Get
type AgreementRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AgreementRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AgreementRequestBuilderGetQueryParameters get entity from agreements by key
type AgreementRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// AgreementRequestBuilderPatchOptions options for Patch
type AgreementRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Agreement;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AgreementRequestBuilder) Acceptances()(*i327fcd640227f503f8954e6b42ff3020c935058951cc9cd59e3d13ef561d4a35.AcceptancesRequestBuilder) {
    return i327fcd640227f503f8954e6b42ff3020c935058951cc9cd59e3d13ef561d4a35.NewAcceptancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.agreements.item.acceptances.item collection
func (m *AgreementRequestBuilder) AcceptancesById(id string)(*i3bee643e3932996f745cbce66aa7b1cb97137f56a746664bf5147711700962e8.AgreementAcceptanceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementAcceptance_id"] = id
    }
    return i3bee643e3932996f745cbce66aa7b1cb97137f56a746664bf5147711700962e8.NewAgreementAcceptanceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAgreementRequestBuilderInternal instantiates a new AgreementRequestBuilder and sets the default values.
func NewAgreementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AgreementRequestBuilder) {
    m := &AgreementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/agreements/{agreement_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAgreementRequestBuilder instantiates a new AgreementRequestBuilder and sets the default values.
func NewAgreementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AgreementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAgreementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from agreements
func (m *AgreementRequestBuilder) CreateDeleteRequestInformation(options *AgreementRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from agreements by key
func (m *AgreementRequestBuilder) CreateGetRequestInformation(options *AgreementRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in agreements
func (m *AgreementRequestBuilder) CreatePatchRequestInformation(options *AgreementRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from agreements
func (m *AgreementRequestBuilder) Delete(options *AgreementRequestBuilderDeleteOptions)(error) {
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
func (m *AgreementRequestBuilder) File()(*ieb15810f180e10f80c2cdb1c4fd1740becd78855ed63b329c016adce90654749.FileRequestBuilder) {
    return ieb15810f180e10f80c2cdb1c4fd1740becd78855ed63b329c016adce90654749.NewFileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AgreementRequestBuilder) Files()(*i7c390e1de4aa07a80ae54615ecee39712acdc4aa46152235bc7e5a120d0c2b3e.FilesRequestBuilder) {
    return i7c390e1de4aa07a80ae54615ecee39712acdc4aa46152235bc7e5a120d0c2b3e.NewFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.agreements.item.files.item collection
func (m *AgreementRequestBuilder) FilesById(id string)(*id2a53a68b053462b5a0a928c19838a8d68e478bd5246d11c35fbb140518f97d5.AgreementFileLocalizationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementFileLocalization_id"] = id
    }
    return id2a53a68b053462b5a0a928c19838a8d68e478bd5246d11c35fbb140518f97d5.NewAgreementFileLocalizationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from agreements by key
func (m *AgreementRequestBuilder) Get(options *AgreementRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Agreement, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAgreement() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Agreement), nil
}
// Patch update entity in agreements
func (m *AgreementRequestBuilder) Patch(options *AgreementRequestBuilderPatchOptions)(error) {
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
