package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i04d76fbaad9f08921b0fefffd4ae909a8fbadb2da7b557711ada64363e2e763d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/profilestatus"
    i9915f3a66e2de398bdfae23c9b58bdf2c4a7c373437f422e62cc19ab1b37c01d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/errors"
    ibd324584f412d6082609e09d36338222a1ea33342c17f036ead958c71c7e57d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/resume"
    ibfdf2171a3f86b83740d33b9eb574ed52ef8e274f8b0c78649787ec4351bfa5d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/pause"
    icea2a0d7fcf828c9af01c7bba86f54c7305fbb4d6c4cea549b5ee5ece1723710 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/uploadurl"
    id814594e10d99c016c54d9229da3f00c037b42a01b52dd4efad6b8da5d373fb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/reset"
    if70ade9d653b89ed0849b345d9e94a2327f9abc57ba5b75b6511e5e4013d1aea "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/start"
    i88e077efbe7af1476bbdd89a3d093be25e50e57c1d6a1e16d2b9f2406ba01c90 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item/errors/item"
)

// EducationSynchronizationProfileRequestBuilder builds and executes requests for operations under \education\synchronizationProfiles\{educationSynchronizationProfile-id}
type EducationSynchronizationProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationSynchronizationProfileRequestBuilderDeleteOptions options for Delete
type EducationSynchronizationProfileRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationSynchronizationProfileRequestBuilderGetOptions options for Get
type EducationSynchronizationProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationSynchronizationProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationSynchronizationProfileRequestBuilderGetQueryParameters get synchronizationProfiles from education
type EducationSynchronizationProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationSynchronizationProfileRequestBuilderPatchOptions options for Patch
type EducationSynchronizationProfileRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSynchronizationProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewEducationSynchronizationProfileRequestBuilderInternal instantiates a new EducationSynchronizationProfileRequestBuilder and sets the default values.
func NewEducationSynchronizationProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSynchronizationProfileRequestBuilder) {
    m := &EducationSynchronizationProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/synchronizationProfiles/{educationSynchronizationProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationSynchronizationProfileRequestBuilder instantiates a new EducationSynchronizationProfileRequestBuilder and sets the default values.
func NewEducationSynchronizationProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSynchronizationProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSynchronizationProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property synchronizationProfiles for education
func (m *EducationSynchronizationProfileRequestBuilder) CreateDeleteRequestInformation(options *EducationSynchronizationProfileRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get synchronizationProfiles from education
func (m *EducationSynchronizationProfileRequestBuilder) CreateGetRequestInformation(options *EducationSynchronizationProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property synchronizationProfiles in education
func (m *EducationSynchronizationProfileRequestBuilder) CreatePatchRequestInformation(options *EducationSynchronizationProfileRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property synchronizationProfiles for education
func (m *EducationSynchronizationProfileRequestBuilder) Delete(options *EducationSynchronizationProfileRequestBuilderDeleteOptions)(error) {
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
func (m *EducationSynchronizationProfileRequestBuilder) Errors()(*i9915f3a66e2de398bdfae23c9b58bdf2c4a7c373437f422e62cc19ab1b37c01d.ErrorsRequestBuilder) {
    return i9915f3a66e2de398bdfae23c9b58bdf2c4a7c373437f422e62cc19ab1b37c01d.NewErrorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ErrorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.synchronizationProfiles.item.errors.item collection
func (m *EducationSynchronizationProfileRequestBuilder) ErrorsById(id string)(*i88e077efbe7af1476bbdd89a3d093be25e50e57c1d6a1e16d2b9f2406ba01c90.EducationSynchronizationErrorRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSynchronizationError_id"] = id
    }
    return i88e077efbe7af1476bbdd89a3d093be25e50e57c1d6a1e16d2b9f2406ba01c90.NewEducationSynchronizationErrorRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get synchronizationProfiles from education
func (m *EducationSynchronizationProfileRequestBuilder) Get(options *EducationSynchronizationProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSynchronizationProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEducationSynchronizationProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSynchronizationProfile), nil
}
// Patch update the navigation property synchronizationProfiles in education
func (m *EducationSynchronizationProfileRequestBuilder) Patch(options *EducationSynchronizationProfileRequestBuilderPatchOptions)(error) {
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
func (m *EducationSynchronizationProfileRequestBuilder) Pause()(*ibfdf2171a3f86b83740d33b9eb574ed52ef8e274f8b0c78649787ec4351bfa5d.PauseRequestBuilder) {
    return ibfdf2171a3f86b83740d33b9eb574ed52ef8e274f8b0c78649787ec4351bfa5d.NewPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSynchronizationProfileRequestBuilder) ProfileStatus()(*i04d76fbaad9f08921b0fefffd4ae909a8fbadb2da7b557711ada64363e2e763d.ProfileStatusRequestBuilder) {
    return i04d76fbaad9f08921b0fefffd4ae909a8fbadb2da7b557711ada64363e2e763d.NewProfileStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSynchronizationProfileRequestBuilder) Reset()(*id814594e10d99c016c54d9229da3f00c037b42a01b52dd4efad6b8da5d373fb6.ResetRequestBuilder) {
    return id814594e10d99c016c54d9229da3f00c037b42a01b52dd4efad6b8da5d373fb6.NewResetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSynchronizationProfileRequestBuilder) Resume()(*ibd324584f412d6082609e09d36338222a1ea33342c17f036ead958c71c7e57d6.ResumeRequestBuilder) {
    return ibd324584f412d6082609e09d36338222a1ea33342c17f036ead958c71c7e57d6.NewResumeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSynchronizationProfileRequestBuilder) Start()(*if70ade9d653b89ed0849b345d9e94a2327f9abc57ba5b75b6511e5e4013d1aea.StartRequestBuilder) {
    return if70ade9d653b89ed0849b345d9e94a2327f9abc57ba5b75b6511e5e4013d1aea.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UploadUrl builds and executes requests for operations under \education\synchronizationProfiles\{educationSynchronizationProfile-id}\microsoft.graph.uploadUrl()
func (m *EducationSynchronizationProfileRequestBuilder) UploadUrl()(*icea2a0d7fcf828c9af01c7bba86f54c7305fbb4d6c4cea549b5ee5ece1723710.UploadUrlRequestBuilder) {
    return icea2a0d7fcf828c9af01c7bba86f54c7305fbb4d6c4cea549b5ee5ece1723710.NewUploadUrlRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
