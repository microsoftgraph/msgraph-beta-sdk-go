package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i44234eb4e1f9cd47b539840bc18fbb166e3fabe51e7c344b8627004aa486ba33 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/androidmanagedappprotections/item/apps"
    i4778d96774d9e2ccca7803ef9b1eb7081dc9a8d0975b583eb2aa707cc4c66855 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/androidmanagedappprotections/item/deploymentsummary"
    i6fdd1eff96c393b127cac2aa1f215abb604a17faa6ac98f59cd063c3cb5f55c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/androidmanagedappprotections/item/apps/item"
)

// AndroidManagedAppProtectionRequestBuilder builds and executes requests for operations under \deviceAppManagement\androidManagedAppProtections\{androidManagedAppProtection-id}
type AndroidManagedAppProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AndroidManagedAppProtectionRequestBuilderDeleteOptions options for Delete
type AndroidManagedAppProtectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AndroidManagedAppProtectionRequestBuilderGetOptions options for Get
type AndroidManagedAppProtectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AndroidManagedAppProtectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AndroidManagedAppProtectionRequestBuilderGetQueryParameters android managed app policies.
type AndroidManagedAppProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AndroidManagedAppProtectionRequestBuilderPatchOptions options for Patch
type AndroidManagedAppProtectionRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidManagedAppProtection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AndroidManagedAppProtectionRequestBuilder) Apps()(*i44234eb4e1f9cd47b539840bc18fbb166e3fabe51e7c344b8627004aa486ba33.AppsRequestBuilder) {
    return i44234eb4e1f9cd47b539840bc18fbb166e3fabe51e7c344b8627004aa486ba33.NewAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.androidManagedAppProtections.item.apps.item collection
func (m *AndroidManagedAppProtectionRequestBuilder) AppsById(id string)(*i6fdd1eff96c393b127cac2aa1f215abb604a17faa6ac98f59cd063c3cb5f55c0.ManagedMobileAppRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp_id"] = id
    }
    return i6fdd1eff96c393b127cac2aa1f215abb604a17faa6ac98f59cd063c3cb5f55c0.NewManagedMobileAppRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAndroidManagedAppProtectionRequestBuilderInternal instantiates a new AndroidManagedAppProtectionRequestBuilder and sets the default values.
func NewAndroidManagedAppProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidManagedAppProtectionRequestBuilder) {
    m := &AndroidManagedAppProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/androidManagedAppProtections/{androidManagedAppProtection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAndroidManagedAppProtectionRequestBuilder instantiates a new AndroidManagedAppProtectionRequestBuilder and sets the default values.
func NewAndroidManagedAppProtectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AndroidManagedAppProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidManagedAppProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation android managed app policies.
func (m *AndroidManagedAppProtectionRequestBuilder) CreateDeleteRequestInformation(options *AndroidManagedAppProtectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation android managed app policies.
func (m *AndroidManagedAppProtectionRequestBuilder) CreateGetRequestInformation(options *AndroidManagedAppProtectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation android managed app policies.
func (m *AndroidManagedAppProtectionRequestBuilder) CreatePatchRequestInformation(options *AndroidManagedAppProtectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete android managed app policies.
func (m *AndroidManagedAppProtectionRequestBuilder) Delete(options *AndroidManagedAppProtectionRequestBuilderDeleteOptions)(error) {
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
func (m *AndroidManagedAppProtectionRequestBuilder) DeploymentSummary()(*i4778d96774d9e2ccca7803ef9b1eb7081dc9a8d0975b583eb2aa707cc4c66855.DeploymentSummaryRequestBuilder) {
    return i4778d96774d9e2ccca7803ef9b1eb7081dc9a8d0975b583eb2aa707cc4c66855.NewDeploymentSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get android managed app policies.
func (m *AndroidManagedAppProtectionRequestBuilder) Get(options *AndroidManagedAppProtectionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidManagedAppProtection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAndroidManagedAppProtection() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AndroidManagedAppProtection), nil
}
// Patch android managed app policies.
func (m *AndroidManagedAppProtectionRequestBuilder) Patch(options *AndroidManagedAppProtectionRequestBuilderPatchOptions)(error) {
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
