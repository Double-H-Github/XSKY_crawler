package xsky

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	getTokenUrl = "https://xskydata.jobs.feishu.cn/api/v1/csrf/token"
	getJobListUrl = "https://xskydata.jobs.feishu.cn/api/v1/search/job/posts"
)

func GetToken() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", getTokenUrl, nil)
	if err != nil {
		log.Printf("[GetToken] http.NewRequest failed err: %v\n", err)
		return "", err
	}
	resp, err := client.Do(req) //发起请求
	if err != nil {
		log.Printf("[GetToken] Do req failed err: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[GetToken] ioutil.ReadAll(resp.Body) failed err: %v\n", err)
		return "", err
	}

	getTokenResp := &GetTokenResp{} //自定义结构体，反序列化返回的数据
	if err := json.Unmarshal(body, getTokenResp); err != nil {
		log.Printf("[GetToken] Unmarshal failed err: %v\n", err)
		return "", err
	}

	if getTokenResp.Code != 0 && getTokenResp.Message != "ok" {
		log.Printf("[GetToken] Response err: %v, resp: %v \n", err, getTokenResp)
		return "", err
	}
	return getTokenResp.Data.Token, nil
}

func GetJobList(token string) ([]*JobInfo, error) {

	client := &http.Client{}
	reqBody := bytes.NewBuffer([]byte("{\"limit\":10000}"))
	req, err := http.NewRequest("POST", getJobListUrl, reqBody) //设置请求Body
	if err != nil {
		log.Printf("[GetJobList] Do req failed err: %v\n", err)
		return nil, err
	}

	//设置请求Header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("cookie", "atsx-csrf-token=" + token)
	req.Header.Set("website-path", "school")
	req.Header.Set("x-csrf-token", token)

	resp, err := client.Do(req) //发起Post请求
	if err != nil {
		log.Printf("[GetJobList] Do req failed err: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	getJobResp := &GetJobResp{} //自定义结构体，反序列化返回的数据
	if err := json.Unmarshal(body, getJobResp); err != nil {
		log.Printf("[GetJobList] Unmarshal failed err: %v, body %v \n", err, string(body))
		return nil, err
	}
	if getJobResp.Code != 0 && getJobResp.Message != "ok" {
		log.Printf("[GetJobList] Response err: %v, resp: %v \n", err, getJobResp)
		return nil, err
	}
	return getJobResp.Data.JobPostList, nil
}

func SaveJson(jobInfoList []*JobInfo) error {
	filePtr, err := os.Create("job_info_list.json")
	if err != nil {
		fmt.Printf("[SaveJson] os.Create failed, err: %v \n", err)
		return err
	}
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(jobInfoList)
	if err != nil {
		fmt.Printf("[SaveJson] encoder.Encode failed, err: %v \n", err)
		return err
	}
	return nil
}