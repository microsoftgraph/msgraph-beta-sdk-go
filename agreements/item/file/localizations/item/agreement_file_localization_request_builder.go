package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1bfeb14b00e82ba273a6ce7f098d8c2fb6b8070aa6baf325ecc73cedd5776311 "github.com/microsoftgraph/msgraph-beta-sdk-go/agreements/item/file/localizations/item/versions"
    i59b8285d8b9da798aec45fee2e33d08b5aefbd44b6ea832ee6113712ad4d9f25 "github.com/microsoftgraph/msgraph-beta-sdk-go/agreements/item/file/localizations/item/versions/item"
)

// AgreementFileLocalizationRequestBuilder builds and executes requests for operations under \agreements\{agreement-id}\file\localizations\{agreementFileLocalization-id}
type AgreementFileLocalizationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AgreementFileLocalizationRequestBuilderDeleteOptions options for Delete
type AgreementFileLocalizationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AgreementFileLocalizationRequestBuilderGetOptions options for Get
type AgreementFileLocalizationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AgreementFileLocalizationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AgreementFileLocalizationRequestBuilderGetQueryParameters the localized version of the terms of use agreement files attached to the agreement.
type AgreementFileLocalizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AgreementFileLocalizationRequestBuilderPatchOptions options for Patch
type AgreementFileLocalizationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AgreementFileLocalization;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAgreementFileLocalizationRequestBuilderInternal instantiates a new AgreementFileLocalizationRequestBuilder and sets the default values.
func NewAgreementFileLocalizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AgreementFileLocalizationRequestBuilder) {
    m := &AgreementFileLocalizationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/agreements/{agreement_id}/file/localizations/{agreementFileLocalization_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAgreementFileLocalizationRequestBuilder instantiates a new AgreementFileLocalizationRequestBuilder and sets the default values.
func NewAgreementFileLocalizationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AgreementFileLocalizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAgreementFileLocalizationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the localized version of the terms of use agreement files attached to the agreement.
func (m *AgreementFileLocalizationRequestBuilder) CreateDeleteRequestInformation(options *AgreementFileLocalizationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the localized version of the terms of use agreement files attached to the agreement.
func (m *AgreementFileLocalizationRequestBuilder) CreateGetRequestInformation(options *AgreementFileLocalizationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the localized version of the terms of use agreement files attached to the agreement.
func (m *AgreementFileLocalizationRequestBuilder) CreatePatchRequestInformation(options *AgreementFileLocalizationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the localized version of the terms of use agreement files attached to the agreement.
func (m *AgreementFileLocalizationRequestBuilder) Delete(options *AgreementFileLocalizationRequestBuilderDeleteOptions)(error) {
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
// Get the localized version of the terms of use agreement files attached to the agreement.
func (m *AgreementFileLocalizationRequestBuilder) Get(options *AgreementFileLocalizationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AgreementFileLocalization, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAgreementFileLocalization() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AgreementFileLocalization), nil
}
// Patch the localized version of the terms of use agreement files attached to the agreement.
func (m *AgreementFileLocalizationRequestBuilder) Patch(options *AgreementFileLocalizationRequestBuilderPatchOptions)(error) {
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
func (m *AgreementFileLocalizationRequestBuilder) Versions()(*i1bfeb14b00e82ba273a6ce7f098d8c2fb6b8070aa6baf325ecc73cedd5776311.VersionsRequestBuilder) {
    return i1bfeb14b00e82ba273a6ce7f098d8c2fb6b8070aa6baf325ecc73cedd5776311.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.agreements.item.file.localizations.item.versions.item collection
func (m *AgreementFileLocalizationRequestBuilder) VersionsById(id string)(*i59b8285d8b9da798aec45fee2e33d08b5aefbd44b6ea832ee6113712ad4d9f25.AgreementFileVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["agreementFileVersion_id"] = id
    }
    return i59b8285d8b9da798aec45fee2e33d08b5aefbd44b6ea832ee6113712ad4d9f25.NewAgreementFileVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
