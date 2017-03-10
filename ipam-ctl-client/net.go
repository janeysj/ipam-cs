package main
import (
	"os"
	"fmt"
	"net"
	"io/ioutil"
	"net/http"
    //"strings"
	"github.com/google/uuid"
)


type Net struct {
	Id      uuid.UUID
	StartIp int
	EndIp   int
	//AllocatedIpList map[string]VipIp /* TODO: list */
}

type VipIp struct  {
	IpAddress string /* or int? */
	localIpList []string /* TODO: list */
}


var netList [10] Net

func CreateNet(netString string) (*Net, error){
	_, ipnet, err := net.ParseCIDR(netString)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(os.Stderr, "%v\n", ipnet)
	net := Net{Id: uuid.New(), StartIp:3, EndIp:6}
	netList[0] = net
	fmt.Fprintf(os.Stderr, "%v\n", netList)
	return &net, nil
}

func GetAllNetworks() error {
	for _,net := range netList {
		fmt.Fprintf(os.Stderr, "%v\n", net)
	}
	return nil
}

func GetIP() (string, error)  {
	ret := httpGet()
	return ret, nil
}

func DeleteIP(ipaddr string) (string, error)  {
	ret := httpDelete(ipaddr)
	return ret, nil
}

func httpGet() string{
    resp, err := http.Get("http://127.0.0.1:8081/ip")
    if err != nil {
        // handle error
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    //fmt.Println(string(body))
	return  string(body)
}

func httpDelete(id string) string{
	//client := &http.Client{}
	//bodystr := ""
    //reqest, _ := http.NewRequest("Del", "http://127.0.0.1:8081/{id=123}", nil)
	//
    //reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    //reqest.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
    //reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
    //reqest.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
    //reqest.Header.Set("Cache-Control","max-age=0")
    //reqest.Header.Set("Connection","keep-alive")
	//
    //response,_ := client.Do(reqest)
    //if response.StatusCode == 200 {
    //    body, _ := ioutil.ReadAll(response.Body)
    //    bodystr = string(body);
    //    fmt.Println(bodystr)
    //}
	//return bodystr
	requestStr := fmt.Sprintf("http://127.0.0.1:8081/ip/%s", id)
	resp, err := http.Get(requestStr)
    if err != nil {
        // handle error
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    //fmt.Println(string(body))
	return  string(body)
}