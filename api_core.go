/*
 * Affix API
 *
 * The affixapi.com API documentation.  # Introduction Affix API is an OAuth 2.1 application that allows developers to access customer data, without developers needing to manage or maintain integrations; or collect login credentials or API keys from users for these third party systems.  # OAuth 2.1 Affix API follows the [OAuth 2.1 spec](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-v2-1-08).  As an OAuth application, Affix API handles not only both the collection of sensitive user credentials or API keys, but also builds and maintains the integrations with the providers, so you don't have to.  # How to obtain an access token in order to get started, you must:   - register a `client_id`   - direct your user to the sign in flow  (`https://connect.affixapi.com`     [with the appropriate query     parameters](https://github.com/affixapi/starter-kit/tree/master/connect))   - capture `authorization_code` we will send to your redirect URI after     the sign in flow is complete and exchange that `authorization_code` for     a Bearer token  # Sandbox keys (payroll) ### dev ``` eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvcGF5cm9sbC9lbXBsb3llZXMiLCIvMjAyMy0wMy0wMS9wYXlyb2xsL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zIiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zLzpwYXlydW5faWQiXSwidG9rZW4iOiI1ZGZmMWRlMy1kZTk2LTQzZWYtYTgwMi0wYWZlNWQ5NzU4MzkiLCJpYXQiOjE2OTM4Mzk1NjIsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUuZGV2LmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6InBheXJvbGwiLCJhdWQiOiIzRkRBRURGOS0xRENBNEY1NC04Nzk0OUY2QS00MTAyNzY0MyJ9.D8i_TwguGZ9vok7RCUuk5FoUJGWEGnI-4VzReI1DAg0ytZaQ8XFmVxC4RxGYBxah1ZJpcrbAsuhssGEDXffIgQ ```  ``` curl --fail \\   \"https://dev.api.affixapi.com/2023-03-01/payroll/employees\" \\   -H \"Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvcGF5cm9sbC9lbXBsb3llZXMiLCIvMjAyMy0wMy0wMS9wYXlyb2xsL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zIiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zLzpwYXlydW5faWQiXSwidG9rZW4iOiI1ZGZmMWRlMy1kZTk2LTQzZWYtYTgwMi0wYWZlNWQ5NzU4MzkiLCJpYXQiOjE2OTM4Mzk1NjIsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUuZGV2LmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6InBheXJvbGwiLCJhdWQiOiIzRkRBRURGOS0xRENBNEY1NC04Nzk0OUY2QS00MTAyNzY0MyJ9.D8i_TwguGZ9vok7RCUuk5FoUJGWEGnI-4VzReI1DAg0ytZaQ8XFmVxC4RxGYBxah1ZJpcrbAsuhssGEDXffIgQ\" \\   -X GET ```  ### prod ``` eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvcGF5cm9sbC9lbXBsb3llZXMiLCIvMjAyMy0wMy0wMS9wYXlyb2xsL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zIiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zLzpwYXlydW5faWQiXSwidG9rZW4iOiJhODdhNGE0My1jMWM2LTQ5ZTktOTUxZS0zNzAzNzRjODhiNjciLCJpYXQiOjE2OTM4Mzk2NjMsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUucHJvZC5lbmdpbmVlcmluZy5hZmZpeGFwaS5jb20iLCJzdWIiOiJwYXlyb2xsIiwiYXVkIjoiMDhCQjA4MUUtRDlBQjREMTQtOERGOTkyMzMtNjY2MTVDRTkifQ.AJPINjuEzd9JH1fpf6l1L6L4XdHNIsVcaDsicFe8Giz69zqK1z6za3usbUz71kKJNmm6Eo8HfLnowCi8bcAnow ```  ``` curl --fail \\   \"https://api.affixapi.com/2023-03-01/payroll/employees\" \\   -H \"Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvcGF5cm9sbC9lbXBsb3llZXMiLCIvMjAyMy0wMy0wMS9wYXlyb2xsL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zIiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zLzpwYXlydW5faWQiXSwidG9rZW4iOiJhODdhNGE0My1jMWM2LTQ5ZTktOTUxZS0zNzAzNzRjODhiNjciLCJpYXQiOjE2OTM4Mzk2NjMsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUucHJvZC5lbmdpbmVlcmluZy5hZmZpeGFwaS5jb20iLCJzdWIiOiJwYXlyb2xsIiwiYXVkIjoiMDhCQjA4MUUtRDlBQjREMTQtOERGOTkyMzMtNjY2MTVDRTkifQ.AJPINjuEzd9JH1fpf6l1L6L4XdHNIsVcaDsicFe8Giz69zqK1z6za3usbUz71kKJNmm6Eo8HfLnowCi8bcAnow\" \\   -X GET ```  # Webhooks An exciting feature for HRIS/Payroll modes are webhooks.  If enabled, your `webhook_uri` is set on your `client_id` for the respective environment: `dev | prod`  Webhooks are configured to make live requests to the underlying integration 1x/hr, and if a difference is detected since the last request, we will send a request to your `webhook_uri` with this shape:  ``` {    added: <api.Employees20230301Response>[     <api.Employee20230301>{       ...,       date_of_birth: '2010-08-06',       display_full_name: 'Daija Rogahn',       employee_number: '57993',       employment_status: 'pending',       employment_type: 'other',       employments: [         {           currency: 'eur',           effective_date: '2022-02-25',           employment_type: 'other',           job_title: 'Dynamic Implementation Manager',           pay_frequency: 'semimonthly',           pay_period: 'YEAR',           pay_rate: 96000,         },       ],       first_name: 'Daija',       ...     }   ],   removed: [],   updated: [     <api.Employee20230301>{       ...,       date_of_birth: '2009-11-09',       display_full_name: 'Lourdes Stiedemann',       employee_number: '63189',       employment_status: 'leave',       employment_type: 'full_time',       employments: [         {           currency: 'gbp',           effective_date: '2023-01-16',           employment_type: 'full_time',           job_title: 'Forward Brand Planner',           pay_frequency: 'semimonthly',           pay_period: 'YEAR',           pay_rate: 86000,         },       ],       first_name: 'Lourdes',     }   ] } ```  the following headers will be sent with webhook requests:  ``` x-affix-api-signature: ab8474e609db95d5df3adc39ea3add7a7544bd215c5c520a30a650ae93a2fba7  x-affix-api-origin:  webhooks-employees-webhook  user-agent:  affixapi.com ```  Before trusting the payload, you should sign the payload and verify the signature matches the signature sent by the `affixapi.com` service.  This secures that the data sent to your `webhook_uri` is from the `affixapi.com` server.  The signature is created by combining the signing secret (your `client_secret`) with the body of the request sent using a standard HMAC-SHA256 keyed hash.  The signature can be created via:   - create an `HMAC` with your `client_secret`   - update the `HMAC` with the payload   - get the hex digest -> this is the signature  Sample `typescript` code that follows this recipe:  ``` import { createHmac } from 'crypto';  export const computeSignature = ({   str,   signingSecret, }: {   signingSecret: string;   str: string; }): string => {   const hmac = createHmac('sha256', signingSecret);   hmac.update(str);   const signature = hmac.digest('hex');    return signature; }; ```  ## Rate limits Open endpoints (not gated by an API key) (applied at endpoint level):   - 15 requests every 1 minute (by IP address)   - 25 requests every 5 minutes (by IP address)  Gated endpoints (require an API key) (applied at endpoint level):   - 40 requests every 1 minute (by IP address)   - 40 requests every 5 minutes (by `client_id`)  Things to keep in mind:   - Open endpoints (not gated by an API key) will likely be called by your     users, not you, so rate limits generally would not apply to you.   - As a developer, rate limits are applied at the endpoint granularity.     - For example, say the rate limits below are 10 requests per minute by ip.       from that same ip, within 1 minute, you get:       - 10 requests per minute on `/orders`,       - another 10 requests per minute on `/items`,       - and another 10 requests per minute on `/identity`,       - for a total of 30 requests per minute. 
 *
 * API version: 2023-03-01
 * Contact: john@affixapi.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// CoreApiService CoreApi service
type CoreApiService service

type ApiProvidersRequest struct {
	ctx _context.Context
	ApiService *CoreApiService
}


func (r ApiProvidersRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.ProvidersExecute(r)
}

/*
 * Providers List of providers
 * Retrieve the api modes (hris, payroll) and providers for the respective
modes

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiProvidersRequest
 */
func (a *CoreApiService) Providers(ctx _context.Context) ApiProvidersRequest {
	return ApiProvidersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []map[string]interface{}
 */
func (a *CoreApiService) ProvidersExecute(r ApiProvidersRequest) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreApiService.Providers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v MessageResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
