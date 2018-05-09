package tools

import (
	"fmt"
	"reflect"
	"encoding/json"
	"encoding/xml"
	"testing"
)

type GetAVFDPolicyByAirLineReq struct {
	XMLName          xml.Name `xml:"GetAVPolicyByAirLineReq"`
	AirCode          string   `xml:"AirCode"`               // 航空公司代号.
	DepartureAirport string   `xml:"DepartureAirport"`      // 出发地三字码.
	ArrivalAirport   string   `xml:"ArrivalAirport"`        // 抵达地三字码.
	FlightSdate      string   `xml:"FlightSDate"`           // 航班开始日期yyyy-mm-dd HH:MM.
	FlightEdate      string   `xml:"FlightEDate,omitempty"` // 回程出发日期,无需设置值,因为目前只允许单程.
	AirType          string   `xml:"AirType"`               //  1-单程,2-往返. 因为目前只允许单程.
	PassengerType    string   `xml:"PassengerType"`         // 乘机人类型 1-成人，2-儿童 3-婴儿.
}

var req GetAVFDPolicyByAirLineReq

func init() {
	req := GetAVFDPolicyByAirLineReq{}
	req.AirCode = "CA"
	req.DepartureAirport = "PEK"
	req.ArrivalAirport = "SHA"
	req.FlightSdate = "2018-05-20 10:00"
	req.AirType = "1"
	req.PassengerType = "1"
}

// 测试任意接口对象和[]byte互转
func TestBytesOfInterface(t *testing.T) {
	bBuf, err := ToBytes(req)
	if err != nil {
		t.Fatal(err)
	}
	bReq := GetAVFDPolicyByAirLineReq{}
	fmt.Printf("bBuf:%x\n", bBuf)
	FromBytes(bBuf, &bReq)
	if !reflect.DeepEqual(req, bReq) {
		t.Fatalf("gob bytes not equal")
	}
}
// 测试任意接口对象和json []byte互转
func TestJson(t *testing.T) {
	jBuf, _ := json.Marshal(req)
	fmt.Printf("jBuf:%x\n", jBuf)
	jBuf1 := []byte(string(jBuf))
	fmt.Printf("jBuf1:%x\n", jBuf1)
	req1 := GetAVFDPolicyByAirLineReq{}
	json.Unmarshal(jBuf1, &req1)
	if !reflect.DeepEqual(req1, req) {
		t.Fatalf("not equal")
	}
}
