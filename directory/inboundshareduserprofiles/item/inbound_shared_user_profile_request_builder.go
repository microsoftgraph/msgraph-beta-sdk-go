package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i03f81bc977f9d90efbd1427f3444ba111baf8b4fe1f097fb2522abf9ca22efc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/inboundshareduserprofiles/item/exportpersonaldata"
    i2a5ab3f902913f407fab17927469dfb7cf3909152c8e32eaa79d54021bc94cb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/directory/inboundshareduserprofiles/item/removepersonaldata"
)

// InboundSharedUserProfileRequestBuilder builds and executes requests for operations under \directory\inboundSharedUserProfiles\{inboundSharedUserProfile-userId}
type InboundSharedUserProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// InboundSharedUserProfileRequestBuilderDeleteOptions options for Delete
type InboundSharedUserProfileRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// InboundSharedUserProfileRequestBuilderGetOptions options for Get
type InboundSharedUserProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *InboundSharedUserProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// InboundSharedUserProfileRequestBuilderGetQueryParameters get inboundSharedUserProfiles from directory
type InboundSharedUserProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// InboundSharedUserProfileRequestBuilderPatchOptions options for Patch
type InboundSharedUserProfileRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InboundSharedUserProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewInboundSharedUserProfileRequestBuilderInternal instantiates a new InboundSharedUserProfileRequestBuilder and sets the default values.
func NewInboundSharedUserProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InboundSharedUserProfileRequestBuilder) {
    m := &InboundSharedUserProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/inboundSharedUserProfiles/{inboundSharedUserProfile_userId}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInboundSharedUserProfileRequestBuilder instantiates a new InboundSharedUserProfileRequestBuilder and sets the default values.
func NewInboundSharedUserProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*InboundSharedUserProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInboundSharedUserProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property inboundSharedUserProfiles for directory
func (m *InboundSharedUserProfileRequestBuilder) CreateDeleteRequestInformation(options *InboundSharedUserProfileRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get inboundSharedUserProfiles from directory
func (m *InboundSharedUserProfileRequestBuilder) CreateGetRequestInformation(options *InboundSharedUserProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property inboundSharedUserProfiles in directory
func (m *InboundSharedUserProfileRequestBuilder) CreatePatchRequestInformation(options *InboundSharedUserProfileRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property inboundSharedUserProfiles for directory
func (m *InboundSharedUserProfileRequestBuilder) Delete(options *InboundSharedUserProfileRequestBuilderDeleteOptions)(error) {
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
func (m *InboundSharedUserProfileRequestBuilder) ExportPersonalData()(*i03f81bc977f9d90efbd1427f3444ba111baf8b4fe1f097fb2522abf9ca22efc4.ExportPersonalDataRequestBuilder) {
    return i03f81bc977f9d90efbd1427f3444ba111baf8b4fe1f097fb2522abf9ca22efc4.NewExportPersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get inboundSharedUserProfiles from directory
func (m *InboundSharedUserProfileRequestBuilder) Get(options *InboundSharedUserProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InboundSharedUserProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewInboundSharedUserProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InboundSharedUserProfile), nil
}
// Patch update the navigation property inboundSharedUserProfiles in directory
func (m *InboundSharedUserProfileRequestBuilder) Patch(options *InboundSharedUserProfileRequestBuilderPatchOptions)(error) {
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
func (m *InboundSharedUserProfileRequestBuilder) RemovePersonalData()(*i2a5ab3f902913f407fab17927469dfb7cf3909152c8e32eaa79d54021bc94cb3.RemovePersonalDataRequestBuilder) {
    return i2a5ab3f902913f407fab17927469dfb7cf3909152c8e32eaa79d54021bc94cb3.NewRemovePersonalDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
