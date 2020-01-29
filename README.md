func SendMessage(Header string, Data string, Severity int, Facility int, ServerAddrAdndPort string) (sentbytes int, err error)

Header - Message header aka Mnemonic
Data - message text
Severity - from 0 to 7
Facility - from 0 to 23
ServerAddrAdndPort - ip or dns and : port like: 192.168.0.1:514
example: 	sent,err := SendMessage("%PowerC2","Test Msg",3,23,"192.168.5.1:514")
return number of sent bytes and error
error == nil if data was sent
