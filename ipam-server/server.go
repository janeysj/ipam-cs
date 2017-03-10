package main

import (
	"os"
	"fmt"
	"net"
    "errors"
    "strconv"
    "strings"
	"github.com/google/uuid"
)

type Net struct {
	Id      uuid.UUID
	StartIp int64
	EndIp   int64
	//AllocatedIpList map[string]VipIp /* TODO: list */
}

type VipIp struct  {
	IpAddress net.IP /* or int? */
	localIpList []net.IP /* TODO: list */
}

var ipnet Net
var ipList map[int64]VipIp

func inet_ntoa(ipnr int64) net.IP {
    var bytes [4]byte
    bytes[0] = byte(ipnr & 0xFF)
    bytes[1] = byte((ipnr >> 8) & 0xFF)
    bytes[2] = byte((ipnr >> 16) & 0xFF)
    bytes[3] = byte((ipnr >> 24) & 0xFF)

    return net.IPv4(bytes[3],bytes[2],bytes[1],bytes[0])
}

func inet_aton(ipnr net.IP) int64 {
    bits := strings.Split(ipnr.String(), ".")

    b0, _ := strconv.Atoi(bits[0])
    b1, _ := strconv.Atoi(bits[1])
    b2, _ := strconv.Atoi(bits[2])
    b3, _ := strconv.Atoi(bits[3])

    var sum int64

    sum += int64(b0) << 24
    sum += int64(b1) << 16
    sum += int64(b2) << 8
    sum += int64(b3)

    return sum
}


func simulateNet() {
	fmt.Fprintf(os.Stderr, "Init simulated net...\n")
    /* TODO: Get data from etcd server. Contains the IP segment and assigned IPs. */
    ip := net.ParseIP( "192.168.1.1")
    startIp := inet_aton(ip)
	ip = net.ParseIP("192.168.1.20")
    endIp := inet_aton(ip)
	ipnet = Net{Id: uuid.New(), StartIp:startIp, EndIp:endIp}
    ipList = make(map[int64]VipIp)
}

func AssignIP() (net.IP, error) {
    var assignedIp int64
    var err error
    for ipInt := ipnet.StartIp; ipInt < ipnet.EndIp; ipInt++ {
        if _, aimIp := ipList[ipInt]; !aimIp {
            assignedIp = ipInt
            break
        }
    }
    if assignedIp!= 0 {
        /* assign IP successfully */
        var localList = [] net.IP{net.ParseIP("192.168.2.1"), net.ParseIP("192.168.2.2")}
        ipList[assignedIp] = VipIp{IpAddress: inet_ntoa(assignedIp), localIpList: localList}
        /* TODO: wrte ipList[assignedIp] to etcd server */
    } else {
        /* TODO: handle error */
        err = errors.New("No available IP.")
    }
    fmt.Fprintf(os.Stderr, "ipList is %v\n", ipList)

    ret := inet_ntoa(assignedIp)
	return ret, err
}

func ReleaseIP(ipstr string) (error) {
    ip := net.ParseIP(ipstr)
    ipInt := inet_aton(ip)
    var err error
    if _, aimIp := ipList[ipInt]; aimIp {
        fmt.Fprintf(os.Stderr, "Delete IP %s\n", ipInt)
        delete(ipList, ipInt)
        /* TODO: delete aimIP from etcd server */
    } else {
        fmt.Fprintf(os.Stderr, "ip %s not found.\n", ipstr)
        /* TODO: handle error */
        err = errors.New("IP not found.")
    }

    fmt.Fprintf(os.Stderr, "ipList is %v\n", ipList)
	return err
}