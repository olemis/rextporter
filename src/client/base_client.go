package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/oliveagle/jsonpath"
	"github.com/simelo/rextporter/src/config"
	"github.com/simelo/rextporter/src/util"
)

// BaseClient have common data to be shared through embedded struct in those type who implement the
// client.Client interface
type BaseClient struct {
	req         *http.Request
	service     config.Service
	token       string
	metricJPath string
}

func (client *BaseClient) resetToken() (err error) {
	const generalScopeErr = "error making resetting the token"
	client.token = ""
	var clientToken *TokenClient
	if clientToken, err = newTokenClient(client.service); err != nil {
		errCause := fmt.Sprintln("can not find a host: ", err.Error())
		return util.ErrorFromThisScope(errCause, generalScopeErr)
	}
	var data []byte
	if data, err = clientToken.getRemoteInfo(); err != nil {
		errCause := fmt.Sprintln("can make the request to get a token: ", err.Error())
		return util.ErrorFromThisScope(errCause, generalScopeErr)
	}
	var jsonData interface{}
	if err = json.Unmarshal(data, &jsonData); err != nil {
		errCause := fmt.Sprintln("can not decode the body: ", string(data), " ", err.Error())
		return util.ErrorFromThisScope(errCause, generalScopeErr)
	}
	var val interface{}
	jPath := "$" + strings.Replace(client.service.TokenKeyFromEndpoint, "/", ".", -1)
	if val, err = jsonpath.JsonPathLookup(jsonData, jPath); err != nil {
		errCause := fmt.Sprintln("can not locate the path: ", err.Error())
		return util.ErrorFromThisScope(errCause, generalScopeErr)
	}
	tk, ok := val.(string)
	if !ok {
		errCause := fmt.Sprintln("unable the get the token as a string value")
		return util.ErrorFromThisScope(errCause, generalScopeErr)
	}
	client.token = tk
	if len(client.token) == 0 {
		errCause := fmt.Sprintln("unable the get a not null(empty) token")
		return util.ErrorFromThisScope(errCause, generalScopeErr)
	}
	return nil
}

func (client *BaseClient) getRemoteInfo() (data []byte, err error) {
	const generalScopeErr = "error making a server request to get metric from remote endpoint"
	client.req.Header.Set(client.service.TokenHeaderKey, client.token)
	getData := func() (data []byte, err error) {
		httpClient := &http.Client{}
		var resp *http.Response
		if resp, err = httpClient.Do(client.req); err != nil {
			errCause := fmt.Sprintln("can not do the request: ", err.Error())
			return nil, util.ErrorFromThisScope(errCause, generalScopeErr)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			errCause := fmt.Sprintf("no success response, status %s", resp.Status)
			return nil, util.ErrorFromThisScope(errCause, generalScopeErr)
		}
		if data, err = ioutil.ReadAll(resp.Body); err != nil {
			errCause := fmt.Sprintln("can not read the body: ", err.Error())
			return nil, util.ErrorFromThisScope(errCause, generalScopeErr)
		}
		return data, nil
	}
	if data, err = getData(); err != nil {
		// log.Println("can not do the request:", err.Error(), "trying with a new token...")
		if err = client.resetToken(); err != nil {
			errCause := fmt.Sprintln("can not reset the token: ", err.Error())
			return nil, util.ErrorFromThisScope(errCause, generalScopeErr)
		}
		if data, err = getData(); err != nil {
			errCause := fmt.Sprintln("can not do the request after a token reset neither: ", err.Error())
			return nil, util.ErrorFromThisScope(errCause, generalScopeErr)
		}
	}
	return data, nil
}