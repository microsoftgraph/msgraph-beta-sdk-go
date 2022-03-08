package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i33749878be357a8a4e7ab7549dfb7ef00a61ae4756862d17fdbe0399ec25776c "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/stop"
    i3f4cdc08d50a52eb36625eb92c9b78c15ed747a07ec2270b3b04bdbeea41d91f "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/provisionondemand"
    i47d716a54a179680138af986e459c792376d7f4c6514002b3e70daf5a1a76fef "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/restart"
    i8400fe465eee24495a8d2b54b414e58dba9e8d403a34045a85f9ef8eb6613f0f "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/pause"
    i9e52a20ec6431f72acddd1559db2e4f0fd21df996b830a7113908ef411982967 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/schema"
    idaec5822ce9a73474ab924b2776d55984f7bb1b0ccd24f511556b91451979237 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/start"
    ifd48d70ab4c59f1f11d77d1818ba8659d5894f3575962284a88e156d2365ee90 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/validatecredentials"
)

// SynchronizationJobItemRequestBuilder provides operations to manage the jobs property of the microsoft.graph.synchronization entity.
type SynchronizationJobItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SynchronizationJobItemRequestBuilderDeleteOptions options for Delete
type SynchronizationJobItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SynchronizationJobItemRequestBuilderGetOptions options for Get
type SynchronizationJobItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SynchronizationJobItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SynchronizationJobItemRequestBuilderGetQueryParameters performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
type SynchronizationJobItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SynchronizationJobItemRequestBuilderPatchOptions options for Patch
type SynchronizationJobItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationJobable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSynchronizationJobItemRequestBuilderInternal instantiates a new SynchronizationJobItemRequestBuilder and sets the default values.
func NewSynchronizationJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationJobItemRequestBuilder) {
    m := &SynchronizationJobItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application_id}/synchronization/jobs/{synchronizationJob_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSynchronizationJobItemRequestBuilder instantiates a new SynchronizationJobItemRequestBuilder and sets the default values.
func NewSynchronizationJobItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property jobs for applications
func (m *SynchronizationJobItemRequestBuilder) CreateDeleteRequestInformation(options *SynchronizationJobItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *SynchronizationJobItemRequestBuilder) CreateGetRequestInformation(options *SynchronizationJobItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property jobs in applications
func (m *SynchronizationJobItemRequestBuilder) CreatePatchRequestInformation(options *SynchronizationJobItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property jobs for applications
func (m *SynchronizationJobItemRequestBuilder) Delete(options *SynchronizationJobItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *SynchronizationJobItemRequestBuilder) Get(options *SynchronizationJobItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationJobable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateSynchronizationJobFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationJobable), nil
}
// Patch update the navigation property jobs in applications
func (m *SynchronizationJobItemRequestBuilder) Patch(options *SynchronizationJobItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *SynchronizationJobItemRequestBuilder) Pause()(*i8400fe465eee24495a8d2b54b414e58dba9e8d403a34045a85f9ef8eb6613f0f.PauseRequestBuilder) {
    return i8400fe465eee24495a8d2b54b414e58dba9e8d403a34045a85f9ef8eb6613f0f.NewPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobItemRequestBuilder) ProvisionOnDemand()(*i3f4cdc08d50a52eb36625eb92c9b78c15ed747a07ec2270b3b04bdbeea41d91f.ProvisionOnDemandRequestBuilder) {
    return i3f4cdc08d50a52eb36625eb92c9b78c15ed747a07ec2270b3b04bdbeea41d91f.NewProvisionOnDemandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobItemRequestBuilder) Restart()(*i47d716a54a179680138af986e459c792376d7f4c6514002b3e70daf5a1a76fef.RestartRequestBuilder) {
    return i47d716a54a179680138af986e459c792376d7f4c6514002b3e70daf5a1a76fef.NewRestartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobItemRequestBuilder) Schema()(*i9e52a20ec6431f72acddd1559db2e4f0fd21df996b830a7113908ef411982967.SchemaRequestBuilder) {
    return i9e52a20ec6431f72acddd1559db2e4f0fd21df996b830a7113908ef411982967.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobItemRequestBuilder) Start()(*idaec5822ce9a73474ab924b2776d55984f7bb1b0ccd24f511556b91451979237.StartRequestBuilder) {
    return idaec5822ce9a73474ab924b2776d55984f7bb1b0ccd24f511556b91451979237.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobItemRequestBuilder) Stop()(*i33749878be357a8a4e7ab7549dfb7ef00a61ae4756862d17fdbe0399ec25776c.StopRequestBuilder) {
    return i33749878be357a8a4e7ab7549dfb7ef00a61ae4756862d17fdbe0399ec25776c.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobItemRequestBuilder) ValidateCredentials()(*ifd48d70ab4c59f1f11d77d1818ba8659d5894f3575962284a88e156d2365ee90.ValidateCredentialsRequestBuilder) {
    return ifd48d70ab4c59f1f11d77d1818ba8659d5894f3575962284a88e156d2365ee90.NewValidateCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
