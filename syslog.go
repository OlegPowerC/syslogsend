package syslogsend

import (
	"fmt"
	"net"
	"time"
)

const SYSLOGTIME = "Mon Jan _2 2006 15:04:05"

/*
Header - Message header aka Mnemonic
Data - message text
Severity - from 0 to 7
Facility - from 0 to 23
ServerAddrAdndPort - ip or dns and : port like: 192.168.0.1:514
example: 	sent,err := SendMessage("%PowerC2","Test Msg",3,23,"192.168.5.1:514")
return number of sent bytes and error
error == nil if data was sent
*/

func SendMessage(Header string, Data string, Severity int, Facility int, ServerAddrAdndPort string) (sentbytes int, err error) {
	var SyslogMessage string
	if Severity > 7 {
		return 0, fmt.Errorf("Invalid severity %d", Severity)
	}
	if Facility > 23 {
		return 0, fmt.Errorf("Invalid severity %d", Facility)
	}

	t := time.Now()
	fmt.Println(t.Format(SYSLOGTIME))
	var Priority int
	Priority = Facility*8 + Severity
	SyslogMessage = fmt.Sprintf("<%d> %s:%s:%s", Priority, t.Format(SYSLOGTIME), Header, Data)
	raddr, errresolv := net.ResolveUDPAddr("udp", ServerAddrAdndPort)
	if errresolv != nil {
		return 0, errresolv
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	defer conn.Close()
	if err != nil {
		return 0, err
	}
	var fmb []byte
	fmb = []byte(SyslogMessage)
	sentbytes, errsend := conn.Write(fmb)
	if errsend != nil {
		return 0, errsend
	}
	return sentbytes, nil
}
