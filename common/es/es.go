package es

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mu-Exchange/Mu-End/common/errors"
	"github.com/opensearch-project/opensearch-go/v2"

	"github.com/Mu-Exchange/Mu-End/common/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/mitchellh/mapstructure"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	requestsigner "github.com/opensearch-project/opensearch-go/v2/signer/awsv2"
)

type ES struct {
	client *opensearch.Client
	ctx    context.Context
}

func New(cfg *config.ES) (*ES, error) {
	if cfg.EnableAws {
		awsCfg, err := awsconfig.LoadDefaultConfig(context.Background(),
			awsconfig.WithRegion(cfg.AwsES.Region),
			awsconfig.WithCredentialsProvider(

				credentials.StaticCredentialsProvider{Value: aws.Credentials{
					AccessKeyID:     cfg.AwsES.AccessKeyId,
					SecretAccessKey: cfg.AwsES.AccessKeySecret,
				}},
			),
		)
		if err != nil {
			return nil, err
		}

		// CreatePublish an AWS request Signer and load AWS configuration using default config folder or env vars.
		signer, err := requestsigner.NewSignerWithService(awsCfg, "es")
		if err != nil {
			return nil, err
		}

		// CreatePublish an opensearch client and use the request-signer
		client, err := opensearch.NewClient(opensearch.Config{
			Addresses: cfg.AwsES.Address,
			Signer:    signer,
		})
		if err != nil {
			return nil, err
		}

		return &ES{
			client: client,
			ctx:    context.Background(),
		}, nil
	}

	es, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.NativeES.SkipTLSVerify},
		},
		Addresses: cfg.NativeES.Address,
		Username:  cfg.NativeES.Username,
		Password:  cfg.NativeES.Password,
	})
	if err != nil {
		return nil, err
	}

	return &ES{
		client: es,
		ctx:    context.Background(),
	}, nil
}

func getCredentialProvider(accessKey, secretAccessKey, token string) aws.CredentialsProviderFunc {
	return func(ctx context.Context) (aws.Credentials, error) {
		c := &aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretAccessKey,
			SessionToken:    token,
		}
		return *c, nil
	}
}

func (m *ES) SetData(index string, id string, data interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return fmt.Errorf("encoding query: %w", err)
	}
	req := opensearchapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       &buf,
	}
	res, err := req.Do(context.Background(), m.client)
	if err != nil {
		return fmt.Errorf("index response: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return fmt.Errorf("error parsing the response body: %w", err)
		} else {
			return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		}
	}

	return nil
}

func (m *ES) GetData(index string, id string, data interface{}) error {
	get := opensearchapi.GetRequest{
		Index:      index,
		DocumentID: id,
	}

	res, err := get.Do(context.Background(), m.client)
	if err != nil {
		return fmt.Errorf("get response: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return errors.ErrESNotFound
	}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return fmt.Errorf("error parsing the response body: %w", err)
		} else {
			return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return fmt.Errorf("error parsing the response body: %w", err)
	}

	if r["found"] == false {
		return errors.ErrESNotFound
	}

	if err := mapstructure.Decode(r["_source"], data); err != nil {
		return fmt.Errorf("error parsing the response body: %w", err)
	}

	return nil
}

func (m *ES) MGetData(index string, ids []string, data interface{}) error {
	//for i := range ids {
	//	ids[i] = fmt.Sprintf(`{"_id": "%s"}`, ids[i])
	//}
	//
	//body := strings.Join(ids, ",")
	//body = fmt.Sprintf(`{ "docs": [%s] }`, body)

	var idsM []map[string]interface{}
	for _, id := range ids {
		idsM = append(idsM, map[string]interface{}{
			"_id": id,
		})
	}

	body := map[string]interface{}{
		"docs": idsM,
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return fmt.Errorf("encoding mget: %w", err)
	}

	mget := opensearchapi.MgetRequest{
		Index: index,
		Body:  &buf,
	}

	res, err := mget.Do(context.Background(), m.client)
	if err != nil {
		return fmt.Errorf("get response: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return fmt.Errorf("error parsing the response body: %w", err)
		} else {
			return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return fmt.Errorf("error parsing the response body: %w", err)
	}

	docs := r["docs"].([]interface{})
	result := make(map[string]interface{})
	for _, doc := range docs {
		d := doc.(map[string]interface{})
		if d["found"] == false {
			continue
		}

		result[d["_id"].(string)] = d["_source"]
	}

	if err := mapstructure.Decode(result, data); err != nil {
		return fmt.Errorf("error parsing the response body: %w", err)
	}

	return nil
}

func (m *ES) DeleteData(index string, id string) error {
	req := opensearchapi.DeleteRequest{
		Index:      index,
		DocumentID: id,
	}

	res, err := req.Do(context.Background(), m.client)
	if err != nil {
		return fmt.Errorf("delete response: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return nil
	}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return fmt.Errorf("error parsing the response body: %w", err)
		} else {
			return fmt.Errorf("[%s] delete data response: %v", res.Status(), e)
		}
	}

	return nil
}

func (m *ES) SearchData(index string, condition map[string]interface{}, data interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(condition); err != nil {
		return fmt.Errorf("encoding query: %w", err)
	}

	req := opensearchapi.SearchRequest{
		Index: []string{index},
		Body:  &buf,
	}

	res, err := req.Do(context.Background(), m.client)
	if err != nil {
		return fmt.Errorf("search response: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return fmt.Errorf("error parsing the response body: %w", err)
		} else {
			return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return fmt.Errorf("error parsing the response body: %w", err)
	}

	total := r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)
	if total == 0 {
		return errors.ErrESNotFound
	}

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	var result []map[string]interface{}
	for _, hit := range hits {
		result = append(result, hit.(map[string]interface{})["_source"].(map[string]interface{}))
	}

	if err := mapstructure.Decode(result, data); err != nil {
		return fmt.Errorf("error parsing the response body: %w", err)
	}

	return nil

}

func (m *ES) SearchDataWithAgg(index string, condition map[string]interface{}, aggKey string, data *[]string) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(condition); err != nil {
		return fmt.Errorf("encoding query: %w", err)
	}

	req := opensearchapi.SearchRequest{
		Index: []string{index},
		Body:  &buf,
	}

	res, err := req.Do(context.Background(), m.client)
	if err != nil {
		return fmt.Errorf("search response: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return fmt.Errorf("error parsing the response body: %w", err)
		} else {
			return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return fmt.Errorf("error parsing the response body: %w", err)
	}

	total := r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)
	if total == 0 {
		return errors.ErrESNotFound
	}

	buckets := r["aggregations"].(map[string]interface{})[aggKey].(map[string]interface{})["buckets"].([]interface{})
	for _, bucket := range buckets {
		*data = append(*data, bucket.(map[string]interface{})["key"].(string))
	}

	return nil
}

func (m *ES) SearchDataWithAgg2(index string, condition map[string]interface{}) (interface{}, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(condition); err != nil {
		return nil, fmt.Errorf("encoding query: %w", err)
	}

	req := opensearchapi.SearchRequest{
		Index: []string{index},
		Body:  &buf,
	}

	res, err := req.Do(context.Background(), m.client)
	if err != nil {
		return nil, fmt.Errorf("search response: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, fmt.Errorf("error parsing the response body: %w", err)
		} else {
			return nil, fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %w", err)
	}

	total := r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)
	if total == 0 {
		return nil, errors.ErrESNotFound
	}

	return r["aggregations"], nil
}

func (m *ES) SearchByKeywords(index string, fields, keywords []string, data interface{}) error {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keywords,
				"fields": fields,
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return fmt.Errorf("encoding query: %w", err)
	}

	req := opensearchapi.SearchRequest{
		Index: []string{index},
		Body:  &buf,
	}

	res, err := req.Do(context.Background(), m.client)
	if err != nil {
		return fmt.Errorf("search response: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return fmt.Errorf("error parsing the response body: %w", err)
		} else {
			return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return fmt.Errorf("error parsing the response body: %w", err)
	}

	if err := mapstructure.Decode(r["hits"].(map[string]interface{})["hits"], data); err != nil {
		return fmt.Errorf("error parsing the response body: %w", err)
	}

	return nil
}

func (m *ES) CreateIndex(index string, mapping map[string]interface{}) error {
	req := opensearchapi.IndicesCreateRequest{
		Index: index,
	}

	if mapping != nil {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(mapping); err != nil {
			return fmt.Errorf("encoding mapping: %w", err)
		}
		req.Body = &buf
	}

	res, err := req.Do(context.Background(), m.client)
	if err != nil {
		return fmt.Errorf("create index response: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return fmt.Errorf("error parsing the response body: %w", err)
		} else {
			return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		}
	}

	return nil
}

func (m *ES) IndexExists(index string) (bool, error) {
	req := opensearchapi.IndicesExistsRequest{
		Index: []string{index},
	}

	res, err := req.Do(context.Background(), m.client)
	if err != nil {
		return false, fmt.Errorf("index exists response: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return false, nil
	}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return false, fmt.Errorf("error parsing the response body: %w", err)
		} else {
			return false, fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		}
	}

	return res.StatusCode == 200, nil
}

func (m *ES) CreateIndexIfNotExists(index string, mapping map[string]interface{}) error {
	exists, err := m.IndexExists(index)
	if err != nil {
		return fmt.Errorf("check if index %s exists: %w", index, err)
	}

	if !exists {
		if err := m.CreateIndex(index, mapping); err != nil {
			return fmt.Errorf("create index: %w", err)
		}
	}

	return nil
}

func (m *ES) DeleteIndex(index string) error {
	req := opensearchapi.IndicesDeleteRequest{
		Index: []string{index},
	}

	res, err := req.Do(context.Background(), m.client)
	if err != nil {
		return fmt.Errorf("delete index response: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return fmt.Errorf("error parsing the response body: %w", err)
		} else {
			return fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		}
	}

	return nil
}
