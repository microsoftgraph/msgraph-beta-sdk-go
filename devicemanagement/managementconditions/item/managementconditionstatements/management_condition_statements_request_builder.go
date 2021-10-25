package managementconditionstatements

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ie653e6d497ebe0b3127cb9b98857f9799213e0b1bd7e4906a4d8cd8a5280330c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/managementconditions/item/managementconditionstatements/ref"
    if8ca9966930d7ad505e9fc8a4fd9fc00488cdc68bc25732da4bff69daa542a27 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/managementconditions/item/managementconditionstatements/getmanagementconditionstatementsforplatformwithplatform"
)

type ManagementConditionStatementsRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ManagementConditionStatementsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select_escpaped []string;
    Skip *int32;
    Top *int32;
}
func NewManagementConditionStatementsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementConditionStatementsRequestBuilder) {
    m := &ManagementConditionStatementsRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/managementConditions/{managementCondition_id}/managementConditionStatements{?top,skip,search,filter,count,orderby,select,expand}";
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
func NewManagementConditionStatementsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementConditionStatementsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagementConditionStatementsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagementConditionStatementsRequestBuilder) CreateGetRequestInformation(q func (value *ManagementConditionStatementsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ManagementConditionStatementsRequestBuilderGetQueryParameters)
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
func (m *ManagementConditionStatementsRequestBuilder) Get(q func (value *ManagementConditionStatementsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*ManagementConditionStatementsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementConditionStatementsResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*ManagementConditionStatementsResponse), nil
}
func (m *ManagementConditionStatementsRequestBuilder) GetManagementConditionStatementsForPlatformWithPlatform(platform *string)(*if8ca9966930d7ad505e9fc8a4fd9fc00488cdc68bc25732da4bff69daa542a27.GetManagementConditionStatementsForPlatformWithPlatformRequestBuilder) {
    return if8ca9966930d7ad505e9fc8a4fd9fc00488cdc68bc25732da4bff69daa542a27.NewGetManagementConditionStatementsForPlatformWithPlatformRequestBuilderInternal(m.pathParameters, m.requestAdapter, platform);
}
func (m *ManagementConditionStatementsRequestBuilder) Ref()(*ie653e6d497ebe0b3127cb9b98857f9799213e0b1bd7e4906a4d8cd8a5280330c.RefRequestBuilder) {
    return ie653e6d497ebe0b3127cb9b98857f9799213e0b1bd7e4906a4d8cd8a5280330c.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
