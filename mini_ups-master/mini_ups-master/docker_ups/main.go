package main

import (
	"fmt"
	"net"
	"html/template"
	psql "database_utils"
	back "server_backend"
	front "server_frontend"
)

func run_back_end_server(channel_main chan bool) {


	for{
		if err := back.Do_connection_to_world(); err == nil {
			break
		}
	}
	fmt.Println("<UConnect>")
	for{
		if uconnObj, err := back.On_connection_to_world(); err == nil {
			back.WorldId = *uconnObj.Worldid
			front.WorldId = *uconnObj.Worldid
			back.Do_send_worldID_to_Amazon(*uconnObj.Worldid)
			fmt.Printf("<UConnected> world id: %d\n",*uconnObj.Worldid)
			break
		}
	}
	//receive from world and amazon
	if _, err := psql.Query_seq_number(back.WorldId); err != nil {
		psql.Insert_seq_number(0)
	}

	channel_world_term := make(chan bool)
	channel_amazon_term := make(chan bool)
	//
	go back.Do_receive_from_world(channel_world_term)
	go back.Do_receive_from_amazon(channel_amazon_term)

	<-channel_world_term


	close(channel_main)
}

func init() {
	front.Tpl = template.Must(template.ParseGlob("server_frontend/templates/*"))
}

func run_front_end_server(channel_main chan bool) {

	front.Run()

	close(channel_main)
}



func main() {
	//------------1.initialize database-------------------------
	if err:= psql.Initialize_database_and_create_tables(); err != nil {
		panic(err)
	}
	defer psql.Close_database()

	channel_front_and_back_term := make(chan bool)//remove the blocking when front-end or back-end terminates
	//-----------2.connection to the world-------------------
	worldConnection, err := net.Dial("tcp", "server:12345")
	if err != nil {
		panic(err)
	}
	wconn := &worldConnection
	back.Wconn = wconn//pass world fd to back-end
	front.Wconn = wconn//pass world fd to front-end

	fmt.Println("<ups connect to world>")

	go run_front_end_server(channel_front_and_back_term)
	///-----------3.connection to Amazon-----------------------
	amazonListener, err := net.Listen("tcp", ":23333")
	if err != nil {
		panic(err)
	}
	fmt.Println("<web listen at 23333>")

	amazonConnection, err := amazonListener.Accept()
	if err != nil {
		panic(err)
	}
	aconn := &amazonConnection
	back.Aconn = aconn//pass amazon fd to back-end
	fmt.Println("<ups connect to amazon>")


	go run_back_end_server(channel_front_and_back_term)


	<-channel_front_and_back_term

	(*wconn).Close()
	(*aconn).Close()
	fmt.Println("<done>")
}
