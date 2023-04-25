package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// InformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilder provides operations to call the evaluateApplication method.
type InformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilderInternal instantiates a new SecurityEvaluateApplicationRequestBuilder and sets the default values.
func NewInformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilder) {
    m := &InformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/informationProtection/sensitivityLabels/security.evaluateApplication", pathParameters),
    }
    return m
}
// NewInformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilder instantiates a new SecurityEvaluateApplicationRequestBuilder and sets the default values.
func NewInformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action evaluateApplication
func (m *InformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilder) Post(ctx context.Context, body InformationProtectionSensitivityLabelsSecurityEvaluateApplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *InformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration)(InformationProtectionSensitivityLabelsSecurityEvaluateApplicationEvaluateApplicationResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInformationProtectionSensitivityLabelsSecurityEvaluateApplicationEvaluateApplicationResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InformationProtectionSensitivityLabelsSecurityEvaluateApplicationEvaluateApplicationResponseable), nil
}
// ToPostRequestInformation invoke action evaluateApplication
func (m *InformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilder) ToPostRequestInformation(ctx context.Context, body InformationProtectionSensitivityLabelsSecurityEvaluateApplicationEvaluateApplicationPostRequestBodyable, requestConfiguration *InformationProtectionSensitivityLabelsSecurityEvaluateApplicationRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
