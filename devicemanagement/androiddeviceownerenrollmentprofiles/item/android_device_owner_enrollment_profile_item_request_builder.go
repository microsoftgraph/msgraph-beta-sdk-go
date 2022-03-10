package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i11bdf5bd909415fbbbbde18debe66a9b2f9f109ca57dc34d589ea53d118fd6bf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androiddeviceownerenrollmentprofiles/item/revoketoken"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    ibf7310ca68d55392bbdbd2da406d3078b3d0933d2dfaa641477e64ac9571c49c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/androiddeviceownerenrollmentprofiles/item/createtoken"
)

// AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
type AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteOptions options for Delete
type AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetOptions options for Get
type AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetQueryParameters android device owner enrollment profile entities.
type AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchOptions options for Patch
type AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidDeviceOwnerEnrollmentProfileable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal instantiates a new AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    m := &AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidDeviceOwnerEnrollmentProfiles/{androidDeviceOwnerEnrollmentProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder instantiates a new AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder and sets the default values.
func NewAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property androidDeviceOwnerEnrollmentProfiles for deviceManagement
func (m *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) CreateDeleteRequestInformation(options *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) CreateGetRequestInformation(options *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property androidDeviceOwnerEnrollmentProfiles in deviceManagement
func (m *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) CreatePatchRequestInformation(options *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) CreateToken()(*ibf7310ca68d55392bbdbd2da406d3078b3d0933d2dfaa641477e64ac9571c49c.CreateTokenRequestBuilder) {
    return ibf7310ca68d55392bbdbd2da406d3078b3d0933d2dfaa641477e64ac9571c49c.NewCreateTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property androidDeviceOwnerEnrollmentProfiles for deviceManagement
func (m *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) Delete(options *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get android device owner enrollment profile entities.
func (m *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) Get(options *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidDeviceOwnerEnrollmentProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidDeviceOwnerEnrollmentProfileable), nil
}
// Patch update the navigation property androidDeviceOwnerEnrollmentProfiles in deviceManagement
func (m *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) Patch(options *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) RevokeToken()(*i11bdf5bd909415fbbbbde18debe66a9b2f9f109ca57dc34d589ea53d118fd6bf.RevokeTokenRequestBuilder) {
    return i11bdf5bd909415fbbbbde18debe66a9b2f9f109ca57dc34d589ea53d118fd6bf.NewRevokeTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
