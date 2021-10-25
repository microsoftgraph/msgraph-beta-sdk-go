package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i111455d8fd8d7f0ddf817ca0fcc2bc0624e3819c5bdee6ec2194f7ae9db42ddf "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/targetapps"
    i4633bb1fbea3b9ff62e44e3637dd657b2fff6cbd90287bd10a013fed2d40860f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assignments"
    i7d0f522bac2e00da7e218c106fc841a9805ca78f7c4bdd7512999c9cfc802d42 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/apps"
    ida25c58687771cbc2e8cd5750db08465dd94e755be0fd37d473f36b0a0502e49 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/deploymentsummary"
    iefd9a3454d070af5a6b3c6f36a2215b5d86ef0d740bac940405f70157b57b8a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assign"
    i6a789da51c915c8f551873e29bac081934811f780b24231053b29a8984f30fb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/apps/item"
    ie924ef4384ce239f82e862609ff7154fd7dc52ac669aa7e01fb54a9dbd903a66 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item/assignments/item"
)

type TargetedManagedAppConfigurationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type TargetedManagedAppConfigurationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *TargetedManagedAppConfigurationRequestBuilder) Apps()(*i7d0f522bac2e00da7e218c106fc841a9805ca78f7c4bdd7512999c9cfc802d42.AppsRequestBuilder) {
    return i7d0f522bac2e00da7e218c106fc841a9805ca78f7c4bdd7512999c9cfc802d42.NewAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TargetedManagedAppConfigurationRequestBuilder) AppsById(id string)(*i6a789da51c915c8f551873e29bac081934811f780b24231053b29a8984f30fb5.ManagedMobileAppRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["managedMobileApp_id"] = id
    }
    return i6a789da51c915c8f551873e29bac081934811f780b24231053b29a8984f30fb5.NewManagedMobileAppRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TargetedManagedAppConfigurationRequestBuilder) Assign()(*iefd9a3454d070af5a6b3c6f36a2215b5d86ef0d740bac940405f70157b57b8a5.AssignRequestBuilder) {
    return iefd9a3454d070af5a6b3c6f36a2215b5d86ef0d740bac940405f70157b57b8a5.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TargetedManagedAppConfigurationRequestBuilder) Assignments()(*i4633bb1fbea3b9ff62e44e3637dd657b2fff6cbd90287bd10a013fed2d40860f.AssignmentsRequestBuilder) {
    return i4633bb1fbea3b9ff62e44e3637dd657b2fff6cbd90287bd10a013fed2d40860f.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TargetedManagedAppConfigurationRequestBuilder) AssignmentsById(id string)(*ie924ef4384ce239f82e862609ff7154fd7dc52ac669aa7e01fb54a9dbd903a66.TargetedManagedAppPolicyAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["targetedManagedAppPolicyAssignment_id"] = id
    }
    return ie924ef4384ce239f82e862609ff7154fd7dc52ac669aa7e01fb54a9dbd903a66.NewTargetedManagedAppPolicyAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewTargetedManagedAppConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppConfigurationRequestBuilder) {
    m := &TargetedManagedAppConfigurationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewTargetedManagedAppConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TargetedManagedAppConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedManagedAppConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *TargetedManagedAppConfigurationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *TargetedManagedAppConfigurationRequestBuilder) CreateGetRequestInformation(q func (value *TargetedManagedAppConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(TargetedManagedAppConfigurationRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *TargetedManagedAppConfigurationRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TargetedManagedAppConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *TargetedManagedAppConfigurationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *TargetedManagedAppConfigurationRequestBuilder) DeploymentSummary()(*ida25c58687771cbc2e8cd5750db08465dd94e755be0fd37d473f36b0a0502e49.DeploymentSummaryRequestBuilder) {
    return ida25c58687771cbc2e8cd5750db08465dd94e755be0fd37d473f36b0a0502e49.NewDeploymentSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TargetedManagedAppConfigurationRequestBuilder) Get(q func (value *TargetedManagedAppConfigurationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TargetedManagedAppConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTargetedManagedAppConfiguration() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TargetedManagedAppConfiguration), nil
}
func (m *TargetedManagedAppConfigurationRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TargetedManagedAppConfiguration, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *TargetedManagedAppConfigurationRequestBuilder) TargetApps()(*i111455d8fd8d7f0ddf817ca0fcc2bc0624e3819c5bdee6ec2194f7ae9db42ddf.TargetAppsRequestBuilder) {
    return i111455d8fd8d7f0ddf817ca0fcc2bc0624e3819c5bdee6ec2194f7ae9db42ddf.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
