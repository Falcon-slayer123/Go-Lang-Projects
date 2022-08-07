package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct { //customer structure
	ticket_id int
	name      string
	number    int
	complaint string
	status    string
	comment   string
}

//var customer_map = map[int]Customer{} //global slice declration
var customer_list = []Customer{}

//ping database

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}

var db, e = sql.Open("mysql", "root:root@/golangapp")

//function- to add customer details

//1--Adding Customer//
func addcust() {
	var new_cust Customer
	new_cust.ticket_id = len(customer_list) + 1
	fmt.Println("Enter the Name --->")
	fmt.Scan(&new_cust.name)
	fmt.Println("Enter the Number --->")
	fmt.Scan(&new_cust.number)
	fmt.Println("Enter the Complaint --->")
	fmt.Scan(&new_cust.complaint)
	new_cust.status = "open" //assising the stataus as open in begining6

	new_cust.comment = "Null"
	new_cust.ticket_id = 0
	new_cust.ticket_id = len(customer_list) + 1
	customer_list = append(customer_list, new_cust)
	fmt.Println("Compliant Added Successfully")

	// INSERT user_data INTO DB
	// prepare

	stmt, e := db.Prepare("INSERT INTO Ticket(name,number,compliant,status,comment) values(?,?,?,?,?)")
	ErrorCheck(e)

	//execute
	res, er := stmt.Exec(new_cust.name, new_cust.number, new_cust.complaint, new_cust.status, new_cust.comment)
	ErrorCheck(er)

	id, e := res.LastInsertId()
	ErrorCheck(e)
	fmt.Println("Customer Addeded with id--->", id)
	fmt.Println("Customer Added Successfully")
}

//function to check error//
func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func view_tickets(cust Customer) {
	fmt.Println("Ticket ID --->", cust.ticket_id)
	fmt.Println("Customer Name --->", cust.name)
	fmt.Println("Customer Number --->", cust.number)
	fmt.Println("Customer Compliant --->", cust.complaint)
	fmt.Println("Current Status --->", cust.status)
	fmt.Println("Comments --->", cust.comment)
}

func list_compliant() {

	var cust = Customer{}
	// query all data
	rows, e := db.Query("select * from Ticket")
	ErrorCheck(e)

	for rows.Next() {
		e = rows.Scan(&cust.ticket_id, &cust.name, &cust.number, &cust.complaint, &cust.status, &cust.comment)
		ErrorCheck(e)
		view_tickets(cust)
	}
}

//function for ticket update

func ticket_update() {
	var ticket_id int
	var cust Customer
	fmt.Println("Enter the ticket id for the Information")
	fmt.Println(&ticket_id)
	for index, val := range customer_list {
		if ticket_id == val.ticket_id {
			var option int
			fmt.Println("1--> Resolved Ticket")
			fmt.Println("2--> Rejected Ticket") //The user will enter the choice on his own
			fmt.Println("Enter the Choice")
			fmt.Scan(&option)
			switch option {
			case 1:
				fmt.Println("Ticket is Resolved and updated")
				fmt.Println("Enter the Solution")
				fmt.Println(&cust.comment)
				//adding the status to the list
				customer_list[index].status = "Resolved"

			case 2:
				fmt.Println("Ticket is Rejected")
				fmt.Println("Enter the closing Comments")
				fmt.Println(&cust.comment)
				//adding the status to the list
				customer_list[index].status = "Rejected"
			}
			break
		}

	}
}

//Search ticket by ID
func search_by_id() {
	var id_new int
	fmt.Println("Enter the Ticket ID for Search")
	fmt.Scan(&id_new)
	qyr := fmt.Sprintf("select * from Product where id=%d", id_new)
	rows, e := db.Query(qyr)
	ErrorCheck(e)
	var cust Customer
	for rows.Next() {
		e = rows.Scan(&cust.ticket_id, &cust.name, &cust.number, &cust.complaint, &cust.status, &cust.comment)
		ErrorCheck(e)
	}

	// if cust.ticket_id != id_new {
	// 	fmt.Println("Product with Id given is not found")
	// 	return
	// }

	// for _, val := range customer_list {
	// 	if val.ticket_id == id_new {
	// 		view_tickets(val)

	// 	}
	// 	//view_tickets(val)

	// }

}

//Search ticket by name
func search_by_name() {
	var name string
	fmt.Println("Enter the name for Search")
	fmt.Scan(&name)
	for _, val := range customer_list {
		if val.name == name {
			view_tickets(val)

		}
		//view_tickets(nam)
	}
}

//loop for status
func statusloop(status_ticket string) {
	for _, val := range customer_list {
		view_tickets(val)
	}
}

//Search by Status
func search_by_status() {
	var option int
	fmt.Println(" 1-->Ticket having Open Status")
	fmt.Println(" 2-->Ticket having Closed Status")
	fmt.Println(" 3-->Ticket having Rejected Status")
	fmt.Println("Enter the Choice")
	fmt.Scan(&option)

	switch option {
	case 1:
		statusloop("OPEN")

	case 2:
		statusloop("RESOLVED")

	case 3:
		statusloop("REJECTED")
	}
}

//Function to write in the File

func main() {
	fmt.Println("<------------------WELCOME TO SHARMA CRM SOFTWARE------------------")
	exit := false
	var choice int
	PingDB(db)

	for {
		fmt.Println("----Enter the Serial Action to be Done----")
		fmt.Println("1.--->Register Compliant")
		fmt.Println("2.--->List Tickets")
		fmt.Println("3.--->Tickets Update")
		fmt.Println("4.--->Tickets Search by Name")
		fmt.Println("5.--->Tickets Search by Status")
		fmt.Println("6.--->Tickets Search by ID")
		fmt.Println("7.--->Exit")
		fmt.Println("Enter your Choice")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addcust()
		case 2:
			list_compliant()
		case 3:
			ticket_update()
		case 4:
			search_by_name()
		case 5:
			search_by_status()
		case 6:
			search_by_id()
		case 7:
			exit = true
			fmt.Println("Thank you ! Application is Closed")
		}

		if exit {
			break
		}
	}

	//PingDB(db)

}
