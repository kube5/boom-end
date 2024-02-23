package social

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test(t *testing.T) {

	url := "https://api.twitter.com/2/tweets/1750465302226247855?tweet.fields=attachments%2Cauthor_id%2Ccreated_at%2Centities%2Cgeo%2Cid%2Cin_reply_to_user_id%2Clang%2Cpossibly_sensitive%2Creferenced_tweets%2Csource%2Ctext%2Cwithheld"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAAAJ3rgEAAAAAqKALKb4ArFquhJdFOHXQk4ACgSU%3D4Go1UufqB7PMuUtDrZ60rfA7Yx4G1qPwlvqQla6chQwLOBNRt2")
	//req.Header.Add("Cookie", "guest_id=v1%3A170435269132382336; guest_id_ads=v1%3A170435269132382336; guest_id_marketing=v1%3A170435269132382336; personalization_id=\"v1_Y+pfgyqZ+6270GYfVUnK/Q==\"")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
