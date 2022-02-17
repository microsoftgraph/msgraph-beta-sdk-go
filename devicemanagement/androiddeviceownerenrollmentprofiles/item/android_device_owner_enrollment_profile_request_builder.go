package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i11bdf5bd909415fbbbbde18debe66a9b2f9f109ca57dc34d589ea53d118fd6bf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androiddeviceownerenrollmentprofiles/item/revoketoken"
    ibf7310ca68d55392bbdbd2da406d3078b3d0933d2dfaa641477e64ac9571c49c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androiddeviceownerenrollmentprofiles/item/createtoken"
)

// AndroidDeviceOwnerEnrollmentProfileRequestBuilder builds and executes requests for operations under \deviceManagement\androidDeviceOwnerEnrollmentProfiles\{androidDeviceOwnerEnrollmentProfile-id}
type AndroidDeviceOwnerEnrollmentProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AndroidDeviceOwnerEnrollmentProfileRequestBuilderDeleteOptions options for Delete
type AndroidDeviceOwnerEnrollmentProfileRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AndroidDeviceOwnerEnrollmentProfileRequestBuilderGetOptions options for Get
type AndroidDeviceOwnerEnrollmentProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AndroidDeviceOwnerEnrollmentProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AndroidDeviceOwnerEnrollmentProfileRequestBuilderGetQueryParameters android device owner enrollment profile entities.
type AndroidDeviceOwnerEnrollmentProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AndroidDeviceOwnerEnrollmentProfileRequestBuilderPatchOptions options for Patch
type AndroidDeviceOwnerEnrollmentProfileRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidDeviceOwnerEnrollmentProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAndroidDeviceOwnerEnrollmentProfileRequestBuilderInternal instantiates a new AndroidDeviceOwnerEnrollmentProfileRequestBuilder and sets the default values.
func NewAndroidDeviceOwnerEnrollmentProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidDeviceOwnerEnrollmentProfileRequestBuilder) {
    m := &AndroidDeviceOwnerEnrollmentProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidDeviceOwnerEnrollmentProfiles/{androidDeviceOwnerEnrollmentProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAndroidDeviceOwnerEnrollmentProfileRequestBuilder instantiates a new AndroidDeviceOwnerEnrollmentProfileRequestBuilder and sets the default values.
func NewAndroidDeviceOwnerEnrollmentProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidDeviceOwnerEnrollmentProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidDeviceOwnerEnrollmentProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation android device owner enrollment profile entities.
func (m *AndroidDeviceOwnerEnrollmentProfileRequestBuilder) CreateDeleteRequestInformation(options *AndroidDeviceOwnerEnrollmentProfileRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation android device owner enrollment profile entities.
func (m *AndroidDeviceOwnerEnrollmentProfileRequestBuilder) CreateGetRequestInformation(options *AndroidDeviceOwnerEnrollmentProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation android device owner enrollment profile entities.
func (m *AndroidDeviceOwnerEnrollmentProfileRequestBuilder) CreatePatchRequestInformation(options *AndroidDeviceOwnerEnrollmentProfileRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AndroidDeviceOwnerEnrollmentProfileRequestBuilder) CreateToken()(*ibf7310ca68d55392bbdbd2da406d3078b3d0933d2dfaa641477e64ac9571c49c.CreateTokenRequestBuilder) {
    return ibf7310ca68d55392bbdbd2da406d3078b3d0933d2dfaa641477e64ac9571c49c.NewCreateTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete android device owner enrollment profile entities.
func (m *AndroidDeviceOwnerEnrollmentProfileRequestBuilder) Delete(options *AndroidDeviceOwnerEnrollmentProfileRequestBuilderDeleteOptions)(error) {
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
// Get android device owner enrollment profile entities.
func (m *AndroidDeviceOwnerEnrollmentProfileRequestBuilder) Get(options *AndroidDeviceOwnerEnrollmentProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidDeviceOwnerEnrollmentProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAndroidDeviceOwnerEnrollmentProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidDeviceOwnerEnrollmentProfile), nil
}
// Patch android device owner enrollment profile entities.
func (m *AndroidDeviceOwnerEnrollmentProfileRequestBuilder) Patch(options *AndroidDeviceOwnerEnrollmentProfileRequestBuilderPatchOptions)(error) {
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
func (m *AndroidDeviceOwnerEnrollmentProfileRequestBuilder) RevokeToken()(*i11bdf5bd909415fbbbbde18debe66a9b2f9f109ca57dc34d589ea53d118fd6bf.RevokeTokenRequestBuilder) {
    return i11bdf5bd909415fbbbbde18debe66a9b2f9f109ca57dc34d589ea53d118fd6bf.NewRevokeTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
