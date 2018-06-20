package sdk

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"harborclient/conf"
	"harborclient/logger"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sort"
	"strings"
	"time"
)

var HClient *HarborClient = new(HarborClient)
var HttpClient *http.Client

type HarborClient struct {
	Protocol string
	Host     string
	UserId   string
	Pwd      string
	//SessionId string
}

// login harbor
func (c *HarborClient) Login() *http.Client {

	loginUrl := fmt.Sprintf("%s://%s/login", c.Protocol, c.Host)
	loginData := url.Values{
		"principal": {c.UserId},
		"password":  {c.Pwd},
	}

	//封装请求
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	//启用cookie
	client.Jar, _ = cookiejar.New(nil)

	//发起post请求
	resp, err := client.PostForm(loginUrl, loginData)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	if strings.Contains("200", resp.Status) {
		logger.Logging.WithFields(logrus.Fields{"resp Status": resp.Status}).Fatal("login fail !")
	}

	return client
}

// get harbor all project id []*int
func (c *HarborClient) GetProjectId(client *http.Client) []int {
	targetUrl := fmt.Sprintf("%s://%s/api/projects", HClient.Protocol, HClient.Host)

	resp, err := client.Get(targetUrl)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	var projectId []int
	if data, err := ioutil.ReadAll(resp.Body); err == nil {
		json.Unmarshal(data, &Projects)
		for _, p := range Projects {
			projectId = append(projectId, p.Project_id)
		}
	}
	return projectId
}

// get harbor all repo []string
func (c *HarborClient) GetRepoName(client *http.Client) []string {
	var targetUrl []string

	projectId := HClient.GetProjectId(client)

	for _, p := range projectId {
		targetUrl = append(targetUrl, fmt.Sprintf("%s://%s/api/repositories?project_id=%d", HClient.Protocol, HClient.Host, p))
	}

	var repo []string
	for _, t := range targetUrl {
		resp, err := client.Get(t)
		defer resp.Body.Close()
		if err != nil {
			panic(err)
		}

		if data, err := ioutil.ReadAll(resp.Body); err == nil {
			json.Unmarshal(data, &Repositorys)
			for _, p := range Repositorys {
				repo = append(repo, p.Name)
			}
		}
	}
	return repo
}

// get harbor repo's tag
func (c *HarborClient) GetTag(client *http.Client, repo string) []string {

	targetUrl := fmt.Sprintf("%s://%s/api/repositories/%s/tags", HClient.Protocol, HClient.Host, repo)

	resp, err := client.Get(targetUrl)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	// 已经按照创建时间排序的tag
	var tagAfterSort []string

	tagMap := make(map[string]tagDate)

	if data, err := ioutil.ReadAll(resp.Body); err == nil {

		json.Unmarshal(data, &DetailedTags)

		for _, p := range DetailedTags {
			// Parse string to time
			tagCreateTime, err := time.Parse(time.RFC3339, p.Created)
			if err != nil {
				logger.Logging.Warn(err, p.Created)
			}

			// 添加到map
			tagMap[p.Name] = tagDate{tagId: p.Name, tagCreateTime: tagCreateTime}
		}
	}

	//将tagMap转化为[]slice并按照创建时间排序
	dateSortReviews := make(tagSlice, 0, len(tagMap))
	for _, d := range tagMap {
		dateSortReviews = append(dateSortReviews, d)
	}

	//按照时间排序
	sort.Sort(dateSortReviews)

	for _, k := range dateSortReviews {
		tagAfterSort = append(tagAfterSort, k.tagId)
	}

	return tagAfterSort
}

// delete harbor repo's tag
func (c *HarborClient) DeleteRepoTag(client *http.Client, repo, tag string) {
	targetUrl := fmt.Sprintf("%s://%s/api/repositories/%s/tags/%s", HClient.Protocol, HClient.Host, repo, tag)

	// Create request
	req, err := http.NewRequest("DELETE", targetUrl, nil)
	if err != nil {
		logger.Logging.Fatal(err)
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		logger.Logging.Fatal(err)
	}
	defer resp.Body.Close()

	if strings.Contains("200", resp.Status) {
		logger.Logging.WithFields(logrus.Fields{"resp Status": resp.Status}).Fatal("delete tag fail !")
	} else {
		//logger.Logging.WithFields(logrus.Fields{"repo:": repo, "tag": tag}).Info("delete tag success")
		logger.Logging.Info("delete ", repo, ":", tag, " success !")
	}
}

func init() {
	envParams := conf.GetEnvParams()
	HClient.Protocol = envParams["harbor_protocol"]
	HClient.Host = envParams["harbor_host"]
	HClient.UserId = envParams["harbor_user"]
	HClient.Pwd = envParams["harbor_password"]
}

type tagDate struct {
	tagId         string
	tagCreateTime time.Time
}

type tagSlice []tagDate

func (p tagSlice) Len() int {
	return len(p)
}

func (p tagSlice) Less(i, j int) bool {
	return p[i].tagCreateTime.Before(p[j].tagCreateTime)
}

func (p tagSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
