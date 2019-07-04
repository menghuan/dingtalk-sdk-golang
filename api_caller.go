package dingding_sdk_golang

import (
	"github.com/polaris-team/dingding-sdk-golang/encrypt"
	"github.com/polaris-team/dingding-sdk-golang/http"
	"github.com/polaris-team/dingding-sdk-golang/json"
	"strconv"
	"time"
)

//Get a common interface for enterprise authentication information
func CorpAuth(url string, accessKey string, suiteTicket string, authCorpId string, agentId int) (string, error) {
	timestamp := time.Now().UnixNano() / 1e6
	nativeSignature := strconv.FormatInt(timestamp, 10) + "\n" + suiteTicket

	afterHmacSHA256 := encrypt.SHA256(nativeSignature)
	afterBase64 := encrypt.BASE64(afterHmacSHA256)
	afterUrlEncode := encrypt.URLEncode(afterBase64)

	params := map[string]string{
		"timestamp":   strconv.FormatInt(timestamp, 10),
		"accessKey":   accessKey,
		"suiteTicket": suiteTicket,
		"signature":   afterUrlEncode,
	}

	bodyParams := map[string]string{
		"auth_corpid": authCorpId,
	}
	if agentId != 0 {
		bodyParams["agentid"] = strconv.Itoa(agentId)
		bodyParams["suite_key"] = accessKey
	}

	body, _ := json.ToJson(bodyParams)
	return http.Post(url, params, body)
}

//Desc: 获取企业凭证
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/hv357q
func GetCorpToken(accessKey string, suiteTicket string, authCorpId string) (GetCorpTokenResp, error) {
	body, err := CorpAuth("https://oapi.dingtalk.com/service/get_corp_token", accessKey, suiteTicket, authCorpId, 0)
	resp := GetCorpTokenResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 获取企业授权信息
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/fmdqvx
func GetAuthInfo(accessKey string, suiteTicket string, authCorpId string) (GetAuthInfoResp, error) {
	body, err := CorpAuth("https://oapi.dingtalk.com/service/get_auth_info", accessKey, suiteTicket, authCorpId, 0)
	resp := GetAuthInfoResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}

//Desc: 获取授权应用信息
//Doc: https://open-doc.dingtalk.com/microapp/serverapi3/vfitg0
func GetAgent(accessKey string, suiteTicket string, authCorpId string, agentId int) (GetAgentResp, error) {
	body, err := CorpAuth("https://oapi.dingtalk.com/service/get_agent", accessKey, suiteTicket, authCorpId, agentId)
	resp := GetAgentResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}