package mdmwindowsinformationprotectionpolicies

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    if5c7a26625cafb004ed0e764dba42822957f316d6ddc05ff25f1effaad016720 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mdmwindowsinformationprotectionpolicies/haspayloadlinks"
    ifc6fca424d2fce33bd51c77e17bc106cfef5325f7e72318232be6188af43400f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mdmwindowsinformationprotectionpolicies/count"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// MdmWindowsInformationProtectionPoliciesRequestBuilder provides operations to manage the mdmWindowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
type MdmWindowsInformationProtectionPoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MdmWindowsInformationProtectionPoliciesRequestBuilderGetOptions options for Get
type MdmWindowsInformationProtectionPoliciesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MdmWindowsInformationProtectionPoliciesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MdmWindowsInformationProtectionPoliciesRequestBuilderGetQueryParameters windows information protection for apps running on devices which are MDM enrolled.
type MdmWindowsInformationProtectionPoliciesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// MdmWindowsInformationProtectionPoliciesRequestBuilderPostOptions options for Post
type MdmWindowsInformationProtectionPoliciesRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MdmWindowsInformationProtectionPolicyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewMdmWindowsInformationProtectionPoliciesRequestBuilderInternal instantiates a new MdmWindowsInformationProtectionPoliciesRequestBuilder and sets the default values.
func NewMdmWindowsInformationProtectionPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MdmWindowsInformationProtectionPoliciesRequestBuilder) {
    m := &MdmWindowsInformationProtectionPoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mdmWindowsInformationProtectionPolicies{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMdmWindowsInformationProtectionPoliciesRequestBuilder instantiates a new MdmWindowsInformationProtectionPoliciesRequestBuilder and sets the default values.
func NewMdmWindowsInformationProtectionPoliciesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MdmWindowsInformationProtectionPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMdmWindowsInformationProtectionPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MdmWindowsInformationProtectionPoliciesRequestBuilder) Count()(*ifc6fca424d2fce33bd51c77e17bc106cfef5325f7e72318232be6188af43400f.CountRequestBuilder) {
    return ifc6fca424d2fce33bd51c77e17bc106cfef5325f7e72318232be6188af43400f.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation windows information protection for apps running on devices which are MDM enrolled.
func (m *MdmWindowsInformationProtectionPoliciesRequestBuilder) CreateGetRequestInformation(options *MdmWindowsInformationProtectionPoliciesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to mdmWindowsInformationProtectionPolicies for deviceAppManagement
func (m *MdmWindowsInformationProtectionPoliciesRequestBuilder) CreatePostRequestInformation(options *MdmWindowsInformationProtectionPoliciesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Get windows information protection for apps running on devices which are MDM enrolled.
func (m *MdmWindowsInformationProtectionPoliciesRequestBuilder) Get(options *MdmWindowsInformationProtectionPoliciesRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MdmWindowsInformationProtectionPolicyCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateMdmWindowsInformationProtectionPolicyCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MdmWindowsInformationProtectionPolicyCollectionResponseable), nil
}
func (m *MdmWindowsInformationProtectionPoliciesRequestBuilder) HasPayloadLinks()(*if5c7a26625cafb004ed0e764dba42822957f316d6ddc05ff25f1effaad016720.HasPayloadLinksRequestBuilder) {
    return if5c7a26625cafb004ed0e764dba42822957f316d6ddc05ff25f1effaad016720.NewHasPayloadLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to mdmWindowsInformationProtectionPolicies for deviceAppManagement
func (m *MdmWindowsInformationProtectionPoliciesRequestBuilder) Post(options *MdmWindowsInformationProtectionPoliciesRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MdmWindowsInformationProtectionPolicyable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateMdmWindowsInformationProtectionPolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MdmWindowsInformationProtectionPolicyable), nil
}
