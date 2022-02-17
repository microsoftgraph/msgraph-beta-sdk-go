package education

import (
    i3c67db4ec1fd2f6c500265c4eb67ba5ae029834f96efe48e7be4b1c86c5e4deb "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me"
    i53e096f34b515b4400cb74425f756b37c554f0dd01122ccbdc851e302a54df1c "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools"
    i58108e2d82e0462b534008f2f016e3d633c124bc780e0445f5b9d139fd18cf22 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles"
    i8c83ac2fb480533d8c9e30080c936573aac55d73adc5a8f1617aed89f50165d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users"
    ib66e30f6b3aeaf350ed09763c0deac43e2a8c5ebd97258eecafd38240fa58d1b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i55d158864f05addf1ddf88c8f6aff3854a29b88e36bc404eb7167a1c13495a9a "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools/item"
    i7bc4bfb8dac45e0b487266018aebf5c4f0660e8547d69a5d9bccd64de3f76c6e "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item"
    ia9a70568b8d421a38915cdf33eea35c2df3432111abdad0bb7d5169fb8c9e26b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item"
    ic55376c4e180e6c30db51b0da6b21653e2bd1c223954aa398f7f85ee7f1ab1fd "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item"
)

// EducationRequestBuilder builds and executes requests for operations under \education
type EducationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationRequestBuilderGetOptions options for Get
type EducationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationRequestBuilderGetQueryParameters get education
type EducationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationRequestBuilderPatchOptions options for Patch
type EducationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationRoot;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EducationRequestBuilder) Classes()(*ib66e30f6b3aeaf350ed09763c0deac43e2a8c5ebd97258eecafd38240fa58d1b.ClassesRequestBuilder) {
    return ib66e30f6b3aeaf350ed09763c0deac43e2a8c5ebd97258eecafd38240fa58d1b.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.classes.item collection
func (m *EducationRequestBuilder) ClassesById(id string)(*ic55376c4e180e6c30db51b0da6b21653e2bd1c223954aa398f7f85ee7f1ab1fd.EducationClassRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass_id"] = id
    }
    return ic55376c4e180e6c30db51b0da6b21653e2bd1c223954aa398f7f85ee7f1ab1fd.NewEducationClassRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationRequestBuilderInternal instantiates a new EducationRequestBuilder and sets the default values.
func NewEducationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationRequestBuilder) {
    m := &EducationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationRequestBuilder instantiates a new EducationRequestBuilder and sets the default values.
func NewEducationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get education
func (m *EducationRequestBuilder) CreateGetRequestInformation(options *EducationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update education
func (m *EducationRequestBuilder) CreatePatchRequestInformation(options *EducationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get education
func (m *EducationRequestBuilder) Get(options *EducationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEducationRoot() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationRoot), nil
}
func (m *EducationRequestBuilder) Me()(*i3c67db4ec1fd2f6c500265c4eb67ba5ae029834f96efe48e7be4b1c86c5e4deb.MeRequestBuilder) {
    return i3c67db4ec1fd2f6c500265c4eb67ba5ae029834f96efe48e7be4b1c86c5e4deb.NewMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update education
func (m *EducationRequestBuilder) Patch(options *EducationRequestBuilderPatchOptions)(error) {
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
func (m *EducationRequestBuilder) Schools()(*i53e096f34b515b4400cb74425f756b37c554f0dd01122ccbdc851e302a54df1c.SchoolsRequestBuilder) {
    return i53e096f34b515b4400cb74425f756b37c554f0dd01122ccbdc851e302a54df1c.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchoolsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.schools.item collection
func (m *EducationRequestBuilder) SchoolsById(id string)(*i55d158864f05addf1ddf88c8f6aff3854a29b88e36bc404eb7167a1c13495a9a.EducationSchoolRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSchool_id"] = id
    }
    return i55d158864f05addf1ddf88c8f6aff3854a29b88e36bc404eb7167a1c13495a9a.NewEducationSchoolRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationRequestBuilder) SynchronizationProfiles()(*i58108e2d82e0462b534008f2f016e3d633c124bc780e0445f5b9d139fd18cf22.SynchronizationProfilesRequestBuilder) {
    return i58108e2d82e0462b534008f2f016e3d633c124bc780e0445f5b9d139fd18cf22.NewSynchronizationProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SynchronizationProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.synchronizationProfiles.item collection
func (m *EducationRequestBuilder) SynchronizationProfilesById(id string)(*ia9a70568b8d421a38915cdf33eea35c2df3432111abdad0bb7d5169fb8c9e26b.EducationSynchronizationProfileRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSynchronizationProfile_id"] = id
    }
    return ia9a70568b8d421a38915cdf33eea35c2df3432111abdad0bb7d5169fb8c9e26b.NewEducationSynchronizationProfileRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationRequestBuilder) Users()(*i8c83ac2fb480533d8c9e30080c936573aac55d73adc5a8f1617aed89f50165d0.UsersRequestBuilder) {
    return i8c83ac2fb480533d8c9e30080c936573aac55d73adc5a8f1617aed89f50165d0.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item collection
func (m *EducationRequestBuilder) UsersById(id string)(*i7bc4bfb8dac45e0b487266018aebf5c4f0660e8547d69a5d9bccd64de3f76c6e.EducationUserRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationUser_id"] = id
    }
    return i7bc4bfb8dac45e0b487266018aebf5c4f0660e8547d69a5d9bccd64de3f76c6e.NewEducationUserRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
