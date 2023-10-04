package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder provides operations to call the findByMethodMode method.
type AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters invoke function findByMethodMode
type AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetQueryParameters
}
// NewAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal instantiates a new FindByMethodModeWithAuthenticationMethodModesRequestBuilder and sets the default values.
func NewAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, authenticationMethodModes *string)(*AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    m := &AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/authenticationStrengthPolicies/findByMethodMode(authenticationMethodModes={authenticationMethodModes}){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    if authenticationMethodModes != nil {
        m.BaseRequestBuilder.PathParameters["authenticationMethodModes"] = *authenticationMethodModes
    }
    return m
}
// NewAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder instantiates a new FindByMethodModeWithAuthenticationMethodModesRequestBuilder and sets the default values.
func NewAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function findByMethodMode
// Deprecated: This method is obsolete. Use GetAsFindByMethodModeWithAuthenticationMethodModesGetResponse instead.
func (m *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesResponseable), nil
}
// GetAsFindByMethodModeWithAuthenticationMethodModesGetResponse invoke function findByMethodMode
// Deprecated: The findByMethodMode function is deprecated. Please use OData filter query instead. as of 2023-02/FindByMethodModeRemove on 2023-02-01 and will be removed 2023-03-31
func (m *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) GetAsFindByMethodModeWithAuthenticationMethodModesGetResponse(ctx context.Context, requestConfiguration *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesGetResponseable), nil
}
// ToGetRequestInformation invoke function findByMethodMode
// Deprecated: The findByMethodMode function is deprecated. Please use OData filter query instead. as of 2023-02/FindByMethodModeRemove on 2023-02-01 and will be removed 2023-03-31
func (m *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The findByMethodMode function is deprecated. Please use OData filter query instead. as of 2023-02/FindByMethodModeRemove on 2023-02-01 and will be removed 2023-03-31
func (m *AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) WithUrl(rawUrl string)(*AuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder) {
    return NewAuthenticationStrengthPoliciesFindByMethodModeWithAuthenticationMethodModesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
