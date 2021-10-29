package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i029c708ccdb6d4ef5dc4da9e27b34aa93e864246d8928a227c86dea55f0ed64c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/devicestatusoverview"
    i3f8fc70c29f6750d091ef750e61ce5245979cb89e4ded10d81e6f29376601e4f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/scheduleactionsforrules"
    i439f777579087aca184f6db254b06b730f0350f86f3b16ab9715855a4d1dba9d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/devicesettingstatesummaries"
    i67fbfff12a38b588d862cac3e0f5919efcd47ec1706ebbebae7f2a9c4cb13f8b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/scheduledactionsforrule"
    i9488f52bd7cc76d580224eb7ca66454394d137f7c9e6dd2acb88d09d9a7e51c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/assignments"
    ib270494a9068f6e72dec9f03b5ba60e5bb8fe5ad2b5cb9fbe2e630750404a4b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/assign"
    ibc404ddee37b91bd7cb4a0ad73dd266370588956b1f15cfc815c8ae7b9470c9d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/userstatuses"
    ibe0820e05500713b9524c00f4582d491a213af40b239abac338d37833fdf7c21 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/userstatusoverview"
    id1ddb4feb2a86c7e00f57540a6ad6d5e4ec28b49c8a7d13acfad4a9b79ffef7a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/devicestatuses"
    i2241d054608ba9643ba9fabeb07e233374bce0842cb061ee1fa65ee48d7cf7fb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/devicestatuses/item"
    i3383893ad7fa9d341ce2de93a744c841a52543d2213e0bfa5f2aa3542cc73fe0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/assignments/item"
    i411a74566de656324f76d6caebf7327fa83cbd750a7c74e112ce1d060b24d414 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/userstatuses/item"
    i5151b97d6c1c450de6501bbc7815d527e3f1cd53abb08be08da23fc61bd291db "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/scheduledactionsforrule/item"
    iae08883b92d515e98cdc6bc2aa90846da99d3b61ad866c8ab0134b07d2effedf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancepolicies/item/devicesettingstatesummaries/item"
)

type DeviceCompliancePolicyRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DeviceCompliancePolicyRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *DeviceCompliancePolicyRequestBuilder) Assign()(*ib270494a9068f6e72dec9f03b5ba60e5bb8fe5ad2b5cb9fbe2e630750404a4b1.AssignRequestBuilder) {
    return ib270494a9068f6e72dec9f03b5ba60e5bb8fe5ad2b5cb9fbe2e630750404a4b1.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) Assignments()(*i9488f52bd7cc76d580224eb7ca66454394d137f7c9e6dd2acb88d09d9a7e51c6.AssignmentsRequestBuilder) {
    return i9488f52bd7cc76d580224eb7ca66454394d137f7c9e6dd2acb88d09d9a7e51c6.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) AssignmentsById(id string)(*i3383893ad7fa9d341ce2de93a744c841a52543d2213e0bfa5f2aa3542cc73fe0.DeviceCompliancePolicyAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyAssignment_id"] = id
    }
    return i3383893ad7fa9d341ce2de93a744c841a52543d2213e0bfa5f2aa3542cc73fe0.NewDeviceCompliancePolicyAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewDeviceCompliancePolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCompliancePolicyRequestBuilder) {
    m := &DeviceCompliancePolicyRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewDeviceCompliancePolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCompliancePolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePolicyRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DeviceCompliancePolicyRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceCompliancePolicyRequestBuilder) CreateGetRequestInformation(q func (value *DeviceCompliancePolicyRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DeviceCompliancePolicyRequestBuilderGetQueryParameters)
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
func (m *DeviceCompliancePolicyRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCompliancePolicy, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceCompliancePolicyRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *DeviceCompliancePolicyRequestBuilder) DeviceSettingStateSummaries()(*i439f777579087aca184f6db254b06b730f0350f86f3b16ab9715855a4d1dba9d.DeviceSettingStateSummariesRequestBuilder) {
    return i439f777579087aca184f6db254b06b730f0350f86f3b16ab9715855a4d1dba9d.NewDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) DeviceSettingStateSummariesById(id string)(*iae08883b92d515e98cdc6bc2aa90846da99d3b61ad866c8ab0134b07d2effedf.SettingStateDeviceSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["settingStateDeviceSummary_id"] = id
    }
    return iae08883b92d515e98cdc6bc2aa90846da99d3b61ad866c8ab0134b07d2effedf.NewSettingStateDeviceSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) DeviceStatuses()(*id1ddb4feb2a86c7e00f57540a6ad6d5e4ec28b49c8a7d13acfad4a9b79ffef7a.DeviceStatusesRequestBuilder) {
    return id1ddb4feb2a86c7e00f57540a6ad6d5e4ec28b49c8a7d13acfad4a9b79ffef7a.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) DeviceStatusesById(id string)(*i2241d054608ba9643ba9fabeb07e233374bce0842cb061ee1fa65ee48d7cf7fb.DeviceComplianceDeviceStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceDeviceStatus_id"] = id
    }
    return i2241d054608ba9643ba9fabeb07e233374bce0842cb061ee1fa65ee48d7cf7fb.NewDeviceComplianceDeviceStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) DeviceStatusOverview()(*i029c708ccdb6d4ef5dc4da9e27b34aa93e864246d8928a227c86dea55f0ed64c.DeviceStatusOverviewRequestBuilder) {
    return i029c708ccdb6d4ef5dc4da9e27b34aa93e864246d8928a227c86dea55f0ed64c.NewDeviceStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) Get(q func (value *DeviceCompliancePolicyRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCompliancePolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceCompliancePolicy() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCompliancePolicy), nil
}
func (m *DeviceCompliancePolicyRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCompliancePolicy, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *DeviceCompliancePolicyRequestBuilder) ScheduleActionsForRules()(*i3f8fc70c29f6750d091ef750e61ce5245979cb89e4ded10d81e6f29376601e4f.ScheduleActionsForRulesRequestBuilder) {
    return i3f8fc70c29f6750d091ef750e61ce5245979cb89e4ded10d81e6f29376601e4f.NewScheduleActionsForRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) ScheduledActionsForRule()(*i67fbfff12a38b588d862cac3e0f5919efcd47ec1706ebbebae7f2a9c4cb13f8b.ScheduledActionsForRuleRequestBuilder) {
    return i67fbfff12a38b588d862cac3e0f5919efcd47ec1706ebbebae7f2a9c4cb13f8b.NewScheduledActionsForRuleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) ScheduledActionsForRuleById(id string)(*i5151b97d6c1c450de6501bbc7815d527e3f1cd53abb08be08da23fc61bd291db.DeviceComplianceScheduledActionForRuleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceScheduledActionForRule_id"] = id
    }
    return i5151b97d6c1c450de6501bbc7815d527e3f1cd53abb08be08da23fc61bd291db.NewDeviceComplianceScheduledActionForRuleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) UserStatuses()(*ibc404ddee37b91bd7cb4a0ad73dd266370588956b1f15cfc815c8ae7b9470c9d.UserStatusesRequestBuilder) {
    return ibc404ddee37b91bd7cb4a0ad73dd266370588956b1f15cfc815c8ae7b9470c9d.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) UserStatusesById(id string)(*i411a74566de656324f76d6caebf7327fa83cbd750a7c74e112ce1d060b24d414.DeviceComplianceUserStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceUserStatus_id"] = id
    }
    return i411a74566de656324f76d6caebf7327fa83cbd750a7c74e112ce1d060b24d414.NewDeviceComplianceUserStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceCompliancePolicyRequestBuilder) UserStatusOverview()(*ibe0820e05500713b9524c00f4582d491a213af40b239abac338d37833fdf7c21.UserStatusOverviewRequestBuilder) {
    return ibe0820e05500713b9524c00f4582d491a213af40b239abac338d37833fdf7c21.NewUserStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}