package doesuserhaveaccesswithuseridwithtenantidwithuserprincipalname

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder provides operations to call the doesUserHaveAccess method.
type DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetOptions options for Get
type DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal instantiates a new DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder and sets the default values.
func NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, tenantId *string, userId *string, userPrincipalName *string)(*DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    m := &DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/team/primaryChannel/microsoft.graph.doesUserHaveAccess(userId='{userId}',tenantId='{tenantId}',userPrincipalName='{userPrincipalName}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if tenantId != nil {
        urlTplParams["tenantId"] = *tenantId
    }
    if userId != nil {
        urlTplParams["userId"] = *userId
    }
    if userPrincipalName != nil {
        urlTplParams["userPrincipalName"] = *userPrincipalName
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder instantiates a new DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder and sets the default values.
func NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// CreateGetRequestInformation invoke function doesUserHaveAccess
func (m *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) CreateGetRequestInformation(options *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function doesUserHaveAccess
func (m *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilder) Get(options *DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameRequestBuilderGetOptions)(DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateDoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(DoesUserHaveAccessWithUserIdWithTenantIdWithUserPrincipalNameResponseable), nil
}
