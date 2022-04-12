package education

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3c67db4ec1fd2f6c500265c4eb67ba5ae029834f96efe48e7be4b1c86c5e4deb "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me"
    i53e096f34b515b4400cb74425f756b37c554f0dd01122ccbdc851e302a54df1c "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools"
    i58108e2d82e0462b534008f2f016e3d633c124bc780e0445f5b9d139fd18cf22 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles"
    i8c83ac2fb480533d8c9e30080c936573aac55d73adc5a8f1617aed89f50165d0 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users"
    ib66e30f6b3aeaf350ed09763c0deac43e2a8c5ebd97258eecafd38240fa58d1b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes"
    i55d158864f05addf1ddf88c8f6aff3854a29b88e36bc404eb7167a1c13495a9a "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools/item"
    i7bc4bfb8dac45e0b487266018aebf5c4f0660e8547d69a5d9bccd64de3f76c6e "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item"
    ia9a70568b8d421a38915cdf33eea35c2df3432111abdad0bb7d5169fb8c9e26b "github.com/microsoftgraph/msgraph-beta-sdk-go/education/synchronizationprofiles/item"
    ic55376c4e180e6c30db51b0da6b21653e2bd1c223954aa398f7f85ee7f1ab1fd "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item"
)

// EducationRequestBuilder provides operations to manage the educationRoot singleton.
type EducationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EducationRequestBuilderGetOptions options for Get
type EducationRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EducationRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EducationRequestBuilderGetQueryParameters get education
type EducationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EducationRequestBuilderPatchOptions options for Patch
type EducationRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationRootable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Classes the classes property
func (m *EducationRequestBuilder) Classes()(*ib66e30f6b3aeaf350ed09763c0deac43e2a8c5ebd97258eecafd38240fa58d1b.ClassesRequestBuilder) {
    return ib66e30f6b3aeaf350ed09763c0deac43e2a8c5ebd97258eecafd38240fa58d1b.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.classes.item collection
func (m *EducationRequestBuilder) ClassesById(id string)(*ic55376c4e180e6c30db51b0da6b21653e2bd1c223954aa398f7f85ee7f1ab1fd.EducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass%2Did"] = id
    }
    return ic55376c4e180e6c30db51b0da6b21653e2bd1c223954aa398f7f85ee7f1ab1fd.NewEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationRequestBuilderInternal instantiates a new EducationRequestBuilder and sets the default values.
func NewEducationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationRequestBuilder) {
    m := &EducationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationRequestBuilder instantiates a new EducationRequestBuilder and sets the default values.
func NewEducationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get education
func (m *EducationRequestBuilder) CreateGetRequestInformation(options *EducationRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update education
func (m *EducationRequestBuilder) CreatePatchRequestInformation(options *EducationRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get get education
func (m *EducationRequestBuilder) Get(options *EducationRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEducationRootFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EducationRootable), nil
}
// Me the me property
func (m *EducationRequestBuilder) Me()(*i3c67db4ec1fd2f6c500265c4eb67ba5ae029834f96efe48e7be4b1c86c5e4deb.MeRequestBuilder) {
    return i3c67db4ec1fd2f6c500265c4eb67ba5ae029834f96efe48e7be4b1c86c5e4deb.NewMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update education
func (m *EducationRequestBuilder) Patch(options *EducationRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Schools the schools property
func (m *EducationRequestBuilder) Schools()(*i53e096f34b515b4400cb74425f756b37c554f0dd01122ccbdc851e302a54df1c.SchoolsRequestBuilder) {
    return i53e096f34b515b4400cb74425f756b37c554f0dd01122ccbdc851e302a54df1c.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchoolsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.schools.item collection
func (m *EducationRequestBuilder) SchoolsById(id string)(*i55d158864f05addf1ddf88c8f6aff3854a29b88e36bc404eb7167a1c13495a9a.EducationSchoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSchool%2Did"] = id
    }
    return i55d158864f05addf1ddf88c8f6aff3854a29b88e36bc404eb7167a1c13495a9a.NewEducationSchoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SynchronizationProfiles the synchronizationProfiles property
func (m *EducationRequestBuilder) SynchronizationProfiles()(*i58108e2d82e0462b534008f2f016e3d633c124bc780e0445f5b9d139fd18cf22.SynchronizationProfilesRequestBuilder) {
    return i58108e2d82e0462b534008f2f016e3d633c124bc780e0445f5b9d139fd18cf22.NewSynchronizationProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SynchronizationProfilesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.synchronizationProfiles.item collection
func (m *EducationRequestBuilder) SynchronizationProfilesById(id string)(*ia9a70568b8d421a38915cdf33eea35c2df3432111abdad0bb7d5169fb8c9e26b.EducationSynchronizationProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSynchronizationProfile%2Did"] = id
    }
    return ia9a70568b8d421a38915cdf33eea35c2df3432111abdad0bb7d5169fb8c9e26b.NewEducationSynchronizationProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Users the users property
func (m *EducationRequestBuilder) Users()(*i8c83ac2fb480533d8c9e30080c936573aac55d73adc5a8f1617aed89f50165d0.UsersRequestBuilder) {
    return i8c83ac2fb480533d8c9e30080c936573aac55d73adc5a8f1617aed89f50165d0.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.education.users.item collection
func (m *EducationRequestBuilder) UsersById(id string)(*i7bc4bfb8dac45e0b487266018aebf5c4f0660e8547d69a5d9bccd64de3f76c6e.EducationUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationUser%2Did"] = id
    }
    return i7bc4bfb8dac45e0b487266018aebf5c4f0660e8547d69a5d9bccd64de3f76c6e.NewEducationUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
