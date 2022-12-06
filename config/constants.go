package config

//const SECRET_JWT = "S3CR3T"
var menteeStatus = [...]string{
	"Interview",
	"Join Class",
	"Unit 1",
	"Unit 2",
	"Unit 3",
	"Repeat Unit 1",
	"Repeat Unit 2",
	"Repeat Unit 3",
	"Placement",
	"Eleminated",
	"Graduated",
}

var menteeEmergencyStatus = [...]string{
	"Orang Tua",
	"Kakek Nenek",
	"Saudara Kandung",
	"Saudara dari Orang Tua",
}

var menteeEducationType = [...]string{
	"Informatic",
	"Non-Informatic",
}

var menteeGener = [...]string{
	"Male",
	"Female",
}

var userTeam = [...]string{
	"Academic",
	"People",
	"Placement",
}

var userRole = [...]string{
	"Super Admin",
	"User",
}

var userStatus = [...]string{
	"Active",
	"Non-Active",
	"Deleted",
}
