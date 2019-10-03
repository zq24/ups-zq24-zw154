package front_server

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/satori/go.uuid"

	"log"
	"net"
	"strconv"

	"html/template"
	"net/http"
	uw "protobuf/pb_uw"
	psql "database_utils"
)

type user struct {
	UserName string
	Password string
	Email string
}

type PackageInfo struct {
	PackageID int64
	Status string
	CreatedAt string
	UpdatedAt string
	X int32
	Y int32
	WhID int32
	TruckID int32
	Username string
	Delivered bool
}

type IndexPage struct {
	Username string
	Pkgs[] PackageInfo
	HasLoggedIn bool
}

func (m *PackageInfo) get_url_packageid() (string) {
	var url string
	url = "/login/?key=" + strconv.FormatInt(m.PackageID, 10)
	return url
}

var Tpl *template.Template
var dbSessions = map[string]string{} // session ID, user ID

const (
	sessionLength = 300
)

var Wconn *net.Conn
var WorldId int64
//func init() {
//	tpl = template.Must(template.ParseGlob("templates/*"))
//}


func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	pkgs, err := psql.Query_table_package_byUsername(u.UserName)
	if err != nil {
		fmt.Println(err)
	}
	var pkgInfos []PackageInfo
	for _, pkg := range pkgs {
		var pkgInfo PackageInfo
		pkgInfo.PackageID = pkg.PackageID
		pkgInfo.Status = pkg.Status
		pkgInfo.X = pkg.X
		pkgInfo.Y = pkg.Y
		pkgInfo.TruckID = pkg.TruckID
		pkgInfo.WhID = pkg.WhID
		pkgInfo.CreatedAt = pkg.CreatedAt.String()
		pkgInfo.UpdatedAt = pkg.UpdatedAt.String()
		//if pkgInfo.Status == "DELIVERED" {
		//	pkgInfo.Delivered = true
		//} else {
		//	pkgInfo.Delivered = false
		//}

		if pkgInfo.Status == "DELIVERRING" || pkgInfo.Status == "LOADED" {
			pkgInfo.Delivered = false
		} else {
			pkgInfo.Delivered = true
		}


		pkgInfos = append(pkgInfos, pkgInfo)
	}
	pkgsInfo := IndexPage{u.UserName, pkgInfos, alreadyLoggedIn(w, req)}
	Tpl.ExecuteTemplate(w, "index.gohtml", pkgsInfo)
}


func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		username := req.FormValue("username")
		passwd := req.FormValue("password")
		email := req.FormValue("email")

		var isExistUserName bool
		var err error
		if isExistUserName, err =  psql.IsExist_username(username); err != nil {
			fmt.Printf("Check existence of an user on error: %s", err)
		}
		if isExistUserName {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		//
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = username
		u = user{username, passwd, email}
		psql.Insert_table_user(username, passwd, email)
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	Tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u user
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")


		var isValid bool
		var err error
		if isValid, err = psql.IsValid_user(un, p); err != nil {
			return
		}
		if !isValid {
			http.Error(w, "invalid user or password", http.StatusForbidden)
			return

		}
		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

		c.MaxAge = sessionLength
		//
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}


	Tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(dbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}


func guest(w http.ResponseWriter, req *http.Request) {
	var pkg psql.Package
	var err error
	if req.Method == http.MethodPost {
		trackID := req.FormValue("tracking")
		trackID_int64, _ := strconv.ParseInt(trackID, 10, 64)
		pkg, err = psql.Query_table_package_byID(trackID_int64)
		if err != nil {
			http.Error(w, "Not a valid tracking number", http.StatusForbidden)
			return
		}
	}
	Tpl.ExecuteTemplate(w, "guest.gohtml", pkg)
}

/*
DESCRIPTION
start point of the front server
USAGE
invoke this function in the main thread
*/
func Run() {
	//psql.Initialize_database_and_create_tables()
	//tpl = template.Must(template.ParseGlob("templates/*"))
	http.HandleFunc("/", index)
	//http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/guest", guest)
	http.HandleFunc("/modifydestination", modifydestination)
	//http.HandleFunc("/currentLocation", currentLocation)
	http.ListenAndServe(":8000", nil)
}



func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u.UserName, u.Password, u.Email, err = psql.Query_table_user(un)
		if err != nil {
			fmt.Println("retrieve user info on error")
		}
	}
	return u
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	un, ok := dbSessions[c.Value]
	if !ok {
		return false
	}

	var isExist bool
	isExist, err = psql.IsExist_username(un)
	if err != nil {
		fmt.Println("check user exist on error")
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return isExist
}


func modifydestination(w http.ResponseWriter, req *http.Request) {
	packageids, ok := req.URL.Query()["packageid"]
	if !ok || len(packageids) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	if req.Method == http.MethodPost {
		x := req.FormValue("xcoordinate")
		y := req.FormValue("ycoordinate")
		x_int64, err := strconv.ParseInt(x, 10, 32)
		if err != nil {
			return
		}
		y_int64, err := strconv.ParseInt(y, 10, 32)
		if err != nil {
			return
		}
		p_int64, err := strconv.ParseInt(packageids[0], 10, 64)
		if err != nil {
			return
		}



		pkg, _ := psql.Query_table_package_byID(p_int64)
		var new_pkg uw.UDeliveryLocation
		var new_pkg_db psql.UDeliveryLocation
		new_pkg.Packageid = proto.Int64(p_int64)
		new_pkg.X = proto.Int32(int32(x_int64))
		new_pkg.Y = proto.Int32(int32(y_int64))
		new_pkg_db.PackId = p_int64
		new_pkg_db.X = int32(x_int64)
		new_pkg_db.Y = int32(y_int64)

		var new_pkgs []*uw.UDeliveryLocation
		var new_pkgs_db []psql.UDeliveryLocation
		new_pkgs = append(new_pkgs, &new_pkg)
		new_pkgs_db = append(new_pkgs_db, new_pkg_db)

		var ugodelivery uw.UGoDeliver
		ugodelivery.Truckid = proto.Int32(pkg.TruckID)
		ugodelivery.Packages = new_pkgs

		var uwcommands uw.UCommands

		latestSeqNum, err := psql.Increment_seq_number(WorldId)
		if err != nil {
			log.Fatal(err)
		} else {
			ugodelivery.Seqnum = proto.Int64(latestSeqNum)
			uwcommands.Deliveries = append(uwcommands.Deliveries, &ugodelivery)
			//3.timer for retransmission
			ugodelivery.WaitAck(Wconn)
			//4.insert request to database
			psql.Insert_request_godeliver(latestSeqNum, *ugodelivery.Truckid, new_pkgs_db)
		}

		//psql.Update_table_package_location(p_int64, int32(x_int64), int32(y_int64))

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	Tpl.ExecuteTemplate(w, "changelocation.html",nil)
}