package main

import (
	"egts_protocol_receiver/internal/egts/packet"
	"encoding/hex"
	"fmt"
)

func main() {
	auth := "0100020B0069000100012562000200000202101800068EFC190B6ABE7BD8CB043E810000552388060005DD0100110400080F0000140500027E002900170200400617020050061702006006170200700617020080061702009006170200A006170200B0061D0D000000000108000000000000000052BD" // авторизация

	pkgBytes, err := hex.DecodeString(auth)
	if err != nil {
		fmt.Print(err)
		//t.Errorf("Error: %s", err.Error())
	}

	pkg := packet.Packet{}

	pkg, err = packet.ReadPacket(pkgBytes)

	fmt.Println(pkg.PacketType)
	/*readPacket, err := packet.ReadPacket(pkgBytes)
	if err != nil {
		return
	}

	z := &readPacket.ServicesFrameData

	fmt.Println(*z)*/

}
