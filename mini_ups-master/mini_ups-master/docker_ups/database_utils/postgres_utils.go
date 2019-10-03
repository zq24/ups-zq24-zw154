package database_utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"sync"
	"time"
)

const(
	host = "db"
	port = 5432
	user = "postgres"
	password = "12345"
	dbname = "postgres"
)

var seqNumberLock sync.Mutex
var truckLock sync.Mutex
var packageLock sync.Mutex

//----------------------------table schema---------------------------------


type User struct {
	Username string `gorm:"NOT NULL;unique_index;FOREIGNKEY:UserUsername;ASSOCIATION_FOREIGNKEY:Username;"`
	Password string	`gorm:"NOT NULL"`
	Email string
}
//---------------------------------------------------------------
type WorldSeq struct {
	WorldId int64 `gorm:"primary_key"`
	SeqNum int64 `gorm:"default:0"`
}

type RequestUGoPickUp struct {
	SeqNumber int64
	TruckId int32
	Whid int32
}

type UDeliveryLocation struct {
	PackId int64
	X int32
	Y int32
	RequestUGoDeliverTruckId int32
}

type RequestUGoDeliver struct {
	SeqNumber int64
	TruckId int32

	Pkgs []UDeliveryLocation `gorm:"FOREIGNKEY:RequestUGoDeliverTruckId;ASSOCIATION_FOREIGNKEY:TruckId;"`
}

type RequestUQuery struct {
	SeqNumber int64
	Truckid int32
}
//--------------------------------------------------------------
type Truck struct {
	TruckId int32 `gorm:"primary_key;INDEX"`//WARNING: no space between ":" and "primary_key"
	X_lasttime int32 `gorm:"NOT NULL"`
	Y_lasttime int32 `gorm:"NOT NULL"`
	Status string `gorm:"DEFAULT:'IDLE'"`

	TowhId int32 //for APurchaseMore without UGoPickUp
	Packages []Package //for deciding which package to deliver next, since priority is considered
	/*
		so field TruckID in Package
	*/
}

type Package struct {
	PackageID int64 `gorm:"primary_key;INDEX"`
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time

	X int32
	Y int32

	WhID int32 `gorm:"NOT NULL"`//from which warehouse
	TruckID int32 `gorm:"NOT NULL"`//on which truck
	UserUsername string
	//Prioirty int32//for UGoPickUp Intervention
}

//-------------------------table manipulations--------------------------------
//--------------------------create--------------------------------------------
//--------------------------insert--------------------------------------------
//--------------------------update--------------------------------------------
//--------------------------delete--------------------------------------------
//--------------------------query---------------------------------------------

var db *(gorm.DB)
var err error
/*
if you are trying to use global variable, don't use ":=" in the function, it will
have no warning or report, but the variable with same name is not the global variable we declared

alternative is to pass the db to each function
*/

func Initialize_database_and_create_tables() (error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = gorm.Open("postgres", psqlInfo);
	if err != nil {
		return err
	}

	if err = db.AutoMigrate(&WorldSeq{}).Error; err != nil {
		return err
	}
	if err = db.AutoMigrate(&Truck{}).Error; err != nil {
		return err
	}
	if err = db.AutoMigrate(&Package{}).Error; err != nil {
		return err
	}
	if err = db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}


	if err = db.AutoMigrate(&RequestUGoPickUp{}).Error; err != nil {
		return err
	}
	if err = db.AutoMigrate(&UDeliveryLocation{}).Error; err != nil {
		return err
	}
	if err = db.AutoMigrate(&RequestUGoDeliver{}).Error; err != nil {
		return err
	}
	if err = db.AutoMigrate(&RequestUQuery{}).Error; err != nil {
		return err
	}
	return nil
}

func Close_database() {
	db.Close()
}


func Insert_seq_number (worldid int64) (error) {
	worldSeq := WorldSeq{WorldId:worldid, SeqNum: 0}
	if err = db.Create(&worldSeq).Error; err != nil {
		return err
	}
	return nil
}

/*
USAGE
user creating account
INPUT
--username
--passwd
OUTPUT
--error
*/
func Insert_table_user( username string, passwd string, email string) (error){
	if err = db.Create(&User{Username:username, Password:passwd, Email:email}).Error; err != nil {
		return err
	}
	return nil
}


/*
USAGE
when trying to UConnect the world, trucks info are provided to the world, and at that time,
trucks info should be inserted into table
the initial state of a truck is "Idle"
TowhID and Packages could be null
INPUT
--truck id
--last location (x, y)
OUTPUT
--error
*/
func Insert_table_truck(truckID int32, x int32, y int32) (error){
	if err = db.Create(&Truck{TruckId: truckID, X_lasttime:x, Y_lasttime: y}).Error; err != nil {
		return err
	}
	return nil
}


/*
USAGE
Before UGoDeliver, recording all the packages bound to a truck
INPUT
--package id
--package status
--username
--from which warehouse
--on which truck
OUTPUT
--error
*/
func Insert_table_package_and_update_truck_towhid(packageID int64, status string, warehouseID int32, x int32, y int32, truckID int32, username string) (error){
	tx := db.Begin()
	if err = tx.Create(&Package{
		PackageID: packageID,
		Status: status,
		WhID: warehouseID,
		X:x,
		Y:y,
		TruckID: truckID,
		UserUsername:username}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("fail to insert package")
	}

	var trucks Truck
	if err = tx.Model(&trucks).Where("truck_id=?", truckID).Updates(
		Truck{TowhId: warehouseID}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("fail to update truck towhid")
	}
	return tx.Commit().Error
}


/*
USAGE
save UGoPickUp request in database
*/
func Insert_request_gopickup(seqNum int64, truckid int32, whid int32) (error){
	if err = db.Create(&RequestUGoPickUp{
		SeqNumber: seqNum,
		TruckId: truckid,
		Whid: whid}).Error; err != nil {
		return err
	}
	return nil
}



/*
USAGE
save UGoDeliver request in database
*/
func Insert_request_godeliver(seqNum int64, truckid int32, pkgs []UDeliveryLocation) (error){
	for _, pkg := range pkgs {
		pkg.RequestUGoDeliverTruckId = truckid
	}
	if err = db.Create(&RequestUGoDeliver{
		SeqNumber: seqNum,
		TruckId: truckid,
		Pkgs:pkgs}).Error; err != nil {
		return err
	}

	return nil
}

/*
USAGE
save UQuery request in database
*/
func Insert_request_query(seqNum int64, truckid int32) (error){
	if err = db.Create(&RequestUQuery{
		SeqNumber: seqNum,
		Truckid: truckid,
	}).Error; err != nil {
		return err
	}
	return nil
}


func Query_seq_number (worldid int64) (int64, error) {
	var world WorldSeq
	if err = db.Where("world_id=?", worldid).Find(&world).Error; err != nil {
		return 0, err
	}
	return world.SeqNum, nil
}

/*
USAGE
when we want to retrieve truck info such as last time location (x, y), truck status,
toward which warehouse
INPUT
--truck id
OUTPUT
--Truck struct
--error
*/
//check query, should we add [] to Truck even though only one tuple will be returned
func Query_table_truck_byID(truckID int32) (Truck, error){
	var trucks []Truck
	if err = db.Where("truck_id=?", truckID).Find(&trucks).Error; err != nil {
		return Truck{}, err
	}
	return trucks[0], nil
}


/*
USAGE
when we want to retrieve truck info such as last time location (x, y), truck status,
toward which warehouse
INPUT
--truck id
OUTPUT
--Truck struct
--error
*/
func Query_table_packages_byTruckId(truckID int32) ([]Package, error){
	var truck Truck
	var pkgs []Package
	if err = db.Model(&truck).Where("truck_id=?", truckID).Related(&truck.Packages).Find(&pkgs).Error; err != nil {
		return pkgs, err
	}
	return pkgs, nil
}


//check query, should we add [] to package even though only one tuple will be returned
func Query_table_package_byID(packageID int64) (Package, error){
	var packages []Package
	if err = db.Where("package_id=?", packageID).Find(&packages).Error; err != nil {
		return Package{},err
	}
	if len(packages) == 0 {
		return Package{}, fmt.Errorf("package not found")
	}
	return packages[0], nil
}

//check query, should we add [] to package even though only one tuple will be returned
func Query_table_package_byUsername(username string) ([]Package, error){
	var packages []Package
	if err = db.Where("user_username=?", username).Find(&packages).Error; err != nil {
		return packages,err
	}
	return packages, nil
}


func Query_request_gopickup(seqNum int64) (RequestUGoPickUp, error){
	var request RequestUGoPickUp
	if err = db.Where("seq_number=?", seqNum).Find(&request).Error; err != nil {
		return request, err
	}
	return request, nil
}

func Query_info_deliverylocation_bySeqNum(truckid int32) ([]UDeliveryLocation, error){
	var request RequestUGoDeliver
	var pkgs []UDeliveryLocation
	if err = db.First(&request).Where("truck_id=?", truckid).Error; err != nil {
		return pkgs, err
	}
	if err = db.Model(&request).Related(&request.Pkgs,"request_u_go_deliver_truck_id").Find(&pkgs).Error; err != nil {
		return pkgs, err
	}
	return pkgs, nil
}

func Query_request_godeliver(seqNum int64) (RequestUGoDeliver, error){
	var request RequestUGoDeliver
	if err = db.Where("seq_number=?", seqNum).Find(&request).Error; err != nil {
		return request, err
	}
	request.Pkgs, err = Query_info_deliverylocation_bySeqNum(request.TruckId)
	return request, nil
}


func Query_request_goquery(seqNum int64) (RequestUQuery, error){
	var request RequestUQuery
	if err = db.Where("seq_number=?", seqNum).Find(&request).Error; err != nil {
		return request, err
	}
	return request, nil
}

func Query_table_user(username string) (string, string, string, error){
	var user []User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return "","","", err
	}
	if len(user) == 0 {
		return "","","", fmt.Errorf("user not exist")
	}
	return user[0].Username, user[0].Password, user[0].Email, nil
}


func Update_seq_number(worldid int64, seqNum int64) (error) {
	var world []WorldSeq
	if err = db.Model(&world).Where("world_id=?", worldid).Updates(
		WorldSeq{SeqNum:seqNum}).Error; err != nil {
		return err
	}
	return nil
}

/*
USAGE
when we want to update the latest location of the truck
INPUT
--truck id
--location (x,y)
OUTPUT
--error
*/
//check query, should we add [] to Truck even though only one tuple will be returned
func Update_table_truck_location(truckID int32, x int32, y int32) (error){
	var trucks []Truck
	if err = db.Model(&trucks).Where("truck_id=?", truckID).Updates(
		Truck{X_lasttime: x, Y_lasttime: y}).Error; err != nil {
		return err
	}
	return nil
}



/*
USAGE
when we want to update the latest location of the truck
INPUT
--truck id
--location (x,y)
OUTPUT
--error
*/
//check query, should we add [] to Truck even though only one tuple will be returned
func Update_table_truck_status(truckID int32, status string) (error){
	var truck []Truck
	if err = db.Model(&truck).Where("truck_id=?", truckID).Updates(Truck{Status: status}).Error; err != nil {
		return err
	}
	return nil
}

func Update_table_truck_location_and_status(truckID int32, x int32, y int32, status string) (error) {
	var truck []Truck
	truckLock.Lock()
	defer truckLock.Unlock()
	if err = db.Model(&truck).Where("truck_id=?", truckID).Updates(Truck{Status: status, X_lasttime:x, Y_lasttime:y}).Error; err != nil {
		return err
	}
	return nil
}


/*
USAGE
when update package status
INPUT
--package id
--package new status
OUTPUT
--error
*/
//check query, should we add [] to Truck even though only one tuple will be returned
func Update_table_package_status(packageID int64, status string) (error){
	var packages []Package
	if err = db.Model(&packages).Where("package_id=?", packageID).Updates(Package{Status: status}).Error; err != nil {
		return err
	}
	return nil
}

func Update_table_package_username(packageID int64, username string) (error){
	var packages []Package
	if err = db.Model(&packages).Where("package_id=?", packageID).Updates(Package{UserUsername:username}).Error; err != nil {
		return err
	}
	return nil
}

func Update_table_package_dest(packageID int64, x int32, y int32) (error){
	var packages []Package
	if err = db.Model(&packages).Where("package_id=?", packageID).Updates(Package{X:x, Y:y}).Error; err != nil {
		return err
	}
	return nil
}

func Update_table_package_location(packageID int64, x int32, y int32) (error){
	var packages []Package
	if err = db.Model(&packages).Where("package_id=?", packageID).Updates(Package{X:x, Y:y}).Error; err != nil {
		return err
	}
	return nil
}

func Delete_request_ugopickup(seqNum int64) (error) {
	var request RequestUGoPickUp
	if err = db.Where("seq_number=?", seqNum).Delete(&request).Error; err != nil {
		return err
	}
	return nil
}

func Delete_request_ugodeliver(seqNum int64) (error) {
	var request RequestUGoDeliver
	if err = db.Where("seq_number=?", seqNum).Delete(&request).Error; err != nil {
		return err
	}
	return nil
}

func Delete_request_uquery(seqNum int64) (error) {
	var request RequestUQuery
	if err = db.Where("seq_number=?", seqNum).Delete(&request).Error; err != nil {
		return err
	}
	return nil
}

func Select_truck(whid int32) (int32, error) {
	var truck Truck
	if err = db.Where("status=? and  towh_id=?", "ARRIVE WAREHOUSE", whid).First(&truck).Error; err != nil {
		//fmt.Println("no arrived warehouse truck")
	} else {
		return truck.TruckId, nil
	}
	if err = db.Where("status=? and  towh_id=?", "TRAVELING", whid).First(&truck).Error; err != nil {
		//fmt.Println("no traveling truck")
	} else {
		return truck.TruckId, nil
	}
	if err = db.Where("status=? and  towh_id=?", "LOADING", whid).First(&truck).Error; err != nil {
		//fmt.Println("no traveling truck")
	} else {
		return truck.TruckId, nil
	}
	if err = db.Where("status=?", "IDLE").First(&truck).Error; err != nil {
		//fmt.Println("no an idle truck")
	} else {
		return truck.TruckId, nil
	}
	if err = db.First(&truck).Error; err != nil {
		//fmt.Println("no an idle truck")
	} else {
		return truck.TruckId, nil
	}


	return 0, fmt.Errorf("no trucks available")
}
/*
 return current seq #
 and increment it by 1 in database
*/
func Increment_seq_number(worldid int64) (int64, error) {
	var latestSeqNumber int64
	seqNumberLock.Lock()
	defer seqNumberLock.Unlock()

	if latestSeqNumber, err = Query_seq_number(worldid); err != nil {
		return 0, fmt.Errorf("fail to query sequence number when doing increment")
	} else {
		if err = Update_seq_number(worldid, latestSeqNumber + 1);  err != nil {
			return 0, fmt.Errorf("fail to update sequence number when doing increment")
		} else {
			return latestSeqNumber, nil
		}

	}

}

func IsExist_username(username string) (bool, error){
	var user []User
	if err = db.Where("username=?", username).Find(&user).Error; err != nil {
		return false, fmt.Errorf("fail to check if a username exists in database")
	}
	if (len(user) == 0) {
		return false, fmt.Errorf("no user found")
	}

	return true, nil
}

func IsValid_user(username string, passwd string) (bool, error) {
	var user []User
	if err = db.Where("username=? and password=?", username, passwd).Find(&user).Error; err != nil {
		return false, fmt.Errorf("fail to valid a user")
	}
	if (len(user) != 0) {
		return true, nil
	} else {
		return false, nil
	}
}

//unit tests
func main() {
	if err = Initialize_database_and_create_tables(); err != nil {
		panic(err)
	}
	fmt.Println("successfully connected to postgres")

	defer db.Close()

	//----------------------table testing----------------------------------

	////insert truck1 to truck table
	//if err = Insert_table_truck(1, 0 ,0); err != nil {
	//	panic(err)
	//}
	//fmt.Println("successfully insert truck")
	//
	////insert package1 to truck1
	//if err = Insert_table_package_and_update_truck_towhid(1,"A", 1, 1, 1, 1, "zw"); err != nil {
	//	panic(err)
	//}
	//fmt.Println("successfully insert package")
	//
	////query truck 1
	//var truck Truck
	//if truck, err = Query_table_truck_byID(1); err != nil {
	//	panic(err)
	//}
	//fmt.Println(truck)
	////query package1
	//var pkg Package
	//if pkg, err = Query_table_package_byID(1); err != nil {
	//	panic(err)
	//}
	//fmt.Println(pkg)
	//
	//
	////insert package2 to truck1
	//if err = Insert_table_package_and_update_truck_towhid(2,"B", 1, 1, 2, 1, "zq"); err != nil {
	//	panic(err)
	//}
	////query package2
	//if pkg, err = Query_table_package_byID(2); err != nil {
	//	panic(err)
	//}
	//fmt.Println("successfully insert package")
	//
	////query all packages in truck1
	//var pkgs []Package
	//if pkgs, err = Query_table_packages_byTruckId(1); err != nil {
	//	panic(err)
	//}
	//fmt.Println(pkgs)

	//--------------------------request testing-----------------------------------
	//if err = Insert_request_gopickup(1, 1 ,1); err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("successfully insert ugopickup")
	//
	//var pkgs []UDeliveryLocation
	//pkgs = append(pkgs, UDeliveryLocation{PackId:1, X:1, Y:1}, UDeliveryLocation{PackId:2, X:0,Y:0})
	//fmt.Println(pkgs)
	//if err = Insert_request_godeliver(2, 1 , pkgs); err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("successfully insert ugodeliver")
	//
	//if err = Insert_request_query(3, 1); err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("successfully insert ugodeliver")
	//
	//
	//if request, err := Query_request_gopickup(1); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(request)
	//}
	//if request, err := Query_request_godeliver(2); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(request)
	//}
	//if request, err := Query_request_goquery(3); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(request)
	//}
	//
	//if err := Delete_request_ugopickup(1); err != nil {
	//	fmt.Println(err)
	//} else {
	//
	//}
	//if request, err := Query_request_gopickup(1); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(request)
	//}
	//
	//if err := Delete_request_ugodeliver(2); err != nil {
	//	fmt.Println(err)
	//} else {
	//
	//}
	//if request, err := Query_request_godeliver(2); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(request)
	//}
	//
	//if err := Delete_request_uquery(3); err != nil {
	//	fmt.Println(err)
	//} else {
	//
	//}
	//if request, err := Query_request_goquery(3); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(request)
	//}
	//-------------test seq number----------------------------
	//if err = Insert_seq_number(1); err !=nil {
	//}
	//fmt.Println("successfully start a the sequence number of a world")
	//
	//var seq int64
	//if seq, err = Increment_seq_number(1); err !=nil {
	//}
	//fmt.Printf("current sequence number: %d\n", seq)
	//
	////var truckid int32
	////if  truckid, err = Select_truck(1); err != nil {
	////
	////} else {
	////	fmt.Printf("select truck id: %d\n", truckid)
	////}

	//if err = Insert_table_user("zw", "1234567", "zuaho.wu@outlook.com"); err != nil {
	//	fmt.Println(err)
	//}
	//var isExist bool
	//if isExist, err = IsExist_username("zw"); err != nil {
	//	fmt.Println(err)
	//}
	//if (isExist) {
	//	fmt.Println("<exist> zw")
	//}
	//
	//if isExist, err = IsExist_username("zq"); err != nil {
	//	fmt.Println(err)
	//}
	//if (isExist) {
	//	fmt.Println("<exist> zq")
	//} else {
	//	fmt.Println("<exist> not exist zq")
	//}
	//
	////-------------------------------------------------------------
	if err = Insert_table_user("zw", "1234567", "zuaho.wu@outlook.com"); err != nil {
		fmt.Println(err)
	}
	//var isValid bool
	//if isValid, err = IsValid_user("zw", "1234567"); err != nil {
	//	fmt.Println(err)
	//}
	//if (isValid) {
	//	fmt.Println("<exist> zw")
	//}
	//
	//if isValid, err = IsValid_user("zw", "7654321"); err != nil {
	//	fmt.Println(err)
	//}
	//if (isValid) {
	//	fmt.Println("<exist> zw")
	//} else {
	//	fmt.Println("<exist> not exist zq")
	//}
	//
	//username, passwd, email, err1 := Query_table_user("zw")
	//if err1 != nil {
	//	fmt.Println("error when query user")
	//} else {
	//	fmt.Println(username + " " + passwd + " " + email)
	//}

	Insert_table_package_and_update_truck_towhid(1,"A",1,1,1,1,"zw")
	Insert_table_package_and_update_truck_towhid(2,"B",1,1,1,1,"zw")

	if pkgs, err :=Query_table_package_byUsername("zw"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pkgs)
	}

}