package server_backend



/*
no consideration for retransmission for now
after testing, we will call Waitack() method for each uw command
*/

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net/smtp"
	"strconv"

	"net"

	uw "protobuf/pb_uw"
	ua "protobuf/pb_ua"
	psql "database_utils"

)

const(
	hostURL = "smtp.gmail.com"
	hostPort = "587"
	password = "wert1234567"
	emailSender = "miniups568@gmail.com"
)


var WorldId int64//world id
var Wconn *net.Conn
var Aconn *net.Conn

//------------------------do and on connection-----------------------------------
func Do_connection_to_world() (error) {
	uconnObj := &uw.UConnect{}
	var i int32
	for i = 1; i < 5 ; i++ {
		//UInitTruck
		utruckObj := &uw.UInitTruck{}
		utruckObj.Id = proto.Int32(i);
		utruckObj.X = proto.Int32(3*i-5);
		utruckObj.Y = proto.Int32(5*i-20);
		uconnObj.IsAmazon = proto.Bool(false);
		uconnObj.Trucks = append(uconnObj.Trucks, utruckObj)
		//insert into database
		psql.Insert_table_truck(*utruckObj.Id, *utruckObj.X, *utruckObj.Y)
	}

	out, err := proto.Marshal(uconnObj)
	if err != nil {
		fmt.Println("Failed to encode UConnect")
	}
	n := proto.Size(uconnObj)
	out = append(proto.EncodeVarint(uint64(n)), out...)
	if _, err = (*Wconn).Write(out); err != nil {
		return err
	}
	return nil
}

func On_connection_to_world() (uw.UConnected, error){
	mesg := make([]byte, 1024)
	size, _ := (*Wconn).Read(mesg)
	mesg = mesg[:size]

	_, n := proto.DecodeVarint(mesg)
	mesg = mesg[n:]
	uconnectedObj := &uw.UConnected{}
	if err := proto.Unmarshal(mesg, uconnectedObj); err != nil {
		return uw.UConnected{}, err
	}
	return *uconnectedObj, nil
}

func Do_send_worldID_to_Amazon(worldId int64) (error) {
	uconnObj := &ua.USendWorldID{}
	uconnObj.Worldid = proto.Int64(worldId)


	out, err := proto.Marshal(uconnObj)
	if err != nil {
		fmt.Println("Failed to encode UConnect")
	}
	n := proto.Size(uconnObj)
	out = append(proto.EncodeVarint(uint64(n)), out...)
	if _, err = (*Aconn).Write(out); err != nil {
		return err
	}
	return nil
}






//-----------------------receive and handle--------------------------------------
func Do_receive_from_world(wchan chan bool) {
	for{
		mesg := make([]byte, 1024)
		size, _ := (*Wconn).Read(mesg)
		if size != 0 {
			mesg = mesg[:size]

			_, n := proto.DecodeVarint(mesg)
			mesg = mesg[n:]
			uresponseObj := &uw.UResponses{}
			if err := proto.Unmarshal(mesg, uresponseObj); err != nil {
				fmt.Println(err)
			} else {
				go On_receive_from_world(*uresponseObj, wchan)
			}
		}
	}
}

func On_receive_from_world(uresponses uw.UResponses, wchan chan bool) {

	var uwcommands uw.UCommands
	var uacommands ua.UtoAResponses
	//check Finished
	if uresponses.Finished != nil && *uresponses.Finished {
		wchan <- true
		return
	}
	//Ack handling
	for _, ack := range uresponses.Acks {
		fmt.Printf("<world> <ack>: %d\n", ack )

		if result, err := psql.Query_request_gopickup(ack); err != nil {
			fmt.Printf("<world> <ack> <UGoPickup> request not found")
		} else {
			fmt.Printf("<world> <ack> <UGoPickup>: %d\n", ack )
			var pkgs_db []psql.Package
			var pkgs []*ua.APackageInfo
			pkgs_db, err = psql.Query_table_packages_byTruckId(result.TruckId)
			if err == nil {
				for _, pkg_db := range pkgs_db {
					var pkg ua.APackageInfo
					pkg.Packageid = proto.Int64(pkg_db.PackageID)
					pkg.X = proto.Int32(pkg_db.X)
					pkg.Y = proto.Int32(pkg_db.Y)
					pkgs = append(pkgs, &pkg)
				}
			} else {
				fmt.Printf("no packages found on this truck")
			}
			//1.set UTrucksend command
			var utrucksent ua.UTruckSent
			utrucksent.Truckid = proto.Int32(result.TruckId)
			utrucksent.Packages = pkgs

			uacommands.Trucksent = append(uacommands.Trucksent, &utrucksent)
			fmt.Printf("<world> send truck %v\n", utrucksent)
			//2.delete corresponding ugopicku command
			psql.Delete_request_ugopickup(ack) //delete from database
		}

		if result, err := psql.Query_request_godeliver(ack); err == nil {
			for _, pkg := range result.Pkgs{

				var udelivered ua.UDelivered
				udelivered.Shipid = proto.Int64(pkg.PackId)
				psql.Update_table_package_location(pkg.PackId, pkg.X, pkg.Y)//user may update destination
				uacommands.Delivered = append(uacommands.Delivered, &udelivered)
			}

			psql.Delete_request_ugodeliver(ack) //delete from database
		}

		if _, err := psql.Query_request_goquery(ack); err == nil {
			psql.Delete_request_uquery(ack)
		}

	}

	//completion 1) arrive warehouse 2) finish all deliveries
	if uresponses.Completions != nil {
		/*
			5 truck status:
			--idle
			--traveling
			--arrive warehouse
			--loading
			--delivering
		*/
		for _, completion := range uresponses.Completions {
			//1.return world ack # same as completion's seq #
			uwcommands.Acks = append(uwcommands.Acks, *completion.Seqnum)

			//2.send UTruckArrived if truck status is arrive warehouse
			fmt.Printf("<world> <completion> <UFinished>, truck id: %d\n", *completion.Truckid)
			fmt.Printf("<world> <completion> <UFinished>, truck status: %s\n", *completion.Status)
			fmt.Printf("<world> <completion> <UFinished>, truck location: (%d,%d)\n", *completion.X, *completion.Y)
			if *completion.Status == "IDLE" {

			} else {//truck status: ARRIVE WAREHOUSE
				var utruckarrived ua.UTruckArrived
				utruckarrived.Truckid = proto.Int32(*completion.Truckid)
				uacommands.Arrived = append(uacommands.Arrived, &utruckarrived)
				fmt.Printf("<world> <UTruckArrived> %v\n", utruckarrived)
			}
			//3.update truck info to database
			psql.Update_table_truck_location_and_status(*completion.Truckid, *completion.X, *completion.Y, *completion.Status)

		}
	}

	if uresponses.Delivered != nil {
		for _, delivered := range uresponses.Delivered {
			fmt.Printf("<world> <delivered> %v\n", delivered)
			//1.return ack
			uwcommands.Acks = append(uwcommands.Acks, *delivered.Seqnum)

			//2.update package status
			psql.Update_table_package_status(*delivered.Packageid, "DELIVERED")

			//3.send email to user
			pkg, err := psql.Query_table_package_byID(*delivered.Packageid)
			_, _, emailReceiver, err := psql.Query_table_user(pkg.UserUsername)
			if err == nil {
				emailAuth := smtp.PlainAuth (
					"",
					emailSender,
					password,
					hostURL)
				msg := []byte("To: " + emailReceiver + "\r\n" + "Subject: Packaged delivered, id:"+ strconv.FormatInt(*delivered.Packageid, 10) +"\r\n")

				err := smtp.SendMail(
					hostURL + ":"+ hostPort,
					emailAuth,
					emailSender,
					[]string{emailReceiver},
					msg)
				if err != nil {
					fmt.Println(err)
				}
			}

		}
	}

	if uresponses.Truckstatus != nil {
		for _, truckst := range uresponses.Truckstatus {
			fmt.Printf("<world> <truct status> %v\n", truckst)
			//1.return ack
			uwcommands.Acks = append(uwcommands.Acks, *truckst.Seqnum)
			//2.update truck info
			psql.Update_table_truck_location_and_status(*truckst.Truckid, *truckst.X, *truckst.Y, *truckst.Status)
		}
	}

	if uresponses.Error != nil {
		for _, err := range uresponses.Error {
			//1.return ack
			uwcommands.Acks = append(uwcommands.Acks, *err.Seqnum)
			//2.print error message
			fmt.Printf("<world> error: %s", err)
		}
	}

	//fmt.Printf("<world> uacommand sent %v\n", uacommands)
	out, err := proto.Marshal(&uacommands)
	if err != nil {
		fmt.Println("Failed to encode UConnect")
	}
	n := proto.Size(&uacommands)
	out = append(proto.EncodeVarint(uint64(n)), out...)
	if _, err = (*Aconn).Write(out); err != nil {
		log.Fatal("fail to send ua commands in on_receive_from_world")
	}

	//fmt.Printf("<world> uwcommand sent %v\n", uwcommands)
	out, err = proto.Marshal(&uwcommands)
	if err != nil {
		fmt.Println("Failed to encode UConnect")
	}
	n = proto.Size(&uwcommands)
	out = append(proto.EncodeVarint(uint64(n)), out...)
	if _, err = (*Wconn).Write(out); err != nil {
		log.Fatal("fail to send uw commands in on_receive_from_world")
	}

}

func Do_receive_from_amazon(achan chan bool) {
	for{
		mesg := make([]byte, 1024)
		size, _ := (*Aconn).Read(mesg)
		if size != 0 {
			mesg = mesg[:size]

			_, n := proto.DecodeVarint(mesg)
			mesg = mesg[n:]
			aturesponseObj := &ua.AtoUCommands{}
			if err := proto.Unmarshal(mesg, aturesponseObj); err != nil {
				fmt.Println(err)
			} else {
				go On_receive_from_amazon(*aturesponseObj,achan)
			}
		}
	}
}


func On_receive_from_amazon(atucommands ua.AtoUCommands, achan chan bool) {
	var uwcommands uw.UCommands
	if atucommands.Sendtrucks != nil {
		for _, sendTruck := range atucommands.Sendtrucks {
			fmt.Printf("<amazon> <sendtrucks> %v\n",sendTruck)

			//select truck
			truckid, err := Do_select_and_pick_truck(*sendTruck.Whs.Id)
			fmt.Printf("<amazon> select truck id %d\n", truckid)

			//1.insert packages info to database
			wareHouse := sendTruck.Whs
			whid := *wareHouse.Id

			for _, pkg := range sendTruck.Packages {
				if pkg.UpsUserName != nil {//since username is optional
					psql.Insert_table_package_and_update_truck_towhid(*pkg.Packageid, "unloaded", whid, *pkg.X, *pkg.Y, truckid, *pkg.UpsUserName)
				} else {
					psql.Insert_table_package_and_update_truck_towhid(*pkg.Packageid, "unloaded", whid, *pkg.X, *pkg.Y, truckid, "")
				}
			}

			//2.set UGoPickup command
			var ugopickup uw.UGoPickup
			ugopickup.Truckid = proto.Int32(truckid)
			ugopickup.Whid = proto.Int32(whid)

			latestSeqNum, err := psql.Increment_seq_number(WorldId)//retrieve latest seq # and update it
			fmt.Printf("<amazon> <GO PICK UP> seq #: %d\n",latestSeqNum)
			if err != nil {
				log.Fatal(err)
			} else {
				ugopickup.Seqnum = proto.Int64(latestSeqNum)
				uwcommands.Pickups  = append(uwcommands.Pickups, &ugopickup)
				//3.timer for retransmission
				ugopickup.WaitAck(Wconn)
				//4.insert request to database
				psql.Insert_request_gopickup(latestSeqNum, *ugopickup.Truckid, *ugopickup.Whid)
			}
		}
	}
	//
	////since Amazon guarantess all packages in AStartDelivery are loaded and identical those in ASendTruck
	////we will not update the databse, and use the information in AStartDelivery directly
	if atucommands.Startdelivery != nil {
		for _, startDelivery := range atucommands.Startdelivery {
			fmt.Printf("<amazon> <startdelivery> %v\n",startDelivery)
			//1.update package info in table
			var pkgs []*uw.UDeliveryLocation
			var pkgs_db []psql.UDeliveryLocation
			for _, pkg := range startDelivery.Packages {
				var deliverylocation uw.UDeliveryLocation
				var deliverylocation_db psql.UDeliveryLocation
				deliverylocation.Packageid = pkg.Packageid
				deliverylocation.X = pkg.X
				deliverylocation.Y = pkg.Y
				deliverylocation_db.PackId = *pkg.Packageid
				deliverylocation_db.X = *pkg.X
				deliverylocation_db.Y = *pkg.Y
				deliverylocation_db.RequestUGoDeliverTruckId = *startDelivery.Truckid
				if pkg.UpsUserName != nil {
					username := pkg.UpsUserName
					psql.Update_table_package_username(*deliverylocation.Packageid, *username)
				}
				psql.Update_table_package_status(*deliverylocation.Packageid, "DELIVERING")
				pkgs = append(pkgs, &deliverylocation)
				pkgs_db = append(pkgs_db, deliverylocation_db)
			}
			//2.send UGoDeliver command
			var ugodelivery uw.UGoDeliver
			ugodelivery.Truckid = proto.Int32(*startDelivery.Truckid)
			ugodelivery.Packages = pkgs
			latestSeqNum, err := psql.Increment_seq_number(WorldId)
			if err != nil {
				log.Fatal(err)
			} else {
				ugodelivery.Seqnum = proto.Int64(latestSeqNum)
				uwcommands.Deliveries = append(uwcommands.Deliveries, &ugodelivery)
				//3.timer for retransmission
				ugodelivery.WaitAck(Wconn)
				//4.insert request to database
				psql.Insert_request_godeliver(latestSeqNum, *ugodelivery.Truckid, pkgs_db)
			}
		}
	}


	out, err := proto.Marshal(&uwcommands)
	if err != nil {
		fmt.Println("Failed to encode UConnect")
	}
	n := proto.Size(&uwcommands)
	out = append(proto.EncodeVarint(uint64(n)), out...)
	if _, err = (*Wconn).Write(out); err != nil {
		log.Fatal("fail to send uw commands in on_receive_from_world")
	}

}

func Do_select_and_pick_truck(whid int32) (int32, error) {
	for{
		if truckid, err := psql.Select_truck(whid); err != nil {

		} else {
			return truckid, nil
		}
	}
}

