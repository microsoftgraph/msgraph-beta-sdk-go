package getrelatedappstateswithuserprincipalnamewithdeviceid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder provides operations to call the getRelatedAppStates method.
type GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetOptions options for Get
type GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal instantiates a new GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder and sets the default values.
func NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deviceId *string, userPrincipalName *string)(*GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    m := &GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/microsoft.graph.getRelatedAppStates(userPrincipalName='{userPrincipalName}',deviceId='{deviceId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if deviceId != nil {
        urlTplParams[""] = *deviceId
    }
    if userPrincipalName != nil {
        urlTplParams[""] = *userPrincipalName
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder instantiates a new GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder and sets the default values.
func NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// CreateGetRequestInformation invoke function getRelatedAppStates
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) CreateGetRequestInformation(options *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
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
// Get invoke function getRelatedAppStates
func (m *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) Get(options *GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderGetOptions)(GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdResponseable), nil
}
