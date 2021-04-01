// Package hashrate is a message type for Simplebuffers generated by miners to broadcast an IP address, a count and
// version number and current height of mining work just completed. This data should be stored in a log file and added
// together to generate hashrate reporting in nodes when their controller is running
package hashrate

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"net"
	"time"
	
	"github.com/niubaoshu/gotiny"
	
	"github.com/p9c/duod/pkg/routeable"
)

var Magic = []byte{'h', 'a', 's', 1}

//
// type Container struct {
// 	simplebuffer.Container
// }

type Hashrate struct {
	Time    time.Time
	IP      net.IP
	Count   int
	Version int32
	Height  int32
	Nonce   int32
	ID      string
}

func Get(count int32, version int32, height int32, id string) []byte {
	nonce := make([]byte, 4)
	var e error
	if _, e = io.ReadFull(rand.Reader, nonce); E.Chk(e) {
	}
	hr := Hashrate{
		Time:    time.Now(),
		IP:      routeable.GetListenable(),
		Count:   int(count),
		Version: version,
		Height:  height,
		Nonce:   int32(binary.LittleEndian.Uint32(nonce)),
		ID:      id,
	}
	srlz := gotiny.Marshal(&hr)
	// D.S(srlz)
	return srlz
	// return Container{*simplebuffer.Serializers{
	// 	Time.New().Put(time.Now()),
	// 	IPs.GetListenable(),
	// 	Int32.New().Put(count),
	// 	Int32.New().Put(version),
	// 	Int32.New().Put(height),
	// 	Int32.New().Put(int32(binary.BigEndian.Uint32(nonce))),
	// 	String.New().Put(id),
	// }.CreateContainer(Magic)}
}

//
// // LoadContainer takes a message byte slice payload and loads it into a container
// // ready to be decoded
// func LoadContainer(b []byte) (out Container) {
// 	out.Data = b
// 	return
// }
//
// func (j *Container) GetTime() time.Time {
// 	return Time.New().DecodeOne(j.Get(0)).Get()
// }
//
// func (j *Container) GetIPs() []*net.IP {
// 	return IPs.New().DecodeOne(j.Get(1)).Get()
// }
//
// func (j *Container) GetCount() int {
// 	return int(Int32.New().DecodeOne(j.Get(2)).Get())
// }
//
// func (j *Container) GetVersion() int32 {
// 	return Int32.New().DecodeOne(j.Get(3)).Get()
// }
//
// func (j *Container) GetHeight() int32 {
// 	return Int32.New().DecodeOne(j.Get(4)).Get()
// }
//
// func (j *Container) GetNonce() int32 {
// 	return Int32.New().DecodeOne(j.Get(5)).Get()
// }
//
// func (j *Container) GetID() string {
// 	return String.New().DecodeOne(j.Get(6)).Get()
// }
//
// func (j *Container) String() (s string) {
// 	s += fmt.Sprint("\ntype '"+string(Magic)+"' elements:", j.Count())
// 	s += "\n"
// 	t := j.GetTime()
// 	s += "1 Time: "
// 	s += fmt.Sprint(t)
// 	s += "\n"
// 	ips := j.GetIPs()
// 	s += "2 IPs:"
// 	for i := range ips {
// 		s += fmt.Sprint(" ", ips[i].String())
// 	}
// 	s += "\n"
// 	count := j.GetCount()
// 	s += "3 Count: "
// 	s += fmt.Sprint(count)
// 	s += "\n"
// 	version := j.GetVersion()
// 	s += "4 Version: "
// 	s += fmt.Sprint(version)
// 	s += "\n"
// 	nonce := j.GetNonce()
// 	s += "5 Nonce: "
// 	s += fmt.Sprint(nonce)
// 	s += "\n"
// 	ID := j.GetID()
// 	s += "6 ID: "
// 	s += fmt.Sprint(ID)
// 	s += "\n"
// 	return
// }
//
// // Struct deserializes the data all in one go by calling the field deserializing functions into a structure containing
// // the fields. The height is given in this report as it is part of the job message and makes it faster for clients to
// // look up the algorithm name according to the block height, which can change between hard fork versions
// func (j *Container) Struct() (out Hashrate) {
// 	out = Hashrate{
// 		Time:    j.GetTime(),
// 		IPs:     j.GetIPs(),
// 		Count:   j.GetCount(),
// 		Version: j.GetVersion(),
// 		Height:  j.GetHeight(),
// 		Nonce:   j.GetNonce(),
// 		ID:      j.GetID(),
// 	}
// 	return
// }
